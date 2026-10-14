package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int32 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int16; F1 uint16; F2 uint8 }

func main() {
    var v1 float64 = ((1.5 + 0.0) / (1.5 * 1e10))
    var v2 int = -9223372036854775808
    v3 := []uint32{65535}
    var v4 uint8 = 2
    var v5 bool = false
    if v5 {
        for fi6 := 0; fi6 < 16; fi6++ {
            fmt.Println(fi6)
            v1 = v1
            if (false && (false || ((v2 / ((v2) | 1)) >= -584))) {
                for fi7 := 0; fi7 < 10; fi7++ {
                    if v5 {
                        if ((!v5) || (float32(0.1) == 0.0)) {
                            switch (fi7 * v2) {
                            case -9223372036854775808:
                                v1 = (((v1 + 0.1) * v1) / -2.25)
                            default:
                                v8 := GenS1{F0: 6, F1: "", F2: 42}
                                _ = v8
                            }
                            v9 := GenS2{F0: 0, F1: (v2 + fi6)}
                            switch fi7 {
                            case -2147483648:
                                v1 = v1
                            case 42:
                                v3 = append(v3, 2147483647)
                            default:
                                switch ("foo bar" + "a") {
                                case "hello":
                                    v10 := GenS1{F0: 4294967295, F1: "foo bar", F2: 0}
                                    _ = v10
                                case "":
                                    if len(v3) > 0 {
                                        v3[0] = uint32(v4)
                                    }
                                default:
                                    if len(v3) > 0 {
                                        v3[0] = 1
                                    }
                                }
                            }
                            _ = v9
                        }
                        v3 = append(v3, 180)
                    }
                    if (42 > 255) {
                        var v11 uint64 = 0
                        _ = v11
                    } else {
                        if ((v1 + (v1 + 0.0)) != (v1 - v1)) {
                            var v12 int = 2147483647
                            var v13 uint64 = 1
                            _ = v12
                            _ = v13
                        }
                        switch 9223372036854775807 {
                        case 2147483647:
                            if true {
                                if len(v3) > 0 {
                                    v3[0] = 4294967295
                                }
                                fmt.Println(v2)
                                v14 := GenS2{F0: 0, F1: -994}
                                fmt.Println(fi7)
                                _ = v14
                            }
                        }
                    }
                    _ = fi7
                }
                v1 = 1e10
            }
            _ = fi6
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v4 += 131
    v1 += v1
    switch int(v4) {
    case 32768:
        var v15 string = "foo bar"
        _ = v15
    case 489:
        fmt.Println(v4, v4, v2)
    case 9223372036854775807:
        var v16 string = ("" + "hello")
        _ = v16
    default:
        v5 = ((v2 ^ v2) > (v2 % ((-736) | 1)))
    }
    v3 = append(v3, 4294967295)
    v4 = (v4 ^ 26)
    if true {
        var v17 uint = 738
        v18 := GenS3{F0: 1, F1: 255, F2: v4}
        var v19 float64 = v1
        _ = v17
        _ = v18
        _ = v19
    }
    fmt.Println(v4)
    v20 := GenS3{F0: 1, F1: 32767, F2: 170}
    var v21 int16 = -655
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
