package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 2147483647
    v3 := make(map[int]string)
    var v4 string = "foo bar"
    v5 := make(map[uint32]int)
    var v6 uint32 = 350
    fmt.Println(v4)
    if false {
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 5; fi8++ {
                v4 = (v4 + "\n\t")
                _ = fi8
            }
            var v9 int64 = -9223372036854775808
            var v10 float32 = -2.25
            _ = fi7
            _ = v10
            _ = v9
        }
        fmt.Println(v4, v6, v6)
    } else {
        switch 2147483647 {
        case 682:
            var v11 uint32 = 65535
            _ = v11
        }
        var v12 int32 = v2
        fmt.Println(v12, v12)
        _ = v12
    }
    var v13 int8 = -128
    v2 = v2
    for fi14 := 0; fi14 < 12; fi14++ {
        switch (fi14 & fi14) {
        case 4294967295:
            fmt.Println(v6)
        case 597:
            _ = len(v1)
        case -9223372036854775808:
            v2 += (v2 / ((-2147483648) | 1))
        }
        _ = fi14
    }
    var v15 uint16 = 42
    for fi16 := 0; fi16 < 1; fi16++ {
        fmt.Println(v6)
        _ = fi16
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 uint64 = 2
        v5[0] = fi17
        _ = fi17
        _ = v18
    }
    v13 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
