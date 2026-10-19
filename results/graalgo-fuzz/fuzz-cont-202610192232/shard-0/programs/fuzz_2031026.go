package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
func genFn0(p0 bool, p1 int64, p2 uint8) int32 {
    for fi6 := 0; fi6 < 2; fi6++ {
        p1 = p1
        if false {
            var v7 uint = 360
            if (p1 < p1) {
                var v8 int8 = -128
                switch ((("a" + "\n\t") + "") + "hello") {
                case "foo bar":
                    for fi9 := 0; fi9 < 13; fi9++ {
                        v10 := GenS1{F0: 893, F1: (^(p2))}
                        _ = fi9
                        _ = v10
                    }
                case "a":
                    if (0 == 42) {
                        fmt.Println(p2, p1, v8)
                        for fi11 := 0; fi11 < 16; fi11++ {
                            v12 := GenS1{F0: uint32(p2), F1: (p2 - (p2 * p2))}
                            if p0 {
                                switch (fi6 << ((4294967295) & 63)) {
                                case 796:
                                    var v13 float32 = float32(0.1)
                                    var v14 uint32 = 65535
                                    _ = v13
                                    _ = v14
                                case 32768:
                                    p0 = p0
                                }
                                v15 := []int{(fi6 << ((v7) & 63))}
                                p1 = p1
                                _ = v15
                            }
                            fmt.Println(p2, p0)
                            _ = fi11
                            _ = v12
                        }
                        var v16 int16 = -32768
                        _ = v16
                    } else {
                        fmt.Println(v8)
                        v7 += 344
                    }
                }
                _ = v8
            }
            var v17 int32 = -255
            _ = v17
            _ = v7
        } else {
            fmt.Println(fi6, p1, fi6)
        }
        var v18 uint8 = 225
        _ = fi6
        _ = v18
    }
    switch "a" {
    case "hello":
        if (!(p0 && (p0 || (216 != p2)))) {
            var v19 uint64 = 18446744073709551615
            _ = v19
        }
    case "a":
        var v20 uint64 = 4294967295
        _ = v20
    }
    v21 := []int{-741}
    var v22 int8 = -126
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return -1
}
func genFn1(p0 int16) uint32 {
    fmt.Println(p0)
    var v23 string = "foo bar"
    v24 := GenS1{F0: 0, F1: 42}
    var v25 uint16 = 32767
    _ = p0
    _ = v23
    _ = v24
    _ = v25
    return 65535
}

func main() {
    v1 := []int{344}
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]bool)
    var v26 uint16 = 255
    v26 += (v26 - (v26 * v26))
    v3["foo bar"] = 269
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 int16 = 13
        v29 := GenS1{F0: 2147483648, F1: 196}
        switch ("\n\t" + "a") {
        case "":
            var v30 uint8 = 127
            _ = v30
        default:
            var v31 uint32 = 2147483647
            _ = v31
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        _ = len(v3)
        _ = fi32
    }
    v4[genFn1(-945)] = 1
    for fi33 := 0; fi33 < 14; fi33++ {
        fmt.Println(v26, v26)
        _ = fi33
    }
    v26 = (v26 % (((^(v26))) | 1))
    var v34 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
