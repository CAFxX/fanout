package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
func genFn0(p0 int16) bool {
    var v7 int16 = (^(p0))
    fmt.Println(v7)
    fmt.Println(v7, p0)
    v7 = v7
    var v8 int32 = -1
    _ = p0
    _ = v7
    _ = v8
    return (((p0 * (p0 % ((-32768) | 1))) >= 0) || false)
}
func genFn1(p0 int32) int32 {
    if ((!(802 < 4294967295)) && false) {
        switch "" {
        case "foo bar":
            v9 := GenS1{F0: 100, F1: 2, F2: ""}
            _ = v9
        case "\n\t":
            var v10 int64 = 390
            _ = v10
        case "a":
            v11 := []int{-1, -2147483648, 9223372036854775807, -2147483648}
            _ = v11
        default:
            p0 = -2147483648
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            p0 = (-(p0))
            var v13 float32 = (0.0 - (1.5 / 1.5))
            _ = fi12
            _ = v13
        }
    }
    if (1e10 == (0.0 / (float64(((1.5 - (3.14159 * -2.25)) / 1e10)) - float64(((0.0 - float32(3.14159)) / 1e10))))) {
        fmt.Println(p0)
        p0 -= (p0 << ((418) & 63))
        if (true || true) {
            for fi14 := 0; fi14 < 7; fi14++ {
                fmt.Println(fi14, fi14)
                v15 := GenS1{F0: 42, F1: 4294967295, F2: ""}
                _ = fi14
                _ = v15
            }
            var v16 uint64 = 4294967295
            _ = v16
        } else {
            p0 = 42
            for fi17 := 0; fi17 < 4; fi17++ {
                var v18 float64 = 1e10
                var v19 uint8 = 2
                var v20 string = "hello"
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            var v21 uint16 = 42
            _ = v21
        }
    }
    _ = p0
    return p0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 32767
    var v3 int16 = 1000
    var v4 uint32 = 65535
    var v5 int16 = (^(v3))
    var v6 int64 = 871
    var v22 string = ""
    v23 := GenS2{F0: 955}
    v2 -= 1
    v24 := GenS2{F0: 4294967295}
    if true {
        v25 := GenS3{F0: -24}
        for fi26 := 0; fi26 < 6; fi26++ {
            v27 := GenS2{F0: 0}
            if (((v22 + v22) <= v22) && true) {
                for fi28 := 0; fi28 < 16; fi28++ {
                    var v29 uint8 = 42
                    _ = fi28
                    _ = v29
                }
                v30 := GenS3{F0: (-(v6))}
                v22 = v22
                _ = v30
            }
            _ = fi26
            _ = v27
        }
        var v31 uint32 = v4
        _ = v25
        _ = v31
    } else {
        v32 := []int{128, 32767}
        _ = v32
    }
    var v33 float64 = 3.14159
    v34 := GenS1{F0: 127, F1: 18446744073709551615, F2: (v22 + (v22 + ""))}
    var v35 float32 = 0.1
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
