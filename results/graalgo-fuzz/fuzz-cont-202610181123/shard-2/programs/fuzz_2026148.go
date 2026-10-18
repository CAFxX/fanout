package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
func genFn0() uint64 {
    var v6 uint = 332
    v7 := GenS1{F0: 897, F1: 1}
    switch "" {
    case "foo bar":
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 float64 = 1e10
            var v10 uint = (v6 / (((v6 % ((v6) | 1))) | 1))
            _ = fi8
            _ = v10
            _ = v9
        }
    case "\n\t":
        var v11 uint64 = 4294967295
        _ = v11
    default:
        var v12 int = 543
        _ = v12
    }
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 uint16 = 0
    v2 := []string{"", (("hello" + ("foo bar" + "")) + ("a" + ("\n\t" + ("\n\t" + "a")))), "\n\t"}
    v3 := make(map[string]string)
    var v4 uint = 65535
    var v5 uint8 = 42
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := GenS1{F0: 0, F1: 0}
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: 100}
    var v16 int64 = 559
    var v17 float32 = (0.1 / (1.5 + (1e10 * (float32(0.1) + (0.1 + -2.25)))))
    var v18 int32 = -1
    fmt.Println(v17, v18, v17)
    for fi19 := 0; fi19 < 19; fi19++ {
        v17 = v17
        var v20 uint64 = genFn0()
        v21 := GenS3{F0: 80}
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        switch int(v5) {
        case -2147483648:
            fmt.Println(fi22)
        }
        _ = fi22
    }
    v23 := GenS2{F0: -34}
    if false {
        v2 = append(v2, (("" + ((("hello" + "\n\t") + "") + "a")) + ""))
        v24 := GenS2{F0: -128}
        v4 -= v4
        _ = v24
    }
    var v25 int32 = (^(v18))
    fmt.Println(v5, v1)
    if (((("foo bar" + "\n\t") != "a") && false) || false) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v1, v1)
    fmt.Println(v25, v5, v4)
    v17 += (v17 - -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
