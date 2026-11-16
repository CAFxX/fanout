package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int; F2 uint32 }

func main() {
    v1 := []int{1}
    var v2 uint32 = 65535
    v3 := []uint32{v2, v2, 427, 2147483647}
    var v4 float64 = 3.14159
    v2 += (v2 & 423)
    v5 := GenS1{F0: 1, F1: (v2 | 101)}
    var v6 uint64 = 4294967295
    v2 = (v2 & 42)
    if true {
        fmt.Println(v2, v6, v2)
        v4 -= (1.5 + v4)
        v7 := GenS2{F0: "a", F1: 0}
        _ = v7
    } else {
        v6 += v6
        v6 -= 0
        fmt.Println(v6)
    }
    v8 := GenS3{F0: 40, F1: -124, F2: (v2 << ((18446744073709551615) & 63))}
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 float32 = 1.5
        switch ("" + "") {
        case "":
            var v11 int64 = -9223372036854775808
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    _ = len(v3)
    fmt.Println(v6)
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint32 = (v2 - v2)
        v2 -= v13
        _ = fi12
        _ = v13
    }
    v4 += v4
    if false {
        v2 = (v2 >> ((149) & 63))
        var v14 float64 = (1e10 / v4)
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS2{F0: (("hello" + ("a" + (("a" + ("a" + (("a" + ("hello" + "\n\t")) + "\n\t"))) + (("a" + "a") + ("" + ""))))) + "foo bar"), F1: 255}
            if false {
                fmt.Println(v14, v6, v14)
                v14 = (v14 - 1.5)
            } else {
                for fi17 := 0; fi17 < 10; fi17++ {
                    _ = len(v1)
                    v14 = v4
                    v18 := GenS2{F0: "\n\t", F1: 255}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v19 uint8 = 128
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            _ = len(v3)
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
