package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 18446744073709551615
    var v3 uint32 = 1
    if true {
        if ((3.14159 + 3.14159) < 3.14159) {
            v3 -= v3
        }
        var v4 int8 = -101
        fmt.Println(v4)
        _ = v4
    }
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(v3, v3)
        if (438 <= 1) {
            var v6 int32 = -974
            fmt.Println(v6, v6)
            v7 := GenS1{F0: 672, F1: 42, F2: int(v6)}
            _ = v6
            _ = v7
        } else {
            var v8 bool = true
            var v9 string = ""
            _ = v8
            _ = v9
        }
        v10 := []int{fi5, (fi5 << (((v2 << ((v2) & 63))) & 63)), (fi5 & fi5)}
        _ = fi5
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        if (true || (499 <= int(v3))) {
            var v12 int16 = 0
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS1{F0: 2, F1: 9223372036854775807, F2: 0}
                v15 := GenS2{F0: 100}
                var v16 int8 = -1
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            _ = v12
        }
        var v17 int8 = -128
        fmt.Println(fi11, v3, v17)
        _ = fi11
        _ = v17
    }
    var v18 int16 = -1
    if false {
        for fi19 := 0; fi19 < 5; fi19++ {
            v2 -= (^(v2))
            for fi20 := 0; fi20 < 10; fi20++ {
                if (!((((v18 == v18) && false) && ((v18 <= (v18 + -32768)) && true)) || (true || (false && (-1 >= 2147483647))))) {
                    var v21 uint16 = 255
                    _ = v21
                }
                if false {
                    v22 := GenS3{F0: 0, F1: 42}
                    _ = v22
                } else {
                    v23 := GenS2{F0: 93}
                    fmt.Println(v18)
                    if false {
                        v24 := []int{(fi20 % (((-(fi20))) | 1)), (fi20 + fi19), 0, 127}
                        _ = v24
                    }
                    _ = v23
                }
                _ = fi20
            }
            var v25 float32 = float32(0.1)
            _ = fi19
            _ = v25
        }
        v26 := v1[2147483647]
        _ = v26
    } else {
        var v27 int32 = -890
        _ = v27
    }
    var v28 float32 = 0.1
    var v29 string = ""
    var v30 int16 = v18
    v31 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
