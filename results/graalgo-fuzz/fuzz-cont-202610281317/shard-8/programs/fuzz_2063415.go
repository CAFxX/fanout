package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int32 }

func main() {
    var v1 int32 = -580
    var v2 uint32 = 102
    var v3 int8 = -128
    v4 := make(map[string]int)
    v5 := []string{}
    var v6 int16 = 32767
    _ = len(v4)
    if (3.14159 > 0.1) {
        switch "foo bar" {
        case "":
            for fi7 := 0; fi7 < 14; fi7++ {
                v6 = -32768
                _ = fi7
            }
        case "foo bar":
            v4[("foo bar" + ((("" + "foo bar") + "\n\t") + "a"))] = 127
        default:
            var v8 float64 = 0.0
            _ = v8
        }
    } else {
        fmt.Println(v1, v6)
        var v9 int64 = -1
        var v10 uint64 = 0
        _ = v10
        _ = v9
    }
    switch 109 {
    case 2147483647:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        fmt.Println(v1)
    }
    v6 = v6
    v3 = v3
    switch -194 {
    case 2147483647:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 128:
        if true {
            for fi11 := 0; fi11 < 1; fi11++ {
                v2 = 78
                if true {
                    if ((18446744073709551615 > 4294967295) && false) {
                        var v12 bool = (0 == 970)
                        for fi13 := 0; fi13 < 4; fi13++ {
                            var v14 float64 = 1e10
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v15 float32 = 0.1
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        if ((-2.25 >= 0.1) || (42 < 2)) {
                            v16 := GenS1{F0: 76}
                            _ = len(v5)
                            _ = v16
                        }
                        _ = v12
                    } else {
                        for fi17 := 0; fi17 < 8; fi17++ {
                            v18 := GenS2{F0: 137, F1: 2147483647}
                            if true {
                                fmt.Println(v1, fi17)
                            } else {
                                var v19 uint64 = 797
                                v6 = (v6 << ((65535) & 63))
                                _ = v19
                            }
                            _ = fi17
                            _ = v18
                        }
                        v2 = v2
                        var v20 uint64 = 833
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 20; fi21++ {
                        switch fi21 {
                        case 128:
                            v2 -= v2
                        case 4294967295:
                            var v22 uint = 324
                            _ = v22
                        default:
                            fmt.Println(fi11, v1)
                        }
                        v6 -= 1000
                        var v23 float32 = 0.0
                        _ = fi21
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 12; fi24++ {
                        switch (fi11 - fi11) {
                        case 127:
                            v2 = v2
                        }
                        _ = fi24
                    }
                } else {
                    v1 += 18
                    var v25 int32 = 1
                    var v26 float64 = 3.14159
                    _ = v25
                    _ = v26
                }
                _ = fi11
            }
            fmt.Println(v1, v3)
        }
    default:
        if true {
            var v27 uint64 = 93
            _ = v27
        } else {
            _ = len(v5)
        }
    }
    delete(v4, "a")
    for fi28 := 0; fi28 < 9; fi28++ {
        v29 := GenS1{F0: v3}
        var v30 uint16 = 934
        _ = fi28
        _ = v29
        _ = v30
    }
    v31 := GenS1{F0: (v3 | (v3 % (((^(v3))) | 1)))}
    switch "a" {
    case "":
        var v32 float32 = (1e10 * -2.25)
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
