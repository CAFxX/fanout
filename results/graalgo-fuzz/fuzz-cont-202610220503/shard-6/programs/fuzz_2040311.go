package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 bool; F2 uint16 }
func genFn0(p0 bool, p1 uint16, p2 int8) uint32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint8 = uint8(p1)
            p0 = ((p1 << ((519) & 63)) == 255)
            v10 := GenS1{F0: p1}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    fmt.Println(p2, p2, p2)
    p1 -= (p1 - (p1 << ((57) & 63)))
    p1 = (p1 | 2)
    for fi11 := 0; fi11 < 17; fi11++ {
        fmt.Println(p2, fi11)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int8 = 42
    var v2 float32 = 0.1
    v3 := make(map[uint32]bool)
    var v4 int16 = 172
    v5 := make(map[uint32]bool)
    var v6 float32 = ((v2 * 0.0) + v2)
    var v12 int64 = 9223372036854775807
    fmt.Println(v4, v1, v1)
    fmt.Println(v6, v6, v12)
    v13 := GenS1{F0: 378}
    var v14 bool = true
    var v15 uint32 = 42
    if v14 {
        var v16 int = -2147483648
        v17 := GenS2{F0: v4, F1: (v16 > -2147483648), F2: 32767}
        _ = v16
        _ = v17
    }
    switch 42 {
    case -9223372036854775808:
        var v18 uint32 = v15
        _ = v18
    default:
        for fi19 := 0; fi19 < 15; fi19++ {
            if (true && (1 <= 226)) {
                var v20 uint32 = 62
                _ = v20
            }
            if v14 {
                switch ("hello" + ((("foo bar" + "hello") + ("" + "")) + ("" + ("hello" + "")))) {
                case "\n\t":
                    var v21 uint16 = 2
                    _ = v21
                case "":
                    var v22 int16 = (^(v4))
                    _ = v22
                default:
                    v2 -= v2
                }
                v23 := GenS2{F0: 1, F1: (-32768 > v4), F2: 727}
                v6 = (v6 - 3.14159)
                _ = v23
            } else {
                v14 = v14
                fmt.Println(v6, v2)
            }
            switch (fi19 & 32767) {
            case 2147483647:
                v12 = v12
            case -734:
                for fi24 := 0; fi24 < 16; fi24++ {
                    var v25 int8 = (v1 + (v1 & v1))
                    _ = fi24
                    _ = v25
                }
            case 32767:
                fmt.Println(v1)
            default:
                v1 += v1
            }
            _ = fi19
        }
    }
    if (-2.25 != 1.5) {
        v26 := []int{128, 2147483647, 32768, 42}
        for fi27 := 0; fi27 < 2; fi27++ {
            var v28 int16 = v4
            _ = fi27
            _ = v28
        }
        fmt.Println(v6, v2)
        _ = v26
    }
    v5[v15] = ((false && ("\n\t" < ("\n\t" + ""))) || (!((v1 & int8(v4)) >= v1)))
    switch "a" {
    case "\n\t":
        switch 522 {
        case -2147483648:
            v2 -= -2.25
        default:
            v29 := []int{632}
            _ = v29
        }
    default:
        fmt.Println(v15, v2, v15)
    }
    var v30 uint = 42
    var v31 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
