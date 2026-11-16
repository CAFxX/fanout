package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint16; F1 string; F2 int16 }

func main() {
    var v1 uint32 = 440
    var v2 uint64 = 18446744073709551615
    var v3 float32 = 3.14159
    var v4 float64 = 0.1
    var v5 int32 = 1
    v4 = v4
    v5 = v5
    v6 := []int{-616, 127, -975, 127}
    v7 := GenS1{F0: 65535}
    fmt.Println(v4, v4, v1)
    var v8 uint8 = 1
    if true {
        if len(v6) > 0 {
            v6[0] = 42
        }
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(v5, v4, v5)
        _ = fi9
    }
    v8 = (v8 % ((v8) | 1))
    for fi10 := 0; fi10 < 2; fi10++ {
        v11 := GenS1{F0: 2}
        v12 := GenS2{F0: 32767, F1: 32767}
        v13 := GenS1{F0: 65535}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    if (!(255 != 136)) {
        switch 128 {
        case -894:
            for fi14 := 0; fi14 < 15; fi14++ {
                v15 := GenS2{F0: 2, F1: (-(fi14))}
                for fi16 := 0; fi16 < 12; fi16++ {
                    if true {
                        var v17 int8 = -42
                        switch "\n\t" {
                        case "\n\t":
                            var v18 uint32 = (v1 ^ (v1 >> ((824) & 63)))
                            _ = v18
                        case "a":
                            v8 = 255
                        case "hello":
                            if true {
                                var v19 uint64 = v2
                                switch "\n\t" {
                                case "hello":
                                    switch (-(fi16)) {
                                    case 0:
                                        var v20 int = 42
                                        _ = v20
                                    }
                                case "":
                                    v5 -= v5
                                }
                                _ = v19
                            }
                        default:
                            var v21 int32 = v5
                            v22 := GenS2{F0: 1, F1: fi14}
                            var v23 float64 = v4
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        _ = v17
                    }
                    switch 9223372036854775807 {
                    case -965:
                        if false {
                            _ = len(v6)
                            fmt.Println(v5)
                        }
                        v4 = v4
                    case 9223372036854775807:
                        v2 += v2
                    case -1:
                        v24 := GenS1{F0: 2}
                        _ = v24
                    default:
                        var v25 int8 = -1
                        _ = v25
                    }
                    v4 = 3.14159
                    var v26 string = "foo bar"
                    if false {
                        if (true && (!true)) {
                            var v27 int16 = 1
                            _ = v27
                        } else {
                            v26 = v26
                            var v28 int32 = (v5 ^ v5)
                            _ = v28
                        }
                        v29 := GenS1{F0: uint(v8)}
                        _ = v29
                    } else {
                        _ = len(v6)
                        _ = len(v6)
                        var v30 string = v26
                        if true {
                            var v31 int8 = 127
                            v2 = 10
                            _ = v31
                        }
                        _ = v30
                    }
                    _ = fi16
                    _ = v26
                }
                v32 := GenS3{F0: 687, F1: "foo bar", F2: 606}
                _ = fi14
                _ = v15
                _ = v32
            }
        default:
            v33 := GenS1{F0: 0}
            _ = v33
        }
    }
    fmt.Println(v2)
    v34 := GenS2{F0: 558, F1: -9223372036854775808}
    _ = len(v6)
    var v35 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
