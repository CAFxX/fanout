package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint, p2 uint32) bool {
    p2 -= (^(p2))
    var v6 int32 = p0
    if ((v6 % (((p0 / (((^(p0))) | 1))) | 1)) != v6) {
        var v7 int32 = (v6 & 0)
        p0 = -1
        _ = v7
    }
    v8 := []int{127, 4294967295, 878, 650}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return (-13 > 100)
}
func genFn1() int16 {
    if false {
        for fi9 := 0; fi9 < 14; fi9++ {
            fmt.Println(fi9)
            v10 := GenS1{F0: "", F1: false}
            var v11 int8 = -1
            _ = fi9
            _ = v10
            _ = v11
        }
        if true {
            for fi12 := 0; fi12 < 15; fi12++ {
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 uint8 = 42
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
            v15 := GenS1{F0: "hello", F1: false}
            switch "\n\t" {
            case "foo bar":
                var v16 int = 128
                _ = v16
            case "":
                if false {
                    var v17 uint16 = 255
                    _ = v17
                }
            }
            _ = v15
        }
    } else {
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS2{F0: 1, F1: 1}
            _ = fi18
            _ = v19
        }
    }
    if (-650 > -971) {
        var v20 bool = (-2.25 >= (-2.25 + 1.5))
        fmt.Println(v20)
        _ = v20
    }
    return 670
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint = 255
    var v3 uint = v2
    var v4 uint16 = 2
    v5 := make(map[string]bool)
    var v21 int32 = 807
    v22 := []int{-600, 10, 127}
    if (((true || (-724 > 441)) && (!false)) || false) {
        for fi23 := 0; fi23 < 15; fi23++ {
            fmt.Println(fi23)
            if (65535 > 2147483648) {
                if true {
                    if false {
                        fmt.Println(v3)
                    }
                    switch int(v4) {
                    case -2147483648:
                        var v24 string = ("foo bar" + ("\n\t" + "foo bar"))
                        _ = v24
                    case -495:
                        switch -923 {
                        case 1:
                            for fi25 := 0; fi25 < 16; fi25++ {
                                switch "hello" {
                                case "hello":
                                    v21 = (v21 >> (((v2 / (((v2 << ((v3) & 63))) | 1))) & 63))
                                case "\n\t":
                                    var v26 uint8 = 255
                                    _ = v26
                                }
                                _ = fi25
                            }
                        default:
                            v4 = (^(v4))
                        }
                    default:
                        var v27 uint8 = 128
                        _ = v27
                    }
                    v1 += 3.14159
                }
                v3 += (v3 ^ 255)
            } else {
                if false {
                    var v28 uint = (v2 | v2)
                    v28 += (v3 + v28)
                    v2 += (v3 ^ v2)
                    _ = v28
                }
                v1 -= -2.25
            }
            switch "a" {
            case "a":
                for fi29 := 0; fi29 < 19; fi29++ {
                    if true {
                        var v30 int64 = -622
                        _ = v30
                    }
                    _ = fi29
                }
            case "foo bar":
                if false {
                    var v31 uint32 = 1
                    _ = v31
                } else {
                    v32 := v5["hello"]
                    switch "" {
                    case "foo bar":
                        if (1 < -362) {
                            v33 := GenS3{F0: -1}
                            _ = v33
                        } else {
                            if v32 {
                                v5[""] = (v32 && true)
                                if (!(!("\n\t" <= ("hello" + "hello")))) {
                                    v34 := GenS4{F0: (v4 ^ (v4 / ((32767) | 1))), F1: (^(v4))}
                                    var v35 string = "hello"
                                    _ = v34
                                    _ = v35
                                }
                            }
                        }
                    }
                    _ = v32
                }
            default:
                v2 += v2
            }
            _ = fi23
        }
        _ = len(v5)
        var v36 int16 = -182
        _ = v36
    }
    v4 = 32767
    var v37 uint32 = 385
    fmt.Println(v37, v21)
    fmt.Println(v2, v21, v21)
    delete(v5, "a")
    v38 := GenS5{F0: uint8(v37)}
    if (!false) {
        fmt.Println(v21, v2)
        var v39 uint64 = 42
        _ = v39
    }
    v40 := GenS1{F0: "\n\t", F1: false}
    var v41 string = "\n\t"
    v4 += (v4 | v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
