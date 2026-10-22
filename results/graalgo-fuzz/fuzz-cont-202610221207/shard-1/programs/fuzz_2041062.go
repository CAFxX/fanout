package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int8; F1 uint; F2 int64 }
type GenS5 struct { F0 uint16; F1 int32; F2 int }
type GenS6 struct { F0 uint16; F1 int64; F2 uint64 }
func genFn0() int32 {
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: 752, F1: ""}
        switch (fi5 / ((2147483647) | 1)) {
        case -2147483648:
            var v7 string = (("foo bar" + "") + ("a" + "a"))
            _ = v7
        case 1:
            fmt.Println(fi5)
        case 42:
            if false {
                for fi8 := 0; fi8 < 3; fi8++ {
                    var v9 uint32 = 2147483647
                    fmt.Println(fi8)
                    _ = fi8
                    _ = v9
                }
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 bool = (3.14159 != -2.25)
                    var v12 int8 = -128
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
        }
        _ = fi5
        _ = v6
    }
    if (18446744073709551615 == 42) {
        v13 := []int{345, 4}
        _ = v13
    } else {
        var v14 uint64 = 927
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(v14, v14)
            v14 = 42
            _ = fi15
        }
        if true {
            var v16 bool = (-2.25 >= 1.5)
            for fi17 := 0; fi17 < 4; fi17++ {
                switch "a" {
                case "foo bar":
                    v14 = 4294967295
                case "":
                    v18 := GenS2{F0: 0, F1: ""}
                    _ = v18
                case "hello":
                    var v19 int8 = -81
                    _ = v19
                default:
                    v20 := GenS3{F0: ("a" + "hello"), F1: 2}
                    _ = v20
                }
                _ = fi17
            }
            _ = v16
        } else {
            if false {
                if true {
                    if (42 >= 793) {
                        v21 := []int{1, -568, -9223372036854775808, -9223372036854775808}
                        if len(v21) > 0 {
                            v21[0] = -1
                        }
                        _ = v21
                    } else {
                        var v22 float64 = -2.25
                        v22 += v22
                        fmt.Println(v22, v14)
                        _ = v22
                    }
                } else {
                    v14 -= v14
                    v23 := []int{944, -9223372036854775808, -1}
                    _ = v23
                }
            }
            switch 4294967295 {
            case -1:
                for fi24 := 0; fi24 < 7; fi24++ {
                    for fi25 := 0; fi25 < 15; fi25++ {
                        if true {
                            v26 := GenS4{F0: -49, F1: 4294967295, F2: 83}
                            _ = v26
                        }
                        v27 := GenS2{F0: 0, F1: ((("\n\t" + "") + "foo bar") + "hello")}
                        _ = fi25
                        _ = v27
                    }
                    _ = fi24
                }
            }
        }
        _ = v14
    }
    if true {
        var v28 int = 42
        _ = v28
    } else {
        var v29 uint64 = 0
        for fi30 := 0; fi30 < 3; fi30++ {
            for fi31 := 0; fi31 < 18; fi31++ {
                fmt.Println(fi31, v29, fi30)
                if false {
                    v29 -= (v29 % (((v29 + 2)) | 1))
                    fmt.Println(fi30)
                    fmt.Println(v29, fi30, fi30)
                } else {
                    switch (fi30 + (fi30 << ((1) & 63))) {
                    case 202:
                        v32 := GenS4{F0: -128, F1: 65535, F2: -1}
                        _ = v32
                    case 128:
                        v33 := []int{(fi31 ^ (fi31 - 2147483647))}
                        v29 -= (v29 % ((42) | 1))
                        v34 := GenS1{F0: 4294967295, F1: (((("a" + "hello") + (("hello" + ("foo bar" + ("a" + "\n\t"))) + (("" + "\n\t") + "foo bar"))) + "hello") + ("hello" + "foo bar"))}
                        var v35 int8 = -15
                        var v36 int32 = 0
                        v33 = append(v33, fi31)
                        _ = v33
                        _ = v34
                        _ = v35
                        _ = v36
                    case 32767:
                        v29 = (v29 << ((18446744073709551615) & 63))
                    default:
                        var v37 int = 4294967295
                        _ = v37
                    }
                    v38 := []int{fi31}
                    _ = v38
                }
                _ = fi31
            }
            v39 := GenS5{F0: 27, F1: 42, F2: (fi30 & fi30)}
            fmt.Println(v29, fi30)
            _ = fi30
            _ = v39
        }
        var v40 int64 = 42
        _ = v29
        _ = v40
    }
    return -2147483648
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 255
    var v3 string = "\n\t"
    v4 := make(map[uint32]uint32)
    v2 -= v2
    fmt.Println(v2)
    v3 = ((v3 + v3) + v3)
    switch 700 {
    case -2147483648:
        if false {
            var v41 float64 = (1e10 / (((-2.25 * -2.25) - 1.5) + (0.1 / 3.14159)))
            v3 = v3
            for fi42 := 0; fi42 < 16; fi42++ {
                for fi43 := 0; fi43 < 6; fi43++ {
                    if (true && (1 <= 42)) {
                        var v44 uint8 = 0
                        _ = v44
                    } else {
                        v45 := []int{(fi43 * -2147483648), (fi43 % (((fi42 - int(v2))) | 1)), fi43, 32767}
                        switch fi42 {
                        case -9223372036854775808:
                            if true {
                                v46 := GenS1{F0: 255, F1: ("foo bar" + (v3 + v3))}
                                _ = v46
                            }
                        case 9223372036854775807:
                            var v47 int = (fi42 + fi43)
                            _ = v47
                        case -18:
                            if false {
                                delete(v1, 1)
                            }
                        }
                        delete(v4, 4294967295)
                        _ = v45
                    }
                    var v48 int = (fi43 & fi42)
                    var v49 uint = 255
                    _ = fi43
                    _ = v48
                    _ = v49
                }
                v2 = (v2 + v2)
                v50 := GenS5{F0: 808, F1: 42, F2: (fi42 << ((18446744073709551615) & 63))}
                _ = fi42
                _ = v50
            }
            _ = v41
        }
    case -546:
        v3 = (v3 + v3)
    default:
        v51 := GenS6{F0: 601, F1: -9223372036854775808, F2: 100}
        _ = v51
    }
    var v52 uint16 = 0
    for fi53 := 0; fi53 < 5; fi53++ {
        switch (fi53 | (fi53 << ((4294967295) & 63))) {
        case 32767:
            switch v3 {
            case "":
                fmt.Println(v2, v2)
            case "foo bar":
                v54 := []int{}
                _ = v54
            default:
                v2 -= (v2 % (((v2 << ((178) & 63))) | 1))
            }
        case -242:
            for fi55 := 0; fi55 < 11; fi55++ {
                v56 := GenS4{F0: 0, F1: 42, F2: 42}
                v57 := []int{int(v52), fi53, 4294967295}
                _ = fi55
                _ = v56
                _ = v57
            }
        case -948:
            var v58 uint32 = 4294967295
            _ = v58
        }
        v59 := GenS1{F0: 65535, F1: ("\n\t" + (v3 + v3))}
        _ = fi53
        _ = v59
    }
    for fi60 := 0; fi60 < 4; fi60++ {
        for fi61 := 0; fi61 < 1; fi61++ {
            if true {
                v52 -= v52
                v3 = v3
            }
            fmt.Println(fi61, v2, v2)
            _ = fi61
        }
        _ = fi60
    }
    v52 = (v52 ^ v52)
    v52 += (^(v52))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
}
