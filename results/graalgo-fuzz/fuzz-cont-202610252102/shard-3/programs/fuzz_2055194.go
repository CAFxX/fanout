package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint32 }
type GenS2 struct { F0 int8; F1 int32; F2 uint32 }
func genFn0(p0 uint16, p1 int) uint8 {
    var v4 bool = false
    v5 := []int{p1, -949}
    fmt.Println(p1, v4)
    var v6 float32 = float32(1.5)
    p0 -= 2
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 127
}
func genFn1(p0 int16, p1 uint16, p2 int64) uint16 {
    if ((!true) || false) {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 float32 = 0.1
            var v9 uint16 = (p1 - (p1 + (^(p1))))
            var v10 string = "hello"
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        if false {
            p0 -= 0
            var v11 uint = 18446744073709551615
            _ = v11
        } else {
            var v12 int8 = 0
            fmt.Println(p1, p0)
            v12 += (v12 & v12)
            _ = v12
        }
        p2 -= p2
    }
    if (false || false) {
        var v13 uint = 2
        _ = v13
    } else {
        p1 = 1
    }
    if (-995 == -1) {
        p0 = 1
        p2 = 316
    }
    var v14 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 0
}

func main() {
    v1 := []string{"foo bar"}
    var v2 uint8 = 42
    var v3 uint16 = 32767
    for fi15 := 0; fi15 < 11; fi15++ {
        if false {
            v2 -= (v2 << ((1) & 63))
            v2 -= (v2 % (((v2 << ((65535) & 63))) | 1))
            var v16 int64 = -450
            _ = v16
        } else {
            fmt.Println(v3, v2)
            v17 := GenS1{F0: 484, F1: -9223372036854775808, F2: 1}
            fmt.Println(v2, v2)
            _ = v17
        }
        fmt.Println(v3)
        var v18 uint16 = v3
        _ = fi15
        _ = v18
    }
    v3 -= (v3 % ((65535) | 1))
    if false {
        v3 = v3
        var v19 int = 4294967295
        _ = v19
    } else {
        var v20 float32 = float32(1.5)
        if true {
            for fi21 := 0; fi21 < 2; fi21++ {
                switch "a" {
                case "":
                    var v22 uint32 = 782
                    _ = v22
                case "a":
                    v20 = (float32(1.5) * (((v20 * 0.0) * v20) + ((((float32(-2.25) - v20) / v20) / v20) + 0.0)))
                case "hello":
                    var v23 int64 = 42
                    _ = v23
                default:
                    v2 = (v2 & (v2 * genFn0((v3 >> ((715) & 63)), (fi21 | 9223372036854775807))))
                }
                for fi24 := 0; fi24 < 19; fi24++ {
                    var v25 bool = (!(true && ((v3 * (v3 - (v3 << ((851) & 63)))) != 0)))
                    v25 = (!(false && false))
                    v26 := GenS1{F0: 2, F1: -1, F2: 2}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                switch "\n\t" {
                case "foo bar":
                    v20 = 1.5
                case "a":
                    fmt.Println(v20, v3)
                case "hello":
                    fmt.Println(v20)
                default:
                    v20 += (v20 - 3.14159)
                }
                _ = fi21
            }
            if false {
                if false {
                    fmt.Println(v20)
                    for fi27 := 0; fi27 < 11; fi27++ {
                        if true {
                            v20 = v20
                        }
                        v3 -= v3
                        _ = fi27
                    }
                }
            }
            var v28 uint = 65535
            _ = v28
        }
        _ = v20
    }
    v2 += (^(v2))
    if false {
        v29 := GenS1{F0: 578, F1: -1, F2: 810}
        switch (("foo bar" + "foo bar") + "") {
        case "a":
            for fi30 := 0; fi30 < 1; fi30++ {
                v2 += v2
                v31 := GenS2{F0: 1, F1: 1, F2: 4294967295}
                var v32 uint32 = 1
                _ = fi30
                _ = v31
                _ = v32
            }
        default:
            if true {
                var v33 uint8 = (v2 & 127)
                _ = v33
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        _ = v29
    } else {
        var v34 uint8 = genFn0(42, 32767)
        _ = v34
    }
    v2 = v2
    fmt.Println(v3)
    fmt.Println(v3, v2)
    var v35 int16 = -32768
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v36 uint64 = 823
    var v37 int16 = (^(v35))
    v38 := GenS2{F0: 1, F1: -580, F2: 119}
    var v39 float32 = (float32(0.1) / 1e10)
    v35 += 53
    var v40 uint8 = v2
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
