package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int8 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    var v3 bool = false
    v3 = v3
    var v4 int32 = -1
    v4 -= (v4 - -1)
    if (!(2 == 255)) {
        if (1.5 <= 3.14159) {
            v1[4294967295] = (!(true && (v3 && true)))
        }
        v5 := []int{32767}
        v6 := GenS1{F0: 1, F1: -1, F2: -1}
        _ = v5
        _ = v6
    }
    fmt.Println(v3, v4)
    if (4294967295 <= 4294967295) {
        var v7 string = "foo bar"
        v8 := GenS1{F0: 18446744073709551615, F1: 880, F2: -127}
        _ = v7
        _ = v8
    }
    fmt.Println(v4, v4)
    var v9 int8 = -77
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 float64 = -2.25
            var v13 uint64 = 2
            var v14 uint32 = 42
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        v4 = v4
        if (1 <= 68) {
            if true {
                var v15 int8 = 85
                fmt.Println(fi10)
                _ = v15
            }
            var v16 int64 = -401
            var v17 int32 = (v4 | v4)
            _ = v16
            _ = v17
        } else {
            v18 := GenS2{F0: 2147483647}
            _ = v18
        }
        _ = fi10
    }
    var v19 int64 = 1
    var v20 uint = 2
    delete(v1, 42)
    if v3 {
        delete(v2, 42)
        var v21 bool = v3
        delete(v2, 592)
        _ = v21
    } else {
        v9 -= 127
        v22 := v2[1]
        _ = v22
    }
    v9 += (v9 << ((v20) & 63))
    if v3 {
        var v23 uint64 = 4294967295
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 int = (fi24 >> ((1) & 63))
            switch v25 {
            case 32768:
                v9 -= (v9 & v9)
            case 2147483647:
                v19 -= 1
            default:
                if ((!true) && ((^(v4)) > (-(v4)))) {
                    for fi26 := 0; fi26 < 2; fi26++ {
                        v19 -= v19
                        _ = fi26
                    }
                    switch fi24 {
                    case 42:
                        for fi27 := 0; fi27 < 9; fi27++ {
                            v28 := GenS2{F0: 227}
                            fmt.Println(v25, fi27)
                            v9 = (-(v9))
                            _ = fi27
                            _ = v28
                        }
                    default:
                        for fi29 := 0; fi29 < 14; fi29++ {
                            var v30 uint8 = 2
                            if (!(!(!(uint32(v30) == 65535)))) {
                                v25 = (fi24 - (v25 + (fi24 * (fi24 % (((^(v25))) | 1)))))
                                var v31 uint64 = 4294967295
                                _ = v31
                            }
                            _ = fi29
                            _ = v30
                        }
                    }
                }
            }
            var v32 string = "a"
            _ = fi24
            _ = v25
            _ = v32
        }
        for fi33 := 0; fi33 < 13; fi33++ {
            fmt.Println(v4, v9)
            fmt.Println(fi33)
            v3 = v3
            _ = fi33
        }
        _ = v23
    }
    v2[0] = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
