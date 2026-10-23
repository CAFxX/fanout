package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint32 }

func main() {
    v1 := []int{-9223372036854775808, 128}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 576
    var v4 string = "hello"
    v5 := make(map[string]int)
    switch 1 {
    case -9223372036854775808:
        var v6 uint = 581
        _ = v6
    case 32767:
        v3 += (v3 ^ v3)
    case -1:
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 string = ""
            _ = fi7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: 32767}
        _ = fi9
        _ = v10
    }
    v4 = (v4 + ("hello" + (v4 + (("a" + v4) + ((v4 + "\n\t") + ("\n\t" + "foo bar"))))))
    fmt.Println(v4, v3)
    v5[(v4 + "hello")] = 128
    v4 = v4
    v11 := GenS1{F0: uint16(v3)}
    if true {
        var v12 int64 = -9223372036854775808
        _ = v12
    } else {
        switch v4 {
        case "foo bar":
            if len(v1) > 0 {
                v1[0] = 127
            }
        }
    }
    var v13 int64 = -186
    switch 32767 {
    case -1:
        var v14 int = 127
        _ = v14
    }
    _ = len(v5)
    for fi15 := 0; fi15 < 9; fi15++ {
        if (false || (!true)) {
            v13 += (v13 / ((-838) | 1))
            switch "foo bar" {
            case "a":
                for fi16 := 0; fi16 < 13; fi16++ {
                    v2[2] = 0
                    v17 := GenS1{F0: 731}
                    _ = fi16
                    _ = v17
                }
            case "foo bar":
                v13 = (v13 / ((42) | 1))
            default:
                if false {
                    switch 4294967295 {
                    case 9223372036854775807:
                        v18 := GenS2{F0: false, F1: 562}
                        _ = v18
                    case 748:
                        for fi19 := 0; fi19 < 4; fi19++ {
                            v20 := v5["hello"]
                            if true {
                                v4 = "foo bar"
                            } else {
                                v21 := GenS1{F0: 32767}
                                _ = v21
                            }
                            v22 := v2[2]
                            _ = fi19
                            _ = v20
                            _ = v22
                        }
                    }
                    v13 -= (v13 - (v13 >> ((1) & 63)))
                    for fi23 := 0; fi23 < 19; fi23++ {
                        var v24 uint8 = 127
                        _ = fi23
                        _ = v24
                    }
                }
            }
            var v25 int32 = 0
            _ = v25
        }
        _ = fi15
    }
    v26 := GenS3{F0: 857, F1: (v3 | 254), F2: 642}
    var v27 uint = 352
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
