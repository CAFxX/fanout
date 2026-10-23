package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint16; F2 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    var v3 string = (("" + "") + ("a" + "hello"))
    var v4 int32 = 1
    var v5 uint16 = 42
    for fi6 := 0; fi6 < 14; fi6++ {
        v4 += (v4 / (((v4 / (((^(v4))) | 1))) | 1))
        switch v3 {
        case "a":
            for fi7 := 0; fi7 < 2; fi7++ {
                if (!(1 >= 9223372036854775807)) {
                    fmt.Println(fi7, fi6, fi6)
                }
                _ = fi7
            }
        case "\n\t":
            if (true || (-9223372036854775808 < -398)) {
                v8 := GenS1{F0: fi6, F1: (fi6 + (fi6 % ((-9223372036854775808) | 1))), F2: false}
                _ = v8
            }
        case "":
            for fi9 := 0; fi9 < 14; fi9++ {
                v10 := []int{}
                _ = fi9
                _ = v10
            }
        default:
            var v11 int8 = 127
            _ = v11
        }
        v12 := []int{(fi6 * 1), (^(fi6))}
        _ = fi6
        _ = v12
    }
    if (!true) {
        v5 -= v5
        v3 = "a"
    }
    if (42 <= -92) {
        var v13 uint = 0
        _ = v13
    }
    v3 = ((v3 + v3) + "a")
    v14 := GenS1{F0: 2147483647, F1: -1, F2: (-33 < 100)}
    var v15 int8 = 127
    v1[368] = 127
    if (false && false) {
        v16 := GenS2{F0: 919}
        v15 -= v15
        var v17 uint = 2
        _ = v16
        _ = v17
    } else {
        var v18 int64 = 495
        _ = v18
    }
    v19 := GenS3{F0: 42, F1: 78, F2: 4294967295}
    if (float32(-2.25) != (1e10 * 0.0)) {
        for fi20 := 0; fi20 < 11; fi20++ {
            if false {
                var v21 float32 = 1.5
                _ = v21
            }
            _ = fi20
        }
        var v22 bool = false
        _ = v22
    } else {
        v23 := GenS4{F0: 18446744073709551615, F1: 65535, F2: (-(v4))}
        if ((159 <= (v5 >> ((645) & 63))) && false) {
            var v24 float32 = 1e10
            fmt.Println(v24)
            _ = v24
        }
        _ = v23
    }
    fmt.Println(v4, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
