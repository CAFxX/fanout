package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 int8; F2 uint }
type GenS5 struct { F0 int8; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint16; F1 int16; F2 bool }
type GenS7 struct { F0 int64; F1 int32 }
func genFn0(p0 bool, p1 uint) uint16 {
    v5 := GenS1{F0: (p1 >> ((923) & 63)), F1: "hello", F2: 0}
    p0 = p0
    switch ("foo bar" + "\n\t") {
    case "":
        if (!((!(!p0)) && p0)) {
            p1 += (p1 - 255)
        }
    case "a":
        var v6 string = "\n\t"
        _ = v6
    }
    p0 = p0
    p0 = p0
    _ = p0
    _ = p1
    _ = v5
    return 647
}
func genFn1(p0 uint16, p1 int16) int {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint64 = 397
        v9 := GenS1{F0: 598, F1: "\n\t", F2: 1}
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        if (2 >= 983) {
            for fi10 := 0; fi10 < 5; fi10++ {
                fmt.Println(p0, p1, fi10)
                p0 = (p0 & genFn0((true || (805 > 2147483647)), 18446744073709551615))
                v11 := GenS1{F0: 344, F1: "a", F2: 658}
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 int8 = -23
                _ = fi12
                _ = v13
            }
        } else {
            for fi14 := 0; fi14 < 12; fi14++ {
                fmt.Println(fi14)
                var v15 int64 = -1
                _ = fi14
                _ = v15
            }
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 uint = 65535
                if true {
                    v17 += uint(p0)
                } else {
                    v18 := GenS2{F0: 42, F1: p0}
                    var v19 uint32 = 629
                    v20 := GenS2{F0: -540, F1: p0}
                    _ = v18
                    _ = v19
                    _ = v20
                }
                _ = fi16
                _ = v17
            }
            if (!true) {
                p1 = 42
            }
        }
    } else {
        if false {
            v21 := GenS3{F0: true}
            _ = v21
        }
        v22 := []int{-554, -368, -543}
        _ = v22
    }
    var v23 uint32 = 42
    p0 = (p0 % (((p0 % (((p0 >> ((42) & 63))) | 1))) | 1))
    _ = p0
    _ = p1
    _ = v23
    return 127
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 42
    var v3 uint8 = 1
    v4 := []uint32{(^(v2))}
    for fi24 := 0; fi24 < 1; fi24++ {
        if (genFn0((!(42 <= 42)), 255) <= 65535) {
            switch (fi24 | (fi24 >> ((uint(v2)) & 63))) {
            case 127:
                var v25 string = "hello"
                _ = v25
            }
        }
        _ = fi24
    }
    var v26 uint8 = v3
    var v27 float32 = 1e10
    fmt.Println(v2)
    v28 := GenS4{F0: "", F1: -128, F2: 450}
    v2 = 42
    v2 += 2147483648
    if (!true) {
        var v29 float64 = 1e10
        for fi30 := 0; fi30 < 17; fi30++ {
            v3 = (v3 & (v3 ^ (v26 / ((v26) | 1))))
            var v31 string = "a"
            var v32 uint8 = (v3 - (v3 + v3))
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = v29
    } else {
        fmt.Println(v2)
        v2 += (^(v2))
    }
    v3 = (v26 % ((128) | 1))
    v33 := GenS5{F0: 13, F1: 773, F2: 32767}
    if true {
        v34 := GenS2{F0: -158, F1: 32767}
        fmt.Println(v26, v27)
        v2 -= uint32(v3)
        _ = v34
    }
    var v35 float64 = 0.0
    if len(v4) > 0 {
        v4[0] = (v2 >> ((65535) & 63))
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v27, v3)
    switch -9223372036854775808 {
    case -9223372036854775808:
        switch genFn1(528, -555) {
        case 127:
            _ = len(v4)
        default:
            if (false || false) {
                for fi36 := 0; fi36 < 17; fi36++ {
                    if (!(0.0 != (v35 / v35))) {
                        delete(v1, "foo bar")
                        v37 := GenS6{F0: 32767, F1: 878, F2: ((true || false) || (1000 == 740))}
                        for fi38 := 0; fi38 < 6; fi38++ {
                            var v39 int32 = 444
                            v40 := GenS3{F0: true}
                            v41 := v1[(("\n\t" + ("" + "a")) + (((("a" + "hello") + "a") + "foo bar") + "foo bar"))]
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        _ = v37
                    }
                    fmt.Println(v35)
                    v42 := GenS7{F0: 795, F1: -497}
                    _ = fi36
                    _ = v42
                }
                v43 := GenS2{F0: -697, F1: 482}
                switch genFn1(32767, -32768) {
                case -503:
                    v27 = v27
                case 127:
                    v35 = float64(0.0)
                default:
                    v35 -= v35
                }
                _ = v43
            } else {
                if len(v4) > 0 {
                    v4[0] = uint32(v3)
                }
                if false {
                    _ = len(v4)
                } else {
                    v35 = v35
                    var v44 bool = (42 < 561)
                    _ = v44
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v35
    _ = v4
}
