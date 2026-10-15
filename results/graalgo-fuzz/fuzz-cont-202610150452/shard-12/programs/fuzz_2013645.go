package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint8 = 1
    v3 := make(map[string]uint32)
    var v4 uint32 = 586
    var v5 uint = 255
    var v6 uint16 = 32767
    var v7 uint8 = (v2 * (v2 + (v2 + v2)))
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 string = "\n\t"
        if false {
            v10 := GenS1{F0: (v2 / (((^(v2))) | 1))}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 7; fi11++ {
                switch -185 {
                case -1:
                    var v12 bool = ((v4 >> ((2) & 63)) > (v4 | (v4 >> ((v5) & 63))))
                    _ = v12
                case 284:
                    v13 := GenS1{F0: (v2 | uint8(v4))}
                    _ = v13
                case 2147483647:
                    if (853 == v6) {
                        var v14 string = v9
                        var v15 float32 = 0.0
                        var v16 int = int(v2)
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        var v17 int64 = v1
                        _ = v17
                    }
                }
                v6 = (v6 + (^(v6)))
                var v18 uint8 = (v2 - (^(v2)))
                _ = fi11
                _ = v18
            }
            var v19 float32 = (float32(1.5) / -2.25)
            delete(v3, "\n\t")
            _ = v19
        }
        _ = fi8
        _ = v9
    }
    var v20 int64 = v1
    var v21 float32 = -2.25
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
