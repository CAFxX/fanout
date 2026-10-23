package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int16 = -850
    var v2 uint8 = 1
    var v3 int64 = 42
    v4 := make(map[int]bool)
    var v5 int64 = 1
    v3 = int64(v1)
    var v6 uint32 = 741
    fmt.Println(v2)
    v7 := []int{-2147483648, 148, 667}
    if true {
        v8 := GenS1{F0: 94}
        _ = v8
    }
    switch "foo bar" {
    case "\n\t":
        var v9 int = 844
        _ = v9
    case "hello":
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    switch 0 {
    case -398:
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int32 = -2147483648
            var v12 uint = 18446744073709551615
            _ = fi10
            _ = v11
            _ = v12
        }
    case 42:
        fmt.Println(v5, v2, v2)
    case 507:
        if true {
            for fi13 := 0; fi13 < 12; fi13++ {
                for fi14 := 0; fi14 < 9; fi14++ {
                    switch "a" {
                    case "":
                        if false {
                            v1 = v1
                            v5 = (v3 - -9223372036854775808)
                            var v15 string = ""
                            _ = v15
                        }
                    case "\n\t":
                        var v16 float32 = float32(float64(3.14159))
                        _ = v16
                    default:
                        var v17 uint = 42
                        _ = v17
                    }
                    _ = fi14
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                if false {
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v5 -= v5
                        v1 = -523
                        switch (fi13 ^ fi13) {
                        case 918:
                            delete(v4, (-(fi13)))
                        case -9223372036854775808:
                            var v19 string = "hello"
                            _ = v19
                        }
                        _ = fi18
                    }
                    v3 = 0
                    v20 := GenS1{F0: 1}
                    _ = v20
                }
                _ = fi13
            }
            if len(v7) > 0 {
                v7[0] = 728
            }
            fmt.Println(v5)
        }
    }
    var v21 float32 = 3.14159
    fmt.Println(v2, v2)
    var v22 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
