package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int32; F1 int64 }
func genFn0() int {
    if (((!((true || (-730 >= 42)) && true)) && (!(true && false))) && (0.0 != 1.5)) {
        var v6 uint16 = 756
        v7 := []int{42, 4294967295, 1, 127}
        _ = v6
        _ = v7
    }
    if false {
        v8 := []int{93, 2147483647}
        var v9 int16 = 0
        _ = v8
        _ = v9
    } else {
        var v10 float32 = -2.25
        for fi11 := 0; fi11 < 2; fi11++ {
            fmt.Println(fi11, fi11)
            _ = fi11
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            if (!(!false)) {
                var v13 uint8 = 128
                v10 -= (3.14159 / (v10 / v10))
                _ = v13
            }
            var v14 int = fi12
            _ = fi12
            _ = v14
        }
        _ = v10
    }
    var v15 bool = true
    fmt.Println(v15)
    v16 := GenS1{F0: 255}
    _ = v15
    _ = v16
    return -2147483648
}
func genFn1() int32 {
    var v17 uint = 2
    v18 := []int{-2147483648, int(v17), 2147483647, -1}
    if false {
        v17 = v17
        fmt.Println(v17)
    } else {
        v17 += v17
        v17 += (v17 << (((^(v17))) & 63))
    }
    _ = v17
    _ = v18
    return 2147483647
}

func main() {
    var v1 uint8 = 128
    var v2 int = 4294967295
    var v3 float32 = -2.25
    v4 := []int{}
    var v5 uint = 4294967295
    fmt.Println(v5, v2)
    fmt.Println(v3, v2, v2)
    fmt.Println(v5, v3)
    fmt.Println(v2, v3, v5)
    var v19 bool = false
    v4 = append(v4, (v2 * -882))
    var v20 int64 = 846
    v21 := GenS2{F0: v20, F1: ("hello" + "foo bar")}
    v22 := GenS3{F0: genFn1(), F1: 1}
    var v23 string = "foo bar"
    v5 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
