package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint16; F2 int }
func genFn0(p0 string, p1 int16, p2 uint32) int16 {
    v5 := GenS1{F0: 2147483647}
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (^(p1))
}
func genFn1(p0 int32, p1 int32, p2 string) int32 {
    p2 = (("" + p2) + "\n\t")
    var v6 uint16 = 475
    fmt.Println(p2)
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}

func main() {
    var v1 bool = (1 < 128)
    var v2 uint64 = 2
    v3 := []uint32{65535, 42}
    var v4 uint16 = 42
    switch -1 {
    case 2147483647:
        var v7 int16 = 942
        _ = v7
    default:
        v2 = v2
    }
    _ = len(v3)
    var v8 uint8 = 217
    var v9 float32 = 0.0
    var v10 int = -496
    v1 = v1
    for fi11 := 0; fi11 < 7; fi11++ {
        switch (("" + "hello") + "hello") {
        case "a":
            var v12 string = ""
            _ = v12
        default:
            if ((v8 > (v8 << ((65535) & 63))) && v1) {
                if v1 {
                    _ = len(v3)
                    v4 += v4
                }
            } else {
                v9 -= v9
            }
        }
        _ = fi11
    }
    _ = len(v3)
    switch "foo bar" {
    case "hello":
        if (false || (true || v1)) {
            if ("" < ("foo bar" + "\n\t")) {
                fmt.Println(v8)
                var v13 string = ("hello" + "")
                var v14 int16 = -1
                _ = v13
                _ = v14
            }
            v8 += v8
            var v15 float64 = 1e10
            _ = v15
        } else {
            var v16 int64 = -950
            for fi17 := 0; fi17 < 17; fi17++ {
                switch fi17 {
                case 2147483647:
                    v18 := GenS1{F0: (-(v10))}
                    _ = v18
                default:
                    v2 = v2
                }
                for fi19 := 0; fi19 < 15; fi19++ {
                    fmt.Println(v9, v8)
                    v9 -= v9
                    _ = fi19
                }
                _ = fi17
            }
            _ = v16
        }
    case "":
        if false {
            v9 = ((3.14159 - -2.25) * v9)
        }
    case "foo bar":
        for fi20 := 0; fi20 < 8; fi20++ {
            v3 = append(v3, 205)
            v21 := GenS2{F0: 42, F1: uint16(v8), F2: (v10 + fi20)}
            _ = fi20
            _ = v21
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
