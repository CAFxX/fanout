package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    var v4 bool = (1e10 < 0.1)
    var v5 int64 = 9223372036854775807
    v4 = (3.14159 < 0.1)
    var v6 int = 32767
    var v7 int = 32768
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(v5)
        var v9 int8 = 127
        _ = fi8
        _ = v9
    }
    if (false && v4) {
        if v4 {
            for fi10 := 0; fi10 < 2; fi10++ {
                if false {
                    fmt.Println(v1, v5)
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                } else {
                    fmt.Println(fi10, v4, v1)
                }
                _ = fi10
            }
        } else {
            switch v7 {
            case 2147483647:
                for fi11 := 0; fi11 < 17; fi11++ {
                    v5 = v5
                    var v12 uint8 = 142
                    v13 := GenS1{F0: 255, F1: 2, F2: (v5 ^ (v5 - (v5 << ((0) & 63))))}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case -385:
                v14 := GenS1{F0: 655, F1: 4294967295, F2: (-(v5))}
                _ = v14
            case 116:
                v3["hello"] = 2
            default:
                switch "a" {
                case "foo bar":
                    if v4 {
                        var v15 uint64 = 450
                        v1 = 2
                        if ((9223372036854775807 > v5) && v4) {
                            v4 = v4
                            v4 = (((!(v4 && (v1 == 678))) && (!v4)) || (65535 == 730))
                            switch "\n\t" {
                            case "\n\t":
                                v7 = (v6 >> ((18446744073709551615) & 63))
                            }
                        } else {
                            for fi16 := 0; fi16 < 13; fi16++ {
                                _ = len(v2)
                                var v17 int8 = 100
                                if (42 <= 1) {
                                    var v18 int32 = 0
                                    _ = v18
                                }
                                _ = fi16
                                _ = v17
                            }
                        }
                        _ = v15
                    } else {
                        var v19 float64 = -2.25
                        _ = v19
                    }
                default:
                    for fi20 := 0; fi20 < 12; fi20++ {
                        if v4 {
                            v6 += -881
                            if (255 >= 151) {
                                v3[""] = 315
                                fmt.Println(v6, fi20, v1)
                            } else {
                                fmt.Println(fi20, v1, v6)
                            }
                            v21 := GenS1{F0: 4294967295, F1: 4294967295, F2: 42}
                            _ = v21
                        }
                        _ = fi20
                    }
                }
            }
            v5 = (v5 | (-(v5)))
            v22 := GenS2{F0: 2, F1: (^(v1))}
            _ = v22
        }
    }
    var v23 bool = (v5 >= 9223372036854775807)
    if v23 {
        v24 := GenS1{F0: 913, F1: 605, F2: 42}
        var v25 uint8 = 34
        var v26 int = v6
        _ = v24
        _ = v25
        _ = v26
    } else {
        if (v23 && v4) {
            delete(v3, "a")
            v4 = (!("hello" <= "hello"))
            var v27 uint8 = 95
            _ = v27
        } else {
            if len(v2) > 0 {
                v2[0] = 959
            }
            v28 := GenS1{F0: 379, F1: 553, F2: (v5 ^ (v5 | v5))}
            var v29 int32 = 2147483647
            _ = v28
            _ = v29
        }
        v7 += v7
    }
    if v4 {
        if false {
            for fi30 := 0; fi30 < 6; fi30++ {
                _ = len(v2)
                _ = fi30
            }
        } else {
            fmt.Println(v5, v5, v5)
            v1 -= (v1 - v1)
        }
        var v31 int = v7
        if (!(721 > 2)) {
            if v23 {
                v6 += (v6 + (v6 | (v6 % (((^(v7))) | 1))))
            }
            v32 := GenS2{F0: 48, F1: v1}
            var v33 string = ("a" + "a")
            _ = v32
            _ = v33
        }
        _ = v31
    }
    var v34 float32 = 0.0
    for fi35 := 0; fi35 < 7; fi35++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
