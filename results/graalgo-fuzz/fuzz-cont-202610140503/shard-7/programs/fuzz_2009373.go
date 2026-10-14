package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 bool }
func genFn0(p0 uint64, p1 uint, p2 int16) uint16 {
    v5 := GenS1{F0: false}
    if true {
        p2 -= -1
        p2 = (p2 - p2)
    } else {
        if (false && (-90 != -2147483648)) {
            var v6 int64 = -1
            _ = v6
        } else {
            if (0 == (p2 << ((p1) & 63))) {
                var v7 uint64 = (p0 / (((p0 + (p0 << (((p1 & (p1 ^ p1))) & 63)))) | 1))
                _ = v7
            } else {
                fmt.Println(p2, p1)
                p1 += (p1 + 262)
            }
            if true {
                switch -203 {
                case 0:
                    var v8 uint16 = 65535
                    _ = v8
                case 2147483647:
                    var v9 float64 = (0.1 * float64(-2.25))
                    _ = v9
                default:
                    for fi10 := 0; fi10 < 12; fi10++ {
                        if (0 > 0) {
                            switch "\n\t" {
                            case "":
                                p2 -= 484
                            }
                            v11 := GenS1{F0: false}
                            v12 := GenS1{F0: true}
                            _ = v11
                            _ = v12
                        }
                        var v13 bool = false
                        _ = fi10
                        _ = v13
                    }
                }
                v14 := GenS2{F0: 2}
                _ = v14
            } else {
                for fi15 := 0; fi15 < 1; fi15++ {
                    p0 = 4294967295
                    fmt.Println(p0)
                    _ = fi15
                }
                fmt.Println(p1, p1)
            }
        }
        fmt.Println(p0, p2, p1)
    }
    var v16 int16 = 925
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS2{F0: 398}
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v5
    return uint16(p1)
}

func main() {
    var v1 uint16 = 32767
    var v2 uint64 = 146
    var v3 string = "foo bar"
    var v4 int32 = 42
    v19 := GenS1{F0: ((("foo bar" + (("hello" + v3) + v3)) <= v3) && (-2147483648 <= 2147483647))}
    v4 -= -2147483648
    v20 := []int{127, 128}
    switch 747 {
    case -405:
        var v21 int = 4294967295
        _ = v21
    case 1:
        var v22 string = (v3 + ("hello" + "\n\t"))
        _ = v22
    }
    v23 := GenS2{F0: 1}
    fmt.Println(v2, v1)
    if true {
        var v24 uint64 = 4294967295
        _ = v24
    }
    switch "a" {
    case "a":
        var v25 int32 = v4
        _ = v25
    case "foo bar":
        v4 -= v4
    case "\n\t":
        if len(v20) > 0 {
            v20[0] = -947
        }
    }
    v26 := GenS3{F0: (v3 + (((v3 + (v3 + "foo bar")) + "") + (v3 + "\n\t")))}
    v2 = v2
    if false {
        fmt.Println(v1, v3)
        _ = len(v20)
        var v27 uint8 = 255
        _ = v27
    }
    var v28 uint64 = (v2 % ((2) | 1))
    if (32767 != v1) {
        for fi29 := 0; fi29 < 6; fi29++ {
            var v30 string = v3
            var v31 int = 9223372036854775807
            switch "a" {
            case "hello":
                v3 = v3
            case "a":
                v31 += (-(v31))
            default:
                v32 := GenS4{F0: 42, F1: (!false)}
                _ = v32
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        if false {
            v2 -= 2
        }
        v2 = v28
    } else {
        fmt.Println(v1, v2)
        var v33 uint32 = 4294967295
        _ = v33
    }
    var v34 float32 = (1e10 - 3.14159)
    v35 := GenS4{F0: 4294967295, F1: false}
    v4 += 772
    v28 -= v28
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v26
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
