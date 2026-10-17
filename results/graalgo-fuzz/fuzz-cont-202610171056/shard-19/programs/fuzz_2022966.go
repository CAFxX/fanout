package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int32; F1 int; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int; F1 int64; F2 uint }
type GenS6 struct { F0 string }
func genFn0(p0 bool, p1 uint64) bool {
    if false {
        v6 := []int{}
        var v7 bool = p0
        p0 = p0
        _ = v6
        _ = v7
    }
    var v8 int64 = 240
    v9 := GenS1{F0: 32, F1: 1, F2: (p1 - (p1 + (p1 - 509)))}
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            switch "a" {
            case "\n\t":
                var v12 uint = uint(p1)
                _ = v12
            case "foo bar":
                v13 := GenS2{F0: -277, F1: 128, F2: 42}
                _ = v13
            }
            switch (fi11 | fi11) {
            case -9223372036854775808:
                if true {
                    var v14 string = (("hello" + "\n\t") + "foo bar")
                    _ = v14
                }
            case -2147483648:
                var v15 uint = 42
                if p0 {
                    var v16 uint16 = 925
                    v16 -= (v16 / ((691) | 1))
                    var v17 string = ("" + "\n\t")
                    var v18 float64 = ((float64(0.0) * 0.0) * (1.5 / 1.0))
                    _ = v16
                    _ = v17
                    _ = v18
                }
                if false {
                    v19 := []int{fi10, 42, fi11, 569}
                    _ = v19
                }
                switch (-(fi11)) {
                case 840:
                    var v20 uint32 = 882
                    _ = v20
                case 806:
                    var v21 int = fi10
                    _ = v21
                case -2147483648:
                    p0 = true
                default:
                    switch int(v8) {
                    case -2147483648:
                        switch fi11 {
                        case 0:
                            v8 -= (v8 + v8)
                            p1 = (p1 - (p1 >> (((^(v15))) & 63)))
                        case 32768:
                            v22 := GenS3{F0: 2}
                            _ = v22
                        default:
                            if (0 <= uint8(v15)) {
                                v23 := GenS4{F0: -128, F1: v8}
                                var v24 uint32 = 42
                                switch ("\n\t" + "") {
                                case "foo bar":
                                    var v25 int = 708
                                    _ = v25
                                case "hello":
                                    var v26 uint8 = 1
                                    v27 := GenS4{F0: 0, F1: (v8 << ((v15) & 63))}
                                    _ = v26
                                    _ = v27
                                case "a":
                                    if false {
                                        v28 := []int{}
                                        switch int(p1) {
                                        case 1:
                                            _ = len(v28)
                                        case 42:
                                            fmt.Println(v24)
                                        case 555:
                                            p1 = (p1 - p1)
                                        }
                                        if (42 < 43) {
                                            fmt.Println(v24)
                                            fmt.Println(fi10, p0)
                                        }
                                        _ = v28
                                    } else {
                                        var v29 float64 = float64(0.1)
                                        _ = v29
                                    }
                                default:
                                    p1 += p1
                                }
                                _ = v23
                                _ = v24
                            }
                        }
                    case -1:
                        v30 := []int{(^(fi11)), 0, (fi10 - (fi11 | -9223372036854775808))}
                        _ = v30
                    }
                }
                v8 -= v8
                _ = v15
            }
            _ = fi11
        }
        _ = fi10
    }
    if (!true) {
        var v31 int16 = 654
        _ = v31
    } else {
        v32 := GenS3{F0: (p1 * p1)}
        var v33 int = 4294967295
        v34 := GenS3{F0: (p1 % (((p1 - 486)) | 1))}
        _ = v32
        _ = v33
        _ = v34
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return p0
}
func genFn1(p0 uint32, p1 uint, p2 uint64) int32 {
    v35 := []int{}
    p0 += (p0 / ((2147483647) | 1))
    if len(v35) > 0 {
        v35[0] = 9223372036854775807
    }
    p1 = 42
    var v36 uint = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    _ = v36
    return -2147483648
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]uint32)
    v3 := []string{"foo bar", "hello", ("" + ("" + "\n\t")), ("" + "")}
    var v4 float32 = -2.25
    v5 := []uint32{2}
    var v37 float64 = ((-2.25 / 3.14159) / 1e10)
    fmt.Println(v4, v4)
    var v38 float32 = 1.5
    var v39 uint8 = 97
    if (516 < genFn1(2147483648, 4294967295, (v1 - (v1 / ((2) | 1))))) {
        _ = len(v2)
    }
    v40 := v2[1]
    v40 -= v40
    v41 := GenS5{F0: 708, F1: 9223372036854775807, F2: 0}
    v42 := GenS6{F0: (("\n\t" + "\n\t") + (("" + (("\n\t" + "\n\t") + "hello")) + ""))}
    v43 := v2[(^(v40))]
    fmt.Println(v38, v1)
    fmt.Println(v1, v38, v39)
    var v44 float32 = float32(v37)
    v1 = v1
    if (1 >= 32767) {
        var v45 uint8 = 0
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
