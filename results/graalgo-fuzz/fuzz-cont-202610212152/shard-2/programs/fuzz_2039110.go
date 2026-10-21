package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 bool = (42 == 739)
    v2 := make(map[string]int)
    var v3 float32 = (0.0 / (3.14159 - 0.1))
    v4 := GenS1{F0: 821}
    var v5 uint64 = 18446744073709551615
    fmt.Println(v5)
    v1 = v1
    for fi6 := 0; fi6 < 7; fi6++ {
        if v1 {
            v7 := []int{fi6, fi6, 4294967295, fi6}
            var v8 bool = false
            switch "hello" {
            case "hello":
                v1 = v8
            case "foo bar":
                v9 := GenS1{F0: 845}
                _ = v9
            }
            _ = v7
            _ = v8
        } else {
            var v10 uint = 382
            v11 := GenS2{F0: 32767}
            _ = v10
            _ = v11
        }
        _ = fi6
    }
    v12 := v2["hello"]
    v13 := v2["a"]
    v3 += 3.14159
    switch v13 {
    case 32767:
        v14 := []int{}
        _ = v14
    default:
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS1{F0: 2}
            fmt.Println(v1, v3, v12)
            _ = len(v2)
            _ = fi15
            _ = v16
        }
    }
    v17 := []int{v13}
    for fi18 := 0; fi18 < 18; fi18++ {
        if v1 {
            switch "foo bar" {
            case "hello":
                fmt.Println(v1, v13)
            }
            v5 = 4294967295
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            fmt.Println(fi18)
            v20 := GenS2{F0: 65535}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    v1 = ((0 == 42) || true)
    fmt.Println(v13, v1, v3)
    v5 += v5
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
