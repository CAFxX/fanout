package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 int; F1 int64; F2 uint }
type GenS6 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS7 struct { F0 uint; F1 uint64; F2 int8 }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 uint32; F1 int16 }
type GenS10 struct { F0 string; F1 uint32 }
type GenS11 struct { F0 uint16; F1 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 127
    var v3 bool = false
    var v4 int64 = -9223372036854775808
    v5 := []int{}
    for fi6 := 0; fi6 < 9; fi6++ {
        switch (fi6 - fi6) {
        case 128:
            var v7 uint64 = 0
            _ = v7
        default:
            var v8 int = 350
            _ = v8
        }
        _ = fi6
    }
    if ((v4 / ((v4) | 1)) <= -9223372036854775808) {
        var v9 uint = 65535
        fmt.Println(v4, v4, v2)
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 int16 = -32768
            var v12 uint16 = 1
            if v3 {
                var v13 int64 = -9223372036854775808
                for fi14 := 0; fi14 < 4; fi14++ {
                    switch "a" {
                    case "foo bar":
                        v2 = (^(v2))
                    case "\n\t":
                        v15 := GenS1{F0: (v12 % ((v12) | 1))}
                        _ = v15
                    case "hello":
                        v16 := GenS1{F0: (v12 | v12)}
                        _ = v16
                    }
                    var v17 uint = (v9 % ((v9) | 1))
                    var v18 bool = ((3.14159 - (1.5 - (3.14159 / 3.14159))) == 0.1)
                    _ = fi14
                    _ = v17
                    _ = v18
                }
                var v19 uint32 = 2147483647
                _ = v13
                _ = v19
            } else {
                for fi20 := 0; fi20 < 4; fi20++ {
                    v9 += (v9 - (v9 ^ v9))
                    _ = len(v1)
                    v5 = append(v5, (fi20 - (fi20 << (((v9 - v9)) & 63))))
                    _ = len(v1)
                    _ = fi20
                }
                for fi21 := 0; fi21 < 13; fi21++ {
                    v22 := GenS2{F0: int64(v2)}
                    fmt.Println(v4)
                    v3 = (981 < 255)
                    v23 := GenS3{F0: false}
                    v11 = int16(v12)
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                if (v3 || false) {
                    var v24 uint64 = 306
                    var v25 uint16 = (v12 | v12)
                    _ = v24
                    _ = v25
                }
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    switch "hello" {
    case "foo bar":
        for fi26 := 0; fi26 < 8; fi26++ {
            v27 := GenS4{F0: v3, F1: 0}
            switch ((("\n\t" + "a") + "hello") + "hello") {
            case "a":
                if ((!false) && (42 == 165)) {
                    var v28 uint8 = v2
                    delete(v1, "")
                    switch "hello" {
                    case "hello":
                        for fi29 := 0; fi29 < 5; fi29++ {
                            fmt.Println(fi26, fi29)
                            var v30 string = ("foo bar" + ("hello" + "foo bar"))
                            var v31 uint8 = 1
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    case "a":
                        for fi32 := 0; fi32 < 11; fi32++ {
                            var v33 string = "a"
                            var v34 uint8 = (^(v2))
                            if (0 > 2147483647) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            var v35 float32 = ((3.14159 - 1.5) / 1.0)
                            var v36 uint = 0
                            _ = fi32
                            _ = v33
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                    }
                    _ = v28
                } else {
                    v2 += v2
                    var v37 int64 = int64(v2)
                    _ = v37
                }
            case "\n\t":
                var v38 string = ("" + ("foo bar" + ""))
                _ = v38
            }
            _ = fi26
            _ = v27
        }
    case "\n\t":
        var v39 uint64 = 18446744073709551615
        _ = v39
    case "a":
        var v40 uint = 2
        _ = v40
    default:
        if true {
            var v41 uint32 = 42
            v42 := GenS5{F0: 42, F1: 729, F2: 255}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v41
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 16; fi43++ {
        v44 := GenS6{F0: 0, F1: 65535, F2: 2}
        if (-2.25 > 0.0) {
            fmt.Println(fi43, fi43)
            var v45 uint64 = 144
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v45
        }
        _ = fi43
        _ = v44
    }
    v3 = (597 == 4294967295)
    fmt.Println(v3)
    for fi46 := 0; fi46 < 13; fi46++ {
        for fi47 := 0; fi47 < 4; fi47++ {
            switch "a" {
            case "\n\t":
                v3 = (424 > 42)
            case "a":
                var v48 float32 = (1.5 + float32(float64((((3.14159 - (-2.25 + 1.5)) - float32(0.0)) * (float32(1.5) / 3.14159)))))
                _ = v48
            case "":
                var v49 uint64 = 42
                _ = v49
            }
            v4 = 9223372036854775807
            if v3 {
                switch "" {
                case "foo bar":
                    v50 := GenS7{F0: 42, F1: 2, F2: 117}
                    _ = v50
                case "a":
                    switch "a" {
                    case "foo bar":
                        v51 := GenS8{F0: fi46, F1: 328}
                        _ = v51
                    case "\n\t":
                        delete(v1, "")
                        switch "hello" {
                        case "\n\t":
                            v3 = (!(!(((1.5 + (0.0 / 3.14159)) * 3.14159) < 1.5)))
                        }
                        v52 := GenS9{F0: 1, F1: -1}
                        _ = v52
                    default:
                        var v53 int16 = 952
                        _ = v53
                    }
                default:
                    fmt.Println(v2)
                }
                v54 := GenS10{F0: "a", F1: 42}
                if v3 {
                    var v55 uint = 289
                    _ = v55
                } else {
                    switch 907 {
                    case 9223372036854775807:
                        v56 := GenS10{F0: "foo bar", F1: 2147483648}
                        _ = v56
                    case -250:
                        if len(v5) > 0 {
                            v5[0] = 32767
                        }
                    }
                    fmt.Println(v3, fi46)
                    var v57 uint = uint(v2)
                    var v58 int64 = -215
                    _ = len(v1)
                    _ = v57
                    _ = v58
                }
                _ = v54
            } else {
                var v59 bool = (v3 && false)
                v60 := GenS11{F0: 32767, F1: 400}
                _ = v59
                _ = v60
            }
            _ = fi47
        }
        if (!(0.0 >= -2.25)) {
            if v3 {
                var v61 uint32 = 65535
                v61 -= 1
                _ = v61
            } else {
                var v62 uint = 185
                _ = v62
            }
        } else {
            v3 = v3
            for fi63 := 0; fi63 < 12; fi63++ {
                v5 = append(v5, (fi63 % (((-(fi46))) | 1)))
                switch 788 {
                case -370:
                    switch (fi63 ^ fi46) {
                    case -1:
                        _ = len(v5)
                    case -945:
                        v64 := GenS10{F0: ("a" + "\n\t"), F1: 71}
                        _ = v64
                    }
                case 628:
                    v3 = v3
                default:
                    var v65 string = "foo bar"
                    _ = v65
                }
                _ = fi63
            }
        }
        _ = len(v5)
        _ = fi46
    }
    var v66 bool = false
    if (true && v3) {
        v3 = false
        var v67 uint16 = 52
        _ = v67
    } else {
        v5 = append(v5, 9223372036854775807)
        v2 += 127
        v3 = (128 != 791)
    }
    for fi68 := 0; fi68 < 4; fi68++ {
        for fi69 := 0; fi69 < 10; fi69++ {
            fmt.Println(v4, v3, v3)
            var v70 float64 = 0.1
            switch (("foo bar" + "\n\t") + "foo bar") {
            case "foo bar":
                _ = len(v5)
            case "a":
                fmt.Println(v3, v2)
            }
            _ = fi69
            _ = v70
        }
        _ = fi68
    }
    var v71 uint8 = (v2 | (v2 | v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v66:", v66)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v71
}
