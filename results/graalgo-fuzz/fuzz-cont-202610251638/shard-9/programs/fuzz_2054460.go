package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 bool }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32, p1 int32) string {
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: 0, F1: 1, F2: ((((0.1 * ((-2.25 + (-2.25 * 0.1)) - -2.25)) / 3.14159) + (0.0 - 1.5)) != ((0.1 * 1.5) * 0.0))}
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int64 = -1
            p0 += (p0 * p0)
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint64 = 42
        p0 += (^(p1))
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    return "hello"
}

func main() {
    var v1 uint = 101
    v2 := []string{("\n\t" + (("a" + ("foo bar" + "\n\t")) + ("foo bar" + ("" + "\n\t")))), "a"}
    var v3 int16 = -614
    v4 := make(map[string]string)
    for fi11 := 0; fi11 < 5; fi11++ {
        fmt.Println(v1, v1)
        _ = fi11
    }
    switch "foo bar" {
    case "hello":
        v12 := GenS1{F0: 42, F1: 951, F2: (33 != uint64(v1))}
        _ = v12
    case "\n\t":
        for fi13 := 0; fi13 < 13; fi13++ {
            v1 += (^(v1))
            fmt.Println(v3)
            _ = fi13
        }
    default:
        v1 = (v1 & v1)
    }
    fmt.Println(v1, v3)
    switch 1 {
    case 2147483647:
        var v14 uint64 = 42
        _ = v14
    case 4294967295:
        var v15 uint8 = 245
        _ = v15
    case 762:
        v1 -= 546
    default:
        for fi16 := 0; fi16 < 14; fi16++ {
            _ = len(v2)
            _ = fi16
        }
    }
    v3 = (v3 + 1000)
    var v17 int16 = 32767
    var v18 int16 = (^(v17))
    var v19 uint64 = 4294967295
    if true {
        fmt.Println(v1)
        fmt.Println(v18, v18, v17)
        var v20 string = "hello"
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 int64 = -1
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 int16 = (v18 | 32767)
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        for fi26 := 0; fi26 < 1; fi26++ {
            fmt.Println(fi25, v18)
            v18 = (v3 - v18)
            var v27 int8 = 42
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 uint8 = 112
            switch "a" {
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "\n\t":
                var v30 uint16 = 1
                _ = v30
            default:
                v31 := GenS1{F0: 353, F1: 542, F2: true}
                fmt.Println(fi28, fi25)
                fmt.Println(v18, v18, v3)
                _ = v31
            }
            v32 := GenS2{F0: v17}
            _ = fi28
            _ = v29
            _ = v32
        }
        switch "hello" {
        case "":
            _ = len(v2)
        case "a":
            if (int8(v1) < -23) {
                fmt.Println(v1)
            } else {
                var v33 int32 = 0
                v1 = 153
                _ = v33
            }
        }
        _ = fi25
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
