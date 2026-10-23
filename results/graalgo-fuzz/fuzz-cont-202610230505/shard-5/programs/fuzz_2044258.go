package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int64 }
type GenS2 struct { F0 bool; F1 int; F2 int16 }
type GenS3 struct { F0 uint32; F1 int; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool; F1 int32 }
func genFn0(p0 int16, p1 uint16, p2 uint8) int8 {
    p0 += (-(p0))
    v6 := []int{239, 4294967295}
    p0 -= (p0 | p0)
    _ = len(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -99
}
func genFn1() int32 {
    if (0 < 255) {
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := GenS1{F0: (fi7 ^ fi7), F1: 127, F2: -9223372036854775808}
            _ = fi7
            _ = v8
        }
        if (!(229 <= -2147483648)) {
            if (1 >= 4294967295) {
                v9 := GenS2{F0: (false || (false && (42 > 161))), F1: 128, F2: -32768}
                var v10 uint8 = 1
                _ = v10
                _ = v9
            }
        }
    } else {
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 14; fi12++ {
                v13 := GenS1{F0: 4294967295, F1: 127, F2: 368}
                var v14 int64 = -186
                if ("hello" >= "a") {
                    v15 := []int{fi12, (fi11 + fi12)}
                    v16 := GenS1{F0: 4294967295, F1: 42, F2: (v14 << ((255) & 63))}
                    _ = v15
                    _ = v16
                } else {
                    v17 := []int{fi11}
                    switch 32767 {
                    case 4294967295:
                        switch ((("foo bar" + "\n\t") + "hello") + "") {
                        case "":
                            v18 := GenS2{F0: (2 >= 2147483647), F1: 647, F2: -545}
                            _ = v18
                        case "a":
                            v17 = append(v17, fi11)
                        }
                    case -1:
                        var v19 bool = (255 == 42)
                        _ = v19
                    }
                    v17 = append(v17, fi11)
                    v17 = append(v17, (fi11 >> ((42) & 63)))
                    if true {
                        v14 = v14
                        v14 = (v14 & -805)
                    }
                    _ = v17
                }
                if (4294967295 == 912) {
                    if false {
                        v20 := []int{0, -551, (fi11 * (fi11 | fi12))}
                        _ = v20
                    } else {
                        v14 = -9223372036854775808
                    }
                    v14 -= (v14 ^ v14)
                    if (2147483647 == 0) {
                        v14 -= (v14 / (((v14 / ((v14) | 1))) | 1))
                        v21 := []int{fi12, fi12}
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                        _ = v21
                    }
                }
                var v22 uint16 = 255
                v22 -= (^(v22))
                _ = fi12
                _ = v13
                _ = v14
                _ = v22
            }
            _ = fi11
        }
        var v23 string = ""
        _ = v23
    }
    var v24 int16 = 42
    v25 := []int{1, -901}
    _ = v24
    _ = v25
    return -1
}

func main() {
    var v1 float32 = 0.1
    var v2 uint32 = 0
    v3 := make(map[int]int)
    var v4 uint = 4294967295
    var v5 int32 = -2147483648
    _ = len(v3)
    var v26 string = ""
    if false {
        v27 := GenS3{F0: (v2 << (((^(v4))) & 63)), F1: 595, F2: 42}
        var v28 uint = v4
        _ = v27
        _ = v28
    } else {
        v5 += 1
        v1 += v1
    }
    v1 = v1
    for fi29 := 0; fi29 < 8; fi29++ {
        if ((1e10 + ((v1 / v1) - v1)) > v1) {
            for fi30 := 0; fi30 < 6; fi30++ {
                var v31 uint64 = 98
                _ = fi30
                _ = v31
            }
        }
        _ = fi29
    }
    var v32 uint = v4
    switch 32768 {
    case 742:
        var v33 int64 = 42
        _ = v33
    case -2147483648:
        var v34 string = (("hello" + (v26 + v26)) + "hello")
        _ = v34
    }
    v35 := GenS4{F0: 42, F1: (^(v4))}
    v2 = 4294967295
    switch "" {
    case "":
        if (!true) {
            v36 := GenS5{F0: 255, F1: (v2 - (^(v2)))}
            _ = v36
        }
    case "hello":
        v4 -= v4
    default:
        for fi37 := 0; fi37 < 1; fi37++ {
            v38 := GenS6{F0: 42}
            _ = fi37
            _ = v38
        }
    }
    if (!true) {
        v39 := []int{}
        v40 := GenS7{F0: 694}
        fmt.Println(v5, v26, v5)
        _ = v39
        _ = v40
    } else {
        delete(v3, 735)
        v1 += (v1 + (float32(0.0) * v1))
        if true {
            for fi41 := 0; fi41 < 13; fi41++ {
                var v42 int64 = -9223372036854775808
                v43 := GenS7{F0: 2}
                v44 := []int{(fi41 << ((823) & 63)), fi41, fi41}
                _ = fi41
                _ = v42
                _ = v43
                _ = v44
            }
            if (v2 == 1) {
                fmt.Println(v1)
            } else {
                fmt.Println(v2, v2, v32)
            }
        }
    }
    switch -2147483648 {
    case 9223372036854775807:
        v32 = (v32 * (v4 ^ uint(v2)))
    case -1:
        v5 -= (v5 >> ((v4) & 63))
    case -594:
        var v45 float64 = 0.0
        _ = v45
    }
    _ = len(v3)
    v46 := GenS8{F0: (437 != 18446744073709551615), F1: v5}
    var v47 uint8 = 186
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
