package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]bool)
    var v3 string = ""
    v4 := []int{32767, 1, -47, 361}
    var v5 string = v3
    if (!true) {
        var v6 int = -59
        _ = v6
    } else {
        fmt.Println(v5)
    }
    v1 = 151
    switch 32768 {
    case -1:
        v1 += 2
    default:
        var v7 int64 = 398
        _ = v7
    }
    var v8 float32 = ((-2.25 * (1.5 + 0.1)) / ((-2.25 + 1.5) / float32(float64(0.1))))
    v3 = v3
    switch "a" {
    case "\n\t":
        fmt.Println(v5, v1)
    default:
        if false {
            var v9 int64 = 42
            for fi10 := 0; fi10 < 13; fi10++ {
                if ((!(!(!true))) && true) {
                    var v11 int32 = -1
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 17; fi12++ {
                        var v13 int8 = -1
                        v5 = v3
                        _ = fi12
                        _ = v13
                    }
                }
                _ = fi10
            }
            for fi14 := 0; fi14 < 14; fi14++ {
                fmt.Println(fi14)
                _ = fi14
            }
            _ = v9
        } else {
            for fi15 := 0; fi15 < 10; fi15++ {
                if (false && false) {
                    v3 = "foo bar"
                    var v16 int64 = 9223372036854775807
                    var v17 int64 = (v16 - (v16 & (^(v16))))
                    _ = v16
                    _ = v17
                }
                _ = fi15
            }
        }
    }
    var v18 int32 = 2147483647
    var v19 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
