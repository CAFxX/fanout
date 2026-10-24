package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 int64 }

func main() {
    v1 := []uint32{}
    var v2 uint8 = 255
    var v3 uint = 255
    _ = len(v1)
    v1 = append(v1, 0)
    _ = len(v1)
    switch 0 {
    case 2147483647:
        v4 := GenS1{F0: -128, F1: (^(v2)), F2: (^(v2))}
        _ = v4
    case 42:
        if (42 > 1000) {
            switch 32767 {
            case -994:
                for fi5 := 0; fi5 < 10; fi5++ {
                    var v6 int16 = 42
                    var v7 uint32 = 2
                    _ = fi5
                    _ = v6
                    _ = v7
                }
            case 625:
                v3 = v3
            }
            fmt.Println(v3, v3)
            v2 = v2
        }
    default:
        if (749 != 2147483647) {
            var v8 int8 = -128
            v2 = 127
            if true {
                v8 = v8
            } else {
                if false {
                    v9 := GenS2{F0: 127}
                    v10 := GenS3{F0: (^(v2)), F1: 4294967295}
                    var v11 float64 = 1e10
                    _ = v10
                    _ = v11
                    _ = v9
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            _ = v8
        }
    }
    var v12 uint16 = 255
    v12 += 1
    if true {
        var v13 float32 = float32(0.0)
        v14 := GenS4{F0: 42}
        v1 = append(v1, 4294967295)
        _ = v13
        _ = v14
    } else {
        switch "\n\t" {
        case "":
            switch 678 {
            case 656:
                for fi15 := 0; fi15 < 4; fi15++ {
                    v2 += (^(v2))
                    fmt.Println(fi15)
                    for fi16 := 0; fi16 < 2; fi16++ {
                        switch (fi16 ^ fi16) {
                        case 32768:
                            fmt.Println(fi16, v2)
                        case -1:
                            _ = len(v1)
                        case 1:
                            var v17 int16 = 570
                            v18 := GenS2{F0: fi16}
                            v17 = v17
                            _ = v17
                            _ = v18
                        default:
                            v3 += (v3 - (v3 ^ v3))
                        }
                        _ = fi16
                    }
                    _ = fi15
                }
            default:
                fmt.Println(v2, v2)
            }
        case "a":
            for fi19 := 0; fi19 < 11; fi19++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi20 := 0; fi20 < 2; fi20++ {
                    switch 32768 {
                    case 127:
                        if len(v1) > 0 {
                            v1[0] = 970
                        }
                    }
                    var v21 int = fi20
                    v12 = v12
                    _ = fi20
                    _ = v21
                }
                _ = fi19
            }
        case "\n\t":
            v3 += 255
        default:
            switch -430 {
            case -9223372036854775808:
                v22 := GenS5{F0: 32767, F1: -9223372036854775808}
                _ = v22
            case -894:
                var v23 uint32 = 2147483647
                _ = v23
            case 966:
                if false {
                    fmt.Println(v12, v12, v12)
                    v24 := GenS2{F0: 1}
                    _ = len(v1)
                    _ = v24
                } else {
                    fmt.Println(v12, v2, v3)
                }
            }
        }
        fmt.Println(v2, v3, v12)
    }
    v3 = 18446744073709551615
    var v25 uint64 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v3
}
