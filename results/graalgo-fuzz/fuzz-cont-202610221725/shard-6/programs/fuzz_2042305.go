package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint8; F1 string }
func genFn0(p0 int8, p1 uint, p2 uint) uint8 {
    fmt.Println(p0, p2, p0)
    switch ((("\n\t" + ("hello" + "\n\t")) + "\n\t") + "hello") {
    case "hello":
        for fi6 := 0; fi6 < 15; fi6++ {
            fmt.Println(p0, fi6)
            var v7 uint64 = 383
            _ = fi6
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 uint8 = 45
    var v2 bool = false
    v3 := make(map[int]int)
    v4 := []string{"foo bar"}
    v5 := make(map[string]uint32)
    var v8 uint32 = 800
    v8 = v8
    v9 := GenS1{F0: "a", F1: (v8 % ((v8) | 1))}
    if false {
        for fi10 := 0; fi10 < 9; fi10++ {
            v8 -= v8
            var v11 uint = 42
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 uint16 = 65535
                v14 := GenS2{F0: (v1 >> (((v11 & 2)) & 63)), F1: (("" + (("foo bar" + ("\n\t" + "\n\t")) + ("hello" + "\n\t"))) + "\n\t")}
                v15 := v5["hello"]
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi10
            _ = v11
        }
    } else {
        v2 = (-2147483648 < 128)
        var v16 int64 = -980
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        v2 = v2
        v8 = v8
        var v18 uint32 = v8
        _ = fi17
        _ = v18
    }
    var v19 string = (("\n\t" + "a") + "foo bar")
    for fi20 := 0; fi20 < 9; fi20++ {
        switch ((v19 + v19) + v19) {
        case "foo bar":
            var v21 int16 = -837
            _ = v21
        case "a":
            fmt.Println(fi20)
        case "":
            var v22 float32 = (-2.25 / 1.0)
            _ = v22
        default:
            v4 = append(v4, v19)
        }
        v1 = 1
        for fi23 := 0; fi23 < 15; fi23++ {
            fmt.Println(v19)
            fmt.Println(v8)
            var v24 uint64 = 42
            _ = fi23
            _ = v24
        }
        _ = fi20
    }
    switch v19 {
    case "foo bar":
        fmt.Println(v19)
    default:
        if v2 {
            if false {
                var v25 int16 = -135
                _ = v25
            }
            v3[-9223372036854775808] = 1
            fmt.Println(v19)
        }
    }
    v1 += (v1 >> ((0) & 63))
    v26 := GenS1{F0: v19, F1: v8}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
