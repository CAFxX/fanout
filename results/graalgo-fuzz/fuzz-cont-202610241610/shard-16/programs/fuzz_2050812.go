package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0(p0 int8, p1 int16, p2 uint) int64 {
    p2 = (^(p2))
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1() uint8 {
    v6 := []int{-439}
    var v7 uint16 = 42
    if (true && true) {
        if false {
            switch ("foo bar" + ("hello" + "foo bar")) {
            case "\n\t":
                if true {
                    v8 := GenS1{F0: 18446744073709551615}
                    _ = len(v6)
                    _ = v8
                } else {
                    v7 -= (v7 % ((2) | 1))
                }
            case "a":
                var v9 uint64 = uint64(v7)
                _ = v9
            }
            for fi10 := 0; fi10 < 18; fi10++ {
                if len(v6) > 0 {
                    v6[0] = (fi10 & (fi10 % ((32767) | 1)))
                }
                v7 = (v7 << ((18446744073709551615) & 63))
                _ = fi10
            }
            if true {
                fmt.Println(v7)
                v7 += (v7 - v7)
            } else {
                if len(v6) > 0 {
                    v6[0] = -1
                }
                if (v7 <= v7) {
                    v7 = (v7 ^ (^(v7)))
                    if (-128 < -128) {
                        v7 = v7
                        _ = len(v6)
                        fmt.Println(v7)
                    } else {
                        var v11 uint16 = v7
                        _ = len(v6)
                        _ = v11
                    }
                    var v12 uint8 = 0
                    _ = v12
                } else {
                    fmt.Println(v7)
                    switch 493 {
                    case 0:
                        fmt.Println(v7)
                    case 32767:
                        v7 += (^(v7))
                    case -2147483648:
                        var v13 bool = (163 > 255)
                        _ = v13
                    default:
                        v14 := GenS1{F0: 2}
                        _ = v14
                    }
                    fmt.Println(v7)
                }
            }
        }
    } else {
        var v15 uint64 = 0
        for fi16 := 0; fi16 < 4; fi16++ {
            v17 := GenS2{F0: fi16, F1: (v15 | v15)}
            var v18 bool = true
            v7 = (v7 - v7)
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    }
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 float64 = ((1.5 + (1e10 - 0.0)) / -2.25)
    var v2 int64 = 42
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]string)
    var v5 float64 = (v1 - 1e10)
    var v19 uint32 = 65535
    fmt.Println(v2, v19, v1)
    var v20 bool = false
    var v21 bool = (!(2 <= 2))
    v20 = (v21 || v21)
    fmt.Println(v20, v2)
    fmt.Println(v5, v21, v21)
    v20 = v21
    v21 = (100 > 0)
    for fi22 := 0; fi22 < 2; fi22++ {
        v5 = -2.25
        _ = fi22
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 11; fi24++ {
            switch "" {
            case "hello":
                fmt.Println(fi24, v19)
            case "a":
                var v25 uint = 683
                v20 = ("\n\t" < "hello")
                _ = v25
            case "foo bar":
                var v26 string = ""
                _ = v26
            }
            var v27 uint8 = 232
            _ = fi24
            _ = v27
        }
        v20 = (fi23 != 0)
        v20 = (255 >= 2)
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
