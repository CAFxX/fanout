package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int32; F2 uint32 }

func main() {
    v1 := []string{}
    var v2 uint8 = 2
    var v3 int16 = 32767
    var v4 uint = 2
    v3 = (v3 ^ (v3 & (v3 | v3)))
    fmt.Println(v2)
    if true {
        for fi5 := 0; fi5 < 11; fi5++ {
            fmt.Println(v4)
            _ = fi5
        }
        v6 := GenS1{F0: true}
        _ = v6
    } else {
        var v7 int64 = 712
        var v8 uint = 983
        _ = v7
        _ = v8
    }
    var v9 int16 = v3
    if (((("a" + "foo bar") + ("hello" + "foo bar")) + "a") >= "hello") {
        v3 = v9
        if true {
            v3 = (-(v3))
            if len(v1) > 0 {
                v1[0] = (("" + "") + "\n\t")
            }
        }
        if (255 < 128) {
            v10 := GenS2{F0: 49, F1: "a"}
            _ = v10
        } else {
            v2 -= (^(v2))
            var v11 int16 = (v3 * 42)
            _ = v11
        }
    } else {
        v12 := GenS3{F0: 0}
        switch 42 {
        case 127:
            v4 += v4
        case 2147483647:
            var v13 int32 = 2147483647
            _ = v13
        case -598:
            v14 := GenS1{F0: ((v4 >> (((^(v4))) & 63)) != (v4 % ((v4) | 1)))}
            _ = v14
        }
        _ = v12
    }
    switch "a" {
    case "\n\t":
        fmt.Println(v3, v2)
    }
    v4 = v4
    if false {
        var v15 uint = 0
        for fi16 := 0; fi16 < 9; fi16++ {
            v9 = (v3 << (((v15 & v15)) & 63))
            for fi17 := 0; fi17 < 10; fi17++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi17
            }
            v9 = v3
            _ = fi16
        }
        if len(v1) > 0 {
            v1[0] = ""
        }
        _ = v15
    } else {
        fmt.Println(v2, v3)
    }
    var v18 float32 = 3.14159
    if true {
        v2 = v2
    }
    if (18446744073709551615 == 313) {
        v19 := GenS4{F0: false, F1: -835, F2: 509}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
