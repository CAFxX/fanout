package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int8; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint32; F2 bool }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS8 struct { F0 uint8; F1 int }
type GenS9 struct { F0 uint; F1 string; F2 uint64 }
func genFn0() int16 {
    var v6 uint8 = 128
    v6 -= v6
    _ = v6
    return 32767
}
func genFn1(p0 bool, p1 int64) int8 {
    p1 -= p1
    if p0 {
        var v7 float32 = 0.1
        v8 := []int{42}
        v9 := GenS1{F0: (0 != 42)}
        _ = v7
        _ = v8
        _ = v9
    } else {
        if true {
            for fi10 := 0; fi10 < 9; fi10++ {
                switch "a" {
                case "\n\t":
                    var v11 uint8 = 218
                    _ = v11
                }
                v12 := GenS2{F0: 255, F1: 42, F2: fi10}
                _ = fi10
                _ = v12
            }
            p0 = p0
        } else {
            v13 := []int{32768, 32767, 9223372036854775807, 42}
            _ = v13
        }
    }
    _ = p0
    _ = p1
    return 35
}

func main() {
    v1 := []string{"", "foo bar", "", "a"}
    var v2 float64 = (((-2.25 - 1e10) / (((3.14159 + (3.14159 - 3.14159)) / 0.1) - ((3.14159 - (1e10 * 0.0)) + 1.5))) * ((((1e10 * 3.14159) - 3.14159) * 1.5) / (float64((1.5 / (1e10 * -2.25))) * (1.5 / 1.5))))
    var v3 uint16 = 2
    v4 := make(map[uint32]bool)
    var v5 int = -1
    v3 -= 979
    var v14 int32 = 174
    var v15 uint = 18446744073709551615
    switch v5 {
    case -769:
        fmt.Println(v14)
    default:
        for fi16 := 0; fi16 < 11; fi16++ {
            v15 += v15
            if ((3.14159 / -2.25) >= float32(1.5)) {
                var v17 int64 = 42
                if (int8(v15) >= int8(v5)) {
                    _ = len(v4)
                    if (2 >= 59) {
                        v3 = 42
                        _ = len(v1)
                        var v18 int16 = -158
                        _ = v18
                    }
                    if true {
                        var v19 uint = 976
                        v20 := GenS3{F0: v17}
                        v21 := GenS1{F0: ((v3 >> ((v19) & 63)) < 255)}
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                }
                _ = v17
            } else {
                for fi22 := 0; fi22 < 15; fi22++ {
                    v15 = (v15 / (((^(v15))) | 1))
                    _ = fi22
                }
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            }
            switch fi16 {
            case 1:
                var v23 uint = 1
                _ = v23
            case 128:
                var v24 float64 = (0.0 * v2)
                _ = v24
            case 127:
                for fi25 := 0; fi25 < 10; fi25++ {
                    if false {
                        v14 = (v14 | v14)
                        switch (v5 % ((1) | 1)) {
                        case -2147483648:
                            var v26 uint = v15
                            _ = v26
                        case 42:
                            switch "hello" {
                            case "hello":
                                if false {
                                    v27 := GenS1{F0: (int8(v3) <= -121)}
                                    v4[240] = (true && (v3 <= 311))
                                    _ = v27
                                }
                            case "":
                                _ = len(v4)
                            }
                        case 32768:
                            if (false && true) {
                                var v28 float64 = (float64((float32(float64(0.1)) + ((((0.1 * 3.14159) / float32(float64((3.14159 + 3.14159)))) / 0.1) + 3.14159))) + (((3.14159 + float64(-2.25)) - ((((1e10 - 0.0) - 1.5) / 1.5) + v2)) + 0.0))
                                v29 := GenS4{F0: uint(v3), F1: 0, F2: (float64(1.5) == float64(-2.25))}
                                _ = v28
                                _ = v29
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                                v30 := GenS1{F0: ((0.1 <= v2) || (1 > (-(v14))))}
                                v3 += v3
                                _ = v30
                            }
                        default:
                            v31 := GenS2{F0: 748, F1: 1, F2: fi16}
                            _ = v31
                        }
                    }
                    _ = fi25
                }
            default:
                if (false || false) {
                    v32 := GenS1{F0: false}
                    if (false && (!true)) {
                        if true {
                            var v33 uint16 = v3
                            var v34 string = "\n\t"
                            _ = v33
                            _ = v34
                        }
                    }
                    _ = v32
                }
            }
            _ = fi16
        }
    }
    v35 := GenS4{F0: v15, F1: 2147483647, F2: false}
    var v36 int8 = -1
    var v37 int32 = v14
    if true {
        _ = len(v4)
        _ = len(v4)
    }
    v38 := v4[2147483647]
    for fi39 := 0; fi39 < 4; fi39++ {
        if len(v1) > 0 {
            v1[0] = ""
        }
        _ = fi39
    }
    v3 -= v3
    v3 += v3
    if true {
        _ = len(v1)
        v38 = v38
        switch "\n\t" {
        case "hello":
            if (91 > 128) {
                if false {
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    var v40 int64 = 57
                    v3 = 255
                    _ = v40
                }
            } else {
                v1 = append(v1, "a")
                var v41 int8 = (^(v36))
                for fi42 := 0; fi42 < 8; fi42++ {
                    var v43 bool = (v5 > v5)
                    _ = fi42
                    _ = v43
                }
                _ = v41
            }
        case "":
            var v44 uint64 = 568
            _ = v44
        }
    } else {
        var v45 int = (v5 ^ int(v15))
        v46 := GenS1{F0: v38}
        switch "a" {
        case "a":
            v47 := GenS5{F0: 824, F1: 42}
            _ = v47
        case "":
            v5 -= v5
        }
        _ = v45
        _ = v46
    }
    if true {
        v3 -= v3
    } else {
        if (((!(true || (v14 <= v14))) || true) || (42 > (v15 >> (((^(v15))) & 63)))) {
            var v48 uint = (^(v15))
            _ = v48
        } else {
            v49 := GenS6{F0: (v36 >> ((v15) & 63))}
            switch int(v37) {
            case 534:
                if v38 {
                    _ = len(v4)
                    var v50 float32 = 1.5
                    v51 := GenS7{F0: (v3 * v3), F1: 4294967295, F2: 0}
                    _ = v50
                    _ = v51
                }
            case 309:
                fmt.Println(v2, v36, v37)
            case 128:
                fmt.Println(v2, v37, v3)
            }
            if (("\n\t" > "a") && ("a" < "")) {
                for fi52 := 0; fi52 < 7; fi52++ {
                    var v53 float32 = (((0.1 / -2.25) / (1.5 * (1e10 - 3.14159))) / 1.0)
                    var v54 float64 = v2
                    fmt.Println(v5, v37)
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                if true {
                    for fi55 := 0; fi55 < 10; fi55++ {
                        if v38 {
                            for fi56 := 0; fi56 < 3; fi56++ {
                                var v57 int = (v5 / ((fi56) | 1))
                                switch -231 {
                                case 116:
                                    var v58 float32 = -2.25
                                    _ = v58
                                case 128:
                                    delete(v4, uint32(v3))
                                case 700:
                                    v59 := GenS8{F0: 127, F1: 0}
                                    _ = v59
                                }
                                var v60 uint64 = uint64(v3)
                                if false {
                                    v61 := GenS9{F0: 2, F1: "foo bar", F2: 0}
                                    if (float32(v2) > (float32(3.14159) + ((float32(v2) - float32((0.1 * 1.5))) * ((1e10 * 1.5) / 0.1)))) {
                                        var v62 string = "\n\t"
                                        _ = v62
                                    }
                                    _ = v61
                                }
                                _ = fi56
                                _ = v57
                                _ = v60
                            }
                            var v63 uint32 = 0
                            _ = v63
                        } else {
                            fmt.Println(v5, v14)
                            v38 = (((false && v38) && v38) || (!(v37 < (v14 ^ (-(v37))))))
                            v15 -= v15
                        }
                        fmt.Println(fi55, v14)
                        _ = fi55
                    }
                    v14 = (v14 * 0)
                    for fi64 := 0; fi64 < 2; fi64++ {
                        v5 = v5
                        var v65 float64 = v2
                        _ = fi64
                        _ = v65
                    }
                } else {
                    var v66 uint64 = uint64(v15)
                    var v67 float32 = ((1.5 - ((0.0 + float32((v2 * -2.25))) / 1e10)) + (((((0.0 * 0.1) / 1e10) + (3.14159 / 0.1)) * float32(v2)) / -2.25))
                    v15 += 65535
                    _ = v66
                    _ = v67
                }
                fmt.Println(v5)
            } else {
                var v68 uint8 = 95
                _ = v68
            }
            _ = v49
        }
    }
    for fi69 := 0; fi69 < 8; fi69++ {
        var v70 int16 = 1000
        switch "hello" {
        case "a":
            v71 := v4[4294967295]
            _ = v71
        case "foo bar":
            for fi72 := 0; fi72 < 19; fi72++ {
                switch 127 {
                case 792:
                    if v38 {
                        switch (v5 | 901) {
                        case -2147483648:
                            v36 -= 100
                            fmt.Println(v14)
                            switch "foo bar" {
                            case "hello":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        }
                        v73 := GenS9{F0: uint(v3), F1: "a", F2: 2}
                        v1 = append(v1, "hello")
                        _ = v73
                    }
                default:
                    _ = len(v1)
                }
                v38 = (true && v38)
                _ = fi72
            }
        case "hello":
            var v74 uint16 = (v3 & v3)
            _ = v74
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi69
        _ = v70
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
