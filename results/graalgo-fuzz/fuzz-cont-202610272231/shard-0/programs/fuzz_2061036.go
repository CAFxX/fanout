package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 bool }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int16, p1 int32, p2 int) int32 {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint64 = 18446744073709551615
        v6 -= 4294967295
        _ = fi5
        _ = v6
    }
    p0 -= 1000
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: 2, F1: -1, F2: ("foo bar" > "")}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint = 0
    var v2 int64 = 9223372036854775807
    var v3 float64 = float64(0.1)
    var v4 uint8 = 80
    v9 := []int{0, -1}
    var v10 bool = (v2 <= v2)
    var v11 int8 = -1
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v12 int16 = 1000
    switch "foo bar" {
    case "a":
        if len(v9) > 0 {
            v9[0] = 443
        }
    case "hello":
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS2{F0: 4294967295, F1: 0, F2: 4294967295}
            _ = fi13
            _ = v14
        }
    }
    var v15 int = 9223372036854775807
    var v16 int64 = 0
    v10 = ((v4 - v4) != 128)
    if v10 {
        for fi17 := 0; fi17 < 16; fi17++ {
            v15 += fi17
            _ = fi17
        }
    }
    v18 := GenS3{F0: v10, F1: (v15 << ((v1) & 63))}
    if (v11 < v11) {
        if true {
            switch ("a" + "a") {
            case "foo bar":
                fmt.Println(v15, v2)
            default:
                fmt.Println(v2, v15)
            }
            v19 := GenS1{F0: 4294967295, F1: (v2 << (((v1 >> ((v1) & 63))) & 63)), F2: (v10 || (1.5 > 0.0))}
            v20 := GenS4{F0: v4}
            _ = v19
            _ = v20
        }
    }
    if true {
        for fi21 := 0; fi21 < 15; fi21++ {
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 int = (-(fi21))
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    }
    switch ((((("" + ("a" + "\n\t")) + (("a" + "hello") + (("hello" + ("\n\t" + "foo bar")) + ("\n\t" + "foo bar")))) + "a") + "\n\t") + "a") {
    case "\n\t":
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    default:
        for fi24 := 0; fi24 < 13; fi24++ {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            v15 -= 32768
            fmt.Println(v1, v2, fi24)
            _ = fi24
        }
    }
    fmt.Println(v15, v2)
    v25 := GenS4{F0: 127}
    for fi26 := 0; fi26 < 1; fi26++ {
        v2 -= v2
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v9
}
