package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint; F1 uint; F2 int8 }
func genFn0() int32 {
    if false {
        v6 := GenS1{F0: 0}
        _ = v6
    } else {
        var v7 uint8 = 99
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: (fi8 >> ((255) & 63))}
            for fi10 := 0; fi10 < 8; fi10++ {
                v11 := GenS2{F0: "\n\t", F1: -1}
                v12 := GenS2{F0: "", F1: -648}
                var v13 float32 = 1e10
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    var v14 float32 = 1e10
    var v15 uint8 = 128
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 int32 = 2147483647
        _ = fi16
        _ = v17
    }
    _ = v14
    _ = v15
    return 2147483647
}

func main() {
    var v1 string = ""
    v2 := make(map[int]uint32)
    var v3 int64 = 42
    var v4 string = ""
    var v5 int32 = 0
    v1 = (v4 + "foo bar")
    if (!(-102 >= 100)) {
        v18 := GenS3{F0: 65535, F1: 1}
        v19 := []int{-1}
        var v20 uint32 = 28
        _ = v18
        _ = v19
        _ = v20
    } else {
        if false {
            if ((556 <= 127) && ((65535 != 174) || true)) {
                var v21 int8 = -1
                if (4294967295 >= 275) {
                    v22 := GenS4{F0: (v5 + (v5 % ((v5) | 1))), F1: "a"}
                    v23 := GenS1{F0: 1}
                    _ = v22
                    _ = v23
                } else {
                    v24 := []int{-545, -1}
                    var v25 int8 = v21
                    v26 := GenS5{F0: 1, F1: 786, F2: (v21 & (v25 >> ((131) & 63)))}
                    _ = v24
                    _ = v25
                    _ = v26
                }
                _ = v21
            } else {
                v27 := GenS2{F0: "", F1: 42}
                _ = v27
            }
        }
        fmt.Println(v4, v1, v3)
        if true {
            v1 = (v4 + v1)
            v5 = (v5 >> ((0) & 63))
            switch 975 {
            case -2147483648:
                v1 = ("a" + "")
            case 4294967295:
                for fi28 := 0; fi28 < 17; fi28++ {
                    fmt.Println(fi28, fi28)
                    _ = fi28
                }
            }
        }
    }
    if false {
        v29 := []int{477, 361, -167}
        var v30 int64 = 9223372036854775807
        _ = v29
        _ = v30
    } else {
        if (((-2.25 + 0.1) == 1.5) && (1.5 < ((0.1 * 1.5) / (((((((3.14159 * (3.14159 / 1.0)) - (1.5 / 1.0)) / 1.0) - 3.14159) * -2.25) - (-2.25 + 0.1)) - -2.25)))) {
            var v31 uint8 = 170
            _ = v31
        } else {
            v1 = ((("hello" + (v1 + ((v4 + v1) + (v1 + (v4 + "a"))))) + "\n\t") + ("\n\t" + ("hello" + "\n\t")))
            fmt.Println(v3)
            for fi32 := 0; fi32 < 7; fi32++ {
                if true {
                    v33 := []int{}
                    _ = v33
                } else {
                    v1 = ("foo bar" + v4)
                }
                for fi34 := 0; fi34 < 5; fi34++ {
                    v1 = ("" + v1)
                    switch -1 {
                    case -9223372036854775808:
                        if ((false && true) || (v1 >= (v1 + v1))) {
                            if ((190 <= 42) || false) {
                                var v35 int64 = 0
                                var v36 int = -1
                                if (!(0.1 < (1.5 - (0.0 - (3.14159 * (1.5 * (1.5 / 1e10))))))) {
                                    var v37 bool = false
                                    fmt.Println(v37, v1)
                                    if true {
                                        v37 = (423 != 257)
                                    } else {
                                        delete(v2, (v36 + (fi32 & (-(fi34)))))
                                    }
                                    _ = v37
                                }
                                _ = v35
                                _ = v36
                            }
                            var v38 int32 = v5
                            _ = v38
                        }
                    }
                    _ = fi34
                }
                v3 = v3
                _ = fi32
            }
        }
        v39 := GenS5{F0: 65535, F1: 4294967295, F2: 42}
        for fi40 := 0; fi40 < 12; fi40++ {
            if ((v5 * (v5 - (v5 | (v5 & -739)))) >= 1) {
                v41 := GenS3{F0: 677, F1: 255}
                var v42 string = "\n\t"
                _ = v41
                _ = v42
            } else {
                for fi43 := 0; fi43 < 14; fi43++ {
                    if true {
                        switch fi40 {
                        case 1:
                            if ((2 <= 1) || true) {
                                v3 -= v3
                                fmt.Println(fi40)
                            }
                        case 42:
                            v3 -= (v3 - 0)
                        }
                        v5 -= 42
                        var v44 float64 = (-2.25 - 1.5)
                        var v45 int = (^(fi43))
                        _ = v44
                        _ = v45
                    }
                    v46 := GenS5{F0: 42, F1: 42, F2: 117}
                    v47 := GenS5{F0: 421, F1: 2, F2: -1}
                    _ = fi43
                    _ = v46
                    _ = v47
                }
                v48 := GenS5{F0: 65535, F1: 4294967295, F2: 127}
                _ = v48
            }
            _ = fi40
        }
        _ = v39
    }
    for fi49 := 0; fi49 < 15; fi49++ {
        v50 := []int{(fi49 & fi49), -57}
        _ = fi49
        _ = v50
    }
    v5 += (v5 - 0)
    v3 += v3
    switch "" {
    case "a":
        for fi51 := 0; fi51 < 1; fi51++ {
            if true {
                var v52 uint = 2
                _ = v52
            }
            _ = fi51
        }
    }
    if (!(!(42 <= 573))) {
        v53 := GenS5{F0: 0, F1: 255, F2: 0}
        var v54 int = 1
        _ = v53
        _ = v54
    } else {
        var v55 uint64 = 908
        fmt.Println(v5, v55)
        v56 := GenS4{F0: (v5 ^ v5), F1: v1}
        _ = v55
        _ = v56
    }
    v1 = ((("a" + v4) + v4) + ((v4 + v1) + "foo bar"))
    v57 := GenS1{F0: 65}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
}
