package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int16; F1 int16; F2 uint8 }
type GenS4 struct { F0 bool; F1 int }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]bool)
    v3 := []uint32{(v1 ^ v1)}
    var v4 uint = 4294967295
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 float64 = (0.1 + (-2.25 * 0.1))
        v1 = 271
        _ = fi5
        _ = v6
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        if true {
            if (42 > 9) {
                v3 = append(v3, (^(v1)))
            } else {
                for fi7 := 0; fi7 < 9; fi7++ {
                    if (100 >= -36) {
                        var v8 int16 = 1000
                        _ = v8
                    } else {
                        v4 += (^(v4))
                        if ((560 > 2) || true) {
                            var v9 bool = false
                            _ = v9
                        }
                        for fi10 := 0; fi10 < 12; fi10++ {
                            v11 := GenS1{F0: (("hello" + "\n\t") + "\n\t")}
                            _ = fi10
                            _ = v11
                        }
                    }
                    v4 -= v4
                    _ = fi7
                }
                for fi12 := 0; fi12 < 10; fi12++ {
                    v13 := GenS2{F0: 1000, F1: 943}
                    var v14 uint16 = 1
                    v4 -= (v4 | v4)
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                v3 = append(v3, 0)
            }
        }
    case "a":
        v15 := GenS2{F0: 0, F1: 255}
        _ = v15
    case "\n\t":
        v16 := GenS1{F0: ("" + "")}
        _ = v16
    default:
        if false {
            if true {
                var v17 uint8 = 127
                _ = v17
            } else {
                v18 := GenS1{F0: "foo bar"}
                _ = v18
            }
        }
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 uint8 = 93
            fmt.Println(v4, fi20)
            var v22 string = ("hello" + "\n\t")
            if false {
                if false {
                    switch v22 {
                    case "\n\t":
                        v21 = v21
                    case "foo bar":
                        v22 = v22
                    case "hello":
                        var v23 int = fi19
                        _ = v23
                    }
                    var v24 int32 = -903
                    fmt.Println(v1)
                    _ = v24
                }
                var v25 uint = (v4 - (v4 | v4))
                var v26 uint = uint(v1)
                var v27 uint = v4
                _ = len(v2)
                if len(v3) > 0 {
                    v3[0] = (v1 >> ((1) & 63))
                }
                _ = v25
                _ = v26
                _ = v27
            } else {
                var v28 uint32 = (v1 | v1)
                _ = v28
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        if (!false) {
            if true {
                _ = len(v3)
            }
            for fi29 := 0; fi29 < 13; fi29++ {
                if ((float32(1.5) * (0.1 + (3.14159 + (1.5 - ((3.14159 + float32((3.14159 / (float64(0.1) + float64(3.14159))))) * (3.14159 - float32(-2.25))))))) != ((1.5 + 0.0) + ((0.1 - (float32(((1.5 - -2.25) - 3.14159)) * (0.1 + 1.5))) + (0.0 - 0.0)))) {
                    switch ("foo bar" + "a") {
                    case "a":
                        switch fi29 {
                        case 127:
                            v1 -= 2147483648
                        case 394:
                            var v30 int16 = -365
                            _ = v30
                        default:
                            var v31 uint32 = v1
                            _ = v31
                        }
                    case "":
                        v4 += v4
                    case "hello":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        var v32 float32 = 0.0
                        _ = v32
                    }
                    switch ("hello" + "hello") {
                    case "":
                        var v33 string = "foo bar"
                        _ = v33
                    }
                    if true {
                        v4 = 140
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v1 = (v1 - v1)
                        v4 -= (^(v4))
                        var v34 string = "foo bar"
                        v4 = (v4 & v4)
                        _ = v34
                    } else {
                        fmt.Println(v4, v1, fi19)
                        if (v1 > (v1 % ((1) | 1))) {
                            var v35 uint16 = 32767
                            _ = v35
                        }
                        if (fi19 == fi29) {
                            v1 = (v1 | uint32(v4))
                            fmt.Println(v4, fi19)
                            v36 := GenS3{F0: -1, F1: -32768, F2: 2}
                            _ = v36
                        } else {
                            fmt.Println(v4)
                            v1 += 0
                            var v37 uint8 = 128
                            v4 = 4294967295
                            _ = v37
                        }
                    }
                } else {
                    fmt.Println(v4, v4, v4)
                    if ((float32(float64(float32(1e10))) - (0.1 * 0.0)) >= 1e10) {
                        fmt.Println(fi29, fi29, v4)
                        delete(v2, fi29)
                    } else {
                        v1 += 2147483648
                        delete(v2, (fi19 / ((351) | 1)))
                        v1 = 4294967295
                    }
                    _ = len(v2)
                }
                var v38 int32 = -2147483648
                var v39 int8 = -1
                _ = fi29
                _ = v38
                _ = v39
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi19
    }
    for fi40 := 0; fi40 < 13; fi40++ {
        v1 += (^(v1))
        if true {
            if false {
                if false {
                    v41 := GenS1{F0: ((("foo bar" + ("" + ("hello" + ""))) + ("" + "")) + "")}
                    _ = v41
                }
                if len(v3) > 0 {
                    v3[0] = (^(v1))
                }
            }
        }
        _ = fi40
    }
    if false {
        var v42 float64 = 3.14159
        _ = v42
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        fmt.Println(v1, fi43)
        _ = fi43
    }
    var v44 uint32 = (^(v1))
    if false {
        var v45 int = 553
        v44 = (v44 + 2147483648)
        _ = v45
    }
    var v46 bool = ((uint32(v4) <= v1) || ((v1 | v44) < v44))
    for fi47 := 0; fi47 < 6; fi47++ {
        v46 = v46
        v48 := GenS4{F0: (!v46), F1: (fi47 / ((fi47) | 1))}
        _ = fi47
        _ = v48
    }
    var v49 uint64 = 511
    var v50 int8 = -128
    if true {
        var v51 int32 = -775
        _ = v51
    } else {
        v50 = 67
        v49 += (v49 << (((^(v4))) & 63))
        v4 += 65535
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v49
    _ = v50
}
