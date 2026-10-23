package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 bool }
func genFn0(p0 int64, p1 int16) uint8 {
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int32 = 42
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, p1)
    if true {
        p1 -= p1
        v7 := []int{950, 42, 2147483647}
        _ = v7
    } else {
        if false {
            v8 := GenS1{F0: 32767, F1: int16(p0), F2: 929}
            p1 = (p1 / ((p1) | 1))
            _ = v8
        }
        p0 += -9223372036854775808
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 string = (("\n\t" + "foo bar") + "foo bar")
            fmt.Println(v10, p1, v10)
            var v11 uint = 2
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 float64 = ((float64(((float32(float64(1.5)) / (3.14159 * (0.1 / 3.14159))) / float32((3.14159 + (-2.25 / 1.0))))) * 0.1) - float64(0.1))
    switch 128 {
    case -2147483648:
        var v13 int8 = 0
        _ = v13
    default:
        v14 := []int{1, 42}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    v1 := []string{"", "", ("a" + ("" + ((("a" + "") + "foo bar") + "\n\t")))}
    v2 := make(map[int]bool)
    var v3 uint8 = 128
    v4 := []uint32{886, 829}
    var v15 bool = true
    fmt.Println(v3)
    if v15 {
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := GenS1{F0: 0, F1: 910, F2: 4294967295}
            _ = fi16
            _ = v17
        }
        switch ("hello" + "hello") {
        case "a":
            if (v15 || (0 >= 1)) {
                for fi18 := 0; fi18 < 19; fi18++ {
                    if (0.1 < 1.5) {
                        v19 := v2[fi18]
                        switch "\n\t" {
                        case "a":
                            v15 = false
                        case "":
                            v15 = ((v3 % ((v3) | 1)) == genFn0(0, 1))
                        case "\n\t":
                            v19 = v19
                        default:
                            fmt.Println(v19, v15, fi18)
                        }
                        for fi20 := 0; fi20 < 7; fi20++ {
                            switch fi18 {
                            case -521:
                                if (!((572 < 4294967295) || false)) {
                                    v19 = true
                                }
                            case 32767:
                                v15 = true
                            case 769:
                                var v21 uint32 = 2147483647
                                _ = v21
                            default:
                                var v22 float64 = float64(0.1)
                                _ = v22
                            }
                            _ = fi20
                        }
                        _ = v19
                    } else {
                        var v23 int32 = -2147483648
                        _ = v23
                    }
                    _ = fi18
                }
                v24 := GenS1{F0: 32767, F1: 42, F2: 65535}
                v25 := v2[32768]
                _ = v24
                _ = v25
            } else {
                for fi26 := 0; fi26 < 14; fi26++ {
                    v15 = (1 < 4294967295)
                    v27 := GenS2{F0: 32767, F1: 4294967295}
                    _ = fi26
                    _ = v27
                }
                for fi28 := 0; fi28 < 2; fi28++ {
                    for fi29 := 0; fi29 < 2; fi29++ {
                        v15 = ((42 < 1) || (v15 || v15))
                        var v30 float64 = 0.1
                        _ = fi29
                        _ = v30
                    }
                    v31 := GenS2{F0: int16(fi28), F1: 2}
                    v3 = (^(v3))
                    _ = fi28
                    _ = v31
                }
            }
        case "foo bar":
            v32 := GenS1{F0: 65535, F1: 1000, F2: 4294967295}
            _ = v32
        case "\n\t":
            for fi33 := 0; fi33 < 6; fi33++ {
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                _ = fi33
            }
        default:
            if v15 {
                var v34 int = 32767
                v35 := GenS3{F0: v15}
                _ = v34
                _ = v35
            } else {
                var v36 int = 127
                _ = v36
            }
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v3 -= (v3 + v3)
    var v37 int32 = 1
    if (true && (-506 >= 9223372036854775807)) {
        fmt.Println(v15)
    }
    v2[1] = (100 != 1)
    var v38 float32 = (1.5 / (((0.1 - -2.25) * 1.5) / float32(1.5)))
    var v39 int16 = 1
    delete(v2, -276)
    v40 := GenS3{F0: v15}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
