package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint32; F1 int64; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 int8; F1 int64; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"foo bar", ("a" + ("hello" + "\n\t")), ("foo bar" + ((("hello" + ("hello" + ("foo bar" + "hello"))) + "hello") + ((((("hello" + ("foo bar" + "hello")) + "a") + "") + (("" + "\n\t") + "foo bar")) + ("hello" + ("a" + "\n\t")))))}
    var v3 int64 = 1
    v4 := make(map[int]string)
    var v5 bool = false
    var v6 uint = 252
    v7 := GenS1{F0: 1, F1: -2147483648}
    switch "hello" {
    case "a":
        for fi8 := 0; fi8 < 2; fi8++ {
            v3 += v3
            v5 = (!false)
            if true {
                switch ("hello" + "") {
                case "a":
                    if v5 {
                        v6 -= v6
                        v9 := GenS1{F0: -1, F1: 42}
                        _ = v9
                    } else {
                        var v10 int32 = 263
                        _ = v10
                    }
                }
            } else {
                switch int(v6) {
                case 1:
                    for fi11 := 0; fi11 < 20; fi11++ {
                        v12 := GenS2{F0: 2147483647, F1: -568, F2: 0}
                        _ = fi11
                        _ = v12
                    }
                }
                for fi13 := 0; fi13 < 16; fi13++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v14 := GenS3{F0: 2, F1: 255}
                    v15 := GenS4{F0: -111, F1: 1, F2: 598}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v16 uint8 = 0
                _ = v16
            }
            _ = fi8
        }
    }
    v17 := v1[912]
    v5 = v5
    v1[-1] = 2147483647
    switch "a" {
    case "foo bar":
        if ((-(v3)) != v3) {
            var v18 int32 = 0
            _ = v18
        }
    case "a":
        var v19 int64 = v3
        _ = v19
    }
    switch "a" {
    case "a":
        var v20 uint64 = 187
        _ = v20
    default:
        if len(v2) > 0 {
            v2[0] = "a"
        }
    }
    fmt.Println(v5, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
