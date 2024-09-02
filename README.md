# `mediastreamcmp`

`mediastreamcmp` is a command-line tool that ensures two media files have identical video and audio streams. This is particularly useful for verifying the success of a remuxing operation (e.g., changing from `.mkv` to `.mp4`) or determining if the differences between two seemingly identical video files are only due to metadata.

# Features

- *Stream Comparison*: Compares video and audio streams between two media files using MD5 hashes to ensure they are identical.
- *Remux Verification*: Helps verify that a remuxing process (changing container formats) was successful without altering the underlying streams.
- *Metadata Ignorance*: Allows you to focus on the content of the streams, ignoring differences in metadata that do not affect the actual media.

# Usage

```bash
mediastreamcmp <file1> <file2>
```
- <file1>: The first media file to compare.
- <file2>: The second media file to compare.

# Example

To compare two video files, `video1.mkv` and `video2.mp4`, run:

```bash
mediastreamcmp video1.mkv video2.mp4
```
If the files have identical video and audio streams, the output will be:
```
The files are identical in terms of audio and video streams.
```
If the files differ in their streams, the output will indicate that they are not identical.

# Installation

## Prerequisites
- *Go*: Make sure you have Go installed on your system. You can download it from golang.org.
- *FFmpeg*: This tool relies on `ffmpeg` and `ffprobe` for stream extraction and hashing. You can install FFmpeg via package managers like `apt`, `brew`, or directly from ffmpeg.org.

## Build from Source
To install `mediastreamcmp`, clone the repository and build the binary using Go:

```bash
git clone <repository-url>
cd mediastreamcmp
go build -o mediastreamcmp
```
This will produce the mediastreamcmp binary in the current directory.

## Installation via Go
Alternatively, you can install the tool directly using go install:

```bash
go install <repository-url>@latest
```
# Use Cases

- `Remux Verification`: After converting a file from one container format to another (e.g., `.mkv` to `.mp4`), use `mediastreamcmp` to ensure the streams are unchanged.
- `Duplicate Detection`: When you have two video files that look almost identical, `mediastreamcmp` can help determine if they differ only in metadata or in their actual streams.
- `Quality Assurance`: Verify that no unintended changes were introduced during video processing or editing workflows.

# Contributing

Contributions are welcome! Please submit pull requests, report issues, or suggest improvements.

# Acknowledgments

`mediastreamcmp` relies on the powerful FFmpeg suite to handle media stream extraction and hashing.
