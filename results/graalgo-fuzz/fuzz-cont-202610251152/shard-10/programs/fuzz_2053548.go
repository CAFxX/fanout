package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint64, p1 bool, p2 uint) int32 {
    v4 := []int{126, 1}
    if len(v4) > 0 {
        v4[0] = 42
    }
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 int8 = 100
            _ = fi6
            _ = v7
        }
        if (0 >= 980) {
            if len(v4) > 0 {
                v4[0] = (-(fi5))
            }
            var v8 uint8 = 128
            var v9 int = (fi5 * 9223372036854775807)
            _ = v8
            _ = v9
        } else {
            var v10 bool = (-70 >= -128)
            for fi11 := 0; fi11 < 2; fi11++ {
                if true {
                    var v12 int8 = -128
                    fmt.Println(fi11)
                    _ = v12
                } else {
                    fmt.Println(p0, fi11)
                    var v13 string = ("\n\t" + "a")
                    _ = v13
                }
                if p1 {
                    if len(v4) > 0 {
                        v4[0] = -1
                    }
                } else {
                    v14 := GenS1{F0: 32767}
                    switch fi11 {
                    case -2147483648:
                        p1 = true
                    case 4294967295:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case 2147483647:
                        var v15 int = (fi11 >> ((p2) & 63))
                        _ = v15
                    default:
                        var v16 int16 = -32768
                        _ = v16
                    }
                    switch (fi5 % ((fi11) | 1)) {
                    case 32768:
                        var v17 bool = (42 >= -128)
                        var v18 int8 = -1
                        p0 -= p0
                        fmt.Println(p0)
                        _ = v17
                        _ = v18
                    case -2147483648:
                        p1 = (true || (fi11 != 472))
                    }
                    _ = v14
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch "hello" {
                case "":
                    p0 -= p0
                case "a":
                    _ = len(v4)
                }
                _ = fi11
            }
            if false {
                if false {
                    var v19 uint8 = 92
                    _ = v19
                }
                v10 = true
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    for fi20 := 0; fi20 < 15; fi20++ {
                        if true {
                            if v10 {
                                p0 = (p0 + (p0 + p0))
                                v21 := GenS1{F0: 1000}
                                _ = len(v4)
                                _ = v21
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = (fi20 ^ int(p0))
                                }
                            }
                            if ((^(p0)) > p0) {
                                p0 += 269
                                switch "\n\t" {
                                case "\n\t":
                                    var v22 int16 = 1
                                    _ = v22
                                case "foo bar":
                                    v23 := GenS1{F0: -1}
                                    _ = v23
                                }
                            }
                        } else {
                            var v24 int64 = 9223372036854775807
                            var v25 float64 = 1.5
                            _ = v24
                            _ = v25
                        }
                        _ = fi20
                    }
                    var v26 bool = (4294967295 == 454)
                    v26 = v26
                    _ = v26
                }
            }
            _ = v10
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 80
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int32 = 680
    v3 := []uint32{(v1 & (v1 % ((4294967295) | 1))), (v1 | 2)}
    var v27 uint64 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = uint32(v27)
    v2 = v2
    v2 -= v2
    for fi28 := 0; fi28 < 20; fi28++ {
        v2 += (v2 % ((1) | 1))
        _ = fi28
    }
    _ = len(v3)
    v29 := GenS2{F0: 2147483647}
    v30 := GenS1{F0: 1}
    var v31 int = 225
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
