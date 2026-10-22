package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 4294967295
    var v3 int = -1
    var v4 uint16 = 255
    if (2 < (v4 / ((v4) | 1))) {
        for fi5 := 0; fi5 < 8; fi5++ {
            v3 -= fi5
            for fi6 := 0; fi6 < 3; fi6++ {
                switch fi5 {
                case 988:
                    v3 -= v3
                case 1:
                    if false {
                        v7 := []int{fi6, fi6}
                        _ = v7
                    }
                }
                _ = fi6
            }
            _ = fi5
        }
        fmt.Println(v4)
        v8 := v1[127]
        _ = v8
    }
    v1[-1] = 2
    for fi9 := 0; fi9 < 3; fi9++ {
        v1[(fi9 >> ((uint(v4)) & 63))] = 65535
        _ = fi9
    }
    if ("" > (("hello" + "foo bar") + ("hello" + (("foo bar" + "") + (("" + ("hello" + "hello")) + "foo bar"))))) {
        switch "\n\t" {
        case "":
            var v10 string = ("hello" + "foo bar")
            _ = v10
        case "hello":
            var v11 uint64 = v2
            _ = v11
        }
        switch int(v4) {
        case -1:
            v12 := []int{9223372036854775807, (-(v3)), -409}
            _ = v12
        case 570:
            switch "" {
            case "\n\t":
                v4 -= 180
            case "hello":
                for fi13 := 0; fi13 < 16; fi13++ {
                    fmt.Println(fi13, v4)
                    _ = fi13
                }
            default:
                _ = len(v1)
            }
        }
    }
    v14 := GenS1{F0: 998, F1: -845, F2: uint8(v4)}
    switch (v3 + v3) {
    case 128:
        if (!(false || false)) {
            var v15 float64 = (3.14159 - 0.1)
            _ = v15
        }
    default:
        var v16 uint64 = (^(v2))
        _ = v16
    }
    fmt.Println(v2)
    v17 := []int{}
    var v18 int = -1
    v19 := GenS1{F0: 65535, F1: -9223372036854775808, F2: 255}
    delete(v1, 127)
    _ = len(v1)
    var v20 uint16 = 32767
    v2 = 4294967295
    v21 := v1[v3]
    v3 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
