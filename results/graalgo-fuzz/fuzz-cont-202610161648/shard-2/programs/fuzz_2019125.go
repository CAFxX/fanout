package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int8 }

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]int)
    var v3 uint32 = 1
    v4 := []string{}
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]bool)
    v7 := GenS1{F0: 729}
    if (1e10 > 1e10) {
        v8 := GenS2{F0: -310, F1: 42, F2: 511}
        v3 -= (^(v3))
        _ = v8
    } else {
        v1 -= uint(v3)
        if false {
            v1 = v1
        } else {
            v3 += 2
        }
        fmt.Println(v1)
    }
    var v9 uint64 = 0
    v3 = v3
    var v10 int32 = 2147483647
    if (1 < 9223372036854775807) {
        for fi11 := 0; fi11 < 5; fi11++ {
            for fi12 := 0; fi12 < 20; fi12++ {
                if false {
                    switch "" {
                    case "hello":
                        if (!((1.5 * 0.0) > 3.14159)) {
                            v13 := GenS3{F0: (v1 | (^(v1))), F1: true}
                            if false {
                                if true {
                                    v9 -= 198
                                }
                                if false {
                                    v3 -= (v3 ^ (v3 * v3))
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = "foo bar"
                                    }
                                    var v14 uint8 = 1
                                    _ = v14
                                }
                                var v15 bool = true
                                _ = v15
                            } else {
                                fmt.Println(fi12, v10, v1)
                            }
                            _ = v13
                        }
                    }
                    switch (fi11 >> (((v1 - (v1 - 18446744073709551615))) & 63)) {
                    case -1:
                        _ = len(v4)
                        var v16 int32 = v10
                        var v17 int16 = 227
                        _ = v16
                        _ = v17
                    case 42:
                        v9 = v9
                    }
                    v18 := GenS4{F0: "hello", F1: 4294967295}
                    _ = v18
                }
                _ = fi12
            }
            switch "" {
            case "hello":
                v10 += 941
            case "\n\t":
                if true {
                    fmt.Println(v9, v3)
                }
            case "":
                fmt.Println(v9, v10)
            }
            v19 := GenS2{F0: 4294967295, F1: 1, F2: (v10 - int32(v9))}
            _ = fi11
            _ = v19
        }
        v1 = v1
        v20 := GenS5{F0: -128}
        _ = v20
    } else {
        v21 := GenS2{F0: 2147483647, F1: 2, F2: v10}
        _ = v21
    }
    if len(v4) > 0 {
        v4[0] = ((("" + ("a" + "foo bar")) + "foo bar") + "foo bar")
    }
    _ = len(v4)
    v22 := GenS2{F0: 1, F1: 255, F2: -1}
    fmt.Println(v9)
    var v23 float32 = float32(0.1)
    var v24 int64 = 818
    if true {
        v23 = v23
        v25 := GenS2{F0: 128, F1: 0, F2: (^(v10))}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
