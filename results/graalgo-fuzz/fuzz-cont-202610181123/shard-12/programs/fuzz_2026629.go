package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int64; F2 uint8 }
type GenS4 struct { F0 bool; F1 string; F2 int16 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 string }
func genFn0() uint8 {
    v7 := []int{}
    var v8 float64 = 1.5
    _ = v7
    _ = v8
    return 4
}

func main() {
    var v1 bool = false
    var v2 uint = 42
    var v3 int64 = -9223372036854775808
    var v4 int64 = v3
    var v5 int16 = 42
    var v6 int32 = 2147483647
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := []int{}
        if len(v10) > 0 {
            v10[0] = (fi9 >> ((460) & 63))
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            v6 -= (v6 / ((v6) | 1))
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    v12 := []int{int(v6), -9223372036854775808}
    v2 = (^(v2))
    if ("foo bar" == ("a" + "\n\t")) {
        v5 += 601
        var v13 int32 = 101
        _ = v13
    } else {
        if (9223372036854775807 != -9223372036854775808) {
            for fi14 := 0; fi14 < 1; fi14++ {
                v15 := GenS1{F0: (v4 * -9223372036854775808), F1: 255, F2: ("hello" + "hello")}
                _ = len(v12)
                _ = fi14
                _ = v15
            }
        } else {
            var v16 uint8 = 2
            var v17 uint64 = 0
            _ = v16
            _ = v17
        }
    }
    if len(v12) > 0 {
        v12[0] = 127
    }
    var v18 int32 = (v6 % ((0) | 1))
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 float64 = 3.14159
        _ = fi19
        _ = v20
    }
    switch "hello" {
    case "a":
        if v1 {
            v1 = (1 > 732)
            var v21 int = 0
            var v22 uint = v2
            _ = v21
            _ = v22
        } else {
            var v23 float32 = (1e10 + (0.0 - ((0.0 - 3.14159) * 3.14159)))
            if len(v12) > 0 {
                v12[0] = 4294967295
            }
            switch 1 {
            case -2147483648:
                if ((v18 ^ (v18 + v6)) < v18) {
                    if false {
                        if false {
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                        }
                        if v1 {
                            v18 -= (v6 + (^(v6)))
                        } else {
                            for fi24 := 0; fi24 < 13; fi24++ {
                                if true {
                                    if (v1 || ("hello" == "hello")) {
                                        fmt.Println(v23)
                                        if len(v12) > 0 {
                                            v12[0] = fi24
                                        }
                                        v25 := GenS2{F0: (v5 ^ (-(v5)))}
                                        _ = v25
                                    }
                                    for fi26 := 0; fi26 < 2; fi26++ {
                                        var v27 int16 = (-(v5))
                                        var v28 uint16 = 32767
                                        _ = fi26
                                        _ = v27
                                        _ = v28
                                    }
                                } else {
                                    v4 += (v4 & v3)
                                    if v1 {
                                        switch "a" {
                                        case "a":
                                            for fi29 := 0; fi29 < 11; fi29++ {
                                                v30 := GenS3{F0: (((!(1 != 1)) && false) && (-52 > -1)), F1: int64(v5), F2: 126}
                                                _ = fi29
                                                _ = v30
                                            }
                                        case "foo bar":
                                            v12 = append(v12, fi24)
                                        }
                                        var v31 string = ("" + "a")
                                        _ = v31
                                    } else {
                                        if (-2.25 >= 0.1) {
                                            for fi32 := 0; fi32 < 5; fi32++ {
                                                if len(v12) >= 2 {
                                                    v12 = v12[1:len(v12)]
                                                }
                                                v33 := GenS4{F0: v1, F1: "hello", F2: v5}
                                                _ = fi32
                                                _ = v33
                                            }
                                            var v34 uint8 = 42
                                            _ = v34
                                        } else {
                                            v4 = (v3 & (v4 * (^(v4))))
                                            v35 := GenS4{F0: false, F1: "a", F2: v5}
                                            v3 = 0
                                            _ = v35
                                        }
                                    }
                                    if len(v12) > 0 {
                                        v12[0] = (fi24 - (fi24 & (fi24 | fi24)))
                                    }
                                }
                                v36 := GenS5{F0: genFn0(), F1: 42}
                                _ = fi24
                                _ = v36
                            }
                            var v37 string = "hello"
                            v2 = v2
                            _ = v37
                        }
                    } else {
                        var v38 int = 4294967295
                        v39 := GenS5{F0: 1, F1: -84}
                        _ = v38
                        _ = v39
                    }
                } else {
                    for fi40 := 0; fi40 < 11; fi40++ {
                        var v41 int8 = 3
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        _ = fi40
                        _ = v41
                    }
                    v4 -= 42
                    var v42 uint32 = 80
                    _ = v42
                }
            case -9223372036854775808:
                var v43 float32 = v23
                _ = v43
            }
            _ = v23
        }
    case "hello":
        var v44 int32 = 42
        _ = v44
    case "\n\t":
        switch 32768 {
        case -2147483648:
            var v45 uint64 = 1
            _ = v45
        default:
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
    default:
        for fi46 := 0; fi46 < 6; fi46++ {
            if v1 {
                var v47 uint64 = 0
                switch "" {
                case "a":
                    switch (^(fi46)) {
                    case 1:
                        for fi48 := 0; fi48 < 11; fi48++ {
                            var v49 int = fi48
                            _ = fi48
                            _ = v49
                        }
                    default:
                        if v1 {
                            v18 = int32(v4)
                        }
                    }
                default:
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                }
                switch (^(fi46)) {
                case 32768:
                    fmt.Println(v3, v18)
                default:
                    fmt.Println(v18, v18)
                }
                _ = v47
            }
            for fi50 := 0; fi50 < 13; fi50++ {
                var v51 int64 = v3
                v18 -= (v6 * 161)
                _ = fi50
                _ = v51
            }
            v1 = ((int8(v5) <= -51) || v1)
            _ = fi46
        }
    }
    v52 := GenS6{F0: 2, F1: (^(v2))}
    v3 -= v4
    v6 = 1
    var v53 uint32 = 2147483647
    v54 := GenS7{F0: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
