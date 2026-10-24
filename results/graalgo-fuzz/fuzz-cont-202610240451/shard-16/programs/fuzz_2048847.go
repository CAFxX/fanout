package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint }
type GenS2 struct { F0 int8; F1 uint8 }

func main() {
    v1 := []uint32{0}
    var v2 uint8 = 60
    v3 := make(map[uint32]int)
    v2 -= 168
    v2 = v2
    if ("" > "foo bar") {
        v3[534] = 9223372036854775807
        fmt.Println(v2)
    } else {
        for fi4 := 0; fi4 < 8; fi4++ {
            var v5 uint64 = 42
            _ = fi4
            _ = v5
        }
        switch 9223372036854775807 {
        case -9223372036854775808:
            for fi6 := 0; fi6 < 12; fi6++ {
                var v7 int = (fi6 / (((fi6 % ((fi6) | 1))) | 1))
                var v8 uint8 = (v2 << ((42) & 63))
                if (!(3.14159 > float64(1.5))) {
                    fmt.Println(v2)
                } else {
                    if true {
                        _ = len(v3)
                    } else {
                        fmt.Println(fi6)
                        v2 += (v2 & v8)
                        var v9 uint32 = 2
                        _ = v9
                    }
                    var v10 uint = 807
                    _ = v10
                }
                _ = fi6
                _ = v7
                _ = v8
            }
        case 795:
            for fi11 := 0; fi11 < 13; fi11++ {
                v12 := GenS1{F0: -1, F1: (3.14159 != float32(-2.25)), F2: 436}
                _ = fi11
                _ = v12
            }
        default:
            fmt.Println(v2)
        }
    }
    v2 -= 69
    var v13 int32 = -430
    var v14 uint16 = 65535
    var v15 float64 = 0.0
    for fi16 := 0; fi16 < 4; fi16++ {
        if (500 == 2) {
            if true {
                if true {
                    var v17 float64 = float64((((1.5 * 1e10) / 1e10) - 1.5))
                    if false {
                        for fi18 := 0; fi18 < 10; fi18++ {
                            v14 = v14
                            v2 = (v2 | v2)
                            if (-1 <= 1) {
                                if (((true && false) || true) && true) {
                                    fmt.Println(v17, v13)
                                } else {
                                    fmt.Println(v14)
                                    v2 = v2
                                    var v19 bool = true
                                    _ = v19
                                }
                            } else {
                                switch (("a" + "hello") + "hello") {
                                case "":
                                    v2 += (^(v2))
                                case "a":
                                    if (0.0 != 3.14159) {
                                        fmt.Println(v2, v17)
                                        delete(v3, 4294967295)
                                        var v20 bool = (!true)
                                        _ = v20
                                    }
                                }
                                var v21 int8 = -48
                                _ = v21
                            }
                            v22 := GenS1{F0: 253, F1: ((!(3.14159 == 3.14159)) || ((v2 << ((42) & 63)) == (v2 + (v2 - (v2 * (v2 / ((v2) | 1))))))), F2: 1}
                            _ = fi18
                            _ = v22
                        }
                        v15 = 1.5
                    } else {
                        var v23 uint16 = (^(v14))
                        var v24 bool = false
                        _ = v23
                        _ = v24
                    }
                    _ = v17
                }
                fmt.Println(fi16)
                v2 = (v2 >> ((65535) & 63))
            }
            fmt.Println(fi16, v14, v14)
            fmt.Println(v2, v2)
        } else {
            switch fi16 {
            case 873:
                v25 := GenS2{F0: 100, F1: (v2 ^ (v2 | (^(v2))))}
                _ = v25
            case 580:
                v13 = 42
            default:
                var v26 uint8 = (v2 * v2)
                _ = v26
            }
            if (!true) {
                var v27 uint64 = 42
                if false {
                    v14 = 32767
                    v28 := GenS2{F0: 75, F1: v2}
                    _ = v28
                } else {
                    if (19 < 18446744073709551615) {
                        if ((v2 >= v2) && (42 < 913)) {
                            var v29 string = "a"
                            _ = len(v3)
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 12; fi30++ {
                            var v31 int16 = -32768
                            var v32 uint32 = 2147483647
                            v31 = v31
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    }
                }
                _ = v27
            }
            for fi33 := 0; fi33 < 20; fi33++ {
                v14 -= 42
                v34 := GenS1{F0: 1, F1: false, F2: 1}
                var v35 string = "hello"
                _ = fi33
                _ = v34
                _ = v35
            }
        }
        _ = fi16
    }
    var v36 bool = false
    if v36 {
        fmt.Println(v2)
    }
    var v37 uint16 = (v14 + (v14 / (((v14 - uint16(v2))) | 1)))
    var v38 int32 = 42
    fmt.Println(v14)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}
