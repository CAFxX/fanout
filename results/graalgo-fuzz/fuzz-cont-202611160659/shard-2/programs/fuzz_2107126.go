package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint; F2 int64 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8; F1 int16; F2 uint16 }
func genFn0(p0 uint, p1 uint8) string {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: p1, F1: -1}
        switch (fi4 ^ -209) {
        case 2:
            if true {
                p1 += p1
            }
        case 518:
            for fi6 := 0; fi6 < 10; fi6++ {
                p0 += 13
                _ = fi6
            }
        default:
            var v7 uint32 = 797
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    p0 -= p0
    fmt.Println(p0, p1)
    p1 += 0
    _ = p0
    _ = p1
    return "foo bar"
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]uint32)
    var v3 int64 = -564
    switch 127 {
    case -492:
        var v8 string = "foo bar"
        _ = v8
    default:
        fmt.Println(v3, v1)
    }
    var v9 uint16 = 286
    v10 := []int{}
    v11 := GenS2{F0: 42, F1: v3}
    fmt.Println(v9, v1, v9)
    for fi12 := 0; fi12 < 9; fi12++ {
        if false {
            var v13 float32 = 3.14159
            _ = v13
        }
        _ = fi12
    }
    if true {
        if (!((!(!((32767 >= -1) || (2 >= 4294967295)))) && false)) {
            var v14 uint64 = 18446744073709551615
            for fi15 := 0; fi15 < 8; fi15++ {
                if (false && (200 == 127)) {
                    v14 += v14
                    v16 := GenS3{F0: 255}
                    v3 -= v3
                    _ = v16
                }
                if (255 <= 65535) {
                    var v17 int = fi15
                    fmt.Println(v3, v17)
                    _ = len(v10)
                    _ = v17
                } else {
                    switch fi15 {
                    case 32768:
                        if (v9 != v9) {
                            if (!("hello" != "foo bar")) {
                                fmt.Println(v1)
                            } else {
                                _ = len(v2)
                                for fi18 := 0; fi18 < 12; fi18++ {
                                    v19 := GenS1{F0: uint8(v14), F1: 1}
                                    _ = fi18
                                    _ = v19
                                }
                                var v20 uint32 = v1
                                _ = v20
                            }
                            var v21 uint16 = v9
                            _ = v21
                        } else {
                            switch "foo bar" {
                            case "a":
                                v14 += v14
                            case "":
                                for fi22 := 0; fi22 < 20; fi22++ {
                                    v23 := GenS3{F0: 255}
                                    if (!true) {
                                        v24 := GenS2{F0: 0, F1: -1}
                                        v3 = 0
                                        _ = v24
                                    }
                                    fmt.Println(fi22)
                                    _ = fi22
                                    _ = v23
                                }
                            case "foo bar":
                                v2[fi15] = v1
                            }
                            fmt.Println(v9, v9)
                        }
                    case 42:
                        v14 -= v14
                    case 2147483647:
                        if (true || (1 >= 0)) {
                            v2[(fi15 ^ fi15)] = v1
                            if (((v1 >= (v1 >> ((2) & 63))) && true) && false) {
                                for fi25 := 0; fi25 < 14; fi25++ {
                                    v26 := GenS4{F0: (^(v9)), F1: 2, F2: (v3 ^ (v3 * v3))}
                                    _ = fi25
                                    _ = v26
                                }
                            }
                            v27 := GenS5{F0: 0, F1: 1}
                            _ = v27
                        }
                    }
                    switch "a" {
                    case "foo bar":
                        v3 = v3
                    case "hello":
                        for fi28 := 0; fi28 < 9; fi28++ {
                            switch -507 {
                            case -932:
                                if (false || (0.1 <= 1e10)) {
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    if false {
                                        v9 -= v9
                                    }
                                } else {
                                    var v29 uint16 = v9
                                    var v30 uint32 = (v1 >> ((785) & 63))
                                    v14 -= (v14 & v14)
                                    v29 -= v9
                                    switch (fi28 / ((fi15) | 1)) {
                                    case 32767:
                                        fmt.Println(v3, v14, v29)
                                    case -1:
                                        if true {
                                            fmt.Println(v1)
                                            switch ((("\n\t" + (("" + "foo bar") + "a")) + ("" + ("a" + "a"))) + "foo bar") {
                                            case "":
                                                v29 -= 589
                                            case "a":
                                                v31 := GenS2{F0: 65535, F1: (v3 ^ (-(v3)))}
                                                _ = v31
                                            default:
                                                delete(v2, fi15)
                                            }
                                            v32 := GenS1{F0: 128, F1: 0}
                                            _ = v32
                                        } else {
                                            v3 = 97
                                            fmt.Println(v3)
                                            v2[fi15] = (v30 | (v1 & v1))
                                        }
                                    default:
                                        var v33 string = "foo bar"
                                        _ = v33
                                    }
                                    _ = v29
                                    _ = v30
                                }
                            }
                            if (false || false) {
                                v34 := GenS3{F0: 127}
                                var v35 uint64 = v14
                                if false {
                                    if (572 < 2147483647) {
                                        if false {
                                            v35 -= 42
                                        } else {
                                            var v36 int32 = -2147483648
                                            if (21 > 1) {
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                                var v37 uint8 = 127
                                                _ = v37
                                            } else {
                                                v38 := GenS5{F0: (v36 & v36), F1: 732}
                                                _ = v38
                                            }
                                            v9 = v9
                                            _ = v36
                                        }
                                    }
                                }
                                _ = v34
                                _ = v35
                            }
                            _ = fi28
                        }
                    case "\n\t":
                        var v39 int8 = 98
                        _ = v39
                    }
                    if true {
                        v3 = -9223372036854775808
                    }
                }
                _ = fi15
            }
            _ = v14
        }
    }
    var v40 int8 = -43
    for fi41 := 0; fi41 < 18; fi41++ {
        if true {
            for fi42 := 0; fi42 < 12; fi42++ {
                if true {
                    var v43 int16 = 92
                    v43 = (-(v43))
                    _ = v43
                } else {
                    fmt.Println(v3)
                }
                v44 := GenS6{F0: 1, F1: -1}
                _ = fi42
                _ = v44
            }
        }
        _ = fi41
    }
    v45 := GenS7{F0: "hello"}
    for fi46 := 0; fi46 < 11; fi46++ {
        _ = len(v10)
        var v47 int16 = 0
        _ = fi46
        _ = v47
    }
    v48 := GenS8{F0: ((("foo bar" + "foo bar") + "hello") + "hello")}
    switch "\n\t" {
    case "":
        v49 := GenS9{F0: 1, F1: -32768, F2: (^(v9))}
        _ = v49
    case "foo bar":
        _ = len(v2)
    case "a":
        if true {
            var v50 float32 = 1e10
            switch "hello" {
            case "":
                fmt.Println(v3, v3, v3)
            case "hello":
                v2[1] = v1
            default:
                if len(v10) > 0 {
                    v10[0] = -9223372036854775808
                }
            }
            for fi51 := 0; fi51 < 10; fi51++ {
                if len(v10) > 0 {
                    v10[0] = fi51
                }
                v52 := v2[(fi51 ^ (fi51 << ((2) & 63)))]
                v1 = v52
                _ = fi51
                _ = v52
            }
            _ = v50
        }
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v40
    _ = v45
    _ = v48
    _ = v9
}
