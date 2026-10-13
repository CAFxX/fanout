package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 uint32; F1 int16; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    var v3 int64 = -9223372036854775808
    var v4 int8 = -66
    var v5 uint = 137
    if true {
        if false {
            var v6 float32 = 1.5
            _ = v6
        } else {
            fmt.Println(v5)
            var v7 uint = 65535
            for fi8 := 0; fi8 < 8; fi8++ {
                v5 += (^(v7))
                if ((!(v4 >= 100)) || (true || true)) {
                    switch 0 {
                    case 32767:
                        var v9 int16 = -1
                        _ = v9
                    case 977:
                        v5 = 65535
                    case 32768:
                        fmt.Println(v4)
                    }
                    switch "" {
                    case "a":
                        for fi10 := 0; fi10 < 19; fi10++ {
                            v11 := GenS1{F0: 712}
                            if true {
                                if false {
                                    var v12 int64 = (^(v3))
                                    var v13 uint = v5
                                    _ = v12
                                    _ = v13
                                } else {
                                    v14 := GenS2{F0: -87, F1: 199}
                                    _ = v14
                                }
                                var v15 int64 = (v3 << ((v7) & 63))
                                v15 += v3
                                _ = v15
                            }
                            _ = fi10
                            _ = v11
                        }
                    case "\n\t":
                        v16 := GenS3{F0: v5}
                        _ = v16
                    }
                    switch (-(fi8)) {
                    case 1:
                        v7 = v5
                    case 4294967295:
                        if (!(true || (false || false))) {
                            v1[155] = 1
                            v17 := GenS4{F0: ("foo bar" >= "foo bar")}
                            _ = v17
                        }
                    case 32767:
                        var v18 uint = (v5 + 1)
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 19; fi19++ {
                            v4 -= v4
                            v4 = (v4 - -128)
                            _ = fi19
                        }
                    }
                } else {
                    var v20 bool = (4294967295 <= 608)
                    _ = v20
                }
                fmt.Println(v7, fi8, v5)
                _ = fi8
            }
            _ = v7
        }
        v21 := GenS5{F0: false, F1: (v5 / ((v5) | 1))}
        _ = v21
    } else {
        v3 = (v3 % (((v3 + (v3 * v3))) | 1))
        delete(v2, "foo bar")
        v5 -= (v5 - v5)
    }
    v3 = (v3 >> ((18446744073709551615) & 63))
    var v22 uint = v5
    switch "\n\t" {
    case "\n\t":
        v23 := GenS4{F0: (1 > 349)}
        _ = v23
    case "a":
        for fi24 := 0; fi24 < 1; fi24++ {
            fmt.Println(v5)
            var v25 int64 = -9223372036854775808
            _ = fi24
            _ = v25
        }
    case "":
        v1[uint32(v5)] = 2147483647
    }
    v26 := []int{-606, -726, int(v4), 32767}
    v27 := GenS6{F0: 373, F1: 1, F2: 0}
    switch 4294967295 {
    case 0:
        var v28 bool = (1 > 2)
        _ = v28
    case 1:
        v2["a"] = -9223372036854775808
    default:
        v5 = 255
    }
    var v29 uint16 = 65535
    var v30 uint64 = 0
    v2["\n\t"] = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
