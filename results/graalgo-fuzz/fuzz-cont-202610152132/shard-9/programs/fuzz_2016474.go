package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int64; F2 int32 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 bool; F2 uint64 }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 int16; F1 int32; F2 bool }
func genFn0() int16 {
    for fi4 := 0; fi4 < 10; fi4++ {
        for fi5 := 0; fi5 < 2; fi5++ {
            if (!((-128 >= -128) && false)) {
                if (65535 <= 0) {
                    var v6 uint64 = 4294967295
                    v6 -= 963
                    _ = v6
                }
                var v7 int32 = 296
                _ = v7
            } else {
                v8 := GenS1{F0: 494}
                _ = v8
            }
            v9 := []int{fi4}
            var v10 float64 = ((((float64(float32(3.14159)) * 1e10) - (float64((0.0 / ((((1e10 * 3.14159) / 1e10) * float32(1e10)) / 1.0))) / 3.14159)) / 0.1) - float64(1.5))
            _ = fi5
            _ = v10
            _ = v9
        }
        _ = fi4
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if false {
            switch "" {
            case "foo bar":
                fmt.Println(fi11)
            default:
                fmt.Println(fi11)
            }
            if (!(true || false)) {
                v12 := []int{(fi11 * -241), -843}
                _ = v12
            }
            switch "a" {
            case "hello":
                var v13 uint8 = 2
                _ = v13
            default:
                var v14 int32 = 79
                _ = v14
            }
        } else {
            for fi15 := 0; fi15 < 5; fi15++ {
                v16 := []int{}
                var v17 uint32 = 2147483647
                v16 = append(v16, fi15)
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 float64 = (1.5 - 1e10)
            _ = v18
        }
        v19 := GenS2{F0: 1, F1: 352, F2: -1}
        v20 := GenS1{F0: 305}
        _ = fi11
        _ = v19
        _ = v20
    }
    if (!true) {
        var v21 int32 = -1
        fmt.Println(v21)
        _ = v21
    } else {
        v22 := GenS2{F0: 2, F1: 0, F2: 42}
        var v23 int64 = 1
        v24 := GenS2{F0: 0, F1: -159, F2: 42}
        _ = v22
        _ = v23
        _ = v24
    }
    return 1
}

func main() {
    var v1 uint16 = 0
    var v2 float64 = 0.1
    var v3 string = "\n\t"
    v25 := GenS3{F0: 42, F1: 4294967295}
    for fi26 := 0; fi26 < 13; fi26++ {
        v27 := GenS4{F0: 1000}
        var v28 string = "\n\t"
        for fi29 := 0; fi29 < 15; fi29++ {
            v30 := []int{85, 264, fi29}
            _ = fi29
            _ = v30
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        fmt.Println(v2, v2)
        switch "\n\t" {
        case "a":
            if (!(510 == 104)) {
                var v32 int64 = 1
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v3, fi31, v2)
                case "":
                    var v33 int16 = -657
                    _ = v33
                case "foo bar":
                    v34 := GenS4{F0: 276}
                    _ = v34
                }
                v1 = (v1 << ((419) & 63))
                _ = v32
            } else {
                var v35 uint32 = 2
                _ = v35
            }
        case "":
            v36 := GenS5{F0: false, F1: (18446744073709551615 > 1), F2: 2}
            _ = v36
        }
        fmt.Println(v3, fi31, fi31)
        _ = fi31
    }
    v2 = 3.14159
    v37 := GenS3{F0: 255, F1: 2147483647}
    var v38 uint32 = 65535
    var v39 uint64 = 4294967295
    fmt.Println(v2, v39)
    var v40 uint = 255
    if true {
        if false {
            v41 := []int{32768, -514, -281, -446}
            if (1 != 1) {
                if true {
                    v41 = append(v41, -759)
                } else {
                    _ = len(v41)
                    v3 = "hello"
                    v3 = v3
                }
                var v42 uint = v40
                _ = v42
            } else {
                v43 := GenS1{F0: v38}
                v38 = (v38 - v38)
                _ = v43
            }
            switch "foo bar" {
            case "a":
                for fi44 := 0; fi44 < 5; fi44++ {
                    _ = len(v41)
                    _ = fi44
                }
            default:
                if (true || (!((^(v1)) != (v1 / (((v1 | v1)) | 1))))) {
                    v45 := GenS6{F0: 1, F1: false}
                    v46 := GenS3{F0: uint8(v40), F1: -9223372036854775808}
                    v39 -= 19
                    _ = v45
                    _ = v46
                } else {
                    _ = len(v41)
                }
            }
            _ = v41
        } else {
            for fi47 := 0; fi47 < 3; fi47++ {
                switch v3 {
                case "":
                    if (true && (false && false)) {
                        fmt.Println(v3, v40, v2)
                    }
                }
                var v48 int64 = 9223372036854775807
                v49 := GenS4{F0: -1}
                _ = fi47
                _ = v48
                _ = v49
            }
        }
    }
    v50 := []int{4294967295, -9223372036854775808, 1, -1}
    if false {
        v50 = append(v50, 0)
        if len(v50) >= 2 {
            v50 = v50[1:len(v50)]
        }
        var v51 uint = v40
        _ = v51
    } else {
        for fi52 := 0; fi52 < 13; fi52++ {
            var v53 float64 = v2
            _ = fi52
            _ = v53
        }
        switch ("hello" + v3) {
        case "foo bar":
            v54 := GenS7{F0: 0, F1: 2147483647, F2: true}
            _ = v54
        case "\n\t":
            if (!false) {
                v1 += (v1 % (((v1 - (v1 << ((255) & 63)))) | 1))
                v55 := GenS3{F0: 0, F1: -9223372036854775808}
                v56 := GenS6{F0: (v1 >> (((^(v40))) & 63)), F1: true}
                _ = v55
                _ = v56
            } else {
                _ = len(v50)
            }
        case "hello":
            var v57 string = v3
            _ = v57
        default:
            v50 = append(v50, -2147483648)
        }
        switch 4294967295 {
        case 632:
            switch (("\n\t" + ("foo bar" + ("" + "hello"))) + "a") {
            case "foo bar":
                switch v3 {
                case "":
                    if len(v50) >= 2 {
                        v50 = v50[1:len(v50)]
                    }
                case "a":
                    if len(v50) > 0 {
                        v50[0] = 32767
                    }
                case "foo bar":
                    v38 -= v38
                }
            case "":
                v39 = 18446744073709551615
            default:
                var v58 uint64 = (v39 << ((v40) & 63))
                _ = v58
            }
        }
    }
    v59 := GenS5{F0: true, F1: (877 <= 2147483648), F2: (v39 + uint64(v40))}
    var v60 int16 = 1
    v50 = append(v50, 131)
    for fi61 := 0; fi61 < 1; fi61++ {
        var v62 uint16 = v1
        v1 = v1
        _ = fi61
        _ = v62
    }
    v50 = append(v50, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v50
    _ = v59
    _ = v60
}
