package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint32 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 string; F1 string; F2 int }
type GenS4 struct { F0 int8; F1 uint32; F2 string }
func genFn0() int64 {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: 100, F1: "a", F2: 2147483648}
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 254
    _ = v8
    return -513
}

func main() {
    v1 := []int{-539}
    v2 := make(map[int]int)
    var v3 float64 = 3.14159
    var v4 float64 = v3
    var v5 int = 42
    switch (v5 - 868) {
    case 0:
        v9 := GenS1{F0: 0, F1: "a", F2: 124}
        _ = v9
    case 295:
        switch "a" {
        case "":
            var v10 uint32 = 1
            _ = v10
        case "foo bar":
            v1 = append(v1, 32767)
        }
    }
    var v11 uint16 = 42
    v4 -= float64(((3.14159 - 0.0) / 3.14159))
    var v12 int = v5
    if false {
        var v13 float64 = v3
        if (v13 != (1e10 + float64((3.14159 + 1.5)))) {
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 int64 = 0
                switch (v5 + fi14) {
                case 2147483647:
                    fmt.Println(v13, v12, fi14)
                default:
                    v16 := GenS2{F0: (2 != 886), F1: -1}
                    _ = v16
                }
                for fi17 := 0; fi17 < 15; fi17++ {
                    v18 := v2[(v5 | (fi14 >> ((65535) & 63)))]
                    v15 = genFn0()
                    v19 := GenS1{F0: 42, F1: ((("" + "hello") + (((("\n\t" + "a") + "") + "") + "hello")) + ""), F2: 106}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi14
                _ = v15
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := GenS3{F0: ("hello" + ("foo bar" + "a")), F1: (("\n\t" + "foo bar") + ""), F2: v5}
                v22 := v2[4294967295]
                _ = fi20
                _ = v21
                _ = v22
            }
        } else {
            fmt.Println(v3, v11, v11)
            v23 := v2[(v12 % ((v5) | 1))]
            fmt.Println(v13, v13, v23)
            _ = v23
        }
        var v24 uint64 = 543
        _ = v13
        _ = v24
    } else {
        var v25 int = 32767
        v5 += 2147483647
        _ = v25
    }
    var v26 uint16 = (v11 / (((v11 / (((v11 * v11)) | 1))) | 1))
    v3 = 1.5
    if (127 != 234) {
        var v27 float32 = 0.1
        v1 = append(v1, v5)
        v2[v12] = (v12 + (-(v12)))
        _ = v27
    }
    v28 := GenS4{F0: 42, F1: 3, F2: ("\n\t" + "\n\t")}
    for fi29 := 0; fi29 < 11; fi29++ {
        v12 -= (v12 / ((v5) | 1))
        _ = fi29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
