package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 int; F1 uint; F2 bool }
func genFn0(p0 uint16) uint8 {
    var v6 int8 = 42
    if false {
        v7 := GenS1{F0: 298, F1: -1, F2: (p0 / ((p0) | 1))}
        v8 := GenS2{F0: -2147483648}
        fmt.Println(p0, v6)
        _ = v7
        _ = v8
    }
    v6 += v6
    _ = p0
    _ = v6
    return 125
}
func genFn1(p0 uint32, p1 int, p2 uint64) uint32 {
    p1 += -9223372036854775808
    for fi9 := 0; fi9 < 9; fi9++ {
        if true {
            fmt.Println(p2, p1, p0)
            switch "hello" {
            case "hello":
                fmt.Println(p1)
            case "\n\t":
                if ((true || ((796 != 1000) && true)) && true) {
                    for fi10 := 0; fi10 < 8; fi10++ {
                        switch "" {
                        case "a":
                            switch fi10 {
                            case 1:
                                var v11 uint8 = 127
                                _ = v11
                            case -198:
                                p2 += 2
                            }
                        default:
                            if true {
                                var v12 uint32 = p0
                                p0 -= (v12 - (^(p0)))
                                var v13 uint16 = 790
                                _ = v12
                                _ = v13
                            }
                        }
                        p1 += p1
                        _ = fi10
                    }
                } else {
                    if false {
                        switch (^(fi9)) {
                        case 603:
                            p0 += p0
                        default:
                            v14 := GenS2{F0: 902}
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 10; fi15++ {
                            v16 := []int{(fi9 ^ p1)}
                            var v17 int16 = 42
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        fmt.Println(p2, fi9)
                    } else {
                        if (1000 <= 32767) {
                            var v18 uint64 = 2
                            v19 := GenS3{F0: (p0 / (((p0 - 281)) | 1)), F1: 80}
                            _ = v18
                            _ = v19
                        }
                    }
                    v20 := []int{(fi9 % (((p1 | p1)) | 1)), 9223372036854775807}
                    _ = v20
                }
            }
        }
        _ = fi9
    }
    p0 = 38
    _ = p0
    _ = p1
    _ = p2
    return (p0 / (((p0 ^ p0)) | 1))
}

func main() {
    v1 := []uint32{0, 250}
    v2 := make(map[uint32]uint32)
    var v3 int64 = 42
    var v4 uint64 = 284
    var v5 int8 = 127
    var v21 uint64 = (v4 - (v4 & 4294967295))
    var v22 int = 125
    var v23 string = "hello"
    for fi24 := 0; fi24 < 13; fi24++ {
        if (false || false) {
            for fi25 := 0; fi25 < 17; fi25++ {
                switch (fi25 >> ((2) & 63)) {
                case 42:
                    v26 := GenS4{F0: fi24, F1: 0, F2: true}
                    _ = v26
                default:
                    delete(v2, 811)
                }
                var v27 float32 = (1e10 + (float32(((1e10 * 0.0) - -2.25)) + (-2.25 - -2.25)))
                _ = fi25
                _ = v27
            }
        }
        var v28 bool = false
        _ = fi24
        _ = v28
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        v22 -= fi29
        _ = fi29
    }
    v23 = v23
    var v30 int64 = 9223372036854775807
    fmt.Println(v3, v23, v22)
    var v31 float64 = -2.25
    v32 := GenS1{F0: 0, F1: (-(v3)), F2: 2}
    var v33 uint16 = 2
    switch 9223372036854775807 {
    case 2147483647:
        var v34 int32 = 0
        _ = v34
    case -554:
        var v35 uint8 = 0
        _ = v35
    case 1:
        v3 = 0
    }
    v1 = append(v1, 42)
    var v36 uint64 = 18446744073709551615
    var v37 uint16 = (v33 << ((517) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
