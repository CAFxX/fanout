package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int; F1 uint16 }
func genFn0(p0 int) string {
    fmt.Println(p0)
    var v4 uint = 692
    for fi5 := 0; fi5 < 4; fi5++ {
        if false {
            for fi6 := 0; fi6 < 5; fi6++ {
                v7 := GenS1{F0: -9223372036854775808}
                switch (p0 ^ (fi5 & fi6)) {
                case 32767:
                    v8 := []int{p0, (fi6 - fi5), fi5}
                    _ = v8
                }
                _ = fi6
                _ = v7
            }
        }
        _ = fi5
    }
    _ = p0
    _ = v4
    return ("hello" + ("" + "hello"))
}
func genFn1(p0 int) int16 {
    p0 = (p0 % ((p0) | 1))
    fmt.Println(p0)
    _ = p0
    return 184
}

func main() {
    var v1 uint = 1
    var v2 int64 = 42
    var v3 string = ("" + ("\n\t" + "foo bar"))
    var v9 string = v3
    if false {
        var v10 int16 = 0
        _ = v10
    } else {
        if true {
            for fi11 := 0; fi11 < 11; fi11++ {
                if false {
                    fmt.Println(v2)
                    v2 -= (v2 << ((v1) & 63))
                    if true {
                        v12 := []int{-2147483648, 0, int(v2), fi11}
                        _ = v12
                    } else {
                        fmt.Println(v1, v9, v9)
                        if false {
                            v13 := GenS1{F0: v2}
                            v3 = v3
                            if (fi11 != fi11) {
                                v3 = v3
                                if (false || ((v1 + v1) < v1)) {
                                    v14 := GenS2{F0: fi11}
                                    if false {
                                        v15 := []int{-9223372036854775808, (fi11 * (fi11 ^ fi11)), fi11}
                                        if len(v15) > 0 {
                                            v15[0] = fi11
                                        }
                                        _ = v15
                                    } else {
                                        v16 := []int{}
                                        for fi17 := 0; fi17 < 19; fi17++ {
                                            var v18 int32 = 1
                                            if false {
                                                var v19 uint8 = 54
                                                if len(v16) >= 2 {
                                                    v16 = v16[1:len(v16)]
                                                }
                                                _ = v19
                                            }
                                            _ = fi17
                                            _ = v18
                                        }
                                        _ = v16
                                    }
                                    _ = v14
                                } else {
                                    v20 := []int{}
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                    v21 := GenS2{F0: (-(fi11))}
                                    _ = v20
                                    _ = v21
                                }
                            } else {
                                var v22 float32 = ((0.0 - 0.0) - 1.5)
                                v3 = ((v3 + "\n\t") + v3)
                                _ = v22
                            }
                            _ = v13
                        }
                        switch 9223372036854775807 {
                        case 0:
                            var v23 int = 120
                            _ = v23
                        }
                    }
                }
                v3 = v9
                var v24 int8 = 0
                _ = fi11
                _ = v24
            }
            for fi25 := 0; fi25 < 5; fi25++ {
                if false {
                    v26 := []int{4294967295, -388, fi25}
                    v27 := GenS1{F0: v2}
                    v1 -= 1
                    _ = v26
                    _ = v27
                }
                if true {
                    v3 = ("\n\t" + "\n\t")
                    var v28 uint16 = 2
                    var v29 uint16 = (v28 + v28)
                    _ = v28
                    _ = v29
                } else {
                    var v30 int16 = 32767
                    _ = v30
                }
                if true {
                    v1 -= (v1 * v1)
                } else {
                    for fi31 := 0; fi31 < 18; fi31++ {
                        if true {
                            var v32 uint = 18446744073709551615
                            _ = v32
                        } else {
                            v9 = (v3 + v3)
                        }
                        _ = fi31
                    }
                    v33 := []int{fi25}
                    _ = v33
                }
                _ = fi25
            }
            if (!false) {
                for fi34 := 0; fi34 < 11; fi34++ {
                    var v35 int16 = 42
                    _ = fi34
                    _ = v35
                }
                var v36 int64 = (^(v2))
                var v37 int16 = genFn1(-2147483648)
                _ = v36
                _ = v37
            } else {
                for fi38 := 0; fi38 < 6; fi38++ {
                    v9 = v9
                    v39 := []int{fi38, fi38, 1}
                    _ = fi38
                    _ = v39
                }
                if true {
                    for fi40 := 0; fi40 < 13; fi40++ {
                        v2 += (v2 ^ (v2 * 42))
                        v41 := GenS1{F0: (^(v2))}
                        if false {
                            v42 := GenS1{F0: (v2 + v2)}
                            fmt.Println(v9, v1, v2)
                            _ = v42
                        } else {
                            if true {
                                v43 := GenS2{F0: fi40}
                                if true {
                                    v2 = -9223372036854775808
                                }
                                v44 := GenS3{F0: v2, F1: fi40}
                                _ = v43
                                _ = v44
                            }
                            v2 -= v2
                        }
                        _ = fi40
                        _ = v41
                    }
                    v45 := GenS4{F0: 0, F1: -40}
                    var v46 bool = false
                    _ = v45
                    _ = v46
                } else {
                    var v47 int16 = 42
                    v48 := GenS2{F0: 0}
                    if false {
                        if true {
                            v9 = "a"
                            v49 := []int{-208}
                            v9 = v3
                            _ = v49
                        } else {
                            var v50 int16 = (v47 | genFn1(9223372036854775807))
                            _ = v50
                        }
                        var v51 uint = v1
                        _ = v51
                    } else {
                        switch -1 {
                        case -918:
                            var v52 bool = (42 < 18446744073709551615)
                            _ = v52
                        }
                    }
                    _ = v47
                    _ = v48
                }
                v53 := GenS4{F0: 385, F1: 116}
                _ = v53
            }
        }
    }
    fmt.Println(v9, v9, v1)
    v9 = "hello"
    var v54 float32 = ((1.5 / ((3.14159 * ((0.1 / 0.1) + (0.1 * (-2.25 / ((float32(1.5) * (float32(1e10) - (-2.25 * 0.0))) + 1.5))))) * ((0.0 + float32(-2.25)) / float32(1e10)))) - 0.1)
    v1 += (v1 + 0)
    if (2 <= 128) {
        v2 -= v2
        for fi55 := 0; fi55 < 6; fi55++ {
            var v56 uint32 = 42
            var v57 int16 = 42
            v58 := GenS4{F0: 695, F1: 0}
            _ = fi55
            _ = v56
            _ = v57
            _ = v58
        }
    } else {
        v59 := GenS5{F0: 2147483648, F1: 2}
        v2 -= v2
        v2 = v2
        _ = v59
    }
    v3 = v3
    for fi60 := 0; fi60 < 17; fi60++ {
        v61 := []int{fi60, fi60, fi60, fi60}
        fmt.Println(fi60)
        _ = fi60
        _ = v61
    }
    for fi62 := 0; fi62 < 19; fi62++ {
        switch ("" + (v9 + v3)) {
        case "":
            fmt.Println(v2)
        default:
            switch fi62 {
            case -96:
                switch fi62 {
                case -595:
                    v63 := GenS6{F0: fi62, F1: 65535}
                    _ = v63
                default:
                    v64 := GenS3{F0: (v2 >> ((42) & 63)), F1: fi62}
                    _ = v64
                }
            }
        }
        var v65 uint32 = 355
        _ = fi62
        _ = v65
    }
    var v66 int = -938
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("v66:", v66)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v66
    _ = v9
}
