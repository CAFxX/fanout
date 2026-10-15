package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 1
    v3 := []int{42, -1}
    var v4 uint16 = 2
    var v5 int = 2147483647
    var v6 int32 = -365
    switch "" {
    case "a":
        switch ("a" + "a") {
        case "\n\t":
            var v7 int = v5
            _ = v7
        case "hello":
            switch ("hello" + "hello") {
            case "foo bar":
                var v8 uint32 = 719
                _ = v8
            default:
                v9 := GenS1{F0: 1}
                _ = v9
            }
        default:
            v4 -= v4
        }
    case "hello":
        var v10 float32 = -2.25
        _ = v10
    default:
        var v11 uint16 = (v4 & v4)
        _ = v11
    }
    var v12 uint8 = 56
    v13 := GenS1{F0: 9223372036854775807}
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 uint8 = v12
        var v16 uint64 = 264
        _ = fi14
        _ = v15
        _ = v16
    }
    if len(v3) > 0 {
        v3[0] = v5
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        fmt.Println(v12)
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v18 uint8 = v12
            _ = v18
        } else {
            if ("" >= "hello") {
                v4 = 42
                v5 += -824
                v2 += (v2 | 41)
            }
            if (4294967295 > 1) {
                switch v5 {
                case 42:
                    v12 -= v12
                case 9223372036854775807:
                    var v19 int32 = (v6 << (((^(v2))) & 63))
                    _ = v19
                case 32767:
                    fmt.Println(v2, v12, v4)
                }
            } else {
                v5 = (fi17 - (v5 % (((v5 * (^(fi17)))) | 1)))
            }
            v3 = append(v3, 2147483647)
        }
        var v20 int16 = 1000
        _ = fi17
        _ = v20
    }
    if (v6 < (v6 & -1)) {
        var v21 float64 = 1e10
        _ = v21
    } else {
        if (true && (!false)) {
            v2 -= (v2 % ((v2) | 1))
            v22 := GenS2{F0: v2}
            _ = v22
        } else {
            v2 = v2
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 float32 = 0.0
                fmt.Println(v4)
                _ = fi23
                _ = v24
            }
        }
    }
    var v25 int8 = 127
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        if false {
            v4 -= 122
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch "a" {
        case "a":
            for fi27 := 0; fi27 < 1; fi27++ {
                var v28 uint = (v2 + 42)
                _ = fi27
                _ = v28
            }
        case "hello":
            var v29 float32 = 1e10
            _ = v29
        }
        _ = fi26
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v12)
    v12 = (^(v12))
    switch "\n\t" {
    case "":
        var v30 string = (("a" + (("" + ("hello" + "")) + "hello")) + "a")
        _ = v30
    case "\n\t":
        fmt.Println(v5, v5, v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
