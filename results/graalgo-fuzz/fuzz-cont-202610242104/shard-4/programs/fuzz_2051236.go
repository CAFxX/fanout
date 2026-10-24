package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 bool; F2 string }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 string }

func main() {
    var v1 int8 = -1
    v2 := []int{740, 32768, 1, 2147483647}
    var v3 uint8 = 169
    var v4 int = 9223372036854775807
    v5 := []int{v4}
    v6 := GenS1{F0: 9223372036854775807}
    v3 += v3
    var v7 int64 = -385
    v3 += 42
    if false {
        switch (("a" + "\n\t") + "\n\t") {
        case "a":
            var v8 float32 = float32((3.14159 / 1.5))
            _ = v8
        }
        if len(v5) > 0 {
            v5[0] = v4
        }
    } else {
        switch v4 {
        case -9223372036854775808:
            v9 := GenS2{F0: "hello"}
            _ = v9
        case 128:
            for fi10 := 0; fi10 < 1; fi10++ {
                if false {
                    for fi11 := 0; fi11 < 18; fi11++ {
                        fmt.Println(v4, fi10)
                        var v12 int = 1
                        _ = fi11
                        _ = v12
                    }
                    v3 -= v3
                    for fi13 := 0; fi13 < 1; fi13++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi13
                    }
                } else {
                    v3 = (v3 + 29)
                    v3 += v3
                    v14 := GenS3{F0: v7, F1: 637}
                    _ = v14
                }
                if true {
                    switch v4 {
                    case 32767:
                        if (!false) {
                            if (!((fi10 & fi10) <= -1)) {
                                for fi15 := 0; fi15 < 1; fi15++ {
                                    var v16 uint = 324
                                    _ = fi15
                                    _ = v16
                                }
                                v5 = append(v5, (^(v4)))
                                v17 := GenS1{F0: (v7 / ((v7) | 1))}
                                _ = v17
                            } else {
                                var v18 int8 = v1
                                _ = v18
                            }
                            if (1 >= 1) {
                                for fi19 := 0; fi19 < 9; fi19++ {
                                    switch fi19 {
                                    case 2147483647:
                                        var v20 uint8 = (v3 / (((^(v3))) | 1))
                                        _ = v20
                                    }
                                    _ = fi19
                                }
                            }
                        } else {
                            if (v3 > 1) {
                                var v21 int = int(v3)
                                _ = v21
                            }
                            if (42 < 32767) {
                                var v22 string = ""
                                fmt.Println(v1, v22, v1)
                                _ = v22
                            } else {
                                if (-2.25 < 0.1) {
                                    v23 := GenS4{F0: (v1 >> ((4294967295) & 63))}
                                    _ = v23
                                }
                                _ = len(v2)
                                var v24 uint16 = 580
                                _ = v24
                            }
                        }
                    case -1:
                        switch ("" + (("a" + (("hello" + "foo bar") + "hello")) + "foo bar")) {
                        case "":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case "\n\t":
                            fmt.Println(v7, fi10, v4)
                        default:
                            fmt.Println(v3, v7)
                        }
                    case -2147483648:
                        var v25 string = "hello"
                        _ = v25
                    default:
                        var v26 uint64 = 756
                        _ = v26
                    }
                    fmt.Println(v7, v3, fi10)
                    if true {
                        v3 = 0
                        for fi27 := 0; fi27 < 2; fi27++ {
                            if (!false) {
                                if len(v5) > 0 {
                                    v5[0] = fi27
                                }
                                v28 := GenS5{F0: -128}
                                v29 := GenS6{F0: "hello", F1: ((^(v1)) > 42), F2: ("\n\t" + "foo bar")}
                                _ = v28
                                _ = v29
                            }
                            switch "foo bar" {
                            case "foo bar":
                                switch "\n\t" {
                                case "hello":
                                    var v30 int8 = (v1 & -128)
                                    _ = v30
                                case "\n\t":
                                    v31 := GenS1{F0: (v7 | (v7 | -1))}
                                    _ = v31
                                }
                            case "\n\t":
                                v5 = append(v5, (v4 | v4))
                            default:
                                v32 := GenS7{F0: 128, F1: 184}
                                _ = v32
                            }
                            fmt.Println(fi27, v3, v1)
                            _ = fi27
                        }
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v7 -= (-(v7))
                            v1 += v1
                        }
                    }
                } else {
                    v4 = (fi10 + 127)
                }
                v1 = v1
                _ = fi10
            }
        case -296:
            var v33 int = v4
            _ = v33
        default:
            v34 := GenS8{F0: (("" + "\n\t") + "foo bar")}
            _ = v34
        }
        var v35 int = (v4 >> ((42) & 63))
        _ = v35
    }
    if ((v1 & -128) == (^(v1))) {
        var v36 bool = true
        if false {
            v7 = (v7 - (v7 << ((2) & 63)))
            if (-1 == v4) {
                var v37 float32 = (float32(3.14159) - 3.14159)
                _ = v37
            }
        }
        _ = v36
    }
    switch "hello" {
    case "a":
        for fi38 := 0; fi38 < 6; fi38++ {
            var v39 float32 = float32((1e10 - 3.14159))
            _ = fi38
            _ = v39
        }
    case "foo bar":
        var v40 float32 = (float32(-2.25) - (0.0 + (0.0 + 0.0)))
        _ = v40
    default:
        switch ("a" + "\n\t") {
        case "hello":
            var v41 float32 = 0.0
            _ = v41
        }
    }
    v42 := GenS5{F0: (^(v1))}
    var v43 int16 = 330
    v44 := GenS3{F0: (v7 % ((-1) | 1)), F1: (v43 & (v43 ^ v43))}
    v3 = v3
    if true {
        v5 = append(v5, (v4 % (((v4 * v4)) | 1)))
        v45 := GenS3{F0: -9223372036854775808, F1: v43}
        v5 = append(v5, (v4 << ((4294967295) & 63)))
        _ = v45
    } else {
        var v46 bool = true
        v2 = append(v2, v4)
        for fi47 := 0; fi47 < 13; fi47++ {
            v1 = v1
            var v48 uint32 = 42
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    v49 := GenS7{F0: (v4 ^ (v4 | 32767)), F1: 42}
    var v50 uint32 = 732
    v50 = (^(v50))
    var v51 int = (v4 / ((128) | 1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v7
}
