package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int32; F1 bool }
func genFn0(p0 uint8, p1 int64) uint64 {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int32 = -2147483648
        v7 := GenS1{F0: 127, F1: 42}
        var v8 int8 = 100
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 uint16 = 150
    v10 := []int{32767, 594}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 4294967295
}
func genFn1(p0 uint32, p1 uint16, p2 bool) uint16 {
    v11 := GenS1{F0: -1, F1: p0}
    fmt.Println(p2)
    switch 0 {
    case -669:
        if (-2.25 == 0.0) {
            v12 := GenS2{F0: 42, F1: 611}
            _ = v12
        }
    case 32767:
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 string = ("a" + "")
            _ = fi13
            _ = v14
        }
    case 1:
        if true {
            fmt.Println(p2, p1)
            var v15 string = ("foo bar" + "foo bar")
            if true {
                p2 = false
                var v16 int16 = 1000
                v17 := []int{-111, 754, 362, -2147483648}
                _ = v16
                _ = v17
            } else {
                var v18 bool = p2
                if ("" > (v15 + (v15 + v15))) {
                    for fi19 := 0; fi19 < 11; fi19++ {
                        v20 := GenS2{F0: 2147483648, F1: 255}
                        v21 := GenS3{F0: -2147483648, F1: (-2.25 < 1e10)}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    fmt.Println(p1)
                    var v22 float64 = 1e10
                    _ = v22
                } else {
                    fmt.Println(p0, v15)
                    var v23 float32 = (-2.25 * 0.0)
                    _ = v23
                }
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 float32 = 1e10
                    if (1 >= -1) {
                        fmt.Println(v25, fi24)
                        for fi26 := 0; fi26 < 20; fi26++ {
                            if true {
                                var v27 int = 127
                                var v28 uint16 = (p1 / (((p1 ^ p1)) | 1))
                                _ = v27
                                _ = v28
                            }
                            v29 := []int{2147483647, (fi24 + (fi26 % (((fi26 | fi24)) | 1))), (fi24 + (fi24 ^ fi24))}
                            var v30 uint64 = 1
                            _ = fi26
                            _ = v29
                            _ = v30
                        }
                        fmt.Println(v15, v15, p0)
                    }
                    v31 := []int{(fi24 << ((759) & 63)), (fi24 / (((fi24 & fi24)) | 1))}
                    _ = fi24
                    _ = v25
                    _ = v31
                }
                _ = v18
            }
            _ = v15
        }
    default:
        v32 := []int{-1, 9223372036854775807, -9223372036854775808}
        _ = v32
    }
    switch ("\n\t" + "\n\t") {
    case "hello":
        var v33 int8 = 0
        _ = v33
    case "a":
        for fi34 := 0; fi34 < 9; fi34++ {
            var v35 int16 = -32768
            _ = fi34
            _ = v35
        }
    default:
        v36 := []int{-1, 155, -9223372036854775808}
        _ = v36
    }
    var v37 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v37
    return (p1 | (p1 << ((65535) & 63)))
}

func main() {
    var v1 uint = 2
    var v2 int = -9223372036854775808
    var v3 int8 = 42
    var v4 uint32 = 2147483648
    fmt.Println(v1)
    fmt.Println(v2, v1)
    var v38 int8 = -128
    switch "" {
    case "foo bar":
        v2 -= (v2 ^ v2)
    case "\n\t":
        var v39 uint = (^(v1))
        _ = v39
    case "hello":
        v40 := []int{v2, 2147483647, v2, (^(v2))}
        _ = v40
    default:
        for fi41 := 0; fi41 < 17; fi41++ {
            fmt.Println(fi41, v3, v38)
            var v42 string = "a"
            _ = fi41
            _ = v42
        }
    }
    var v43 float64 = 0.0
    v44 := []int{v2}
    v3 -= v3
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v43
    _ = v44
}
