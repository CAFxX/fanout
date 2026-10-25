package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 252
    var v3 float64 = (0.1 + -2.25)
    var v4 bool = false
    var v5 bool = false
    var v6 int16 = 0
    var v7 int64 = 1
    if (v5 && v5) {
        v8 := GenS1{F0: v6, F1: 240}
        _ = v8
    } else {
        v4 = v4
        v9 := GenS2{F0: 274}
        if (v5 && (1 < 0)) {
            var v10 int32 = -1
            _ = v10
        }
        _ = v9
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        switch -1 {
        case 127:
            if v4 {
                switch "a" {
                case "":
                    v2 = (v2 << ((18446744073709551615) & 63))
                }
            } else {
                v12 := GenS3{F0: (v6 / ((636) | 1))}
                _ = v12
            }
        case 32768:
            v13 := GenS3{F0: 729}
            _ = v13
        case 1:
            for fi14 := 0; fi14 < 13; fi14++ {
                v1[""] = (v3 >= v3)
                v3 -= float64(1e10)
                var v15 int16 = (v6 | v6)
                _ = fi14
                _ = v15
            }
        default:
            v5 = (761 < 18446744073709551615)
        }
        v16 := GenS3{F0: (v6 & (v6 * v6))}
        if v5 {
            var v17 uint64 = 18446744073709551615
            var v18 uint16 = (v2 * v2)
            var v19 bool = v5
            _ = v17
            _ = v18
            _ = v19
        } else {
            switch ("a" + (("a" + ("a" + "hello")) + "")) {
            case "foo bar":
                for fi20 := 0; fi20 < 12; fi20++ {
                    switch "foo bar" {
                    case "\n\t":
                        if (v4 && true) {
                            v21 := v1["foo bar"]
                            _ = v21
                        } else {
                            if (!false) {
                                var v22 uint8 = 255
                                _ = v22
                            } else {
                                fmt.Println(v3)
                                var v23 uint64 = 1
                                _ = v23
                            }
                            v5 = v4
                            v24 := v1["hello"]
                            _ = v24
                        }
                    default:
                        v25 := GenS1{F0: 0, F1: uint8(v2)}
                        _ = v25
                    }
                    v3 += 0.1
                    if v4 {
                        var v26 bool = v4
                        _ = v26
                    }
                    var v27 uint8 = 127
                    var v28 int64 = v7
                    _ = fi20
                    _ = v27
                    _ = v28
                }
            }
        }
        _ = fi11
        _ = v16
    }
    v2 = (v2 & 32767)
    v29 := GenS4{F0: 2}
    v30 := GenS3{F0: (v6 - (^(v6)))}
    var v31 int8 = -128
    for fi32 := 0; fi32 < 10; fi32++ {
        if (((fi32 >> ((65535) & 63)) >= -485) && true) {
            if (-2147483648 > -303) {
                v33 := GenS3{F0: (^(v6))}
                _ = v33
            } else {
                if true {
                    v6 = v6
                    var v34 float64 = v3
                    fmt.Println(v34, v7)
                    _ = v34
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        v6 = v6
                    case "a":
                        var v35 uint64 = 42
                        _ = v35
                    default:
                        fmt.Println(v2, v4)
                    }
                }
                v6 += (v6 & v6)
                fmt.Println(v31, v4)
            }
            v6 -= v6
        }
        switch (("" + "a") + "\n\t") {
        case "hello":
            v3 -= 0.0
        }
        for fi36 := 0; fi36 < 7; fi36++ {
            switch ("hello" + "foo bar") {
            case "hello":
                v37 := []int{fi32}
                _ = v37
            case "":
                v6 += (v6 << ((0) & 63))
            case "foo bar":
                fmt.Println(v3, fi32)
            default:
                v4 = v5
            }
            _ = fi36
        }
        _ = fi32
    }
    v2 += 0
    var v38 uint = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
