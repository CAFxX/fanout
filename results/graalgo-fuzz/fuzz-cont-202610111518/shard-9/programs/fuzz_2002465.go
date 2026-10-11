package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 bool) int {
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return 1
}
func genFn1(p0 uint8, p1 int8) int {
    var v7 uint16 = 42
    var v8 bool = true
    p0 += (p0 / (((^(p0))) | 1))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    v4 := make(map[int]uint32)
    var v5 int = 552
    v6 := make(map[uint32]string)
    v5 += -9223372036854775808
    switch ("a" + ("\n\t" + "")) {
    case "a":
        v5 = 128
    case "\n\t":
        fmt.Println(v5)
    }
    fmt.Println(v5)
    var v9 int32 = 0
    if false {
        if (!(false && false)) {
            v10 := GenS1{F0: 255}
            switch (v5 - v5) {
            case -460:
                v11 := v3["foo bar"]
                _ = v11
            case 32767:
                var v12 float32 = (((3.14159 - 0.1) + ((-2.25 * (1.5 * 1e10)) * 3.14159)) / (float32(1e10) + 0.1))
                _ = v12
            }
            v13 := []int{v5, 32767, 0, v5}
            _ = v10
            _ = v13
        } else {
            fmt.Println(v5, v9)
        }
        var v14 float64 = ((float64(((0.1 + (-2.25 + (-2.25 + 3.14159))) * 0.1)) + 1e10) * 1e10)
        _ = v14
    }
    var v15 uint64 = 658
    switch int(v15) {
    case 127:
        for fi16 := 0; fi16 < 16; fi16++ {
            if (!(!(!true))) {
                v17 := GenS1{F0: 128}
                switch ("hello" + "\n\t") {
                case "":
                    var v18 int32 = v9
                    _ = v18
                default:
                    var v19 bool = (false || false)
                    _ = v19
                }
                v9 = -2147483648
                _ = v17
            } else {
                v20 := []int{(v5 + genFn0((-1 < 672))), 1}
                _ = v20
            }
            _ = fi16
        }
    case 42:
        var v21 uint8 = 187
        _ = v21
    case -9223372036854775808:
        v22 := v2[v5]
        _ = v22
    default:
        v23 := GenS2{F0: 0}
        _ = v23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 float32 = 0.1
        _ = fi24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
