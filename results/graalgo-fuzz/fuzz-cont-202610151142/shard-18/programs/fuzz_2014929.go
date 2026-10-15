package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 float64 = -2.25
    var v2 uint = 2
    var v3 float32 = 1.5
    var v4 int32 = -2147483648
    var v5 float32 = ((0.1 / ((v3 + v3) - v3)) / 0.1)
    var v6 string = "hello"
    if (true && (((94 > -1) && false) && false)) {
        switch v6 {
        case "":
            var v7 int64 = -1
            _ = v7
        default:
            v8 := GenS1{F0: 1}
            _ = v8
        }
        if false {
            v9 := []int{4294967295, 42, 127, -895}
            v4 = (v4 * (v4 ^ (v4 * (v4 + (v4 * 2147483647)))))
            v1 = 3.14159
            _ = v9
        }
    } else {
        v10 := GenS1{F0: 42}
        v11 := []int{522, 9223372036854775807}
        _ = v10
        _ = v11
    }
    var v12 uint64 = 0
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 int32 = v4
            v16 := []int{}
            var v17 uint32 = 0
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        switch v6 {
        case "":
            for fi19 := 0; fi19 < 1; fi19++ {
                if true {
                    var v20 int8 = 0
                    _ = v20
                }
                _ = fi19
            }
        }
        _ = fi18
    }
    var v21 uint32 = 2147483647
    fmt.Println(v2, v2, v21)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
