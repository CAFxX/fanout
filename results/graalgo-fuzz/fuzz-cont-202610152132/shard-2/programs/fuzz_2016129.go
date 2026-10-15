package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
func genFn0(p0 int64, p1 int8) uint16 {
    v4 := GenS1{F0: "hello"}
    v5 := GenS2{F0: -1}
    v6 := []int{-2147483648, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 65535
}
func genFn1() int32 {
    if (1 != -339) {
        var v7 string = "a"
        _ = v7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := []int{(fi8 * fi8)}
        var v10 int64 = -1
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 string = ("hello" + "")
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := []int{1, (-(fi12)), fi13, (fi12 ^ fi13)}
            if (-128 > 0) {
                if false {
                    var v15 float32 = 1e10
                    _ = v15
                }
            }
            fmt.Println(fi13, fi12)
            fmt.Println(fi13)
            _ = fi13
            _ = v14
        }
        var v16 int16 = -222
        v11 = v11
        _ = fi12
        _ = v16
    }
    v11 = ((v11 + ("foo bar" + (("hello" + ("foo bar" + "foo bar")) + "a"))) + "a")
    _ = v11
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 0
    var v3 uint = 0
    v3 = v3
    switch (v2 * (v2 % (((v2 ^ (v2 + (v2 & v2)))) | 1))) {
    case 1:
        var v17 int64 = 42
        _ = v17
    default:
        var v18 uint8 = uint8(v3)
        _ = v18
    }
    var v19 int8 = -44
    if (65535 < 84) {
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 int8 = (^(v19))
            _ = fi20
            _ = v21
        }
    } else {
        v19 = v19
        switch 127 {
        case -1:
            for fi22 := 0; fi22 < 9; fi22++ {
                switch "\n\t" {
                case "a":
                    var v23 uint64 = 0
                    _ = v23
                }
                if (1 != -185) {
                    fmt.Println(fi22, v2, v2)
                    var v24 float32 = 3.14159
                    var v25 int8 = v19
                    _ = v24
                    _ = v25
                }
                fmt.Println(v3, v2, v2)
                _ = fi22
            }
        case 1:
            var v26 float64 = 3.14159
            _ = v26
        case 653:
            fmt.Println(v2)
        }
        v27 := []int{v2, 32767, (v2 ^ (v2 & v2))}
        _ = v27
    }
    fmt.Println(v19, v3, v19)
    fmt.Println(v19)
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 float32 = -2.25
        v1[fi28] = true
        _ = fi28
        _ = v29
    }
    var v30 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
}
