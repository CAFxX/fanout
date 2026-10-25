package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint32; F2 string }
type GenS3 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int32 = 2147483647
    v3 := make(map[int]uint32)
    fmt.Println(v2)
    fmt.Println(v2, v2)
    var v4 float64 = 0.1
    for fi5 := 0; fi5 < 7; fi5++ {
        if (!true) {
            var v6 int = fi5
            if (851 <= 2147483647) {
                for fi7 := 0; fi7 < 13; fi7++ {
                    delete(v3, v6)
                    if (!false) {
                        v2 = v2
                    } else {
                        if (18446744073709551615 < 4294967295) {
                            v8 := GenS1{F0: false}
                            var v9 uint8 = 247
                            _ = v8
                            _ = v9
                        } else {
                            fmt.Println(v6, v2)
                            switch (v6 * (fi5 << ((65535) & 63))) {
                            case 2147483647:
                                v10 := []int{v6, (fi7 & 32768), 0}
                                _ = v10
                            default:
                                v6 = (v6 * (v6 ^ (-(v6))))
                            }
                            v2 -= (v2 | (v2 << ((255) & 63)))
                            if (true || true) {
                                switch ((("\n\t" + ("" + "\n\t")) + "") + "\n\t") {
                                case "foo bar":
                                    _ = len(v3)
                                case "":
                                    if true {
                                        v2 = (v2 << ((18446744073709551615) & 63))
                                        v1 -= (-(v1))
                                    } else {
                                        fmt.Println(fi7, v4, v1)
                                    }
                                }
                                fmt.Println(v2, v4)
                            }
                            if false {
                                var v11 int8 = -21
                                _ = v11
                            }
                            var v12 int16 = -32768
                            _ = len(v3)
                            var v13 int8 = 1
                            v14 := GenS2{F0: v13, F1: 91, F2: "foo bar"}
                            v15 := GenS1{F0: (!false)}
                            v4 -= (float64(0.0) * v4)
                            _ = v12
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                        var v16 int64 = 1
                        if (!((fi5 <= (v6 * v6)) && false)) {
                            v4 += -2.25
                        }
                        _ = v16
                    }
                    _ = fi7
                }
                if false {
                    v17 := GenS3{F0: 80, F1: 42}
                    if (32767 > 877) {
                        v6 = v6
                        delete(v3, fi5)
                        v4 -= v4
                    }
                    _ = v17
                }
            }
            var v18 int64 = v1
            _ = v18
            _ = v6
        } else {
            _ = len(v3)
            if (65535 == 255) {
                v19 := GenS2{F0: -1, F1: 2, F2: "a"}
                _ = v19
            }
        }
        var v20 uint32 = 2147483648
        var v21 float32 = -2.25
        _ = fi5
        _ = v20
        _ = v21
    }
    if true {
        v4 = (v4 - (float64((-2.25 / 1.5)) / v4))
        _ = len(v3)
        v2 -= (v2 + (v2 ^ -2147483648))
    }
    v1 += v1
    delete(v3, 42)
    var v22 uint8 = 127
    var v23 string = "\n\t"
    for fi24 := 0; fi24 < 2; fi24++ {
        v22 = 110
        _ = fi24
    }
    var v25 int16 = -125
    v26 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
