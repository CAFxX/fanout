package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int; F2 uint8 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 bool; F1 int32; F2 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = -1
    var v3 uint64 = 1
    var v4 float32 = 1e10
    var v5 bool = (((322 < 2) || true) && (0.1 > 1e10))
    v6 := make(map[uint32]bool)
    if (255 < 2) {
        fmt.Println(v2, v4)
        if (!((4294967295 != 32767) && true)) {
            for fi7 := 0; fi7 < 17; fi7++ {
                for fi8 := 0; fi8 < 7; fi8++ {
                    v9 := []int{}
                    _ = len(v6)
                    _ = fi8
                    _ = v9
                }
                fmt.Println(v4, v4, v5)
                _ = fi7
            }
        }
    }
    if v5 {
        var v10 float64 = float64((1e10 / v4))
        v10 = float64(v4)
        v3 -= v3
        _ = v10
    }
    v2 = (v2 / ((v2) | 1))
    v2 -= v2
    var v11 int8 = -1
    if false {
        v4 -= 0.0
        v12 := GenS1{F0: 65535, F1: "a"}
        _ = v12
    } else {
        fmt.Println(v2)
        if false {
            v13 := GenS2{F0: 1}
            if (520 <= 255) {
                for fi14 := 0; fi14 < 19; fi14++ {
                    switch (("\n\t" + "a") + ("a" + ("" + "foo bar"))) {
                    case "foo bar":
                        if v5 {
                            if ((!(0 == 100)) && false) {
                                v11 = 0
                                var v15 uint32 = 2147483647
                                v5 = (42 == 65535)
                                _ = v15
                            }
                            if v5 {
                                fmt.Println(v11, v11)
                                v16 := v1[((((("foo bar" + "\n\t") + "a") + "foo bar") + "a") + "\n\t")]
                                for fi17 := 0; fi17 < 19; fi17++ {
                                    var v18 float64 = 0.1
                                    _ = fi17
                                    _ = v18
                                }
                                _ = v16
                            }
                        }
                    }
                    switch (("foo bar" + "a") + "\n\t") {
                    case "a":
                        var v19 uint8 = 117
                        _ = v19
                    case "foo bar":
                        if true {
                            if (100 == (v11 + v11)) {
                                switch "hello" {
                                case "a":
                                    if (367 >= uint(v3)) {
                                        if false {
                                            v2 += (^(v2))
                                            if (!v5) {
                                                for fi20 := 0; fi20 < 10; fi20++ {
                                                    if true {
                                                        v21 := GenS1{F0: 255, F1: "foo bar"}
                                                        var v22 int8 = (v11 % ((v11) | 1))
                                                        v6[65535] = (834 != int32(fi14))
                                                        _ = v21
                                                        _ = v22
                                                    } else {
                                                        v23 := []int{1, fi14}
                                                        var v24 float64 = 1.5
                                                        _ = v23
                                                        _ = v24
                                                    }
                                                    _ = fi20
                                                }
                                            }
                                        }
                                        var v25 int = (-(fi14))
                                        v26 := GenS3{F0: ("\n\t" + ("a" + "")), F1: int(v2), F2: 1}
                                        _ = v25
                                        _ = v26
                                    } else {
                                        fmt.Println(v5, v3, v3)
                                    }
                                }
                                v11 += v11
                                if v5 {
                                    switch (fi14 >> ((121) & 63)) {
                                    case 42:
                                        _ = len(v1)
                                    case 643:
                                        v4 = v4
                                    default:
                                        v27 := GenS1{F0: 17, F1: "a"}
                                        _ = v27
                                    }
                                    for fi28 := 0; fi28 < 1; fi28++ {
                                        v29 := GenS4{F0: 65535, F1: (v2 / ((v2) | 1))}
                                        _ = fi28
                                        _ = v29
                                    }
                                } else {
                                    var v30 uint = 2
                                    v11 += -128
                                    _ = v30
                                }
                            }
                            var v31 int16 = -131
                            fmt.Println(v2)
                            _ = v31
                        } else {
                            v32 := GenS5{F0: (v4 == v4), F1: -1, F2: 36}
                            _ = v32
                        }
                    }
                    _ = fi14
                }
                if v5 {
                    switch 32768 {
                    case 42:
                        for fi33 := 0; fi33 < 17; fi33++ {
                            var v34 uint16 = 42
                            if ((^(v2)) != 1) {
                                fmt.Println(fi33, v4)
                                v34 += (v34 / (((v34 - v34)) | 1))
                                fmt.Println(v3, v11)
                            }
                            _ = fi33
                            _ = v34
                        }
                    case 4294967295:
                        v35 := []int{2147483647, -1, 0}
                        _ = v35
                    case 127:
                        v4 -= ((3.14159 + (v4 / (v4 * (v4 - v4)))) - (3.14159 - (v4 + float32(0.1))))
                    }
                } else {
                    v36 := GenS6{F0: false}
                    v5 = v5
                    _ = v36
                }
                fmt.Println(v4)
            }
            _ = v13
        }
    }
    v4 -= v4
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        if v5 {
            var v37 uint64 = (^(v3))
            v38 := GenS6{F0: true}
            var v39 int = 9223372036854775807
            _ = v37
            _ = v38
            _ = v39
        } else {
            var v40 string = ("" + ("foo bar" + "a"))
            v1["\n\t"] = 355
            _ = v40
        }
    case "foo bar":
        v5 = (v3 >= (v3 & 1))
    case "hello":
        fmt.Println(v4, v11)
    }
    v11 = (v11 << ((409) & 63))
    for fi41 := 0; fi41 < 4; fi41++ {
        if v5 {
            v42 := GenS5{F0: v5, F1: -287, F2: 1}
            _ = v42
        }
        if (1 <= 4294967295) {
            switch 4294967295 {
            case -233:
                fmt.Println(fi41, v2)
            case -9:
                for fi43 := 0; fi43 < 1; fi43++ {
                    var v44 uint64 = (v3 % (((v3 * 0)) | 1))
                    _ = fi43
                    _ = v44
                }
            case 9223372036854775807:
                fmt.Println(v5, v11)
            default:
                fmt.Println(v5, fi41)
            }
            for fi45 := 0; fi45 < 4; fi45++ {
                switch "hello" {
                case "hello":
                    v3 = (v3 >> ((42) & 63))
                case "\n\t":
                    v46 := GenS6{F0: v5}
                    _ = v46
                case "":
                    var v47 float64 = ((-2.25 * -2.25) / (0.1 - (0.0 + 1.5)))
                    _ = v47
                }
                v48 := GenS7{F0: 1, F1: 744}
                _ = fi45
                _ = v48
            }
        }
        _ = fi41
    }
    for fi49 := 0; fi49 < 4; fi49++ {
        for fi50 := 0; fi50 < 5; fi50++ {
            var v51 uint8 = 1
            var v52 float64 = (0.1 - ((1e10 / -2.25) + 0.0))
            fmt.Println(v5, v4)
            v53 := GenS5{F0: ((fi49 >> ((106) & 63)) >= 256), F1: int32(v3), F2: v51}
            _ = len(v1)
            var v54 bool = (false || false)
            _ = fi50
            _ = v51
            _ = v52
            _ = v53
            _ = v54
        }
        _ = fi49
    }
    v55 := GenS4{F0: 4294967295, F1: 42}
    v11 -= v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
