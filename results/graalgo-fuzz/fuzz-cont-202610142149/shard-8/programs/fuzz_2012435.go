package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint32; F1 string }
func genFn0() uint16 {
    v7 := []int{42, 127, 32768}
    var v8 string = ((("hello" + "\n\t") + "hello") + (((("" + ("" + ("foo bar" + "\n\t"))) + "\n\t") + ("hello" + ("" + ""))) + "a"))
    v8 = v8
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    v3 := []string{"hello", "a"}
    v4 := []int{}
    v5 := []uint32{2147483647, 42}
    var v6 int64 = 1
    var v9 int = -9223372036854775808
    var v10 string = ("" + "a")
    v11 := GenS1{F0: 255}
    if true {
        v12 := GenS1{F0: 2}
        v1[v10] = ((((v10 + ("\n\t" + v10)) + ((v10 + "") + (v10 + (v10 + v10)))) + v10) + v10)
        _ = v12
    }
    _ = len(v4)
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "foo bar")
    }
    if ((!false) || (!((!true) || (2147483647 > 583)))) {
        var v13 uint64 = 2
        _ = v13
    } else {
        v14 := GenS2{F0: 1, F1: (v9 | (v9 * (v9 & v9)))}
        v15 := GenS3{F0: 4294967295, F1: "a"}
        _ = v14
        _ = v15
    }
    if true {
        var v16 uint16 = 2
        v17 := GenS1{F0: 809}
        _ = v16
        _ = v17
    }
    fmt.Println(v6, v10, v6)
    if len(v3) > 0 {
        v3[0] = v10
    }
    if (4294967295 > 42) {
        if len(v3) > 0 {
            v3[0] = v10
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
