package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint64) uint16 {
    v4 := GenS1{F0: -5}
    var v5 int32 = int32(p0)
    v6 := []int{-256, 128}
    p0 -= (p0 + (^(p0)))
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 int64, p1 int64, p2 string) string {
    v7 := GenS1{F0: 127}
    fmt.Println(p1)
    var v8 float64 = 0.0
    var v9 uint = 18446744073709551615
    var v10 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (("" + "hello") + p2)
}

func main() {
    v1 := []uint32{2147483648, 856, 65535}
    var v2 int16 = 986
    v3 := make(map[int]int)
    if ((false && (2 > 65535)) && (false && false)) {
        v2 += (v2 * (v2 & v2))
    } else {
        var v11 float32 = 0.1
        _ = v11
    }
    if ("hello" > ("foo bar" + "hello")) {
        for fi12 := 0; fi12 < 8; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                v14 := GenS2{F0: "foo bar", F1: 85}
                v15 := GenS3{F0: -921}
                v16 := GenS3{F0: 42}
                v17 := GenS1{F0: 1}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi12
        }
        v18 := GenS3{F0: -2147483648}
        _ = v18
    }
    if (v2 == v2) {
        if ("foo bar" < ("" + ((("\n\t" + ("a" + "a")) + "hello") + "foo bar"))) {
            if (!(!(1 <= -51))) {
                switch ("a" + ("" + "")) {
                case "hello":
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                case "foo bar":
                    v2 -= v2
                }
                for fi19 := 0; fi19 < 5; fi19++ {
                    fmt.Println(fi19, v2)
                    _ = fi19
                }
            }
        }
    }
    v20 := GenS2{F0: "\n\t", F1: 87}
    var v21 int32 = 42
    var v22 int = 1
    _ = len(v1)
    var v23 int = (v22 - 343)
    var v24 int8 = 127
    for fi25 := 0; fi25 < 17; fi25++ {
        if true {
            v22 = (v23 / ((9223372036854775807) | 1))
            if true {
                switch (v22 & fi25) {
                case -9223372036854775808:
                    if (0.0 >= (3.14159 - 1e10)) {
                        if ("foo bar" >= "\n\t") {
                            if ((!(v24 <= v24)) && true) {
                                v2 += -43
                            } else {
                                var v26 uint64 = 42
                                v2 = 42
                                for fi27 := 0; fi27 < 10; fi27++ {
                                    var v28 int64 = int64(v24)
                                    if (-2.25 != 1.5) {
                                        var v29 uint32 = 0
                                        v23 = v23
                                        _ = v29
                                    }
                                    switch 32767 {
                                    case 9223372036854775807:
                                        switch (fi27 & -850) {
                                        case 77:
                                            v2 -= (v2 / ((v2) | 1))
                                        default:
                                            fmt.Println(fi27)
                                        }
                                    default:
                                        fmt.Println(fi25, v28)
                                        var v30 int16 = v2
                                        _ = v30
                                    }
                                    v3[1] = (fi25 % ((fi27) | 1))
                                    _ = fi27
                                    _ = v28
                                }
                                _ = v26
                            }
                            v24 += (v24 ^ (v24 >> ((4294967295) & 63)))
                        }
                    }
                case 1:
                    for fi31 := 0; fi31 < 19; fi31++ {
                        if true {
                            v32 := GenS4{F0: true, F1: 1}
                            var v33 string = ""
                            _ = v32
                            _ = v33
                        }
                        var v34 string = "hello"
                        if ((-(fi31)) <= v23) {
                            if ((fi25 >> ((487) & 63)) != (fi25 ^ fi25)) {
                                switch (^(v22)) {
                                case 42:
                                    var v35 int = (v23 >> ((255) & 63))
                                    _ = v35
                                }
                                switch 1 {
                                case 42:
                                    if (!((-(v24)) < (-(v24)))) {
                                        var v36 int = (v23 >> ((255) & 63))
                                        _ = v36
                                    } else {
                                        v37 := GenS1{F0: (v24 & v24)}
                                        _ = v37
                                    }
                                case -9223372036854775808:
                                    v21 = v21
                                case -361:
                                    switch ((v34 + v34) + v34) {
                                    case "":
                                        var v38 uint = 18446744073709551615
                                        _ = v38
                                    case "a":
                                        fmt.Println(v22)
                                    case "foo bar":
                                        v39 := GenS5{F0: v34, F1: 42}
                                        _ = v39
                                    default:
                                        if false {
                                            var v40 bool = false
                                            _ = v40
                                        }
                                        delete(v3, (fi25 * (fi25 ^ 271)))
                                    }
                                default:
                                    v22 -= 64
                                }
                            }
                        }
                        _ = fi31
                        _ = v34
                    }
                case -451:
                    var v41 float64 = -2.25
                    _ = v41
                default:
                    if (((0.1 - (1e10 + 1.5)) - ((1e10 / (((0.0 / 0.1) / (0.1 + (float32((float64(1e10) - (1e10 - ((float64(((((0.1 + 3.14159) - float32(float64(1.5))) + 0.1) + 0.1)) - (((3.14159 * 3.14159) + (1e10 / (1e10 + 3.14159))) * 3.14159)) - -2.25)))) / -2.25))) + 0.1)) / (0.0 - (0.0 - 3.14159)))) > 0.0) {
                        v42 := v3[v23]
                        _ = v42
                    } else {
                        for fi43 := 0; fi43 < 18; fi43++ {
                            v24 = 42
                            _ = fi43
                        }
                    }
                }
                _ = len(v3)
                switch 398 {
                case 804:
                    switch ((("a" + "a") + "") + "") {
                    case "\n\t":
                        fmt.Println(v24, fi25)
                    default:
                        if true {
                            for fi44 := 0; fi44 < 11; fi44++ {
                                var v45 int32 = v21
                                var v46 bool = (!true)
                                _ = fi44
                                _ = v45
                                _ = v46
                            }
                            for fi47 := 0; fi47 < 8; fi47++ {
                                if true {
                                    v2 += v2
                                } else {
                                    v21 = v21
                                }
                                v23 = -125
                                if true {
                                    v23 = fi25
                                }
                                _ = fi47
                            }
                            for fi48 := 0; fi48 < 9; fi48++ {
                                var v49 float32 = ((0.1 / -2.25) - (-2.25 + (3.14159 / 1.0)))
                                _ = fi48
                                _ = v49
                            }
                        } else {
                            if (0 != 0) {
                                var v50 uint32 = 2
                                _ = v50
                            }
                        }
                    }
                case 2147483647:
                    v51 := GenS6{F0: 127, F1: 65535}
                    _ = v51
                }
            }
        } else {
            v2 -= (-(v2))
            var v52 int16 = v2
            _ = v52
        }
        _ = fi25
    }
    switch "\n\t" {
    case "foo bar":
        for fi53 := 0; fi53 < 3; fi53++ {
            v24 -= 119
            var v54 int16 = v2
            for fi55 := 0; fi55 < 8; fi55++ {
                v1 = append(v1, 2147483647)
                _ = fi55
            }
            _ = fi53
            _ = v54
        }
    }
    if false {
        v56 := GenS3{F0: 751}
        for fi57 := 0; fi57 < 16; fi57++ {
            var v58 uint16 = genFn0(42)
            _ = fi57
            _ = v58
        }
        fmt.Println(v24)
        _ = v56
    } else {
        var v59 int8 = v24
        v59 = (v24 / ((-42) | 1))
        _ = v59
    }
    for fi60 := 0; fi60 < 16; fi60++ {
        v24 = v24
        v24 += (v24 >> ((65535) & 63))
        _ = fi60
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
