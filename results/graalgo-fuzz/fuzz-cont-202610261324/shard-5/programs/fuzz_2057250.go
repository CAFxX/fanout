package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int8; F1 int; F2 int16 }
func genFn0(p0 int16) uint16 {
    var v4 float32 = ((0.1 + (0.1 - float32((3.14159 - 0.0)))) / 1e10)
    var v5 int = 42
    _ = p0
    _ = v4
    _ = v5
    return 128
}

func main() {
    var v1 bool = true
    var v2 float32 = 1e10
    v3 := make(map[string]string)
    var v6 uint64 = 2
    v1 = (!(255 == 0))
    v2 = 1.5
    switch "" {
    case "\n\t":
        if v1 {
            for fi7 := 0; fi7 < 14; fi7++ {
                for fi8 := 0; fi8 < 16; fi8++ {
                    fmt.Println(v2)
                    v1 = (genFn0(-350) <= 65535)
                    var v9 uint16 = 612
                    switch "foo bar" {
                    case "\n\t":
                        var v10 uint64 = (v6 << ((407) & 63))
                        _ = v10
                    }
                    _ = fi8
                    _ = v9
                }
                v11 := GenS1{F0: 65535, F1: 100}
                _ = fi7
                _ = v11
            }
        } else {
            switch -10 {
            case 2147483647:
                v6 += v6
            case 128:
                if v1 {
                    if v1 {
                        if (255 > 2) {
                            var v12 uint16 = 32767
                            fmt.Println(v2, v1)
                            v13 := []int{}
                            _ = v12
                            _ = v13
                        }
                        v14 := GenS1{F0: 7, F1: -128}
                        for fi15 := 0; fi15 < 1; fi15++ {
                            fmt.Println(v2)
                            v1 = false
                            for fi16 := 0; fi16 < 19; fi16++ {
                                v17 := []int{(-(fi16)), fi15}
                                _ = fi16
                                _ = v17
                            }
                            _ = fi15
                        }
                        _ = v14
                    }
                    var v18 uint32 = 65535
                    var v19 int32 = -926
                    _ = v18
                    _ = v19
                } else {
                    if (!(false || (127 < 0))) {
                        v2 = -2.25
                        switch -289 {
                        case 629:
                            _ = len(v3)
                        case 32767:
                            v2 = -2.25
                        default:
                            v20 := []int{610}
                            _ = v20
                        }
                    } else {
                        v2 = (0.1 / (v2 / (-2.25 / ((v2 - float32(-2.25)) * (v2 + v2)))))
                        if v1 {
                            v21 := GenS2{F0: 9223372036854775807, F1: ("hello" + "a")}
                            v22 := GenS2{F0: 42, F1: "a"}
                            _ = v21
                            _ = v22
                        }
                    }
                    v23 := GenS1{F0: 1, F1: 19}
                    var v24 string = ""
                    _ = v23
                    _ = v24
                }
            default:
                var v25 float32 = v2
                _ = v25
            }
        }
    case "a":
        v6 -= (^(v6))
    default:
        v26 := GenS3{F0: 0, F1: 42, F2: -600}
        _ = v26
    }
    fmt.Println(v6, v6, v1)
    var v27 uint64 = v6
    v6 = v27
    _ = len(v3)
    for fi28 := 0; fi28 < 11; fi28++ {
        v1 = false
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v6
}
