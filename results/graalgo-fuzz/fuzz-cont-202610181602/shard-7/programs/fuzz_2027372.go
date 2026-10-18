package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 bool }
func genFn0(p0 int16) uint8 {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint = 249
        v9 := []int{fi7, 128, 1, (fi7 - fi7)}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        p0 -= p0
        switch "\n\t" {
        case "foo bar":
            if true {
                for fi11 := 0; fi11 < 5; fi11++ {
                    v12 := GenS1{F0: 42}
                    v13 := GenS2{F0: 171, F1: 127, F2: (!false)}
                    fmt.Println(p0)
                    var v14 uint64 = 1
                    var v15 uint64 = (v14 * 42)
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                fmt.Println(p0)
            }
        case "a":
            var v16 uint16 = 42
            _ = v16
        case "\n\t":
            fmt.Println(fi10, fi10)
        default:
            v17 := []int{fi10}
            _ = v17
        }
        v18 := GenS2{F0: 2, F1: 127, F2: true}
        _ = fi10
        _ = v18
    }
    _ = p0
    return 1
}
func genFn1() uint64 {
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 int8 = 122
        v21 := GenS1{F0: int8(fi19)}
        v20 += (v20 << ((71) & 63))
        _ = fi19
        _ = v20
        _ = v21
    }
    switch "foo bar" {
    case "":
        v22 := []int{127, -207}
        _ = v22
    default:
        switch ("" + (("hello" + "hello") + "hello")) {
        case "\n\t":
            var v23 uint8 = 127
            _ = v23
        }
    }
    v24 := []int{-1, 1}
    var v25 uint8 = 127
    _ = v24
    _ = v25
    return 54
}

func main() {
    var v1 bool = false
    var v2 uint16 = 205
    var v3 int16 = 1
    v4 := []int{-9223372036854775808}
    var v5 uint16 = 255
    var v6 float32 = 1e10
    for fi26 := 0; fi26 < 20; fi26++ {
        var v27 uint16 = 763
        if len(v4) > 0 {
            v4[0] = fi26
        }
        _ = fi26
        _ = v27
    }
    _ = len(v4)
    var v28 uint64 = 18446744073709551615
    v29 := GenS2{F0: genFn0(v3), F1: genFn0((v3 | v3)), F2: true}
    if ((true && ((^(v28)) != (v28 * genFn1()))) || (4294967295 < 4294967295)) {
        var v30 string = ""
        var v31 bool = ((10 <= 8) && (!(0 != -2147483648)))
        v4 = append(v4, 0)
        _ = v30
        _ = v31
    } else {
        switch "a" {
        case "foo bar":
            if v1 {
                for fi32 := 0; fi32 < 2; fi32++ {
                    switch "hello" {
                    case "foo bar":
                        for fi33 := 0; fi33 < 3; fi33++ {
                            var v34 int32 = 0
                            _ = fi33
                            _ = v34
                        }
                    case "hello":
                        v28 -= 779
                    }
                    if (152 > -1) {
                        fmt.Println(v5, fi32, v2)
                    }
                    _ = fi32
                }
                var v35 bool = true
                v36 := GenS1{F0: 42}
                _ = v35
                _ = v36
            }
        }
    }
    if (104 < v5) {
        if len(v4) > 0 {
            v4[0] = 283
        }
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
    }
    v6 += (1e10 - v6)
    fmt.Println(v5)
    fmt.Println(v28, v3)
    var v37 int64 = -91
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
