package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint32 }
func genFn0() uint8 {
    v7 := GenS1{F0: 65535, F1: 2, F2: 2147483647}
    v8 := GenS1{F0: 65535, F1: 559, F2: 2}
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ("hello" + "a")
    var v3 int16 = 1000
    var v4 uint8 = 163
    var v5 uint16 = 2
    v6 := make(map[uint32]uint32)
    switch ("" + v2) {
    case "\n\t":
        var v9 uint8 = (v4 - (v4 - (v4 | (^(v4)))))
        _ = v9
    }
    var v10 uint = 679
    v5 = (v5 + (v5 ^ (v5 / ((v5) | 1))))
    var v11 uint = v10
    _ = len(v6)
    v2 = ((v2 + v2) + v2)
    var v12 string = (v2 + (v2 + v2))
    v13 := []int{42, -9223372036854775808, -2147483648}
    v5 -= v5
    if ((-(v3)) <= v3) {
        delete(v6, 65535)
        var v14 int32 = 0
        for fi15 := 0; fi15 < 16; fi15++ {
            if (4294967295 != 42) {
                v2 = ("\n\t" + "")
                v13 = append(v13, fi15)
            } else {
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 uint = v11
                    if false {
                        if (float64(float32(0.0)) >= (1e10 / 1.0)) {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            v10 += (v17 >> ((642) & 63))
                            _ = len(v13)
                        }
                    } else {
                        switch (fi15 & fi16) {
                        case -967:
                            var v18 uint64 = 4294967295
                            _ = v18
                        }
                        switch 128 {
                        case -592:
                            switch 0 {
                            case 128:
                                v4 -= (^(v4))
                            }
                        case 4294967295:
                            var v19 int = fi15
                            _ = v19
                        case -2147483648:
                            var v20 float64 = (0.1 - -2.25)
                            _ = v20
                        default:
                            if true {
                                v21 := GenS2{F0: 287, F1: 2147483648}
                                _ = v21
                            }
                        }
                        v12 = ("" + (("a" + v12) + "hello"))
                    }
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi15
        }
        _ = v14
    } else {
        var v22 int64 = 9223372036854775807
        fmt.Println(v3, v11, v11)
        v11 -= (v11 >> ((v10) & 63))
        _ = v22
    }
    if ((false && false) || true) {
        var v23 float32 = 0.0
        _ = v23
    } else {
        switch (v2 + v12) {
        case "foo bar":
            for fi24 := 0; fi24 < 17; fi24++ {
                if false {
                    fmt.Println(fi24, v2)
                } else {
                    switch fi24 {
                    case 9223372036854775807:
                        v4 = genFn0()
                    }
                    v13 = append(v13, (fi24 >> (((v10 ^ (^(v10)))) & 63)))
                }
                _ = fi24
            }
        }
        v25 := GenS3{F0: 65535}
        _ = v25
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        if (((!false) || false) && ((v3 + (v3 / ((v3) | 1))) < (-(v3)))) {
            var v27 bool = (!(((1.5 + (0.0 * 0.1)) * 3.14159) < -2.25))
            var v28 uint32 = 773
            _ = v27
            _ = v28
        } else {
            v2 = v12
            switch ("" + ("\n\t" + v2)) {
            case "":
                if true {
                    var v29 float32 = (1e10 * (3.14159 - (3.14159 * (-2.25 + 0.1))))
                    _ = v29
                } else {
                    var v30 int64 = 1
                    _ = v30
                }
            default:
                for fi31 := 0; fi31 < 9; fi31++ {
                    var v32 int = fi26
                    v33 := GenS4{F0: 136, F1: 4294967295}
                    v34 := GenS3{F0: 4294967295}
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
            }
            var v35 uint16 = v5
            _ = v35
        }
        var v36 int64 = 559
        _ = fi26
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
