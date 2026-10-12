package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int64 }
type GenS2 struct { F0 uint64; F1 string; F2 uint64 }
func genFn0(p0 int64) uint64 {
    if (!false) {
        v5 := []int{32768, 2147483647, 32767}
        _ = v5
    }
    fmt.Println(p0)
    if (42 == 32767) {
        if false {
            v6 := GenS1{F0: -109, F1: 95, F2: p0}
            fmt.Println(p0)
            var v7 int = -736
            _ = v6
            _ = v7
        } else {
            var v8 float32 = (float32((3.14159 - float64(0.1))) / (0.1 * float32(1.5)))
            fmt.Println(p0, p0)
            _ = v8
        }
        v9 := []int{2147483647, 9223372036854775807, 1, -301}
        _ = v9
    } else {
        v10 := []int{}
        fmt.Println(p0)
        if true {
            v11 := GenS1{F0: -45, F1: 2, F2: (p0 >> ((4294967295) & 63))}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 9; fi12++ {
                p0 = (p0 % (((-(p0))) | 1))
                v13 := GenS2{F0: 588, F1: ("a" + (((("\n\t" + "") + "hello") + ("foo bar" + ((("foo bar" + "a") + (("a" + ("foo bar" + ((("hello" + "\n\t") + "foo bar") + ""))) + "\n\t")) + "\n\t"))) + ("a" + "foo bar"))), F2: 42}
                _ = fi12
                _ = v13
            }
        }
        _ = v10
    }
    _ = p0
    return 976
}
func genFn1(p0 uint32, p1 uint64, p2 uint32) bool {
    var v14 int = 4294967295
    for fi15 := 0; fi15 < 15; fi15++ {
        switch "" {
        case "\n\t":
            var v16 uint64 = 18446744073709551615
            _ = v16
        case "foo bar":
            v17 := []int{2147483647, 697}
            _ = v17
        case "a":
            var v18 int8 = -46
            _ = v18
        default:
            var v19 uint32 = (p2 & 2147483648)
            _ = v19
        }
        var v20 float64 = 3.14159
        v21 := GenS1{F0: 2147483647, F1: 65535, F2: 1}
        _ = fi15
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (true || ((42 != 732) || true))
}

func main() {
    v1 := make(map[int]string)
    var v2 int = -9223372036854775808
    v3 := []string{"\n\t", ""}
    v4 := make(map[int]int)
    var v22 int16 = 0
    v23 := GenS2{F0: 724, F1: "foo bar", F2: 2}
    var v24 float64 = 0.0
    var v25 float32 = 0.0
    fmt.Println(v2, v25, v24)
    if (v24 > 3.14159) {
        var v26 uint32 = 235
        _ = v26
    }
    if len(v3) > 0 {
        v3[0] = (("\n\t" + "\n\t") + "a")
    }
    var v27 uint32 = 2147483648
    for fi28 := 0; fi28 < 14; fi28++ {
        var v29 uint8 = 1
        _ = fi28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
}
