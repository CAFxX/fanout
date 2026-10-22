package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int16; F1 int64; F2 int8 }

func main() {
    var v1 uint32 = 528
    v2 := make(map[uint32]uint32)
    var v3 uint = 2
    if false {
        var v4 int = 757
        _ = v4
    }
    if ((v3 > (v3 * v3)) && (52 <= -76)) {
        for fi5 := 0; fi5 < 12; fi5++ {
            v6 := []int{fi5}
            _ = fi5
            _ = v6
        }
        v7 := []int{486, -226}
        _ = v7
    }
    var v8 int = -159
    var v9 uint32 = 489
    var v10 float64 = 1.5
    delete(v2, 2)
    v1 -= 42
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := v2[v1]
            if (127 < 162) {
                v13 = 2147483648
                v14 := GenS1{F0: 0, F1: 255}
                v3 -= v3
                if (true || false) {
                    v15 := GenS2{F0: -9223372036854775808, F1: fi11}
                    _ = v15
                } else {
                    var v16 int = v8
                    if true {
                        if ((^(v3)) < (v3 - v3)) {
                            _ = len(v2)
                            v3 = 594
                            v17 := GenS1{F0: 42, F1: 128}
                            _ = v17
                        } else {
                            var v18 bool = (-1 != 1)
                            _ = v18
                        }
                        v19 := []int{0}
                        v19 = append(v19, (v16 << ((v3) & 63)))
                        _ = v19
                    } else {
                        fmt.Println(v10, v10)
                    }
                    var v20 int16 = 0
                    _ = v16
                    _ = v20
                }
                var v21 string = "hello"
                fmt.Println(v13, v9)
                _ = v14
                _ = v21
            } else {
                v22 := GenS3{F0: 32767, F1: 42, F2: 1}
                v13 = v9
                _ = v22
            }
            _ = fi12
            _ = v13
        }
        v23 := GenS3{F0: 1, F1: -564, F2: 127}
        _ = fi11
        _ = v23
    }
    var v24 int8 = -55
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v8
    _ = v9
}
