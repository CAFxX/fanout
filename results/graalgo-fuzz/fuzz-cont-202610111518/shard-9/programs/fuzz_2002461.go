package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 bool }

func main() {
    var v1 float32 = float32(((-2.25 - 0.0) / float64(0.1)))
    v2 := []uint32{}
    v3 := []string{"foo bar", "a", ("" + "a")}
    v1 += v1
    switch "a" {
    case "foo bar":
        if false {
            v1 = v1
            v1 -= 1e10
            if true {
                var v4 int = 1
                v1 += 3.14159
                _ = v4
            }
        } else {
            if true {
                switch ("\n\t" + (("\n\t" + "\n\t") + "")) {
                case "":
                    fmt.Println(v1)
                default:
                    var v5 uint32 = 0
                    _ = v5
                }
                for fi6 := 0; fi6 < 6; fi6++ {
                    fmt.Println(fi6)
                    if (false || true) {
                        fmt.Println(v1, v1)
                        var v7 float32 = (0.1 + 0.1)
                        _ = v7
                    } else {
                        v1 = v1
                        var v8 int32 = -1
                        _ = v8
                    }
                    _ = fi6
                }
                fmt.Println(v1)
            }
            fmt.Println(v1)
        }
    }
    if len(v2) > 0 {
        v2[0] = 159
    }
    v2 = append(v2, 2)
    var v9 bool = (276 <= 774)
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 float64 = 1e10
        if ((int16(fi10) > -1) && (float32(v11) != ((v1 / 3.14159) - v1))) {
            v11 += (float64(v1) * v11)
            if (1000 <= 1) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch ("hello" + ("\n\t" + "")) {
                case "":
                    v1 = v1
                case "hello":
                    v9 = v9
                case "\n\t":
                    fmt.Println(v9, v1)
                }
            } else {
                fmt.Println(v9, v11, v1)
            }
            var v12 uint = 189
            _ = v12
        } else {
            v1 = 0.1
        }
        _ = fi10
        _ = v11
    }
    if len(v3) > 0 {
        v3[0] = ("" + "\n\t")
    }
    if v9 {
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 int64 = 1
            for fi15 := 0; fi15 < 3; fi15++ {
                v9 = false
                _ = fi15
            }
            v3 = append(v3, (("a" + "\n\t") + "hello"))
            _ = fi13
            _ = v14
        }
    }
    v16 := GenS1{F0: 1, F1: 773}
    if (187 > 0) {
        fmt.Println(v1, v1)
        fmt.Println(v9, v9)
    } else {
        switch -418 {
        case 128:
            v2 = append(v2, 65535)
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v17 := GenS1{F0: 0, F1: 2147483647}
    fmt.Println(v9)
    v18 := GenS1{F0: 0, F1: 4294967295}
    var v19 uint64 = 18446744073709551615
    v20 := GenS2{F0: (128 > -310)}
    var v21 int64 = -1
    var v22 int = 617
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v9
}
