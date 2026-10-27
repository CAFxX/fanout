package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool; F1 int16 }
func genFn0(p0 uint32) int {
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(fi5, p0)
        p0 -= p0
        _ = fi5
    }
    v6 := GenS1{F0: 9223372036854775807}
    for fi7 := 0; fi7 < 5; fi7++ {
        p0 -= (p0 + p0)
        _ = fi7
    }
    var v8 int8 = -84
    _ = p0
    _ = v6
    _ = v8
    return 127
}
func genFn1(p0 bool, p1 uint16) int64 {
    v9 := GenS1{F0: 42}
    fmt.Println(p1, p1)
    switch ("" + ("a" + "\n\t")) {
    case "foo bar":
        if p0 {
            var v10 uint = 0
            v10 = v10
            fmt.Println(v10)
            _ = v10
        }
    }
    fmt.Println(p1, p0)
    p0 = p0
    _ = p0
    _ = p1
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 string = "a"
    var v2 int = 0
    var v3 int16 = -32768
    var v4 uint64 = 18446744073709551615
    switch v1 {
    case "a":
        var v11 uint16 = 1
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := []int{(fi12 % (((v2 / ((9223372036854775807) | 1))) | 1))}
            v14 := GenS2{F0: -1}
            _ = fi12
            _ = v13
            _ = v14
        }
    case "\n\t":
        v1 = (((v1 + (v1 + (((v1 + (v1 + "a")) + v1) + "foo bar"))) + "hello") + (((v1 + v1) + "") + v1))
    default:
        var v15 int64 = 9223372036854775807
        _ = v15
    }
    var v16 uint16 = 766
    if (!true) {
        var v17 int8 = 127
        var v18 uint = 42
        _ = v17
        _ = v18
    } else {
        v19 := GenS3{F0: 1, F1: v3, F2: 129}
        _ = v19
    }
    var v20 float32 = 1.5
    v21 := GenS4{F0: (v16 + 32767), F1: genFn1(true, v16)}
    v22 := []int{4294967295, v2, genFn0(769), -9223372036854775808}
    if true {
        var v23 float64 = ((3.14159 * 1e10) - 1e10)
        _ = v23
    }
    fmt.Println(v2, v16)
    v24 := GenS5{F0: 42}
    var v25 uint16 = (v16 / ((v16) | 1))
    switch 107 {
    case 110:
        if (1 > 42) {
            var v26 float32 = (v20 - (1.5 / (v20 + v20)))
            _ = v26
        } else {
            var v27 float64 = float64((float32(0.0) - float32(3.14159)))
            v16 = v16
            if ((!true) && (!false)) {
                _ = len(v22)
            }
            _ = v27
        }
    case 32768:
        switch (v1 + "a") {
        case "hello":
            var v28 uint8 = 195
            _ = v28
        case "foo bar":
            fmt.Println(v2, v25, v4)
        }
    default:
        v25 += (v25 | v25)
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        for fi30 := 0; fi30 < 19; fi30++ {
            var v31 int16 = (v3 * (v3 << ((65535) & 63)))
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        v33 := GenS3{F0: 1, F1: (v3 / ((int16(v4)) | 1)), F2: 251}
        _ = fi32
        _ = v33
    }
    v20 = float32((0.0 + 1e10))
    for fi34 := 0; fi34 < 18; fi34++ {
        var v35 int64 = int64(v25)
        if (true && (255 == 2)) {
            var v36 string = "foo bar"
            _ = v36
        } else {
            if false {
                var v37 bool = false
                if (v25 != (v16 >> ((1) & 63))) {
                    fmt.Println(v37)
                    if false {
                        v38 := GenS6{F0: v3, F1: (v4 / (((v4 % ((0) | 1))) | 1))}
                        var v39 int32 = 1
                        var v40 uint16 = (v25 & v16)
                        _ = v38
                        _ = v39
                        _ = v40
                    } else {
                        var v41 uint32 = 2147483648
                        _ = v41
                    }
                } else {
                    switch (fi34 ^ 42) {
                    case -9223372036854775808:
                        switch "" {
                        case "foo bar":
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                        case "":
                            var v42 int32 = -2147483648
                            _ = v42
                        case "\n\t":
                            var v43 uint16 = 1
                            _ = v43
                        }
                    case -1:
                        if false {
                            for fi44 := 0; fi44 < 14; fi44++ {
                                var v45 float64 = 3.14159
                                if (v16 > (v25 / (((v16 ^ (v25 / ((0) | 1)))) | 1))) {
                                    var v46 uint = 65535
                                    _ = v46
                                } else {
                                    if len(v22) > 0 {
                                        v22[0] = v2
                                    }
                                    switch v1 {
                                    case "foo bar":
                                        var v47 uint16 = (^(v25))
                                        _ = v47
                                    case "\n\t":
                                        var v48 int16 = (v3 + 517)
                                        v49 := GenS7{F0: 13}
                                        var v50 uint = 1
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    }
                                }
                                fmt.Println(v16, v20, v20)
                                _ = fi44
                                _ = v45
                            }
                        }
                    case 32768:
                        fmt.Println(v35)
                    default:
                        v51 := GenS8{F0: (v37 && ((!((v4 % (((v4 + v4)) | 1)) <= (v4 << ((0) & 63)))) && ((v35 + v35) != 678))), F1: int16(v25)}
                        _ = v51
                    }
                    var v52 int = int(v4)
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    _ = v52
                }
                _ = v37
            }
            if false {
                v1 = v1
                var v53 string = v1
                for fi54 := 0; fi54 < 12; fi54++ {
                    switch (fi34 * (v2 / ((fi34) | 1))) {
                    case 9223372036854775807:
                        if true {
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                        } else {
                            var v55 uint32 = 544
                            _ = v55
                        }
                    case 1:
                        v3 += v3
                    case 42:
                        fmt.Println(v25, v35)
                    default:
                        v22 = append(v22, (v2 | (v2 - genFn0(uint32(v4)))))
                    }
                    switch ((v53 + v1) + (v1 + (v1 + (v1 + ("foo bar" + v1))))) {
                    case "a":
                        v2 = fi34
                    case "hello":
                        if true {
                            v56 := GenS4{F0: v16, F1: 0}
                            var v57 uint64 = (v4 >> ((18446744073709551615) & 63))
                            _ = v56
                            _ = v57
                        }
                    }
                    v1 = ""
                    _ = fi54
                }
                _ = v53
            }
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
