package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS4 struct { F0 int; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint64; F1 int; F2 uint }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 bool; F1 string }
func genFn0() int {
    v6 := GenS1{F0: "\n\t"}
    v7 := []int{9223372036854775807, 42, -305}
    var v8 uint64 = 4294967295
    var v9 int32 = 1
    if true {
        if (32767 >= 42) {
            v9 -= (v9 >> ((2) & 63))
        }
        v10 := GenS1{F0: "foo bar"}
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 int = 583
    var v2 int = (^(v1))
    var v3 int64 = 9223372036854775807
    v4 := []uint32{927}
    var v5 uint = 0
    if (("hello" + "a") != "foo bar") {
        var v11 uint32 = 42
        _ = v11
    }
    if len(v4) > 0 {
        v4[0] = 65535
    }
    var v12 uint32 = 2
    v4 = append(v4, (v12 + (^(v12))))
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 float32 = 1e10
        _ = fi13
        _ = v14
    }
    v3 = v3
    var v15 bool = false
    var v16 bool = (((v3 + -9223372036854775808) != v3) || v15)
    v16 = (3.14159 == 3.14159)
    v15 = true
    v17 := GenS1{F0: "hello"}
    var v18 bool = ((v5 + v5) >= (v5 - 411))
    if v16 {
        if ("a" == ("a" + "")) {
            v18 = v18
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 uint16 = 65535
                v21 := GenS2{F0: -72}
                if (-829 > -32768) {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        v23 := GenS3{F0: (v12 + v12), F1: 9, F2: v3}
                        _ = fi22
                        _ = v23
                    }
                    switch (v1 * (^(v1))) {
                    case 32768:
                        if ((!(true && (true && v15))) && (("a" + "hello") <= ("" + "hello"))) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if v16 {
                                var v24 uint32 = v12
                                switch "foo bar" {
                                case "\n\t":
                                    var v25 bool = v15
                                    _ = v25
                                case "foo bar":
                                    if ((v15 && v15) || (v20 >= (v20 / ((v20) | 1)))) {
                                        var v26 int = 553
                                        _ = v26
                                    }
                                case "a":
                                    switch (v1 >> ((1) & 63)) {
                                    case -2147483648:
                                        switch v2 {
                                        case 42:
                                            v24 += 42
                                        case -327:
                                            v27 := GenS4{F0: (v2 * fi19), F1: (v20 - (^(v20))), F2: 255}
                                            _ = v27
                                        case -9223372036854775808:
                                            if v18 {
                                                if (2147483647 != -2147483648) {
                                                    for fi28 := 0; fi28 < 4; fi28++ {
                                                        v1 -= (-(v1))
                                                        _ = fi28
                                                    }
                                                } else {
                                                    v29 := GenS5{F0: 223, F1: v1, F2: 395}
                                                    var v30 int8 = 1
                                                    _ = v29
                                                    _ = v30
                                                }
                                                if v18 {
                                                    fmt.Println(v3, v16)
                                                    fmt.Println(v20, v3, v15)
                                                } else {
                                                    for fi31 := 0; fi31 < 11; fi31++ {
                                                        if ((v20 % ((v20) | 1)) != 65535) {
                                                            switch (fi31 - -102) {
                                                            case -2147483648:
                                                                v16 = v15
                                                            default:
                                                                v32 := GenS3{F0: 1, F1: 65535, F2: v3}
                                                                _ = v32
                                                            }
                                                            v15 = (((v24 > v24) || (uint32(v20) <= 42)) && (1000 == 42))
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        fmt.Println(v16, v15)
                                                        var v33 string = "\n\t"
                                                        _ = fi31
                                                        _ = v33
                                                    }
                                                    var v34 float32 = 0.1
                                                    _ = v34
                                                }
                                            }
                                        default:
                                            switch (("\n\t" + "a") + "a") {
                                            case "\n\t":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case "":
                                                if v15 {
                                                    v3 -= v3
                                                    for fi35 := 0; fi35 < 10; fi35++ {
                                                        fmt.Println(v20, v1, v20)
                                                        v5 = uint(v24)
                                                        _ = fi35
                                                    }
                                                } else {
                                                    var v36 int8 = 42
                                                    _ = v36
                                                }
                                            }
                                        }
                                    case -1:
                                        if (42 > -2147483648) {
                                            if (false || (!((^(v20)) == (^(v20))))) {
                                                var v37 int32 = 0
                                                _ = v37
                                            }
                                            switch genFn0() {
                                            case 32767:
                                                v18 = (v5 > 0)
                                            default:
                                                v5 += uint(v20)
                                            }
                                        }
                                    }
                                }
                                _ = v24
                            }
                        } else {
                            var v38 int = v2
                            for fi39 := 0; fi39 < 6; fi39++ {
                                switch ("\n\t" + "\n\t") {
                                case "a":
                                    var v40 uint32 = 748
                                    _ = v40
                                case "hello":
                                    fmt.Println(v3, v5, v12)
                                case "foo bar":
                                    v41 := GenS6{F0: 2147483647, F1: 127}
                                    _ = v41
                                default:
                                    v16 = true
                                }
                                var v42 uint16 = v20
                                _ = fi39
                                _ = v42
                            }
                            _ = v38
                        }
                    default:
                        v43 := GenS7{F0: (2147483647 <= 1), F1: (("" + ("a" + ("\n\t" + ("hello" + "foo bar")))) + ("hello" + "\n\t"))}
                        _ = v43
                    }
                }
                _ = fi19
                _ = v20
                _ = v21
            }
            fmt.Println(v3, v5)
        } else {
            switch "foo bar" {
            case "a":
                if v15 {
                    var v44 string = ""
                    if ((v5 % ((v5) | 1)) <= v5) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        for fi45 := 0; fi45 < 19; fi45++ {
                            var v46 uint32 = 2
                            for fi47 := 0; fi47 < 16; fi47++ {
                                var v48 uint64 = 18446744073709551615
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                fmt.Println(fi45, v3)
                                _ = fi47
                                _ = v48
                            }
                            _ = fi45
                            _ = v46
                        }
                    } else {
                        var v49 int32 = -347
                        v3 = (-(v3))
                        v4 = append(v4, v12)
                        _ = v49
                    }
                    _ = v44
                }
            default:
                if (v16 && ((uint16(v5) <= 292) || (42 >= 916))) {
                    v1 += (v1 / ((v2) | 1))
                    fmt.Println(v3)
                    var v50 uint8 = 255
                    _ = v50
                } else {
                    if len(v4) > 0 {
                        v4[0] = 65535
                    }
                }
            }
            switch "hello" {
            case "foo bar":
                v2 += v2
            }
        }
    } else {
        v51 := GenS2{F0: 1}
        v16 = (true || ((false || (127 <= 127)) || v18))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v51
    }
    var v52 uint8 = 168
    if v15 {
        var v53 int8 = 100
        _ = v53
    } else {
        _ = len(v4)
    }
    switch "a" {
    case "a":
        if false {
            var v54 int32 = 2147483647
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v54
        } else {
            var v55 int16 = int16(v12)
            if v16 {
                v56 := GenS4{F0: genFn0(), F1: uint16(v12), F2: v52}
                fmt.Println(v12, v55, v5)
                _ = v56
            }
            var v57 uint8 = (v52 % ((235) | 1))
            _ = v55
            _ = v57
        }
    case "\n\t":
        var v58 uint16 = 0
        _ = v58
    }
    if len(v4) > 0 {
        v4[0] = (^(v12))
    }
    if len(v4) > 0 {
        v4[0] = (v12 * v12)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
