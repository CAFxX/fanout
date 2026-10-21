package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int64; F1 bool; F2 int32 }
type GenS5 struct { F0 int64 }
func genFn0() int64 {
    v7 := GenS1{F0: 1, F1: (true || (1 >= -887)), F2: 42}
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 float32 = ((-2.25 - 0.1) * 1e10)
        v9 += float32((-2.25 + ((0.1 + 3.14159) * (1e10 * (3.14159 + 3.14159)))))
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 245, F1: true, F2: 0}
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 uint8 = 0
        _ = fi11
        _ = v12
    }
    _ = v10
    _ = v7
    return -226
}

func main() {
    v1 := []int{128, 127, -120}
    var v2 int16 = 32767
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]string)
    v6 := []string{(("a" + "foo bar") + "hello"), ("foo bar" + "foo bar"), "hello", "a"}
    var v13 int32 = -468
    if (0 < 107) {
        switch 32767 {
        case 87:
            v14 := GenS2{F0: (v2 >> ((2) & 63))}
            _ = v14
        default:
            fmt.Println(v13)
        }
    } else {
        var v15 bool = true
        _ = v15
    }
    var v16 uint16 = 644
    var v17 uint8 = 42
    v13 -= (v13 % ((v13) | 1))
    if (4294967295 <= 879) {
        v18 := GenS1{F0: v17, F1: true, F2: (^(v16))}
        _ = v18
    }
    if (!false) {
        v19 := GenS1{F0: (v17 / ((101) | 1)), F1: true, F2: (^(v16))}
        var v20 uint16 = v16
        _ = v19
        _ = v20
    } else {
        var v21 bool = (((0.0 * (1e10 / 1.5)) * 0.0) <= 3.14159)
        _ = len(v6)
        v22 := GenS3{F0: (!v21), F1: (v13 / ((v13) | 1))}
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 uint = 860
            if ((-2.25 + 0.1) >= float32(1.5)) {
                fmt.Println(v13)
                var v26 uint32 = 2147483648
                _ = v26
            } else {
                v3[32767] = "hello"
                v27 := GenS4{F0: 9223372036854775807, F1: false, F2: v13}
                v25 += v25
                if (true || (!false)) {
                    if (!false) {
                        switch "" {
                        case "foo bar":
                            switch fi23 {
                            case 127:
                                fmt.Println(v2, v16)
                            }
                        }
                    }
                    v17 = (v17 >> (((^(v25))) & 63))
                    var v28 uint64 = 2
                    _ = v28
                } else {
                    if (("hello" + ("hello" + "")) == "\n\t") {
                        var v29 float32 = -2.25
                        fmt.Println(v16, v16)
                        switch (fi24 >> ((v25) & 63)) {
                        case 648:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            if false {
                                var v30 bool = true
                                var v31 int32 = (v13 - -1)
                                _ = v30
                                _ = v31
                            }
                        }
                        v16 = (v16 | (v16 - (^(v16))))
                        v32 := GenS2{F0: (-(v2))}
                        delete(v3, (fi23 / (((fi23 | 0)) | 1)))
                        _ = v29
                        _ = v32
                    }
                    fmt.Println(v2, v13, v25)
                }
                _ = v27
            }
            _ = fi24
            _ = v25
        }
        var v33 int16 = v2
        var v34 int16 = 1000
        _ = fi23
        _ = v33
        _ = v34
    }
    if (446 == 9223372036854775807) {
        v16 = (^(v16))
        switch "\n\t" {
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    v16 += (v16 << ((uint(v17)) & 63))
    for fi35 := 0; fi35 < 13; fi35++ {
        if (175 >= 628) {
            v16 = v16
            var v36 bool = (fi35 >= fi35)
            _ = v36
        } else {
            if false {
                v37 := GenS5{F0: 42}
                switch (-(fi35)) {
                case 127:
                    _ = len(v1)
                case 32767:
                    switch fi35 {
                    case 422:
                        v13 -= (v13 + v13)
                    case 127:
                        var v38 uint64 = 2
                        _ = v38
                    case 958:
                        v17 += v17
                    default:
                        var v39 float64 = 0.0
                        _ = v39
                    }
                }
                _ = v37
            } else {
                if len(v1) > 0 {
                    v1[0] = fi35
                }
                if ((false && true) && false) {
                    var v40 float32 = 0.0
                    switch 4294967295 {
                    case -9223372036854775808:
                        v41 := GenS4{F0: -1, F1: true, F2: (v13 ^ 379)}
                        _ = v41
                    case -1:
                        var v42 uint64 = 650
                        _ = v42
                    default:
                        for fi43 := 0; fi43 < 20; fi43++ {
                            switch "\n\t" {
                            case "\n\t":
                                _ = len(v5)
                            case "foo bar":
                                var v44 float64 = float64(v40)
                                _ = v44
                            case "":
                                fmt.Println(v17, v2, fi43)
                            }
                            v1 = append(v1, fi35)
                            switch ("a" + "") {
                            case "":
                                var v45 uint32 = 2147483647
                                _ = v45
                            case "a":
                                v16 = v16
                            case "foo bar":
                                v46 := GenS2{F0: (v2 - (v2 >> ((175) & 63)))}
                                _ = v46
                            default:
                                v13 = 42
                            }
                            _ = fi43
                        }
                    }
                    var v47 uint = 255
                    _ = v40
                    _ = v47
                } else {
                    v16 += v16
                }
            }
            v3[(fi35 % (((fi35 / ((fi35) | 1))) | 1))] = ("" + "")
            switch "" {
            case "foo bar":
                v1 = append(v1, fi35)
            case "a":
                v17 -= v17
            default:
                fmt.Println(v2, v13)
            }
        }
        _ = fi35
    }
    v17 -= (^(v17))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
