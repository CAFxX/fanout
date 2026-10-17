package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 string; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 bool; F2 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint; F1 uint16 }
func genFn0() int8 {
    v4 := GenS1{F0: 0, F1: 42}
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int = fi5
        var v7 uint = 212
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint8 = 128
    v8 = (v8 * v8)
    v8 += 2
    _ = v4
    _ = v8
    return -128
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 668
    var v3 uint64 = 2
    for fi9 := 0; fi9 < 8; fi9++ {
        if false {
            for fi10 := 0; fi10 < 7; fi10++ {
                if true {
                    v11 := GenS2{F0: 128, F1: v3}
                    v1["foo bar"] = fi10
                    fmt.Println(v2)
                    _ = v11
                }
                v2 -= v2
                _ = fi10
            }
            for fi12 := 0; fi12 < 19; fi12++ {
                switch "" {
                case "\n\t":
                    v13 := GenS1{F0: int32(v3), F1: (^(v3))}
                    _ = v13
                }
                v2 = 9223372036854775807
                _ = fi12
            }
        } else {
            v2 += (v2 ^ 881)
        }
        _ = fi9
    }
    v14 := GenS2{F0: 127, F1: 610}
    v2 += (v2 % (((v2 ^ (v2 ^ v2))) | 1))
    if true {
        var v15 int16 = 1000
        v16 := []int{9223372036854775807, 128}
        _ = v15
        _ = v16
    }
    v17 := GenS3{F0: "hello", F1: (v2 / ((v2) | 1)), F2: 0}
    if (0 > 255) {
        var v18 uint = 18446744073709551615
        switch -2147483648 {
        case 128:
            switch 32767 {
            case -498:
                if false {
                    v19 := GenS2{F0: 2, F1: uint64(v18)}
                    for fi20 := 0; fi20 < 15; fi20++ {
                        var v21 uint32 = 0
                        var v22 uint = (v18 << ((v18) & 63))
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 6; fi23++ {
                        v1["a"] = 991
                        _ = fi23
                    }
                    _ = v19
                }
            case 42:
                if false {
                    if (v3 == v3) {
                        _ = len(v1)
                    } else {
                        v3 -= v3
                    }
                    v24 := []int{128}
                    _ = v24
                } else {
                    v18 = (v18 << ((0) & 63))
                }
            case 840:
                for fi25 := 0; fi25 < 6; fi25++ {
                    var v26 int32 = 0
                    _ = fi25
                    _ = v26
                }
            default:
                v27 := []int{32767, 152, 32768, 2147483647}
                _ = v27
            }
        case -9223372036854775808:
            switch ("hello" + "hello") {
            case "\n\t":
                var v28 int8 = -128
                _ = v28
            case "a":
                v3 = (v3 ^ (v3 & v3))
            default:
                if true {
                    fmt.Println(v2)
                    v18 = (v18 / ((v18) | 1))
                }
            }
        }
        if ((!false) || (1.5 != 1.5)) {
            var v29 float64 = ((0.1 + 3.14159) + -2.25)
            _ = v29
        }
        _ = v18
    } else {
        v30 := []int{2147483647, -754, -509}
        _ = v30
    }
    if (124 == 100) {
        for fi31 := 0; fi31 < 2; fi31++ {
            if (false && true) {
                _ = len(v1)
            } else {
                v2 = int64(v3)
                var v32 uint64 = (v3 + 42)
                _ = v32
            }
            for fi33 := 0; fi33 < 1; fi33++ {
                var v34 float32 = 3.14159
                if false {
                    var v35 string = ("\n\t" + "hello")
                    _ = v35
                }
                _ = fi33
                _ = v34
            }
            switch fi31 {
            case -1:
                fmt.Println(v3)
            case 899:
                for fi36 := 0; fi36 < 5; fi36++ {
                    v37 := GenS1{F0: int32(v2), F1: (^(v3))}
                    _ = fi36
                    _ = v37
                }
            case 32768:
                v38 := []int{fi31, -548}
                _ = v38
            default:
                if false {
                    var v39 uint = 18446744073709551615
                    var v40 string = ""
                    v41 := []int{(fi31 * (-(fi31))), (fi31 & fi31), 167, (fi31 | (fi31 ^ fi31))}
                    _ = v39
                    _ = v40
                    _ = v41
                } else {
                    switch (fi31 | (fi31 + (fi31 + (fi31 / ((fi31) | 1))))) {
                    case 4294967295:
                        fmt.Println(fi31, fi31)
                    case 42:
                        delete(v1, ("" + ("hello" + "")))
                    default:
                        var v42 uint32 = 2
                        _ = v42
                    }
                }
            }
            _ = fi31
        }
        v43 := GenS4{F0: 195, F1: 1, F2: 252}
        _ = v43
    } else {
        if false {
            v44 := []int{128, -398}
            v45 := GenS5{F0: v2}
            var v46 uint16 = 65535
            _ = v44
            _ = v45
            _ = v46
        } else {
            var v47 uint8 = 1
            v3 -= 1
            delete(v1, ("foo bar" + "a"))
            _ = v47
        }
        var v48 uint = 4294967295
        var v49 int16 = -1
        _ = v48
        _ = v49
    }
    fmt.Println(v3, v3)
    v50 := GenS3{F0: "", F1: v2, F2: 1}
    for fi51 := 0; fi51 < 3; fi51++ {
        v52 := []int{fi51, fi51, fi51, fi51}
        if (561 != 42) {
            v2 += (v2 & v2)
            var v53 uint64 = (v3 >> ((42) & 63))
            _ = v53
        } else {
            v54 := GenS6{F0: false, F1: (1 >= 867), F2: 42}
            for fi55 := 0; fi55 < 16; fi55++ {
                v3 = (v3 + v3)
                v56 := GenS7{F0: int32(v3)}
                var v57 float32 = 1.5
                var v58 uint16 = 2
                switch (fi51 / (((fi51 << ((18446744073709551615) & 63))) | 1)) {
                case 4294967295:
                    if (true || false) {
                        fmt.Println(v2, v2)
                        _ = len(v52)
                        v59 := GenS8{F0: 4294967295, F1: (v58 + v58)}
                        _ = v59
                    } else {
                        v58 -= (v58 / ((32767) | 1))
                    }
                default:
                    v57 += v57
                }
                var v60 int8 = 1
                _ = fi55
                _ = v56
                _ = v57
                _ = v58
                _ = v60
            }
            _ = v54
        }
        var v61 int32 = -1
        _ = fi51
        _ = v52
        _ = v61
    }
    for fi62 := 0; fi62 < 5; fi62++ {
        var v63 int64 = (v2 * v2)
        _ = fi62
        _ = v63
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v50
}
