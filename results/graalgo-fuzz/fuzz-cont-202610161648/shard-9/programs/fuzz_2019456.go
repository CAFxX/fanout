package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float64 = 0.1
    var v2 bool = (9223372036854775807 >= 473)
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 9; fi4++ {
        if true {
            v2 = ((((0.0 + 1.5) < (1.5 / float32(-2.25))) && v2) || (!v2))
            switch (fi4 / ((fi4) | 1)) {
            case 32767:
                var v5 uint16 = 530
                _ = v5
            }
        } else {
            fmt.Println(v2, v2)
        }
        v1 -= (1.5 + v1)
        _ = fi4
    }
    if true {
        v3[65535] = "hello"
        v2 = (245 >= 1)
        v2 = v2
    }
    delete(v3, 42)
    for fi6 := 0; fi6 < 6; fi6++ {
        v2 = (100 != 0)
        v7 := []int{fi6, (fi6 % ((2147483647) | 1))}
        fmt.Println(fi6, fi6)
        _ = fi6
        _ = v7
    }
    delete(v3, 1)
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 float64 = v1
        delete(v3, 0)
        _ = fi8
        _ = v9
    }
    var v10 int64 = 42
    switch 2147483647 {
    case -9223372036854775808:
        var v11 uint16 = 65535
        _ = v11
    case 4294967295:
        v12 := GenS1{F0: 1000}
        _ = v12
    case 38:
        fmt.Println(v2, v1)
    default:
        v13 := []int{157, -274, 0}
        _ = v13
    }
    var v14 uint32 = 1
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
}
