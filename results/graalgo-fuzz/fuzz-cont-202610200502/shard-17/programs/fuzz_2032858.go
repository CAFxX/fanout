package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint64; F2 string }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 uint64; F1 bool }
func genFn0(p0 uint64, p1 uint8, p2 uint) uint8 {
    var v6 uint16 = 247
    v7 := []int{}
    v6 -= 933
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return uint8(p0)
}
func genFn1(p0 string, p1 int8, p2 string) int {
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int16 = 32767
        _ = fi8
        _ = v9
    }
    var v10 int32 = 1
    v11 := []int{4294967295, -121}
    var v12 uint = 255
    for fi13 := 0; fi13 < 9; fi13++ {
        p1 = (p1 + 100)
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 126
}

func main() {
    var v1 bool = (827 <= 42)
    var v2 string = "hello"
    var v3 int32 = 137
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    var v14 string = v2
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(fi15)
        _ = fi15
    }
    v16 := GenS1{F0: -9223372036854775808}
    var v17 int16 = 32767
    var v18 bool = v1
    v19 := []int{-9223372036854775808, 42, 247}
    var v20 uint32 = 65535
    v21 := GenS1{F0: -9223372036854775808}
    v20 = 2147483648
    v22 := GenS1{F0: 566}
    for fi23 := 0; fi23 < 7; fi23++ {
        if v1 {
            for fi24 := 0; fi24 < 17; fi24++ {
                fmt.Println(v18)
                _ = fi24
            }
            var v25 int16 = int16(fi23)
            var v26 int16 = 32767
            _ = v25
            _ = v26
        }
        if true {
            v4[v20] = ((v2 + (v14 + v14)) + v2)
            v27 := GenS2{F0: true, F1: 1, F2: 65535}
            if v18 {
                fmt.Println(v18, v20)
                switch (fi23 | (fi23 ^ -250)) {
                case 602:
                    if (v18 && v1) {
                        v2 = v2
                        var v28 uint = 42
                        _ = v28
                    } else {
                        if v1 {
                            var v29 uint = 0
                            var v30 int32 = (v3 * (^(v3)))
                            for fi31 := 0; fi31 < 3; fi31++ {
                                var v32 uint16 = 42
                                if v18 {
                                    if v1 {
                                        var v33 string = (v14 + ("\n\t" + "\n\t"))
                                        if true {
                                            var v34 int = (fi23 << (((v29 / (((^(v29))) | 1))) & 63))
                                            v35 := GenS3{F0: (v18 || (100 == -1))}
                                            _ = v34
                                            _ = v35
                                        } else {
                                            if v18 {
                                                v36 := v4[v20]
                                                var v37 uint8 = 2
                                                v29 = 65535
                                                _ = v36
                                                _ = v37
                                            }
                                        }
                                        _ = v33
                                    } else {
                                        switch fi31 {
                                        case 9223372036854775807:
                                            if v1 {
                                                var v38 float32 = ((((0.1 * 0.1) / -2.25) * (1.5 - float32(-2.25))) + 1e10)
                                                switch ("" + "hello") {
                                                case "\n\t":
                                                    v30 += -2147483648
                                                    var v39 int64 = 0
                                                    v40 := GenS4{F0: 799, F1: 2, F2: v2}
                                                    _ = v39
                                                    _ = v40
                                                case "a":
                                                    v41 := v4[v20]
                                                    _ = v41
                                                }
                                                _ = v38
                                            } else {
                                                if v18 {
                                                    v42 := GenS1{F0: 9223372036854775807}
                                                    v43 := GenS5{F0: v30, F1: (v17 & v17)}
                                                    _ = v42
                                                    _ = v43
                                                }
                                                var v44 float64 = 1e10
                                                _ = v44
                                            }
                                        }
                                        v45 := GenS4{F0: 612, F1: 436, F2: v2}
                                        var v46 uint16 = (^(v32))
                                        _ = v45
                                        _ = v46
                                    }
                                }
                                if (!true) {
                                    v30 = 0
                                    v47 := GenS1{F0: 1}
                                    _ = v47
                                }
                                if v1 {
                                    v20 -= (v20 >> ((77) & 63))
                                } else {
                                    var v48 float32 = ((1e10 - 0.0) + 3.14159)
                                    _ = v48
                                }
                                var v49 int16 = int16(v20)
                                _ = fi31
                                _ = v32
                                _ = v49
                            }
                            _ = v29
                            _ = v30
                        } else {
                            var v50 uint32 = v20
                            _ = v50
                        }
                    }
                default:
                    _ = len(v19)
                }
                v17 -= v17
            } else {
                v17 -= (-(v17))
                switch fi23 {
                case -428:
                    v3 = v3
                case 9223372036854775807:
                    if len(v19) > 0 {
                        v19[0] = genFn1("", -41, v14)
                    }
                case 0:
                    v2 = (v2 + "foo bar")
                default:
                    var v51 uint16 = 718
                    _ = v51
                }
            }
            _ = v27
        }
        v14 = "foo bar"
        _ = fi23
    }
    var v52 uint = 4294967295
    v53 := v5[(v2 + "")]
    var v54 int32 = (v3 ^ (v3 - (v3 << ((v52) & 63))))
    for fi55 := 0; fi55 < 13; fi55++ {
        v56 := GenS6{F0: 701, F1: true}
        for fi57 := 0; fi57 < 1; fi57++ {
            var v58 uint = v52
            switch (fi57 % ((genFn1(((("a" + "hello") + (v14 + v2)) + "hello"), 1, "\n\t")) | 1)) {
            case 651:
                if len(v19) > 0 {
                    v19[0] = fi55
                }
            case 2147483647:
                var v59 uint8 = 2
                _ = v59
            }
            fmt.Println(fi57, fi55)
            _ = fi57
            _ = v58
        }
        v1 = (0 < 18446744073709551615)
        _ = fi55
        _ = v56
    }
    var v60 uint32 = (v20 >> (((v52 & 18446744073709551615)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v60
}
