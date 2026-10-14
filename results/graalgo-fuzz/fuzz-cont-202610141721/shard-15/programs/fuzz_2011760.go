package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int64, p1 int64, p2 string) uint64 {
    p1 += p0
    p1 = 300
    v5 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}

func main() {
    var v1 int64 = -203
    v2 := make(map[string]uint32)
    v3 := []uint32{65535, 2147483647, 4294967295}
    v4 := []string{}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        v1 = (v1 & 1)
        for fi6 := 0; fi6 < 16; fi6++ {
            v1 = v1
            switch "hello" {
            case "\n\t":
                var v7 int = (^(fi6))
                _ = v7
            case "":
                if false {
                    for fi8 := 0; fi8 < 20; fi8++ {
                        _ = len(v2)
                        _ = fi8
                    }
                } else {
                    for fi9 := 0; fi9 < 5; fi9++ {
                        var v10 bool = (true || true)
                        _ = fi9
                        _ = v10
                    }
                    if false {
                        v1 = (v1 * v1)
                    }
                    var v11 int32 = 1
                    _ = v11
                }
            }
            _ = fi6
        }
    }
    fmt.Println(v1)
    var v12 int16 = 32767
    v1 -= (-(v1))
    var v13 int8 = -43
    if true {
        fmt.Println(v13, v13)
    }
    v14 := GenS1{F0: 2}
    fmt.Println(v12, v12, v12)
    for fi15 := 0; fi15 < 7; fi15++ {
        v16 := GenS1{F0: 42}
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
