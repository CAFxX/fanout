package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int64 }

func main() {
    var v1 string = (("hello" + "") + ((("a" + "a") + "hello") + (("hello" + "a") + ((("foo bar" + "a") + "\n\t") + ""))))
    var v2 uint64 = 18446744073709551615
    var v3 float64 = 1e10
    var v4 int32 = 0
    var v5 uint = 237
    var v6 uint = v5
    v7 := GenS1{F0: 0, F1: 4, F2: 42}
    v8 := []int{-9223372036854775808}
    var v9 int32 = v4
    _ = len(v8)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v10 uint8 = 192
    if true {
        if true {
            switch "\n\t" {
            case "":
                _ = len(v8)
            case "\n\t":
                var v11 bool = true
                _ = v11
            case "hello":
                v8 = append(v8, 42)
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS1{F0: 221, F1: 0, F2: 1}
                _ = fi12
                _ = v13
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        } else {
            var v14 uint32 = 4294967295
            _ = v14
        }
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        fmt.Println(v5)
        var v16 uint8 = 128
        switch "\n\t" {
        case "\n\t":
            var v17 uint = uint(v2)
            _ = v17
        case "hello":
            for fi18 := 0; fi18 < 16; fi18++ {
                v19 := GenS1{F0: 65535, F1: 999, F2: -9223372036854775808}
                if false {
                    v5 += (^(v6))
                    v5 -= v5
                }
                v1 = ("foo bar" + "foo bar")
                _ = fi18
                _ = v19
            }
        case "a":
            var v20 uint8 = (v16 >> ((v6) & 63))
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        if false {
            v8 = append(v8, int(v5))
        } else {
            switch v1 {
            case "":
                if len(v8) > 0 {
                    v8[0] = (fi21 >> (((v5 >> ((v6) & 63))) & 63))
                }
            case "hello":
                var v22 uint = (^(v6))
                _ = v22
            case "foo bar":
                var v23 uint16 = 255
                _ = v23
            default:
                v24 := GenS1{F0: 779, F1: 0, F2: 0}
                _ = v24
            }
            var v25 int16 = 0
            _ = v25
        }
        fmt.Println(fi21)
        _ = len(v8)
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
