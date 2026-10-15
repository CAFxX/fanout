package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS2 struct { F0 int16; F1 bool; F2 int8 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int16; F1 uint64 }
func genFn0() int64 {
    var v7 uint = 255
    var v8 int8 = 123
    var v9 bool = (32767 <= 0)
    v10 := GenS1{F0: 4294967295, F1: 42, F2: -1}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]string)
    var v3 float32 = (0.1 / 0.1)
    var v4 int16 = -116
    var v5 uint16 = 850
    var v6 bool = ((true || false) && false)
    v11 := []int{-843}
    for fi12 := 0; fi12 < 14; fi12++ {
        delete(v2, 4294967295)
        v11 = append(v11, 9223372036854775807)
        delete(v2, 45)
        _ = fi12
    }
    v13 := GenS1{F0: 2, F1: 169, F2: v4}
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 int32 = -2147483648
        v16 := GenS1{F0: 0, F1: 0, F2: (v4 & (v4 << ((771) & 63)))}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint16 = (v5 ^ v5)
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int = fi18
        v3 -= float32(-2.25)
        var v20 uint16 = (v17 + (^(v5)))
        _ = fi18
        _ = v19
        _ = v20
    }
    switch "\n\t" {
    case "hello":
        var v21 int16 = (v4 + v4)
        _ = v21
    case "foo bar":
        var v22 int8 = (v1 | -125)
        _ = v22
    }
    v23 := GenS1{F0: 4294967295, F1: 139, F2: (v4 * v4)}
    v24 := GenS1{F0: 18446744073709551615, F1: 255, F2: -132}
    v17 -= v5
    fmt.Println(v17)
    var v25 uint64 = uint64(v5)
    switch "" {
    case "hello":
        v26 := GenS1{F0: 1, F1: 149, F2: 0}
        _ = v26
    case "":
        if v6 {
            if false {
                v6 = v6
                delete(v2, 4294967295)
            } else {
                var v27 bool = v6
                _ = v27
            }
            var v28 bool = (!v6)
            _ = v28
        }
    default:
        v29 := GenS2{F0: v4, F1: true, F2: v1}
        _ = v29
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        if true {
            v6 = (!v6)
            var v31 string = "a"
            _ = v31
        }
        _ = fi30
    }
    var v32 int = 646
    if false {
        if false {
            v33 := GenS3{F0: v25, F1: 150}
            v3 += float32((((float64(1.5) * 1.5) + (1e10 * 3.14159)) - 3.14159))
            for fi34 := 0; fi34 < 2; fi34++ {
                var v35 bool = ((1 < 393) && (-1 != 42))
                v1 = v1
                _ = fi34
                _ = v35
            }
            _ = v33
        } else {
            if true {
                if v6 {
                    switch ("foo bar" + "\n\t") {
                    case "hello":
                        var v36 float32 = ((float32(1e10) * 0.1) / float32(0.1))
                        _ = v36
                    case "foo bar":
                        switch "\n\t" {
                        case "\n\t":
                            var v37 int = 2147483647
                            _ = v37
                        case "hello":
                            for fi38 := 0; fi38 < 4; fi38++ {
                                v1 -= 64
                                fmt.Println(v5, v6)
                                switch ("\n\t" + ("\n\t" + ("a" + ""))) {
                                case "foo bar":
                                    v3 += (v3 / 1.0)
                                case "":
                                    v32 += fi38
                                case "a":
                                    v1 -= (v1 | v1)
                                }
                                _ = fi38
                            }
                        }
                    default:
                        v11 = append(v11, (v32 ^ 9223372036854775807))
                    }
                    if v6 {
                        switch (v32 << ((1) & 63)) {
                        case 32767:
                            var v39 int16 = 1
                            _ = v39
                        case -1:
                            fmt.Println(v6, v25, v6)
                        }
                        var v40 float32 = (1.5 + (-2.25 + -2.25))
                        if v6 {
                            v32 = (^(v32))
                            var v41 int32 = 0
                            var v42 uint = 2
                            _ = v41
                            _ = v42
                        } else {
                            if (v6 || ("hello" == "foo bar")) {
                                fmt.Println(v25, v25)
                            }
                            var v43 int16 = (v4 ^ (^(v4)))
                            _ = v43
                        }
                        _ = v40
                    }
                }
                fmt.Println(v4, v17)
                v11 = append(v11, v32)
            } else {
                fmt.Println(v32)
                v11 = append(v11, (v32 * 32767))
            }
            v3 += (float32(-2.25) / 1.0)
            if (-1 <= v1) {
                _ = len(v11)
                v17 -= (v5 * (v17 + (v17 << ((634) & 63))))
            } else {
                for fi44 := 0; fi44 < 16; fi44++ {
                    v45 := v2[1]
                    var v46 float32 = (v3 - 1e10)
                    v25 = v25
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                v47 := GenS4{F0: v4, F1: (v25 - v25)}
                if len(v11) > 0 {
                    v11[0] = v32
                }
                _ = v47
            }
        }
        v17 += (v17 >> ((uint(v17)) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
