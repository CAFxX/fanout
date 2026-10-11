package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int; F2 int }
type GenS4 struct { F0 bool }
func genFn0(p0 uint16) int32 {
    v6 := GenS1{F0: (((("\n\t" + ((("" + "hello") + "\n\t") + "\n\t")) + "hello") + "hello") + ("a" + "")), F1: -1}
    for fi7 := 0; fi7 < 12; fi7++ {
        p0 = p0
        _ = fi7
    }
    var v8 int = 2147483647
    _ = p0
    _ = v6
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{1, 1, 275, 780}
    v2 := []uint32{287, 743, 2147483647, 2147483648}
    var v3 uint64 = 4294967295
    var v4 float32 = 1.5
    var v5 int16 = 1
    switch "a" {
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "hello":
        var v9 int = -732
        _ = v9
    }
    var v10 int8 = 127
    if ("a" <= ("" + "")) {
        var v11 uint32 = 42
        _ = v11
    }
    v12 := GenS2{F0: (805 >= 128)}
    v4 = float32(0.0)
    if true {
        if (false && false) {
            var v13 uint64 = v3
            var v14 string = ""
            _ = v13
            _ = v14
        } else {
            v3 = (v3 / ((v3) | 1))
            v15 := GenS3{F0: 0, F1: 0, F2: 128}
            v10 = v10
            _ = v15
        }
        var v16 int64 = 121
        for fi17 := 0; fi17 < 6; fi17++ {
            if true {
                if (((v3 * (v3 << ((255) & 63))) > v3) && (65535 >= 255)) {
                    _ = len(v1)
                    var v18 uint8 = 255
                    fmt.Println(v16, v5, v5)
                    _ = v18
                }
                v19 := GenS1{F0: ("" + "\n\t"), F1: (v16 % ((v16) | 1))}
                var v20 int8 = (v10 + int8(fi17))
                _ = v19
                _ = v20
            } else {
                v10 = 106
            }
            _ = fi17
        }
        _ = v16
    }
    switch -889 {
    case 355:
        if (v3 != 0) {
            v21 := GenS3{F0: 521, F1: 128, F2: 4294967295}
            var v22 uint32 = 2
            v22 -= (v22 | 733)
            _ = v21
            _ = v22
        } else {
            var v23 uint = 4294967295
            v3 -= (v3 ^ v3)
            var v24 string = "a"
            _ = v23
            _ = v24
        }
    default:
        fmt.Println(v5, v5)
    }
    var v25 string = "\n\t"
    v26 := GenS4{F0: (465 != 42)}
    switch 128 {
    case 98:
        switch (((("a" + "hello") + "a") + (v25 + v25)) + "") {
        case "foo bar":
            var v27 int16 = (v5 | (v5 ^ (v5 - 0)))
            _ = v27
        case "\n\t":
            v4 = 0.1
        default:
            for fi28 := 0; fi28 < 2; fi28++ {
                v1 = append(v1, 65535)
                fmt.Println(v25)
                _ = fi28
            }
        }
    }
    v4 += 0.1
    _ = len(v2)
    var v29 int16 = (v5 * v5)
    var v30 uint = 1
    for fi31 := 0; fi31 < 19; fi31++ {
        v25 = ""
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
