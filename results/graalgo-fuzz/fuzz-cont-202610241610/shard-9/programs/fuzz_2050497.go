package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []string{"hello"}
    var v2 int = -9223372036854775808
    var v3 float32 = 0.0
    for fi4 := 0; fi4 < 7; fi4++ {
        fmt.Println(v3, v3)
        v2 = fi4
        _ = fi4
    }
    v2 -= -893
    v1 = append(v1, (("a" + "foo bar") + "\n\t"))
    for fi5 := 0; fi5 < 5; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            if (0.0 != ((float32(3.14159) - 0.0) * 1e10)) {
                var v7 uint32 = 231
                _ = v7
            }
            _ = fi6
        }
        var v8 float64 = (3.14159 * 1e10)
        _ = fi5
        _ = v8
    }
    v9 := GenS1{F0: 18446744073709551615}
    var v10 int16 = -32768
    v11 := GenS1{F0: 4294967295}
    fmt.Println(v10, v10)
    var v12 int = (v2 + 42)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
