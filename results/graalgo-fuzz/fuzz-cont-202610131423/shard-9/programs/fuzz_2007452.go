package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int; F1 uint32; F2 uint8 }
func genFn0(p0 bool, p1 uint64, p2 uint64) string {
    switch ("\n\t" + "hello") {
    case "\n\t":
        for fi4 := 0; fi4 < 15; fi4++ {
            var v5 int = (-(fi4))
            var v6 int64 = 1
            _ = fi4
            _ = v5
            _ = v6
        }
    case "":
        if p0 {
            var v7 uint16 = 32767
            var v8 uint8 = uint8(p1)
            fmt.Println(v8, v8)
            _ = v7
            _ = v8
        } else {
            p0 = (!("hello" > ""))
        }
    case "hello":
        var v9 float32 = ((float32(0.0) - (0.0 - -2.25)) + 0.1)
        _ = v9
    }
    var v10 int64 = -250
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return ("foo bar" + ("hello" + "foo bar"))
}

func main() {
    var v1 uint16 = 42
    var v2 uint = 1
    v3 := []string{""}
    switch "hello" {
    case "":
        switch 9223372036854775807 {
        case 32768:
            v11 := GenS1{F0: (("\n\t" + ("foo bar" + "foo bar")) >= "hello")}
            _ = v11
        case 128:
            var v12 uint64 = 42
            _ = v12
        }
    }
    v13 := GenS2{F0: (!((!false) && false)), F1: 0}
    var v14 int64 = 0
    v1 = 881
    var v15 int64 = 0
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := GenS3{F0: fi16, F1: 95, F2: 87}
        _ = fi16
        _ = v17
    }
    var v18 uint8 = uint8(v1)
    var v19 int16 = 1
    v20 := GenS3{F0: int(v2), F1: 2147483647, F2: 42}
    v21 := GenS2{F0: false, F1: 9223372036854775807}
    if (!true) {
        v1 = v1
    } else {
        var v22 int16 = (v19 + (v19 - (v19 % ((v19) | 1))))
        _ = v22
    }
    var v23 uint64 = 0
    if (302 >= -1) {
        var v24 float32 = 1e10
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
}
