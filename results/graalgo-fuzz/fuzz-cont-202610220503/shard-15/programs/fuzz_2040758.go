package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int8; F2 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8; F1 uint32; F2 uint }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint8) int32 {
    switch 127 {
    case 1:
        for fi7 := 0; fi7 < 1; fi7++ {
            if true {
                if (!(!(true || (!(18446744073709551615 <= 18446744073709551615))))) {
                    if (-1 < -839) {
                        v8 := GenS1{F0: -1, F1: 18446744073709551615}
                        v9 := GenS2{F0: ("" + "")}
                        _ = v8
                        _ = v9
                    } else {
                        v10 := GenS3{F0: -32768, F1: 21, F2: -891}
                        _ = v10
                    }
                    v11 := GenS4{F0: (2147483648 < 885)}
                    _ = v11
                } else {
                    v12 := GenS4{F0: true}
                    v13 := []int{(-(fi7)), (fi7 - (-(fi7)))}
                    _ = v12
                    _ = v13
                }
                var v14 uint16 = 0
                _ = v14
            }
            _ = fi7
        }
    case 127:
        var v15 int32 = 42
        _ = v15
    default:
        var v16 float32 = (0.1 / 1.5)
        _ = v16
    }
    p0 = 200
    _ = p0
    return 718
}
func genFn1(p0 int) uint {
    var v17 uint32 = 42
    v17 = 51
    v18 := GenS5{F0: int8(v17), F1: 2, F2: 1}
    fmt.Println(v17, v17)
    fmt.Println(v17, v17)
    _ = p0
    _ = v17
    _ = v18
    return 970
}

func main() {
    var v1 int32 = 99
    var v2 int = -2147483648
    v3 := make(map[int]int)
    var v4 int32 = 1
    v5 := []string{"hello"}
    var v6 uint32 = 1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v19 uint = 679
    var v20 bool = ((!(1 >= 946)) || (0.1 > -2.25))
    var v21 uint64 = 42
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 uint16 = 65535
        for fi24 := 0; fi24 < 15; fi24++ {
            if (127 != 127) {
                if (2 == uint16(v6)) {
                    _ = len(v5)
                    delete(v3, int(v1))
                }
                if true {
                    fmt.Println(v2, v21)
                    var v25 int = v2
                    _ = v25
                }
                if false {
                    if (v6 > (v6 << ((695) & 63))) {
                        v26 := GenS4{F0: ((^(v23)) <= (v23 & v23))}
                        fmt.Println(v21, v19)
                        var v27 uint64 = uint64(v6)
                        switch "\n\t" {
                        case "foo bar":
                            var v28 uint64 = (v27 - 257)
                            _ = v28
                        default:
                            v29 := GenS6{F0: (v23 & (v23 - v23)), F1: 206}
                            _ = v29
                        }
                        _ = v26
                        _ = v27
                    }
                    delete(v3, (^(fi22)))
                }
                v2 += (v2 - (fi22 * fi22))
            } else {
                if (v20 || false) {
                    var v30 float32 = 1.5
                    var v31 int32 = 1
                    _ = v30
                    _ = v31
                }
                v21 += v21
                switch "\n\t" {
                case "":
                    v21 = (v21 & (v21 + 18446744073709551615))
                }
            }
            _ = fi24
        }
        v6 -= (^(v6))
        _ = fi22
        _ = v23
    }
    v19 = v19
    var v32 float64 = 1e10
    for fi33 := 0; fi33 < 20; fi33++ {
        for fi34 := 0; fi34 < 3; fi34++ {
            if ((1 > -1) && true) {
                if true {
                    v35 := GenS7{F0: 65535}
                    if v20 {
                        v36 := GenS3{F0: 1, F1: 1, F2: -32768}
                        var v37 int16 = 42
                        v38 := GenS2{F0: ("a" + "\n\t")}
                        v2 += (v2 / (((fi33 % (((fi33 | (fi34 + (v2 | -612)))) | 1))) | 1))
                        delete(v3, v2)
                        var v39 int16 = (v37 << (((v19 + 4294967295)) & 63))
                        _ = v36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    v1 += 2147483647
                    _ = v35
                }
                v19 = uint(v6)
            }
            _ = fi34
        }
        _ = fi33
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        var v41 int = 636
        v5 = append(v5, (((("hello" + "hello") + "\n\t") + "hello") + ""))
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
