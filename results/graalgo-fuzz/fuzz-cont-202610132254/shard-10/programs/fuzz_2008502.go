package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = false
    var v3 int64 = -297
    var v4 float32 = 1.5
    v5 := []int{}
    v4 += v4
    v6 := GenS1{F0: "\n\t", F1: 861}
    v7 := GenS1{F0: ("\n\t" + "a"), F1: 18446744073709551615}
    for fi8 := 0; fi8 < 2; fi8++ {
        v3 -= v3
        if false {
            var v9 uint = 4294967295
            v10 := v1[(fi8 | (fi8 >> ((v9) & 63)))]
            _ = v10
            _ = v9
        } else {
            var v11 uint64 = 9
            _ = v11
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint64 = 0
        v14 := GenS2{F0: 802, F1: v13}
        delete(v1, (^(fi12)))
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 string = "\n\t"
    if v2 {
        var v16 int32 = -1
        _ = v16
    } else {
        var v17 uint64 = 18446744073709551615
        v18 := v1[128]
        _ = v17
        _ = v18
    }
    _ = len(v1)
    v4 += -2.25
    if false {
        var v19 uint = 4294967295
        _ = v19
    } else {
        fmt.Println(v3, v3, v4)
        fmt.Println(v4, v3)
    }
    fmt.Println(v15, v2, v2)
    var v20 uint32 = 1
    v20 = 1
    fmt.Println(v15)
    var v21 int = -2147483648
    for fi22 := 0; fi22 < 4; fi22++ {
        for fi23 := 0; fi23 < 4; fi23++ {
            v20 = v20
            v15 = v15
            v21 = fi23
            _ = fi23
        }
        switch v15 {
        case "":
            for fi24 := 0; fi24 < 10; fi24++ {
                if (v2 || (true || v2)) {
                    switch ((v15 + ((v15 + v15) + ("foo bar" + v15))) + v15) {
                    case "a":
                        switch v15 {
                        case "a":
                            v21 = (fi22 + (^(fi24)))
                        case "":
                            fmt.Println(fi22, v21, v15)
                        case "hello":
                            if (2147483647 != -1) {
                                fmt.Println(v4, v21)
                                _ = len(v5)
                            } else {
                                if true {
                                    v25 := GenS1{F0: v15, F1: 0}
                                    _ = v25
                                }
                                fmt.Println(v20, v15, v4)
                                var v26 float32 = 1.5
                                _ = v26
                            }
                        default:
                            switch 32767 {
                            case 2147483647:
                                v27 := GenS1{F0: (v15 + v15), F1: uint64(v20)}
                                v4 += v4
                                _ = v27
                            case 127:
                                if (v20 == v20) {
                                    var v28 int = (^(fi24))
                                    fmt.Println(v20)
                                    _ = v28
                                }
                            }
                        }
                    case "foo bar":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        fmt.Println(fi24, fi22)
                    }
                } else {
                    if v2 {
                        var v29 bool = (42 == 1)
                        v2 = v2
                        _ = v29
                    }
                    var v30 uint32 = v20
                    v3 = v3
                    _ = v30
                }
                _ = fi24
            }
        case "\n\t":
            var v31 uint64 = 355
            _ = v31
        case "a":
            var v32 int = v21
            _ = v32
        }
        _ = fi22
    }
    if v2 {
        var v33 float64 = (float64((((v4 * ((v4 - 3.14159) - ((0.1 - v4) * v4))) + 3.14159) / v4)) + 0.1)
        v34 := GenS3{F0: 127}
        fmt.Println(v2)
        _ = v33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
