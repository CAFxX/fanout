package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint) int8 {
    var v4 uint64 = 4294967295
    var v5 int64 = -9223372036854775808
    p0 = p0
    fmt.Println(p0, p0)
    for fi6 := 0; fi6 < 17; fi6++ {
        p0 += (p0 % ((65535) | 1))
        v7 := GenS1{F0: 2}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := []int{32767}
    var v2 int = 128
    var v3 int = 106
    var v8 int8 = 42
    v9 := GenS1{F0: 0}
    fmt.Println(v8, v2)
    if false {
        if (v8 == (v8 | genFn0(255))) {
            var v10 uint = 2
            for fi11 := 0; fi11 < 7; fi11++ {
                if (true || true) {
                    fmt.Println(v3, v8, v3)
                    var v12 uint32 = 781
                    _ = v12
                }
                var v13 uint = v10
                _ = fi11
                _ = v13
            }
            _ = v10
        } else {
            var v14 bool = (1 < 127)
            _ = v14
        }
    } else {
        var v15 string = ""
        _ = v15
    }
    fmt.Println(v2)
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 int32 = 0
        _ = fi16
        _ = v17
    }
    fmt.Println(v3, v2, v3)
    if (false && true) {
        v3 += 127
        fmt.Println(v3, v2, v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
