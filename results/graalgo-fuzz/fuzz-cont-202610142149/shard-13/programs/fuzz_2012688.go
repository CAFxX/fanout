package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint8, p1 uint, p2 int8) string {
    var v7 float32 = 1.5
    var v8 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ((("hello" + "foo bar") + ("a" + ((("" + (("" + "hello") + (("a" + "\n\t") + "hello"))) + "a") + ("hello" + ("foo bar" + ("a" + "")))))) + "hello")
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    var v3 uint32 = 2147483648
    v4 := []int{717, 195, 0}
    var v5 uint8 = 1
    v6 := make(map[int]uint32)
    fmt.Println(v5, v3)
    v5 = v5
    v9 := GenS1{F0: 100, F1: true, F2: 0}
    switch 773 {
    case -1:
        v3 -= (v3 % ((v3) | 1))
    case 493:
        v10 := GenS2{F0: 0}
        _ = v10
    case 128:
        for fi11 := 0; fi11 < 19; fi11++ {
            v5 = (v5 ^ (v5 ^ v5))
            _ = fi11
        }
    default:
        var v12 uint32 = v3
        _ = v12
    }
    fmt.Println(v5, v3)
    if ((false && true) && ((((v5 * v5) < (v5 & v5)) && false) || false)) {
        for fi13 := 0; fi13 < 8; fi13++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v5 = v5
            v14 := GenS3{F0: 946}
            _ = fi13
            _ = v14
        }
    } else {
        var v15 uint = 923
        if true {
            v15 = (v15 << (((v15 * (^(v15)))) & 63))
        } else {
            fmt.Println(v5)
            v15 = 565
            switch 951 {
            case -742:
                for fi16 := 0; fi16 < 14; fi16++ {
                    if false {
                        var v17 float64 = 0.0
                        v5 -= v5
                        if (false && (-1 < 42)) {
                            var v18 bool = false
                            _ = v18
                        }
                        _ = v17
                    }
                    _ = fi16
                }
            case 32768:
                for fi19 := 0; fi19 < 7; fi19++ {
                    var v20 int8 = -128
                    v3 = 0
                    if (!true) {
                        v15 -= (v15 + 4294967295)
                    }
                    _ = fi19
                    _ = v20
                }
            case -407:
                v15 = (^(v15))
            }
        }
        v21 := GenS4{F0: (2147483647 >= 0)}
        _ = v15
        _ = v21
    }
    switch "hello" {
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "foo bar":
        var v22 int32 = 0
        _ = v22
    default:
        var v23 string = ((("\n\t" + ("a" + ("\n\t" + ""))) + "a") + ("a" + "hello"))
        _ = v23
    }
    _ = len(v1)
    if false {
        v24 := GenS3{F0: 65535}
        for fi25 := 0; fi25 < 9; fi25++ {
            v3 += v3
            var v26 bool = (false || (0 != -737))
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
