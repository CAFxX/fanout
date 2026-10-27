package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint32; F1 string; F2 bool }

func main() {
    var v1 uint = 1
    v2 := []uint32{4294967295, 674, 65535}
    var v3 uint8 = 2
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 float32 = (((((3.14159 * -2.25) / 0.1) + ((0.1 * 3.14159) + -2.25)) + 1e10) + 0.1)
        var v6 int16 = -1
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        v3 = (v3 % (((v3 << ((42) & 63))) | 1))
        _ = fi7
    }
    v1 += v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v8 uint64 = 0
    var v9 uint8 = v3
    if (v3 < (v9 | 42)) {
        var v10 int = 32768
        v10 -= -9223372036854775808
        v8 = v8
        _ = v10
    }
    switch "foo bar" {
    case "":
        _ = len(v2)
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        v11 := GenS1{F0: (42 == -461), F1: true}
        _ = v11
    default:
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS1{F0: false, F1: ((fi12 / (((fi12 << ((65535) & 63))) | 1)) > (fi12 + fi12))}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi12
            _ = v13
        }
    }
    v3 += v9
    v1 = v1
    if (v3 >= (v9 & (v3 | v9))) {
        var v14 int64 = -552
        _ = v14
    }
    if false {
        _ = len(v2)
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v15 := GenS2{F0: (^(v1))}
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    switch "hello" {
    case "\n\t":
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 string = ""
            v18 := GenS3{F0: -1, F1: 271, F2: 65535}
            if true {
                for fi19 := 0; fi19 < 18; fi19++ {
                    var v20 uint64 = (v8 & (v8 << ((0) & 63)))
                    var v21 int64 = 1
                    if false {
                        v1 = 2
                        v20 -= 4294967295
                        v9 = (^(v9))
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                for fi22 := 0; fi22 < 14; fi22++ {
                    v17 = ((v17 + "foo bar") + (v17 + v17))
                    switch 32768 {
                    case 2147483647:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -9223372036854775808:
                        if false {
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                            var v23 int16 = -622
                            _ = v23
                        }
                    }
                    _ = fi22
                }
                _ = len(v2)
            }
            _ = fi16
            _ = v17
            _ = v18
        }
    case "":
        if ((497 >= 1) && (2147483647 < -1)) {
            v24 := GenS1{F0: false, F1: (2147483647 == -575)}
            _ = v24
        } else {
            if ((727 != 42) || true) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if false {
                    v3 = (^(v9))
                    for fi25 := 0; fi25 < 17; fi25++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        for fi26 := 0; fi26 < 2; fi26++ {
                            v27 := GenS4{F0: v3, F1: fi26, F2: 1}
                            var v28 uint8 = (v9 << ((v1) & 63))
                            fmt.Println(v1, v8, v9)
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        for fi29 := 0; fi29 < 14; fi29++ {
                            fmt.Println(fi25)
                            v30 := GenS5{F0: v8, F1: uint16(v1)}
                            _ = fi29
                            _ = v30
                        }
                        _ = fi25
                    }
                }
                for fi31 := 0; fi31 < 12; fi31++ {
                    v32 := GenS3{F0: -168, F1: 42, F2: 65535}
                    v9 += v9
                    fmt.Println(v1, v1)
                    _ = fi31
                    _ = v32
                }
            } else {
                switch 32768 {
                case 42:
                    var v33 float32 = 0.0
                    _ = v33
                case -716:
                    if (((v3 != 0) || (1000 >= -385)) || (677 >= -1)) {
                        v3 -= (v3 >> ((v1) & 63))
                    } else {
                        if len(v2) > 0 {
                            v2[0] = 0
                        }
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            if (int64(v1) > -9223372036854775808) {
                                v34 := GenS6{F0: 2147483647, F1: (v8 | (v8 - uint64(v1))), F2: v8}
                                if (false || (118 <= 100)) {
                                    var v35 int64 = 749
                                    fmt.Println(v8, v8, v3)
                                    _ = v35
                                }
                                v3 += 1
                                _ = v34
                            }
                            fmt.Println(v3)
                            var v36 int64 = 847
                            _ = v36
                        }
                    }
                case 2147483647:
                    for fi37 := 0; fi37 < 15; fi37++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v38 float32 = 1.5
                        if (1e10 < (v38 / 1e10)) {
                            switch -1 {
                            case -152:
                                v3 = v9
                            case 708:
                                fmt.Println(v1)
                            case 42:
                                var v39 uint16 = 32767
                                _ = v39
                            }
                            fmt.Println(v3)
                        }
                        _ = fi37
                        _ = v38
                    }
                }
                fmt.Println(v3, v3, v9)
            }
            _ = len(v2)
        }
    case "a":
        v2 = append(v2, 733)
    }
    switch (("a" + "\n\t") + ("" + "foo bar")) {
    case "a":
        fmt.Println(v8, v3)
    case "foo bar":
        if false {
            fmt.Println(v9, v3, v9)
            for fi40 := 0; fi40 < 18; fi40++ {
                if (!(3.14159 == (-2.25 * (1e10 * 1e10)))) {
                    v41 := GenS3{F0: 32767, F1: 1, F2: 11}
                    if (65535 >= 0) {
                        v42 := GenS5{F0: v8, F1: 32767}
                        if len(v2) > 0 {
                            v2[0] = 657
                        }
                        v2 = append(v2, 760)
                        _ = v42
                    } else {
                        v9 -= v3
                    }
                    _ = v41
                }
                _ = fi40
            }
        }
    }
    v43 := GenS7{F0: uint32(v8), F1: "\n\t", F2: false}
    v2 = append(v2, 2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v43
    _ = v8
    _ = v9
}
