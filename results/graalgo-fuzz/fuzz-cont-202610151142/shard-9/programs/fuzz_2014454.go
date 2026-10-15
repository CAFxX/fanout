package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int16, p1 int, p2 int16) int16 {
    switch p1 {
    case 980:
        v6 := []int{p1, p1, (p1 >> ((91) & 63)), (p1 ^ 4294967295)}
        _ = v6
    case 0:
        v7 := []int{(^(p1))}
        _ = v7
    case 42:
        var v8 int8 = 100
        _ = v8
    default:
        var v9 uint8 = 174
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS1{F0: (fi10 + fi10)}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 ^ 1000)
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 0
    v3 := make(map[string]string)
    var v4 uint16 = uint16(v2)
    v5 := make(map[uint32]int)
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS1{F0: (-(fi12))}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        if false {
            switch "a" {
            case "\n\t":
                for fi15 := 0; fi15 < 13; fi15++ {
                    v4 = v4
                    if true {
                        var v16 uint = 1
                        if (!(uint32(v2) == 190)) {
                            switch fi15 {
                            case 127:
                                v4 = v4
                            default:
                                if true {
                                    switch "foo bar" {
                                    case "hello":
                                        v17 := v3["\n\t"]
                                        _ = v17
                                    case "":
                                        v18 := GenS1{F0: -2147483648}
                                        _ = v18
                                    }
                                    var v19 int32 = 456
                                    v20 := GenS1{F0: fi15}
                                    _ = v19
                                    _ = v20
                                } else {
                                    fmt.Println(v4, fi14, v16)
                                    v21 := GenS2{F0: 59}
                                    _ = v21
                                }
                            }
                            v4 = 331
                        }
                        switch fi15 {
                        case 0:
                            switch fi15 {
                            case -874:
                                var v22 uint64 = 12
                                _ = v22
                            }
                        case 2147483647:
                            v2 += (v2 & (v2 >> (((v16 * (v16 >> ((v16) & 63)))) & 63)))
                        }
                        _ = v16
                    }
                    var v23 int8 = 100
                    _ = fi15
                    _ = v23
                }
            case "":
                v2 = v2
            case "foo bar":
                for fi24 := 0; fi24 < 10; fi24++ {
                    if true {
                        v2 += 1
                        if true {
                            fmt.Println(fi14, v4, fi14)
                            var v25 int8 = 74
                            if (v4 < v4) {
                                if false {
                                    v25 = (-(v25))
                                    if ((float32(0.0) - (1e10 - (-2.25 / (float32(3.14159) + -2.25)))) <= 0.1) {
                                        if (!false) {
                                            var v26 int32 = -2147483648
                                            var v27 uint8 = (^(v2))
                                            var v28 bool = (0.1 >= 1e10)
                                            v4 = v4
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        } else {
                                            v4 = v4
                                        }
                                        if false {
                                            fmt.Println(v25, fi14, fi24)
                                            if false {
                                                var v29 int16 = -32768
                                                v30 := GenS3{F0: -92, F1: v4}
                                                if true {
                                                    v25 -= int8(fi24)
                                                    switch fi24 {
                                                    case 32768:
                                                        fmt.Println(fi24, v29)
                                                    }
                                                }
                                                var v31 string = "foo bar"
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            }
                                            v25 -= (-(v25))
                                        } else {
                                            v2 += (v2 ^ v2)
                                            switch fi14 {
                                            case 9223372036854775807:
                                                fmt.Println(fi24)
                                            case 32767:
                                                v32 := []int{(fi24 & 9223372036854775807)}
                                                _ = v32
                                            }
                                            if (!(true && false)) {
                                                v2 += v2
                                            }
                                            var v33 int64 = 870
                                            v2 = v2
                                            _ = v33
                                        }
                                    } else {
                                        v25 = (v25 * (v25 - 100))
                                        if (true && false) {
                                            var v34 uint64 = 1
                                            v35 := GenS4{F0: v4, F1: (^(v4))}
                                            _ = v34
                                            _ = v35
                                        }
                                    }
                                    var v36 bool = false
                                    _ = v36
                                }
                                v37 := GenS3{F0: -1, F1: (v4 ^ (v4 * v4))}
                                _ = v37
                            }
                            _ = v25
                        } else {
                            v38 := GenS1{F0: (fi14 ^ (fi24 >> ((853) & 63)))}
                            _ = v38
                        }
                    }
                    v39 := []int{}
                    v4 = v4
                    _ = fi24
                    _ = v39
                }
            default:
                v4 = 42
            }
        }
        var v40 int16 = 42
        v40 += -717
        _ = fi14
        _ = v40
    }
    var v41 float64 = 1e10
    _ = len(v3)
    v4 -= v4
    v42 := []int{9223372036854775807}
    v43 := GenS4{F0: (v4 >> ((876) & 63)), F1: v4}
    v41 = v41
    for fi44 := 0; fi44 < 4; fi44++ {
        fmt.Println(v41, v41, v41)
        for fi45 := 0; fi45 < 9; fi45++ {
            fmt.Println(v41, v2, v2)
            v3[(("" + "foo bar") + ("" + "foo bar"))] = "foo bar"
            _ = fi45
        }
        var v46 int32 = 1
        _ = fi44
        _ = v46
    }
    v41 = (0.0 - (float64(float32(3.14159)) + v41))
    v47 := GenS3{F0: 1, F1: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v47
    _ = v5
}
