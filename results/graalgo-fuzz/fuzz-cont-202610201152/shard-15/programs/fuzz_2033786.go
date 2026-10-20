package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint32 }

func main() {
    var v1 string = "a"
    var v2 uint16 = 26
    var v3 int8 = 127
    var v4 int8 = v3
    v5 := []int{}
    v1 = v1
    for fi6 := 0; fi6 < 1; fi6++ {
        _ = len(v5)
        _ = fi6
    }
    var v7 uint64 = 1
    fmt.Println(v3, v7, v3)
    _ = len(v5)
    if false {
        switch v1 {
        case "a":
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 int64 = -1
                if (2147483647 == 318) {
                    v10 := GenS1{F0: v9, F1: true, F2: 567}
                    var v11 int16 = 0
                    if true {
                        v12 := GenS1{F0: 9223372036854775807, F1: (!(0.1 <= 3.14159)), F2: 371}
                        fmt.Println(v2, v7, v9)
                        _ = v12
                    }
                    _ = v10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
        }
    }
    v7 += v7
    v3 += (-(v3))
    var v13 int64 = 9223372036854775807
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
