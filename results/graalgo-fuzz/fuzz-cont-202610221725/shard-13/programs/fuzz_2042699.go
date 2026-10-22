package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int16; F1 string; F2 int64 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint64; F1 uint64; F2 int32 }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[uint32]string)
    var v3 uint = 0
    v3 = 65535
    v4 := GenS1{F0: -577, F1: ("foo bar" + "a")}
    switch -1 {
    case 32767:
        for fi5 := 0; fi5 < 17; fi5++ {
            var v6 int32 = 0
            _ = fi5
            _ = v6
        }
    case 32768:
        if (((-2.25 + 1e10) + (0.1 + -2.25)) != float64((float32(3.14159) - v1))) {
            var v7 uint32 = 4294967295
            switch 106 {
            case 0:
                var v8 float32 = v1
                _ = v8
            case 222:
                var v9 int32 = 42
                _ = v9
            default:
                if false {
                    var v10 int8 = 0
                    switch "\n\t" {
                    case "hello":
                        v11 := v2[2]
                        _ = v11
                    case "\n\t":
                        v3 -= v3
                    case "":
                        v12 := GenS1{F0: -32768, F1: (("foo bar" + ("hello" + ("hello" + ("foo bar" + "hello")))) + (("foo bar" + "\n\t") + "a"))}
                        _ = v12
                    }
                    v7 += (v7 - (v7 + (v7 / ((uint32(v3)) | 1))))
                    _ = v10
                }
            }
            _ = v7
        }
    }
    switch "hello" {
    case "foo bar":
        for fi13 := 0; fi13 < 14; fi13++ {
            if false {
                v14 := GenS1{F0: -1, F1: "hello"}
                _ = v14
            } else {
                if (100 != 100) {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        var v16 uint = (v3 << ((18446744073709551615) & 63))
                        v17 := GenS1{F0: -32768, F1: "hello"}
                        v1 = float32(((1e10 / (1.5 / (-2.25 - 0.1))) * 3.14159))
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    switch "" {
                    case "hello":
                        for fi18 := 0; fi18 < 11; fi18++ {
                            v1 = (v1 / v1)
                            var v19 uint16 = 32767
                            if true {
                                var v20 float64 = float64(float32(((0.0 / 1.0) - -2.25)))
                                v1 -= v1
                                _ = v20
                            } else {
                                var v21 uint = v3
                                var v22 uint = (v3 - v3)
                                _ = v21
                                _ = v22
                            }
                            _ = fi18
                            _ = v19
                        }
                    case "a":
                        var v23 bool = (false && ((v3 % (((v3 << ((v3) & 63))) | 1)) < (v3 & v3)))
                        _ = v23
                    case "":
                        fmt.Println(v3, fi13, v1)
                    default:
                        delete(v2, 42)
                    }
                }
                v3 -= v3
            }
            _ = fi13
        }
    }
    var v24 bool = false
    if (87 != 127) {
        fmt.Println(v24)
        v25 := GenS1{F0: 0, F1: (("a" + (("" + "foo bar") + "a")) + (("hello" + "hello") + (("\n\t" + "") + ("foo bar" + "a"))))}
        var v26 uint = 4294967295
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: -32768, F1: ("" + "a")}
    v28 := GenS1{F0: -1, F1: "hello"}
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 uint8 = 127
        for fi31 := 0; fi31 < 9; fi31++ {
            v24 = ((fi31 / ((fi31) | 1)) == 9223372036854775807)
            v32 := []int{(fi29 * fi31), fi29, (fi31 >> ((v3) & 63))}
            v32 = append(v32, (fi31 - -765))
            _ = fi31
            _ = v32
        }
        v30 = (v30 % ((v30) | 1))
        _ = fi29
        _ = v30
    }
    v33 := GenS1{F0: 0, F1: "hello"}
    v34 := GenS1{F0: 1, F1: ("foo bar" + "a")}
    var v35 string = ("hello" + ("\n\t" + "\n\t"))
    for fi36 := 0; fi36 < 15; fi36++ {
        if v24 {
            v37 := GenS2{F0: 1000, F1: "\n\t", F2: 0}
            for fi38 := 0; fi38 < 15; fi38++ {
                var v39 float64 = -2.25
                if v24 {
                    var v40 int16 = -32768
                    _ = v40
                } else {
                    fmt.Println(fi36, fi36, fi36)
                    v41 := GenS3{F0: 157, F1: ((v1 + v1) != 1e10)}
                    v35 = "\n\t"
                    _ = v41
                }
                _ = fi38
                _ = v39
            }
            _ = v37
        }
        v42 := []int{}
        var v43 uint64 = 0
        _ = fi36
        _ = v42
        _ = v43
    }
    if (!false) {
        v24 = (!(4294967295 == 0))
        var v44 int16 = -732
        fmt.Println(v1, v24, v44)
        _ = v44
    }
    if v24 {
        v45 := GenS4{F0: 4294967295, F1: 466, F2: 2147483647}
        var v46 uint32 = 2147483647
        _ = v45
        _ = v46
    } else {
        var v47 bool = (v1 < v1)
        if (42 == 255) {
            v24 = v24
        }
        _ = v47
    }
    var v48 int32 = 1
    for fi49 := 0; fi49 < 7; fi49++ {
        var v50 string = (v35 + ((v35 + ((("a" + v35) + v35) + v35)) + v35))
        v48 = -1
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v48
}
