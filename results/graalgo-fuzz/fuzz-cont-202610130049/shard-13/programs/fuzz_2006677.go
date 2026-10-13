package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 0
    var v2 int32 = 811
    var v3 uint = 405
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    var v6 float32 = 1.5
    var v7 int16 = 583
    v3 -= (v3 ^ v3)
    var v8 int64 = 9223372036854775807
    var v9 bool = (true && true)
    if (!v9) {
        var v10 float32 = v6
        fmt.Println(v3, v9, v10)
        _ = v10
    } else {
        if false {
            var v11 int64 = v8
            v5[((("foo bar" + "foo bar") + ("foo bar" + "a")) + ("" + ""))] = (v3 == v3)
            if v9 {
                if (!true) {
                    v7 += (v7 * -1)
                }
                switch "foo bar" {
                case "\n\t":
                    var v12 uint = v3
                    _ = v12
                }
                v6 = v6
            }
            _ = v11
        }
    }
    if (v9 || v9) {
        v1 += v1
    }
    var v13 uint = v3
    v7 = v7
    v9 = v9
    v9 = false
    var v14 uint16 = 438
    for fi15 := 0; fi15 < 8; fi15++ {
        if ((v1 + 96) == v1) {
            fmt.Println(v3)
        } else {
            var v16 string = (((("\n\t" + "foo bar") + "foo bar") + "a") + ((((("foo bar" + ("foo bar" + (("\n\t" + "a") + ""))) + "a") + "foo bar") + "") + "foo bar"))
            _ = v16
        }
        switch "a" {
        case "\n\t":
            var v17 uint = 42
            _ = v17
        default:
            v14 -= 781
        }
        var v18 string = "a"
        _ = fi15
        _ = v18
    }
    var v19 uint16 = (v14 % ((v14) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
