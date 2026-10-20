package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int8 }
func genFn0(p0 uint16) int64 {
    v7 := []int{}
    v8 := GenS1{F0: 18446744073709551615, F1: "hello", F2: 127}
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint32, p1 string, p2 string) int16 {
    var v9 int8 = 127
    p2 = ((p2 + "a") + p2)
    v10 := GenS1{F0: 31, F1: p1, F2: v9}
    v11 := []int{-1, 127}
    var v12 int32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (1 >= 42)
    var v3 float64 = 0.1
    var v4 uint16 = 42
    var v5 float32 = 0.1
    v6 := []uint32{65535}
    var v13 int16 = -319
    fmt.Println(v4)
    switch "hello" {
    case "a":
        _ = len(v1)
    case "\n\t":
        var v14 int = -822
        _ = v14
    default:
        if v2 {
            v15 := v1[749]
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 bool = true
                _ = fi16
                _ = v17
            }
            _ = v15
        }
    }
    v18 := v1[65535]
    var v19 string = "a"
    var v20 int8 = int8(v13)
    if (213 == 127) {
        var v21 bool = v2
        _ = v21
    }
    switch -2147483648 {
    case -386:
        fmt.Println(v13)
    default:
        switch (v19 + "\n\t") {
        case "a":
            v19 = "\n\t"
        case "hello":
            fmt.Println(v13)
        case "\n\t":
            for fi22 := 0; fi22 < 7; fi22++ {
                for fi23 := 0; fi23 < 1; fi23++ {
                    delete(v1, 586)
                    var v24 bool = (255 == 4294967295)
                    var v25 bool = (-1 > -362)
                    v19 = v18
                    var v26 float32 = 1e10
                    fmt.Println(v2, v24, v26)
                    var v27 int8 = (v20 & v20)
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = fi22
            }
        default:
            var v28 bool = (!(!(genFn1(101, v18, ("a" + ("foo bar" + v19))) <= 42)))
            _ = v28
        }
    }
    fmt.Println(v5, v13, v20)
    var v29 bool = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
