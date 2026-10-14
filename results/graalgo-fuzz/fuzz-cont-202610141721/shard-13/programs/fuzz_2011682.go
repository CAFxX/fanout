package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int }

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]uint32)
    var v3 int16 = 32767
    v4 := make(map[string]bool)
    var v5 uint = (v1 + 4294967295)
    var v6 uint8 = 255
    for fi7 := 0; fi7 < 7; fi7++ {
        if true {
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := []int{1, (fi7 << ((255) & 63)), (fi8 ^ (fi7 ^ fi8))}
                var v10 uint32 = 2147483647
                v3 = -1
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        _ = fi7
    }
    switch "hello" {
    case "":
        fmt.Println(v3)
    default:
        v11 := GenS1{F0: -537, F1: (844 == 9223372036854775807), F2: 736}
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 float64 = 1.5
        v14 := []int{}
        v15 := GenS1{F0: 42, F1: false, F2: 4294967295}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    fmt.Println(v6, v1)
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 int8 = 102
        for fi18 := 0; fi18 < 7; fi18++ {
            if false {
                var v19 int32 = -2147483648
                var v20 int = (fi16 & -1)
                var v21 int16 = 1
                if (true && true) {
                    v4[("" + ("a" + ""))] = (-69 >= v21)
                }
                var v22 uint8 = (v6 * v6)
                _ = v19
                _ = v20
                _ = v21
                _ = v22
            }
            v6 += (v6 >> (((v1 * v1)) & 63))
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    v6 -= (v6 ^ (v6 << ((v5) & 63)))
    var v23 string = ""
    if (true && false) {
        fmt.Println(v6)
        v6 += v6
    }
    v4[v23] = true
    var v24 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
