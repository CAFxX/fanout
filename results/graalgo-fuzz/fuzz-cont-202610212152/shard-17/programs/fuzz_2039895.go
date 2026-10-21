package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint32; F2 int32 }
func genFn0(p0 uint, p1 int32, p2 string) uint16 {
    if false {
        switch 42 {
        case 811:
            v7 := []int{-2147483648, int(p1)}
            _ = v7
        default:
            var v8 uint16 = 42
            _ = v8
        }
    }
    var v9 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}
func genFn1(p0 uint8) int16 {
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 uint16 = 466
        _ = fi10
        _ = v11
    }
    v12 := []int{128, 32768}
    _ = p0
    _ = v12
    return 32767
}

func main() {
    var v1 bool = true
    v2 := make(map[string]bool)
    var v3 float32 = 1e10
    var v4 uint8 = 0
    var v5 int64 = -300
    v6 := []uint32{2147483647, 2147483648}
    v5 -= 1
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := GenS1{F0: "\n\t"}
        _ = fi13
        _ = v14
    }
    if (!v1) {
        v3 -= float32(3.14159)
        v1 = (v5 == 42)
    } else {
        if ((float32(1.5) == v3) && v1) {
            v15 := GenS2{F0: 42, F1: 2147483647, F2: 42}
            var v16 string = "hello"
            _ = v15
            _ = v16
        }
        fmt.Println(v4, v1, v4)
        if false {
            v1 = v1
            var v17 uint32 = 429
            _ = len(v2)
            _ = v17
        }
    }
    switch 789 {
    case 127:
        if false {
            for fi18 := 0; fi18 < 20; fi18++ {
                var v19 float32 = v3
                for fi20 := 0; fi20 < 5; fi20++ {
                    v5 = v5
                    var v21 uint64 = 145
                    _ = fi20
                    _ = v21
                }
                _ = fi18
                _ = v19
            }
        } else {
            if false {
                v22 := GenS1{F0: ("a" + "foo bar")}
                _ = v22
            } else {
                for fi23 := 0; fi23 < 18; fi23++ {
                    for fi24 := 0; fi24 < 11; fi24++ {
                        var v25 string = (("\n\t" + "foo bar") + "a")
                        _ = fi24
                        _ = v25
                    }
                    v26 := GenS1{F0: ""}
                    _ = fi23
                    _ = v26
                }
            }
        }
    case -2147483648:
        v5 = (v5 - (v5 | (v5 - v5)))
    default:
        v27 := GenS2{F0: -1, F1: 1, F2: -1}
        _ = v27
    }
    switch 28 {
    case -1:
        fmt.Println(v5, v4, v1)
    case 1:
        v3 = (v3 + v3)
    }
    var v28 uint64 = 2
    switch "\n\t" {
    case "\n\t":
        v5 = v5
    case "foo bar":
        if (!(!(!(v28 >= (v28 ^ v28))))) {
            v28 += v28
        }
    case "a":
        switch "hello" {
        case "\n\t":
            v1 = (612 == 2147483648)
        case "hello":
            v4 -= v4
        case "":
            v29 := GenS1{F0: "\n\t"}
            _ = v29
        }
    }
    v1 = (18446744073709551615 >= 18446744073709551615)
    var v30 uint32 = 446
    switch 117 {
    case 1:
        var v31 uint16 = 2
        _ = v31
    case -2147483648:
        for fi32 := 0; fi32 < 13; fi32++ {
            var v33 string = "a"
            _ = fi32
            _ = v33
        }
    case 177:
        v30 -= (v30 % (((^(v30))) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
