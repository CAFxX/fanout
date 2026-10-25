package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint16; F1 int; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 bool; F1 int8; F2 bool }
type GenS6 struct { F0 uint; F1 int32; F2 int64 }
type GenS7 struct { F0 uint8; F1 int8; F2 int16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int; F1 string }
func genFn0(p0 int32, p1 uint, p2 uint32) uint64 {
    p2 += p2
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 string = ""
        p1 = p1
        v8 := []int{(fi6 & fi6)}
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1() int64 {
    v9 := GenS1{F0: 65535}
    v10 := GenS1{F0: 210}
    switch "\n\t" {
    case "a":
        v11 := GenS1{F0: 255}
        _ = v11
    case "foo bar":
        switch -125 {
        case 127:
            switch 9223372036854775807 {
            case -9223372036854775808:
                var v12 uint = 65535
                _ = v12
            case 0:
                v13 := GenS2{F0: 84, F1: 4294967295, F2: 525}
                _ = v13
            default:
                v14 := GenS3{F0: 102, F1: -9223372036854775808, F2: 960}
                _ = v14
            }
        case -179:
            var v15 int64 = 94
            _ = v15
        case 0:
            var v16 float32 = 1e10
            _ = v16
        default:
            switch "a" {
            case "\n\t":
                v17 := []int{0, 223, 4294967295, -1}
                _ = v17
            case "hello":
                var v18 uint32 = 65535
                _ = v18
            }
        }
    default:
        v19 := []int{32767, -371, -20}
        _ = v19
    }
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 bool = true
    v2 := []uint32{42}
    var v3 bool = (false || false)
    var v4 int32 = 42
    v5 := make(map[int]string)
    _ = len(v2)
    fmt.Println(v4)
    var v20 bool = (2 > 18446744073709551615)
    fmt.Println(v20, v4)
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    _ = len(v5)
    switch 349 {
    case 42:
        v21 := GenS2{F0: 2147483648, F1: 673, F2: -9223372036854775808}
        _ = v21
    case 2147483647:
        var v22 int32 = (-(v4))
        _ = v22
    }
    fmt.Println(v4, v3)
    switch 42 {
    case -2147483648:
        v23 := GenS4{F0: 255, F1: 0}
        _ = v23
    case -1:
        var v24 int32 = (v4 >> ((42) & 63))
        _ = v24
    default:
        var v25 uint8 = 0
        _ = v25
    }
    v5[379] = ""
    if true {
        v3 = true
        v26 := GenS5{F0: true, F1: 82, F2: true}
        v27 := GenS6{F0: 4294967295, F1: v4, F2: genFn1()}
        _ = v26
        _ = v27
    } else {
        if ((127 >= -43) && (!(!(!false)))) {
            v20 = true
        } else {
            if (true && false) {
                v3 = (!v3)
                switch 128 {
                case -919:
                    for fi28 := 0; fi28 < 4; fi28++ {
                        for fi29 := 0; fi29 < 15; fi29++ {
                            if ((fi28 | fi29) == fi28) {
                                fmt.Println(fi28, fi29, v4)
                                v1 = ((fi29 + (fi29 ^ (fi29 % ((int(v4)) | 1)))) == fi29)
                                if v20 {
                                    if v1 {
                                        fmt.Println(fi28, v1)
                                        v1 = (42 <= 9223372036854775807)
                                        v3 = v20
                                        var v30 int8 = -1
                                        v30 -= v30
                                        _ = v30
                                    }
                                    v31 := GenS7{F0: 255, F1: 127, F2: 1000}
                                    v4 -= (v4 / (((v4 & v4)) | 1))
                                    _ = v31
                                } else {
                                    if v1 {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    var v32 int64 = 42
                                    v4 = -2147483648
                                    _ = v32
                                }
                                v33 := GenS2{F0: 65535, F1: 336, F2: 1}
                                _ = v33
                            } else {
                                _ = len(v5)
                                if (255 <= 144) {
                                    var v34 uint = 1
                                    v1 = (4294967295 <= 4294967295)
                                    _ = v34
                                }
                            }
                            _ = fi29
                        }
                        for fi35 := 0; fi35 < 11; fi35++ {
                            var v36 int8 = 100
                            fmt.Println(v4, fi35)
                            _ = fi35
                            _ = v36
                        }
                        _ = fi28
                    }
                case -494:
                    fmt.Println(v3, v1)
                case 4294967295:
                    v37 := GenS3{F0: 255, F1: -9223372036854775808, F2: 65535}
                    _ = v37
                default:
                    var v38 int64 = -9223372036854775808
                    _ = v38
                }
            } else {
                var v39 uint = 1
                v40 := v5[127]
                var v41 bool = v1
                _ = v39
                _ = v40
                _ = v41
            }
            v4 = v4
            for fi42 := 0; fi42 < 15; fi42++ {
                var v43 int8 = 33
                var v44 bool = ((v4 ^ v4) <= 1)
                _ = fi42
                _ = v43
                _ = v44
            }
        }
        switch 76 {
        case 2147483647:
            for fi45 := 0; fi45 < 1; fi45++ {
                v46 := GenS8{F0: ("\n\t" + (("foo bar" + "a") + "foo bar"))}
                _ = fi45
                _ = v46
            }
        case 32767:
            var v47 string = "hello"
            _ = v47
        case 128:
            switch 127 {
            case 1:
                fmt.Println(v1)
            case -9223372036854775808:
                switch "foo bar" {
                case "a":
                    var v48 uint = 4294967295
                    _ = v48
                case "hello":
                    _ = len(v2)
                }
            default:
                v49 := GenS9{F0: 32767, F1: "hello"}
                _ = v49
            }
        }
        if (0.1 > -2.25) {
            v50 := GenS5{F0: (100 != -1), F1: 1, F2: (-2.25 >= (float64(1e10) - (0.0 + 1.5)))}
            _ = v50
        }
    }
    v51 := v5[-2147483648]
    v4 = (-(v4))
    v3 = (65535 != 1)
    for fi52 := 0; fi52 < 6; fi52++ {
        switch fi52 {
        case -9223372036854775808:
            var v53 int64 = -946
            _ = v53
        }
        _ = fi52
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
