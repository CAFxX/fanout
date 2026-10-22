package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 string; F2 int }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int8, p1 int32, p2 uint32) int16 {
    var v5 uint64 = 251
    var v6 uint16 = uint16(p2)
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            p0 -= -32
            v9 := GenS1{F0: 255, F1: 65535, F2: uint64(p2)}
            _ = fi8
            _ = v9
        }
        fmt.Println(p1)
        _ = fi7
    }
    if (!(((("foo bar" + "hello") + "") >= "a") && (((0.0 - 1e10) * 0.1) != 1.5))) {
        v10 := []int{}
        if false {
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 int64 = -9223372036854775808
                var v13 int64 = v12
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            v14 := GenS1{F0: uint8(v5), F1: 102, F2: 2}
            _ = v14
        }
        _ = v10
    } else {
        v15 := GenS1{F0: 128, F1: 2, F2: v5}
        _ = v15
    }
    var v16 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    var v3 int8 = 42
    v4 := make(map[int]uint32)
    if true {
        switch ((("foo bar" + "a") + "hello") + "a") {
        case "":
            v3 = (v3 / ((v3) | 1))
        case "foo bar":
            if false {
                var v17 uint16 = 1
                v17 = v17
                delete(v1, 757)
                _ = v17
            }
        default:
            v18 := GenS2{F0: 4294967295}
            _ = v18
        }
        v3 = (v3 & 100)
        fmt.Println(v3)
    } else {
        v19 := GenS2{F0: 18446744073709551615}
        if (float64(0.0) > 1e10) {
            switch "\n\t" {
            case "":
                var v20 float32 = 0.1
                _ = v20
            case "\n\t":
                if (!true) {
                    v3 = 127
                    if ("hello" >= "a") {
                        fmt.Println(v3)
                        _ = len(v2)
                        var v21 uint64 = 957
                        _ = v21
                    } else {
                        for fi22 := 0; fi22 < 4; fi22++ {
                            v3 -= v3
                            v23 := GenS1{F0: 84, F1: 1, F2: 946}
                            _ = fi22
                            _ = v23
                        }
                    }
                    v3 = (v3 ^ v3)
                }
            case "foo bar":
                for fi24 := 0; fi24 < 13; fi24++ {
                    v25 := GenS3{F0: 15, F1: ("" + (("hello" + ("hello" + ("hello" + "hello"))) + ("" + "foo bar"))), F2: (^(fi24))}
                    v26 := GenS4{F0: ("a" + "foo bar"), F1: 1}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            v27 := []int{323}
            var v28 bool = (!(true || false))
            _ = v27
            _ = v28
        } else {
            for fi29 := 0; fi29 < 5; fi29++ {
                v3 = v3
                v3 -= 1
                _ = fi29
            }
        }
        if (true || false) {
            delete(v4, -686)
            v30 := GenS5{F0: 853}
            v1[4294967295] = 65535
            _ = v30
        } else {
            v3 = (v3 | (-(v3)))
        }
        _ = v19
    }
    v31 := GenS1{F0: 1, F1: 1, F2: 758}
    if true {
        var v32 bool = true
        var v33 int16 = -1
        _ = v32
        _ = v33
    }
    var v34 int32 = -2147483648
    var v35 int8 = v3
    var v36 int = 32768
    v37 := GenS6{F0: (-(v35)), F1: "hello"}
    v38 := GenS6{F0: v3, F1: "\n\t"}
    v39 := v2[v36]
    switch ("\n\t" + (("\n\t" + "") + "\n\t")) {
    case "a":
        switch (v36 << ((183) & 63)) {
        case -11:
            fmt.Println(v34, v34)
        case 9223372036854775807:
            if (!(3 > 54)) {
                switch v39 {
                case 4294967295:
                    var v40 int32 = int32(v36)
                    _ = v40
                default:
                    var v41 int16 = -1
                    _ = v41
                }
                v36 = v36
            } else {
                if true {
                    if (!((!(true && (!true))) && (!true))) {
                        fmt.Println(v3, v3)
                        v3 = -1
                        _ = len(v2)
                    } else {
                        if (v39 != (v36 + v39)) {
                            for fi42 := 0; fi42 < 1; fi42++ {
                                fmt.Println(v35, v39)
                                var v43 uint32 = 992
                                _ = fi42
                                _ = v43
                            }
                            v44 := GenS1{F0: 13, F1: 2, F2: 0}
                            v3 = v35
                            _ = v44
                        }
                        for fi45 := 0; fi45 < 7; fi45++ {
                            switch v39 {
                            case -2147483648:
                                switch "hello" {
                                case "a":
                                    v39 -= v36
                                default:
                                    var v46 int32 = 1
                                    _ = v46
                                }
                            case 1:
                                v47 := GenS7{F0: 31}
                                _ = v47
                            case 42:
                                var v48 bool = true
                                _ = v48
                            default:
                                var v49 uint32 = 65535
                                _ = v49
                            }
                            _ = fi45
                        }
                    }
                }
            }
        case 659:
            if (v39 == (v36 * v36)) {
                if true {
                    for fi50 := 0; fi50 < 9; fi50++ {
                        v34 += 0
                        v39 -= (-(v39))
                        var v51 uint32 = 2
                        _ = fi50
                        _ = v51
                    }
                    var v52 uint64 = 4294967295
                    _ = v52
                } else {
                    v35 += 42
                    v3 = v35
                }
            }
        }
    case "":
        var v53 int = 32768
        _ = v53
    default:
        v39 = 127
    }
    var v54 uint64 = 4294967295
    v55 := GenS5{F0: v54}
    var v56 uint64 = v54
    v56 -= 422
    if (!("foo bar" >= (((("a" + "foo bar") + (("foo bar" + "\n\t") + "hello")) + "\n\t") + "a"))) {
        var v57 uint32 = 65535
        fmt.Println(v34, v35)
        if ((v57 << ((uint(v54)) & 63)) > 65535) {
            if (-1 <= -9223372036854775808) {
                v58 := v2[(v39 - 32768)]
                switch 42 {
                case -2147483648:
                    v59 := GenS6{F0: (v3 - v3), F1: "foo bar"}
                    _ = v59
                case 32767:
                    var v60 bool = false
                    _ = v60
                default:
                    switch int(v57) {
                    case 711:
                        var v61 uint64 = 4294967295
                        _ = v61
                    case 42:
                        for fi62 := 0; fi62 < 18; fi62++ {
                            v34 -= v34
                            var v63 int16 = genFn0(v3, (v34 + (v34 << ((4294967295) & 63))), v57)
                            var v64 uint8 = 185
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                    case 840:
                        _ = len(v1)
                    default:
                        v65 := v1[v36]
                        _ = v65
                    }
                }
                _ = v58
            }
            var v66 int16 = 32767
            _ = len(v1)
            _ = v66
        } else {
            if (32767 == 157) {
                for fi67 := 0; fi67 < 1; fi67++ {
                    v68 := []int{v36, (v36 | (fi67 * v36)), 9223372036854775807}
                    _ = fi67
                    _ = v68
                }
                v57 = (v57 ^ (v57 * (v57 >> ((65535) & 63))))
            } else {
                fmt.Println(v54, v34, v54)
                fmt.Println(v34)
            }
            v54 = (v54 + v54)
        }
        _ = v57
    } else {
        for fi69 := 0; fi69 < 18; fi69++ {
            var v70 int32 = v34
            _ = fi69
            _ = v70
        }
    }
    v71 := []int{(v39 | v36), v36}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v71:", len(v71), v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v54
    _ = v55
    _ = v56
    _ = v71
}
