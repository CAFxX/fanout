package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint64; F1 int8; F2 int64 }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 uint32; F1 bool; F2 int16 }
type GenS9 struct { F0 string; F1 uint64; F2 int32 }
type GenS10 struct { F0 int16 }
func genFn0(p0 int32, p1 uint8, p2 int) int8 {
    switch (p2 << ((65535) & 63)) {
    case 2147483647:
        v6 := GenS1{F0: -9223372036854775808, F1: 0, F2: (p1 << ((255) & 63))}
        _ = v6
    case 128:
        var v7 int = (p2 + (p2 % (((p2 + (p2 * p2))) | 1)))
        _ = v7
    case -1:
        if true {
            var v8 string = ("a" + "")
            for fi9 := 0; fi9 < 20; fi9++ {
                v8 = v8
                v10 := []int{}
                _ = fi9
                _ = v10
            }
            for fi11 := 0; fi11 < 11; fi11++ {
                if false {
                    v12 := GenS2{F0: v8}
                    v13 := []int{-265}
                    v14 := GenS1{F0: 9223372036854775807, F1: 2147483648, F2: p1}
                    _ = v12
                    _ = v13
                    _ = v14
                }
                if (p2 == p2) {
                    v8 = (v8 + (v8 + v8))
                    v15 := GenS3{F0: fi11}
                    v16 := GenS3{F0: p2}
                    _ = v15
                    _ = v16
                } else {
                    switch "" {
                    case "foo bar":
                        var v17 float64 = 0.1
                        _ = v17
                    default:
                        v18 := GenS1{F0: -1, F1: 2, F2: 0}
                        _ = v18
                    }
                }
                if false {
                    v19 := GenS2{F0: v8}
                    fmt.Println(p0, p1, fi11)
                    p0 -= (p0 / ((p0) | 1))
                    _ = v19
                } else {
                    v20 := GenS2{F0: v8}
                    _ = v20
                }
                _ = fi11
            }
            _ = v8
        }
    }
    v21 := []int{(^(p2)), (-(p2)), p2, p2}
    v22 := GenS2{F0: "foo bar"}
    var v23 float32 = ((0.0 / (1.5 * 0.1)) - 1.5)
    for fi24 := 0; fi24 < 4; fi24++ {
        for fi25 := 0; fi25 < 18; fi25++ {
            v26 := GenS4{F0: (p0 % (((p0 ^ 1)) | 1)), F1: 0}
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 1; fi27++ {
            fmt.Println(p1, p0, p2)
            var v28 int8 = 127
            var v29 int = (fi24 % ((fi27) | 1))
            _ = fi27
            _ = v28
            _ = v29
        }
        fmt.Println(fi24)
        _ = fi24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    _ = v23
    return 0
}

func main() {
    v1 := []uint32{86, 1}
    v2 := make(map[int]int)
    var v3 float64 = 1.5
    var v4 float32 = 3.14159
    var v5 int32 = -461
    var v30 int64 = 589
    switch ("" + ("hello" + "a")) {
    case "hello":
        var v31 float32 = (float32((v3 + v3)) / v4)
        _ = v31
    case "":
        if (!((0 != 2) || (-2.25 != (v3 - 3.14159)))) {
            v32 := GenS5{F0: 2, F1: genFn0(-2147483648, 255, 1), F2: (v30 * v30)}
            v4 = ((3.14159 - v4) - v4)
            v5 += (v5 | (v5 + (v5 * (v5 & 210))))
            _ = v32
        }
    default:
        v4 -= (v4 * (1e10 + (v4 - 3.14159)))
    }
    v33 := GenS3{F0: 133}
    v4 = (1e10 + 0.0)
    for fi34 := 0; fi34 < 9; fi34++ {
        if (!false) {
            var v35 int8 = 127
            _ = v35
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v2)
            var v37 int = fi36
            _ = fi36
            _ = v37
        }
        _ = fi34
    }
    fmt.Println(v3, v5)
    switch "hello" {
    case "":
        var v38 uint64 = 1
        _ = v38
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    fmt.Println(v5)
    v39 := GenS6{F0: -128, F1: 0}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch ("foo bar" + ("foo bar" + "\n\t")) {
    case "\n\t":
        if ((3.14159 / (1e10 - ((float32(0.0) / (v4 + float32((((v3 + (-2.25 * (v3 * v3))) / v3) + (v3 + float64(v4)))))) / (v4 - float32(3.14159))))) < (1.5 - v4)) {
            var v40 uint32 = 2147483647
            _ = v40
        } else {
            v1 = append(v1, 2147483647)
            fmt.Println(v30, v3, v3)
        }
    case "a":
        v41 := GenS7{F0: 1, F1: 170}
        _ = v41
    case "hello":
        v42 := GenS8{F0: 173, F1: false, F2: -1}
        _ = v42
    default:
        for fi43 := 0; fi43 < 10; fi43++ {
            if true {
                if ("hello" < "") {
                    fmt.Println(v4, v30)
                    v44 := GenS5{F0: 4294967295, F1: 42, F2: (^(v30))}
                    _ = v44
                } else {
                    var v45 uint64 = 4294967295
                    fmt.Println(v30, v3)
                    fmt.Println(v3, v30, v30)
                    _ = v45
                }
                var v46 uint = 42
                for fi47 := 0; fi47 < 12; fi47++ {
                    if true {
                        if (1 == v46) {
                            if false {
                                var v48 int8 = int8(v46)
                                _ = v48
                            } else {
                                switch fi47 {
                                case 42:
                                    var v49 uint16 = 1
                                    _ = v49
                                default:
                                    var v50 float32 = v4
                                    _ = v50
                                }
                            }
                            if false {
                                switch "" {
                                case "":
                                    var v51 string = ""
                                    _ = v51
                                case "\n\t":
                                    switch ("hello" + "a") {
                                    case "foo bar":
                                        fmt.Println(fi43)
                                    default:
                                        if (-1 == (v5 | v5)) {
                                            var v52 float32 = 1e10
                                            var v53 int64 = v30
                                            var v54 uint = 1
                                            v5 += v5
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                        }
                                    }
                                default:
                                    delete(v2, (^(fi43)))
                                }
                                fmt.Println(v5, fi47, v5)
                            } else {
                                v55 := GenS9{F0: "hello", F1: 350, F2: (v5 >> (((v46 - (v46 >> ((0) & 63)))) & 63))}
                                switch (fi47 % (((fi47 * (fi43 - 32767))) | 1)) {
                                case 637:
                                    if false {
                                        v3 += v3
                                        fmt.Println(v3, v46, v5)
                                    }
                                }
                                fmt.Println(v30, v46)
                                _ = v55
                            }
                        }
                    }
                    if (!((1e10 + float64(3.14159)) == v3)) {
                        var v56 uint16 = 65535
                        switch "foo bar" {
                        case "foo bar":
                            var v57 string = ("foo bar" + "foo bar")
                            if true {
                                var v58 float64 = v3
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v5 = (-(v5))
                                _ = v58
                            } else {
                                fmt.Println(v30)
                                v2[fi47] = fi43
                            }
                            var v59 uint32 = 2
                            _ = v57
                            _ = v59
                        default:
                            v46 = (v46 - v46)
                        }
                        switch fi47 {
                        case -9223372036854775808:
                            v60 := v2[int(v56)]
                            _ = v60
                        case 4294967295:
                            v4 -= (v4 - 1.5)
                        case 0:
                            v61 := GenS4{F0: 0, F1: -128}
                            _ = v61
                        }
                        _ = v56
                    }
                    v62 := v2[127]
                    _ = fi47
                    _ = v62
                }
                _ = v46
            }
            v63 := GenS10{F0: 42}
            _ = len(v2)
            _ = fi43
            _ = v63
        }
    }
    switch -547 {
    case 2147483647:
        v4 = v4
    case 9223372036854775807:
        v64 := GenS1{F0: -319, F1: 4294967295, F2: 173}
        _ = v64
    default:
        v5 = 0
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v39
    _ = v4
    _ = v5
}
