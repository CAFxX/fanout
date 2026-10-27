package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint32 }
type GenS2 struct { F0 int; F1 int; F2 int32 }
type GenS3 struct { F0 int16; F1 int64 }

func main() {
    var v1 uint = 0
    var v2 uint64 = 42
    v3 := []int{-1, 65, -634, 597}
    var v4 float64 = 3.14159
    var v5 int64 = 287
    v5 = (v5 >> (((v1 >> (((v1 / ((255) | 1))) & 63))) & 63))
    v5 = (v5 - (v5 << ((4294967295) & 63)))
    v1 = v1
    var v6 int64 = (v5 ^ v5)
    v2 += (v2 | v2)
    for fi7 := 0; fi7 < 10; fi7++ {
        switch (fi7 - fi7) {
        case -740:
            v8 := GenS1{F0: 972, F1: fi7, F2: 1}
            _ = v8
        case -1:
            fmt.Println(fi7)
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        var v9 int = fi7
        _ = fi7
        _ = v9
    }
    if false {
        var v10 string = "a"
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                if false {
                    if (128 != 26) {
                        _ = len(v3)
                    } else {
                        fmt.Println(v2)
                        var v13 uint32 = 605
                        _ = v13
                    }
                }
                _ = fi12
            }
            var v14 string = v10
            _ = fi11
            _ = v14
        }
        _ = v10
    } else {
        var v15 uint16 = 42
        v6 += v5
        switch "foo bar" {
        case "":
            var v16 bool = (4294967295 >= 2147483647)
            _ = v16
        case "hello":
            var v17 int = 2147483647
            _ = v17
        case "foo bar":
            if (832 > -1) {
                switch ("foo bar" + "\n\t") {
                case "hello":
                    v18 := GenS2{F0: -864, F1: 32767, F2: -2147483648}
                    _ = v18
                case "\n\t":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v3 = append(v3, -2147483648)
                var v19 string = (((("" + "hello") + "a") + ("a" + "foo bar")) + ("a" + ""))
                _ = v19
            } else {
                v20 := GenS3{F0: 775, F1: v6}
                _ = v20
            }
        default:
            for fi21 := 0; fi21 < 1; fi21++ {
                if len(v3) > 0 {
                    v3[0] = (fi21 / ((2147483647) | 1))
                }
                _ = fi21
            }
        }
        _ = v15
    }
    var v22 int32 = 0
    for fi23 := 0; fi23 < 19; fi23++ {
        if (false || true) {
            fmt.Println(fi23, fi23)
        } else {
            v24 := GenS1{F0: 1000, F1: (fi23 / (((fi23 - (^(fi23)))) | 1)), F2: 0}
            _ = v24
        }
        _ = fi23
    }
    var v25 float32 = 0.0
    v25 -= (v25 / v25)
    switch int(v2) {
    case -1:
        v5 += (v5 % ((504) | 1))
    case -154:
        if true {
            v22 = (v22 / ((2147483647) | 1))
            for fi26 := 0; fi26 < 19; fi26++ {
                v27 := GenS2{F0: 0, F1: fi26, F2: (v22 + (v22 | 2147483647))}
                var v28 int32 = 1
                v2 = (v2 ^ v2)
                _ = fi26
                _ = v27
                _ = v28
            }
            v2 += v2
        } else {
            v3 = append(v3, -473)
            if ((718 < 498) && true) {
                for fi29 := 0; fi29 < 7; fi29++ {
                    var v30 int32 = 0
                    v3 = append(v3, fi29)
                    var v31 float32 = -2.25
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                fmt.Println(v22, v1)
            } else {
                if false {
                    for fi32 := 0; fi32 < 10; fi32++ {
                        fmt.Println(v6, v25, v4)
                        _ = fi32
                    }
                } else {
                    v6 += 42
                }
            }
            switch -150 {
            case 2147483647:
                _ = len(v3)
            default:
                v33 := GenS3{F0: 42, F1: (v6 | (v5 - 42))}
                _ = v33
            }
        }
    }
    var v34 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
