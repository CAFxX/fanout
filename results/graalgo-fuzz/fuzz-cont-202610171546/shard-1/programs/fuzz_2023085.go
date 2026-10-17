package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }

func main() {
    var v1 bool = (4294967295 == 4294967295)
    var v2 uint = 978
    var v3 uint64 = 469
    var v4 string = "\n\t"
    v5 := make(map[int]string)
    var v6 uint8 = 1
    var v7 uint8 = v6
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 int = fi8
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    fmt.Println(v4, v1, v3)
    v7 -= v7
    switch "\n\t" {
    case "":
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 uint16 = 2
            _ = fi11
            _ = v12
        }
    case "a":
        var v13 int32 = 0
        _ = v13
    default:
        v14 := []int{128, 32768, 908}
        _ = v14
    }
    v4 = "hello"
    v15 := GenS1{F0: -1, F1: (v7 | v7)}
    switch 32768 {
    case 32768:
        v4 = (v4 + v4)
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            if ((v3 <= v3) || v1) {
                fmt.Println(v6, fi16)
            }
            if ("foo bar" < "a") {
                v18 := GenS1{F0: -35, F1: (^(v7))}
                v1 = true
                var v19 uint = (^(v2))
                _ = v18
                _ = v19
            }
            v20 := GenS1{F0: -128, F1: v7}
            _ = fi17
            _ = v20
        }
        v7 += (v7 % ((128) | 1))
        _ = fi16
    }
    var v21 int8 = 50
    var v22 int = 128
    v23 := GenS1{F0: 100, F1: (^(v7))}
    var v24 int8 = -1
    var v25 bool = false
    var v26 uint8 = v7
    var v27 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
