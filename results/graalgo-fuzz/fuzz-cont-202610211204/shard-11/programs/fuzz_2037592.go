package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int64; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 string; F2 int64 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 uint64 = 4294967295
        v5 += (v5 | v5)
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 int16 = 48
            v9 := GenS1{F0: 128, F1: 1}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v10 := []int{0}
    _ = v10
    return 819
}
func genFn1(p0 uint64, p1 int64, p2 uint64) int {
    p2 += 18446744073709551615
    p1 = genFn0()
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    v1 := []int{2147483647, 1, 127}
    var v2 int16 = 1
    var v3 int16 = v2
    v11 := GenS2{F0: -18}
    var v12 uint16 = 32767
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        switch ("a" + "hello") {
        case "foo bar":
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 int = (fi13 >> ((934) & 63))
                _ = fi14
                _ = v15
            }
        }
        if len(v1) > 0 {
            v1[0] = int(v2)
        }
        _ = fi13
    }
    v1 = append(v1, 329)
    if true {
        if ((255 > 0) || (((v12 != (v12 / (((v12 * 1)) | 1))) || (true && (97 == 255))) && false)) {
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 int = (fi16 % ((genFn1(865, -1, 10)) | 1))
                for fi18 := 0; fi18 < 16; fi18++ {
                    var v19 uint8 = 1
                    if (false || false) {
                        _ = len(v1)
                    }
                    _ = fi18
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
            var v20 bool = false
            var v21 uint8 = 42
            _ = v20
            _ = v21
        } else {
            if true {
                var v22 int8 = -1
                _ = v22
            } else {
                var v23 uint = 641
                if len(v1) > 0 {
                    v1[0] = -718
                }
                _ = v23
            }
            v3 -= v2
            var v24 int32 = -125
            _ = v24
        }
        var v25 float64 = 1e10
        var v26 int16 = -1
        _ = v25
        _ = v26
    } else {
        v2 = (^(v2))
    }
    fmt.Println(v2, v2)
    v27 := GenS3{F0: 813, F1: -1, F2: 100}
    v28 := GenS4{F0: v12, F1: 2, F2: 42}
    fmt.Println(v12, v3, v2)
    if true {
        if false {
            if true {
                switch -9223372036854775808 {
                case 2147483647:
                    v3 += 1
                case -9223372036854775808:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    var v29 string = "\n\t"
                    _ = v29
                }
            }
            v12 = v12
            switch -159 {
            case -571:
                switch 9223372036854775807 {
                case 652:
                    switch 32768 {
                    case 4294967295:
                        var v30 bool = false
                        _ = v30
                    }
                }
            case 546:
                var v31 uint64 = 877
                _ = v31
            case 789:
                v1 = append(v1, 785)
            }
        }
        for fi32 := 0; fi32 < 5; fi32++ {
            v1 = append(v1, genFn1(18446744073709551615, 1, 18446744073709551615))
            v33 := GenS3{F0: 65535, F1: -158, F2: 87}
            _ = fi32
            _ = v33
        }
    } else {
        if false {
            v3 = v2
            switch 1 {
            case -2147483648:
                var v34 uint8 = 158
                _ = v34
            default:
                if false {
                    for fi35 := 0; fi35 < 17; fi35++ {
                        v36 := GenS2{F0: 48}
                        for fi37 := 0; fi37 < 8; fi37++ {
                            v38 := GenS5{F0: 1, F1: -2147483648}
                            _ = fi37
                            _ = v38
                        }
                        if (128 <= fi35) {
                            if len(v1) > 0 {
                                v1[0] = fi35
                            }
                        }
                        _ = fi35
                        _ = v36
                    }
                }
            }
        }
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        var v40 int64 = 1
        if false {
            switch "hello" {
            case "hello":
                for fi41 := 0; fi41 < 19; fi41++ {
                    var v42 uint64 = 2
                    _ = fi41
                    _ = v42
                }
            case "\n\t":
                var v43 int32 = -1
                _ = v43
            }
            if len(v1) > 0 {
                v1[0] = genFn1(2, genFn0(), 2)
            }
            v44 := GenS6{F0: v12}
            _ = v44
        }
        _ = fi39
        _ = v40
    }
    for fi45 := 0; fi45 < 5; fi45++ {
        var v46 int64 = 0
        fmt.Println(v3)
        _ = fi45
        _ = v46
    }
    var v47 int64 = -1
    if ((!(false && (65535 >= 1))) && (uint(v12) < 1)) {
        fmt.Println(v12)
        if (false && true) {
            var v48 int32 = 278
            v49 := GenS7{F0: true, F1: ("\n\t" + ("foo bar" + ("foo bar" + ("hello" + "")))), F2: 1}
            _ = v48
            _ = v49
        }
        switch ("a" + ("\n\t" + "foo bar")) {
        case "a":
            var v50 string = "foo bar"
            _ = v50
        case "\n\t":
            _ = len(v1)
        case "":
            switch ("a" + "foo bar") {
            case "":
                v51 := GenS3{F0: 0, F1: v47, F2: 22}
                _ = v51
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        default:
            var v52 uint64 = 167
            _ = v52
        }
    } else {
        var v53 uint8 = 20
        _ = v53
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v47
}
