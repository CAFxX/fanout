package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int8; F1 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 2
    v3 := []string{"hello"}
    var v4 int32 = 923
    var v5 int = -2147483648
    _ = len(v1)
    v6 := GenS1{F0: 809, F1: 308}
    v4 += v4
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 65535, F1: -1}
        _ = fi7
        _ = v8
    }
    switch (("" + "") + ("a" + "\n\t")) {
    case "a":
        var v9 uint16 = 51
        _ = v9
    case "hello":
        var v10 int32 = (-(v4))
        _ = v10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        if len(v3) > 0 {
            v3[0] = ""
        }
        v4 = v4
        _ = fi11
    }
    if (!false) {
        if ((("a" < "\n\t") || false) || (-110 >= 1)) {
            v3 = append(v3, "\n\t")
            var v12 float32 = 3.14159
            _ = v12
        }
        var v13 int64 = -9223372036854775808
        v14 := GenS2{F0: 120, F1: 4294967295}
        _ = v13
        _ = v14
    } else {
        var v15 int32 = v4
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 int64 = 663
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    if true {
        var v18 int32 = v4
        var v19 int = v5
        fmt.Println(v18, v19, v2)
        _ = v18
        _ = v19
    } else {
        var v20 int8 = 0
        v2 += 4294967295
        _ = v20
    }
    if (!(32 == 42)) {
        var v21 int = v5
        _ = v21
    }
    var v22 float32 = 1e10
    v1[-885] = ""
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
