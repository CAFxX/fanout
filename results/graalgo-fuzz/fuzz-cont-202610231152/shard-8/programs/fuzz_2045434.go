package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint }
func genFn0() uint32 {
    var v7 uint = 1
    var v8 int16 = 42
    switch "hello" {
    case "":
        var v9 uint16 = 1
        _ = v9
    case "hello":
        var v10 uint = v7
        _ = v10
    default:
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := []int{(fi11 << ((v7) & 63)), (fi11 / ((fi11) | 1))}
            v8 = 42
            _ = fi11
            _ = v12
        }
    }
    _ = v7
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 4294967295
    var v3 float32 = 3.14159
    var v4 int64 = 0
    var v5 uint = 4294967295
    v6 := []int{}
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint16 = 0
            v16 := GenS1{F0: (("" + "\n\t") + ((("foo bar" + "\n\t") + ("foo bar" + ("foo bar" + "\n\t"))) + "foo bar"))}
            _ = fi14
            _ = v15
            _ = v16
        }
        v2 += 42
        if true {
            v5 -= (v5 + 696)
            delete(v1, "foo bar")
        }
        _ = fi13
    }
    v3 -= float32(1.5)
    if len(v6) > 0 {
        v6[0] = 0
    }
    v17 := v1["\n\t"]
    for fi18 := 0; fi18 < 3; fi18++ {
        fmt.Println(v17, fi18, v17)
        var v19 uint16 = 302
        v19 += (v19 + v19)
        _ = fi18
        _ = v19
    }
    var v20 float64 = 1.5
    v21 := GenS1{F0: ""}
    fmt.Println(v4, v4)
    if (!(("\n\t" + "hello") < "hello")) {
        v4 = 0
        v17 -= 899
        v22 := GenS2{F0: (^(v2)), F1: 86, F2: (^(v2))}
        _ = v22
    }
    v23 := GenS2{F0: (v2 | v2), F1: 5, F2: 800}
    switch (v17 % (((v17 / (((-(v17))) | 1))) | 1)) {
    case -9223372036854775808:
        v24 := GenS3{F0: -525, F1: 0, F2: v5}
        _ = v24
    }
    var v25 uint = v5
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        fmt.Println(v4, v4, fi26)
        fmt.Println(v20)
        _ = fi26
    }
    switch -1 {
    case 314:
        var v27 int = (^(v17))
        _ = v27
    case 127:
        for fi28 := 0; fi28 < 20; fi28++ {
            var v29 uint8 = 65
            if true {
                if false {
                    v5 -= (^(v25))
                    v29 -= 128
                }
            }
            var v30 bool = true
            _ = fi28
            _ = v29
            _ = v30
        }
    case 9223372036854775807:
        _ = len(v1)
    default:
        var v31 string = "foo bar"
        _ = v31
    }
    var v32 bool = true
    var v33 uint16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
