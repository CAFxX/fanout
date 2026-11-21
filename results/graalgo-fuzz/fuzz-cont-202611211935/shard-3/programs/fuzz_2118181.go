package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 string }
func genFn0() uint64 {
    var v5 float32 = 1.5
    var v6 bool = false
    v6 = false
    var v7 uint64 = 1
    _ = v5
    _ = v6
    _ = v7
    return (v7 << ((1) & 63))
}
func genFn1(p0 int64) uint64 {
    v8 := GenS1{F0: 2147483648}
    p0 += p0
    p0 += p0
    for fi9 := 0; fi9 < 4; fi9++ {
        if false {
            var v10 int64 = -509
            _ = v10
        }
        _ = fi9
    }
    var v11 int32 = 0
    _ = p0
    _ = v11
    _ = v8
    return 2
}

func main() {
    v1 := []string{"", ("hello" + ""), "", "foo bar"}
    var v2 bool = true
    var v3 int16 = 42
    v4 := make(map[int]uint32)
    var v12 int16 = (v3 + v3)
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint = 2
        if (!(!v2)) {
            var v15 int32 = -1
            _ = v15
        }
        switch fi13 {
        case 367:
            if true {
                v1 = append(v1, ("\n\t" + "foo bar"))
            } else {
                fmt.Println(v3)
                fmt.Println(fi13)
            }
        default:
            var v16 float32 = 1.5
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    var v17 bool = (-2147483648 == 188)
    if v2 {
        v18 := GenS2{F0: 0, F1: ("foo bar" + ("foo bar" + "\n\t"))}
        var v19 string = "a"
        v19 = v19
        _ = v18
        _ = v19
    }
    var v20 string = "\n\t"
    v3 += 826
    var v21 int64 = 0
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 16; fi23++ {
            switch ((("" + (("foo bar" + ((v20 + "\n\t") + v20)) + (("foo bar" + v20) + "\n\t"))) + v20) + v20) {
            case "foo bar":
                fmt.Println(fi22)
                _ = len(v4)
                var v24 uint8 = 1
                _ = v24
            case "\n\t":
                v12 = (^(v3))
            }
            v25 := GenS1{F0: 2147483647}
            switch v20 {
            case "foo bar":
                v26 := GenS1{F0: 71}
                _ = v26
            }
            _ = fi23
            _ = v25
        }
        _ = fi22
    }
    v12 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
