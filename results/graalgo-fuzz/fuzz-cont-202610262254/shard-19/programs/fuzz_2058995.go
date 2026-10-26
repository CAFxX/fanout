package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 2
    v3 := []uint32{}
    var v4 int = 234
    var v5 int16 = -32768
    v6 := GenS1{F0: "", F1: int(v5)}
    if ((-75 >= 1) && (672 >= 496)) {
        if true {
            switch ("hello" + "") {
            case "a":
                if false {
                    v2 -= (v2 % ((v2) | 1))
                }
            case "":
                for fi7 := 0; fi7 < 8; fi7++ {
                    fmt.Println(v5)
                    v8 := GenS2{F0: 65535}
                    switch (v4 + int(v2)) {
                    case 32767:
                        v1[2] = v4
                    default:
                        v4 = fi7
                    }
                    _ = fi7
                    _ = v8
                }
            }
        } else {
            switch -9223372036854775808 {
            case 842:
                var v9 float64 = 0.1
                _ = v9
            case 32767:
                for fi10 := 0; fi10 < 9; fi10++ {
                    for fi11 := 0; fi11 < 2; fi11++ {
                        var v12 uint32 = 0
                        if false {
                            var v13 int8 = 127
                            _ = v13
                        } else {
                            v2 += v2
                            var v14 int8 = int8(v2)
                            _ = v14
                        }
                        delete(v1, 2147483648)
                        _ = fi11
                        _ = v12
                    }
                    var v15 int16 = (^(v5))
                    switch (fi10 + fi10) {
                    case 128:
                        _ = len(v1)
                    case 4294967295:
                        switch (v4 << ((42) & 63)) {
                        case 2147483647:
                            v16 := GenS2{F0: 120}
                            _ = v16
                        }
                    case -578:
                        delete(v1, 4294967295)
                    default:
                        v4 = -2147483648
                    }
                    _ = fi10
                    _ = v15
                }
            default:
                v17 := GenS3{F0: "hello", F1: 439, F2: 2}
                _ = v17
            }
        }
        if (1 < 337) {
            v4 = v4
        }
        var v18 uint8 = 82
        _ = v18
    } else {
        v19 := GenS3{F0: "foo bar", F1: 42, F2: 65}
        _ = len(v1)
        _ = v19
    }
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    v20 := GenS4{F0: v5}
    var v21 int16 = 0
    _ = len(v3)
    var v22 int8 = -93
    v23 := GenS4{F0: 1000}
    v24 := v1[42]
    fmt.Println(v21, v24)
    fmt.Println(v24, v21)
    v21 -= v5
    if (0 < 2) {
        for fi25 := 0; fi25 < 15; fi25++ {
            for fi26 := 0; fi26 < 2; fi26++ {
                v2 = v2
                v2 = v2
                v27 := GenS5{F0: (^(v22))}
                v22 = v22
                _ = fi26
                _ = v27
            }
            v28 := v1[2147483647]
            v4 = (v4 % ((-387) | 1))
            _ = fi25
            _ = v28
        }
        var v29 bool = (false || false)
        v22 += (-(v22))
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
