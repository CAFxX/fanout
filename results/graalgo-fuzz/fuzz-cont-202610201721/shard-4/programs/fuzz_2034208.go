package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    v4 := []string{"hello", (("" + "foo bar") + ("a" + ("" + "\n\t"))), "foo bar", "a"}
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 2
    v7 := GenS1{F0: "hello"}
    if (true && (true || (-2.25 <= 0.0))) {
        v6 += v6
    } else {
        var v8 uint16 = 1
        v9 := GenS1{F0: "foo bar"}
        v8 -= (v8 ^ v8)
        _ = v8
        _ = v9
    }
    v6 += v6
    switch "\n\t" {
    case "foo bar":
        if true {
            var v10 uint32 = v6
            for fi11 := 0; fi11 < 7; fi11++ {
                for fi12 := 0; fi12 < 5; fi12++ {
                    var v13 int8 = -128
                    v14 := GenS2{F0: -9223372036854775808}
                    fmt.Println(v6, fi12)
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = fi11
            }
            _ = v10
        }
    }
    var v15 string = "\n\t"
    var v16 int32 = 1
    v17 := GenS3{F0: 42, F1: 127, F2: -32768}
    for fi18 := 0; fi18 < 4; fi18++ {
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        var v19 int32 = (v16 + int32(fi18))
        _ = fi18
        _ = v19
    }
    v16 -= (-(v16))
    v5[1] = (^(v6))
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 int16 = -223
        v16 = (v16 & 2147483647)
        if (int32(v6) > -1) {
            v1[(("foo bar" + (v15 + v15)) + "foo bar")] = ((!false) || true)
            for fi22 := 0; fi22 < 6; fi22++ {
                if true {
                    v15 = (v15 + v15)
                } else {
                    if false {
                        v6 = (v6 - (v6 + (v6 * 65535)))
                    }
                }
                _ = fi22
            }
        }
        _ = fi20
        _ = v21
    }
    v1[v15] = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
