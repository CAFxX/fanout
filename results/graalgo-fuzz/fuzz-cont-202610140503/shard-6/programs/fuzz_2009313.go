package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 152
    v3 := []uint32{0, 42, 65535}
    var v4 uint16 = 65535
    switch ("a" + "foo bar") {
    case "a":
        v4 = 890
    case "hello":
        v4 += v4
    default:
        switch -333 {
        case -624:
            fmt.Println(v4, v4)
        case 42:
            v4 = (v4 | (v4 - 0))
        case -455:
            var v5 int8 = 0
            _ = v5
        default:
            fmt.Println(v4, v2)
        }
    }
    var v6 int8 = 0
    switch "\n\t" {
    case "a":
        _ = len(v1)
    case "foo bar":
        fmt.Println(v4)
    }
    var v7 uint16 = 736
    var v8 float32 = -2.25
    v4 += v4
    var v9 string = "hello"
    v7 += v7
    var v10 bool = (9223372036854775807 != 308)
    switch (v9 + "\n\t") {
    case "foo bar":
        var v11 uint64 = 4294967295
        _ = v11
    case "a":
        v6 += (v6 - 42)
    case "\n\t":
        var v12 uint32 = 65535
        _ = v12
    }
    if v10 {
        v7 = v7
        if v10 {
            fmt.Println(v4, v4)
            v2 -= v2
            _ = len(v1)
        }
    } else {
        v7 = v7
        switch int(v6) {
        case 1:
            fmt.Println(v7, v9)
        case 127:
            fmt.Println(v9, v6, v9)
        case 9223372036854775807:
            v13 := GenS1{F0: v10}
            _ = v13
        }
        v14 := GenS1{F0: (-312 != -9223372036854775808)}
        _ = v14
    }
    v15 := v1[686]
    for fi16 := 0; fi16 < 4; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            v3 = append(v3, 1)
            _ = len(v3)
            var v18 string = v15
            _ = fi17
            _ = v18
        }
        if false {
            var v19 uint = 42
            _ = v19
        }
        switch ("\n\t" + "a") {
        case "\n\t":
            v20 := v1[65535]
            _ = v20
        default:
            var v21 uint16 = (^(v7))
            _ = v21
        }
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
