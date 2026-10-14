package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint16; F2 int }

func main() {
    v1 := []string{"foo bar", "foo bar", "hello", ("hello" + ("a" + ("hello" + ("hello" + ("" + "")))))}
    var v2 uint16 = 32767
    var v3 uint = 2
    var v4 uint32 = 3
    if true {
        fmt.Println(v3)
        var v5 int16 = 1
        var v6 int64 = int64(v3)
        _ = v5
        _ = v6
    }
    var v7 bool = ((4294967295 > 885) && false)
    var v8 bool = (0.0 <= float64(0.0))
    v8 = (-516 > -1)
    switch ("" + "hello") {
    case "a":
        var v9 int16 = 1000
        _ = v9
    case "foo bar":
        var v10 int16 = 1
        _ = v10
    }
    var v11 uint16 = (v2 ^ 255)
    v2 = 505
    switch ("\n\t" + ("\n\t" + ((((("foo bar" + "") + "hello") + "") + "\n\t") + ("foo bar" + "")))) {
    case "hello":
        v12 := GenS1{F0: 2}
        _ = v12
    case "":
        v13 := GenS2{F0: 969}
        _ = v13
    }
    if (v3 <= v3) {
        for fi14 := 0; fi14 < 15; fi14++ {
            v1 = append(v1, "\n\t")
            v11 = v2
            _ = fi14
        }
        var v15 uint32 = v4
        _ = v15
    } else {
        var v16 string = ""
        v1 = append(v1, v16)
        v4 += (v4 * v4)
        _ = v16
    }
    v17 := GenS3{F0: 42, F1: v2, F2: int(v3)}
    var v18 uint64 = 18446744073709551615
    switch 32767 {
    case 0:
        _ = len(v1)
    case 32768:
        var v19 int64 = 0
        _ = v19
    case -229:
        if (42 != 191) {
            fmt.Println(v18)
            switch 0 {
            case -1:
                v7 = v8
            case 989:
                if true {
                    v2 = v2
                }
            case 1:
                if v7 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v7 = v8
                    if v7 {
                        v18 += (v18 ^ (v18 ^ v18))
                        if (!false) {
                            v3 = 18446744073709551615
                            var v20 int8 = -91
                            _ = v20
                        }
                    }
                }
            }
            v21 := GenS2{F0: 0}
            _ = v21
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
