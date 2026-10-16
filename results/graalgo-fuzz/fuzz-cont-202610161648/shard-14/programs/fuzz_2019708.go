package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() bool {
    v4 := []int{}
    var v5 uint16 = 42
    fmt.Println(v5)
    v5 = (v5 ^ (v5 | 32767))
    if true {
        v5 = (v5 << ((4294967295) & 63))
    }
    _ = v4
    _ = v5
    return (!false)
}
func genFn1(p0 int8, p1 string) int {
    p0 -= (p0 - (p0 - 80))
    var v6 uint32 = 65535
    if false {
        var v7 uint64 = 2
        for fi8 := 0; fi8 < 12; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 bool = (-9223372036854775808 <= -675)
                p1 = "foo bar"
                if ((fi9 ^ 353) < 0) {
                    v11 := []int{(fi9 & 42)}
                    _ = v11
                } else {
                    if ((!true) && true) {
                        fmt.Println(v6, fi8, p0)
                        fmt.Println(p0, fi9)
                        var v12 uint32 = 2
                        _ = v12
                    } else {
                        fmt.Println(v6, p0, v6)
                        v6 += (^(v6))
                        v13 := GenS1{F0: 2}
                        _ = v13
                    }
                }
                _ = fi9
                _ = v10
            }
            p1 = p1
            _ = fi8
        }
        _ = v7
    }
    v6 += (v6 - v6)
    p1 = "foo bar"
    _ = p0
    _ = p1
    _ = v6
    return -859
}

func main() {
    v1 := []string{"hello", ("\n\t" + "")}
    var v2 uint64 = 42
    var v3 int8 = -96
    switch 9223372036854775807 {
    case 9223372036854775807:
        v2 -= (^(v2))
    default:
        fmt.Println(v3, v2)
    }
    switch 128 {
    case 9223372036854775807:
        for fi14 := 0; fi14 < 4; fi14++ {
            v3 += 58
            _ = fi14
        }
    case -2147483648:
        var v15 uint8 = 200
        _ = v15
    }
    if ((genFn1((v3 + (v3 ^ 0)), "\n\t") != -545) || false) {
        v2 += v2
        if false {
            v3 = v3
        } else {
            var v16 uint = 255
            _ = v16
        }
        var v17 uint16 = 1
        _ = v17
    }
    if (18446744073709551615 != 516) {
        switch 243 {
        case 425:
            var v18 uint32 = 1
            _ = v18
        }
        v19 := GenS1{F0: 42}
        _ = v19
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = v2
    fmt.Println(v3, v3)
    if true {
        v2 = (v2 / (((v2 - (v2 << ((2) & 63)))) | 1))
    } else {
        var v20 uint16 = 0
        switch "\n\t" {
        case "":
            v2 = 145
        }
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
