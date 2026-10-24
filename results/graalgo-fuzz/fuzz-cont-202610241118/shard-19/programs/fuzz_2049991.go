package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int16; F2 string }
func genFn0(p0 uint8) uint64 {
    v6 := []int{-2147483648, 32767, -9223372036854775808}
    v7 := GenS1{F0: 529}
    fmt.Println(p0)
    _ = p0
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 uint32, p1 string, p2 int16) int64 {
    v8 := []int{2147483647}
    switch 0 {
    case -9223372036854775808:
        var v9 string = (p1 + p1)
        _ = v9
    case -226:
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS1{F0: 32767}
        var v12 int16 = p2
        _ = fi10
        _ = v11
        _ = v12
    }
    v8 = append(v8, 127)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -981
}

func main() {
    v1 := []string{("hello" + "foo bar")}
    var v2 uint8 = 128
    var v3 float64 = float64(3.14159)
    var v4 bool = false
    v5 := make(map[uint32]int)
    if (((0.1 * (-2.25 * (float64(0.0) * 3.14159))) * ((3.14159 * (((v3 / v3) - float64(1e10)) * v3)) / v3)) != v3) {
        if v4 {
            v2 = (v2 ^ 32)
        } else {
            var v13 int16 = 42
            switch 127 {
            case 30:
                var v14 int16 = v13
                _ = v14
            case 2147483647:
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 uint = 65535
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi15
                    _ = v16
                }
            default:
                v17 := GenS1{F0: 2}
                _ = v17
            }
            _ = v13
        }
    }
    var v18 uint32 = 5
    v4 = v4
    if (("hello" + "") <= (("\n\t" + "") + "a")) {
        _ = len(v5)
    }
    switch "\n\t" {
    case "":
        var v19 float32 = 0.0
        _ = v19
    case "a":
        v3 = 3.14159
    case "foo bar":
        var v20 bool = (!v4)
        _ = v20
    default:
        fmt.Println(v4, v4)
    }
    v21 := GenS1{F0: 255}
    v5[(v18 << ((129) & 63))] = -9223372036854775808
    v22 := GenS1{F0: 32767}
    v3 += (v3 + 0.0)
    var v23 int8 = -128
    switch 2147483647 {
    case 52:
        switch "\n\t" {
        case "":
            var v24 uint16 = 87
            _ = v24
        case "hello":
            v25 := GenS1{F0: 372}
            _ = v25
        default:
            var v26 float32 = (0.1 - 1e10)
            _ = v26
        }
    case -629:
        switch 9223372036854775807 {
        case -9223372036854775808:
            v4 = (18446744073709551615 <= uint64(v2))
        }
    case 645:
        var v27 int8 = v23
        _ = v27
    }
    if v4 {
        for fi28 := 0; fi28 < 9; fi28++ {
            if v4 {
                v18 -= 451
                if (float64(3.14159) >= v3) {
                    var v29 float64 = -2.25
                    var v30 int8 = (v23 % (((v23 << ((1) & 63))) | 1))
                    _ = v29
                    _ = v30
                } else {
                    var v31 int64 = -1
                    _ = v31
                }
                v2 -= v2
            } else {
                v32 := GenS1{F0: 903}
                _ = v32
            }
            for fi33 := 0; fi33 < 15; fi33++ {
                var v34 uint8 = 128
                fmt.Println(v18, fi28, v3)
                var v35 int16 = 32767
                _ = fi33
                _ = v34
                _ = v35
            }
            var v36 int16 = 42
            _ = fi28
            _ = v36
        }
        switch 433 {
        case 1:
            fmt.Println(v3)
        case 32768:
            var v37 int64 = 1
            _ = v37
        default:
            var v38 uint32 = (v18 >> ((0) & 63))
            _ = v38
        }
    } else {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        v39 := GenS2{F0: 1, F1: int16(v2), F2: "hello"}
        v2 = (v2 % ((204) | 1))
        _ = v39
    }
    v40 := GenS1{F0: 1}
    var v41 uint32 = v18
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
