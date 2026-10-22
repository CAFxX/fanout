package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 string = "hello"
    var v2 bool = false
    v3 := []string{(v1 + v1), v1}
    v4 := make(map[string]int)
    var v5 uint64 = 4294967295
    var v6 uint8 = 59
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    v7 := GenS1{F0: 337}
    fmt.Println(v6, v6, v1)
    for fi8 := 0; fi8 < 17; fi8++ {
        if (909 > 42) {
            v2 = false
            for fi9 := 0; fi9 < 20; fi9++ {
                v5 = (^(v5))
                _ = len(v4)
                _ = fi9
            }
        }
        switch (fi8 & int(v6)) {
        case 32767:
            v2 = (v1 < v1)
        case -9223372036854775808:
            for fi10 := 0; fi10 < 15; fi10++ {
                v4[(v1 + v1)] = (fi8 + fi10)
                _ = fi10
            }
        }
        _ = fi8
    }
    fmt.Println(v1)
    v4["hello"] = 2147483647
    var v11 uint16 = 2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        switch fi12 {
        case 127:
            _ = len(v4)
        case 1:
            if false {
                for fi13 := 0; fi13 < 13; fi13++ {
                    if (v2 || (false && false)) {
                        if false {
                            var v14 uint8 = (v6 - 127)
                            _ = v14
                        }
                    }
                    var v15 uint32 = 42
                    _ = fi13
                    _ = v15
                }
                for fi16 := 0; fi16 < 5; fi16++ {
                    fmt.Println(v1, v11)
                    _ = fi16
                }
            } else {
                if (163 < 311) {
                    var v17 int8 = -57
                    v5 -= 71
                    _ = v17
                } else {
                    v2 = ((v1 > v1) || (float64(1.5) >= 0.1))
                    switch fi12 {
                    case 258:
                        if v2 {
                            var v18 uint16 = v11
                            for fi19 := 0; fi19 < 12; fi19++ {
                                v20 := v4[(v1 + v1)]
                                var v21 uint = 42
                                v21 -= (v21 << (((v21 ^ (v21 << (((v21 & v21)) & 63)))) & 63))
                                v22 := GenS2{F0: 68}
                                v23 := GenS2{F0: 0}
                                _ = len(v3)
                                _ = fi19
                                _ = v20
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            var v24 int = fi12
                            _ = v18
                            _ = v24
                        }
                    default:
                        fmt.Println(v1, fi12, v11)
                    }
                    v2 = (0 != 4294967295)
                }
                switch -1 {
                case -9223372036854775808:
                    _ = len(v4)
                default:
                    _ = len(v3)
                }
            }
        case 32767:
            var v25 int = (fi12 / (((^(fi12))) | 1))
            _ = v25
        }
        v2 = v2
        if (-802 > 1) {
            v6 = (v6 - 69)
            fmt.Println(v6)
        } else {
            v4[((v1 + v1) + "foo bar")] = 32768
            fmt.Println(fi12, v1, v1)
            v4["hello"] = 127
        }
        _ = fi12
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        if (((float64((0.0 + 1e10)) < 1e10) && v2) || true) {
            v11 = uint16(v5)
            if v2 {
                var v27 float32 = -2.25
                fmt.Println(fi26, v1, fi26)
                v28 := GenS3{F0: v6}
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 17; fi29++ {
                v11 = 32767
                if len(v3) > 0 {
                    v3[0] = (v1 + v1)
                }
                _ = fi29
            }
        } else {
            v1 = (v1 + "")
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
