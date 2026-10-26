package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -9223372036854775808
    var v3 float64 = (-2.25 - 3.14159)
    var v4 int8 = 42
    var v5 uint16 = 32767
    var v6 string = "\n\t"
    var v7 uint8 = 2
    v8 := GenS1{F0: 9223372036854775807}
    v9 := GenS1{F0: int(v5)}
    if false {
        if (!(32767 < 9223372036854775807)) {
            v10 := GenS1{F0: 0}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 20; fi11++ {
                var v12 int = fi11
                var v13 int16 = -32768
                if ((992 < 1) || true) {
                    switch "" {
                    case "\n\t":
                        var v14 float32 = 0.0
                        _ = v14
                    case "a":
                        var v15 float64 = (v3 * 1.5)
                        _ = v15
                    case "":
                        var v16 int64 = v2
                        _ = v16
                    default:
                        fmt.Println(v3, v4, v6)
                    }
                    v17 := GenS1{F0: (v12 >> ((65535) & 63))}
                    _ = v17
                } else {
                    v18 := GenS1{F0: (v12 ^ 4294967295)}
                    var v19 uint64 = 0
                    _ = v18
                    _ = v19
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            var v20 int16 = 229
            _ = v20
        }
        v21 := GenS1{F0: -124}
        _ = v21
    }
    if (v7 == v7) {
        v22 := []int{9223372036854775807, 32768, 532, -1}
        fmt.Println(v5, v4)
        var v23 bool = ((((4294967295 <= 18446744073709551615) || false) || false) && ((v6 + "hello") >= v6))
        _ = v22
        _ = v23
    }
    v1[int(v2)] = 65535
    if true {
        fmt.Println(v7)
    }
    v1[254] = 2147483648
    var v24 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
