package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS7 struct { F0 uint64; F1 uint16 }
func genFn0() int {
    var v7 uint8 = 162
    var v8 float64 = 1e10
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int32, p1 bool) uint {
    if true {
        v9 := GenS1{F0: 65535, F1: 22, F2: ((0 == 2) || (-32768 >= 1))}
        var v10 int64 = -9223372036854775808
        switch "a" {
        case "\n\t":
            switch 42 {
            case -14:
                var v11 uint32 = 2147483648
                _ = v11
            case 2147483647:
                var v12 uint = 686
                _ = v12
            }
        case "foo bar":
            v13 := GenS2{F0: 564}
            _ = v13
        case "hello":
            for fi14 := 0; fi14 < 15; fi14++ {
                v15 := []int{(-(fi14)), fi14, genFn0()}
                v10 = (v10 / ((v10) | 1))
                _ = fi14
                _ = v15
            }
        default:
            p0 = (^(p0))
        }
        _ = v10
        _ = v9
    }
    p1 = p1
    var v16 float32 = (0.0 - 0.1)
    v17 := []int{}
    _ = len(v17)
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 42
}

func main() {
    var v1 float32 = 1e10
    var v2 int32 = 0
    v3 := make(map[string]string)
    var v4 int = 1
    var v5 int32 = (v2 % (((v2 * v2)) | 1))
    v6 := []string{"foo bar", "hello", "\n\t", "foo bar"}
    fmt.Println(v4)
    for fi18 := 0; fi18 < 3; fi18++ {
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := v3["\n\t"]
            var v21 uint32 = 2147483648
            if false {
                if false {
                    v4 = fi19
                    if len(v6) > 0 {
                        v6[0] = v20
                    }
                    switch v20 {
                    case "foo bar":
                        var v22 int16 = 665
                        _ = v22
                    case "hello":
                        var v23 uint16 = 2
                        _ = v23
                    case "":
                        v24 := GenS3{F0: -1}
                        _ = v24
                    default:
                        switch (v20 + (((("a" + ((v20 + ("a" + v20)) + "a")) + (v20 + v20)) + v20) + ("foo bar" + ((v20 + v20) + (((v20 + v20) + v20) + ("hello" + "hello")))))) {
                        case "foo bar":
                            if true {
                                v25 := GenS4{F0: (v2 % ((v5) | 1)), F1: 65535, F2: 432}
                                delete(v3, "a")
                                v5 -= (v5 & -2147483648)
                                _ = v25
                            }
                        case "hello":
                            var v26 uint64 = 0
                            _ = v26
                        }
                    }
                    fmt.Println(fi19, v5)
                } else {
                    _ = len(v6)
                    if ((-2.25 > (((v1 - (3.14159 - 3.14159)) / float32(-2.25)) + v1)) || (-32768 <= 1000)) {
                        fmt.Println(v5, v20)
                        if false {
                            fmt.Println(v20)
                            var v27 float64 = ((3.14159 * -2.25) * 1.5)
                            switch (fi19 & -2147483648) {
                            case 42:
                                if false {
                                    delete(v3, ((((v20 + v20) + (v20 + v20)) + v20) + (v20 + v20)))
                                    v20 = "hello"
                                    var v28 int = fi18
                                    _ = v28
                                }
                            case 32767:
                                if false {
                                    v29 := GenS2{F0: 0}
                                    _ = v29
                                }
                            }
                            _ = v27
                        }
                    } else {
                        if false {
                            var v30 int8 = 100
                            if true {
                                v21 = (v21 * 55)
                            }
                            var v31 int64 = -583
                            v31 = (v31 - (v31 ^ v31))
                            v20 = ("a" + v20)
                            _ = v30
                            _ = v31
                        } else {
                            var v32 uint64 = 927
                            var v33 float64 = 1.5
                            _ = v32
                            _ = v33
                        }
                    }
                    v34 := GenS5{F0: (!(1 != 2))}
                    v6 = append(v6, v20)
                    _ = v34
                }
            } else {
                switch 968 {
                case -9223372036854775808:
                    v3["hello"] = "a"
                case 414:
                    fmt.Println(v21)
                case 32768:
                    v21 += (^(v21))
                default:
                    v35 := GenS3{F0: -462}
                    _ = v35
                }
            }
            var v36 uint32 = (v21 | 2)
            _ = fi19
            _ = v20
            _ = v21
            _ = v36
        }
        fmt.Println(v2, v4, v5)
        for fi37 := 0; fi37 < 12; fi37++ {
            var v38 bool = (482 > 964)
            fmt.Println(fi37, v2, v5)
            v4 = genFn0()
            _ = fi37
            _ = v38
        }
        _ = fi18
    }
    for fi39 := 0; fi39 < 15; fi39++ {
        if (true || true) {
            v40 := GenS4{F0: (v2 + v2), F1: 1, F2: 4294967295}
            if true {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v41 := GenS6{F0: 127, F1: 4294967295, F2: 0}
                v42 := GenS3{F0: -832}
                _ = v41
                _ = v42
            }
            for fi43 := 0; fi43 < 17; fi43++ {
                fmt.Println(fi39)
                if (0.1 < -2.25) {
                    v4 = (fi43 >> ((genFn1((v5 << ((42) & 63)), false)) & 63))
                    if (65535 != 42) {
                        var v44 float64 = (float64((float32((-2.25 + 3.14159)) + (3.14159 * 0.0))) + (3.14159 + ((float64(v1) + 0.1) * 3.14159)))
                        _ = v44
                    } else {
                        v4 = fi39
                        _ = len(v6)
                    }
                } else {
                    var v45 int = fi39
                    v46 := GenS6{F0: 255, F1: 624, F2: 128}
                    fmt.Println(v4, v2, v5)
                    _ = v45
                    _ = v46
                }
                _ = fi43
            }
            _ = v40
        }
        var v47 int16 = 1000
        _ = fi39
        _ = v47
    }
    var v48 int16 = -1
    if true {
        var v49 uint8 = 174
        var v50 uint = 0
        _ = v49
        _ = v50
    }
    if false {
        v51 := GenS1{F0: 2, F1: -114, F2: false}
        v48 -= v48
        switch genFn0() {
        case -931:
            v5 += (v5 * (v2 | v5))
        case 11:
            v52 := GenS3{F0: v48}
            _ = v52
        default:
            if (true && false) {
                var v53 int32 = (v5 / (((v5 | 0)) | 1))
                var v54 uint64 = 0
                _ = v53
                _ = v54
            }
        }
        _ = v51
    } else {
        v55 := v3["foo bar"]
        var v56 int32 = int32(v48)
        _ = v55
        _ = v56
    }
    var v57 uint32 = 2
    v4 += v4
    switch (((("foo bar" + "\n\t") + "") + "hello") + "a") {
    case "\n\t":
        var v58 int64 = -31
        _ = v58
    case "foo bar":
        v59 := GenS7{F0: 414, F1: 532}
        _ = v59
    default:
        v5 -= v5
    }
    var v60 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v57
    _ = v6
    _ = v60
}
