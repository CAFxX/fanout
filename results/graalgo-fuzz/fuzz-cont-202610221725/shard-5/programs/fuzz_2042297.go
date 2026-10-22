package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int16 = 1000
    var v2 int64 = -942
    var v3 int64 = -9223372036854775808
    v4 := []int{538, -1}
    var v5 bool = (42 != 835)
    var v6 uint64 = 1
    if v5 {
        v1 += (v1 % (((v1 & (v1 - (v1 >> ((0) & 63))))) | 1))
        var v7 int64 = v3
        switch 154 {
        case 1:
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
        case 42:
            for fi8 := 0; fi8 < 10; fi8++ {
                if false {
                    v6 -= 346
                    for fi9 := 0; fi9 < 1; fi9++ {
                        if (458 != 2147483647) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v4 = append(v4, fi8)
                        } else {
                            var v10 string = "a"
                            _ = v10
                        }
                        v11 := GenS1{F0: fi8, F1: int64(fi8)}
                        _ = fi9
                        _ = v11
                    }
                } else {
                    v12 := GenS1{F0: (^(fi8)), F1: 1}
                    var v13 int16 = 281
                    v14 := GenS1{F0: fi8, F1: (v7 | v3)}
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi8
            }
        case 540:
            fmt.Println(v1)
        default:
            var v15 int8 = 127
            _ = v15
        }
        _ = v7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, -1)
    fmt.Println(v5, v3, v3)
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := GenS2{F0: -1}
        _ = fi16
        _ = v17
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v18 string = "foo bar"
    fmt.Println(v5)
    for fi19 := 0; fi19 < 16; fi19++ {
        if true {
            var v20 bool = v5
            _ = v20
        } else {
            for fi21 := 0; fi21 < 12; fi21++ {
                fmt.Println(fi19, v3, v3)
                fmt.Println(fi21, v18, v6)
                _ = fi21
            }
            var v22 bool = false
            var v23 float64 = 1.5
            _ = v22
            _ = v23
        }
        if (!true) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            for fi24 := 0; fi24 < 13; fi24++ {
                fmt.Println(v3, fi24, v3)
                switch (v18 + "") {
                case "a":
                    switch (-(fi24)) {
                    case 32768:
                        if len(v4) > 0 {
                            v4[0] = -186
                        }
                    case 0:
                        fmt.Println(v3)
                    }
                case "":
                    v25 := GenS1{F0: (fi19 << ((4294967295) & 63)), F1: (^(v3))}
                    _ = v25
                }
                _ = fi24
            }
        }
        _ = fi19
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 bool = false
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 int64 = (v3 / ((v2) | 1))
            var v30 float64 = (0.1 / 1e10)
            v4 = append(v4, -9223372036854775808)
            switch fi28 {
            case 1:
                var v31 uint16 = 1
                _ = v31
            case 634:
                if v27 {
                    _ = len(v4)
                    var v32 bool = v27
                    if v27 {
                        var v33 uint8 = 106
                        _ = v33
                    }
                    _ = v32
                } else {
                    var v34 float32 = -2.25
                    _ = v34
                }
            }
            if true {
                fmt.Println(fi28, fi26, v5)
            }
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi26
        _ = v27
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
