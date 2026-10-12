package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint8) int {
    v6 := []int{4294967295}
    if false {
        fmt.Println(p0, p1)
        var v7 int8 = 100
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return 699
}

func main() {
    var v1 int32 = 0
    v2 := []uint32{}
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    v5 := []string{("hello" + "a"), "\n\t", ""}
    fmt.Println(v1)
    delete(v4, 2147483647)
    v8 := v3[""]
    var v9 uint64 = 4294967295
    if (127 != 103) {
        for fi10 := 0; fi10 < 6; fi10++ {
            delete(v3, "foo bar")
            _ = fi10
        }
    }
    var v11 int8 = 1
    v11 = (v11 + v11)
    v3[("\n\t" + ("hello" + "\n\t"))] = v8
    v12 := GenS1{F0: 165}
    v11 = 100
    if (1 == 42) {
        switch 42 {
        case 128:
            fmt.Println(v9, v9)
        case 0:
            delete(v3, "a")
        default:
            if len(v2) > 0 {
                v2[0] = 2
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
