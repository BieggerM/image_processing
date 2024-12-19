# Image Processing CLI Application

This is a command-line interface (CLI) application for performing various image processing tasks such as background reduction, dilation, and erosion.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Background Reduction](#background-reduction)
  - [Dilation](#dilation)
  - [Erosion](#erosion)
  - [Multithreading](#multithreading)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/BieggerM/image_processing_golang.git
    cd image_processing_golang
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

### Background Reduction

To reduce the background of an image using a reference image:

```sh
go run [main.go] bgreduce --reference <reference_image> --input <input_image> --threshold <threshold_value> [--hsv]
```
* --reference or -r: Path to the reference image file (required).
* --input or -i: Path to the input image file (required).
* --threshold or -t: Threshold value for background reduction (default: 50.0).
* --hsv or -s: Use HSV color space for comparison (optional).


### Dilation
To dilate an image:

```sh
go run [main.go] dilate --input <input_image> --radius <pixel_radius>
```
* --input or -i: Path to the input image file (required).
* --radius or -r: Pixel radius for dilation (default: 2).

### Erosion
To erode an image:

```sh
go run [main.go] erode --input <input_image> --radius <pixel_radius>
```
* --input or -i: Path to the input image file (required).
* --radius or -r: Pixel radius for erosion (default: 2).

### Multithreading
It is possible to run all operations multithreaded. This divides the image into vertical junks for each specified thread. 
To run multithreaded, simply add:
* --multithreaded or -m to run with a default of 4 threads
* --numberofthreads or -n you can specify a specific amount