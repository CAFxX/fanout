package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int = -479
    var v2 uint64 = 4294967295
    var v3 int64 = 1
    var v4 float64 = 1e10
    v5 := []int{(-(v1))}
    var v6 int32 = -726
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(v4)
        _ = fi8
    }
    var v9 int32 = v6
    v1 = (v1 << ((2) & 63))
    switch ("hello" + "a") {
    case "a":
        var v10 int8 = 25
        _ = v10
    case "foo bar":
        v5 = append(v5, (^(v1)))
    default:
        if true {
            var v11 uint64 = v2
            for fi12 := 0; fi12 < 15; fi12++ {
                v2 += (v2 | (v11 | 673))
                var v13 int8 = 100
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 int8 = (v13 & v13)
                    _ = fi14
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
            _ = v11
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    switch v1 {
    case 42:
        v16 := GenS1{F0: 128}
        _ = v16
    case -775:
        var v17 uint64 = (v2 + v2)
        _ = v17
    case 32768:
        v4 += 1.5
    }
    fmt.Println(v4, v4, v6)
    v7 = -513
    var v18 int32 = (v9 + v9)
    var v19 float64 = -2.25
    v3 += -9223372036854775808
    v20 := GenS1{F0: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
