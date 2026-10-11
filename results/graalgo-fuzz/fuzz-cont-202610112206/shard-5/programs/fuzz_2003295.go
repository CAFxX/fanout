package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int8; F1 int }
func genFn0() int {
    v4 := GenS1{F0: 199, F1: 18446744073709551615, F2: -333}
    var v5 bool = false
    v5 = ((696 == 42) || (0.1 < 1e10))
    if (!(1 < 2)) {
        v5 = true
        fmt.Println(v5)
    } else {
        v5 = (3.14159 >= (0.0 + float32(0.0)))
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 int = fi6
            _ = fi6
            _ = v7
        }
        if ((-361 == -1) || v5) {
            fmt.Println(v5)
            v8 := GenS2{F0: 65535, F1: 599}
            var v9 int = 2147483647
            _ = v8
            _ = v9
        } else {
            if false {
                v10 := GenS1{F0: 255, F1: 2, F2: 42}
                _ = v10
            }
            v11 := []int{-1}
            for fi12 := 0; fi12 < 3; fi12++ {
                for fi13 := 0; fi13 < 10; fi13++ {
                    v5 = ("\n\t" >= ("foo bar" + ("\n\t" + ("a" + ("a" + ("\n\t" + "\n\t"))))))
                    v5 = true
                    if false {
                        var v14 bool = ("foo bar" != "a")
                        _ = v14
                    }
                    switch fi13 {
                    case 9223372036854775807:
                        fmt.Println(fi12)
                    case 128:
                        var v15 int8 = 127
                        _ = v15
                    case -252:
                        fmt.Println(fi13, fi12, fi12)
                    default:
                        if (100 < 42) {
                            var v16 int8 = 100
                            _ = v16
                        } else {
                            var v17 uint16 = 65535
                            var v18 uint8 = 234
                            _ = v17
                            _ = v18
                        }
                    }
                    _ = fi13
                }
                v5 = (0 < 783)
                _ = fi12
            }
            _ = v11
        }
    }
    _ = v4
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 float32 = -2.25
    v2 := []string{}
    v3 := make(map[string]int)
    v1 += v1
    v1 -= 3.14159
    var v19 uint64 = 2
    v20 := GenS3{F0: 42, F1: 1}
    _ = len(v2)
    switch "\n\t" {
    case "foo bar":
        var v21 int32 = 797
        _ = v21
    case "\n\t":
        for fi22 := 0; fi22 < 1; fi22++ {
            switch genFn0() {
            case -9223372036854775808:
                v3["foo bar"] = fi22
            }
            _ = fi22
        }
    }
    fmt.Println(v1, v1)
    for fi23 := 0; fi23 < 18; fi23++ {
        v19 = 18446744073709551615
        var v24 int = fi23
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
