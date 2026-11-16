package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint32, p1 uint8, p2 uint) int64 {
    if true {
        v5 := GenS1{F0: (p1 * (p1 << ((2) & 63))), F1: (p1 + p1)}
        var v6 float32 = 1.5
        v7 := []int{-337}
        _ = v5
        _ = v6
        _ = v7
    } else {
        var v8 string = ("hello" + (("a" + "\n\t") + ""))
        _ = v8
    }
    switch 61 {
    case -1:
        p1 = p1
    case 32767:
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 uint64 = 4294967295
            if ("a" >= "a") {
                p0 += (^(p0))
            }
            p2 += 4294967295
            _ = fi9
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        p2 = p2
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    return -66
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 860
    var v3 uint32 = 42
    v4 := []int{-2147483648, -60}
    switch 9223372036854775807 {
    case 9223372036854775807:
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS1{F0: 255, F1: 14}
            _ = fi12
            _ = v13
        }
    }
    var v14 string = ""
    v3 = 1
    var v15 bool = false
    var v16 uint = 1
    delete(v1, -9223372036854775808)
    for fi17 := 0; fi17 < 20; fi17++ {
        if v15 {
            for fi18 := 0; fi18 < 4; fi18++ {
                _ = len(v1)
                _ = fi18
            }
            v19 := GenS2{F0: 1, F1: 937}
            _ = v19
        } else {
            switch "\n\t" {
            case "\n\t":
                if true {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        fmt.Println(v3)
                        var v21 float32 = -2.25
                        _ = fi20
                        _ = v21
                    }
                }
            default:
                var v22 int64 = -288
                _ = v22
            }
            v23 := GenS2{F0: -1, F1: 580}
            _ = v23
        }
        if (2 >= 2) {
            if len(v4) > 0 {
                v4[0] = fi17
            }
            var v24 int32 = 870
            var v25 uint64 = 68
            _ = v24
            _ = v25
        }
        if (true || (!(!(false && (1 < (v16 ^ 65535)))))) {
            v2 = -1
            var v26 uint8 = 1
            v15 = false
            _ = v26
        }
        _ = fi17
    }
    switch "foo bar" {
    case "":
        if v15 {
            var v27 uint = 693
            var v28 uint8 = 150
            _ = v27
            _ = v28
        } else {
            v4 = append(v4, -2147483648)
        }
    case "\n\t":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
