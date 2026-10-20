package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 bool; F2 int16 }
func genFn0(p0 int64, p1 int) int8 {
    fmt.Println(p1)
    switch (p1 % ((p1) | 1)) {
    case 128:
        p1 += (p1 & p1)
    }
    v5 := []int{-2147483648, 32768, (p1 | (p1 >> ((65535) & 63)))}
    p0 -= (p0 >> ((4294967295) & 63))
    _ = p0
    _ = p1
    _ = v5
    return -67
}

func main() {
    var v1 bool = (2147483647 > 1)
    var v2 uint8 = 1
    var v3 uint16 = 42
    var v4 bool = v1
    var v6 string = (("foo bar" + ("hello" + "foo bar")) + "foo bar")
    v7 := GenS1{F0: 1000, F1: 127, F2: -9223372036854775808}
    var v8 int64 = -811
    if v1 {
        switch "foo bar" {
        case "a":
            var v9 float32 = 1.5
            _ = v9
        case "hello":
            switch 87 {
            case 117:
                if (!(v1 && (42 > 18446744073709551615))) {
                    for fi10 := 0; fi10 < 9; fi10++ {
                        for fi11 := 0; fi11 < 17; fi11++ {
                            v12 := []int{fi10, (fi10 >> ((18446744073709551615) & 63))}
                            _ = fi11
                            _ = v12
                        }
                        _ = fi10
                    }
                }
            case 1:
                v13 := []int{32768, 42, -546, 556}
                _ = v13
            default:
                v3 = (v3 - v3)
            }
        }
        var v14 bool = (!(v1 || v1))
        _ = v14
    }
    v6 = "hello"
    switch "foo bar" {
    case "hello":
        var v15 bool = ((235 > 0) || (3.14159 <= 0.0))
        _ = v15
    case "\n\t":
        v16 := GenS2{F0: (!(32768 >= -1))}
        _ = v16
    }
    var v17 float64 = 1e10
    v18 := GenS3{F0: (!(3.14159 != 0.1)), F1: v4, F2: -527}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
