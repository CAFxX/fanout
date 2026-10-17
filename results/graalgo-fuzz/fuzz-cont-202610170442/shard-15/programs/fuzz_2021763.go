package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint16; F2 int8 }

func main() {
    var v1 int8 = 1
    var v2 uint = 53
    v3 := make(map[uint32]bool)
    var v4 int = 430
    var v5 uint8 = 234
    var v6 float32 = float32(3.14159)
    v7 := GenS1{F0: 4294967295, F1: 42, F2: 293}
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: (^(v2)), F1: 2147483647, F2: 42}
        fmt.Println(v6)
        var v10 int8 = v1
        _ = fi8
        _ = v10
        _ = v9
    }
    if ((3.14159 / 0.1) < ((0.1 / 1e10) + 1e10)) {
        switch (("hello" + "") + "") {
        case "a":
            switch (-(v4)) {
            case -9223372036854775808:
                v4 += 783
            case 2147483647:
                switch (v4 | 127) {
                case 4294967295:
                    v11 := GenS1{F0: v2, F1: -616, F2: 1}
                    _ = v11
                case 127:
                    v1 -= v1
                default:
                    if false {
                        for fi12 := 0; fi12 < 13; fi12++ {
                            v13 := []int{(v4 | fi12), -546, v4, (fi12 & (fi12 & (v4 & (v4 + fi12))))}
                            _ = fi12
                            _ = v13
                        }
                    } else {
                        _ = len(v3)
                        for fi14 := 0; fi14 < 16; fi14++ {
                            v3[2147483647] = (((fi14 | -9223372036854775808) >= -770) || true)
                            fmt.Println(v4, v2, v2)
                            for fi15 := 0; fi15 < 18; fi15++ {
                                v16 := GenS2{F0: 255}
                                v5 = (v5 ^ v5)
                                var v17 bool = true
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                            _ = fi14
                        }
                        v2 += 1
                    }
                }
            case 127:
                if ((("\n\t" + "") == ("" + "foo bar")) && ((-1 <= 0) || (1 == 32767))) {
                    v4 -= 128
                    v1 = v1
                } else {
                    if true {
                        for fi18 := 0; fi18 < 1; fi18++ {
                            var v19 int32 = -2147483648
                            v19 += (-(v19))
                            for fi20 := 0; fi20 < 14; fi20++ {
                                var v21 uint = (v2 / ((v2) | 1))
                                _ = fi20
                                _ = v21
                            }
                            _ = fi18
                            _ = v19
                        }
                        var v22 uint8 = v5
                        delete(v3, 65535)
                        _ = v22
                    } else {
                        v2 = v2
                        v23 := GenS2{F0: 645}
                        v24 := GenS3{F0: -838, F1: 248, F2: v1}
                        _ = v23
                        _ = v24
                    }
                }
            }
        case "hello":
            fmt.Println(v4)
        default:
            v1 = -128
        }
        var v25 string = ((("\n\t" + "") + "a") + "")
        v4 += (v4 | (v4 | int(v5)))
        _ = v25
    } else {
        fmt.Println(v2)
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        v4 = v4
        var v27 string = ""
        var v28 bool = (1e10 > (1.5 * 3.14159))
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v4)
    v1 = (v1 % ((119) | 1))
    v5 = (v5 & v5)
    if true {
        var v29 float32 = 1.5
        for fi30 := 0; fi30 < 18; fi30++ {
            switch "" {
            case "\n\t":
                var v31 uint = 231
                _ = v31
            case "foo bar":
                switch (v4 >> (((^(v2))) & 63)) {
                case 32768:
                    if false {
                        var v32 string = "a"
                        var v33 uint64 = 4294967295
                        v34 := []int{fi30, (fi30 - v4)}
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                case 4294967295:
                    if false {
                        if (v2 <= v2) {
                            var v35 string = "hello"
                            v36 := []int{128, (v4 + v4), (v4 ^ 0)}
                            fmt.Println(fi30)
                            _ = v35
                            _ = v36
                        }
                        v1 += (v1 << ((618) & 63))
                        var v37 uint = (^(v2))
                        _ = v37
                    }
                }
            case "a":
                switch "a" {
                case "foo bar":
                    v38 := GenS3{F0: -1, F1: 0, F2: v1}
                    _ = v38
                case "a":
                    v39 := GenS2{F0: 65535}
                    _ = v39
                default:
                    var v40 int32 = 0
                    _ = v40
                }
            }
            v41 := []int{fi30, v4, 42, (-(v4))}
            _ = fi30
            _ = v41
        }
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
