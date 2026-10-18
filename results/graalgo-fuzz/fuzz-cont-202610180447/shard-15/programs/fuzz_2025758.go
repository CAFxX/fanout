package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint }
func genFn0() int {
    if (63 > -876) {
        if true {
            v6 := GenS1{F0: 437, F1: 0, F2: 1}
            _ = v6
        } else {
            if (true || (9223372036854775807 != 9223372036854775807)) {
                for fi7 := 0; fi7 < 6; fi7++ {
                    var v8 int16 = 1000
                    var v9 uint32 = 2
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                var v10 int16 = 1
                _ = v10
            }
            var v11 string = "a"
            v12 := []int{0}
            _ = v11
            _ = v12
        }
        v13 := GenS2{F0: 786}
        var v14 int32 = -863
        _ = v13
        _ = v14
    } else {
        var v15 uint32 = 666
        var v16 int16 = -32768
        for fi17 := 0; fi17 < 13; fi17++ {
            for fi18 := 0; fi18 < 15; fi18++ {
                v16 = (^(v16))
                fmt.Println(fi18, fi18)
                _ = fi18
            }
            _ = fi17
        }
        _ = v15
        _ = v16
    }
    var v19 int32 = -216
    for fi20 := 0; fi20 < 2; fi20++ {
        v19 = v19
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 uint64 = 1
            v22 -= v22
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    _ = v19
    return -1
}
func genFn1(p0 int8, p1 int8, p2 uint16) uint {
    var v23 uint32 = 83
    var v24 int = 42
    for fi25 := 0; fi25 < 13; fi25++ {
        if false {
            var v26 bool = true
            _ = v26
        } else {
            if false {
                switch v24 {
                case 32767:
                    v27 := GenS2{F0: uint16(v23)}
                    _ = v27
                case 2147483647:
                    switch 4294967295 {
                    case 244:
                        for fi28 := 0; fi28 < 12; fi28++ {
                            v29 := GenS1{F0: 196, F1: 18446744073709551615, F2: -1}
                            v30 := []int{}
                            v30 = append(v30, genFn0())
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    case 127:
                        switch fi25 {
                        case 127:
                            v31 := []int{(v24 % ((genFn0()) | 1)), fi25}
                            _ = v31
                        case -388:
                            v32 := GenS1{F0: 4294967295, F1: 42, F2: 32767}
                            _ = v32
                        case 9223372036854775807:
                            if (!(1 == 1000)) {
                                v33 := GenS1{F0: 642, F1: 0, F2: 0}
                                var v34 int16 = 0
                                _ = v33
                                _ = v34
                            } else {
                                v24 += -1
                            }
                        }
                    case 128:
                        var v35 uint = 65535
                        _ = v35
                    }
                case 673:
                    p1 += int8(v23)
                }
            } else {
                var v36 int = (fi25 + v24)
                var v37 uint8 = 255
                var v38 uint64 = 1
                _ = v36
                _ = v37
                _ = v38
            }
            switch "foo bar" {
            case "hello":
                v23 += v23
            default:
                p0 += (p1 / ((p1) | 1))
            }
            fmt.Println(v23)
        }
        _ = fi25
    }
    fmt.Println(p0, p0)
    if (((^(v24)) < v24) && (((1.5 - 0.0) > 0.0) || false)) {
        var v39 bool = (1 >= 695)
        for fi40 := 0; fi40 < 18; fi40++ {
            for fi41 := 0; fi41 < 3; fi41++ {
                var v42 int64 = 42
                _ = fi41
                _ = v42
            }
            var v43 int64 = -9223372036854775808
            v44 := GenS1{F0: 42, F1: 4294967295, F2: 0}
            _ = fi40
            _ = v43
            _ = v44
        }
        _ = v39
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    return 255
}

func main() {
    v1 := []int{}
    v2 := []uint32{544, 74}
    v3 := []uint32{}
    v4 := make(map[string]int)
    var v5 int16 = 754
    v45 := GenS3{F0: 42}
    for fi46 := 0; fi46 < 10; fi46++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v47 := GenS2{F0: 32767}
        _ = fi46
        _ = v47
    }
    if (1.5 == 0.0) {
        var v48 int64 = -563
        v48 = v48
        _ = v48
    }
    v49 := GenS4{F0: "a", F1: 2}
    if (-2147483648 > 0) {
        fmt.Println(v5)
        if false {
            switch "\n\t" {
            case "":
                var v50 int = 2147483647
                _ = v50
            case "\n\t":
                if (42 >= 42) {
                    if true {
                        if ((!false) && false) {
                            if true {
                                v5 = (v5 + (-(v5)))
                                v5 += (v5 % ((42) | 1))
                                _ = len(v2)
                            }
                            v5 -= v5
                        } else {
                            switch "" {
                            case "hello":
                                var v51 float32 = 3.14159
                                _ = v51
                            }
                        }
                    } else {
                        for fi52 := 0; fi52 < 14; fi52++ {
                            v4["a"] = fi52
                            _ = fi52
                        }
                        fmt.Println(v5)
                    }
                    var v53 int8 = 42
                    _ = v53
                } else {
                    var v54 uint16 = 547
                    _ = v54
                }
            }
            var v55 float64 = (1e10 * -2.25)
            _ = v55
        }
        switch "a" {
        case "\n\t":
            v56 := GenS1{F0: 18446744073709551615, F1: 2, F2: (v5 ^ v5)}
            _ = v56
        case "a":
            switch 32768 {
            case 42:
                var v57 int8 = 127
                _ = v57
            case 0:
                for fi58 := 0; fi58 < 13; fi58++ {
                    v5 += (v5 + (^(v5)))
                    _ = fi58
                }
            case 128:
                var v59 int64 = 0
                _ = v59
            default:
                v60 := GenS1{F0: genFn1(1, 119, 1), F1: 879, F2: v5}
                _ = v60
            }
        }
    }
    fmt.Println(v5)
    v61 := GenS5{F0: 2147483647, F1: 2147483647}
    var v62 int8 = 0
    v63 := GenS6{F0: 298}
    var v64 uint8 = 2
    fmt.Println(v5, v62)
    v65 := v4["a"]
    var v66 uint16 = 1
    for fi67 := 0; fi67 < 5; fi67++ {
        for fi68 := 0; fi68 < 18; fi68++ {
            var v69 uint16 = uint16(v64)
            delete(v4, "hello")
            _ = fi68
            _ = v69
        }
        var v70 int32 = -2147483648
        if (((1e10 * -2.25) <= (float32(3.14159) * (0.1 - ((((3.14159 / 3.14159) + -2.25) - 0.1) + 1.5)))) && (int64(v62) == 0)) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi67
        _ = v70
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v49
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v65
    _ = v66
}
