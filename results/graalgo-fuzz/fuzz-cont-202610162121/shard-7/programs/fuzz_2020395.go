package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 bool = (true || true)
    var v2 int = 903
    var v3 uint32 = 524
    var v4 int = (v2 - 32768)
    var v5 int16 = 1000
    var v6 int32 = -763
    var v7 bool = (!v1)
    v2 = (v2 >> ((0) & 63))
    if true {
        v4 = (v4 & v2)
        fmt.Println(v1, v2, v1)
        v8 := []int{32768}
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        v1 = v1
        v10 := GenS1{F0: true}
        _ = fi9
        _ = v10
    }
    fmt.Println(v3, v2)
    v1 = v1
    switch ("hello" + "foo bar") {
    case "foo bar":
        v11 := GenS1{F0: (42 < 0)}
        _ = v11
    case "\n\t":
        switch -199 {
        case -9223372036854775808:
            switch ("foo bar" + "a") {
            case "hello":
                v12 := []int{(v4 - 9223372036854775807)}
                _ = v12
            case "":
                v4 = v4
            }
        case -968:
            var v13 float32 = 0.0
            _ = v13
        default:
            var v14 uint8 = 128
            _ = v14
        }
    case "":
        v4 += 1
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 string = ("" + ("a" + "\n\t"))
        v17 := []int{v2, (v4 >> ((65535) & 63)), 9223372036854775807, 42}
        v5 -= (v5 * (v5 << ((705) & 63)))
        _ = fi15
        _ = v16
        _ = v17
    }
    if (v2 <= (v2 ^ 1)) {
        var v18 int8 = -79
        v1 = (!false)
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
