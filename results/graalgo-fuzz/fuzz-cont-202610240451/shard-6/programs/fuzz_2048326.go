package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 int64; F2 bool }
func genFn0(p0 string, p1 uint32, p2 int16) int8 {
    v5 := GenS1{F0: 42, F1: 765, F2: 1}
    fmt.Println(p1, p2)
    v6 := []int{-9223372036854775808, 32768, 9223372036854775807, -9223372036854775808}
    p0 = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 98
}

func main() {
    v1 := []int{525, 2147483647, 813}
    v2 := []uint32{918}
    var v3 int16 = 0
    var v4 uint = 381
    var v7 uint16 = 1
    for fi8 := 0; fi8 < 15; fi8++ {
        if ((v7 ^ v7) > v7) {
            v3 += v3
            v9 := GenS2{F0: 1}
            _ = v9
        }
        fmt.Println(v3)
        _ = fi8
    }
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        v10 := GenS3{F0: -59}
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(v3, v7)
            var v12 int = -1
            _ = fi11
            _ = v12
        }
    default:
        switch ("\n\t" + "hello") {
        case "hello":
            if true {
                v3 = v3
            } else {
                v13 := GenS4{F0: (v3 & v3), F1: 1, F2: 0}
                for fi14 := 0; fi14 < 14; fi14++ {
                    if (true && (false || (!false))) {
                        if true {
                            v3 = (v3 + v3)
                            fmt.Println(v7, v7, v4)
                            var v15 bool = (42 != 116)
                            _ = v15
                        }
                        fmt.Println(v3)
                    } else {
                        if (1e10 != (-2.25 / 1.0)) {
                            var v16 uint64 = 2
                            for fi17 := 0; fi17 < 12; fi17++ {
                                var v18 uint16 = (v7 - v7)
                                if (fi14 != fi14) {
                                    v3 = 42
                                    fmt.Println(v4)
                                }
                                var v19 uint64 = 18446744073709551615
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            _ = v16
                        }
                        fmt.Println(v3)
                        v4 -= v4
                    }
                    v3 = -1
                    _ = fi14
                }
                _ = v13
            }
        default:
            fmt.Println(v4, v7)
        }
    }
    var v20 int8 = 127
    if (42 < -107) {
        v21 := GenS3{F0: v20}
        var v22 float64 = -2.25
        _ = v21
        _ = v22
    }
    v4 -= v4
    v23 := GenS5{F0: -227}
    var v24 float64 = 3.14159
    switch 9223372036854775807 {
    case 9223372036854775807:
        if false {
            for fi25 := 0; fi25 < 17; fi25++ {
                v24 += float64(((float32(1e10) + 1.5) - 0.0))
                _ = fi25
            }
            var v26 uint16 = v7
            _ = v26
        }
    case 232:
        if ((((v20 >> (((v4 >> ((v4) & 63))) & 63)) <= v20) || false) || (!("hello" <= "a"))) {
            switch 127 {
            case -595:
                v7 += (v7 + (v7 * v7))
            }
            for fi27 := 0; fi27 < 4; fi27++ {
                v2 = append(v2, 202)
                v28 := GenS6{F0: "foo bar", F1: 434, F2: false}
                var v29 uint = v4
                _ = fi27
                _ = v28
                _ = v29
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v7
}
