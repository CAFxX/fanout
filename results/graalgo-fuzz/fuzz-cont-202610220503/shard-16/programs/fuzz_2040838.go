package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint32; F2 int32 }
type GenS4 struct { F0 int; F1 int8; F2 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool }
func genFn0() uint64 {
    var v6 uint16 = 2
    var v7 int32 = -2147483648
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint = 794
    v2 := make(map[int]int)
    var v3 uint = v1
    var v4 float64 = 3.14159
    var v5 string = (("foo bar" + "hello") + "a")
    if (true || ((true || (((v4 / float64((-2.25 + (3.14159 - 1.5)))) - v4) >= v4)) || false)) {
        var v8 uint16 = 32767
        if (!((float32(float64(((((1e10 / 1.0) / (-2.25 * 0.1)) / (float32(0.0) - (0.1 / 1.0))) - 0.1))) / 1.0) > ((-2.25 * -2.25) / (0.0 - (0.1 - (float32(float64(1e10)) - (0.0 + 0.1))))))) {
            var v9 uint32 = 1
            fmt.Println(v3, v3)
            _ = v9
        } else {
            switch ("foo bar" + "foo bar") {
            case "foo bar":
                v10 := GenS1{F0: v3}
                _ = v10
            case "":
                v11 := []int{42, -2147483648}
                _ = v11
            case "a":
                fmt.Println(v5, v4)
            default:
                var v12 uint = 161
                _ = v12
            }
            var v13 int64 = -301
            _ = v13
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 string = "hello"
            v1 = v3
            v16 := GenS2{F0: v8}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v8
    } else {
        switch v5 {
        case "\n\t":
            if (42 < 42) {
                if (314 < -2147483648) {
                    if true {
                        for fi17 := 0; fi17 < 15; fi17++ {
                            for fi18 := 0; fi18 < 18; fi18++ {
                                v19 := v2[fi17]
                                _ = fi18
                                _ = v19
                            }
                            v20 := []int{fi17, fi17}
                            _ = fi17
                            _ = v20
                        }
                        v21 := []int{32768, 4294967295, -913, 357}
                        v22 := GenS3{F0: 0, F1: 884, F2: 2147483647}
                        _ = v21
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 17; fi23++ {
                            v5 = "foo bar"
                            fmt.Println(v3, v4)
                            v3 -= v1
                            _ = fi23
                        }
                        var v24 uint64 = 313
                        _ = v24
                    }
                    v3 += 1
                }
                fmt.Println(v4)
            } else {
                v3 = v1
            }
        case "hello":
            if (823 != -640) {
                if true {
                    var v25 uint32 = 2147483647
                    for fi26 := 0; fi26 < 10; fi26++ {
                        var v27 int16 = 424
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                } else {
                    if true {
                        if false {
                            v28 := v2[-9223372036854775808]
                            if (!false) {
                                v3 = (v1 * 0)
                                for fi29 := 0; fi29 < 1; fi29++ {
                                    var v30 uint64 = 0
                                    var v31 uint32 = 562
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                            }
                            v28 = v28
                            _ = v28
                        }
                        fmt.Println(v3, v3)
                    }
                    fmt.Println(v3, v1, v4)
                    switch -1 {
                    case 128:
                        if false {
                            _ = len(v2)
                            fmt.Println(v4, v4, v3)
                            var v32 int64 = 505
                            _ = v32
                        } else {
                            switch -260 {
                            case 0:
                                v1 = 695
                            case -2147483648:
                                for fi33 := 0; fi33 < 7; fi33++ {
                                    v4 -= v4
                                    fmt.Println(v4, v4, v4)
                                    switch (v5 + "\n\t") {
                                    case "foo bar":
                                        var v34 int32 = 756
                                        _ = v34
                                    case "hello":
                                        v35 := GenS1{F0: (v3 + 207)}
                                        _ = v35
                                    }
                                    _ = fi33
                                }
                            }
                            for fi36 := 0; fi36 < 2; fi36++ {
                                switch v5 {
                                case "\n\t":
                                    var v37 int32 = -1
                                    _ = v37
                                case "hello":
                                    if true {
                                        v38 := GenS1{F0: v1}
                                        switch ((v5 + "foo bar") + ("hello" + "foo bar")) {
                                        case "hello":
                                            v3 = v3
                                        case "\n\t":
                                            v4 -= v4
                                        case "a":
                                            v39 := GenS2{F0: 436}
                                            _ = v39
                                        }
                                        _ = v38
                                    } else {
                                        v3 += 0
                                    }
                                case "foo bar":
                                    v40 := []int{(fi36 >> (((v1 % ((v3) | 1))) & 63)), (fi36 * 313), (fi36 << (((v1 >> ((v1) & 63))) & 63)), fi36}
                                    _ = v40
                                }
                                _ = fi36
                            }
                            v41 := []int{800, -718, 2147483647, 127}
                            _ = v41
                        }
                    case -2147483648:
                        var v42 uint8 = 1
                        _ = v42
                    case -1:
                        var v43 float64 = (((((v4 - v4) - v4) * 0.1) - float64((3.14159 / (1e10 * 3.14159)))) / v4)
                        _ = v43
                    default:
                        _ = len(v2)
                    }
                }
                var v44 float32 = 0.1
                _ = v44
            }
        default:
            v45 := []int{30, 32767, -166}
            _ = v45
        }
        var v46 float32 = (float32((float64(3.14159) / 3.14159)) / (-2.25 + 1e10))
        for fi47 := 0; fi47 < 2; fi47++ {
            var v48 uint32 = 828
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    v3 += v3
    v2[127] = 128
    v3 = v3
    v49 := GenS4{F0: -2147483648, F1: 0, F2: 42}
    switch 127 {
    case 332:
        v5 = (v5 + "\n\t")
    case -2147483648:
        if false {
            v50 := GenS5{F0: 42}
            fmt.Println(v1, v1, v1)
            v51 := GenS6{F0: 2}
            _ = v50
            _ = v51
        }
    }
    switch -669 {
    case 128:
        for fi52 := 0; fi52 < 19; fi52++ {
            v2[fi52] = -9223372036854775808
            _ = fi52
        }
    case -78:
        var v53 int8 = 42
        _ = v53
    default:
        var v54 uint8 = 0
        _ = v54
    }
    if ((!true) || false) {
        var v55 float64 = ((v4 + 0.0) + v4)
        switch v5 {
        case "hello":
            switch v5 {
            case "":
                if false {
                    var v56 int8 = 127
                    _ = v56
                }
            case "a":
                v55 += 0.1
            }
        case "foo bar":
            v1 += v1
        case "a":
            v1 = 693
        }
        if (true && (18446744073709551615 != 0)) {
            for fi57 := 0; fi57 < 4; fi57++ {
                for fi58 := 0; fi58 < 18; fi58++ {
                    v5 = (v5 + ((v5 + (((v5 + (v5 + v5)) + v5) + v5)) + "\n\t"))
                    switch v5 {
                    case "foo bar":
                        v5 = ""
                    case "hello":
                        if (1000 != -1) {
                            v5 = v5
                        }
                    case "":
                        v59 := GenS7{F0: false}
                        _ = v59
                    }
                    _ = fi58
                }
                _ = fi57
            }
            v1 += (v3 - 660)
            v55 = v4
        } else {
            var v60 bool = false
            v5 = (v5 + ((v5 + v5) + ("foo bar" + "\n\t")))
            _ = v60
        }
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
}
