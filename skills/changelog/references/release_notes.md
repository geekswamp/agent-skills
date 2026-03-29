# Generate Release Notes
Follow this workflow to generate `RELEASE_NOTES`.

## 1. Read the Latest Release from CHANGELOG.md
Open `CHANGELOG.md` and locate the **most recent version** (the first release section).

Use only the changes listed under that version as the source of truth.

## 2. Select the Most Important Changes
If the version contains many changes:
- Select only the **most important or user-visible improvements**.
- Ignore minor internal updates, refactoring, dependency updates, or technical changes that users would not notice.
- Keep the release notes concise and easy to understand.
- Include **no more than 5 improvements** in the final release notes.

## 3. Rewrite Changes into Friendly Explanations
Transform the selected changes into short explanations that:
- Are **clear and friendly**
- Use **simple language**
- Are understandable by **non-technical users**, especially older users
- Avoid technical jargon whenever possible
- Start each bullet with a **capital letter**

Before listing the changes, add a **short opening sentence** that introduces the update.

Rules for the opening sentence:
- Minimum **150 characters**
- Maximum **2 sentences**
- Keep it **friendly and concise**
- Avoid technical terms
- Make it easy to understand for general users

Tone guidelines:
- Warm and welcoming
- Relaxed and conversational
- Easy to read
- No emojis

## 4. Language Requirements
Generate the release notes in **multiple languages** (at least Indonesian and English):

### Indonesian
- Maximum **2500 characters**
- Friendly and easy to understand
- Natural Indonesian wording

### English
- Maximum **2500 characters**
- Clear and natural English
- Avoid overly technical phrasing

Additional languages (optional):
- Follow the same tone and structure
- Keep wording natural for native speakers

## 5. Write the Result Using the Release Note Script
After preparing the introduction and the selected improvements, write the final result using the [write-release-notes.py](../scripts/write-release-notes.py) script:

```
python3 scripts/write-release-notes.py
```

Provide parameters using **language blocks**:
- `--lang` = Language code (e.g., `id`, `en`, `fr`)
- `--intro` = Opening sentence for that language
- `--items` = List of improvements for that language

Notes:
- Each language must be defined using its own `--lang` block.
- Each improvement can be written using `-` bullets or numbering.
- The script will automatically:
  - Normalize formatting
  - Validate changes against `CHANGELOG.md`
  - Rank changes by importance
- The script will overwrite or create the `RELEASE_NOTES` file.

## 6. Example Usage
Example command:

```bash
python3 scripts/write-release-notes.py \
--lang id \
--intro "Kami menghadirkan beberapa pembaruan penting untuk membuat aplikasi lebih nyaman, stabil, dan mudah digunakan dalam aktivitas sehari-hari. Pembaruan ini juga membantu meningkatkan keandalan aplikasi agar pengalaman penggunaan menjadi lebih lancar." \
--items "
- Menambahkan fitur ekspor laporan agar data dapat diunduh dengan lebih mudah
- Memperbaiki masalah login yang kadang membuat pengguna gagal masuk
- Meningkatkan stabilitas aplikasi agar berjalan lebih lancar
" \
--lang en \
--intro "This update brings several improvements designed to make the app easier, more stable, and more comfortable to use in everyday situations. It also enhances overall reliability so the experience feels smoother and more dependable." \
--items "
- Added a report export feature so users can download their data more easily
- Fixed a login issue that sometimes prevented users from signing in
- Improved overall app stability for a smoother experience
"
```

The script will generate or update the file:

```
RELEASE_NOTES
```

With the following structure:

```md
# Release Notes

## ID
<opening sentence>

- <important improvement>
- <important improvement>
- <important improvement>

## EN
<opening sentence>

- <important improvement>
- <important improvement>
- <important improvement>
```