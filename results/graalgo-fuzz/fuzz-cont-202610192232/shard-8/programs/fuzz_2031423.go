package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint; F1 uint16; F2 uint64 }
func genFn0() string {
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            if true {
                v8 := GenS1{F0: true, F1: 2, F2: 255}
                _ = v8
            }
            v9 := GenS2{F0: 1, F1: 42}
            _ = fi7
            _ = v9
        }
        var v10 int16 = 1000
        var v11 int16 = (v10 & (v10 >> ((18446744073709551615) & 63)))
        _ = fi6
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        for fi13 := 0; fi13 < 6; fi13++ {
            v14 := GenS2{F0: 371, F1: 1}
            fmt.Println(fi13)
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v15 := []int{2147483647, 127, 35, -112}
    _ = v15
    return "foo bar"
}
func genFn1(p0 uint32, p1 int8, p2 int) uint {
    var v16 uint8 = 1
    switch "" {
    case "foo bar":
        switch "hello" {
        case "a":
            p1 += (p1 + (p1 << ((255) & 63)))
        default:
            for fi17 := 0; fi17 < 18; fi17++ {
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 int64 = 9223372036854775807
                    v16 = v16
                    _ = fi18
                    _ = v19
                }
                _ = fi17
            }
        }
    default:
        var v20 uint64 = 42
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := []int{-447, (fi21 | (fi21 & fi21)), fi21, fi21}
        var v23 int8 = (-(p1))
        var v24 int32 = 42
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    if (9223372036854775807 > 1) {
        for fi25 := 0; fi25 < 17; fi25++ {
            fmt.Println(p2)
            for fi26 := 0; fi26 < 9; fi26++ {
                var v27 uint64 = 2
                if (42 != uint16(p0)) {
                    v28 := GenS2{F0: (v27 >> ((4294967295) & 63)), F1: 1}
                    v16 += 42
                    p2 += (p2 | p2)
                    _ = v28
                }
                var v29 int = p2
                _ = fi26
                _ = v27
                _ = v29
            }
            _ = fi25
        }
        if (false || (128 < p2)) {
            fmt.Println(p1, p0)
        }
        p2 = 32767
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 255
}

func main() {
    var v1 uint32 = 980
    v2 := []int{9223372036854775807, 4294967295}
    var v3 uint16 = 2
    var v4 int = 0
    v5 := make(map[uint32]string)
    fmt.Println(v1, v1)
    var v30 int = (v4 ^ v4)
    fmt.Println(v3, v30)
    for fi31 := 0; fi31 < 5; fi31++ {
        v32 := GenS3{F0: 0, F1: 1}
        _ = fi31
        _ = v32
    }
    _ = len(v2)
    v5[111] = ""
    v1 += 42
    v1 -= 479
    if false {
        var v33 int16 = 42
        var v34 int32 = 2147483647
        var v35 int = (-(v30))
        _ = v33
        _ = v34
        _ = v35
    } else {
        var v36 int = 439
        _ = v36
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        v38 := v5[(v1 >> ((547) & 63))]
        var v39 uint32 = 4294967295
        for fi40 := 0; fi40 < 13; fi40++ {
            var v41 uint32 = v39
            v42 := GenS4{F0: 693, F1: ("a" + "a")}
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = fi37
        _ = v38
        _ = v39
    }
    v4 = -2147483648
    v3 = (v3 & v3)
    fmt.Println(v1, v4, v1)
    if (!((v1 % (((v1 >> ((uint(v3)) & 63))) | 1)) == (v1 >> ((uint(v3)) & 63)))) {
        fmt.Println(v1, v1)
    }
    for fi43 := 0; fi43 < 12; fi43++ {
        var v44 float64 = 0.0
        _ = fi43
        _ = v44
    }
    for fi45 := 0; fi45 < 3; fi45++ {
        fmt.Println(v30, v3)
        switch (v30 & 128) {
        case 0:
            var v46 int16 = -32768
            _ = v46
        case -1:
            var v47 int64 = int64(v3)
            _ = v47
        case -9223372036854775808:
            var v48 uint16 = 764
            _ = v48
        default:
            if (!(127 > 42)) {
                v49 := GenS3{F0: -29, F1: 2}
                _ = v49
            } else {
                if true {
                    v50 := GenS2{F0: 18446744073709551615, F1: 0}
                    var v51 int64 = 42
                    for fi52 := 0; fi52 < 5; fi52++ {
                        if ((fi52 & fi45) >= v30) {
                            if len(v2) > 0 {
                                v2[0] = -949
                            }
                            v1 += (v1 ^ (v1 % (((v1 << ((2) & 63))) | 1)))
                            if (3.14159 == 1e10) {
                                var v53 uint32 = 4294967295
                                var v54 uint32 = (^(v1))
                                _ = v53
                                _ = v54
                            }
                            v51 = (v51 + (v51 + v51))
                        } else {
                            switch ("" + "\n\t") {
                            case "":
                                var v55 float64 = 0.0
                                _ = v55
                            case "\n\t":
                                v3 = 32767
                            case "a":
                                v5[(v1 * 2147483647)] = "hello"
                            }
                            v56 := GenS5{F0: 107, F1: (^(v3)), F2: 565}
                            var v57 float64 = 0.1
                            _ = v56
                            _ = v57
                        }
                        _ = fi52
                    }
                    _ = v50
                    _ = v51
                }
                v4 = (-(v30))
            }
        }
        v1 = (v1 % (((v1 * 684)) | 1))
        _ = fi45
    }
    var v58 float64 = 1e10
    switch (v30 + v4) {
    case 332:
        switch (^(v4)) {
        case 128:
            var v59 uint = 65535
            _ = v59
        }
    case 4:
        v4 += v4
    case -609:
        v60 := GenS3{F0: 42, F1: 1}
        _ = v60
    default:
        fmt.Println(v30)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v58
}
