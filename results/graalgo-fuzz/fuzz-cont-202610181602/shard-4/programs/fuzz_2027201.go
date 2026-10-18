package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint32, p2 uint16) uint {
    for fi4 := 0; fi4 < 16; fi4++ {
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 uint8 = 2
            p0 -= (p0 % ((p2) | 1))
            var v7 uint8 = v6
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    var v8 int32 = 295
    var v9 int64 = 42
    fmt.Println(v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 431
}

func main() {
    var v1 int8 = -1
    v2 := []int{}
    var v3 string = "a"
    var v10 string = "hello"
    v1 += -1
    v10 = "foo bar"
    var v11 float64 = (3.14159 / ((-2.25 * 1e10) * (-2.25 - 3.14159)))
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: 651}
        _ = fi12
        _ = v13
    }
    fmt.Println(v10, v10)
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            switch (fi15 % ((fi15) | 1)) {
            case 42:
                v1 = v1
            }
            _ = fi15
        }
        v3 = "hello"
        _ = fi14
    }
    var v16 uint = 808
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 int8 = v1
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 int64 = 1
        var v21 float64 = (0.1 / 1.0)
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
}
