package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 string }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int16; F1 bool; F2 int32 }
type GenS4 struct { F0 string; F1 uint16; F2 int16 }

func main() {
    var v1 int64 = 1
    var v2 float64 = 1.5
    var v3 int = 1
    var v4 uint = 18446744073709551615
    v5 := GenS1{F0: 255, F1: -3, F2: "foo bar"}
    v6 := []int{}
    v7 := GenS1{F0: 876, F1: -1, F2: ("hello" + (("a" + "a") + "hello"))}
    for fi8 := 0; fi8 < 18; fi8++ {
        v1 += (v1 ^ (v1 + (v1 % ((v1) | 1))))
        v2 += 1e10
        _ = fi8
    }
    fmt.Println(v2)
    if false {
        var v9 uint16 = 0
        v2 += (3.14159 - v2)
        var v10 uint32 = 1
        _ = v10
        _ = v9
    }
    v4 = v4
    switch "" {
    case "foo bar":
        if (2147483647 < 2147483648) {
            if (false || true) {
                v11 := GenS1{F0: 0, F1: 42, F2: (("a" + "a") + "hello")}
                if len(v6) > 0 {
                    v6[0] = 4294967295
                }
                var v12 uint64 = 22
                _ = v11
                _ = v12
            } else {
                var v13 float64 = v2
                var v14 float64 = 0.0
                for fi15 := 0; fi15 < 17; fi15++ {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        v17 := GenS2{F0: "hello", F1: -116}
                        _ = fi16
                        _ = v17
                    }
                    fmt.Println(v1)
                    _ = fi15
                }
                _ = v13
                _ = v14
            }
            v2 -= ((v2 - v2) * 3.14159)
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 uint64 = 791
                v20 := GenS3{F0: 576, F1: false, F2: 1}
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    }
    v3 = (v3 % ((-385) | 1))
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := GenS4{F0: "a", F1: 2, F2: -714}
        _ = fi21
        _ = v22
    }
    v1 += 1
    _ = len(v6)
    var v23 uint = v4
    v23 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
