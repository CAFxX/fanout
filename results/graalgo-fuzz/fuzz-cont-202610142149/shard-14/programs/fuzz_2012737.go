package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint16; F2 uint32 }
func genFn0(p0 string, p1 int) bool {
    var v7 bool = (true || (!(((-845 < -1) || false) && false)))
    for fi8 := 0; fi8 < 19; fi8++ {
        v7 = true
        switch (p1 * (fi8 ^ 32767)) {
        case 42:
            fmt.Println(fi8, p1, p0)
        }
        _ = fi8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        if v7 {
            v10 := GenS1{F0: 42, F1: 615, F2: -1}
            v7 = (100 == 127)
            _ = v10
        }
        v11 := GenS1{F0: 203, F1: 155, F2: 32767}
        if v7 {
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 float32 = (-2.25 + 0.1)
                _ = fi12
                _ = v13
            }
            v7 = v7
        }
        _ = fi9
        _ = v11
    }
    switch (p0 + "a") {
    case "":
        v14 := []int{(^(p1)), (p1 & (^(p1)))}
        _ = v14
    case "foo bar":
        var v15 float32 = ((1.5 / 1e10) - (0.0 + (float32((0.1 + (float64((float32((0.0 + 1e10)) - float32(float64((1e10 * (0.0 + 0.1)))))) / 1.5))) * 3.14159)))
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v7
    return (-1 < 0)
}

func main() {
    v1 := []string{"hello", ("hello" + ("hello" + ("foo bar" + "foo bar"))), "hello"}
    var v2 uint = 2
    var v3 int = -2147483648
    var v4 int16 = 0
    var v5 uint = (v2 | (v2 + (^(v2))))
    var v6 uint64 = 912
    v4 = (v4 % ((v4) | 1))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = (^(v3))
    for fi16 := 0; fi16 < 7; fi16++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            v6 -= v6
            switch (^(fi16)) {
            case -9223372036854775808:
                _ = len(v1)
            default:
                v4 += (v4 - (v4 - v4))
            }
            _ = fi17
        }
        _ = fi16
    }
    _ = len(v1)
    if true {
        v2 = (v2 * (v5 & v2))
        if (!((((true && false) && (1 < 42)) || false) || false)) {
            fmt.Println(v4, v2, v5)
        } else {
            fmt.Println(v5, v3)
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 float64 = 1.5
                _ = fi18
                _ = v19
            }
            if ((v4 * (-(v4))) > v4) {
                if true {
                    for fi20 := 0; fi20 < 11; fi20++ {
                        for fi21 := 0; fi21 < 6; fi21++ {
                            fmt.Println(v5)
                            switch "hello" {
                            case "hello":
                                switch fi21 {
                                case 1:
                                    var v22 string = "\n\t"
                                    _ = v22
                                case 932:
                                    v23 := GenS1{F0: 255, F1: 0, F2: v4}
                                    _ = v23
                                case 2147483647:
                                    if true {
                                        v1 = append(v1, "")
                                        if true {
                                            switch "" {
                                            case "a":
                                                var v24 float64 = (1.5 - (float64((1.5 / -2.25)) / 1.0))
                                                _ = v24
                                            case "\n\t":
                                                if true {
                                                    v5 = 4294967295
                                                    var v25 uint32 = 65535
                                                    v26 := GenS2{F0: 42}
                                                    _ = v25
                                                    _ = v26
                                                }
                                            case "":
                                                var v27 int8 = 0
                                                _ = v27
                                            }
                                            var v28 float32 = 0.1
                                            var v29 int8 = 67
                                            _ = v28
                                            _ = v29
                                        }
                                    }
                                }
                            case "\n\t":
                                fmt.Println(v3, fi20)
                            case "a":
                                var v30 int8 = -128
                                _ = v30
                            default:
                                fmt.Println(v5, v6)
                            }
                            _ = fi21
                        }
                        v31 := GenS1{F0: 128, F1: 255, F2: v4}
                        _ = fi20
                        _ = v31
                    }
                    var v32 bool = true
                    v32 = (!v32)
                    _ = v32
                }
            }
        }
        var v33 uint32 = 195
        _ = v33
    } else {
        fmt.Println(v2, v3)
        fmt.Println(v5, v2)
    }
    v5 = v2
    if (false || true) {
        for fi34 := 0; fi34 < 18; fi34++ {
            var v35 int8 = int8(v4)
            var v36 int64 = -192
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    switch "a" {
    case "a":
        var v37 float64 = ((1e10 - 0.1) / ((3.14159 * (3.14159 - float64((3.14159 / 1e10)))) / -2.25))
        _ = v37
    case "hello":
        if ("a" < "foo bar") {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v38 := GenS3{F0: 0, F1: -1, F2: 42}
                _ = v38
            } else {
                var v39 float32 = (0.0 - (0.1 - 0.1))
                _ = v39
            }
            var v40 int16 = (v4 >> (((^(v2))) & 63))
            _ = v40
        }
    }
    if false {
        var v41 uint16 = 1
        _ = v41
    }
    if ((-(v3)) == -2147483648) {
        v5 += (v2 << (((v2 ^ 2)) & 63))
        for fi42 := 0; fi42 < 12; fi42++ {
            v43 := GenS1{F0: 127, F1: 255, F2: (^(v4))}
            switch fi42 {
            case 9223372036854775807:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                var v44 int = fi42
                _ = v44
            }
            var v45 float64 = 1.5
            _ = fi42
            _ = v43
            _ = v45
        }
        if true {
            v46 := GenS4{F0: 1, F1: 0, F2: 1}
            _ = v46
        }
    } else {
        _ = len(v1)
        var v47 float32 = (-2.25 - (0.1 - (1.5 + -2.25)))
        var v48 bool = (v2 <= (v5 * 65535))
        _ = v47
        _ = v48
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
