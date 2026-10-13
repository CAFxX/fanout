package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 string) uint8 {
    v6 := GenS1{F0: 2147483647}
    fmt.Println(p0)
    switch p0 {
    case "a":
        fmt.Println(p0)
    default:
        v7 := []int{}
        _ = v7
    }
    var v8 uint8 = 1
    _ = p0
    _ = v6
    _ = v8
    return (v8 % ((2) | 1))
}
func genFn1() uint16 {
    var v9 int64 = 9223372036854775807
    var v10 string = "foo bar"
    v11 := []int{}
    _ = v10
    _ = v11
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    var v3 float64 = 3.14159
    var v4 float64 = v3
    var v5 int64 = -9223372036854775808
    v12 := GenS2{F0: int8(v5)}
    v5 = v5
    for fi13 := 0; fi13 < 6; fi13++ {
        if false {
            v4 = (v3 + (-2.25 / 3.14159))
            v4 += v3
        }
        _ = fi13
    }
    v5 = 9223372036854775807
    var v14 int16 = -541
    if false {
        fmt.Println(v5, v5)
        v1[-100] = "\n\t"
        var v15 int8 = -128
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        fmt.Println(v14)
        v17 := v1[(fi16 >> ((42) & 63))]
        var v18 uint8 = 0
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint32 = 800
    if false {
        v4 = 1e10
    }
    v3 -= v3
    v4 += ((float64(-2.25) - v3) / v4)
    if (true && true) {
        if (v19 > v19) {
            var v20 bool = true
            _ = v20
        } else {
            v21 := GenS1{F0: 233}
            var v22 float32 = -2.25
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 uint32 = (v19 >> ((uint(v19)) & 63))
            v25 := []int{(fi23 / ((fi23) | 1)), 76, fi23, fi23}
            v14 -= int16(fi23)
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 string = ("\n\t" + ("hello" + (("hello" + "a") + "hello")))
        _ = v26
    } else {
        v27 := []int{128}
        _ = v27
    }
    fmt.Println(v3)
    if ((0 == 0) || (!(0 == -1))) {
        for fi28 := 0; fi28 < 15; fi28++ {
            v14 = -1
            if true {
                v1[fi28] = ("a" + "a")
                if (!true) {
                    if true {
                        var v29 int16 = (v14 >> ((4294967295) & 63))
                        v3 -= 1e10
                        v5 += (v5 >> ((2) & 63))
                        _ = v29
                    } else {
                        var v30 string = ""
                        v31 := []int{fi28}
                        _ = v30
                        _ = v31
                    }
                    v32 := GenS2{F0: 127}
                    _ = v32
                }
                v33 := GenS2{F0: 112}
                _ = v33
            } else {
                v3 = float64(-2.25)
            }
            _ = fi28
        }
    } else {
        var v34 int = 711
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
