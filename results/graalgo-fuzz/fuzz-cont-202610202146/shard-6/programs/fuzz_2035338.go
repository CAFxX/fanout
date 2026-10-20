package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int }
type GenS3 struct { F0 int32; F1 uint64; F2 uint32 }
func genFn0(p0 int16) string {
    v7 := GenS1{F0: 0}
    v8 := GenS1{F0: (p0 & p0)}
    _ = p0
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 255
    var v3 bool = (!true)
    v4 := []uint32{0, 0}
    v5 := make(map[uint32]string)
    v6 := make(map[int]bool)
    if v3 {
        for fi9 := 0; fi9 < 9; fi9++ {
            if len(v4) > 0 {
                v4[0] = 782
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = len(v5)
            _ = fi9
        }
        v3 = true
        v2 = (v2 - v2)
    }
    var v10 int32 = 1
    v3 = ("foo bar" > ("a" + "a"))
    if ((3.14159 * 1.5) >= 1e10) {
        var v11 float32 = 0.1
        v2 += (v2 << ((982) & 63))
        _ = v11
    } else {
        switch ("\n\t" + ((("" + "") + (("a" + "\n\t") + ("a" + (("" + (("" + "hello") + "a")) + "hello")))) + "\n\t")) {
        case "\n\t":
            var v12 float32 = 1e10
            _ = v12
        case "":
            v10 -= (v10 % ((v10) | 1))
        case "a":
            var v13 uint8 = 255
            _ = v13
        default:
            fmt.Println(v2, v10, v2)
        }
        var v14 int16 = 213
        _ = v14
    }
    switch ("" + (("" + (("a" + ("\n\t" + ("foo bar" + "foo bar"))) + "hello")) + "a")) {
    case "hello":
        v15 := GenS2{F0: 65535, F1: 4294967295, F2: 0}
        _ = v15
    case "":
        for fi16 := 0; fi16 < 11; fi16++ {
            v2 = (v2 * (^(v2)))
            for fi17 := 0; fi17 < 2; fi17++ {
                v18 := v6[32768]
                v10 = v10
                v2 -= v2
                fmt.Println(fi17, fi17)
                delete(v5, 1)
                var v19 uint32 = 2147483648
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi16
        }
    case "\n\t":
        v20 := GenS2{F0: 65535, F1: 2, F2: 806}
        _ = v20
    }
    var v21 uint8 = 127
    switch -247 {
    case 0:
        v3 = v3
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        fmt.Println(fi22, v3)
        if false {
            var v23 int = (fi22 + (fi22 & fi22))
            v3 = (1.5 > -2.25)
            v24 := GenS2{F0: 420, F1: 947, F2: v23}
            _ = v23
            _ = v24
        } else {
            v25 := GenS3{F0: (v10 & 2147483647), F1: 18446744073709551615, F2: 4294967295}
            fmt.Println(v21)
            v3 = (65535 >= 2)
            _ = v25
        }
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
