package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int16; F1 int64 }

func main() {
    var v1 string = ("a" + ("foo bar" + ("" + "hello")))
    var v2 float64 = (3.14159 + float64(float32(0.0)))
    v3 := make(map[int]bool)
    var v4 uint = 170
    v5 := []int{127, 42, 2147483647}
    v2 = (v2 + float64(3.14159))
    v6 := GenS1{F0: 2147483647}
    v5 = append(v5, 127)
    v7 := GenS1{F0: 2147483647}
    fmt.Println(v1, v4)
    v8 := v3[-9223372036854775808]
    v4 += (v4 * 884)
    if ((255 <= 42) || (127 != -1)) {
        if false {
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 uint16 = uint16(v4)
                var v11 string = ("foo bar" + "foo bar")
                _ = fi9
                _ = v10
                _ = v11
            }
            if v8 {
                v2 = v2
            } else {
                var v12 float32 = float32(-2.25)
                v13 := GenS2{F0: 32767, F1: 164, F2: -343}
                if v8 {
                    var v14 int64 = 147
                    var v15 uint8 = 255
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v1 = v1
                        _ = fi16
                    }
                    _ = v14
                    _ = v15
                } else {
                    var v17 string = "\n\t"
                    v18 := v3[-2147483648]
                    v8 = v18
                    _ = v17
                    _ = v18
                }
                _ = v12
                _ = v13
            }
            var v19 uint8 = 127
            _ = v19
        }
        fmt.Println(v2, v1, v4)
    }
    v20 := GenS3{F0: 42}
    fmt.Println(v4, v4, v4)
    v21 := GenS4{F0: 42, F1: uint8(v4)}
    v22 := GenS5{F0: -444, F1: 9223372036854775807}
    fmt.Println(v2, v4, v2)
    v23 := GenS1{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
