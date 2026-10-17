package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
func genFn0(p0 uint8, p1 uint32) uint {
    if true {
        var v7 int8 = 1
        p1 = p1
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int8 = -1
        _ = fi8
        _ = v9
    }
    v10 := []int{}
    _ = p0
    _ = p1
    _ = v10
    return 65535
}

func main() {
    var v1 int64 = 1
    var v2 uint64 = 940
    v3 := make(map[int]string)
    var v4 string = "a"
    var v5 uint8 = 26
    var v6 string = (v4 + "a")
    var v11 int16 = -305
    v11 = (v11 / ((v11) | 1))
    v12 := GenS1{F0: -726, F1: genFn0(v5, 65535)}
    for fi13 := 0; fi13 < 10; fi13++ {
        v6 = (("foo bar" + "hello") + ((v6 + v6) + v4))
        switch "hello" {
        case "":
            var v14 int16 = (-(v11))
            _ = v14
        default:
            v3[fi13] = (("" + "foo bar") + "hello")
        }
        _ = fi13
    }
    v6 = v6
    v4 = (v6 + v4)
    if false {
        for fi15 := 0; fi15 < 8; fi15++ {
            if (!(!((v1 ^ (v1 / (((v1 & (^(v1)))) | 1))) != int64(v11)))) {
                v2 += (v2 & (v2 >> ((878) & 63)))
                if false {
                    fmt.Println(v2, fi15)
                    fmt.Println(v1, v4)
                } else {
                    var v16 int16 = v11
                    v4 = v4
                    _ = v16
                }
                for fi17 := 0; fi17 < 12; fi17++ {
                    v18 := GenS1{F0: 1, F1: 42}
                    v19 := GenS1{F0: 2147483647, F1: genFn0(v5, 42)}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            v4 = "hello"
            if (!((v11 >> ((2) & 63)) > v11)) {
                var v20 uint = 0
                fmt.Println(v6)
                delete(v3, (fi15 - fi15))
                _ = v20
            }
            _ = fi15
        }
        fmt.Println(v11, v1)
    }
    if (true || true) {
        var v21 uint8 = (v5 << ((255) & 63))
        v22 := []int{601, 128, 32767, 2147483647}
        fmt.Println(v11, v21)
        _ = v21
        _ = v22
    } else {
        switch 32767 {
        case 9223372036854775807:
            var v23 uint16 = 32767
            _ = v23
        case -899:
            switch 42 {
            case 384:
                v24 := []int{-1, 4294967295}
                _ = v24
            case -2147483648:
                v6 = "\n\t"
            default:
                v6 = v6
            }
        case 4294967295:
            v1 -= (v1 * (v1 | (v1 % (((v1 | int64(v5))) | 1))))
        }
    }
    switch (((("" + v6) + "") + "") + "\n\t") {
    case "":
        var v25 uint64 = (^(v2))
        _ = v25
    case "foo bar":
        fmt.Println(v11)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
