package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v4 int = -347
    v5 := []uint32{2, 4294967295}
    v1[213] = 2
    var v6 int32 = 1
    if (true || (1000 < 42)) {
        delete(v3, 184)
        v7 := GenS1{F0: (!true), F1: (!(v6 >= (v6 ^ v6)))}
        _ = v7
    }
    var v8 uint8 = 83
    fmt.Println(v8, v4)
    v6 += -651
    switch (("\n\t" + "") + ("foo bar" + ("" + ("" + ("a" + ("a" + "a")))))) {
    case "foo bar":
        v9 := GenS2{F0: 1}
        _ = v9
    case "a":
        if ("hello" == "\n\t") {
            switch "hello" {
            case "\n\t":
                v5 = append(v5, 65535)
            case "":
                var v10 uint32 = 4294967295
                _ = v10
            case "a":
                var v11 uint64 = 198
                _ = v11
            }
            switch "\n\t" {
            case "\n\t":
                switch v4 {
                case -725:
                    _ = len(v5)
                case 71:
                    var v12 uint32 = 990
                    _ = v12
                }
            case "":
                var v13 int = (v4 % (((^(v4))) | 1))
                _ = v13
            }
        } else {
            fmt.Println(v4, v4, v4)
        }
    case "":
        v4 -= v4
    }
    v14 := GenS1{F0: (65535 == 763), F1: (v4 != v4)}
    fmt.Println(v6)
    v15 := v1[4294967295]
    fmt.Println(v4)
    for fi16 := 0; fi16 < 14; fi16++ {
        v17 := GenS3{F0: 127, F1: 65535}
        var v18 uint32 = v15
        var v19 float64 = 1.5
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 int16 = 32767
    v8 = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
