package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int8 }
func genFn0(p0 int) bool {
    v7 := []int{(p0 ^ (p0 ^ (p0 | (-(p0))))), p0, p0}
    if false {
        if true {
            if (100 <= 0) {
                var v8 float64 = 0.1
                var v9 uint16 = 32767
                _ = v8
                _ = v9
            }
            if len(v7) > 0 {
                v7[0] = (-(p0))
            }
            p0 -= p0
        } else {
            v7 = append(v7, (p0 >> ((905) & 63)))
        }
    } else {
        for fi10 := 0; fi10 < 12; fi10++ {
            _ = len(v7)
            _ = fi10
        }
        p0 -= p0
    }
    var v11 int16 = -32768
    _ = p0
    _ = v11
    _ = v7
    return (false && true)
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]int)
    var v3 bool = ((1.5 + (((v1 - v1) * float32(1.5)) + float32(0.1))) < 1.5)
    var v4 int64 = 655
    var v5 float32 = v1
    var v6 uint16 = 0
    v6 = v6
    fmt.Println(v5, v4, v3)
    if ((1 >= 0) && ((!v3) || true)) {
        var v12 uint32 = 929
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint16 = v6
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 uint8 = 102
                if false {
                    var v18 float32 = (v5 * ((((1.5 * 3.14159) * v1) + float32(0.0)) * v5))
                    var v19 uint = 980
                    _ = v18
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
            fmt.Println(v1, v5, v12)
            _ = fi15
        }
        _ = v12
    } else {
        switch 2147483647 {
        case 42:
            v4 = 42
        case 0:
            v20 := GenS1{F0: 9223372036854775807}
            _ = v20
        default:
            if (-430 == v4) {
                v21 := GenS1{F0: (^(v4))}
                fmt.Println(v4, v3)
                if ((v4 | v4) != -9223372036854775808) {
                    v22 := GenS2{F0: 358}
                    v3 = (v3 && ("hello" == ((("hello" + ("hello" + "foo bar")) + "hello") + "")))
                    _ = v22
                }
                _ = v21
            } else {
                fmt.Println(v6)
                v23 := GenS1{F0: (v4 & v4)}
                v24 := v2[2147483648]
                _ = v23
                _ = v24
            }
        }
        var v25 int64 = (v4 | (v4 + (v4 * 0)))
        v26 := v2[1]
        _ = v25
        _ = v26
    }
    v27 := []int{-40, 32767, 390, 0}
    var v28 uint64 = 661
    for fi29 := 0; fi29 < 18; fi29++ {
        v30 := GenS2{F0: 635}
        var v31 float32 = 1e10
        for fi32 := 0; fi32 < 13; fi32++ {
            v4 = v4
            _ = fi32
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    delete(v2, 1)
    v28 += v28
    v6 += 2
    fmt.Println(v1, v3)
    switch -2147483648 {
    case 620:
        for fi33 := 0; fi33 < 19; fi33++ {
            if ((0 == 137) && true) {
                for fi34 := 0; fi34 < 18; fi34++ {
                    v27 = append(v27, fi34)
                    _ = fi34
                }
                _ = len(v27)
                if ("hello" <= "") {
                    if (v4 == int64(fi33)) {
                        var v35 uint16 = (^(v6))
                        _ = v35
                    } else {
                        var v36 int64 = int64(v28)
                        _ = v36
                    }
                    v5 -= (v1 - v1)
                    var v37 float32 = (0.0 - float32(((3.14159 / 0.1) + 3.14159)))
                    _ = v37
                } else {
                    v38 := GenS3{F0: 127, F1: 127}
                    _ = v38
                }
            }
            _ = fi33
        }
    default:
        if true {
            fmt.Println(v3)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
