package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint8; F1 int; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string; F1 string; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = ((0.0 - 0.0) * (0.1 / 1.0))
    v3 := make(map[string]bool)
    var v4 uint16 = 42
    switch 32767 {
    case 568:
        v5 := GenS1{F0: v4}
        _ = v5
    case 127:
        v6 := GenS1{F0: 32767}
        _ = v6
    }
    v4 = (^(v4))
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        if ((true || true) && (true || true)) {
            v2 = ((v2 + (v2 * (v2 / 1.0))) * v2)
        } else {
            if ((!true) && false) {
                if ((-2147483648 < -471) && true) {
                    if true {
                        v2 = v2
                        var v7 int8 = -1
                        _ = v7
                    }
                    v2 = v2
                    var v8 uint = 2
                    _ = v8
                }
                if true {
                    var v9 uint64 = 1
                    if (-1 <= -9223372036854775808) {
                        switch ("foo bar" + "a") {
                        case "a":
                            v10 := GenS1{F0: v4}
                            _ = v10
                        default:
                            if true {
                                v4 = 1
                            } else {
                                for fi11 := 0; fi11 < 12; fi11++ {
                                    for fi12 := 0; fi12 < 13; fi12++ {
                                        v4 -= v4
                                        v13 := GenS1{F0: (^(v4))}
                                        fmt.Println(v2, fi11)
                                        _ = fi12
                                        _ = v13
                                    }
                                    if true {
                                        v14 := GenS2{F0: v9}
                                        switch "a" {
                                        case "hello":
                                            var v15 float64 = -2.25
                                            _ = v15
                                        case "":
                                            var v16 uint8 = 127
                                            _ = v16
                                        case "\n\t":
                                            for fi17 := 0; fi17 < 2; fi17++ {
                                                fmt.Println(v4, v4, v9)
                                                v2 = 0.0
                                                v18 := GenS3{F0: (fi17 <= fi11), F1: 91}
                                                _ = fi17
                                                _ = v18
                                            }
                                        }
                                        _ = v14
                                    } else {
                                        v19 := []int{fi11, fi11}
                                        v4 = v4
                                        var v20 string = ("hello" + "foo bar")
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = fi11
                                }
                                var v21 float64 = ((0.0 - 0.0) / 1.0)
                                fmt.Println(v2, v21)
                                _ = v21
                            }
                        }
                    } else {
                        var v22 uint16 = 32767
                        _ = v22
                    }
                    _ = v9
                } else {
                    for fi23 := 0; fi23 < 15; fi23++ {
                        if true {
                            var v24 uint = 983
                            var v25 int16 = -655
                            _ = v24
                            _ = v25
                        }
                        var v26 uint8 = 1
                        _ = fi23
                        _ = v26
                    }
                    if (0.0 <= 3.14159) {
                        delete(v3, "a")
                        if (591 <= 380) {
                            fmt.Println(v2, v2)
                            switch "" {
                            case "foo bar":
                                v2 += 0.1
                            case "\n\t":
                                v27 := GenS2{F0: 44}
                                _ = v27
                            case "a":
                                switch "hello" {
                                case "hello":
                                    v4 = 0
                                }
                            }
                        }
                        for fi28 := 0; fi28 < 19; fi28++ {
                            v2 -= (v2 / (((float32(0.0) + 1.5) - ((v2 / (v2 - 0.0)) * (v2 - (v2 - (v2 + 0.1))))) + (0.0 / v2)))
                            for fi29 := 0; fi29 < 9; fi29++ {
                                v4 -= 1
                                v30 := GenS1{F0: 2}
                                switch "a" {
                                case "a":
                                    var v31 int16 = -807
                                    _ = v31
                                }
                                _ = fi29
                                _ = v30
                            }
                            _ = fi28
                        }
                    } else {
                        fmt.Println(v4)
                        switch "" {
                        case "hello":
                            if true {
                                var v32 float64 = float64(float32((1.5 * -2.25)))
                                _ = v32
                            }
                        case "":
                            v33 := []int{-340, 127, 2147483647, 0}
                            _ = v33
                        default:
                            var v34 int8 = -1
                            _ = v34
                        }
                        v35 := []int{}
                        _ = v35
                    }
                    var v36 bool = (false && (true || true))
                    _ = v36
                }
                fmt.Println(v4)
            } else {
                for fi37 := 0; fi37 < 1; fi37++ {
                    var v38 uint32 = 218
                    _ = fi37
                    _ = v38
                }
                var v39 int32 = -2147483648
                for fi40 := 0; fi40 < 20; fi40++ {
                    v39 = (v39 / ((v39) | 1))
                    _ = fi40
                }
                _ = v39
            }
            var v41 uint64 = 42
            _ = v41
        }
    }
    if true {
        v42 := v3["hello"]
        v2 = v2
        _ = v42
    } else {
        var v43 int = 1
        if (true && false) {
            fmt.Println(v43)
            v2 = (3.14159 - ((float32((((float64(v2) * 0.1) * (0.0 / 1.0)) / 1.0)) - (v2 / ((-2.25 + (float32(1e10) - ((0.1 / v2) + (v2 + v2)))) - 1.5))) / (v2 * 0.0)))
            v44 := GenS4{F0: 1, F1: v43, F2: -1}
            _ = v44
        }
        for fi45 := 0; fi45 < 18; fi45++ {
            var v46 int64 = 0
            v46 += (v46 + -829)
            for fi47 := 0; fi47 < 17; fi47++ {
                v48 := GenS2{F0: 42}
                v49 := v1[-645]
                v2 -= v2
                _ = fi47
                _ = v48
                _ = v49
            }
            _ = fi45
            _ = v46
        }
        _ = v43
    }
    v50 := GenS5{F0: 520}
    if (1 == 1) {
        var v51 uint32 = 4294967295
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v52 uint8 = 1
            _ = v52
        }
        _ = v51
    }
    for fi53 := 0; fi53 < 14; fi53++ {
        var v54 uint32 = 42
        v55 := []int{(^(fi53))}
        if true {
            v54 += 0
        } else {
            var v56 int8 = 0
            _ = v56
        }
        _ = fi53
        _ = v54
        _ = v55
    }
    v2 -= v2
    v57 := GenS1{F0: v4}
    fmt.Println(v4)
    v4 = 65535
    var v58 float64 = 1e10
    for fi59 := 0; fi59 < 3; fi59++ {
        for fi60 := 0; fi60 < 3; fi60++ {
            v61 := GenS6{F0: "foo bar", F1: (("hello" + ("hello" + "")) + (((("foo bar" + "foo bar") + "\n\t") + "") + ("" + ("\n\t" + "foo bar")))), F2: 32767}
            if (true && (86 != -238)) {
                if true {
                    fmt.Println(v2)
                    v62 := []int{128, 2147483647}
                    fmt.Println(v4)
                    _ = v62
                } else {
                    var v63 uint8 = uint8(v4)
                    _ = v63
                }
            }
            _ = fi60
            _ = v61
        }
        _ = fi59
    }
    var v64 bool = (v2 >= (0.1 * v2))
    v65 := GenS5{F0: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v58:", v58)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v57
    _ = v58
    _ = v64
    _ = v65
}
