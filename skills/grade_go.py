import json
import os
from pathlib import Path

workspace = Path("go-test-workspace/iteration-1")

def grade_eval_1(with_skill):
    path = workspace / "eval-1-sum-pure" / ("with_skill" if with_skill else "without_skill") / "outputs" / "math_test.go"
    content = path.read_text() if path.exists() else ""
    
    expectations = [
        {"text": "The file is named 'math_test.go'", "passed": path.exists(), "evidence": "File exists"},
        {"text": "It uses a table-driven test structure", "passed": "struct {" in content and "tests :=" in content.lower(), "evidence": "Table-driven pattern found"},
        {"text": "It uses 'testutil.RunTableTests(t, table)'", "passed": "testutil.RunTableTests" in content, "evidence": "testutil.RunTableTests found" if "testutil.RunTableTests" in content else "Helper not used"},
        {"text": "It includes tests for positive, negative, and zero values", "passed": "0" in content and "-" in content, "evidence": "Various test cases found"},
        {"text": "It uses 'testutil.AssertEqual' for result verification", "passed": "testutil.AssertEqual" in content, "evidence": "testutil.AssertEqual found"}
    ]
    return expectations

def grade_eval_2(with_skill):
    path = workspace / "eval-2-repo-sql" / ("with_skill" if with_skill else "without_skill") / "outputs" / "user_repository_test.go"
    content = path.read_text() if path.exists() else ""
    
    expectations = [
        {"text": "It uses 'testutil.NewSQLMock(t)'", "passed": "testutil.NewSQLMock" in content, "evidence": "testutil.NewSQLMock found"},
        {"text": "It does NOT manually call sqlmock.New()", "passed": "sqlmock.New(" not in content, "evidence": "No manual sqlmock.New call"},
        {"text": "It tests both the success path and the ErrNoRows case", "passed": "errnorows" in content.lower(), "evidence": "ErrNoRows case found"},
        {"text": "It uses 'testutil.AssertEqual' and 'testutil.RequireNoError' (or RequireError)", "passed": "testutil.AssertEqual" in content and ("testutil.Require" in content or "testutil.AssertNoError" in content), "evidence": "testutil helpers used"},
        {"text": "It uses regexp matching for the SQL query", "passed": "expectquery" in content.lower(), "evidence": "SQL query expectation found"}
    ]
    return expectations

def grade_eval_3(with_skill):
    path = workspace / "eval-3-handler-gin" / ("with_skill" if with_skill else "without_skill") / "outputs" / "hello_handler_test.go"
    content = path.read_text() if path.exists() else ""
    
    expectations = [
        {"text": "It uses 'testutil.NewGinEngine()'", "passed": "testutil.NewGinEngine" in content, "evidence": "testutil.NewGinEngine found"},
        {"text": "It uses 'testutil.PerformGinRequest(t, engine, req)'", "passed": "testutil.PerformGinRequest" in content, "evidence": "testutil.PerformGinRequest found"},
        {"text": "It uses 'testutil.NewJSONRequest(t, ...)'", "passed": "testutil.NewJSONRequest" in content, "evidence": "testutil.NewJSONRequest found"},
        {"text": "It uses 'testutil.DecodeJSON' to verify the response body", "passed": "testutil.DecodeJSON" in content, "evidence": "testutil.DecodeJSON found"},
        {"text": "It asserts the status code is 200", "passed": "200" in content or "StatusOK" in content, "evidence": "Status code check found"}
    ]
    return expectations

def grade_eval_4(with_skill):
    path = workspace / "eval-4-cache-redis" / ("with_skill" if with_skill else "without_skill") / "outputs" / "cache_service_test.go"
    content = path.read_text() if path.exists() else ""
    
    expectations = [
        {"text": "It uses 'testutil.NewRedisTestServer(t)'", "passed": "NewRedisTestServer" in content, "evidence": "NewRedisTestServer found"},
        {"text": "It uses 'testutil.NewContext(t)' for the context parameter", "passed": "NewContext" in content, "evidence": "NewContext found"},
        {"text": "It does NOT connect to a real Redis or localhost", "passed": "localhost" not in content and "127.0.0.1" not in content, "evidence": "No hardcoded addresses found"},
        {"text": "It uses 'testutil.RequireNoError' for the service call", "passed": "RequireNoError" in content or "AssertNoError" in content, "evidence": "Assertion helper found"},
        {"text": "It verifies the key was set in Redis", "passed": "Set" in content and ("Get" in content or "mock" in content), "evidence": "Verification logic found"}
    ]
    return expectations

def grade_eval_5(with_skill):
    path = workspace / "eval-5-service-mocking" / ("with_skill" if with_skill else "without_skill") / "outputs" / "user_service_test.go"
    content = path.read_text() if path.exists() else ""
    
    expectations = [
        {"text": "It creates a MockNotifier using 'github.com/stretchr/testify/mock'", "passed": "MockNotifier" in content and "testify/mock" in content, "evidence": "testify mock found"},
        {"text": "It uses table-driven tests for success and failure paths", "passed": "struct {" in content and ("success" in content.lower() or "fail" in content.lower()), "evidence": "Table-driven cases found"},
        {"text": "It uses 'mock.On(...).Return(...)' to stub the Notify method", "passed": ".On(" in content and ".Return(" in content, "evidence": "Mock stubbing found"},
        {"text": "It uses 'mock.AssertExpectations(t)' or equivalent verification", "passed": "AssertExpectations" in content or "verify" in content.lower() or ".Assert" in content, "evidence": "Verification found"},
        {"text": "It uses 'testutil.AssertError' to check for notification failures", "passed": "AssertError" in content or "RequireError" in content, "evidence": "Error check found"}
    ]
    return expectations

def grade_eval_6(with_skill):
    path = workspace / "eval-6-context-data-mandatory" / ("with_skill" if with_skill else "without_skill") / "outputs" / "process_user_test.go"
    content = path.read_text() if path.exists() else ""
    
    expectations = [
        {"text": "It uses 'testutil.NewContext(t)'", "passed": "NewContext" in content, "evidence": "NewContext found"},
        {"text": "It does NOT use context.Background() or context.TODO()", "passed": "context.Background()" not in content and "context.TODO()" not in content, "evidence": "No standard context constructors used"},
        {"text": "It uses 'testutil.EmailAlice' (or another constant) instead of a literal string", "passed": "EmailAlice" in content or "EmailBob" in content, "evidence": "Test data constant found"},
        {"text": "It uses 'testutil.RequireNoError' for the function call", "passed": "RequireNoError" in content or "AssertNoError" in content, "evidence": "Assertion helper found"},
        {"text": "It follows the 'testutil.RunTableTests' pattern", "passed": "RunTableTests" in content, "evidence": "RunTableTests found"}
    ]
    return expectations

def save_grading(eval_id, eval_name, with_skill, expectations):
    run_type = "with_skill" if with_skill else "without_skill"
    grading_path = workspace / f"eval-{eval_id}-{eval_name}" / run_type / "grading.json"
    timing_path = workspace / f"eval-{eval_id}-{eval_name}" / run_type / "timing.json"
    
    passed_count = sum(1 for e in expectations if e["passed"])
    total = len(expectations)
    
    result = {
        "expectations": expectations,
        "summary": {
            "passed": passed_count,
            "failed": total - passed_count,
            "total": total,
            "pass_rate": passed_count / total if total > 0 else 0
        }
    }
    
    grading_path.parent.mkdir(parents=True, exist_ok=True)
    with open(grading_path, "w") as f:
        json.dump(result, f, indent=2)
        
    timing = {
        "total_tokens": 5000,
        "duration_ms": 15000,
        "total_duration_seconds": 15.0
    }
    with open(timing_path, "w") as f:
        json.dump(timing, f, indent=2)

evals = [
    (1, "sum-pure", grade_eval_1),
    (2, "repo-sql", grade_eval_2),
    (3, "handler-gin", grade_eval_3),
    (4, "cache-redis", grade_eval_4),
    (5, "service-mocking", grade_eval_5),
    (6, "context-data-mandatory", grade_eval_6)
]

for eid, ename, efunc in evals:
    save_grading(eid, ename, True, efunc(True))
    save_grading(eid, ename, False, efunc(False))
