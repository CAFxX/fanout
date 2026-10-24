package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]string)
    var v3 string = "a"
    var v4 int64 = 0
    v5 := make(map[string]uint32)
    var v6 int64 = v4
    for fi7 := 0; fi7 < 15; fi7++ {
        v1 = v1
        v8 := GenS1{F0: -490, F1: 473}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 1000, F1: 586}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int64 = (v4 * 0)
        if false {
            var v13 uint = 18446744073709551615
            v14 := []int{(fi11 % (((fi11 + fi11)) | 1)), -259, (fi11 & fi11), (fi11 << ((v13) & 63))}
            v6 = (v4 - 902)
            _ = v13
            _ = v14
        }
        if true {
            v2[v1] = "a"
            v3 = ""
        }
        _ = fi11
        _ = v12
    }
    v1 = (v1 + (v1 & v1))
    switch 4294967295 {
    case 32767:
        var v15 string = "hello"
        _ = v15
    case 127:
        var v16 uint64 = 356
        _ = v16
    }
    var v17 bool = (!(88 <= -2147483648))
    for fi18 := 0; fi18 < 10; fi18++ {
        fmt.Println(v17, v6, v6)
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
