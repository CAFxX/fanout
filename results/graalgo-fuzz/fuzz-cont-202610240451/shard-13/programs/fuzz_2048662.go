package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() string {
    if true {
        var v4 string = ""
        v5 := []int{}
        _ = v4
        _ = v5
    }
    var v6 bool = (!(-58 > 42))
    _ = v6
    return "\n\t"
}
func genFn1(p0 int32, p1 string) uint16 {
    v7 := []int{}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    p0 = (p0 % ((p0) | 1))
    var v8 bool = ((!(3.14159 != (((3.14159 + 3.14159) + (1e10 - -2.25)) - 0.0))) || ("a" > p1))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 int = 415
    var v2 int64 = 485
    v3 := make(map[uint32]bool)
    v2 = v2
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := GenS1{F0: -1}
        if false {
            if ((-128 <= 0) && (-422 >= -1)) {
                v11 := GenS1{F0: -28}
                _ = v11
            }
            fmt.Println(v2, v2)
            fmt.Println(v2)
        }
        v12 := []int{fi9, (-(v1)), 369, fi9}
        _ = fi9
        _ = v10
        _ = v12
    }
    var v13 int = v1
    var v14 uint16 = 0
    v15 := GenS1{F0: 42}
    fmt.Println(v2, v14, v1)
    v16 := GenS1{F0: 48}
    fmt.Println(v14, v14)
    var v17 int64 = v2
    v18 := []int{}
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 int64 = (v2 >> ((0) & 63))
        for fi21 := 0; fi21 < 4; fi21++ {
            fmt.Println(v20)
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
