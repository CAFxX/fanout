package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 bool; F1 int64; F2 int32 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]string)
    var v3 int64 = 0
    var v4 int64 = (-(v3))
    v4 = v3
    var v5 uint64 = 130
    for fi6 := 0; fi6 < 1; fi6++ {
        _ = len(v2)
        switch "a" {
        case "foo bar":
            if (v3 < (v4 - v3)) {
                if ("foo bar" >= "a") {
                    var v7 string = "a"
                    v8 := []int{fi6, fi6, 128}
                    _ = v7
                    _ = v8
                }
                fmt.Println(v1, fi6, v4)
            } else {
                if true {
                    var v9 uint32 = v1
                    _ = v9
                } else {
                    var v10 int8 = 42
                    _ = v10
                }
                var v11 bool = (1 != 32767)
                _ = v11
            }
        case "\n\t":
            switch (fi6 - (fi6 - -1)) {
            case 2147483647:
                switch fi6 {
                case -2147483648:
                    for fi12 := 0; fi12 < 14; fi12++ {
                        var v13 bool = true
                        if v13 {
                            var v14 int = (-(fi12))
                            _ = v14
                        }
                        v13 = true
                        _ = fi12
                        _ = v13
                    }
                default:
                    fmt.Println(v3)
                }
            case -454:
                if false {
                    var v15 uint16 = 2
                    _ = v15
                } else {
                    if true {
                        v4 -= (v4 % ((v3) | 1))
                        var v16 uint8 = 128
                        _ = v16
                    }
                    switch (fi6 % ((fi6) | 1)) {
                    case -1:
                        switch (fi6 / ((fi6) | 1)) {
                        case 32767:
                            var v17 float32 = 0.1
                            _ = v17
                        }
                    case 127:
                        var v18 uint16 = 1
                        _ = v18
                    case -373:
                        var v19 int16 = 32767
                        _ = v19
                    default:
                        switch "\n\t" {
                        case "hello":
                            var v20 int64 = (v4 + (v3 + v4))
                            _ = v20
                        case "a":
                            if (-128 <= 83) {
                                v1 += 367
                                if true {
                                    for fi21 := 0; fi21 < 17; fi21++ {
                                        switch "" {
                                        case "":
                                            var v22 string = "a"
                                            _ = v22
                                        }
                                        _ = fi21
                                    }
                                    fmt.Println(fi6, fi6, v5)
                                    if (!(float32(3.14159) == float32(3.14159))) {
                                        for fi23 := 0; fi23 < 4; fi23++ {
                                            delete(v2, 42)
                                            var v24 uint = 1
                                            v25 := v2[(v1 >> ((v24) & 63))]
                                            _ = fi23
                                            _ = v24
                                            _ = v25
                                        }
                                    } else {
                                        v5 = v5
                                        v4 = (v4 | (v3 / ((v4) | 1)))
                                    }
                                } else {
                                    var v26 int64 = v3
                                    _ = v26
                                }
                            }
                        case "foo bar":
                            v1 = (^(v1))
                        }
                    }
                }
            }
        }
        v27 := GenS1{F0: (v1 ^ v1)}
        _ = fi6
        _ = v27
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        v5 = v5
        if (65535 <= 802) {
            for fi29 := 0; fi29 < 15; fi29++ {
                v4 += 1
                v30 := GenS2{F0: fi29, F1: 0}
                var v31 int8 = 0
                _ = fi29
                _ = v30
                _ = v31
            }
        } else {
            for fi32 := 0; fi32 < 13; fi32++ {
                if (-128 == 1) {
                    v5 -= 890
                    v33 := GenS2{F0: (fi32 / (((fi32 + (-(fi28)))) | 1)), F1: -32768}
                    v2[v1] = "foo bar"
                    _ = v33
                }
                _ = fi32
            }
            var v34 int32 = -1
            var v35 float64 = 0.1
            _ = v34
            _ = v35
        }
        var v36 int = 292
        _ = fi28
        _ = v36
    }
    var v37 uint32 = (^(v1))
    v37 = v1
    v38 := GenS3{F0: false, F1: (v3 - (^(v3))), F2: -2147483648}
    v39 := GenS1{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
