package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 int = 9223372036854775807
    var v2 int8 = 1
    var v3 string = "\n\t"
    var v4 uint64 = 2
    v5 := make(map[uint32]string)
    var v6 string = (v3 + "")
    var v7 uint32 = 1
    fmt.Println(v2, v3, v6)
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int64 = 177
        _ = fi8
        _ = v9
    }
    var v10 int32 = 279
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 2; fi12++ {
            if (1 < 525) {
                var v13 float32 = 3.14159
                if (!false) {
                    if true {
                        switch -788 {
                        case -9223372036854775808:
                            v14 := []int{v1, -9223372036854775808, v1}
                            _ = v14
                        }
                        v3 = v6
                    }
                    v15 := []int{fi11, (v1 | fi12), (fi12 ^ (fi12 + (fi11 - v1))), fi12}
                    _ = v15
                } else {
                    v5[(v7 & 231)] = ("a" + ("\n\t" + ""))
                }
                fmt.Println(v6)
                fmt.Println(v7)
                _ = v13
            }
            var v16 int8 = v2
            v16 = v16
            _ = fi12
            _ = v16
        }
        _ = fi11
    }
    v17 := GenS1{F0: 255, F1: v7}
    if true {
        var v18 float64 = 3.14159
        v19 := GenS1{F0: 529, F1: (^(v7))}
        _ = v18
        _ = v19
    } else {
        fmt.Println(v4, v1, v1)
    }
    var v20 int = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
