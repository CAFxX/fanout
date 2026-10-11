package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0(p0 uint) int16 {
    switch ((("\n\t" + "a") + (("foo bar" + "foo bar") + "")) + ("" + "foo bar")) {
    case "":
        var v7 bool = false
        _ = v7
    case "\n\t":
        p0 -= 0
    }
    if false {
        if false {
            var v8 uint8 = 1
            _ = v8
        }
    }
    if (127 == -1) {
        p0 = (p0 << (((p0 >> (((p0 | 4294967295)) & 63))) & 63))
    } else {
        v9 := GenS1{F0: 0, F1: "a"}
        _ = v9
    }
    _ = p0
    return 936
}
func genFn1(p0 int8, p1 uint32) uint {
    var v10 int64 = 1
    fmt.Println(p1)
    switch 42 {
    case 9223372036854775807:
        p0 = 1
    case 128:
        for fi11 := 0; fi11 < 9; fi11++ {
            fmt.Println(v10, p0, p1)
            _ = fi11
        }
    default:
        switch ("a" + "\n\t") {
        case "a":
            p1 = (p1 << ((18446744073709551615) & 63))
        default:
            var v12 uint16 = 42
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v10
    return 255
}

func main() {
    v1 := []string{("a" + "a"), "", "hello", ("\n\t" + ("a" + "a"))}
    var v2 string = ("\n\t" + "\n\t")
    v3 := []int{}
    v4 := []string{v2, ""}
    v5 := make(map[string]uint32)
    v6 := make(map[string]int)
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 uint16 = 584
        fmt.Println(v14, v14, v2)
        _ = fi13
        _ = v14
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    if false {
        fmt.Println(v2)
    } else {
        fmt.Println(v2)
        fmt.Println(v2)
        switch 1 {
        case 483:
            delete(v6, (v2 + v2))
        case 733:
            delete(v6, "")
        default:
            var v15 float64 = float64(1e10)
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        v2 = (v2 + "foo bar")
        _ = fi16
    }
    _ = len(v6)
    v2 = ""
    if ((42 != 720) || (491 > 0)) {
        fmt.Println(v2)
    } else {
        fmt.Println(v2)
        var v17 uint = 274
        _ = v17
    }
    fmt.Println(v2)
    v18 := v6[""]
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 int32 = -277
        _ = fi19
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
