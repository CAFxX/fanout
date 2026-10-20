package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int16 }
type GenS2 struct { F0 bool }
func genFn0(p0 bool, p1 int32) bool {
    for fi7 := 0; fi7 < 18; fi7++ {
        p0 = (127 != 95)
        switch 42 {
        case 951:
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 int32 = -507
                _ = fi8
                _ = v9
            }
        case 32767:
            p0 = (("a" < "a") || (p0 || p0))
        default:
            p0 = (-3 >= 127)
        }
        _ = fi7
    }
    switch ("\n\t" + (("foo bar" + "hello") + "")) {
    case "a":
        if ((true && ((!true) && true)) || (!(3.14159 != (1.5 / 1.0)))) {
            var v10 int = 218
            p1 -= p1
            _ = v10
        } else {
            fmt.Println(p1, p0)
        }
    }
    v11 := GenS1{F0: 2147483647, F1: (702 <= 42), F2: 1000}
    p0 = p0
    var v12 bool = p0
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return v12
}
func genFn1(p0 uint16, p1 uint8, p2 int64) int8 {
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS2{F0: ((fi13 != (fi13 ^ (fi13 << ((2) & 63)))) && true)}
        fmt.Println(p2, p2, p0)
        p0 -= 1
        _ = fi13
        _ = v14
    }
    var v15 int64 = (-(p2))
    p0 = 32767
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 int8 = 38
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 0
}

func main() {
    var v1 int = 127
    var v2 uint = 432
    v3 := make(map[uint32]bool)
    var v4 uint16 = 32767
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]uint32)
    var v18 float32 = -2.25
    v19 := GenS1{F0: 1, F1: false, F2: -1}
    v20 := v3[2147483647]
    var v21 float64 = float64(v18)
    fmt.Println(v4, v21, v2)
    v22 := GenS1{F0: 4294967295, F1: v20, F2: 1}
    v1 -= -1
    v2 = v2
    if (255 != 205) {
        var v23 string = ("a" + (("hello" + ("\n\t" + "hello")) + (((("" + ("" + "hello")) + "") + "foo bar") + ("" + ("" + "\n\t")))))
        _ = v23
    } else {
        if (-9223372036854775808 > -9223372036854775808) {
            v2 -= 65535
            if v20 {
                var v24 bool = (2147483647 <= 42)
                _ = v24
            }
            var v25 uint32 = 4294967295
            _ = v25
        } else {
            if (!(!(v4 > v4))) {
                for fi26 := 0; fi26 < 2; fi26++ {
                    if v20 {
                        var v27 float64 = float64(v18)
                        _ = v27
                    }
                    _ = fi26
                }
            }
            v28 := GenS2{F0: false}
            switch (v1 + (v1 ^ 1)) {
            case 42:
                v2 = (^(v2))
            case 577:
                switch "hello" {
                case "hello":
                    delete(v6, 535)
                default:
                    v21 += v21
                }
            case -1:
                v20 = v20
            }
            _ = v28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
