package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float32 = (1e10 * ((3.14159 * (1e10 - (((float32(-2.25) / float32(1e10)) * 0.0) * ((float32(1e10) * 0.0) + float32(1e10))))) / 1.5))
    v2 := make(map[int]bool)
    var v3 float32 = v1
    v4 := make(map[string]uint32)
    v5 := []string{}
    var v6 bool = false
    v6 = false
    delete(v4, ("hello" + ""))
    var v7 int16 = 32767
    var v8 int32 = -828
    var v9 int64 = 42
    v10 := GenS1{F0: 4294967295, F1: 127}
    for fi11 := 0; fi11 < 20; fi11++ {
        if (!(2 != 255)) {
            v12 := GenS1{F0: 2, F1: 126}
            for fi13 := 0; fi13 < 12; fi13++ {
                if v6 {
                    v1 -= (v1 * 0.1)
                    fmt.Println(v9, fi11, v9)
                    v1 -= 1.5
                } else {
                    v14 := v4["hello"]
                    v15 := GenS2{F0: 581}
                    _ = v14
                    _ = v15
                }
                var v16 uint16 = 255
                _ = len(v5)
                _ = fi13
                _ = v16
            }
            _ = v12
        } else {
            v17 := GenS1{F0: 4294967295, F1: 0}
            v1 += (v3 * v1)
            _ = v17
        }
        _ = fi11
    }
    switch -689 {
    case -194:
        var v18 bool = (v8 != v8)
        _ = v18
    case 128:
        delete(v2, 166)
    default:
        v9 += (v9 | 1)
    }
    var v19 uint64 = 2
    v5 = append(v5, "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
