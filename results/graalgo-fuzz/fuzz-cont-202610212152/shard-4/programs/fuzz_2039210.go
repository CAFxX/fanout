package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int8; F1 string; F2 bool }
type GenS4 struct { F0 int8; F1 uint; F2 int32 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{32767}
    var v3 int = -2147483648
    switch 351 {
    case 0:
        for fi4 := 0; fi4 < 1; fi4++ {
            if false {
                switch (-(fi4)) {
                case 230:
                    fmt.Println(v3)
                case 9223372036854775807:
                    var v5 uint32 = 42
                    _ = v5
                }
            }
            _ = fi4
        }
    default:
        switch "foo bar" {
        case "foo bar":
            v3 += (^(v3))
        case "":
            switch ("a" + "foo bar") {
            case "foo bar":
                v6 := GenS1{F0: -48}
                _ = v6
            }
        case "a":
            var v7 int64 = 1
            _ = v7
        }
    }
    _ = len(v2)
    switch "a" {
    case "\n\t":
        v8 := v1[187]
        _ = v8
    }
    switch "hello" {
    case "a":
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint64 = 4294967295
            v1[649] = ("" + "\n\t")
            switch ("foo bar" + "\n\t") {
            case "hello":
                switch ("foo bar" + "") {
                case "":
                    v11 := GenS1{F0: 0}
                    _ = v11
                case "hello":
                    for fi12 := 0; fi12 < 2; fi12++ {
                        if (!true) {
                            var v13 float64 = ((1e10 * float64((-2.25 - 1.5))) / 0.1)
                            var v14 int16 = 1
                            v3 += v3
                            v15 := GenS1{F0: 100}
                            var v16 string = "foo bar"
                            delete(v1, 65535)
                            v17 := GenS1{F0: 100}
                            if (0 > 159) {
                                if true {
                                    if false {
                                        var v18 int = 619
                                        v19 := GenS1{F0: -13}
                                        _ = v18
                                        _ = v19
                                    } else {
                                        switch v16 {
                                        case "hello":
                                            v20 := GenS2{F0: -9223372036854775808, F1: 65535}
                                            _ = v20
                                        default:
                                            if false {
                                                var v21 uint8 = 2
                                                v3 = 1
                                                if (v14 < v14) {
                                                    v22 := GenS2{F0: 1, F1: 531}
                                                    var v23 uint64 = (^(v10))
                                                    _ = v22
                                                    _ = v23
                                                }
                                                v10 = uint64(v21)
                                                if ((v21 / (((v21 - 128)) | 1)) < (v21 | 167)) {
                                                    if true {
                                                        var v24 uint = 4294967295
                                                        _ = v24
                                                    } else {
                                                        switch (^(v3)) {
                                                        case 2147483647:
                                                            if true {
                                                                switch v16 {
                                                                case "\n\t":
                                                                    var v25 uint16 = 2
                                                                    _ = v25
                                                                case "":
                                                                    v26 := GenS2{F0: 175, F1: 65535}
                                                                    _ = v26
                                                                }
                                                                v27 := GenS3{F0: 31, F1: v16, F2: (-1 > -1)}
                                                                _ = v27
                                                            } else {
                                                                switch ((v16 + (v16 + v16)) + "") {
                                                                case "foo bar":
                                                                    v16 = v16
                                                                default:
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                                v3 -= v3
                                                                var v28 int16 = v14
                                                                _ = v28
                                                            }
                                                        case 127:
                                                            var v29 int32 = -1
                                                            _ = v29
                                                        }
                                                    }
                                                    fmt.Println(fi12, v10, v21)
                                                }
                                                _ = v21
                                            }
                                        }
                                        v16 = v16
                                        v13 -= float64(1.5)
                                    }
                                }
                            }
                            v10 += v10
                            _ = v13
                            _ = v14
                            _ = v15
                            _ = v16
                            _ = v17
                        } else {
                            if (42 == 42) {
                                v10 -= (^(v10))
                                var v30 uint8 = 255
                                _ = v30
                            } else {
                                _ = len(v1)
                                v2 = append(v2, (v3 | (v3 << ((1) & 63))))
                            }
                            var v31 string = "foo bar"
                            _ = v31
                        }
                        v2 = append(v2, (-(fi12)))
                        v1[uint32(v10)] = "hello"
                        _ = fi12
                    }
                case "foo bar":
                    if ("foo bar" >= "\n\t") {
                        if len(v2) > 0 {
                            v2[0] = (v3 & -628)
                        }
                        v32 := GenS4{F0: 127, F1: 1, F2: -1}
                        _ = v32
                    }
                }
            case "a":
                v10 = v10
            }
            _ = fi9
            _ = v10
        }
    case "":
        v3 = 0
    case "foo bar":
        fmt.Println(v3)
    }
    var v33 int16 = -212
    v34 := GenS5{F0: 1, F1: -279}
    if true {
        v35 := GenS6{F0: (("a" + "foo bar") + ((("foo bar" + "foo bar") + "") + "a")), F1: 42}
        var v36 uint16 = 65535
        v36 = v36
        _ = v35
        _ = v36
    } else {
        var v37 float32 = (((1e10 * (0.0 / (1.5 - 0.1))) - ((1.5 - (0.0 + 0.0)) / -2.25)) * 1e10)
        var v38 string = "\n\t"
        v33 = 558
        _ = v37
        _ = v38
    }
    switch "" {
    case "hello":
        fmt.Println(v33)
    }
    switch (v3 % (((v3 * (^(v3)))) | 1)) {
    case -638:
        v3 = 127
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        _ = len(v2)
        switch "" {
        case "a":
            var v40 uint = 255
            _ = v40
        case "\n\t":
            var v41 uint64 = 1
            _ = v41
        case "foo bar":
            var v42 uint8 = 128
            _ = v42
        default:
            if (2147483647 != 0) {
                v3 = v3
                var v43 int = v3
                var v44 int32 = -1
                _ = v43
                _ = v44
            }
        }
        _ = fi39
    }
    var v45 uint64 = 2
    for fi46 := 0; fi46 < 16; fi46++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v47 uint = 255
        _ = fi46
        _ = v47
    }
    v48 := GenS6{F0: ("foo bar" + "foo bar"), F1: -2147483648}
    v49 := v1[65535]
    v33 = (v33 - -32768)
    fmt.Println(v3)
    if (!true) {
        var v50 float64 = ((1.5 + (3.14159 + float64((1.5 / float32(1e10))))) / 3.14159)
        _ = v50
    } else {
        fmt.Println(v3, v33)
        var v51 bool = (false && (v33 < (^(v33))))
        switch v3 {
        case 0:
            v52 := v1[935]
            _ = v52
        case -276:
            for fi53 := 0; fi53 < 3; fi53++ {
                var v54 int8 = 0
                var v55 int16 = (v33 | v33)
                v51 = (867 >= 544)
                _ = fi53
                _ = v54
                _ = v55
            }
        default:
            for fi56 := 0; fi56 < 5; fi56++ {
                var v57 float32 = 1.5
                if len(v2) > 0 {
                    v2[0] = int(v45)
                }
                _ = fi56
                _ = v57
            }
        }
        _ = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v45
    _ = v48
    _ = v49
}
