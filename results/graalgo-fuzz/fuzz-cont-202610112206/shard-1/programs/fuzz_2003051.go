package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint32 }
func genFn0() int {
    var v7 int16 = 1000
    var v8 int64 = 1
    switch "a" {
    case "\n\t":
        fmt.Println(v8)
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        v8 = v8
        fmt.Println(v7, fi9, v8)
        _ = fi9
    }
    _ = v7
    _ = v8
    return -970
}
func genFn1(p0 uint16, p1 uint32, p2 string) string {
    var v10 int8 = -128
    var v11 int16 = -32768
    p1 += (p1 / (((^(p1))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return ("a" + p2)
}

func main() {
    var v1 uint32 = 65535
    v2 := []string{("\n\t" + "hello"), (("foo bar" + "foo bar") + ((("hello" + "\n\t") + "hello") + "a"))}
    v3 := make(map[string]int)
    var v4 int16 = 1
    var v5 int64 = -1
    v6 := make(map[uint32]bool)
    if false {
        v1 = (v1 / (((v1 >> ((420) & 63))) | 1))
    }
    var v12 float32 = 0.0
    var v13 string = ""
    fmt.Println(v12)
    for fi14 := 0; fi14 < 5; fi14++ {
        fmt.Println(v1)
        if false {
            var v15 float64 = (0.0 - 0.1)
            if true {
                v1 = 42
                switch fi14 {
                case 840:
                    v13 = v13
                }
                switch (fi14 & -889) {
                case -923:
                    for fi16 := 0; fi16 < 20; fi16++ {
                        v17 := GenS1{F0: 127}
                        v18 := GenS2{F0: v5}
                        fmt.Println(v15)
                        v5 = -9223372036854775808
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                case 32768:
                    v13 = (v13 + (v13 + ("foo bar" + (((v13 + v13) + v13) + v13))))
                case -379:
                    v13 = v13
                default:
                    v1 = (v1 >> ((255) & 63))
                }
            }
            _ = v15
        }
        v13 = ""
        _ = fi14
    }
    v1 = v1
    v4 += v4
    v4 = -835
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 float32 = 0.0
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v5 -= v5
        fmt.Println(v13, fi21, fi21)
        v22 := GenS3{F0: -2147483648, F1: v1}
        _ = fi21
        _ = v22
    }
    delete(v3, ("a" + v13))
    var v23 bool = (false && (0 <= -35))
    v24 := GenS2{F0: v5}
    var v25 int = -292
    var v26 uint = 964
    if (v23 || (int32(v5) > 353)) {
        v27 := GenS3{F0: 400, F1: v1}
        v26 -= v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
