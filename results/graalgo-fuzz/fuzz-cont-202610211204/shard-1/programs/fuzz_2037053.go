package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []string{("hello" + "\n\t"), "\n\t", "\n\t"}
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    v4 := GenS1{F0: 1}
    switch 689 {
    case 128:
        for fi5 := 0; fi5 < 3; fi5++ {
            v6 := GenS1{F0: -9223372036854775808}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v7 := GenS1{F0: -829}
            _ = fi5
            _ = v6
            _ = v7
        }
    case 268:
        var v8 int16 = 893
        _ = v8
    case 4294967295:
        var v9 int = -950
        _ = v9
    }
    v10 := GenS1{F0: 995}
    var v11 int8 = 127
    var v12 int64 = -901
    fmt.Println(v12, v11)
    var v13 float64 = 0.1
    v14 := GenS1{F0: -505}
    var v15 uint = 65535
    fmt.Println(v15, v13)
    v16 := GenS1{F0: (v12 % (((v12 * v12)) | 1))}
    v13 = 1e10
    v17 := GenS1{F0: (v12 | 172)}
    v11 = v11
    v15 -= (v15 | 42)
    v11 += v11
    v1 = append(v1, "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
