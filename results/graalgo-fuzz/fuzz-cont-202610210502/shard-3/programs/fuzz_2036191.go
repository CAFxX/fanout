package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int64; F1 int32; F2 uint16 }
func genFn0(p0 int8, p1 int64) int8 {
    p0 = (p0 >> ((18446744073709551615) & 63))
    if false {
        var v6 int64 = (p1 / ((0) | 1))
        var v7 bool = true
        _ = v6
        _ = v7
    } else {
        if true {
            if false {
                switch (("foo bar" + "") + "hello") {
                case "foo bar":
                    var v8 string = (("hello" + "\n\t") + (("\n\t" + "") + ("\n\t" + ("a" + (("a" + "\n\t") + "foo bar")))))
                    _ = v8
                case "":
                    v9 := []int{348, 245}
                    _ = v9
                case "a":
                    p1 = (p1 ^ (p1 - 9223372036854775807))
                }
                var v10 uint16 = 723
                v11 := []int{}
                _ = v10
                _ = v11
            }
            if true {
                if false {
                    switch -2147483648 {
                    case 4294967295:
                        for fi12 := 0; fi12 < 7; fi12++ {
                            if (true || (!false)) {
                                for fi13 := 0; fi13 < 2; fi13++ {
                                    var v14 int64 = 1
                                    _ = fi13
                                    _ = v14
                                }
                                fmt.Println(fi12, p0)
                                p0 = -47
                            } else {
                                fmt.Println(fi12, p0)
                                for fi15 := 0; fi15 < 12; fi15++ {
                                    fmt.Println(fi12)
                                    p1 += -1
                                    _ = fi15
                                }
                                for fi16 := 0; fi16 < 19; fi16++ {
                                    v17 := GenS1{F0: (p1 >> ((2) & 63)), F1: -1, F2: -831}
                                    v18 := []int{}
                                    fmt.Println(fi12, p1)
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                }
                            }
                            p0 -= (p0 * -34)
                            v19 := GenS1{F0: (p1 / ((p1) | 1)), F1: 83, F2: 1}
                            _ = fi12
                            _ = v19
                        }
                    case 127:
                        p0 = (p0 + 42)
                    }
                } else {
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        var v20 int8 = 89
                        _ = v20
                    }
                }
                if false {
                    var v21 uint32 = 232
                    v21 = v21
                    _ = v21
                }
            } else {
                v22 := GenS2{F0: -748, F1: "hello"}
                _ = v22
            }
            p1 += (p1 - 0)
        }
        p0 = (^(p0))
    }
    var v23 uint64 = 242
    var v24 int32 = 42
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return p0
}

func main() {
    var v1 float32 = (-2.25 + (1e10 / 0.1))
    var v2 string = "hello"
    var v3 string = (v2 + (v2 + (v2 + v2)))
    var v4 float64 = 1.5
    v5 := make(map[uint32]bool)
    for fi25 := 0; fi25 < 2; fi25++ {
        v1 = (v1 / (v1 + (v1 + v1)))
        fmt.Println(fi25, v4)
        _ = fi25
    }
    var v26 int64 = 42
    if (false || (((!false) || ("a" >= v3)) && true)) {
        var v27 uint64 = 1
        for fi28 := 0; fi28 < 5; fi28++ {
            v29 := []int{}
            _ = fi28
            _ = v29
        }
        _ = v27
    } else {
        var v30 float32 = (0.0 / 1.0)
        _ = v30
    }
    if true {
        var v31 float32 = 0.0
        var v32 int8 = 100
        _ = v31
        _ = v32
    }
    v3 = v3
    var v33 int32 = 0
    v34 := GenS1{F0: -368, F1: -575, F2: -1}
    v35 := v5[994]
    var v36 float64 = ((0.1 - 0.0) - (0.1 / v4))
    v37 := GenS3{F0: (v26 | v26), F1: 664, F2: 32767}
    var v38 int8 = 22
    v26 -= (^(v26))
    if v35 {
        v38 -= (v38 / (((v38 ^ (v38 % (((^(v38))) | 1)))) | 1))
        v4 = (((0.1 - (-2.25 + (v4 / v36))) + float64(v1)) / 1e10)
    } else {
        v2 = "hello"
    }
    v5[555] = true
    if (1 <= 1) {
        if v35 {
            v39 := []int{-2147483648}
            v35 = false
            _ = v39
        } else {
            if (2147483648 < 65535) {
                switch (v3 + v3) {
                case "\n\t":
                    fmt.Println(v38, v38)
                case "foo bar":
                    var v40 uint8 = 255
                    _ = v40
                default:
                    var v41 int16 = 315
                    _ = v41
                }
                v42 := []int{}
                _ = v42
            }
            v3 = ((((("" + "hello") + "hello") + v2) + v3) + v2)
        }
        var v43 uint16 = 467
        _ = v43
    }
    var v44 int8 = (v38 & 1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v44
    _ = v5
}
