package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 uint8; F1 uint16; F2 bool }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int8 }
func genFn0(p0 uint8, p1 int, p2 uint8) int16 {
    var v5 uint8 = 127
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := []int{fi6, (p1 / (((fi6 | p1)) | 1)), (^(fi6)), fi6}
        var v8 int32 = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    if (310 == -9223372036854775808) {
        var v9 float32 = 1.5
        _ = v9
    } else {
        if true {
            for fi10 := 0; fi10 < 10; fi10++ {
                for fi11 := 0; fi11 < 17; fi11++ {
                    fmt.Println(fi10)
                    _ = fi11
                }
                for fi12 := 0; fi12 < 7; fi12++ {
                    var v13 uint8 = 210
                    var v14 uint64 = 799
                    var v15 uint16 = 0
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                _ = fi10
            }
            v16 := []int{128, 42, -9223372036854775808}
            for fi17 := 0; fi17 < 18; fi17++ {
                if (((0 > 569) || (90 != 42)) && (true && false)) {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        v5 = v5
                        v5 = (p0 ^ p0)
                        _ = fi18
                    }
                }
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 int16 = 1000
                    _ = fi19
                    _ = v20
                }
                _ = len(v16)
                _ = fi17
            }
            _ = v16
        }
        p1 = (p1 - p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    v4 := []uint32{347, 0, 65535}
    switch "hello" {
    case "":
        fmt.Println(v1)
    case "hello":
        for fi21 := 0; fi21 < 6; fi21++ {
            if ((!false) && false) {
                v22 := GenS1{F0: 2}
                v23 := GenS1{F0: 2}
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 11; fi24++ {
                fmt.Println(fi21)
                v1 += (^(v1))
                _ = len(v4)
                _ = fi24
            }
            _ = fi21
        }
    case "a":
        if true {
            fmt.Println(v1)
        } else {
            v1 = (v1 & -128)
            var v25 int8 = v1
            var v26 string = (("\n\t" + ("a" + "")) + "\n\t")
            _ = v25
            _ = v26
        }
    }
    v27 := GenS2{F0: 42, F1: 42}
    v1 -= v1
    v1 = (v1 % ((v1) | 1))
    switch 32767 {
    case -2147483648:
        for fi28 := 0; fi28 < 15; fi28++ {
            fmt.Println(v1)
            _ = fi28
        }
    case 4294967295:
        var v29 float64 = 0.0
        _ = v29
    default:
        var v30 uint = 18446744073709551615
        _ = v30
    }
    _ = len(v4)
    if (("\n\t" == "") || (42 <= genFn0(177, 1, 1))) {
        var v31 uint64 = 487
        if (0 >= 42) {
            switch 0 {
            case -2147483648:
                v1 += 0
            default:
                var v32 uint8 = 255
                _ = v32
            }
            switch "a" {
            case "a":
                v33 := GenS1{F0: 6}
                _ = v33
            case "foo bar":
                v31 = (v31 % (((v31 ^ (v31 % ((1) | 1)))) | 1))
            default:
                var v34 uint32 = 65535
                _ = v34
            }
        } else {
            v35 := GenS3{F0: 1, F1: 78}
            v36 := GenS4{F0: "\n\t"}
            if false {
                var v37 uint = 121
                v38 := GenS2{F0: 890, F1: -9223372036854775808}
                v31 -= v31
                _ = v37
                _ = v38
            }
            _ = v35
            _ = v36
        }
        v1 -= v1
        _ = v31
    }
    if len(v4) > 0 {
        v4[0] = 2
    }
    v39 := GenS5{F0: -1, F1: 140}
    for fi40 := 0; fi40 < 1; fi40++ {
        for fi41 := 0; fi41 < 18; fi41++ {
            if ((fi40 * -9223372036854775808) == 0) {
                switch "hello" {
                case "foo bar":
                    v1 -= v1
                case "":
                    fmt.Println(v1, fi41)
                case "hello":
                    v42 := GenS2{F0: 0, F1: 802}
                    _ = v42
                }
            } else {
                fmt.Println(fi41, fi40)
                fmt.Println(v1, fi41)
                var v43 uint8 = 152
                _ = v43
            }
            if true {
                fmt.Println(v1, fi41)
                if false {
                    v44 := GenS6{F0: (fi41 % ((128) | 1)), F1: 4294967295}
                    var v45 uint = 65535
                    if false {
                        fmt.Println(v45, v1, fi40)
                    }
                    _ = v44
                    _ = v45
                }
                var v46 uint8 = 1
                _ = v46
            }
            if false {
                fmt.Println(fi40, v1, fi40)
                fmt.Println(v1, v1, fi40)
                switch "foo bar" {
                case "hello":
                    var v47 float32 = 3.14159
                    _ = v47
                case "\n\t":
                    if len(v4) > 0 {
                        v4[0] = 821
                    }
                case "foo bar":
                    var v48 uint64 = 399
                    _ = v48
                }
                if (-194 < 911) {
                    v49 := GenS7{F0: 109, F1: 32767, F2: false}
                    if (255 < 255) {
                        v50 := GenS7{F0: 215, F1: 88, F2: false}
                        _ = v50
                    }
                    _ = v49
                } else {
                    if (!true) {
                        v1 = (v1 % ((v1) | 1))
                    }
                    v1 += v1
                }
                v51 := GenS1{F0: 791}
                _ = v51
            } else {
                v52 := v2[(("a" + "foo bar") + "foo bar")]
                _ = v52
            }
            _ = fi41
        }
        v53 := GenS8{F0: v1}
        var v54 float32 = 0.1
        _ = fi40
        _ = v53
        _ = v54
    }
    v1 -= v1
    v1 += (v1 % ((127) | 1))
    var v55 float64 = -2.25
    for fi56 := 0; fi56 < 13; fi56++ {
        v57 := GenS3{F0: 512, F1: -641}
        v55 = ((v55 / (v55 - ((0.1 * ((0.1 / (v55 + -2.25)) + v55)) * v55))) - v55)
        switch (fi56 / ((fi56) | 1)) {
        case 32768:
            v58 := GenS9{F0: 0}
            _ = v58
        case -2147483648:
            var v59 int16 = -768
            _ = v59
        }
        _ = fi56
        _ = v57
    }
    delete(v3, 65535)
    fmt.Println(v1, v1)
    v1 += -128
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v55
}
