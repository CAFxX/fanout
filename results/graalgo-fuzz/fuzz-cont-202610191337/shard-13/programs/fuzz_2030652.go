package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 bool }
func genFn0(p0 int, p1 uint, p2 int64) string {
    var v5 uint16 = 849
    v6 := []int{}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint8 = 54
        v5 -= 1
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return "\n\t"
}

func main() {
    v1 := []uint32{4294967295, 2147483648, 800, 4294967295}
    var v2 int64 = -486
    v3 := []int{-467, 32768, -9223372036854775808, -2147483648}
    var v4 string = "foo bar"
    if false {
        v2 -= (v2 / (((v2 / ((v2) | 1))) | 1))
        if true {
            v2 -= -386
        } else {
            if (!true) {
                switch "foo bar" {
                case "a":
                    fmt.Println(v2)
                case "":
                    for fi9 := 0; fi9 < 3; fi9++ {
                        var v10 float64 = 0.1
                        var v11 bool = (!false)
                        fmt.Println(v11, v10)
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                case "hello":
                    v12 := GenS1{F0: 0, F1: 643, F2: (109 == 127)}
                    _ = v12
                }
                var v13 int8 = -1
                if (42 >= 625) {
                    for fi14 := 0; fi14 < 16; fi14++ {
                        v15 := GenS1{F0: 1, F1: 42, F2: (-1 == -1)}
                        _ = fi14
                        _ = v15
                    }
                    v2 = (-(v2))
                }
                _ = v13
            }
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v4 = v4
            }
            var v16 float64 = float64(1e10)
            _ = v16
        }
        _ = len(v3)
    } else {
        var v17 uint64 = 18446744073709551615
        for fi18 := 0; fi18 < 9; fi18++ {
            if (!(true && false)) {
                v2 += v2
                v19 := GenS1{F0: -1, F1: fi18, F2: false}
                _ = v19
            } else {
                v20 := GenS1{F0: 0, F1: -2147483648, F2: (!true)}
                for fi21 := 0; fi21 < 13; fi21++ {
                    if true {
                        switch 80 {
                        case 32768:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 766:
                            var v22 float32 = 0.0
                            _ = v22
                        case 645:
                            if true {
                                var v23 uint8 = 1
                                v2 += (v2 % ((1) | 1))
                                var v24 int16 = 42
                                _ = v23
                                _ = v24
                            } else {
                                _ = len(v1)
                                if (!(0 == 149)) {
                                    var v25 uint = 706
                                    if false {
                                        fmt.Println(v25, fi21, v25)
                                    } else {
                                        switch (fi21 % ((617) | 1)) {
                                        case 9223372036854775807:
                                            switch (fi18 ^ 731) {
                                            case 1:
                                                var v26 uint = v25
                                                _ = v26
                                            case 32767:
                                                var v27 int64 = (-(v2))
                                                _ = v27
                                            }
                                        case 584:
                                            v25 = (v25 / (((v25 << ((65535) & 63))) | 1))
                                        }
                                    }
                                    var v28 uint = (v25 % ((940) | 1))
                                    var v29 uint64 = v17
                                    _ = v25
                                    _ = v28
                                    _ = v29
                                }
                                v30 := GenS1{F0: -2147483648, F1: -828, F2: (!(true || true))}
                                v4 = ((v4 + (v4 + v4)) + ((v4 + v4) + "foo bar"))
                                _ = v30
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v31 int = (-(fi21))
                        _ = v31
                    } else {
                        var v32 uint8 = 115
                        _ = v32
                    }
                    v4 = v4
                    v17 -= 4294967295
                    _ = fi21
                }
                _ = v20
            }
            var v33 int = (fi18 & 4294967295)
            _ = fi18
            _ = v33
        }
        _ = v17
    }
    if true {
        v2 += -516
    }
    for fi34 := 0; fi34 < 3; fi34++ {
        var v35 int64 = -1
        var v36 int32 = 85
        if true {
            v35 = v2
        } else {
            v35 = 9223372036854775807
            v2 = v35
        }
        _ = fi34
        _ = v35
        _ = v36
    }
    v4 = "\n\t"
    v2 = (v2 << ((4294967295) & 63))
    switch v4 {
    case "a":
        var v37 string = "\n\t"
        _ = v37
    case "foo bar":
        var v38 bool = (1 >= 0)
        _ = v38
    case "":
        v39 := GenS1{F0: 42, F1: int(v2), F2: false}
        _ = v39
    }
    var v40 uint8 = 0
    var v41 string = (v4 + (v4 + "hello"))
    var v42 uint64 = 4294967295
    if false {
        var v43 int16 = 0
        _ = v43
    } else {
        v42 = v42
        v44 := GenS1{F0: -1, F1: 42, F2: ((!false) || (!true))}
        _ = v44
    }
    var v45 uint16 = 143
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v45
}
