package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint32, p1 int, p2 int16) int {
    var v4 string = "a"
    v5 := []int{(p1 & 289)}
    if ("a" <= v4) {
        var v6 bool = false
        fmt.Println(v6)
        for fi7 := 0; fi7 < 13; fi7++ {
            p1 = 1
            var v8 uint64 = 18446744073709551615
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (p1 / ((-182) | 1))
}
func genFn1() uint16 {
    switch ("" + ("hello" + "\n\t")) {
    case "\n\t":
        var v9 int = -673
        _ = v9
    }
    var v10 bool = false
    _ = v10
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = -2147483648
    v3 := make(map[string]string)
    var v11 uint64 = 42
    var v12 int64 = -9223372036854775808
    var v13 int8 = -128
    delete(v1, "\n\t")
    var v14 uint64 = 2
    v15 := GenS1{F0: 2}
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 uint32 = 65535
        switch ("" + "") {
        case "":
            v12 -= v12
        case "a":
            switch ("hello" + ("" + "\n\t")) {
            case "a":
                fmt.Println(v17, v13, fi16)
            case "foo bar":
                var v18 uint64 = v11
                _ = v18
            case "hello":
                _ = len(v3)
            default:
                var v19 int16 = -1
                _ = v19
            }
        default:
            var v20 float64 = float64((1e10 + 3.14159))
            _ = v20
        }
        v13 = v13
        _ = fi16
        _ = v17
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        v22 := GenS1{F0: v14}
        var v23 int16 = -1
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 int64 = v12
    var v25 int32 = 2147483647
    fmt.Println(v13)
    var v26 float64 = 1.5
    var v27 float32 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
