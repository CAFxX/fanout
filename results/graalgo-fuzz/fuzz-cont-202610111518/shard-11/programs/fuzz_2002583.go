package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := []uint32{42, 42, 2, 1}
    v2 := make(map[int]uint32)
    var v3 bool = true
    var v4 int8 = -90
    for fi5 := 0; fi5 < 16; fi5++ {
        if v3 {
            var v6 uint32 = 2147483647
            if v3 {
                v2[(fi5 % ((fi5) | 1))] = (^(v6))
                v6 += 2
            }
            _ = v6
        } else {
            var v7 int64 = 42
            v7 += v7
            _ = v7
        }
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint64 = 42
            _ = fi8
            _ = v9
        }
        _ = fi5
    }
    v1 = append(v1, 685)
    for fi10 := 0; fi10 < 16; fi10++ {
        v4 = (v4 / ((v4) | 1))
        for fi11 := 0; fi11 < 4; fi11++ {
            v4 += v4
            _ = fi11
        }
        switch fi10 {
        case 1:
            if ((1 < 2) && v3) {
                for fi12 := 0; fi12 < 19; fi12++ {
                    _ = len(v2)
                    v3 = (((1.5 / float32((float64(1e10) * (-2.25 - float64(1e10))))) / float32(-2.25)) == 0.0)
                    _ = fi12
                }
                v4 += (-(v4))
                var v13 uint16 = 255
                _ = v13
            }
        default:
            v14 := GenS1{F0: 0, F1: 2, F2: 734}
            _ = v14
        }
        _ = fi10
    }
    if (v3 || v3) {
        if v3 {
            var v15 int8 = v4
            if (!v3) {
                switch "foo bar" {
                case "hello":
                    v3 = v3
                case "":
                    for fi16 := 0; fi16 < 8; fi16++ {
                        for fi17 := 0; fi17 < 13; fi17++ {
                            v18 := v2[(fi17 / ((fi17) | 1))]
                            _ = fi17
                            _ = v18
                        }
                        switch fi16 {
                        case 2147483647:
                            if v3 {
                                v3 = (65535 == 2147483647)
                                fmt.Println(fi16, v15)
                                var v19 uint8 = 42
                                _ = v19
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = 4294967295
                                }
                                v20 := GenS2{F0: 128}
                                v3 = true
                                _ = v20
                            }
                        default:
                            fmt.Println(v4, v3, v15)
                        }
                        var v21 uint16 = 923
                        _ = fi16
                        _ = v21
                    }
                case "\n\t":
                    var v22 uint8 = 0
                    _ = v22
                }
                var v23 uint64 = 0
                v24 := GenS3{F0: 2}
                _ = v23
                _ = v24
            } else {
                v3 = true
            }
            _ = v15
        }
    }
    var v25 uint16 = 2
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := GenS4{F0: (v25 / (((v25 / ((1) | 1))) | 1))}
        v28 := GenS1{F0: 65535, F1: 2, F2: 833}
        _ = fi26
        _ = v27
        _ = v28
    }
    if v3 {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        _ = len(v1)
        v3 = (0.1 <= (1e10 / float32((1e10 - 3.14159))))
    }
    v3 = false
    v25 = 32767
    var v29 int32 = 2147483647
    var v30 int64 = 1
    v30 = v30
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
