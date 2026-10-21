package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 int16) int32 {
    for fi6 := 0; fi6 < 9; fi6++ {
        p0 -= (p0 | 42)
        var v7 uint32 = 65535
        p0 -= p0
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int64 = 1
        switch "\n\t" {
        case "\n\t":
            v9 = v9
        default:
            v10 := GenS1{F0: 0, F1: (true && false), F2: 42}
            _ = v10
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 bool = (!true)
            if false {
                var v13 float32 = -2.25
                _ = v13
            }
            _ = fi11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(p0)
    var v14 int8 = 0
    _ = p0
    _ = v14
    return 2147483647
}
func genFn1(p0 uint) int16 {
    v15 := GenS1{F0: 157, F1: (!true), F2: 255}
    switch ("hello" + "hello") {
    case "hello":
        fmt.Println(p0)
    case "a":
        var v16 int16 = 249
        _ = v16
    default:
        if false {
            for fi17 := 0; fi17 < 11; fi17++ {
                p0 = 255
                fmt.Println(fi17)
                v18 := GenS1{F0: 6, F1: ((0 <= uint16(p0)) || (!(!false))), F2: 994}
                _ = fi17
                _ = v18
            }
            var v19 uint16 = 1
            _ = v19
        }
    }
    _ = p0
    _ = v15
    return 1
}

func main() {
    v1 := []string{"\n\t"}
    var v2 int64 = -1
    var v3 int32 = -2147483648
    var v4 string = ("hello" + "")
    var v5 uint32 = 750
    for fi20 := 0; fi20 < 7; fi20++ {
        fmt.Println(v5, v5, v3)
        v21 := GenS2{F0: true}
        v22 := GenS2{F0: true}
        _ = fi20
        _ = v21
        _ = v22
    }
    v1 = append(v1, ("" + "\n\t"))
    fmt.Println(v5, v4, v2)
    fmt.Println(v2, v3)
    if false {
        v2 = 9223372036854775807
    }
    switch -2147483648 {
    case 336:
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 uint = uint(v5)
            _ = fi23
            _ = v24
        }
    case 9223372036854775807:
        v5 = v5
    default:
        _ = len(v1)
    }
    if (false && false) {
        v1 = append(v1, v4)
        var v25 int = 42
        _ = v25
    } else {
        v5 = v5
        for fi26 := 0; fi26 < 19; fi26++ {
            for fi27 := 0; fi27 < 15; fi27++ {
                var v28 float32 = -2.25
                _ = fi27
                _ = v28
            }
            if len(v1) > 0 {
                v1[0] = ("a" + "\n\t")
            }
            fmt.Println(v5, fi26)
            _ = fi26
        }
    }
    if (true || (1e10 < (((3.14159 / 3.14159) - -2.25) * 1.5))) {
        v29 := GenS1{F0: 0, F1: (0 < 1), F2: 902}
        _ = v29
    } else {
        v4 = v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
