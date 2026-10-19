package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint }

func main() {
    v1 := []uint32{}
    var v2 string = "\n\t"
    v3 := []string{}
    var v4 int = 32768
    var v5 uint16 = 1
    v6 := make(map[int]bool)
    if (!true) {
        for fi7 := 0; fi7 < 6; fi7++ {
            if (true || true) {
                v8 := GenS1{F0: 32767, F1: 0, F2: 1}
                var v9 uint8 = 1
                _ = v8
                _ = v9
            }
            _ = fi7
        }
        v10 := v6[v4]
        var v11 float64 = 3.14159
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v4, v2, v4)
        case "":
            var v13 int32 = -2147483648
            _ = v13
        case "foo bar":
            v14 := GenS1{F0: v4, F1: 498, F2: 42}
            _ = v14
        default:
            var v15 float32 = 1e10
            _ = v15
        }
        for fi16 := 0; fi16 < 5; fi16++ {
            if false {
                fmt.Println(v5, fi16, fi16)
            } else {
                fmt.Println(v2, v5, fi16)
            }
            v5 -= 944
            var v17 int64 = 9223372036854775807
            v4 = fi12
            switch v2 {
            case "foo bar":
                v18 := GenS1{F0: v4, F1: 104, F2: 0}
                _ = v18
            }
            switch (fi16 | (^(fi16))) {
            case 42:
                if len(v1) > 0 {
                    v1[0] = 711
                }
            }
            _ = fi16
            _ = v17
        }
        _ = fi12
    }
    v1 = append(v1, 4294967295)
    if false {
        var v19 uint32 = 732
        switch "" {
        case "":
            switch v2 {
            case "":
                v2 = v2
            case "a":
                var v20 int32 = 1
                _ = v20
            case "\n\t":
                var v21 int16 = -772
                _ = v21
            default:
                switch "hello" {
                case "\n\t":
                    v22 := GenS1{F0: 32767, F1: 733, F2: 127}
                    _ = v22
                case "a":
                    v4 += v4
                case "hello":
                    if false {
                        v19 -= 65535
                    }
                default:
                    for fi23 := 0; fi23 < 1; fi23++ {
                        v24 := GenS1{F0: v4, F1: 0, F2: 42}
                        v6[(fi23 ^ (v4 << ((1) & 63)))] = true
                        _ = fi23
                        _ = v24
                    }
                }
            }
        case "foo bar":
            v19 = uint32(v5)
        }
        v25 := GenS2{F0: 13}
        _ = v19
        _ = v25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        if (100 > 127) {
            v27 := GenS1{F0: 4294967295, F1: 715, F2: 127}
            _ = v27
        }
        var v28 uint32 = 42
        switch v2 {
        case "a":
            switch v4 {
            case -1:
                v4 = v4
            default:
                v28 += 981
            }
        case "":
            var v29 int = (v4 * fi26)
            _ = v29
        case "foo bar":
            switch fi26 {
            case 32768:
                for fi30 := 0; fi30 < 9; fi30++ {
                    var v31 int64 = 753
                    if false {
                        var v32 int16 = 1000
                        _ = v32
                    } else {
                        v5 += (v5 << ((255) & 63))
                    }
                    _ = fi30
                    _ = v31
                }
            }
        default:
            if ((false && true) || ((!(!true)) || ((((float64(0.1) + 1.5) == (1.5 + 3.14159)) && true) || true))) {
                v4 = (fi26 / (((v4 + v4)) | 1))
                v33 := GenS1{F0: 42, F1: 42, F2: 232}
                v28 = 4294967295
                _ = v33
            } else {
                fmt.Println(v4, fi26, v28)
                v34 := GenS2{F0: 42}
                _ = v34
            }
        }
        _ = fi26
        _ = v28
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        var v36 int8 = 101
        _ = fi35
        _ = v36
    }
    if true {
        var v37 uint64 = 4294967295
        _ = v37
    }
    v2 = v2
    switch v4 {
    case 32768:
        var v38 bool = (false || (!(1 <= 18446744073709551615)))
        _ = v38
    default:
        for fi39 := 0; fi39 < 13; fi39++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = len(v1)
            switch 42 {
            case -505:
                v4 = (fi39 & v4)
            case -1:
                if ((true || true) || false) {
                    for fi40 := 0; fi40 < 17; fi40++ {
                        switch (fi40 + v4) {
                        case 9223372036854775807:
                            var v41 int64 = 9223372036854775807
                            _ = v41
                        case 828:
                            v5 -= v5
                        case 71:
                            v5 -= (v5 & (^(v5)))
                        }
                        _ = fi40
                    }
                    var v42 bool = false
                    _ = v42
                }
            }
            _ = fi39
        }
    }
    v2 = (v2 + v2)
    v4 -= v4
    v5 -= (v5 >> ((2) & 63))
    for fi43 := 0; fi43 < 6; fi43++ {
        v44 := GenS3{F0: 838, F1: 542}
        v45 := GenS3{F0: 822, F1: 4294967295}
        _ = fi43
        _ = v44
        _ = v45
    }
    fmt.Println(v2, v4)
    v46 := GenS2{F0: 38}
    v2 = "a"
    if false {
        if (954 == 921) {
            v47 := GenS2{F0: 2}
            _ = v47
        } else {
            v5 -= v5
            switch ((("\n\t" + v2) + "") + v2) {
            case "foo bar":
                var v48 uint16 = 1
                _ = v48
            default:
                for fi49 := 0; fi49 < 17; fi49++ {
                    if (!false) {
                        fmt.Println(v4)
                        v5 -= 0
                        var v50 uint = 4294967295
                        _ = v50
                    } else {
                        var v51 int32 = -2147483648
                        var v52 uint8 = 128
                        v51 -= (v51 << ((2) & 63))
                        _ = v51
                        _ = v52
                    }
                    v53 := v6[(^(fi49))]
                    _ = fi49
                    _ = v53
                }
            }
            var v54 int64 = -213
            _ = v54
        }
        v4 = v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
