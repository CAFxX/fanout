package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 bool; F2 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 int; F2 int32 }
func genFn0(p0 int64) bool {
    p0 -= -541
    v6 := GenS1{F0: 42}
    _ = p0
    _ = v6
    return (2 > 2147483647)
}
func genFn1(p0 int, p1 int32) uint16 {
    v7 := []int{}
    p1 = (p1 + p1)
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint32 = 2
    v3 := []int{}
    var v4 bool = true
    var v5 bool = v4
    for fi8 := 0; fi8 < 5; fi8++ {
        if v5 {
            var v9 uint16 = 555
            var v10 uint8 = 127
            _ = v10
            _ = v9
        }
        _ = fi8
    }
    var v11 int = 370
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 int64 = -9223372036854775808
        switch "a" {
        case "a":
            if true {
                var v14 float64 = (1e10 / -2.25)
                _ = v14
            } else {
                var v15 int16 = -1
                _ = v15
            }
        }
        switch (v11 | (v11 - (fi12 << ((657) & 63)))) {
        case 915:
            if v5 {
                switch 42 {
                case 293:
                    v11 = (v11 << ((42) & 63))
                case 9223372036854775807:
                    switch "foo bar" {
                    case "foo bar":
                        if v4 {
                            for fi16 := 0; fi16 < 2; fi16++ {
                                var v17 uint = 162
                                if (!(int8(v17) >= 127)) {
                                    v18 := GenS1{F0: 0}
                                    _ = v18
                                }
                                v19 := GenS2{F0: 42}
                                var v20 int32 = v1
                                var v21 int = v11
                                _ = fi16
                                _ = v17
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                        } else {
                            v22 := GenS3{F0: -2147483648, F1: true, F2: ("\n\t" + "")}
                            _ = v22
                        }
                    case "hello":
                        var v23 float64 = (1e10 - 0.0)
                        _ = v23
                    default:
                        v4 = (18446744073709551615 < 873)
                    }
                default:
                    fmt.Println(v11, v11)
                }
            } else {
                v11 += (v11 % (((v11 - (-(v11)))) | 1))
                v24 := GenS4{F0: 2}
                v2 += 4294967295
                _ = v24
            }
        case -9223372036854775808:
            v2 += (v2 ^ 65535)
        case 9223372036854775807:
            v11 = 32767
        default:
            if true {
                if len(v3) > 0 {
                    v3[0] = (fi12 << ((52) & 63))
                }
            }
        }
        _ = fi12
        _ = v13
    }
    var v25 int64 = 0
    v2 -= v2
    v26 := GenS5{F0: 939, F1: (v11 - -250), F2: -2147483648}
    for fi27 := 0; fi27 < 14; fi27++ {
        if len(v3) > 0 {
            v3[0] = 218
        }
        _ = fi27
    }
    switch (v11 | v11) {
    case -1:
        v1 = v1
    case -2147483648:
        var v28 int16 = -1
        _ = v28
    case 4294967295:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
