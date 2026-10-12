package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int16; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS8 struct { F0 int; F1 int64 }
func genFn0() uint {
    v6 := []int{295}
    v7 := GenS1{F0: "hello"}
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        var v8 string = "a"
        _ = v8
    }
    var v9 uint32 = 42
    _ = v6
    _ = v7
    _ = v9
    return 286
}

func main() {
    var v1 int8 = 90
    var v2 int16 = -1
    v3 := make(map[uint32]int)
    var v4 uint64 = 4294967295
    v5 := []string{"", "a"}
    fmt.Println(v4, v2)
    if true {
        v10 := GenS2{F0: 2147483647, F1: (v2 << ((4294967295) & 63)), F2: -106}
        for fi11 := 0; fi11 < 11; fi11++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v4 = v4
            if false {
                for fi12 := 0; fi12 < 8; fi12++ {
                    switch fi11 {
                    case 42:
                        switch "hello" {
                        case "\n\t":
                            v13 := GenS3{F0: 65535, F1: 1}
                            _ = v13
                        default:
                            v14 := v3[2147483647]
                            _ = v14
                        }
                    default:
                        v4 += v4
                    }
                    var v15 bool = true
                    var v16 uint16 = 0
                    switch "hello" {
                    case "a":
                        if true {
                            var v17 uint8 = 42
                            _ = v17
                        }
                    }
                    _ = fi12
                    _ = v15
                    _ = v16
                }
                switch ("foo bar" + ("hello" + "foo bar")) {
                case "foo bar":
                    v4 += v4
                }
            } else {
                switch (fi11 + fi11) {
                case 260:
                    var v18 int16 = 1000
                    _ = v18
                case 32768:
                    v2 = v2
                case -1:
                    switch (-(fi11)) {
                    case 42:
                        v2 = (v2 << ((308) & 63))
                    case 1:
                        v3[688] = 32767
                    case -9223372036854775808:
                        var v19 int32 = 0
                        _ = v19
                    default:
                        if false {
                            if false {
                                if false {
                                    if ((fi11 / ((fi11) | 1)) != 32768) {
                                        v2 = (v2 / (((v2 ^ (v2 / ((v2) | 1)))) | 1))
                                    }
                                }
                            } else {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        } else {
                            fmt.Println(v4, v4, fi11)
                            var v20 bool = true
                            v20 = false
                            _ = v20
                        }
                    }
                }
            }
            _ = fi11
        }
        _ = v10
    }
    _ = len(v3)
    if len(v5) > 0 {
        v5[0] = ""
    }
    var v21 uint16 = 32767
    fmt.Println(v4)
    v22 := GenS1{F0: "foo bar"}
    var v23 bool = (1 >= (v21 / ((v21) | 1)))
    for fi24 := 0; fi24 < 19; fi24++ {
        if len(v5) > 0 {
            v5[0] = ("foo bar" + ("foo bar" + "foo bar"))
        }
        _ = fi24
    }
    if v23 {
        if (32768 != -1) {
            if v23 {
                switch "a" {
                case "foo bar":
                    v25 := GenS4{F0: 104}
                    _ = v25
                default:
                    v4 += v4
                }
                v26 := GenS4{F0: 127}
                v27 := GenS4{F0: 110}
                _ = v26
                _ = v27
            }
            var v28 uint8 = 1
            v28 = (v28 << ((261) & 63))
            _ = v28
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if v23 {
            switch 222 {
            case 127:
                for fi29 := 0; fi29 < 16; fi29++ {
                    if false {
                        switch "hello" {
                        case "hello":
                            v30 := GenS2{F0: 0, F1: v2, F2: 100}
                            _ = v30
                        default:
                            var v31 int8 = int8(v4)
                            _ = v31
                        }
                        if false {
                            var v32 uint16 = 32767
                            _ = v32
                        } else {
                            for fi33 := 0; fi33 < 9; fi33++ {
                                v34 := GenS1{F0: (("\n\t" + "foo bar") + "a")}
                                if true {
                                    var v35 uint8 = 42
                                    _ = v35
                                }
                                var v36 uint8 = 6
                                _ = fi33
                                _ = v34
                                _ = v36
                            }
                            var v37 float64 = -2.25
                            v1 += 127
                            _ = v37
                        }
                    }
                    _ = fi29
                }
            case 1:
                delete(v3, 2)
            case 32767:
                if true {
                    if v23 {
                        v23 = v23
                        if (!(!false)) {
                            v1 = v1
                            var v38 int64 = 1
                            switch ("hello" + "a") {
                            case "foo bar":
                                if false {
                                    v39 := GenS5{F0: false, F1: (v2 - int16(v1))}
                                    _ = v39
                                }
                            case "a":
                                fmt.Println(v23, v23, v21)
                            case "hello":
                                v1 = -128
                            default:
                                fmt.Println(v2, v2)
                            }
                            _ = v38
                        }
                    }
                    v40 := GenS6{F0: (true || (!v23))}
                    var v41 string = "\n\t"
                    _ = v40
                    _ = v41
                }
            default:
                v3[42] = 2147483647
            }
        }
    }
    for fi42 := 0; fi42 < 15; fi42++ {
        v43 := v3[2147483647]
        _ = fi42
        _ = v43
    }
    var v44 float32 = float32((-2.25 / (1.5 - -2.25)))
    fmt.Println(v21, v44)
    v45 := GenS7{F0: (v4 | v4), F1: 1, F2: 2}
    v23 = (v23 || ((-935 == -134) && v23))
    switch 127 {
    case 127:
        var v46 int8 = (^(v1))
        _ = v46
    case 9223372036854775807:
        for fi47 := 0; fi47 < 10; fi47++ {
            for fi48 := 0; fi48 < 12; fi48++ {
                var v49 int = (fi47 / ((2147483647) | 1))
                v2 += v2
                _ = fi48
                _ = v49
            }
            fmt.Println(fi47, fi47)
            v50 := GenS8{F0: (fi47 % ((42) | 1)), F1: 60}
            _ = fi47
            _ = v50
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}
