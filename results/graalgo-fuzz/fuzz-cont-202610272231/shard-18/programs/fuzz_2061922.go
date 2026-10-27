package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint }
func genFn0() uint16 {
    var v6 uint = 395
    var v7 uint32 = 2
    fmt.Println(v7, v6)
    switch ("foo bar" + (("hello" + "foo bar") + "")) {
    case "a":
        for fi8 := 0; fi8 < 3; fi8++ {
            if (true || false) {
                v9 := GenS1{F0: ("" <= "foo bar"), F1: 194, F2: 0}
                for fi10 := 0; fi10 < 19; fi10++ {
                    v6 = 575
                    _ = fi10
                }
                _ = v9
            }
            _ = fi8
        }
    default:
        var v11 string = "hello"
        _ = v11
    }
    if true {
        v7 -= v7
        v7 = 2147483648
        fmt.Println(v7)
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := GenS1{F0: (-66 > 70), F1: 127, F2: 105}
            v14 := GenS1{F0: ((("\n\t" + "\n\t") + "") != ("foo bar" + "a")), F1: 255, F2: 128}
            fmt.Println(fi12, fi12, v6)
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := GenS2{F0: -128, F1: 74, F2: v6}
        if (false && (!(false || (-593 < 4294967295)))) {
            var v16 bool = false
            _ = v16
        } else {
            v17 := []int{128, -605}
            v7 -= (v7 * 1)
            _ = v17
        }
        _ = v15
    }
    _ = v6
    _ = v7
    return 845
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[string]bool)
    v3 := []uint32{399, 2147483648}
    var v4 uint32 = 65535
    var v5 uint16 = 485
    _ = len(v2)
    if len(v3) > 0 {
        v3[0] = (^(v4))
    }
    v4 -= v4
    v5 = genFn0()
    v5 += 1
    var v18 uint8 = 128
    for fi19 := 0; fi19 < 6; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            v5 += v5
            _ = fi20
        }
        v1 -= v1
        _ = fi19
    }
    fmt.Println(v5, v18)
    var v21 int32 = 444
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
