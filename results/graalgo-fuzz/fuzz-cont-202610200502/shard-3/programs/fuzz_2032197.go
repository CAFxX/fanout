package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint16) int16 {
    p0 -= (p0 | p0)
    if true {
        v7 := []int{128, 1, 1, 4294967295}
        if (42 <= 127) {
            var v8 uint32 = 0
            _ = v8
        }
        _ = v7
    }
    _ = p0
    return -77
}
func genFn1() int {
    var v9 int64 = -1
    var v10 float64 = 0.0
    v9 = 9223372036854775807
    _ = v10
    _ = v9
    return 435
}

func main() {
    var v1 uint64 = 6
    var v2 int16 = -1
    var v3 uint8 = 1
    var v4 uint16 = 125
    var v5 float32 = -2.25
    var v6 uint16 = (v4 + v4)
    var v11 int8 = 101
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := []int{}
        v3 = v3
        v14 := GenS1{F0: 0}
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v4, v1, v5)
    fmt.Println(v5, v1, v4)
    fmt.Println(v6, v2)
    var v15 int16 = v2
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 string = (("" + "\n\t") + ("hello" + "foo bar"))
        v18 := []int{fi16, 42, (^(fi16)), (fi16 >> ((4294967295) & 63))}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
