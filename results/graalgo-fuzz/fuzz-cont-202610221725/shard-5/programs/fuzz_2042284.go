package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint }
type GenS2 struct { F0 bool; F1 int32; F2 int16 }
type GenS3 struct { F0 bool; F1 bool; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 bool = (0 == -1)
    var v2 uint8 = 128
    var v3 float32 = (float32(float64((((-2.25 + (0.0 + 0.1)) - float32(0.1)) / 1.5))) + (1.5 / float32(0.1)))
    var v4 float32 = 1.5
    var v5 uint16 = 0
    switch "hello" {
    case "hello":
        var v6 bool = (!false)
        _ = v6
    case "a":
        var v7 int8 = 45
        _ = v7
    case "\n\t":
        switch 964 {
        case -468:
            v8 := GenS1{F0: "", F1: 2, F2: 326}
            _ = v8
        case 127:
            for fi9 := 0; fi9 < 19; fi9++ {
                v4 -= v3
                fmt.Println(v2, v1)
                if (254 != v2) {
                    v3 += (v4 * float32((float64(v4) - -2.25)))
                    var v10 uint = 845
                    for fi11 := 0; fi11 < 7; fi11++ {
                        var v12 int = 2147483647
                        var v13 int64 = 1
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    _ = v10
                } else {
                    if v1 {
                        var v14 uint8 = v2
                        _ = v14
                    }
                    v5 = (v5 ^ (v5 & v5))
                    var v15 float32 = (-2.25 / 3.14159)
                    _ = v15
                }
                _ = fi9
            }
        default:
            var v16 float32 = (v4 * 0.0)
            _ = v16
        }
    default:
        if (uint64(v5) != 4294967295) {
            var v17 uint8 = (v2 / ((127) | 1))
            switch "foo bar" {
            case "hello":
                v18 := GenS2{F0: (!(1 == -796)), F1: 1, F2: 551}
                _ = v18
            }
            _ = v17
        }
    }
    v19 := []int{}
    for fi20 := 0; fi20 < 3; fi20++ {
        switch ("a" + "a") {
        case "":
            if len(v19) > 0 {
                v19[0] = fi20
            }
        }
        _ = fi20
    }
    v1 = (2 != v2)
    for fi21 := 0; fi21 < 1; fi21++ {
        v4 -= (v3 * (float32(0.0) * v3))
        v4 -= (v4 * (v4 - -2.25))
        _ = fi21
    }
    if (v1 && v1) {
        var v22 int32 = -514
        var v23 int64 = 0
        _ = v22
        _ = v23
    } else {
        fmt.Println(v3)
        v3 = (float32(((-2.25 - 1e10) * 1e10)) / 0.1)
        for fi24 := 0; fi24 < 17; fi24++ {
            var v25 bool = (42 != v2)
            for fi26 := 0; fi26 < 14; fi26++ {
                var v27 uint64 = 2
                var v28 int16 = -847
                if v25 {
                    v5 = (^(v5))
                } else {
                    fmt.Println(v28, v4, v3)
                }
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi24
            _ = v25
        }
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        var v30 int16 = 1000
        for fi31 := 0; fi31 < 5; fi31++ {
            v4 = ((v3 - (v3 / 1e10)) - ((1e10 * v4) - ((((0.0 + 1e10) - (3.14159 + 1.5)) + (0.0 * 1.5)) + v4)))
            v32 := GenS1{F0: "foo bar", F1: 65535, F2: uint(v2)}
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 15; fi33++ {
            if ((float32(0.0) - v3) >= ((-2.25 / (0.0 * ((v4 / 1.0) + v3))) - v3)) {
                v34 := GenS3{F0: (fi33 > (fi29 + 688)), F1: (1.5 <= (3.14159 - 3.14159)), F2: -10}
                _ = v34
            } else {
                if true {
                    v2 = 2
                } else {
                    v2 = v2
                }
                v35 := GenS2{F0: v1, F1: 2147483647, F2: v30}
                switch (fi33 & fi33) {
                case -9223372036854775808:
                    v36 := GenS3{F0: ((-9223372036854775808 != 9223372036854775807) || v1), F1: (2147483647 < 2), F2: -60}
                    _ = v36
                }
                if v1 {
                    v5 = (v5 << ((4294967295) & 63))
                    var v37 bool = v1
                    v3 = v3
                    var v38 bool = (!false)
                    _ = v37
                    _ = v38
                } else {
                    fmt.Println(v30)
                }
                _ = v35
            }
            if true {
                fmt.Println(v1, v3, v3)
            } else {
                v30 = v30
            }
            switch fi29 {
            case 32767:
                var v39 bool = (true || (!v1))
                _ = v39
            case -731:
                switch "hello" {
                case "a":
                    var v40 float32 = v3
                    _ = v40
                case "foo bar":
                    if (!(-1 != 1)) {
                        fmt.Println(v5)
                        v41 := GenS3{F0: (4294967295 <= 2), F1: (((v30 == (v30 * v30)) && (4294967295 != 1)) || v1), F2: -30}
                        v19 = append(v19, 1)
                        _ = v41
                    }
                }
            default:
                var v42 uint = 571
                _ = v42
            }
            _ = fi33
        }
        _ = fi29
        _ = v30
    }
    switch -368 {
    case 0:
        var v43 uint = 4294967295
        _ = v43
    case 42:
        fmt.Println(v1, v1, v2)
    }
    if (v5 != 733) {
        var v44 uint8 = (v2 ^ (^(v2)))
        var v45 uint32 = 614
        v46 := GenS4{F0: v5, F1: (v45 - (v45 - (v45 / (((v45 ^ v45)) | 1))))}
        _ = v44
        _ = v45
        _ = v46
    }
    var v47 string = (("" + ("a" + "\n\t")) + "a")
    var v48 uint64 = 42
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
