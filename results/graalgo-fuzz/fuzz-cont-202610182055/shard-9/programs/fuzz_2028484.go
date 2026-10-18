package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 string; F2 uint8 }
func genFn0(p0 int, p1 uint16) string {
    v7 := []int{}
    p1 += (p1 * p1)
    _ = p0
    _ = p1
    _ = v7
    return "hello"
}

func main() {
    var v1 uint32 = 2
    var v2 bool = (406 > 50)
    var v3 int32 = -2147483648
    var v4 float64 = (-2.25 + 0.0)
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]string)
    switch -52 {
    case 128:
        var v8 float32 = (0.1 * (1e10 / 3.14159))
        _ = v8
    default:
        fmt.Println(v4, v4)
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        if (1 == -23) {
            v10 := []int{(fi9 - 9223372036854775807), 42, -451, -896}
            for fi11 := 0; fi11 < 15; fi11++ {
                v3 += (v3 | 0)
                _ = fi11
            }
            var v12 int64 = -9223372036854775808
            _ = v10
            _ = v12
        } else {
            delete(v6, (v1 ^ (v1 ^ v1)))
            v2 = (159 < 255)
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := []int{}
                _ = fi13
                _ = v14
            }
        }
        v15 := GenS1{F0: (v1 >> ((65535) & 63)), F1: int64(fi9)}
        var v16 int = 32768
        _ = fi9
        _ = v15
        _ = v16
    }
    switch 32767 {
    case 32768:
        var v17 int64 = 1
        _ = v17
    case -9223372036854775808:
        var v18 string = ((("\n\t" + (("hello" + "") + ("" + "a"))) + "") + "\n\t")
        _ = v18
    case 9223372036854775807:
        switch "foo bar" {
        case "foo bar":
            var v19 int16 = 233
            _ = v19
        }
    }
    if v2 {
        v20 := GenS1{F0: 2147483648, F1: 178}
        if (-854 < 0) {
            v21 := GenS1{F0: v1, F1: -967}
            var v22 int8 = 65
            switch "" {
            case "a":
                delete(v5, 1)
            case "":
                if (v2 && ((v22 & int8(v3)) <= 100)) {
                    v23 := []int{}
                    _ = v23
                }
            default:
                var v24 string = "hello"
                _ = v24
            }
            _ = v21
            _ = v22
        } else {
            fmt.Println(v3, v4)
        }
        v3 = (v3 / ((-1) | 1))
        _ = v20
    }
    v25 := GenS2{F0: "\n\t"}
    if (69 < 1) {
        if v2 {
            fmt.Println(v2)
            v26 := GenS2{F0: "\n\t"}
            v1 = 2
            _ = v26
        }
        var v27 float64 = float64(0.1)
        v3 -= v3
        _ = v27
    }
    switch -364 {
    case -9223372036854775808:
        for fi28 := 0; fi28 < 19; fi28++ {
            v29 := GenS3{F0: v1, F1: "hello", F2: 253}
            v30 := []int{(fi28 | (fi28 | fi28))}
            _ = fi28
            _ = v29
            _ = v30
        }
    case -300:
        var v31 float32 = ((1e10 * float32((-2.25 / v4))) - (((float32(1e10) + float32(v4)) - -2.25) * float32(-2.25)))
        _ = v31
    case 343:
        switch "a" {
        case "a":
            var v32 uint8 = 128
            _ = v32
        case "\n\t":
            v33 := GenS2{F0: ((("hello" + "foo bar") + "a") + "")}
            _ = v33
        case "":
            var v34 int64 = 0
            _ = v34
        }
    }
    if (false || false) {
        if v2 {
            v35 := []int{460}
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
