package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int64 }

func main() {
    var v1 bool = (true || false)
    var v2 uint8 = 0
    var v3 uint16 = 402
    v4 := make(map[string]int)
    v5 := make(map[uint32]uint32)
    var v6 int8 = -1
    var v7 string = "hello"
    for fi8 := 0; fi8 < 6; fi8++ {
        if v1 {
            switch "" {
            case "a":
                for fi9 := 0; fi9 < 3; fi9++ {
                    if (4294967295 <= 4294967295) {
                        if v1 {
                            fmt.Println(fi8, v2, v1)
                            fmt.Println(v1)
                            v1 = true
                        }
                        v10 := GenS1{F0: 0, F1: -1, F2: 0}
                        v7 = "\n\t"
                        _ = v10
                    } else {
                        var v11 uint64 = 2
                        v2 = (v2 - (v2 << ((18446744073709551615) & 63)))
                        v1 = ((v11 <= v11) || (!(!true)))
                        _ = v11
                    }
                    delete(v4, v7)
                    _ = fi9
                }
            case "foo bar":
                delete(v4, "foo bar")
            case "":
                v6 = (^(v6))
            default:
                var v12 uint8 = 2
                _ = v12
            }
            fmt.Println(v6, fi8)
            for fi13 := 0; fi13 < 20; fi13++ {
                switch "" {
                case "":
                    var v14 uint64 = 390
                    _ = v14
                case "\n\t":
                    v4[v7] = fi8
                case "foo bar":
                    if false {
                        if (!v1) {
                            v1 = true
                        }
                    }
                default:
                    v6 += (v6 + (v6 | v6))
                }
                _ = fi13
            }
        }
        v15 := []int{fi8, (fi8 + fi8)}
        _ = fi8
        _ = v15
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        if (!v1) {
            v2 = 170
            v6 = -128
        } else {
            var v17 uint = 65535
            _ = v17
        }
        _ = fi16
    }
    var v18 int = 877
    var v19 float32 = 1.5
    var v20 int8 = -1
    v21 := []int{}
    v3 = (v3 / ((v3) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
