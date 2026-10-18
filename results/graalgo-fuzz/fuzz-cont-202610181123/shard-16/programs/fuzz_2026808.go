package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint8; F1 int; F2 int8 }
type GenS3 struct { F0 string; F1 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 1
    var v3 int16 = 1
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]bool)
    v6 := []int{-427, 42}
    v4[2] = "hello"
    switch 127 {
    case 32767:
        fmt.Println(v3)
    case -457:
        v2 += v2
    default:
        if (float64(-2.25) <= float64(-2.25)) {
            switch "a" {
            case "foo bar":
                v2 = v2
            case "a":
                if true {
                    for fi7 := 0; fi7 < 6; fi7++ {
                        v8 := GenS1{F0: 1, F1: 1}
                        var v9 uint = 2
                        for fi10 := 0; fi10 < 16; fi10++ {
                            var v11 int16 = 1
                            var v12 int16 = v11
                            var v13 float64 = 1e10
                            _ = fi10
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    for fi14 := 0; fi14 < 19; fi14++ {
                        v3 -= 413
                        _ = fi14
                    }
                }
            case "hello":
                switch "hello" {
                case "\n\t":
                    v3 += (v3 + v3)
                case "hello":
                    switch 1 {
                    case 687:
                        v15 := GenS1{F0: 843, F1: v3}
                        _ = v15
                    case 42:
                        v3 = -1
                    default:
                        var v16 int8 = 1
                        _ = v16
                    }
                }
            }
            v17 := GenS2{F0: 1, F1: 127, F2: -1}
            _ = v17
        } else {
            v2 -= (^(v2))
        }
    }
    if false {
        for fi18 := 0; fi18 < 8; fi18++ {
            for fi19 := 0; fi19 < 9; fi19++ {
                v3 += 189
                _ = fi19
            }
            _ = fi18
        }
        for fi20 := 0; fi20 < 5; fi20++ {
            if (1 == 74) {
                _ = len(v6)
                v3 += -102
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v21 int32 = -1
            _ = fi20
            _ = v21
        }
    }
    v3 = v3
    v3 = v3
    v2 -= 2
    v1[65535] = true
    v1[2147483648] = false
    var v22 uint8 = 93
    fmt.Println(v3, v2)
    var v23 int16 = 1000
    var v24 int = 128
    if (0.1 >= ((3.14159 - 3.14159) / 0.1)) {
        for fi25 := 0; fi25 < 6; fi25++ {
            for fi26 := 0; fi26 < 8; fi26++ {
                var v27 bool = (65535 != 881)
                if v27 {
                    v2 -= v2
                    var v28 string = "a"
                    v29 := GenS2{F0: 255, F1: (-(fi25)), F2: -61}
                    _ = v28
                    _ = v29
                } else {
                    var v30 int32 = -2147483648
                    var v31 uint = 474
                    var v32 int8 = 100
                    _ = v30
                    _ = v31
                    _ = v32
                }
                v23 = v23
                _ = fi26
                _ = v27
            }
            var v33 int8 = 100
            _ = fi25
            _ = v33
        }
        v34 := GenS1{F0: 2147483647, F1: v3}
        _ = v34
    } else {
        for fi35 := 0; fi35 < 8; fi35++ {
            if (v22 == v22) {
                var v36 int8 = 100
                v37 := GenS1{F0: 770, F1: v3}
                v38 := GenS1{F0: 2147483647, F1: v23}
                _ = v36
                _ = v37
                _ = v38
            } else {
                fmt.Println(v3, v23)
            }
            if (!true) {
                if len(v6) > 0 {
                    v6[0] = (fi35 + v24)
                }
                v22 = (v22 & 255)
                fmt.Println(v24, v22)
            } else {
                var v39 uint64 = 4294967295
                _ = v39
            }
            _ = fi35
        }
        v40 := GenS3{F0: (("\n\t" + "hello") + ""), F1: true}
        var v41 uint8 = (v22 - (v22 ^ v22))
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        if false {
            v22 = 1
            v43 := GenS3{F0: "", F1: (true || true)}
            _ = v43
        }
        var v44 int64 = int64(fi42)
        _ = fi42
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
