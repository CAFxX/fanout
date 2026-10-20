package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 int32 }
func genFn0(p0 uint8) int16 {
    var v4 uint64 = 18446744073709551615
    v5 := []int{32767}
    v4 += 2
    var v6 uint64 = v4
    v7 := GenS1{F0: (v6 & v4), F1: "foo bar"}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return -975
}

func main() {
    var v1 int32 = 1
    var v2 string = "\n\t"
    v3 := make(map[uint32]int)
    var v8 uint = 393
    v1 = v1
    if (2 <= 4294967295) {
        v9 := GenS1{F0: 424, F1: v2}
        _ = v9
    } else {
        delete(v3, 0)
        for fi10 := 0; fi10 < 16; fi10++ {
            if (515 <= 822) {
                if false {
                    v3[1] = (fi10 & (fi10 & (fi10 / ((fi10) | 1))))
                    var v11 uint = v8
                    fmt.Println(v8, v11, v11)
                    _ = v11
                }
                v12 := []int{2147483647}
                _ = v12
            } else {
                fmt.Println(v8)
                for fi13 := 0; fi13 < 7; fi13++ {
                    switch v2 {
                    case "\n\t":
                        v1 -= v1
                    case "foo bar":
                        var v14 string = "hello"
                        _ = v14
                    default:
                        if true {
                            var v15 float32 = (0.1 * 1.5)
                            v16 := GenS2{F0: 32767, F1: 156}
                            if false {
                                if (true || false) {
                                    v8 = 547
                                    switch "a" {
                                    case "\n\t":
                                        v17 := GenS3{F0: 0}
                                        _ = v17
                                    case "hello":
                                        var v18 int64 = 0
                                        _ = v18
                                    case "foo bar":
                                        v19 := GenS1{F0: 125, F1: "hello"}
                                        _ = v19
                                    default:
                                        fmt.Println(fi10, v1)
                                    }
                                }
                            }
                            _ = v15
                            _ = v16
                        }
                    }
                    v20 := []int{}
                    _ = fi13
                    _ = v20
                }
                v21 := []int{fi10, fi10, fi10}
                _ = v21
            }
            _ = fi10
        }
    }
    var v22 string = "hello"
    for fi23 := 0; fi23 < 11; fi23++ {
        fmt.Println(v8, v1, fi23)
        switch "\n\t" {
        case "a":
            switch (("\n\t" + "hello") + "\n\t") {
            case "a":
                var v24 uint64 = 2
                _ = v24
            case "":
                fmt.Println(v8, fi23)
            case "\n\t":
                v8 = (v8 / ((v8) | 1))
            }
        case "hello":
            v25 := GenS2{F0: 32767, F1: 255}
            _ = v25
        case "\n\t":
            v26 := []int{(fi23 * fi23), fi23}
            _ = v26
        default:
            if (!((fi23 != (fi23 >> ((4294967295) & 63))) && false)) {
                if false {
                    var v27 uint = v8
                    v28 := []int{(fi23 >> (((v8 ^ v27)) & 63)), 128}
                    v29 := GenS4{F0: 18446744073709551615, F1: v1}
                    _ = v27
                    _ = v28
                    _ = v29
                } else {
                    var v30 float64 = 0.1
                    _ = v30
                }
                var v31 int64 = 0
                if true {
                    fmt.Println(v8, fi23, v22)
                }
                _ = v31
            } else {
                var v32 int8 = 42
                v8 = v8
                _ = v32
            }
        }
        var v33 int8 = 30
        _ = fi23
        _ = v33
    }
    v2 = v22
    fmt.Println(v8, v2)
    v34 := []int{9223372036854775807}
    fmt.Println(v8, v8, v2)
    switch 4294967295 {
    case -606:
        v35 := GenS4{F0: v8, F1: (v1 - v1)}
        _ = v35
    }
    v2 = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v8
}
