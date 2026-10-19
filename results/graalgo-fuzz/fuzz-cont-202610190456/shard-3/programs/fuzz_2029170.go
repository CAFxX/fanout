package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]uint32)
    var v3 uint = 0
    var v4 float64 = 0.1
    if (-554 >= 0) {
        var v5 float32 = float32(3.14159)
        v6 := []int{42, 42, -2147483648, 1}
        switch "" {
        case "foo bar":
            var v7 int8 = -1
            _ = v7
        case "":
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 uint32 = 2147483647
                for fi10 := 0; fi10 < 5; fi10++ {
                    if false {
                        var v11 uint8 = 0
                        _ = v11
                    } else {
                        var v12 int = 4294967295
                        if (!false) {
                            delete(v2, v9)
                            if true {
                                if (v9 != v9) {
                                    v3 = (v3 / ((52) | 1))
                                }
                                v9 += (v9 / (((v9 & v9)) | 1))
                            }
                            var v13 uint16 = 1
                            _ = v13
                        } else {
                            var v14 uint8 = 2
                            fmt.Println(v4)
                            v15 := GenS1{F0: v14}
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if len(v6) > 0 {
                                v6[0] = (v12 % ((v12) | 1))
                            }
                            _ = v14
                            _ = v15
                        }
                        v4 += v4
                        _ = v12
                    }
                    v6 = append(v6, fi8)
                    _ = fi10
                }
                v16 := v2[(^(v9))]
                _ = fi8
                _ = v16
                _ = v9
            }
        }
        _ = v5
        _ = v6
    } else {
        switch ((("\n\t" + "hello") + "") + ("hello" + "\n\t")) {
        case "":
            fmt.Println(v1, v1, v1)
        }
        v1 = (v1 % ((4294967295) | 1))
        _ = len(v2)
    }
    v4 -= v4
    fmt.Println(v3, v4, v4)
    var v17 uint32 = 4294967295
    fmt.Println(v3, v4, v3)
    var v18 int64 = 42
    var v19 uint64 = v1
    if false {
        if true {
            v20 := []int{127, 0, 42}
            if ((v1 % ((2) | 1)) != v1) {
                if true {
                    var v21 uint32 = 4294967295
                    _ = v21
                }
                fmt.Println(v17, v18)
            } else {
                v20 = append(v20, -1)
                v19 += v1
                switch -359 {
                case 2147483647:
                    fmt.Println(v3)
                case 127:
                    v18 += (v18 & 0)
                default:
                    v3 -= (v3 % ((v3) | 1))
                }
            }
            _ = v20
        } else {
            v22 := v2[uint32(v19)]
            if true {
                v23 := GenS1{F0: 128}
                _ = v23
            } else {
                var v24 float64 = 0.1
                v22 = (v22 - v22)
                _ = v24
            }
            _ = v22
        }
    }
    var v25 bool = (true || ((1e10 - 0.0) < (((0.0 + (1e10 - 0.0)) + 1.5) + 0.1)))
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 int8 = 103
        _ = fi26
        _ = v27
    }
    v19 = v1
    if ((v3 / ((42) | 1)) >= 18446744073709551615) {
        v19 -= (v1 - (v1 - v1))
    } else {
        if true {
            var v28 uint32 = v17
            _ = v28
        } else {
            for fi29 := 0; fi29 < 3; fi29++ {
                v4 = v4
                var v30 uint32 = 12
                _ = fi29
                _ = v30
            }
            v3 += uint(v1)
            for fi31 := 0; fi31 < 5; fi31++ {
                switch "\n\t" {
                case "hello":
                    v4 = v4
                default:
                    var v32 float32 = 1.5
                    _ = v32
                }
                _ = fi31
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
