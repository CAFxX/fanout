package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint16; F2 bool }
type GenS3 struct { F0 uint64; F1 uint64; F2 string }
func genFn0() uint16 {
    var v5 int64 = -9223372036854775808
    v5 = (v5 & (v5 - 99))
    _ = v5
    return 255
}
func genFn1() uint16 {
    var v6 int16 = 1000
    fmt.Println(v6)
    if (!false) {
        var v7 int = 32768
        v7 += (v7 << ((18446744073709551615) & 63))
        var v8 int8 = 51
        _ = v7
        _ = v8
    }
    _ = v6
    return 32767
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    var v4 uint64 = 4294967295
    var v9 int = 9223372036854775807
    var v10 uint16 = genFn0()
    var v11 float32 = 1.5
    var v12 string = (("hello" + ("hello" + ("hello" + "foo bar"))) + "")
    v4 = (^(v4))
    var v13 int32 = -789
    v14 := GenS1{F0: 42, F1: (v12 + "foo bar"), F2: 0}
    var v15 uint32 = 65535
    v10 -= (^(v10))
    if true {
        for fi16 := 0; fi16 < 15; fi16++ {
            fmt.Println(fi16, v4, v9)
            _ = fi16
        }
        var v17 string = ("hello" + ("hello" + v12))
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := GenS1{F0: uint(v10), F1: (v12 + v12), F2: 210}
        for fi20 := 0; fi20 < 4; fi20++ {
            v12 = v12
            var v21 bool = false
            v21 = false
            v4 = v4
            switch "\n\t" {
            case "a":
                v9 = (fi20 / (((fi18 & (v9 | fi18))) | 1))
            case "":
                v22 := GenS2{F0: 1, F1: genFn1(), F2: (v21 && true)}
                _ = v22
            case "hello":
                var v23 float64 = 0.0
                _ = v23
            default:
                var v24 int32 = 42
                _ = v24
            }
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        fmt.Println(v4)
        _ = fi25
    }
    v26 := []int{(v9 | (v9 << ((65535) & 63)))}
    if true {
        fmt.Println(v10, v11, v15)
        fmt.Println(v11, v1, v13)
    } else {
        v27 := GenS2{F0: v1, F1: 2, F2: false}
        if true {
            for fi28 := 0; fi28 < 10; fi28++ {
                if true {
                    _ = len(v3)
                    for fi29 := 0; fi29 < 7; fi29++ {
                        delete(v3, fi28)
                        _ = fi29
                    }
                    for fi30 := 0; fi30 < 4; fi30++ {
                        v31 := GenS1{F0: 990, F1: v12, F2: 53}
                        v9 = (^(fi30))
                        _ = fi30
                        _ = v31
                    }
                } else {
                    for fi32 := 0; fi32 < 20; fi32++ {
                        v1 = (^(v1))
                        _ = fi32
                    }
                }
                v15 = (v15 ^ 65535)
                v10 += (v10 + 1)
                _ = fi28
            }
        }
        v33 := GenS2{F0: v1, F1: 876, F2: (((((v12 + "a") + v12) + v12) + v12) >= v12)}
        _ = v27
        _ = v33
    }
    if false {
        var v34 uint64 = v4
        for fi35 := 0; fi35 < 20; fi35++ {
            if (768 != v4) {
                v36 := GenS2{F0: (^(v1)), F1: 32767, F2: (v12 <= v12)}
                _ = len(v26)
                _ = v36
            } else {
                if false {
                    switch "hello" {
                    case "\n\t":
                        if true {
                            var v37 int16 = v1
                            if false {
                                v2[(v15 + v15)] = 4294967295
                                v11 += v11
                            }
                            if false {
                                v38 := GenS3{F0: (v34 - 502), F1: v34, F2: (("" + "hello") + "")}
                                fmt.Println(v4, v15)
                                var v39 int32 = (v13 | -2147483648)
                                _ = v38
                                _ = v39
                            }
                            _ = v37
                        }
                    case "a":
                        fmt.Println(fi35, v9, v15)
                    }
                    var v40 int16 = 663
                    v40 -= (v40 / ((v1) | 1))
                    _ = v40
                } else {
                    var v41 uint8 = 214
                    _ = v41
                }
            }
            var v42 uint16 = uint16(v15)
            _ = fi35
            _ = v42
        }
        v11 = float32(1.5)
        _ = v34
    } else {
        v43 := v3[v9]
        v44 := GenS2{F0: v1, F1: (v10 * genFn0()), F2: false}
        var v45 string = ("a" + "hello")
        _ = v43
        _ = v44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v9
}
