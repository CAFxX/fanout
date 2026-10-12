package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 uint32 }
func genFn0() uint16 {
    v7 := []int{127, -908}
    var v8 uint16 = 255
    v9 := GenS1{F0: 255, F1: 231, F2: 2}
    _ = v7
    _ = v8
    _ = v9
    return (v8 | 2)
}

func main() {
    v1 := []uint32{}
    var v2 int8 = 127
    v3 := make(map[uint32]bool)
    var v4 string = ("foo bar" + ("a" + "a"))
    var v5 uint16 = 0
    var v6 float64 = 0.1
    var v10 float32 = float32(v6)
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 19; fi12++ {
            v1 = append(v1, 424)
            var v13 bool = true
            _ = fi12
            _ = v13
        }
        v14 := GenS2{F0: 287, F1: 1000, F2: 42}
        var v15 uint8 = 0
        _ = fi11
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 42, F1: 845, F2: 127}
    var v17 bool = (-2147483648 <= 2147483647)
    v18 := GenS1{F0: 2, F1: 1, F2: 0}
    if v17 {
        var v19 int64 = 1
        var v20 uint = 2
        v19 = v19
        _ = v19
        _ = v20
    } else {
        if (0 >= 285) {
            fmt.Println(v10)
            var v21 int32 = -2147483648
            _ = v21
        } else {
            for fi22 := 0; fi22 < 4; fi22++ {
                v10 = (v10 / (float32(float64(float32(v6))) - v10))
                if (fi22 >= -9223372036854775808) {
                    var v23 uint8 = 147
                    var v24 string = "\n\t"
                    _ = v23
                    _ = v24
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v25 bool = (1 >= 459)
                    _ = v25
                }
                var v26 uint64 = 1
                _ = fi22
                _ = v26
            }
            var v27 int8 = (v2 % ((-74) | 1))
            fmt.Println(v5)
            _ = v27
        }
        if (((v10 + (v10 - v10)) / (v10 + float32(v6))) == v10) {
            if v17 {
                for fi28 := 0; fi28 < 7; fi28++ {
                    var v29 uint8 = 0
                    v5 = (v5 + v5)
                    _ = fi28
                    _ = v29
                }
                switch v4 {
                case "a":
                    v30 := GenS2{F0: uint32(v5), F1: 0, F2: 785}
                    _ = v30
                case "hello":
                    v2 -= 5
                case "foo bar":
                    var v31 uint32 = 947
                    _ = v31
                }
                if v17 {
                    switch "a" {
                    case "foo bar":
                        v3[2147483648] = (true && (!((^(v2)) < v2)))
                    }
                    var v32 int16 = 1
                    if (v17 || false) {
                        fmt.Println(v5, v10, v6)
                        var v33 uint32 = 2
                        _ = v33
                    } else {
                        if (v17 || v17) {
                            fmt.Println(v32, v10)
                        } else {
                            v34 := v3[514]
                            _ = v34
                        }
                        v5 -= genFn0()
                    }
                    _ = v32
                }
            }
        } else {
            for fi35 := 0; fi35 < 20; fi35++ {
                v36 := GenS3{F0: 1, F1: -365}
                var v37 uint32 = 2
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        switch 128 {
        case 185:
            v38 := GenS2{F0: 349, F1: 1, F2: 2147483647}
            _ = v38
        default:
            v39 := GenS4{F0: ((v4 + "\n\t") + ("\n\t" + "hello")), F1: 18446744073709551615}
            _ = v39
        }
    }
    v10 -= (3.14159 - ((v10 / v10) + float32((0.1 - (v6 - (float64((v10 - 3.14159)) * 3.14159))))))
    fmt.Println(v5, v4)
    if (1 >= 198) {
        if v17 {
            for fi40 := 0; fi40 < 19; fi40++ {
                var v41 int = (fi40 / ((127) | 1))
                v1 = append(v1, 4294967295)
                if (9223372036854775807 >= int64(v5)) {
                    if (156 < 65535) {
                        var v42 uint64 = 4294967295
                        _ = v42
                    }
                    var v43 string = (("" + v4) + v4)
                    _ = v43
                }
                _ = fi40
                _ = v41
            }
            var v44 uint16 = genFn0()
            _ = v44
        } else {
            fmt.Println(v5, v17, v4)
            switch v4 {
            case "a":
                var v45 float32 = float32(1.5)
                _ = v45
            case "\n\t":
                v2 = v2
            case "hello":
                var v46 uint16 = genFn0()
                _ = v46
            default:
                v2 = (v2 << ((42) & 63))
            }
        }
        _ = len(v3)
    } else {
        if (true || v17) {
            if v17 {
                if len(v1) > 0 {
                    v1[0] = 2
                }
                var v47 uint64 = 18446744073709551615
                _ = v47
            }
            if (!true) {
                switch ((((v4 + "foo bar") + (("a" + (v4 + ("hello" + v4))) + (("\n\t" + "foo bar") + v4))) + v4) + v4) {
                case "hello":
                    var v48 int = 671
                    _ = v48
                default:
                    fmt.Println(v2, v17)
                }
                var v49 uint64 = 1
                _ = v49
            } else {
                var v50 uint8 = 2
                v1 = append(v1, 0)
                var v51 uint16 = (v5 * (^(v5)))
                _ = v50
                _ = v51
            }
        } else {
            v52 := GenS5{F0: -1}
            _ = v52
        }
        fmt.Println(v6, v4)
    }
    for fi53 := 0; fi53 < 14; fi53++ {
        v54 := GenS6{F0: 65535, F1: 225}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            for fi55 := 0; fi55 < 4; fi55++ {
                var v56 string = v4
                v17 = v17
                _ = fi55
                _ = v56
            }
            var v57 int8 = v2
            _ = v57
        } else {
            for fi58 := 0; fi58 < 18; fi58++ {
                var v59 uint16 = (v5 << ((936) & 63))
                _ = fi58
                _ = v59
            }
            var v60 float64 = ((0.1 - (float64(0.0) / v6)) + ((float64((v10 - v10)) / 3.14159) * v6))
            _ = v60
        }
        _ = fi53
        _ = v54
    }
    v61 := GenS6{F0: 857, F1: 2147483647}
    v2 = 1
    if (!v17) {
        fmt.Println(v5, v4)
        var v62 float64 = (0.1 - v6)
        _ = v62
    }
    v63 := GenS3{F0: 461, F1: -366}
    fmt.Println(v4, v2, v6)
    for fi64 := 0; fi64 < 6; fi64++ {
        v65 := GenS3{F0: 18446744073709551615, F1: -292}
        _ = fi64
        _ = v65
    }
    for fi66 := 0; fi66 < 6; fi66++ {
        v1 = append(v1, 0)
        v10 += v10
        if (0 < 255) {
            v17 = true
        }
        _ = fi66
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v63
}
