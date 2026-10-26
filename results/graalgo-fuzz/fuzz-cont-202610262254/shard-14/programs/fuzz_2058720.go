package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 bool; F2 int32 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 string; F1 int; F2 uint64 }
func genFn0(p0 bool, p1 int16) uint {
    if false {
        var v5 uint64 = 718
        var v6 int32 = 588
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: "foo bar"}
        v8 := []int{4294967295, 127, 1, 387}
        _ = v7
        _ = v8
    }
    p0 = (false || ((0 != 128) && (32768 < 128)))
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    v1 := []int{2147483647}
    v2 := make(map[string]bool)
    var v3 uint16 = 0
    var v4 int = 17
    var v9 int32 = 2147483647
    var v10 string = "a"
    switch 9223372036854775807 {
    case 32768:
        v11 := GenS2{F0: 127}
        _ = v11
    }
    v10 = ("" + "a")
    v12 := GenS1{F0: v10}
    var v13 int64 = -1
    var v14 string = v10
    if (v3 == 255) {
        if true {
            v13 = v13
            switch v4 {
            case -2147483648:
                v13 += 923
            case -401:
                fmt.Println(v13, v3, v4)
            case 128:
                var v15 int16 = -628
                _ = v15
            default:
                v16 := GenS3{F0: 971, F1: (v9 >= (v9 + (v9 % ((-2147483648) | 1)))), F2: 42}
                _ = v16
            }
        }
    } else {
        delete(v2, (v10 + ""))
        v13 -= (v13 - int64(v3))
        switch "foo bar" {
        case "hello":
            v3 = v3
        case "a":
            for fi17 := 0; fi17 < 1; fi17++ {
                switch v4 {
                case 4294967295:
                    var v18 uint64 = 4294967295
                    _ = v18
                default:
                    if false {
                        if true {
                            var v19 uint8 = 255
                            for fi20 := 0; fi20 < 7; fi20++ {
                                var v21 uint8 = v19
                                var v22 float64 = 0.0
                                v3 -= v3
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                            _ = v19
                        }
                    }
                }
                _ = fi17
            }
        }
    }
    if ("\n\t" <= ("" + (((v14 + "foo bar") + (v14 + v14)) + "hello"))) {
        v14 = v14
        switch int(v9) {
        case 32767:
            for fi23 := 0; fi23 < 14; fi23++ {
                switch ("a" + "a") {
                case "a":
                    var v24 string = (v14 + v10)
                    _ = v24
                default:
                    v3 += v3
                }
                _ = fi23
            }
        case 128:
            v13 = 9223372036854775807
        }
    } else {
        var v25 uint32 = 0
        v25 = v25
        _ = v25
    }
    _ = len(v2)
    if (915 > 2) {
        var v26 int16 = -32768
        var v27 float32 = float32(1e10)
        _ = v26
        _ = v27
    }
    if ((0.0 * (-2.25 + 1e10)) == float32(0.0)) {
        for fi28 := 0; fi28 < 5; fi28++ {
            var v29 uint16 = (v3 << ((genFn0(false, 0)) & 63))
            v14 = v14
            _ = fi28
            _ = v29
        }
    }
    v14 = "hello"
    if false {
        switch (v10 + ("\n\t" + "a")) {
        case "hello":
            v30 := v2[v10]
            _ = v30
        case "foo bar":
            if (false && (640 < 1)) {
                var v31 bool = true
                _ = v31
            }
        case "":
            v10 = v14
        }
        if true {
            var v32 bool = true
            fmt.Println(v9)
            _ = v32
        } else {
            switch (v4 * 42) {
            case 138:
                for fi33 := 0; fi33 < 5; fi33++ {
                    if true {
                        if false {
                            delete(v2, v14)
                        }
                        var v34 uint16 = (v3 % ((396) | 1))
                        _ = v34
                    } else {
                        var v35 int64 = v13
                        _ = v35
                    }
                    v36 := GenS2{F0: 25}
                    if (!((fi33 / (((fi33 | fi33)) | 1)) == (-(v4)))) {
                        if len(v1) > 0 {
                            v1[0] = (-(v4))
                        }
                        for fi37 := 0; fi37 < 14; fi37++ {
                            v38 := GenS1{F0: (v14 + "hello")}
                            var v39 int32 = v9
                            var v40 uint32 = 305
                            v14 = v14
                            _ = fi37
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        var v41 bool = false
                        _ = v41
                    }
                    _ = fi33
                    _ = v36
                }
            case 32767:
                v42 := GenS4{F0: 42, F1: int(v9)}
                _ = v42
            case -487:
                for fi43 := 0; fi43 < 2; fi43++ {
                    v13 = 98
                    _ = fi43
                }
            }
            v44 := GenS1{F0: "hello"}
            _ = v44
        }
        v45 := GenS3{F0: genFn0(false, int16(v9)), F1: false, F2: v9}
        _ = v45
    } else {
        if (65535 >= v3) {
            v9 = (^(v9))
        }
        v46 := GenS5{F0: v10, F1: 0, F2: 344}
        var v47 uint = genFn0((!true), -32768)
        _ = v46
        _ = v47
    }
    for fi48 := 0; fi48 < 3; fi48++ {
        fmt.Println(v10, v13, v14)
        fmt.Println(v13, fi48, v9)
        var v49 int8 = 0
        _ = fi48
        _ = v49
    }
    v50 := GenS3{F0: 1, F1: (!true), F2: (v9 >> ((801) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v9
}
