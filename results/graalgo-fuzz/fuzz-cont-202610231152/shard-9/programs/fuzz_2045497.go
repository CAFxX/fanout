package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint }
func genFn0(p0 uint, p1 uint8, p2 uint32) int16 {
    fmt.Println(p1, p1, p1)
    p1 += (p1 << (((p0 << ((p0) & 63))) & 63))
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 0
    var v3 int8 = 57
    var v4 bool = false
    if (-925 < 32767) {
        v3 = v3
    } else {
        for fi5 := 0; fi5 < 4; fi5++ {
            if ((!(1 == 65535)) && true) {
                var v6 uint16 = 0
                switch "\n\t" {
                case "hello":
                    switch (fi5 / ((-9223372036854775808) | 1)) {
                    case 780:
                        var v7 float32 = (3.14159 / ((-2.25 / 3.14159) + 0.0))
                        _ = v7
                    case 1:
                        v6 += (^(v6))
                    default:
                        for fi8 := 0; fi8 < 10; fi8++ {
                            v3 += (^(v3))
                            switch fi5 {
                            case 543:
                                var v9 string = "\n\t"
                                _ = v9
                            default:
                                v4 = v4
                            }
                            v3 -= v3
                            if true {
                                var v10 bool = true
                                _ = v10
                            }
                            _ = fi8
                        }
                    }
                default:
                    var v11 int = fi5
                    _ = v11
                }
                switch -971 {
                case 9223372036854775807:
                    v6 -= v6
                case -277:
                    v4 = v4
                }
                _ = v6
            }
            _ = fi5
        }
        if v4 {
            switch 128 {
            case 0:
                for fi12 := 0; fi12 < 4; fi12++ {
                    for fi13 := 0; fi13 < 19; fi13++ {
                        var v14 int = 1
                        var v15 string = ("" + "")
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    v4 = true
                    _ = fi12
                }
            case 915:
                v16 := GenS1{F0: 116}
                _ = v16
            }
        } else {
            if (v3 != (v3 ^ v3)) {
                switch "" {
                case "foo bar":
                    _ = len(v1)
                case "":
                    for fi17 := 0; fi17 < 5; fi17++ {
                        for fi18 := 0; fi18 < 3; fi18++ {
                            v19 := GenS2{F0: 0, F1: 1, F2: 42}
                            _ = fi18
                            _ = v19
                        }
                        var v20 int32 = 690
                        _ = fi17
                        _ = v20
                    }
                }
            }
        }
        var v21 float32 = 1e10
        _ = v21
    }
    v3 = v3
    v4 = false
    v3 += v3
    fmt.Println(v3, v3, v3)
    fmt.Println(v4, v4, v4)
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 int32 = 1
        _ = fi22
        _ = v23
    }
    var v24 uint16 = 658
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
