package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 bool; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 int8; F1 bool; F2 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint16 }
func genFn0(p0 int32) int64 {
    v7 := GenS1{F0: 128}
    v8 := []int{32767, -9223372036854775808}
    if false {
        var v9 uint32 = 2147483647
        v9 += 500
        _ = v9
    } else {
        if len(v8) > 0 {
            v8[0] = 912
        }
    }
    if (("\n\t" + "") <= "foo bar") {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        switch ("hello" + "") {
        case "a":
            if len(v8) > 0 {
                v8[0] = -349
            }
        case "hello":
            p0 = 276
        case "\n\t":
            if false {
                fmt.Println(p0)
            }
        default:
            if ("foo bar" <= "a") {
                v10 := GenS1{F0: 127}
                p0 = p0
                fmt.Println(p0)
                _ = v10
            } else {
                v8 = append(v8, -1)
                for fi11 := 0; fi11 < 13; fi11++ {
                    if (true && false) {
                        if (true || (0 == 0)) {
                            for fi12 := 0; fi12 < 16; fi12++ {
                                if (!true) {
                                    fmt.Println(fi12, fi11)
                                    p0 += (^(p0))
                                    if (0.1 > ((1e10 * 0.0) * 3.14159)) {
                                        v13 := GenS1{F0: 13}
                                        v14 := GenS2{F0: 2}
                                        if false {
                                            v15 := GenS1{F0: 1}
                                            v16 := GenS3{F0: 50, F1: false, F2: false}
                                            fmt.Println(fi11, fi12)
                                            _ = v15
                                            _ = v16
                                        } else {
                                            if false {
                                                var v17 uint8 = 177
                                                fmt.Println(v17, v17)
                                                if (1 != 0) {
                                                    p0 = -831
                                                }
                                                var v18 uint8 = (v17 % ((v17) | 1))
                                                _ = v17
                                                _ = v18
                                            }
                                        }
                                        _ = v13
                                        _ = v14
                                    } else {
                                        v19 := GenS2{F0: 753}
                                        var v20 string = ""
                                        var v21 float32 = (3.14159 - -2.25)
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                    }
                                    p0 = p0
                                } else {
                                    fmt.Println(fi11, p0, fi11)
                                    switch "a" {
                                    case "\n\t":
                                        if false {
                                            if false {
                                                p0 = p0
                                                v22 := GenS4{F0: 255}
                                                _ = v22
                                            }
                                            p0 += (p0 << ((65535) & 63))
                                            var v23 uint32 = 487
                                            _ = v23
                                        }
                                    case "":
                                        if len(v8) > 0 {
                                            v8[0] = (fi11 >> ((0) & 63))
                                        }
                                    default:
                                        var v24 int8 = 53
                                        _ = v24
                                    }
                                    p0 -= (^(p0))
                                }
                                if (4294967295 > 1) {
                                    switch (fi11 ^ 4294967295) {
                                    case 2147483647:
                                        var v25 float32 = 1e10
                                        _ = v25
                                    case 668:
                                        if ((!false) && false) {
                                            var v26 int64 = 253
                                            var v27 uint64 = 885
                                            _ = v26
                                            _ = v27
                                        }
                                    }
                                    p0 = p0
                                }
                                _ = fi12
                            }
                        }
                    } else {
                        v28 := GenS3{F0: 91, F1: true, F2: false}
                        var v29 uint8 = 127
                        _ = v28
                        _ = v29
                    }
                    p0 += p0
                    _ = fi11
                }
                v30 := GenS5{F0: "", F1: "hello"}
                _ = v30
            }
        }
        v31 := GenS5{F0: "\n\t", F1: ("hello" + "\n\t")}
        _ = v31
    } else {
        switch 345 {
        case 32767:
            v32 := GenS3{F0: 1, F1: true, F2: (2 >= 2)}
            _ = v32
        case -846:
            v33 := GenS6{F0: -1, F1: (1 >= -927), F2: 4294967295}
            _ = v33
        case -153:
            _ = len(v8)
        default:
            if len(v8) > 0 {
                v8[0] = -9223372036854775808
            }
        }
    }
    p0 += (p0 * -117)
    _ = p0
    _ = v7
    _ = v8
    return 387
}

func main() {
    var v1 float32 = (0.0 / ((float32(-2.25) - 1.5) / 1.5))
    var v2 uint = 756
    var v3 int16 = -1
    var v4 uint8 = 127
    var v5 uint64 = uint64(v4)
    var v6 bool = (796 == 255)
    v34 := GenS7{F0: (266 > uint16(v5))}
    switch 9223372036854775807 {
    case 9223372036854775807:
        v35 := GenS8{F0: genFn0(42), F1: false}
        _ = v35
    default:
        if (1000 < (v3 * v3)) {
            v36 := GenS8{F0: 595, F1: false}
            fmt.Println(v3)
            _ = v36
        }
    }
    v37 := GenS9{F0: 94}
    v4 -= v4
    v38 := []int{-249, -2147483648, 0}
    var v39 uint = v2
    v6 = (1e10 > 1e10)
    if (9223372036854775807 > int64(v4)) {
        var v40 bool = (false || v6)
        _ = v40
    } else {
        if ((1 != 0) && v6) {
            if len(v38) >= 2 {
                v38 = v38[1:len(v38)]
            }
            for fi41 := 0; fi41 < 1; fi41++ {
                switch int(v4) {
                case 42:
                    var v42 uint32 = 2147483648
                    _ = v42
                case 0:
                    for fi43 := 0; fi43 < 4; fi43++ {
                        v1 = ((1e10 * (float32(0.0) - 1e10)) / 1.5)
                        fmt.Println(v4, v4, v3)
                        if true {
                            var v44 float32 = v1
                            _ = v44
                        }
                        _ = fi43
                    }
                default:
                    if v6 {
                        fmt.Println(v4)
                        for fi45 := 0; fi45 < 18; fi45++ {
                            var v46 uint64 = (v5 - (v5 * (v5 * (v5 ^ v5))))
                            _ = fi45
                            _ = v46
                        }
                    } else {
                        v47 := GenS4{F0: 32767}
                        for fi48 := 0; fi48 < 13; fi48++ {
                            var v49 int8 = 1
                            var v50 uint32 = 42
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                        _ = v47
                    }
                }
                _ = fi41
            }
        }
        if (!false) {
            for fi51 := 0; fi51 < 3; fi51++ {
                if (v4 >= (^(v4))) {
                    v52 := GenS10{F0: 255}
                    _ = v52
                }
                _ = fi51
            }
        } else {
            var v53 string = "\n\t"
            v38 = append(v38, 2147483647)
            switch 204 {
            case 1:
                v4 = v4
            default:
                fmt.Println(v53, v3, v39)
            }
            _ = v53
        }
    }
    v4 = 127
    v38 = append(v38, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
