package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint16 }
type GenS2 struct { F0 int32; F1 int16; F2 int32 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint16; F1 string; F2 int8 }

func main() {
    var v1 uint8 = 2
    v2 := []int{32767, -849}
    v3 := make(map[uint32]string)
    var v4 int = -2147483648
    v5 := make(map[uint32]bool)
    switch "a" {
    case "foo bar":
        var v6 float64 = 1.5
        _ = v6
    default:
        v7 := GenS1{F0: 1, F1: 0, F2: 2}
        _ = v7
    }
    var v8 float32 = ((0.0 + 1.5) * float32(0.0))
    var v9 uint64 = 4294967295
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 10; fi11++ {
            if (0 != 1) {
                v1 -= 128
                var v12 uint64 = (^(v9))
                _ = v12
            } else {
                switch ("hello" + "hello") {
                case "a":
                    var v13 uint16 = 2
                    _ = v13
                case "\n\t":
                    delete(v5, 65535)
                case "foo bar":
                    var v14 int32 = 1
                    _ = v14
                }
                var v15 int = (v4 & (fi10 - fi11))
                v4 = (fi11 + v4)
                var v16 string = "hello"
                _ = v15
                _ = v16
            }
            _ = fi11
        }
        for fi17 := 0; fi17 < 1; fi17++ {
            delete(v3, 1)
            _ = fi17
        }
        _ = fi10
    }
    v18 := GenS1{F0: 100, F1: 32767, F2: 255}
    v9 = (^(v9))
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := GenS1{F0: 100, F1: -1, F2: 32767}
        var v21 int64 = -9223372036854775808
        _ = fi19
        _ = v20
        _ = v21
    }
    if true {
        v8 = float32(float64(-2.25))
        v9 += 4294967295
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 int64 = 9223372036854775807
        for fi24 := 0; fi24 < 15; fi24++ {
            if (false && (0 > 255)) {
                var v25 uint8 = v1
                v23 += (v23 | 0)
                if len(v2) > 0 {
                    v2[0] = 9223372036854775807
                }
                _ = v25
            }
            var v26 int64 = v23
            var v27 float64 = 3.14159
            _ = fi24
            _ = v26
            _ = v27
        }
        var v28 int16 = int16(v1)
        _ = fi22
        _ = v23
        _ = v28
    }
    var v29 float64 = (-2.25 * 0.1)
    var v30 uint32 = 2
    for fi31 := 0; fi31 < 20; fi31++ {
        v5[(v30 / ((0) | 1))] = true
        _ = fi31
    }
    switch v4 {
    case 723:
        switch "" {
        case "foo bar":
            v4 -= (v4 + (v4 * (v4 & v4)))
        case "hello":
            v29 += -2.25
        }
    case 32768:
        switch (-(v4)) {
        case 32768:
            if (false || (18446744073709551615 > 4294967295)) {
                var v32 uint = 2
                _ = v32
            } else {
                v8 = ((float32((v29 * (v29 + float64(v8)))) / (3.14159 / v8)) / v8)
                for fi33 := 0; fi33 < 8; fi33++ {
                    var v34 bool = false
                    for fi35 := 0; fi35 < 15; fi35++ {
                        v4 = (fi35 - 32768)
                        if len(v2) > 0 {
                            v2[0] = (v4 << ((255) & 63))
                        }
                        _ = fi35
                    }
                    var v36 int64 = -368
                    _ = fi33
                    _ = v34
                    _ = v36
                }
            }
        case 9223372036854775807:
            var v37 uint64 = uint64(v30)
            _ = v37
        case -370:
            var v38 int16 = 138
            _ = v38
        default:
            v5[v30] = (v4 >= (v4 | v4))
        }
    case 742:
        for fi39 := 0; fi39 < 10; fi39++ {
            var v40 string = ("a" + "a")
            _ = fi39
            _ = v40
        }
    }
    v41 := GenS2{F0: 0, F1: -1, F2: -1}
    var v42 int16 = 0
    switch v4 {
    case -859:
        for fi43 := 0; fi43 < 4; fi43++ {
            if true {
                var v44 bool = (918 > 1)
                fmt.Println(fi43, v30, v8)
                if true {
                    var v45 float64 = ((v29 + 3.14159) - float64((v8 - float32((((v29 * v29) - float64(v8)) - 3.14159)))))
                    if v44 {
                        if (!v44) {
                            v45 += 0.1
                            if v44 {
                                var v46 uint16 = 0
                                _ = v46
                            }
                            switch "foo bar" {
                            case "\n\t":
                                v45 = v45
                            case "":
                                switch v4 {
                                case 20:
                                    for fi47 := 0; fi47 < 15; fi47++ {
                                        fmt.Println(v9, v45)
                                        fmt.Println(v9, v30, fi43)
                                        var v48 uint = 42
                                        _ = fi47
                                        _ = v48
                                    }
                                case 943:
                                    v9 = (v9 + (v9 << ((uint(v30)) & 63)))
                                }
                            default:
                                var v49 int = v4
                                _ = v49
                            }
                        }
                        var v50 uint = 1
                        _ = v50
                    } else {
                        if false {
                            v4 = v4
                            for fi51 := 0; fi51 < 15; fi51++ {
                                v52 := GenS3{F0: v1, F1: ((!(0 != 649)) && ((v9 << ((951) & 63)) > 2))}
                                v4 = (v4 - 127)
                                var v53 uint8 = 2
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                            v54 := GenS4{F0: 2, F1: "a", F2: 113}
                            _ = v54
                        }
                        v8 = 3.14159
                    }
                    _ = v45
                } else {
                    var v55 float32 = v8
                    v4 -= v4
                    v56 := GenS3{F0: (v1 % (((v1 - 1)) | 1)), F1: (909 > 32767)}
                    _ = v55
                    _ = v56
                }
                _ = v44
            }
            _ = fi43
        }
    case 0:
        fmt.Println(v9)
    }
    v57 := GenS2{F0: 2147483647, F1: v42, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v57
    _ = v8
    _ = v9
}
