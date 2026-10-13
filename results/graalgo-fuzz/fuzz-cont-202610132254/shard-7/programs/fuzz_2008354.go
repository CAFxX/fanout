package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int; F2 string }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint64, p1 int, p2 int16) uint32 {
    v6 := GenS1{F0: "foo bar"}
    var v7 float32 = 1.5
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 int8 = 0
    var v2 int8 = (v1 >> ((42) & 63))
    var v3 float64 = -2.25
    v4 := []string{"foo bar", ("" + "foo bar"), "foo bar", ""}
    var v5 float64 = 3.14159
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint16 = 183
        var v10 int = (^(fi8))
        _ = fi8
        _ = v10
        _ = v9
    }
    if (!(!(!false))) {
        v11 := GenS2{F0: 2147483647, F1: 502, F2: "hello"}
        if true {
            for fi12 := 0; fi12 < 13; fi12++ {
                v4 = append(v4, ("" + "hello"))
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 uint64 = 577
                    fmt.Println(fi12, v2, v14)
                    fmt.Println(v5)
                    _ = fi13
                    _ = v14
                }
                var v15 float32 = -2.25
                _ = fi12
                _ = v15
            }
            switch "" {
            case "":
                v1 -= v1
            }
            if len(v4) > 0 {
                v4[0] = (("\n\t" + ("\n\t" + "a")) + (("foo bar" + "foo bar") + "a"))
            }
        }
        var v16 uint32 = 4294967295
        _ = v11
        _ = v16
    } else {
        v17 := GenS3{F0: 1, F1: 854}
        var v18 int8 = (v2 ^ v1)
        v3 += 3.14159
        _ = v17
        _ = v18
    }
    var v19 uint = 4294967295
    var v20 int16 = 0
    fmt.Println(v20, v2)
    for fi21 := 0; fi21 < 8; fi21++ {
        if true {
            var v22 int16 = 1000
            if true {
                var v23 float32 = (0.0 + (0.1 - 0.0))
                _ = v23
            } else {
                fmt.Println(v2, v1)
            }
            v19 = (^(v19))
            _ = v22
        }
        _ = fi21
    }
    v24 := GenS2{F0: 4294967295, F1: 2147483647, F2: ("\n\t" + "foo bar")}
    v25 := GenS4{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
