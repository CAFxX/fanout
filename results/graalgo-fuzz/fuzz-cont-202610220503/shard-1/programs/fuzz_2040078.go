package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 string }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint; F2 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool; F1 int64 }
func genFn0(p0 uint8) int32 {
    p0 += (p0 * 42)
    p0 += p0
    _ = p0
    return 42
}
func genFn1() int16 {
    var v4 float64 = 0.0
    for fi5 := 0; fi5 < 4; fi5++ {
        switch "foo bar" {
        case "\n\t":
            v4 = v4
        case "hello":
            var v6 uint8 = 203
            _ = v6
        case "foo bar":
            v4 = (v4 + float64(0.0))
        default:
            if (fi5 != fi5) {
                fmt.Println(v4)
                switch fi5 {
                case 9223372036854775807:
                    var v7 int8 = -128
                    _ = v7
                case 318:
                    v8 := GenS1{F0: 16, F1: 1, F2: ((("a" + "\n\t") + "foo bar") + "a")}
                    _ = v8
                default:
                    v9 := []int{}
                    _ = v9
                }
            }
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            if true {
                v11 := []int{(fi5 >> ((2) & 63)), fi10}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = v11
            } else {
                v12 := GenS1{F0: 18446744073709551615, F1: 1, F2: ""}
                if false {
                    switch (fi5 - (fi5 % ((127) | 1))) {
                    case -2147483648:
                        var v13 int8 = 119
                        _ = v13
                    case -1:
                        v14 := GenS2{F0: -603, F1: (fi10 >> ((1) & 63))}
                        _ = v14
                    }
                    v4 -= v4
                }
                if false {
                    v15 := GenS3{F0: fi10, F1: 1000, F2: 18446744073709551615}
                    v16 := GenS4{F0: 340, F1: 496, F2: -379}
                    v17 := GenS4{F0: 374, F1: 503, F2: 9223372036854775807}
                    _ = v15
                    _ = v16
                    _ = v17
                }
                _ = v12
            }
            _ = fi10
        }
        v18 := []int{0}
        _ = fi5
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 uint8 = 1
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := []int{(fi19 * (fi19 | fi21))}
            v4 = (v4 / 0.1)
            fmt.Println(v20)
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    var v23 int8 = 42
    _ = v23
    _ = v4
    return 207
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 2
    var v3 float32 = (-2.25 + ((1e10 - ((-2.25 * (1e10 / ((1e10 / 0.1) - -2.25))) + (1.5 / 1.5))) + 1e10))
    for fi24 := 0; fi24 < 11; fi24++ {
        if false {
            var v25 uint16 = 701
            delete(v1, (fi24 << ((964) & 63)))
            if ((42 == 762) && false) {
                v2 += v2
            }
            _ = v25
        }
        v3 -= (float32(0.0) + (v3 / 1.0))
        var v26 uint32 = 42
        _ = fi24
        _ = v26
    }
    if true {
        switch -1 {
        case 32768:
            var v27 string = "a"
            _ = v27
        case 9223372036854775807:
            var v28 uint64 = 1
            _ = v28
        case 756:
            fmt.Println(v3, v3)
        default:
            for fi29 := 0; fi29 < 19; fi29++ {
                v30 := []int{fi29}
                var v31 int8 = 100
                var v32 uint16 = 65535
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
        }
        v3 += (v3 * 0.1)
    } else {
        _ = len(v1)
        switch -93 {
        case 9223372036854775807:
            v33 := GenS5{F0: 65535}
            _ = v33
        case 128:
            var v34 int = 772
            _ = v34
        case -597:
            var v35 uint = 953
            _ = v35
        default:
            v36 := []int{288, 32768, -1}
            _ = v36
        }
        v3 = v3
    }
    v3 += ((((float32(float64(v3)) + v3) / (v3 * 1e10)) * 0.1) / v3)
    switch 42 {
    case -1:
        v3 = v3
    case 194:
        if (0 == 128) {
            v2 = (^(v2))
            var v37 bool = ((42 >= 1) && (true && false))
            _ = v37
        }
    case -941:
        fmt.Println(v2)
    default:
        var v38 int32 = 806
        _ = v38
    }
    v39 := GenS1{F0: 616, F1: 255, F2: "foo bar"}
    fmt.Println(v3, v2)
    if (!(!true)) {
        v40 := GenS6{F0: 42, F1: 1}
        if false {
            v41 := GenS4{F0: 32767, F1: 764, F2: -9223372036854775808}
            _ = v41
        } else {
            v42 := []int{0, -443, 4294967295, -9223372036854775808}
            _ = v42
        }
        _ = v40
    }
    var v43 float32 = ((1e10 + (v3 + (0.1 / v3))) - (float32(0.0) * v3))
    v44 := GenS7{F0: 18446744073709551615}
    v2 = v2
    for fi45 := 0; fi45 < 14; fi45++ {
        _ = len(v1)
        fmt.Println(v2, fi45)
        _ = fi45
    }
    for fi46 := 0; fi46 < 12; fi46++ {
        switch "a" {
        case "\n\t":
            var v47 uint32 = (v2 % (((v2 + 42)) | 1))
            _ = v47
        case "a":
            switch ("hello" + "foo bar") {
            case "a":
                v48 := GenS2{F0: 370, F1: (fi46 % ((128) | 1))}
                _ = v48
            case "":
                var v49 uint32 = 836
                _ = v49
            default:
                fmt.Println(v43)
            }
        case "hello":
            if false {
                for fi50 := 0; fi50 < 3; fi50++ {
                    v51 := GenS7{F0: 1}
                    v52 := GenS8{F0: ("foo bar" == (("" + "hello") + (("" + ("a" + "\n\t")) + "hello"))), F1: 9223372036854775807}
                    fmt.Println(fi50, v3, v2)
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                v1[fi46] = 1
                var v53 uint32 = v2
                _ = v53
            }
        default:
            if ((!false) || true) {
                var v54 uint16 = 0
                v55 := []int{(fi46 & fi46)}
                var v56 float32 = v3
                _ = v54
                _ = v55
                _ = v56
            }
        }
        v2 += (v2 >> ((18446744073709551615) & 63))
        var v57 int16 = 1000
        _ = fi46
        _ = v57
    }
    for fi58 := 0; fi58 < 15; fi58++ {
        v2 = 235
        _ = fi58
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v43
    _ = v44
}
