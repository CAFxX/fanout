package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int; F1 uint8; F2 int }
type GenS3 struct { F0 string; F1 int; F2 uint8 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint; F1 int64 }
func genFn0(p0 int16, p1 int64, p2 uint64) uint32 {
    var v6 float32 = (1e10 * (1.5 / 3.14159))
    v7 := []int{}
    p1 += 502
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 340
}

func main() {
    v1 := []int{2147483647, 223, 820, 9223372036854775807}
    v2 := []string{}
    var v3 int16 = -1
    var v4 bool = (255 >= 1)
    var v5 uint16 = 65535
    v8 := GenS1{F0: "foo bar", F1: "foo bar"}
    var v9 uint8 = 1
    v5 = 721
    fmt.Println(v3)
    switch 378 {
    case 1:
        v10 := GenS2{F0: -487, F1: (v9 & v9), F2: 978}
        _ = v10
    case -154:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        var v11 uint = uint(v5)
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 float32 = 0.0
        _ = fi12
        _ = v13
    }
    if v4 {
        v3 -= v3
    } else {
        if v4 {
            v3 = (v3 / (((-(v3))) | 1))
            for fi14 := 0; fi14 < 16; fi14++ {
                if true {
                    var v15 int8 = 42
                    v2 = append(v2, "hello")
                    v16 := GenS3{F0: "", F1: (fi14 << ((4294967295) & 63)), F2: (v9 - 109)}
                    _ = v15
                    _ = v16
                }
                v17 := GenS1{F0: "foo bar", F1: "a"}
                switch (((("hello" + ("a" + (("\n\t" + ("hello" + "foo bar")) + "a"))) + ("\n\t" + "\n\t")) + ("foo bar" + (("hello" + "hello") + ""))) + "\n\t") {
                case "":
                    var v18 uint64 = 4294967295
                    _ = v18
                default:
                    fmt.Println(v4, v9)
                }
                _ = fi14
                _ = v17
            }
            fmt.Println(v5)
        } else {
            switch -263 {
            case -414:
                var v19 uint64 = 205
                _ = v19
            default:
                for fi20 := 0; fi20 < 12; fi20++ {
                    v21 := GenS3{F0: "foo bar", F1: (fi20 | fi20), F2: 1}
                    v4 = v4
                    v9 = (v9 / (((v9 ^ 1)) | 1))
                    _ = fi20
                    _ = v21
                }
            }
            v22 := GenS4{F0: -32768, F1: -292}
            for fi23 := 0; fi23 < 15; fi23++ {
                v24 := GenS5{F0: 18446744073709551615, F1: (^(v3)), F2: 4294967295}
                _ = fi23
                _ = v24
            }
            _ = v22
        }
        if (-2.25 <= float64(3.14159)) {
            v25 := GenS6{F0: 2, F1: 42}
            var v26 uint8 = 42
            _ = v25
            _ = v26
        }
        var v27 uint64 = 61
        _ = v27
    }
    var v28 int16 = (v3 - (v3 % (((v3 & (^(v3)))) | 1)))
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 bool = (v9 == v9)
        v28 = v28
        for fi31 := 0; fi31 < 7; fi31++ {
            v32 := GenS5{F0: 4294967295, F1: v3, F2: 1}
            fmt.Println(v5, v3)
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    fmt.Println(v9, v9, v5)
    v3 = (v3 % ((-1) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
