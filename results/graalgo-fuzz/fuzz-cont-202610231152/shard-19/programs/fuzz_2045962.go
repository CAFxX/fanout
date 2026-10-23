package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int8, p1 int64, p2 uint32) uint32 {
    p2 = 65535
    fmt.Println(p1)
    var v4 bool = (int64(p0) <= (p1 >> ((4294967295) & 63)))
    v5 := GenS1{F0: 9223372036854775807}
    var v6 int8 = -62
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return p2
}
func genFn1(p0 uint, p1 int8) uint64 {
    switch (("a" + "a") + "hello") {
    case "hello":
        v7 := []int{-60, -9223372036854775808, 0}
        _ = v7
    }
    switch 0 {
    case 9223372036854775807:
        switch ("a" + ("hello" + ("\n\t" + "a"))) {
        case "a":
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 uint32 = 423
                var v10 string = "foo bar"
                if false {
                    p1 = int8(fi8)
                    var v11 float64 = 0.0
                    _ = v11
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        case "foo bar":
            var v12 int64 = 9223372036854775807
            _ = v12
        case "":
            for fi13 := 0; fi13 < 8; fi13++ {
                p1 = (p1 % ((0) | 1))
                for fi14 := 0; fi14 < 9; fi14++ {
                    p1 = p1
                    _ = fi14
                }
                fmt.Println(fi13, p1)
                _ = fi13
            }
        default:
            var v15 uint16 = 32767
            _ = v15
        }
    case 4294967295:
        v16 := GenS2{F0: 2147483647}
        _ = v16
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 int = -482
    var v2 int64 = -450
    var v3 uint64 = 445
    var v17 int = (v1 % ((32767) | 1))
    v17 = v17
    v1 += (v17 % ((v1) | 1))
    switch "foo bar" {
    case "hello":
        var v18 uint = 423
        _ = v18
    case "\n\t":
        var v19 float32 = 0.0
        _ = v19
    case "a":
        for fi20 := 0; fi20 < 20; fi20++ {
            v1 += (fi20 % ((fi20) | 1))
            fmt.Println(v17)
            _ = fi20
        }
    default:
        var v21 float64 = 3.14159
        _ = v21
    }
    v22 := GenS3{F0: genFn0(int8(v3), v2, 2147483647)}
    var v23 bool = (v2 <= (v2 ^ v2))
    fmt.Println(v23, v23, v17)
    var v24 uint32 = 509
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
