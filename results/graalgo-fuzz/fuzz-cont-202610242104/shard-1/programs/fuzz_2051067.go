package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int; F2 int64 }

func main() {
    var v1 int8 = 108
    var v2 float32 = 0.0
    var v3 uint = 255
    v4 := GenS1{F0: "hello"}
    v2 = 3.14159
    if false {
        var v5 float32 = (v2 / (0.1 / 0.1))
        _ = v5
    } else {
        if true {
            v3 -= (v3 | (v3 & (^(v3))))
            fmt.Println(v2, v3, v1)
        }
    }
    switch "" {
    case "":
        fmt.Println(v1, v1, v3)
    case "hello":
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS2{F0: "hello", F1: 0}
            v8 := []int{fi6}
            var v9 int16 = 32767
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    case "\n\t":
        if ((!true) || true) {
            var v10 int8 = 127
            _ = v10
        }
    default:
        switch 2147483647 {
        case 128:
            v11 := GenS3{F0: 2}
            _ = v11
        case 157:
            v1 -= (v1 ^ 42)
        default:
            v12 := GenS2{F0: (("foo bar" + ((("hello" + "\n\t") + "foo bar") + "hello")) + ("" + "a")), F1: -353}
            _ = v12
        }
    }
    v13 := []int{-2147483648}
    v3 = (v3 >> ((v3) & 63))
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 uint32 = 0
        v16 := GenS2{F0: "\n\t", F1: 42}
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS4{F0: 128, F1: 42, F2: -1}
    v18 := GenS4{F0: 128, F1: -1, F2: -410}
    var v19 float32 = ((v2 / (0.0 + v2)) / v2)
    v13 = append(v13, 127)
    v20 := GenS4{F0: 9223372036854775807, F1: -623, F2: 942}
    v1 -= -90
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 5; fi22++ {
            if len(v13) > 0 {
                v13[0] = (fi21 << (((v3 - v3)) & 63))
            }
            fmt.Println(v3, fi22, fi22)
            v19 -= float32((-2.25 * (0.0 - 3.14159)))
            v19 = ((0.1 - 1.5) * v2)
            switch (^(fi21)) {
            case 9223372036854775807:
                v19 = (v2 / ((-2.25 * v19) + v2))
            case 1:
                var v23 uint16 = 414
                v24 := GenS1{F0: ("foo bar" + "hello")}
                _ = len(v13)
                _ = v23
                _ = v24
            }
            _ = fi22
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 string = ("a" + "foo bar")
            var v27 uint64 = 18446744073709551615
            v28 := GenS4{F0: (fi21 - 2147483647), F1: (fi25 + (^(fi21))), F2: 9223372036854775807}
            var v29 bool = (float64((v2 + (v19 / v2))) > ((1e10 - 1.5) + 1e10))
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
            _ = v29
        }
        switch "a" {
        case "a":
            if false {
                v3 -= v3
            }
        }
        _ = fi21
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
