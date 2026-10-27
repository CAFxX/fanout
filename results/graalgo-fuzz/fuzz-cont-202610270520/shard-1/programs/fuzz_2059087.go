package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int8 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := []string{("a" + "hello"), (("hello" + "\n\t") + "a"), "hello", ("" + "hello")}
    var v2 bool = ((42 >= 42) && true)
    v3 := []string{"\n\t", "\n\t", "foo bar"}
    v4 := make(map[int]string)
    var v5 uint8 = 1
    v6 := GenS1{F0: 100, F1: 2, F2: -128}
    if len(v1) > 0 {
        v1[0] = ""
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int32 = -2147483648
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: 42, F1: 0}
    var v10 string = ("foo bar" + "hello")
    v11 := GenS3{F0: 684, F1: -994}
    var v12 uint32 = 316
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v13 int64 = 785
    v10 = v10
    switch 32767 {
    case -813:
        v14 := GenS1{F0: 100, F1: v12, F2: 0}
        _ = v14
    case 209:
        var v15 int32 = 1
        _ = v15
    case -547:
        v12 += 567
    default:
        v10 = (v10 + "foo bar")
    }
    v5 = (^(v5))
    delete(v4, -64)
    v16 := GenS4{F0: 42, F1: (v12 * uint32(v5))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
