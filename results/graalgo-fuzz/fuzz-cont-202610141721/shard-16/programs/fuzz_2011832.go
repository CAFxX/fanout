package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 bool }
type GenS2 struct { F0 uint; F1 bool; F2 uint }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int; F1 uint }

func main() {
    v1 := []int{322, 2147483647, -1, 32767}
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 42
    var v4 int16 = 521
    v5 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 8; fi6++ {
        if (4294967295 != 1) {
            if ((v4 | (v4 * (v4 ^ (^(v4))))) == 0) {
                _ = len(v1)
                var v7 int = fi6
                _ = v7
            } else {
                if (42 == 65535) {
                    v4 = v4
                    for fi8 := 0; fi8 < 15; fi8++ {
                        v9 := GenS1{F0: -1, F1: 223, F2: (42 >= 1)}
                        _ = fi8
                        _ = v9
                    }
                    v4 = -1
                } else {
                    var v10 float64 = 0.0
                    var v11 uint8 = 1
                    if (!true) {
                        if (42 >= 0) {
                            var v12 int = -693
                            if (660 == 0) {
                                var v13 uint = 0
                                for fi14 := 0; fi14 < 15; fi14++ {
                                    var v15 uint32 = 2147483648
                                    _ = fi14
                                    _ = v15
                                }
                                _ = v13
                            } else {
                                _ = len(v2)
                                v4 = (v4 % ((v4) | 1))
                                for fi16 := 0; fi16 < 13; fi16++ {
                                    var v17 int64 = 0
                                    _ = fi16
                                    _ = v17
                                }
                            }
                            _ = v12
                        }
                    }
                    _ = v10
                    _ = v11
                }
                switch (-(fi6)) {
                case 128:
                    fmt.Println(fi6)
                }
            }
            var v18 uint16 = 1
            v18 = uint16(v3)
            _ = v18
        }
        var v19 int64 = 296
        _ = fi6
        _ = v19
    }
    var v20 uint16 = 810
    _ = len(v5)
    switch "foo bar" {
    case "\n\t":
        for fi21 := 0; fi21 < 15; fi21++ {
            for fi22 := 0; fi22 < 11; fi22++ {
                var v23 int32 = 0
                v2[0] = 258
                v3 = (v3 * 1)
                _ = fi22
                _ = v23
            }
            if (502 <= 0) {
                v24 := GenS2{F0: 280, F1: (692 < 477), F2: 1}
                _ = v24
            }
            v20 += (v20 >> ((611) & 63))
            _ = fi21
        }
    case "a":
        v25 := GenS2{F0: 65535, F1: true, F2: 2}
        _ = v25
    case "foo bar":
        if true {
            switch ("hello" + "hello") {
            case "":
                v20 = (v20 * 2)
            default:
                var v26 float32 = (3.14159 / (3.14159 * 3.14159))
                _ = v26
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            var v27 uint8 = 0
            _ = v27
        }
    default:
        v20 = 0
    }
    v3 += v3
    var v28 string = "foo bar"
    v28 = v28
    delete(v2, 4294967295)
    for fi29 := 0; fi29 < 20; fi29++ {
        if len(v1) > 0 {
            v1[0] = fi29
        }
        if (false && (false && (3.14159 == -2.25))) {
            fmt.Println(v20)
            switch 269 {
            case -498:
                v30 := GenS3{F0: 513, F1: 65535}
                _ = v30
            case 42:
                for fi31 := 0; fi31 < 1; fi31++ {
                    v5[42] = (0 >= 0)
                    _ = len(v1)
                    _ = fi31
                }
            default:
                switch (v28 + ("a" + (v28 + "hello"))) {
                case "foo bar":
                    fmt.Println(v28, v20)
                case "":
                    for fi32 := 0; fi32 < 8; fi32++ {
                        if (57 > 42) {
                            v33 := GenS1{F0: 107, F1: v3, F2: false}
                            fmt.Println(fi29)
                            _ = v33
                        } else {
                            switch (fi32 ^ (fi32 * -1)) {
                            case 127:
                                var v34 int8 = 127
                                if false {
                                    v35 := GenS1{F0: 243, F1: (^(v3)), F2: (0 < (v4 >> ((2) & 63)))}
                                    v20 = (v20 + v20)
                                    v36 := GenS2{F0: 69, F1: (true || false), F2: 455}
                                    _ = v35
                                    _ = v36
                                } else {
                                    var v37 uint32 = 427
                                    _ = v37
                                }
                                v3 = (v3 - v3)
                                v38 := GenS1{F0: -9223372036854775808, F1: 0, F2: true}
                                _ = v34
                                _ = v38
                            case 4294967295:
                                _ = len(v1)
                            }
                            var v39 uint64 = 4294967295
                            var v40 int16 = (-(v4))
                            _ = v39
                            _ = v40
                        }
                        var v41 int = fi29
                        _ = fi32
                        _ = v41
                    }
                case "hello":
                    v3 = (v3 % ((v3) | 1))
                default:
                    if true {
                        _ = len(v1)
                    }
                }
            }
            fmt.Println(v4, v20, v4)
        }
        _ = fi29
    }
    v20 = v20
    v42 := GenS2{F0: 1, F1: (1e10 != 0.1), F2: 255}
    v43 := GenS1{F0: -1, F1: (v3 & (v3 & (v3 % ((v3) | 1)))), F2: (("foo bar" + v28) <= "\n\t")}
    v44 := GenS1{F0: 9223372036854775807, F1: v3, F2: true}
    for fi45 := 0; fi45 < 5; fi45++ {
        if false {
            for fi46 := 0; fi46 < 7; fi46++ {
                v47 := v2[4294967295]
                _ = fi46
                _ = v47
            }
            v48 := GenS2{F0: 4294967295, F1: (-9223372036854775808 != 1), F2: 410}
            _ = v48
        } else {
            switch v28 {
            case "hello":
                v4 += 32767
            default:
                v28 = v28
            }
        }
        delete(v5, 2147483647)
        if (!((v28 + v28) == ("hello" + ("a" + v28)))) {
            var v49 uint = 403
            v28 = (v28 + "")
            switch fi45 {
            case 265:
                switch v28 {
                case "\n\t":
                    v3 -= (v3 ^ uint8(v20))
                case "hello":
                    v4 = v4
                case "":
                    if true {
                        v50 := GenS4{F0: 66, F1: (v49 << ((250) & 63))}
                        v28 = "a"
                        _ = v50
                    }
                }
            case 9223372036854775807:
                delete(v5, 42)
            }
            _ = v49
        }
        _ = fi45
    }
    fmt.Println(v3)
    var v51 int32 = 549
    for fi52 := 0; fi52 < 14; fi52++ {
        for fi53 := 0; fi53 < 3; fi53++ {
            v28 = v28
            var v54 int = (fi52 % ((fi52) | 1))
            switch ("\n\t" + (("" + v28) + "hello")) {
            case "\n\t":
                if true {
                    v28 = (v28 + "\n\t")
                }
            case "hello":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "foo bar":
                _ = len(v5)
                v54 += fi52
                fmt.Println(v4, v28, v54)
            }
            _ = fi53
            _ = v54
        }
        _ = fi52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v51
}
