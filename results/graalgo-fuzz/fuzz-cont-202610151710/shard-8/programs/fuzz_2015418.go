package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint64 }

func main() {
    var v1 int = 32768
    var v2 uint8 = 128
    v3 := make(map[int]string)
    var v4 int32 = 1
    for fi5 := 0; fi5 < 14; fi5++ {
        if false {
            v1 += fi5
        }
        _ = fi5
    }
    switch "\n\t" {
    case "\n\t":
        if (0.1 <= (((-2.25 + (-2.25 + 3.14159)) / (-2.25 * 0.1)) * ((3.14159 * 0.0) * 0.0))) {
            var v6 int16 = -32768
            var v7 int8 = 127
            _ = v6
            _ = v7
        } else {
            v2 -= 1
            var v8 int64 = int64(v1)
            _ = v8
        }
    }
    v1 -= v1
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: (v4 ^ int32(v2))}
        _ = fi9
        _ = v10
    }
    switch 19 {
    case -597:
        switch v1 {
        case 0:
            v11 := GenS2{F0: (("\n\t" + ("\n\t" + "")) + "foo bar"), F1: 4294967295}
            _ = v11
        case 32767:
            _ = len(v3)
        case 979:
            v12 := GenS1{F0: -108}
            _ = v12
        }
    case 2147483647:
        v4 -= 739
    }
    v2 = v2
    v3[(v1 % (((v1 | v1)) | 1))] = ""
    v1 = (v1 * (v1 | (^(v1))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
