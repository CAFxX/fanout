package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint16; F2 uint64 }
func genFn0(p0 int8) bool {
    p0 = p0
    var v6 float64 = (3.14159 * 1e10)
    var v7 uint8 = 0
    _ = p0
    _ = v6
    _ = v7
    return (53 <= 471)
}
func genFn1(p0 uint8) string {
    var v8 uint = 917
    v9 := GenS1{F0: (v8 & (^(v8))), F1: 0}
    v8 += 2
    p0 -= (p0 & p0)
    var v10 uint8 = p0
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return ""
}

func main() {
    var v1 float64 = 1.5
    var v2 int16 = 32767
    v3 := []string{"hello"}
    var v4 int8 = -128
    var v5 int64 = 0
    fmt.Println(v4, v1, v5)
    for fi11 := 0; fi11 < 9; fi11++ {
        v2 -= (v2 >> ((2) & 63))
        var v12 uint64 = 1
        _ = fi11
        _ = v12
    }
    var v13 uint16 = 42
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS1{F0: 18446744073709551615, F1: 60}
            _ = fi15
            _ = v16
        }
        fmt.Println(fi14)
        if len(v3) > 0 {
            v3[0] = ("\n\t" + ("a" + ("\n\t" + "a")))
        }
        _ = fi14
    }
    var v17 int8 = -104
    for fi18 := 0; fi18 < 18; fi18++ {
        v1 += v1
        v2 = v2
        if true {
            for fi19 := 0; fi19 < 15; fi19++ {
                v1 = (1.5 - (0.0 - ((v1 / (v1 + (v1 - 1.5))) / v1)))
                _ = len(v3)
                _ = fi19
            }
            v4 -= (v17 + (v4 << ((65535) & 63)))
            v3 = append(v3, "hello")
        }
        _ = fi18
    }
    var v20 int8 = (v4 << ((632) & 63))
    var v21 float64 = 1e10
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 float64 = v1
        _ = fi22
        _ = v23
    }
    fmt.Println(v1, v13)
    var v24 uint32 = 101
    var v25 int64 = v5
    var v26 int8 = v4
    for fi27 := 0; fi27 < 18; fi27++ {
        v28 := GenS2{F0: 4294967295}
        _ = fi27
        _ = v28
    }
    switch -2147483648 {
    case 127:
        if true {
            _ = len(v3)
        } else {
            v29 := GenS3{F0: (v2 >> ((725) & 63)), F1: v13, F2: 224}
            _ = v29
        }
    }
    v17 = int8(v13)
    var v30 int32 = -195
    switch "a" {
    case "foo bar":
        v30 = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
