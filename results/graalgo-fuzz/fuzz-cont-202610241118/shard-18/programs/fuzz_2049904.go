package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint8 }

func main() {
    v1 := []uint32{2147483647, 2147483648, 2147483647}
    var v2 int32 = 1
    var v3 int = 42
    var v4 uint64 = 4294967295
    for fi5 := 0; fi5 < 12; fi5++ {
        v4 -= (v4 % (((v4 ^ (^(v4)))) | 1))
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 float64 = 0.1
            v8 := GenS1{F0: 0}
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 float32 = 1.5
        _ = fi5
        _ = v9
    }
    v10 := GenS1{F0: 0}
    v2 = 2147483647
    switch "\n\t" {
    case "a":
        switch "hello" {
        case "hello":
            var v11 int = 1
            _ = v11
        case "\n\t":
            if true {
                switch "" {
                case "a":
                    for fi12 := 0; fi12 < 3; fi12++ {
                        switch int(v4) {
                        case -356:
                            for fi13 := 0; fi13 < 18; fi13++ {
                                var v14 string = "a"
                                _ = fi13
                                _ = v14
                            }
                        default:
                            var v15 int64 = -959
                            _ = v15
                        }
                        _ = fi12
                    }
                case "foo bar":
                    var v16 int32 = (v2 - 52)
                    _ = v16
                }
                v2 = v2
                v17 := GenS1{F0: 42}
                _ = v17
            } else {
                v18 := GenS1{F0: 65535}
                if true {
                    fmt.Println(v4, v3)
                } else {
                    var v19 uint32 = 1
                    var v20 uint = 0
                    _ = v19
                    _ = v20
                }
                switch "\n\t" {
                case "foo bar":
                    v21 := GenS1{F0: 886}
                    _ = v21
                case "a":
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                default:
                    v4 -= 403
                }
                _ = v18
            }
        default:
            v4 += v4
        }
    case "\n\t":
        v3 += int(v2)
    default:
        var v22 int16 = 42
        _ = v22
    }
    if len(v1) > 0 {
        v1[0] = 36
    }
    v4 -= 18446744073709551615
    var v23 int8 = -56
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS2{F0: v4}
        fmt.Println(v4, v3)
        for fi26 := 0; fi26 < 17; fi26++ {
            fmt.Println(fi24)
            v3 -= (-(fi26))
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        v1 = append(v1, 0)
        for fi28 := 0; fi28 < 18; fi28++ {
            v1 = append(v1, 42)
            var v29 int16 = -32768
            v30 := GenS2{F0: v4}
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi27
    }
    switch ("\n\t" + "hello") {
    case "\n\t":
        fmt.Println(v23)
    }
    if true {
        var v31 float32 = 3.14159
        var v32 int64 = 42
        _ = v31
        _ = v32
    } else {
        var v33 int64 = 1
        if len(v1) > 0 {
            v1[0] = 616
        }
        _ = v33
    }
    v3 = v3
    _ = len(v1)
    var v34 uint32 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v34 += 65535
    v35 := GenS3{F0: 18446744073709551615, F1: 869, F2: 2}
    var v36 uint64 = (v4 ^ 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
