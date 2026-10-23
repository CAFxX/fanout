package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint32, p1 int8, p2 uint32) int64 {
    p2 = p2
    var v4 uint = 797
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]int)
    var v3 bool = false
    var v5 float32 = (1.5 / (((0.1 - 1e10) / 1e10) / 3.14159))
    var v6 int = 1
    v7 := []int{(v6 - -783), (v6 * 128), v6, v6}
    var v8 int32 = -2147483648
    if v3 {
        for fi9 := 0; fi9 < 13; fi9++ {
            v5 = 1.5
            for fi10 := 0; fi10 < 13; fi10++ {
                v3 = (-808 < fi9)
                v11 := v2[-1]
                var v12 int = fi10
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
        var v13 float64 = 0.1
        _ = v13
    } else {
        v2[(v6 << ((1) & 63))] = (v6 + (v6 % ((2147483647) | 1)))
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 float32 = v5
            var v16 uint64 = 4294967295
            v17 := GenS1{F0: uint(v1)}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        var v18 string = ("" + ("" + ("foo bar" + "foo bar")))
        _ = v18
    }
    delete(v2, 2147483647)
    if v3 {
        v19 := GenS1{F0: 18446744073709551615}
        _ = v19
    } else {
        if v3 {
            var v20 bool = ((v6 & v6) != (v6 << ((773) & 63)))
            v6 = 0
            _ = v20
        }
        v6 += v6
        fmt.Println(v6, v3, v3)
    }
    fmt.Println(v8)
    if v3 {
        var v21 string = (("" + "a") + ("foo bar" + ""))
        v6 += (v6 % ((v6) | 1))
        delete(v2, (v6 % ((-1) | 1)))
        _ = v21
    } else {
        v7 = append(v7, (^(v6)))
        var v22 int16 = -926
        var v23 int32 = v8
        _ = v22
        _ = v23
    }
    if (v3 && v3) {
        var v24 int8 = -121
        _ = v24
    } else {
        v25 := GenS1{F0: 822}
        v3 = v3
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
