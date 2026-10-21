package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int; F2 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 42
    v3 := []string{}
    v4 := make(map[string]string)
    v5 := make(map[int]uint32)
    var v6 float32 = (1.5 / float32(3.14159))
    if false {
        _ = len(v4)
        if (!(32767 >= -133)) {
            fmt.Println(v2, v2)
            v6 -= float32(3.14159)
        }
    } else {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint64 = 42
            fmt.Println(v2, v6)
            _ = fi7
            _ = v8
        }
        switch 834 {
        case 1:
            v2 = v2
        default:
            if ((!(v2 <= v2)) || ("" >= ("foo bar" + "foo bar"))) {
                v2 = (v2 * v2)
                for fi9 := 0; fi9 < 1; fi9++ {
                    for fi10 := 0; fi10 < 15; fi10++ {
                        fmt.Println(v6, v6, v6)
                        fmt.Println(fi10, v2)
                        _ = fi10
                    }
                    var v11 int = fi9
                    _ = fi9
                    _ = v11
                }
                if (true && true) {
                    switch "a" {
                    case "foo bar":
                        var v12 float64 = 1.5
                        _ = v12
                    case "a":
                        switch -1 {
                        case 9223372036854775807:
                            var v13 int16 = 585
                            _ = v13
                        case 4294967295:
                            v6 -= 1e10
                        case 32767:
                            _ = len(v5)
                        }
                    case "\n\t":
                        _ = len(v3)
                    }
                }
            }
        }
    }
    var v14 uint16 = 65535
    if (1e10 == ((-2.25 / -2.25) - (3.14159 / ((3.14159 - 1.5) * 1e10)))) {
        v2 = v2
    }
    v14 -= 253
    fmt.Println(v6, v2, v6)
    delete(v1, 4294967295)
    if (1 < 505) {
        v15 := GenS1{F0: (v2 != (v2 / ((127) | 1)))}
        v16 := GenS1{F0: false}
        _ = v15
        _ = v16
    } else {
        if false {
            switch 9223372036854775807 {
            case 42:
                fmt.Println(v14)
            case 32768:
                fmt.Println(v14, v2, v14)
            case 127:
                v2 -= v2
            default:
                v17 := GenS1{F0: false}
                _ = v17
            }
            if (487 <= 805) {
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 uint32 = uint32(v14)
                    if true {
                        for fi20 := 0; fi20 < 13; fi20++ {
                            var v21 float64 = 3.14159
                            v6 += v6
                            _ = fi20
                            _ = v21
                        }
                        if (18446744073709551615 >= 1) {
                            v22 := GenS2{F0: 0, F1: 0}
                            for fi23 := 0; fi23 < 6; fi23++ {
                                v24 := v5[(^(fi23))]
                                v14 -= v14
                                _ = fi23
                                _ = v24
                            }
                            _ = v22
                        } else {
                            var v25 uint = 4294967295
                            for fi26 := 0; fi26 < 6; fi26++ {
                                var v27 float32 = (v6 + ((1.5 - v6) + ((((1e10 * -2.25) + ((v6 * v6) / v6)) - float32(((((3.14159 + 1e10) - 1e10) * (((1.5 * 0.1) / ((float64(v6) * (1.5 - 0.0)) / float64(v6))) / 1e10)) + 3.14159))) * 1.5)))
                                _ = fi26
                                _ = v27
                            }
                            switch "hello" {
                            case "":
                                _ = len(v3)
                            case "a":
                                var v28 string = "\n\t"
                                _ = v28
                            default:
                                var v29 bool = false
                                _ = v29
                            }
                            _ = v25
                        }
                    } else {
                        var v30 int8 = -13
                        _ = v30
                    }
                    v6 = 1e10
                    _ = fi18
                    _ = v19
                }
                if (1 == -1) {
                    _ = len(v3)
                    v31 := GenS3{F0: 429}
                    _ = v31
                } else {
                    var v32 uint64 = 188
                    var v33 int16 = -86
                    _ = v32
                    _ = v33
                }
                v14 += 65535
            }
        }
        v34 := GenS2{F0: 1000, F1: 42}
        _ = v34
    }
    var v35 uint32 = 2147483647
    for fi36 := 0; fi36 < 14; fi36++ {
        var v37 uint64 = 2
        _ = fi36
        _ = v37
    }
    v38 := GenS3{F0: 2}
    switch (("" + "a") + "") {
    case "hello":
        var v39 int32 = 2147483647
        _ = v39
    default:
        v6 = v6
    }
    if true {
        var v40 int16 = -605
        v41 := GenS4{F0: -1, F1: 1, F2: ("hello" + "foo bar")}
        _ = v40
        _ = v41
    }
    v42 := GenS2{F0: 42, F1: -818}
    var v43 float32 = v6
    var v44 bool = false
    if (v35 != (v35 | v35)) {
        v14 -= (v14 % ((v14) | 1))
        var v45 int32 = 197
        v3 = append(v3, ("a" + ""))
        _ = v45
    } else {
        for fi46 := 0; fi46 < 6; fi46++ {
            v35 += v35
            fmt.Println(v35, v14)
            var v47 int16 = 1000
            _ = fi46
            _ = v47
        }
    }
    var v48 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v48
    _ = v5
    _ = v6
}
