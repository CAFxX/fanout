package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() string {
    var v5 int8 = 100
    var v6 int8 = (v5 - 1)
    _ = v5
    _ = v6
    return ("hello" + "")
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    var v4 uint64 = 620
    v1 = v1
    v7 := GenS1{F0: 849}
    var v8 uint16 = 255
    var v9 string = "\n\t"
    delete(v2, 1)
    v10 := []int{4294967295, 325}
    if ((1e10 != ((-2.25 / ((1e10 - 3.14159) * 0.1)) * (1e10 - -2.25))) || false) {
        fmt.Println(v1, v9)
        switch 461 {
        case 32767:
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 uint16 = v8
                _ = fi11
                _ = v12
            }
        case 1:
            v8 = uint16(v4)
        }
    }
    v4 = 268
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
