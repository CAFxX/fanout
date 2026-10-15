package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 bool; F2 string }
type GenS4 struct { F0 int64; F1 bool; F2 int8 }
func genFn0() uint64 {
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 uint64 = 1
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 0
    v7 += 4294967295
    v7 = v7
    if (!(0 != 1)) {
        v8 := GenS1{F0: 32767, F1: 173}
        v9 := []int{94, -2147483648}
        var v10 uint64 = v7
        _ = v10
        _ = v8
        _ = v9
    } else {
        fmt.Println(v7)
    }
    _ = v7
    return (v7 << ((255) & 63))
}

func main() {
    var v1 bool = true
    v2 := make(map[int]string)
    v3 := make(map[uint32]string)
    var v4 bool = (float32(1e10) <= float32(0.1))
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 int32 = -322
        v13 := GenS1{F0: -153, F1: 255}
        fmt.Println(v4, v4, v1)
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int8 = -128
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := []int{fi15, fi15, -1, (-(fi15))}
        v14 = 17
        v17 := GenS1{F0: -1, F1: 99}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint8 = 255
    fmt.Println(v14, v1)
    var v19 uint32 = 1
    v1 = false
    v1 = false
    v20 := GenS2{F0: 1}
    fmt.Println(v18)
    switch "\n\t" {
    case "a":
        fmt.Println(v19, v14)
    default:
        if false {
            var v21 uint64 = 0
            if false {
                if v4 {
                    v4 = ((0.1 * (1.5 + ((((float32(float64(-2.25)) / 1.0) + (0.1 * (-2.25 * 0.1))) - (3.14159 * 1e10)) + (-2.25 / 3.14159)))) > 1.5)
                } else {
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        v22 := GenS2{F0: 2}
                        _ = v22
                    case "a":
                        switch -9223372036854775808 {
                        case 639:
                            switch ("foo bar" + ("a" + ("a" + ""))) {
                            case "\n\t":
                                for fi23 := 0; fi23 < 1; fi23++ {
                                    switch (fi23 ^ fi23) {
                                    case 42:
                                        fmt.Println(v4, fi23)
                                    case 127:
                                        fmt.Println(v4, v19, v1)
                                    case 32767:
                                        if (!true) {
                                            for fi24 := 0; fi24 < 12; fi24++ {
                                                var v25 int = fi23
                                                _ = fi24
                                                _ = v25
                                            }
                                            for fi26 := 0; fi26 < 12; fi26++ {
                                                v1 = v4
                                                _ = fi26
                                            }
                                            v1 = v4
                                        } else {
                                            v1 = (("a" + "hello") != ("a" + "\n\t"))
                                            fmt.Println(v21)
                                        }
                                    }
                                    _ = fi23
                                }
                            case "":
                                v27 := GenS3{F0: (("a" + "") + ""), F1: (false || (!(false && v4))), F2: "foo bar"}
                                _ = v27
                            }
                        case -9223372036854775808:
                            v14 = v14
                        case -461:
                            v28 := GenS3{F0: "hello", F1: false, F2: "a"}
                            _ = v28
                        }
                    case "":
                        v1 = ((!(!false)) && false)
                    default:
                        var v29 uint64 = 42
                        _ = v29
                    }
                }
                fmt.Println(v18, v14)
            }
            switch "hello" {
            case "":
                var v30 string = "a"
                _ = v30
            case "a":
                for fi31 := 0; fi31 < 8; fi31++ {
                    switch fi31 {
                    case 128:
                        fmt.Println(v21, v4, fi31)
                    case -516:
                        for fi32 := 0; fi32 < 4; fi32++ {
                            var v33 int16 = 388
                            _ = fi32
                            _ = v33
                        }
                    case 32768:
                        fmt.Println(v18, fi31)
                    }
                    _ = fi31
                }
            default:
                v18 += v18
            }
            _ = v21
        } else {
            v34 := GenS3{F0: ("\n\t" + "a"), F1: (((^(v19)) != 2147483647) || (946 == -99)), F2: "foo bar"}
            v4 = v1
            v35 := GenS4{F0: -1, F1: true, F2: (^(v14))}
            _ = v34
            _ = v35
        }
    }
    fmt.Println(v14, v14)
    for fi36 := 0; fi36 < 6; fi36++ {
        v4 = true
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
