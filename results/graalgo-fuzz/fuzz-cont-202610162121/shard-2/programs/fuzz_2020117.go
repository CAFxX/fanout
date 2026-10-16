package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int; F1 int64 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    var v4 string = "\n\t"
    v5 := make(map[string]int)
    var v6 uint8 = 255
    var v7 uint = 2
    v8 := []int{-774, 289, -9223372036854775808, 342}
    v9 := GenS1{F0: 0}
    if (!true) {
        var v10 int8 = -18
        _ = v10
    } else {
        var v11 uint16 = 65535
        if false {
            if len(v8) > 0 {
                v8[0] = -748
            }
            var v12 string = v4
            v13 := GenS1{F0: 666}
            _ = v12
            _ = v13
        }
        _ = v11
    }
    fmt.Println(v7)
    v6 += (v6 + v6)
    for fi14 := 0; fi14 < 16; fi14++ {
        if false {
            v15 := GenS1{F0: int64(fi14)}
            _ = v15
        } else {
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 int16 = v1
                _ = fi16
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 13; fi18++ {
            v7 -= v7
            var v19 uint16 = 42
            _ = fi18
            _ = v19
        }
        v7 -= (v7 % ((532) | 1))
        _ = fi14
    }
    var v20 uint = v7
    if (-1 == -9223372036854775808) {
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 string = "foo bar"
            var v23 uint8 = v6
            v23 -= (v23 / (((v6 ^ (v23 | (v23 / ((v6) | 1))))) | 1))
            _ = fi21
            _ = v22
            _ = v23
        }
        if (-561 <= 128) {
            for fi24 := 0; fi24 < 14; fi24++ {
                v20 -= v20
                v25 := GenS2{F0: 556, F1: 988}
                if len(v8) > 0 {
                    v8[0] = fi24
                }
                _ = fi24
                _ = v25
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
    } else {
        var v26 int16 = -1
        _ = v26
    }
    if true {
        if true {
            if false {
                v7 -= v20
            } else {
                for fi27 := 0; fi27 < 18; fi27++ {
                    v7 += 0
                    if (42 > uint16(v20)) {
                        v6 -= 191
                    } else {
                        v2[("\n\t" + v4)] = fi27
                        var v28 uint8 = v6
                        var v29 int32 = 2147483647
                        _ = v28
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 20; fi30++ {
                        var v31 float64 = 0.1
                        _ = fi30
                        _ = v31
                    }
                    _ = fi27
                }
            }
            v6 = v6
            if false {
                var v32 uint = (v7 << ((966) & 63))
                if (18 == 100) {
                    var v33 bool = true
                    _ = v33
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = v32
            }
        }
    } else {
        v34 := GenS3{F0: -2147483648, F1: -155}
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
