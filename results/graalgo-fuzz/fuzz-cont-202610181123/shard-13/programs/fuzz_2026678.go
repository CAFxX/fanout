package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint; F1 int8; F2 uint }
func genFn0(p0 uint8) int16 {
    v5 := GenS1{F0: 42, F1: 42}
    p0 = (^(p0))
    fmt.Println(p0)
    var v6 int32 = 2147483647
    v7 := GenS1{F0: 0, F1: 42}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -32768
}

func main() {
    var v1 uint8 = 127
    var v2 int32 = 0
    v3 := []string{"hello", ("\n\t" + "")}
    v4 := make(map[string]string)
    var v8 int8 = 31
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS2{F0: 18446744073709551615, F1: (v8 * (v8 >> ((4294967295) & 63))), F2: 2}
            fmt.Println(v1, v2, v2)
            var v11 uint32 = 421
            _ = fi9
            _ = v10
            _ = v11
        }
    } else {
        v8 += v8
    }
    v3 = append(v3, "foo bar")
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(fi12)
        _ = fi12
    }
    var v13 uint32 = 2147483648
    fmt.Println(v2, v13)
    var v14 int64 = -2
    var v15 uint64 = 1
    for fi16 := 0; fi16 < 4; fi16++ {
        v15 = v15
        _ = fi16
    }
    var v17 float32 = 1.5
    if (!(-32768 == 1)) {
        var v18 string = "hello"
        var v19 int32 = v2
        _ = v18
        _ = v19
    }
    v13 += 102
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
