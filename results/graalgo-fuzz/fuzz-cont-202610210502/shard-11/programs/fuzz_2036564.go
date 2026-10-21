package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 255
    var v3 uint32 = 42
    var v4 int = 128
    var v5 uint16 = 1
    var v6 uint8 = 204
    fmt.Println(v4, v2, v2)
    v3 -= (^(v3))
    var v7 uint = (^(v2))
    if (v6 == (v6 + v6)) {
        v8 := []int{}
        fmt.Println(v6, v4, v4)
        _ = v8
    } else {
        var v9 int32 = -153
        var v10 int32 = (v9 >> ((uint(v5)) & 63))
        fmt.Println(v6)
        _ = v10
        _ = v9
    }
    var v11 uint64 = 4294967295
    v6 -= 255
    v2 = v2
    var v12 bool = false
    v13 := GenS1{F0: (v7 >= v7), F1: "foo bar", F2: v6}
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := GenS1{F0: v12, F1: "", F2: 128}
        if (v2 <= v2) {
            v1[(("hello" + "\n\t") + ("\n\t" + "a"))] = uint32(v2)
        } else {
            if v12 {
                v12 = (-77 <= 127)
                v16 := GenS1{F0: v12, F1: "a", F2: 113}
                _ = v16
            } else {
                v17 := GenS1{F0: (!v12), F1: ("hello" + ((("\n\t" + ("\n\t" + "hello")) + ("\n\t" + "foo bar")) + (("a" + "\n\t") + "foo bar"))), F2: (v6 >> (((v7 + (v2 | 2))) & 63))}
                var v18 int16 = -247
                _ = v17
                _ = v18
            }
            v19 := GenS1{F0: (-9223372036854775808 < v4), F1: "a", F2: v6}
            _ = v19
        }
        v20 := []int{}
        _ = fi14
        _ = v15
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        v6 -= (v6 & (v6 / (((v6 & v6)) | 1)))
        switch "foo bar" {
        case "\n\t":
            v7 += (^(v7))
        case "hello":
            var v22 int = 0
            _ = v22
        case "foo bar":
            var v23 int16 = -1
            _ = v23
        }
        v5 += 1
        _ = fi21
    }
    if v12 {
        v2 += 42
        v7 += (v2 * v2)
    } else {
        v24 := GenS1{F0: ((v12 || true) || false), F1: ("\n\t" + "foo bar"), F2: v6}
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
