package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint16; F1 bool; F2 int }
type GenS7 struct { F0 uint32; F1 uint64; F2 int }
func genFn0(p0 uint32, p1 uint8) int32 {
    var v4 float32 = float32(0.1)
    v5 := []int{}
    if (false || true) {
        var v6 bool = (0 != 0)
        switch 42 {
        case -1:
            v7 := GenS1{F0: -526, F1: 372}
            _ = v7
        case 889:
            var v8 int64 = -1
            _ = v8
        case -860:
            var v9 float64 = (1.5 - ((-2.25 + 3.14159) - 1.5))
            _ = v9
        default:
            v10 := GenS2{F0: 9223372036854775807, F1: 18446744073709551615, F2: (^(p1))}
            _ = v10
        }
        _ = v6
    }
    if (32767 < 0) {
        p1 -= 2
        v11 := GenS2{F0: -1, F1: 107, F2: p1}
        if false {
            p1 = (p1 / ((2) | 1))
            p1 = (p1 & p1)
        }
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 258
}
func genFn1(p0 uint, p1 int64, p2 int64) uint8 {
    var v12 uint32 = 15
    switch (("" + "a") + "hello") {
    case "\n\t":
        v13 := []int{2147483647, 32767}
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := GenS3{F0: (fi14 + 101), F1: 22}
        p2 += -17
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 uint = p0
        switch (("hello" + ((("" + (("hello" + ("hello" + "hello")) + ("foo bar" + ""))) + "") + "hello")) + "hello") {
        case "hello":
            var v18 bool = true
            _ = v18
        case "foo bar":
            if true {
                switch int(p1) {
                case -236:
                    v19 := GenS3{F0: (fi16 + -2147483648), F1: -74}
                    _ = v19
                default:
                    v20 := []int{(fi16 & (fi16 % (((^(fi16))) | 1))), fi16, 9223372036854775807}
                    _ = v20
                }
                v21 := GenS1{F0: 1, F1: 616}
                for fi22 := 0; fi22 < 6; fi22++ {
                    var v23 uint = p0
                    switch (fi16 * (fi22 % ((fi16) | 1))) {
                    case -1:
                        if (p0 <= p0) {
                            var v24 uint16 = 42
                            var v25 uint8 = 42
                            var v26 int8 = 1
                            v27 := GenS4{F0: 255}
                            _ = v24
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    }
                    v17 += (v23 * p0)
                    _ = fi22
                    _ = v23
                }
                _ = v21
            }
        case "":
            switch (fi16 % ((22) | 1)) {
            case 127:
                for fi28 := 0; fi28 < 2; fi28++ {
                    switch fi28 {
                    case 0:
                        v29 := GenS2{F0: (p1 / ((1) | 1)), F1: 177, F2: 193}
                        _ = v29
                    case 2147483647:
                        v12 -= (v12 + (v12 % (((v12 - v12)) | 1)))
                    default:
                        v17 = (v17 % ((p0) | 1))
                    }
                    _ = fi28
                }
            case 42:
                v17 = (v17 - 2)
            default:
                p0 += 2
            }
        default:
            var v30 bool = (true || ((0 < 65535) && true))
            _ = v30
        }
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 325
    var v3 float32 = ((0.1 * 3.14159) / 1.0)
    v2 += (v2 >> ((607) & 63))
    if false {
        switch "hello" {
        case "hello":
            v2 = (v2 + (v2 | v2))
        case "foo bar":
            for fi31 := 0; fi31 < 4; fi31++ {
                v32 := GenS5{F0: uint32(v2), F1: (v2 + (^(v2))), F2: -533}
                var v33 bool = (-1 <= -32768)
                _ = fi31
                _ = v32
                _ = v33
            }
        case "":
            var v34 uint32 = 2147483647
            _ = v34
        }
        v35 := []int{127, 42, 442}
        _ = v35
    } else {
        var v36 uint32 = 645
        _ = v36
    }
    var v37 int16 = -1
    if false {
        v38 := GenS1{F0: 1, F1: v2}
        var v39 uint16 = 2
        _ = v38
        _ = v39
    }
    var v40 uint64 = v2
    var v41 uint64 = v2
    var v42 int32 = genFn0(65535, 0)
    fmt.Println(v37)
    for fi43 := 0; fi43 < 2; fi43++ {
        v44 := GenS1{F0: (-(v42)), F1: v41}
        _ = fi43
        _ = v44
    }
    v37 = v37
    var v45 int64 = 629
    for fi46 := 0; fi46 < 6; fi46++ {
        for fi47 := 0; fi47 < 10; fi47++ {
            v48 := GenS1{F0: (v42 / ((v42) | 1)), F1: v2}
            _ = fi47
            _ = v48
        }
        v49 := []int{(fi46 >> ((4294967295) & 63)), fi46, 127, fi46}
        v40 -= v2
        _ = fi46
        _ = v49
    }
    for fi50 := 0; fi50 < 10; fi50++ {
        if (false || (!(true || ((true || false) && true)))) {
            var v51 uint64 = 0
            var v52 float32 = v3
            _ = v51
            _ = v52
        } else {
            v53 := GenS3{F0: fi50, F1: 1}
            fmt.Println(v41)
            _ = v53
        }
        delete(v1, 792)
        var v54 int64 = 0
        _ = fi50
        _ = v54
    }
    v55 := GenS6{F0: 32767, F1: (((v45 & v45) != v45) || (true && ((255 != 255) || ((32767 == v37) && true)))), F2: 625}
    v37 -= v37
    for fi56 := 0; fi56 < 20; fi56++ {
        switch fi56 {
        case 2147483647:
            v57 := GenS7{F0: 480, F1: v2, F2: (fi56 ^ 42)}
            _ = v57
        case 128:
            v40 += v40
        default:
            var v58 float32 = v3
            _ = v58
        }
        _ = fi56
    }
    for fi59 := 0; fi59 < 18; fi59++ {
        var v60 int16 = v37
        if true {
            v61 := []int{-1}
            switch -9223372036854775808 {
            case 9223372036854775807:
                v62 := v1[1]
                _ = v62
            case 128:
                v40 += (v41 ^ v40)
            }
            _ = v61
        }
        _ = fi59
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v40
    _ = v41
    _ = v42
    _ = v45
    _ = v55
}
