package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 string; F1 uint; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 uint64 }
func genFn0() uint64 {
    var v4 uint16 = 2
    if false {
        fmt.Println(v4)
        v5 := []int{9223372036854775807, 0}
        _ = v5
    }
    v4 = (v4 & v4)
    fmt.Println(v4)
    v4 = (v4 - (v4 << ((42) & 63)))
    _ = v4
    return 39
}

func main() {
    var v1 bool = false
    var v2 uint16 = 1
    v3 := []int{128, 32768, 42}
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 1}
            v9 := GenS1{F0: -1}
            fmt.Println(fi6, fi7, v1)
            _ = fi7
            _ = v8
            _ = v9
        }
        if (42 > -1) {
            fmt.Println(v2, v1, v1)
            for fi10 := 0; fi10 < 5; fi10++ {
                v2 -= v2
                _ = fi10
            }
        } else {
            var v11 int = fi6
            _ = v11
        }
        _ = fi6
    }
    v2 += (v2 >> ((808) & 63))
    v12 := GenS2{F0: (v2 * (v2 << ((2) & 63))), F1: 733}
    if (!true) {
        var v13 uint = 2
        if (!((0.1 - ((float32(1.5) / 0.1) - 0.0)) == 1e10)) {
            if len(v3) > 0 {
                v3[0] = 32767
            }
        }
        if v1 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = v13
    } else {
        for fi14 := 0; fi14 < 18; fi14++ {
            v2 = 992
            v15 := GenS3{F0: "foo bar", F1: 255, F2: 42}
            _ = fi14
            _ = v15
        }
        var v16 string = "hello"
        var v17 int = -1
        _ = v16
        _ = v17
    }
    if true {
        switch 32767 {
        case 2147483647:
            v2 = (v2 << ((4294967295) & 63))
        case -903:
            v1 = v1
        default:
            switch 128 {
            case 32767:
                v1 = v1
            }
        }
    }
    var v18 bool = v1
    var v19 float64 = (3.14159 * (0.0 - 1e10))
    if false {
        var v20 float32 = 0.0
        v1 = (!v1)
        var v21 uint32 = 546
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 17; fi22++ {
            if (v18 || v1) {
                var v23 int8 = 100
                if len(v3) > 0 {
                    v3[0] = fi22
                }
                _ = v23
            }
            var v24 bool = (0.0 < (-2.25 - (0.0 + (-2.25 + ((float32(0.0) * 1e10) / 0.1)))))
            if (1 <= 529) {
                for fi25 := 0; fi25 < 7; fi25++ {
                    if v18 {
                        if (!(v2 < v2)) {
                            var v26 uint64 = 854
                            var v27 uint64 = (v26 ^ genFn0())
                            var v28 uint = 65535
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            if (478 >= 396) {
                                if (42 > 1) {
                                    if (!(611 != 42)) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v1 = v1
                                        var v29 uint64 = 4294967295
                                        _ = v29
                                    }
                                    v18 = true
                                    switch 127 {
                                    case -2147483648:
                                        if ((!(!(4294967295 != 118))) || (true || v18)) {
                                            if true {
                                                if len(v3) > 0 {
                                                    v3[0] = 32767
                                                }
                                                if (!(v19 < -2.25)) {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if v1 {
                                                        var v30 uint64 = 42
                                                        _ = v30
                                                    }
                                                } else {
                                                    var v31 int16 = int16(fi25)
                                                    var v32 int32 = 1
                                                    switch -9223372036854775808 {
                                                    case 32768:
                                                        if (!(-2147483648 != 445)) {
                                                            v33 := GenS1{F0: v31}
                                                            if (922 < 2) {
                                                                v34 := GenS3{F0: "", F1: 1, F2: 201}
                                                                var v35 int32 = (v32 << ((42) & 63))
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                            fmt.Println(v18, v19, v31)
                                                            _ = v33
                                                        } else {
                                                            switch -495 {
                                                            case 0:
                                                                var v36 uint64 = 450
                                                                _ = v36
                                                            }
                                                            fmt.Println(v1, v19, v24)
                                                            if true {
                                                                v3 = append(v3, (fi25 | fi25))
                                                            }
                                                        }
                                                    }
                                                    _ = v31
                                                    _ = v32
                                                }
                                                var v37 int32 = 1
                                                _ = v37
                                            }
                                        } else {
                                            var v38 int32 = 0
                                            v39 := GenS1{F0: 1000}
                                            var v40 int64 = -553
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        }
                                    }
                                } else {
                                    var v41 uint64 = genFn0()
                                    v42 := GenS1{F0: int16(fi22)}
                                    if (!((0 < 0) && ((v18 && v18) && v1))) {
                                        v43 := GenS4{F0: (3.14159 < ((float32(v19) * (((-2.25 * (0.0 - (1.5 + -2.25))) / 1.0) - 0.0)) - 0.1)), F1: v2, F2: 2147483647}
                                        var v44 float64 = (1e10 / ((v19 * -2.25) + 0.0))
                                        _ = v43
                                        _ = v44
                                    } else {
                                        v3 = append(v3, (fi25 % (((fi22 >> ((694) & 63))) | 1)))
                                        v45 := GenS5{F0: v1}
                                        var v46 int16 = 1
                                        _ = v45
                                        _ = v46
                                    }
                                    _ = v41
                                    _ = v42
                                }
                                v47 := GenS3{F0: "foo bar", F1: 42, F2: 85}
                                if (("" == "a") || (!(0 > (fi22 >> ((2) & 63))))) {
                                    switch fi22 {
                                    case 0:
                                        if v24 {
                                            switch "" {
                                            case "hello":
                                                v48 := GenS1{F0: 32767}
                                                _ = v48
                                            case "foo bar":
                                                var v49 int16 = -1
                                                _ = v49
                                            default:
                                                var v50 float32 = 3.14159
                                                _ = v50
                                            }
                                        } else {
                                            v19 = (((-2.25 - float64(-2.25)) * v19) - 0.0)
                                            var v51 int32 = 2147483647
                                            var v52 uint8 = 127
                                            _ = v51
                                            _ = v52
                                        }
                                    }
                                    v53 := GenS6{F0: 32767, F1: 676}
                                    _ = v53
                                }
                                var v54 float64 = 0.1
                                _ = v47
                                _ = v54
                            }
                            v2 += (v2 ^ (v2 * 679))
                        }
                    }
                    var v55 int32 = 0
                    _ = fi25
                    _ = v55
                }
            } else {
                v24 = false
            }
            _ = fi22
            _ = v24
        }
    }
    var v56 int16 = -1
    if false {
        _ = len(v3)
        v3 = append(v3, 0)
    }
    for fi57 := 0; fi57 < 10; fi57++ {
        for fi58 := 0; fi58 < 5; fi58++ {
            fmt.Println(v2, v56, v19)
            _ = fi58
        }
        v1 = false
        _ = fi57
    }
    v59 := GenS7{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v56
    _ = v59
}
