package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 uint8 }

func main() {
    var v1 uint8 = 128
    var v2 string = "foo bar"
    var v3 int64 = -617
    var v4 int = 9223372036854775807
    var v5 uint16 = 32767
    var v6 float32 = 0.0
    v2 = ""
    var v7 int64 = 1
    var v8 uint8 = (v1 & (v1 % ((v1) | 1)))
    var v9 uint16 = v5
    v10 := []int{(^(v4)), v4, 2147483647}
    v5 = v9
    switch v2 {
    case "foo bar":
        switch v4 {
        case 128:
            if true {
                var v11 int64 = v3
                if ((v8 % ((v1) | 1)) <= v8) {
                    var v12 uint = 65535
                    v13 := GenS1{F0: (v5 ^ 0)}
                    v4 -= v4
                    _ = v12
                    _ = v13
                } else {
                    _ = len(v10)
                    v10 = append(v10, -9223372036854775808)
                    var v14 uint32 = 4294967295
                    _ = v14
                }
                fmt.Println(v8, v11)
                _ = v11
            } else {
                var v15 int = 833
                for fi16 := 0; fi16 < 12; fi16++ {
                    if (!true) {
                        if true {
                            v17 := GenS2{F0: (v1 ^ (^(v8)))}
                            var v18 uint64 = 760
                            _ = v17
                            _ = v18
                        }
                        var v19 uint32 = 0
                        _ = v19
                    }
                    var v20 uint64 = 4294967295
                    _ = fi16
                    _ = v20
                }
                fmt.Println(v15, v2, v1)
                _ = v15
            }
        case 42:
            for fi21 := 0; fi21 < 13; fi21++ {
                v22 := GenS3{F0: 65535}
                _ = len(v10)
                _ = fi21
                _ = v22
            }
        case -955:
            v9 = (^(v9))
        }
    case "hello":
        v2 = ("foo bar" + v2)
    }
    var v23 bool = (!((v4 / ((-123) | 1)) < v4))
    v1 -= v8
    fmt.Println(v8)
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 uint32 = 4294967295
        _ = fi24
        _ = v25
    }
    v26 := GenS4{F0: 105, F1: (v1 * (v8 + (v1 - 112)))}
    if len(v10) > 0 {
        v10[0] = (v4 % ((v4) | 1))
    }
    v10 = append(v10, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
