package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint8 }
func genFn0() uint64 {
    var v5 float32 = (1e10 - 0.0)
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: 2147483648}
        _ = fi6
        _ = v7
    }
    _ = v5
    return 42
}
func genFn1(p0 uint64, p1 int16, p2 int) uint16 {
    p2 = p2
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 int32 = 0
    var v2 int = -2147483648
    v3 := []int{}
    v4 := []int{}
    v8 := GenS1{F0: 0}
    var v9 float64 = (3.14159 - (0.1 * 1.5))
    v2 -= v2
    switch "a" {
    case "\n\t":
        if (((true || (("a" > "a") || true)) && false) || true) {
            switch "" {
            case "\n\t":
                v10 := GenS2{F0: 169}
                _ = v10
            case "foo bar":
                fmt.Println(v1, v9, v1)
            case "a":
                var v11 uint64 = 807
                _ = v11
            default:
                if (v2 == (v2 % ((205) | 1))) {
                    v12 := GenS1{F0: 42}
                    switch ("a" + "\n\t") {
                    case "hello":
                        var v13 int64 = 0
                        _ = v13
                    case "":
                        v2 += (v2 + (v2 ^ v2))
                    case "a":
                        v14 := GenS2{F0: 1}
                        _ = v14
                    default:
                        v9 = (v9 + v9)
                    }
                    _ = v12
                } else {
                    v9 += 0.0
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v9 = (v9 - float64((1.5 * 3.14159)))
                }
            }
        } else {
            v4 = append(v4, v2)
        }
    case "hello":
        switch "" {
        case "a":
            var v15 uint32 = 1
            _ = v15
        default:
            fmt.Println(v1)
        }
    case "foo bar":
        for fi16 := 0; fi16 < 9; fi16++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v4 = append(v4, v2)
            if (-1 != 32767) {
                v1 = v1
            } else {
                v9 += v9
                if (!(fi16 < v2)) {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        var v18 uint8 = 127
                        _ = fi17
                        _ = v18
                    }
                    v3 = append(v3, 42)
                    var v19 uint = 65535
                    _ = v19
                } else {
                    v2 = v2
                }
            }
            _ = fi16
        }
    default:
        var v20 uint64 = 4294967295
        _ = v20
    }
    var v21 float64 = (((v9 - 1.5) - (3.14159 + v9)) * (v9 + v9))
    if (!(32767 <= 2)) {
        if len(v3) > 0 {
            v3[0] = 32768
        }
        var v22 float32 = 1e10
        var v23 bool = (!(false && true))
        _ = v22
        _ = v23
    } else {
        var v24 int = 4294967295
        _ = v24
    }
    v1 = 42
    v25 := GenS2{F0: 1}
    v26 := GenS1{F0: 2147483647}
    v1 += 42
    fmt.Println(v21, v1, v21)
    if false {
        for fi27 := 0; fi27 < 6; fi27++ {
            for fi28 := 0; fi28 < 10; fi28++ {
                v29 := GenS3{F0: fi28, F1: fi28}
                v30 := GenS4{F0: 33}
                _ = fi28
                _ = v29
                _ = v30
            }
            fmt.Println(v2, fi27, v9)
            v1 = (v1 >> ((641) & 63))
            _ = fi27
        }
    } else {
        switch "a" {
        case "\n\t":
            fmt.Println(v9)
        case "":
            if len(v3) > 0 {
                v3[0] = v2
            }
        case "a":
            var v31 uint64 = 4294967295
            _ = v31
        }
        var v32 float64 = (-2.25 - v21)
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
