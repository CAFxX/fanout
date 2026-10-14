package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 bool; F1 int8; F2 int8 }
type GenS3 struct { F0 int; F1 uint8; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 string; F2 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 bool }
type GenS8 struct { F0 uint; F1 uint }
type GenS9 struct { F0 int32; F1 int; F2 int }
func genFn0(p0 int) int32 {
    for fi4 := 0; fi4 < 7; fi4++ {
        p0 -= p0
        var v5 float32 = 1.5
        if (true && (42 < -901)) {
            for fi6 := 0; fi6 < 12; fi6++ {
                p0 = (fi6 ^ fi4)
                v7 := []int{fi4, (fi4 | 2147483647)}
                var v8 uint32 = 1
                _ = fi6
                _ = v7
                _ = v8
            }
            var v9 int64 = -703
            _ = v9
        }
        _ = fi4
        _ = v5
    }
    fmt.Println(p0)
    if false {
        v10 := GenS1{F0: (p0 | p0), F1: p0}
        for fi11 := 0; fi11 < 17; fi11++ {
            switch "\n\t" {
            case "\n\t":
                switch fi11 {
                case 0:
                    v12 := GenS2{F0: (false || true), F1: -1, F2: -26}
                    _ = v12
                case 32768:
                    fmt.Println(p0, fi11)
                case 2147483647:
                    fmt.Println(p0, fi11)
                default:
                    v13 := []int{(fi11 / ((fi11) | 1)), 424, p0}
                    _ = v13
                }
            default:
                var v14 uint = 255
                _ = v14
            }
            var v15 uint64 = 0
            v16 := []int{p0, (fi11 / (((fi11 - (p0 % ((p0) | 1)))) | 1))}
            _ = fi11
            _ = v15
            _ = v16
        }
        var v17 uint32 = 4294967295
        _ = v10
        _ = v17
    } else {
        p0 += (p0 + 0)
    }
    p0 -= p0
    _ = p0
    return -1
}
func genFn1(p0 uint64) uint {
    v18 := GenS3{F0: 1, F1: 10, F2: uint8(p0)}
    p0 -= 189
    v19 := []int{659, -1, -1}
    _ = p0
    _ = v18
    _ = v19
    return 286
}

func main() {
    var v1 int16 = 86
    var v2 int32 = 2147483647
    v3 := make(map[uint32]bool)
    var v20 int16 = -32768
    if false {
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 bool = false
            _ = fi21
            _ = v22
        }
        v23 := GenS4{F0: (v1 >= int16(v2))}
        _ = v23
    }
    if (42 >= -9223372036854775808) {
        v2 = v2
    } else {
        for fi24 := 0; fi24 < 11; fi24++ {
            if ((!(!true)) || (true || true)) {
                var v25 uint64 = 0
                v2 = (-(v2))
                v20 = (v1 + 32767)
                _ = v25
            }
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 uint8 = 127
                var v28 uint32 = 42
                v29 := GenS5{F0: ((false && (!(1.5 == 1.5))) || ((true && (true && true)) && (v20 == (v1 & -826)))), F1: (((("\n\t" + ("" + "")) + "a") + "foo bar") + "a"), F2: v27}
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            v1 += (v20 >> ((0) & 63))
            _ = fi24
        }
        v30 := []int{550}
        _ = v30
    }
    var v31 uint64 = 238
    if true {
        var v32 int8 = -66
        v33 := GenS6{F0: (v31 == v31)}
        _ = v32
        _ = v33
    } else {
        var v34 bool = false
        v1 -= 1000
        _ = v34
    }
    var v35 int64 = -9223372036854775808
    v1 -= (-(v1))
    var v36 int8 = 100
    switch 2147483647 {
    case 2147483647:
        var v37 uint16 = 0
        _ = v37
    case -966:
        v38 := GenS2{F0: ((!(v35 != v35)) || (1 > (v35 ^ (v35 % (((v35 >> ((4294967295) & 63))) | 1))))), F1: (v36 & (v36 % (((v36 % ((-128) | 1))) | 1))), F2: v36}
        _ = v38
    case 32767:
        if false {
            delete(v3, 353)
        } else {
            switch -221 {
            case -542:
                fmt.Println(v20, v20)
            }
            var v39 float64 = (-2.25 + 1.5)
            v35 += 9223372036854775807
            _ = v39
        }
    }
    if (false && true) {
        v20 = v20
        if ((v1 - v20) > v1) {
            var v40 uint32 = 65535
            for fi41 := 0; fi41 < 6; fi41++ {
                v1 = (v1 & (v20 / ((42) | 1)))
                _ = fi41
            }
            v42 := []int{int(v2)}
            _ = v40
            _ = v42
        } else {
            var v43 uint32 = 2147483647
            _ = v43
        }
        v31 += v31
    }
    var v44 int8 = (^(v36))
    if ((v20 & (v20 / ((-603) | 1))) > (-(v20))) {
        v45 := GenS7{F0: 32767, F1: ((("\n\t" > "a") || (-2.25 != float32(-2.25))) || (4294967295 <= 255))}
        _ = v45
    } else {
        var v46 uint8 = 128
        _ = v46
    }
    v47 := GenS8{F0: 2, F1: 690}
    v48 := GenS5{F0: false, F1: "foo bar", F2: 255}
    if ((!(2 >= 2)) || (!(true && (2 > 127)))) {
        v49 := []int{-48, 2147483647}
        _ = v49
    } else {
        v31 += 4294967295
        switch 42 {
        case 0:
            for fi50 := 0; fi50 < 8; fi50++ {
                for fi51 := 0; fi51 < 4; fi51++ {
                    v52 := GenS3{F0: fi51, F1: 255, F2: 44}
                    var v53 uint8 = 42
                    _ = fi51
                    _ = v52
                    _ = v53
                }
                v44 = (^(v36))
                _ = fi50
            }
        case 2147483647:
            var v54 uint8 = 128
            _ = v54
        case 881:
            v3[789] = ((!true) || false)
        }
        var v55 uint8 = 91
        _ = v55
    }
    v56 := GenS9{F0: genFn0(127), F1: 32767, F2: -607}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v44
    _ = v47
    _ = v48
    _ = v56
}
