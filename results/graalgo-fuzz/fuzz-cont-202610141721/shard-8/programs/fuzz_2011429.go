package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int32 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int }
func genFn0(p0 uint32) uint16 {
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            fmt.Println(fi5, p0, fi6)
            _ = fi6
        }
        for fi7 := 0; fi7 < 2; fi7++ {
            if ("foo bar" > "\n\t") {
                var v8 float64 = 3.14159
                if true {
                    switch fi7 {
                    case -9223372036854775808:
                        v9 := GenS1{F0: -845, F1: fi7, F2: 1}
                        p0 = 612
                        _ = v9
                    }
                    v10 := []int{fi5, 32767, fi7, fi7}
                    _ = v10
                } else {
                    var v11 float32 = 1e10
                    if (false && true) {
                        v12 := []int{(fi7 - (fi7 / ((42) | 1))), (fi7 | fi7)}
                        fmt.Println(fi7, v11, fi7)
                        if (618 < 255) {
                            var v13 int16 = 0
                            v11 -= v11
                            _ = v13
                        } else {
                            var v14 int = fi5
                            v15 := GenS2{F0: (fi7 / ((v14) | 1)), F1: 0}
                            if true {
                                if len(v12) > 0 {
                                    v12[0] = 2147483647
                                }
                                v14 = -756
                                p0 += 856
                            }
                            _ = v14
                            _ = v15
                        }
                        v16 := GenS1{F0: 1, F1: fi5, F2: 75}
                        if len(v12) > 0 {
                            v12[0] = (fi5 + 42)
                        }
                        _ = v12
                        _ = v16
                    }
                    fmt.Println(fi5)
                    _ = v11
                }
                _ = v8
            }
            var v17 int16 = 1
            var v18 uint8 = 127
            _ = fi7
            _ = v17
            _ = v18
        }
        v19 := GenS2{F0: (fi5 | fi5), F1: 2}
        _ = fi5
        _ = v19
    }
    var v20 float32 = 0.0
    _ = p0
    _ = v20
    return 837
}

func main() {
    var v1 uint32 = 2
    var v2 int32 = 1
    v3 := []int{-9223372036854775808, -1}
    var v4 float32 = 1.5
    fmt.Println(v1)
    v21 := GenS3{F0: 42, F1: 1, F2: 2}
    var v22 bool = true
    var v23 int64 = 1
    v23 -= (v23 - v23)
    if (v22 && (true && v22)) {
        v24 := GenS2{F0: 42, F1: 42}
        var v25 int16 = 32767
        for fi26 := 0; fi26 < 18; fi26++ {
            var v27 int8 = -128
            var v28 uint8 = 2
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = v24
        _ = v25
    } else {
        var v29 uint8 = uint8(v1)
        var v30 uint32 = 182
        var v31 int8 = 0
        _ = v29
        _ = v30
        _ = v31
    }
    if v22 {
        switch "\n\t" {
        case "a":
            for fi32 := 0; fi32 < 14; fi32++ {
                v33 := GenS1{F0: (v23 | int64(v2)), F1: (fi32 & -94), F2: -1}
                if true {
                    var v34 bool = v22
                    _ = v34
                }
                _ = fi32
                _ = v33
            }
        case "\n\t":
            _ = len(v3)
        }
        var v35 bool = v22
        _ = v35
    } else {
        var v36 uint32 = (v1 / ((v1) | 1))
        if (-2.25 < -2.25) {
            for fi37 := 0; fi37 < 8; fi37++ {
                switch "" {
                case "hello":
                    v36 -= v1
                case "":
                    var v38 int16 = 791
                    _ = v38
                default:
                    switch (("foo bar" + ("foo bar" + "\n\t")) + "a") {
                    case "\n\t":
                        for fi39 := 0; fi39 < 19; fi39++ {
                            v36 += (v36 / (((^(v1))) | 1))
                            _ = fi39
                        }
                    }
                }
                if (v23 <= 42) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v36 = (v1 << ((2) & 63))
                    switch fi37 {
                    case 203:
                        fmt.Println(v22, v23)
                    default:
                        fmt.Println(v22)
                    }
                }
                var v40 int32 = 1
                _ = fi37
                _ = v40
            }
            v36 += (v36 / (((v36 ^ 42)) | 1))
        } else {
            v36 = v36
            for fi41 := 0; fi41 < 12; fi41++ {
                var v42 float32 = v4
                _ = fi41
                _ = v42
            }
            var v43 uint = 255
            _ = v43
        }
        v36 += (v1 >> ((0) & 63))
        _ = v36
    }
    v44 := GenS4{F0: 4294967295}
    if v22 {
        var v45 int64 = -563
        switch "hello" {
        case "hello":
            var v46 int32 = 999
            _ = v46
        case "foo bar":
            v47 := GenS3{F0: (v45 ^ int64(v1)), F1: 2, F2: 0}
            _ = v47
        }
        if false {
            if len(v3) > 0 {
                v3[0] = 32768
            }
        } else {
            fmt.Println(v4)
        }
        _ = v45
    }
    switch ("hello" + (("a" + "a") + ((("" + "foo bar") + "foo bar") + ("" + "hello")))) {
    case "a":
        fmt.Println(v22, v2, v2)
    case "\n\t":
        if (4294967295 >= 18446744073709551615) {
            v48 := GenS1{F0: (v23 * v23), F1: -665, F2: -2147483648}
            var v49 int64 = (v23 ^ (v23 + v23))
            if v22 {
                if len(v3) > 0 {
                    v3[0] = -9223372036854775808
                }
            }
            _ = v48
            _ = v49
        }
    case "foo bar":
        var v50 int16 = 0
        _ = v50
    default:
        var v51 int16 = 0
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v44
}
