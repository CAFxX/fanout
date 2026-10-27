package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 float64 = 0.1
    var v2 int16 = 525
    v3 := make(map[string]int)
    var v4 uint8 = 56
    var v5 bool = true
    var v6 int = 9223372036854775807
    fmt.Println(v1)
    v7 := GenS1{F0: 640, F1: 42, F2: -6}
    fmt.Println(v4, v6, v6)
    v6 -= (v6 + (v6 % ((v6) | 1)))
    v8 := GenS1{F0: 18446744073709551615, F1: 2, F2: 100}
    switch -9223372036854775808 {
    case 4294967295:
        v6 = v6
    }
    switch (v6 % ((v6) | 1)) {
    case 42:
        v5 = (!v5)
    case 702:
        var v9 int = -359
        _ = v9
    case 1:
        delete(v3, (((("\n\t" + (("hello" + (("\n\t" + (("hello" + "foo bar") + "foo bar")) + "")) + "")) + "hello") + "a") + "foo bar"))
    }
    delete(v3, "a")
    if v5 {
        switch "" {
        case "a":
            v10 := []int{(v6 << ((255) & 63)), (v6 / (((v6 * v6)) | 1))}
            _ = v10
        case "foo bar":
            switch -9223372036854775808 {
            case 769:
                for fi11 := 0; fi11 < 7; fi11++ {
                    v4 = 208
                    _ = fi11
                }
            case -1:
                if (!v5) {
                    var v12 float64 = (1e10 / v1)
                    _ = v12
                } else {
                    var v13 uint = 255
                    if v5 {
                        var v14 int64 = int64(v13)
                        for fi15 := 0; fi15 < 14; fi15++ {
                            v1 -= 1.5
                            _ = fi15
                        }
                        _ = v14
                    } else {
                        fmt.Println(v2, v1, v6)
                        v5 = true
                        v4 = v4
                    }
                    if (!((v13 ^ v13) <= (v13 % ((4294967295) | 1)))) {
                        v6 = -1
                        fmt.Println(v2, v4)
                    }
                    _ = v13
                }
            case -980:
                if v5 {
                    var v16 uint = 448
                    _ = v16
                } else {
                    var v17 uint64 = 77
                    v18 := v3[(("a" + "hello") + "hello")]
                    v18 += (v6 | v18)
                    _ = v17
                    _ = v18
                }
            default:
                v5 = ((v4 | v4) != v4)
            }
        default:
            v1 = v1
        }
        switch v6 {
        case -1:
            var v19 int = v6
            _ = v19
        case -532:
            _ = len(v3)
        default:
            switch ("foo bar" + "") {
            case "":
                var v20 int8 = 71
                _ = v20
            default:
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 uint64 = 0
                    for fi23 := 0; fi23 < 11; fi23++ {
                        if (!false) {
                            fmt.Println(v2, v5)
                            v24 := []int{4294967295, (fi23 & v6)}
                            if ((!v5) && (-178 > 1)) {
                                v4 += (v4 * (^(v4)))
                                v2 = v2
                                _ = len(v24)
                                v25 := GenS1{F0: 0, F1: 199, F2: -1}
                                v26 := GenS2{F0: true}
                                v24 = append(v24, (v6 / ((fi23) | 1)))
                                _ = v25
                                _ = v26
                            }
                            _ = v24
                        }
                        _ = fi23
                    }
                    _ = fi21
                    _ = v22
                }
            }
        }
        v1 -= (v1 / (v1 + v1))
    } else {
        v1 -= v1
        var v27 float32 = float32(-2.25)
        _ = v27
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 uint64 = 2
        v2 += v2
        _ = fi28
        _ = v29
    }
    v30 := GenS3{F0: (^(v4)), F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
