package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int32; F1 int16; F2 int64 }

func main() {
    v1 := []string{}
    var v2 float64 = float64(1e10)
    var v3 uint8 = 0
    v4 := []uint32{2147483647, 444}
    v5 := make(map[uint32]string)
    v3 -= (v3 - (^(v3)))
    v3 = 1
    v6 := GenS1{F0: 32, F1: 1}
    for fi7 := 0; fi7 < 7; fi7++ {
        v2 += v2
        v8 := GenS1{F0: -959, F1: -128}
        fmt.Println(v3, v3, v3)
        _ = fi7
        _ = v8
    }
    if true {
        v9 := GenS1{F0: 42, F1: 0}
        v5[949] = ""
        for fi10 := 0; fi10 < 4; fi10++ {
            if (338 == 0) {
                var v11 int32 = 459
                _ = v11
            }
            if ((0.0 / 1e10) != ((1.5 * (1e10 * 3.14159)) - 1.5)) {
                var v12 uint16 = 32767
                var v13 int = (^(fi10))
                v3 = (v3 | v3)
                _ = v12
                _ = v13
            } else {
                fmt.Println(fi10)
            }
            _ = fi10
        }
        _ = v9
    } else {
        switch "hello" {
        case "hello":
            switch -512 {
            case 979:
                v14 := GenS1{F0: 0, F1: 0}
                _ = v14
            case 32768:
                fmt.Println(v3, v3)
            case 127:
                switch ("foo bar" + "a") {
                case "foo bar":
                    switch 69 {
                    case -2147483648:
                        switch "" {
                        case "a":
                            v15 := v5[816]
                            _ = v15
                        case "":
                            var v16 int16 = 1
                            _ = v16
                        case "foo bar":
                            v17 := GenS2{F0: false}
                            _ = v17
                        default:
                            fmt.Println(v2)
                        }
                    default:
                        v2 = 0.1
                    }
                }
            }
        case "\n\t":
            var v18 uint64 = 2
            _ = v18
        }
    }
    var v19 int64 = 42
    delete(v5, 322)
    for fi20 := 0; fi20 < 6; fi20++ {
        if false {
            if (-1 < 1) {
                v19 = v19
            } else {
                v21 := GenS3{F0: v19, F1: 108}
                _ = v21
            }
        }
        var v22 int16 = 0
        _ = fi20
        _ = v22
    }
    v23 := v5[2147483647]
    for fi24 := 0; fi24 < 20; fi24++ {
        v2 = v2
        _ = fi24
    }
    fmt.Println(v23)
    v19 += v19
    fmt.Println(v19, v2)
    for fi25 := 0; fi25 < 2; fi25++ {
        v4 = append(v4, 2147483647)
        var v26 int16 = 0
        var v27 int16 = 1000
        _ = fi25
        _ = v26
        _ = v27
    }
    if false {
        if true {
            if (!(false || false)) {
                var v28 uint = 42
                if false {
                    fmt.Println(v19, v3)
                }
                _ = v28
            } else {
                for fi29 := 0; fi29 < 13; fi29++ {
                    v2 = v2
                    v3 -= 255
                    if false {
                        v30 := GenS1{F0: v19, F1: int8(fi29)}
                        var v31 float64 = float64(float32(v2))
                        _ = v30
                        _ = v31
                    } else {
                        fmt.Println(v19, fi29)
                        if ((v19 / (((-(v19))) | 1)) < v19) {
                            v2 += (0.1 + 1e10)
                        } else {
                            v23 = ((((v23 + ("foo bar" + "\n\t")) + ("hello" + v23)) + v23) + v23)
                            v32 := GenS4{F0: 891, F1: 0, F2: int64(fi29)}
                            _ = v32
                        }
                    }
                    _ = fi29
                }
            }
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
            var v33 int16 = -1
            _ = v33
        }
    } else {
        if (100 == 42) {
            if (!false) {
                if true {
                    var v34 int16 = 1
                    _ = v34
                }
            } else {
                v23 = ((((("\n\t" + ("foo bar" + v23)) + v23) + v23) + ((v23 + (v23 + (v23 + v23))) + ("foo bar" + ("" + v23)))) + v23)
            }
        }
        if (-280 != v19) {
            if false {
                var v35 uint64 = 4294967295
                _ = v35
            } else {
                var v36 float32 = 0.0
                var v37 uint = 42
                _ = v36
                _ = v37
            }
        }
    }
    var v38 uint16 = 1
    for fi39 := 0; fi39 < 16; fi39++ {
        fmt.Println(v2)
        var v40 bool = true
        _ = fi39
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
