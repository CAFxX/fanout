package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint32 }

func main() {
    var v1 uint16 = 1
    var v2 int16 = 1
    v3 := make(map[uint32]int)
    var v4 string = "foo bar"
    if (186 < 2) {
        v5 := []int{128, 32767}
        _ = v5
    } else {
        if false {
            if (v4 > ((v4 + v4) + "\n\t")) {
                v6 := GenS1{F0: (v1 << ((0) & 63)), F1: (v2 | v2), F2: 341}
                v7 := GenS1{F0: v1, F1: v2, F2: (^(v1))}
                _ = v6
                _ = v7
            }
            switch v4 {
            case "hello":
                if ((true && (-1 == -108)) || false) {
                    v8 := GenS1{F0: v1, F1: (^(v2)), F2: (v1 ^ (v1 | 32767))}
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 13; fi9++ {
                        switch fi9 {
                        case -84:
                            var v10 uint = 0
                            _ = v10
                        }
                        if (1 <= -71) {
                            for fi11 := 0; fi11 < 14; fi11++ {
                                _ = len(v3)
                                _ = fi11
                            }
                            v2 += v2
                        }
                        _ = fi9
                    }
                }
            default:
                var v12 string = ((v4 + v4) + ((("" + "foo bar") + v4) + v4))
                _ = v12
            }
        }
        var v13 int16 = (v2 | (v2 ^ (v2 * v2)))
        _ = v13
    }
    v1 = (^(v1))
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            if false {
                switch "\n\t" {
                case "foo bar":
                    v1 = v1
                case "":
                    v16 := GenS2{F0: -1}
                    _ = v16
                case "a":
                    var v17 int32 = -1
                    _ = v17
                default:
                    v18 := GenS3{F0: true, F1: 382}
                    _ = v18
                }
            } else {
                var v19 uint16 = 65535
                v20 := GenS1{F0: 177, F1: -1, F2: v1}
                if true {
                    v4 = v4
                    var v21 int16 = 32767
                    _ = v21
                } else {
                    fmt.Println(fi14, v1)
                    v22 := GenS2{F0: 2147483647}
                    v23 := GenS3{F0: false, F1: 2}
                    _ = v22
                    _ = v23
                }
                _ = v19
                _ = v20
            }
            _ = fi15
        }
        for fi24 := 0; fi24 < 7; fi24++ {
            v4 = (v4 + "\n\t")
            v4 = v4
            var v25 int16 = -857
            delete(v3, 734)
            _ = fi24
            _ = v25
        }
        v1 = v1
        _ = fi14
    }
    v26 := GenS1{F0: v1, F1: 683, F2: v1}
    switch v4 {
    case "a":
        var v27 float32 = -2.25
        _ = v27
    default:
        v1 = v1
    }
    v2 = -1
    v28 := GenS1{F0: 0, F1: 32767, F2: 65535}
    if true {
        fmt.Println(v2, v1, v1)
        fmt.Println(v1, v2)
        v29 := GenS2{F0: 42}
        _ = v29
    }
    fmt.Println(v4, v2)
    var v30 int16 = v2
    var v31 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
