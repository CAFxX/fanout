package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int16 }
func genFn0(p0 int) int8 {
    v7 := []int{p0}
    p0 -= (p0 / ((9223372036854775807) | 1))
    p0 = (^(p0))
    _ = p0
    _ = v7
    return -1
}

func main() {
    v1 := []string{("foo bar" + "hello")}
    var v2 uint = 271
    var v3 float32 = 1e10
    var v4 int64 = 703
    v5 := []int{836, -9223372036854775808, -681}
    var v6 uint64 = 18446744073709551615
    if len(v1) > 0 {
        v1[0] = ((("a" + "a") + "a") + "\n\t")
    }
    var v8 uint32 = 2147483647
    if false {
        for fi9 := 0; fi9 < 11; fi9++ {
            v4 += v4
            v10 := GenS1{F0: (((("foo bar" + "\n\t") + "\n\t") + ("hello" + ("foo bar" + ("a" + (("hello" + ("hello" + "a")) + "\n\t"))))) + "\n\t"), F1: "", F2: 42}
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: "\n\t", F1: ((("" + "") + ("\n\t" + "\n\t")) + "hello"), F2: 42}
        var v12 uint64 = uint64(v2)
        _ = v11
        _ = v12
    }
    var v13 float32 = float32(1e10)
    var v14 bool = false
    fmt.Println(v13, v13, v2)
    v15 := GenS1{F0: "foo bar", F1: "\n\t", F2: 0}
    var v16 uint32 = (^(v8))
    v17 := GenS1{F0: ("hello" + (("a" + ("hello" + ((("\n\t" + "hello") + "foo bar") + "hello"))) + ("a" + ("" + ("\n\t" + "a"))))), F1: "hello", F2: 32767}
    var v18 float64 = 0.1
    if ((v16 != 0) && v14) {
        var v19 uint32 = v16
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 int64 = (v4 | v4)
            _ = len(v1)
            _ = fi20
            _ = v21
        }
        v5 = append(v5, 32768)
        _ = v19
    }
    var v22 float32 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
