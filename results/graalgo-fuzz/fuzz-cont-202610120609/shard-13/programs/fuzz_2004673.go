package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int32; F1 bool; F2 bool }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int16; F1 int32; F2 uint32 }
type GenS5 struct { F0 string }
func genFn0(p0 int8, p1 int, p2 string) int8 {
    p2 = ((p2 + "a") + "")
    v5 := []int{}
    _ = len(v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}

func main() {
    v1 := []int{-1, -9223372036854775808}
    v2 := []string{}
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    v6 := GenS1{F0: 42, F1: 241}
    v7 := GenS1{F0: 9223372036854775807, F1: 2147483648}
    var v8 int16 = 0
    v9 := GenS2{F0: 0, F1: (534 >= 183), F2: false}
    v8 -= v8
    for fi10 := 0; fi10 < 16; fi10++ {
        switch 32767 {
        case 32767:
            fmt.Println(fi10)
        case 128:
            var v11 uint = 1
            _ = v11
        default:
            var v12 float64 = -2.25
            _ = v12
        }
        if (!(!false)) {
            for fi13 := 0; fi13 < 20; fi13++ {
                if false {
                    var v14 int16 = (v8 + (v8 << ((2) & 63)))
                    _ = v14
                }
                if (false || (((3.14159 * float64((((0.1 - 3.14159) + 0.0) + (3.14159 + float32(0.1))))) + (1.5 / 3.14159)) > -2.25)) {
                    var v15 int32 = 0
                    _ = v15
                } else {
                    v3["foo bar"] = true
                }
                _ = fi13
            }
            fmt.Println(v8, v8)
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi10
    }
    if true {
        v2 = append(v2, "")
        switch (("\n\t" + "a") + ("a" + "\n\t")) {
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v8 -= v8
        }
        if ((!true) || (625 >= 2)) {
            var v16 uint64 = 243
            switch "a" {
            case "\n\t":
                v17 := GenS3{F0: 108, F1: 255}
                _ = v17
            case "":
                if (!true) {
                    var v18 int = -442
                    if (v8 > v8) {
                        v8 = v8
                    } else {
                        fmt.Println(v18, v18, v18)
                        v19 := GenS3{F0: -84, F1: 313}
                        var v20 float64 = 1e10
                        _ = v19
                        _ = v20
                    }
                    fmt.Println(v18)
                    _ = v18
                }
            default:
                v16 -= (v16 + v16)
            }
            var v21 uint8 = 51
            _ = v16
            _ = v21
        } else {
            var v22 float32 = 0.1
            _ = v22
        }
    }
    v8 += -32768
    if false {
        if (992 >= 0) {
            if true {
                v23 := GenS4{F0: (v8 & 0), F1: 42, F2: 65535}
                v24 := GenS2{F0: -2147483648, F1: ("a" <= ""), F2: (!false)}
                _ = v23
                _ = v24
            } else {
                v8 = 595
                if false {
                    fmt.Println(v8)
                } else {
                    v8 = v8
                    v25 := GenS5{F0: ("hello" + "\n\t")}
                    var v26 int8 = 1
                    _ = v25
                    _ = v26
                }
                v8 += (v8 / ((v8) | 1))
            }
            v27 := GenS4{F0: v8, F1: 0, F2: 42}
            var v28 uint8 = 127
            _ = v27
            _ = v28
        } else {
            v8 = v8
            v29 := GenS1{F0: 473, F1: 4294967295}
            fmt.Println(v8)
            _ = v29
        }
        v8 -= 1
    } else {
        v30 := GenS5{F0: ""}
        delete(v4, "\n\t")
        fmt.Println(v8)
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
