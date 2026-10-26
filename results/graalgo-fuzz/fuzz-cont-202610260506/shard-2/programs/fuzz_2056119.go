package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }

func main() {
    var v1 int32 = 510
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    var v4 int32 = (v1 ^ v1)
    var v5 int16 = -406
    v6 := []int{42, 4294967295, 522, 2147483647}
    v7 := GenS1{F0: 0}
    v1 += v1
    var v8 uint16 = 65535
    switch 32767 {
    case 32768:
        v9 := GenS2{F0: -549}
        _ = v9
    case 329:
        v8 = 255
    case 127:
        var v10 uint64 = 2
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint8 = 42
        _ = fi11
        _ = v12
    }
    if false {
        fmt.Println(v1, v4)
        var v13 int32 = (v1 & -1)
        v1 += v4
        _ = v13
    }
    var v14 float64 = 0.1
    v1 -= 169
    var v15 string = "hello"
    v2[4294967295] = 111
    if true {
        if (1 >= 212) {
            fmt.Println(v5)
            _ = len(v2)
            var v16 int = 42
            _ = v16
        }
        v4 = (-(v1))
        var v17 uint = 0
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        delete(v2, 1)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v15 = v15
        _ = fi18
    }
    if (v5 != (v5 % ((32767) | 1))) {
        for fi19 := 0; fi19 < 14; fi19++ {
            v8 = (v8 | (^(v8)))
            for fi20 := 0; fi20 < 14; fi20++ {
                switch v15 {
                case "\n\t":
                    v14 = (v14 / v14)
                default:
                    var v21 uint = 42
                    _ = v21
                }
                _ = fi20
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                switch ("a" + v15) {
                case "foo bar":
                    fmt.Println(v14, v14)
                case "":
                    if (false || true) {
                        v6 = append(v6, fi22)
                        fmt.Println(v15)
                        v14 = float64(1.5)
                    } else {
                        var v23 int32 = v4
                        _ = v23
                    }
                default:
                    fmt.Println(v5)
                }
                if (true || true) {
                    v15 = "a"
                    var v24 int64 = -9223372036854775808
                    _ = v24
                }
                _ = fi22
            }
            _ = fi19
        }
        v15 = ("\n\t" + "")
    }
    if (3.14159 <= v14) {
        v25 := v3[42]
        delete(v3, 845)
        _ = v25
    } else {
        if false {
            var v26 uint64 = 2
            _ = v26
        }
        var v27 int64 = 0
        _ = v27
    }
    var v28 int8 = 0
    if (18446744073709551615 != 65535) {
        var v29 float32 = 0.1
        switch 560 {
        case 4294967295:
            var v30 float32 = (v29 - v29)
            _ = v30
        }
        _ = v29
    } else {
        fmt.Println(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
