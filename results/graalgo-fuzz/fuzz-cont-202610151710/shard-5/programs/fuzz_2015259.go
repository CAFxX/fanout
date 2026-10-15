package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint }

func main() {
    var v1 uint64 = 42
    v2 := []int{42, 127, 127, 1}
    var v3 int32 = int32(v1)
    var v4 float64 = 0.1
    var v5 uint32 = 908
    var v6 uint16 = 2
    switch -199 {
    case -9223372036854775808:
        var v7 bool = true
        _ = v7
    case -600:
        switch int(v3) {
        case 525:
            for fi8 := 0; fi8 < 8; fi8++ {
                if len(v2) > 0 {
                    v2[0] = 1
                }
                var v9 uint64 = (v1 | (v1 % ((57) | 1)))
                _ = fi8
                _ = v9
            }
        }
    case -1:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v1 += (v1 + (v1 + v1))
    var v10 uint64 = 693
    switch ("a" + "foo bar") {
    case "a":
        for fi11 := 0; fi11 < 8; fi11++ {
            if false {
                var v12 string = ("a" + "a")
                v13 := GenS1{F0: false, F1: 2147483648, F2: 447}
                var v14 int16 = 32767
                _ = v12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 7; fi15++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi15
            }
            _ = len(v2)
            _ = fi11
        }
    case "hello":
        if len(v2) > 0 {
            v2[0] = 397
        }
    case "foo bar":
        v10 = (v1 >> ((uint(v6)) & 63))
    }
    fmt.Println(v5)
    var v16 uint = 985
    _ = len(v2)
    var v17 float32 = ((3.14159 / (float32(3.14159) + 0.0)) - 1e10)
    v5 = v5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        fmt.Println(v5, v6)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v18 := GenS1{F0: false, F1: v5, F2: 475}
        _ = v18
    } else {
        var v19 int8 = int8(v1)
        v20 := GenS1{F0: true, F1: 2, F2: 83}
        v17 -= 1e10
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
