package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS8 struct { F0 int16; F1 bool; F2 int8 }
type GenS9 struct { F0 uint; F1 int16; F2 uint32 }
type GenS10 struct { F0 uint16 }
func genFn0(p0 int) int {
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := []int{}
        switch (-(p0)) {
        case -813:
            p0 = (p0 >> ((65535) & 63))
        case -1:
            var v8 uint64 = 18446744073709551615
            _ = v8
        case 820:
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                v9 := GenS1{F0: 4294967295}
                _ = v9
            case "":
                var v10 uint = 992
                _ = v10
            case "a":
                p0 = (p0 << ((2) & 63))
            default:
                var v11 uint64 = 402
                _ = v11
            }
        default:
            var v12 float64 = (1e10 - 3.14159)
            _ = v12
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    fmt.Println(p0)
    v13 := []int{-736}
    _ = p0
    _ = v13
    return (-(p0))
}
func genFn1(p0 int32, p1 int8, p2 int) uint8 {
    if true {
        v14 := GenS1{F0: 42}
        fmt.Println(p2, p0)
        _ = v14
    } else {
        v15 := GenS1{F0: 42}
        v16 := []int{(p2 % ((32768) | 1)), (p2 / ((p2) | 1)), 32767, genFn0(genFn0(p2))}
        _ = v15
        _ = v16
    }
    var v17 uint32 = 42
    switch "" {
    case "foo bar":
        var v18 int8 = 0
        _ = v18
    case "hello":
        v17 = (v17 + v17)
    }
    if false {
        v17 = (v17 % ((v17) | 1))
        p2 -= (^(p2))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 0
}

func main() {
    v1 := []int{0, 128, -1}
    v2 := make(map[uint32]string)
    var v3 uint = 4294967295
    var v4 bool = false
    var v5 int16 = 32767
    switch 4294967295 {
    case 758:
        var v19 float32 = (0.0 * 1.5)
        _ = v19
    case -293:
        var v20 uint8 = uint8(v3)
        _ = v20
    case 1:
        var v21 uint16 = 0
        _ = v21
    default:
        for fi22 := 0; fi22 < 19; fi22++ {
            v4 = (v4 && true)
            _ = fi22
        }
    }
    v23 := GenS2{F0: 0}
    v5 = v5
    if (65535 >= 942) {
        v24 := GenS2{F0: -187}
        for fi25 := 0; fi25 < 19; fi25++ {
            v5 += v5
            _ = fi25
        }
        _ = v24
    }
    var v26 string = ""
    for fi27 := 0; fi27 < 2; fi27++ {
        switch ((v26 + (((v26 + "") + v26) + ("hello" + v26))) + v26) {
        case "a":
            for fi28 := 0; fi28 < 4; fi28++ {
                if ((0.1 / 1e10) <= float64((1e10 * 1.5))) {
                    if v4 {
                        v29 := GenS3{F0: 42}
                        _ = v29
                    }
                } else {
                    v4 = v4
                    if false {
                        if true {
                            v30 := GenS4{F0: (v5 % (((v5 - (v5 | (v5 & -1)))) | 1)), F1: (^(fi28))}
                            _ = v30
                        }
                        var v31 float64 = 0.0
                        _ = v31
                    } else {
                        switch genFn0((fi28 ^ -9223372036854775808)) {
                        case 0:
                            var v32 int32 = 0
                            _ = v32
                        case -222:
                            var v33 string = v26
                            _ = v33
                        }
                    }
                }
                switch v26 {
                case "foo bar":
                    v5 -= (v5 >> ((1) & 63))
                    v26 = ((v26 + "") + (v26 + "\n\t"))
                }
                var v34 bool = true
                v34 = ((1 == 391) || (true || false))
                _ = fi28
                _ = v34
            }
        case "foo bar":
            v35 := GenS5{F0: v5, F1: 467}
            _ = v35
        case "\n\t":
            v36 := v2[2147483648]
            _ = v36
        default:
            for fi37 := 0; fi37 < 5; fi37++ {
                v38 := GenS2{F0: 2147483647}
                _ = fi37
                _ = v38
            }
        }
        if true {
            v39 := GenS6{F0: 1}
            v40 := GenS7{F0: 71, F1: 32767, F2: -2147483648}
            if ((4294967295 != 2) && v4) {
                var v41 int8 = 42
                v5 = 211
                if (v4 || v4) {
                    fmt.Println(v41, v3, v5)
                }
                _ = v41
            } else {
                v4 = v4
                var v42 float32 = 1.5
                v3 += v3
                _ = v42
            }
            _ = v39
            _ = v40
        } else {
            v26 = v26
            switch (v26 + "a") {
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = (fi27 * genFn0(32768))
                }
            }
        }
        _ = fi27
    }
    v5 = v5
    if (603 <= 0) {
        if (!v4) {
            for fi43 := 0; fi43 < 7; fi43++ {
                var v44 int64 = 9223372036854775807
                fmt.Println(fi43)
                _ = fi43
                _ = v44
            }
        }
        for fi45 := 0; fi45 < 13; fi45++ {
            v5 = v5
            _ = fi45
        }
        if (995 >= 65535) {
            if len(v1) > 0 {
                v1[0] = 42
            }
            _ = len(v1)
        }
    }
    var v46 float64 = 1.5
    v47 := GenS7{F0: 668, F1: 755, F2: 992}
    if (255 >= 32767) {
        v48 := v2[2147483648]
        _ = v48
    }
    v49 := GenS8{F0: int16(v3), F1: (134 == -782), F2: 73}
    if len(v1) > 0 {
        v1[0] = int(v5)
    }
    v26 = v26
    v50 := GenS9{F0: (v3 - 255), F1: (^(v5)), F2: 1}
    v51 := GenS10{F0: 2}
    v1 = append(v1, genFn0(-1))
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
