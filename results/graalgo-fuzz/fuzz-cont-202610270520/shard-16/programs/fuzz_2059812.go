package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 string; F1 string; F2 int64 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 uint }
func genFn0(p0 int32) uint8 {
    p0 -= 655
    switch ("a" + "") {
    case "":
        for fi4 := 0; fi4 < 13; fi4++ {
            v5 := GenS1{F0: -1}
            _ = fi4
            _ = v5
        }
    }
    v6 := GenS2{F0: 100, F1: "hello"}
    var v7 uint64 = 1
    _ = p0
    _ = v6
    _ = v7
    return 179
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 127
    var v3 uint = 65535
    switch "\n\t" {
    case "\n\t":
        delete(v1, "hello")
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS3{F0: "hello", F1: "a", F2: 9223372036854775807}
        v10 := GenS3{F0: ("a" + "hello"), F1: "", F2: 1}
        switch (("hello" + "hello") + "hello") {
        case "a":
            _ = len(v1)
        case "\n\t":
            _ = len(v1)
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    switch -2147483648 {
    case -1:
        var v11 uint8 = uint8(v3)
        _ = v11
    default:
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := []int{fi12}
            v3 = v3
            _ = fi12
            _ = v13
        }
    }
    if false {
        v14 := GenS1{F0: 1}
        for fi15 := 0; fi15 < 14; fi15++ {
            for fi16 := 0; fi16 < 1; fi16++ {
                fmt.Println(fi16, v3)
                v2 = (v2 | (^(v2)))
                _ = fi16
            }
            _ = fi15
        }
        _ = v14
    }
    delete(v1, "hello")
    var v17 string = "foo bar"
    switch (v17 + v17) {
    case "a":
        v18 := v1[v17]
        _ = v18
    case "hello":
        switch "" {
        case "foo bar":
            fmt.Println(v2)
        case "\n\t":
            var v19 string = ((v17 + v17) + v17)
            _ = v19
        }
    case "\n\t":
        switch 392 {
        case 32767:
            for fi20 := 0; fi20 < 11; fi20++ {
                delete(v1, v17)
                var v21 int = fi20
                _ = fi20
                _ = v21
            }
        case 9223372036854775807:
            v22 := v1[(v17 + v17)]
            _ = v22
        default:
            v17 = (((((v17 + (v17 + "\n\t")) + v17) + "") + "") + "hello")
        }
    }
    v2 += genFn0(42)
    if ((18446744073709551615 != 0) && (((1.5 / 0.1) / 1e10) <= (0.0 + float32((1e10 + 3.14159))))) {
        switch v17 {
        case "hello":
            if false {
                if true {
                    fmt.Println(v17, v2, v17)
                    v17 = v17
                }
                v2 += (v2 >> ((v3) & 63))
            }
        }
        v3 = (^(v3))
        if false {
            if (!true) {
                v23 := GenS4{F0: (18446744073709551615 == 42), F1: 94}
                _ = v23
            } else {
                switch 42 {
                case 0:
                    var v24 uint32 = 2147483647
                    _ = v24
                case 32768:
                    v2 -= (v2 >> ((v3) & 63))
                case -9223372036854775808:
                    var v25 float64 = (0.1 - float64((1.5 + 3.14159)))
                    _ = v25
                }
                if true {
                    v2 += genFn0(-2147483648)
                } else {
                    for fi26 := 0; fi26 < 17; fi26++ {
                        switch "foo bar" {
                        case "a":
                            switch v17 {
                            case "foo bar":
                                v2 = 101
                            }
                        case "":
                            if (1 <= -2147483648) {
                                v27 := []int{(fi26 << (((^(v3))) & 63)), 9223372036854775807, (fi26 | (fi26 >> ((255) & 63)))}
                                if ((uint16(v2) <= 0) && false) {
                                    if true {
                                        _ = len(v1)
                                        _ = len(v1)
                                        v2 = v2
                                    } else {
                                        var v28 uint16 = 42
                                        var v29 uint = (v3 / ((65535) | 1))
                                        _ = v28
                                        _ = v29
                                    }
                                } else {
                                    if (false || false) {
                                        var v30 string = v17
                                        _ = v30
                                    }
                                }
                                _ = v27
                            }
                        case "\n\t":
                            if (-1 >= -729) {
                                v31 := GenS3{F0: v17, F1: v17, F2: -1}
                                switch v17 {
                                case "hello":
                                    var v32 int32 = -566
                                    _ = v32
                                case "a":
                                    for fi33 := 0; fi33 < 9; fi33++ {
                                        v34 := GenS5{F0: 1, F1: 42, F2: -480}
                                        v35 := GenS6{F0: false, F1: 42}
                                        if true {
                                            v3 -= 18446744073709551615
                                        } else {
                                            v1[v17] = 42
                                            v36 := GenS3{F0: v17, F1: (v17 + v17), F2: 9223372036854775807}
                                            _ = v36
                                        }
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                    }
                                default:
                                    for fi37 := 0; fi37 < 12; fi37++ {
                                        v38 := GenS7{F0: (^(v3))}
                                        var v39 uint16 = 481
                                        fmt.Println(fi26, v39)
                                        _ = fi37
                                        _ = v38
                                        _ = v39
                                    }
                                }
                                v17 = v17
                                _ = v31
                            } else {
                                fmt.Println(v17, v2)
                            }
                        }
                        v2 += 102
                        _ = fi26
                    }
                }
            }
            v40 := GenS5{F0: 0, F1: 1, F2: -2147483648}
            v41 := []int{}
            _ = v40
            _ = v41
        }
    }
    if true {
        v42 := GenS3{F0: "", F1: ("\n\t" + v17), F2: 42}
        for fi43 := 0; fi43 < 19; fi43++ {
            for fi44 := 0; fi44 < 2; fi44++ {
                var v45 bool = true
                var v46 float64 = 1.5
                if v45 {
                    var v47 float64 = float64(1.5)
                    v48 := []int{(fi44 | fi43), 9223372036854775807, 32768}
                    _ = len(v1)
                    _ = v47
                    _ = v48
                }
                _ = fi44
                _ = v45
                _ = v46
            }
            if (fi43 != 128) {
                fmt.Println(fi43, fi43, v3)
                switch ((v17 + ("a" + "\n\t")) + v17) {
                case "a":
                    v3 -= 42
                case "hello":
                    var v49 int16 = -1
                    _ = v49
                default:
                    for fi50 := 0; fi50 < 3; fi50++ {
                        var v51 uint64 = 42
                        if (true || false) {
                            var v52 uint64 = (^(v51))
                            v53 := []int{(fi43 * fi43), int(v3)}
                            switch v17 {
                            case "\n\t":
                                var v54 uint32 = 255
                                _ = v54
                            default:
                                v52 -= 215
                            }
                            _ = v52
                            _ = v53
                        }
                        v55 := GenS2{F0: -92, F1: ""}
                        _ = fi50
                        _ = v51
                        _ = v55
                    }
                }
            } else {
                v56 := []int{fi43, fi43, fi43, fi43}
                _ = v56
            }
            _ = fi43
        }
        v1[((("" + ((v17 + "foo bar") + v17)) + v17) + ("\n\t" + v17))] = 65535
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
}
