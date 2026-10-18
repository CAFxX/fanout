package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int16 }
type GenS2 struct { F0 int; F1 uint; F2 string }
type GenS3 struct { F0 uint; F1 int; F2 int64 }
type GenS4 struct { F0 uint; F1 uint16; F2 int }
type GenS5 struct { F0 int8; F1 int16; F2 uint8 }
type GenS6 struct { F0 int }
func genFn0(p0 int16, p1 bool, p2 uint8) bool {
    v6 := []int{-380, 127, 32768, -379}
    p2 += (^(p2))
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 1, F1: 0, F2: 1000}
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 string = ""
            if len(v6) > 0 {
                v6[0] = (fi9 & (fi7 / ((-2147483648) | 1)))
            }
            fmt.Println(v10, p1, v10)
            _ = fi9
            _ = v10
        }
        switch -1 {
        case -169:
            for fi11 := 0; fi11 < 8; fi11++ {
                p2 -= (p2 % ((p2) | 1))
                var v12 uint32 = 2
                var v13 int64 = 42
                if (2 > 1) {
                    var v14 int64 = v13
                    v13 = 1
                    _ = v14
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        _ = fi7
        _ = v8
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 uint32 = 4294967295
        var v17 uint64 = 18446744073709551615
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v6
    return (false || false)
}
func genFn1(p0 bool, p1 uint64, p2 int16) bool {
    if p0 {
        for fi19 := 0; fi19 < 10; fi19++ {
            fmt.Println(p1, p1, p0)
            _ = fi19
        }
    }
    switch "a" {
    case "hello":
        if p0 {
            var v20 float64 = 1.5
            var v21 float32 = 0.1
            _ = v20
            _ = v21
        }
    case "":
        var v22 uint64 = 654
        _ = v22
    }
    var v23 uint8 = 93
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return p0
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 715
    var v3 int8 = -1
    v4 := make(map[int]int)
    var v5 uint16 = 42
    var v24 uint32 = 2147483648
    fmt.Println(v2, v5)
    v2 -= v2
    v4[87] = 32768
    var v25 int64 = 1
    v26 := GenS1{F0: 65535, F1: (^(v25)), F2: 42}
    v27 := v4[32767]
    fmt.Println(v3, v25)
    var v28 int32 = -2147483648
    var v29 uint64 = (^(v2))
    switch ("foo bar" + "foo bar") {
    case "a":
        v30 := v4[v27]
        _ = v30
    case "hello":
        for fi31 := 0; fi31 < 2; fi31++ {
            var v32 bool = (float64((-2.25 * 1e10)) == (0.0 - 1.5))
            for fi33 := 0; fi33 < 14; fi33++ {
                v34 := GenS1{F0: 973, F1: (v25 - 42), F2: 290}
                fmt.Println(v32, v25, v29)
                switch fi33 {
                case 401:
                    var v35 int8 = v3
                    delete(v1, (fi33 | (fi31 / (((fi33 >> ((1) & 63))) | 1))))
                    var v36 float64 = 0.0
                    v32 = (10 != 118)
                    _ = v35
                    _ = v36
                case 128:
                    if false {
                        var v37 string = (("" + "hello") + ("" + "\n\t"))
                        _ = v37
                    } else {
                        var v38 uint16 = (v5 ^ 255)
                        _ = v38
                    }
                }
                _ = fi33
                _ = v34
            }
            fmt.Println(v3)
            _ = fi31
            _ = v32
        }
    }
    var v39 uint8 = 128
    if (0 >= 4294967295) {
        for fi40 := 0; fi40 < 18; fi40++ {
            v41 := GenS2{F0: v27, F1: 271, F2: "a"}
            fmt.Println(v27, v2)
            _ = fi40
            _ = v41
        }
        v2 += (v29 + uint64(v39))
        v24 = v24
    }
    if ((v25 - v25) == (v25 + v25)) {
        var v42 int = (v27 << ((69) & 63))
        if (("a" + ((("foo bar" + "a") + "") + "a")) != "\n\t") {
            for fi43 := 0; fi43 < 7; fi43++ {
                fmt.Println(fi43, v3)
                for fi44 := 0; fi44 < 18; fi44++ {
                    var v45 string = "hello"
                    if false {
                        v3 = v3
                    } else {
                        v46 := GenS2{F0: -860, F1: 255, F2: "hello"}
                        v47 := GenS1{F0: uint(v2), F1: (v25 / ((v25) | 1)), F2: 42}
                        if true {
                            _ = len(v1)
                            if (v45 >= v45) {
                                switch (v27 | v27) {
                                case -84:
                                    v48 := GenS2{F0: -1, F1: uint(v5), F2: (((v45 + (v45 + ("a" + "\n\t"))) + v45) + (v45 + "foo bar"))}
                                    fmt.Println(v24)
                                    v25 = v25
                                    v27 += int(v5)
                                    if ((-308 == v28) || (706 == 2)) {
                                        v49 := GenS1{F0: 489, F1: v25, F2: -385}
                                        _ = v49
                                    }
                                    v45 = v45
                                    _ = v48
                                case 0:
                                    if true {
                                        v50 := GenS3{F0: 4294967295, F1: (fi43 - (v42 >> ((0) & 63))), F2: v25}
                                        fmt.Println(v25, v42)
                                        var v51 int16 = 1000
                                        _ = v50
                                        _ = v51
                                    } else {
                                        v29 -= 4294967295
                                        v5 = (v5 ^ 1)
                                        var v52 uint = 690
                                        _ = v52
                                    }
                                default:
                                    var v53 int64 = 0
                                    _ = v53
                                }
                                if true {
                                    var v54 int8 = v3
                                    fmt.Println(v25, v3)
                                    var v55 float64 = 1e10
                                    _ = v54
                                    _ = v55
                                } else {
                                    v56 := []int{(^(fi43)), (v42 / ((v27) | 1)), 4294967295}
                                    fmt.Println(v29, v25)
                                    v57 := GenS4{F0: 498, F1: v5, F2: (v27 | 797)}
                                    _ = v56
                                    _ = v57
                                }
                                v58 := GenS5{F0: (v3 % (((-(v3))) | 1)), F1: 0, F2: 232}
                                v5 = uint16(v24)
                                var v59 uint32 = (v24 * 2147483647)
                                _ = v58
                                _ = v59
                            } else {
                                if true {
                                    v60 := []int{fi44, (fi44 - 284), v27}
                                    _ = v60
                                }
                                v45 = v45
                            }
                        } else {
                            delete(v4, fi43)
                            var v61 uint64 = v2
                            _ = v61
                        }
                        _ = v46
                        _ = v47
                    }
                    _ = fi44
                    _ = v45
                }
                _ = fi43
            }
        } else {
            if false {
                v62 := []int{(v42 ^ v42), (v42 * 128)}
                var v63 int = v42
                fmt.Println(v27, v42)
                _ = v62
                _ = v63
            }
        }
        v25 -= v25
        _ = v42
    } else {
        v64 := GenS6{F0: int(v25)}
        var v65 int64 = v25
        v29 += (v2 / ((v2) | 1))
        _ = v64
        _ = v65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
