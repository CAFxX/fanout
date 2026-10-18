package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint16; F2 string }

func main() {
    var v1 bool = (65535 <= 2147483648)
    var v2 bool = ((!false) || v1)
    var v3 int16 = 558
    v4 := GenS1{F0: v3, F1: 682}
    v3 = (v3 * (v3 ^ 158))
    switch "a" {
    case "foo bar":
        if v1 {
            switch "" {
            case "\n\t":
                fmt.Println(v2, v2)
            case "foo bar":
                var v5 uint32 = 1
                _ = v5
            }
            for fi6 := 0; fi6 < 8; fi6++ {
                for fi7 := 0; fi7 < 7; fi7++ {
                    if (!((("foo bar" + "\n\t") + "foo bar") != ("hello" + "a"))) {
                        if v2 {
                            var v8 float32 = 1.5
                            fmt.Println(fi6, v3)
                            _ = v8
                        }
                        v9 := GenS1{F0: v3, F1: 128}
                        _ = v9
                    }
                    if (-99 >= -557) {
                        switch "\n\t" {
                        case "a":
                            v10 := GenS2{F0: (2147483648 == 390)}
                            _ = v10
                        case "":
                            v11 := GenS2{F0: (("a" + "\n\t") < "foo bar")}
                            _ = v11
                        case "foo bar":
                            if ((!(!true)) && v2) {
                                v1 = ((fi7 - -1) < -196)
                                var v12 int16 = 42
                                var v13 uint64 = 820
                                _ = v12
                                _ = v13
                            }
                        }
                        switch -666 {
                        case 127:
                            v14 := []int{862, 237, fi7, (fi6 % ((fi7) | 1))}
                            _ = v14
                        case -20:
                            v15 := GenS2{F0: v2}
                            _ = v15
                        default:
                            v16 := []int{(fi6 / ((fi7) | 1)), 32768, fi6, 42}
                            fmt.Println(v2, v2)
                            _ = v16
                        }
                    }
                    var v17 float32 = ((3.14159 + 1.5) / float32(((float64((0.0 + ((float32(1e10) / 1e10) * 1e10))) * 0.1) + 0.1)))
                    v1 = (!(175 <= 959))
                    _ = fi7
                    _ = v17
                }
                _ = fi6
            }
        } else {
            if v2 {
                v18 := GenS3{F0: 42, F1: 1, F2: 17}
                if (!(!v2)) {
                    if v2 {
                        var v19 string = "a"
                        v20 := GenS3{F0: 51, F1: 127, F2: 1}
                        v21 := GenS3{F0: 18446744073709551615, F1: 73, F2: 171}
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    v1 = v1
                }
                _ = v18
            } else {
                var v22 string = "hello"
                var v23 float64 = ((0.1 - (1.5 * 0.1)) / 3.14159)
                _ = v22
                _ = v23
            }
            v2 = true
        }
    case "":
        v24 := GenS2{F0: v2}
        _ = v24
    default:
        for fi25 := 0; fi25 < 11; fi25++ {
            if (993 < 65535) {
                var v26 float32 = (1e10 * 3.14159)
                var v27 bool = (false && (true && (32767 > 0)))
                if v2 {
                    var v28 int = (-(fi25))
                    v29 := []int{int(v3)}
                    _ = v28
                    _ = v29
                }
                _ = v26
                _ = v27
            }
            v1 = (234 > 1)
            _ = fi25
        }
    }
    var v30 uint8 = 255
    v1 = (-255 < 42)
    v31 := GenS4{F0: 4294967295, F1: 32767, F2: "foo bar"}
    var v32 bool = (685 < 0)
    fmt.Println(v30, v30)
    v33 := GenS3{F0: 0, F1: -1, F2: 0}
    v1 = v2
    var v34 float64 = (-2.25 + float64((float32(-2.25) / 0.1)))
    for fi35 := 0; fi35 < 13; fi35++ {
        v36 := []int{(-(fi35)), fi35, 32767, fi35}
        var v37 int16 = v3
        v38 := GenS3{F0: 223, F1: -1, F2: (v30 & 84)}
        _ = fi35
        _ = v36
        _ = v37
        _ = v38
    }
    v39 := GenS2{F0: (!false)}
    var v40 bool = (!(!((int32(v3) > -1) || v2)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v40
}
