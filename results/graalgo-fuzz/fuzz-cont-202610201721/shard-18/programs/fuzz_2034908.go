package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 string = (("\n\t" + "foo bar") + "\n\t")
    v3 := make(map[int]uint32)
    fmt.Println(v2)
    for fi4 := 0; fi4 < 12; fi4++ {
        if (1e10 != ((float32(float64(3.14159)) * -2.25) * -2.25)) {
            v5 := GenS1{F0: 0, F1: 0, F2: 0}
            _ = v5
        }
        for fi6 := 0; fi6 < 16; fi6++ {
            fmt.Println(fi6, v2, fi6)
            v7 := []int{fi6, -267}
            _ = fi6
            _ = v7
        }
        var v8 int64 = 9223372036854775807
        _ = fi4
        _ = v8
    }
    v2 = v2
    v2 = (v2 + (("\n\t" + "a") + ("" + (v2 + ("" + v2)))))
    delete(v1, "")
    v9 := GenS1{F0: 4294967295, F1: -1, F2: 350}
    if false {
        for fi10 := 0; fi10 < 3; fi10++ {
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 uint16 = 42
                var v13 int = fi10
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
    }
    var v14 bool = (-1 > 948)
    v14 = (false || (18 >= 42))
    var v15 int64 = -991
    var v16 string = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
