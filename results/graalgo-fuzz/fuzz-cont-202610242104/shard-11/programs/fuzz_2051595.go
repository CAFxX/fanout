package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 string }

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]int)
    var v3 uint64 = 0
    var v4 uint8 = 127
    v5 := []uint32{812, 525, 2147483648}
    fmt.Println(v4, v3, v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        if ((3.14159 / 1.0) < 1e10) {
            for fi7 := 0; fi7 < 11; fi7++ {
                if false {
                    v1 -= v1
                    fmt.Println(fi7, fi7)
                    var v8 int = (^(fi7))
                    _ = v8
                }
                v9 := GenS1{F0: 4294967295, F1: ("a" + "")}
                v10 := GenS2{F0: true, F1: 0}
                _ = fi7
                _ = v10
                _ = v9
            }
            var v11 int8 = 100
            _ = v11
        }
        switch ("a" + "\n\t") {
        case "":
            v1 = 42
        case "\n\t":
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 int8 = -43
                fmt.Println(v1, v13, fi12)
                _ = fi12
                _ = v13
            }
        case "foo bar":
            v3 += 4294967295
        }
        var v14 int32 = 0
        _ = fi6
        _ = v14
    }
    switch "" {
    case "hello":
        var v15 uint = 4294967295
        _ = v15
    }
    if false {
        v16 := v2[1]
        v3 = 305
        _ = v16
    }
    if ((v1 | v1) <= (v1 % ((v1) | 1))) {
        v1 = v1
        var v17 string = ("" + ("\n\t" + ("\n\t" + "")))
        _ = v17
    } else {
        var v18 uint16 = 7
        var v19 uint64 = (^(v3))
        v20 := GenS3{F0: 2147483647}
        _ = v18
        _ = v19
        _ = v20
    }
    v21 := GenS4{F0: v4, F1: 128}
    var v22 uint32 = 65535
    v3 += v3
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch 0 {
    case 32767:
        switch ("foo bar" + ((("" + "") + "") + "")) {
        case "":
            if (0 == -28) {
                _ = len(v5)
                v23 := GenS5{F0: (("hello" + "") + (("a" + "\n\t") + "hello"))}
                _ = v23
            } else {
                fmt.Println(v22, v4, v3)
            }
        case "foo bar":
            var v24 uint8 = v4
            _ = v24
        default:
            v4 += (v4 >> ((447) & 63))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
