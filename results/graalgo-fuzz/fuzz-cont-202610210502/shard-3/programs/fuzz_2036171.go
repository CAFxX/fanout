package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int8; F1 string; F2 int16 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint16; F2 bool }
type GenS6 struct { F0 int32 }
func genFn0() uint {
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(fi6)
        v7 := []int{fi6, fi6, fi6, (fi6 - (fi6 << ((72) & 63)))}
        switch "foo bar" {
        case "\n\t":
            for fi8 := 0; fi8 < 13; fi8++ {
                switch "" {
                case "a":
                    var v9 int16 = 721
                    _ = v9
                }
                fmt.Println(fi8)
                _ = fi8
            }
        case "hello":
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 float64 = 0.0
                v11 += 3.14159
                _ = fi10
                _ = v11
            }
        default:
            if false {
                var v12 uint32 = 42
                switch "" {
                case "foo bar":
                    switch -2147483648 {
                    case 1:
                        for fi13 := 0; fi13 < 19; fi13++ {
                            var v14 int = fi13
                            _ = len(v7)
                            _ = len(v7)
                            _ = fi13
                            _ = v14
                        }
                    case 4294967295:
                        for fi15 := 0; fi15 < 20; fi15++ {
                            switch (fi15 ^ 2147483647) {
                            case 128:
                                var v16 int8 = 127
                                _ = v16
                            default:
                                var v17 int8 = 100
                                if false {
                                    switch (fi15 - fi15) {
                                    case 330:
                                        var v18 float64 = 0.1
                                        _ = v18
                                    }
                                    v19 := GenS1{F0: false, F1: "a"}
                                    v20 := GenS2{F0: v17, F1: "\n\t", F2: -557}
                                    var v21 int = fi15
                                    var v22 float64 = (3.14159 - 0.1)
                                    var v23 int32 = int32(v21)
                                    v24 := GenS3{F0: 42, F1: v21}
                                    fmt.Println(v21, v23, v17)
                                    var v25 int64 = -9223372036854775808
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                                v17 -= 127
                                _ = v17
                            }
                            fmt.Println(fi6, fi6, v12)
                            _ = fi15
                        }
                    }
                case "a":
                    v26 := GenS3{F0: 9223372036854775807, F1: (fi6 << ((0) & 63))}
                    _ = v26
                case "hello":
                    if (66 != -125) {
                        for fi27 := 0; fi27 < 10; fi27++ {
                            if (!(255 > 2)) {
                                v28 := GenS3{F0: 9223372036854775807, F1: fi6}
                                if len(v7) > 0 {
                                    v7[0] = fi6
                                }
                                _ = v28
                            }
                            v12 -= (v12 | v12)
                            var v29 float64 = 0.0
                            _ = fi27
                            _ = v29
                        }
                        v12 += (v12 + (^(v12)))
                    } else {
                        for fi30 := 0; fi30 < 4; fi30++ {
                            v31 := GenS4{F0: 65535, F1: 0}
                            _ = fi30
                            _ = v31
                        }
                        fmt.Println(fi6, fi6)
                    }
                }
                _ = v12
            }
        }
        _ = fi6
        _ = v7
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        var v33 string = "\n\t"
        _ = fi32
        _ = v33
    }
    if false {
        var v34 string = ""
        var v35 uint32 = 298
        _ = v34
        _ = v35
    }
    return 753
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]string)
    var v3 float32 = 0.1
    var v4 string = "\n\t"
    var v5 uint16 = 32767
    for fi36 := 0; fi36 < 5; fi36++ {
        for fi37 := 0; fi37 < 7; fi37++ {
            v3 = (v3 / float32(float64(v3)))
            if true {
                if true {
                    var v38 float32 = v3
                    switch fi37 {
                    case -967:
                        if (v1 > (v1 * (v1 / ((uint8(v5)) | 1)))) {
                            v4 = "a"
                        }
                    }
                    _ = v38
                }
            }
            _ = fi37
        }
        for fi39 := 0; fi39 < 17; fi39++ {
            v40 := GenS5{F0: 1, F1: (v5 / (((v5 / ((v5) | 1))) | 1)), F2: (!true)}
            _ = fi39
            _ = v40
        }
        switch fi36 {
        case 2147483647:
            _ = len(v2)
        case 32767:
            var v41 int16 = 1000
            _ = v41
        case 127:
            v1 -= 93
        }
        _ = fi36
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        for fi43 := 0; fi43 < 5; fi43++ {
            fmt.Println(fi42, v3, v5)
            _ = fi43
        }
        _ = fi42
    }
    switch v4 {
    case "a":
        v4 = "\n\t"
    case "hello":
        fmt.Println(v1, v3)
    case "\n\t":
        fmt.Println(v5, v1, v5)
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        v1 += v1
        for fi45 := 0; fi45 < 6; fi45++ {
            v2[65535] = v4
            _ = fi45
        }
        v3 = v3
        _ = fi44
    }
    var v46 uint8 = (v1 | 0)
    v46 += 127
    switch -857 {
    case 127:
        var v47 bool = (v46 != v46)
        _ = v47
    default:
        v48 := []int{-2147483648, -827, 42, -497}
        _ = v48
    }
    switch v4 {
    case "\n\t":
        switch 127 {
        case -2147483648:
            var v49 int = 127
            _ = v49
        }
    case "":
        v50 := GenS6{F0: -2147483648}
        _ = v50
    }
    switch "foo bar" {
    case "\n\t":
        for fi51 := 0; fi51 < 8; fi51++ {
            v46 = v46
            var v52 int16 = 41
            _ = fi51
            _ = v52
        }
    default:
        v53 := []int{-873, -9223372036854775808, 1}
        _ = v53
    }
    var v54 int64 = -1
    var v55 float64 = 0.0
    var v56 string = (("foo bar" + "hello") + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v54
    _ = v55
    _ = v56
}
