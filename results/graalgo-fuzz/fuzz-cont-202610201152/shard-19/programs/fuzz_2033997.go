package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int8; F1 int; F2 uint8 }
type GenS3 struct { F0 bool; F1 int64; F2 uint }

func main() {
    var v1 int8 = 100
    v2 := []string{("\n\t" + "a")}
    v3 := make(map[int]int)
    v4 := []string{("a" + "\n\t"), "\n\t", "a", ""}
    v5 := make(map[int]uint32)
    if true {
        fmt.Println(v1)
    }
    var v6 uint64 = 18446744073709551615
    if (603 < 2) {
        v7 := GenS1{F0: v1, F1: -2147483648}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 3; fi8++ {
            v1 = (-(v1))
            v3[-9223372036854775808] = fi8
            v1 = 73
            _ = fi8
        }
        v6 = v6
        delete(v5, 925)
    }
    var v9 int16 = 1
    if false {
        v9 = 42
        if ((-128 <= int8(v9)) || ((false && true) && false)) {
            v10 := GenS2{F0: (-(v1)), F1: -2147483648, F2: 255}
            _ = v10
        } else {
            if false {
                var v11 float64 = (-2.25 + 0.0)
                _ = v11
            }
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 uint32 = uint32(v6)
            var v14 int64 = 157
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        var v15 uint8 = 42
        switch (("hello" + ("" + ("\n\t" + ""))) + "") {
        case "\n\t":
            var v16 bool = true
            _ = v16
        default:
            var v17 int64 = -9223372036854775808
            _ = v17
        }
        _ = v15
    }
    v4 = append(v4, "hello")
    fmt.Println(v1, v6, v9)
    if (1 == 42) {
        v6 = 174
        if (v6 < (v6 ^ (v6 % ((800) | 1)))) {
            fmt.Println(v6, v6, v6)
            var v18 uint32 = 2147483648
            _ = v18
        } else {
            v1 += v1
            fmt.Println(v9, v1, v6)
            var v19 uint64 = (v6 ^ v6)
            _ = v19
        }
        var v20 int = 128
        _ = v20
    }
    _ = len(v5)
    if (1 == 486) {
        var v21 int = -878
        v21 -= -2147483648
        _ = v21
    } else {
        if true {
            v3[127] = 1
            v22 := GenS3{F0: (204 <= 2), F1: 995, F2: 4294967295}
            if false {
                v6 -= v6
                for fi23 := 0; fi23 < 10; fi23++ {
                    switch fi23 {
                    case -1:
                        for fi24 := 0; fi24 < 14; fi24++ {
                            v6 -= 2
                            var v25 int32 = -2147483648
                            v9 = 1000
                            if (v25 == 1) {
                                _ = len(v4)
                                switch 127 {
                                case 462:
                                    var v26 uint8 = 2
                                    _ = v26
                                case -9223372036854775808:
                                    v9 += (v9 & v9)
                                }
                                var v27 bool = true
                                _ = v27
                            }
                            _ = fi24
                            _ = v25
                        }
                    case 4294967295:
                        for fi28 := 0; fi28 < 12; fi28++ {
                            var v29 int32 = int32(v9)
                            var v30 uint = 18446744073709551615
                            v31 := GenS3{F0: ((v30 | v30) < 323), F1: 663, F2: v30}
                            var v32 float32 = 1e10
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    }
                    _ = len(v2)
                    var v33 string = "a"
                    _ = fi23
                    _ = v33
                }
            }
            _ = v22
        }
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
