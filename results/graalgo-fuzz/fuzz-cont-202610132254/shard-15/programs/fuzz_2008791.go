package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint }
type GenS2 struct { F0 uint32; F1 uint32; F2 int32 }

func main() {
    var v1 string = "hello"
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[int]bool)
    v5 := v2[-356]
    switch "hello" {
    case "":
        v5 = false
    }
    var v6 int = -1
    switch "foo bar" {
    case "hello":
        v7 := []int{v6, -648, (v6 * v6)}
        _ = v7
    case "a":
        switch v6 {
        case 9223372036854775807:
            var v8 uint16 = 4
            _ = v8
        case 128:
            var v9 uint16 = 854
            _ = v9
        }
    }
    fmt.Println(v5)
    for fi10 := 0; fi10 < 9; fi10++ {
        if v5 {
            v1 = ""
        }
        var v11 uint16 = 1
        _ = fi10
        _ = v11
    }
    v5 = ((200 > 155) && ((v1 + v1) <= ((v1 + "foo bar") + (v1 + v1))))
    v12 := GenS1{F0: -1, F1: 0, F2: 509}
    for fi13 := 0; fi13 < 7; fi13++ {
        v1 = ((v1 + v1) + "a")
        if v5 {
            v14 := GenS2{F0: 4294967295, F1: 2147483648, F2: 0}
            switch v1 {
            case "a":
                fmt.Println(fi13, fi13)
            case "":
                switch v6 {
                case -9223372036854775808:
                    v6 = fi13
                case 9223372036854775807:
                    if false {
                        v4[fi13] = true
                        v1 = (v1 + ("" + v1))
                    }
                default:
                    v5 = (18446744073709551615 < 532)
                }
            default:
                fmt.Println(v1, fi13)
            }
            if v5 {
                v1 = "hello"
                v1 = "foo bar"
            } else {
                for fi15 := 0; fi15 < 4; fi15++ {
                    switch "foo bar" {
                    case "\n\t":
                        var v16 uint = 2
                        _ = v16
                    case "hello":
                        v17 := []int{501, (v6 >> ((18446744073709551615) & 63)), (fi15 | fi15), -1}
                        _ = v17
                    }
                    v6 += (-(fi13))
                    _ = fi15
                }
                v6 -= -1
            }
            _ = v14
        }
        _ = fi13
    }
    if (v5 || true) {
        var v18 int64 = 0
        _ = v18
    } else {
        v1 = v1
    }
    if false {
        var v19 float64 = (((3.14159 + 0.1) - 1.5) + 1e10)
        v19 = 0.0
        _ = v19
    } else {
        var v20 uint = 540
        v21 := []int{v6}
        _ = v20
        _ = v21
    }
    fmt.Println(v6, v6, v6)
    var v22 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
