package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint }
func genFn0(p0 int64, p1 string) uint32 {
    if true {
        v4 := []int{32768, -798}
        _ = v4
    } else {
        fmt.Println(p0, p0)
        v5 := GenS1{F0: (-128 <= -128), F1: 795, F2: int16(p0)}
        fmt.Println(p0, p0)
        _ = v5
    }
    p0 = 704
    if true {
        switch -2147483648 {
        case 9223372036854775807:
            switch 127 {
            case 2147483647:
                if true {
                    v6 := []int{-9223372036854775808, -80}
                    _ = v6
                }
            }
        case -2147483648:
            v7 := []int{585, -843}
            _ = v7
        case 32767:
            v8 := GenS1{F0: ((true && (255 > 0)) && (!(-25 != int16(p0)))), F1: 42, F2: 1}
            _ = v8
        }
    } else {
        fmt.Println(p1, p1)
        p1 = "a"
        var v9 int64 = (^(p0))
        _ = v9
    }
    p0 += p0
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint32 = 65535
    var v3 uint8 = 0
    v1 -= 1.5
    v1 = float32(0.0)
    if (-329 >= 2147483647) {
        for fi10 := 0; fi10 < 16; fi10++ {
            v1 -= float32(1e10)
            _ = fi10
        }
        var v11 int32 = 760
        _ = v11
    } else {
        v12 := []int{684, 200, -2147483648}
        _ = v12
    }
    v13 := []int{}
    var v14 int = 173
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint16 = 42
            v13 = append(v13, fi15)
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v13 = append(v13, v14)
    v1 = 0.1
    switch "hello" {
    case "foo bar":
        v2 = v2
    }
    v18 := GenS1{F0: (!(127 != -1)), F1: v2, F2: -464}
    var v19 uint8 = (^(v3))
    var v20 float64 = 0.1
    switch v14 {
    case 9223372036854775807:
        if false {
            switch "\n\t" {
            case "\n\t":
                var v21 uint32 = (v2 + (v2 - 65535))
                _ = v21
            default:
                for fi22 := 0; fi22 < 7; fi22++ {
                    if (18446744073709551615 < 1) {
                        var v23 string = (("\n\t" + (("hello" + "\n\t") + "\n\t")) + ("a" + ""))
                        fmt.Println(v1, v2, v3)
                        _ = v23
                    }
                    _ = fi22
                }
            }
        } else {
            for fi24 := 0; fi24 < 17; fi24++ {
                v14 -= (fi24 ^ v14)
                _ = fi24
            }
            var v25 uint32 = v2
            v13 = append(v13, v14)
            _ = v25
        }
    case -925:
        v19 += 1
    case 151:
        if len(v13) > 0 {
            v13[0] = v14
        }
    default:
        fmt.Println(v3)
    }
    v19 = (^(v19))
    if (false && (!(v2 <= 1))) {
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            for fi26 := 0; fi26 < 13; fi26++ {
                v27 := GenS1{F0: true, F1: (v2 << ((284) & 63)), F2: 0}
                fmt.Println(v1, fi26)
                _ = fi26
                _ = v27
            }
        }
    } else {
        v28 := GenS2{F0: 18446744073709551615}
        fmt.Println(v3, v14, v2)
        _ = v28
    }
    var v29 uint16 = 440
    var v30 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
}
