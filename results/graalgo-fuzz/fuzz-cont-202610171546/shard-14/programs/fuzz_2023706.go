package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int32, p1 string) string {
    if true {
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 int16 = -32768
            _ = fi5
            _ = v6
        }
        var v7 uint8 = 0
        fmt.Println(p1, p0)
        _ = v7
    } else {
        var v8 int = -998
        v9 := GenS1{F0: 42}
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 1}
    v11 := []int{-1, 32768, 0, 9223372036854775807}
    var v12 int = 4294967295
    fmt.Println(p1, p0, v12)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return p1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 42
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    var v13 int = -9223372036854775808
    var v14 int32 = 42
    _ = len(v1)
    v15 := []int{(v13 / ((-435) | 1)), 253, 882}
    var v16 uint32 = 1
    var v17 uint = 0
    v13 = 825
    switch (v13 >> (((v17 * v17)) & 63)) {
    case 2147483647:
        v4[v16] = -291
    case 42:
        var v18 uint = v17
        _ = v18
    }
    v19 := GenS1{F0: 42}
    for fi20 := 0; fi20 < 1; fi20++ {
        v21 := GenS1{F0: v14}
        if (!false) {
            if len(v15) > 0 {
                v15[0] = (v13 | (v13 >> ((v17) & 63)))
            }
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            fmt.Println(v2, fi20)
        }
        _ = fi20
        _ = v21
    }
    var v22 float32 = 1e10
    var v23 int16 = 1
    v15 = append(v15, 42)
    for fi24 := 0; fi24 < 19; fi24++ {
        v25 := v4[(v16 + 2)]
        var v26 uint8 = 120
        delete(v3, (("" + "\n\t") + ((("\n\t" + "\n\t") + "a") + "a")))
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println(v22, v23)
    var v27 int64 = 0
    v23 = 861
    v13 -= v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
}
