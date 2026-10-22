package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int }
func genFn0(p0 string, p1 int32, p2 bool) bool {
    p0 = (p0 + p0)
    v7 := GenS1{F0: 1, F1: 248}
    v8 := GenS2{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ((p2 && false) || p2)
}

func main() {
    v1 := []string{"a"}
    v2 := make(map[string]string)
    v3 := []int{32768}
    v4 := []string{"hello", "hello"}
    v5 := make(map[uint32]string)
    var v6 uint8 = 1
    var v9 float64 = float64((1e10 - 1.5))
    var v10 uint8 = 255
    v11 := GenS1{F0: 192, F1: 42}
    if false {
        var v12 uint64 = 18446744073709551615
        v13 := GenS2{F0: 65535}
        fmt.Println(v9, v12)
        _ = v12
        _ = v13
    }
    v3 = append(v3, 4)
    switch ("hello" + "hello") {
    case "":
        v14 := GenS3{F0: 1, F1: 978}
        _ = v14
    }
    v15 := GenS3{F0: 207, F1: -2147483648}
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 string = ((("" + "a") + "\n\t") + "")
        var v18 bool = (1 == 1)
        _ = fi16
        _ = v17
        _ = v18
    }
    switch ("\n\t" + "") {
    case "a":
        v10 = (v10 + 39)
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        v19 := GenS1{F0: 1, F1: 1}
        _ = v19
    }
    var v20 uint64 = 1
    switch 1 {
    case -2147483648:
        if false {
            var v21 int32 = -306
            _ = v21
        }
    default:
        var v22 int8 = 0
        _ = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
