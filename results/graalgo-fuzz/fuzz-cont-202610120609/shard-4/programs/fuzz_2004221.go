package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS6 struct { F0 string; F1 uint16; F2 int64 }
type GenS7 struct { F0 uint16; F1 uint32; F2 bool }
type GenS8 struct { F0 uint32; F1 uint8; F2 uint }
type GenS9 struct { F0 uint32; F1 int; F2 string }
func genFn0() int16 {
    v7 := GenS1{F0: 97, F1: -486, F2: 4294967295}
    if ((0.0 + 1e10) == 1.5) {
        var v8 int16 = -32768
        v9 := []int{}
        _ = v8
        _ = v9
    }
    if (1.5 == 1e10) {
        if true {
            var v10 int8 = 127
            var v11 int64 = -443
            v12 := GenS2{F0: 42, F1: -546}
            _ = v10
            _ = v11
            _ = v12
        } else {
            v13 := []int{}
            switch 32767 {
            case -1:
                var v14 bool = false
                _ = v14
            case 336:
                if (0 != 0) {
                    for fi15 := 0; fi15 < 20; fi15++ {
                        switch "foo bar" {
                        case "\n\t":
                            var v16 uint8 = 2
                            _ = v16
                        }
                        v17 := GenS3{F0: -32768}
                        for fi18 := 0; fi18 < 14; fi18++ {
                            v19 := GenS2{F0: 1, F1: 664}
                            var v20 string = "foo bar"
                            _ = len(v13)
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        _ = fi15
                        _ = v17
                    }
                    if (0 >= -128) {
                        var v21 int16 = -32768
                        v21 = v21
                        _ = v21
                    } else {
                        var v22 float32 = 0.0
                        _ = v22
                    }
                }
            case 42:
                for fi23 := 0; fi23 < 1; fi23++ {
                    if false {
                        switch (("\n\t" + "foo bar") + "") {
                        case "a":
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                        case "\n\t":
                            var v24 uint = 18446744073709551615
                            _ = v24
                        case "foo bar":
                            v25 := GenS4{F0: 65535}
                            _ = v25
                        default:
                            var v26 uint32 = 65535
                            _ = v26
                        }
                    }
                    fmt.Println(fi23)
                    if true {
                        var v27 uint32 = 0
                        var v28 float32 = (3.14159 - 3.14159)
                        v28 = float32(1e10)
                        _ = v27
                        _ = v28
                    }
                    _ = fi23
                }
            default:
                var v29 int = -1
                _ = v29
            }
            var v30 int16 = 905
            _ = v13
            _ = v30
        }
        for fi31 := 0; fi31 < 19; fi31++ {
            var v32 int16 = -1
            var v33 int = -443
            _ = fi31
            _ = v32
            _ = v33
        }
        for fi34 := 0; fi34 < 17; fi34++ {
            var v35 float64 = -2.25
            if false {
                var v36 int64 = -9223372036854775808
                switch ("" + "hello") {
                case "hello":
                    fmt.Println(v35, v36, v36)
                }
                _ = v36
            } else {
                fmt.Println(fi34, fi34)
            }
            v37 := []int{}
            _ = fi34
            _ = v35
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 6; fi38++ {
        v39 := GenS5{F0: 2147483647, F1: 229, F2: 274}
        v40 := []int{(fi38 >> ((0) & 63))}
        var v41 uint16 = 2
        _ = fi38
        _ = v39
        _ = v40
        _ = v41
    }
    _ = v7
    return -1
}

func main() {
    var v1 uint = 42
    var v2 int32 = -281
    var v3 int64 = 42
    var v4 int64 = (v3 ^ (v3 * v3))
    v5 := []int{641, 473}
    v6 := make(map[uint32]int)
    if true {
        v42 := GenS6{F0: "foo bar", F1: 0, F2: v4}
        v3 = v3
        v1 += v1
        _ = v42
    }
    v1 -= (v1 - (^(v1)))
    for fi43 := 0; fi43 < 19; fi43++ {
        _ = len(v5)
        _ = fi43
    }
    v1 = (v1 << ((v1) & 63))
    for fi44 := 0; fi44 < 5; fi44++ {
        fmt.Println(fi44)
        v4 = (v3 & (v3 * -93))
        _ = fi44
    }
    if ((v1 ^ 65535) != (v1 / ((v1) | 1))) {
        if false {
            v4 = (^(v4))
        } else {
            if false {
                if false {
                    var v45 int64 = (v3 % ((v3) | 1))
                    var v46 int64 = -1
                    var v47 int = 127
                    _ = v45
                    _ = v46
                    _ = v47
                } else {
                    var v48 bool = true
                    _ = v48
                }
            } else {
                v49 := GenS7{F0: 662, F1: 1, F2: false}
                if len(v5) > 0 {
                    v5[0] = 42
                }
                fmt.Println(v4)
                _ = v49
            }
        }
    } else {
        fmt.Println(v2, v1)
        v50 := GenS8{F0: 658, F1: 2, F2: 255}
        if (!false) {
            if false {
                v51 := GenS9{F0: 2, F1: -455, F2: ("hello" + "hello")}
                v3 += v4
                _ = v51
            }
        }
        _ = v50
    }
    var v52 int = -600
    var v53 int = (v52 + int(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
}
