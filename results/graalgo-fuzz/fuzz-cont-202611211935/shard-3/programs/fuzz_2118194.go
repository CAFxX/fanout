package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int16 }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0() int32 {
    var v5 uint64 = 1
    if (0 < 2) {
        var v6 int16 = -1
        _ = v6
    }
    fmt.Println(v5)
    _ = v5
    return 2147483647
}
func genFn1(p0 bool) int64 {
    if (0 <= 0) {
        var v7 uint16 = 317
        p0 = p0
        v8 := GenS1{F0: -1, F1: "foo bar", F2: 1000}
        _ = v7
        _ = v8
    } else {
        if (133 < 2) {
            var v9 int32 = 42
            switch 0 {
            case 32767:
                v10 := []int{2147483647, -63, -130}
                _ = v10
            case 72:
                v11 := []int{50}
                _ = v11
            default:
                var v12 uint = 35
                _ = v12
            }
            _ = v9
        }
    }
    v13 := []int{}
    p0 = p0
    _ = p0
    _ = v13
    return 591
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 344
    var v3 float32 = ((-2.25 / 1.0) + 0.0)
    var v4 int64 = 1
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint16 = 1
        if true {
            v16 := []int{(fi14 / ((fi14) | 1))}
            v2 = v2
            _ = v16
        }
        var v17 uint16 = v15
        _ = fi14
        _ = v15
        _ = v17
    }
    v3 = v3
    var v18 uint32 = 419
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := []int{-883}
        _ = fi19
        _ = v20
    }
    var v21 bool = ((!false) || false)
    fmt.Println(v2, v4, v2)
    var v22 uint8 = 128
    v23 := GenS2{F0: 100, F1: (^(v18))}
    v3 -= 1e10
    var v24 int = -2147483648
    var v25 uint16 = 520
    var v26 int = v24
    if true {
        var v27 uint16 = (v25 * (v25 % ((32767) | 1)))
        _ = v27
    } else {
        if v21 {
            v28 := GenS2{F0: 1, F1: 4294967295}
            _ = v28
        } else {
            switch ("" + "\n\t") {
            case "\n\t":
                if (((v18 <= (v18 >> ((65535) & 63))) && (v21 && true)) && false) {
                    for fi29 := 0; fi29 < 6; fi29++ {
                        for fi30 := 0; fi30 < 20; fi30++ {
                            v25 = (v25 >> ((831) & 63))
                            _ = fi30
                        }
                        _ = fi29
                    }
                } else {
                    v31 := v1["\n\t"]
                    _ = v31
                }
            case "foo bar":
                v32 := GenS2{F0: 0, F1: 65535}
                _ = v32
            default:
                if (-753 != 0) {
                    v33 := GenS1{F0: 350, F1: "hello", F2: -32768}
                    _ = v33
                }
            }
        }
        v3 += (float32(((1e10 * float64(v3)) + (-2.25 - (3.14159 / -2.25)))) - (float32(3.14159) / 1.0))
        v34 := []int{v24, v26, (v26 << ((1) & 63))}
        _ = v34
    }
    if (true && (true || v21)) {
        var v35 bool = v21
        fmt.Println(v3, v22)
        _ = v35
    } else {
        delete(v1, "hello")
        if (!(((v2 % ((v2) | 1)) > uint64(v18)) || true)) {
            v2 -= v2
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
