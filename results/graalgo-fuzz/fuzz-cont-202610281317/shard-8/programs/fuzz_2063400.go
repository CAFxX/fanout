package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 bool, p1 int8) uint {
    v7 := []int{}
    if p0 {
        v7 = append(v7, 0)
        var v8 int32 = -251
        _ = v8
    } else {
        p1 = p1
        v9 := GenS1{F0: "a", F1: 1, F2: 0}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 255
}
func genFn1(p0 int8, p1 uint16) int8 {
    fmt.Println(p1)
    for fi10 := 0; fi10 < 10; fi10++ {
        if false {
            p0 += p0
        } else {
            for fi11 := 0; fi11 < 7; fi11++ {
                v12 := []int{}
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 float64 = (((3.14159 + (-2.25 + -2.25)) / 1.0) / (1e10 * 3.14159))
                var v15 int = fi13
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        v16 := GenS1{F0: (("" + "a") + "hello"), F1: uint32(p1), F2: 42}
        _ = fi10
        _ = v16
    }
    p1 -= p1
    var v17 uint64 = 2
    _ = p0
    _ = p1
    _ = v17
    return p0
}

func main() {
    var v1 int8 = -17
    var v2 string = "a"
    var v3 uint = 893
    var v4 uint32 = 1
    v5 := make(map[string]string)
    var v6 int16 = 32767
    var v18 float32 = (0.0 + -2.25)
    v3 = 18446744073709551615
    var v19 string = v2
    if false {
        switch v19 {
        case "hello":
            if (42 > -9223372036854775808) {
                v20 := GenS2{F0: 0}
                if true {
                    var v21 int8 = v1
                    v19 = ""
                    var v22 uint32 = (v4 >> (((v3 % ((470) | 1))) & 63))
                    _ = v21
                    _ = v22
                } else {
                    v19 = "\n\t"
                    if ((v3 / ((genFn0(((828 >= -145) && true), v1)) | 1)) == 845) {
                        var v23 bool = false
                        var v24 bool = ((0.0 - ((float64(float32(1e10)) * float64(float32((0.0 * 0.1)))) / 3.14159)) > 1e10)
                        _ = v23
                        _ = v24
                    }
                }
                var v25 int16 = v6
                _ = v20
                _ = v25
            }
        default:
            var v26 int32 = 2147483647
            _ = v26
        }
        var v27 float32 = v18
        var v28 float32 = v18
        _ = v27
        _ = v28
    }
    v18 = (0.0 + 3.14159)
    v29 := GenS3{F0: v4}
    var v30 int32 = 42
    v31 := v5[v19]
    fmt.Println(v3, v30, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
