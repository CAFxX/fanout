package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint8) bool {
    v4 := []int{1}
    if (42 < 42) {
        var v5 string = (("hello" + "a") + ("hello" + "hello"))
        var v6 uint32 = 2147483648
        for fi7 := 0; fi7 < 9; fi7++ {
            if p0 {
                var v8 uint64 = 4294967295
                v9 := GenS1{F0: (v5 + "\n\t")}
                if (2147483647 == 2147483647) {
                    var v10 uint8 = 128
                    _ = v10
                }
                _ = v8
                _ = v9
            }
            switch 1 {
            case 4294967295:
                p0 = (!p0)
            default:
                p0 = ((42 == 127) && p0)
            }
            _ = fi7
        }
        _ = v5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return true
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]uint32)
    v3 := []string{"foo bar", "foo bar"}
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 uint8 = 42
        _ = fi11
        _ = v12
    }
    var v13 float64 = ((3.14159 + -2.25) * ((float64((-2.25 / 1.0)) - 0.1) + 3.14159))
    v13 += v13
    switch -2147483648 {
    case 2147483647:
        switch (("" + "foo bar") + "\n\t") {
        case "":
            if len(v3) > 0 {
                v3[0] = ("foo bar" + ("a" + ("\n\t" + "a")))
            }
        default:
            if true {
                for fi14 := 0; fi14 < 2; fi14++ {
                    v15 := GenS1{F0: "\n\t"}
                    _ = fi14
                    _ = v15
                }
                var v16 uint64 = 0
                if (-398 > 0) {
                    if true {
                        fmt.Println(v13)
                        v17 := GenS1{F0: "\n\t"}
                        switch 127 {
                        case 32768:
                            switch 798 {
                            case 127:
                                var v18 uint16 = 639
                                _ = v18
                            case 4294967295:
                                v19 := GenS1{F0: "\n\t"}
                                _ = v19
                            case 351:
                                v13 = 1e10
                            default:
                                v16 = v16
                            }
                        }
                        _ = v17
                    } else {
                        if (v13 != 0.1) {
                            v13 = ((1.5 + float64(0.1)) - (v13 - v13))
                        }
                        fmt.Println(v1)
                        switch 32768 {
                        case 42:
                            var v20 int = 1
                            _ = v20
                        case 2147483647:
                            if false {
                                for fi21 := 0; fi21 < 16; fi21++ {
                                    var v22 int = fi21
                                    var v23 int = 32767
                                    fmt.Println(v13, v13, v23)
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                                if true {
                                    v24 := GenS2{F0: (v1 - uint(v16)), F1: v1}
                                    var v25 uint8 = 128
                                    v26 := GenS3{F0: v16}
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                                v27 := GenS3{F0: (v16 + 1)}
                                _ = v27
                            }
                        case 405:
                            _ = len(v3)
                        default:
                            switch 0 {
                            case -946:
                                var v28 uint64 = (^(v16))
                                _ = v28
                            case 42:
                                fmt.Println(v13, v16, v16)
                            }
                        }
                    }
                    v29 := v2[127]
                    _ = v29
                } else {
                    var v30 uint32 = 1
                    v16 = (v16 | (v16 | (v16 % ((v16) | 1))))
                    var v31 uint64 = uint64(v30)
                    _ = v30
                    _ = v31
                }
                _ = v16
            }
        }
    case 418:
        v1 += v1
    case 32768:
        for fi32 := 0; fi32 < 2; fi32++ {
            var v33 float32 = 0.0
            _ = fi32
            _ = v33
        }
    }
    v34 := GenS1{F0: (("a" + "") + ("hello" + ("foo bar" + "a")))}
    for fi35 := 0; fi35 < 16; fi35++ {
        v13 += -2.25
        v36 := GenS2{F0: 432, F1: (v1 >> ((v1) & 63))}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi35
        _ = v36
    }
    var v37 uint8 = 127
    v38 := GenS1{F0: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v38
}
