package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }

func main() {
    var v1 bool = (!false)
    var v2 int16 = 1000
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    v5 := make(map[int]string)
    var v6 int16 = 32767
    delete(v3, "hello")
    for fi7 := 0; fi7 < 13; fi7++ {
        if false {
            var v8 int16 = (v2 * (v6 ^ (v2 | (v2 & (-(v6))))))
            _ = v8
        } else {
            var v9 int8 = 1
            _ = v9
        }
        var v10 uint32 = 2147483648
        var v11 uint = 42
        _ = fi7
        _ = v10
        _ = v11
    }
    v2 += v6
    fmt.Println(v1, v2, v1)
    switch "foo bar" {
    case "\n\t":
        v2 = v2
    default:
        v12 := GenS1{F0: 164}
        _ = v12
    }
    v1 = v1
    if (-1 != 1) {
        v6 = v2
        for fi13 := 0; fi13 < 10; fi13++ {
            v1 = v1
            var v14 int16 = 906
            _ = fi13
            _ = v14
        }
    } else {
        if v1 {
            v15 := []int{}
            var v16 bool = (v1 && v1)
            var v17 int = 1
            _ = v15
            _ = v16
            _ = v17
        }
        v18 := GenS2{F0: "hello"}
        _ = v18
    }
    v5[-9223372036854775808] = ("\n\t" + "")
    var v19 int64 = -91
    delete(v5, 616)
    if v1 {
        v20 := []int{729, 0, 32768, 1}
        fmt.Println(v6, v19, v19)
        _ = v20
    }
    v6 -= v6
    if (((1.5 * (0.1 + -2.25)) < 3.14159) || (v1 && false)) {
        if (128 >= 127) {
            if v1 {
                v6 += (v2 ^ (v2 & 926))
                var v21 bool = false
                _ = v21
            }
        } else {
            v4[("foo bar" + "a")] = v1
            for fi22 := 0; fi22 < 10; fi22++ {
                if ("" <= ("" + "\n\t")) {
                    v2 = (v6 + v6)
                } else {
                    v23 := GenS3{F0: 348}
                    v6 -= v6
                    _ = v23
                }
                for fi24 := 0; fi24 < 11; fi24++ {
                    delete(v3, "hello")
                    _ = fi24
                }
                var v25 uint32 = 2
                _ = fi22
                _ = v25
            }
        }
    } else {
        v1 = (((true && true) && v1) || ((v6 % ((v2) | 1)) != v2))
        var v26 uint = 0
        _ = v26
    }
    v27 := []int{32767, -702, 297}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
