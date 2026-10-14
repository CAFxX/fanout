package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int64) int64 {
    v4 := []int{128}
    p0 = p0
    p0 += -454
    p0 = p0
    _ = p0
    _ = v4
    return (p0 + (p0 + p0))
}
func genFn1(p0 int) uint8 {
    var v5 uint = 42
    v6 := []int{(p0 ^ -303)}
    var v7 float64 = -2.25
    fmt.Println(v5, p0)
    v5 = v5
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int16 = 937
    v2 := make(map[uint32]bool)
    var v3 uint16 = 255
    var v8 int64 = int64(v1)
    v8 = v8
    var v9 int64 = v8
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int32 = 42
            _ = fi11
            _ = v12
        }
        if (!true) {
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 string = ("a" + "foo bar")
                _ = fi13
                _ = v14
            }
            var v15 int32 = 42
            if true {
                _ = len(v2)
                v16 := GenS1{F0: 749}
                _ = v16
            }
            _ = v15
        }
        _ = fi10
    }
    v17 := GenS1{F0: 926}
    v3 = v3
    v8 = (v9 ^ genFn0((v8 % ((v9) | 1))))
    v3 += v3
    v9 += genFn0(-1)
    v1 = v1
    var v18 uint = 255
    var v19 int8 = 1
    fmt.Println(v9, v19, v8)
    var v20 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
