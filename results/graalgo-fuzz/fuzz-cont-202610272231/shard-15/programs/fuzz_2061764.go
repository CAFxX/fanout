package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int = -9223372036854775808
    var v2 int32 = 42
    var v3 uint64 = 0
    var v4 string = "a"
    v5 := make(map[string]uint32)
    var v6 bool = false
    var v7 int = (v1 % (((v1 & v1)) | 1))
    delete(v5, v4)
    if (("foo bar" + v4) <= v4) {
        var v8 float64 = (3.14159 / 1.0)
        fmt.Println(v8, v3)
        v9 := []int{}
        _ = v8
        _ = v9
    }
    v2 = (^(v2))
    if ((-767 == -32768) || v6) {
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 float32 = 1.5
            _ = fi10
            _ = v11
        }
    } else {
        var v12 bool = (42 <= 0)
        var v13 uint16 = 42
        _ = v12
        _ = v13
    }
    if v6 {
        var v14 uint16 = 0
        var v15 bool = true
        var v16 int64 = -540
        _ = v14
        _ = v15
        _ = v16
    } else {
        var v17 float64 = -2.25
        _ = v17
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 uint16 = 235
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: v3}
    switch ((v4 + v4) + (("foo bar" + (("" + v4) + v4)) + v4)) {
    case "\n\t":
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := GenS1{F0: (^(v3))}
            v23 := GenS1{F0: 2}
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    var v24 int64 = 629
    if v6 {
        v6 = v6
        var v25 uint16 = 2
        _ = v25
    }
    _ = len(v5)
    var v26 uint16 = 32767
    v27 := GenS1{F0: (v3 + (v3 << ((42) & 63)))}
    fmt.Println(v4)
    v3 += v3
    fmt.Println(v26, v4, v24)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
