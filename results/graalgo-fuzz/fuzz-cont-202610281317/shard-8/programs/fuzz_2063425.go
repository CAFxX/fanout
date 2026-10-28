package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int64 }

func main() {
    var v1 float64 = float64(float32(1.5))
    var v2 bool = false
    var v3 int8 = -65
    var v4 uint8 = 1
    var v5 int16 = 32767
    var v6 float32 = float32(float64(float32(v1)))
    v7 := []int{395, 42, 32768}
    v2 = (!(1 > -9223372036854775808))
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(v5)
        _ = fi8
    }
    v5 += -562
    if v2 {
        for fi9 := 0; fi9 < 14; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 uint16 = 32767
                if v2 {
                    var v12 int64 = 1
                    v3 = 108
                    _ = v12
                } else {
                    var v13 uint = 18446744073709551615
                    v14 := GenS1{F0: "foo bar", F1: uint64(v4)}
                    if len(v7) > 0 {
                        v7[0] = (^(fi10))
                    }
                    _ = v13
                    _ = v14
                }
                v15 := GenS2{F0: -9223372036854775808}
                _ = fi10
                _ = v11
                _ = v15
            }
            _ = fi9
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            for fi17 := 0; fi17 < 18; fi17++ {
                _ = len(v7)
                var v18 bool = v2
                _ = fi17
                _ = v18
            }
            fmt.Println(v2)
            v5 = v5
            _ = fi16
        }
        fmt.Println(v3, v3)
    }
    if (32768 < -2147483648) {
        v19 := GenS3{F0: 4294967295, F1: 42}
        _ = v19
    } else {
        var v20 int = int(v3)
        _ = v20
    }
    var v21 uint8 = 217
    for fi22 := 0; fi22 < 18; fi22++ {
        fmt.Println(v6)
        _ = fi22
    }
    switch (("hello" + ((("foo bar" + "hello") + "a") + ("hello" + "hello"))) + ("hello" + "a")) {
    case "a":
        v23 := GenS3{F0: 384, F1: -553}
        _ = v23
    case "":
        if (("a" + (("\n\t" + "\n\t") + "\n\t")) < ("" + "a")) {
            v7 = append(v7, -2147483648)
            v24 := GenS3{F0: 162, F1: -224}
            _ = v24
        }
    default:
        v25 := GenS1{F0: "foo bar", F1: 66}
        _ = v25
    }
    _ = len(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
