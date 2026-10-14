package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 bool; F1 int8; F2 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = -2.25
    var v3 int32 = 0
    var v4 string = ""
    var v5 int = -9223372036854775808
    v6 := GenS1{F0: 32767}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 float32 = -2.25
        _ = fi7
        _ = v8
    }
    v9 := []int{(v5 - (v5 % (((v5 & v5)) | 1))), (v5 >> ((18446744073709551615) & 63)), (v5 - (^(v5))), (v5 | (v5 % ((v5) | 1)))}
    v2 = (v2 * v2)
    if (!true) {
        var v10 uint8 = 255
        _ = v10
    } else {
        v5 = v5
        v11 := GenS2{F0: 9223372036854775807, F1: 1}
        if true {
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 int32 = v3
                fmt.Println(v5, v13)
                if false {
                    if (2147483648 != 65535) {
                        v3 = -1
                        var v14 int64 = 9223372036854775807
                        var v15 uint = 816
                        _ = v14
                        _ = v15
                    } else {
                        v3 += v3
                        var v16 float64 = float64(-2.25)
                        _ = v16
                    }
                    v3 += v13
                } else {
                    fmt.Println(v13, fi12)
                    v17 := GenS3{F0: (!(0 < 774)), F1: 0, F2: -174}
                    var v18 float64 = 0.0
                    _ = v17
                    _ = v18
                }
                _ = fi12
                _ = v13
            }
        }
        _ = v11
    }
    var v19 uint64 = 188
    if (!true) {
        v20 := GenS4{F0: -9223372036854775808}
        _ = v20
    }
    fmt.Println(v2)
    v2 -= ((v2 + 0.1) - (-2.25 + (0.0 * 1.5)))
    var v21 float32 = (3.14159 * 1.5)
    v2 = (v2 * float64((((v21 * 1.5) - v21) * 1.5)))
    if (1 != 100) {
        var v22 int32 = (v3 % ((v3) | 1))
        v22 = (v3 | (v22 - 1))
        fmt.Println(v3, v3, v4)
        _ = v22
    }
    if ((3.14159 + v21) >= v21) {
        v23 := GenS3{F0: false, F1: -128, F2: -1}
        if false {
            if (2 != 2147483648) {
                v3 = v3
            } else {
                if (v2 != v2) {
                    if false {
                        v24 := GenS5{F0: -1}
                        v5 = (v5 % ((v5) | 1))
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            var v26 int16 = -407
                            var v27 bool = (0 < int8(v5))
                            if false {
                                if v27 {
                                    delete(v1, "foo bar")
                                    for fi28 := 0; fi28 < 13; fi28++ {
                                        v27 = (v4 <= v4)
                                        if v27 {
                                            fmt.Println(fi25, fi28, v26)
                                        }
                                        var v29 uint8 = 10
                                        v19 = (v19 >> ((926) & 63))
                                        v19 += v19
                                        _ = fi28
                                        _ = v29
                                    }
                                    switch (v4 + v4) {
                                    case "foo bar":
                                        var v30 uint = 1
                                        _ = v30
                                    case "\n\t":
                                        v31 := GenS6{F0: 127}
                                        _ = v31
                                    case "hello":
                                        for fi32 := 0; fi32 < 8; fi32++ {
                                            var v33 int32 = -681
                                            switch v4 {
                                            case "":
                                                v2 = (0.0 * v2)
                                            case "hello":
                                                v34 := GenS1{F0: (v26 - (v26 ^ v26))}
                                                _ = v34
                                            }
                                            v4 = v4
                                            var v35 uint16 = 32767
                                            _ = fi32
                                            _ = v33
                                            _ = v35
                                        }
                                    default:
                                        v1[(v4 + "foo bar")] = v5
                                    }
                                } else {
                                    fmt.Println(v3)
                                    var v36 uint8 = 255
                                    for fi37 := 0; fi37 < 16; fi37++ {
                                        v2 = 0.1
                                        v19 += (^(v19))
                                        _ = fi37
                                    }
                                    _ = v36
                                }
                            }
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                    if (false && (1 != (-(v3)))) {
                        for fi38 := 0; fi38 < 16; fi38++ {
                            v39 := GenS7{F0: 255}
                            v21 -= 0.1
                            _ = fi38
                            _ = v39
                        }
                        v5 = v5
                        v19 = 0
                    }
                } else {
                    var v40 bool = false
                    _ = v40
                }
                if (true || (!false)) {
                    for fi41 := 0; fi41 < 16; fi41++ {
                        v21 -= v21
                        switch v5 {
                        case 440:
                            if (((255 != 255) || false) && true) {
                                if ((v3 + v3) > v3) {
                                    v19 += (v19 & 0)
                                    var v42 uint32 = 4294967295
                                    v4 = "\n\t"
                                    _ = v42
                                } else {
                                    v43 := GenS4{F0: 0}
                                    v3 = v3
                                    _ = v43
                                }
                            } else {
                                for fi44 := 0; fi44 < 8; fi44++ {
                                    v45 := GenS8{F0: -9223372036854775808}
                                    var v46 uint32 = 155
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                }
                            }
                        default:
                            for fi47 := 0; fi47 < 4; fi47++ {
                                v4 = v4
                                _ = fi47
                            }
                        }
                        _ = fi41
                    }
                    var v48 uint32 = 366
                    var v49 int8 = 36
                    _ = v48
                    _ = v49
                } else {
                    if ((((v21 + ((0.1 - v21) + (((v21 / v21) / float32(v2)) - v21))) != (0.0 + (v21 + -2.25))) || (true && (!true))) || true) {
                        v19 -= v19
                    } else {
                        for fi50 := 0; fi50 < 11; fi50++ {
                            var v51 uint = 4294967295
                            _ = fi50
                            _ = v51
                        }
                        fmt.Println(v5)
                        fmt.Println(v5, v21, v21)
                    }
                    v21 = 0.1
                }
            }
            if (128 < 42) {
                if false {
                    v52 := v1[("a" + "foo bar")]
                    v53 := v1[v4]
                    _ = v52
                    _ = v53
                } else {
                    _ = len(v1)
                    v4 = ((v4 + v4) + v4)
                }
                if ((v19 & v19) >= (v19 | 0)) {
                    _ = len(v9)
                    v3 -= (v3 + v3)
                } else {
                    v54 := GenS1{F0: 42}
                    v2 += ((v2 - (1e10 + v2)) + (float64(3.14159) / v2))
                    _ = len(v1)
                    _ = v54
                }
                fmt.Println(v21)
            } else {
                v55 := GenS2{F0: 128, F1: 80}
                v56 := GenS1{F0: 1}
                _ = v55
                _ = v56
            }
        }
        var v57 float32 = v21
        _ = v23
        _ = v57
    }
    v58 := GenS8{F0: 652}
    fmt.Println(v2, v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v9
}
