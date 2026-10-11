package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 string; F2 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 bool; F2 int16 }

func main() {
    var v1 int16 = -141
    v2 := make(map[string]int)
    var v3 uint16 = 0
    var v4 uint32 = 842
    var v5 uint = uint(v4)
    var v6 int64 = -702
    for fi7 := 0; fi7 < 12; fi7++ {
        if (true || (83 == 128)) {
            for fi8 := 0; fi8 < 18; fi8++ {
                switch (("foo bar" + ("a" + ("\n\t" + (("hello" + "\n\t") + "foo bar")))) + ("hello" + "a")) {
                case "\n\t":
                    fmt.Println(v4, v5)
                case "":
                    v9 := GenS1{F0: (v1 ^ v1)}
                    _ = v9
                case "foo bar":
                    if true {
                        if false {
                            var v10 int16 = v1
                            _ = v10
                        }
                    } else {
                        delete(v2, "")
                        switch "" {
                        case "\n\t":
                            v3 = 42
                        case "a":
                            v5 += v5
                        case "hello":
                            v11 := []int{(-(fi7)), (fi8 ^ -241), fi7, (fi8 * (fi7 * (fi8 >> (((v5 & 42)) & 63))))}
                            _ = v11
                        }
                    }
                }
                _ = fi8
            }
            if false {
                var v12 int8 = -1
                var v13 uint = v5
                if (1 >= 7) {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v15 := GenS2{F0: v6, F1: (("foo bar" + "a") + ""), F2: fi14}
                        var v16 int64 = (v6 - (v6 ^ -509))
                        if ((1e10 + 3.14159) > 1.5) {
                            var v17 float32 = -2.25
                            _ = v17
                        }
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    if false {
                        v6 += -9223372036854775808
                        v13 -= (v5 * v5)
                        v1 = v1
                    } else {
                        v18 := GenS2{F0: v6, F1: "a", F2: (fi7 ^ fi7)}
                        fmt.Println(fi7)
                        for fi19 := 0; fi19 < 11; fi19++ {
                            v13 = uint(v4)
                            fmt.Println(v3)
                            fmt.Println(v4)
                            _ = fi19
                        }
                        _ = v18
                    }
                }
                _ = v12
                _ = v13
            }
        }
        fmt.Println(v4)
        v20 := GenS1{F0: int16(v5)}
        _ = fi7
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 float32 = (((1.5 * 1e10) * (((3.14159 - float32(1.5)) + float32((-2.25 * 0.0))) - 3.14159)) * -2.25)
        _ = fi21
        _ = v22
    }
    v5 = uint(v3)
    if true {
        var v23 int8 = 42
        for fi24 := 0; fi24 < 16; fi24++ {
            v25 := []int{fi24, (fi24 - (fi24 / (((-(fi24))) | 1))), (fi24 + 9223372036854775807), (fi24 & (^(fi24)))}
            _ = fi24
            _ = v25
        }
        _ = v23
    } else {
        var v26 int32 = -1
        _ = v26
    }
    var v27 int = 9223372036854775807
    switch "foo bar" {
    case "foo bar":
        v28 := []int{(v27 << ((v5) & 63))}
        _ = v28
    }
    if true {
        v29 := []int{(v27 | (v27 >> ((65535) & 63))), v27}
        _ = v29
    } else {
        var v30 bool = false
        var v31 int8 = 127
        if v30 {
            var v32 float64 = (-2.25 * 1e10)
            v33 := GenS1{F0: v1}
            if (v1 <= v1) {
                switch (v27 ^ (^(v27))) {
                case 4294967295:
                    var v34 string = ""
                    _ = v34
                case -874:
                    v35 := []int{}
                    _ = v35
                case 2147483647:
                    var v36 uint8 = 128
                    _ = v36
                default:
                    if v30 {
                        v27 = (v27 | 404)
                    } else {
                        switch (v27 % ((int(v31)) | 1)) {
                        case 951:
                            v37 := GenS3{F0: 42}
                            _ = v37
                        case 1:
                            var v38 string = "foo bar"
                            _ = v38
                        default:
                            v27 = v27
                        }
                        var v39 int8 = (v31 - v31)
                        _ = v39
                    }
                }
                var v40 int64 = v6
                _ = v40
            }
            _ = v32
            _ = v33
        }
        _ = v30
        _ = v31
    }
    if ((false && (false || (!(65535 != (v4 | 1))))) || false) {
        v27 -= v27
        v41 := []int{(v27 * v27), v27, int(v1), (-(v27))}
        _ = v41
    } else {
        v42 := GenS2{F0: (v6 ^ (^(v6))), F1: ("a" + "a"), F2: (v27 >> (((^(v5))) & 63))}
        _ = v42
    }
    var v43 uint64 = 924
    if ((v43 * 18446744073709551615) >= 1) {
        v6 = (v6 ^ v6)
        fmt.Println(v1, v43, v4)
    }
    var v44 float32 = ((float32(1.5) + (((0.0 - 1.5) / 1.5) + (0.0 * 0.1))) - float32(float64((((1e10 - 0.1) - 3.14159) - (float32(0.1) * 1e10)))))
    var v45 float64 = (3.14159 * 1.5)
    switch 0 {
    case 128:
        for fi46 := 0; fi46 < 13; fi46++ {
            if false {
                v1 = v1
                var v47 int16 = v1
                _ = v47
            } else {
                for fi48 := 0; fi48 < 20; fi48++ {
                    var v49 uint = (^(v5))
                    _ = fi48
                    _ = v49
                }
                fmt.Println(v3)
            }
            if ("a" <= "foo bar") {
                if (100 > 100) {
                    v50 := []int{32767, -86}
                    _ = v50
                }
                var v51 uint8 = 128
                v52 := GenS3{F0: 4294967295}
                _ = v51
                _ = v52
            }
            _ = fi46
        }
    case 4294967295:
        v53 := GenS4{F0: 255, F1: true, F2: (v1 % ((int16(v3)) | 1))}
        _ = v53
    case -326:
        _ = len(v2)
    default:
        if (v3 >= 2) {
            if (!true) {
                var v54 bool = (!true)
                for fi55 := 0; fi55 < 5; fi55++ {
                    for fi56 := 0; fi56 < 2; fi56++ {
                        var v57 uint16 = (v3 << (((v5 >> (((v5 | 268)) & 63))) & 63))
                        _ = fi56
                        _ = v57
                    }
                    v3 = (v3 + v3)
                    _ = fi55
                }
                switch "\n\t" {
                case "\n\t":
                    v6 = -1
                default:
                    var v58 int64 = int64(v1)
                    _ = v58
                }
                _ = v54
            } else {
                fmt.Println(v45)
                switch "a" {
                case "foo bar":
                    v27 -= int(v1)
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
