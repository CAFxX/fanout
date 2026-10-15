package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int, p1 uint) int16 {
    p1 -= 972
    fmt.Println(p0)
    v4 := GenS1{F0: 0, F1: (p1 | (p1 / ((6) | 1)))}
    p1 = (p1 + (p1 * (p1 | (p1 % ((p1) | 1)))))
    _ = p0
    _ = p1
    _ = v4
    return 508
}

func main() {
    var v1 uint32 = 2147483647
    var v2 float32 = 1.5
    v3 := make(map[uint32]bool)
    if true {
        var v5 int = -642
        _ = v5
    } else {
        v1 = v1
    }
    switch "\n\t" {
    case "a":
        v6 := []int{42, 1}
        _ = v6
    case "\n\t":
        v7 := GenS2{F0: 0}
        _ = v7
    case "foo bar":
        var v8 int = -186
        _ = v8
    default:
        for fi9 := 0; fi9 < 19; fi9++ {
            if true {
                v10 := GenS2{F0: 776}
                var v11 string = ""
                for fi12 := 0; fi12 < 1; fi12++ {
                    v2 = v2
                    _ = fi12
                }
                _ = v10
                _ = v11
            }
            _ = fi9
        }
    }
    var v13 int16 = 865
    var v14 int64 = 9223372036854775807
    v2 += v2
    fmt.Println(v13, v14)
    for fi15 := 0; fi15 < 7; fi15++ {
        if true {
            var v16 uint8 = 128
            v17 := GenS1{F0: (-(v14)), F1: 18446744073709551615}
            _ = v16
            _ = v17
        } else {
            switch "\n\t" {
            case "foo bar":
                v18 := []int{32767, fi15, fi15}
                _ = v18
            case "a":
                fmt.Println(v14, v1)
            case "":
                v13 += (-(v13))
            }
            switch (fi15 + (fi15 << ((348) & 63))) {
            case 2147483647:
                if true {
                    switch (fi15 - (-(fi15))) {
                    case 127:
                        v2 += v2
                    case 32767:
                        v13 -= 1
                    case 92:
                        var v19 float32 = v2
                        _ = v19
                    default:
                        var v20 int16 = (-(v13))
                        _ = v20
                    }
                } else {
                    if (!true) {
                        v2 += v2
                    } else {
                        if false {
                            var v21 string = ""
                            var v22 uint16 = 996
                            _ = v21
                            _ = v22
                        }
                        var v23 uint = 18446744073709551615
                        _ = v23
                    }
                    fmt.Println(fi15, v1, v14)
                    if false {
                        v13 = v13
                        for fi24 := 0; fi24 < 17; fi24++ {
                            v13 = v13
                            v25 := []int{}
                            v25 = append(v25, (fi24 | fi24))
                            v26 := GenS1{F0: v14, F1: 4294967295}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                }
            case -1:
                if false {
                    v27 := []int{}
                    var v28 int32 = -8
                    _ = v27
                    _ = v28
                } else {
                    fmt.Println(fi15, fi15, v13)
                    var v29 uint16 = 2
                    _ = v29
                }
            case -2147483648:
                delete(v3, (^(v1)))
            default:
                var v30 bool = (0 > 0)
                _ = v30
            }
            if true {
                switch (fi15 % (((^(fi15))) | 1)) {
                case 32768:
                    for fi31 := 0; fi31 < 16; fi31++ {
                        v14 = v14
                        if false {
                            v13 = v13
                        } else {
                            var v32 uint64 = 18446744073709551615
                            v3[(v1 * (v1 + 708))] = true
                            v33 := []int{fi15}
                            switch "" {
                            case "":
                                var v34 float64 = (-2.25 - (0.1 * (((3.14159 * (1e10 * 0.0)) + 1.5) + 1e10)))
                                _ = v34
                            }
                            v3[1] = false
                            _ = v32
                            _ = v33
                        }
                        _ = fi31
                    }
                case 505:
                    if (!(-1 <= (^(v13)))) {
                        var v35 string = ""
                        for fi36 := 0; fi36 < 6; fi36++ {
                            v13 = v13
                            _ = fi36
                        }
                        _ = v35
                    }
                case 22:
                    v13 = (v13 << ((65535) & 63))
                default:
                    v37 := GenS1{F0: (v14 / (((v14 % (((^(v14))) | 1))) | 1)), F1: 2}
                    _ = v37
                }
                for fi38 := 0; fi38 < 19; fi38++ {
                    v2 += 3.14159
                    v14 = (v14 | 882)
                    _ = fi38
                }
                if true {
                    var v39 int16 = 1
                    _ = v39
                }
            } else {
                v1 -= v1
                var v40 int = fi15
                _ = v40
            }
        }
        v41 := GenS2{F0: 42}
        v42 := []int{(-(fi15)), (fi15 * fi15), fi15, fi15}
        _ = fi15
        _ = v41
        _ = v42
    }
    v14 = (-(v14))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
