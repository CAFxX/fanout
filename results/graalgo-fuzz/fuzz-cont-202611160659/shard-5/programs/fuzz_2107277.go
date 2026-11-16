package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 uint64; F2 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint8) uint64 {
    if false {
        var v6 int64 = 0
        fmt.Println(v6, p0, v6)
        _ = v6
    } else {
        p1 += (p1 / ((p1) | 1))
    }
    v7 := GenS1{F0: 42}
    if false {
        if false {
            var v8 int16 = 1
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 int = fi9
                _ = fi9
                _ = v10
            }
            _ = v8
        } else {
            v11 := []int{42, 0, -9223372036854775808}
            if true {
                fmt.Println(p0, p1)
                if (!true) {
                    fmt.Println(p0)
                    _ = len(v11)
                    for fi12 := 0; fi12 < 10; fi12++ {
                        p0 -= (-(p0))
                        _ = fi12
                    }
                }
            } else {
                var v13 float32 = 0.0
                _ = v13
            }
            var v14 uint64 = 518
            _ = v11
            _ = v14
        }
        var v15 int16 = 777
        fmt.Println(p1, p1, v15)
        _ = v15
    } else {
        switch 42 {
        case 0:
            v16 := []int{42, 744}
            _ = v16
        case 684:
            var v17 string = ("foo bar" + "\n\t")
            _ = v17
        default:
            v18 := GenS1{F0: 1}
            _ = v18
        }
        switch -308 {
        case 9223372036854775807:
            fmt.Println(p1, p1)
        case -1:
            p1 -= 118
        case -9223372036854775808:
            p1 -= 0
        }
        var v19 uint32 = 816
        _ = v19
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1() uint8 {
    v20 := GenS1{F0: 42}
    var v21 uint8 = 42
    if false {
        v21 -= (v21 | (v21 >> ((656) & 63)))
        var v22 int64 = 42
        var v23 bool = ((true && true) && (!(true || true)))
        _ = v22
        _ = v23
    } else {
        fmt.Println(v21)
        var v24 int16 = 1
        _ = v24
    }
    var v25 bool = ("foo bar" != "\n\t")
    _ = v20
    _ = v21
    _ = v25
    return v21
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -9223372036854775808
    v3 := make(map[string]string)
    var v4 int8 = -128
    var v5 uint16 = 0
    var v26 uint = 4294967295
    for fi27 := 0; fi27 < 14; fi27++ {
        fmt.Println(fi27, v5, v5)
        if false {
            _ = len(v1)
            v2 = fi27
        } else {
            switch "\n\t" {
            case "foo bar":
                var v28 int64 = 1
                _ = v28
            case "\n\t":
                v5 = v5
            }
            v29 := GenS2{F0: false}
            _ = v29
        }
        _ = fi27
    }
    v30 := GenS3{F0: 0, F1: 0}
    delete(v3, "")
    v31 := GenS4{F0: 54}
    switch "foo bar" {
    case "":
        if false {
            if (!(955 < 1)) {
                var v32 float64 = 3.14159
                var v33 int = 487
                _ = v32
                _ = v33
            } else {
                if (uint8(v5) == 2) {
                    var v34 uint16 = v5
                    if false {
                        var v35 bool = false
                        _ = v35
                    }
                    v2 = (v2 / ((v2) | 1))
                    _ = v34
                } else {
                    var v36 float64 = -2.25
                    _ = len(v1)
                    switch "a" {
                    case "a":
                        var v37 int64 = 9223372036854775807
                        _ = v37
                    case "\n\t":
                        var v38 bool = false
                        _ = v38
                    }
                    _ = v36
                }
            }
        }
    case "foo bar":
        v39 := GenS3{F0: 1, F1: -996}
        _ = v39
    default:
        var v40 uint16 = v5
        _ = v40
    }
    var v41 int16 = 1
    switch "foo bar" {
    case "foo bar":
        for fi42 := 0; fi42 < 13; fi42++ {
            v26 = 42
            v4 += (v4 * 127)
            _ = len(v1)
            _ = fi42
        }
    case "hello":
        if false {
            if false {
                switch (v2 * -950) {
                case 128:
                    v2 += v2
                case 4294967295:
                    v43 := GenS2{F0: true}
                    _ = v43
                case 2147483647:
                    var v44 uint8 = 42
                    _ = v44
                default:
                    v45 := GenS5{F0: (v26 * (v26 & (v26 >> ((v26) & 63)))), F1: 42, F2: ("" + "a")}
                    _ = v45
                }
            } else {
                v46 := GenS6{F0: (^(v26))}
                _ = v46
            }
            for fi47 := 0; fi47 < 5; fi47++ {
                v48 := []int{-9223372036854775808, fi47, v2, v2}
                for fi49 := 0; fi49 < 15; fi49++ {
                    v2 = fi49
                    if len(v48) > 0 {
                        v48[0] = (fi49 & fi49)
                    }
                    var v50 uint = (v26 & v26)
                    if len(v48) >= 2 {
                        v48 = v48[1:len(v48)]
                    }
                    var v51 float64 = ((((1.5 / (((1e10 - float64(3.14159)) + ((0.0 / (3.14159 * -2.25)) - float64(0.1))) + -2.25)) - -2.25) / 1e10) / 3.14159)
                    _ = fi49
                    _ = v50
                    _ = v51
                }
                v52 := GenS3{F0: 42, F1: v41}
                _ = fi47
                _ = v48
                _ = v52
            }
        } else {
            v2 = 9223372036854775807
        }
    default:
        var v53 int32 = 0
        _ = v53
    }
    for fi54 := 0; fi54 < 13; fi54++ {
        v5 = (v5 ^ (^(v5)))
        fmt.Println(v5, fi54)
        _ = fi54
    }
    if (966 <= 1) {
        var v55 float32 = (1.5 * -2.25)
        _ = v55
    }
    var v56 uint8 = genFn1()
    var v57 bool = (!(42 <= 803))
    _ = len(v3)
    v58 := GenS7{F0: 1}
    v2 += v2
    var v59 uint64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v41
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
