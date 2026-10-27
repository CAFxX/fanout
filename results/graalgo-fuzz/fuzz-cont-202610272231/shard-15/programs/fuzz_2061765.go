package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 bool; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint16 }

func main() {
    v1 := []int{-2147483648}
    v2 := []uint32{4294967295, 884}
    v3 := []int{-9223372036854775808}
    switch "a" {
    case "\n\t":
        v4 := GenS1{F0: 1, F1: "hello", F2: 622}
        _ = v4
    default:
        var v5 uint8 = 0
        _ = v5
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(fi6)
        for fi7 := 0; fi7 < 7; fi7++ {
            if false {
                switch ("foo bar" + ("\n\t" + ("foo bar" + "a"))) {
                case "foo bar":
                    var v8 int64 = -660
                    _ = v8
                case "a":
                    if false {
                        var v9 uint8 = 0
                        var v10 int8 = 100
                        _ = v10
                        _ = v9
                    } else {
                        fmt.Println(fi6, fi6)
                        var v11 int = 0
                        _ = v11
                    }
                }
                fmt.Println(fi6, fi6)
                switch fi6 {
                case 42:
                    var v12 float64 = float64(3.14159)
                    _ = v12
                }
            } else {
                fmt.Println(fi7, fi6)
                fmt.Println(fi6, fi6)
                fmt.Println(fi7, fi7)
                v13 := GenS1{F0: 32767, F1: "foo bar", F2: 65535}
                var v14 float32 = -2.25
                _ = v13
                _ = v14
            }
            var v15 uint = 65535
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi7
            _ = v15
        }
        if false {
            v16 := GenS2{F0: 2147483648, F1: 42}
            switch (fi6 << ((930) & 63)) {
            case 127:
                var v17 string = ("hello" + (("foo bar" + "") + "foo bar"))
                _ = v17
            default:
                var v18 float32 = 3.14159
                _ = v18
            }
            _ = v16
        } else {
            for fi19 := 0; fi19 < 20; fi19++ {
                switch (("" + "\n\t") + "hello") {
                case "":
                    v20 := GenS2{F0: 0, F1: 1}
                    _ = v20
                default:
                    var v21 int = fi19
                    _ = v21
                }
                _ = fi19
            }
            var v22 uint = 55
            var v23 uint16 = 923
            _ = v22
            _ = v23
        }
        _ = fi6
    }
    v24 := GenS3{F0: ("" < "hello"), F1: 1, F2: 0}
    v25 := GenS4{F0: 960, F1: 255, F2: 255}
    _ = len(v2)
    switch "foo bar" {
    case "\n\t":
        switch "a" {
        case "\n\t":
            var v26 uint8 = 49
            _ = v26
        case "a":
            var v27 uint16 = 255
            _ = v27
        }
    case "a":
        var v28 bool = false
        _ = v28
    case "hello":
        var v29 float32 = (-2.25 * 1e10)
        _ = v29
    }
    v30 := GenS4{F0: 843, F1: 255, F2: 305}
    if (-32768 > 131) {
        var v31 uint = 1
        _ = v31
    }
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    var v32 int8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v32
}
