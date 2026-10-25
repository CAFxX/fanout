package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int; F1 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -1
    var v3 int = 736
    for fi4 := 0; fi4 < 1; fi4++ {
        _ = len(v1)
        _ = fi4
    }
    var v5 string = "foo bar"
    v6 := v1[2147483647]
    v2 = (v2 * 141)
    v2 += (v2 * v2)
    var v7 string = ""
    if (!true) {
        switch 32768 {
        case 2147483647:
            v8 := GenS1{F0: int64(v3), F1: true}
            _ = v8
        default:
            v9 := []int{4294967295, (v3 - v3), (v3 % ((-534) | 1)), (v3 >> ((65535) & 63))}
            _ = v9
        }
        switch (v3 - v3) {
        case -2147483648:
            if ((false && true) || ((((1.5 > (1.5 / (-2.25 + 1.5))) && false) && true) || true)) {
                switch (v3 ^ 630) {
                case 807:
                    var v10 uint = 18446744073709551615
                    _ = v10
                default:
                    var v11 int8 = -121
                    _ = v11
                }
                var v12 uint32 = 1
                _ = v12
            }
        }
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            for fi14 := 0; fi14 < 17; fi14++ {
                v5 = v5
                switch (fi14 & 82) {
                case -284:
                    v3 += (fi13 | fi13)
                case 127:
                    if true {
                        switch fi14 {
                        case 9223372036854775807:
                            if true {
                                v3 = 2147483647
                                v15 := GenS1{F0: -9223372036854775808, F1: false}
                                _ = v15
                            } else {
                                v16 := GenS1{F0: -9223372036854775808, F1: (!false)}
                                _ = v16
                            }
                        }
                        if ((1e10 != 1.5) && false) {
                            v17 := v1[2147483648]
                            if true {
                                var v18 float32 = 1.5
                                switch "a" {
                                case "a":
                                    if ((32767 < 194) && true) {
                                        v19 := GenS1{F0: 1, F1: true}
                                        v20 := GenS1{F0: -9223372036854775808, F1: ((fi13 * (v3 - 9223372036854775807)) < (fi13 + (fi13 & -231)))}
                                        _ = v19
                                        _ = v20
                                    } else {
                                        fmt.Println(fi14, fi14)
                                        v17 = v17
                                    }
                                }
                                v21 := GenS2{F0: 255}
                                _ = v18
                                _ = v21
                            } else {
                                v22 := v1[125]
                                var v23 uint64 = 0
                                _ = v22
                                _ = v23
                            }
                            if false {
                                var v24 int = (fi13 / ((42) | 1))
                                var v25 uint64 = 0
                                _ = v24
                                _ = v25
                            }
                            v17 = (v5 + v5)
                            if false {
                                fmt.Println(fi13)
                            } else {
                                var v26 int16 = v2
                                if (!false) {
                                    v27 := GenS3{F0: fi14, F1: 65535}
                                    fmt.Println(v6)
                                    v28 := GenS4{F0: fi14, F1: 4294967295}
                                    _ = v27
                                    _ = v28
                                }
                                v29 := GenS2{F0: 65535}
                                _ = v26
                                _ = v29
                            }
                            var v30 int16 = v2
                            var v31 string = (v5 + (v17 + v6))
                            _ = v17
                            _ = v30
                            _ = v31
                        } else {
                            fmt.Println(v6)
                            var v32 float64 = 0.0
                            _ = v32
                        }
                        fmt.Println(v7)
                    }
                default:
                    v7 = v7
                }
                _ = fi14
            }
            for fi33 := 0; fi33 < 13; fi33++ {
                var v34 float64 = 1e10
                _ = fi33
                _ = v34
            }
            _ = fi13
        }
        v6 = ""
    }
    if true {
        v35 := v1[65535]
        v1[2147483648] = v7
        fmt.Println(v3)
        _ = v35
    }
    v36 := GenS3{F0: v3, F1: 848}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v5
    _ = v6
    _ = v7
}
