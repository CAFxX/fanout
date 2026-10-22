package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int8; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint16 }
func genFn0(p0 int) string {
    v4 := GenS1{F0: 65535}
    p0 -= 570
    _ = p0
    _ = v4
    return "a"
}

func main() {
    var v1 int64 = -910
    var v2 uint16 = 0
    v3 := []int{604, 32767, -730, 32767}
    switch ("" + ("a" + "hello")) {
    case "\n\t":
        for fi5 := 0; fi5 < 12; fi5++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) > 0 {
                v3[0] = (fi5 ^ (fi5 << ((196) & 63)))
            }
            _ = fi5
        }
    case "hello":
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 string = ("foo bar" + ("foo bar" + ("foo bar" + (("\n\t" + "a") + "a"))))
            if true {
                v8 := GenS2{F0: 4294967295, F1: -1}
                _ = v8
            }
            switch "foo bar" {
            case "foo bar":
                v2 -= 32767
            }
            _ = fi6
            _ = v7
        }
    case "foo bar":
        var v9 float64 = 0.1
        _ = v9
    }
    fmt.Println(v1)
    v10 := GenS3{F0: 1, F1: 42, F2: 0}
    v2 = (v2 ^ (v2 - v2))
    var v11 int16 = 32767
    var v12 int16 = v11
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := GenS2{F0: 2, F1: 1}
        v15 := GenS1{F0: 2}
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS3{F0: -128, F1: (v1 >> ((435) & 63)), F2: 18446744073709551615}
    var v17 float32 = float32((1.5 - 3.14159))
    v3 = append(v3, 4294967295)
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 int64 = v1
        v2 -= 761
        _ = fi18
        _ = v19
    }
    switch "foo bar" {
    case "hello":
        for fi20 := 0; fi20 < 8; fi20++ {
            v2 += (^(v2))
            _ = fi20
        }
    default:
        v21 := GenS4{F0: 2, F1: (-(v1)), F2: v2}
        _ = v21
    }
    v22 := GenS1{F0: 4294967295}
    switch ("\n\t" + "") {
    case "a":
        fmt.Println(v1, v12, v12)
    default:
        fmt.Println(v1, v12)
    }
    var v23 uint = 732
    v24 := GenS4{F0: 193, F1: v1, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
