package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }
func genFn0() bool {
    v5 := []int{32768, 32767, 69}
    var v6 float64 = -2.25
    if len(v5) > 0 {
        v5[0] = 42
    }
    v6 -= (1.5 - 3.14159)
    switch ("" + "a") {
    case "\n\t":
        v6 -= -2.25
    case "hello":
        fmt.Println(v6)
    case "":
        if (((0 < 65535) && (128 == 245)) && ("a" < "")) {
            var v7 int16 = 1
            _ = v7
        } else {
            switch 9223372036854775807 {
            case 286:
                v8 := GenS1{F0: "hello", F1: "hello"}
                _ = v8
            case 127:
                v6 = (((3.14159 / (1.5 / -2.25)) / v6) / 3.14159)
            }
            var v9 uint32 = 202
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 3; fi11++ {
                fmt.Println(fi11)
                v6 = v6
                if (848 == 1) {
                    if (42 >= 789) {
                        fmt.Println(v6, v6)
                        if false {
                            if true {
                                var v12 uint8 = 2
                                if len(v5) > 0 {
                                    v5[0] = (fi11 & (fi10 & fi10))
                                }
                                _ = v12
                            }
                        }
                    }
                    fmt.Println(fi10, v6)
                }
                v6 -= float64(1e10)
                fmt.Println(v6, v6)
                _ = fi11
            }
            fmt.Println(v6)
            _ = fi10
        }
    }
    _ = v5
    _ = v6
    return ((!false) && true)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"", (("foo bar" + "hello") + "foo bar")}
    var v3 uint32 = 2147483647
    var v4 uint32 = 4294967295
    v13 := v1[(v3 << ((827) & 63))]
    v4 = v3
    v4 += (v3 | v3)
    if (-463 <= 9223372036854775807) {
        if false {
            if false {
                v13 = v13
                for fi14 := 0; fi14 < 16; fi14++ {
                    v4 = v4
                    if true {
                        for fi15 := 0; fi15 < 13; fi15++ {
                            v16 := GenS2{F0: 255}
                            _ = fi15
                            _ = v16
                        }
                        v1[(^(v3))] = (v13 + (((v13 + (((v13 + "foo bar") + "foo bar") + "foo bar")) + (v13 + v13)) + "a"))
                        _ = len(v2)
                    }
                    if (v13 > (v13 + ("" + (((v13 + "\n\t") + v13) + "a")))) {
                        v17 := GenS3{F0: 42}
                        _ = len(v2)
                        _ = v17
                    } else {
                        fmt.Println(v3, fi14, v4)
                        for fi18 := 0; fi18 < 16; fi18++ {
                            v19 := v1[623]
                            _ = fi18
                            _ = v19
                        }
                        v20 := GenS3{F0: 42}
                        _ = v20
                    }
                    _ = fi14
                }
                fmt.Println(v4, v3)
            }
            var v21 uint8 = 127
            _ = v21
        } else {
            var v22 int32 = 1
            var v23 uint64 = 397
            _ = v22
            _ = v23
        }
        if (42 != 1) {
            v3 = 542
            v3 += v4
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 int16 = 42
                _ = fi24
                _ = v25
            }
        } else {
            _ = len(v2)
            for fi26 := 0; fi26 < 1; fi26++ {
                if len(v2) > 0 {
                    v2[0] = v13
                }
                switch fi26 {
                case 128:
                    v3 = v4
                case 858:
                    v4 = (^(v3))
                default:
                    fmt.Println(fi26, v4)
                }
                _ = fi26
            }
        }
        v3 = v4
    }
    fmt.Println(v3, v13, v4)
    _ = len(v1)
    v13 = ""
    v13 = ""
    v27 := GenS2{F0: 474}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
