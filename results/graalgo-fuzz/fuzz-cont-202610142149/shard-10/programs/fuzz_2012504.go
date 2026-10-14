package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 string }
func genFn0(p0 int16) bool {
    v6 := GenS1{F0: "a", F1: "a", F2: "hello"}
    v7 := []int{128}
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: "a", F1: "", F2: ("hello" + "")}
            _ = fi9
            _ = v10
        }
        v7 = append(v7, int(p0))
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS1{F0: "a", F1: ("\n\t" + "\n\t"), F2: "hello"}
            _ = fi11
            _ = v12
        }
        _ = fi8
    }
    fmt.Println(p0)
    p0 += (p0 << ((482) & 63))
    _ = p0
    _ = v6
    _ = v7
    return true
}

func main() {
    var v1 int32 = 0
    var v2 uint = 42
    var v3 uint64 = 2
    var v4 uint32 = 897
    var v5 int8 = 127
    var v13 float64 = 1e10
    var v14 uint32 = 65535
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := []int{(fi15 - (^(fi15))), fi15, (^(fi15)), fi15}
        _ = fi15
        _ = v16
    }
    var v17 int64 = 1
    v13 = 1.5
    var v18 int16 = 0
    v19 := []int{128, 127}
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
