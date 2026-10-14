package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int64 }

func main() {
    v1 := []uint32{508, 65535, 2147483648, 0}
    var v2 float64 = 3.14159
    var v3 int = 127
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    v2 -= v2
    v4 := GenS1{F0: 1, F1: -1, F2: -9223372036854775808}
    fmt.Println(v3, v3)
    v5 := GenS1{F0: 2, F1: 0, F2: -984}
    v3 += 32768
    var v6 int = v3
    v7 := GenS1{F0: 268, F1: 0, F2: 9223372036854775807}
    v2 -= float64((3.14159 - 1.5))
    if (!true) {
        var v8 int16 = 937
        if len(v1) > 0 {
            v1[0] = 65535
        }
        if true {
            v8 += v8
            v9 := GenS1{F0: 32767, F1: 100, F2: -831}
            _ = v9
        }
        _ = v8
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: 42, F1: 127, F2: 801}
            var v12 uint = 0
            _ = fi10
            _ = v11
            _ = v12
        }
        v2 -= 0.0
    }
    if false {
        v13 := GenS1{F0: 0, F1: 127, F2: 967}
        var v14 uint = 601
        var v15 int8 = 0
        _ = v13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        v6 += (v6 + v3)
        v17 := GenS1{F0: 578, F1: 0, F2: 0}
        _ = fi16
        _ = v17
    }
    if false {
        var v18 uint64 = 797
        _ = v18
    } else {
        switch "foo bar" {
        case "hello":
            var v19 int64 = 1
            _ = v19
        }
    }
    _ = len(v1)
    var v20 int = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
