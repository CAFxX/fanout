package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint16 }

func main() {
    var v1 int8 = 42
    var v2 uint64 = 0
    var v3 uint8 = 2
    var v4 bool = ((-2.25 < 3.14159) && (!true))
    var v5 int32 = 42
    if (!(v4 || false)) {
        v6 := GenS1{F0: 0, F1: 32767}
        var v7 int32 = (v5 - (-(v5)))
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: -266, F1: 65535}
    for fi9 := 0; fi9 < 4; fi9++ {
        v2 = (v2 & (v2 * (v2 << ((2) & 63))))
        v10 := GenS1{F0: (-(fi9)), F1: 42}
        fmt.Println(fi9, fi9, fi9)
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: -741, F1: 725}
    switch 32767 {
    case 0:
        for fi12 := 0; fi12 < 15; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                v4 = ((!(!(float64(0.0) == (-2.25 - 1e10)))) && v4)
                _ = fi13
            }
            var v14 float32 = (((1e10 + 1e10) - 1.5) / ((1.5 / -2.25) - 0.1))
            var v15 uint64 = 896
            _ = fi12
            _ = v14
            _ = v15
        }
    case -2147483648:
        switch "a" {
        case "hello":
            v16 := GenS2{F0: v4, F1: 128, F2: 630}
            _ = v16
        }
    case 48:
        v4 = ((v1 != 127) && v4)
    }
    if v4 {
        switch "a" {
        case "a":
            var v17 float64 = -2.25
            _ = v17
        case "foo bar":
            v18 := []int{32767}
            _ = v18
        default:
            v2 = 42
        }
        var v19 string = "hello"
        if v4 {
            v1 -= int8(v5)
            if ((!((v1 << ((1) & 63)) > (v1 & v1))) && (v19 != v19)) {
                if v4 {
                    var v20 uint64 = (v2 ^ 948)
                    _ = v20
                }
                v21 := []int{128, 127, 42}
                var v22 bool = (2147483647 < 796)
                _ = v21
                _ = v22
            }
            v23 := []int{-200, -766, -932, 2147483647}
            _ = v23
        } else {
            var v24 string = v19
            _ = v24
        }
        _ = v19
    }
    if true {
        for fi25 := 0; fi25 < 12; fi25++ {
            v5 = (v5 << ((255) & 63))
            v26 := []int{}
            _ = fi25
            _ = v26
        }
        switch "a" {
        case "\n\t":
            switch 0 {
            case -745:
                v3 -= (v3 & 128)
            case 42:
                fmt.Println(v5, v1, v5)
            case -2147483648:
                v1 += (v1 | 100)
            default:
                var v27 int = 127
                _ = v27
            }
        case "a":
            for fi28 := 0; fi28 < 20; fi28++ {
                v2 = v2
                v3 += v3
                _ = fi28
            }
        case "hello":
            var v29 uint8 = v3
            _ = v29
        }
        switch "foo bar" {
        case "hello":
            if v4 {
                fmt.Println(v5, v2)
            } else {
                switch "hello" {
                case "":
                    var v30 bool = true
                    _ = v30
                case "foo bar":
                    v31 := []int{int(v5), -297, -557, -9223372036854775808}
                    _ = v31
                default:
                    var v32 uint = 0
                    _ = v32
                }
                v5 -= (v5 - (v5 / ((v5) | 1)))
                var v33 uint16 = 32767
                _ = v33
            }
        default:
            var v34 int8 = (v1 % ((v1) | 1))
            _ = v34
        }
    }
    fmt.Println(v1, v2, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
