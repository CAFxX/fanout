package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
func genFn0() int64 {
    var v4 uint8 = 127
    v4 = 0
    var v5 int32 = 97
    _ = v4
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 int = 128
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 127
    v6 := v2[586]
    delete(v2, (^(v6)))
    fmt.Println(v6)
    delete(v2, v6)
    for fi7 := 0; fi7 < 8; fi7++ {
        if (!(("" + "") >= "")) {
            if false {
                fmt.Println(v6, fi7, v6)
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    if (!false) {
                        if (1e10 != 0.1) {
                            fmt.Println(fi7, v1, fi7)
                            for fi8 := 0; fi8 < 19; fi8++ {
                                if ((!((!false) || (42 > 1))) || false) {
                                    v9 := GenS1{F0: -119, F1: ("a" + "a")}
                                    _ = v9
                                } else {
                                    v3 = v3
                                    var v10 bool = (0.0 != (-2.25 / 3.14159))
                                    _ = v10
                                }
                                v6 -= v6
                                v3 = v3
                                _ = fi8
                            }
                        }
                        v11 := GenS2{F0: 0}
                        var v12 uint = 255
                        _ = v11
                        _ = v12
                    } else {
                        var v13 bool = ((!false) && (!(!(123 > fi7))))
                        fmt.Println(v3)
                        _ = v13
                    }
                case "a":
                    var v14 int8 = 100
                    _ = v14
                default:
                    var v15 int64 = 42
                    _ = v15
                }
            }
        }
        v16 := []int{-639, fi7}
        if true {
            if false {
                fmt.Println(v3, v3, v6)
            } else {
                switch v1 {
                case 9223372036854775807:
                    var v17 uint16 = 0
                    _ = v17
                case 2147483647:
                    var v18 string = "foo bar"
                    _ = v18
                default:
                    var v19 int64 = 211
                    _ = v19
                }
                _ = len(v16)
            }
        }
        _ = fi7
        _ = v16
    }
    var v20 int16 = 963
    switch ("foo bar" + "foo bar") {
    case "hello":
        v6 = 4294967295
    case "foo bar":
        v6 += 361
    case "a":
        if true {
            fmt.Println(v20, v1, v6)
            var v21 uint16 = 65535
            _ = v21
        }
    }
    var v22 float64 = 1.5
    var v23 int32 = -1
    var v24 uint64 = 4294967295
    if ((-142 < (^(v1))) || (!true)) {
        for fi25 := 0; fi25 < 2; fi25++ {
            v2[v6] = (v6 - 0)
            var v26 float32 = 0.1
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 8; fi27++ {
            v3 = v3
            v3 += v3
            if true {
                var v28 int32 = -1
                v24 -= 1
                switch ("a" + "a") {
                case "foo bar":
                    for fi29 := 0; fi29 < 9; fi29++ {
                        v6 += v6
                        _ = fi29
                    }
                default:
                    v30 := GenS1{F0: -1, F1: (("\n\t" + (("foo bar" + ("foo bar" + "a")) + "\n\t")) + "foo bar")}
                    _ = v30
                }
                _ = v28
            }
            _ = fi27
        }
        if false {
            if ((v1 / ((v1) | 1)) <= 127) {
                _ = len(v2)
            } else {
                for fi31 := 0; fi31 < 7; fi31++ {
                    fmt.Println(v3)
                    v32 := GenS3{F0: v23}
                    _ = fi31
                    _ = v32
                }
                v24 -= (v24 | (v24 / ((v24) | 1)))
            }
        }
    } else {
        v33 := GenS2{F0: v23}
        if (((v6 % (((v6 >> ((1) & 63))) | 1)) < (v6 % (((v6 * v6)) | 1))) && false) {
            v3 -= 255
            v22 -= (v22 + v22)
        } else {
            if false {
                var v34 int8 = int8(v1)
                _ = v34
            } else {
                delete(v2, 0)
            }
        }
        _ = v33
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        delete(v2, v6)
        var v36 uint16 = 0
        var v37 uint16 = v36
        _ = fi35
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        var v39 int64 = 9223372036854775807
        var v40 uint8 = uint8(v6)
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v6
}
