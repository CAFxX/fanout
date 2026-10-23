package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint32 }
type GenS2 struct { F0 int32; F1 int64; F2 bool }
func genFn0() uint16 {
    v7 := []int{128}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = v7
    return 42
}
func genFn1(p0 int) int64 {
    if true {
        var v8 int16 = 764
        if (!false) {
            var v9 uint = 813
            _ = v9
        }
        _ = v8
    } else {
        if (-100 != 42) {
            v10 := []int{1, p0, (p0 ^ (p0 + p0)), p0}
            if (274 > 2147483648) {
                if false {
                    p0 += 127
                    for fi11 := 0; fi11 < 3; fi11++ {
                        p0 = 1
                        if (117 > 42) {
                            p0 = (p0 - (p0 & -9223372036854775808))
                        }
                        for fi12 := 0; fi12 < 18; fi12++ {
                            var v13 int32 = 1
                            p0 -= p0
                            switch -400 {
                            case 4294967295:
                                p0 = fi11
                            default:
                                p0 = -9223372036854775808
                            }
                            var v14 uint = 4294967295
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        _ = fi11
                    }
                    p0 = (p0 + p0)
                }
                _ = len(v10)
                var v15 int8 = 127
                _ = v15
            }
            var v16 int = (^(p0))
            _ = v10
            _ = v16
        }
        p0 = -1
        var v17 int64 = 9223372036854775807
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        p0 -= 42
        fmt.Println(p0, p0)
        _ = fi18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        fmt.Println(fi19, fi19)
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := []int{}
            v22 := GenS1{F0: -16, F1: 1, F2: 65535}
            _ = len(v21)
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    if false {
        var v23 int = (p0 % ((p0) | 1))
        _ = v23
    }
    _ = p0
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 42
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    var v5 int16 = 1000
    var v6 uint64 = (v2 % ((1) | 1))
    v5 += v5
    var v24 string = ""
    for fi25 := 0; fi25 < 12; fi25++ {
        for fi26 := 0; fi26 < 15; fi26++ {
            v24 = v24
            fmt.Println(fi25, v24)
            var v27 float32 = 1e10
            _ = fi26
            _ = v27
        }
        v28 := GenS2{F0: -2147483648, F1: genFn1((fi25 & fi25)), F2: true}
        _ = fi25
        _ = v28
    }
    var v29 int = 2147483647
    for fi30 := 0; fi30 < 10; fi30++ {
        fmt.Println(v6, v29, v24)
        _ = fi30
    }
    if true {
        fmt.Println(v5)
    } else {
        if (1 > 4294967295) {
            var v31 int = (v29 << ((874) & 63))
            _ = v31
        } else {
            v32 := GenS2{F0: 42, F1: -9223372036854775808, F2: true}
            _ = v32
        }
        switch v24 {
        case "hello":
            var v33 int32 = 2147483647
            _ = v33
        case "":
            fmt.Println(v5)
        default:
            v5 = v5
        }
    }
    fmt.Println(v2, v24, v29)
    var v34 uint = 42
    if false {
        for fi35 := 0; fi35 < 19; fi35++ {
            if (!false) {
                var v36 uint64 = 524
                _ = len(v3)
                for fi37 := 0; fi37 < 5; fi37++ {
                    if (!(true || false)) {
                        v38 := GenS2{F0: 136, F1: 0, F2: true}
                        fmt.Println(v2)
                        v6 += v36
                        _ = v38
                    } else {
                        var v39 int8 = -128
                        if false {
                            v40 := []int{(v29 | 407), fi37}
                            fmt.Println(fi37, v34, fi35)
                            var v41 int32 = 0
                            v24 = v24
                            v34 -= 65535
                            _ = v40
                            _ = v41
                        }
                        _ = v39
                    }
                    switch "\n\t" {
                    case "foo bar":
                        var v42 int8 = int8(v6)
                        _ = v42
                    case "\n\t":
                        v6 = (v36 | 791)
                    default:
                        v5 += v5
                    }
                    v43 := []int{}
                    _ = fi37
                    _ = v43
                }
                _ = v36
            }
            v29 -= 1
            switch (v29 + fi35) {
            case -103:
                var v44 bool = (true || false)
                _ = v44
            }
            _ = fi35
        }
    }
    v34 += v34
    switch (v29 - (v29 | (v29 - (v29 << (((v34 | (v34 ^ v34))) & 63))))) {
    case 2147483647:
        var v45 uint8 = uint8(v34)
        _ = v45
    case -1:
        v46 := GenS1{F0: 100, F1: 0, F2: 1}
        _ = v46
    case 128:
        v6 = (^(v6))
    default:
        var v47 float64 = float64((-2.25 * 0.1))
        _ = v47
    }
    fmt.Println(v34, v24)
    var v48 int16 = (v5 * (v5 % (((^(v5))) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
