package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() uint16 {
    switch "foo bar" {
    case "foo bar":
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: -80}
            _ = fi7
            _ = v8
        }
    case "hello":
        var v9 int8 = 1
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 int8 = -128
            v11 = v11
            var v12 string = "a"
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 int = 9223372036854775807
    v14 := GenS1{F0: 42}
    var v15 string = "a"
    v15 = v15
    _ = v13
    _ = v14
    _ = v15
    return 59
}

func main() {
    var v1 bool = true
    v2 := []uint32{65535}
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    v5 := []uint32{2147483648, 4294967295, 966}
    v6 := make(map[int]int)
    for fi16 := 0; fi16 < 17; fi16++ {
        if ("" > "a") {
            fmt.Println(fi16, fi16)
            fmt.Println(v1, v1)
            for fi17 := 0; fi17 < 15; fi17++ {
                v1 = ("\n\t" < ("a" + "foo bar"))
                v1 = v1
                var v18 uint64 = 18446744073709551615
                _ = fi17
                _ = v18
            }
        }
        _ = fi16
    }
    if len(v5) > 0 {
        v5[0] = 0
    }
    var v19 int16 = 0
    for fi20 := 0; fi20 < 9; fi20++ {
        v2 = append(v2, 0)
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 float64 = 0.0
            var v23 float32 = float32(v22)
            var v24 bool = v1
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        _ = fi20
    }
    v19 = (v19 + (v19 << ((18446744073709551615) & 63)))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v25 := GenS1{F0: 946}
    v1 = (!true)
    if len(v2) > 0 {
        v2[0] = 2
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 int64 = 1
        fmt.Println(v1, v27)
        v28 := GenS1{F0: v27}
        _ = fi26
        _ = v27
        _ = v28
    }
    v19 = (v19 | 0)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
