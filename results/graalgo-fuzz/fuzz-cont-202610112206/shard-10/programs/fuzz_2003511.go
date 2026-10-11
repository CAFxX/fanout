package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int32 }
func genFn0() uint {
    v5 := GenS1{F0: 6}
    var v6 float32 = 1e10
    v7 := GenS1{F0: 4294967295}
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1() int {
    v8 := []int{}
    var v9 uint16 = 2
    _ = v8
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 int32 = 1
    v2 := make(map[int]int)
    var v3 uint8 = 1
    var v4 int32 = 0
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(fi10)
        v1 = 42
        _ = fi10
    }
    v1 -= v4
    v11 := []int{332, 128}
    fmt.Println(v3, v3)
    fmt.Println(v4, v3)
    v4 += (v4 % ((v4) | 1))
    var v12 int64 = -111
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            if true {
                v15 := GenS1{F0: 4294967295}
                fmt.Println(v4, fi13, v1)
                switch (fi13 / (((fi14 % (((fi13 & (-(fi14)))) | 1))) | 1)) {
                case 9223372036854775807:
                    if (true && ((fi13 >> ((2) & 63)) > fi13)) {
                        switch "foo bar" {
                        case "hello":
                            var v16 float64 = 1e10
                            _ = v16
                        case "a":
                            switch fi13 {
                            case -733:
                                v3 = 148
                            }
                        case "foo bar":
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        }
                    }
                case 999:
                    v3 -= (v3 + 127)
                default:
                    var v17 int32 = v1
                    _ = v17
                }
                _ = v15
            }
            _ = fi14
        }
        for fi18 := 0; fi18 < 16; fi18++ {
            if len(v11) > 0 {
                v11[0] = fi13
            }
            fmt.Println(v3)
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi18
        }
        _ = fi13
    }
    if (!(true || (32767 == uint16(v3)))) {
        v4 = 0
        if true {
            var v19 int8 = -101
            _ = v19
        }
        _ = len(v11)
    } else {
        if false {
            for fi20 := 0; fi20 < 16; fi20++ {
                if false {
                    if ("a" <= ("hello" + (("a" + "foo bar") + ""))) {
                        v1 += v1
                        if (65535 == 1) {
                            for fi21 := 0; fi21 < 17; fi21++ {
                                if true {
                                    if (2147483648 < 0) {
                                        var v22 int16 = 300
                                        v12 += v12
                                        _ = v22
                                    } else {
                                        v12 = (v12 / ((v12) | 1))
                                        var v23 uint16 = 42
                                        _ = v23
                                    }
                                    fmt.Println(fi20)
                                    if (false && true) {
                                        v12 = 0
                                        v12 = v12
                                        var v24 uint8 = 0
                                        _ = v24
                                    }
                                } else {
                                    v1 = v1
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    v3 -= v3
                                }
                                _ = fi21
                            }
                            v4 = (v1 % ((v1) | 1))
                            v25 := GenS1{F0: 4294967295}
                            _ = v25
                        }
                        fmt.Println(v3, fi20)
                    }
                }
                _ = fi20
            }
            var v26 float32 = 0.0
            _ = v26
        }
        _ = len(v11)
        fmt.Println(v12)
    }
    v27 := GenS2{F0: 100, F1: v1}
    v4 -= v1
    var v28 bool = false
    for fi29 := 0; fi29 < 14; fi29++ {
        for fi30 := 0; fi30 < 8; fi30++ {
            if (false && v28) {
                if ((v28 && v28) || true) {
                    var v31 int32 = 1
                    _ = v31
                }
                if len(v11) > 0 {
                    v11[0] = -1
                }
            } else {
                fmt.Println(fi30, v4, v3)
                v32 := GenS2{F0: 0, F1: v4}
                v28 = (genFn0() == genFn0())
                _ = v32
            }
            _ = fi30
        }
        var v33 int32 = (v1 & v4)
        _ = fi29
        _ = v33
    }
    fmt.Println(v4, v28)
    for fi34 := 0; fi34 < 18; fi34++ {
        switch (-(fi34)) {
        case 2147483647:
            var v35 uint8 = v3
            _ = v35
        case -1:
            var v36 string = ""
            _ = v36
        case 4294967295:
            var v37 uint64 = 9
            _ = v37
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
