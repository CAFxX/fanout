package main

import (
    "fmt"
)


func main() {
    var v1 int = 32768
    var v2 uint16 = 0
    var v3 int16 = 42
    var v4 float32 = (-2.25 + -2.25)
    var v5 uint = 0
    var v6 float64 = (-2.25 * 1e10)
    fmt.Println(v5)
    v5 += v5
    fmt.Println(v4, v3, v2)
    v6 = 1e10
    v2 += (^(v2))
    switch v1 {
    case 127:
        var v7 int32 = -308
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(v6, fi8)
        v9 := []int{v1, fi8}
        _ = fi8
        _ = v9
    }
    if false {
        v1 = 42
    } else {
        var v10 uint16 = 0
        _ = v10
    }
    v6 += (v6 - v6)
    var v11 float32 = -2.25
    if (v1 == 833) {
        fmt.Println(v11)
        v12 := []int{128, v1}
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 int = (fi13 | (^(v1)))
            _ = fi13
            _ = v14
        }
        _ = v12
    } else {
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 float32 = 0.1
            var v17 bool = (false && ((false || (100 <= -21)) && false))
            fmt.Println(v17)
            _ = fi15
            _ = v16
            _ = v17
        }
        switch "a" {
        case "\n\t":
            v18 := []int{(v1 % (((v1 | (v1 % (((v1 + v1)) | 1)))) | 1))}
            _ = v18
        }
        v1 -= v1
    }
    v3 = v3
    var v19 float32 = v4
    switch ("" + "a") {
    case "a":
        if false {
            var v20 string = ""
            _ = v20
        }
    case "\n\t":
        var v21 string = "hello"
        _ = v21
    default:
        var v22 int32 = -523
        _ = v22
    }
    fmt.Println(v5, v1)
    var v23 int16 = int16(v2)
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
