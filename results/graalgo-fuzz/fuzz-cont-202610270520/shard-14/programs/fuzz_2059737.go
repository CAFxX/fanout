package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint64; F2 string }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int32; F1 int16; F2 uint32 }
func genFn0() uint {
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 int = 9223372036854775807
            v9 = v9
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := []int{52, 42}
    _ = v10
    return 0
}

func main() {
    var v1 uint8 = 42
    var v2 uint16 = 0
    var v3 float64 = (0.0 - ((3.14159 / ((-2.25 - 1e10) + 1.5)) * (-2.25 + 3.14159)))
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    var v6 int64 = -906
    v3 -= 0.1
    v11 := GenS1{F0: v6}
    v12 := GenS2{F0: (^(v6)), F1: 2, F2: ("hello" + "\n\t")}
    switch -676 {
    case -1:
        if false {
            var v13 uint16 = v2
            v1 -= v1
            _ = v13
        } else {
            var v14 uint = 4294967295
            if false {
                for fi15 := 0; fi15 < 19; fi15++ {
                    v14 = (^(v14))
                    fmt.Println(fi15)
                    v16 := []int{}
                    _ = fi15
                    _ = v16
                }
                if false {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        var v18 uint8 = 127
                        _ = fi17
                        _ = v18
                    }
                }
                for fi19 := 0; fi19 < 8; fi19++ {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        switch "hello" {
                        case "hello":
                            v21 := v5[(fi19 * fi20)]
                            _ = v21
                        }
                        switch fi20 {
                        case 2147483647:
                            v14 += (^(v14))
                        case -1:
                            v1 = (v1 * v1)
                        case 32767:
                            if false {
                                v6 = (v6 ^ v6)
                            }
                            var v22 int = fi19
                            var v23 int32 = 1
                            _ = v22
                            _ = v23
                        }
                        _ = fi20
                    }
                    _ = fi19
                }
            }
            v24 := []int{-795, -2147483648, 319, -166}
            _ = v14
            _ = v24
        }
    case 42:
        for fi25 := 0; fi25 < 13; fi25++ {
            v26 := GenS3{F0: (v1 - (^(v1))), F1: 124, F2: 1}
            var v27 int = (fi25 >> ((genFn0()) & 63))
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        var v29 uint = 805
        v30 := []int{fi28, 32767, (fi28 >> (((v29 - v29)) & 63))}
        var v31 int16 = 411
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    if false {
        if false {
            if false {
                v1 -= (^(v1))
                for fi32 := 0; fi32 < 10; fi32++ {
                    v33 := v4[(fi32 * (fi32 >> ((42) & 63)))]
                    var v34 int = (fi32 * 32768)
                    for fi35 := 0; fi35 < 4; fi35++ {
                        v36 := GenS3{F0: v1, F1: (v1 + 167), F2: 1}
                        if (false && v33) {
                            fmt.Println(v34, v33, fi32)
                        } else {
                            fmt.Println(fi35, v2, v34)
                        }
                        v37 := GenS2{F0: v6, F1: 1, F2: ("a" + "foo bar")}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                _ = len(v5)
            } else {
                v2 = (v2 - 2)
            }
        } else {
            fmt.Println(v3, v1, v6)
            var v38 uint64 = 4294967295
            for fi39 := 0; fi39 < 2; fi39++ {
                v38 = (v38 | (v38 / ((42) | 1)))
                var v40 int = (fi39 | fi39)
                _ = fi39
                _ = v40
            }
            _ = v38
        }
    } else {
        if (v6 > (^(v6))) {
            fmt.Println(v2, v3, v6)
            var v41 string = "foo bar"
            _ = v41
        } else {
            delete(v5, 42)
            if false {
                v3 = v3
                var v42 int16 = 1000
                if true {
                    for fi43 := 0; fi43 < 9; fi43++ {
                        var v44 int = (fi43 >> ((65535) & 63))
                        switch v44 {
                        case -1:
                            v6 += v6
                        case 1:
                            v3 = v3
                        default:
                            v3 -= v3
                        }
                        v45 := GenS4{F0: 42, F1: int16(v1), F2: 2147483648}
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    for fi46 := 0; fi46 < 1; fi46++ {
                        for fi47 := 0; fi47 < 19; fi47++ {
                            if true {
                                switch fi46 {
                                case 128:
                                    if (!true) {
                                        v48 := []int{0, (fi47 - 165), -985}
                                        _ = v48
                                    } else {
                                        fmt.Println(v1, v3, v42)
                                    }
                                }
                                v49 := GenS3{F0: v1, F1: 223, F2: 659}
                                _ = v49
                            }
                            v5[(fi46 << ((65535) & 63))] = 42
                            var v50 float32 = (1.5 / (-2.25 - (0.0 / 0.1)))
                            fmt.Println(fi46, v3)
                            _ = fi47
                            _ = v50
                        }
                        _ = fi46
                    }
                } else {
                    v42 -= 1000
                    switch ("hello" + (("\n\t" + "") + ((("hello" + "") + (("foo bar" + "a") + "hello")) + ""))) {
                    case "":
                        v51 := GenS2{F0: v6, F1: 2, F2: ("foo bar" + "")}
                        _ = v51
                    default:
                        for fi52 := 0; fi52 < 4; fi52++ {
                            delete(v4, fi52)
                            v6 = int64(v42)
                            fmt.Println(v1, v3, v1)
                            _ = fi52
                        }
                    }
                    v1 = v1
                }
                _ = v42
            } else {
                var v53 int64 = v6
                delete(v4, 42)
                _ = v53
            }
        }
        if true {
            var v54 int64 = (-(v6))
            var v55 bool = (v54 == v54)
            var v56 string = ("" + (("hello" + ("hello" + ("\n\t" + ((("a" + (("a" + ("" + "hello")) + "hello")) + "") + "")))) + ("foo bar" + "foo bar")))
            _ = v54
            _ = v55
            _ = v56
        }
        var v57 uint = 999
        _ = v57
    }
    if (2 >= 42) {
        var v58 float64 = float64(3.14159)
        var v59 int64 = (v6 % ((v6) | 1))
        _ = v58
        _ = v59
    }
    var v60 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
