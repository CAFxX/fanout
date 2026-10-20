package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int32 = -1
    var v2 bool = false
    var v3 int8 = -1
    var v4 int = -1
    var v5 int = v4
    var v6 uint32 = 4294967295
    var v7 uint16 = 42
    fmt.Println(v1)
    v8 := GenS1{F0: v3}
    fmt.Println(v4, v2, v4)
    for fi9 := 0; fi9 < 17; fi9++ {
        v6 = 42
        switch -2147483648 {
        case 1:
            for fi10 := 0; fi10 < 6; fi10++ {
                if ((v1 * v1) < v1) {
                    switch 32768 {
                    case 4294967295:
                        var v11 uint32 = (v6 << ((65535) & 63))
                        _ = v11
                    }
                    switch ("\n\t" + "") {
                    case "\n\t":
                        v1 += v1
                    default:
                        var v12 uint = 18446744073709551615
                        _ = v12
                    }
                }
                var v13 bool = v2
                v1 -= 0
                var v14 int16 = 32767
                switch v5 {
                case 9223372036854775807:
                    v15 := GenS1{F0: 107}
                    _ = v15
                default:
                    var v16 uint = 255
                    _ = v16
                }
                _ = fi10
                _ = v13
                _ = v14
            }
        case -1:
            v6 += (v6 ^ 832)
        }
        _ = fi9
    }
    if (!(255 == 166)) {
        v17 := GenS1{F0: (v3 + v3)}
        _ = v17
    }
    var v18 float32 = float32(0.0)
    var v19 uint32 = (v6 / (((v6 >> ((65535) & 63))) | 1))
    var v20 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
