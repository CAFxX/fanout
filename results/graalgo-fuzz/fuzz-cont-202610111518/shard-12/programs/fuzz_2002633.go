package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint32 = 117
    var v2 uint64 = 0
    v3 := []uint32{}
    var v4 int16 = 0
    var v5 uint = 4294967295
    var v6 uint16 = 2
    if len(v3) > 0 {
        v3[0] = 24
    }
    if (!("a" > "")) {
        var v7 string = ""
        _ = v7
    } else {
        if false {
            if false {
                v4 = v4
            } else {
                switch 127 {
                case 2147483647:
                    fmt.Println(v4, v6)
                case 128:
                    v8 := GenS1{F0: v1}
                    _ = v8
                case 9223372036854775807:
                    var v9 uint8 = uint8(v5)
                    _ = v9
                }
                var v10 string = "\n\t"
                var v11 int = 32768
                _ = v10
                _ = v11
            }
            if len(v3) > 0 {
                v3[0] = (v1 * 2147483647)
            }
        } else {
            switch 127 {
            case -691:
                var v12 int64 = 0
                _ = v12
            }
            var v13 bool = true
            _ = v13
        }
        v14 := GenS1{F0: (^(v1))}
        _ = v14
    }
    var v15 int64 = 42
    v6 += (v6 << (((v5 ^ (v5 | (v5 / (((v5 | v5)) | 1))))) & 63))
    var v16 int16 = 313
    var v17 bool = true
    var v18 uint64 = v2
    v18 = 282
    var v19 uint32 = 722
    v20 := GenS1{F0: 4294967295}
    var v21 float32 = -2.25
    var v22 float32 = float32((1e10 / (1e10 - (0.1 / 0.1))))
    var v23 uint32 = v19
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v24 := GenS1{F0: (^(v23))}
    fmt.Println(v18)
    fmt.Println(v17, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
