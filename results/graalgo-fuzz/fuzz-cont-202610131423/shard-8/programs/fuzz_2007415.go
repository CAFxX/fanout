package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 int8 }
func genFn0() int64 {
    var v6 uint64 = 713
    v7 := []int{423, int(v6)}
    _ = v6
    _ = v7
    return 895
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{""}
    v3 := make(map[string]string)
    var v4 int64 = 0
    var v5 int8 = -128
    v5 += -110
    v8 := GenS1{F0: v4, F1: 1000}
    if (42 == 4294967295) {
        switch "\n\t" {
        case "":
            switch -2147483648 {
            case 127:
                var v9 uint16 = 255
                _ = v9
            case -221:
                v10 := GenS1{F0: (v4 - 9223372036854775807), F1: 1}
                _ = v10
            case -9223372036854775808:
                delete(v3, ("foo bar" + "a"))
            }
        case "hello":
            v1[(((("hello" + (("" + "foo bar") + "a")) + "a") + "\n\t") + "a")] = 65535
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            if false {
                v4 = 1
                for fi12 := 0; fi12 < 13; fi12++ {
                    _ = len(v2)
                    v13 := v3[(("foo bar" + "a") + "a")]
                    _ = fi12
                    _ = v13
                }
            }
            switch "hello" {
            case "a":
                v4 = (v4 * genFn0())
            default:
                for fi14 := 0; fi14 < 8; fi14++ {
                    if (-2.25 < float64((1e10 / float32((float64(3.14159) / 1.5))))) {
                        v5 = v5
                    }
                    _ = fi14
                }
            }
            if (1 <= (-(fi11))) {
                switch ("\n\t" + "\n\t") {
                case "foo bar":
                    switch 42 {
                    case 127:
                        v5 += (v5 / (((v5 + (v5 ^ (-(v5))))) | 1))
                    default:
                        var v15 int64 = genFn0()
                        _ = v15
                    }
                case "":
                    for fi16 := 0; fi16 < 4; fi16++ {
                        var v17 float64 = -2.25
                        _ = fi16
                        _ = v17
                    }
                }
                switch "\n\t" {
                case "hello":
                    switch fi11 {
                    case 32768:
                        var v18 uint64 = 2
                        _ = v18
                    case 1:
                        v19 := GenS1{F0: v4, F1: 0}
                        _ = v19
                    }
                case "a":
                    switch "hello" {
                    case "":
                        var v20 uint16 = 1
                        _ = v20
                    case "a":
                        var v21 uint = 734
                        _ = v21
                    default:
                        for fi22 := 0; fi22 < 6; fi22++ {
                            _ = len(v3)
                            if ((!true) && (!true)) {
                                v5 = (^(v5))
                            }
                            var v23 int = (^(fi11))
                            _ = fi22
                            _ = v23
                        }
                    }
                case "\n\t":
                    for fi24 := 0; fi24 < 13; fi24++ {
                        if len(v2) > 0 {
                            v2[0] = ("a" + (((((("a" + "a") + "") + "a") + "a") + ("\n\t" + "\n\t")) + ""))
                        }
                        _ = fi24
                    }
                default:
                    fmt.Println(fi11)
                }
            }
            _ = fi11
        }
        v2 = append(v2, ("a" + "hello"))
    } else {
        fmt.Println(v5, v5)
        v25 := GenS1{F0: 9223372036854775807, F1: 32767}
        var v26 int32 = -316
        _ = v25
        _ = v26
    }
    if false {
        var v27 uint8 = 127
        _ = v27
    }
    if true {
        for fi28 := 0; fi28 < 18; fi28++ {
            switch fi28 {
            case 852:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -1:
                switch fi28 {
                case 185:
                    v5 += v5
                case 42:
                    fmt.Println(fi28, v5, v4)
                }
            default:
                v5 -= (^(v5))
            }
            _ = fi28
        }
        v5 = (v5 - (v5 % ((v5) | 1)))
        v29 := GenS2{F0: -1}
        _ = v29
    }
    v5 = (v5 | (v5 >> ((637) & 63)))
    for fi30 := 0; fi30 < 8; fi30++ {
        switch ("\n\t" + ("\n\t" + "")) {
        case "\n\t":
            v31 := GenS3{F0: 128, F1: 555, F2: 961}
            _ = v31
        case "":
            switch fi30 {
            case 1:
                var v32 uint64 = 1
                _ = v32
            case 0:
                v33 := GenS4{F0: genFn0()}
                _ = v33
            case 613:
                var v34 uint32 = 2
                _ = v34
            }
        default:
            v35 := GenS5{F0: 1, F1: (v5 << ((65535) & 63))}
            _ = v35
        }
        v2 = append(v2, "")
        _ = fi30
    }
    var v36 uint64 = 18446744073709551615
    for fi37 := 0; fi37 < 5; fi37++ {
        v5 -= (v5 ^ 30)
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v8
}
