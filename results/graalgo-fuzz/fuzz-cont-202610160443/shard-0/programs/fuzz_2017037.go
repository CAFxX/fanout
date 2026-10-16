package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0(p0 int) uint16 {
    for fi4 := 0; fi4 < 18; fi4++ {
        for fi5 := 0; fi5 < 15; fi5++ {
            switch (^(fi4)) {
            case 32767:
                var v6 string = (("\n\t" + ("a" + "foo bar")) + "foo bar")
                _ = v6
            }
            if true {
                var v7 int64 = -1
                var v8 uint32 = 2147483647
                _ = v7
                _ = v8
            }
            _ = fi5
        }
        v9 := []int{}
        _ = fi4
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := []int{}
        _ = fi10
        _ = v11
    }
    _ = p0
    return 157
}
func genFn1(p0 int64) int16 {
    v12 := GenS1{F0: 65535, F1: 100}
    var v13 float32 = (float32((1.5 * -2.25)) + 1e10)
    _ = p0
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := []string{"a", "\n\t", ("foo bar" + ("a" + ("hello" + "\n\t")))}
    v2 := make(map[int]uint32)
    var v3 string = ("foo bar" + "\n\t")
    v1 = append(v1, "hello")
    for fi14 := 0; fi14 < 18; fi14++ {
        fmt.Println(fi14, fi14)
        var v15 uint = 2
        _ = fi14
        _ = v15
    }
    switch v3 {
    case "foo bar":
        var v16 uint = 255
        _ = v16
    case "a":
        _ = len(v1)
    case "":
        v2[1] = 2
    default:
        v3 = v3
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        v18 := GenS1{F0: 0, F1: 0}
        var v19 uint8 = 128
        var v20 float32 = 0.1
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    _ = len(v1)
    var v21 float32 = (3.14159 / (1.5 - (0.1 / 1.0)))
    fmt.Println(v3)
    var v22 int32 = -1
    var v23 uint32 = 0
    var v24 uint = 18446744073709551615
    v25 := GenS1{F0: v23, F1: 52}
    v3 = (v3 + v3)
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
