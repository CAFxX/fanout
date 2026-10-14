package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 2
    var v3 float32 = (3.14159 / (((((3.14159 / (float32((((0.1 + 1e10) - 0.1) + (0.1 / 1e10))) / 1.5)) - (1e10 + 1e10)) - 1e10) + 1e10) - (0.0 - float32(3.14159))))
    var v4 int32 = 42
    v5 := make(map[int]uint32)
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: v2, F1: "\n\t"}
        if false {
            if false {
                fmt.Println(fi7, v3)
                var v9 int = (fi7 | -737)
                _ = v9
            } else {
                v2 = 255
            }
            v10 := GenS2{F0: 0}
            _ = v10
        } else {
            v11 := GenS3{F0: ""}
            var v12 uint8 = 224
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    if (359 == 1) {
        v3 = (((v3 + 0.1) - (1.5 + 1.5)) * v3)
        var v13 uint64 = 460
        _ = len(v1)
        _ = v13
    }
    if (!false) {
        v14 := []int{128, 4294967295, -2147483648, 0}
        _ = v14
    }
    var v15 uint = 18446744073709551615
    v16 := GenS2{F0: 0}
    switch 0 {
    case 4294967295:
        for fi17 := 0; fi17 < 19; fi17++ {
            v4 -= int32(fi17)
            _ = fi17
        }
    case 42:
        var v18 string = ("a" + "\n\t")
        _ = v18
    }
    var v19 uint = (v2 >> ((v2) & 63))
    v4 = (v4 % (((v4 & -2147483648)) | 1))
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 bool = (!false)
        for fi22 := 0; fi22 < 2; fi22++ {
            v23 := GenS1{F0: v19, F1: ("\n\t" + ("foo bar" + ""))}
            _ = fi22
            _ = v23
        }
        var v24 float32 = v3
        _ = fi20
        _ = v21
        _ = v24
    }
    var v25 bool = true
    var v26 uint64 = 0
    var v27 uint32 = 1
    v27 -= (v27 * v27)
    v4 = (v4 | (-(v4)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
