package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
func genFn0() uint8 {
    var v7 int16 = -1
    v7 = v7
    if false {
        v7 = 353
    }
    v7 += 1000
    if ((-923 < -1) && true) {
        v8 := []int{-1, 32767, 654, 1}
        if (42 >= 740) {
            if (int8(v7) != 42) {
                var v9 int64 = 1
                _ = v9
            }
            var v10 int = 0
            _ = v10
        }
        var v11 uint16 = 32767
        _ = v11
        _ = v8
    } else {
        var v12 int8 = 100
        var v13 uint8 = 0
        _ = v12
        _ = v13
    }
    _ = v7
    return 1
}

func main() {
    var v1 int32 = -1
    var v2 int16 = -1
    v3 := []uint32{4294967295, 42, 552}
    var v4 int32 = v1
    v5 := make(map[int]uint32)
    var v6 string = (("a" + ("foo bar" + "hello")) + "")
    var v14 uint8 = 129
    v1 = (^(v4))
    fmt.Println(v14, v1, v14)
    if ((1.5 + 0.0) <= (3.14159 + -2.25)) {
        v15 := GenS1{F0: 42, F1: 605}
        v3 = append(v3, 4294967295)
        _ = v15
    }
    var v16 int = -456
    for fi17 := 0; fi17 < 3; fi17++ {
        if true {
            fmt.Println(v14, v14, v14)
            fmt.Println(v1)
            var v18 int32 = int32(v14)
            _ = v18
        } else {
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 float64 = float64(-2.25)
                _ = fi19
                _ = v20
            }
            _ = len(v5)
            v2 = -1
        }
        _ = fi17
    }
    var v21 string = v6
    v4 = -2147483648
    for fi22 := 0; fi22 < 9; fi22++ {
        v14 -= 70
        _ = fi22
    }
    var v23 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
