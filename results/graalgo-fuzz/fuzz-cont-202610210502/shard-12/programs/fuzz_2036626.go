package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 bool }
type GenS2 struct { F0 uint16; F1 int8; F2 uint16 }
func genFn0() string {
    v7 := GenS1{F0: 42, F1: 2, F2: true}
    var v8 float32 = -2.25
    _ = v7
    _ = v8
    return "a"
}
func genFn1(p0 string) int32 {
    v9 := []int{-348, -1, -9223372036854775808}
    p0 = "foo bar"
    switch 9223372036854775807 {
    case 32768:
        var v10 int16 = 0
        _ = v10
    case 42:
        fmt.Println(p0)
    case 127:
        v11 := GenS2{F0: 1, F1: -1, F2: 2}
        _ = v11
    default:
        p0 = ((p0 + "\n\t") + ("hello" + "a"))
    }
    _ = p0
    _ = v9
    return 757
}

func main() {
    var v1 int32 = -680
    v2 := make(map[uint32]uint32)
    var v3 int8 = 100
    var v4 int64 = 0
    var v5 float32 = float32((1.5 - float64(1e10)))
    v6 := make(map[string]string)
    for fi12 := 0; fi12 < 3; fi12++ {
        v5 = v5
        switch ("a" + "a") {
        case "a":
            v13 := GenS2{F0: 2, F1: (^(v3)), F2: 32767}
            _ = v13
        default:
            if (fi12 > (fi12 + fi12)) {
                v14 := []int{(fi12 << ((1) & 63)), (fi12 ^ 423)}
                var v15 int = (^(fi12))
                _ = v14
                _ = v15
            } else {
                var v16 int32 = (-(v1))
                fmt.Println(v5, v3, v5)
                var v17 string = (("a" + "\n\t") + "a")
                _ = v16
                _ = v17
            }
        }
        var v18 int32 = int32(v4)
        _ = fi12
        _ = v18
    }
    if (!(-767 <= -680)) {
        var v19 uint32 = 460
        _ = v19
    } else {
        v20 := []int{-211, 1, -618, 127}
        switch -148 {
        case 32767:
            var v21 bool = false
            _ = v21
        default:
            v22 := GenS2{F0: 665, F1: v3, F2: 57}
            _ = v22
        }
        _ = v20
    }
    v23 := GenS1{F0: 913, F1: 1, F2: false}
    var v24 int8 = v3
    switch "" {
    case "foo bar":
        var v25 int64 = v4
        _ = v25
    case "a":
        switch "foo bar" {
        case "hello":
            var v26 int64 = 0
            _ = v26
        case "foo bar":
            switch "hello" {
            case "hello":
                v27 := []int{-2147483648, 2147483647, 0, 936}
                _ = v27
            case "a":
                v5 = (0.1 / v5)
            default:
                if false {
                    v28 := v2[361]
                    _ = v28
                } else {
                    if (((259 >= 4294967295) && (-2.25 <= ((-2.25 + ((3.14159 - (-2.25 + 1.5)) / 3.14159)) / 3.14159))) && (1 >= -1)) {
                        _ = len(v2)
                        v1 = v1
                        var v29 int8 = 0
                        _ = v29
                    } else {
                        v30 := GenS1{F0: 18446744073709551615, F1: 0, F2: false}
                        _ = v30
                    }
                    v31 := []int{0, -2147483648, 127}
                    _ = v31
                }
            }
        case "\n\t":
            if (!false) {
                for fi32 := 0; fi32 < 14; fi32++ {
                    v4 -= -311
                    switch "a" {
                    case "\n\t":
                        var v33 int = (^(fi32))
                        _ = v33
                    case "a":
                        var v34 uint32 = 2147483647
                        _ = v34
                    default:
                        fmt.Println(v24, fi32)
                    }
                    v35 := GenS2{F0: 2, F1: (v24 * v3), F2: 255}
                    _ = fi32
                    _ = v35
                }
            } else {
                var v36 uint64 = 252
                _ = v36
            }
        default:
            switch 128 {
            case 42:
                v1 += (v1 - int32(v4))
            default:
                var v37 bool = false
                _ = v37
            }
        }
    case "\n\t":
        var v38 int16 = 1
        _ = v38
    default:
        var v39 uint64 = 382
        _ = v39
    }
    if (!(-32768 > 0)) {
        var v40 uint = 835
        var v41 bool = (-2147483648 < int(v3))
        v5 = v5
        _ = v40
        _ = v41
    }
    var v42 float32 = v5
    var v43 uint16 = 185
    v4 += v4
    v44 := GenS1{F0: 1, F1: v43, F2: (!(((v1 + v1) < (-(v1))) || (!((18446744073709551615 > 2) || false))))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
