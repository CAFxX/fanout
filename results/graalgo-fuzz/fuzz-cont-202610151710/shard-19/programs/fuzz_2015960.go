package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint; F1 int16; F2 int64 }
type GenS7 struct { F0 int16; F1 uint64 }
func genFn0() int8 {
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int16 = 0
        var v7 bool = false
        _ = fi5
        _ = v6
        _ = v7
    }
    if (!true) {
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: 128, F1: "hello", F2: -1}
            _ = fi8
            _ = v9
        }
        var v10 uint16 = 255
        v10 += (^(v10))
        _ = v10
    } else {
        for fi11 := 0; fi11 < 12; fi11++ {
            switch fi11 {
            case 9223372036854775807:
                v12 := []int{(-(fi11)), (^(fi11))}
                _ = v12
            }
            var v13 string = "a"
            _ = fi11
            _ = v13
        }
        var v14 int8 = -128
        _ = v14
    }
    v15 := GenS1{F0: 0, F1: "foo bar", F2: 1}
    v16 := []int{578, -9223372036854775808}
    _ = v15
    _ = v16
    return 84
}
func genFn1(p0 int8, p1 uint32, p2 bool) int8 {
    p1 = (p1 + 1)
    v17 := GenS2{F0: 853}
    if (32767 != 355) {
        if false {
            switch -785 {
            case 32767:
                v18 := []int{}
                _ = v18
            case 868:
                var v19 int32 = 126
                _ = v19
            case 0:
                v20 := GenS1{F0: 48, F1: ((("\n\t" + "a") + "") + "foo bar"), F2: -56}
                _ = v20
            }
        } else {
            var v21 int32 = -1
            var v22 float64 = 0.0
            _ = v21
            _ = v22
        }
        var v23 uint64 = 0
        if (p1 > p1) {
            v23 = (v23 >> ((4294967295) & 63))
        }
        _ = v23
    } else {
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 uint16 = 255
            var v26 int16 = 0
            p0 += (p0 & (p0 & (p0 ^ (p0 ^ (p0 | 42)))))
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    p2 = (genFn0() > 100)
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return (p0 / ((genFn0()) | 1))
}

func main() {
    v1 := []string{}
    v2 := make(map[string]int)
    var v3 int64 = -844
    v4 := make(map[uint32]int)
    if (!(!((!true) || true))) {
        v27 := GenS3{F0: v3}
        _ = v27
    } else {
        v4[2] = 42
    }
    v28 := GenS4{F0: 16, F1: 2}
    for fi29 := 0; fi29 < 5; fi29++ {
        fmt.Println(fi29)
        v30 := GenS5{F0: 58, F1: (^(v3))}
        _ = fi29
        _ = v30
    }
    switch "foo bar" {
    case "a":
        for fi31 := 0; fi31 < 2; fi31++ {
            for fi32 := 0; fi32 < 5; fi32++ {
                var v33 int16 = 32767
                var v34 float64 = -2.25
                v33 -= 32767
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi31
        }
    case "hello":
        if true {
            v3 -= (-(v3))
            if len(v1) > 0 {
                v1[0] = ""
            }
        }
    default:
        v3 += v3
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        var v36 int64 = v3
        switch "\n\t" {
        case "hello":
            var v37 float64 = 1.5
            _ = v37
        case "a":
            var v38 uint16 = 1
            _ = v38
        default:
            var v39 int64 = v36
            _ = v39
        }
        _ = fi35
        _ = v36
    }
    if (((42 >= 91) && (((2 >= 477) || true) || (4294967295 < 0))) && false) {
        var v40 int16 = 1000
        for fi41 := 0; fi41 < 18; fi41++ {
            v42 := GenS6{F0: 0, F1: v40, F2: 1}
            v43 := GenS6{F0: 449, F1: (v40 | v40), F2: (v3 % ((v3) | 1))}
            var v44 uint32 = 42
            _ = fi41
            _ = v42
            _ = v43
            _ = v44
        }
        var v45 int32 = 2147483647
        _ = v40
        _ = v45
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v46 bool = (true || false)
        _ = v46
    }
    if false {
        v47 := GenS2{F0: 100}
        fmt.Println(v3)
        _ = v47
    }
    var v48 float64 = -2.25
    var v49 uint64 = 309
    fmt.Println(v3, v3)
    var v50 int32 = 235
    for fi51 := 0; fi51 < 10; fi51++ {
        for fi52 := 0; fi52 < 19; fi52++ {
            var v53 float64 = (0.0 / v48)
            var v54 float32 = float32(1e10)
            v55 := GenS5{F0: 42, F1: -1}
            v56 := GenS4{F0: -1, F1: 283}
            v49 += v49
            v3 -= (v3 | v3)
            v48 -= float64(v54)
            fmt.Println(v3)
            v50 = v50
            v57 := GenS6{F0: 65535, F1: 103, F2: -508}
            _ = fi52
            _ = v53
            _ = v54
            _ = v55
            _ = v56
            _ = v57
        }
        if false {
            delete(v4, 2147483648)
            v3 = (v3 - (v3 >> ((4294967295) & 63)))
        } else {
            var v58 int16 = -1
            _ = v58
        }
        _ = fi51
    }
    var v59 float32 = 3.14159
    for fi60 := 0; fi60 < 16; fi60++ {
        _ = len(v4)
        if ((false || (true || false)) && true) {
            v61 := GenS7{F0: 329, F1: 4294967295}
            _ = v61
        }
        _ = fi60
    }
    fmt.Println(v50, v48, v50)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v50
    _ = v59
}
