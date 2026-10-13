package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int8; F1 uint32; F2 uint64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 float64 = ((0.0 + 1e10) + (-2.25 - ((3.14159 * 1.5) - (3.14159 - 0.1))))
    var v4 int32 = -1
    var v5 int = -1
    fmt.Println(v5, v4)
    if (0 < 4294967295) {
        switch "\n\t" {
        case "":
            if true {
                v6 := []int{v5}
                _ = v6
            }
        default:
            for fi7 := 0; fi7 < 1; fi7++ {
                v4 -= v4
                _ = fi7
            }
        }
    } else {
        if false {
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := []int{}
                _ = fi8
                _ = v9
            }
            v3 -= 1e10
            var v10 float32 = (1e10 + 3.14159)
            _ = v10
        }
        v11 := v1["\n\t"]
        _ = v11
    }
    var v12 int16 = -32768
    v12 += v12
    if false {
        if false {
            var v13 uint = 0
            _ = v13
        }
        for fi14 := 0; fi14 < 3; fi14++ {
            switch v5 {
            case -868:
                var v15 int16 = v12
                _ = v15
            case 986:
                fmt.Println(v3, v5, fi14)
            }
            v16 := []int{(fi14 * fi14), fi14}
            fmt.Println(v3, v4, fi14)
            _ = fi14
            _ = v16
        }
        var v17 int8 = -63
        _ = v17
    }
    fmt.Println(v12, v4)
    v18 := []int{v5, (v5 >> ((255) & 63)), v5}
    var v19 float32 = 0.1
    v20 := GenS1{F0: (v12 << ((18446744073709551615) & 63))}
    if true {
        _ = len(v2)
    } else {
        if ((75 >= -128) || true) {
            v21 := GenS1{F0: v12}
            var v22 bool = true
            v23 := GenS1{F0: v12}
            _ = v21
            _ = v22
            _ = v23
        }
    }
    v24 := GenS1{F0: v12}
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := GenS1{F0: (v12 + 0)}
        _ = fi25
        _ = v26
    }
    fmt.Println(v19, v12)
    var v27 int64 = -439
    for fi28 := 0; fi28 < 6; fi28++ {
        if true {
            for fi29 := 0; fi29 < 10; fi29++ {
                var v30 int64 = (^(v27))
                switch "\n\t" {
                case "foo bar":
                    if true {
                        var v31 uint16 = 0
                        v12 = v12
                        _ = v31
                    } else {
                        v4 = v4
                        v30 += (-(v27))
                    }
                case "\n\t":
                    v32 := GenS1{F0: v12}
                    _ = v32
                }
                _ = fi29
                _ = v30
            }
            for fi33 := 0; fi33 < 4; fi33++ {
                v3 += 1e10
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                _ = fi33
            }
            for fi34 := 0; fi34 < 1; fi34++ {
                var v35 int64 = v27
                v5 -= (v5 >> ((1) & 63))
                fmt.Println(v27, v19)
                _ = fi34
                _ = v35
            }
        } else {
            if true {
                for fi36 := 0; fi36 < 13; fi36++ {
                    if ((fi28 - 9223372036854775807) >= (v5 + v5)) {
                        v3 = 3.14159
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                    } else {
                        var v37 uint32 = 2147483648
                        _ = v37
                    }
                    switch "hello" {
                    case "":
                        if true {
                            v38 := GenS2{F0: int8(v27)}
                            _ = v38
                        }
                    case "foo bar":
                        var v39 float64 = (-2.25 + float64(v19))
                        _ = v39
                    default:
                        if true {
                            v3 = v3
                            v3 -= 0.0
                        } else {
                            v5 += 0
                            if true {
                                switch (fi28 + -9223372036854775808) {
                                case -1:
                                    v3 = 0.0
                                default:
                                    delete(v2, "a")
                                }
                                var v40 uint64 = 2
                                _ = v40
                            }
                        }
                    }
                    switch 0 {
                    case 4294967295:
                        var v41 int8 = 100
                        _ = v41
                    default:
                        var v42 int32 = v4
                        _ = v42
                    }
                    _ = fi36
                }
                var v43 int = 2147483647
                if false {
                    fmt.Println(v27, v5, v12)
                } else {
                    for fi44 := 0; fi44 < 16; fi44++ {
                        v3 -= (v3 * (v3 + 0.0))
                        v4 = v4
                        v45 := GenS3{F0: -1}
                        if false {
                            v3 = v3
                            v43 = fi28
                        } else {
                            delete(v1, "\n\t")
                            var v46 uint = 4294967295
                            _ = v46
                        }
                        _ = fi44
                        _ = v45
                    }
                    v2["a"] = false
                }
                _ = v43
            } else {
                if false {
                    v47 := GenS3{F0: -1}
                    delete(v1, "hello")
                    v48 := GenS4{F0: ("" + "a"), F1: "hello"}
                    _ = v47
                    _ = v48
                }
            }
            switch (^(fi28)) {
            case 32768:
                for fi49 := 0; fi49 < 13; fi49++ {
                    _ = len(v2)
                    v50 := v1["\n\t"]
                    _ = fi49
                    _ = v50
                }
            case -2147483648:
                fmt.Println(v27, v4, v12)
            case 32767:
                fmt.Println(v19, v27, v19)
            default:
                if false {
                    var v51 int16 = -905
                    fmt.Println(v27, v5)
                    v4 = (v4 | (v4 | (v4 + v4)))
                    _ = v51
                } else {
                    v52 := GenS5{F0: 100, F1: 0, F2: 0}
                    _ = v52
                }
            }
            switch (("" + ("" + "\n\t")) + "a") {
            case "hello":
                fmt.Println(v12, v3)
            case "a":
                var v53 uint64 = 350
                _ = v53
            default:
                v12 += v12
            }
        }
        v4 -= v4
        fmt.Println(v12)
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
