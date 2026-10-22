package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint8 }
type GenS2 struct { F0 string }
func genFn0(p0 int8, p1 int8) int8 {
    switch "hello" {
    case "a":
        switch "a" {
        case "hello":
            for fi6 := 0; fi6 < 16; fi6++ {
                var v7 uint64 = 4294967295
                p1 -= (p1 << ((0) & 63))
                _ = fi6
                _ = v7
            }
        case "foo bar":
            if (true || true) {
                fmt.Println(p1, p1)
                v8 := GenS1{F0: 32767, F1: 1, F2: 1}
                for fi9 := 0; fi9 < 17; fi9++ {
                    var v10 float64 = 1e10
                    if true {
                        v10 -= v10
                        p0 -= -126
                    }
                    _ = fi9
                    _ = v10
                }
                _ = v8
            } else {
                var v11 uint16 = 42
                p0 -= p1
                _ = v11
            }
        }
    case "":
        fmt.Println(p0, p1)
    default:
        var v12 uint16 = 65535
        _ = v12
    }
    p1 -= p0
    var v13 int16 = 32767
    var v14 uint16 = 255
    fmt.Println(v13, p0, v14)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return (p0 + p0)
}
func genFn1(p0 string) uint {
    var v15 uint16 = 2
    if false {
        v16 := []int{1, -9223372036854775808, -214}
        var v17 int32 = -559
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 18; fi18++ {
            v15 = v15
            _ = fi18
        }
        p0 = ""
        p0 = "hello"
    }
    _ = p0
    _ = v15
    return 0
}

func main() {
    var v1 bool = false
    var v2 float32 = (1.5 / -2.25)
    var v3 float64 = 0.0
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]uint32)
    v1 = false
    v1 = ((2 != 1) || ((42 >= 58) || ((v1 || (((1 <= 2147483648) && false) && ((42 >= 1) || false))) || (-2.25 < (v3 * v3)))))
    fmt.Println(v1, v2)
    for fi19 := 0; fi19 < 11; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            fmt.Println(v1)
            _ = fi20
        }
        switch (((("\n\t" + "hello") + "foo bar") + ("\n\t" + "foo bar")) + "\n\t") {
        case "foo bar":
            v21 := GenS2{F0: ("foo bar" + "foo bar")}
            _ = v21
        case "":
            var v22 string = ""
            _ = v22
        case "\n\t":
            var v23 string = "foo bar"
            _ = v23
        default:
            v1 = v1
        }
        switch fi19 {
        case 112:
            v24 := []int{(fi19 % ((32767) | 1)), 739, (fi19 ^ fi19), -1}
            _ = v24
        case -594:
            if true {
                v25 := []int{(-(fi19)), fi19, (^(fi19)), fi19}
                v26 := GenS2{F0: "hello"}
                _ = v25
                _ = v26
            }
        case 912:
            v3 = (v3 + v3)
        default:
            var v27 float32 = ((v2 - v2) + float32(v3))
            _ = v27
        }
        _ = fi19
    }
    v28 := GenS2{F0: ""}
    if v1 {
        v3 -= ((float64(((v2 + (float32(1e10) * v2)) * (v2 + v2))) * 3.14159) + v3)
        var v29 int16 = 809
        _ = v29
    }
    var v30 int64 = -331
    v2 -= 3.14159
    v31 := []int{32767, 32767, 42}
    var v32 int = 4294967295
    v30 += (v30 >> ((0) & 63))
    fmt.Println(v3)
    if v1 {
        v33 := GenS2{F0: ("\n\t" + "a")}
        _ = v33
    } else {
        v2 = ((v2 - v2) * v2)
        v3 += (0.1 * ((v3 / float64(v2)) * 1e10))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
