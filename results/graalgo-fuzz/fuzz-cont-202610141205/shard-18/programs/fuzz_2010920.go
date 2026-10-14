package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int16; F1 bool }
func genFn0(p0 int16, p1 uint8) int8 {
    var v6 string = (("" + "foo bar") + ("" + ("foo bar" + (("\n\t" + (("hello" + "\n\t") + "\n\t")) + "\n\t"))))
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{}
        if false {
            var v9 uint16 = 957
            _ = v9
        } else {
            var v10 float64 = (3.14159 - (1e10 - 0.1))
            switch fi7 {
            case -650:
                _ = len(v8)
            default:
                for fi11 := 0; fi11 < 5; fi11++ {
                    v12 := GenS1{F0: -2147483648, F1: (p1 * (^(p1)))}
                    if (!(float64(1.5) >= v10)) {
                        switch fi7 {
                        case 127:
                            var v13 int32 = -2147483648
                            _ = v13
                        case 32767:
                            var v14 int8 = -1
                            _ = v14
                        default:
                            if (0 < 922) {
                                var v15 uint64 = 792
                                _ = v15
                            } else {
                                var v16 uint64 = 0
                                _ = v16
                            }
                        }
                    }
                    _ = fi11
                    _ = v12
                }
            }
            v8 = append(v8, (fi7 % (((fi7 + (fi7 / ((884) | 1)))) | 1)))
            _ = v10
        }
        var v17 int16 = 1000
        _ = fi7
        _ = v17
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return -128
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = v1
    var v3 bool = true
    v4 := []int{-1, 32768}
    var v5 int8 = 42
    v18 := GenS2{F0: 1, F1: (1e10 == 1e10)}
    fmt.Println(v3)
    v19 := GenS2{F0: -32768, F1: v3}
    for fi20 := 0; fi20 < 13; fi20++ {
        v4 = append(v4, fi20)
        v5 = (v5 + 127)
        v21 := GenS1{F0: 699, F1: 1}
        _ = fi20
        _ = v21
    }
    switch 911 {
    case 128:
        v22 := GenS2{F0: 32767, F1: v3}
        _ = v22
    }
    v23 := GenS1{F0: 0, F1: 128}
    var v24 uint8 = 127
    switch ("a" + "hello") {
    case "":
        v25 := GenS2{F0: 838, F1: (v3 || v3)}
        _ = v25
    case "hello":
        switch "" {
        case "a":
            var v26 int32 = -398
            _ = v26
        }
    case "\n\t":
        v1 = (v2 ^ (v2 >> ((4294967295) & 63)))
    default:
        var v27 int = 32768
        _ = v27
    }
    if ((v5 * (v5 << ((3) & 63))) > (v5 ^ v5)) {
        fmt.Println(v5, v3, v24)
    } else {
        var v28 int64 = (v1 | (v2 % (((v2 + v2)) | 1)))
        var v29 int8 = -128
        if ((v1 % ((42) | 1)) <= (^(v1))) {
            for fi30 := 0; fi30 < 9; fi30++ {
                for fi31 := 0; fi31 < 10; fi31++ {
                    var v32 int32 = -125
                    _ = fi31
                    _ = v32
                }
                v33 := GenS2{F0: 0, F1: ("" > "\n\t")}
                var v34 int8 = v29
                _ = fi30
                _ = v33
                _ = v34
            }
            if (-1 <= 1) {
                var v35 uint8 = 255
                for fi36 := 0; fi36 < 3; fi36++ {
                    if len(v4) > 0 {
                        v4[0] = fi36
                    }
                    _ = fi36
                }
                _ = v35
            } else {
                switch 32768 {
                case 266:
                    v4 = append(v4, 128)
                default:
                    v29 += v5
                }
                var v37 int = 128
                _ = v37
            }
            for fi38 := 0; fi38 < 15; fi38++ {
                switch 510 {
                case -491:
                    v39 := GenS2{F0: 705, F1: (!(-2.25 >= -2.25))}
                    _ = v39
                case -64:
                    fmt.Println(v28)
                case 127:
                    v40 := GenS1{F0: 1, F1: 42}
                    _ = v40
                }
                if v3 {
                    if v3 {
                        switch 4294967295 {
                        case -1:
                            _ = len(v4)
                        default:
                            for fi41 := 0; fi41 < 15; fi41++ {
                                var v42 int = fi41
                                _ = fi41
                                _ = v42
                            }
                        }
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v28 += (v2 << ((4294967295) & 63))
                }
                if (false && true) {
                    var v43 int = fi38
                    var v44 uint = 959
                    if (210 <= -1) {
                        v44 = v44
                        fmt.Println(v3)
                        for fi45 := 0; fi45 < 17; fi45++ {
                            v5 += (v29 << (((v44 | v44)) & 63))
                            _ = fi45
                        }
                    } else {
                        v46 := GenS1{F0: 2147483647, F1: (v24 / ((1) | 1))}
                        _ = v46
                    }
                    _ = v43
                    _ = v44
                }
                _ = fi38
            }
        } else {
            v47 := GenS1{F0: -1, F1: v24}
            fmt.Println(v5)
            _ = v47
        }
        _ = v28
        _ = v29
    }
    var v48 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v48
    _ = v5
}
