package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int16; F2 int }
type GenS3 struct { F0 int8 }
func genFn0() int8 {
    v6 := GenS1{F0: 65535}
    v7 := []int{-965, -9223372036854775808, -2147483648, 127}
    var v8 uint8 = 127
    _ = len(v7)
    _ = v6
    _ = v7
    _ = v8
    return -128
}
func genFn1(p0 int8, p1 int16, p2 int) int16 {
    p1 += -805
    v9 := GenS1{F0: 991}
    p2 -= (p2 & (p2 * (p2 * p2)))
    switch p2 {
    case -160:
        if true {
            p1 = p1
        } else {
            var v10 uint8 = 127
            _ = v10
        }
    case -540:
        var v11 int64 = 0
        _ = v11
    default:
        if (1000 != p1) {
            if false {
                var v12 float32 = (1.5 / 1.5)
                _ = v12
            } else {
                switch ((("hello" + "\n\t") + "") + "\n\t") {
                case "":
                    if (2 < 42) {
                        v13 := GenS1{F0: 33}
                        v14 := GenS1{F0: 2}
                        _ = v13
                        _ = v14
                    } else {
                        var v15 uint16 = 619
                        _ = v15
                    }
                default:
                    var v16 uint64 = 2
                    _ = v16
                }
                var v17 int8 = p0
                if true {
                    for fi18 := 0; fi18 < 13; fi18++ {
                        for fi19 := 0; fi19 < 1; fi19++ {
                            p0 -= (p0 >> ((65535) & 63))
                            p1 = p1
                            _ = fi19
                        }
                        v17 += p0
                        for fi20 := 0; fi20 < 5; fi20++ {
                            v21 := []int{0, (fi18 + p2)}
                            v22 := GenS2{F0: 855, F1: p1, F2: (^(p2))}
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        _ = fi18
                    }
                } else {
                    switch (p2 << ((0) & 63)) {
                    case 128:
                        var v23 int = (p2 ^ p2)
                        _ = v23
                    case -9223372036854775808:
                        p1 = (p1 ^ p1)
                    default:
                        var v24 uint64 = 42
                        _ = v24
                    }
                }
                _ = v17
            }
        } else {
            var v25 string = "a"
            _ = v25
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p1
}

func main() {
    var v1 uint8 = 182
    var v2 int8 = -1
    var v3 int = -11
    var v4 uint = 1
    var v5 float64 = 0.0
    var v26 bool = (((v4 * (v4 - (v4 | v4))) > 2) && false)
    if (false || v26) {
        v27 := GenS2{F0: 0, F1: -1, F2: v3}
        v28 := []int{v3}
        _ = v27
        _ = v28
    } else {
        v26 = true
    }
    v26 = (1 <= v2)
    var v29 uint64 = 1
    var v30 uint = 919
    fmt.Println(v2, v26)
    v31 := GenS1{F0: v30}
    if v26 {
        v29 = (^(v29))
        for fi32 := 0; fi32 < 9; fi32++ {
            var v33 int16 = 503
            _ = fi32
            _ = v33
        }
    }
    v34 := GenS3{F0: genFn0()}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
