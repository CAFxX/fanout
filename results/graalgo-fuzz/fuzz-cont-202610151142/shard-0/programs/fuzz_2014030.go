package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int8) uint16 {
    var v6 float32 = (3.14159 + float32((-2.25 + -2.25)))
    v6 = v6
    var v7 uint32 = 65535
    fmt.Println(v7)
    var v8 float64 = -2.25
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{4294967295, 65535, 2, 0}
    var v2 bool = true
    var v3 uint32 = 1
    v4 := make(map[string]uint32)
    var v5 int = -2147483648
    delete(v4, "a")
    v3 -= v3
    v9 := GenS1{F0: 100}
    v2 = (127 == 255)
    var v10 int8 = -115
    v5 += 42
    if (42 != 255) {
        var v11 float32 = float32(1.5)
        _ = v11
    } else {
        v12 := GenS1{F0: (v10 << ((18446744073709551615) & 63))}
        _ = v12
    }
    v2 = false
    v10 = (v10 / ((v10) | 1))
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS2{F0: (893 != 2), F1: 521}
        _ = fi13
        _ = v14
    }
    delete(v4, "hello")
    if (105 >= 184) {
        v10 = (v10 - v10)
    } else {
        switch "a" {
        case "hello":
            fmt.Println(v5, v2)
        case "":
            v5 += (v5 | (v5 % ((v5) | 1)))
        }
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        if (464 == -1) {
            fmt.Println(v5)
            v1 = append(v1, v3)
            v16 := GenS3{F0: v3}
            _ = v16
        } else {
            var v17 float64 = (1.5 - 0.0)
            if true {
                fmt.Println(v17, v17, v3)
                v18 := v4[("\n\t" + ("hello" + ""))]
                v17 = 3.14159
                _ = v18
            } else {
                fmt.Println(v10, v2)
                if (!v2) {
                    for fi19 := 0; fi19 < 13; fi19++ {
                        var v20 int32 = -576
                        _ = fi19
                        _ = v20
                    }
                    var v21 int32 = 84
                    _ = v21
                }
            }
            _ = v17
        }
        _ = fi15
    }
    var v22 bool = false
    if v22 {
        fmt.Println(v5, v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
