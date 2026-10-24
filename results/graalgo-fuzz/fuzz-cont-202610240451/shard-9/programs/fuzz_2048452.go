package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int32 }
type GenS2 struct { F0 bool; F1 int32; F2 int8 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint) uint64 {
    var v7 float32 = 3.14159
    v8 := []int{-2147483648, 32768}
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: "foo bar", F1: p0, F2: -676}
        var v11 int64 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 uint32, p1 int64) uint {
    switch 1 {
    case 32768:
        if (false && true) {
            p0 = p0
            for fi12 := 0; fi12 < 11; fi12++ {
                p1 = (p1 ^ (p1 / ((p1) | 1)))
                _ = fi12
            }
        } else {
            p1 -= -526
            for fi13 := 0; fi13 < 7; fi13++ {
                for fi14 := 0; fi14 < 10; fi14++ {
                    fmt.Println(p1, fi14)
                    _ = fi14
                }
                var v15 int64 = (p1 * (p1 | p1))
                _ = fi13
                _ = v15
            }
            fmt.Println(p1)
        }
    case 127:
        if (0.1 > 3.14159) {
            fmt.Println(p0, p0)
        }
    }
    p1 += 1
    if true {
        v16 := []int{128, -336, -9223372036854775808}
        v17 := GenS1{F0: "foo bar", F1: 18446744073709551615, F2: 2147483647}
        for fi18 := 0; fi18 < 9; fi18++ {
            if false {
                p1 += p1
                switch (fi18 & (fi18 ^ fi18)) {
                case 128:
                    var v19 float32 = -2.25
                    _ = v19
                }
                p0 = (p0 >> ((393) & 63))
            }
            switch "hello" {
            case "":
                v20 := GenS1{F0: ("\n\t" + ("" + ("" + ("hello" + ("a" + (("" + (("" + "hello") + (("hello" + "\n\t") + "hello"))) + ("" + "hello"))))))), F1: 4294967295, F2: -2147483648}
                _ = v20
            default:
                fmt.Println(p1)
            }
            _ = fi18
        }
        _ = v16
        _ = v17
    }
    v21 := []int{127, 564}
    _ = p0
    _ = p1
    _ = v21
    return 284
}

func main() {
    var v1 uint16 = 666
    v2 := []uint32{2147483648}
    v3 := []int{914}
    v4 := make(map[uint32]uint32)
    var v5 float32 = 1.5
    var v6 int8 = 127
    if (false || true) {
        var v22 float32 = v5
        for fi23 := 0; fi23 < 7; fi23++ {
            v1 += 716
            fmt.Println(v22, v1)
            _ = fi23
        }
        v1 = 2
        _ = v22
    }
    if true {
        var v24 int32 = 42
        _ = v24
    }
    switch "a" {
    case "foo bar":
        v6 += (-(v6))
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v25 int = 0
    for fi26 := 0; fi26 < 8; fi26++ {
        if false {
            var v27 uint16 = 65535
            v28 := GenS1{F0: "a", F1: 119, F2: 1}
            var v29 uint64 = 2
            _ = v27
            _ = v28
            _ = v29
        } else {
            var v30 uint16 = (v1 % (((v1 * v1)) | 1))
            _ = v30
        }
        var v31 uint16 = 42
        v32 := v4[0]
        _ = fi26
        _ = v31
        _ = v32
    }
    v6 = v6
    v33 := GenS1{F0: "a", F1: 0, F2: -1}
    _ = len(v3)
    v1 = (v1 % ((v1) | 1))
    v25 -= v25
    var v34 uint = uint(v1)
    v2 = append(v2, 4294967295)
    var v35 int16 = -1
    v2 = append(v2, 42)
    if false {
        v36 := GenS1{F0: "hello", F1: v34, F2: 2147483647}
        v37 := v4[720]
        _ = v36
        _ = v37
    } else {
        if false {
            if false {
                v38 := GenS1{F0: "foo bar", F1: 255, F2: 42}
                var v39 int = 32768
                if true {
                    v40 := v4[876]
                    v6 = 1
                    _ = v40
                } else {
                    var v41 uint = 1
                    var v42 int32 = 872
                    _ = v41
                    _ = v42
                }
                _ = v38
                _ = v39
            } else {
                if (false || (0 <= int64(v25))) {
                    var v43 uint64 = 4294967295
                    _ = v43
                } else {
                    for fi44 := 0; fi44 < 12; fi44++ {
                        v45 := GenS2{F0: (false && true), F1: -130, F2: 1}
                        v46 := v4[407]
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                    if true {
                        var v47 float64 = (3.14159 * 1e10)
                        if ((v47 <= float64(3.14159)) && false) {
                            fmt.Println(v34)
                            for fi48 := 0; fi48 < 13; fi48++ {
                                v49 := GenS3{F0: uint32(v1), F1: v35}
                                v6 += v6
                                switch (^(fi48)) {
                                case -2147483648:
                                    if true {
                                        var v50 uint64 = 342
                                        fmt.Println(v47)
                                        _ = v50
                                    }
                                default:
                                    delete(v4, 2)
                                }
                                _ = fi48
                                _ = v49
                            }
                        }
                        _ = v47
                    } else {
                        v51 := GenS4{F0: (!((v1 / ((v1) | 1)) != v1))}
                        for fi52 := 0; fi52 < 13; fi52++ {
                            fmt.Println(v34, v6, v25)
                            _ = fi52
                        }
                        for fi53 := 0; fi53 < 1; fi53++ {
                            v54 := GenS2{F0: false, F1: 296, F2: 0}
                            v55 := GenS3{F0: 2, F1: 850}
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                        _ = v51
                    }
                    var v56 uint64 = 0
                    _ = v56
                }
            }
            _ = len(v4)
        } else {
            delete(v4, 543)
        }
    }
    var v57 uint16 = v1
    if (0.0 != 3.14159) {
        for fi58 := 0; fi58 < 10; fi58++ {
            if (false || true) {
                for fi59 := 0; fi59 < 8; fi59++ {
                    v25 += (fi58 + (fi59 << ((65535) & 63)))
                    _ = fi59
                }
                if (18446744073709551615 != 4294967295) {
                    for fi60 := 0; fi60 < 18; fi60++ {
                        v61 := v4[2]
                        var v62 int32 = 215
                        _ = fi60
                        _ = v61
                        _ = v62
                    }
                    var v63 bool = (127 > 42)
                    _ = v63
                }
            } else {
                fmt.Println(v1, v6)
                delete(v4, 65535)
            }
            _ = fi58
        }
    }
    var v64 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v64
}
