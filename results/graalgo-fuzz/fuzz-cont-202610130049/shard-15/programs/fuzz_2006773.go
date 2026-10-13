package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int64; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 int }

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]int)
    var v3 float32 = 0.0
    var v4 float64 = 0.1
    v5 := []string{""}
    if (-1 <= 467) {
        var v6 bool = (1 >= 260)
        _ = v6
    } else {
        for fi7 := 0; fi7 < 11; fi7++ {
            v1 -= v1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi7
        }
    }
    fmt.Println(v4, v1, v3)
    fmt.Println(v4, v3)
    v8 := GenS1{F0: 42}
    if ((0.0 / (0.0 - v3)) <= v3) {
        switch 664 {
        case -1:
            var v9 bool = (0 < 25)
            _ = v9
        case 128:
            v10 := GenS1{F0: 32767}
            _ = v10
        default:
            if false {
                for fi11 := 0; fi11 < 12; fi11++ {
                    v12 := GenS2{F0: (!(!true)), F1: 9223372036854775807, F2: true}
                    var v13 int32 = 733
                    v14 := GenS3{F0: 615}
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
            }
        }
        var v15 uint = 2
        if ((("\n\t" + "a") != "foo bar") && ((27 > 9223372036854775807) && (false || (true && false)))) {
            for fi16 := 0; fi16 < 19; fi16++ {
                if (v1 == v15) {
                    var v17 int32 = -294
                    v3 = float32(float64(v3))
                    _ = v17
                }
                var v18 int16 = 1000
                var v19 uint = (v15 ^ 553)
                _ = fi16
                _ = v18
                _ = v19
            }
            if false {
                delete(v2, 2147483648)
            } else {
                v20 := GenS4{F0: 273}
                _ = v20
            }
        }
        _ = v15
    } else {
        v21 := GenS5{F0: 1, F1: 9223372036854775807}
        var v22 int = 4294967295
        for fi23 := 0; fi23 < 17; fi23++ {
            fmt.Println(v4)
            for fi24 := 0; fi24 < 18; fi24++ {
                fmt.Println(v3, fi24, fi24)
                fmt.Println(v3, v22)
                _ = fi24
            }
            _ = fi23
        }
        _ = v21
        _ = v22
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        v4 = v4
        if ("a" > "") {
            var v25 uint8 = 190
            if len(v5) > 0 {
                v5[0] = "hello"
            }
            if (!(!true)) {
                for fi26 := 0; fi26 < 3; fi26++ {
                    v27 := GenS4{F0: fi26}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v3 = (v3 / v3)
                    _ = fi26
                    _ = v27
                }
                if true {
                    fmt.Println(v3)
                } else {
                    for fi28 := 0; fi28 < 16; fi28++ {
                        for fi29 := 0; fi29 < 17; fi29++ {
                            fmt.Println(fi29)
                            _ = fi29
                        }
                        var v30 bool = (42 <= (^(v1)))
                        _ = fi28
                        _ = v30
                    }
                }
            }
            _ = v25
        } else {
            switch 0 {
            case 2147483647:
                v4 += 0.0
            case 9223372036854775807:
                v31 := GenS2{F0: (!(2147483648 > 0)), F1: -9223372036854775808, F2: false}
                _ = v31
            }
        }
        v4 -= v4
    }
    v32 := GenS1{F0: 128}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v8
}
