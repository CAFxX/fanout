package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int8 }
func genFn0(p0 string) string {
    v7 := []int{143}
    if (p0 <= p0) {
        var v8 float64 = 0.1
        p0 = ((("" + "") + ((p0 + (((p0 + "foo bar") + "") + ("foo bar" + ""))) + p0)) + p0)
        v9 := GenS1{F0: p0, F1: -1, F2: 2147483648}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v7
    return ""
}
func genFn1(p0 int64) int {
    if (!(1 != 0)) {
        v10 := GenS2{F0: 1, F1: 2, F2: -128}
        p0 += p0
        _ = v10
    } else {
        v11 := []int{-373, 127, 0, -9223372036854775808}
        _ = v11
    }
    v12 := GenS2{F0: 694, F1: 308, F2: -84}
    p0 = p0
    if false {
        p0 += (p0 + -1)
        v13 := GenS2{F0: 1, F1: 0, F2: 0}
        _ = v13
    }
    _ = p0
    _ = v12
    return -39
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{}
    var v3 int64 = (v1 + v1)
    v4 := make(map[uint32]string)
    var v5 int32 = 42
    var v6 float64 = (-2.25 / 0.1)
    v4[2] = ("a" + ("foo bar" + ("foo bar" + "a")))
    v4[783] = "hello"
    _ = len(v2)
    for fi14 := 0; fi14 < 9; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            if false {
                v4[2147483647] = ("\n\t" + "")
            } else {
                var v16 string = ("\n\t" + "")
                switch "hello" {
                case "hello":
                    var v17 uint8 = 0
                    _ = v17
                case "":
                    if ((-128 != 127) || (true && (!false))) {
                        fmt.Println(v1, v5, v1)
                    }
                    var v18 int64 = 79
                    _ = v18
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v19 float32 = 0.0
                _ = v16
                _ = v19
            }
            _ = fi15
        }
        v2 = append(v2, genFn1(v3))
        switch "foo bar" {
        case "\n\t":
            for fi20 := 0; fi20 < 10; fi20++ {
                fmt.Println(v5, v3)
                v5 += v5
                _ = fi20
            }
        case "":
            var v21 int8 = -1
            _ = v21
        case "a":
            v4[65535] = "\n\t"
        default:
            fmt.Println(v5, v5)
        }
        _ = fi14
    }
    if ((1.5 / (0.1 + 3.14159)) == 1.5) {
        fmt.Println(v5, v5)
        var v22 bool = true
        _ = v22
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS2{F0: 156, F1: 92, F2: 42}
        if true {
            for fi25 := 0; fi25 < 4; fi25++ {
                v1 += (^(v3))
                _ = fi25
            }
            switch (fi23 << ((737) & 63)) {
            case 4294967295:
                v1 = v3
            default:
                v1 += 42
            }
        }
        _ = fi23
        _ = v24
    }
    var v26 int = -127
    v3 += 1
    var v27 uint16 = 774
    var v28 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
