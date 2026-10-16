package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 bool; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint64 }
func genFn0(p0 uint32, p1 uint64) uint32 {
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{fi6, -9223372036854775808, 646, (fi6 << ((0) & 63))}
        _ = fi6
        _ = v7
    }
    p0 -= p0
    v8 := []int{-9223372036854775808}
    for fi9 := 0; fi9 < 6; fi9++ {
        p0 -= (p0 >> ((1) & 63))
        v10 := GenS1{F0: -1}
        p0 = p0
        _ = fi9
        _ = v10
    }
    if false {
        if ((1 < -1) && (false || (false || true))) {
            fmt.Println(p1, p1)
            if false {
                v11 := GenS1{F0: -408}
                v8 = append(v8, -429)
                _ = v11
            } else {
                if true {
                    v8 = append(v8, -1)
                    for fi12 := 0; fi12 < 3; fi12++ {
                        if false {
                            fmt.Println(p1, fi12, fi12)
                            if false {
                                for fi13 := 0; fi13 < 1; fi13++ {
                                    v14 := GenS1{F0: 488}
                                    p1 = (p1 | 42)
                                    _ = fi13
                                    _ = v14
                                }
                                if len(v8) > 0 {
                                    v8[0] = 60
                                }
                            }
                            fmt.Println(fi12, fi12)
                        } else {
                            fmt.Println(p1)
                            for fi15 := 0; fi15 < 11; fi15++ {
                                fmt.Println(fi12)
                                p1 = 0
                                _ = fi15
                            }
                            for fi16 := 0; fi16 < 19; fi16++ {
                                var v17 uint16 = 32767
                                var v18 float64 = -2.25
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        }
                        _ = fi12
                    }
                    if true {
                        p0 += 794
                        fmt.Println(p1)
                    } else {
                        for fi19 := 0; fi19 < 14; fi19++ {
                            var v20 float64 = (-2.25 / 3.14159)
                            v21 := GenS1{F0: 7}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        var v22 uint16 = uint16(p1)
                        _ = v22
                    }
                }
                v23 := GenS1{F0: -2147483648}
                _ = v23
            }
            if (!true) {
                v24 := GenS2{F0: 423}
                _ = v24
            }
        } else {
            p0 -= (p0 | p0)
            for fi25 := 0; fi25 < 20; fi25++ {
                var v26 int64 = 9223372036854775807
                if len(v8) > 0 {
                    v8[0] = fi25
                }
                if (1 != (p1 >> ((uint(p1)) & 63))) {
                    var v27 int8 = -128
                    v8 = append(v8, (fi25 + fi25))
                    _ = v27
                } else {
                    fmt.Println(p0)
                    p1 -= (^(p1))
                    var v28 int = -499
                    _ = v28
                }
                _ = fi25
                _ = v26
            }
        }
    } else {
        fmt.Println(p0, p1)
        var v29 int16 = -1
        _ = v29
    }
    _ = p0
    _ = p1
    _ = v8
    return p0
}
func genFn1() uint32 {
    var v30 float64 = 0.0
    fmt.Println(v30)
    v31 := []int{-604}
    _ = v30
    _ = v31
    return genFn0(genFn0(4294967295, 1), 4294967295)
}

func main() {
    var v1 uint32 = 1
    var v2 bool = true
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    v5 := []uint32{v1, v1, v1}
    if (-2147483648 < 967) {
        fmt.Println(v2, v1)
        fmt.Println(v2)
        if (128 < 1) {
            v2 = true
            switch ("a" + "foo bar") {
            case "a":
                for fi32 := 0; fi32 < 9; fi32++ {
                    v33 := GenS2{F0: v1}
                    _ = fi32
                    _ = v33
                }
            case "hello":
                for fi34 := 0; fi34 < 1; fi34++ {
                    if v2 {
                        for fi35 := 0; fi35 < 17; fi35++ {
                            switch ("" + "a") {
                            case "hello":
                                v2 = (127 != int8(fi35))
                            case "foo bar":
                                v1 += (v1 * v1)
                            case "\n\t":
                                v36 := GenS1{F0: 356}
                                _ = v36
                            }
                            var v37 int32 = -1
                            _ = fi35
                            _ = v37
                        }
                    }
                    _ = fi34
                }
            }
            v2 = ((-1 >= -32768) || (v2 && (false && v2)))
        }
    } else {
        v4[(v1 + genFn0((^(v1)), 4294967295))] = ("hello" + "foo bar")
        v38 := GenS3{F0: 970, F1: v2, F2: 1000}
        switch -2147483648 {
        case -9223372036854775808:
            var v39 float32 = float32((0.1 / 1.5))
            _ = v39
        case 4294967295:
            fmt.Println(v2, v2)
        default:
            var v40 int = 32768
            _ = v40
        }
        _ = v38
    }
    var v41 uint64 = 4294967295
    v41 -= 115
    var v42 int32 = 93
    v43 := GenS4{F0: v41, F1: 127, F2: (v41 ^ (v41 ^ (v41 << ((4294967295) & 63))))}
    switch (("foo bar" + "a") + "\n\t") {
    case "\n\t":
        v44 := GenS3{F0: (v42 >> ((18446744073709551615) & 63)), F1: (!v2), F2: 70}
        _ = v44
    }
    delete(v3, 0)
    v45 := GenS2{F0: v1}
    var v46 uint16 = 0
    var v47 bool = (false || v2)
    var v48 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
}
