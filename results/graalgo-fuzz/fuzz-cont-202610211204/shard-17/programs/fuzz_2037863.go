package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int64) uint64 {
    var v6 int64 = 1
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 float64 = 1.5
        v9 := GenS1{F0: 643}
        v10 := []int{fi7, 0}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint8 = 16
    v12 := []int{-1, -896, -1}
    v13 := GenS1{F0: 18446744073709551615}
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    _ = v6
    return 0
}
func genFn1() uint {
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := []int{(fi14 | (fi15 * (fi14 << ((48) & 63)))), fi15, fi15}
            _ = len(v16)
            if len(v16) > 0 {
                v16[0] = (fi15 + (fi14 + 32768))
            }
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            var v17 int64 = 19
            v17 -= 9223372036854775807
            if (127 == 0) {
                v17 -= (v17 & (v17 << ((18446744073709551615) & 63)))
                v17 += v17
            }
            fmt.Println(v17, fi14, fi15)
            v18 := GenS1{F0: 305}
            v17 = -9223372036854775808
            switch "foo bar" {
            case "foo bar":
                var v19 float64 = 3.14159
                _ = v19
            case "\n\t":
                v20 := GenS1{F0: 4294967295}
                _ = v20
            default:
                var v21 uint16 = 42
                _ = v21
            }
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi14
    }
    v22 := []int{42, 9223372036854775807}
    if len(v22) > 0 {
        v22[0] = 0
    }
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    var v23 float64 = 0.1
    _ = v22
    _ = v23
    return 866
}

func main() {
    var v1 int16 = 1000
    var v2 uint32 = 1
    var v3 uint64 = 883
    var v4 bool = (!false)
    v5 := make(map[string]bool)
    v24 := []int{}
    var v25 uint16 = 0
    fmt.Println(v25, v1, v2)
    for fi26 := 0; fi26 < 19; fi26++ {
        var v27 bool = v4
        _ = len(v24)
        _ = fi26
        _ = v27
    }
    switch ("" + "foo bar") {
    case "a":
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
    case "hello":
        var v28 float64 = (0.1 - 1e10)
        _ = v28
    case "":
        var v29 int64 = 1
        _ = v29
    }
    _ = len(v5)
    if (v2 <= uint32(v3)) {
        fmt.Println(v25)
        v25 -= v25
    } else {
        fmt.Println(v3, v3)
        var v30 float64 = 0.1
        var v31 uint = 65535
        _ = v30
        _ = v31
    }
    var v32 uint32 = v2
    for fi33 := 0; fi33 < 5; fi33++ {
        if (!false) {
            fmt.Println(v25, fi33)
            if len(v24) > 0 {
                v24[0] = 42
            }
        } else {
            for fi34 := 0; fi34 < 13; fi34++ {
                var v35 bool = ((("a" + "hello") == "foo bar") && (!(v4 && v4)))
                _ = fi34
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 11; fi36++ {
            v25 -= 255
            _ = fi36
        }
        _ = fi33
    }
    var v37 int = 2147483647
    switch v37 {
    case 2147483647:
        v38 := v5["foo bar"]
        _ = v38
    case -991:
        var v39 bool = v4
        _ = v39
    case -578:
        fmt.Println(v37)
    default:
        v25 -= 255
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v5
}
