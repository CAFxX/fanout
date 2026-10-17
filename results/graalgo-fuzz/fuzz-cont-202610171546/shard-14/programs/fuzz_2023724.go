package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 string = ""
    var v4 int64 = 9223372036854775807
    var v5 int64 = v4
    var v6 uint16 = 255
    v5 = (v4 << ((672) & 63))
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 float64 = -2.25
        v9 := GenS1{F0: ((0 < 4294967295) && (1 == -73)), F1: true, F2: fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{42, 2147483647}
    if false {
        for fi11 := 0; fi11 < 17; fi11++ {
            fmt.Println(v3, fi11, v4)
            fmt.Println(v3)
            delete(v2, v3)
            _ = fi11
        }
    } else {
        fmt.Println(v3)
    }
    v12 := GenS1{F0: true, F1: false, F2: -186}
    switch v3 {
    case "hello":
        v5 += -9223372036854775808
    }
    v4 += (v5 << ((2) & 63))
    fmt.Println(v5, v3)
    switch 0 {
    case 42:
        v2[(v3 + (v3 + v3))] = 0
    case 2147483647:
        switch -833 {
        case 0:
            if false {
                v4 -= v5
                delete(v1, 0)
            }
        case 781:
            v5 -= v4
        case 4294967295:
            var v13 uint32 = 0
            _ = v13
        default:
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 uint16 = v6
                _ = fi14
                _ = v15
            }
        }
    default:
        _ = len(v1)
    }
    var v16 uint16 = v6
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v10 = append(v10, -303)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
