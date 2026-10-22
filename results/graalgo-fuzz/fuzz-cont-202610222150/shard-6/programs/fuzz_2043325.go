package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }
type GenS2 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint64, p1 int16) int8 {
    for fi5 := 0; fi5 < 9; fi5++ {
        if false {
            p1 = p1
            var v6 uint64 = 1
            _ = v6
        } else {
            p0 -= (p0 ^ 18446744073709551615)
        }
        var v7 float64 = float64(-2.25)
        _ = fi5
        _ = v7
    }
    var v8 uint8 = 38
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = false
    var v3 uint64 = 42
    v4 := make(map[uint32]int)
    var v9 uint16 = 763
    v9 -= 0
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 string = "hello"
        v9 = (v9 >> ((983) & 63))
        _ = len(v1)
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: -46, F1: -34, F2: -1}
    if (false || false) {
        for fi13 := 0; fi13 < 13; fi13++ {
            fmt.Println(fi13)
            switch (^(fi13)) {
            case 0:
                fmt.Println(v3, v3)
            default:
                fmt.Println(v3, fi13)
            }
            var v14 int16 = 32767
            _ = fi13
            _ = v14
        }
    }
    if ((!((v9 % ((v9) | 1)) <= v9)) || ((32 >= 40) || true)) {
        delete(v1, 394)
        switch -486 {
        case 1:
            if v2 {
                for fi15 := 0; fi15 < 12; fi15++ {
                    v16 := GenS1{F0: (fi15 >> ((42) & 63)), F1: -9223372036854775808, F2: 42}
                    _ = fi15
                    _ = v16
                }
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 int16 = int16(v3)
                    for fi19 := 0; fi19 < 16; fi19++ {
                        switch "a" {
                        case "":
                            switch "" {
                            case "\n\t":
                                var v20 int8 = 127
                                _ = v20
                            }
                        }
                        switch "\n\t" {
                        case "a":
                            if true {
                                v21 := GenS1{F0: (fi19 / (((fi17 / ((fi19) | 1))) | 1)), F1: 42, F2: genFn0(v3, (v18 + 525))}
                                if true {
                                    v3 = (v3 | (v3 << ((990) & 63)))
                                }
                                var v22 uint16 = 2
                                _ = v21
                                _ = v22
                            }
                        case "\n\t":
                            v23 := GenS2{F0: 1, F1: 1}
                            _ = v23
                        case "foo bar":
                            v3 += v3
                        }
                        var v24 bool = (!v2)
                        _ = fi19
                        _ = v24
                    }
                    _ = fi17
                    _ = v18
                }
            } else {
                var v25 float32 = 0.1
                var v26 uint64 = (v3 - v3)
                _ = v25
                _ = v26
            }
        case 32768:
            for fi27 := 0; fi27 < 6; fi27++ {
                v3 -= (v3 & uint64(v9))
                var v28 bool = (!(!v2))
                _ = fi27
                _ = v28
            }
        case 32767:
            v3 = v3
        default:
            if v2 {
                for fi29 := 0; fi29 < 3; fi29++ {
                    fmt.Println(v9, v2)
                    _ = fi29
                }
                fmt.Println(v3)
                v9 = (v9 << ((1) & 63))
            }
        }
        v9 += v9
    }
    v30 := GenS2{F0: -1, F1: 0}
    for fi31 := 0; fi31 < 16; fi31++ {
        for fi32 := 0; fi32 < 17; fi32++ {
            v3 -= v3
            _ = fi32
        }
        _ = fi31
    }
    if ((0.1 * 1.5) == -2.25) {
        var v33 int16 = -309
        for fi34 := 0; fi34 < 12; fi34++ {
            var v35 int32 = -893
            v4[908] = int(v33)
            _ = len(v4)
            _ = fi34
            _ = v35
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            var v37 int16 = 1000
            _ = fi36
            _ = v37
        }
        _ = v33
    } else {
        var v38 int16 = -32768
        _ = v38
    }
    v2 = false
    v9 += (v9 + uint16(v3))
    for fi39 := 0; fi39 < 10; fi39++ {
        _ = len(v4)
        fmt.Println(v3)
        fmt.Println(v3, fi39, v9)
        _ = fi39
    }
    var v40 float32 = 1e10
    var v41 bool = (!(v2 && (-1 <= 32767)))
    if v41 {
        v42 := GenS1{F0: 32767, F1: 9223372036854775807, F2: 42}
        v9 -= (v9 ^ (v9 % (((^(v9))) | 1)))
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v41
    _ = v9
}
