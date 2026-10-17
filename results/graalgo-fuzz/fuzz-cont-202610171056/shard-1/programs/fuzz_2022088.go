package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 string }
func genFn0(p0 uint16, p1 uint64, p2 uint8) string {
    v7 := GenS1{F0: 2}
    p0 += (^(p0))
    if (!true) {
        p1 = p1
        p1 -= 2
        p2 += p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (((("foo bar" + "") + "") + "hello") + "foo bar")
}

func main() {
    var v1 int32 = 383
    v2 := make(map[int]uint32)
    var v3 int8 = 100
    v4 := make(map[uint32]bool)
    var v5 uint = 652
    var v6 uint = v5
    v6 += v6
    v8 := GenS1{F0: (v6 ^ 2)}
    v5 -= v5
    v9 := GenS1{F0: (v5 - (v5 - 4294967295))}
    v10 := v2[-853]
    if true {
        v6 += 221
        var v11 int64 = -491
        var v12 float64 = 0.1
        _ = v11
        _ = v12
    } else {
        v3 = v3
        for fi13 := 0; fi13 < 9; fi13++ {
            if (42 <= (v10 & 96)) {
                var v14 uint64 = 557
                var v15 float64 = 1e10
                v10 -= uint32(v14)
                _ = v14
                _ = v15
            } else {
                if (!(v1 <= v1)) {
                    v16 := GenS1{F0: (v5 / ((v6) | 1))}
                    _ = v16
                }
            }
            v1 = (v1 & (v1 / ((v1) | 1)))
            _ = fi13
        }
    }
    v6 -= (v6 % ((v5) | 1))
    v1 = v1
    if (v10 <= (^(v10))) {
        v2[-161] = (v10 % ((v10) | 1))
        switch ("\n\t" + "") {
        case "\n\t":
            for fi17 := 0; fi17 < 14; fi17++ {
                if false {
                    v6 = v6
                    var v18 uint16 = uint16(v6)
                    if (v18 < (v18 & v18)) {
                        var v19 int8 = v3
                        for fi20 := 0; fi20 < 15; fi20++ {
                            var v21 int = (fi17 >> ((4294967295) & 63))
                            v5 += uint(v18)
                            _ = fi20
                            _ = v21
                        }
                        _ = v19
                    } else {
                        fmt.Println(v5, v6, v3)
                    }
                    _ = v18
                } else {
                    if false {
                        fmt.Println(v10, v10, v5)
                        switch "" {
                        case "a":
                            switch fi17 {
                            case 1:
                                v22 := []int{}
                                _ = v22
                            case -367:
                                fmt.Println(v3)
                            case 128:
                                if true {
                                    v23 := GenS1{F0: v5}
                                    _ = v23
                                }
                            }
                        case "foo bar":
                            v24 := v2[fi17]
                            _ = v24
                        case "":
                            if (65535 <= 574) {
                                v25 := []int{fi17, (fi17 | (^(fi17))), -2147483648}
                                _ = v25
                            } else {
                                v26 := []int{fi17}
                                _ = v26
                            }
                        default:
                            v3 += (v3 / ((v3) | 1))
                        }
                        var v27 int = fi17
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 1; fi28++ {
                            v6 = (v6 / ((v6) | 1))
                            _ = fi28
                        }
                        var v29 float64 = ((0.0 + (-2.25 / ((0.0 / float64(((0.1 - 0.0) + 0.1))) - ((-2.25 + 0.0) - (-2.25 - (3.14159 / 1.0)))))) * 1e10)
                        _ = v29
                    }
                    v30 := []int{fi17}
                    _ = v30
                }
                _ = fi17
            }
        case "hello":
            v31 := GenS1{F0: (v5 & (v6 << (((^(v6))) & 63)))}
            _ = v31
        case "foo bar":
            if true {
                for fi32 := 0; fi32 < 9; fi32++ {
                    var v33 float64 = 1.5
                    _ = fi32
                    _ = v33
                }
                for fi34 := 0; fi34 < 11; fi34++ {
                    v35 := []int{fi34, (fi34 << ((v5) & 63)), (-(fi34)), (fi34 ^ fi34)}
                    for fi36 := 0; fi36 < 4; fi36++ {
                        fmt.Println(v6, v3, v6)
                        _ = fi36
                    }
                    _ = fi34
                    _ = v35
                }
                v37 := []int{}
                _ = v37
            } else {
                switch (("a" + "hello") + ("a" + "a")) {
                case "foo bar":
                    var v38 int64 = -9223372036854775808
                    _ = v38
                case "hello":
                    var v39 uint16 = 600
                    _ = v39
                case "a":
                    v5 = (v6 & v5)
                }
                v40 := GenS1{F0: 42}
                _ = v40
            }
        }
    } else {
        v41 := GenS2{F0: uint16(v6), F1: "foo bar"}
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
