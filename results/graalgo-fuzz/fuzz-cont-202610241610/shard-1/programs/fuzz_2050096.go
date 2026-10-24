package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 bool) string {
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 bool = p0
            var v9 uint32 = 4294967295
            if (true || p0) {
                if v8 {
                    v8 = ((255 != 255) || v8)
                    var v10 uint32 = v9
                    _ = v10
                } else {
                    fmt.Println(p0, v8)
                    var v11 uint8 = 1
                    var v12 int8 = -67
                    _ = v11
                    _ = v12
                }
                v13 := []int{fi6, fi7, 621}
                v14 := GenS1{F0: false}
                _ = v13
                _ = v14
            } else {
                v8 = (!v8)
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        v15 := GenS1{F0: true}
        v16 := GenS2{F0: -6, F1: 42, F2: 42}
        _ = fi6
        _ = v15
        _ = v16
    }
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return "hello"
}
func genFn1(p0 uint, p1 bool, p2 int16) string {
    var v17 int = 127
    v18 := []int{(v17 / (((^(v17))) | 1)), 42}
    if len(v18) > 0 {
        v18[0] = v17
    }
    p0 = p0
    switch v17 {
    case 32768:
        for fi19 := 0; fi19 < 11; fi19++ {
            if (218 <= 165) {
                fmt.Println(p1)
                v20 := GenS2{F0: 13, F1: 42, F2: 255}
                p2 += 42
                _ = v20
            } else {
                var v21 bool = false
                if len(v18) > 0 {
                    v18[0] = v17
                }
                _ = v21
            }
            if ((p1 || ((((3.14159 >= 0.1) && true) && (p1 || (!(-1 > 258)))) || p1)) || ((^(v17)) <= fi19)) {
                if (163 == 127) {
                    if len(v18) > 0 {
                        v18[0] = fi19
                    }
                } else {
                    switch v17 {
                    case 4294967295:
                        for fi22 := 0; fi22 < 15; fi22++ {
                            var v23 float64 = (((((float64(1e10) + (3.14159 + (0.0 / (-2.25 / ((1.5 + 1e10) * (-2.25 / float64((0.1 * 3.14159)))))))) + (0.0 - 1e10)) / 1.5) - 3.14159) - 1.5)
                            _ = fi22
                            _ = v23
                        }
                    }
                    var v24 uint64 = 2
                    if p1 {
                        if (p1 || p1) {
                            fmt.Println(p0, p1, v24)
                        } else {
                            for fi25 := 0; fi25 < 3; fi25++ {
                                if p1 {
                                    _ = len(v18)
                                    if p1 {
                                        v17 = (fi25 ^ (fi25 ^ (^(v17))))
                                        var v26 uint16 = 0
                                        _ = v26
                                    }
                                    v27 := GenS3{F0: false, F1: 2147483647}
                                    if p1 {
                                        v28 := GenS2{F0: 100, F1: (v24 << ((2) & 63)), F2: 42}
                                        _ = v28
                                    } else {
                                        fmt.Println(fi25, v17, v24)
                                    }
                                    _ = v27
                                }
                                var v29 bool = p1
                                _ = fi25
                                _ = v29
                            }
                            fmt.Println(fi19)
                        }
                        for fi30 := 0; fi30 < 4; fi30++ {
                            v31 := GenS4{F0: 48}
                            if len(v18) > 0 {
                                v18[0] = fi19
                            }
                            v24 += (v24 % ((1) | 1))
                            _ = fi30
                            _ = v31
                        }
                    }
                    _ = v24
                }
            } else {
                if len(v18) > 0 {
                    v18[0] = (^(fi19))
                }
            }
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            _ = fi19
        }
    case 128:
        if (1 <= 177) {
            var v32 bool = true
            switch "" {
            case "\n\t":
                v33 := GenS5{F0: 9223372036854775807, F1: 4294967295}
                _ = v33
            case "a":
                if (!false) {
                    v17 -= v17
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                } else {
                    switch v17 {
                    case -775:
                        if ((float64(1e10) > (float64(float32(3.14159)) - (0.1 / 1.5))) && p1) {
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                        } else {
                            v32 = (!false)
                            var v34 int = 42
                            for fi35 := 0; fi35 < 4; fi35++ {
                                fmt.Println(v17, p1, p1)
                                _ = fi35
                            }
                            _ = v34
                        }
                    case 536:
                        if (((1000 != (p0 * 1)) && false) && p1) {
                            v36 := GenS1{F0: false}
                            _ = v36
                        }
                    }
                }
            case "foo bar":
                var v37 int64 = -717
                _ = v37
            default:
                if true {
                    var v38 string = ""
                    _ = v38
                }
            }
            _ = v32
        } else {
            v39 := GenS6{F0: p2}
            p1 = p1
            _ = v39
        }
    case 715:
        v17 += -522
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return "hello"
}

func main() {
    var v1 float64 = ((1.5 + (1.5 + (0.0 / 1.5))) * (-2.25 / (1e10 * 1.5)))
    var v2 uint16 = 1
    var v3 uint16 = (v2 | v2)
    var v4 int32 = -2147483648
    var v5 uint = 811
    fmt.Println(v3, v1)
    v5 += v5
    if ("\n\t" != "") {
        v40 := []int{2147483647, 774, -2147483648}
        v1 -= 1e10
        fmt.Println(v1, v1, v1)
        _ = v40
    } else {
        v41 := GenS5{F0: 9223372036854775807, F1: 42}
        var v42 int8 = 1
        v43 := GenS6{F0: 0}
        _ = v41
        _ = v42
        _ = v43
    }
    if (true || true) {
        fmt.Println(v4, v2)
        if true {
            var v44 int16 = -961
            v45 := []int{44, -1, 9223372036854775807}
            _ = v44
            _ = v45
        }
    }
    v46 := GenS7{F0: 733}
    v47 := []int{32767, 542, 32767, -9223372036854775808}
    _ = len(v47)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v2, v4)
    case "a":
        v3 = v3
    }
    var v48 uint64 = 4294967295
    v49 := GenS1{F0: true}
    var v50 int64 = -996
    v51 := GenS8{F0: 2}
    v50 += (^(v50))
    var v52 int64 = (v50 << (((^(v5))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
