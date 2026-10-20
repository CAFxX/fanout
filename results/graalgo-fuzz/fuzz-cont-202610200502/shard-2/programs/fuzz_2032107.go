package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint }
func genFn0(p0 uint8) string {
    var v7 int64 = 42
    v7 += v7
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: "\n\t", F1: (p0 ^ 1)}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return "foo bar"
}

func main() {
    var v1 uint64 = 0
    var v2 int8 = 103
    v3 := []int{}
    v4 := make(map[string]bool)
    var v5 uint32 = 2147483647
    var v6 int16 = -1
    v2 -= (v2 | 127)
    if (int8(v6) > (-(v2))) {
        v10 := v4["a"]
        if true {
            var v11 string = "a"
            _ = v11
        } else {
            var v12 int8 = v2
            v5 += (^(v5))
            v10 = ((uint32(v1) <= v5) && (1 < -266))
            _ = v12
        }
        for fi13 := 0; fi13 < 13; fi13++ {
            v14 := GenS1{F0: "hello", F1: 1}
            var v15 uint64 = (v1 ^ v1)
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v10
    }
    v1 = v1
    var v16 uint16 = 0
    var v17 uint8 = 42
    var v18 uint64 = (v1 | (v1 & (v1 + (v1 + v1))))
    var v19 uint32 = (v5 << ((18446744073709551615) & 63))
    v1 = (^(v1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        for fi21 := 0; fi21 < 11; fi21++ {
            fmt.Println(fi21)
            if ((^(v19)) < v19) {
                delete(v4, "foo bar")
                var v22 int = 127
                switch ((("foo bar" + ("foo bar" + ("foo bar" + ""))) + ("hello" + "foo bar")) + "foo bar") {
                case "":
                    _ = len(v4)
                default:
                    if (false && false) {
                        fmt.Println(v19, v1)
                        v1 = v18
                    } else {
                        var v23 bool = false
                        _ = v23
                    }
                }
                v1 = v18
                switch fi20 {
                case 32767:
                    var v24 int16 = (v6 * 1000)
                    _ = v24
                case -940:
                    v1 = (v18 << ((4294967295) & 63))
                case -9223372036854775808:
                    v5 = v19
                default:
                    var v25 bool = true
                    _ = v25
                }
                if false {
                    var v26 uint8 = (^(v17))
                    v17 += (v17 - v26)
                    v27 := GenS1{F0: "", F1: v26}
                    _ = v26
                    _ = v27
                }
                _ = v22
            } else {
                _ = len(v3)
            }
            _ = fi21
        }
        _ = fi20
    }
    switch int(v16) {
    case -512:
        for fi28 := 0; fi28 < 19; fi28++ {
            for fi29 := 0; fi29 < 19; fi29++ {
                v2 = 47
                if true {
                    if (-1 >= (fi29 * fi29)) {
                        if (true || (42 <= -1)) {
                            if len(v3) > 0 {
                                v3[0] = -9223372036854775808
                            }
                        }
                        v30 := GenS2{F0: "\n\t"}
                        _ = v30
                    } else {
                        v31 := GenS3{F0: ("" + ""), F1: 65535}
                        delete(v4, "foo bar")
                        v2 = v2
                        _ = v31
                    }
                    v32 := GenS2{F0: "\n\t"}
                    _ = v32
                }
                _ = fi29
            }
            _ = fi28
        }
    case 4294967295:
        v33 := GenS3{F0: "", F1: 255}
        _ = v33
    }
    switch ("" + "\n\t") {
    case "\n\t":
        _ = len(v4)
    case "foo bar":
        if ((v16 | 385) < 65535) {
            fmt.Println(v16, v6, v18)
            var v34 uint = uint(v16)
            _ = v34
        }
    }
    v2 += (v2 % (((v2 << ((798) & 63))) | 1))
    var v35 uint = 866
    var v36 float64 = (((3.14159 * 1e10) * -2.25) / 0.1)
    v17 += (v17 >> ((114) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
