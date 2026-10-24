package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int32 }
type GenS2 struct { F0 int; F1 uint64; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 string; F2 uint64 }
type GenS6 struct { F0 uint32; F1 string; F2 int8 }
func genFn0(p0 int16, p1 int8) int16 {
    v4 := GenS1{F0: 42, F1: 1, F2: -127}
    switch ((((("" + "hello") + "a") + ("\n\t" + ("" + "\n\t"))) + "a") + ("\n\t" + "foo bar")) {
    case "hello":
        fmt.Println(p0, p1)
    case "":
        if (false || (496 > 0)) {
            var v5 uint64 = 2
            _ = v5
        }
    }
    if (-432 <= -1) {
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := GenS2{F0: int(p0), F1: 0, F2: fi6}
            switch -1 {
            case 128:
                for fi8 := 0; fi8 < 16; fi8++ {
                    v9 := []int{fi6, fi8}
                    if len(v9) > 0 {
                        v9[0] = (fi6 - int(p0))
                    }
                    var v10 float32 = 1.5
                    p1 += 0
                    var v11 uint16 = 673
                    switch (fi6 >> ((0) & 63)) {
                    case -9223372036854775808:
                        v11 = (v11 | (v11 / (((v11 + v11)) | 1)))
                    default:
                        switch (fi8 - 4294967295) {
                        case 279:
                            v12 := GenS3{F0: ("\n\t" + "foo bar")}
                            _ = v12
                        case 0:
                            v13 := GenS3{F0: "a"}
                            _ = v13
                        default:
                            if false {
                                fmt.Println(p1, fi8)
                                var v14 uint = 2
                                p0 += p0
                                _ = v14
                            }
                        }
                    }
                    var v15 int64 = 0
                    if false {
                        v9 = append(v9, fi6)
                        v15 -= v15
                    }
                    var v16 uint = 0
                    v15 -= v15
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v15
                    _ = v16
                    _ = v9
                }
            default:
                switch fi6 {
                case -1:
                    if false {
                        switch (fi6 ^ fi6) {
                        case -815:
                            for fi17 := 0; fi17 < 14; fi17++ {
                                switch (fi6 ^ (fi6 % ((fi6) | 1))) {
                                case 0:
                                    v18 := []int{}
                                    _ = v18
                                case 32767:
                                    var v19 int16 = -1
                                    _ = v19
                                case 4294967295:
                                    var v20 int8 = -91
                                    _ = v20
                                }
                                switch "\n\t" {
                                case "hello":
                                    v21 := GenS1{F0: 693, F1: (p1 + (p1 >> ((2) & 63))), F2: 793}
                                    p0 += (p0 - (p0 * (p0 % ((p0) | 1))))
                                    fmt.Println(fi6)
                                    _ = v21
                                case "\n\t":
                                    switch "" {
                                    case "a":
                                        if true {
                                            fmt.Println(fi6)
                                        }
                                    case "hello":
                                        v22 := []int{}
                                        _ = v22
                                    case "\n\t":
                                        if ((!(p0 >= (p0 ^ p0))) && (868 == (-(fi17)))) {
                                            v23 := []int{(fi6 - (fi6 >> ((18446744073709551615) & 63))), (^(fi6)), (fi17 % ((fi6) | 1)), 76}
                                            v24 := GenS4{F0: 32767}
                                            _ = v23
                                            _ = v24
                                        }
                                    }
                                case "":
                                    p1 = p1
                                }
                                v25 := []int{(fi17 | (fi17 & fi17)), fi17}
                                _ = fi17
                                _ = v25
                            }
                        case 519:
                            v26 := []int{(fi6 + fi6), 1, fi6}
                            _ = v26
                        }
                        if true {
                            v27 := GenS4{F0: p0}
                            var v28 int64 = -1
                            _ = v27
                            _ = v28
                        } else {
                            v29 := []int{(-(fi6)), (fi6 ^ (^(fi6)))}
                            _ = v29
                        }
                        var v30 uint32 = 42
                        _ = v30
                    } else {
                        p1 += p1
                        var v31 int64 = 0
                        _ = v31
                    }
                case 9223372036854775807:
                    for fi32 := 0; fi32 < 15; fi32++ {
                        if (!false) {
                            p1 -= p1
                            var v33 int32 = 1
                            _ = v33
                        }
                        v34 := GenS3{F0: ""}
                        v35 := []int{(fi32 - (fi6 * -2147483648)), fi32, (fi6 * fi32), fi6}
                        _ = fi32
                        _ = v34
                        _ = v35
                    }
                default:
                    var v36 int64 = 9223372036854775807
                    _ = v36
                }
            }
            _ = fi6
            _ = v7
        }
    }
    var v37 int16 = (^(p0))
    _ = p0
    _ = p1
    _ = v37
    _ = v4
    return v37
}

func main() {
    v1 := []uint32{2147483648, 447}
    v2 := make(map[string]string)
    var v3 int16 = 42
    v3 = (-(v3))
    var v38 bool = true
    if true {
        var v39 uint64 = 0
        _ = len(v1)
        _ = v39
    }
    var v40 float64 = 1e10
    v41 := GenS5{F0: 42, F1: ("" + ""), F2: 42}
    fmt.Println(v40, v38, v38)
    for fi42 := 0; fi42 < 4; fi42++ {
        var v43 uint = 4294967295
        delete(v2, ("a" + ""))
        _ = fi42
        _ = v43
    }
    v44 := GenS2{F0: 672, F1: 1, F2: 128}
    for fi45 := 0; fi45 < 18; fi45++ {
        if (114 != 127) {
            var v46 bool = (!(115 != 444))
            fmt.Println(fi45, v3, v46)
            _ = v46
        } else {
            var v47 int64 = -9223372036854775808
            fmt.Println(fi45)
            v48 := GenS6{F0: 42, F1: "a", F2: 42}
            _ = v47
            _ = v48
        }
        switch (("" + ("foo bar" + ("a" + ("" + "foo bar")))) + "\n\t") {
        case "a":
            v1 = append(v1, 711)
        case "foo bar":
            if (2 == 89) {
                var v49 uint32 = 1
                v38 = v38
                v49 += 65535
                _ = v49
            }
        }
        _ = fi45
    }
    v38 = v38
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v40
    _ = v41
    _ = v44
}
