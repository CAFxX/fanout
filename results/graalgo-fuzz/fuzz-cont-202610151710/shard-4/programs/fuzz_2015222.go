package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint = 65535
    var v2 int16 = 1000
    var v3 int8 = 64
    var v4 uint32 = 2
    var v5 uint8 = 114
    if (uint32(v5) >= v4) {
        if false {
            v6 := GenS1{F0: 1000}
            _ = v6
        }
        if false {
            var v7 int = 923
            _ = v7
        } else {
            switch 1 {
            case 127:
                if false {
                    var v8 uint64 = uint64(v5)
                    var v9 bool = true
                    var v10 int64 = -44
                    _ = v10
                    _ = v8
                    _ = v9
                }
            case -2147483648:
                if true {
                    v4 += 0
                    for fi11 := 0; fi11 < 13; fi11++ {
                        fmt.Println(v3, v2)
                        _ = fi11
                    }
                }
            case 4294967295:
                fmt.Println(v2, v5)
            default:
                var v12 uint = 637
                _ = v12
            }
        }
        v13 := GenS1{F0: (^(v2))}
        _ = v13
    }
    var v14 int64 = 9223372036854775807
    var v15 string = "foo bar"
    var v16 uint16 = 1
    v16 = (v16 << (((^(v1))) & 63))
    var v17 bool = true
    fmt.Println(v2, v2)
    var v18 float32 = 1.5
    v19 := []int{-9223372036854775808, 127}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
