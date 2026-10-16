package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int = 42
    var v2 int = v1
    var v3 int32 = 42
    var v4 string = "\n\t"
    v1 -= v1
    fmt.Println(v3, v4)
    switch "" {
    case "":
        if ("" == (v4 + ("hello" + ((v4 + v4) + v4)))) {
            v5 := GenS1{F0: false}
            var v6 int16 = 32767
            v7 := []int{(-(v2))}
            _ = v5
            _ = v6
            _ = v7
        } else {
            v4 = (v4 + (v4 + v4))
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 int16 = 490
                fmt.Println(v4, v9, v9)
                var v10 string = v4
                _ = fi8
                _ = v10
                _ = v9
            }
        }
    case "hello":
        var v11 int16 = 0
        _ = v11
    case "a":
        if false {
            if true {
                switch v2 {
                case 952:
                    fmt.Println(v2, v1)
                case 2147483647:
                    var v12 uint8 = 0
                    _ = v12
                }
            }
            switch -606 {
            case -926:
                v3 = v3
            default:
                if false {
                    if false {
                        v1 += v2
                        fmt.Println(v2, v1)
                        if false {
                            fmt.Println(v1)
                        } else {
                            var v13 uint16 = 42
                            var v14 uint16 = v13
                            _ = v13
                            _ = v14
                        }
                    }
                }
            }
        }
    default:
        v2 = (v1 & 4294967295)
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        if (true || (1000 <= 1)) {
            v16 := []int{(v2 * v2), (v2 << ((255) & 63)), (^(fi15)), (v2 + -9)}
            _ = v16
        } else {
            v17 := []int{(v1 - (v2 / ((fi15) | 1))), (v1 % ((v2) | 1)), fi15, 127}
            var v18 int32 = (v3 & v3)
            v1 += v2
            _ = v17
            _ = v18
        }
        var v19 uint64 = 2
        v20 := GenS1{F0: false}
        _ = fi15
        _ = v19
        _ = v20
    }
    v2 = v1
    var v21 int16 = -32768
    v4 = ((v4 + v4) + "foo bar")
    v2 = v1
    var v22 int64 = -395
    var v23 int8 = 127
    v24 := GenS1{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
