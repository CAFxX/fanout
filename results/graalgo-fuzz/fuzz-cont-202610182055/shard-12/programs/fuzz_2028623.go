package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int; F1 int; F2 uint16 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint, p1 bool, p2 uint) bool {
    v7 := GenS1{F0: -1}
    var v8 int16 = 0
    v9 := []int{}
    v10 := GenS1{F0: (-(v8))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return false
}

func main() {
    var v1 uint16 = 0
    v2 := []uint32{}
    v3 := []string{"\n\t", "a", (("" + ("\n\t" + "foo bar")) + ("\n\t" + "foo bar")), "foo bar"}
    v4 := make(map[int]string)
    var v5 int16 = 542
    var v6 uint16 = v1
    fmt.Println(v1, v6, v5)
    var v11 string = "foo bar"
    v12 := GenS1{F0: 0}
    v5 -= v5
    if false {
        for fi13 := 0; fi13 < 19; fi13++ {
            if false {
                if (0 == 42) {
                    var v14 int = fi13
                    v1 -= 2
                    _ = v14
                }
                for fi15 := 0; fi15 < 17; fi15++ {
                    v11 = "foo bar"
                    if (!true) {
                        v2 = append(v2, 706)
                        if len(v2) > 0 {
                            v2[0] = 155
                        }
                        fmt.Println(fi15)
                    }
                    v16 := GenS2{F0: "a", F1: (v6 & v1)}
                    _ = fi15
                    _ = v16
                }
            }
            var v17 float32 = float32(0.0)
            _ = fi13
            _ = v17
        }
        switch -784 {
        case 9223372036854775807:
            v6 -= (v6 << ((2) & 63))
        case -9223372036854775808:
            fmt.Println(v5)
        case 1:
            v18 := GenS1{F0: v5}
            _ = v18
        }
    } else {
        fmt.Println(v1, v6, v6)
        var v19 int16 = (v5 * 1)
        var v20 int64 = 1
        _ = v19
        _ = v20
    }
    fmt.Println(v5, v6)
    if (1 <= -612) {
        switch 32768 {
        case 2147483647:
            var v21 string = ((("foo bar" + ((v11 + "foo bar") + v11)) + v11) + ((v11 + v11) + (v11 + "")))
            _ = v21
        default:
            v22 := GenS1{F0: (v5 | v5)}
            _ = v22
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            v5 = v5
            v1 -= (v1 ^ 42)
            _ = fi23
        }
    } else {
        for fi24 := 0; fi24 < 3; fi24++ {
            var v25 int64 = -1
            v11 = v11
            _ = fi24
            _ = v25
        }
    }
    v26 := GenS3{F0: 42, F1: -9223372036854775808, F2: v1}
    var v27 int64 = 42
    var v28 uint32 = 4294967295
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v6)
    if false {
        var v29 string = "hello"
        if false {
            v1 = 255
            v30 := GenS1{F0: v5}
            _ = v30
        } else {
            if true {
                v29 = "hello"
                for fi31 := 0; fi31 < 10; fi31++ {
                    fmt.Println(v29)
                    _ = fi31
                }
            } else {
                if (((^(v28)) == (v28 | v28)) && true) {
                    var v32 int8 = -1
                    _ = v32
                } else {
                    var v33 string = v29
                    if true {
                        v28 = v28
                        v34 := GenS4{F0: 136}
                        var v35 int16 = (v5 - v5)
                        _ = v34
                        _ = v35
                    }
                    _ = v33
                }
                fmt.Println(v29)
            }
            _ = len(v3)
            v27 += v27
        }
        _ = v29
    }
    v11 = ((v11 + ("a" + (v11 + ((v11 + v11) + v11)))) + (v11 + v11))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
