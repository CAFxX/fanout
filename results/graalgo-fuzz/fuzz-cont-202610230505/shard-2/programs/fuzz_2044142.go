package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int8; F1 int32; F2 string }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 uint16; F1 uint32 }
func genFn0() int8 {
    v4 := GenS1{F0: 1, F1: 0}
    var v5 float64 = 3.14159
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]uint32)
    var v3 float32 = 0.0
    if false {
        v6 := GenS2{F0: 4294967295, F1: 127}
        _ = v6
    }
    var v7 int64 = -9223372036854775808
    v8 := GenS3{F0: 985, F1: -378, F2: 2147483647}
    v9 := GenS4{F0: 18446744073709551615, F1: 0}
    v7 += v7
    var v10 int = 9223372036854775807
    var v11 uint = 18446744073709551615
    _ = len(v2)
    if true {
        var v12 float64 = 1.5
        delete(v2, (-(v10)))
        var v13 int = (v10 - v10)
        _ = v12
        _ = v13
    }
    v11 -= 280
    for fi14 := 0; fi14 < 7; fi14++ {
        switch 161 {
        case 9223372036854775807:
            if (-40 >= 100) {
                switch fi14 {
                case 820:
                    v15 := GenS3{F0: 4294967295, F1: (v7 / (((v7 + v7)) | 1)), F2: 2147483647}
                    _ = v15
                case 42:
                    v3 += v3
                case 618:
                    v10 += (^(fi14))
                }
                v16 := GenS4{F0: 987, F1: 1}
                var v17 uint16 = 3
                _ = v16
                _ = v17
            }
        case 32768:
            var v18 int32 = 1
            _ = v18
        case -1:
            if true {
                var v19 uint32 = 631
                for fi20 := 0; fi20 < 11; fi20++ {
                    switch (v1 + "") {
                    case "a":
                        v11 += (v11 - 2)
                    default:
                        v21 := GenS3{F0: (v19 + 719), F1: (^(v7)), F2: (v19 + (v19 | v19))}
                        _ = v21
                    }
                    v22 := []int{(fi14 << ((v11) & 63)), -9223372036854775808, fi20}
                    _ = fi20
                    _ = v22
                }
                for fi23 := 0; fi23 < 17; fi23++ {
                    v1 = ("a" + v1)
                    var v24 uint16 = 2
                    _ = fi23
                    _ = v24
                }
                _ = v19
            }
        }
        _ = fi14
    }
    switch "" {
    case "":
        if (34 != 255) {
            if true {
                var v25 int16 = -32768
                delete(v2, v10)
                _ = v25
            }
            v26 := GenS3{F0: 550, F1: (^(v7)), F2: 901}
            for fi27 := 0; fi27 < 3; fi27++ {
                for fi28 := 0; fi28 < 14; fi28++ {
                    if (true || false) {
                        v29 := []int{(fi28 | (v10 % ((fi28) | 1))), fi28, -2147483648}
                        fmt.Println(v7, fi27)
                        switch fi27 {
                        case 1:
                            v3 -= -2.25
                        case 9223372036854775807:
                            switch ("\n\t" + "a") {
                            case "":
                                if ((2 > v11) || (!false)) {
                                    switch v1 {
                                    case "":
                                        fmt.Println(v3, v7, v3)
                                        delete(v2, (-(v10)))
                                    }
                                } else {
                                    var v30 int64 = int64(v11)
                                    _ = v30
                                }
                            }
                        case -824:
                            _ = len(v29)
                        }
                        _ = v29
                    }
                    v3 -= v3
                    _ = fi28
                }
                v7 = v7
                if true {
                    _ = len(v2)
                    v31 := GenS4{F0: 18446744073709551615, F1: 32767}
                    v32 := GenS5{F0: genFn0(), F1: -2147483648, F2: (("" + v1) + v1)}
                    _ = v31
                    _ = v32
                }
                _ = fi27
            }
            _ = v26
        }
    default:
        v33 := v2[v10]
        _ = v33
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        for fi35 := 0; fi35 < 19; fi35++ {
            v7 -= v7
            if true {
                var v36 int64 = v7
                v37 := GenS2{F0: v11, F1: 255}
                _ = v36
                _ = v37
            } else {
                var v38 string = (v1 + (v1 + "hello"))
                delete(v2, fi35)
                if false {
                    var v39 uint64 = 836
                    _ = v39
                } else {
                    v10 = int(v7)
                    var v40 uint8 = 128
                    v41 := GenS6{F0: 1, F1: 4294967295}
                    _ = v40
                    _ = v41
                }
                v3 = 3.14159
                _ = v38
            }
            var v42 float32 = v3
            _ = fi35
            _ = v42
        }
        _ = fi34
    }
    for fi43 := 0; fi43 < 14; fi43++ {
        var v44 int = fi43
        for fi45 := 0; fi45 < 19; fi45++ {
            fmt.Println(v44, fi45, v11)
            v46 := GenS5{F0: -32, F1: 2147483647, F2: ("hello" + ((v1 + "foo bar") + v1))}
            _ = fi45
            _ = v46
        }
        v11 += (v11 | (v11 << (((v11 / ((v11) | 1))) & 63)))
        _ = fi43
        _ = v44
    }
    v47 := GenS7{F0: 1, F1: 42}
    fmt.Println(v3, v3, v11)
    v7 = 1
    var v48 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v7
    _ = v8
    _ = v9
}
