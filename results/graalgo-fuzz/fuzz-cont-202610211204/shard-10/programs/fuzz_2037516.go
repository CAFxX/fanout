package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int; F1 uint64; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 float32 = -2.25
    var v2 uint32 = 2147483648
    var v3 float32 = v1
    var v4 uint8 = 201
    switch 32768 {
    case -9223372036854775808:
        v5 := GenS1{F0: -170, F1: "foo bar"}
        _ = v5
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 string = "hello"
        _ = fi6
        _ = v7
    }
    var v8 int = 32768
    switch (v8 | 4294967295) {
    case 41:
        v3 = float32(3.14159)
    default:
        var v9 int16 = 0
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{(v8 - (^(v8))), (v8 * -9223372036854775808)}
        for fi12 := 0; fi12 < 18; fi12++ {
            switch "foo bar" {
            case "a":
                switch ("foo bar" + "foo bar") {
                case "hello":
                    if true {
                        var v13 int8 = 127
                        _ = v13
                    } else {
                        var v14 uint32 = v2
                        if len(v11) > 0 {
                            v11[0] = 42
                        }
                        _ = v14
                    }
                default:
                    fmt.Println(v1)
                }
            }
            _ = len(v11)
            fmt.Println(v3, fi10)
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    v2 = (^(v2))
    v15 := GenS1{F0: -1, F1: "\n\t"}
    fmt.Println(v3)
    var v16 int8 = -85
    switch -330 {
    case 4294967295:
        var v17 uint64 = 1
        _ = v17
    case -736:
        v18 := GenS2{F0: (v8 / ((v8) | 1)), F1: 1, F2: v16}
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        if true {
            if true {
                var v20 int32 = 42
                _ = v20
            } else {
                switch (("\n\t" + ("a" + "\n\t")) + "\n\t") {
                case "\n\t":
                    v21 := GenS3{F0: "\n\t"}
                    _ = v21
                default:
                    if true {
                        v22 := []int{(v8 >> ((355) & 63)), 1, -461, (v8 >> ((4294967295) & 63))}
                        _ = v22
                    }
                }
                for fi23 := 0; fi23 < 8; fi23++ {
                    var v24 uint16 = 37
                    fmt.Println(fi23, v3)
                    _ = fi23
                    _ = v24
                }
                var v25 int8 = (v16 & v16)
                _ = v25
            }
            var v26 float32 = v3
            v16 += 100
            _ = v26
        } else {
            v27 := GenS4{F0: 18446744073709551615, F1: 42}
            var v28 int32 = 2147483647
            fmt.Println(v2)
            _ = v27
            _ = v28
        }
        _ = fi19
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        v30 := []int{(v8 * v8), (v8 + v8)}
        fmt.Println(v1, fi29, v3)
        _ = fi29
        _ = v30
    }
    v8 = (v8 + (-(v8)))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
