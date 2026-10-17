package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }

func main() {
    var v1 uint32 = 2147483647
    var v2 uint8 = 42
    var v3 float32 = 1e10
    var v4 float32 = v3
    var v5 float32 = v3
    for fi6 := 0; fi6 < 20; fi6++ {
        switch fi6 {
        case 873:
            var v7 bool = false
            _ = v7
        case 1:
            var v8 uint32 = v1
            _ = v8
        case 127:
            if (-1 < 1) {
                if true {
                    v9 := GenS1{F0: "foo bar", F1: 255}
                    _ = v9
                }
            } else {
                var v10 int = (fi6 * 9223372036854775807)
                for fi11 := 0; fi11 < 13; fi11++ {
                    v2 += (^(v2))
                    v12 := []int{v10, (fi6 * (v10 & (fi11 * (fi11 - fi11)))), (^(fi6)), fi6}
                    var v13 int64 = 9223372036854775807
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = v10
            }
        }
        var v14 string = "a"
        _ = fi6
        _ = v14
    }
    v1 -= (v1 * v1)
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 uint32 = v1
        fmt.Println(v4)
        _ = fi15
        _ = v16
    }
    fmt.Println(v3, v3)
    var v17 uint = 42
    var v18 string = "hello"
    var v19 int64 = -1
    if true {
        var v20 int16 = -1
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
