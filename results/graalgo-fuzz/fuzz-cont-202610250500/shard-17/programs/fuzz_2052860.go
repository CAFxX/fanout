package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
func genFn0(p0 uint32) int32 {
    p0 += p0
    var v5 int32 = -898
    if (-519 > -85) {
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := []int{(fi6 - (fi6 >> ((276) & 63))), fi6, (fi6 + fi6)}
            var v8 int8 = 100
            var v9 float64 = -2.25
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        fmt.Println(v5, v5)
    } else {
        v10 := GenS1{F0: 9223372036854775807}
        _ = v10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        p0 = p0
        v5 += v5
        _ = fi11
    }
    _ = p0
    _ = v5
    return v5
}
func genFn1(p0 int, p1 int64, p2 uint) uint32 {
    p2 = (p2 | p2)
    switch 9223372036854775807 {
    case -312:
        var v12 uint = (p2 ^ 18446744073709551615)
        _ = v12
    }
    var v13 bool = (true && (true && (351 >= 32767)))
    var v14 int8 = -1
    fmt.Println(p1, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 66
}

func main() {
    v1 := []string{}
    var v2 uint = 0
    var v3 uint = 0
    v4 := make(map[uint32]int)
    for fi15 := 0; fi15 < 13; fi15++ {
        switch "" {
        case "":
            if false {
                if false {
                    v2 = v3
                    _ = len(v4)
                    for fi16 := 0; fi16 < 10; fi16++ {
                        fmt.Println(fi15)
                        if (29 < 142) {
                            var v17 uint32 = 2
                            v17 -= 2147483648
                            _ = v17
                        }
                        _ = fi16
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = ("foo bar" + ("" + "foo bar"))
                    }
                }
                var v18 float64 = 0.0
                _ = v18
            } else {
                _ = len(v1)
                v3 = v2
                v19 := GenS2{F0: false}
                _ = v19
            }
        case "\n\t":
            switch 42 {
            case -1:
                for fi20 := 0; fi20 < 1; fi20++ {
                    v2 = v2
                    switch 1 {
                    case 2147483647:
                        v1 = append(v1, "hello")
                    case 32768:
                        var v21 int32 = genFn0(2147483648)
                        var v22 int = fi20
                        switch "" {
                        case "":
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                        case "foo bar":
                            switch fi15 {
                            case 2147483647:
                                fmt.Println(v21, v21)
                            default:
                                if (false || false) {
                                    v2 += 0
                                }
                            }
                        default:
                            v23 := GenS2{F0: false}
                            _ = v23
                        }
                        _ = v21
                        _ = v22
                    case 128:
                        if true {
                            if (80 > 703) {
                                var v24 float32 = float32(0.1)
                                _ = v24
                            } else {
                                var v25 uint16 = 42
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                                v26 := GenS3{F0: (^(v2))}
                                _ = v25
                                _ = v26
                            }
                            v2 += v2
                            if false {
                                v3 += (v3 * v3)
                                fmt.Println(v2, fi15)
                            }
                        }
                    default:
                        v27 := GenS3{F0: v2}
                        _ = v27
                    }
                    _ = fi20
                }
            default:
                v28 := GenS1{F0: fi15}
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            v3 -= (v2 / ((v3) | 1))
            v30 := GenS2{F0: true}
            var v31 uint32 = 1
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = len(v4)
        _ = fi15
    }
    switch 1 {
    case 605:
        v3 -= v2
    case -514:
        var v32 int = 32767
        _ = v32
    }
    v33 := v4[0]
    v3 = (^(v3))
    var v34 int8 = 100
    v35 := GenS4{F0: v33}
    v33 -= (v33 >> (((v3 % ((4294967295) | 1))) & 63))
    var v36 int64 = 604
    delete(v4, 42)
    var v37 int32 = -940
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
