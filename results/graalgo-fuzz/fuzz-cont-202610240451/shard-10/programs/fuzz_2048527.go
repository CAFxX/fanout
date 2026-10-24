package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 bool }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint = 42
    var v2 uint32 = 42
    var v3 uint = (^(v1))
    var v4 uint64 = 166
    v5 := []int{32768, 32767}
    fmt.Println(v3)
    v2 += (v2 % ((v2) | 1))
    if (true || true) {
        var v6 string = ("a" + "")
        var v7 uint = (v1 << ((v3) & 63))
        v8 := GenS1{F0: 100, F1: 514, F2: (!false)}
        _ = v6
        _ = v7
        _ = v8
    }
    v2 -= 1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    v3 += v1
    if (!(1 <= -32768)) {
        for fi9 := 0; fi9 < 20; fi9++ {
            v3 = 42
            _ = fi9
        }
        if false {
            if (-1 > -1) {
                if len(v5) > 0 {
                    v5[0] = int(v2)
                }
                var v10 int64 = 1
                _ = v10
            } else {
                v1 -= v1
                if true {
                    if (false && false) {
                        v4 -= 42
                    }
                    var v11 float64 = (((1.5 * 1.5) / 1.0) * 1.5)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 8; fi12++ {
                        v13 := GenS2{F0: 346}
                        _ = fi12
                        _ = v13
                    }
                    if len(v5) > 0 {
                        v5[0] = 4294967295
                    }
                }
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 int16 = -1
                _ = fi14
                _ = v15
            }
            var v16 uint64 = (^(v4))
            _ = v16
        } else {
            v17 := GenS1{F0: 100, F1: 2, F2: true}
            fmt.Println(v1, v3, v2)
            for fi18 := 0; fi18 < 2; fi18++ {
                fmt.Println(v3)
                v2 -= uint32(v4)
                _ = fi18
            }
            _ = v17
        }
        if (false || true) {
            var v19 uint64 = (v4 - v4)
            if len(v5) > 0 {
                v5[0] = int(v4)
            }
            v20 := GenS2{F0: uint(v19)}
            _ = v19
            _ = v20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
