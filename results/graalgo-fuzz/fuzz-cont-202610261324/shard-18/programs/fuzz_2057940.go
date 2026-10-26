package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 string }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    if ((((float32(float64(1e10)) - -2.25) - ((0.1 * 0.1) + 0.1)) / (float32(3.14159) * 0.1)) != 1e10) {
        v1 -= 255
        v4 := GenS1{F0: 42, F1: (^(v1)), F2: "foo bar"}
        _ = v4
    } else {
        var v5 float64 = ((0.1 / ((0.1 + -2.25) / 3.14159)) / -2.25)
        v1 += 122
        _ = v5
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        if ("foo bar" != "\n\t") {
            v1 = (v1 + v1)
        }
        v7 := GenS1{F0: 127, F1: v1, F2: ("a" + ("hello" + ""))}
        _ = fi6
        _ = v7
    }
    v8 := v2[127]
    var v9 float32 = (1e10 * 0.0)
    if (true && (2 == 65535)) {
        var v10 int8 = -1
        v11 := GenS2{F0: 227, F1: ""}
        if true {
            var v12 uint64 = 2
            var v13 float64 = float64(v9)
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 int64 = 42
                var v16 uint8 = (v1 & v1)
                v10 = v10
                _ = fi14
                _ = v15
                _ = v16
            }
            if ("\n\t" == "\n\t") {
                if ((!false) || false) {
                    var v17 string = v8
                    _ = v17
                }
            } else {
                v9 += 0.0
                fmt.Println(v10)
            }
        }
        _ = v10
        _ = v11
    }
    v18 := GenS2{F0: v1, F1: (v8 + "a")}
    v19 := GenS3{F0: -860}
    var v20 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
