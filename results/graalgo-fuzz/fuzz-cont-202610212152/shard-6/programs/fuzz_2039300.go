package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint }
type GenS2 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS3 struct { F0 string }
func genFn0(p0 uint64, p1 uint8) int8 {
    v6 := []int{128}
    var v7 uint64 = 546
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := []uint32{}
    v2 := []string{("hello" + "\n\t")}
    var v3 uint32 = 65535
    v4 := make(map[string]string)
    var v5 float64 = (1.5 * 3.14159)
    v8 := GenS1{F0: 65535, F1: 18446744073709551615, F2: 646}
    v9 := GenS2{F0: 18446744073709551615, F1: 255, F2: (^(v3))}
    v10 := GenS2{F0: 0, F1: 128, F2: 42}
    fmt.Println(v5, v5)
    var v11 uint16 = 255
    var v12 int8 = 42
    var v13 float32 = -2.25
    if true {
        var v14 float32 = v13
        for fi15 := 0; fi15 < 3; fi15++ {
            if false {
                v16 := GenS3{F0: "a"}
                _ = v16
            } else {
                if true {
                    if true {
                        v13 += (-2.25 - 3.14159)
                        v14 += (float32(v5) - v13)
                        v11 = (v11 & 463)
                    }
                }
                fmt.Println(v13, v12, v3)
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 int8 = v12
                _ = fi17
                _ = v18
            }
            _ = fi15
        }
        v14 = 3.14159
        _ = v14
    } else {
        var v19 float32 = v13
        _ = v19
    }
    var v20 uint8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
