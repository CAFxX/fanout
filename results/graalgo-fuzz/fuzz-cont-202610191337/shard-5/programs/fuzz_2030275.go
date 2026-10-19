package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint }
type GenS3 struct { F0 int8; F1 bool; F2 uint8 }
func genFn0(p0 int8, p1 uint16) int8 {
    v6 := GenS1{F0: -9223372036854775808, F1: "foo bar", F2: 0}
    var v7 string = "foo bar"
    if (true || (1 <= -1)) {
        v8 := GenS2{F0: (true && (546 != 2147483647)), F1: p1, F2: 255}
        _ = v8
    } else {
        switch 32768 {
        case 0:
            p0 += (p0 - p0)
        default:
            if true {
                for fi9 := 0; fi9 < 17; fi9++ {
                    if (18446744073709551615 != 2) {
                        if true {
                            for fi10 := 0; fi10 < 8; fi10++ {
                                fmt.Println(fi10, p1)
                                _ = fi10
                            }
                            v7 = "foo bar"
                        } else {
                            var v11 int64 = -213
                            _ = v11
                        }
                        v12 := []int{fi9, (fi9 ^ 1), fi9, fi9}
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 5; fi13++ {
                            var v14 int16 = 1000
                            var v15 string = ("\n\t" + v7)
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        v16 := GenS3{F0: (p0 << ((1) & 63)), F1: true, F2: 1}
                        _ = v16
                    }
                    v17 := []int{fi9, fi9, (fi9 >> ((695) & 63))}
                    _ = fi9
                    _ = v17
                }
                var v18 int16 = -853
                _ = v18
            }
        }
    }
    var v19 uint32 = 399
    if true {
        for fi20 := 0; fi20 < 8; fi20++ {
            if true {
                switch (v7 + ("hello" + "a")) {
                case "a":
                    var v21 float64 = (float64(float32(1e10)) * float64(0.0))
                    _ = v21
                case "\n\t":
                    switch "\n\t" {
                    case "\n\t":
                        var v22 int = fi20
                        _ = v22
                    case "hello":
                        for fi23 := 0; fi23 < 10; fi23++ {
                            fmt.Println(fi23)
                            v24 := []int{154, (fi20 >> ((48) & 63))}
                            _ = fi23
                            _ = v24
                        }
                    }
                case "":
                    var v25 uint8 = 127
                    _ = v25
                default:
                    var v26 uint32 = 720
                    _ = v26
                }
                var v27 int8 = p0
                _ = v27
            }
            fmt.Println(p0, v19, v19)
            _ = fi20
        }
    } else {
        p0 += (p0 ^ p0)
        var v28 uint = 2
        _ = v28
    }
    _ = p0
    _ = p1
    _ = v19
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 bool, p1 bool, p2 int32) int {
    p1 = p1
    v29 := []int{127, -2147483648, 4294967295}
    var v30 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v30
    return int(p2)
}

func main() {
    v1 := []uint32{0, 1, 160, 2147483648}
    var v2 bool = true
    var v3 uint8 = 1
    var v4 string = "foo bar"
    var v5 uint8 = v3
    fmt.Println(v2)
    var v31 float64 = 0.0
    switch "foo bar" {
    case "hello":
        switch v4 {
        case "a":
            v31 += float64((1.5 + 1.5))
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    if ((-2.25 * ((float64(0.0) + (((v31 / v31) / (0.1 / v31)) / float64((float32(float64(1.5)) * 0.1)))) + ((1e10 * ((v31 / float64(1.5)) * v31)) * 0.0))) <= 3.14159) {
        for fi32 := 0; fi32 < 4; fi32++ {
            v5 = (v5 - v3)
            v1 = append(v1, 42)
            v5 -= (v5 + (^(v5)))
            _ = fi32
        }
        var v33 int32 = 42
        _ = v33
    }
    var v34 int32 = 1
    if (!(genFn0(100, 164) <= 127)) {
        var v35 uint16 = 65535
        _ = v35
    }
    v3 = v5
    for fi36 := 0; fi36 < 10; fi36++ {
        v5 += (v3 % ((v3) | 1))
        var v37 bool = v2
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        var v39 bool = true
        var v40 uint32 = 65535
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
