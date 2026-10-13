package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 uint64, p1 string, p2 string) string {
    switch p1 {
    case "hello":
        switch (p2 + p1) {
        case "foo bar":
            var v4 int16 = 499
            _ = v4
        default:
            var v5 int32 = -842
            _ = v5
        }
    }
    p2 = "\n\t"
    v6 := []int{}
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}
func genFn1(p0 uint64, p1 int32, p2 uint8) int8 {
    fmt.Println(p2)
    v7 := GenS1{F0: 1}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 float64 = (float64(1.5) - (1e10 / 1.0))
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = 1000
    var v3 float32 = float32(0.1)
    v1 = ("a" + "")
    var v10 int16 = (v2 & v2)
    var v11 int8 = -1
    v12 := []int{127, 32767, -2147483648, 32767}
    var v13 float64 = 3.14159
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 int = -9223372036854775808
            v17 := GenS1{F0: 0}
            if (1.5 > float64(3.14159)) {
                v18 := GenS1{F0: 65535}
                fmt.Println(fi15, v13)
                v1 = ("" + v1)
                switch fi15 {
                case 762:
                    fmt.Println(v13)
                }
                _ = v18
            }
            switch 2147483647 {
            case 1:
                fmt.Println(fi15)
            }
            v13 -= 3.14159
            _ = fi15
            _ = v16
            _ = v17
        }
        if (!false) {
            var v19 int16 = (v2 / (((-(v2))) | 1))
            v19 -= (-(v2))
            _ = v19
        }
        _ = fi14
    }
    var v20 bool = (!true)
    if v20 {
        var v21 uint = 4294967295
        _ = v21
    } else {
        var v22 string = ""
        var v23 uint = 255
        v24 := GenS1{F0: 42}
        _ = v22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        v26 := GenS2{F0: 127, F1: int(v11)}
        if len(v12) > 0 {
            v12[0] = (fi25 * 4294967295)
        }
        fmt.Println(v20, v3)
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
}
