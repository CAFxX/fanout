package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint32 }
func genFn0(p0 int32, p1 int16) uint64 {
    var v4 float32 = -2.25
    fmt.Println(p1, v4, p0)
    _ = p0
    _ = p1
    _ = v4
    return 18446744073709551615
}
func genFn1(p0 uint64) int16 {
    fmt.Println(p0)
    p0 = (p0 / ((18446744073709551615) | 1))
    p0 = (p0 >> ((2) & 63))
    var v5 string = (("\n\t" + ("hello" + ("" + "hello"))) + ("\n\t" + ("\n\t" + "foo bar")))
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 int64 = -1
    var v2 string = "foo bar"
    var v3 uint = 1
    var v6 uint16 = 32767
    if false {
        v1 = v1
        switch v2 {
        case "a":
            for fi7 := 0; fi7 < 18; fi7++ {
                v3 += v3
                _ = fi7
            }
        }
        v6 -= (^(v6))
    } else {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := []int{100, 1, (-(fi8)), -9223372036854775808}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            v6 = uint16(v3)
            _ = fi8
            _ = v9
        }
        v10 := []int{2147483647, 0}
        _ = v10
    }
    switch (v2 + ("\n\t" + v2)) {
    case "hello":
        v11 := []int{128, 42}
        _ = v11
    default:
        if (("\n\t" != ("hello" + v2)) || (-1 >= 42)) {
            v12 := []int{-1, 32767, -935, 4294967295}
            var v13 int8 = 100
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: 1}
            var v15 int64 = (v1 + v1)
            _ = v14
            _ = v15
        }
    }
    switch 9223372036854775807 {
    case 32768:
        v16 := GenS1{F0: -128}
        _ = v16
    default:
        fmt.Println(v1, v2, v6)
    }
    fmt.Println(v2, v1)
    if false {
        v2 = (v2 + v2)
        v17 := GenS2{F0: 0, F1: 413, F2: 65535}
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 20; fi19++ {
                v20 := []int{}
                var v21 float32 = 0.1
                _ = fi19
                _ = v20
                _ = v21
            }
            if (!false) {
                var v22 int8 = -128
                v23 := GenS1{F0: (v22 & int8(v6))}
                fmt.Println(fi18)
                _ = v22
                _ = v23
            }
            _ = fi18
        }
        _ = v17
    }
    var v24 int = -831
    switch ("foo bar" + ("" + v2)) {
    case "\n\t":
        v1 -= (v1 << (((v3 << ((1) & 63))) & 63))
    }
    v25 := GenS1{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v6
}
