package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 string; F2 int8 }
type GenS4 struct { F0 bool }
func genFn0(p0 string, p1 uint32, p2 string) uint16 {
    v5 := GenS1{F0: 42}
    p1 = (p1 & p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 679
}

func main() {
    var v1 int16 = 0
    v2 := []string{"foo bar", "\n\t", (("a" + "\n\t") + "hello"), "foo bar"}
    var v3 float32 = (float32(1.5) - (3.14159 + float32(1.5)))
    var v4 bool = (2147483647 != 2147483648)
    var v6 int64 = 9223372036854775807
    var v7 float64 = -2.25
    fmt.Println(v1, v3, v3)
    fmt.Println(v7)
    fmt.Println(v3)
    switch 127 {
    case 9223372036854775807:
        v8 := GenS2{F0: -1}
        _ = v8
    case -942:
        v2 = append(v2, (("a" + (("a" + "\n\t") + ("" + "foo bar"))) + "foo bar"))
    default:
        for fi9 := 0; fi9 < 15; fi9++ {
            if v4 {
                var v10 int32 = 680
                var v11 bool = ((float64(v3) != (v7 * float64(-2.25))) && false)
                _ = v10
                _ = v11
            }
            _ = fi9
        }
    }
    v4 = true
    var v12 string = ("" + "")
    var v13 uint8 = 127
    if (177 >= (v13 % (((v13 * (v13 ^ (v13 * 28)))) | 1))) {
        if len(v2) > 0 {
            v2[0] = v12
        }
        var v14 float64 = v7
        v15 := GenS1{F0: (v13 >> ((586) & 63))}
        _ = v14
        _ = v15
    } else {
        var v16 int8 = 127
        v13 += 57
        _ = v16
    }
    var v17 string = ""
    v4 = (v4 && true)
    v1 += (v1 - v1)
    v7 += v7
    _ = len(v2)
    v6 += (v6 * (^(v6)))
    if (2147483647 <= 127) {
        if (611 < 1) {
            for fi18 := 0; fi18 < 4; fi18++ {
                v17 = "hello"
                fmt.Println(v13, v17, v13)
                for fi19 := 0; fi19 < 1; fi19++ {
                    switch fi19 {
                    case 4294967295:
                        v13 += v13
                    case -9223372036854775808:
                        switch (fi19 * fi19) {
                        case -211:
                            if (0 > 0) {
                                v3 = v3
                                v20 := GenS1{F0: (^(v13))}
                                _ = v20
                            } else {
                                switch fi18 {
                                case -46:
                                    v21 := GenS1{F0: (^(v13))}
                                    _ = v21
                                default:
                                    v7 = 0.0
                                }
                                var v22 int32 = -1
                                if (127 > fi18) {
                                    if v4 {
                                        v4 = (((v13 * 128) != 127) && v4)
                                        fmt.Println(v17, v17, v4)
                                        if v4 {
                                            if (!(true && false)) {
                                                switch 1 {
                                                case 2147483647:
                                                    v3 += 1e10
                                                default:
                                                    _ = len(v2)
                                                }
                                            } else {
                                                var v23 int16 = 0
                                                if true {
                                                    if (!(v22 <= v22)) {
                                                        v2 = append(v2, v17)
                                                    } else {
                                                        v24 := GenS1{F0: v13}
                                                        _ = v24
                                                    }
                                                    v25 := GenS3{F0: fi18, F1: ((v17 + v12) + v12), F2: int8(v22)}
                                                    if ((^(fi18)) == fi18) {
                                                        var v26 float32 = float32(v7)
                                                        var v27 int16 = v23
                                                        _ = v26
                                                        _ = v27
                                                    }
                                                    var v28 float32 = 1.5
                                                    _ = v25
                                                    _ = v28
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                if v4 {
                                                    var v29 uint64 = 4294967295
                                                    if false {
                                                        var v30 uint32 = 65535
                                                        fmt.Println(v17, v23, v3)
                                                        if (true && v4) {
                                                            v4 = false
                                                            var v31 uint8 = v13
                                                            switch (fi19 << ((42) & 63)) {
                                                            case 32767:
                                                                if (v4 || true) {
                                                                    var v32 string = ((v17 + v17) + "hello")
                                                                    _ = v32
                                                                } else {
                                                                    var v33 float64 = float64(0.0)
                                                                    _ = v33
                                                                }
                                                                var v34 int8 = 127
                                                                _ = v34
                                                            case -118:
                                                                var v35 uint16 = genFn0(v12, (v30 * v30), ("\n\t" + "hello"))
                                                                _ = v35
                                                            default:
                                                                if v4 {
                                                                    v6 = int64(v30)
                                                                    v1 = (v1 | v1)
                                                                    v23 = v23
                                                                    fmt.Println(v17, v6)
                                                                }
                                                            }
                                                            _ = v31
                                                        } else {
                                                            if ((int64(v13) == v6) && (false && v4)) {
                                                                v2 = append(v2, v12)
                                                            } else {
                                                                var v36 int = 0
                                                                _ = v36
                                                            }
                                                        }
                                                        fmt.Println(v4)
                                                        _ = v30
                                                    } else {
                                                        switch ("foo bar" + ((v12 + "\n\t") + "\n\t")) {
                                                        case "hello":
                                                            if (true && false) {
                                                                fmt.Println(v3, v23)
                                                                var v37 uint16 = 65535
                                                                _ = v37
                                                            } else {
                                                                var v38 uint64 = (v29 * (v29 * (v29 + (v29 | 4294967295))))
                                                                v22 = (v22 << ((2) & 63))
                                                                _ = v38
                                                            }
                                                        case "foo bar":
                                                            v22 = int32(fi18)
                                                        case "a":
                                                            v39 := GenS3{F0: fi19, F1: "\n\t", F2: -68}
                                                            _ = v39
                                                        }
                                                        v40 := GenS4{F0: v4}
                                                        v12 = ((v12 + "a") + "")
                                                        _ = v40
                                                    }
                                                    _ = v29
                                                }
                                                _ = v23
                                            }
                                        }
                                    }
                                    fmt.Println(v13, v3, v17)
                                    fmt.Println(v17)
                                    v41 := GenS3{F0: fi18, F1: v12, F2: 99}
                                    _ = v41
                                }
                                _ = v22
                            }
                        case 9223372036854775807:
                            _ = len(v2)
                        case 127:
                            var v42 uint32 = 49
                            _ = v42
                        default:
                            var v43 uint8 = (v13 ^ v13)
                            _ = v43
                        }
                    }
                    var v44 uint = 659
                    if (false || (true || true)) {
                        v6 += v6
                    }
                    v44 -= v44
                    if (14 >= (v13 & (^(v13)))) {
                        var v45 uint16 = 65535
                        v46 := GenS1{F0: v13}
                        _ = v45
                        _ = v46
                    }
                    var v47 int = fi19
                    _ = fi19
                    _ = v44
                    _ = v47
                }
                _ = fi18
            }
        } else {
            var v48 bool = ((float64(float32(3.14159)) / -2.25) < v7)
            v6 += 319
            _ = v48
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
