package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int; F1 uint8; F2 uint }
type GenS3 struct { F0 int16; F1 bool }

func main() {
    var v1 int8 = 42
    v2 := []int{}
    v3 := make(map[int]bool)
    v4 := GenS1{F0: (-1 > -1), F1: 0}
    var v5 uint64 = 4294967295
    switch ("" + "\n\t") {
    case "":
        if true {
            v1 -= v1
            if false {
                v6 := GenS1{F0: (127 == 255), F1: v5}
                var v7 int16 = int16(v5)
                _ = v6
                _ = v7
            } else {
                v8 := GenS2{F0: 32767, F1: 127, F2: 0}
                _ = v8
            }
        }
    }
    if (false && false) {
        var v9 float64 = ((3.14159 - 1.5) - (float64((1e10 * 3.14159)) / (float64(0.1) - float64((-2.25 + -2.25)))))
        var v10 uint64 = v5
        var v11 string = "foo bar"
        _ = v10
        _ = v11
        _ = v9
    }
    fmt.Println(v5)
    if (-1 > 1000) {
        for fi12 := 0; fi12 < 18; fi12++ {
            for fi13 := 0; fi13 < 7; fi13++ {
                switch "foo bar" {
                case "hello":
                    switch (fi12 / ((fi12) | 1)) {
                    case 4294967295:
                        v14 := GenS2{F0: fi13, F1: 127, F2: 242}
                        _ = v14
                    case 0:
                        var v15 uint8 = 2
                        _ = v15
                    case 1:
                        v16 := GenS1{F0: true, F1: (v5 + v5)}
                        _ = v16
                    default:
                        var v17 uint64 = 18446744073709551615
                        _ = v17
                    }
                case "a":
                    v5 += (v5 ^ v5)
                }
                var v18 string = ("\n\t" + ("foo bar" + (("hello" + "a") + "hello")))
                var v19 uint16 = 3
                _ = fi13
                _ = v18
                _ = v19
            }
            _ = len(v3)
            _ = fi12
        }
        switch "\n\t" {
        case "":
            var v20 bool = (!true)
            _ = v20
        case "\n\t":
            var v21 bool = false
            _ = v21
        case "hello":
            if true {
                switch 183 {
                case 652:
                    for fi22 := 0; fi22 < 13; fi22++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi22
                    }
                }
            }
        default:
            fmt.Println(v5, v1)
        }
        _ = len(v3)
    }
    if (true || (255 >= 2)) {
        var v23 uint64 = v5
        _ = v23
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v24 uint64 = (v5 ^ 18446744073709551615)
    v1 = -1
    v24 -= v5
    _ = len(v2)
    v5 = v5
    v25 := GenS3{F0: -32768, F1: ((!false) || ("a" >= "hello"))}
    var v26 string = "\n\t"
    var v27 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
