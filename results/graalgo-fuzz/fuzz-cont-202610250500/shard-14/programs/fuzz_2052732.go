package main

import (
    "fmt"
)

func genFn0(p0 uint64) uint8 {
    fmt.Println(p0)
    v7 := []int{445, 4294967295, 42}
    _ = p0
    _ = v7
    return 223
}
func genFn1(p0 int16, p1 int64, p2 string) int16 {
    switch p2 {
    case "hello":
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                v10 := []int{(fi9 / ((fi9) | 1)), (fi8 * (fi9 % ((fi8) | 1)))}
                var v11 uint64 = 79
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
    case "\n\t":
        var v12 uint32 = 337
        _ = v12
    }
    p0 -= 32767
    var v13 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return (^(p0))
}

func main() {
    var v1 uint = 42
    v2 := make(map[int]uint32)
    v3 := []string{}
    v4 := []int{32768, 0, -2147483648}
    var v5 uint64 = 4294967295
    var v6 int32 = -901
    v6 = 665
    fmt.Println(v5, v5, v5)
    v14 := v2[423]
    fmt.Println(v1)
    v15 := v2[667]
    var v16 int8 = -86
    for fi17 := 0; fi17 < 10; fi17++ {
        if (!(false || (int(v6) <= fi17))) {
            delete(v2, -9223372036854775808)
            var v18 float64 = (1.5 * float64(-2.25))
            fmt.Println(fi17, v5)
            _ = v18
        }
        _ = fi17
    }
    delete(v2, 914)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
