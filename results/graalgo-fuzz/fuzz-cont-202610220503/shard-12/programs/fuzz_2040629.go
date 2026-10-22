package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 string; F1 uint64 }
func genFn0(p0 uint, p1 bool, p2 uint32) uint32 {
    v6 := GenS1{F0: "hello", F1: 42}
    p1 = false
    p2 += 2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483647
}
func genFn1(p0 uint8) int8 {
    p0 += (p0 & 128)
    switch -22 {
    case -828:
        p0 = p0
    default:
        p0 = p0
    }
    v7 := []int{32768, -824, -438, 9223372036854775807}
    if false {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    } else {
        var v8 int16 = 0
        _ = v8
    }
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[uint32]bool)
    var v3 uint8 = 247
    var v4 int8 = 100
    v5 := make(map[uint32]bool)
    v3 += (v3 % (((^(v3))) | 1))
    delete(v5, 65535)
    if (42 != uint32(v3)) {
        if (!true) {
            fmt.Println(v1)
            if (true && false) {
                v2[0] = false
            } else {
                v3 = (v3 - v3)
            }
            v9 := GenS1{F0: "a", F1: -2147483648}
            _ = v9
        } else {
            var v10 float64 = 1e10
            fmt.Println(v4, v10)
            _ = v10
        }
        v1 = 1
    } else {
        if true {
            fmt.Println(v1)
        } else {
            for fi11 := 0; fi11 < 9; fi11++ {
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 float64 = -2.25
                    _ = fi12
                    _ = v13
                }
                var v14 int = -1
                v1 += (^(v1))
                _ = fi11
                _ = v14
            }
            for fi15 := 0; fi15 < 11; fi15++ {
                switch "" {
                case "\n\t":
                    if true {
                        for fi16 := 0; fi16 < 16; fi16++ {
                            var v17 uint8 = v3
                            var v18 bool = true
                            fmt.Println(v1, v4, v4)
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        v1 = (v1 + (v1 | (v1 + (v1 * 65535))))
                        v19 := GenS2{F0: 940}
                        _ = v19
                    } else {
                        fmt.Println(v4, v1, v1)
                    }
                default:
                    var v20 bool = ((v1 / ((2) | 1)) == 42)
                    _ = v20
                }
                v21 := GenS2{F0: 1}
                v22 := []int{}
                _ = fi15
                _ = v21
                _ = v22
            }
            if ((610 < int(v4)) && false) {
                v23 := GenS2{F0: -32768}
                fmt.Println(v1, v1)
                _ = v23
            }
        }
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 string = (((((("foo bar" + "a") + "hello") + "\n\t") + "a") + "foo bar") + (("" + (("a" + (((("foo bar" + "a") + "") + "hello") + ((("hello" + (("a" + "") + "a")) + ("hello" + "hello")) + "a"))) + "")) + "hello"))
        _ = fi24
        _ = v25
    }
    if (18446744073709551615 <= 2) {
        if false {
            var v26 int8 = (v4 ^ (v4 & (^(v4))))
            var v27 int8 = (-(v4))
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 11; fi28++ {
                fmt.Println(v4)
                v29 := v2[478]
                _ = fi28
                _ = v29
            }
            var v30 int32 = 2147483647
            v1 = (v1 & 83)
            _ = v30
        }
        var v31 uint16 = (v1 % (((^(v1))) | 1))
        var v32 uint8 = (^(v3))
        _ = v31
        _ = v32
    } else {
        v33 := GenS3{F0: "foo bar", F1: (!(!true))}
        if (false || (((!false) && false) && (!(42 < 4294967295)))) {
            if ((1e10 + float64(1e10)) <= 0.0) {
                switch -9223372036854775808 {
                case 32768:
                    var v34 string = ("foo bar" + "\n\t")
                    _ = v34
                default:
                    if (((v4 < v4) || ((1000 == 327) && true)) && true) {
                        v1 = (v1 << ((4294967295) & 63))
                    } else {
                        if (!true) {
                            var v35 string = "foo bar"
                            if false {
                                v5[388] = ((127 > 1) && ("hello" == (((v35 + (v35 + "")) + v35) + v35)))
                            }
                            _ = len(v2)
                            _ = v35
                        } else {
                            v1 = 233
                        }
                    }
                }
                v3 -= (v3 | (^(v3)))
                v36 := GenS4{F0: ("" + "a"), F1: 1}
                _ = v36
            }
            for fi37 := 0; fi37 < 8; fi37++ {
                var v38 int16 = 1
                _ = fi37
                _ = v38
            }
        }
        _ = v33
    }
    v39 := GenS1{F0: "hello", F1: -2147483648}
    v4 = (v4 | genFn1(v3))
    v40 := []int{}
    var v41 uint = 1
    fmt.Println(v3)
    v42 := GenS4{F0: (("\n\t" + "a") + "a"), F1: 701}
    var v43 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
