package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int32 }
type GenS2 struct { F0 bool; F1 int32; F2 uint }

func main() {
    var v1 uint8 = 127
    var v2 uint16 = 255
    var v3 string = "hello"
    var v4 uint64 = 42
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            switch -2147483648 {
            case 4294967295:
                switch int(v2) {
                case -223:
                    if false {
                        v3 = "foo bar"
                        v7 := GenS1{F0: (v2 ^ 32767), F1: 4294967295, F2: -1}
                        _ = v7
                    }
                default:
                    v3 = ((v3 + ("foo bar" + (v3 + v3))) + v3)
                }
                v8 := GenS1{F0: (v2 & (^(v2))), F1: 0, F2: 0}
                _ = v8
            }
            var v9 uint64 = v4
            v2 -= (v2 ^ v2)
            _ = fi6
            _ = v9
        }
        _ = fi5
    }
    v10 := []int{-2147483648, -561, -9223372036854775808}
    var v11 int64 = 216
    if (1e10 != float64(0.0)) {
        v12 := GenS2{F0: (-389 != -578), F1: 951, F2: 255}
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 int64 = (-(v11))
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    v3 = ""
    v4 -= 18446744073709551615
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            v4 = v4
            var v17 string = v3
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v18 := GenS1{F0: 42, F1: 4294967295, F2: 370}
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
