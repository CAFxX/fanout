package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
func genFn0(p0 int64, p1 int64) int32 {
    v5 := GenS1{F0: (18446744073709551615 < 681), F1: 42}
    switch "\n\t" {
    case "a":
        p0 += p1
    case "\n\t":
        v6 := []int{}
        _ = v6
    case "":
        p0 -= (p0 | 0)
    }
    p0 = p1
    p1 += p0
    _ = p0
    _ = p1
    _ = v5
    return 261
}
func genFn1(p0 int, p1 uint64) int8 {
    p1 = (p1 / ((p1) | 1))
    switch (-(p0)) {
    case 166:
        v7 := GenS1{F0: true, F1: 30}
        _ = v7
    case 1:
        v8 := []int{}
        _ = v8
    case 4294967295:
        if ((0 == 2) || (695 > 255)) {
            var v9 uint16 = 1
            _ = v9
        } else {
            if false {
                p0 = 0
                if (9223372036854775807 > -277) {
                    p1 += (^(p1))
                } else {
                    for fi10 := 0; fi10 < 18; fi10++ {
                        p0 -= (fi10 << ((4294967295) & 63))
                        _ = fi10
                    }
                    fmt.Println(p1)
                    p1 = (p1 | (^(p1)))
                }
            } else {
                fmt.Println(p1, p0)
                if false {
                    p0 += (p0 | (-(p0)))
                }
            }
        }
    default:
        p0 = -2147483648
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 string = ""
    var v2 uint32 = 2
    var v3 bool = false
    var v4 uint16 = 2
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS1{F0: v3, F1: 2}
        v2 = (v2 % ((v2) | 1))
        _ = fi11
        _ = v12
    }
    var v13 float64 = 3.14159
    v1 = v1
    v3 = (0.1 != ((((1.5 * 3.14159) + -2.25) / (-2.25 / 1e10)) * ((1.5 * 1e10) + float32(v13))))
    switch "hello" {
    case "":
        var v14 string = v1
        _ = v14
    case "foo bar":
        if (42 > -9223372036854775808) {
            v4 = v4
        } else {
            v15 := GenS1{F0: ((!(-2147483648 == 0)) || v3), F1: 20}
            fmt.Println(v13, v3, v3)
            _ = v15
        }
    }
    if (774 < -1) {
        var v16 uint = 350
        v17 := []int{1, 2147483647}
        _ = v16
        _ = v17
    }
    var v18 uint64 = 471
    var v19 uint64 = (v18 | v18)
    v3 = false
    v19 = (v18 + (v18 ^ (v19 & 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
