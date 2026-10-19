package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int16; F2 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 18446744073709551615
    var v3 int64 = 42
    for fi4 := 0; fi4 < 6; fi4++ {
        fmt.Println(v3, v2)
        _ = fi4
    }
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 int32 = 42
            v8 := GenS1{F0: 2147483647}
            v9 := GenS1{F0: 375}
            if (1 > 1) {
                fmt.Println(fi6)
            } else {
                if (false || ((true && (0.0 == 0.1)) && (!("\n\t" == ("" + ("hello" + "\n\t")))))) {
                    v2 -= v2
                }
                v10 := GenS2{F0: "a", F1: 32767, F2: 2}
                _ = v10
            }
            switch 49 {
            case -9223372036854775808:
                v11 := GenS1{F0: 0}
                _ = v11
            case -651:
                v12 := GenS2{F0: "", F1: -32768, F2: 0}
                _ = v12
            case -526:
                if (fi5 > (fi6 % ((fi5) | 1))) {
                    if true {
                        v13 := GenS2{F0: "\n\t", F1: 0, F2: 496}
                        _ = v13
                    }
                    if (0.0 > (0.1 * 1e10)) {
                        var v14 uint64 = 2
                        _ = v14
                    } else {
                        if false {
                            if true {
                                switch "a" {
                                case "\n\t":
                                    v3 += v3
                                    var v15 int16 = 1000
                                    v2 = v2
                                    _ = v15
                                case "foo bar":
                                    v1[4294967295] = 94
                                }
                            }
                            if false {
                                var v16 uint16 = 1
                                v16 += v16
                                var v17 int = fi5
                                _ = v16
                                _ = v17
                            } else {
                                if true {
                                    fmt.Println(v2)
                                    v18 := GenS1{F0: 430}
                                    var v19 uint16 = 532
                                    _ = v18
                                    _ = v19
                                } else {
                                    var v20 float32 = 0.0
                                    fmt.Println(v20, fi6, fi5)
                                    _ = v20
                                }
                                v7 = 780
                                v7 -= v7
                                var v21 bool = true
                                _ = v21
                            }
                        }
                        v3 -= 9223372036854775807
                        switch fi6 {
                        case 4294967295:
                            v7 = v7
                        }
                    }
                    var v22 uint64 = 42
                    _ = v22
                }
            }
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        _ = len(v1)
        _ = fi5
    }
    if true {
        v3 += 42
    }
    var v23 int = 0
    v24 := GenS2{F0: "\n\t", F1: 749, F2: 2}
    fmt.Println(v3)
    var v25 int32 = 2147483647
    v25 = v25
    var v26 bool = (3.14159 == ((0.1 - float32(0.1)) / -2.25))
    var v27 uint = 42
    var v28 float64 = -2.25
    var v29 int = v23
    if false {
        var v30 int = v29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
