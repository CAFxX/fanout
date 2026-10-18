package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int32; F1 int32 }
func genFn0() int16 {
    var v4 int = 127
    v5 := GenS1{F0: 255, F1: 65535}
    v4 = 0
    fmt.Println(v4)
    v6 := []int{4294967295, v4}
    _ = v4
    _ = v5
    _ = v6
    return 1000
}

func main() {
    var v1 uint8 = 0
    var v2 uint16 = 255
    var v3 bool = false
    var v7 string = ("a" + "hello")
    v8 := GenS1{F0: 1, F1: 773}
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := []int{-2147483648, fi9}
        for fi11 := 0; fi11 < 6; fi11++ {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            switch "a" {
            case "a":
                fmt.Println(v3, v2)
            case "hello":
                switch v7 {
                case "a":
                    fmt.Println(v2, v2, fi11)
                case "foo bar":
                    v10 = append(v10, fi11)
                    var v12 float64 = 0.0
                    _ = v12
                }
            }
            _ = fi11
        }
        fmt.Println(v3)
        _ = fi9
        _ = v10
    }
    var v13 int = 9223372036854775807
    if v3 {
        fmt.Println(v1, v2, v3)
        v14 := GenS1{F0: v2, F1: 2}
        _ = v14
    } else {
        var v15 uint8 = 127
        _ = v15
    }
    v16 := GenS2{F0: -1, F1: -2147483648}
    var v17 int8 = 42
    var v18 int8 = v17
    if v3 {
        switch ("foo bar" + (v7 + v7)) {
        case "foo bar":
            v3 = (!((v2 << ((uint(v1)) & 63)) == v2))
        default:
            var v19 uint32 = 2
            _ = v19
        }
    }
    switch 519 {
    case 4294967295:
        for fi20 := 0; fi20 < 6; fi20++ {
            v21 := []int{(v13 + (fi20 ^ 32768)), v13, 42, (-(fi20))}
            var v22 int = fi20
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    var v23 uint16 = 32767
    var v24 string = ((v7 + v7) + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v7
    _ = v8
}
