package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 string; F2 int64 }
type GenS4 struct { F0 int32; F1 int32 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 float32 = 3.14159
    v4 := make(map[string]string)
    v5 := GenS1{F0: 42, F1: 1, F2: 18446744073709551615}
    delete(v4, "")
    var v6 int8 = 0
    switch 580 {
    case 4294967295:
        for fi7 := 0; fi7 < 19; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                if true {
                    v9 := GenS2{F0: 26}
                    if false {
                        switch (fi8 - fi8) {
                        case -2147483648:
                            switch "" {
                            case "hello":
                                v6 = v6
                            case "a":
                                var v10 uint64 = 2
                                switch fi8 {
                                case 2147483647:
                                    var v11 int32 = 95
                                    _ = v11
                                case -9223372036854775808:
                                    if (1e10 < float64(0.1)) {
                                        switch ("hello" + "hello") {
                                        case "":
                                            var v12 float32 = (v3 * v3)
                                            _ = v12
                                        }
                                        _ = len(v4)
                                    } else {
                                        v3 = ((0.0 / (1.5 + v3)) + v3)
                                        var v13 int64 = 9223372036854775807
                                        v14 := GenS2{F0: 255}
                                        switch fi8 {
                                        case -219:
                                            var v15 float64 = 1e10
                                            _ = v15
                                        case 4294967295:
                                            if true {
                                                var v16 uint8 = 2
                                                _ = v16
                                            }
                                        }
                                        _ = v13
                                        _ = v14
                                    }
                                }
                                _ = len(v1)
                                switch fi7 {
                                case 42:
                                    fmt.Println(v3, v6, v10)
                                }
                                _ = v10
                            case "foo bar":
                                switch ("\n\t" + ("a" + "hello")) {
                                case "":
                                    v17 := GenS1{F0: -1, F1: 0, F2: 21}
                                    _ = v17
                                }
                            default:
                                v18 := GenS2{F0: 128}
                                _ = v18
                            }
                        case 4294967295:
                            v6 = -128
                        default:
                            if false {
                                var v19 uint64 = 4294967295
                                var v20 float64 = 1.5
                                _ = v19
                                _ = v20
                            }
                        }
                    }
                    _ = v9
                } else {
                    if (-239 >= 1000) {
                        v3 = -2.25
                        v21 := []int{(fi7 ^ (fi8 >> ((805) & 63))), (fi7 << ((18446744073709551615) & 63)), fi8}
                        fmt.Println(fi7, fi8)
                        _ = v21
                    }
                    var v22 float64 = 3.14159
                    _ = v22
                }
                v3 += (v3 - v3)
                _ = fi8
            }
            _ = fi7
        }
    case 127:
        fmt.Println(v6)
    case -9223372036854775808:
        if false {
            var v23 int16 = 822
            _ = len(v2)
            _ = v23
        }
    }
    v24 := GenS1{F0: -32768, F1: 2, F2: 1}
    _ = len(v2)
    var v25 int32 = 1
    var v26 uint64 = 0
    v27 := GenS3{F0: 32767, F1: "", F2: 890}
    for fi28 := 0; fi28 < 20; fi28++ {
        var v29 bool = (128 <= 255)
        _ = fi28
        _ = v29
    }
    var v30 uint = 0
    v31 := GenS4{F0: (v25 ^ v25), F1: (v25 ^ (v25 - (v25 >> ((v30) & 63))))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
