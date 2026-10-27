package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint8 }
func genFn0(p0 int64, p1 bool, p2 int32) int16 {
    v7 := GenS1{F0: 1, F1: p2}
    v8 := []int{2147483647, 128, 9223372036854775807, 181}
    _ = len(v8)
    v9 := GenS1{F0: 343, F1: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 bool = false
    var v2 float32 = (0.1 / (1e10 / 3.14159))
    var v3 uint8 = 127
    var v4 int16 = 1000
    v5 := []string{}
    var v6 int32 = 1
    v5 = append(v5, "a")
    if (false || v1) {
        var v10 int = int(v3)
        _ = v10
    } else {
        v5 = append(v5, ("\n\t" + ("" + "a")))
        v11 := GenS2{F0: (v4 ^ 1000), F1: 672}
        if (!v1) {
            var v12 int = -1
            fmt.Println(v12, v4)
            _ = v12
        } else {
            if ((v1 || (("" >= "") || (606 < 65535))) && (!true)) {
                v6 -= v6
            } else {
                for fi13 := 0; fi13 < 20; fi13++ {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        var v15 uint16 = 1
                        v16 := GenS3{F0: 914, F1: v3}
                        v4 -= -32768
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    _ = fi13
                }
                v6 = (-(v6))
            }
        }
        _ = v11
    }
    var v17 uint8 = (v3 - 2)
    v18 := GenS3{F0: 0, F1: 128}
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 int32 = -381
        _ = fi19
        _ = v20
    }
    var v21 int32 = -2147483648
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v22 uint64 = 42
    var v23 int8 = 1
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
