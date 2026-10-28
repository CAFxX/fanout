package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 float32 = (float32(float64((1.5 + 0.1))) - 3.14159)
    var v2 int8 = 123
    v3 := make(map[int]string)
    var v4 bool = true
    if (32767 < 128) {
        var v5 int32 = -2147483648
        v6 := []int{}
        v5 -= v5
        _ = v5
        _ = v6
    }
    var v7 float32 = v1
    for fi8 := 0; fi8 < 18; fi8++ {
        if (v4 && v4) {
            _ = len(v3)
            v9 := GenS1{F0: 1}
            if (v4 && false) {
                switch fi8 {
                case 118:
                    if (!(false || (!v4))) {
                        var v10 bool = (55 != 559)
                        _ = v10
                    }
                }
                v11 := []int{}
                _ = v11
            }
            _ = v9
        } else {
            v7 -= (float32(3.14159) - ((0.0 + 0.0) + v7))
            v4 = ((("hello" + ("hello" + "\n\t")) + "\n\t") != "\n\t")
            var v12 bool = (18446744073709551615 < 0)
            _ = v12
        }
        _ = fi8
    }
    var v13 int64 = 1
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := []int{}
        _ = fi14
        _ = v15
    }
    v13 = (^(v13))
    var v16 bool = (2 < 2147483647)
    fmt.Println(v16, v1, v4)
    var v17 float64 = 3.14159
    v18 := GenS1{F0: 18446744073709551615}
    fmt.Println(v17, v2, v1)
    var v19 int64 = v13
    v20 := []int{1}
    var v21 int8 = (v2 & v2)
    fmt.Println(v21)
    var v22 int = 1
    v13 = -9223372036854775808
    var v23 bool = (1 >= v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v7
}
