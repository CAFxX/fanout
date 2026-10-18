package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int64; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 0
    v3 := make(map[int]int)
    var v4 int16 = 1000
    var v5 int16 = (v4 + v4)
    v6 := GenS1{F0: 253}
    if false {
        v7 := GenS2{F0: 374}
        v8 := GenS3{F0: (-2.25 == 0.0), F1: false}
        _ = v7
        _ = v8
    }
    if (!(949 == 2147483647)) {
        v3[-1] = 32768
        var v9 bool = (-128 == 127)
        v4 += 1000
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            v5 += v4
            _ = fi11
        }
        var v12 float32 = 0.0
        _ = fi10
        _ = v12
    }
    fmt.Println(v4)
    v5 += v5
    var v13 uint64 = 334
    v14 := []int{-9223372036854775808}
    v4 = int16(v13)
    var v15 uint16 = 2
    switch -191 {
    case 127:
        switch -808 {
        case 825:
            v13 += v13
        case 947:
            switch "hello" {
            case "":
                v5 += (v4 << ((uint(v13)) & 63))
            }
        case 128:
            var v16 float32 = (0.1 / 0.1)
            _ = v16
        }
    case 818:
        for fi17 := 0; fi17 < 3; fi17++ {
            v18 := v3[127]
            _ = fi17
            _ = v18
        }
    }
    v5 -= (v5 % ((v5) | 1))
    fmt.Println(v4)
    for fi19 := 0; fi19 < 15; fi19++ {
        switch (fi19 | fi19) {
        case 314:
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 uint64 = (v13 % ((v13) | 1))
                if ("foo bar" > "\n\t") {
                    v21 -= (^(v13))
                    var v22 bool = true
                    var v23 int32 = -2147483648
                    _ = v22
                    _ = v23
                } else {
                    v2 = v2
                    var v24 float32 = (1e10 * 0.0)
                    _ = v24
                }
                if true {
                    v25 := v3[fi19]
                    switch "foo bar" {
                    case "a":
                        var v26 string = ""
                        _ = v26
                    case "hello":
                        if (-49 <= 100) {
                            v27 := GenS1{F0: -9223372036854775808}
                            _ = v27
                        } else {
                            if false {
                                v15 = (^(v15))
                            } else {
                                var v28 float64 = (0.0 * 0.0)
                                _ = v28
                            }
                            v1[("" + ("a" + ""))] = (fi20 % ((v25) | 1))
                            switch (^(fi19)) {
                            case 1:
                                if true {
                                    switch (fi19 + 743) {
                                    case 610:
                                        if false {
                                            if len(v14) > 0 {
                                                v14[0] = (^(fi19))
                                            }
                                            var v29 bool = true
                                            switch "" {
                                            case "hello":
                                                var v30 uint8 = 134
                                                _ = v30
                                            case "foo bar":
                                                v29 = true
                                            }
                                            fmt.Println(v21, v4)
                                            var v31 uint8 = 255
                                            _ = v29
                                            _ = v31
                                        } else {
                                            if (!((^(v15)) != 255)) {
                                                v32 := GenS4{F0: 1, F1: 822, F2: v15}
                                                var v33 string = ("foo bar" + "a")
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                    case 32768:
                                        v5 = v5
                                    case 4294967295:
                                        v14 = append(v14, int(v2))
                                    default:
                                        v34 := GenS5{F0: 127}
                                        _ = v34
                                    }
                                } else {
                                    var v35 int64 = int64(v2)
                                    v36 := GenS6{F0: v13}
                                    switch (fi20 * fi20) {
                                    case 42:
                                        v5 = (v4 / ((0) | 1))
                                    case 0:
                                        var v37 bool = false
                                        _ = v37
                                    case 32768:
                                        v38 := GenS4{F0: (v35 & 769), F1: v2, F2: v15}
                                        _ = v38
                                    }
                                    v35 = -557
                                    _ = v35
                                    _ = v36
                                }
                            case 260:
                                switch 202 {
                                case 4294967295:
                                    if false {
                                        v13 = uint64(v2)
                                        var v39 string = ""
                                        v21 = (v21 & uint64(v15))
                                        _ = v39
                                    }
                                }
                            case 127:
                                v4 -= (^(v4))
                            default:
                                var v40 uint64 = v21
                                if ((!(-2.25 < ((1.5 - 1.5) * -2.25))) || (((^(v40)) == 2) || (32767 <= 1000))) {
                                    if ((((((float64(3.14159) / (float64(((3.14159 - 3.14159) + float32((3.14159 / (((float64((((0.1 + (0.0 * 1e10)) * float32((((-2.25 / 1.0) + float64(0.1)) * (float64(1.5) * -2.25)))) / float32(0.1))) - 1.5) + 1.5) - 0.1))))) + float64((-2.25 / -2.25)))) / 1.0) / float64(float32(1.5))) * float64(float32((0.1 * 1.5)))) + float64((0.1 - 1e10))) <= 3.14159) {
                                        v41 := GenS6{F0: 18446744073709551615}
                                        v25 = -2147483648
                                        v13 = (^(v40))
                                        v40 = v21
                                        _ = v41
                                    } else {
                                        v42 := GenS7{F0: v2}
                                        switch ("hello" + "\n\t") {
                                        case "hello":
                                            fmt.Println(v21, v40)
                                        default:
                                            v13 = v40
                                        }
                                        _ = v42
                                    }
                                    var v43 string = "\n\t"
                                    v44 := GenS3{F0: ((!true) || (!false)), F1: (1e10 <= -2.25)}
                                    _ = v43
                                    _ = v44
                                }
                                _ = v40
                            }
                        }
                    default:
                        var v45 uint = (v2 / (((v2 % ((v2) | 1))) | 1))
                        _ = v45
                    }
                    fmt.Println(fi19, v2)
                    _ = v25
                } else {
                    if (!true) {
                        switch fi19 {
                        case 258:
                            v46 := GenS7{F0: (^(v2))}
                            v4 = (v4 << (((v2 << (((v2 / ((v2) | 1))) & 63))) & 63))
                            _ = len(v14)
                            _ = v46
                        case -2147483648:
                            var v47 float32 = (-2.25 - 0.0)
                            _ = v47
                        }
                        v4 = -32768
                        v48 := GenS5{F0: 0}
                        _ = v48
                    } else {
                        switch (fi20 ^ fi20) {
                        case 16:
                            fmt.Println(fi19, v13, v2)
                            _ = len(v14)
                            var v49 int64 = -824
                            _ = v49
                        }
                        v3[fi20] = 127
                    }
                    fmt.Println(fi20)
                    v5 = v5
                    v13 -= (v13 << ((255) & 63))
                    if false {
                        var v50 uint16 = 746
                        v51 := GenS8{F0: v21, F1: 42}
                        _ = v50
                        _ = v51
                    }
                }
                _ = fi20
                _ = v21
            }
        }
        if (!((-689 < 670) || ((v15 - v15) >= v15))) {
            v2 = v2
        }
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = fi19
    }
    var v52 uint = (v2 + (v2 + v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
