package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int; F2 uint }

func main() {
    var v1 uint8 = 0
    var v2 uint32 = 2
    var v3 uint32 = v2
    v4 := make(map[string]bool)
    var v5 uint16 = 255
    var v6 int16 = -32768
    var v7 uint8 = v1
    var v8 uint = 255
    v9 := GenS1{F0: -65, F1: 127, F2: "hello"}
    v10 := []int{-2147483648}
    if false {
        var v11 int64 = 42
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 uint64 = 0
        var v14 int32 = -647
        if (934 >= 606) {
            for fi15 := 0; fi15 < 20; fi15++ {
                v10 = append(v10, fi12)
                v4["hello"] = false
                v16 := GenS2{F0: (v8 | (v8 | 0))}
                _ = fi15
                _ = v16
            }
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 int16 = v6
                switch (fi12 ^ 9223372036854775807) {
                case -97:
                    var v19 string = "hello"
                    _ = v19
                case -2147483648:
                    v13 = (v13 | v13)
                case 4294967295:
                    switch "hello" {
                    case "hello":
                        fmt.Println(v6)
                    }
                }
                var v20 uint8 = 149
                _ = fi17
                _ = v18
                _ = v20
            }
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v21 int64 = 9223372036854775807
    delete(v4, "foo bar")
    if true {
        var v22 int64 = (v21 * (v21 << (((v8 + v8)) & 63)))
        fmt.Println(v8)
        v7 = (v1 << ((331) & 63))
        _ = v22
    }
    v1 = (v7 << (((^(v8))) & 63))
    v23 := GenS3{F0: ("foo bar" + ""), F1: -27, F2: v8}
    switch ("" + "\n\t") {
    case "":
        if false {
            for fi24 := 0; fi24 < 11; fi24++ {
                v25 := GenS3{F0: "a", F1: (fi24 >> ((v8) & 63)), F2: (v8 >> ((2) & 63))}
                for fi26 := 0; fi26 < 12; fi26++ {
                    var v27 bool = false
                    v2 = v3
                    _ = fi26
                    _ = v27
                }
                v21 = 0
                _ = fi24
                _ = v25
            }
        }
    case "\n\t":
        v28 := GenS3{F0: ("" + "a"), F1: -1, F2: 42}
        _ = v28
    case "a":
        v5 -= v5
    }
    fmt.Println(v8, v21, v3)
    var v29 bool = true
    v30 := GenS3{F0: "hello", F1: 717, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
