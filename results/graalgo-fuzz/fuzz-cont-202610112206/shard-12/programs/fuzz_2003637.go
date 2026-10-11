package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int64, p1 int32, p2 uint) int {
    v5 := []int{-342, -333}
    if len(v5) > 0 {
        v5[0] = 42
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 32768
}

func main() {
    var v1 int32 = 0
    var v2 uint32 = 359
    var v3 bool = (284 < 4294967295)
    var v4 uint32 = 636
    v2 = v2
    v6 := []int{42, 4294967295, 607}
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int16 = -32768
        _ = fi7
        _ = v8
    }
    var v9 uint16 = 0
    if false {
        var v10 int32 = 2147483647
        _ = v10
    } else {
        var v11 string = (("\n\t" + "foo bar") + "a")
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 string = v11
            _ = fi12
            _ = v13
        }
        if (!true) {
            v4 = (v2 >> ((68) & 63))
            var v14 uint = 42
            if v3 {
                if (!v3) {
                    v15 := GenS1{F0: v1, F1: v1, F2: uint8(v4)}
                    _ = v15
                } else {
                    var v16 float32 = (1e10 * 3.14159)
                    _ = v16
                }
            } else {
                var v17 int16 = -32768
                v11 = v11
                _ = v17
            }
            _ = v14
        } else {
            for fi18 := 0; fi18 < 16; fi18++ {
                v3 = (42 != 836)
                v2 = v4
                v19 := GenS1{F0: (v1 & v1), F1: (v1 + v1), F2: 127}
                _ = fi18
                _ = v19
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        _ = v11
    }
    var v20 int16 = -810
    v2 -= (^(v2))
    v21 := GenS2{F0: 255, F1: 1}
    fmt.Println(v4)
    if true {
        fmt.Println(v2, v9, v1)
        if false {
            if true {
                fmt.Println(v2, v2, v4)
                fmt.Println(v2, v2, v9)
            }
        } else {
            if (!(("hello" < ((("a" + ("foo bar" + "\n\t")) + "foo bar") + "foo bar")) && false)) {
                fmt.Println(v20, v1)
                var v22 uint = 255
                _ = v22
            } else {
                switch genFn0(-9223372036854775808, v1, 2) {
                case 32767:
                    if (uint64(v4) != 855) {
                        v4 -= (v2 << ((59) & 63))
                        if (4294967295 <= 42) {
                            if v3 {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                v23 := GenS3{F0: 284}
                                if len(v6) > 0 {
                                    v6[0] = 127
                                }
                                _ = v23
                            }
                        }
                    } else {
                        var v24 int = 32767
                        _ = v24
                    }
                }
                var v25 int8 = -128
                v25 -= v25
                _ = v25
            }
            if v3 {
                v4 = (v2 & v4)
                if v3 {
                    v3 = v3
                    v9 = (v9 | (v9 % ((v9) | 1)))
                } else {
                    switch ("hello" + "a") {
                    case "a":
                        v26 := GenS2{F0: 255, F1: int64(v20)}
                        _ = v26
                    }
                    v2 += (v4 % ((v4) | 1))
                }
                for fi27 := 0; fi27 < 2; fi27++ {
                    var v28 float64 = (0.0 * 3.14159)
                    for fi29 := 0; fi29 < 2; fi29++ {
                        v30 := GenS2{F0: 128, F1: -9223372036854775808}
                        v31 := GenS3{F0: 4294967295}
                        switch 0 {
                        case 128:
                            v9 = 238
                        default:
                            v1 = v1
                        }
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = fi27
                    _ = v28
                }
            }
        }
    } else {
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 int32 = -2147483648
            _ = fi32
            _ = v33
        }
    }
    v34 := GenS4{F0: (v1 * (v1 << ((416) & 63)))}
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v6
    _ = v9
}
