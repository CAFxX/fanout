package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint16 }

func main() {
    var v1 uint8 = 255
    var v2 float64 = ((((((-2.25 / 1.0) + 1e10) + (3.14159 - (0.1 * (-2.25 + (0.1 * 0.1))))) - 0.0) - 1e10) * 1e10)
    var v3 uint32 = 576
    v4 := make(map[uint32]string)
    var v5 uint16 = 773
    fmt.Println(v1)
    var v6 string = ("foo bar" + ((((("hello" + "\n\t") + "foo bar") + ((("foo bar" + "\n\t") + ("foo bar" + "\n\t")) + (("a" + ("a" + "a")) + "\n\t"))) + ("a" + ("a" + ""))) + ("a" + "hello")))
    v7 := []int{4294967295, 127, -555}
    _ = len(v4)
    v2 -= (v2 / 1e10)
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint8 = (v1 & (v1 * (v1 / (((v1 + 128)) | 1))))
        if len(v7) > 0 {
            v7[0] = fi8
        }
        var v10 string = v6
        _ = fi8
        _ = v10
        _ = v9
    }
    if (!((v3 << ((703) & 63)) == v3)) {
        for fi11 := 0; fi11 < 9; fi11++ {
            v2 = ((v2 - v2) - v2)
            _ = fi11
        }
        switch 127 {
        case -1:
            v2 = v2
        case 1:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        default:
            for fi12 := 0; fi12 < 17; fi12++ {
                v3 -= v3
                var v13 int16 = -739
                _ = len(v7)
                _ = fi12
                _ = v13
            }
        }
    } else {
        v2 = 3.14159
        if false {
            for fi14 := 0; fi14 < 3; fi14++ {
                if true {
                    var v15 uint32 = (v3 ^ (v3 % ((v3) | 1)))
                    var v16 uint32 = (v3 % ((351) | 1))
                    _ = v15
                    _ = v16
                }
                var v17 int16 = 42
                var v18 int32 = 1
                _ = fi14
                _ = v17
                _ = v18
            }
            fmt.Println(v6)
            v7 = append(v7, 0)
        }
    }
    if (42 < -136) {
        var v19 uint8 = v1
        for fi20 := 0; fi20 < 18; fi20++ {
            if (100 != 42) {
                var v21 int64 = 0
                v3 = 0
                _ = v21
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                fmt.Println(v19)
                v5 = (v5 % (((v5 & v5)) | 1))
                _ = fi22
            }
            v6 = (v6 + (("a" + v6) + (((v6 + v6) + v6) + v6)))
            _ = fi20
        }
        v7 = append(v7, -2147483648)
        _ = v19
    } else {
        v23 := GenS1{F0: ("a" + (v6 + ((v6 + "foo bar") + "hello"))), F1: 42, F2: v5}
        v24 := GenS1{F0: v6, F1: (^(v1)), F2: (^(v5))}
        for fi25 := 0; fi25 < 19; fi25++ {
            fmt.Println(v3)
            _ = fi25
        }
        _ = v23
        _ = v24
    }
    fmt.Println(v5, v2, v2)
    v1 += (v1 % ((v1) | 1))
    var v26 uint16 = 713
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
