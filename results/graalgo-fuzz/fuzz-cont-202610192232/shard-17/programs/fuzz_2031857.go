package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int64; F1 uint32 }
type GenS9 struct { F0 uint64 }
func genFn0(p0 int, p1 int32) uint32 {
    var v7 uint64 = 4294967295
    v8 := []int{2147483647, p0}
    var v9 bool = true
    if (42 != 25) {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 uint32 = 2147483647
            fmt.Println(fi10)
            var v12 uint64 = v7
            _ = fi10
            _ = v11
            _ = v12
        }
    } else {
        var v13 float64 = ((1e10 - ((3.14159 / 1.0) + 3.14159)) * 1.5)
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 1000
    var v3 uint64 = 2
    var v4 float32 = 0.1
    var v5 int = 4294967295
    var v6 int16 = -476
    switch (("hello" + "a") + (("foo bar" + "\n\t") + "")) {
    case "hello":
        v2 = (v2 * v6)
    default:
        for fi14 := 0; fi14 < 8; fi14++ {
            v15 := []int{(fi14 << ((0) & 63))}
            fmt.Println(v3, v2, v5)
            _ = fi14
            _ = v15
        }
    }
    switch v5 {
    case -200:
        fmt.Println(v4, v2, v3)
    }
    if false {
        for fi16 := 0; fi16 < 6; fi16++ {
            v2 += 0
            v17 := GenS1{F0: "a", F1: ""}
            for fi18 := 0; fi18 < 5; fi18++ {
                v19 := GenS1{F0: "a", F1: ("a" + "foo bar")}
                if ((float64(v4) + (((3.14159 / -2.25) * float64((v4 + 1e10))) + 0.0)) != ((-2.25 - float64(0.0)) / -2.25)) {
                    delete(v1, 1)
                    v4 = v4
                } else {
                    if (-247 > 1) {
                        v20 := GenS2{F0: 42}
                        _ = v20
                    } else {
                        if true {
                            v21 := []int{}
                            var v22 float32 = 3.14159
                            _ = v21
                            _ = v22
                        }
                        var v23 int8 = -1
                        var v24 uint16 = 571
                        _ = v23
                        _ = v24
                    }
                }
                switch int(v3) {
                case -2147483648:
                    var v25 int16 = v6
                    _ = v25
                case -734:
                    var v26 uint8 = 222
                    _ = v26
                case 32767:
                    v27 := []int{}
                    _ = v27
                default:
                    var v28 int64 = -9223372036854775808
                    _ = v28
                }
                _ = fi18
                _ = v19
            }
            _ = fi16
            _ = v17
        }
        v29 := GenS3{F0: genFn0(v5, 42), F1: -1}
        var v30 uint8 = 42
        _ = v29
        _ = v30
    } else {
        v5 += 690
        v31 := GenS1{F0: "hello", F1: "a"}
        _ = v31
    }
    if false {
        var v32 uint = 40
        fmt.Println(v6, v32)
        _ = v32
    }
    if false {
        var v33 int64 = -1
        _ = v33
    }
    var v34 string = ""
    switch (-(v5)) {
    case 50:
        v4 -= 1.5
    case 127:
        fmt.Println(v5, v6, v6)
    default:
        if (!false) {
            switch v34 {
            case "\n\t":
                for fi35 := 0; fi35 < 20; fi35++ {
                    if (!((!false) || true)) {
                        v36 := GenS4{F0: 0, F1: (v3 ^ v3)}
                        v37 := GenS5{F0: v3, F1: false}
                        _ = v36
                        _ = v37
                    }
                    if (0.0 < ((((1.5 / 1.0) * 3.14159) * 1e10) * ((0.1 - 0.1) + -2.25))) {
                        switch "hello" {
                        case "a":
                            var v38 bool = (v5 >= 999)
                            _ = v38
                        case "hello":
                            delete(v1, 2147483647)
                        }
                    } else {
                        for fi39 := 0; fi39 < 16; fi39++ {
                            v40 := GenS4{F0: 4294967295, F1: v3}
                            v34 = v34
                            _ = fi39
                            _ = v40
                        }
                        var v41 bool = true
                        _ = v41
                    }
                    var v42 int64 = -9223372036854775808
                    _ = fi35
                    _ = v42
                }
            }
            for fi43 := 0; fi43 < 19; fi43++ {
                var v44 uint = 583
                var v45 uint32 = 42
                _ = fi43
                _ = v44
                _ = v45
            }
        } else {
            var v46 int32 = 0
            v5 -= (v5 >> ((62) & 63))
            fmt.Println(v4)
            _ = v46
        }
    }
    if ((100 != 100) || (false && false)) {
        v47 := GenS3{F0: genFn0(v5, 986), F1: 42}
        var v48 uint8 = 0
        for fi49 := 0; fi49 < 10; fi49++ {
            v50 := []int{(fi49 >> ((739) & 63)), v5, -844, v5}
            for fi51 := 0; fi51 < 10; fi51++ {
                v48 += 42
                _ = fi51
            }
            v6 = -579
            _ = fi49
            _ = v50
        }
        _ = v47
        _ = v48
    } else {
        if false {
            v5 = v5
            var v52 uint = 0
            _ = v52
        }
        switch 444 {
        case 0:
            v3 = v3
        case -298:
            var v53 int32 = 2147483647
            _ = v53
        default:
            v54 := v1[249]
            _ = v54
        }
    }
    if false {
        for fi55 := 0; fi55 < 15; fi55++ {
            v1[42] = v5
            _ = fi55
        }
        var v56 int64 = 799
        _ = v56
    } else {
        if true {
            if (!(true && ((true || (0 < 127)) && (1 == genFn0(v5, 51))))) {
                for fi57 := 0; fi57 < 18; fi57++ {
                    switch (v34 + "hello") {
                    case "":
                        var v58 bool = true
                        _ = v58
                    case "foo bar":
                        v59 := GenS6{F0: 366, F1: 42}
                        _ = v59
                    default:
                        fmt.Println(v6, v34, v2)
                    }
                    _ = fi57
                }
                v60 := GenS2{F0: 0}
                var v61 int8 = 0
                _ = v60
                _ = v61
            }
        } else {
            for fi62 := 0; fi62 < 14; fi62++ {
                for fi63 := 0; fi63 < 18; fi63++ {
                    v64 := v1[2147483647]
                    var v65 bool = (v3 >= v3)
                    var v66 uint = 42
                    fmt.Println(v3)
                    if v65 {
                        switch "\n\t" {
                        case "\n\t":
                            var v67 int32 = 606
                            var v68 float64 = 1e10
                            v65 = v65
                            if true {
                                v69 := []int{(v5 / ((v64) | 1)), v64, (v64 + fi63)}
                                switch v34 {
                                case "foo bar":
                                    var v70 uint32 = 0
                                    _ = v70
                                }
                                _ = len(v69)
                                _ = v69
                            }
                            var v71 uint32 = 65535
                            fmt.Println(v3)
                            _ = v67
                            _ = v68
                            _ = v71
                        case "a":
                            fmt.Println(v34, fi62, v3)
                            if (0.1 > 1e10) {
                                v5 = (v5 + (v64 + (v5 - (fi62 >> (((v66 * (v66 << (((v66 * (v66 & v66))) & 63)))) & 63)))))
                                switch fi62 {
                                case 610:
                                    v64 -= int(v66)
                                }
                                v4 = float32(1e10)
                            }
                            if ((0.1 <= (1e10 * 0.0)) || v65) {
                                fmt.Println(v3, fi62)
                            } else {
                                var v72 float32 = v4
                                var v73 int8 = -6
                                _ = v72
                                _ = v73
                            }
                        case "":
                            v74 := GenS7{F0: 71}
                            _ = v74
                        }
                        v2 = -32768
                    }
                    v3 += v3
                    _ = fi63
                    _ = v64
                    _ = v65
                    _ = v66
                }
                if ((v3 != (v3 / ((590) | 1))) || false) {
                    var v75 bool = false
                    _ = v75
                }
                if true {
                    v76 := GenS2{F0: -128}
                    if false {
                        switch (v34 + v34) {
                        case "hello":
                            v77 := GenS8{F0: 880, F1: 0}
                            _ = v77
                        case "":
                            if true {
                                var v78 int16 = 655
                                var v79 int16 = (v2 % ((v78) | 1))
                                _ = v78
                                _ = v79
                            }
                        default:
                            v4 -= float32(1.5)
                        }
                    } else {
                        v80 := GenS7{F0: -58}
                        _ = v80
                    }
                    _ = v76
                }
                _ = fi62
            }
            v4 -= 0.0
        }
        v3 -= v3
        var v81 int16 = 32767
        _ = v81
    }
    var v82 int64 = 181
    fmt.Println(v34)
    v34 = "a"
    for fi83 := 0; fi83 < 1; fi83++ {
        for fi84 := 0; fi84 < 17; fi84++ {
            v85 := GenS3{F0: genFn0((v5 << ((18446744073709551615) & 63)), -2147483648), F1: -1}
            v86 := GenS9{F0: (v3 & v3)}
            _ = fi84
            _ = v85
            _ = v86
        }
        var v87 int8 = 100
        _ = fi83
        _ = v87
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v82
}
