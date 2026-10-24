package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int; F2 uint64 }
type GenS3 struct { F0 string; F1 int32; F2 int8 }
type GenS4 struct { F0 uint; F1 uint; F2 uint16 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS7 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint8, p1 uint32, p2 uint8) int32 {
    v4 := GenS1{F0: p1}
    v5 := GenS2{F0: 18446744073709551615, F1: 694, F2: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -2147483648
}
func genFn1(p0 int16) uint16 {
    p0 = p0
    v6 := GenS3{F0: ("a" + "a"), F1: 1, F2: -128}
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 uint = 0
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    delete(v3, "")
    if (0 < 0) {
        var v7 bool = true
        _ = v7
    } else {
        v1 += (v1 ^ v1)
        var v8 string = "foo bar"
        switch v8 {
        case "hello":
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS4{F0: v1, F1: v1, F2: genFn1(42)}
                for fi11 := 0; fi11 < 5; fi11++ {
                    switch v8 {
                    case "hello":
                        if false {
                            var v12 uint64 = 18446744073709551615
                            var v13 float32 = -2.25
                            _ = v12
                            _ = v13
                        }
                    case "":
                        var v14 uint64 = 18446744073709551615
                        _ = v14
                    }
                    var v15 float64 = 1.5
                    _ = fi11
                    _ = v15
                }
                switch v8 {
                case "":
                    var v16 int8 = 98
                    _ = v16
                case "\n\t":
                    if (!true) {
                        v1 -= v1
                        v17 := GenS1{F0: 2147483648}
                        var v18 int16 = 32767
                        _ = v17
                        _ = v18
                    }
                case "a":
                    for fi19 := 0; fi19 < 16; fi19++ {
                        delete(v3, "\n\t")
                        _ = fi19
                    }
                }
                _ = fi9
                _ = v10
            }
        case "a":
            v20 := GenS5{F0: 32768, F1: (v8 + ("\n\t" + (v8 + v8)))}
            _ = v20
        default:
            v8 = (v8 + "\n\t")
        }
        _ = v8
    }
    v21 := GenS3{F0: "", F1: 1, F2: -116}
    _ = len(v3)
    var v22 uint32 = 127
    v23 := GenS6{F0: v22, F1: 2147483648, F2: (^(v22))}
    var v24 uint16 = genFn1(42)
    var v25 int16 = -570
    var v26 uint64 = 492
    if (true && false) {
        var v27 string = (("\n\t" + "hello") + "\n\t")
        fmt.Println(v26, v26, v24)
        v3[v27] = v27
        _ = v27
    }
    switch (("a" + ("a" + "hello")) + ("a" + "hello")) {
    case "hello":
        var v28 uint32 = v22
        _ = v28
    default:
        v29 := GenS5{F0: -975, F1: ""}
        _ = v29
    }
    var v30 uint8 = 128
    v31 := GenS7{F0: uint(v24), F1: 32767}
    var v32 uint32 = v22
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
