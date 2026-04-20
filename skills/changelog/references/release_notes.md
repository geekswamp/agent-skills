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
- Minimum **100 characters**
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
- Maximum **500 characters** per language
- Clear and natural language
- Avoid overly technical phrasing
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
- **CRITICAL**: To generate the final output, you MUST ALWAYS use the `scripts/write-release-notes.py` script with the `--lang`, `--intro`, and `--items` parameters. Never create or edit the `RELEASE_NOTES` file manually.

## 6. Example Usage
Example command:

```bash
python3 scripts/write-release-notes.py \
--lang id \
--intro "Pembaruan ini hadir untuk membuat aplikasi lebih nyaman, stabil, dan mudah digunakan. Kami terus meningkatkan keandalan sistem agar pengalaman Anda lebih lancar." \
--items "
- Tambah fitur ekspor laporan untuk unduh data
- Perbaiki masalah saat login ke aplikasi
- Tingkatkan stabilitas performa aplikasi
" \
--lang en \
--intro "This update makes the app more comfortable, stable, and easy to use. We continuously improve system reliability to ensure a smoother experience for you." \
--items "
- Added simple report export feature
- Fixed issues when signing into the app
- Improved overall application stability
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