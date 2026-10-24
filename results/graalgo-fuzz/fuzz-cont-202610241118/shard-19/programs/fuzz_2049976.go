package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 bool }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 bool; F1 int8; F2 uint }
type GenS4 struct { F0 bool }
func genFn0(p0 int16) uint16 {
    p0 = p0
    p0 += p0
    for fi4 := 0; fi4 < 9; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            if true {
                var v6 int32 = 411
                var v7 int = fi4
                _ = v6
                _ = v7
            }
            if (p0 >= p0) {
                var v8 int = (fi4 - fi4)
                v8 += (fi4 ^ -1)
                _ = v8
            }
            v9 := []int{-932, 965}
            p0 = (p0 >> ((65535) & 63))
            _ = fi5
            _ = v9
        }
        _ = fi4
    }
    p0 += (-(p0))
    _ = p0
    return 42
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]int)
    var v3 uint16 = 2
    v3 += (^(v3))
    var v10 int64 = 343
    v11 := GenS1{F0: uint64(v3), F1: (^(v10)), F2: (-175 < 0)}
    for fi12 := 0; fi12 < 12; fi12++ {
        switch (fi12 - (fi12 * (fi12 >> ((1) & 63)))) {
        case 32767:
            if false {
                switch "hello" {
                case "a":
                    var v13 float32 = (1e10 + 3.14159)
                    _ = v13
                }
                v1 = v1
            }
        case 4294967295:
            v14 := GenS2{F0: v3, F1: (v3 * v3), F2: 42}
            _ = v14
        default:
            fmt.Println(v3)
        }
        fmt.Println(v10)
        if (412 <= 442) {
            v2[(("" + "") + ("\n\t" + ("hello" + "\n\t")))] = fi12
            v1 = (v1 * (^(v1)))
            for fi15 := 0; fi15 < 10; fi15++ {
                fmt.Println(fi12)
                v10 -= (v10 % (((^(v10))) | 1))
                switch (fi15 * (-(fi12))) {
                case -9223372036854775808:
                    v1 += (v1 >> ((42) & 63))
                case 1:
                    v16 := GenS1{F0: 18446744073709551615, F1: (v10 ^ -872), F2: (!true)}
                    _ = v16
                default:
                    fmt.Println(v3, v10)
                }
                _ = fi15
            }
        }
        _ = fi12
    }
    if true {
        v17 := GenS1{F0: 4294967295, F1: (v10 << ((904) & 63)), F2: (true || (0.0 >= ((float32(0.0) / (1e10 * 1e10)) + ((((1.5 + 1.5) / 1.0) + 1.5) - float32(0.1)))))}
        if true {
            var v18 int64 = (v10 << ((945) & 63))
            v1 += 42
            _ = v18
        }
        if (435 != 9223372036854775807) {
            if false {
                var v19 uint = 2
                for fi20 := 0; fi20 < 14; fi20++ {
                    v21 := GenS3{F0: false, F1: 1, F2: v19}
                    if (1 >= 127) {
                        fmt.Println(v10, v3, fi20)
                        for fi22 := 0; fi22 < 9; fi22++ {
                            var v23 int8 = 42
                            v24 := v2["\n\t"]
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    }
                    _ = fi20
                    _ = v21
                }
                _ = v19
            } else {
                var v25 int = 127
                v26 := GenS4{F0: true}
                if true {
                    var v27 uint64 = 4294967295
                    _ = v27
                }
                _ = v25
                _ = v26
            }
            v10 = v10
            v10 = v10
        }
        _ = v17
    }
    v10 = v10
    var v28 float64 = 1e10
    v29 := []int{32767, 543}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
