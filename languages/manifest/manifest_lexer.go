// Code generated from Manifest.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 44, 612,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 394, 10, 2, 12, 2, 14, 2,
	397, 11, 2, 3, 2, 3, 2, 3, 2, 7, 2, 402, 10, 2, 12, 2, 14, 2, 405, 11,
	2, 7, 2, 407, 10, 2, 12, 2, 14, 2, 410, 11, 2, 3, 2, 3, 2, 7, 2, 414, 10,
	2, 12, 2, 14, 2, 417, 11, 2, 3, 2, 3, 2, 3, 2, 7, 2, 422, 10, 2, 12, 2,
	14, 2, 425, 11, 2, 3, 2, 3, 2, 3, 2, 7, 2, 430, 10, 2, 12, 2, 14, 2, 433,
	11, 2, 7, 2, 435, 10, 2, 12, 2, 14, 2, 438, 11, 2, 5, 2, 440, 10, 2, 3,
	3, 3, 3, 5, 3, 444, 10, 3, 3, 3, 7, 3, 447, 10, 3, 12, 3, 14, 3, 450, 11,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 6, 4, 456, 10, 4, 13, 4, 14, 4, 457, 6, 4, 460,
	10, 4, 13, 4, 14, 4, 461, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30,
	3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3,
	36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40,
	3, 41, 5, 41, 546, 10, 41, 3, 41, 3, 41, 5, 41, 550, 10, 41, 3, 41, 3,
	41, 3, 42, 3, 42, 3, 42, 7, 42, 557, 10, 42, 12, 42, 14, 42, 560, 11, 42,
	3, 42, 3, 42, 3, 43, 3, 43, 7, 43, 566, 10, 43, 12, 43, 14, 43, 569, 11,
	43, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 575, 10, 44, 3, 44, 5, 44, 578,
	10, 44, 3, 44, 3, 44, 3, 44, 6, 44, 583, 10, 44, 13, 44, 14, 44, 584, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 593, 10, 44, 3, 45, 3, 45,
	3, 46, 3, 46, 3, 47, 3, 47, 5, 47, 601, 10, 47, 3, 48, 3, 48, 7, 48, 605,
	10, 48, 12, 48, 14, 48, 608, 11, 48, 3, 48, 5, 48, 611, 10, 48, 2, 2, 49,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41,
	22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59,
	31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77,
	40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 2, 89, 2, 91, 2, 93, 2, 95, 2,
	3, 2, 14, 8, 2, 38, 39, 42, 43, 45, 45, 47, 49, 62, 62, 97, 97, 3, 2, 67,
	92, 4, 2, 11, 11, 34, 34, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 10, 2,
	36, 36, 41, 41, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118,
	3, 2, 50, 53, 3, 2, 50, 57, 4, 2, 45, 45, 49, 49, 5, 2, 50, 59, 67, 72,
	99, 104, 4, 2, 67, 92, 99, 124, 3, 2, 50, 59, 4, 2, 50, 59, 97, 97, 2,
	654, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2,
	33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2,
	2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2,
	2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2,
	2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3,
	2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71,
	3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2,
	79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2,
	3, 439, 3, 2, 2, 2, 5, 441, 3, 2, 2, 2, 7, 459, 3, 2, 2, 2, 9, 463, 3,
	2, 2, 2, 11, 465, 3, 2, 2, 2, 13, 467, 3, 2, 2, 2, 15, 469, 3, 2, 2, 2,
	17, 471, 3, 2, 2, 2, 19, 473, 3, 2, 2, 2, 21, 475, 3, 2, 2, 2, 23, 477,
	3, 2, 2, 2, 25, 479, 3, 2, 2, 2, 27, 481, 3, 2, 2, 2, 29, 483, 3, 2, 2,
	2, 31, 485, 3, 2, 2, 2, 33, 487, 3, 2, 2, 2, 35, 489, 3, 2, 2, 2, 37, 491,
	3, 2, 2, 2, 39, 493, 3, 2, 2, 2, 41, 495, 3, 2, 2, 2, 43, 498, 3, 2, 2,
	2, 45, 501, 3, 2, 2, 2, 47, 504, 3, 2, 2, 2, 49, 507, 3, 2, 2, 2, 51, 510,
	3, 2, 2, 2, 53, 513, 3, 2, 2, 2, 55, 516, 3, 2, 2, 2, 57, 519, 3, 2, 2,
	2, 59, 521, 3, 2, 2, 2, 61, 523, 3, 2, 2, 2, 63, 525, 3, 2, 2, 2, 65, 527,
	3, 2, 2, 2, 67, 529, 3, 2, 2, 2, 69, 531, 3, 2, 2, 2, 71, 533, 3, 2, 2,
	2, 73, 535, 3, 2, 2, 2, 75, 537, 3, 2, 2, 2, 77, 540, 3, 2, 2, 2, 79, 542,
	3, 2, 2, 2, 81, 545, 3, 2, 2, 2, 83, 553, 3, 2, 2, 2, 85, 563, 3, 2, 2,
	2, 87, 592, 3, 2, 2, 2, 89, 594, 3, 2, 2, 2, 91, 596, 3, 2, 2, 2, 93, 600,
	3, 2, 2, 2, 95, 602, 3, 2, 2, 2, 97, 98, 7, 79, 2, 2, 98, 99, 7, 99, 2,
	2, 99, 100, 7, 112, 2, 2, 100, 101, 7, 107, 2, 2, 101, 102, 7, 104, 2,
	2, 102, 103, 7, 103, 2, 2, 103, 104, 7, 117, 2, 2, 104, 105, 7, 118, 2,
	2, 105, 106, 7, 47, 2, 2, 106, 107, 7, 88, 2, 2, 107, 108, 7, 103, 2, 2,
	108, 109, 7, 116, 2, 2, 109, 110, 7, 117, 2, 2, 110, 111, 7, 107, 2, 2,
	111, 112, 7, 113, 2, 2, 112, 440, 7, 112, 2, 2, 113, 114, 7, 68, 2, 2,
	114, 115, 7, 119, 2, 2, 115, 116, 7, 112, 2, 2, 116, 117, 7, 102, 2, 2,
	117, 118, 7, 110, 2, 2, 118, 119, 7, 103, 2, 2, 119, 120, 7, 47, 2, 2,
	120, 121, 7, 67, 2, 2, 121, 122, 7, 101, 2, 2, 122, 123, 7, 118, 2, 2,
	123, 124, 7, 107, 2, 2, 124, 125, 7, 120, 2, 2, 125, 126, 7, 99, 2, 2,
	126, 127, 7, 118, 2, 2, 127, 128, 7, 113, 2, 2, 128, 440, 7, 116, 2, 2,
	129, 130, 7, 69, 2, 2, 130, 131, 7, 116, 2, 2, 131, 132, 7, 103, 2, 2,
	132, 133, 7, 99, 2, 2, 133, 134, 7, 118, 2, 2, 134, 135, 7, 103, 2, 2,
	135, 136, 7, 102, 2, 2, 136, 137, 7, 47, 2, 2, 137, 138, 7, 68, 2, 2, 138,
	440, 7, 123, 2, 2, 139, 140, 7, 75, 2, 2, 140, 141, 7, 111, 2, 2, 141,
	142, 7, 114, 2, 2, 142, 143, 7, 113, 2, 2, 143, 144, 7, 116, 2, 2, 144,
	145, 7, 118, 2, 2, 145, 146, 7, 47, 2, 2, 146, 147, 7, 82, 2, 2, 147, 148,
	7, 99, 2, 2, 148, 149, 7, 101, 2, 2, 149, 150, 7, 109, 2, 2, 150, 151,
	7, 99, 2, 2, 151, 152, 7, 105, 2, 2, 152, 440, 7, 103, 2, 2, 153, 154,
	7, 71, 2, 2, 154, 155, 7, 122, 2, 2, 155, 156, 7, 114, 2, 2, 156, 157,
	7, 113, 2, 2, 157, 158, 7, 116, 2, 2, 158, 159, 7, 118, 2, 2, 159, 160,
	7, 47, 2, 2, 160, 161, 7, 82, 2, 2, 161, 162, 7, 99, 2, 2, 162, 163, 7,
	101, 2, 2, 163, 164, 7, 109, 2, 2, 164, 165, 7, 99, 2, 2, 165, 166, 7,
	105, 2, 2, 166, 440, 7, 103, 2, 2, 167, 168, 7, 68, 2, 2, 168, 169, 7,
	119, 2, 2, 169, 170, 7, 112, 2, 2, 170, 171, 7, 102, 2, 2, 171, 172, 7,
	110, 2, 2, 172, 173, 7, 103, 2, 2, 173, 174, 7, 47, 2, 2, 174, 175, 7,
	88, 2, 2, 175, 176, 7, 103, 2, 2, 176, 177, 7, 116, 2, 2, 177, 178, 7,
	117, 2, 2, 178, 179, 7, 107, 2, 2, 179, 180, 7, 113, 2, 2, 180, 440, 7,
	112, 2, 2, 181, 182, 7, 68, 2, 2, 182, 183, 7, 119, 2, 2, 183, 184, 7,
	112, 2, 2, 184, 185, 7, 102, 2, 2, 185, 186, 7, 110, 2, 2, 186, 187, 7,
	103, 2, 2, 187, 188, 7, 47, 2, 2, 188, 189, 7, 80, 2, 2, 189, 190, 7, 99,
	2, 2, 190, 191, 7, 111, 2, 2, 191, 440, 7, 103, 2, 2, 192, 193, 7, 68,
	2, 2, 193, 194, 7, 119, 2, 2, 194, 195, 7, 112, 2, 2, 195, 196, 7, 102,
	2, 2, 196, 197, 7, 110, 2, 2, 197, 198, 7, 103, 2, 2, 198, 199, 7, 47,
	2, 2, 199, 200, 7, 70, 2, 2, 200, 201, 7, 103, 2, 2, 201, 202, 7, 117,
	2, 2, 202, 203, 7, 101, 2, 2, 203, 204, 7, 116, 2, 2, 204, 205, 7, 107,
	2, 2, 205, 206, 7, 114, 2, 2, 206, 207, 7, 118, 2, 2, 207, 208, 7, 107,
	2, 2, 208, 209, 7, 113, 2, 2, 209, 440, 7, 112, 2, 2, 210, 211, 7, 68,
	2, 2, 211, 212, 7, 119, 2, 2, 212, 213, 7, 112, 2, 2, 213, 214, 7, 102,
	2, 2, 214, 215, 7, 110, 2, 2, 215, 216, 7, 103, 2, 2, 216, 217, 7, 47,
	2, 2, 217, 218, 7, 70, 2, 2, 218, 219, 7, 113, 2, 2, 219, 220, 7, 101,
	2, 2, 220, 221, 7, 87, 2, 2, 221, 222, 7, 84, 2, 2, 222, 440, 7, 78, 2,
	2, 223, 224, 7, 68, 2, 2, 224, 225, 7, 119, 2, 2, 225, 226, 7, 112, 2,
	2, 226, 227, 7, 102, 2, 2, 227, 228, 7, 110, 2, 2, 228, 229, 7, 103, 2,
	2, 229, 230, 7, 47, 2, 2, 230, 231, 7, 79, 2, 2, 231, 232, 7, 99, 2, 2,
	232, 233, 7, 112, 2, 2, 233, 234, 7, 107, 2, 2, 234, 235, 7, 104, 2, 2,
	235, 236, 7, 103, 2, 2, 236, 237, 7, 117, 2, 2, 237, 238, 7, 118, 2, 2,
	238, 239, 7, 88, 2, 2, 239, 240, 7, 103, 2, 2, 240, 241, 7, 116, 2, 2,
	241, 242, 7, 117, 2, 2, 242, 243, 7, 107, 2, 2, 243, 244, 7, 113, 2, 2,
	244, 440, 7, 112, 2, 2, 245, 246, 7, 68, 2, 2, 246, 247, 7, 119, 2, 2,
	247, 248, 7, 112, 2, 2, 248, 249, 7, 102, 2, 2, 249, 250, 7, 110, 2, 2,
	250, 251, 7, 103, 2, 2, 251, 252, 7, 47, 2, 2, 252, 253, 7, 88, 2, 2, 253,
	254, 7, 103, 2, 2, 254, 255, 7, 112, 2, 2, 255, 256, 7, 102, 2, 2, 256,
	257, 7, 113, 2, 2, 257, 440, 7, 116, 2, 2, 258, 259, 7, 68, 2, 2, 259,
	260, 7, 119, 2, 2, 260, 261, 7, 112, 2, 2, 261, 262, 7, 102, 2, 2, 262,
	263, 7, 110, 2, 2, 263, 264, 7, 103, 2, 2, 264, 265, 7, 47, 2, 2, 265,
	266, 7, 85, 2, 2, 266, 267, 7, 123, 2, 2, 267, 268, 7, 111, 2, 2, 268,
	269, 7, 100, 2, 2, 269, 270, 7, 113, 2, 2, 270, 271, 7, 110, 2, 2, 271,
	272, 7, 107, 2, 2, 272, 273, 7, 101, 2, 2, 273, 274, 7, 80, 2, 2, 274,
	275, 7, 99, 2, 2, 275, 276, 7, 111, 2, 2, 276, 440, 7, 103, 2, 2, 277,
	278, 7, 75, 2, 2, 278, 279, 7, 111, 2, 2, 279, 280, 7, 114, 2, 2, 280,
	281, 7, 110, 2, 2, 281, 282, 7, 103, 2, 2, 282, 283, 7, 111, 2, 2, 283,
	284, 7, 103, 2, 2, 284, 285, 7, 112, 2, 2, 285, 286, 7, 118, 2, 2, 286,
	287, 7, 99, 2, 2, 287, 288, 7, 118, 2, 2, 288, 289, 7, 107, 2, 2, 289,
	290, 7, 113, 2, 2, 290, 291, 7, 112, 2, 2, 291, 292, 7, 47, 2, 2, 292,
	293, 7, 88, 2, 2, 293, 294, 7, 103, 2, 2, 294, 295, 7, 116, 2, 2, 295,
	296, 7, 117, 2, 2, 296, 297, 7, 107, 2, 2, 297, 298, 7, 113, 2, 2, 298,
	440, 7, 112, 2, 2, 299, 300, 7, 75, 2, 2, 300, 301, 7, 111, 2, 2, 301,
	302, 7, 114, 2, 2, 302, 303, 7, 110, 2, 2, 303, 304, 7, 103, 2, 2, 304,
	305, 7, 111, 2, 2, 305, 306, 7, 103, 2, 2, 306, 307, 7, 112, 2, 2, 307,
	308, 7, 118, 2, 2, 308, 309, 7, 99, 2, 2, 309, 310, 7, 118, 2, 2, 310,
	311, 7, 107, 2, 2, 311, 312, 7, 113, 2, 2, 312, 313, 7, 112, 2, 2, 313,
	314, 7, 47, 2, 2, 314, 315, 7, 86, 2, 2, 315, 316, 7, 107, 2, 2, 316, 317,
	7, 118, 2, 2, 317, 318, 7, 110, 2, 2, 318, 440, 7, 103, 2, 2, 319, 320,
	7, 67, 2, 2, 320, 321, 7, 112, 2, 2, 321, 322, 7, 118, 2, 2, 322, 323,
	7, 47, 2, 2, 323, 324, 7, 88, 2, 2, 324, 325, 7, 103, 2, 2, 325, 326, 7,
	116, 2, 2, 326, 327, 7, 117, 2, 2, 327, 328, 7, 107, 2, 2, 328, 329, 7,
	113, 2, 2, 329, 440, 7, 112, 2, 2, 330, 331, 7, 85, 2, 2, 331, 332, 7,
	114, 2, 2, 332, 333, 7, 116, 2, 2, 333, 334, 7, 107, 2, 2, 334, 335, 7,
	112, 2, 2, 335, 336, 7, 105, 2, 2, 336, 337, 7, 47, 2, 2, 337, 338, 7,
	88, 2, 2, 338, 339, 7, 103, 2, 2, 339, 340, 7, 116, 2, 2, 340, 341, 7,
	117, 2, 2, 341, 342, 7, 107, 2, 2, 342, 343, 7, 113, 2, 2, 343, 440, 7,
	112, 2, 2, 344, 345, 7, 80, 2, 2, 345, 346, 7, 99, 2, 2, 346, 347, 7, 111,
	2, 2, 347, 440, 7, 103, 2, 2, 348, 349, 7, 69, 2, 2, 349, 350, 7, 99, 2,
	2, 350, 351, 7, 112, 2, 2, 351, 352, 7, 47, 2, 2, 352, 353, 7, 84, 2, 2,
	353, 354, 7, 103, 2, 2, 354, 355, 7, 102, 2, 2, 355, 356, 7, 103, 2, 2,
	356, 357, 7, 104, 2, 2, 357, 358, 7, 107, 2, 2, 358, 359, 7, 112, 2, 2,
	359, 360, 7, 103, 2, 2, 360, 361, 7, 47, 2, 2, 361, 362, 7, 69, 2, 2, 362,
	363, 7, 110, 2, 2, 363, 364, 7, 99, 2, 2, 364, 365, 7, 117, 2, 2, 365,
	366, 7, 117, 2, 2, 366, 367, 7, 103, 2, 2, 367, 440, 7, 117, 2, 2, 368,
	369, 7, 82, 2, 2, 369, 370, 7, 99, 2, 2, 370, 371, 7, 101, 2, 2, 371, 372,
	7, 109, 2, 2, 372, 373, 7, 99, 2, 2, 373, 374, 7, 105, 2, 2, 374, 440,
	7, 103, 2, 2, 375, 376, 7, 85, 2, 2, 376, 377, 7, 74, 2, 2, 377, 378, 7,
	67, 2, 2, 378, 379, 7, 51, 2, 2, 379, 380, 7, 47, 2, 2, 380, 381, 7, 70,
	2, 2, 381, 382, 7, 107, 2, 2, 382, 383, 7, 105, 2, 2, 383, 384, 7, 103,
	2, 2, 384, 385, 7, 117, 2, 2, 385, 440, 7, 118, 2, 2, 386, 387, 7, 99,
	2, 2, 387, 388, 7, 112, 2, 2, 388, 389, 7, 118, 2, 2, 389, 390, 3, 2, 2,
	2, 390, 391, 7, 47, 2, 2, 391, 395, 5, 77, 39, 2, 392, 394, 5, 91, 46,
	2, 393, 392, 3, 2, 2, 2, 394, 397, 3, 2, 2, 2, 395, 393, 3, 2, 2, 2, 395,
	396, 3, 2, 2, 2, 396, 408, 3, 2, 2, 2, 397, 395, 3, 2, 2, 2, 398, 399,
	7, 47, 2, 2, 399, 403, 5, 77, 39, 2, 400, 402, 5, 91, 46, 2, 401, 400,
	3, 2, 2, 2, 402, 405, 3, 2, 2, 2, 403, 401, 3, 2, 2, 2, 403, 404, 3, 2,
	2, 2, 404, 407, 3, 2, 2, 2, 405, 403, 3, 2, 2, 2, 406, 398, 3, 2, 2, 2,
	407, 410, 3, 2, 2, 2, 408, 406, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409,
	440, 3, 2, 2, 2, 410, 408, 3, 2, 2, 2, 411, 415, 5, 77, 39, 2, 412, 414,
	5, 91, 46, 2, 413, 412, 3, 2, 2, 2, 414, 417, 3, 2, 2, 2, 415, 413, 3,
	2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 418, 3, 2, 2, 2, 417, 415, 3, 2, 2,
	2, 418, 419, 7, 47, 2, 2, 419, 423, 5, 77, 39, 2, 420, 422, 5, 91, 46,
	2, 421, 420, 3, 2, 2, 2, 422, 425, 3, 2, 2, 2, 423, 421, 3, 2, 2, 2, 423,
	424, 3, 2, 2, 2, 424, 436, 3, 2, 2, 2, 425, 423, 3, 2, 2, 2, 426, 427,
	7, 47, 2, 2, 427, 431, 5, 77, 39, 2, 428, 430, 5, 91, 46, 2, 429, 428,
	3, 2, 2, 2, 430, 433, 3, 2, 2, 2, 431, 429, 3, 2, 2, 2, 431, 432, 3, 2,
	2, 2, 432, 435, 3, 2, 2, 2, 433, 431, 3, 2, 2, 2, 434, 426, 3, 2, 2, 2,
	435, 438, 3, 2, 2, 2, 436, 434, 3, 2, 2, 2, 436, 437, 3, 2, 2, 2, 437,
	440, 3, 2, 2, 2, 438, 436, 3, 2, 2, 2, 439, 97, 3, 2, 2, 2, 439, 113, 3,
	2, 2, 2, 439, 129, 3, 2, 2, 2, 439, 139, 3, 2, 2, 2, 439, 153, 3, 2, 2,
	2, 439, 167, 3, 2, 2, 2, 439, 181, 3, 2, 2, 2, 439, 192, 3, 2, 2, 2, 439,
	210, 3, 2, 2, 2, 439, 223, 3, 2, 2, 2, 439, 245, 3, 2, 2, 2, 439, 258,
	3, 2, 2, 2, 439, 277, 3, 2, 2, 2, 439, 299, 3, 2, 2, 2, 439, 319, 3, 2,
	2, 2, 439, 330, 3, 2, 2, 2, 439, 344, 3, 2, 2, 2, 439, 348, 3, 2, 2, 2,
	439, 368, 3, 2, 2, 2, 439, 375, 3, 2, 2, 2, 439, 386, 3, 2, 2, 2, 439,
	411, 3, 2, 2, 2, 440, 4, 3, 2, 2, 2, 441, 448, 5, 7, 4, 2, 442, 444, 5,
	79, 40, 2, 443, 442, 3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 445, 3, 2,
	2, 2, 445, 447, 5, 7, 4, 2, 446, 443, 3, 2, 2, 2, 447, 450, 3, 2, 2, 2,
	448, 446, 3, 2, 2, 2, 448, 449, 3, 2, 2, 2, 449, 6, 3, 2, 2, 2, 450, 448,
	3, 2, 2, 2, 451, 460, 9, 2, 2, 2, 452, 453, 7, 64, 2, 2, 453, 460, 7, 63,
	2, 2, 454, 456, 5, 93, 47, 2, 455, 454, 3, 2, 2, 2, 456, 457, 3, 2, 2,
	2, 457, 455, 3, 2, 2, 2, 457, 458, 3, 2, 2, 2, 458, 460, 3, 2, 2, 2, 459,
	451, 3, 2, 2, 2, 459, 452, 3, 2, 2, 2, 459, 455, 3, 2, 2, 2, 460, 461,
	3, 2, 2, 2, 461, 459, 3, 2, 2, 2, 461, 462, 3, 2, 2, 2, 462, 8, 3, 2, 2,
	2, 463, 464, 7, 60, 2, 2, 464, 10, 3, 2, 2, 2, 465, 466, 7, 42, 2, 2, 466,
	12, 3, 2, 2, 2, 467, 468, 7, 43, 2, 2, 468, 14, 3, 2, 2, 2, 469, 470, 7,
	125, 2, 2, 470, 16, 3, 2, 2, 2, 471, 472, 7, 127, 2, 2, 472, 18, 3, 2,
	2, 2, 473, 474, 7, 93, 2, 2, 474, 20, 3, 2, 2, 2, 475, 476, 7, 95, 2, 2,
	476, 22, 3, 2, 2, 2, 477, 478, 7, 61, 2, 2, 478, 24, 3, 2, 2, 2, 479, 480,
	7, 46, 2, 2, 480, 26, 3, 2, 2, 2, 481, 482, 7, 48, 2, 2, 482, 28, 3, 2,
	2, 2, 483, 484, 7, 63, 2, 2, 484, 30, 3, 2, 2, 2, 485, 486, 7, 64, 2, 2,
	486, 32, 3, 2, 2, 2, 487, 488, 7, 62, 2, 2, 488, 34, 3, 2, 2, 2, 489, 490,
	7, 35, 2, 2, 490, 36, 3, 2, 2, 2, 491, 492, 7, 128, 2, 2, 492, 38, 3, 2,
	2, 2, 493, 494, 7, 65, 2, 2, 494, 40, 3, 2, 2, 2, 495, 496, 7, 63, 2, 2,
	496, 497, 7, 63, 2, 2, 497, 42, 3, 2, 2, 2, 498, 499, 7, 62, 2, 2, 499,
	500, 7, 63, 2, 2, 500, 44, 3, 2, 2, 2, 501, 502, 7, 64, 2, 2, 502, 503,
	7, 63, 2, 2, 503, 46, 3, 2, 2, 2, 504, 505, 7, 35, 2, 2, 505, 506, 7, 63,
	2, 2, 506, 48, 3, 2, 2, 2, 507, 508, 7, 40, 2, 2, 508, 509, 7, 40, 2, 2,
	509, 50, 3, 2, 2, 2, 510, 511, 7, 126, 2, 2, 511, 512, 7, 126, 2, 2, 512,
	52, 3, 2, 2, 2, 513, 514, 7, 45, 2, 2, 514, 515, 7, 45, 2, 2, 515, 54,
	3, 2, 2, 2, 516, 517, 7, 47, 2, 2, 517, 518, 7, 47, 2, 2, 518, 56, 3, 2,
	2, 2, 519, 520, 7, 45, 2, 2, 520, 58, 3, 2, 2, 2, 521, 522, 7, 47, 2, 2,
	522, 60, 3, 2, 2, 2, 523, 524, 7, 44, 2, 2, 524, 62, 3, 2, 2, 2, 525, 526,
	7, 49, 2, 2, 526, 64, 3, 2, 2, 2, 527, 528, 7, 40, 2, 2, 528, 66, 3, 2,
	2, 2, 529, 530, 7, 126, 2, 2, 530, 68, 3, 2, 2, 2, 531, 532, 7, 96, 2,
	2, 532, 70, 3, 2, 2, 2, 533, 534, 7, 39, 2, 2, 534, 72, 3, 2, 2, 2, 535,
	536, 7, 36, 2, 2, 536, 74, 3, 2, 2, 2, 537, 538, 7, 60, 2, 2, 538, 539,
	7, 63, 2, 2, 539, 76, 3, 2, 2, 2, 540, 541, 9, 3, 2, 2, 541, 78, 3, 2,
	2, 2, 542, 543, 9, 4, 2, 2, 543, 80, 3, 2, 2, 2, 544, 546, 7, 15, 2, 2,
	545, 544, 3, 2, 2, 2, 545, 546, 3, 2, 2, 2, 546, 547, 3, 2, 2, 2, 547,
	549, 7, 12, 2, 2, 548, 550, 7, 34, 2, 2, 549, 548, 3, 2, 2, 2, 549, 550,
	3, 2, 2, 2, 550, 551, 3, 2, 2, 2, 551, 552, 8, 41, 2, 2, 552, 82, 3, 2,
	2, 2, 553, 558, 7, 36, 2, 2, 554, 557, 10, 5, 2, 2, 555, 557, 5, 87, 44,
	2, 556, 554, 3, 2, 2, 2, 556, 555, 3, 2, 2, 2, 557, 560, 3, 2, 2, 2, 558,
	556, 3, 2, 2, 2, 558, 559, 3, 2, 2, 2, 559, 561, 3, 2, 2, 2, 560, 558,
	3, 2, 2, 2, 561, 562, 7, 36, 2, 2, 562, 84, 3, 2, 2, 2, 563, 567, 5, 91,
	46, 2, 564, 566, 5, 93, 47, 2, 565, 564, 3, 2, 2, 2, 566, 569, 3, 2, 2,
	2, 567, 565, 3, 2, 2, 2, 567, 568, 3, 2, 2, 2, 568, 86, 3, 2, 2, 2, 569,
	567, 3, 2, 2, 2, 570, 571, 7, 94, 2, 2, 571, 593, 9, 6, 2, 2, 572, 577,
	7, 94, 2, 2, 573, 575, 9, 7, 2, 2, 574, 573, 3, 2, 2, 2, 574, 575, 3, 2,
	2, 2, 575, 576, 3, 2, 2, 2, 576, 578, 9, 8, 2, 2, 577, 574, 3, 2, 2, 2,
	577, 578, 3, 2, 2, 2, 578, 579, 3, 2, 2, 2, 579, 593, 9, 8, 2, 2, 580,
	582, 7, 94, 2, 2, 581, 583, 7, 119, 2, 2, 582, 581, 3, 2, 2, 2, 583, 584,
	3, 2, 2, 2, 584, 582, 3, 2, 2, 2, 584, 585, 3, 2, 2, 2, 585, 586, 3, 2,
	2, 2, 586, 587, 5, 89, 45, 2, 587, 588, 5, 89, 45, 2, 588, 589, 5, 89,
	45, 2, 589, 590, 5, 89, 45, 2, 590, 593, 3, 2, 2, 2, 591, 593, 9, 9, 2,
	2, 592, 570, 3, 2, 2, 2, 592, 572, 3, 2, 2, 2, 592, 580, 3, 2, 2, 2, 592,
	591, 3, 2, 2, 2, 593, 88, 3, 2, 2, 2, 594, 595, 9, 10, 2, 2, 595, 90, 3,
	2, 2, 2, 596, 597, 9, 11, 2, 2, 597, 92, 3, 2, 2, 2, 598, 601, 5, 91, 46,
	2, 599, 601, 9, 12, 2, 2, 600, 598, 3, 2, 2, 2, 600, 599, 3, 2, 2, 2, 601,
	94, 3, 2, 2, 2, 602, 610, 9, 12, 2, 2, 603, 605, 9, 13, 2, 2, 604, 603,
	3, 2, 2, 2, 605, 608, 3, 2, 2, 2, 606, 604, 3, 2, 2, 2, 606, 607, 3, 2,
	2, 2, 607, 609, 3, 2, 2, 2, 608, 606, 3, 2, 2, 2, 609, 611, 9, 12, 2, 2,
	610, 606, 3, 2, 2, 2, 610, 611, 3, 2, 2, 2, 611, 96, 3, 2, 2, 2, 28, 2,
	395, 403, 408, 415, 423, 431, 436, 439, 443, 448, 457, 459, 461, 545, 549,
	556, 558, 567, 574, 577, 584, 592, 600, 606, 610, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "':'", "'('", "')'", "'{'", "'}'", "'['", "']'", "';'",
	"','", "'.'", "'='", "'>'", "'<'", "'!'", "'~'", "'?'", "'=='", "'<='",
	"'>='", "'!='", "'&&'", "'||'", "'++'", "'--'", "'+'", "'-'", "'*'", "'/'",
	"'&'", "'|'", "'^'", "'%'", "'\"'", "':='",
}

var lexerSymbolicNames = []string{
	"", "Key", "OTHERS", "ValueText", "COLON", "LPAREN", "RPAREN", "LBRACE",
	"RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA", "DOT", "ASSIGN", "GT", "LT",
	"BANG", "TILDE", "QUESTION", "EQUAL", "LE", "GE", "NOTEQUAL", "AND", "OR",
	"INC", "DEC", "ADD", "SUB", "MUL", "DIV", "BITAND", "BITOR", "CARET", "MOD",
	"DQUOTE", "SEQUAL", "Uppercase", "SPACE", "NL", "STRING_LITERAL", "IDENTIFIER",
}

var lexerRuleNames = []string{
	"Key", "OTHERS", "ValueText", "COLON", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
	"LBRACK", "RBRACK", "SEMI", "COMMA", "DOT", "ASSIGN", "GT", "LT", "BANG",
	"TILDE", "QUESTION", "EQUAL", "LE", "GE", "NOTEQUAL", "AND", "OR", "INC",
	"DEC", "ADD", "SUB", "MUL", "DIV", "BITAND", "BITOR", "CARET", "MOD", "DQUOTE",
	"SEQUAL", "Uppercase", "SPACE", "NL", "STRING_LITERAL", "IDENTIFIER", "EscapeSequence",
	"HexDigit", "Letter", "LetterOrDigit", "Digits",
}

type ManifestLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewManifestLexer(input antlr.CharStream) *ManifestLexer {

	l := new(ManifestLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Manifest.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ManifestLexer tokens.
const (
	ManifestLexerKey            = 1
	ManifestLexerOTHERS         = 2
	ManifestLexerValueText      = 3
	ManifestLexerCOLON          = 4
	ManifestLexerLPAREN         = 5
	ManifestLexerRPAREN         = 6
	ManifestLexerLBRACE         = 7
	ManifestLexerRBRACE         = 8
	ManifestLexerLBRACK         = 9
	ManifestLexerRBRACK         = 10
	ManifestLexerSEMI           = 11
	ManifestLexerCOMMA          = 12
	ManifestLexerDOT            = 13
	ManifestLexerASSIGN         = 14
	ManifestLexerGT             = 15
	ManifestLexerLT             = 16
	ManifestLexerBANG           = 17
	ManifestLexerTILDE          = 18
	ManifestLexerQUESTION       = 19
	ManifestLexerEQUAL          = 20
	ManifestLexerLE             = 21
	ManifestLexerGE             = 22
	ManifestLexerNOTEQUAL       = 23
	ManifestLexerAND            = 24
	ManifestLexerOR             = 25
	ManifestLexerINC            = 26
	ManifestLexerDEC            = 27
	ManifestLexerADD            = 28
	ManifestLexerSUB            = 29
	ManifestLexerMUL            = 30
	ManifestLexerDIV            = 31
	ManifestLexerBITAND         = 32
	ManifestLexerBITOR          = 33
	ManifestLexerCARET          = 34
	ManifestLexerMOD            = 35
	ManifestLexerDQUOTE         = 36
	ManifestLexerSEQUAL         = 37
	ManifestLexerUppercase      = 38
	ManifestLexerSPACE          = 39
	ManifestLexerNL             = 40
	ManifestLexerSTRING_LITERAL = 41
	ManifestLexerIDENTIFIER     = 42
)
