package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := []int{-2147483648}
    var v3 bool = (!true)
    var v4 int8 = -44
    v5 := GenS1{F0: "\n\t"}
    v6 := GenS2{F0: 42}
    var v7 int64 = 42
    v2 = append(v2, -515)
    var v8 bool = (42 > 1000)
    for fi9 := 0; fi9 < 16; fi9++ {
        switch (("\n\t" + "a") + "a") {
        case "a":
            for fi10 := 0; fi10 < 11; fi10++ {
                fmt.Println(fi10, v3)
                if v8 {
                    switch "a" {
                    case "a":
                        if ((^(fi9)) <= fi10) {
                            v4 -= (v4 & v4)
                            v7 = -1
                            var v11 uint16 = 149
                            delete(v1, ("hello" + "hello"))
                            switch fi9 {
                            case 4294967295:
                                v12 := GenS1{F0: "a"}
                                _ = v12
                            case -590:
                                v13 := GenS3{F0: 42, F1: false}
                                _ = v13
                            }
                            _ = v11
                        }
                    case "foo bar":
                        if ((v7 - v7) == (v7 % ((42) | 1))) {
                            var v14 uint32 = 381
                            fmt.Println(fi9)
                            switch "\n\t" {
                            case "a":
                                var v15 int16 = 1
                                _ = v15
                            case "hello":
                                var v16 uint = 239
                                _ = v16
                            case "foo bar":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            _ = v14
                        } else {
                            var v17 float32 = 3.14159
                            v3 = (2147483647 > -1)
                            v3 = false
                            _ = v17
                        }
                    case "":
                        if ((-2.25 <= (3.14159 + 1e10)) && ((fi9 + (fi9 & fi9)) == fi9)) {
                            var v18 uint = 0
                            _ = v18
                        } else {
                            fmt.Println(fi10, v8, v8)
                        }
                    default:
                        switch "hello" {
                        case "a":
                            switch -1 {
                            case 2147483647:
                                v19 := GenS4{F0: 42, F1: 589, F2: 571}
                                _ = v19
                            }
                        case "hello":
                            v8 = v3
                        default:
                            delete(v1, ("a" + "hello"))
                        }
                    }
                }
                var v20 int = 32767
                _ = fi10
                _ = v20
            }
        case "foo bar":
            if v3 {
                v21 := GenS3{F0: 176, F1: (-128 < -1)}
                _ = v21
            }
        case "":
            v22 := v1[""]
            _ = v22
        }
        v23 := GenS2{F0: 191}
        _ = fi9
        _ = v23
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        v25 := GenS2{F0: v7}
        v4 += int8(fi24)
        v3 = (v8 && v8)
        _ = fi24
        _ = v25
    }
    switch "" {
    case "a":
        for fi26 := 0; fi26 < 10; fi26++ {
            if (1000 > -32768) {
                var v27 bool = v8
                var v28 bool = v3
                _ = v27
                _ = v28
            }
            _ = fi26
        }
    default:
        v4 = 53
    }
    v4 = 1
    if (v4 == v4) {
        var v29 int32 = 468
        v30 := GenS4{F0: 1, F1: (v29 | (v29 / ((v29) | 1))), F2: 269}
        v31 := GenS5{F0: 1000}
        _ = v29
        _ = v30
        _ = v31
    }
    if (345 >= 503) {
        v7 -= (v7 ^ (v7 | v7))
        var v32 int = 128
        _ = v32
    }
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
