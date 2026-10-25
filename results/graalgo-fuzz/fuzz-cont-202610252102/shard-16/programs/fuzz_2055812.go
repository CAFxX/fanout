package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    var v3 uint16 = 0
    var v4 uint32 = 302
    v5 := make(map[int]string)
    for fi6 := 0; fi6 < 3; fi6++ {
        switch fi6 {
        case 128:
            var v7 uint64 = 42
            _ = v7
        case 4294967295:
            switch "foo bar" {
            case "a":
                if (true && (!(0 == 656))) {
                    var v8 uint16 = v3
                    fmt.Println(fi6, v4, fi6)
                    var v9 uint = 255
                    _ = v8
                    _ = v9
                }
            default:
                v4 -= 65535
            }
        case 32767:
            var v10 float32 = (0.1 - 1.5)
            _ = v10
        }
        _ = fi6
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        fmt.Println(fi11, v4)
        v4 = 2147483648
        _ = fi11
    }
    switch 127 {
    case -954:
        v3 = (v3 & v3)
    default:
        switch "foo bar" {
        case "hello":
            v12 := GenS1{F0: 32767, F1: 42}
            _ = v12
        case "\n\t":
            delete(v5, -343)
        }
    }
    fmt.Println(v3)
    if true {
        v13 := GenS2{F0: 4294967295}
        _ = v13
    } else {
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 uint32 = v4
            var v16 float32 = 0.1
            v17 := GenS2{F0: fi14}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            if (42 != 9223372036854775807) {
                v3 = (v3 & (^(v3)))
                if ((!(-2147483648 > -1)) && false) {
                    v19 := v1[("" + "a")]
                    _ = v19
                }
            } else {
                v20 := GenS1{F0: (fi18 | 345), F1: 42}
                v3 = (v3 >> ((713) & 63))
                _ = v20
            }
            if true {
                var v21 int64 = -969
                v4 -= v4
                v21 = v21
                _ = v21
            }
            if ((3.14159 + 3.14159) < 0.1) {
                for fi22 := 0; fi22 < 3; fi22++ {
                    v4 = (^(v4))
                    if (1000 != 8) {
                        var v23 int64 = 1
                        v3 = 1
                        var v24 uint8 = 127
                        _ = v23
                        _ = v24
                    } else {
                        switch fi22 {
                        case 42:
                            if (42 < -1) {
                                var v25 uint = 1
                                switch fi22 {
                                case 128:
                                    v26 := GenS2{F0: -900}
                                    _ = v26
                                }
                                _ = v25
                            } else {
                                v3 -= (v3 % ((65535) | 1))
                                var v27 float32 = (0.1 * 1e10)
                                _ = v27
                            }
                        }
                        v28 := GenS3{F0: (v3 << ((63) & 63)), F1: (v4 + v4)}
                        _ = v28
                    }
                    v29 := GenS4{F0: 32767}
                    _ = fi22
                    _ = v29
                }
            }
            _ = fi18
        }
        v30 := GenS1{F0: 728, F1: -2147483648}
        _ = v30
    }
    v31 := GenS5{F0: 128}
    v32 := []int{}
    var v33 int = -882
    var v34 uint8 = 255
    v33 = (-(v33))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
