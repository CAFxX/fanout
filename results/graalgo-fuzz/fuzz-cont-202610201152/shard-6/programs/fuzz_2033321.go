package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 int32 = 1
    v5 := make(map[int]int)
    if (3.14159 != 3.14159) {
        var v6 int = 465
        v6 = v6
        _ = v6
    } else {
        switch "\n\t" {
        case "hello":
            switch 127 {
            case 127:
                var v7 uint = 390
                _ = v7
            case -395:
                fmt.Println(v1, v4)
            }
        }
        fmt.Println(v4, v4)
        if v1 {
            v4 -= 42
            v8 := GenS1{F0: "foo bar"}
            v4 -= v4
            _ = v8
        }
    }
    v4 -= (v4 << ((1) & 63))
    if v1 {
        v9 := GenS2{F0: 128, F1: 18446744073709551615, F2: 32767}
        v10 := []int{127, -1}
        if true {
            var v11 int8 = 60
            if ("a" < ("hello" + "a")) {
                if ((("hello" + "") + ("hello" + "foo bar")) == (("foo bar" + "hello") + (("foo bar" + ("" + "foo bar")) + "foo bar"))) {
                    fmt.Println(v4, v11)
                    v12 := GenS3{F0: 2}
                    _ = v12
                } else {
                    if false {
                        if v1 {
                            fmt.Println(v4, v4, v1)
                            var v13 int32 = v4
                            for fi14 := 0; fi14 < 11; fi14++ {
                                var v15 bool = true
                                v15 = (true || (!v15))
                                _ = fi14
                                _ = v15
                            }
                            _ = v13
                        } else {
                            v11 += 72
                            for fi16 := 0; fi16 < 11; fi16++ {
                                var v17 float64 = 0.1
                                _ = fi16
                                _ = v17
                            }
                        }
                        switch "a" {
                        case "foo bar":
                            if (!(1 <= 65535)) {
                                var v18 float32 = (1e10 - 1e10)
                                for fi19 := 0; fi19 < 18; fi19++ {
                                    var v20 uint32 = 709
                                    _ = fi19
                                    _ = v20
                                }
                                switch ("\n\t" + "foo bar") {
                                case "foo bar":
                                    v21 := GenS2{F0: 2147483647, F1: 126, F2: 168}
                                    _ = v21
                                case "\n\t":
                                    switch -1 {
                                    case 128:
                                        if len(v10) > 0 {
                                            v10[0] = -9223372036854775808
                                        }
                                    case 1:
                                        if v1 {
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                        }
                                    case 130:
                                        fmt.Println(v4, v18, v4)
                                    }
                                default:
                                    v4 -= 2147483647
                                }
                                _ = v18
                            }
                        }
                    } else {
                        for fi22 := 0; fi22 < 18; fi22++ {
                            var v23 int64 = 11
                            v24 := GenS2{F0: (fi22 * fi22), F1: 282, F2: 923}
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 15; fi25++ {
                            v11 += (v11 * 100)
                            v2[1] = 9223372036854775807
                            _ = fi25
                        }
                    }
                    v26 := GenS1{F0: ""}
                    var v27 bool = v1
                    _ = v26
                    _ = v27
                }
            } else {
                var v28 int16 = -41
                _ = v28
            }
            v1 = false
            _ = v11
        }
        _ = v10
        _ = v9
    }
    var v29 float32 = float32(float64(float32((0.1 * 1e10))))
    var v30 uint16 = 32767
    v31 := []int{}
    _ = len(v5)
    v32 := GenS4{F0: 65535}
    for fi33 := 0; fi33 < 16; fi33++ {
        v1 = ("a" < "")
        switch fi33 {
        case 2147483647:
            v30 = (v30 % (((v30 * v30)) | 1))
        }
        _ = fi33
    }
    v1 = v1
    v34 := GenS2{F0: 4294967295, F1: 4294967295, F2: 1}
    for fi35 := 0; fi35 < 6; fi35++ {
        fmt.Println(fi35, v4, fi35)
        if ((1.5 < float64(v29)) && v1) {
            var v36 string = "\n\t"
            var v37 float32 = ((1e10 + 0.1) - (v29 * 0.0))
            v38 := GenS3{F0: 2}
            _ = v36
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 15; fi39++ {
            v1 = v1
            _ = fi39
        }
        _ = fi35
    }
    v30 += v30
    switch "hello" {
    case "hello":
        v40 := GenS2{F0: 128, F1: 42, F2: 912}
        _ = v40
    case "":
        for fi41 := 0; fi41 < 17; fi41++ {
            if (v1 || (v4 < 0)) {
                switch fi41 {
                case 775:
                    var v42 uint = 42
                    _ = v42
                default:
                    v1 = (1e10 == 3.14159)
                }
                switch fi41 {
                case 1:
                    var v43 int32 = v4
                    _ = v43
                case -677:
                    switch 4294967295 {
                    case 0:
                        fmt.Println(fi41, v4, v1)
                    case 42:
                        fmt.Println(fi41, v1, v1)
                    case 455:
                        _ = len(v3)
                    default:
                        v4 += (v4 << ((0) & 63))
                    }
                default:
                    var v44 uint32 = 2147483648
                    _ = v44
                }
            }
            _ = fi41
        }
    case "a":
        v1 = v1
    }
    if v1 {
        v1 = false
    } else {
        delete(v2, 280)
    }
    if v1 {
        v45 := GenS5{F0: 636, F1: int(v30)}
        v29 -= v29
        if true {
            v29 = (float32((3.14159 * ((3.14159 * 0.1) / -2.25))) - v29)
        } else {
            var v46 uint32 = 4294967295
            switch -1 {
            case 635:
                v47 := v2[1]
                _ = v47
            case 42:
                v29 -= (float32(0.0) + (0.0 + v29))
            }
            _ = v46
        }
        _ = v45
    } else {
        var v48 uint = 65535
        switch "a" {
        case "foo bar":
            switch 301 {
            case -1:
                _ = len(v3)
            default:
                if (!true) {
                    v48 = uint(v30)
                    var v49 uint = v48
                    _ = v49
                } else {
                    var v50 int8 = 48
                    switch 127 {
                    case 32767:
                        fmt.Println(v4)
                    default:
                        for fi51 := 0; fi51 < 12; fi51++ {
                            var v52 uint = 110
                            fmt.Println(v50, v29)
                            _ = fi51
                            _ = v52
                        }
                    }
                    v48 = v48
                    _ = v50
                }
            }
        case "":
            for fi53 := 0; fi53 < 9; fi53++ {
                v48 += 4294967295
                for fi54 := 0; fi54 < 1; fi54++ {
                    var v55 float32 = v29
                    v56 := GenS5{F0: 4294967295, F1: int(v30)}
                    v57 := GenS6{F0: uint64(v48)}
                    _ = fi54
                    _ = v55
                    _ = v56
                    _ = v57
                }
                var v58 int16 = 42
                _ = fi53
                _ = v58
            }
        case "hello":
            var v59 uint = 65535
            _ = v59
        }
        var v60 bool = ((2147483647 == (v4 % ((v4) | 1))) || v1)
        _ = v48
        _ = v60
    }
    v31 = append(v31, -1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
}
