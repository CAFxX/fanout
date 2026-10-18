package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint16 }

func main() {
    var v1 int8 = 0
    v2 := make(map[uint32]int)
    var v3 int32 = -1
    v4 := []string{((((("foo bar" + "foo bar") + "foo bar") + "") + ("" + "foo bar")) + ""), "foo bar"}
    var v5 int32 = v3
    switch ("hello" + "foo bar") {
    case "":
        var v6 int = 9223372036854775807
        _ = v6
    }
    var v7 uint32 = 39
    v8 := GenS1{F0: 165}
    var v9 int = 479
    fmt.Println(v3, v9)
    switch (("a" + "") + (("" + ("a" + "foo bar")) + (("\n\t" + "foo bar") + ("foo bar" + "")))) {
    case "":
        v10 := v2[(v7 / ((107) | 1))]
        _ = v10
    case "a":
        fmt.Println(v5, v3, v9)
    case "hello":
        v9 -= 128
    }
    fmt.Println(v9, v9)
    if true {
        if (((79 < (v1 * v1)) || true) || ((-(v1)) != v1)) {
            if true {
                v11 := GenS2{F0: 1, F1: 128, F2: 65535}
                var v12 int = (v9 - (v9 & 4294967295))
                _ = v11
                _ = v12
            }
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v1 += (v1 << ((2) & 63))
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 uint32 = (v7 ^ 198)
            _ = fi13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = len(v4)
            switch 42 {
            case 1:
                var v17 float64 = 0.1
                _ = v17
            case 248:
                var v18 string = "hello"
                _ = v18
            default:
                if len(v4) > 0 {
                    v4[0] = "a"
                }
            }
            _ = fi16
        }
        _ = fi15
    }
    delete(v2, v7)
    var v19 uint32 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
