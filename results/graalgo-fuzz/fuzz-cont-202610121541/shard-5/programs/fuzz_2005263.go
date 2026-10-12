package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint }
type GenS2 struct { F0 int64; F1 uint16; F2 int64 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int16; F1 bool; F2 uint32 }
func genFn0() uint16 {
    if (-2147483648 <= 313) {
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 int32 = 42
            _ = fi6
            _ = v7
        }
        var v8 string = (("\n\t" + "a") + "\n\t")
        if (636 == 151) {
            fmt.Println(v8)
            fmt.Println(v8)
        } else {
            var v9 float32 = 1e10
            for fi10 := 0; fi10 < 6; fi10++ {
                for fi11 := 0; fi11 < 4; fi11++ {
                    v12 := GenS1{F0: (32767 != 781), F1: false, F2: 0}
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
            _ = v9
        }
        _ = v8
    }
    var v13 float32 = (1.5 / 1.0)
    v14 := []int{42, 4294967295, 171, 9223372036854775807}
    _ = v13
    _ = v14
    return 32767
}

func main() {
    var v1 bool = (18446744073709551615 >= 2)
    var v2 uint32 = 781
    v3 := make(map[uint32]int)
    var v4 bool = false
    var v5 uint = 65535
    v15 := []int{}
    v1 = (42 == 42)
    if (((4294967295 <= 42) || v1) || false) {
        v4 = (false || v4)
        fmt.Println(v2)
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 string = ("a" + "hello")
            var v18 uint = 255
            switch "" {
            case "hello":
                var v19 float64 = ((0.0 + (0.0 * -2.25)) - 0.1)
                _ = v19
            case "\n\t":
                if (!v4) {
                    v4 = (float32(0.1) < (-2.25 / float32(1.5)))
                } else {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        v17 = v17
                        _ = fi20
                    }
                    var v21 float32 = 0.1
                    _ = v21
                }
            case "foo bar":
                var v22 int = fi16
                _ = v22
            default:
                v4 = (false && v1)
            }
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        var v23 string = "foo bar"
        v23 = (v23 + "hello")
        _ = v23
    }
    v1 = v1
    if ((false || (v5 < v5)) || (221 != 17)) {
        if v1 {
            v2 = (v2 ^ 2147483648)
            if (255 >= 2) {
                switch (("" + "\n\t") + "a") {
                case "":
                    switch "\n\t" {
                    case "foo bar":
                        if v4 {
                            v24 := v3[(v2 % (((^(v2))) | 1))]
                            v1 = (341 >= 65535)
                            _ = v24
                        }
                    case "":
                        switch ("\n\t" + ("a" + ("" + ("hello" + "a")))) {
                        case "foo bar":
                            if v1 {
                                if false {
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                } else {
                                    delete(v3, 42)
                                    var v25 string = (("\n\t" + "foo bar") + "\n\t")
                                    var v26 uint32 = 2
                                    _ = v25
                                    _ = v26
                                }
                            }
                        case "hello":
                            for fi27 := 0; fi27 < 13; fi27++ {
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                var v28 int16 = -32768
                                _ = fi27
                                _ = v28
                            }
                        case "":
                            var v29 uint64 = 1
                            _ = v29
                        }
                    default:
                        if v1 {
                            var v30 uint32 = 1
                            _ = v30
                        }
                    }
                case "\n\t":
                    fmt.Println(v1, v1)
                }
                switch 128 {
                case 0:
                    v31 := GenS2{F0: 1, F1: genFn0(), F2: 42}
                    _ = v31
                }
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
            } else {
                var v32 float32 = (((-2.25 / -2.25) / 0.1) * 1.5)
                v5 -= (v5 << (((^(v5))) & 63))
                _ = v32
            }
        }
        fmt.Println(v2, v1, v5)
    }
    switch ("hello" + ("foo bar" + "")) {
    case "a":
        v33 := GenS3{F0: 9223372036854775807, F1: 9223372036854775807}
        _ = v33
    case "":
        v34 := GenS1{F0: v4, F1: v4, F2: v5}
        _ = v34
    default:
        v5 = v5
    }
    switch 127 {
    case -367:
        var v35 int8 = 0
        _ = v35
    case 32767:
        v36 := GenS4{F0: 32767, F1: false, F2: (^(v2))}
        _ = v36
    case 127:
        var v37 bool = v1
        _ = v37
    }
    v2 = 800
    if true {
        for fi38 := 0; fi38 < 1; fi38++ {
            v15 = append(v15, fi38)
            v1 = ((2147483647 != -66) && (287 <= 187))
            var v39 uint64 = 42
            _ = fi38
            _ = v39
        }
        v40 := GenS3{F0: -708, F1: -1}
        _ = v40
    }
    v15 = append(v15, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
