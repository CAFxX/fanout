package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint64, p1 int8, p2 uint8) bool {
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 float32 = float32(-2.25)
        _ = fi5
        _ = v6
    }
    p0 = 2
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    var v1 uint64 = 42
    var v2 int64 = -9223372036854775808
    var v3 int = 42
    var v4 bool = true
    v7 := []int{9223372036854775807, v3, (v3 ^ -2147483648)}
    var v8 int8 = 127
    var v9 int32 = 42
    if v4 {
        v10 := GenS1{F0: "hello", F1: (-(v9))}
        _ = v10
    } else {
        v1 = v1
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            if (0.0 > (0.0 * 3.14159)) {
                var v13 int16 = 1000
                v14 := GenS1{F0: "\n\t", F1: int32(v1)}
                v15 := GenS2{F0: 42}
                _ = v13
                _ = v14
                _ = v15
            } else {
                v7 = append(v7, (-(fi11)))
                if (v2 < (^(v2))) {
                    if v4 {
                        var v16 string = ("\n\t" + ("" + ("" + "foo bar")))
                        _ = v16
                    }
                } else {
                    v2 += v2
                }
            }
            _ = fi12
        }
        _ = fi11
    }
    v7 = append(v7, v3)
    var v17 int8 = (v8 ^ 42)
    switch "hello" {
    case "":
        var v18 bool = ((v1 & (v1 - v1)) != (^(v1)))
        _ = v18
    case "a":
        v19 := GenS3{F0: (v3 < v3), F1: 1}
        _ = v19
    case "foo bar":
        v17 = 1
    default:
        switch (("a" + ("hello" + "")) + "foo bar") {
        case "foo bar":
            v1 -= (v1 * (v1 - 1))
        case "hello":
            switch v3 {
            case 127:
                var v20 int8 = v17
                _ = v20
            case 4294967295:
                v3 = (v3 & 419)
            }
        case "":
            v4 = (!true)
        default:
            v8 += (v17 | (v8 ^ (v17 >> ((42) & 63))))
        }
    }
    var v21 float64 = 1.5
    v7 = append(v7, v3)
    v22 := GenS2{F0: 911}
    var v23 uint64 = v1
    if (!v4) {
        v7 = append(v7, (v3 % (((^(v3))) | 1)))
    }
    var v24 int16 = -1
    v25 := GenS4{F0: v17}
    _ = len(v7)
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 uint8 = 255
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
