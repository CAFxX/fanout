package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int16; F2 string }

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("foo bar" + ("a" + ((("" + "a") + "a") + ("" + ""))))
    var v3 int8 = 42
    v4 := []int{}
    switch v2 {
    case "foo bar":
        v5 := GenS1{F0: -2147483648}
        _ = v5
    default:
        var v6 int8 = v3
        _ = v6
    }
    var v7 uint64 = 42
    switch 738 {
    case 2147483647:
        var v8 uint16 = 2
        _ = v8
    case -9223372036854775808:
        v4 = append(v4, 32767)
    case -1:
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
    }
    v9 := GenS1{F0: 0}
    v2 = ("foo bar" + ("foo bar" + ((v2 + "\n\t") + "hello")))
    v10 := GenS1{F0: -2147483648}
    fmt.Println(v2, v2, v2)
    var v11 string = v2
    var v12 string = (((v11 + ("hello" + v2)) + v2) + v11)
    if false {
        v13 := GenS1{F0: -923}
        var v14 int32 = 42
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS1{F0: int32(fi15)}
        _ = fi15
        _ = v16
    }
    if (255 < 65535) {
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int = fi17
            v19 := GenS1{F0: -2147483648}
            v11 = v12
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    v3 -= v3
    for fi20 := 0; fi20 < 1; fi20++ {
        delete(v1, fi20)
        _ = fi20
    }
    if true {
        var v21 int8 = int8(v7)
        if true {
            var v22 float32 = -2.25
            v11 = ""
            delete(v1, 2147483647)
            _ = v22
        } else {
            var v23 int = int(v3)
            _ = v23
        }
        _ = v21
    } else {
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := GenS2{F0: (v3 | (v3 % ((v3) | 1))), F1: 42, F2: "\n\t"}
            _ = fi24
            _ = v25
        }
        if (false || true) {
            fmt.Println(v11)
            v12 = "\n\t"
            v26 := GenS1{F0: 42}
            _ = v26
        }
    }
    var v27 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
