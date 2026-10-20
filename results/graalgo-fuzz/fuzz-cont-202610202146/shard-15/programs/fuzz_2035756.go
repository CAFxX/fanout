package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }
type GenS2 struct { F0 int; F1 bool; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 int8; F1 uint16; F2 uint }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 uint64 }
func genFn0(p0 uint8, p1 string, p2 int) bool {
    p2 = (^(p2))
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 bool = false
        v9 := GenS1{F0: -716, F1: 4294967295, F2: 1}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := []int{p2, p2, fi10}
        p0 -= (p0 / ((42) | 1))
        var v12 uint8 = 255
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := []int{(p2 >> ((4294967295) & 63)), (p2 | (-(p2))), (^(p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return false
}
func genFn1(p0 int) int32 {
    v14 := []int{p0}
    if false {
        var v15 uint64 = 187
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        p0 += p0
        _ = v15
    }
    switch "a" {
    case "a":
        p0 = (^(p0))
    case "":
        for fi16 := 0; fi16 < 20; fi16++ {
            for fi17 := 0; fi17 < 17; fi17++ {
                p0 -= (fi17 / ((fi16) | 1))
                var v18 int8 = 55
                _ = fi17
                _ = v18
            }
            var v19 int = p0
            var v20 uint16 = 42
            _ = fi16
            _ = v19
            _ = v20
        }
    case "foo bar":
        if len(v14) > 0 {
            v14[0] = p0
        }
    }
    p0 = (p0 >> ((255) & 63))
    v14 = append(v14, 115)
    _ = p0
    _ = v14
    return 1
}

func main() {
    var v1 int = -1
    v2 := make(map[uint32]int)
    v3 := []string{}
    var v4 uint8 = 0
    var v5 uint8 = v4
    var v6 uint8 = v4
    var v21 uint64 = 3
    v6 += (v4 / (((v6 - 109)) | 1))
    v6 -= (v6 + (v5 - (v6 % ((1) | 1))))
    switch 42 {
    case 912:
        if ((178 > 128) || (4294967295 < 65535)) {
            v3 = append(v3, ("foo bar" + "foo bar"))
            v6 = v6
            fmt.Println(v4)
        } else {
            var v22 int32 = -406
            var v23 bool = false
            _ = v22
            _ = v23
        }
    case 558:
        if false {
            v3 = append(v3, "")
        } else {
            for fi24 := 0; fi24 < 2; fi24++ {
                v25 := v2[1]
                v25 += v1
                _ = fi24
                _ = v25
            }
        }
    }
    v26 := GenS1{F0: (v1 + v1), F1: 18446744073709551615, F2: (v21 | (^(v21)))}
    if true {
        v27 := GenS2{F0: -635, F1: false, F2: "\n\t"}
        _ = len(v3)
        _ = v27
    }
    v28 := GenS3{F0: ("foo bar" + "foo bar")}
    for fi29 := 0; fi29 < 20; fi29++ {
        switch ("hello" + "") {
        case "foo bar":
            v30 := GenS4{F0: 256, F1: 679}
            _ = v30
        case "":
            if false {
                v5 += 127
                v2[2147483648] = 103
                fmt.Println(v1)
            } else {
                if (-388 != 0) {
                    v31 := GenS1{F0: (v1 >> ((2) & 63)), F1: 4294967295, F2: v21}
                    if (false && true) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v32 uint = 42
                        _ = v32
                    } else {
                        v4 -= 42
                        var v33 float64 = 0.1
                        _ = v33
                    }
                    var v34 int32 = 2147483647
                    _ = v31
                    _ = v34
                }
                if false {
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                    fmt.Println(fi29, fi29, v1)
                }
            }
        default:
            var v35 int8 = 1
            _ = v35
        }
        var v36 int = v1
        _ = fi29
        _ = v36
    }
    v5 += (v4 | (v6 ^ v6))
    var v37 float32 = (1e10 - 0.1)
    var v38 int32 = 2147483647
    v39 := GenS5{F0: v21, F1: -964}
    var v40 uint64 = (v21 - 0)
    if (0.0 == 1.5) {
        v37 = 1.5
        v41 := GenS6{F0: 127, F1: 32767, F2: 886}
        _ = v41
    }
    v42 := GenS7{F0: 1, F1: v21}
    v40 = (v40 / ((v21) | 1))
    v40 += v21
    v43 := GenS8{F0: (v40 << ((580) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
