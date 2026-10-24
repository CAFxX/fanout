package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int = 1
    var v2 float64 = float64(0.1)
    var v3 uint64 = 18446744073709551615
    v4 := []string{("\n\t" + "\n\t"), ("hello" + "hello"), "foo bar"}
    var v5 bool = true
    fmt.Println(v2, v2)
    var v6 int8 = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v7 := GenS1{F0: 127}
    var v8 uint64 = (v3 - (v3 << ((513) & 63)))
    v3 -= (v3 + 2)
    v2 += v2
    var v9 uint32 = 763
    v10 := GenS2{F0: v6}
    _ = len(v4)
    _ = len(v4)
    if (!true) {
        switch 0 {
        case 4294967295:
            if v5 {
                for fi11 := 0; fi11 < 19; fi11++ {
                    v12 := GenS2{F0: (^(v6))}
                    if (!false) {
                        var v13 int16 = -32768
                        if v5 {
                            var v14 uint16 = 2
                            if false {
                                v8 = 188
                            } else {
                                for fi15 := 0; fi15 < 3; fi15++ {
                                    v16 := GenS1{F0: 1}
                                    _ = fi15
                                    _ = v16
                                }
                                v9 = v9
                                for fi17 := 0; fi17 < 18; fi17++ {
                                    fmt.Println(v3)
                                    v4 = append(v4, "hello")
                                    fmt.Println(v5, v3)
                                    _ = fi17
                                }
                            }
                            v18 := GenS2{F0: v6}
                            _ = v14
                            _ = v18
                        }
                        _ = v13
                    } else {
                        switch 39 {
                        case 128:
                            var v19 int = fi11
                            _ = v19
                        case 32768:
                            v5 = (255 != 2)
                        default:
                            v20 := GenS1{F0: uint8(v9)}
                            _ = v20
                        }
                    }
                    _ = fi11
                    _ = v12
                }
                for fi21 := 0; fi21 < 8; fi21++ {
                    v2 = ((v2 - 0.1) - float64(float32(0.0)))
                    v2 += v2
                    for fi22 := 0; fi22 < 17; fi22++ {
                        switch fi22 {
                        case 1:
                            v23 := GenS1{F0: 128}
                            _ = v23
                        case 9223372036854775807:
                            v9 += 214
                        case 127:
                            if len(v4) > 0 {
                                v4[0] = ("" + "a")
                            }
                        }
                        _ = fi22
                    }
                    _ = fi21
                }
            } else {
                switch v1 {
                case -1:
                    var v24 int64 = 1
                    _ = v24
                case 127:
                    v2 = v2
                case 42:
                    v1 -= (v1 + v1)
                default:
                    if v5 {
                        v6 += v6
                    }
                }
                var v25 float32 = 0.0
                _ = v25
            }
        case 1:
            v4 = append(v4, "a")
        }
        for fi26 := 0; fi26 < 8; fi26++ {
            v1 = 0
            _ = fi26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
