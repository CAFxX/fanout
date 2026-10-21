package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int64; F2 int32 }
func genFn0(p0 bool) uint16 {
    v6 := []int{-195, 1}
    fmt.Println(p0)
    p0 = false
    var v7 uint8 = 1
    if false {
        v7 = v7
    } else {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 int8 = -127
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 uint8 = 2
    var v2 int32 = 42
    v3 := make(map[int]string)
    v4 := []int{}
    var v5 int8 = int8(v1)
    v5 = (v5 >> ((0) & 63))
    fmt.Println(v5, v1, v1)
    v5 += v5
    var v10 int = 0
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint16 = 42
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    switch v10 {
    case 4294967295:
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 string = "a"
            if false {
                switch fi14 {
                case -9223372036854775808:
                    var v16 bool = false
                    _ = v16
                }
                if (!false) {
                    v5 = (v5 << ((353) & 63))
                    if (2 < 42) {
                        for fi17 := 0; fi17 < 13; fi17++ {
                            if (!(((3.14159 / 1e10) != 0.0) || (48 != 125))) {
                                v10 = (^(v10))
                                v2 = v2
                            }
                            v15 = "foo bar"
                            _ = fi17
                        }
                        delete(v3, -218)
                        switch v15 {
                        case "hello":
                            for fi18 := 0; fi18 < 5; fi18++ {
                                v19 := GenS1{F0: v15}
                                _ = fi18
                                _ = v19
                            }
                        case "foo bar":
                            if false {
                                var v20 int = v10
                                switch fi14 {
                                case 4294967295:
                                    var v21 int = fi14
                                    _ = v21
                                case -9223372036854775808:
                                    for fi22 := 0; fi22 < 12; fi22++ {
                                        var v23 uint64 = 2
                                        _ = fi22
                                        _ = v23
                                    }
                                case 546:
                                    v24 := GenS1{F0: v15}
                                    _ = v24
                                }
                                fmt.Println(v20, v2, v5)
                                _ = v20
                            } else {
                                for fi25 := 0; fi25 < 9; fi25++ {
                                    if (1.5 < (float32(0.1) / float32(0.1))) {
                                        _ = len(v3)
                                        fmt.Println(fi14, fi14)
                                        if (false || true) {
                                            v26 := GenS1{F0: ((v15 + ("" + ("a" + ((v15 + v15) + ("hello" + v15))))) + "")}
                                            fmt.Println(v1, v2, v2)
                                            _ = v26
                                        }
                                    }
                                    _ = fi25
                                }
                            }
                        default:
                            v2 -= v2
                        }
                    } else {
                        v27 := GenS1{F0: ((v15 + v15) + "foo bar")}
                        _ = v27
                    }
                    v3[-2147483648] = v15
                }
            }
            switch 0 {
            case -441:
                v1 -= 255
            case 32767:
                switch (v15 + v15) {
                case "":
                    var v28 uint32 = 912
                    _ = v28
                case "foo bar":
                    var v29 int64 = -1
                    _ = v29
                default:
                    var v30 uint32 = 2
                    _ = v30
                }
            default:
                v10 = v10
            }
            _ = fi14
            _ = v15
        }
    case 32768:
        v31 := GenS1{F0: "foo bar"}
        _ = v31
    }
    v32 := GenS2{F0: 1, F1: 9223372036854775807, F2: -288}
    for fi33 := 0; fi33 < 13; fi33++ {
        for fi34 := 0; fi34 < 16; fi34++ {
            if false {
                v1 += (v1 & (v1 * 174))
            } else {
                fmt.Println(v1, v5, v5)
                var v35 int = v10
                _ = v35
            }
            v4 = append(v4, (^(fi33)))
            v10 += (^(v10))
            v36 := GenS1{F0: ""}
            if (867 < -640) {
                fmt.Println(v5, v5)
                var v37 string = "\n\t"
                fmt.Println(fi33)
                _ = v37
            } else {
                v2 = (v2 ^ (v2 * v2))
                if len(v4) > 0 {
                    v4[0] = (fi34 & 4294967295)
                }
                v38 := GenS1{F0: "a"}
                _ = v38
            }
            _ = fi34
            _ = v36
        }
        v39 := GenS2{F0: 42, F1: 9223372036854775807, F2: v2}
        v3[128] = (("\n\t" + (((("a" + ("\n\t" + ((("\n\t" + "a") + (("\n\t" + "") + ("\n\t" + (("" + ("\n\t" + ("foo bar" + "hello"))) + "foo bar")))) + "foo bar"))) + "") + "hello") + "\n\t")) + "\n\t")
        _ = fi33
        _ = v39
    }
    v5 = v5
    var v40 bool = false
    if ((1 >= 42) && (!v40)) {
        delete(v3, -9223372036854775808)
    } else {
        var v41 uint16 = 1
        v4 = append(v4, v10)
        var v42 uint8 = 185
        _ = v41
        _ = v42
    }
    v2 = (v2 % ((int32(v10)) | 1))
    v43 := GenS1{F0: "hello"}
    v40 = v40
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v43
    _ = v5
}
