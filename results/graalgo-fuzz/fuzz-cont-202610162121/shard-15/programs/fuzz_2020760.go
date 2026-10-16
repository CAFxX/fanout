package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    var v3 float32 = (3.14159 / ((1.5 - 0.1) - (1.5 * -2.25)))
    var v4 int64 = 0
    if true {
        for fi5 := 0; fi5 < 4; fi5++ {
            v2 = v2
            _ = fi5
        }
        switch ("hello" + ("hello" + ("a" + "a"))) {
        case "a":
            fmt.Println(v2)
        case "hello":
            var v6 uint8 = 2
            _ = v6
        case "\n\t":
            v7 := GenS1{F0: 42, F1: 42}
            _ = v7
        default:
            if ("hello" < "") {
                v8 := GenS1{F0: 2, F1: 532}
                if v2 {
                    v2 = ((((-1 >= 127) && (42 != 2147483647)) || v2) || (42 <= 284))
                    fmt.Println(v2, v3)
                    v3 = v3
                } else {
                    var v9 uint32 = 2147483648
                    _ = v9
                }
                _ = v8
            } else {
                if (2 == 1) {
                    v3 -= 0.0
                    var v10 float32 = v3
                    _ = v10
                }
                var v11 int = 480
                _ = v11
            }
        }
    }
    if false {
        v12 := GenS2{F0: 58, F1: 212}
        v3 -= 1.5
        _ = v12
    }
    var v13 string = "hello"
    v14 := GenS2{F0: -1, F1: 42}
    delete(v1, 1)
    v3 = 0.1
    var v15 int32 = 1
    switch (("foo bar" + v13) + v13) {
    case "":
        v16 := []int{}
        _ = v16
    case "foo bar":
        if (218 == 0) {
            v17 := GenS3{F0: 781}
            if (18446744073709551615 <= 65535) {
                v2 = (v4 == (v4 << ((89) & 63)))
                v15 += v15
                var v18 uint64 = 752
                _ = v18
            }
            _ = v17
        } else {
            v13 = ((v13 + (v13 + ("a" + v13))) + v13)
            v4 = v4
        }
    }
    v19 := GenS3{F0: 1}
    v20 := []int{}
    fmt.Println(v2)
    v2 = v2
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 uint64 = 42
        var v23 int8 = 127
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 float32 = float32(-2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
