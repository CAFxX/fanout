package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 string }
func genFn0(p0 int64, p1 uint64, p2 int32) int64 {
    v7 := GenS1{F0: 4294967295, F1: p2, F2: "a"}
    v8 := GenS1{F0: 127, F1: 2147483647, F2: ("a" + ("a" + "\n\t"))}
    v9 := GenS2{F0: "foo bar"}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = 0.0
    var v3 string = ("hello" + "hello")
    var v4 uint = 570
    var v5 int = 4294967295
    v6 := make(map[int]bool)
    var v10 uint64 = 900
    var v11 int8 = 31
    v12 := []int{(v5 - 32768)}
    for fi13 := 0; fi13 < 13; fi13++ {
        v3 = (v3 + "foo bar")
        fmt.Println(v2)
        _ = fi13
    }
    v14 := GenS3{F0: -2147483648}
    switch v5 {
    case 4294967295:
        v4 = (v4 % (((v4 - v4)) | 1))
    case 32767:
        switch 517 {
        case 32767:
            if ((!(false && (!true))) || false) {
                switch v5 {
                case 127:
                    v15 := GenS4{F0: true, F1: 9223372036854775807}
                    _ = v15
                default:
                    if ("foo bar" == "hello") {
                        v11 -= int8(v4)
                        v3 = "hello"
                        switch v5 {
                        case 127:
                            var v16 uint = (^(v4))
                            _ = v16
                        }
                    } else {
                        v2 -= (((v2 * -2.25) - float32(-2.25)) * 0.1)
                    }
                }
            }
        case 571:
            var v17 uint32 = 4294967295
            _ = v17
        default:
            v4 -= (v4 / ((v4) | 1))
        }
    default:
        v2 += (1.5 - float32(1.5))
    }
    v3 = v3
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 uint64 = 0
        var v20 bool = false
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 string = v3
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    if (false && (!(v21 > v21))) {
        if ("foo bar" <= (v3 + v3)) {
            v3 = v3
        } else {
            v21 = "hello"
        }
        switch v21 {
        case "\n\t":
            fmt.Println(v4, v21, v4)
        default:
            var v22 int64 = -9223372036854775808
            _ = v22
        }
        if false {
            v23 := GenS5{F0: v3}
            fmt.Println(v21, v11)
            _ = v23
        }
    } else {
        if true {
            _ = len(v12)
        } else {
            for fi24 := 0; fi24 < 5; fi24++ {
                if false {
                    if false {
                        v11 += (v11 ^ 50)
                    }
                    v25 := GenS2{F0: (v21 + v21)}
                    _ = v25
                } else {
                    var v26 uint = (v4 ^ 707)
                    _ = v26
                }
                v3 = v21
                _ = fi24
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
