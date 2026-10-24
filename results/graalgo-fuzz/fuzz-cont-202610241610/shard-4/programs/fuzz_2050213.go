package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = float64((0.0 / (0.0 - 1.5)))
    var v3 int8 = 100
    var v4 float64 = v2
    var v5 int8 = -128
    switch 32767 {
    case -9223372036854775808:
        var v6 int16 = 862
        _ = v6
    case -1:
        if true {
            v4 += v4
        } else {
            var v7 int64 = -436
            var v8 int16 = 247
            if true {
                for fi9 := 0; fi9 < 16; fi9++ {
                    v10 := []int{(-(fi9)), (fi9 | fi9), -2147483648, -167}
                    var v11 int64 = (^(v7))
                    v7 = (v11 - -9223372036854775808)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                v3 += (v3 * v3)
            } else {
                var v12 uint64 = 42
                var v13 int64 = (v7 & 352)
                fmt.Println(v3, v12)
                _ = v12
                _ = v13
            }
            _ = v7
            _ = v8
        }
    case 32768:
        fmt.Println(v3)
    }
    switch "foo bar" {
    case "hello":
        var v14 float64 = v2
        _ = v14
    case "":
        var v15 uint64 = 80
        _ = v15
    default:
        v16 := GenS1{F0: 936}
        _ = v16
    }
    v17 := GenS1{F0: 710}
    if (1 < 65535) {
        var v18 float32 = 0.1
        if (42 != 4294967295) {
            var v19 int = 127
            if false {
                delete(v1, 2)
                v5 -= 27
                for fi20 := 0; fi20 < 2; fi20++ {
                    for fi21 := 0; fi21 < 15; fi21++ {
                        var v22 bool = (v5 == v5)
                        if v22 {
                            fmt.Println(fi20, fi21)
                        } else {
                            v23 := v1[402]
                            v18 -= 0.0
                            _ = v23
                        }
                        _ = fi21
                        _ = v22
                    }
                    var v24 uint = 4294967295
                    _ = fi20
                    _ = v24
                }
            } else {
                v25 := []int{-9223372036854775808, 962}
                fmt.Println(v2, v2)
                _ = v25
            }
            _ = v19
        }
        var v26 float64 = v4
        _ = v18
        _ = v26
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        for fi28 := 0; fi28 < 4; fi28++ {
            v3 -= -1
            switch fi28 {
            case 42:
                var v29 string = "\n\t"
                _ = v29
            case -347:
                v1[1] = (fi28 - fi28)
            }
            _ = fi28
        }
        _ = fi27
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        for fi31 := 0; fi31 < 15; fi31++ {
            v3 -= (v3 | v5)
            if true {
                v3 = -1
            }
            v3 += v3
            _ = fi31
        }
        switch "hello" {
        case "\n\t":
            if false {
                v32 := GenS1{F0: 2}
                for fi33 := 0; fi33 < 5; fi33++ {
                    var v34 int = (fi33 << ((18446744073709551615) & 63))
                    v35 := GenS1{F0: 65535}
                    if false {
                        var v36 int8 = v5
                        v34 -= fi30
                        v3 += v3
                        fmt.Println(v2, v5)
                        _ = v36
                    } else {
                        if (345 < -2147483648) {
                            v5 = (v5 / ((v5) | 1))
                            v4 += -2.25
                        } else {
                            v5 -= v3
                        }
                    }
                    switch (^(fi33)) {
                    case 31:
                        switch "foo bar" {
                        case "\n\t":
                            v3 += (^(v3))
                        case "foo bar":
                            v34 += (fi33 + fi30)
                        case "":
                            v3 = (v5 - (v5 | v5))
                        }
                    default:
                        if (false || (93 != 255)) {
                            if false {
                                v2 = ((v4 - v4) + v4)
                                switch ("a" + "hello") {
                                case "foo bar":
                                    var v37 float32 = (0.1 + 0.0)
                                    _ = v37
                                case "\n\t":
                                    v38 := GenS2{F0: 1}
                                    _ = v38
                                default:
                                    v39 := []int{128}
                                    _ = v39
                                }
                                v3 -= v5
                                var v40 uint8 = 128
                                _ = v40
                            }
                        } else {
                            var v41 int64 = 42
                            v5 -= (^(v5))
                            var v42 uint = 42
                            _ = v41
                            _ = v42
                        }
                    }
                    if true {
                        v43 := GenS2{F0: 922}
                        _ = v43
                    }
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                _ = v32
            }
        }
        _ = fi30
    }
    v2 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
