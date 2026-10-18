package main

import (
    "fmt"
)


func main() {
    var v1 string = ("a" + "hello")
    var v2 uint16 = 32767
    var v3 uint64 = 241
    var v4 string = (v1 + ((v1 + "hello") + "\n\t"))
    var v5 uint8 = 127
    v6 := make(map[uint32]bool)
    switch "a" {
    case "foo bar":
        for fi7 := 0; fi7 < 12; fi7++ {
            _ = len(v6)
            _ = fi7
        }
    case "":
        var v8 uint64 = v3
        _ = v8
    default:
        if false {
            fmt.Println(v1)
            v2 = (v2 % ((v2) | 1))
        }
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        switch (-(fi9)) {
        case -665:
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 uint16 = 643
                var v12 uint8 = (v5 % ((v5) | 1))
                if true {
                    var v13 float64 = (((((0.0 - 1e10) / (1e10 * ((3.14159 * 0.1) + (0.0 / 1.5)))) - (1e10 * -2.25)) + (1e10 * 0.0)) * -2.25)
                    v14 := []int{128, fi9}
                    _ = v13
                    _ = v14
                } else {
                    var v15 int8 = -9
                    _ = v15
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        case 1:
            v5 += 165
        case 32768:
            switch 1 {
            case 127:
                var v16 int16 = 188
                _ = v16
            }
        }
        _ = len(v6)
        _ = fi9
    }
    var v17 int32 = 0
    var v18 float32 = 1.5
    var v19 uint64 = 4294967295
    var v20 int = 32768
    v20 -= 744
    if false {
        var v21 int64 = 0
        _ = v21
    }
    delete(v6, 42)
    v5 += (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
