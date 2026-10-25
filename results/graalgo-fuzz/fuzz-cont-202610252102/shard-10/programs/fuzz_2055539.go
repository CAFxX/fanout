package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
func genFn0(p0 uint8) int64 {
    p0 = 2
    var v4 uint64 = 0
    for fi5 := 0; fi5 < 13; fi5++ {
        v4 = (v4 & v4)
        v6 := GenS1{F0: (fi5 & 9223372036854775807), F1: "a"}
        v7 := GenS1{F0: 128, F1: ""}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int16 = 902
    _ = p0
    _ = v4
    _ = v8
    return -1
}
func genFn1(p0 uint32, p1 int32, p2 bool) uint64 {
    var v9 uint16 = 313
    if (((p1 & -1) != (p1 | p1)) && (-32768 < 1000)) {
        v10 := []int{127, 4294967295}
        v9 = (v9 >> ((255) & 63))
        for fi11 := 0; fi11 < 1; fi11++ {
            p1 = 1
            if len(v10) > 0 {
                v10[0] = (^(fi11))
            }
            _ = fi11
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 int64 = -145
    v3 = (v3 >> ((255) & 63))
    var v12 int = 32767
    if false {
        v13 := []int{(-(v12))}
        var v14 uint64 = 2
        delete(v1, (("a" + "foo bar") + ("\n\t" + "\n\t")))
        _ = v13
        _ = v14
    }
    fmt.Println(v3)
    v15 := GenS1{F0: v12, F1: "foo bar"}
    v16 := GenS1{F0: v12, F1: ""}
    var v17 int = (^(v12))
    var v18 int16 = 42
    for fi19 := 0; fi19 < 13; fi19++ {
        switch "" {
        case "":
            v20 := GenS1{F0: fi19, F1: "\n\t"}
            _ = v20
        case "hello":
            var v21 int64 = (v3 - genFn0(65))
            _ = v21
        case "\n\t":
            v17 += (fi19 << ((899) & 63))
        default:
            var v22 uint16 = 255
            _ = v22
        }
        _ = fi19
    }
    v3 += (v3 & genFn0(255))
    v23 := GenS1{F0: v17, F1: "hello"}
    var v24 uint8 = 115
    var v25 int32 = 2147483647
    var v26 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
