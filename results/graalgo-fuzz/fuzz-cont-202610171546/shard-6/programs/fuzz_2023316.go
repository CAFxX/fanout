package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint64; F2 int8 }
type GenS3 struct { F0 string; F1 int32; F2 int16 }
func genFn0(p0 int64) uint64 {
    p0 -= (p0 / ((p0) | 1))
    p0 -= 1
    _ = p0
    return 131
}

func main() {
    var v1 int16 = 42
    var v2 int16 = (v1 % ((v1) | 1))
    var v3 bool = true
    v4 := make(map[string]string)
    v5 := []int{32768, 2147483647, 42}
    var v6 uint8 = 0
    if v3 {
        for fi7 := 0; fi7 < 15; fi7++ {
            for fi8 := 0; fi8 < 15; fi8++ {
                v9 := GenS1{F0: 50}
                var v10 uint16 = 1
                if len(v5) > 0 {
                    v5[0] = (fi7 | fi8)
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
        v3 = false
        delete(v4, "a")
    } else {
        _ = len(v5)
        v11 := GenS1{F0: 255}
        _ = v11
    }
    var v12 int8 = 26
    delete(v4, "\n\t")
    v13 := GenS1{F0: 897}
    if v3 {
        v14 := GenS1{F0: 0}
        _ = v14
    }
    v15 := v4["a"]
    for fi16 := 0; fi16 < 1; fi16++ {
        fmt.Println(v15)
        v17 := GenS2{F0: (!v3), F1: 663, F2: (v12 / ((v12) | 1))}
        v15 = v15
        _ = fi16
        _ = v17
    }
    v15 = ("" + v15)
    v4["foo bar"] = ("foo bar" + "foo bar")
    var v18 int16 = -32768
    for fi19 := 0; fi19 < 11; fi19++ {
        v6 = 127
        if (-540 == 1) {
            fmt.Println(v1, v1)
        } else {
            v4[v15] = ""
        }
        _ = fi19
    }
    v20 := GenS3{F0: ((v15 + "\n\t") + "foo bar"), F1: 42, F2: v2}
    v21 := GenS2{F0: (-383 > 32768), F1: 2, F2: (v12 & v12)}
    var v22 uint16 = 1
    v6 = (v6 - v6)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
