package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint; F2 uint32 }
type GenS4 struct { F0 string }

func main() {
    v1 := []string{"", "a"}
    var v2 int16 = -689
    var v3 int8 = -128
    v4 := make(map[string]string)
    var v5 bool = true
    if (255 <= 801) {
        v3 = v3
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    fmt.Println(v2, v3)
    fmt.Println(v3)
    fmt.Println(v3, v3)
    if true {
        v6 := GenS1{F0: 2147483647}
        v7 := GenS1{F0: 585}
        v3 -= (v3 >> ((851) & 63))
        _ = v6
        _ = v7
    }
    v3 -= v3
    for fi8 := 0; fi8 < 2; fi8++ {
        v2 = 32767
        v4[("foo bar" + ("\n\t" + ("" + "")))] = (("foo bar" + ("foo bar" + "hello")) + "")
        switch (("\n\t" + "") + "") {
        case "foo bar":
            if (-2.25 != 1.5) {
                v3 -= v3
                for fi9 := 0; fi9 < 19; fi9++ {
                    delete(v4, "")
                    var v10 float64 = 0.1
                    _ = fi9
                    _ = v10
                }
            }
        }
        _ = fi8
    }
    var v11 uint16 = 1
    if (v5 && v5) {
        if (-2.25 > -2.25) {
            v3 -= 127
            var v12 int = 42
            v13 := GenS2{F0: 1}
            _ = v12
            _ = v13
        } else {
            _ = len(v4)
        }
        switch "" {
        case "foo bar":
            v2 += 533
        default:
            if v5 {
                for fi14 := 0; fi14 < 8; fi14++ {
                    for fi15 := 0; fi15 < 2; fi15++ {
                        v2 += 1000
                        v16 := GenS3{F0: -128, F1: 65535, F2: 128}
                        _ = fi15
                        _ = v16
                    }
                    v17 := GenS4{F0: "\n\t"}
                    v5 = v5
                    _ = fi14
                    _ = v17
                }
                fmt.Println(v5, v3, v11)
            } else {
                v5 = v5
            }
        }
    }
    var v18 bool = false
    if ((uint64(v11) == 971) || v5) {
        for fi19 := 0; fi19 < 8; fi19++ {
            fmt.Println(v3)
            _ = fi19
        }
        switch "" {
        case "a":
            v1 = append(v1, "a")
        }
    } else {
        if ((42 >= 0) || false) {
            fmt.Println(v5)
            v20 := GenS1{F0: -74}
            v21 := GenS4{F0: "hello"}
            _ = v20
            _ = v21
        }
        var v22 float64 = float64(1.5)
        var v23 uint64 = 923
        _ = v22
        _ = v23
    }
    var v24 bool = true
    v3 -= (v3 + (v3 | (v3 << ((478) & 63))))
    v3 = (v3 / ((-39) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
