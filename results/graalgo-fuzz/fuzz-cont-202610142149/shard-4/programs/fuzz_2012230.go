package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int8 }
type GenS2 struct { F0 bool; F1 string; F2 uint32 }

func main() {
    var v1 int = 0
    v2 := make(map[int]bool)
    var v3 int32 = 1
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := []int{(^(fi5)), (^(v1))}
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(v3)
            v1 -= (-(fi5))
            if false {
                switch ("" + ("hello" + ("a" + ("\n\t" + "\n\t")))) {
                case "":
                    v3 = (v3 / (((v3 >> ((255) & 63))) | 1))
                case "foo bar":
                    if true {
                        v3 -= (v3 & v3)
                        v1 = fi7
                        var v8 uint64 = 42
                        v8 = v8
                        _ = v8
                    } else {
                        v6 = append(v6, (-(fi7)))
                    }
                }
                fmt.Println(v3)
                var v9 uint8 = 128
                _ = v9
            }
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    v3 -= v3
    v1 += v1
    switch 42 {
    case 816:
        v3 = 523
    case 42:
        v3 = v3
    case -2147483648:
        if (1e10 != 1.5) {
            v10 := v4[2147483648]
            _ = v10
        } else {
            if ((!true) || (true || (((!true) || true) || (false && (4294967295 == 2))))) {
                var v11 float32 = 0.1
                v3 += 2147483647
                for fi12 := 0; fi12 < 10; fi12++ {
                    if true {
                        v13 := GenS1{F0: v3, F1: 2, F2: 42}
                        fmt.Println(fi12, v3)
                        _ = v13
                    } else {
                        v3 = -2147483648
                    }
                    _ = fi12
                }
                _ = v11
            } else {
                var v14 uint8 = 255
                v3 += v3
                v15 := GenS1{F0: v3, F1: 255, F2: 1}
                _ = v14
                _ = v15
            }
        }
    default:
        fmt.Println(v1)
    }
    delete(v2, int(v3))
    v3 += 1
    for fi16 := 0; fi16 < 8; fi16++ {
        switch "hello" {
        case "hello":
            if true {
                fmt.Println(fi16, v3, fi16)
                v3 -= -2147483648
                v3 = 42
            }
        case "":
            v3 += v3
        }
        v17 := GenS2{F0: ((((0.0 - (1.5 - 1e10)) - 0.1) / ((0.1 * float64(float32(0.1))) + 3.14159)) > 3.14159), F1: "hello", F2: 2147483648}
        _ = fi16
        _ = v17
    }
    _ = len(v4)
    v3 += (v3 | (v3 | v3))
    v3 = 234
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
