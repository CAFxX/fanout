package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 string }
type GenS2 struct { F0 uint64; F1 uint; F2 int16 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 uint8 = 182
    v2 := make(map[string]int)
    var v3 float32 = 1.5
    v4 := make(map[uint32]string)
    v5 := []int{32768, -1, 815}
    for fi6 := 0; fi6 < 9; fi6++ {
        v1 = v1
        if len(v5) > 0 {
            v5[0] = 9223372036854775807
        }
        var v7 int16 = -428
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: (v1 << ((42) & 63)), F1: v1, F2: ""}
    var v9 uint = 18446744073709551615
    var v10 float64 = (3.14159 + -2.25)
    v11 := GenS1{F0: v1, F1: (v1 >> (((v9 >> ((v9) & 63))) & 63)), F2: (("\n\t" + "") + "foo bar")}
    for fi12 := 0; fi12 < 20; fi12++ {
        if false {
            if (!true) {
                v10 -= 0.0
                for fi13 := 0; fi13 < 2; fi13++ {
                    var v14 uint16 = 2
                    var v15 int16 = 42
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                if len(v5) > 0 {
                    v5[0] = fi12
                }
            } else {
                v3 += v3
            }
        } else {
            for fi16 := 0; fi16 < 7; fi16++ {
                v4[4294967295] = ""
                _ = fi16
            }
            v17 := GenS1{F0: 42, F1: v1, F2: ((("foo bar" + "\n\t") + "") + ("\n\t" + ((((((((("a" + "\n\t") + "a") + "") + "foo bar") + ("hello" + "a")) + "") + "hello") + "hello") + "")))}
            _ = v17
        }
        v18 := GenS2{F0: 42, F1: 255, F2: -32768}
        _ = fi12
        _ = v18
    }
    if true {
        v19 := GenS1{F0: v1, F1: 255, F2: "foo bar"}
        v9 = v9
        _ = v19
    } else {
        v5 = append(v5, -9223372036854775808)
        var v20 uint64 = 42
        v10 = 1.5
        _ = v20
    }
    v21 := GenS3{F0: -117}
    fmt.Println(v3, v1, v1)
    for fi22 := 0; fi22 < 4; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := GenS3{F0: 100}
            _ = fi23
            _ = v24
        }
        var v25 int = fi22
        _ = fi22
        _ = v25
    }
    if false {
        v1 = (^(v1))
        v3 += (v3 / v3)
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        v27 := GenS3{F0: 100}
        fmt.Println(v9)
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
