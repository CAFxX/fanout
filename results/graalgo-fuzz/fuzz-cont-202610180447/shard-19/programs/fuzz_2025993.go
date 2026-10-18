package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int16 = -76
    v2 := make(map[string]uint32)
    var v3 uint = 42
    v4 := make(map[string]string)
    var v5 uint8 = 2
    var v6 int8 = 100
    v7 := GenS1{F0: 0}
    var v8 int8 = v6
    var v9 uint64 = 42
    if (false || (-9223372036854775808 >= 653)) {
        for fi10 := 0; fi10 < 13; fi10++ {
            if true {
                if true {
                    v11 := GenS1{F0: 255}
                    v12 := v4[("" + "\n\t")]
                    _ = v11
                    _ = v12
                }
                v1 = v1
                fmt.Println(fi10)
            }
            v13 := GenS1{F0: 42}
            _ = fi10
            _ = v13
        }
    }
    if (!true) {
        v14 := GenS2{F0: 1}
        for fi15 := 0; fi15 < 10; fi15++ {
            fmt.Println(v1)
            _ = fi15
        }
        _ = v14
    }
    fmt.Println(v8, v8, v8)
    if true {
        v16 := GenS2{F0: 0}
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := GenS2{F0: 42}
        var v19 uint64 = 139
        v20 := GenS2{F0: -1}
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    if (655 == 2147483648) {
        fmt.Println(v3)
        fmt.Println(v5, v5)
        v21 := []int{}
        _ = v21
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
