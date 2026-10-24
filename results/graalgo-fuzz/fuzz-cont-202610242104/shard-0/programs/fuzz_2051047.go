package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint16; F2 string }
func genFn0() uint32 {
    var v5 uint64 = 18446744073709551615
    v6 := []int{-2147483648, 1, -2147483648}
    var v7 uint16 = 65535
    var v8 int = 899
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS1{F0: "", F1: v7}
        _ = fi9
        _ = v10
    }
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]bool)
    var v3 float64 = ((1e10 - -2.25) - 0.1)
    v4 := make(map[string]bool)
    v11 := GenS1{F0: v1, F1: 42}
    var v12 string = ((v1 + ("\n\t" + v1)) + "")
    if true {
        v13 := GenS2{F0: 1, F1: 471, F2: "hello"}
        var v14 uint8 = 255
        for fi15 := 0; fi15 < 5; fi15++ {
            v1 = (v1 + (v1 + v1))
            v16 := []int{fi15, fi15}
            _ = fi15
            _ = v16
        }
        _ = v13
        _ = v14
    } else {
        v17 := v4["a"]
        _ = v17
    }
    v1 = "\n\t"
    if true {
        v12 = "hello"
        for fi18 := 0; fi18 < 7; fi18++ {
            fmt.Println(v1, v12, v1)
            var v19 uint = 65535
            v20 := []int{-255, fi18}
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v3 += ((((v3 + ((v3 * (0.1 + float64(1.5))) * v3)) - float64(1.5)) - 0.1) - 0.1)
    switch v1 {
    case "\n\t":
        delete(v2, "a")
    case "foo bar":
        v21 := []int{109, -9223372036854775808, 159}
        _ = v21
    }
    v3 = (v3 + (v3 - (v3 / v3)))
    v22 := []int{-402, -1}
    v23 := GenS2{F0: -146, F1: 403, F2: "foo bar"}
    delete(v4, v12)
    _ = len(v22)
    var v24 uint16 = 484
    v24 = v24
    var v25 int64 = -1
    v12 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
