package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 string; F2 uint16 }
type GenS5 struct { F0 int16; F1 int; F2 int64 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint8) uint64 {
    v7 := GenS1{F0: 749}
    v8 := []int{1, 172, 678}
    if ((true && true) && true) {
        v8 = append(v8, 9223372036854775807)
        v9 := GenS1{F0: 559}
        v10 := GenS2{F0: 1, F1: 4294967295}
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 19; fi11++ {
            if len(v8) > 0 {
                v8[0] = 2147483647
            }
            if (p1 != p1) {
                for fi12 := 0; fi12 < 6; fi12++ {
                    if (846 > 65535) {
                        p0 -= p0
                        p1 += p1
                        p1 = p1
                    }
                    _ = fi12
                }
            } else {
                p1 = p1
                var v13 string = "\n\t"
                _ = v13
            }
            p0 = 785
            _ = fi11
        }
        _ = len(v8)
    }
    var v14 bool = (2 > 4294967295)
    _ = p0
    _ = p1
    _ = v14
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 bool, p1 int8, p2 int) int64 {
    switch (p2 + (p2 % ((p2) | 1))) {
    case -508:
        if (74 == genFn0(4294967295, 127)) {
            p1 = p1
            p0 = true
        } else {
            v15 := GenS3{F0: 2}
            _ = v15
        }
    case -66:
        fmt.Println(p1, p2, p2)
    case 128:
        fmt.Println(p2, p2, p1)
    default:
        var v16 uint16 = 747
        _ = v16
    }
    fmt.Println(p0)
    v17 := GenS2{F0: (p2 | (^(p2))), F1: 580}
    v18 := []int{p2, (p2 ^ (-(p2))), (p2 + p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 0
}

func main() {
    var v1 bool = (1 > 4294967295)
    v2 := make(map[int]string)
    var v3 int8 = 0
    var v4 bool = v1
    v5 := make(map[string]int)
    var v6 uint16 = 32767
    v19 := []int{-466, 32767, -154}
    if len(v19) > 0 {
        v19[0] = 9223372036854775807
    }
    if v4 {
        v6 -= 32767
        for fi20 := 0; fi20 < 12; fi20++ {
            if ((v6 & v6) >= (v6 + (v6 - (v6 / (((^(v6))) | 1))))) {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
            } else {
                if v4 {
                    fmt.Println(v3, v1, fi20)
                } else {
                    v21 := GenS2{F0: (fi20 >> ((178) & 63)), F1: 255}
                    if (-1 == 0) {
                        var v22 int64 = 405
                        v19 = append(v19, (-(fi20)))
                        _ = v22
                    } else {
                        var v23 int32 = 2147483647
                        _ = v23
                    }
                    fmt.Println(v1, v4)
                    _ = v21
                }
            }
            switch (fi20 & (fi20 ^ -342)) {
            case 128:
                v3 = (^(v3))
            case -208:
                switch "\n\t" {
                case "":
                    v24 := GenS4{F0: 9223372036854775807, F1: ((("" + ("a" + "foo bar")) + ("hello" + ("" + "hello"))) + ("\n\t" + "hello")), F2: v6}
                    _ = v24
                case "hello":
                    if len(v19) > 0 {
                        v19[0] = 84
                    }
                case "\n\t":
                    var v25 uint8 = 126
                    _ = v25
                default:
                    var v26 uint16 = 2
                    _ = v26
                }
            }
            v4 = v4
            _ = fi20
        }
        fmt.Println(v4)
    } else {
        var v27 uint = 0
        v3 = -109
        _ = v27
    }
    fmt.Println(v4, v1)
    switch ("hello" + "") {
    case "":
        v1 = v4
    case "a":
        v6 = (v6 - (v6 ^ (v6 - v6)))
    case "hello":
        v28 := v2[2147483647]
        _ = v28
    default:
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 string = "hello"
            _ = fi29
            _ = v30
        }
    }
    v31 := GenS2{F0: -1, F1: 2}
    var v32 int32 = -2147483648
    v33 := GenS5{F0: 42, F1: 0, F2: 42}
    if v4 {
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 int32 = v32
            v35 = (v32 * 42)
            v35 += (v35 | (v35 - (v32 % (((v32 >> ((18446744073709551615) & 63))) | 1))))
            _ = fi34
            _ = v35
        }
        var v36 bool = (42 <= 1)
        if len(v19) > 0 {
            v19[0] = 42
        }
        _ = v36
    } else {
        fmt.Println(v6, v1)
        var v37 float32 = 1e10
        _ = v37
    }
    switch 2147483647 {
    case 32768:
        v3 += 64
    }
    v6 += (v6 + (^(v6)))
    v38 := GenS6{F0: -75}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
