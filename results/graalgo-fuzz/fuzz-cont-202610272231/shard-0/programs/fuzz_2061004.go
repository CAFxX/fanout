package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint16 = 887
    var v2 int32 = 2147483647
    v3 := []uint32{798, 1, 4294967295}
    var v4 uint = 886
    var v5 bool = false
    var v6 float32 = (float32(1e10) - 0.1)
    if v5 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v1 = 32767
    for fi7 := 0; fi7 < 16; fi7++ {
        if (0.0 == float32(1.5)) {
            v2 = (v2 & v2)
            var v8 string = ("a" + "a")
            _ = v8
        } else {
            v9 := GenS1{F0: -465}
            if true {
                v5 = true
                v2 += v2
                var v10 uint8 = 181
                _ = v10
            }
            _ = v9
        }
        v1 += v1
        fmt.Println(v2, fi7, v5)
        _ = fi7
    }
    var v11 float64 = -2.25
    v11 = (0.1 + ((v11 * v11) / 1e10))
    if v5 {
        if ((v2 ^ (^(v2))) == (v2 & 1)) {
            var v12 uint8 = 42
            var v13 int32 = v2
            _ = v12
            _ = v13
        }
    }
    if v5 {
        v5 = v5
        for fi14 := 0; fi14 < 6; fi14++ {
            switch fi14 {
            case -930:
                for fi15 := 0; fi15 < 9; fi15++ {
                    v6 = (3.14159 * v6)
                    fmt.Println(v4, v1)
                    v2 = v2
                    _ = fi15
                }
            case -888:
                v16 := GenS1{F0: 1}
                _ = v16
            case 42:
                v6 = v6
            default:
                v17 := GenS1{F0: 1}
                _ = v17
            }
            _ = fi14
        }
        var v18 uint64 = 4294967295
        _ = v18
    } else {
        var v19 int64 = 141
        v3 = append(v3, 65535)
        _ = v19
    }
    v11 = 0.1
    var v20 float32 = float32(v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
