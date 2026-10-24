package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 bool; F2 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 int8; F2 uint64 }

func main() {
    var v1 int8 = 127
    var v2 int32 = 1
    v3 := make(map[string]int)
    var v4 uint64 = 1
    var v5 float64 = 1e10
    var v6 uint8 = 128
    _ = len(v3)
    var v7 float64 = (float64(((1e10 / -2.25) * (float32(-2.25) + ((3.14159 / 3.14159) * 3.14159)))) / 3.14159)
    v8 := GenS1{F0: 787}
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: 0}
        var v11 int64 = 1
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint16 = 1
    switch 0 {
    case 2147483647:
        fmt.Println(v4, v5, v4)
    case -2147483648:
        switch 2147483647 {
        case -701:
            v13 := []int{}
            _ = v13
        case 42:
            var v14 bool = (951 <= 24)
            _ = v14
        case 127:
            if (42 <= 603) {
                v15 := []int{int(v1)}
                delete(v3, "")
                var v16 uint32 = 2147483647
                _ = v15
                _ = v16
            } else {
                var v17 uint = uint(v6)
                fmt.Println(v2, v4)
                switch "" {
                case "\n\t":
                    var v18 uint = (v17 + 0)
                    _ = v18
                case "hello":
                    delete(v3, "hello")
                }
                _ = v17
            }
        }
    default:
        v19 := GenS1{F0: -797}
        _ = v19
    }
    if (!false) {
        var v20 string = "\n\t"
        if false {
            v21 := GenS1{F0: -1}
            switch 127 {
            case 42:
                switch 32767 {
                case 127:
                    if (!(v5 != 1.5)) {
                        for fi22 := 0; fi22 < 5; fi22++ {
                            v1 = (v1 * (v1 / ((v1) | 1)))
                            _ = fi22
                        }
                        fmt.Println(v7, v1)
                        var v23 float32 = float32(0.0)
                        _ = v23
                    } else {
                        if false {
                            v12 -= 32767
                            var v24 uint32 = 2147483647
                            if true {
                                var v25 int16 = -860
                                _ = v25
                            } else {
                                for fi26 := 0; fi26 < 7; fi26++ {
                                    for fi27 := 0; fi27 < 2; fi27++ {
                                        v2 = (^(v2))
                                        v7 = ((v7 + 3.14159) / (float64(0.0) + ((v5 / (0.1 / float64(((-2.25 - 0.0) - 3.14159)))) - 1.5)))
                                        _ = fi27
                                    }
                                    v1 += (v1 ^ 100)
                                    _ = fi26
                                }
                            }
                            _ = v24
                        } else {
                            fmt.Println(v12, v1)
                            var v28 int8 = v1
                            _ = v28
                        }
                        v29 := GenS1{F0: -1}
                        _ = v29
                    }
                }
            case 9223372036854775807:
                var v30 int16 = 0
                _ = v30
            default:
                v31 := GenS2{F0: 128}
                _ = v31
            }
            _ = v21
        } else {
            delete(v3, v20)
        }
        _ = v20
    } else {
        if (false && ("\n\t" == ("hello" + ((("a" + "\n\t") + "\n\t") + "")))) {
            var v32 uint = 255
            if false {
                v1 += 42
                if (v7 >= 0.0) {
                    var v33 uint32 = 1
                    v1 -= v1
                    _ = v33
                } else {
                    v34 := GenS2{F0: 410}
                    v5 += float64(-2.25)
                    v6 = (v6 + (v6 + v6))
                    _ = v34
                }
                v5 = v5
            } else {
                for fi35 := 0; fi35 < 1; fi35++ {
                    switch "\n\t" {
                    case "foo bar":
                        v2 = v2
                    case "a":
                        var v36 int16 = 1
                        _ = v36
                    }
                    _ = fi35
                }
            }
            var v37 int16 = 1
            _ = v32
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        v39 := GenS2{F0: fi38}
        v40 := GenS3{F0: 128, F1: false, F2: int(v4)}
        v5 = v5
        _ = fi38
        _ = v39
        _ = v40
    }
    v6 -= (v6 >> ((4294967295) & 63))
    var v41 bool = (true || false)
    delete(v3, "hello")
    fmt.Println(v1)
    var v42 uint8 = (^(v6))
    fmt.Println(v1, v1)
    if v41 {
        if ((v6 % (((v42 & (v42 << ((452) & 63)))) | 1)) >= v42) {
            var v43 int64 = 1
            if (!v41) {
                for fi44 := 0; fi44 < 7; fi44++ {
                    v45 := GenS4{F0: int32(v12)}
                    _ = fi44
                    _ = v45
                }
            }
            fmt.Println(v5, v2)
            _ = v43
        }
        v12 = v12
    } else {
        v46 := GenS5{F0: -9223372036854775808, F1: (v1 * v1), F2: (v4 | (v4 * 1))}
        _ = v46
    }
    v2 += 42
    v6 = (v42 << ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
