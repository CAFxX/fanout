package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 string }
type GenS2 struct { F0 int; F1 uint; F2 uint16 }

func main() {
    v1 := []string{(("foo bar" + ("a" + "a")) + "hello"), "foo bar"}
    var v2 int16 = 32767
    var v3 uint8 = 250
    v4 := []int{9223372036854775807, 405}
    switch 32768 {
    case 1:
        v2 = v2
    case 9223372036854775807:
        var v5 int = 0
        _ = v5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            v3 -= v3
            _ = fi7
        }
        var v8 bool = false
        fmt.Println(v2, v2, fi6)
        _ = fi6
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 2147483647, F2: ("" + "")}
    v2 = (v2 >> ((0) & 63))
    for fi10 := 0; fi10 < 14; fi10++ {
        v2 += (v2 ^ v2)
        var v11 uint = 127
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS1{F0: 42, F1: fi12, F2: "hello"}
        v14 := GenS2{F0: fi12, F1: 65535, F2: 0}
        var v15 int16 = (v2 - (v2 >> ((uint(v3)) & 63)))
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 0, F1: -383, F2: (((("foo bar" + "foo bar") + "") + "\n\t") + "a")}
    var v17 uint = 42
    if false {
        var v18 int64 = -1
        _ = v18
    }
    v3 -= v3
    if (true || (127 == 10)) {
        switch "\n\t" {
        case "foo bar":
            v2 = v2
        default:
            var v19 uint64 = 0
            _ = v19
        }
        var v20 int32 = 42
        if false {
            v2 = v2
            v3 += 72
        }
        _ = v20
    }
    if false {
        v21 := GenS2{F0: 726, F1: (v17 << ((v17) & 63)), F2: 65535}
        _ = v21
    }
    var v22 uint8 = v3
    var v23 uint16 = 244
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
