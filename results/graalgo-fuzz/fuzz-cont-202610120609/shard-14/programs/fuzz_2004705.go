package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint16; F1 uint64; F2 bool }
type GenS5 struct { F0 string; F1 int16; F2 int16 }
func genFn0() string {
    var v6 float32 = 0.1
    var v7 uint64 = 18446744073709551615
    switch ("a" + ("" + "hello")) {
    case "a":
        switch -9223372036854775808 {
        case 42:
            fmt.Println(v6, v6)
        default:
            v7 = (^(v7))
        }
    case "hello":
        var v8 int64 = 205
        _ = v8
    case "":
        if false {
            v6 += 3.14159
            var v9 float32 = 1e10
            _ = v9
        }
    default:
        v7 = v7
    }
    v10 := []int{}
    _ = v10
    _ = v6
    _ = v7
    return ""
}
func genFn1() uint64 {
    switch "foo bar" {
    case "\n\t":
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 string = "hello"
            var v13 int32 = 2147483647
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 int16 = 0
                if (v15 > v15) {
                    v13 += (v13 | (-(v13)))
                } else {
                    var v16 int8 = 39
                    v17 := []int{1}
                    var v18 bool = false
                    _ = v16
                    _ = v17
                    _ = v18
                }
                switch fi14 {
                case -533:
                    v12 = ((v12 + "foo bar") + v12)
                }
                v19 := []int{(fi11 >> ((4294967295) & 63))}
                var v20 bool = (!(3.14159 != ((0.0 - 1e10) / -2.25)))
                _ = fi14
                _ = v15
                _ = v19
                _ = v20
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    case "hello":
        v21 := []int{-1}
        _ = v21
    case "a":
        var v22 uint64 = 42
        _ = v22
    }
    var v23 int16 = 0
    switch "a" {
    case "\n\t":
        v23 = 883
    case "":
        if true {
            fmt.Println(v23)
            v24 := GenS1{F0: 255, F1: 874}
            switch "\n\t" {
            case "foo bar":
                switch -892 {
                case -1:
                    switch "\n\t" {
                    case "a":
                        for fi25 := 0; fi25 < 7; fi25++ {
                            v23 -= (v23 - 0)
                            v26 := []int{}
                            _ = fi25
                            _ = v26
                        }
                    case "":
                        var v27 uint = 139
                        _ = v27
                    }
                case -409:
                    if (((float32((float64(((float32(0.0) * -2.25) + (float32(0.1) / 1.0))) + 1.5)) - (0.1 - 0.1)) - (1.5 / 3.14159)) <= (1.5 * (3.14159 - -2.25))) {
                        var v28 float32 = 0.0
                        var v29 float32 = float32(1.5)
                        _ = v28
                        _ = v29
                    } else {
                        v30 := GenS2{F0: "foo bar"}
                        var v31 string = ("hello" + "foo bar")
                        _ = v30
                        _ = v31
                    }
                case -2147483648:
                    v23 += (v23 * v23)
                default:
                    if (-45 < -95) {
                        if false {
                            v23 -= v23
                            var v32 float64 = ((1.5 - (0.0 / 3.14159)) * (1.5 / 0.1))
                            for fi33 := 0; fi33 < 13; fi33++ {
                                var v34 uint32 = 65535
                                v34 -= (v34 * (v34 ^ v34))
                                _ = fi33
                                _ = v34
                            }
                            _ = v32
                        }
                    }
                }
            }
            _ = v24
        } else {
            fmt.Println(v23)
            v35 := GenS3{F0: 2, F1: 2, F2: -2147483648}
            _ = v35
        }
    default:
        var v36 string = "hello"
        _ = v36
    }
    _ = v23
    return 55
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]int)
    v3 := []uint32{2}
    var v4 float32 = v1
    var v5 int = 9223372036854775807
    for fi37 := 0; fi37 < 10; fi37++ {
        v38 := GenS2{F0: "a"}
        _ = fi37
        _ = v38
    }
    v39 := GenS4{F0: 0, F1: 18446744073709551615, F2: (128 > 255)}
    switch v5 {
    case 9223372036854775807:
        switch "" {
        case "foo bar":
            v40 := GenS2{F0: ((("foo bar" + ((("hello" + "a") + "a") + "\n\t")) + "") + "foo bar")}
            _ = v40
        case "\n\t":
            if true {
                var v41 int8 = 127
                _ = v41
            } else {
                v42 := GenS1{F0: 128, F1: 255}
                var v43 int64 = 9223372036854775807
                _ = v42
                _ = v43
            }
        default:
            for fi44 := 0; fi44 < 20; fi44++ {
                v45 := v2[4294967295]
                _ = fi44
                _ = v45
            }
        }
    default:
        var v46 uint16 = 42
        _ = v46
    }
    fmt.Println(v1, v1)
    fmt.Println(v5)
    switch "foo bar" {
    case "\n\t":
        var v47 string = "a"
        _ = v47
    case "foo bar":
        if false {
            fmt.Println(v1)
        } else {
            v5 += (v5 >> ((42) & 63))
        }
    case "hello":
        if false {
            v48 := GenS4{F0: 32767, F1: 1, F2: false}
            _ = v48
        }
    default:
        for fi49 := 0; fi49 < 5; fi49++ {
            if len(v3) > 0 {
                v3[0] = 42
            }
            _ = fi49
        }
    }
    fmt.Println(v1)
    v3 = append(v3, 706)
    var v50 uint16 = 42
    switch v5 {
    case 4294967295:
        for fi51 := 0; fi51 < 2; fi51++ {
            v52 := GenS2{F0: "a"}
            if false {
                v1 -= v1
                _ = len(v2)
                for fi53 := 0; fi53 < 9; fi53++ {
                    var v54 int = 42
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    switch ("a" + "foo bar") {
                    case "\n\t":
                        var v55 int32 = 42
                        _ = v55
                    case "foo bar":
                        if true {
                            var v56 float32 = 3.14159
                            var v57 float64 = 3.14159
                            _ = v56
                            _ = v57
                        }
                    case "hello":
                        var v58 int64 = 557
                        _ = v58
                    default:
                        var v59 string = ("a" + ("hello" + ""))
                        _ = v59
                    }
                    _ = fi53
                    _ = v54
                }
            } else {
                if (true && (!(fi51 <= (v5 % (((fi51 % ((-625) | 1))) | 1))))) {
                    var v60 uint64 = 77
                    fmt.Println(v4, v5)
                    _ = v60
                } else {
                    v1 += v1
                }
            }
            switch "hello" {
            case "\n\t":
                v5 = v5
            case "":
                var v61 uint16 = v50
                _ = v61
            }
            _ = fi51
            _ = v52
        }
    case 131:
        if true {
            v50 -= v50
            var v62 string = ("\n\t" + "hello")
            v2[1] = (v5 - (v5 - v5))
            _ = v62
        }
    default:
        var v63 uint8 = 42
        _ = v63
    }
    v64 := GenS2{F0: ""}
    v65 := GenS5{F0: "a", F1: 1000, F2: 42}
    v66 := GenS2{F0: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v50
    _ = v64
    _ = v65
    _ = v66
}
