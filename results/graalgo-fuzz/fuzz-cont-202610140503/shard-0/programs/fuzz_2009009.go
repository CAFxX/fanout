package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 string }
func genFn0(p0 int16) uint16 {
    fmt.Println(p0)
    var v7 int32 = 42
    v7 += -2147483648
    _ = p0
    _ = v7
    return 65535
}
func genFn1() uint {
    var v8 int8 = -128
    var v9 float32 = 0.1
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := []string{"\n\t"}
    var v2 int16 = 32767
    var v3 float64 = (3.14159 - 1e10)
    v4 := make(map[uint32]string)
    var v5 uint64 = 1
    var v6 uint16 = 255
    v10 := GenS1{F0: 42, F1: v6, F2: (("foo bar" + ("a" + ("\n\t" + "a"))) + "hello")}
    v2 = 32767
    var v11 uint8 = 2
    v11 = (^(v11))
    var v12 uint64 = v5
    var v13 string = (("" + (("foo bar" + "a") + "\n\t")) + "")
    v11 += 128
    switch v13 {
    case "\n\t":
        if (false || false) {
            if (true || (255 == 985)) {
                for fi14 := 0; fi14 < 6; fi14++ {
                    fmt.Println(fi14, v6)
                    v15 := GenS1{F0: 2147483648, F1: 255, F2: (("a" + v13) + v13)}
                    _ = fi14
                    _ = v15
                }
                if (false || (!(!((42 < 30) && false)))) {
                    var v16 uint = genFn1()
                    v12 = 847
                    v17 := GenS1{F0: 2147483648, F1: v6, F2: (((v13 + (v13 + "")) + v13) + ("hello" + ("hello" + v13)))}
                    _ = v16
                    _ = v17
                } else {
                    if (18446744073709551615 == 4294967295) {
                        for fi18 := 0; fi18 < 8; fi18++ {
                            if false {
                                var v19 string = v13
                                v2 = v2
                                for fi20 := 0; fi20 < 3; fi20++ {
                                    var v21 uint = 0
                                    _ = fi20
                                    _ = v21
                                }
                                _ = v19
                            } else {
                                if (!false) {
                                    switch fi18 {
                                    case 0:
                                        fmt.Println(v12, v5, v6)
                                    }
                                    _ = len(v1)
                                } else {
                                    if (true && true) {
                                        fmt.Println(v11)
                                        for fi22 := 0; fi22 < 13; fi22++ {
                                            switch (fi22 % (((fi22 & fi22)) | 1)) {
                                            case 128:
                                                if true {
                                                    var v23 int8 = 37
                                                    v12 -= 0
                                                    _ = v23
                                                }
                                            default:
                                                v6 -= (v6 - (v6 << ((925) & 63)))
                                            }
                                            fmt.Println(v11, v6)
                                            _ = fi22
                                        }
                                    }
                                }
                                if (v6 <= v6) {
                                    if false {
                                        var v24 string = ((v13 + v13) + v13)
                                        _ = v24
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            var v25 uint64 = uint64(v6)
                                            _ = v25
                                        case "":
                                            v1 = append(v1, ((("hello" + v13) + (v13 + v13)) + ""))
                                        default:
                                            v11 += v11
                                        }
                                        fmt.Println(v12, v5, v11)
                                        var v26 uint = 756
                                        _ = v26
                                    }
                                } else {
                                    v11 -= v11
                                    v11 = (^(v11))
                                    var v27 uint = 42
                                    _ = v27
                                }
                            }
                            var v28 float32 = (0.1 * 0.0)
                            _ = fi18
                            _ = v28
                        }
                        var v29 string = v13
                        _ = v29
                    }
                    var v30 int8 = -1
                    _ = v30
                }
            }
            switch "hello" {
            case "foo bar":
                var v31 uint64 = v5
                _ = v31
            default:
                for fi32 := 0; fi32 < 20; fi32++ {
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v5 = 858
                        v34 := v4[0]
                        _ = fi33
                        _ = v34
                    }
                    v1 = append(v1, v13)
                    v2 = (v2 << ((42) & 63))
                    _ = fi32
                }
            }
        } else {
            v1 = append(v1, (("" + (("hello" + v13) + (v13 + ("" + v13)))) + (((v13 + v13) + ((v13 + (("\n\t" + "hello") + v13)) + v13)) + "")))
        }
    default:
        if true {
            var v35 uint32 = 296
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v35
        }
    }
    fmt.Println(v12, v2, v6)
    v12 -= (^(v5))
    var v36 int32 = -653
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
