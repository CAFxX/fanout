package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint; F1 int64; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int64; F2 int16 }

func main() {
    var v1 float64 = (((0.1 + ((3.14159 * -2.25) + (0.1 * ((0.1 - -2.25) - float64(1.5))))) / ((float64(3.14159) / -2.25) + (0.1 + (float64(0.1) / 1.0)))) * 3.14159)
    v2 := []string{(("hello" + "a") + (("foo bar" + "a") + (("foo bar" + "hello") + "foo bar")))}
    v3 := []int{127, 987}
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 0
    for fi6 := 0; fi6 < 11; fi6++ {
        switch (-(fi6)) {
        case 42:
            if false {
                v1 = (((v1 / v1) * v1) * v1)
            } else {
                delete(v4, 2147483648)
                v7 := GenS1{F0: (v5 * (v5 << ((573) & 63))), F1: 127}
                v1 = (float64(0.0) / float64(-2.25))
                _ = v7
            }
        }
        var v8 string = ""
        var v9 int32 = int32(v5)
        _ = fi6
        _ = v8
        _ = v9
    }
    if true {
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: v5, F1: -54}
            v12 := GenS1{F0: 65535, F1: -1}
            v13 := GenS2{F0: 1, F1: 9223372036854775807, F2: 65535}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    }
    v5 = v5
    var v14 uint8 = 2
    v5 = v5
    if (3.14159 != 1e10) {
        if (!false) {
            v15 := GenS3{F0: 4294967295}
            var v16 int64 = 282
            _ = v15
            _ = v16
        } else {
            if (127 < 0) {
                for fi17 := 0; fi17 < 16; fi17++ {
                    switch 2147483647 {
                    case 32767:
                        v5 -= 378
                    default:
                        switch fi17 {
                        case 9223372036854775807:
                            var v18 uint = 627
                            _ = v18
                        case 42:
                            if (1 <= 2) {
                                for fi19 := 0; fi19 < 15; fi19++ {
                                    var v20 int16 = 32767
                                    _ = fi19
                                    _ = v20
                                }
                                var v21 int8 = 127
                                fmt.Println(v5, v1, v21)
                                _ = v21
                            } else {
                                if ((0.0 - 3.14159) >= 0.0) {
                                    v14 += v14
                                    for fi22 := 0; fi22 < 15; fi22++ {
                                        if true {
                                            v23 := v4[377]
                                            _ = v23
                                        } else {
                                            fmt.Println(fi22)
                                            fmt.Println(v5)
                                            v5 = v5
                                        }
                                        v5 += 0
                                        _ = fi22
                                    }
                                } else {
                                    _ = len(v4)
                                    var v24 int16 = -1
                                    _ = v24
                                }
                            }
                        case 0:
                            v4[720] = 542
                        }
                    }
                    _ = fi17
                }
                var v25 uint64 = 1
                _ = v25
            } else {
                var v26 int32 = 420
                v1 += 0.0
                _ = v26
            }
            v1 = v1
        }
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        var v28 int = fi27
        _ = fi27
        _ = v28
    }
    v5 -= 32767
    v14 = (v14 & (^(v14)))
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 int = 42
        if ((127 >= int8(v5)) || (!(!((0 == 100) || false)))) {
            v31 := GenS1{F0: (^(v5)), F1: 1}
            delete(v4, 2)
            _ = v31
        } else {
            for fi32 := 0; fi32 < 14; fi32++ {
                var v33 int = (^(fi29))
                v30 += (fi32 ^ (fi32 / ((-2147483648) | 1)))
                var v34 int64 = 0
                _ = fi32
                _ = v33
                _ = v34
            }
        }
        _ = fi29
        _ = v30
    }
    v4[958] = 2147483648
    v35 := GenS4{F0: 4294967295, F1: 42, F2: 7}
    fmt.Println(v14, v14)
    v36 := GenS4{F0: 42, F1: 9223372036854775807, F2: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
