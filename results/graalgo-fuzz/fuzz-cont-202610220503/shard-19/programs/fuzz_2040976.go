package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int32; F2 bool }

func main() {
    v1 := []uint32{556, 1, 2147483648}
    v2 := []string{"foo bar", (("\n\t" + ("a" + ("hello" + "hello"))) + "hello"), "hello", "hello"}
    var v3 float32 = 0.0
    var v4 bool = ((-1 == 100) && true)
    var v5 uint64 = 42
    v6 := make(map[uint32]int)
    var v7 uint16 = 65535
    if true {
        var v8 int32 = -795
        _ = v8
    }
    v9 := GenS1{F0: uint32(v5), F1: -617}
    var v10 uint64 = v5
    for fi11 := 0; fi11 < 18; fi11++ {
        switch "foo bar" {
        case "foo bar":
            v4 = (!(((0.1 - 0.1) == 1.5) && (572 < 65535)))
        }
        fmt.Println(v4, fi11)
        _ = fi11
    }
    switch "foo bar" {
    case "foo bar":
        v7 = v7
    }
    var v12 uint8 = 128
    if v4 {
        v13 := GenS2{F0: -1}
        var v14 uint16 = 65535
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 11; fi15++ {
            switch 127 {
            case -9223372036854775808:
                v16 := GenS2{F0: 0}
                _ = v16
            }
            _ = fi15
        }
        if v4 {
            if v4 {
                v7 = 255
            }
            for fi17 := 0; fi17 < 19; fi17++ {
                v18 := GenS1{F0: 2147483648, F1: 347}
                v6[2147483647] = 32768
                _ = fi17
                _ = v18
            }
            var v19 int = -585
            _ = v19
        }
        if len(v2) > 0 {
            v2[0] = (("hello" + (("" + "") + ("hello" + "foo bar"))) + ("foo bar" + "\n\t"))
        }
    }
    switch (("\n\t" + "\n\t") + "") {
    case "hello":
        switch (("hello" + ((("foo bar" + "") + ("\n\t" + ("\n\t" + "a"))) + "a")) + "\n\t") {
        case "":
            v10 -= v10
        case "foo bar":
            fmt.Println(v5, v7, v12)
        case "hello":
            var v20 uint32 = 65535
            _ = v20
        default:
            for fi21 := 0; fi21 < 16; fi21++ {
                v4 = false
                _ = fi21
            }
        }
    case "":
        v22 := GenS3{F0: 42, F1: 2147483647, F2: v4}
        _ = v22
    case "\n\t":
        v12 += v12
    default:
        if len(v2) > 0 {
            v2[0] = ("hello" + "hello")
        }
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    v6[2147483648] = 872
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
