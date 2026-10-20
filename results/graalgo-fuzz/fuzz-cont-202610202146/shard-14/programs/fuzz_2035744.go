package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }
type GenS2 struct { F0 uint8; F1 string; F2 int64 }
type GenS3 struct { F0 int; F1 bool }

func main() {
    var v1 int = 185
    var v2 int = (v1 + v1)
    var v3 int8 = -1
    v4 := []string{("foo bar" + "hello")}
    v1 -= 380
    var v5 uint = 0
    if false {
        v3 -= v3
        for fi6 := 0; fi6 < 12; fi6++ {
            if true {
                fmt.Println(v5)
                if true {
                    if true {
                        v7 := GenS1{F0: 1, F1: 2, F2: (!false)}
                        _ = v7
                    }
                } else {
                    for fi8 := 0; fi8 < 3; fi8++ {
                        fmt.Println(fi6)
                        v3 = 81
                        var v9 uint = (v5 | (^(v5)))
                        _ = fi8
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 1; fi10++ {
                        v11 := GenS2{F0: 2, F1: (((("a" + ("hello" + "")) + "foo bar") + ("hello" + ("foo bar" + "hello"))) + "a"), F2: 1}
                        var v12 uint16 = 1
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    var v13 int32 = 42
                    _ = v13
                }
            }
            v14 := GenS2{F0: 255, F1: "a", F2: 9223372036854775807}
            _ = fi6
            _ = v14
        }
        switch (("hello" + "hello") + "") {
        case "hello":
            fmt.Println(v5, v2)
        case "foo bar":
            for fi15 := 0; fi15 < 5; fi15++ {
                v2 = v2
                _ = len(v4)
                _ = fi15
            }
        default:
            var v16 uint64 = 1
            _ = v16
        }
    } else {
        v1 = v2
    }
    if false {
        v17 := GenS3{F0: (v1 + (v1 + v2)), F1: true}
        fmt.Println(v2)
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 uint16 = 348
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    _ = len(v4)
    if (2147483647 < -195) {
        if false {
            v20 := GenS2{F0: 1, F1: "", F2: 9223372036854775807}
            var v21 int8 = (v3 & 127)
            if false {
                var v22 uint16 = 2
                if true {
                    var v23 uint = (v5 + v5)
                    v24 := GenS2{F0: 128, F1: "", F2: -1}
                    v25 := GenS2{F0: 2, F1: "\n\t", F2: -9223372036854775808}
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    _ = len(v4)
                }
                v5 = v5
                _ = v22
            } else {
                var v26 int8 = (-(v21))
                if false {
                    fmt.Println(v1)
                    var v27 float32 = ((1.5 - (3.14159 / (3.14159 + 0.1))) - -2.25)
                    _ = v27
                }
                var v28 int16 = -238
                _ = v26
                _ = v28
            }
            _ = v20
            _ = v21
        }
    }
    v5 += (v5 ^ 4294967295)
    if true {
        var v29 int = (v1 & (v2 - v2))
        _ = len(v4)
        var v30 int32 = 1
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
