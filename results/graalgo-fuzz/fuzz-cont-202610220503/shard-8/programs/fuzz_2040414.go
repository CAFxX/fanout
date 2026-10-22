package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int64 = 1
    var v2 uint32 = 65535
    var v3 int8 = 0
    v4 := []int{-2147483648, -1, -653, 128}
    switch 42 {
    case 127:
        fmt.Println(v2, v1)
    case 93:
        fmt.Println(v2)
    case 32767:
        fmt.Println(v2, v1)
    }
    if (2 > 18446744073709551615) {
        switch 127 {
        case 9223372036854775807:
            for fi5 := 0; fi5 < 14; fi5++ {
                v1 -= (^(v1))
                v3 += v3
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    v3 -= (v3 / ((v3) | 1))
                case "hello":
                    if true {
                        fmt.Println(v1, v1)
                        v1 = (v1 - -1)
                        v6 := GenS1{F0: 255}
                        _ = v6
                    }
                case "":
                    var v7 int64 = v1
                    _ = v7
                default:
                    fmt.Println(fi5, v3, v1)
                }
                _ = fi5
            }
        case -9223372036854775808:
            switch "\n\t" {
            case "\n\t":
                v2 = v2
            case "":
                for fi8 := 0; fi8 < 11; fi8++ {
                    if (true && true) {
                        if (-9223372036854775808 >= 0) {
                            var v9 uint32 = 1
                            _ = v9
                        } else {
                            var v10 string = ("" + "")
                            v11 := GenS1{F0: 730}
                            if len(v4) > 0 {
                                v4[0] = (fi8 * fi8)
                            }
                            _ = v10
                            _ = v11
                        }
                    }
                    if len(v4) > 0 {
                        v4[0] = fi8
                    }
                    _ = fi8
                }
            case "hello":
                var v12 int64 = (-(v1))
                _ = v12
            default:
                var v13 string = "foo bar"
                _ = v13
            }
        case 32768:
            v1 += 1
        }
        fmt.Println(v3)
    }
    v3 -= 42
    fmt.Println(v1)
    v1 += (v1 / ((0) | 1))
    var v14 uint16 = 874
    if (42 > 42) {
        v15 := GenS1{F0: (v14 + (v14 / ((32767) | 1)))}
        v1 = 1
        if len(v4) > 0 {
            v4[0] = -1
        }
        _ = v15
    }
    if true {
        fmt.Println(v3, v1, v14)
    } else {
        var v16 int64 = 713
        for fi17 := 0; fi17 < 17; fi17++ {
            switch fi17 {
            case 42:
                var v18 int16 = 32767
                _ = v18
            default:
                for fi19 := 0; fi19 < 20; fi19++ {
                    _ = len(v4)
                    var v20 string = ("foo bar" + "\n\t")
                    switch v20 {
                    case "":
                        fmt.Println(v2, v1, v20)
                    case "a":
                        var v21 uint16 = (v14 / ((157) | 1))
                        var v22 uint64 = 18446744073709551615
                        var v23 int32 = 2147483647
                        _ = v21
                        _ = v22
                        _ = v23
                    case "\n\t":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        fmt.Println(v2, fi19)
                    }
                    _ = fi19
                    _ = v20
                }
            }
            if true {
                fmt.Println(v16, v14)
            } else {
                for fi24 := 0; fi24 < 16; fi24++ {
                    if (-32768 <= 32767) {
                        switch 2147483647 {
                        case -1:
                            if true {
                                var v25 int = -820
                                _ = v25
                            }
                        case 787:
                            if (!((-1 < 1000) && true)) {
                                if true {
                                    v1 += int64(fi17)
                                }
                                v14 = v14
                            }
                        }
                        fmt.Println(v16, fi17, v1)
                    } else {
                        v1 = 1
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v14, v1, v3)
                        v4 = append(v4, fi24)
                        v3 -= 0
                    }
                    _ = fi24
                }
                v14 -= 65535
            }
            if (!(128 == fi17)) {
                var v26 int64 = (v16 + v16)
                _ = v26
            } else {
                v3 = -81
                v4 = append(v4, 128)
                for fi27 := 0; fi27 < 2; fi27++ {
                    var v28 uint64 = 684
                    fmt.Println(v3, fi17, fi17)
                    if ((4294967295 < fi17) || true) {
                        v16 -= int64(v28)
                    }
                    _ = fi27
                    _ = v28
                }
            }
            _ = fi17
        }
        _ = v16
    }
    var v29 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
