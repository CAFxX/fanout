package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() int16 {
    var v6 string = "\n\t"
    v6 = v6
    v6 = v6
    _ = v6
    return 1000
}
func genFn1() uint16 {
    v7 := GenS1{F0: 100}
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := GenS1{F0: -128}
            v11 := []int{fi8, fi9, fi9}
            _ = len(v11)
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 int8 = 127
        _ = fi8
        _ = v12
    }
    v13 := []int{127, -2147483648, 127, 543}
    _ = v13
    _ = v7
    return 1
}

func main() {
    var v1 int8 = 42
    var v2 int32 = 42
    var v3 bool = (-267 >= -1)
    var v4 uint32 = 2147483647
    var v5 uint8 = 128
    if ("hello" < "foo bar") {
        v14 := []int{127, -2147483648, 9223372036854775807}
        _ = v14
    }
    var v15 uint8 = 128
    v16 := GenS1{F0: (v1 & (v1 >> ((4294967295) & 63)))}
    var v17 float32 = 1.5
    v18 := []int{int(v2)}
    v19 := GenS1{F0: (v1 * -53)}
    v18 = append(v18, 0)
    v3 = v3
    for fi20 := 0; fi20 < 10; fi20++ {
        if len(v18) > 0 {
            v18[0] = 2147483647
        }
        _ = fi20
    }
    fmt.Println(v3, v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
