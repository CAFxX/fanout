package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int }

func main() {
    var v1 uint16 = 0
    v2 := []int{42}
    v3 := []int{32767, 887, 226, 1}
    var v4 int8 = 100
    v5 := []uint32{2, 615, 116}
    var v6 uint32 = 0
    fmt.Println(v4, v1, v6)
    if false {
        if true {
            var v7 bool = (v1 != 255)
            if (4294967295 == (v6 | (v6 ^ (v6 - (^(v6)))))) {
                var v8 uint64 = 0
                v1 += v1
                var v9 uint8 = uint8(v1)
                _ = v8
                _ = v9
            } else {
                if len(v3) > 0 {
                    v3[0] = 9223372036854775807
                }
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 float32 = float32((float64(3.14159) * (1e10 * (-2.25 + float64((-2.25 * ((1e10 - 1.5) / 1.5)))))))
                    var v12 uint64 = 4294967295
                    for fi13 := 0; fi13 < 8; fi13++ {
                        var v14 float32 = 0.1
                        v7 = v7
                        var v15 uint64 = (v12 | v12)
                        var v16 int = fi13
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
            _ = v7
        }
        switch -9223372036854775808 {
        case -1:
            if false {
                var v17 uint16 = (^(v1))
                _ = v17
            } else {
                for fi18 := 0; fi18 < 10; fi18++ {
                    for fi19 := 0; fi19 < 2; fi19++ {
                        v4 = v4
                        v4 -= v4
                        fmt.Println(v6)
                        _ = fi19
                    }
                    var v20 uint16 = (v1 / ((19) | 1))
                    _ = fi18
                    _ = v20
                }
                switch "foo bar" {
                case "\n\t":
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                case "":
                    if true {
                        for fi21 := 0; fi21 < 20; fi21++ {
                            v6 -= 911
                            _ = fi21
                        }
                        fmt.Println(v1, v4, v1)
                        v22 := GenS1{F0: 1000, F1: 0, F2: 0}
                        _ = v22
                    } else {
                        v5 = append(v5, 10)
                    }
                default:
                    var v23 uint16 = (v1 % ((v1) | 1))
                    _ = v23
                }
            }
        }
    } else {
        if false {
            fmt.Println(v4, v1, v4)
        } else {
            var v24 int64 = 42
            v6 = v6
            _ = v24
        }
    }
    v6 -= v6
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 int16 = 42
        _ = fi25
        _ = v26
    }
    var v27 uint32 = (v6 | v6)
    v27 += (v6 + v27)
    fmt.Println(v1, v27, v6)
    v4 -= 1
    var v28 string = "hello"
    if ((548 == 32768) && ((v1 - v1) != v1)) {
        fmt.Println(v27, v28)
        if (true && true) {
            if (false && false) {
                var v29 uint16 = (v1 & v1)
                _ = len(v3)
                v6 = (v27 + (v27 - 893))
                _ = v29
            }
            v30 := GenS1{F0: -1, F1: -9223372036854775808, F2: 32768}
            for fi31 := 0; fi31 < 20; fi31++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi31
            }
            _ = v30
        } else {
            var v32 uint8 = 1
            for fi33 := 0; fi33 < 4; fi33++ {
                switch (v28 + v28) {
                case "":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "foo bar":
                    fmt.Println(v27)
                }
                var v34 int32 = 549
                _ = fi33
                _ = v34
            }
            v32 += (^(v32))
            _ = v32
        }
    } else {
        fmt.Println(v28, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
