package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint = 42
    v2 := []string{}
    var v3 int = -827
    v4 := make(map[string]bool)
    v5 := make(map[uint32]uint32)
    switch (("a" + "") + "foo bar") {
    case "\n\t":
        switch (-(v3)) {
        case 32767:
            v6 := GenS1{F0: 4294967295}
            _ = v6
        case -9223372036854775808:
            _ = len(v5)
        case 4294967295:
            v7 := GenS1{F0: 2147483648}
            _ = v7
        default:
            v8 := GenS1{F0: 1}
            _ = v8
        }
    case "hello":
        v9 := GenS1{F0: 652}
        _ = v9
    case "foo bar":
        v3 -= v3
    default:
        switch "hello" {
        case "a":
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 int32 = -450
                v12 := GenS2{F0: 41}
                v13 := v4["hello"]
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
        default:
            delete(v4, ("foo bar" + "foo bar"))
        }
    }
    var v14 int = v3
    fmt.Println(v1, v3, v3)
    var v15 int = -1
    switch "foo bar" {
    case "":
        if true {
            var v16 bool = true
            _ = v16
        } else {
            v17 := GenS1{F0: 468}
            v14 -= v3
            var v18 string = "a"
            _ = v17
            _ = v18
        }
    case "\n\t":
        v19 := GenS2{F0: v14}
        _ = v19
    }
    switch v15 {
    case 963:
        for fi20 := 0; fi20 < 1; fi20++ {
            for fi21 := 0; fi21 < 18; fi21++ {
                if (463 <= 724) {
                    v22 := GenS3{F0: 210}
                    _ = v22
                }
                _ = fi21
            }
            _ = fi20
        }
    default:
        for fi23 := 0; fi23 < 10; fi23++ {
            v14 += v15
            fmt.Println(v1)
            _ = fi23
        }
    }
    v3 = v3
    v24 := GenS4{F0: (0 > 100)}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
