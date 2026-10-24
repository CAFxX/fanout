package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }

func main() {
    var v1 float64 = (0.1 / -2.25)
    var v2 string = ("" + "")
    v3 := make(map[uint32]bool)
    v1 = 1e10
    v4 := GenS1{F0: 997, F1: 1}
    v2 = (v2 + ((("foo bar" + v2) + ("" + "\n\t")) + v2))
    switch 0 {
    case 4294967295:
        v5 := GenS1{F0: 4294967295, F1: 42}
        _ = v5
    default:
        v6 := GenS1{F0: 127, F1: 1000}
        _ = v6
    }
    switch v2 {
    case "foo bar":
        switch -9223372036854775808 {
        case 687:
            switch -9223372036854775808 {
            case -479:
                delete(v3, 65535)
            case 32768:
                v7 := GenS1{F0: 592, F1: -1}
                _ = v7
            case -1:
                var v8 int64 = 42
                _ = v8
            }
        }
    case "hello":
        switch -9223372036854775808 {
        case 1:
            v9 := GenS1{F0: 396, F1: -32768}
            _ = v9
        case 325:
            fmt.Println(v1)
        case 32768:
            v10 := []int{}
            _ = v10
        default:
            var v11 uint = 2
            _ = v11
        }
    }
    v2 = v2
    var v12 string = v2
    if (!(2 <= 990)) {
        if ((!true) || true) {
            if (127 == 103) {
                fmt.Println(v2, v12, v2)
                v12 = ((v2 + v2) + v2)
                v12 = ((v2 + v12) + v2)
            }
            for fi13 := 0; fi13 < 14; fi13++ {
                v1 += 1e10
                var v14 uint16 = 256
                _ = fi13
                _ = v14
            }
        } else {
            var v15 float32 = 0.0
            _ = v15
        }
    } else {
        var v16 bool = true
        var v17 float64 = v1
        v12 = "foo bar"
        _ = v16
        _ = v17
    }
    var v18 int32 = 569
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
