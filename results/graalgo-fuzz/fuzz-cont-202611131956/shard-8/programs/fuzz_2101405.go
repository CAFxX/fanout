package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int; F2 int32 }
type GenS3 struct { F0 int32 }
func genFn0() uint64 {
    var v7 uint16 = 1
    var v8 int = 127
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 string) int {
    v9 := []int{-9223372036854775808}
    p0 = "foo bar"
    var v10 int64 = -740
    var v11 bool = true
    var v12 string = ("foo bar" + p0)
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return -417
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{865}
    var v3 bool = false
    v4 := make(map[string]int)
    v5 := make(map[int]int)
    var v6 bool = false
    v6 = v6
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        if v6 {
            var v14 int32 = 1
            v15 := GenS1{F0: 1}
            var v16 bool = false
            _ = v14
            _ = v15
            _ = v16
        } else {
            var v17 uint = 64
            var v18 uint64 = 490
            v6 = (v18 < 4294967295)
            _ = v17
            _ = v18
        }
        _ = fi13
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v2 = append(v2, genFn1(("a" + "\n\t")))
        if (-1 <= 226) {
            for fi20 := 0; fi20 < 11; fi20++ {
                v21 := GenS1{F0: 138}
                fmt.Println(fi19)
                _ = fi20
                _ = v21
            }
            var v22 uint = 4294967295
            _ = v22
        }
        v6 = false
        _ = fi19
    }
    v3 = (42 >= 169)
    v6 = v3
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 int32 = -194
        _ = fi23
        _ = v24
    }
    v3 = false
    v25 := GenS2{F0: 42, F1: 834, F2: -791}
    fmt.Println(v6, v6)
    if (1 != -40) {
        if true {
            v26 := GenS3{F0: 80}
            for fi27 := 0; fi27 < 20; fi27++ {
                for fi28 := 0; fi28 < 20; fi28++ {
                    fmt.Println(v6, fi27)
                    v1["hello"] = 65535
                    _ = fi28
                }
                var v29 float64 = ((1e10 + 0.1) * (((0.0 * -2.25) + 3.14159) / (0.1 / ((-2.25 / 1.5) * (((1.5 / 1.0) * 0.0) - 1e10)))))
                _ = fi27
                _ = v29
            }
            _ = v26
        }
    } else {
        v3 = v3
    }
    v6 = true
    v6 = true
    var v30 uint32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
