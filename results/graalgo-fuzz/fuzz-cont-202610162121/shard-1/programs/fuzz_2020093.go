package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int64, p1 int, p2 string) bool {
    for fi4 := 0; fi4 < 5; fi4++ {
        if false {
            p1 = (fi4 & fi4)
            p1 = (p1 << ((51) & 63))
        }
        v5 := []int{}
        _ = fi4
        _ = v5
    }
    v6 := []int{(^(p1)), (p1 * 32767), (p1 << ((753) & 63)), p1}
    for fi7 := 0; fi7 < 7; fi7++ {
        p0 -= (p0 | (p0 & (p0 ^ p0)))
        _ = fi7
    }
    var v8 int = (p1 >> ((1) & 63))
    var v9 uint32 = 989
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    _ = v9
    return true
}
func genFn1(p0 int8, p1 uint32, p2 uint) uint32 {
    p0 -= (p0 << (((p2 - 562)) & 63))
    var v10 bool = (true && (4294967295 > 495))
    var v11 uint32 = p1
    switch "foo bar" {
    case "hello":
        if (true || true) {
            var v12 float64 = (0.1 - 3.14159)
            var v13 int = 32767
            v13 = v13
            _ = v12
            _ = v13
        } else {
            var v14 uint16 = 225
            _ = v14
        }
    case "":
        var v15 uint32 = 261
        _ = v15
    case "a":
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int = 208
            _ = fi16
            _ = v17
        }
    }
    v11 -= (v11 - v11)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return (v11 / ((v11) | 1))
}

func main() {
    var v1 int = 2147483647
    v2 := []string{"foo bar", "a", "foo bar"}
    var v3 int = -371
    var v18 uint64 = 1
    var v19 uint16 = 1
    var v20 uint8 = 240
    if true {
        v2 = append(v2, "\n\t")
        switch "hello" {
        case "hello":
            v21 := GenS1{F0: 1}
            _ = v21
        case "foo bar":
            if true {
                var v22 int = 0
                v2 = append(v2, "")
                _ = len(v2)
                _ = v22
            }
        case "\n\t":
            if ("foo bar" != "hello") {
                if (-2.25 < 0.1) {
                    for fi23 := 0; fi23 < 17; fi23++ {
                        var v24 uint8 = (v20 - (v20 ^ v20))
                        v25 := GenS1{F0: 425}
                        fmt.Println(v24)
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v3 += 0
                        _ = fi26
                    }
                } else {
                    fmt.Println(v19)
                }
            }
        default:
            v2 = append(v2, "")
        }
    }
    v1 = 1
    v19 = v19
    if len(v2) > 0 {
        v2[0] = ("\n\t" + "")
    }
    if true {
        v20 -= 0
        v1 += (v1 * -1)
    }
    if false {
        if len(v2) > 0 {
            v2[0] = ("a" + "")
        }
        v27 := GenS1{F0: 344}
        _ = v27
    }
    var v28 float64 = 0.1
    v29 := GenS1{F0: 862}
    v1 = int(v18)
    v20 -= (^(v20))
    if len(v2) > 0 {
        v2[0] = ("" + "foo bar")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
}
