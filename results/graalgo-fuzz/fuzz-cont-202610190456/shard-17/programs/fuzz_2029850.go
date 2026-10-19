package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    var v3 float32 = 1.5
    var v4 int16 = 1
    var v5 int8 = -1
    var v6 uint64 = 18446744073709551615
    _ = len(v1)
    if (!(-311 == -301)) {
        switch ("foo bar" + (("" + "hello") + (("a" + ("foo bar" + "")) + "a"))) {
        case "hello":
            for fi7 := 0; fi7 < 18; fi7++ {
                if false {
                    var v8 uint16 = 32767
                    var v9 int32 = -474
                    switch ("a" + "") {
                    case "":
                        if false {
                            var v10 int64 = 356
                            _ = v10
                        } else {
                            v6 -= 18446744073709551615
                            v5 += (v5 * v5)
                            v3 = ((0.0 + (0.0 - 1e10)) * 0.0)
                        }
                    }
                    _ = v8
                    _ = v9
                } else {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v2[0] = "hello"
                        v6 -= (^(v6))
                        _ = fi11
                    }
                    var v12 uint8 = 0
                    v13 := GenS1{F0: fi7, F1: fi7}
                    _ = v12
                    _ = v13
                }
                if (9223372036854775807 == -1) {
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v6, v4, v4)
                    }
                    fmt.Println(fi7)
                } else {
                    v3 = 0.0
                    v3 = float32((float64(v3) / -2.25))
                }
                _ = fi7
            }
        case "a":
            var v14 int64 = 9223372036854775807
            _ = v14
        case "":
            if (2 < 255) {
                v15 := GenS1{F0: 42, F1: 128}
                _ = v15
            }
        }
    }
    v2[65535] = ("a" + "foo bar")
    v3 += v3
    var v16 uint32 = 123
    v17 := GenS2{F0: 1}
    v18 := v2[(^(v16))]
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 int64 = 0
        fmt.Println(v20, v5)
        for fi21 := 0; fi21 < 7; fi21++ {
            switch "foo bar" {
            case "foo bar":
                var v22 uint64 = 841
                _ = v22
            }
            fmt.Println(v4, fi19)
            v23 := GenS2{F0: (v20 & 9223372036854775807)}
            _ = fi21
            _ = v23
        }
        _ = fi19
        _ = v20
    }
    switch ("foo bar" + ((v18 + (v18 + v18)) + v18)) {
    case "":
        var v24 uint16 = 65535
        _ = v24
    }
    var v25 uint = 499
    v18 = "foo bar"
    var v26 int64 = -9223372036854775808
    v4 += v4
    var v27 int8 = (v5 + v5)
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 bool = false
        var v30 int64 = 0
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
