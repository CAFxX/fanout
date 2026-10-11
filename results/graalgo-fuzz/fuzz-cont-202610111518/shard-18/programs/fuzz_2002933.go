package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int; F2 int64 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = true
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    var v5 uint64 = 2
    var v6 int = 32767
    var v7 int16 = 42
    v8 := GenS1{F0: ""}
    for fi9 := 0; fi9 < 20; fi9++ {
        v2 = (v2 && (11 >= 255))
        _ = len(v3)
        v7 -= 42
        _ = fi9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        switch 127 {
        case 128:
            v5 = v5
        }
        _ = fi10
    }
    v6 = (-(v6))
    v7 = v7
    var v11 int32 = 42
    if v2 {
        v12 := GenS2{F0: 90, F1: v6, F2: -815}
        _ = v12
    } else {
        v2 = v2
        switch "foo bar" {
        case "foo bar":
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 int16 = 42
                v15 := GenS1{F0: "a"}
                v16 := []int{0, fi13}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        }
        var v17 int = v6
        _ = v17
    }
    if true {
        v18 := []int{v6}
        fmt.Println(v7)
        _ = v18
    } else {
        var v19 int32 = (v11 + 42)
        if false {
            var v20 float64 = 0.1
            for fi21 := 0; fi21 < 11; fi21++ {
                v7 -= 42
                delete(v1, "foo bar")
                _ = fi21
            }
            if (!v2) {
                v22 := GenS2{F0: v5, F1: v6, F2: 42}
                v20 = v20
                _ = v22
            }
            _ = v20
        } else {
            fmt.Println(v7, v7, v6)
            fmt.Println(v5)
            delete(v3, "a")
        }
        v6 = (v6 % (((v6 & (v6 >> ((66) & 63)))) | 1))
        _ = v19
    }
    v23 := GenS3{F0: ("\n\t" + "foo bar")}
    var v24 float32 = (1.5 / 1e10)
    var v25 uint64 = 18446744073709551615
    for fi26 := 0; fi26 < 12; fi26++ {
        v27 := v1["foo bar"]
        _ = fi26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
