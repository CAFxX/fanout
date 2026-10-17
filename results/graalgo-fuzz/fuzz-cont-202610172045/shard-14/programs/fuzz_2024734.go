package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint8 }
type GenS2 struct { F0 bool }
func genFn0() uint8 {
    var v4 uint64 = 94
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: (!(((9223372036854775807 == -9223372036854775808) || false) && (!(!(!((0.1 - 0.0) < 0.1)))))), F1: 42, F2: 42}
        _ = fi5
        _ = v6
    }
    var v7 int16 = 659
    _ = v4
    _ = v7
    return 42
}
func genFn1() int32 {
    var v8 uint = 65535
    var v9 uint16 = 1
    _ = v8
    _ = v9
    return -200
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "\n\t"
    var v3 float32 = 1.5
    v3 = 0.1
    if (true || (42 >= -689)) {
        var v10 uint = 2
        _ = v10
    } else {
        v11 := []int{}
        switch 9223372036854775807 {
        case 42:
            switch 0 {
            case -9223372036854775808:
                _ = len(v1)
            case 886:
                if (1000 > 32767) {
                    fmt.Println(v3, v2)
                    v2 = ((v2 + "") + v2)
                }
            case 32767:
                for fi12 := 0; fi12 < 10; fi12++ {
                    fmt.Println(v2, v2)
                    v13 := GenS2{F0: (0 == 1)}
                    _ = fi12
                    _ = v13
                }
            }
        default:
            var v14 bool = ((!(-9223372036854775808 > 1)) || false)
            _ = v14
        }
        _ = v11
    }
    v3 += v3
    for fi15 := 0; fi15 < 7; fi15++ {
        _ = len(v1)
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 float32 = v3
            v2 = (v2 + ("a" + v2))
            v18 := GenS2{F0: (615 != 42)}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    switch (v2 + v2) {
    case "a":
        var v19 uint8 = 2
        _ = v19
    default:
        v2 = v2
    }
    switch 1 {
    case -9223372036854775808:
        v3 -= v3
    case 127:
        for fi20 := 0; fi20 < 18; fi20++ {
            fmt.Println(v3, v3, fi20)
            _ = fi20
        }
    default:
        fmt.Println(v2, v3)
    }
    if ((!true) && (42 >= 0)) {
        for fi21 := 0; fi21 < 10; fi21++ {
            v3 = 0.1
            _ = fi21
        }
        var v22 int32 = -1
        _ = v22
    }
    var v23 int8 = 1
    v24 := []int{}
    fmt.Println(v3, v23, v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
