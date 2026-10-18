package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
type GenS2 struct { F0 int32; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0() uint8 {
    var v5 string = ""
    var v6 uint64 = 2
    v7 := GenS1{F0: 127, F1: -1, F2: 349}
    v6 = 4294967295
    v8 := GenS1{F0: 100, F1: -32768, F2: 616}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 bool = (false || ((42 < 9223372036854775807) || true))
    var v2 int16 = 1000
    v3 := make(map[int]uint32)
    v4 := make(map[string]string)
    var v9 int8 = 77
    var v10 float32 = -2.25
    v1 = (1 >= 0)
    var v11 bool = v1
    v12 := GenS1{F0: 17, F1: 32767, F2: -2147483648}
    v13 := []int{-2147483648, 42, -1}
    switch -311 {
    case 398:
        fmt.Println(v1)
    case -2147483648:
        v10 -= 0.1
    case -373:
        _ = len(v13)
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        if true {
            var v15 uint = 4294967295
            _ = v15
        } else {
            if (4294967295 != 0) {
                var v16 int64 = 92
                v17 := GenS1{F0: v9, F1: v2, F2: (fi14 & 42)}
                _ = v16
                _ = v17
            } else {
                v4["hello"] = "foo bar"
                if (906 < 838) {
                    fmt.Println(v9, v10)
                } else {
                    if (2147483647 <= 2) {
                        v18 := GenS1{F0: int8(fi14), F1: (v2 + v2), F2: fi14}
                        v19 := GenS2{F0: -1, F1: 0, F2: 1}
                        _ = v18
                        _ = v19
                    }
                    if true {
                        fmt.Println(v11, v1, v1)
                        v10 = v10
                        for fi20 := 0; fi20 < 10; fi20++ {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            _ = fi20
                        }
                    }
                }
                v1 = (true || true)
            }
            v21 := GenS3{F0: 201}
            _ = v21
        }
        _ = fi14
    }
    v1 = false
    v13 = append(v13, 9223372036854775807)
    v1 = v11
    v22 := GenS3{F0: 1}
    var v23 int32 = 2147483647
    var v24 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
