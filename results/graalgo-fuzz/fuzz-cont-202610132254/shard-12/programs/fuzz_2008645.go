package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = ((((0.0 * 1.5) + (0.1 * 0.1)) + float32((float64(-2.25) / 1.0))) + 1e10)
    v3 := []string{"hello", (("a" + (("" + ("a" + "\n\t")) + (("" + "\n\t") + (("" + (("hello" + ("" + ("\n\t" + "hello"))) + "a")) + "hello")))) + (("" + "foo bar") + "\n\t"))}
    v4 := make(map[int]string)
    v2 = 0.0
    v2 -= v2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch (("a" + ("hello" + "\n\t")) + "a") {
    case "hello":
        var v5 uint32 = 1
        _ = v5
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        fmt.Println(fi6)
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(v2)
            _ = fi7
        }
        _ = fi6
    }
    if true {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint8 = 2
            var v10 int = (fi8 >> ((24) & 63))
            _ = fi8
            _ = v10
            _ = v9
        }
        v2 = v2
        v11 := GenS1{F0: 42}
        _ = v11
    } else {
        v2 = 3.14159
        var v12 int8 = 100
        _ = v12
    }
    v2 -= v2
    v2 -= v2
    v13 := GenS1{F0: -1}
    v2 = v2
    v4[-945] = (("hello" + "") + "\n\t")
    var v14 int = 127
    v15 := GenS2{F0: 65535, F1: 65535}
    for fi16 := 0; fi16 < 1; fi16++ {
        switch ("\n\t" + "a") {
        case "hello":
            v2 += (v2 - (v2 * float32(-2.25)))
        case "":
            for fi17 := 0; fi17 < 10; fi17++ {
                fmt.Println(v2, v2, v2)
                _ = fi17
            }
        case "foo bar":
            v14 = (v14 - (v14 & (v14 ^ (v14 & fi16))))
        default:
            delete(v4, v14)
        }
        v18 := GenS1{F0: -902}
        _ = fi16
        _ = v18
    }
    v2 -= 3.14159
    v14 += v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
