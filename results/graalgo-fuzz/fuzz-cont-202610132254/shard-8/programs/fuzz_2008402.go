package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int8; F2 string }
type GenS3 struct { F0 string; F1 string; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int64; F2 bool }
type GenS5 struct { F0 int16; F1 bool }

func main() {
    var v1 int16 = 42
    v2 := []string{(("a" + "hello") + "")}
    v3 := make(map[int]int)
    v4 := []uint32{65535}
    var v5 int = 2147483647
    var v6 uint16 = 350
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 float64 = ((0.1 / ((3.14159 / 1.0) / (1.5 - 1e10))) / 3.14159)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: -264}
    fmt.Println(v5, v5, v5)
    var v11 uint16 = 32767
    v12 := GenS1{F0: 9223372036854775807}
    for fi13 := 0; fi13 < 2; fi13++ {
        switch "" {
        case "hello":
            v14 := GenS2{F0: v1, F1: -52, F2: ("a" + "")}
            _ = v14
        case "a":
            for fi15 := 0; fi15 < 5; fi15++ {
                if true {
                    v6 = 255
                    v16 := GenS3{F0: ("" + "\n\t"), F1: "a", F2: 225}
                    _ = v16
                }
                switch ((("a" + "a") + "hello") + "") {
                case "\n\t":
                    var v17 uint64 = 577
                    _ = v17
                default:
                    var v18 int16 = v1
                    _ = v18
                }
                var v19 int8 = 100
                switch -1 {
                case 127:
                    if true {
                        if true {
                            if (4294967295 >= 2) {
                                var v20 uint8 = 255
                                var v21 int8 = v19
                                var v22 uint32 = 466
                                _ = v20
                                _ = v21
                                _ = v22
                            } else {
                                _ = len(v3)
                                if false {
                                    var v23 int64 = 1
                                    v4 = append(v4, 65535)
                                    _ = v23
                                }
                                v24 := GenS4{F0: 128, F1: -9223372036854775808, F2: true}
                                _ = v24
                            }
                            v25 := GenS3{F0: "foo bar", F1: "", F2: 42}
                            var v26 uint8 = 223
                            _ = v25
                            _ = v26
                        }
                    }
                }
                _ = fi15
                _ = v19
            }
        case "foo bar":
            var v27 float64 = 0.1
            _ = v27
        }
        _ = fi13
    }
    switch (("hello" + ("" + "hello")) + "a") {
    case "foo bar":
        if true {
            v11 = v11
            delete(v3, 327)
            fmt.Println(v5, v6)
        }
    case "a":
        v28 := GenS5{F0: 424, F1: false}
        _ = v28
    default:
        var v29 uint64 = 1
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
