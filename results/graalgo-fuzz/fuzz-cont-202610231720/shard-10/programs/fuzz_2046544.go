package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint; F1 int32; F2 int64 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 uint16 }
func genFn0() int8 {
    var v6 float64 = 3.14159
    v7 := []int{0, 32768, 2147483647}
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 bool, p1 uint64) int {
    v8 := GenS1{F0: (float64(0.1) > (1.5 + 0.1)), F1: 0}
    v9 := GenS1{F0: false, F1: 979}
    p1 = p1
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 382
}

func main() {
    var v1 uint = 291
    var v2 bool = (775 == 716)
    var v3 int32 = -1
    var v4 int64 = -9223372036854775808
    v5 := make(map[string]uint32)
    var v10 float32 = ((0.0 / 1.0) + 1e10)
    var v11 uint = v1
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := v5["foo bar"]
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 8; fi15++ {
            v16 := v5[""]
            var v17 string = "a"
            switch 42 {
            case 1:
                if ((-2.25 + (v10 - float32((0.0 - (1e10 + (-2.25 / 1.0)))))) != float32(float64(v10))) {
                    if v2 {
                        delete(v5, v17)
                        v18 := []int{fi12, (-(fi15)), (fi15 >> (((v1 >> ((v11) & 63))) & 63)), (fi15 ^ (-(fi15)))}
                        delete(v5, "hello")
                        _ = v18
                    } else {
                        switch (^(fi12)) {
                        case 2147483647:
                            if (42 == 42) {
                                v11 = v11
                            }
                        case 127:
                            v10 = 0.0
                            v3 = (^(v3))
                            var v19 uint16 = 502
                            _ = v19
                        case 9223372036854775807:
                            var v20 bool = (1 == 1000)
                            _ = v20
                        }
                        fmt.Println(v4, v4, fi12)
                        fmt.Println(fi12, v2, v2)
                    }
                    if (32767 > 2) {
                        var v21 uint32 = 0
                        var v22 int16 = int16(v21)
                        if (194 != 0) {
                            v21 = v21
                            v3 = -2147483648
                            switch fi15 {
                            case 888:
                                var v23 int = fi15
                                _ = v23
                            case 2147483647:
                                if v2 {
                                    if true {
                                        if v2 {
                                            var v24 bool = v2
                                            v25 := []int{(fi15 * fi12), 128}
                                            v26 := GenS2{F0: v11, F1: -1, F2: v4}
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        } else {
                                            switch (fi12 | fi12) {
                                            case 128:
                                                var v27 uint32 = v16
                                                _ = v27
                                            default:
                                                fmt.Println(v4, v16, fi15)
                                            }
                                            var v28 int32 = v3
                                            _ = v28
                                        }
                                        fmt.Println(v16)
                                    } else {
                                        if v2 {
                                            v29 := []int{(fi15 % (((fi12 - fi12)) | 1)), (fi12 % ((fi12) | 1))}
                                            var v30 uint64 = 727
                                            switch fi12 {
                                            case -147:
                                                if v2 {
                                                    switch fi15 {
                                                    case 2147483647:
                                                        delete(v5, "a")
                                                    case 7:
                                                        var v31 uint64 = (v30 % (((v30 << ((v11) & 63))) | 1))
                                                        _ = v31
                                                    }
                                                }
                                            case -153:
                                                switch genFn1((!(!((0 >= 127) && v2))), v30) {
                                                case 0:
                                                    var v32 uint16 = 924
                                                    _ = v32
                                                }
                                            default:
                                                switch fi15 {
                                                case 32768:
                                                    if ((v30 % (((v30 + 134)) | 1)) >= (v30 * (v30 + v30))) {
                                                        v16 = (v21 * 0)
                                                    } else {
                                                        switch (^(fi12)) {
                                                        case 1:
                                                            if (2 <= (v1 & (v11 % ((255) | 1)))) {
                                                                fmt.Println(fi12)
                                                                var v33 uint64 = v30
                                                                var v34 float32 = v10
                                                                v33 = (^(v30))
                                                                _ = v33
                                                                _ = v34
                                                            } else {
                                                                fmt.Println(fi15, fi12)
                                                            }
                                                        case 32767:
                                                            if v2 {
                                                                var v35 uint8 = 127
                                                                _ = v35
                                                            } else {
                                                                var v36 uint16 = 0
                                                                v37 := GenS2{F0: v11, F1: (v3 - (v3 >> ((uint(v36)) & 63))), F2: v4}
                                                                _ = v36
                                                                _ = v37
                                                            }
                                                        }
                                                    }
                                                case 0:
                                                    v38 := GenS2{F0: v1, F1: v3, F2: v4}
                                                    _ = v38
                                                case 860:
                                                    v29 = append(v29, 128)
                                                default:
                                                    v2 = v2
                                                }
                                            }
                                            delete(v5, "\n\t")
                                            v39 := GenS3{F0: 42, F1: genFn1((!false), 18446744073709551615)}
                                            _ = v29
                                            _ = v30
                                            _ = v39
                                        } else {
                                            v10 -= v10
                                            v40 := []int{328, (fi12 * (fi12 >> ((v1) & 63))), fi12, fi15}
                                            _ = len(v5)
                                            var v41 string = v17
                                            var v42 int = genFn1(v2, 0)
                                            v42 = -103
                                            var v43 uint16 = 2
                                            _ = v40
                                            _ = v41
                                            _ = v42
                                            _ = v43
                                        }
                                        if v2 {
                                            v44 := GenS4{F0: v21, F1: 0}
                                            v5[(v17 + (v17 + "a"))] = 851
                                            _ = v44
                                        }
                                        v21 += (^(v16))
                                    }
                                    v45 := GenS5{F0: v17, F1: v2}
                                    _ = v45
                                }
                            case -701:
                                var v46 int16 = (v22 | (v22 * v22))
                                _ = v46
                            }
                        }
                        _ = v21
                        _ = v22
                    } else {
                        if v2 {
                            var v47 int = fi12
                            v17 = v17
                            switch v47 {
                            case 32768:
                                fmt.Println(v1)
                            case 32767:
                                var v48 float64 = 0.1
                                _ = v48
                            case 844:
                                var v49 uint16 = 341
                                _ = v49
                            default:
                                if v2 {
                                    switch (fi15 ^ (^(fi12))) {
                                    case 127:
                                        var v50 int32 = -244
                                        v5[v17] = v16
                                        _ = v50
                                    case -1:
                                        var v51 string = v17
                                        _ = v51
                                    case -594:
                                        switch -338 {
                                        case 127:
                                            v52 := []int{}
                                            _ = v52
                                        case 415:
                                            v17 = "foo bar"
                                        }
                                    }
                                    var v53 uint64 = 18446744073709551615
                                    _ = v53
                                }
                            }
                            _ = v47
                        }
                        v17 = (v17 + v17)
                        v11 -= v11
                    }
                    v10 -= 3.14159
                }
            default:
                v3 += v3
            }
            var v54 string = ((v17 + v17) + "a")
            v55 := GenS2{F0: 33, F1: v3, F2: v4}
            _ = fi15
            _ = v16
            _ = v17
            _ = v54
            _ = v55
        }
        _ = fi12
    }
    v5[""] = 2
    if (v2 && (42 != 42)) {
        for fi56 := 0; fi56 < 5; fi56++ {
            switch ("a" + ("\n\t" + ("" + "\n\t"))) {
            case "foo bar":
                v1 += v1
            case "hello":
                var v57 uint8 = 127
                _ = v57
            }
            v58 := []int{(fi56 - 373), genFn1(v2, 338), (fi56 | (fi56 / ((fi56) | 1)))}
            v59 := GenS6{F0: 73}
            _ = fi56
            _ = v58
            _ = v59
        }
    } else {
        v60 := []int{}
        v60 = append(v60, -2147483648)
        v10 -= (float32(1.5) - (((3.14159 / -2.25) - v10) / v10))
        _ = v60
    }
    v61 := GenS2{F0: v1, F1: 859, F2: 141}
    v5["a"] = 2147483647
    v1 -= (v1 | v11)
    v62 := GenS7{F0: (v4 % ((v4) | 1)), F1: 2}
    for fi63 := 0; fi63 < 15; fi63++ {
        _ = len(v5)
        var v64 int16 = 32767
        _ = fi63
        _ = v64
    }
    v10 = v10
    var v65 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v65
}
