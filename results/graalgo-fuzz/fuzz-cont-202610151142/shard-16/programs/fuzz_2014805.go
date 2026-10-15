package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint64 = 939
    var v2 string = ("a" + ("hello" + "hello"))
    var v3 uint = 0
    var v4 int16 = 611
    var v5 int64 = -9223372036854775808
    var v6 float64 = 1.5
    if (0 >= 177) {
        fmt.Println(v3)
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int = fi7
            v8 = -797
            _ = fi7
            _ = v8
        }
    } else {
        var v9 bool = false
        v4 += v4
        _ = v9
    }
    switch 584 {
    case 32767:
        v6 += (float64((0.0 - 1e10)) * float64(0.1))
    case 632:
        var v10 uint16 = 245
        _ = v10
    case -181:
        v6 = v6
    default:
        fmt.Println(v5, v2)
    }
    v4 -= (v4 >> ((v3) & 63))
    var v11 uint16 = 42
    var v12 int8 = 0
    var v13 float64 = v6
    var v14 int32 = 1
    if (true && ((v14 / ((v14) | 1)) == (v14 % ((-2147483648) | 1)))) {
        switch v2 {
        case "":
            var v15 uint64 = (v1 % (((v1 << ((v3) & 63))) | 1))
            _ = v15
        default:
            v16 := []int{}
            _ = v16
        }
        var v17 string = v2
        v18 := GenS1{F0: v1, F1: -2147483648}
        _ = v17
        _ = v18
    } else {
        v19 := GenS2{F0: v4, F1: 127}
        _ = v19
    }
    switch int(v11) {
    case 42:
        var v20 bool = (v5 > (v5 * -933))
        _ = v20
    default:
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := []int{fi21, int(v12), (fi21 * int(v4))}
            for fi23 := 0; fi23 < 12; fi23++ {
                v5 += v5
                _ = len(v22)
                v24 := GenS3{F0: v4, F1: (v5 + v5)}
                _ = fi23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
    }
    v4 += (v4 << ((65535) & 63))
    for fi25 := 0; fi25 < 12; fi25++ {
        for fi26 := 0; fi26 < 3; fi26++ {
            if ((!false) || (v11 <= (v11 << ((v3) & 63)))) {
                var v27 uint8 = 1
                var v28 int64 = (v5 ^ v5)
                v3 = (^(v3))
                _ = v27
                _ = v28
            }
            v29 := GenS3{F0: 32767, F1: (v5 << ((885) & 63))}
            v30 := GenS3{F0: v4, F1: v5}
            _ = fi26
            _ = v29
            _ = v30
        }
        _ = fi25
    }
    var v31 int = 32768
    for fi32 := 0; fi32 < 19; fi32++ {
        switch (v31 << (((v3 << (((^(v3))) & 63))) & 63)) {
        case 127:
            if ((v14 / (((-(v14))) | 1)) >= v14) {
                switch fi32 {
                case 1000:
                    for fi33 := 0; fi33 < 13; fi33++ {
                        v34 := []int{}
                        var v35 uint16 = (v11 >> ((0) & 63))
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                case 2147483647:
                    switch v2 {
                    case "hello":
                        fmt.Println(fi32)
                    case "a":
                        switch 0 {
                        case 42:
                            switch v2 {
                            case "hello":
                                for fi36 := 0; fi36 < 1; fi36++ {
                                    v14 = (^(v14))
                                    switch 4294967295 {
                                    case 2147483647:
                                        fmt.Println(v3, v14)
                                    }
                                    _ = fi36
                                }
                            }
                        case -741:
                            v37 := GenS4{F0: false}
                            _ = v37
                        default:
                            for fi38 := 0; fi38 < 9; fi38++ {
                                var v39 bool = (v11 >= 547)
                                _ = fi38
                                _ = v39
                            }
                        }
                    case "foo bar":
                        var v40 int64 = v5
                        _ = v40
                    }
                default:
                    fmt.Println(v4, v11)
                }
                for fi41 := 0; fi41 < 14; fi41++ {
                    v11 = 1
                    v42 := []int{fi41, (-(fi32))}
                    v6 -= (v13 / (v6 / (v13 * (v13 * (((v6 + ((v6 * (0.1 - (float64(float32(v6)) * v13))) + float64((0.0 + -2.25)))) - (0.0 + v13)) + -2.25)))))
                    _ = fi41
                    _ = v42
                }
            }
        case 659:
            v12 = 109
        case -1:
            fmt.Println(v12)
        }
        _ = fi32
    }
    v2 = v2
    v43 := GenS2{F0: (v4 & v4), F1: v31}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
