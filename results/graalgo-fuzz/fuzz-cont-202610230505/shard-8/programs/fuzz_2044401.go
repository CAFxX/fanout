package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 int8 }
func genFn0() uint {
    v5 := GenS1{F0: 0, F1: 0}
    v6 := GenS1{F0: 269, F1: 42}
    switch -907 {
    case 9223372036854775807:
        var v7 int8 = 127
        _ = v7
    case 1:
        var v8 float32 = float32(0.1)
        _ = v8
    case -183:
        v9 := []int{42}
        _ = v9
    }
    var v10 uint = 65535
    _ = v10
    _ = v5
    _ = v6
    return v10
}

func main() {
    v1 := []uint32{42, 42, 42}
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    var v4 uint = 808
    var v11 uint32 = 93
    for fi12 := 0; fi12 < 6; fi12++ {
        v3[v11] = true
        v11 -= v11
        _ = fi12
    }
    var v13 uint8 = 127
    v14 := v2[v11]
    if false {
        v11 -= (v11 / ((1) | 1))
    } else {
        _ = len(v3)
        v15 := GenS2{F0: uint8(v4)}
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 8; fi17++ {
            _ = len(v3)
            v18 := GenS2{F0: 56}
            var v19 uint64 = 0
            _ = fi17
            _ = v18
            _ = v19
        }
        v20 := GenS3{F0: (v14 + "foo bar")}
        _ = fi16
        _ = v20
    }
    v11 += (^(v11))
    var v21 int16 = -32768
    v11 = v11
    switch v14 {
    case "a":
        fmt.Println(v4, v4)
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = (v11 / (((v11 - v11)) | 1))
        }
    case "hello":
        v4 -= v4
    }
    v22 := GenS4{F0: (v4 + (v4 ^ (v4 * (v4 - v4))))}
    v11 = 2
    v23 := GenS2{F0: 67}
    var v24 int64 = -839
    delete(v3, (^(v11)))
    if (!false) {
        if (!(128 != 127)) {
            var v25 int = 1
            var v26 uint16 = 0
            _ = v25
            _ = v26
        } else {
            v21 += -32768
            for fi27 := 0; fi27 < 6; fi27++ {
                v28 := GenS1{F0: 2, F1: 674}
                var v29 uint8 = 42
                delete(v3, v11)
                _ = fi27
                _ = v28
                _ = v29
            }
            var v30 bool = (false || (true || (v13 < v13)))
            _ = v30
        }
    }
    _ = len(v3)
    if false {
        for fi31 := 0; fi31 < 20; fi31++ {
            switch (fi31 - fi31) {
            case -1:
                if true {
                    v14 = v14
                    v24 = 290
                } else {
                    for fi32 := 0; fi32 < 13; fi32++ {
                        var v33 int8 = -78
                        fmt.Println(v14)
                        _ = fi32
                        _ = v33
                    }
                    if ((v13 << (((v4 | 42)) & 63)) == v13) {
                        fmt.Println(v4, v24, fi31)
                        v2[v11] = v14
                        if false {
                            if true {
                                for fi34 := 0; fi34 < 17; fi34++ {
                                    if (float32((1.5 / (((((((0.1 + 1e10) / (3.14159 - (-2.25 * ((1e10 + (((float64(0.0) + 0.0) + 0.0) - 0.1)) - 1e10)))) * (3.14159 * 3.14159)) / 0.1) + 1e10) - 0.1) / -2.25))) < (0.0 * 1.5)) {
                                        v35 := GenS5{F0: v13, F1: 85}
                                        _ = v35
                                    }
                                    v36 := GenS1{F0: 363, F1: -478}
                                    v37 := GenS1{F0: 1, F1: 2147483647}
                                    _ = fi34
                                    _ = v36
                                    _ = v37
                                }
                                switch ("" + (v14 + ("foo bar" + "foo bar"))) {
                                case "":
                                    var v38 int16 = v21
                                    _ = v38
                                }
                            } else {
                                fmt.Println(v11)
                            }
                            v39 := GenS3{F0: "hello"}
                            delete(v2, (v11 ^ 953))
                            _ = v39
                        } else {
                            v14 = ("hello" + (v14 + (v14 + v14)))
                            delete(v2, (v11 << ((genFn0()) & 63)))
                        }
                    } else {
                        v4 = (v4 - (v4 << (((v4 << ((65535) & 63))) & 63)))
                    }
                }
            }
            v24 += v24
            _ = fi31
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
