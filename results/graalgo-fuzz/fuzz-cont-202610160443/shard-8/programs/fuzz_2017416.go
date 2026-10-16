package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint16, p1 uint16) uint16 {
    switch "foo bar" {
    case "\n\t":
        var v6 int32 = 1
        _ = v6
    case "":
        switch 128 {
        case 2147483647:
            p0 = 65535
        case 0:
            v7 := []int{}
            _ = v7
        }
    }
    var v8 string = (("foo bar" + "a") + "a")
    var v9 int32 = 42
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []string{"", "a", (("a" + ("hello" + ("a" + ""))) + "hello"), ("" + "a")}
    v2 := []string{"hello"}
    var v3 int32 = -2147483648
    var v4 int32 = v3
    v5 := make(map[uint32]int)
    fmt.Println(v3, v4)
    var v10 int64 = 42
    var v11 int32 = 0
    v12 := v5[2147483648]
    var v13 int32 = (v11 * v4)
    v3 = (v4 >> ((1) & 63))
    if true {
        v14 := GenS1{F0: 127, F1: 1, F2: 2}
        _ = v14
    } else {
        fmt.Println(v4, v11, v13)
    }
    v15 := v5[287]
    var v16 uint8 = 158
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 float32 = 1.5
        _ = fi17
        _ = v18
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v10 = 1
    var v19 int = (v15 / (((v15 << ((1) & 63))) | 1))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        v21 := GenS2{F0: false, F1: (v16 / (((v16 | v16)) | 1))}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi20
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
