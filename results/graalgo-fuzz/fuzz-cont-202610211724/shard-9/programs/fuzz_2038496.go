package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint32 }
func genFn0() int {
    v6 := []int{}
    v7 := GenS1{F0: -1, F1: 1}
    v8 := GenS1{F0: -9223372036854775808, F1: 42}
    v9 := GenS2{F0: 737, F1: 0}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 128
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 622
    v3 := []string{("hello" + ""), (("hello" + ((("" + "\n\t") + ("foo bar" + "foo bar")) + "hello")) + ((("\n\t" + "hello") + "\n\t") + (("hello" + "a") + "\n\t"))), "foo bar", ""}
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    delete(v5, "a")
    fmt.Println(v2)
    var v10 float64 = 3.14159
    var v11 int16 = 323
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            delete(v1, 4294967295)
            _ = fi13
        }
        var v14 bool = false
        _ = fi12
        _ = v14
    }
    _ = len(v3)
    v11 += (v11 | v11)
    switch 128 {
    case 821:
        v2 -= v2
    case -2147483648:
        v2 = v2
    case 517:
        v11 = v11
    default:
        v10 -= v10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
