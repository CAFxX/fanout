package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 uint16; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 uint; F2 uint }
type GenS7 struct { F0 uint8 }

func main() {
    var v1 string = "a"
    var v2 int8 = 100
    v3 := make(map[uint32]bool)
    var v4 int64 = 1
    var v5 int32 = 42
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: 65535}
        if (4294967295 >= 4294967295) {
            v1 = v1
            v8 := GenS2{F0: ("a" + (v1 + (v1 + v1)))}
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    v5 += v5
    v9 := GenS2{F0: (((v1 + "foo bar") + ("a" + "a")) + v1)}
    var v10 int = -2147483648
    if (!false) {
        v5 -= (v5 ^ v5)
        switch v1 {
        case "foo bar":
            for fi11 := 0; fi11 < 3; fi11++ {
                v2 = (v2 + (v2 * (v2 >> ((18446744073709551615) & 63))))
                if (true && true) {
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v13 := GenS3{F0: fi11, F1: 202, F2: v10}
                        var v14 int = fi12
                        fmt.Println(v5, v10, v1)
                        v10 -= (v10 | v10)
                        v10 = v10
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v16 := GenS4{F0: 533}
                        _ = fi15
                        _ = v16
                    }
                    var v17 bool = (0.0 >= 1e10)
                    _ = v17
                }
                v10 = 4294967295
                _ = fi11
            }
        case "\n\t":
            v4 += v4
        default:
            for fi18 := 0; fi18 < 16; fi18++ {
                v19 := GenS4{F0: 42}
                v2 = v2
                if true {
                    v20 := GenS5{F0: 42}
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v10 = -940
                        var v22 uint64 = 42
                        _ = fi21
                        _ = v22
                    }
                    if (!(v4 == v4)) {
                        fmt.Println(v4)
                        switch fi18 {
                        case 2147483647:
                            v4 += -1
                        case 678:
                            fmt.Println(v10, fi18, v5)
                        default:
                            switch "foo bar" {
                            case "a":
                                v23 := GenS4{F0: 1}
                                _ = v23
                            default:
                                for fi24 := 0; fi24 < 9; fi24++ {
                                    var v25 string = ("\n\t" + "a")
                                    v2 += 35
                                    v26 := []int{fi24, v10}
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                            }
                        }
                        switch v10 {
                        case 216:
                            var v27 float64 = (1.5 - ((float64(((0.0 + (1e10 + (((3.14159 - 0.1) - 0.1) - float32(0.1)))) - 0.0)) + float64(1e10)) + (float64(float32(0.0)) / (1e10 + 1.5))))
                            _ = v27
                        case 0:
                            var v28 int = int(v5)
                            _ = v28
                        }
                    }
                    v2 = 127
                }
                _ = fi18
                _ = v19
            }
        }
        if true {
            v3[65535] = (!(!false))
            if (!(false && (true && false))) {
                v2 -= (v2 + 42)
                v29 := GenS6{F0: 65535, F1: 773, F2: 42}
                _ = v29
            }
            for fi30 := 0; fi30 < 4; fi30++ {
                v31 := GenS7{F0: 110}
                _ = fi30
                _ = v31
            }
        }
    } else {
        if (!(true && false)) {
            v5 += (v5 ^ -681)
        } else {
            if (1 <= 142) {
                var v32 uint32 = 855
                fmt.Println(v5, v4, v2)
                _ = v32
            }
        }
        var v33 int = (^(v10))
        _ = v33
    }
    fmt.Println(v4)
    var v34 uint32 = 212
    fmt.Println(v10, v5)
    var v35 uint32 = v34
    v10 = (v10 | -73)
    for fi36 := 0; fi36 < 18; fi36++ {
        v37 := GenS2{F0: "hello"}
        _ = len(v3)
        v1 = (((("foo bar" + (v1 + (v1 + (v1 + v1)))) + ((v1 + v1) + v1)) + v1) + "hello")
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v9
}
