package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int }
func genFn0(p0 int, p1 uint) int {
    var v6 uint16 = 42
    v7 := GenS1{F0: 42, F1: ("" + (("foo bar" + ("a" + "a")) + "\n\t"))}
    v6 += (v6 << ((p1) & 63))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (^(p0))
}
func genFn1() uint8 {
    var v8 uint32 = 65535
    if (127 >= 128) {
        var v9 uint32 = v8
        var v10 int16 = 32767
        _ = v10
        _ = v9
    }
    _ = v8
    return 1
}

func main() {
    var v1 int8 = 101
    v2 := []int{2147483647}
    v3 := []uint32{2, 0, 2, 183}
    v4 := make(map[string]int)
    var v5 int16 = 525
    v1 -= (v1 | v1)
    if true {
        if true {
            var v11 float32 = 3.14159
            v12 := GenS2{F0: -503}
            var v13 bool = true
            _ = v11
            _ = v12
            _ = v13
        }
        delete(v4, "hello")
    }
    var v14 uint32 = 2147483648
    var v15 uint = 4294967295
    switch 1 {
    case 32767:
        v5 = 446
    case -786:
        switch "foo bar" {
        case "":
            delete(v4, "a")
        case "hello":
            v14 = (v14 | (v14 >> ((v15) & 63)))
        }
    case 42:
        for fi16 := 0; fi16 < 16; fi16++ {
            v4["foo bar"] = 1
            var v17 uint32 = v14
            var v18 int16 = v5
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    var v19 int32 = int32(v1)
    v5 -= (v5 - (-(v5)))
    var v20 uint = v15
    var v21 string = "a"
    var v22 uint8 = 127
    v20 += v20
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
