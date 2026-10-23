package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 string }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    var v3 uint = 893
    v4 := make(map[uint32]uint32)
    var v5 uint = (v3 ^ (v3 - (v3 | (v3 % ((42) | 1)))))
    v3 -= v3
    var v6 uint8 = 83
    fmt.Println(v5, v6)
    var v7 uint8 = (v6 % (((^(v6))) | 1))
    v8 := []int{1, 4294967295}
    if true {
        delete(v4, 0)
        if (("a" + ("\n\t" + "")) == "a") {
            switch 1 {
            case 32768:
                var v9 float32 = (0.1 - 3.14159)
                _ = v9
            case -1:
                _ = len(v8)
            case -393:
                _ = len(v8)
            }
            v5 = 42
            v7 = v7
        } else {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            v3 = 255
            var v10 float64 = float64(1.5)
            _ = v10
        }
    } else {
        fmt.Println(v3, v3)
    }
    var v11 int16 = 42
    var v12 float64 = (0.0 * 3.14159)
    for fi13 := 0; fi13 < 16; fi13++ {
        v3 += 2
        _ = fi13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := GenS1{F0: (fi14 / ((1) | 1)), F1: 668, F2: "foo bar"}
        _ = fi14
        _ = v15
    }
    var v16 uint8 = v6
    if false {
        if ((false && ((-(v11)) != -993)) || (true || true)) {
            v17 := GenS2{F0: -2147483648, F1: 100}
            if true {
                var v18 uint16 = 1
                v7 = 230
                _ = v18
            }
            _ = v17
        }
        if (!false) {
            v19 := GenS1{F0: 32767, F1: 0, F2: (("a" + "hello") + "")}
            _ = v19
        }
    } else {
        v20 := GenS2{F0: -910, F1: -84}
        var v21 int16 = (v11 % (((v11 * v11)) | 1))
        for fi22 := 0; fi22 < 6; fi22++ {
            fmt.Println(v11, v16, v3)
            v23 := GenS3{F0: v5}
            _ = fi22
            _ = v23
        }
        _ = v20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
