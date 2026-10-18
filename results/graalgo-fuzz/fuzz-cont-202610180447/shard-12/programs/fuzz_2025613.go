package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint8; F1 string; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 int64; F1 uint32; F2 bool }
func genFn0(p0 int64, p1 int) uint8 {
    fmt.Println(p1)
    if (2 == 42) {
        p1 += 399
    }
    _ = p0
    _ = p1
    return 127
}
func genFn1(p0 uint16) string {
    switch "" {
    case "\n\t":
        p0 = 2
    }
    p0 = (p0 | 65535)
    _ = p0
    return "\n\t"
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]uint32)
    var v3 bool = true
    v4 := []int{9223372036854775807}
    v4 = append(v4, -421)
    if (65535 <= 255) {
        fmt.Println(v1)
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := GenS1{F0: (fi5 | 32768), F1: 2, F2: 802}
            v7 := GenS1{F0: fi5, F1: 65535, F2: 0}
            for fi8 := 0; fi8 < 5; fi8++ {
                fmt.Println(fi5, v1)
                v1 = v1
                _ = len(v4)
                var v9 int8 = 0
                fmt.Println(v9)
                if ((("a" + v1) + (v1 + (v1 + "foo bar"))) >= v1) {
                    v1 = v1
                    var v10 int16 = 532
                    _ = v10
                } else {
                    var v11 float64 = 1e10
                    switch 576 {
                    case 405:
                        fmt.Println(fi8, fi8, v11)
                    case 127:
                        if (true && v3) {
                            if len(v4) > 0 {
                                v4[0] = 128
                            }
                            var v12 uint = 255
                            v11 -= float64(1.5)
                            _ = v12
                        } else {
                            if (fi5 != fi5) {
                                if false {
                                    v13 := GenS1{F0: (fi8 % ((fi8) | 1)), F1: 255, F2: 0}
                                    var v14 int8 = 127
                                    if false {
                                        _ = len(v4)
                                        var v15 int16 = 1000
                                        _ = v15
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = fi5
                                        }
                                    }
                                    _ = v13
                                    _ = v14
                                }
                                var v16 uint32 = 0
                                switch "\n\t" {
                                case "":
                                    v2[85] = (v16 / ((2147483648) | 1))
                                case "hello":
                                    switch (fi8 ^ (fi5 ^ fi5)) {
                                    case 42:
                                        v11 -= v11
                                    case -485:
                                        var v17 uint8 = 255
                                        _ = v17
                                    case 992:
                                        v9 += (v9 | (v9 ^ 1))
                                    default:
                                        var v18 uint = 0
                                        _ = v18
                                    }
                                }
                                if (true && (true && v3)) {
                                    v19 := GenS2{F0: genFn0(9223372036854775807, fi8), F1: "hello", F2: (v9 * -128)}
                                    var v20 uint = 65535
                                    v1 = v1
                                    _ = v19
                                    _ = v20
                                } else {
                                    v2[(v16 >> ((728) & 63))] = v16
                                    var v21 int32 = 42
                                    fmt.Println(v3, v11, v16)
                                    _ = v21
                                }
                                var v22 uint16 = 882
                                _ = v16
                                _ = v22
                            } else {
                                v23 := GenS3{F0: "a"}
                                v2[4294967295] = 0
                                v24 := GenS2{F0: 0, F1: (v1 + v1), F2: 0}
                                _ = v23
                                _ = v24
                            }
                            v9 += (v9 + v9)
                            var v25 int = fi5
                            _ = v25
                        }
                    }
                    v26 := GenS3{F0: (v1 + v1)}
                    fmt.Println(fi8)
                    v27 := GenS1{F0: -9223372036854775808, F1: 0, F2: 592}
                    var v28 uint64 = 0
                    var v29 uint8 = 255
                    _ = v11
                    _ = v26
                    _ = v27
                    _ = v28
                    _ = v29
                }
                _ = fi8
                _ = v9
            }
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    var v30 int64 = -1
    var v31 uint32 = 2147483647
    var v32 uint = 847
    if v3 {
        if (v3 || v3) {
            var v33 uint = v32
            if (127 > genFn0(v30, -2147483648)) {
                _ = len(v2)
                fmt.Println(v3, v1, v1)
            } else {
                var v34 float32 = 1e10
                _ = v34
            }
            v1 = "\n\t"
            _ = v33
        }
        switch ("\n\t" + "hello") {
        case "\n\t":
            v35 := v2[42]
            _ = v35
        }
    }
    v30 = -709
    for fi36 := 0; fi36 < 8; fi36++ {
        v30 -= -1
        _ = fi36
    }
    if (42 <= 2) {
        switch (("foo bar" + (((v1 + "foo bar") + v1) + "\n\t")) + "") {
        case "foo bar":
            for fi37 := 0; fi37 < 5; fi37++ {
                delete(v2, v31)
                var v38 uint8 = uint8(v31)
                switch (((v1 + v1) + v1) + "\n\t") {
                case "hello":
                    v31 += (v31 % ((939) | 1))
                }
                _ = fi37
                _ = v38
            }
        case "":
            var v39 int64 = v30
            _ = v39
        default:
            v30 = 0
        }
        v40 := GenS4{F0: 128}
        for fi41 := 0; fi41 < 15; fi41++ {
            if (v31 > (v31 + 42)) {
                if (v30 <= (v30 + (v30 ^ v30))) {
                    fmt.Println(v32, v32, v1)
                    v42 := GenS5{F0: 0}
                    v31 += (v31 << ((v32) & 63))
                    _ = v42
                }
                switch v1 {
                case "foo bar":
                    v43 := GenS3{F0: v1}
                    _ = v43
                case "hello":
                    var v44 uint8 = 42
                    _ = v44
                default:
                    if v3 {
                        var v45 uint32 = (v31 & (v31 >> ((v32) & 63)))
                        fmt.Println(v3)
                        _ = v45
                    } else {
                        fmt.Println(v1, v31)
                        if (true || (v3 && v3)) {
                            v31 = (v31 + (v31 / ((2147483647) | 1)))
                            switch fi41 {
                            case 152:
                                v1 = (v1 + "a")
                            case 4294967295:
                                v1 = v1
                            default:
                                v46 := GenS6{F0: -1, F1: 32767}
                                _ = v46
                            }
                            if v3 {
                                var v47 uint64 = 421
                                v30 = (v30 + v30)
                                _ = v47
                            } else {
                                var v48 uint16 = 0
                                fmt.Println(v1, fi41)
                                fmt.Println(v48, v31, v3)
                                _ = v48
                            }
                        }
                    }
                }
            }
            for fi49 := 0; fi49 < 5; fi49++ {
                v3 = v3
                var v50 float32 = 1e10
                v51 := v2[(v31 + (v31 / (((v31 ^ 2147483647)) | 1)))]
                _ = fi49
                _ = v50
                _ = v51
            }
            _ = fi41
        }
        _ = v40
    }
    if true {
        var v52 uint8 = 117
        _ = v52
    }
    v3 = (v30 < v30)
    if (((true && true) || (-1 <= 2147483647)) || ((v30 + v30) != (v30 >> (((v32 & (v32 * v32))) & 63)))) {
        v53 := GenS7{F0: -155, F1: v31, F2: false}
        _ = v53
    } else {
        for fi54 := 0; fi54 < 16; fi54++ {
            v1 = v1
            _ = fi54
        }
        var v55 int32 = 700
        _ = v55
    }
    fmt.Println(v30, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
