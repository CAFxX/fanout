package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int16; F1 string; F2 int32 }

func main() {
    v1 := []int{-2147483648, -646}
    var v2 uint16 = 65535
    v3 := []int{-9223372036854775808}
    var v4 float32 = (((float32(0.1) / 3.14159) * 0.0) - 1e10)
    v5 := []uint32{4294967295, 45}
    if false {
        for fi6 := 0; fi6 < 4; fi6++ {
            fmt.Println(fi6)
            _ = fi6
        }
    } else {
        if false {
            var v7 int16 = 32767
            switch "a" {
            case "":
                v2 = (v2 - 0)
            case "a":
                fmt.Println(v7)
            }
            var v8 uint64 = 18446744073709551615
            _ = v7
            _ = v8
        }
    }
    v9 := GenS1{F0: 309, F1: 32768}
    if (false || false) {
        _ = len(v1)
        _ = len(v1)
    } else {
        if (94 == -18) {
            fmt.Println(v2, v4)
        }
    }
    var v10 bool = (((!(!((!(!true)) && true))) || (18446744073709551615 > 83)) && false)
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v11 := GenS2{F0: 784, F1: ("foo bar" + ""), F2: -2147483648}
    v2 += 42
    if len(v5) > 0 {
        v5[0] = 694
    }
    v4 = v4
    _ = len(v5)
    var v12 int8 = 42
    v12 -= 100
    switch "a" {
    case "a":
        fmt.Println(v12)
    case "foo bar":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 float64 = (0.1 - ((3.14159 - 3.14159) / float64(0.1)))
        if (!(false && v10)) {
            var v15 uint32 = 2147483648
            if ((v4 <= (v4 + float32(1.5))) && ((("" + "") == "foo bar") && false)) {
                fmt.Println(v10, v10)
                v16 := GenS1{F0: 319, F1: fi13}
                _ = v16
            }
            v17 := GenS2{F0: -403, F1: "hello", F2: -1}
            _ = v15
            _ = v17
        } else {
            var v18 uint32 = 0
            _ = v18
        }
        var v19 int64 = 332
        _ = fi13
        _ = v14
        _ = v19
    }
    if len(v5) > 0 {
        v5[0] = 0
    }
    v20 := GenS1{F0: 582, F1: -1}
    v10 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
