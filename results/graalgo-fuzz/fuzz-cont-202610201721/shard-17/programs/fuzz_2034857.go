package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS2 struct { F0 int16; F1 int }

func main() {
    var v1 uint8 = 0
    var v2 int64 = 1
    v3 := []string{"\n\t", "hello", "\n\t"}
    var v4 int = 2147483647
    var v5 float64 = ((-2.25 + float64(3.14159)) / ((1.5 * ((1.5 / 0.1) - 0.1)) - 0.0))
    v6 := GenS1{F0: (^(v1)), F1: 1, F2: 88}
    switch "foo bar" {
    case "":
        switch -2147483648 {
        case 32767:
            v7 := GenS1{F0: 211, F1: 1, F2: -128}
            _ = v7
        default:
            v3 = append(v3, "a")
        }
    case "hello":
        if true {
            switch "hello" {
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        } else {
            for fi8 := 0; fi8 < 4; fi8++ {
                if false {
                    var v9 bool = false
                    _ = v9
                }
                fmt.Println(v1, fi8, v2)
                _ = fi8
            }
            if false {
                if true {
                    fmt.Println(v5, v1, v2)
                    for fi10 := 0; fi10 < 4; fi10++ {
                        v3 = append(v3, "")
                        _ = fi10
                    }
                    if (true && false) {
                        fmt.Println(v5)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        v1 -= (^(v1))
                        var v11 float64 = 1e10
                        v5 = (float64(float32((v5 * v5))) * ((v5 * ((float64(0.0) + v5) / v5)) / float64(-2.25)))
                        _ = v11
                    }
                }
                var v12 float64 = v5
                _ = v12
            }
        }
    }
    _ = len(v3)
    var v13 uint16 = 65535
    fmt.Println(v1, v1, v4)
    switch v4 {
    case 4294967295:
        for fi14 := 0; fi14 < 16; fi14++ {
            if false {
                var v15 uint16 = (v13 | v13)
                _ = v15
            } else {
                var v16 int8 = -10
                v3 = append(v3, "\n\t")
                if (!true) {
                    var v17 uint32 = 2147483648
                    _ = v17
                } else {
                    if false {
                        if false {
                            v18 := GenS2{F0: -1, F1: fi14}
                            v5 -= 1.5
                            var v19 int8 = v16
                            _ = v18
                            _ = v19
                        } else {
                            var v20 float32 = 1e10
                            _ = v20
                        }
                    } else {
                        v1 -= (^(v1))
                        var v21 int8 = v16
                        var v22 int32 = 1
                        _ = v21
                        _ = v22
                    }
                }
                _ = v16
            }
            var v23 uint32 = 45
            _ = fi14
            _ = v23
        }
    case 356:
        switch "foo bar" {
        case "\n\t":
            var v24 uint64 = 0
            _ = v24
        default:
            if true {
                fmt.Println(v1)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    case 1:
        var v25 float64 = -2.25
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        fmt.Println(fi26, fi26)
        var v27 int8 = 42
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 int16 = 0
        _ = fi28
        _ = v29
    }
    v30 := GenS1{F0: 0, F1: v13, F2: -1}
    if false {
        var v31 uint64 = 915
        v4 -= v4
        var v32 uint32 = 108
        _ = v31
        _ = v32
    } else {
        v13 += v13
        v2 = (v2 >> ((18446744073709551615) & 63))
        var v33 uint = 18446744073709551615
        _ = v33
    }
    if true {
        if false {
            for fi34 := 0; fi34 < 18; fi34++ {
                if false {
                    for fi35 := 0; fi35 < 6; fi35++ {
                        var v36 bool = (v13 != (v13 / ((v13) | 1)))
                        _ = fi35
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 3; fi37++ {
                        var v38 uint64 = 45
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi37
                        _ = v38
                    }
                    switch (v4 % ((2147483647) | 1)) {
                    case -234:
                        var v39 bool = (2147483647 < 2147483648)
                        _ = v39
                    case 437:
                        var v40 uint8 = v1
                        _ = v40
                    default:
                        switch (v4 - 127) {
                        case 32767:
                            if false {
                                v13 -= v13
                                var v41 int64 = (v2 << ((1) & 63))
                                _ = v41
                            }
                        case -2147483648:
                            v42 := GenS1{F0: (v1 + v1), F1: v13, F2: -97}
                            _ = v42
                        }
                    }
                }
                var v43 uint = 255
                _ = fi34
                _ = v43
            }
        }
        var v44 uint8 = v1
        _ = v44
    } else {
        v3 = append(v3, "\n\t")
        if (42 >= 480) {
            v45 := GenS1{F0: 255, F1: v13, F2: 100}
            v2 = (v2 % (((-(v2))) | 1))
            _ = v45
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
