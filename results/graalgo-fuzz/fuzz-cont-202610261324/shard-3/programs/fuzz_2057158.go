package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 string; F2 string }
type GenS4 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = (3.14159 + 0.1)
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v5 uint8 = 42
    var v6 uint16 = 117
    switch "" {
    case "foo bar":
        var v7 int = -1
        _ = v7
    case "a":
        switch "" {
        case "foo bar":
            if true {
                var v8 float64 = 1.5
                switch "a" {
                case "hello":
                    var v9 string = "foo bar"
                    _ = v9
                }
                _ = v8
            } else {
                var v10 int64 = 9223372036854775807
                _ = v10
            }
        case "hello":
            switch 948 {
            case 1:
                var v11 int64 = -1
                _ = v11
            case 817:
                v2 = v2
            case 127:
                for fi12 := 0; fi12 < 10; fi12++ {
                    fmt.Println(v6, fi12)
                    _ = fi12
                }
            }
        default:
            var v13 int = 128
            _ = v13
        }
    case "\n\t":
        fmt.Println(v5, v5)
    }
    v5 -= v5
    v5 -= (v5 | v5)
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := v4[fi14]
        var v16 bool = false
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        fmt.Println(v6, fi17)
        v18 := []int{fi17, (^(fi17))}
        for fi19 := 0; fi19 < 1; fi19++ {
            _ = len(v3)
            var v20 uint32 = 2147483647
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        v6 += (v6 * v6)
        switch ("hello" + "hello") {
        case "":
            fmt.Println(v2)
        case "foo bar":
            fmt.Println(fi21, fi21, v2)
        case "hello":
            if true {
                var v22 string = "a"
                _ = v22
            } else {
                v5 += v5
            }
        default:
            v23 := GenS1{F0: -1, F1: 1}
            _ = v23
        }
        v24 := GenS2{F0: 0}
        _ = fi21
        _ = v24
    }
    v25 := []int{32767}
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 float32 = -2.25
        v6 = 65535
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        for fi29 := 0; fi29 < 1; fi29++ {
            if (true || (true || true)) {
                var v30 float32 = float32((1e10 + ((1.5 + (1.5 - (0.1 * (((float64(v2) - 3.14159) * 1e10) / 1.0)))) * 3.14159)))
                _ = v30
            }
            var v31 int64 = 0
            v5 = v5
            _ = fi29
            _ = v31
        }
        var v32 uint32 = 0
        for fi33 := 0; fi33 < 17; fi33++ {
            v34 := GenS1{F0: 32767, F1: -1}
            var v35 uint64 = 761
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi28
        _ = v32
    }
    if (int(v6) >= 2147483647) {
        _ = len(v4)
        var v36 uint64 = 994
        _ = v36
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        switch (fi37 % ((fi37) | 1)) {
        case -671:
            _ = len(v3)
        case -2147483648:
            v38 := GenS1{F0: -626, F1: -15}
            _ = v38
        case 127:
            var v39 float64 = 1.5
            _ = v39
        default:
            v2 -= v2
        }
        _ = fi37
    }
    switch "foo bar" {
    case "":
        for fi40 := 0; fi40 < 7; fi40++ {
            if len(v25) > 0 {
                v25[0] = fi40
            }
            _ = fi40
        }
    case "\n\t":
        for fi41 := 0; fi41 < 10; fi41++ {
            v5 += (v5 >> ((548) & 63))
            _ = len(v25)
            _ = fi41
        }
    case "a":
        v42 := GenS3{F0: 607, F1: "\n\t", F2: "hello"}
        _ = v42
    default:
        if ((65535 < 2) && (!false)) {
            v3[-616] = 571
            if false {
                if len(v25) > 0 {
                    v25[0] = 128
                }
                _ = len(v25)
                fmt.Println(v2)
            } else {
                v5 = 127
                v25 = append(v25, 128)
                for fi43 := 0; fi43 < 7; fi43++ {
                    var v44 uint64 = 1
                    if true {
                        var v45 uint64 = (v44 - (^(v44)))
                        v46 := GenS3{F0: 0, F1: "\n\t", F2: ("hello" + ("a" + "foo bar"))}
                        _ = v45
                        _ = v46
                    } else {
                        var v47 int16 = 23
                        var v48 uint8 = v5
                        for fi49 := 0; fi49 < 17; fi49++ {
                            v47 += (v47 | v47)
                            _ = fi49
                        }
                        _ = v47
                        _ = v48
                    }
                    _ = fi43
                    _ = v44
                }
            }
            var v50 uint8 = 127
            _ = v50
        }
    }
    fmt.Println(v6, v5, v2)
    switch 0 {
    case -610:
        fmt.Println(v6, v5, v5)
    case 463:
        if false {
            v51 := GenS4{F0: 0, F1: 18446744073709551615}
            v5 += 233
            _ = v51
        } else {
            v6 = uint16(v5)
        }
    case 9223372036854775807:
        if (18446744073709551615 != 855) {
            switch -1 {
            case -347:
                v2 = (-2.25 - (-2.25 / (1.5 + (v2 / ((v2 + 1.5) - v2)))))
            case 42:
                v52 := GenS4{F0: 0, F1: 255}
                _ = v52
            case 1:
                var v53 uint = 123
                _ = v53
            }
            v5 = 127
            v2 = v2
        } else {
            for fi54 := 0; fi54 < 4; fi54++ {
                for fi55 := 0; fi55 < 16; fi55++ {
                    var v56 int = fi54
                    var v57 string = ("a" + "hello")
                    v25 = append(v25, (v56 - 127))
                    _ = fi55
                    _ = v56
                    _ = v57
                }
                _ = fi54
            }
        }
    default:
        _ = len(v25)
    }
    if ((42 == 16) && true) {
        var v58 string = (("foo bar" + (("foo bar" + "") + "")) + "\n\t")
        _ = v58
    }
    var v59 bool = false
    if v59 {
        for fi60 := 0; fi60 < 6; fi60++ {
            if (1 != 323) {
                v5 = v5
            }
            v6 = (v6 * 32767)
            v1[("" + "a")] = 0
            _ = fi60
        }
    } else {
        v61 := GenS2{F0: 680}
        _ = v61
    }
    v59 = (524 <= 831)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
