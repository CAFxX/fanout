package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint; F2 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8; F1 int64; F2 bool }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int }
func genFn0(p0 int8, p1 int8) uint8 {
    var v5 bool = ((-493 != 1000) || (!("hello" >= "a")))
    if v5 {
        v6 := []int{1, -175}
        if len(v6) > 0 {
            v6[0] = 128
        }
        v7 := GenS1{F0: 1, F1: ("hello" + "hello"), F2: -100}
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 string = "\n\t"
            var v10 uint64 = 0
            var v11 string = "hello"
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        if (127 >= 1) {
            for fi12 := 0; fi12 < 3; fi12++ {
                for fi13 := 0; fi13 < 12; fi13++ {
                    fmt.Println(p0, p0)
                    if (127 == 128) {
                        if true {
                            var v14 uint8 = 122
                            _ = v14
                        }
                        var v15 int = fi12
                        v16 := GenS1{F0: 1, F1: ("foo bar" + ("" + "foo bar")), F2: -1}
                        _ = v15
                        _ = v16
                    } else {
                        var v17 int16 = 1000
                        _ = v17
                    }
                    v18 := []int{(fi13 >> ((65535) & 63))}
                    _ = fi13
                    _ = v18
                }
                var v19 int = 0
                var v20 int8 = p0
                _ = fi12
                _ = v19
                _ = v20
            }
            var v21 uint64 = 0
            _ = v21
        } else {
            for fi22 := 0; fi22 < 11; fi22++ {
                v23 := []int{(fi22 / ((fi22) | 1)), (fi22 & fi22), (^(fi22)), (-(fi22))}
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                _ = fi22
                _ = v23
            }
            var v24 uint32 = 2
            fmt.Println(v5)
            _ = v24
        }
        p0 = p0
    }
    v25 := []int{998, 9223372036854775807, 1, 42}
    p1 += (p1 * p0)
    _ = p0
    _ = p1
    _ = v25
    _ = v5
    return 127
}
func genFn1(p0 bool) string {
    if false {
        var v26 int = -965
        v27 := []int{(v26 % (((-(v26))) | 1)), -1, v26}
        var v28 uint = 65535
        _ = v26
        _ = v27
        _ = v28
    }
    fmt.Println(p0)
    v29 := GenS1{F0: 2147483647, F1: "\n\t", F2: 42}
    p0 = p0
    p0 = (false && p0)
    _ = p0
    _ = v29
    return ("a" + "")
}

func main() {
    var v1 uint64 = 58
    var v2 int8 = -128
    var v3 int16 = 42
    v4 := []int{-1, 4294967295, -328, 42}
    v4 = append(v4, 9223372036854775807)
    if true {
        switch (("" + "\n\t") + (("a" + "") + (("" + "\n\t") + ("hello" + "")))) {
        case "foo bar":
            fmt.Println(v2)
        case "hello":
            var v30 float32 = 3.14159
            _ = v30
        case "":
            var v31 uint16 = 586
            _ = v31
        }
    } else {
        v32 := GenS1{F0: uint32(v1), F1: "a", F2: (v3 | 42)}
        _ = v32
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        v1 = v1
        _ = fi33
    }
    v34 := GenS2{F0: 1}
    var v35 uint8 = 204
    switch "hello" {
    case "foo bar":
        v35 = v35
    case "a":
        var v36 float32 = 0.0
        _ = v36
    case "":
        v37 := GenS1{F0: 909, F1: "hello", F2: 42}
        _ = v37
    default:
        if false {
            if false {
                for fi38 := 0; fi38 < 2; fi38++ {
                    v1 -= 1
                    v39 := GenS3{F0: 743, F1: 2147483647}
                    _ = fi38
                    _ = v39
                }
                fmt.Println(v35)
                v35 = v35
            }
            fmt.Println(v3)
            v40 := GenS2{F0: (v1 / ((366) | 1))}
            _ = v40
        } else {
            for fi41 := 0; fi41 < 11; fi41++ {
                var v42 float64 = 0.0
                v2 -= v2
                _ = fi41
                _ = v42
            }
            if (!(2147483647 < 2147483648)) {
                _ = len(v4)
                if ((v3 << ((65535) & 63)) == v3) {
                    v35 += (v35 - (v35 & v35))
                    var v43 bool = (3.14159 < 0.0)
                    _ = v43
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(v2, v1, v2)
                }
            }
            v44 := GenS4{F0: 65535, F1: 1, F2: 2147483648}
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        for fi46 := 0; fi46 < 8; fi46++ {
            v2 = (v2 | v2)
            if true {
                if (true || true) {
                    v1 -= v1
                    var v47 bool = true
                    fmt.Println(v1, v35, v3)
                    _ = v47
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                var v48 int8 = (v2 % ((v2) | 1))
                v4 = append(v4, (fi46 * fi45))
                v3 = v3
                _ = v48
            }
            v49 := GenS2{F0: v1}
            if false {
                fmt.Println(fi46, fi45)
                v50 := GenS5{F0: 42}
                _ = v50
            }
            _ = fi46
            _ = v49
        }
        v35 = 2
        var v51 uint64 = (v1 * v1)
        _ = fi45
        _ = v51
    }
    var v52 uint64 = v1
    for fi53 := 0; fi53 < 5; fi53++ {
        v2 = 111
        _ = fi53
    }
    var v54 int16 = (v3 * -401)
    for fi55 := 0; fi55 < 11; fi55++ {
        if false {
            var v56 uint8 = (v35 | 131)
            var v57 uint16 = 1
            _ = v56
            _ = v57
        } else {
            for fi58 := 0; fi58 < 20; fi58++ {
                v59 := GenS6{F0: (^(v35)), F1: -9223372036854775808, F2: true}
                fmt.Println(v54)
                v4 = append(v4, 32767)
                var v60 int16 = (v3 << ((65535) & 63))
                _ = fi58
                _ = v59
                _ = v60
            }
        }
        v52 = v52
        if (false || false) {
            v52 += (v52 * (^(v1)))
            var v61 int64 = 42
            _ = v61
        }
        _ = fi55
    }
    var v62 bool = (true && false)
    var v63 float64 = float64((3.14159 - (1.5 + (((3.14159 - 0.1) / -2.25) + 0.0))))
    fmt.Println(v1, v54, v2)
    v52 += 2
    switch 32767 {
    case 32768:
        _ = len(v4)
    case 123:
        if len(v4) > 0 {
            v4[0] = -701
        }
    default:
        if ("foo bar" <= "hello") {
            v64 := GenS7{F0: v35, F1: -677}
            v52 += (v52 & (v1 << ((468) & 63)))
            _ = v64
        } else {
            v65 := GenS8{F0: int16(v2)}
            for fi66 := 0; fi66 < 7; fi66++ {
                for fi67 := 0; fi67 < 1; fi67++ {
                    if v62 {
                        var v68 int16 = (-(v3))
                        if true {
                            var v69 uint64 = (v1 * v52)
                            fmt.Println(fi66)
                            _ = v69
                        } else {
                            switch ("foo bar" + ("" + "hello")) {
                            case "foo bar":
                                switch "foo bar" {
                                case "hello":
                                    if len(v4) > 0 {
                                        v4[0] = 405
                                    }
                                    v3 += int16(fi66)
                                case "\n\t":
                                    switch "" {
                                    case "\n\t":
                                        var v70 int = fi67
                                        _ = v70
                                    default:
                                        var v71 int64 = -951
                                        _ = v71
                                    }
                                case "foo bar":
                                    v68 += (v68 % (((v54 & 157)) | 1))
                                default:
                                    var v72 uint8 = (v35 | (^(v35)))
                                    _ = v72
                                }
                            }
                            v4 = append(v4, fi66)
                            var v73 float32 = float32(float64(0.0))
                            var v74 uint8 = v35
                            fmt.Println(v62, v74, v2)
                            fmt.Println(v73)
                            v75 := GenS9{F0: fi67}
                            _ = v73
                            _ = v74
                            _ = v75
                        }
                        _ = v68
                    }
                    var v76 int16 = int16(v2)
                    _ = fi67
                    _ = v76
                }
                var v77 int8 = 0
                _ = fi66
                _ = v77
            }
            var v78 int8 = v2
            _ = v65
            _ = v78
        }
    }
    var v79 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v52
    _ = v54
    _ = v62
    _ = v63
    _ = v79
}
