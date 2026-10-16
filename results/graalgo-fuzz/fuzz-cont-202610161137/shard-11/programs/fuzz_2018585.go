package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int64; F2 int64 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 float64 = float64((3.14159 - float32(-2.25)))
    v2 := make(map[uint32]string)
    var v3 uint = 18446744073709551615
    v4 := make(map[string]string)
    if false {
        v5 := GenS1{F0: 42, F1: 4294967295}
        switch 128 {
        case 127:
            if false {
                fmt.Println(v3)
                v6 := GenS1{F0: 65535, F1: 458}
                _ = v6
            } else {
                v1 = 1e10
                fmt.Println(v1)
                for fi7 := 0; fi7 < 12; fi7++ {
                    if (false && false) {
                        v3 = v3
                        var v8 int = 42
                        v9 := GenS1{F0: 0, F1: 2147483647}
                        _ = v8
                        _ = v9
                    } else {
                        var v10 string = "a"
                        _ = v10
                    }
                    var v11 int32 = -43
                    _ = fi7
                    _ = v11
                }
            }
        case -259:
            v1 -= ((1e10 * (v1 * v1)) / float64(-2.25))
        case 0:
            switch "hello" {
            case "\n\t":
                var v12 int64 = -9223372036854775808
                _ = v12
            case "hello":
                v3 = (v3 * (^(v3)))
            case "foo bar":
                var v13 uint = v3
                _ = v13
            default:
                v14 := GenS2{F0: 0, F1: (v3 ^ v3), F2: 1}
                _ = v14
            }
        }
        if false {
            fmt.Println(v3, v1)
        } else {
            if true {
                v15 := GenS2{F0: 2, F1: (v3 & 1), F2: 127}
                if false {
                    if (!(!(23 <= 0))) {
                        var v16 int16 = -32768
                        v2[2147483648] = "hello"
                        fmt.Println(v1, v3)
                        _ = v16
                    }
                    v17 := GenS1{F0: 366, F1: 0}
                    for fi18 := 0; fi18 < 6; fi18++ {
                        v19 := []int{-86, fi18, (fi18 % (((fi18 - fi18)) | 1))}
                        switch (fi18 ^ 679) {
                        case 127:
                            var v20 bool = ((!(1.5 > 0.0)) && true)
                            _ = v20
                        default:
                            v21 := GenS1{F0: 450, F1: 0}
                            _ = v21
                        }
                        _ = fi18
                        _ = v19
                    }
                    _ = v17
                }
                v22 := GenS1{F0: 32767, F1: 65535}
                _ = v15
                _ = v22
            }
            if (104 > 115) {
                v3 += v3
            } else {
                switch ("foo bar" + "") {
                case "":
                    var v23 uint8 = 2
                    _ = v23
                case "foo bar":
                    v3 = 2
                }
                var v24 uint64 = 0
                v1 += ((0.1 / (-2.25 / v1)) - v1)
                _ = v24
            }
        }
        _ = v5
    }
    if (false || true) {
        v25 := v2[2147483647]
        var v26 float32 = 1e10
        fmt.Println(v3, v3, v26)
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        var v28 uint = (^(v3))
        v1 -= v1
        _ = fi27
        _ = v28
    }
    v1 -= 1e10
    var v29 string = ""
    v30 := []int{}
    var v31 uint = (v3 >> (((^(v3))) & 63))
    v32 := GenS3{F0: 0, F1: -9223372036854775808, F2: int64(v3)}
    switch 0 {
    case 910:
        for fi33 := 0; fi33 < 3; fi33++ {
            v34 := GenS4{F0: -408}
            if true {
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(fi33, v31, v3)
                }
                for fi35 := 0; fi35 < 16; fi35++ {
                    fmt.Println(v29, v1)
                    if (!true) {
                        switch fi35 {
                        case 9223372036854775807:
                            if len(v30) > 0 {
                                v30[0] = 1
                            }
                        case -365:
                            if len(v30) >= 2 {
                                v30 = v30[1:len(v30)]
                            }
                        case -9223372036854775808:
                            if len(v30) > 0 {
                                v30[0] = (^(fi33))
                            }
                        }
                        var v36 uint32 = 2
                        switch v29 {
                        case "hello":
                            switch -1 {
                            case 233:
                                _ = len(v30)
                            default:
                                if len(v30) >= 2 {
                                    v30 = v30[1:len(v30)]
                                }
                            }
                        default:
                            if (!true) {
                                var v37 uint16 = 255
                                _ = v37
                            } else {
                                if (int16(fi35) > -1) {
                                    v38 := v4[("hello" + (v29 + v29))]
                                    _ = v38
                                }
                            }
                            v3 += 255
                            v2[(v36 * 0)] = v29
                            if (v3 == v31) {
                                v36 = v36
                            }
                        }
                        _ = v36
                    }
                    var v39 uint32 = 2
                    if (255 <= 1) {
                        delete(v2, 717)
                        v40 := v4[""]
                        _ = v40
                    } else {
                        if len(v30) > 0 {
                            v30[0] = (-(fi35))
                        }
                    }
                    _ = fi35
                    _ = v39
                }
            }
            _ = fi33
            _ = v34
        }
    case 42:
        v41 := v4[""]
        _ = v41
    default:
        _ = len(v30)
    }
    v31 -= 0
    v4[v29] = v29
    v42 := GenS3{F0: 1, F1: 9223372036854775807, F2: 214}
    var v43 bool = (false || true)
    if (1 == -1) {
        for fi44 := 0; fi44 < 6; fi44++ {
            if ((false || true) || v43) {
                if (true || ("hello" != v29)) {
                    if v43 {
                        v45 := GenS2{F0: 255, F1: 4294967295, F2: 128}
                        _ = v45
                    } else {
                        v31 += (^(v3))
                    }
                    var v46 uint64 = 0
                    _ = v46
                } else {
                    var v47 int32 = 139
                    if len(v30) >= 2 {
                        v30 = v30[1:len(v30)]
                    }
                    if len(v30) >= 2 {
                        v30 = v30[1:len(v30)]
                    }
                    _ = v47
                }
            } else {
                if (v43 || ((!v43) || (true && (490 == 1)))) {
                    v48 := GenS1{F0: 152, F1: 594}
                    var v49 uint32 = 2
                    switch fi44 {
                    case 9223372036854775807:
                        v49 = 2147483647
                    case 32768:
                        var v50 float64 = (v1 / ((float64(float32(((v1 * 0.0) - (1.5 + v1)))) * float64((3.14159 * 3.14159))) * 0.1))
                        _ = v50
                    }
                    _ = v48
                    _ = v49
                } else {
                    v3 -= v31
                }
            }
            _ = fi44
        }
        for fi51 := 0; fi51 < 5; fi51++ {
            v4["\n\t"] = v29
            for fi52 := 0; fi52 < 15; fi52++ {
                v53 := GenS1{F0: 42, F1: 4294967295}
                _ = fi52
                _ = v53
            }
            _ = fi51
        }
        fmt.Println(v3, v1)
    }
    v54 := GenS3{F0: 369, F1: 1, F2: 42}
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v42
    _ = v43
    _ = v54
}
