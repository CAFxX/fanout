package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int; F1 int64 }
func genFn0(p0 int) uint16 {
    p0 -= p0
    v6 := []int{32767, -9223372036854775808, 42, 303}
    _ = p0
    _ = v6
    return 105
}

func main() {
    var v1 bool = false
    v2 := []string{}
    var v3 string = (("foo bar" + "a") + "foo bar")
    v4 := make(map[uint32]bool)
    var v5 float64 = (1e10 + 1e10)
    switch v3 {
    case "foo bar":
        var v7 int16 = 32767
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 14; fi8++ {
            v3 = (v3 + v3)
            switch fi8 {
            case 4294967295:
                for fi9 := 0; fi9 < 6; fi9++ {
                    _ = len(v2)
                    _ = fi9
                }
            case 9223372036854775807:
                fmt.Println(v5, fi8, fi8)
            default:
                delete(v4, 329)
            }
            switch fi8 {
            case 204:
                for fi10 := 0; fi10 < 7; fi10++ {
                    v11 := GenS1{F0: 229, F1: -2147483648}
                    _ = fi10
                    _ = v11
                }
            case 0:
                switch v3 {
                case "hello":
                    v4[42] = true
                case "":
                    if (!((114 < 127) || (125 == -845))) {
                        var v12 string = "\n\t"
                        var v13 float32 = 3.14159
                        var v14 uint64 = 206
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        _ = len(v4)
                    }
                case "a":
                    var v15 int16 = -827
                    _ = v15
                }
            case -9223372036854775808:
                var v16 uint16 = 32767
                _ = v16
            default:
                var v17 int32 = 2147483647
                _ = v17
            }
            _ = fi8
        }
    }
    var v18 string = v3
    for fi19 := 0; fi19 < 8; fi19++ {
        if ((-2.25 - (1e10 * 0.1)) < 0.0) {
            var v20 string = "a"
            v2 = append(v2, ("hello" + ("" + v3)))
            if (694 == 65535) {
                v5 -= v5
            }
            _ = v20
        } else {
            var v21 uint = 418
            _ = v21
        }
        var v22 uint32 = 2147483648
        _ = fi19
        _ = v22
    }
    fmt.Println(v1, v3, v5)
    for fi23 := 0; fi23 < 17; fi23++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (-9223372036854775808 >= -9223372036854775808) {
            v24 := GenS1{F0: 79, F1: -501}
            if v1 {
                for fi25 := 0; fi25 < 1; fi25++ {
                    if v1 {
                        var v26 float64 = 0.1
                        v27 := GenS1{F0: 92, F1: -889}
                        if false {
                            var v28 bool = true
                            _ = v28
                        }
                        _ = v26
                        _ = v27
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(v5, v3, fi23)
                        if ((v5 / (((((float64(1e10) / 1.0) + 1e10) * v5) * v5) * -2.25)) >= 0.1) {
                            if v1 {
                                var v29 int = fi25
                                var v30 int16 = 1000
                                _ = v29
                                _ = v30
                            } else {
                                v31 := GenS1{F0: 1, F1: 1}
                                if v1 {
                                    v2 = append(v2, "foo bar")
                                } else {
                                    fmt.Println(v18, v5)
                                }
                                delete(v4, 220)
                                fmt.Println(fi23, v5)
                                _ = v31
                            }
                            switch (fi25 >> ((65535) & 63)) {
                            case 1:
                                v1 = v1
                            case 32768:
                                var v32 int16 = -32768
                                _ = v32
                            default:
                                v3 = v18
                            }
                        }
                    }
                    var v33 string = v18
                    _ = fi25
                    _ = v33
                }
            }
            _ = v24
        }
        _ = fi23
    }
    switch v18 {
    case "a":
        v34 := GenS2{F0: 2147483647, F1: 9223372036854775807}
        _ = v34
    case "\n\t":
        var v35 uint8 = 128
        _ = v35
    }
    v36 := GenS2{F0: 32767, F1: -9223372036854775808}
    if (v5 <= (v5 * v5)) {
        v37 := GenS1{F0: 1, F1: 23}
        v4[353] = v1
        _ = v37
    } else {
        for fi38 := 0; fi38 < 14; fi38++ {
            v18 = "hello"
            if v1 {
                for fi39 := 0; fi39 < 2; fi39++ {
                    fmt.Println(fi39, v3)
                    v18 = "\n\t"
                    _ = fi39
                }
            }
            _ = fi38
        }
        v40 := v4[2147483648]
        _ = v40
    }
    fmt.Println(v1)
    v41 := GenS1{F0: 42, F1: 42}
    for fi42 := 0; fi42 < 6; fi42++ {
        v43 := GenS1{F0: 255, F1: -634}
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v41
    _ = v5
}
