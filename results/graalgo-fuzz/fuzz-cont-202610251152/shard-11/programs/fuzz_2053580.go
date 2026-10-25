package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
func genFn0(p0 uint8, p1 int16, p2 uint16) uint16 {
    for fi6 := 0; fi6 < 10; fi6++ {
        if true {
            v7 := GenS1{F0: 68, F1: -67, F2: (p0 ^ p0)}
            _ = v7
        } else {
            if true {
                var v8 int32 = -2147483648
                v8 += v8
                v9 := []int{(fi6 + -9223372036854775808), 910}
                _ = v8
                _ = v9
            } else {
                var v10 uint16 = (p2 >> ((255) & 63))
                v10 = 42
                v11 := []int{(fi6 >> ((65535) & 63))}
                _ = v10
                _ = v11
            }
        }
        var v12 float64 = (1e10 - 1e10)
        _ = fi6
        _ = v12
    }
    if true {
        var v13 int = 128
        fmt.Println(p0, v13, v13)
        p1 -= (p1 % (((^(p1))) | 1))
        _ = v13
    } else {
        p1 += (^(p1))
        var v14 uint16 = 0
        _ = v14
    }
    v15 := []int{-529, 109, -150}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{42, -2147483648, 42}
    v3 := []int{127, 4294967295, 9223372036854775807}
    v4 := make(map[string]string)
    v5 := []int{127, 32767, -285}
    var v16 uint64 = 0
    var v17 bool = false
    fmt.Println(v17, v16)
    if v17 {
        for fi18 := 0; fi18 < 8; fi18++ {
            v16 = v16
            _ = fi18
        }
    }
    var v19 uint32 = 4294967295
    if ((201 != 128) && (0 > -269)) {
        var v20 uint16 = 65535
        v21 := GenS1{F0: 2147483648, F1: 42, F2: 0}
        _ = v20
        _ = v21
    } else {
        if false {
            var v22 bool = v17
            v23 := GenS2{F0: "a"}
            _ = v22
            _ = v23
        } else {
            v24 := GenS3{F0: 228}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v24
        }
        for fi25 := 0; fi25 < 2; fi25++ {
            v19 += v19
            v26 := GenS2{F0: "foo bar"}
            _ = fi25
            _ = v26
        }
    }
    var v27 string = "foo bar"
    var v28 uint = 0
    v19 = 1
    fmt.Println(v17)
    var v29 uint = (^(v28))
    switch v27 {
    case "a":
        fmt.Println(v17, v17, v29)
    case "":
        if (-32768 >= -826) {
            if true {
                v28 = v28
                if v17 {
                    v17 = (4294967295 > 127)
                    v16 = 353
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    v19 = 162
                }
                fmt.Println(v16)
            } else {
                if (32767 == -731) {
                    for fi30 := 0; fi30 < 4; fi30++ {
                        var v31 string = ((((((((("foo bar" + "foo bar") + ("\n\t" + v27)) + "foo bar") + (v27 + v27)) + (v27 + ("" + ""))) + (v27 + (((v27 + "a") + v27) + v27))) + ("hello" + (v27 + v27))) + v27) + ("" + v27))
                        var v32 uint64 = (v16 | v16)
                        v17 = v17
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                } else {
                    v17 = (v17 && (-43 >= 1))
                    v33 := GenS1{F0: v19, F1: 100, F2: 0}
                    _ = v33
                }
                v16 -= v16
                _ = len(v1)
            }
            v27 = v27
            for fi34 := 0; fi34 < 20; fi34++ {
                var v35 int = fi34
                var v36 uint32 = 2147483647
                _ = fi34
                _ = v35
                _ = v36
            }
        }
    default:
        v29 += (v28 * (v29 ^ 1))
    }
    var v37 float32 = float32(0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
