package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 bool = false
    var v2 int8 = -122
    var v3 string = ""
    var v4 uint64 = 0
    for fi5 := 0; fi5 < 3; fi5++ {
        switch "hello" {
        case "\n\t":
            if (1 > 0) {
                var v6 float32 = 0.0
                v7 := GenS1{F0: 2147483647}
                _ = v6
                _ = v7
            }
        default:
            v8 := []int{(fi5 << ((42) & 63))}
            _ = v8
        }
        v2 += -14
        var v9 uint16 = 863
        _ = fi5
        _ = v9
    }
    v1 = (true && (655 == v4))
    if (v2 < 0) {
        if v1 {
            if (!v1) {
                var v10 string = v3
                v11 := []int{993, 4294967295, 2147483647}
                var v12 uint64 = 4294967295
                _ = v10
                _ = v11
                _ = v12
            } else {
                v13 := GenS2{F0: -1}
                _ = v13
            }
            fmt.Println(v3, v3)
        }
    }
    var v14 uint64 = v4
    v14 = (v4 * (^(v4)))
    fmt.Println(v4, v14, v4)
    v3 = ((v3 + (("" + "\n\t") + "foo bar")) + "\n\t")
    var v15 bool = (65535 > 757)
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(v4, fi16, fi16)
        _ = fi16
    }
    var v17 string = v3
    v18 := GenS2{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
