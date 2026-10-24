package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint16, p1 int16, p2 int32) int {
    v7 := GenS1{F0: 255}
    var v8 int32 = (p2 * 2147483647)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 224
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 uint8 = (v1 | 127)
    var v5 float32 = 3.14159
    var v6 float64 = (-2.25 + -2.25)
    v9 := GenS1{F0: (v4 * (v4 * (^(v4))))}
    var v10 uint32 = 0
    var v11 uint64 = 0
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint8 = uint8(v11)
        _ = fi12
        _ = v13
    }
    var v14 uint64 = v11
    switch 32768 {
    case 0:
        switch "foo bar" {
        case "a":
            v15 := GenS1{F0: (v4 - v1)}
            _ = v15
        case "foo bar":
            if (v5 != 1e10) {
                for fi16 := 0; fi16 < 5; fi16++ {
                    v17 := GenS1{F0: v1}
                    switch "foo bar" {
                    case "foo bar":
                        v4 = v4
                    case "":
                        v18 := GenS1{F0: v4}
                        _ = v18
                    }
                    _ = fi16
                    _ = v17
                }
            } else {
                if true {
                    v2[-9223372036854775808] = ""
                    v19 := GenS1{F0: v1}
                    var v20 float64 = float64((v5 - v5))
                    _ = v19
                    _ = v20
                } else {
                    v1 += v4
                }
            }
        default:
            v1 = uint8(v11)
        }
    }
    switch genFn0(2, 32767, 11) {
    case 128:
        v4 = v1
    default:
        v21 := []int{}
        _ = v21
    }
    var v22 int8 = 0
    if ((0 > 26) && true) {
        if false {
            v23 := GenS2{F0: v14}
            _ = v23
        }
    } else {
        v24 := GenS2{F0: v11}
        v22 -= v22
        _ = v24
    }
    switch ("" + "") {
    case "a":
        for fi25 := 0; fi25 < 13; fi25++ {
            if ((54 > (v14 + 4294967295)) || (true || ((32767 > 40) && false))) {
                var v26 int64 = -1
                var v27 uint8 = 0
                _ = v26
                _ = v27
            }
            fmt.Println(v5, v14, v1)
            v28 := []int{}
            _ = fi25
            _ = v28
        }
    case "foo bar":
        var v29 int = 32768
        _ = v29
    case "":
        var v30 int16 = 32767
        _ = v30
    default:
        v1 = v1
    }
    v31 := GenS3{F0: 994, F1: 42}
    for fi32 := 0; fi32 < 19; fi32++ {
        for fi33 := 0; fi33 < 20; fi33++ {
            _ = len(v2)
            v34 := []int{genFn0(1, 0, -2147483648), genFn0(uint16(v4), 1000, 1), (fi33 >> ((18446744073709551615) & 63)), (fi33 >> ((603) & 63))}
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    var v35 uint32 = 42
    v6 += (v6 / (v6 / v6))
    v10 = (v10 & (^(v35)))
    var v36 int32 = -936
    if false {
        if ((^(v11)) != (v11 / (((v14 << ((18446744073709551615) & 63))) | 1))) {
            v5 -= v5
            for fi37 := 0; fi37 < 12; fi37++ {
                switch ("hello" + "hello") {
                case "a":
                    var v38 float64 = v6
                    _ = v38
                case "\n\t":
                    var v39 uint8 = 42
                    _ = v39
                }
                for fi40 := 0; fi40 < 17; fi40++ {
                    var v41 uint64 = (v14 % (((^(v14))) | 1))
                    _ = fi40
                    _ = v41
                }
                if false {
                    var v42 int32 = (v36 % ((415) | 1))
                    fmt.Println(v35)
                    fmt.Println(v4)
                    _ = v42
                } else {
                    v43 := []int{525, -2147483648, fi37, fi37}
                    if false {
                        v44 := GenS2{F0: (v11 << ((18446744073709551615) & 63))}
                        _ = v44
                    }
                    var v45 uint16 = 547
                    _ = v43
                    _ = v45
                }
                _ = fi37
            }
            v6 -= -2.25
        }
        fmt.Println(v1)
    } else {
        var v46 int8 = (v22 & v22)
        _ = v46
    }
    if (!false) {
        v5 = v5
        v22 = v22
    } else {
        v14 = (v11 ^ 0)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
