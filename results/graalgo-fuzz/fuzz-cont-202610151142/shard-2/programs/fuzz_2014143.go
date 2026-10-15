package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int64 }
type GenS2 struct { F0 int8; F1 string; F2 bool }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int16 = 42
    var v2 int64 = -9223372036854775808
    var v3 float64 = 1e10
    var v4 uint32 = 2147483648
    var v5 bool = true
    var v6 uint = 4294967295
    if v5 {
        var v7 uint16 = 3
        v8 := GenS1{F0: -1, F1: v1, F2: (v2 * -454)}
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                v11 := []int{fi10, fi9, (fi9 & fi10)}
                fmt.Println(v4, v2, v4)
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 10; fi13++ {
                fmt.Println(fi13, fi12, fi13)
                if ("" >= "foo bar") {
                    if (v4 >= 137) {
                        var v14 uint32 = v4
                        v15 := GenS1{F0: -656, F1: v1, F2: -963}
                        _ = v14
                        _ = v15
                    }
                    var v16 float32 = (1.5 * 3.14159)
                    if true {
                        var v17 int = (fi12 | 128)
                        v18 := GenS1{F0: -622, F1: v1, F2: int64(v4)}
                        _ = v17
                        _ = v18
                    } else {
                        if (v1 < 32767) {
                            v19 := GenS1{F0: int32(v6), F1: 32767, F2: (v2 + v2)}
                            switch ("" + ("hello" + "hello")) {
                            case "a":
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v3, fi13)
                                case "hello":
                                    fmt.Println(v6, v5, v1)
                                case "\n\t":
                                    v20 := []int{(^(fi12)), fi13}
                                    _ = v20
                                default:
                                    var v21 uint = v6
                                    _ = v21
                                }
                                switch (-(fi12)) {
                                case 303:
                                    switch (("" + "\n\t") + "hello") {
                                    case "\n\t":
                                        v22 := []int{(fi13 / ((fi13) | 1)), 42, (fi13 & (^(fi13)))}
                                        _ = v22
                                    }
                                case 870:
                                    v23 := GenS1{F0: int32(fi13), F1: int16(v6), F2: 1}
                                    _ = v23
                                }
                                v24 := []int{}
                                _ = v24
                            }
                            _ = v19
                        } else {
                            var v25 int32 = -1
                            fmt.Println(v6)
                            _ = v25
                        }
                        v26 := []int{}
                        fmt.Println(fi13, v5)
                        v2 = int64(v4)
                        v6 = (v6 >> ((471) & 63))
                        _ = v26
                    }
                    v5 = true
                    v27 := []int{(fi12 | fi13)}
                    _ = v16
                    _ = v27
                } else {
                    v28 := GenS2{F0: -128, F1: "", F2: (!v5)}
                    fmt.Println(v4, v1, v6)
                    _ = v28
                }
                var v29 uint64 = 18446744073709551615
                if (!v5) {
                    v2 += -1
                    var v30 string = "a"
                    _ = v30
                } else {
                    v4 = 0
                }
                fmt.Println(v4, v2, v6)
                _ = fi13
                _ = v29
            }
            _ = fi12
        }
    }
    v31 := GenS3{F0: 324}
    fmt.Println(v3)
    var v32 uint32 = v4
    v1 -= (v1 << (((v6 % ((v6) | 1))) & 63))
    v4 = v32
    v1 = 697
    switch "a" {
    case "foo bar":
        v4 -= v32
    }
    var v33 uint = uint(v32)
    v33 = 343
    v34 := []int{32768, 42}
    var v35 string = ("foo bar" + "foo bar")
    fmt.Println(v4, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
