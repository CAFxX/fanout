package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 uint8 }
type GenS3 struct { F0 int64; F1 string }

func main() {
    v1 := []int{-9223372036854775808, 42}
    v2 := make(map[uint32]uint32)
    var v3 string = "a"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        v3 = v3
        v4 := GenS1{F0: 462}
        if true {
            if true {
                switch ((v3 + v3) + "\n\t") {
                case "\n\t":
                    fmt.Println(v3)
                case "hello":
                    v5 := GenS2{F0: 12, F1: 619, F2: 226}
                    _ = v5
                }
            }
            var v6 int8 = 100
            v6 += v6
            _ = v6
        } else {
            v7 := GenS1{F0: 1}
            v3 = ((("foo bar" + "\n\t") + v3) + ((v3 + v3) + ""))
            var v8 int = 865
            _ = v7
            _ = v8
        }
        _ = v4
    } else {
        var v9 string = ((("a" + "foo bar") + ("foo bar" + (v3 + v3))) + v3)
        var v10 int64 = 92
        _ = v10
        _ = v9
    }
    if false {
        fmt.Println(v3)
    } else {
        v11 := GenS1{F0: 836}
        _ = v11
    }
    _ = len(v1)
    v12 := v2[828]
    v12 = (v12 + (v12 / (((v12 * v12)) | 1)))
    var v13 int64 = 42
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 uint = 42
        switch ("\n\t" + "") {
        case "hello":
            var v16 int8 = 29
            _ = v16
        default:
            for fi17 := 0; fi17 < 5; fi17++ {
                v18 := GenS3{F0: v13, F1: "a"}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v3, fi14)
                _ = fi17
                _ = v18
            }
        }
        var v19 bool = (282 > 862)
        _ = fi14
        _ = v15
        _ = v19
    }
    delete(v2, (v12 * (v12 | (v12 + (v12 >> ((255) & 63))))))
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 int32 = 0
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        v2[v12] = 2147483648
        _ = fi22
    }
    var v23 float64 = (3.14159 - 3.14159)
    var v24 float64 = (0.1 / 1.5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
