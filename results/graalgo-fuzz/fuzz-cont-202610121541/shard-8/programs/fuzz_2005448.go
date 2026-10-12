package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []string{"a", ("\n\t" + ("" + ("\n\t" + "\n\t")))}
    var v2 uint32 = 525
    var v3 int16 = 1000
    var v4 uint64 = 18446744073709551615
    var v5 float32 = ((1.5 / 1e10) - (float32((0.0 + ((float64(3.14159) + (3.14159 / (3.14159 * 1.5))) - float64(3.14159)))) * (1e10 - (-2.25 / 1.0))))
    var v6 int16 = v3
    var v7 uint64 = (^(v4))
    var v8 int32 = 1
    if ((-307 != -1) || false) {
        v9 := GenS1{F0: 228}
        _ = v9
    } else {
        v10 := GenS1{F0: 1}
        if (false || true) {
            v3 -= -32768
        } else {
            var v11 int64 = 9223372036854775807
            v12 := GenS1{F0: 91}
            _ = v11
            _ = v12
        }
        _ = v10
    }
    switch -244 {
    case 350:
        _ = len(v1)
    case 42:
        v4 = v7
    default:
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 float64 = (float64((0.0 * -2.25)) + (1.5 + 3.14159))
            v15 := GenS2{F0: 1}
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 uint8 = 255
                v14 -= (v14 - 1e10)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi16
                _ = v17
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    var v18 int32 = v8
    fmt.Println(v7, v3, v5)
    var v19 int8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
