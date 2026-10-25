package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int8) uint8 {
    var v5 uint8 = 106
    v6 := GenS1{F0: 4294967295}
    var v7 uint16 = 1
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v5
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -2147483648
    var v3 bool = (32767 != 0)
    var v4 uint64 = 2
    var v8 bool = v3
    v9 := []int{-612, int(v2), -9223372036854775808, 9223372036854775807}
    v4 -= (v4 << ((18446744073709551615) & 63))
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            switch ("hello" + "a") {
            case "hello":
                v2 += (v2 % ((v2) | 1))
            default:
                var v12 int32 = 42
                v2 = v2
                _ = len(v1)
                _ = v12
            }
            if v3 {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            } else {
                v2 -= v2
                var v13 string = "foo bar"
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                if v8 {
                    var v14 uint16 = 1
                    _ = len(v1)
                    _ = v14
                }
                if (1e10 == (3.14159 / 0.1)) {
                    fmt.Println(v13, fi10)
                } else {
                    var v15 uint16 = 32767
                    if (394 <= 1) {
                        v15 -= (v15 % ((v15) | 1))
                    }
                    _ = v15
                }
                _ = v13
            }
            var v16 float64 = 0.0
            fmt.Println(v4)
            v3 = v3
            _ = fi11
            _ = v16
        }
        fmt.Println(fi10)
        var v17 int64 = 9223372036854775807
        _ = fi10
        _ = v17
    }
    v18 := GenS1{F0: 65535}
    var v19 float32 = 1.5
    v20 := GenS1{F0: 86}
    v19 -= -2.25
    fmt.Println(v2, v2, v3)
    var v21 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
