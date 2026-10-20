package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64; F1 int }
func genFn0(p0 uint64) int16 {
    v7 := []int{640}
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 bool = true
        _ = len(v7)
        v9 = (!(!(!(-9223372036854775808 != -1))))
        _ = fi8
        _ = v9
    }
    _ = len(v7)
    var v10 float32 = 1.5
    _ = p0
    _ = v10
    _ = v7
    return 1000
}
func genFn1(p0 int64, p1 int8) uint32 {
    v11 := GenS1{F0: 18446744073709551615}
    var v12 bool = (-1 > 453)
    var v13 uint64 = 18446744073709551615
    var v14 string = "\n\t"
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS1{F0: 977}
        if v12 {
            v13 = (^(v13))
            var v17 uint32 = 0
            v18 := GenS2{F0: 42}
            _ = v17
            _ = v18
        } else {
            v19 := []int{}
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 0
}

func main() {
    var v1 int8 = -19
    v2 := make(map[uint32]bool)
    var v3 uint8 = 42
    var v4 int = 42
    var v5 uint8 = 127
    v6 := make(map[uint32]bool)
    v2[2] = true
    v4 = (v4 | (v4 | (v4 - (^(v4)))))
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 uint8 = (v3 >> ((255) & 63))
            var v23 int64 = 42
            v24 := []int{(fi21 % ((-202) | 1))}
            v3 = (v22 - v3)
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        var v25 uint32 = 1
        if (v25 < v25) {
            switch "foo bar" {
            case "":
                var v26 uint32 = (v25 ^ 1)
                _ = v26
            }
            v27 := GenS2{F0: 0}
            fmt.Println(fi20, v25)
            _ = v27
        } else {
            switch (-(v4)) {
            case 1:
                v28 := GenS3{F0: 0}
                _ = v28
            case 900:
                v4 = (fi20 % ((fi20) | 1))
            case -967:
                v29 := GenS4{F0: 1}
                _ = v29
            }
            for fi30 := 0; fi30 < 12; fi30++ {
                v31 := v2[genFn1(9223372036854775807, v1)]
                _ = fi30
                _ = v31
            }
        }
        _ = fi20
        _ = v25
    }
    var v32 uint32 = 340
    var v33 string = "a"
    for fi34 := 0; fi34 < 4; fi34++ {
        v35 := GenS5{F0: (^(v5)), F1: (^(v32))}
        v4 -= 42
        _ = fi34
        _ = v35
    }
    if true {
        v3 += 127
        if (!(802 <= 111)) {
            if (-1 <= 42) {
                switch "" {
                case "\n\t":
                    if (42 != 18446744073709551615) {
                        var v36 uint64 = 123
                        _ = v36
                    }
                case "foo bar":
                    v37 := v2[(v32 % (((v32 ^ genFn1(42, (v1 / (((v1 / (((v1 * (v1 >> ((1) & 63)))) | 1))) | 1))))) | 1))]
                    _ = v37
                }
                fmt.Println(v1, v1)
            } else {
                v1 -= (v1 / ((v1) | 1))
                if (v33 >= "foo bar") {
                    var v38 int8 = v1
                    for fi39 := 0; fi39 < 11; fi39++ {
                        var v40 string = (("\n\t" + v33) + v33)
                        _ = fi39
                        _ = v40
                    }
                    var v41 int64 = -429
                    _ = v38
                    _ = v41
                }
            }
            v42 := GenS6{F0: 386}
            fmt.Println(v1)
            _ = v42
        } else {
            var v43 int8 = -57
            if true {
                switch (((("hello" + v33) + "hello") + v33) + "foo bar") {
                case "a":
                    if false {
                        var v44 string = "a"
                        _ = v44
                    } else {
                        var v45 int32 = -976
                        _ = v45
                    }
                case "\n\t":
                    if ((!true) && ((true || ((v33 + ((v33 + "a") + (("foo bar" + ("hello" + "a")) + (v33 + (("hello" + v33) + v33))))) != ("a" + (((v33 + v33) + (v33 + v33)) + v33)))) && (!(18446744073709551615 < 255)))) {
                        v43 += (v1 & (v1 / ((0) | 1)))
                    } else {
                        v46 := []int{-1}
                        _ = v46
                    }
                default:
                    for fi47 := 0; fi47 < 2; fi47++ {
                        var v48 int32 = 1
                        _ = fi47
                        _ = v48
                    }
                }
                v49 := []int{(v4 ^ v4), 2147483647}
                _ = v49
            }
            var v50 string = (v33 + v33)
            _ = v43
            _ = v50
        }
        var v51 uint16 = 1
        _ = v51
    } else {
        for fi52 := 0; fi52 < 16; fi52++ {
            v1 += v1
            _ = fi52
        }
        if true {
            var v53 float32 = 0.1
            _ = v53
        } else {
            v2[(v32 * 4294967295)] = false
        }
    }
    for fi54 := 0; fi54 < 5; fi54++ {
        for fi55 := 0; fi55 < 18; fi55++ {
            v56 := GenS7{F0: 1, F1: (^(v4))}
            _ = fi55
            _ = v56
        }
        v57 := GenS6{F0: 42}
        _ = fi54
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
