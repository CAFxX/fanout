package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint8; F1 int; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 uint64; F2 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 int16; F2 uint8 }
func genFn0() int8 {
    v5 := GenS1{F0: (((("\n\t" + "") + ("" + "a")) + ("hello" + "")) + ""), F1: "foo bar"}
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS2{F0: 2, F1: (fi6 << ((65535) & 63)), F2: 42}
        v8 := []int{}
        v9 := GenS3{F0: 0}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    switch "a" {
    case "foo bar":
        var v10 int16 = -32768
        _ = v10
    case "":
        v11 := GenS4{F0: -303, F1: 18446744073709551615, F2: 314}
        _ = v11
    }
    v12 := []int{0, 4294967295, 32767}
    if (0.1 != (0.1 - 1e10)) {
        v13 := GenS2{F0: 2, F1: -9223372036854775808, F2: 848}
        _ = v13
    }
    _ = v12
    _ = v5
    return -1
}

func main() {
    var v1 float64 = float64(0.1)
    v2 := make(map[string]bool)
    var v3 uint8 = 255
    v4 := make(map[int]string)
    switch 110 {
    case 2147483647:
        fmt.Println(v3, v3)
    case 200:
        if (1 > 1) {
            switch "\n\t" {
            case "":
                for fi14 := 0; fi14 < 15; fi14++ {
                    var v15 float32 = 1.5
                    v16 := GenS5{F0: 117}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "a":
                var v17 int32 = -2147483648
                _ = v17
            default:
                v3 = (^(v3))
            }
            fmt.Println(v3, v3)
            for fi18 := 0; fi18 < 4; fi18++ {
                v19 := []int{}
                if (false || (-777 <= -3)) {
                    var v20 uint64 = 18446744073709551615
                    _ = v20
                } else {
                    if true {
                        var v21 string = "hello"
                        v3 = (v3 + (^(v3)))
                        var v22 float64 = (v1 * float64(3.14159))
                        _ = v21
                        _ = v22
                    }
                }
                for fi23 := 0; fi23 < 11; fi23++ {
                    switch (fi18 ^ (fi18 << ((0) & 63))) {
                    case -9223372036854775808:
                        delete(v2, "hello")
                    case -836:
                        switch "\n\t" {
                        case "hello":
                            var v24 int = fi18
                            v24 = 9223372036854775807
                            v24 -= 127
                            _ = v24
                        case "":
                            if (!(!true)) {
                                switch (fi23 & 1) {
                                case 0:
                                    if false {
                                        if false {
                                            v25 := GenS1{F0: "foo bar", F1: ("a" + ("" + ("foo bar" + "foo bar")))}
                                            v3 = (v3 >> ((1) & 63))
                                            _ = v25
                                        }
                                        var v26 uint = 1
                                        _ = v26
                                    } else {
                                        v1 += float64(1e10)
                                    }
                                default:
                                    fmt.Println(fi23, v1, fi23)
                                }
                                if false {
                                    var v27 uint32 = 42
                                    var v28 uint = 230
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v29 uint8 = v3
                                    v1 = 1.5
                                    _ = v29
                                }
                            } else {
                                var v30 uint32 = 2
                                _ = v30
                            }
                        default:
                            v3 = 127
                        }
                    case 4294967295:
                        v1 += ((v1 / v1) + v1)
                    }
                    switch "foo bar" {
                    case "foo bar":
                        switch fi18 {
                        case 712:
                            v31 := GenS6{F0: -9223372036854775808, F1: 32767, F2: 128}
                            _ = v31
                        default:
                            var v32 uint8 = 128
                            _ = v32
                        }
                    case "":
                        var v33 uint8 = (v3 & v3)
                        _ = v33
                    case "a":
                        var v34 uint8 = v3
                        _ = v34
                    }
                    _ = fi23
                }
                _ = fi18
                _ = v19
            }
        } else {
            fmt.Println(v3, v1)
            var v35 uint64 = uint64(v3)
            if true {
                for fi36 := 0; fi36 < 19; fi36++ {
                    for fi37 := 0; fi37 < 2; fi37++ {
                        _ = len(v4)
                        _ = fi37
                    }
                    fmt.Println(v1, fi36, v35)
                    var v38 uint = 65535
                    _ = fi36
                    _ = v38
                }
                switch "a" {
                case "a":
                    var v39 bool = false
                    _ = v39
                case "hello":
                    var v40 uint64 = v35
                    _ = v40
                default:
                    fmt.Println(v3, v35, v3)
                }
            } else {
                v41 := GenS6{F0: 1, F1: 1, F2: 128}
                v35 += (v35 * 2)
                _ = v41
            }
            _ = v35
        }
    default:
        if (2147483648 != 1) {
            delete(v2, "a")
        }
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        fmt.Println(v1)
        _ = fi42
    }
    for fi43 := 0; fi43 < 13; fi43++ {
        v44 := []int{fi43, fi43}
        var v45 int = (fi43 ^ (fi43 << ((750) & 63)))
        v3 -= v3
        _ = fi43
        _ = v44
        _ = v45
    }
    v1 = v1
    if true {
        var v46 uint16 = 42
        var v47 uint = 879
        switch 1 {
        case -335:
            if true {
                v46 = (^(v46))
                v48 := GenS4{F0: -709, F1: 1, F2: 2147483647}
                _ = v48
            } else {
                var v49 uint32 = 1
                v46 += (^(v46))
                v50 := GenS6{F0: -9223372036854775808, F1: 42, F2: (^(v3))}
                _ = v49
                _ = v50
            }
        case 1:
            var v51 uint32 = 65535
            _ = v51
        default:
            v46 = 0
        }
        _ = v46
        _ = v47
    } else {
        var v52 int64 = -1
        _ = v52
    }
    var v53 uint16 = 611
    if (!(128 > (v3 * v3))) {
        v54 := []int{9223372036854775807}
        _ = len(v2)
        v55 := GenS4{F0: -1, F1: 1, F2: 2147483647}
        _ = v54
        _ = v55
    } else {
        v3 = v3
    }
    v56 := v2[""]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v53:", v53)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v56
}
