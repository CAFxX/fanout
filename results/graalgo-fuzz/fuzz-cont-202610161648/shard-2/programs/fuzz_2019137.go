package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint8 }
func genFn0() string {
    var v6 uint32 = 2
    fmt.Println(v6)
    v7 := GenS1{F0: -1, F1: ((v6 >> ((2) & 63)) > (v6 >> ((1) & 63))), F2: 2147483647}
    _ = v6
    _ = v7
    return "foo bar"
}
func genFn1(p0 uint64) uint {
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return 4294967295
}

func main() {
    var v1 int = 128
    var v2 uint = 18446744073709551615
    var v3 int8 = -128
    v4 := make(map[int]string)
    var v5 string = ""
    v8 := GenS1{F0: 1000, F1: false, F2: 2}
    if false {
        delete(v4, v1)
    }
    v4[42] = v5
    var v9 uint = (v2 ^ v2)
    if (int32(v2) == 1) {
        v10 := GenS2{F0: 9223372036854775807, F1: 128, F2: 42}
        v11 := GenS2{F0: 233, F1: 255, F2: 92}
        fmt.Println(v1, v9, v1)
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: -1, F1: uint8(v9), F2: 216}
    v4[v1] = (v5 + v5)
    for fi13 := 0; fi13 < 13; fi13++ {
        switch ("\n\t" + (v5 + "foo bar")) {
        case "a":
            fmt.Println(v5)
        default:
            var v14 uint64 = 1
            _ = v14
        }
        v15 := []int{9223372036854775807}
        v5 = ""
        _ = fi13
        _ = v15
    }
    switch ((v5 + (((v5 + (v5 + v5)) + v5) + ("a" + "a"))) + v5) {
    case "hello":
        fmt.Println(v9, v2, v1)
    default:
        v16 := GenS1{F0: 1, F1: true, F2: 65535}
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 int16 = -32768
        v1 = (v1 & fi17)
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: 42, F1: false, F2: uint32(v2)}
    v20 := []int{v1, v1, (v1 + (v1 >> ((genFn1(571)) & 63)))}
    v4[v1] = (("\n\t" + ((v5 + "a") + "\n\t")) + v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
