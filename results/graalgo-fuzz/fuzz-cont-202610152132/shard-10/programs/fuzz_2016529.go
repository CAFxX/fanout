package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint16 }
func genFn0() uint64 {
    var v4 int8 = 127
    if true {
        v4 -= 1
        if (1 == 9223372036854775807) {
            for fi5 := 0; fi5 < 11; fi5++ {
                v6 := GenS1{F0: 32767, F1: (42 < 185)}
                _ = fi5
                _ = v6
            }
        } else {
            var v7 float64 = 1.5
            _ = v7
        }
    } else {
        var v8 int8 = v4
        for fi9 := 0; fi9 < 11; fi9++ {
            v4 -= (v4 / (((v4 | -10)) | 1))
            v10 := []int{fi9, 2147483647, -9223372036854775808, fi9}
            switch fi9 {
            case -2147483648:
                var v11 uint8 = 92
                _ = v11
            case 2147483647:
                var v12 int8 = (v8 & (-(v4)))
                _ = v12
            case 830:
                var v13 float64 = 1.5
                _ = v13
            default:
                v8 = v4
            }
            _ = fi9
            _ = v10
        }
        if (18446744073709551615 < 4294967295) {
            if (false && (!true)) {
                v14 := []int{-9223372036854775808}
                v8 -= v8
                if ((!false) || (-2147483648 < 42)) {
                    var v15 int16 = 1000
                    var v16 float64 = 1.5
                    var v17 uint32 = 2147483647
                    _ = v15
                    _ = v16
                    _ = v17
                }
                _ = v14
            }
            for fi18 := 0; fi18 < 7; fi18++ {
                var v19 int16 = -1
                if true {
                    var v20 string = "a"
                    for fi21 := 0; fi21 < 18; fi21++ {
                        var v22 float64 = -2.25
                        var v23 uint = 65535
                        v24 := GenS1{F0: (v19 | (-(v19))), F1: (v23 == 899)}
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    _ = v20
                } else {
                    if false {
                        if ((65535 >= 699) || true) {
                            var v25 string = "a"
                            var v26 uint64 = 4294967295
                            _ = v25
                            _ = v26
                        }
                    } else {
                        for fi27 := 0; fi27 < 6; fi27++ {
                            var v28 float32 = (float32(0.0) / 1.0)
                            _ = fi27
                            _ = v28
                        }
                        if true {
                            fmt.Println(fi18, v4, v19)
                        }
                        switch -9223372036854775808 {
                        case -730:
                            v29 := GenS1{F0: v19, F1: true}
                            _ = v29
                        case 42:
                            v8 = (v8 ^ (v4 & -1))
                        }
                    }
                }
                _ = fi18
                _ = v19
            }
            for fi30 := 0; fi30 < 20; fi30++ {
                v8 += -95
                _ = fi30
            }
        } else {
            for fi31 := 0; fi31 < 3; fi31++ {
                var v32 int64 = 1
                v33 := GenS1{F0: -742, F1: true}
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        _ = v8
    }
    v4 = -128
    for fi34 := 0; fi34 < 12; fi34++ {
        v35 := []int{-111, fi34}
        var v36 uint64 = 2
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 uint32 = 2147483648
    _ = v37
    _ = v4
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "hello"
    v3 := make(map[uint32]uint32)
    v2 = (v2 + "foo bar")
    if false {
        fmt.Println(v2)
        if false {
            v38 := v3[749]
            _ = v38
        } else {
            var v39 int64 = 857
            _ = v39
        }
        switch "\n\t" {
        case "":
            var v40 uint32 = 2
            _ = v40
        }
    } else {
        v41 := []int{410, 32767, -1}
        v42 := GenS2{F0: 333}
        _ = v41
        _ = v42
    }
    v43 := GenS2{F0: 1}
    switch "foo bar" {
    case "":
        if false {
            v2 = "\n\t"
        } else {
            v2 = ("\n\t" + v2)
            var v44 int32 = 278
            var v45 bool = (0 < 2)
            _ = v44
            _ = v45
        }
    }
    if false {
        fmt.Println(v2)
        switch v2 {
        case "\n\t":
            v2 = v2
        case "hello":
            for fi46 := 0; fi46 < 12; fi46++ {
                v47 := GenS2{F0: 65535}
                _ = fi46
                _ = v47
            }
        case "foo bar":
            if true {
                fmt.Println(v2)
            }
        }
        v48 := GenS1{F0: 114, F1: (4294967295 > 434)}
        _ = v48
    } else {
        v2 = "\n\t"
        var v49 int64 = -1
        switch 32768 {
        case 42:
            fmt.Println(v2, v2)
        case 32767:
            v49 = 42
        }
        _ = v49
    }
    if (42 <= 79) {
        if false {
            fmt.Println(v2)
            v50 := v1[0]
            v2 = v2
            _ = v50
        }
        if false {
            var v51 uint16 = 1
            v52 := GenS3{F0: 42, F1: (^(v51))}
            var v53 int64 = 42
            _ = v51
            _ = v52
            _ = v53
        }
    } else {
        if true {
            var v54 int = 0
            _ = v54
        }
        v55 := GenS3{F0: 2, F1: 255}
        _ = v55
    }
    v56 := v3[44]
    switch v2 {
    case "a":
        fmt.Println(v56, v56)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v56
}
