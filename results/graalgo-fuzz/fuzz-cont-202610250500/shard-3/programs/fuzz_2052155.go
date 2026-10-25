package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 int16, p1 int8) int8 {
    if (!true) {
        switch ("\n\t" + "\n\t") {
        case "a":
            v7 := GenS1{F0: ""}
            _ = v7
        case "hello":
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS2{F0: p1}
                var v10 int32 = 42
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        var v11 int8 = (p1 - 25)
        _ = v11
    } else {
        var v12 int32 = 110
        _ = v12
    }
    var v13 int = 42
    _ = p0
    _ = p1
    _ = v13
    return (-(p1))
}
func genFn1() int16 {
    v14 := GenS2{F0: genFn0(-32768, -127)}
    if false {
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 int16 = 42
            _ = fi15
            _ = v16
        }
        var v17 float64 = 0.1
        _ = v17
    }
    _ = v14
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 0.1
    var v3 int64 = -9223372036854775808
    v4 := []string{("\n\t" + "hello"), ("hello" + "a"), "hello", "hello"}
    var v5 int16 = -1
    var v6 uint64 = 18446744073709551615
    switch "hello" {
    case "\n\t":
        v5 = v5
    default:
        if true {
            fmt.Println(v3, v6, v2)
            var v18 string = "a"
            _ = v18
        } else {
            if len(v4) > 0 {
                v4[0] = ("foo bar" + "a")
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                v1["foo bar"] = 0
                v20 := GenS1{F0: "a"}
                for fi21 := 0; fi21 < 13; fi21++ {
                    v22 := GenS2{F0: genFn0((v5 % ((v5) | 1)), -128)}
                    var v23 uint8 = 42
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = fi19
                _ = v20
            }
        }
    }
    var v24 uint16 = 1
    v3 -= (v3 | v3)
    v2 = ((-2.25 + 0.1) / v2)
    v25 := GenS3{F0: (-(v3))}
    var v26 int32 = 764
    var v27 int8 = genFn0(v5, 42)
    v2 = 0.0
    v28 := GenS4{F0: "", F1: (v24 * v24)}
    if true {
        for fi29 := 0; fi29 < 17; fi29++ {
            v30 := GenS2{F0: -56}
            _ = fi29
            _ = v30
        }
        v31 := GenS3{F0: int64(v24)}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v31
    }
    v27 = genFn0(739, v27)
    v32 := GenS1{F0: ("\n\t" + "a")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
