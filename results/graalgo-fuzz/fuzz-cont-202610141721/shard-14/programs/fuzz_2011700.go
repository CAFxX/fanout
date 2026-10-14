package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int) bool {
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := []int{969, (-(fi7))}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: -2147483648, F1: 1, F2: fi7}
        p0 -= (fi7 - 128)
        _ = fi7
        _ = v10
    }
    fmt.Println(p0)
    p0 -= 127
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint64 = 2
        _ = fi11
        _ = v12
    }
    p0 += p0
    _ = p0
    return true
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 96
    var v3 uint32 = 490
    var v4 uint16 = 2
    v5 := make(map[int]bool)
    v6 := make(map[string]string)
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint64 = 18446744073709551615
        v14 -= (v14 >> ((18446744073709551615) & 63))
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 string = "hello"
            var v17 bool = true
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    v18 := GenS2{F0: 65, F1: v4, F2: -515}
    v2 = (v2 >> ((647) & 63))
    fmt.Println(v3, v3)
    var v19 uint16 = 2
    v20 := GenS1{F0: 1, F1: 410, F2: -9223372036854775808}
    v21 := GenS2{F0: 127, F1: v4, F2: 627}
    switch 32768 {
    case -349:
        v22 := GenS2{F0: 3, F1: (^(v19)), F2: 42}
        _ = v22
    case 826:
        for fi23 := 0; fi23 < 16; fi23++ {
            delete(v1, fi23)
            _ = fi23
        }
    }
    if (!(true || true)) {
        v19 = (v4 - v19)
        var v24 float32 = 0.1
        v25 := GenS1{F0: 355, F1: 4294967295, F2: -383}
        _ = v24
        _ = v25
    } else {
        fmt.Println(v4)
        if true {
            for fi26 := 0; fi26 < 10; fi26++ {
                v3 += (v3 | (v3 / ((v3) | 1)))
                v3 = 168
                _ = fi26
            }
            if false {
                if (-2147483648 < 0) {
                    fmt.Println(v3, v19)
                    if ("a" > (("foo bar" + "\n\t") + ("hello" + (("foo bar" + "foo bar") + "\n\t")))) {
                        var v27 int8 = 82
                        v4 = v19
                        v3 = v3
                        _ = v27
                    } else {
                        v28 := GenS3{F0: -611, F1: 128, F2: 1}
                        delete(v1, 81)
                        var v29 uint16 = v4
                        _ = v28
                        _ = v29
                    }
                    if ((v2 * v2) <= v2) {
                        var v30 uint16 = (v19 & v19)
                        var v31 int8 = -128
                        v32 := GenS3{F0: 431, F1: v2, F2: 42}
                        _ = v30
                        _ = v31
                        _ = v32
                    } else {
                        switch "\n\t" {
                        case "a":
                            for fi33 := 0; fi33 < 5; fi33++ {
                                var v34 uint8 = v2
                                var v35 uint8 = (^(v2))
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                        case "":
                            v36 := GenS4{F0: (-711 != -9223372036854775808)}
                            _ = v36
                        case "hello":
                            v37 := []int{-815, -898, 386}
                            _ = v37
                        default:
                            switch "foo bar" {
                            case "\n\t":
                                v38 := GenS3{F0: 9223372036854775807, F1: 42, F2: 4294967295}
                                _ = v38
                            case "":
                                if true {
                                    for fi39 := 0; fi39 < 19; fi39++ {
                                        var v40 int = fi39
                                        switch -314 {
                                        case 128:
                                            for fi41 := 0; fi41 < 17; fi41++ {
                                                var v42 bool = false
                                                var v43 bool = true
                                                _ = fi41
                                                _ = v42
                                                _ = v43
                                            }
                                        case -823:
                                            var v44 int8 = -1
                                            _ = v44
                                        }
                                        _ = fi39
                                        _ = v40
                                    }
                                    var v45 int8 = -1
                                    var v46 int32 = 2147483647
                                    _ = v45
                                    _ = v46
                                }
                            }
                        }
                    }
                } else {
                    for fi47 := 0; fi47 < 18; fi47++ {
                        fmt.Println(v2)
                        for fi48 := 0; fi48 < 7; fi48++ {
                            var v49 int8 = 100
                            var v50 uint32 = v3
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                        _ = fi47
                    }
                }
                if true {
                    v3 = uint32(v2)
                    delete(v1, 0)
                } else {
                    if true {
                        v51 := GenS5{F0: v2}
                        v52 := GenS1{F0: -798, F1: 42, F2: -1}
                        _ = v51
                        _ = v52
                    }
                    v19 = (v4 / (((v4 * 2)) | 1))
                    v53 := GenS5{F0: (v2 / (((v2 >> ((807) & 63))) | 1))}
                    _ = v53
                }
                delete(v5, 411)
            } else {
                switch -1 {
                case -364:
                    v54 := GenS5{F0: v2}
                    _ = v54
                }
                v55 := []int{705, 348, 2147483647, -2147483648}
                _ = v55
            }
            for fi56 := 0; fi56 < 6; fi56++ {
                v3 -= (v3 << ((255) & 63))
                _ = fi56
            }
        }
        switch "" {
        case "hello":
            if (-1 == int8(v4)) {
                v57 := GenS5{F0: 220}
                _ = v57
            }
        default:
            var v58 uint16 = (v19 << ((255) & 63))
            _ = v58
        }
    }
    v4 += v19
    for fi59 := 0; fi59 < 12; fi59++ {
        switch (-(fi59)) {
        case -1:
            var v60 uint8 = 156
            _ = v60
        case 1:
            var v61 int = (fi59 | 4294967295)
            _ = v61
        }
        _ = fi59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
