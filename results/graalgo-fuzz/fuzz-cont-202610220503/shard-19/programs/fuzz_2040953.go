package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS2 struct { F0 string; F1 uint64; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 bool; F1 int64 }
func genFn0(p0 uint8) int16 {
    p0 += p0
    v7 := []int{42, 4294967295, 0}
    v7 = append(v7, 32767)
    _ = p0
    _ = v7
    return -772
}

func main() {
    var v1 int32 = 0
    var v2 uint32 = 0
    var v3 int32 = (v1 / (((v1 ^ (-(v1)))) | 1))
    var v4 uint8 = 2
    v5 := make(map[uint32]int)
    var v6 float32 = ((0.0 * float32(((0.1 - (-2.25 - 0.1)) / 1.0))) - (-2.25 - (0.1 + ((0.0 + -2.25) - 1e10))))
    if true {
        v8 := GenS1{F0: (42 != 4294967295), F1: 569, F2: 187}
        if true {
            if true {
                v9 := GenS1{F0: (((v2 | v2) > (v2 ^ 780)) && true), F1: 18446744073709551615, F2: v4}
                _ = v9
            }
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 uint8 = 218
                for fi12 := 0; fi12 < 3; fi12++ {
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        if (v11 != v11) {
                            v1 += 1
                            if false {
                                v13 := []int{(fi10 | int(v1)), 32768, fi12, fi10}
                                v14 := GenS2{F0: (("foo bar" + "") + "foo bar"), F1: 0, F2: (fi12 - (fi10 * fi12))}
                                _ = v13
                                _ = v14
                            } else {
                                var v15 uint64 = 42
                                if ("a" == ("foo bar" + "a")) {
                                    switch (fi12 & 0) {
                                    case 42:
                                        var v16 uint32 = v2
                                        _ = v16
                                    }
                                    var v17 string = "hello"
                                    fmt.Println(v4, v11)
                                    _ = v17
                                } else {
                                    var v18 float64 = 1.5
                                    _ = v18
                                }
                                v5[v2] = (fi10 * (-(fi12)))
                                _ = v15
                            }
                            v1 += v1
                        } else {
                            var v19 float64 = 0.1
                            fmt.Println(fi12, v6, fi10)
                            _ = v19
                        }
                    case "a":
                        v3 = (v1 >> ((42) & 63))
                    case "":
                        if (false && (true || (42 > 1))) {
                            v20 := GenS3{F0: -1}
                            var v21 int32 = (-(v1))
                            v6 = v6
                            _ = v20
                            _ = v21
                        } else {
                            if (42 != 620) {
                                var v22 uint = 2
                                _ = v22
                            }
                            v4 += 229
                        }
                    default:
                        v23 := GenS4{F0: 255, F1: "hello"}
                        _ = v23
                    }
                    _ = fi12
                }
                var v24 uint8 = 255
                _ = fi10
                _ = v11
                _ = v24
            }
            v3 += v1
        }
        _ = v8
    }
    v6 = float32(float64((v6 + v6)))
    v4 += 127
    var v25 int64 = -1
    var v26 int32 = v1
    v27 := []int{int(v25), -1}
    if false {
        v4 += 109
        for fi28 := 0; fi28 < 1; fi28++ {
            v2 -= 65535
            v29 := GenS5{F0: -52, F1: (v2 << ((42) & 63))}
            fmt.Println(v4, v4)
            _ = fi28
            _ = v29
        }
        if true {
            var v30 string = ("hello" + "\n\t")
            switch v30 {
            case "a":
                _ = len(v27)
            }
            _ = v30
        } else {
            fmt.Println(v26, v4, v1)
            if false {
                v31 := GenS1{F0: false, F1: 42, F2: v4}
                _ = v31
            }
        }
    }
    if ((v26 - v3) > v3) {
        v26 += v1
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 bool = ((("\n\t" + "foo bar") + ("foo bar" + "foo bar")) <= (("" + "a") + ("foo bar" + "\n\t")))
            for fi34 := 0; fi34 < 13; fi34++ {
                switch 127 {
                case 1:
                    v35 := GenS5{F0: -32768, F1: (^(v2))}
                    _ = v35
                case 128:
                    if false {
                        var v36 float64 = (-2.25 + 0.0)
                        _ = v36
                    }
                case 2147483647:
                    var v37 uint8 = (v4 << ((4294967295) & 63))
                    _ = v37
                default:
                    v3 -= -2147483648
                }
                switch 0 {
                case 1:
                    if (((!true) || (33 < -896)) && ((v2 < v2) || v33)) {
                        _ = len(v5)
                        delete(v5, v2)
                        if len(v27) >= 2 {
                            v27 = v27[1:len(v27)]
                        }
                        var v38 uint16 = 32767
                        v5[(v2 | v2)] = fi34
                        _ = len(v5)
                        _ = v38
                    }
                case 2147483647:
                    if len(v27) > 0 {
                        v27[0] = fi32
                    }
                default:
                    v33 = false
                }
                _ = fi34
            }
            if false {
                fmt.Println(v2)
                for fi39 := 0; fi39 < 11; fi39++ {
                    delete(v5, (^(v2)))
                    if true {
                        switch fi39 {
                        case 532:
                            v4 -= v4
                        case -9223372036854775808:
                            var v40 uint = 255
                            _ = v40
                        default:
                            v2 -= (v2 | (v2 & v2))
                        }
                        switch "\n\t" {
                        case "\n\t":
                            v5[v2] = (fi39 | -9223372036854775808)
                        case "a":
                            switch fi39 {
                            case 9223372036854775807:
                                var v41 int64 = v25
                                _ = v41
                            }
                        default:
                            var v42 uint = 255
                            _ = v42
                        }
                    } else {
                        var v43 float32 = 0.0
                        _ = v43
                    }
                    _ = fi39
                }
                var v44 float64 = 1e10
                _ = v44
            }
            _ = fi32
            _ = v33
        }
        v4 -= v4
    } else {
        if ("hello" >= ("hello" + "foo bar")) {
            v27 = append(v27, -322)
            if ((v25 + v25) == (v25 >> ((46) & 63))) {
                v4 -= (v4 ^ 42)
            }
            switch ("hello" + "") {
            case "hello":
                for fi45 := 0; fi45 < 2; fi45++ {
                    for fi46 := 0; fi46 < 12; fi46++ {
                        fmt.Println(v2, fi46, v1)
                        v47 := GenS1{F0: ("hello" <= "hello"), F1: 18446744073709551615, F2: (^(v4))}
                        v1 = 42
                        _ = fi46
                        _ = v47
                    }
                    v48 := GenS4{F0: 882, F1: ("hello" + "\n\t")}
                    _ = fi45
                    _ = v48
                }
            case "a":
                v4 = 0
            case "":
                var v49 uint8 = 213
                _ = v49
            default:
                v3 = (v26 % ((v26) | 1))
            }
        } else {
            _ = len(v27)
            v50 := GenS6{F0: (!false), F1: v25}
            _ = v50
        }
        v3 -= v26
        fmt.Println(v1, v2)
    }
    v3 = v26
    switch -2147483648 {
    case -9223372036854775808:
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v51 string = ("foo bar" + "foo bar")
            _ = v51
        default:
            v52 := GenS6{F0: (true && true), F1: (v25 - (v25 ^ v25))}
            _ = v52
        }
    case -1:
        fmt.Println(v25, v2)
    default:
        var v53 int32 = (v3 << ((435) & 63))
        _ = v53
    }
    if false {
        v54 := GenS4{F0: uint(v4), F1: "\n\t"}
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
