package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint; F1 uint; F2 string }
type GenS6 struct { F0 int; F1 int }
func genFn0() uint16 {
    var v4 int16 = -534
    if false {
        fmt.Println(v4)
        v5 := GenS1{F0: 2, F1: 1, F2: -1}
        _ = v5
    }
    var v6 float32 = (((0.1 - -2.25) + ((float32(0.0) * ((1e10 + (1.5 / (0.1 + (0.1 / 3.14159)))) - ((-2.25 / -2.25) + (0.1 - 1e10)))) * (float32(-2.25) / (1e10 * 3.14159)))) * -2.25)
    v6 -= float32(-2.25)
    var v7 uint16 = 42
    _ = v4
    _ = v6
    _ = v7
    return v7
}
func genFn1() uint {
    v8 := []int{740, -658, -42}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v9 int = -618
    if true {
        var v10 int16 = 0
        var v11 string = (("\n\t" + "foo bar") + "foo bar")
        v8 = append(v8, 128)
        _ = v10
        _ = v11
    } else {
        fmt.Println(v9)
        if false {
            v9 -= (v9 & (v9 | v9))
            for fi12 := 0; fi12 < 16; fi12++ {
                switch ("a" + ("hello" + ("hello" + ("" + "\n\t")))) {
                case "":
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v9 -= fi13
                        _ = fi13
                    }
                case "foo bar":
                    for fi14 := 0; fi14 < 12; fi14++ {
                        if true {
                            v15 := GenS2{F0: 554, F1: 2}
                            _ = v15
                        }
                        v9 -= 2147483647
                        fmt.Println(fi14, v9)
                        _ = fi14
                    }
                }
                if false {
                    v16 := GenS2{F0: 275, F1: 128}
                    _ = v16
                }
                _ = fi12
            }
        }
        v9 += v9
    }
    var v17 float32 = 1e10
    _ = v17
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int16 = 42
    v3 := make(map[string]string)
    for fi18 := 0; fi18 < 5; fi18++ {
        if false {
            fmt.Println(v1)
            v1 -= v1
            switch ("foo bar" + ("\n\t" + "hello")) {
            case "\n\t":
                fmt.Println(fi18, v1)
            case "a":
                var v19 int8 = -3
                _ = v19
            }
        }
        v2 = v2
        _ = fi18
    }
    var v20 int64 = -1
    v21 := GenS3{F0: false, F1: 32767}
    switch "" {
    case "":
        if (0 > 128) {
            fmt.Println(v20, v1, v1)
            for fi22 := 0; fi22 < 10; fi22++ {
                var v23 uint16 = 2
                for fi24 := 0; fi24 < 7; fi24++ {
                    v25 := GenS4{F0: fi22, F1: 1, F2: 42}
                    var v26 int32 = 0
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi22
                _ = v23
            }
            v27 := GenS4{F0: 42, F1: (v2 - v2), F2: 506}
            _ = v27
        } else {
            switch 2147483647 {
            case -435:
                switch "a" {
                case "":
                    if true {
                        var v28 int32 = 122
                        _ = v28
                    }
                }
            case 128:
                switch 127 {
                case 32768:
                    v3["foo bar"] = "foo bar"
                default:
                    var v29 int32 = 1
                    _ = v29
                }
            }
            var v30 float64 = 0.0
            _ = v30
        }
    case "foo bar":
        var v31 int32 = int32(v2)
        _ = v31
    case "\n\t":
        v32 := GenS5{F0: (v1 | (v1 >> (((v1 - (v1 - 42))) & 63))), F1: (v1 % ((18446744073709551615) | 1)), F2: "\n\t"}
        _ = v32
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        for fi34 := 0; fi34 < 9; fi34++ {
            v2 = (v2 * 42)
            var v35 int8 = 100
            v36 := GenS2{F0: 255, F1: 0}
            _ = fi34
            _ = v35
            _ = v36
        }
        v37 := v3[""]
        v38 := GenS6{F0: 127, F1: fi33}
        _ = fi33
        _ = v37
        _ = v38
    }
    var v39 uint64 = 18446744073709551615
    if ((!(((118 != 69) || false) && (-9223372036854775808 <= 9223372036854775807))) || ((-2.25 + (((-2.25 / ((float64(3.14159) / -2.25) - (float64((-2.25 / (1.5 + 1e10))) + float64(float32(3.14159))))) / (-2.25 + 0.0)) + 1.5)) != 0.0)) {
        var v40 float64 = float64(3.14159)
        _ = v40
    }
    fmt.Println(v2, v20)
    var v41 float64 = 1.5
    fmt.Println(v2)
    var v42 uint32 = 128
    if (-226 < -703) {
        for fi43 := 0; fi43 < 20; fi43++ {
            if (fi43 == fi43) {
                v39 += v39
            } else {
                for fi44 := 0; fi44 < 12; fi44++ {
                    v39 -= v39
                    var v45 bool = false
                    _ = fi44
                    _ = v45
                }
            }
            _ = fi43
        }
    }
    var v46 float32 = (0.0 * 1e10)
    var v47 uint32 = (v42 % ((4294967295) | 1))
    if false {
        if true {
            var v48 int = int(v1)
            for fi49 := 0; fi49 < 2; fi49++ {
                fmt.Println(v20, v46)
                switch (fi49 - (^(fi49))) {
                case 1:
                    v48 = (-(v48))
                default:
                    if true {
                        fmt.Println(v39, v42, v48)
                        v2 = v2
                    }
                }
                _ = fi49
            }
            _ = v48
        }
        var v50 float32 = v46
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v39
    _ = v41
    _ = v42
    _ = v46
    _ = v47
}
