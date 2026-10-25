package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int }
func genFn0(p0 uint, p1 uint16, p2 int32) uint8 {
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(p0, p1)
        p0 -= (p0 ^ 1)
        _ = fi5
    }
    switch ("foo bar" + (("foo bar" + "\n\t") + ("foo bar" + ("hello" + "")))) {
    case "a":
        var v6 int8 = 127
        _ = v6
    case "foo bar":
        if false {
            var v7 int64 = -1
            _ = v7
        }
    default:
        v8 := []int{813}
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint64 = 1
        for fi11 := 0; fi11 < 10; fi11++ {
            switch (fi11 | fi11) {
            case 1:
                fmt.Println(p2, fi9, p2)
            case -566:
                switch (fi11 & fi9) {
                case -2147483648:
                    var v12 int = fi11
                    switch "" {
                    case "hello":
                        fmt.Println(v12)
                    case "foo bar":
                        v13 := []int{0, (fi11 & (fi9 & (^(fi11)))), (^(fi11))}
                        _ = v13
                    default:
                        p0 = (^(p0))
                    }
                    _ = v12
                case 1:
                    switch 790 {
                    case -1:
                        switch "\n\t" {
                        case "hello":
                            switch 127 {
                            case -2147483648:
                                var v14 float64 = 0.0
                                _ = v14
                            case 563:
                                if true {
                                    v15 := []int{-1}
                                    var v16 int32 = (p2 << ((42) & 63))
                                    if false {
                                        var v17 int = 4294967295
                                        v18 := GenS1{F0: 115}
                                        v19 := GenS2{F0: (p1 - 0)}
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                    } else {
                                        v20 := GenS1{F0: 100}
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                        v10 += (v10 ^ (^(v10)))
                                        v21 := GenS3{F0: 42}
                                        _ = v20
                                        _ = v21
                                    }
                                    _ = v15
                                    _ = v16
                                } else {
                                    v22 := GenS4{F0: (fi11 & fi9), F1: ""}
                                    fmt.Println(p1)
                                    _ = v22
                                }
                            }
                        case "\n\t":
                            var v23 int64 = 42
                            _ = v23
                        default:
                            v24 := GenS5{F0: p2, F1: 4294967295, F2: 128}
                            _ = v24
                        }
                    case 152:
                        fmt.Println(p2)
                    case 475:
                        v25 := []int{fi11, fi11, fi9}
                        _ = v25
                    }
                case 932:
                    v26 := []int{fi9, fi11, fi9}
                    _ = v26
                }
            case 127:
                if false {
                    var v27 int = -2147483648
                    fmt.Println(fi9, p2, fi9)
                    _ = v27
                } else {
                    var v28 uint8 = 128
                    v29 := GenS3{F0: 0}
                    var v30 uint64 = 175
                    _ = v28
                    _ = v29
                    _ = v30
                }
            default:
                fmt.Println(p2, fi9, p2)
            }
            _ = fi11
        }
        p2 = 631
        _ = fi9
        _ = v10
    }
    if false {
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 int16 = 42
            var v33 int32 = 0
            v34 := GenS6{F0: -9223372036854775808, F1: v32}
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        if (true || false) {
            p1 = (p1 | (^(p1)))
        } else {
            switch (("a" + "hello") + "") {
            case "hello":
                p1 -= p1
            default:
                v35 := GenS7{F0: 128}
                _ = v35
            }
            var v36 int32 = p2
            if true {
                p1 += 1
                var v37 bool = (true && (1 != 9223372036854775807))
                var v38 uint8 = 184
                _ = v37
                _ = v38
            }
            _ = v36
        }
        switch "\n\t" {
        case "hello":
            var v39 uint16 = p1
            _ = v39
        }
    } else {
        var v40 uint8 = 187
        for fi41 := 0; fi41 < 13; fi41++ {
            var v42 int64 = -24
            fmt.Println(p0)
            fmt.Println(p0, v40, v42)
            _ = fi41
            _ = v42
        }
        _ = v40
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 int32, p1 uint16, p2 uint16) int {
    p2 += 255
    var v43 int32 = (p0 + (-(p0)))
    _ = p0
    _ = p1
    _ = p2
    _ = v43
    return 132
}

func main() {
    v1 := []string{"hello", "a", "foo bar", ""}
    v2 := []int{127}
    var v3 int32 = 1
    var v4 int64 = -150
    for fi44 := 0; fi44 < 12; fi44++ {
        var v45 float32 = float32(1.5)
        var v46 int8 = 1
        var v47 bool = (v3 < (v3 | (-(v3))))
        _ = fi44
        _ = v45
        _ = v46
        _ = v47
    }
    var v48 int64 = v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 = 9223372036854775807
    for fi49 := 0; fi49 < 16; fi49++ {
        if ((!(!(2 < 42))) && (!true)) {
            var v50 float64 = 0.0
            _ = v50
        }
        var v51 float32 = ((0.1 - float32((((0.0 - 0.0) - float64(float32(((0.0 + float64(0.0)) / 0.1)))) + (0.1 / (1e10 - 3.14159))))) + 1.5)
        _ = fi49
        _ = v51
    }
    v48 -= v4
    if (false || (0.0 >= (3.14159 + float32(0.1)))) {
        v48 -= v48
        v52 := GenS5{F0: -1, F1: 65535, F2: 255}
        _ = v52
    }
    switch 4294967295 {
    case 1:
        var v53 float64 = (((3.14159 - 0.0) / 1e10) * 0.0)
        _ = v53
    case 42:
        if (v3 == v3) {
            var v54 float32 = 3.14159
            var v55 uint16 = 573
            switch ("a" + "a") {
            case "foo bar":
                fmt.Println(v4, v54, v54)
            case "\n\t":
                switch genFn1(v3, v55, (v55 + 32767)) {
                case 77:
                    v48 = v4
                case 42:
                    for fi56 := 0; fi56 < 15; fi56++ {
                        fmt.Println(v3, v54, v48)
                        for fi57 := 0; fi57 < 5; fi57++ {
                            v54 += (float32(1.5) / 1.0)
                            v48 = v4
                            var v58 int16 = 32767
                            if len(v1) > 0 {
                                v1[0] = ("hello" + "\n\t")
                            }
                            _ = fi57
                            _ = v58
                        }
                        _ = fi56
                    }
                }
            case "hello":
                v1 = append(v1, (("" + ("hello" + (("\n\t" + "hello") + "hello"))) + "\n\t"))
            }
            _ = v54
            _ = v55
        } else {
            var v59 float32 = 0.0
            _ = v59
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
}
