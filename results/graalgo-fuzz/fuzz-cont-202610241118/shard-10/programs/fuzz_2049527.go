package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint64 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint32; F1 int16; F2 int32 }
func genFn0() string {
    var v4 uint32 = 2147483648
    v5 := GenS1{F0: true, F1: 42, F2: 0}
    var v6 uint16 = 136
    if true {
        if false {
            switch 128 {
            case -631:
                var v7 bool = (!true)
                _ = v7
            case 0:
                v4 = v4
            }
            v4 -= v4
            v8 := []int{-2147483648, 389}
            _ = v8
        } else {
            switch 427 {
            case 2147483647:
                var v9 int = 32767
                _ = v9
            case 127:
                if true {
                    var v10 int64 = 42
                    _ = v10
                } else {
                    fmt.Println(v4, v4)
                    v6 = v6
                    v6 = 42
                }
            default:
                var v11 uint64 = 804
                _ = v11
            }
            fmt.Println(v6)
            switch "" {
            case "a":
                if false {
                    var v12 float64 = 1e10
                    _ = v12
                } else {
                    var v13 int8 = 42
                    _ = v13
                }
            case "":
                if (42 > 127) {
                    v14 := []int{127}
                    var v15 bool = true
                    _ = v14
                    _ = v15
                } else {
                    var v16 int32 = 2147483647
                    _ = v16
                }
            case "foo bar":
                fmt.Println(v6, v4)
            default:
                switch 1 {
                case 42:
                    if true {
                        v17 := GenS1{F0: true, F1: 42, F2: 42}
                        var v18 float64 = 0.0
                        if true {
                            if ((-2.25 >= v18) || (!(-32768 > 1000))) {
                                v19 := []int{-9223372036854775808, 32767}
                                _ = v19
                            }
                            if (true || false) {
                                var v20 int16 = 42
                                v6 = v6
                                _ = v20
                            }
                            v21 := []int{2147483647, 128, -476}
                            _ = v21
                        } else {
                            if true {
                                for fi22 := 0; fi22 < 20; fi22++ {
                                    fmt.Println(v4, v6)
                                    _ = fi22
                                }
                                v23 := []int{127, -1, 42}
                                _ = v23
                            }
                        }
                        _ = v17
                        _ = v18
                    } else {
                        for fi24 := 0; fi24 < 3; fi24++ {
                            v25 := GenS1{F0: false, F1: 32767, F2: 109}
                            for fi26 := 0; fi26 < 12; fi26++ {
                                v27 := GenS2{F0: -9223372036854775808, F1: true}
                                v4 -= v4
                                var v28 uint = uint(v4)
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            v6 += v6
                            _ = fi24
                            _ = v25
                        }
                        var v29 uint = 255
                        for fi30 := 0; fi30 < 14; fi30++ {
                            v29 = 42
                            _ = fi30
                        }
                        _ = v29
                    }
                case 1:
                    if true {
                        var v31 bool = (true || ((v6 << ((65535) & 63)) == v6))
                        v6 -= (v6 & v6)
                        _ = v31
                    }
                case 32768:
                    v32 := GenS2{F0: int64(v6), F1: true}
                    _ = v32
                }
            }
        }
    }
    _ = v4
    _ = v5
    _ = v6
    return "a"
}

func main() {
    v1 := []uint32{296, 396}
    var v2 int64 = 1
    var v3 int64 = v2
    if true {
        var v33 int = 681
        _ = len(v1)
        switch v33 {
        case -518:
            v34 := GenS2{F0: (v2 << ((42) & 63)), F1: (false || true)}
            _ = v34
        case -975:
            if (-128 == -1) {
                v35 := GenS1{F0: true, F1: 1, F2: 1}
                v33 = 2147483647
                _ = len(v1)
                _ = v35
            } else {
                v33 += (-(v33))
                v33 = (v33 & v33)
            }
        }
        _ = v33
    }
    v36 := GenS3{F0: 351, F1: 1000, F2: -2147483648}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v37 := GenS3{F0: 994, F1: -259, F2: -2147483648}
    if true {
        var v38 int64 = (v3 >> ((65535) & 63))
        var v39 uint16 = 32767
        for fi40 := 0; fi40 < 15; fi40++ {
            v38 -= v38
            v39 -= 89
            fmt.Println(v2, fi40, fi40)
            _ = fi40
        }
        _ = v38
        _ = v39
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        if (v3 <= -1) {
            v3 = (^(v2))
            v2 += v2
            var v42 int = (fi41 % ((fi41) | 1))
            _ = v42
        }
        for fi43 := 0; fi43 < 10; fi43++ {
            switch 128 {
            case 423:
                var v44 uint8 = 7
                _ = v44
            case -193:
                fmt.Println(fi41, v3, fi41)
            case 128:
                v3 -= v2
            }
            _ = fi43
        }
        _ = fi41
    }
    v3 = v3
    var v45 string = ((("hello" + ("foo bar" + ("\n\t" + "foo bar"))) + ("hello" + ("foo bar" + (("\n\t" + "hello") + "hello")))) + ((("" + ((("a" + "hello") + ("hello" + ("hello" + ""))) + (((("\n\t" + ((("" + "\n\t") + "\n\t") + ("a" + (((("" + (("foo bar" + (("\n\t" + ("hello" + "")) + ("foo bar" + ("" + ((("foo bar" + ("" + ("a" + ""))) + "") + ""))))) + "")) + "hello") + ("hello" + "hello")) + ((("" + ("" + "hello")) + "") + "hello"))))) + "") + ("hello" + ("a" + ("" + "foo bar")))) + (("hello" + "foo bar") + "hello")))) + "hello") + ("a" + "hello")))
    var v46 uint32 = 42
    fmt.Println(v45, v46, v2)
    if (!true) {
        v2 += (v2 / ((-581) | 1))
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v45 = v45
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v45
    _ = v46
}
