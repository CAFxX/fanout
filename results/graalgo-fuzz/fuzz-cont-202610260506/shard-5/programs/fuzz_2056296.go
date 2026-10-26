package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 uint16; F2 int64 }
type GenS5 struct { F0 int16; F1 int32; F2 int32 }
func genFn0() bool {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{42, -1, fi7}
        v9 := GenS1{F0: "hello"}
        v10 := GenS1{F0: "a"}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if (127 <= 255) {
        var v11 int16 = 20
        _ = v11
    }
    var v12 float64 = float64(0.0)
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := []int{fi13, fi13, -512}
        for fi15 := 0; fi15 < 14; fi15++ {
            if true {
                if len(v14) > 0 {
                    v14[0] = fi13
                }
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    _ = v12
    return false
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]uint32)
    var v3 int16 = 1000
    var v4 int16 = 42
    v5 := make(map[uint32]uint32)
    var v6 int = 4294967295
    var v16 int = (v6 / ((v6) | 1))
    v3 += -32768
    v17 := []int{v16, v6}
    _ = len(v2)
    if false {
        switch "\n\t" {
        case "\n\t":
            v18 := GenS1{F0: ""}
            _ = v18
        default:
            if (v1 >= v1) {
                switch 42 {
                case 127:
                    for fi19 := 0; fi19 < 19; fi19++ {
                        v1 -= (v1 + v1)
                        var v20 uint = 2
                        _ = fi19
                        _ = v20
                    }
                case -9223372036854775808:
                    if false {
                        fmt.Println(v16)
                        delete(v5, 888)
                        for fi21 := 0; fi21 < 18; fi21++ {
                            v4 = v3
                            switch "hello" {
                            case "foo bar":
                                if false {
                                    for fi22 := 0; fi22 < 12; fi22++ {
                                        switch 32768 {
                                        case -871:
                                            if (((^(v1)) <= 2) && (!false)) {
                                                v23 := GenS2{F0: true, F1: 32767}
                                                v6 -= fi21
                                                _ = v23
                                            } else {
                                                fmt.Println(v6, v16, v1)
                                                v24 := GenS1{F0: ("" + "a")}
                                                _ = v24
                                            }
                                        }
                                        _ = fi22
                                    }
                                    _ = len(v17)
                                    for fi25 := 0; fi25 < 17; fi25++ {
                                        fmt.Println(v6, v4, v4)
                                        fmt.Println(fi21, fi25)
                                        _ = fi25
                                    }
                                } else {
                                    var v26 int16 = v3
                                    _ = v26
                                }
                            case "a":
                                fmt.Println(v1, v3)
                            case "":
                                switch "a" {
                                case "\n\t":
                                    var v27 uint64 = 42
                                    _ = v27
                                case "hello":
                                    if true {
                                        fmt.Println(v4)
                                        if ((v1 << ((65535) & 63)) < (v1 ^ v1)) {
                                            var v28 int = fi21
                                            _ = v28
                                        } else {
                                            v29 := GenS1{F0: "a"}
                                            var v30 int16 = 42
                                            _ = v29
                                            _ = v30
                                        }
                                        var v31 int16 = (v4 / (((v4 << ((91) & 63))) | 1))
                                        _ = v31
                                    }
                                }
                            }
                            _ = fi21
                        }
                    } else {
                        fmt.Println(v4)
                        if (!true) {
                            var v32 bool = false
                            for fi33 := 0; fi33 < 5; fi33++ {
                                var v34 int64 = -9223372036854775808
                                if (true && v32) {
                                    for fi35 := 0; fi35 < 8; fi35++ {
                                        v16 += v6
                                        _ = fi35
                                    }
                                    var v36 int16 = v3
                                    _ = v36
                                }
                                if false {
                                    var v37 float32 = (1.5 / (0.1 / float32(float64(1.5))))
                                    v32 = (0 > 4294967295)
                                    for fi38 := 0; fi38 < 5; fi38++ {
                                        v1 += v1
                                        _ = fi38
                                    }
                                    _ = v37
                                }
                                _ = fi33
                                _ = v34
                            }
                            if (32767 == 42) {
                                for fi39 := 0; fi39 < 11; fi39++ {
                                    switch "foo bar" {
                                    case "":
                                        var v40 uint64 = 0
                                        _ = v40
                                    case "a":
                                        var v41 uint16 = 255
                                        _ = v41
                                    case "foo bar":
                                        fmt.Println(v1)
                                    }
                                    v32 = v32
                                    _ = fi39
                                }
                                v42 := GenS2{F0: false, F1: (^(v3))}
                                v43 := GenS2{F0: (!(-1 > -9223372036854775808)), F1: v3}
                                _ = v42
                                _ = v43
                            }
                            _ = v32
                        }
                    }
                case 128:
                    switch "a" {
                    case "a":
                        v3 = 0
                    case "hello":
                        var v44 int32 = -2147483648
                        _ = v44
                    default:
                        var v45 uint8 = v1
                        _ = v45
                    }
                }
                var v46 int32 = -569
                var v47 uint = 65535
                _ = v46
                _ = v47
            }
        }
        if true {
            fmt.Println(v4, v4)
        } else {
            v1 += (v1 & (v1 & (v1 & 127)))
        }
        _ = len(v17)
    }
    delete(v2, (v6 << ((42) & 63)))
    var v48 uint64 = 988
    var v49 uint64 = 42
    var v50 int = 4294967295
    var v51 int = (v6 - (v6 & (v6 % ((-1) | 1))))
    v52 := GenS3{F0: -1}
    _ = len(v17)
    if true {
        v50 -= v50
        fmt.Println(v49, v3)
    } else {
        if false {
            for fi53 := 0; fi53 < 1; fi53++ {
                var v54 uint8 = (v1 | (v1 - (v1 >> ((269) & 63))))
                _ = fi53
                _ = v54
            }
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            for fi55 := 0; fi55 < 10; fi55++ {
                switch (v51 - fi55) {
                case -664:
                    switch (v50 >> ((65535) & 63)) {
                    case 272:
                        _ = len(v17)
                    case 127:
                        var v56 string = "a"
                        _ = v56
                    }
                case 128:
                    v49 = v49
                case -139:
                    var v57 uint16 = 873
                    _ = v57
                }
                var v58 int16 = -656
                for fi59 := 0; fi59 < 18; fi59++ {
                    v3 -= 42
                    v60 := GenS4{F0: 1, F1: 0, F2: 9223372036854775807}
                    _ = fi59
                    _ = v60
                }
                _ = fi55
                _ = v58
            }
        }
        var v61 float32 = 0.1
        _ = v61
    }
    v51 = 32768
    for fi62 := 0; fi62 < 7; fi62++ {
        for fi63 := 0; fi63 < 10; fi63++ {
            var v64 float64 = (float64(3.14159) / (1e10 * 1.5))
            _ = fi63
            _ = v64
        }
        _ = fi62
    }
    v65 := GenS5{F0: -32768, F1: 138, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
    _ = v65
}
