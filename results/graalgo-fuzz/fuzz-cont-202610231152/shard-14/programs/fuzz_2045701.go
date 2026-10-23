package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
func genFn0(p0 bool) uint {
    p0 = false
    var v5 uint32 = 0
    _ = p0
    _ = v5
    return 731
}
func genFn1(p0 int, p1 int, p2 int32) int8 {
    var v6 uint32 = 2147483648
    switch "\n\t" {
    case "a":
        var v7 int8 = 100
        _ = v7
    default:
        v8 := GenS1{F0: ((false || (401 > 42)) && (false && (!(((float32(1.5) / ((float32((0.1 * (0.0 + 1.5))) - (0.1 - 0.0)) / 1.5)) - ((0.1 + 3.14159) * -2.25)) <= -2.25))))}
        _ = v8
    }
    var v9 uint16 = 1
    v10 := []int{(p1 << ((984) & 63)), p1, (p1 >> ((1) & 63)), (p0 + p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v9
    return -24
}

func main() {
    var v1 bool = ((0 != 531) || true)
    v2 := make(map[uint32]bool)
    var v3 int32 = 634
    var v4 uint = 189
    if (((880 != 2) && false) || v1) {
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := v2[2147483648]
            _ = fi11
            _ = v12
        }
    } else {
        switch 128 {
        case 32767:
            v3 += 1
        case 42:
            fmt.Println(v3, v3, v3)
        }
        v3 -= 2147483647
    }
    v4 = (^(v4))
    v2[2] = ((127 >= 128) || v1)
    v13 := GenS1{F0: v1}
    for fi14 := 0; fi14 < 3; fi14++ {
        v4 = v4
        fmt.Println(v4)
        _ = fi14
    }
    switch 9223372036854775807 {
    case 0:
        var v15 int16 = -32768
        _ = v15
    default:
        v16 := v2[2147483648]
        _ = v16
    }
    var v17 bool = (!v1)
    v3 -= (v3 / ((2147483647) | 1))
    v18 := GenS1{F0: (v3 == v3)}
    _ = len(v2)
    if (true || (42 < -1)) {
        _ = len(v2)
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 int32 = (v3 % ((v3) | 1))
            switch ("" + "a") {
            case "foo bar":
                _ = len(v2)
            case "":
                for fi21 := 0; fi21 < 18; fi21++ {
                    switch "foo bar" {
                    case "\n\t":
                        var v22 uint32 = 462
                        v23 := GenS2{F0: "\n\t"}
                        v3 -= v20
                        _ = v22
                        _ = v23
                    case "":
                        switch "foo bar" {
                        case "hello":
                            var v24 bool = ((v20 - (v3 & -374)) > v20)
                            _ = v24
                        }
                    }
                    _ = fi21
                }
            default:
                v25 := []int{(^(fi19)), 1}
                _ = v25
            }
            _ = fi19
            _ = v20
        }
        var v26 float32 = 3.14159
        _ = v26
    } else {
        switch -146 {
        case -9223372036854775808:
            for fi27 := 0; fi27 < 18; fi27++ {
                v3 -= -1
                _ = fi27
            }
        }
    }
    if false {
        var v28 bool = v1
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
