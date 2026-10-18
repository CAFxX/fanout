package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint32; F2 bool }

func main() {
    var v1 float32 = 0.0
    var v2 int = -9223372036854775808
    var v3 int64 = 916
    var v4 bool = true
    v5 := []int{(v2 + -845)}
    var v6 uint8 = 255
    if v4 {
        if (!(!v4)) {
            if false {
                v6 += (v6 - v6)
                for fi7 := 0; fi7 < 9; fi7++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi7
                }
            }
            var v8 int64 = v3
            if (true || false) {
                switch (v2 * -9223372036854775808) {
                case 2147483647:
                    var v9 uint = 531
                    _ = v9
                case 32768:
                    v6 = 2
                case -477:
                    for fi10 := 0; fi10 < 20; fi10++ {
                        for fi11 := 0; fi11 < 5; fi11++ {
                            var v12 uint8 = v6
                            _ = fi11
                            _ = v12
                        }
                        var v13 uint = 586
                        _ = fi10
                        _ = v13
                    }
                }
            }
            _ = v8
        } else {
            var v14 int8 = 42
            _ = v14
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            if false {
                switch fi15 {
                case 32767:
                    if v4 {
                        for fi16 := 0; fi16 < 3; fi16++ {
                            switch "hello" {
                            case "a":
                                v6 -= 224
                            default:
                                v17 := GenS1{F0: -2147483648}
                                _ = v17
                            }
                            _ = fi16
                        }
                        if (1.5 >= v1) {
                            v1 -= float32(-2.25)
                        }
                    }
                case 127:
                    if v4 {
                        var v18 int = (fi15 - (-(fi15)))
                        v3 = v3
                        fmt.Println(v3, v18, v2)
                        _ = v18
                    } else {
                        switch "" {
                        case "hello":
                            if (v4 && v4) {
                                var v19 uint = 606
                                _ = v19
                            }
                        }
                        if v4 {
                            var v20 uint16 = 0
                            _ = v20
                        } else {
                            var v21 uint32 = 565
                            var v22 uint32 = v21
                            v23 := GenS1{F0: -2147483648}
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    }
                default:
                    var v24 uint16 = 65535
                    _ = v24
                }
                v4 = (fi15 <= 1)
            }
            _ = fi15
        }
        v25 := GenS2{F0: 42, F1: 18446744073709551615}
        _ = v25
    } else {
        v26 := GenS1{F0: 995}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if len(v5) > 0 {
            v5[0] = (v2 % ((v2) | 1))
        }
        _ = v26
    }
    v27 := GenS3{F0: 18446744073709551615}
    if true {
        v28 := GenS3{F0: 65535}
        _ = v28
    }
    v2 -= 127
    v29 := GenS2{F0: 103, F1: 4294967295}
    v2 -= (v2 & v2)
    if (3.14159 == 1.5) {
        var v30 int8 = -1
        _ = v30
    } else {
        v31 := GenS4{F0: 42, F1: 2, F2: (0 == 42)}
        var v32 uint16 = 438
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
