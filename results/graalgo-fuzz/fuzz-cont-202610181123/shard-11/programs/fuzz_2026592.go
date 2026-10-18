package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int8; F1 int64 }
func genFn0() int32 {
    if false {
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS1{F0: 1, F1: 135, F2: 65535}
            _ = fi7
            _ = v8
        }
        v9 := []int{32767, -1, 498, 128}
        _ = v9
    }
    switch ("foo bar" + "") {
    case "":
        var v10 uint64 = 18446744073709551615
        _ = v10
    default:
        for fi11 := 0; fi11 < 12; fi11++ {
            if (635 <= 305) {
                var v12 float64 = 3.14159
                v13 := []int{fi11, (fi11 | (fi11 >> ((65535) & 63)))}
                _ = v12
                _ = v13
            }
            _ = fi11
        }
    }
    var v14 uint32 = 2147483648
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 int64 = 1
        fmt.Println(v14, fi15, v14)
        var v17 uint16 = 909
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = v14
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 int8 = 42
    var v4 int64 = -34
    var v5 int16 = -149
    var v6 uint = 18446744073709551615
    v18 := GenS1{F0: v5, F1: 2147483648, F2: 954}
    v6 += v6
    v5 = v5
    for fi19 := 0; fi19 < 5; fi19++ {
        fmt.Println(v6)
        fmt.Println(v5, v4)
        v20 := GenS1{F0: (^(v5)), F1: 2147483647, F2: 65535}
        _ = fi19
        _ = v20
    }
    if (255 < 65535) {
        v21 := GenS1{F0: (v5 | (v5 | v5)), F1: 4294967295, F2: 853}
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 int = fi22
            v4 += v4
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    v24 := []int{}
    if (true || false) {
        v3 += v3
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 float64 = ((3.14159 + (3.14159 - 1.5)) - (float64(0.0) + 0.1))
            switch (fi25 % ((-289) | 1)) {
            case 0:
                switch "" {
                case "\n\t":
                    var v27 float32 = float32(v26)
                    _ = v27
                case "foo bar":
                    for fi28 := 0; fi28 < 5; fi28++ {
                        switch (^(fi25)) {
                        case 42:
                            var v29 uint32 = 138
                            _ = v29
                        case -455:
                            switch "foo bar" {
                            case "hello":
                                var v30 bool = (!(!true))
                                _ = v30
                            case "a":
                                fmt.Println(v4, v26, v26)
                            default:
                                fmt.Println(fi28, fi25)
                            }
                        case -1:
                            if true {
                                if (v5 > v5) {
                                    v4 = v4
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v26, fi28)
                                    case "foo bar":
                                        v3 -= v3
                                    default:
                                        fmt.Println(v6)
                                    }
                                }
                                v31 := GenS1{F0: (^(v5)), F1: 2147483647, F2: 817}
                                _ = v31
                            }
                        default:
                            switch fi25 {
                            case 42:
                                fmt.Println(v5)
                                if (1 <= genFn0()) {
                                    switch "hello" {
                                    case "hello":
                                        delete(v1, -127)
                                    default:
                                        v26 += v26
                                    }
                                    var v32 bool = false
                                    _ = len(v24)
                                    _ = v32
                                } else {
                                    v33 := GenS1{F0: v5, F1: 42, F2: 1}
                                    v26 = v26
                                    delete(v1, (fi28 | (fi25 % ((int(v4)) | 1))))
                                    _ = v33
                                }
                            }
                        }
                        _ = fi28
                    }
                default:
                    for fi34 := 0; fi34 < 19; fi34++ {
                        v5 = -1
                        v3 += -118
                        var v35 int16 = -32768
                        _ = fi34
                        _ = v35
                    }
                }
            case -820:
                v26 = 1.5
            case 770:
                switch "" {
                case "\n\t":
                    var v36 string = "foo bar"
                    _ = v36
                default:
                    var v37 uint32 = 42
                    _ = v37
                }
            }
            _ = fi25
            _ = v26
        }
    } else {
        var v38 int64 = (v4 * -662)
        v3 += 100
        _ = v38
    }
    switch 2147483647 {
    case -2147483648:
        if (!(!false)) {
            if (((255 != 65535) || true) || false) {
                v24 = append(v24, 1)
            }
            fmt.Println(v5)
            v6 = 255
        } else {
            v39 := GenS2{F0: (v3 / (((^(v3))) | 1)), F1: v4}
            _ = v39
        }
    case -9223372036854775808:
        v6 += 0
    }
    v4 -= (v4 * (-(v4)))
    var v40 string = ("\n\t" + "a")
    var v41 float32 = 3.14159
    if true {
        fmt.Println(v5)
        var v42 bool = true
        _ = v42
    } else {
        for fi43 := 0; fi43 < 18; fi43++ {
            fmt.Println(v5)
            v6 = v6
            _ = fi43
        }
        var v44 int64 = v4
        _ = v44
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
