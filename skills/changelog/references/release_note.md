# Generate Release Notes

Follow this workflow to generate `RELEASE_NOTE`.

## 1. Read the Latest Release from CHANGELOG.md
Open `CHANGELOG.md` and locate the **most recent version** (the first release section).

Use only the changes listed under that version as the source of truth.

## 2. Select the Most Important Changes
If the version contains many changes:
- Select only the **most important or user-visible improvements**.
- Ignore minor internal updates, refactoring, dependency updates, or technical changes that users would not notice.
- Keep the release notes concise and easy to understand.

## 3. Rewrite Changes into Friendly Explanations
Transform the selected changes into short explanations that:
- Are **clear and friendly**
- Use **simple language**
- Are understandable by **non-technical users**, especially older users
- Avoid technical jargon whenever possible
- Start each bullet with a **capital letter**

Before listing the changes, add a **short opening sentence** that introduces the update.

Tone guidelines:
- Warm and welcoming
- Relaxed and conversational
- Easy to read
- No emojis

## 4. Language Requirements
Generate the release notes in **two languages**:

### Indonesian
- Maximum **1000 characters**
- Friendly and easy to understand
- Natural Indonesian wording

### English
- Maximum **1000 characters**
- Clear and natural English
- Avoid overly technical phrasing

## 5. Output Format

Create a file named `RELEASE_NOTE` with this structure:

```md
# Release Notes

## Bahasa Indonesia
<opening sentence>

- <important improvement>
- <important improvement>
- <important improvement>

## English
<opening sentence>

- <important improvement>
- <important improvement>
- <important improvement>
```