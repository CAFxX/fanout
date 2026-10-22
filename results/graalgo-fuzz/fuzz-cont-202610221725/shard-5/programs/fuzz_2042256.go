package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int8, p1 int16) int {
    var v6 int64 = 0
    v7 := GenS1{F0: v6}
    var v8 uint = 720
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: 673}
        var v11 int = fi9
        v11 = 1
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 328
}
func genFn1() bool {
    if true {
        var v12 float64 = float64(1e10)
        var v13 int = 32768
        v14 := GenS1{F0: 657}
        _ = v12
        _ = v13
        _ = v14
    }
    v15 := []int{genFn0(100, -504), 32768, -694}
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := GenS2{F0: ""}
        switch (("\n\t" + "hello") + "\n\t") {
        case "foo bar":
            var v18 uint64 = 822
            _ = v18
        case "hello":
            var v19 int64 = -9223372036854775808
            _ = v19
        default:
            v20 := GenS2{F0: ""}
            _ = v20
        }
        var v21 uint16 = 65535
        _ = fi16
        _ = v17
        _ = v21
    }
    var v22 uint16 = 2
    _ = v15
    _ = v22
    return (!(true || false))
}

func main() {
    var v1 string = ((("\n\t" + "foo bar") + ("hello" + ("" + (("foo bar" + "foo bar") + ("hello" + ""))))) + "a")
    var v2 uint16 = 65535
    v3 := make(map[uint32]int)
    var v4 uint32 = 563
    var v5 uint8 = 147
    var v23 bool = (false || false)
    if true {
        fmt.Println(v1, v2, v23)
        var v24 string = (v1 + "\n\t")
        _ = v24
    } else {
        var v25 float64 = 1e10
        v26 := GenS2{F0: ("foo bar" + "hello")}
        _ = v25
        _ = v26
    }
    v27 := []int{}
    fmt.Println(v5, v1, v2)
    v1 = v1
    v28 := GenS3{F0: 636, F1: 255}
    var v29 int32 = -370
    v5 = (v5 % ((0) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
