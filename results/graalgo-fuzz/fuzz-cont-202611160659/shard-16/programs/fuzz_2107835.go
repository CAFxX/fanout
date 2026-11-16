package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint8; F2 string }
type GenS3 struct { F0 uint8; F1 string }

func main() {
    var v1 uint8 = 237
    v2 := make(map[string]bool)
    var v3 int32 = 1
    var v4 int64 = 0
    v5 := []int{128, 9223372036854775807, -699}
    if false {
        v4 = (v4 / ((v4) | 1))
        v1 += (v1 - v1)
        v3 += v3
    }
    v3 += v3
    v6 := GenS1{F0: 0}
    var v7 float32 = -2.25
    if ("" > "foo bar") {
        v8 := GenS1{F0: 127}
        switch "a" {
        case "\n\t":
            if false {
                if (((v7 < v7) || (1 == 683)) || true) {
                    var v9 uint32 = 217
                    v10 := GenS2{F0: 0, F1: 41, F2: "a"}
                    _ = v10
                    _ = v9
                }
            }
        case "hello":
            switch "\n\t" {
            case "foo bar":
                v4 = (-(v4))
            case "hello":
                for fi11 := 0; fi11 < 19; fi11++ {
                    var v12 uint16 = 65535
                    v12 = 1
                    _ = fi11
                    _ = v12
                }
            case "a":
                v3 = 931
            }
        default:
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 uint16 = 745
                _ = fi13
                _ = v14
            }
        }
        if true {
            var v15 int64 = (v4 - (v4 | v4))
            _ = v15
        }
        _ = v8
    }
    v3 -= v3
    for fi16 := 0; fi16 < 11; fi16++ {
        if true {
            var v17 int = -30
            _ = len(v2)
            if (!(!(242 >= 5))) {
                v18 := GenS3{F0: 42, F1: ("a" + "hello")}
                var v19 uint64 = 42
                _ = v18
                _ = v19
            }
            _ = v17
        } else {
            if false {
                fmt.Println(fi16, v3)
            }
        }
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
