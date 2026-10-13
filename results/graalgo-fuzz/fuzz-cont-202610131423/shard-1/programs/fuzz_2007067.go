package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int, p1 uint64, p2 int) uint8 {
    p0 = p0
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(p1, p2, p0)
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return 249
}

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    var v4 float32 = (3.14159 - 0.1)
    v5 := make(map[string]string)
    if (42 == 32767) {
        v7 := []int{32768}
        var v8 int8 = 100
        v9 := GenS1{F0: -551}
        _ = v7
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS1{F0: 579}
            var v12 int64 = 42
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 2; fi13++ {
            for fi14 := 0; fi14 < 19; fi14++ {
                v15 := []int{}
                _ = fi14
                _ = v15
            }
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 string = (("\n\t" + "a") + ("hello" + ("hello" + "hello")))
                v1 = v1
                v18 := GenS1{F0: 42}
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi13
        }
    }
    switch 32768 {
    case -2147483648:
        v4 = 1e10
    case 0:
        var v19 string = (("\n\t" + ("a" + "hello")) + "hello")
        _ = v19
    case 1:
        v20 := GenS1{F0: -1}
        _ = v20
    }
    v4 = float32(0.1)
    v1 = true
    v21 := []int{}
    var v22 uint16 = 28
    v4 -= (v4 + v4)
    var v23 int = 967
    if false {
        var v24 uint32 = 1
        fmt.Println(v23, v23, v4)
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
