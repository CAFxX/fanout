package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint; F1 int }
func genFn0() uint {
    var v4 uint8 = 127
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: -989, F1: (fi5 >> ((998) & 63))}
            if (!(!(4294967295 == 678))) {
                switch "foo bar" {
                case "a":
                    var v8 int8 = 1
                    _ = v8
                case "foo bar":
                    v9 := GenS2{F0: 42, F1: (fi6 | (fi5 * (fi5 * -2147483648)))}
                    _ = v9
                default:
                    var v10 int8 = 17
                    _ = v10
                }
                v4 += (v4 * 128)
                var v11 uint16 = 636
                _ = v11
            } else {
                if (!(!false)) {
                    v12 := GenS2{F0: 2, F1: (fi5 >> ((42) & 63))}
                    v13 := GenS1{F0: -279, F1: (fi6 + (fi5 >> ((261) & 63)))}
                    v14 := []int{}
                    _ = v12
                    _ = v13
                    _ = v14
                }
            }
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    if false {
        var v15 uint64 = 0
        var v16 int32 = 42
        _ = v15
        _ = v16
    } else {
        switch 127 {
        case 32767:
            fmt.Println(v4)
        }
        v17 := GenS1{F0: -1, F1: 32768}
        if true {
            v4 = (v4 >> ((42) & 63))
            fmt.Println(v4)
        }
        _ = v17
    }
    _ = v4
    return 65535
}
func genFn1() uint64 {
    switch "\n\t" {
    case "a":
        switch 32768 {
        case 128:
            var v18 uint8 = 2
            _ = v18
        case 1:
            var v19 int64 = 9223372036854775807
            _ = v19
        default:
            if true {
                var v20 int32 = -2147483648
                _ = v20
            } else {
                for fi21 := 0; fi21 < 14; fi21++ {
                    v22 := GenS1{F0: 2147483647, F1: -9223372036854775808}
                    v23 := []int{(fi21 << ((4294967295) & 63)), fi21, fi21}
                    if true {
                        if ((0 >= 587) && (!(-549 > -1))) {
                            var v24 int32 = -2147483648
                            if len(v23) > 0 {
                                v23[0] = 4294967295
                            }
                            _ = v24
                        } else {
                            var v25 string = ("" + "a")
                            _ = v25
                        }
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                if true {
                    for fi26 := 0; fi26 < 8; fi26++ {
                        for fi27 := 0; fi27 < 11; fi27++ {
                            var v28 int16 = 32767
                            if (false && false) {
                                fmt.Println(fi26)
                                v29 := []int{fi27, fi27}
                                _ = v29
                            } else {
                                if true {
                                    var v30 int = 9223372036854775807
                                    _ = v30
                                } else {
                                    v28 = -32768
                                    v28 -= -58
                                }
                                fmt.Println(fi26)
                            }
                            if (true && false) {
                                switch "\n\t" {
                                case "foo bar":
                                    v28 = v28
                                default:
                                    v28 -= 1
                                }
                            }
                            _ = fi27
                            _ = v28
                        }
                        fmt.Println(fi26)
                        _ = fi26
                    }
                }
            }
        }
    }
    v31 := GenS2{F0: 18446744073709551615, F1: 9223372036854775807}
    _ = v31
    return 139
}

func main() {
    v1 := []int{128, 2147483647, -803, 9223372036854775807}
    var v2 uint8 = 0
    var v3 int16 = -1
    v32 := GenS1{F0: 0, F1: 127}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v33 int = -1
    v33 += 0
    if true {
        fmt.Println(v3, v33)
    }
    if ("" < "hello") {
        var v34 string = "foo bar"
        var v35 uint = 42
        v33 = (-(v33))
        _ = v34
        _ = v35
    } else {
        v36 := GenS2{F0: uint(v2), F1: (v33 + (v33 >> ((727) & 63)))}
        _ = v36
    }
    v37 := GenS2{F0: genFn0(), F1: -2147483648}
    fmt.Println(v3, v2, v2)
    var v38 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
}
