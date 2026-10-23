package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }
type GenS3 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (((0.1 - (float32(1e10) - 3.14159)) + (1e10 - float32(3.14159))) - 1e10)
    var v3 bool = true
    v3 = (true || (!v3))
    fmt.Println(v2, v3)
    var v4 uint = 18446744073709551615
    var v5 int16 = 32767
    if true {
        v6 := v1[9223372036854775807]
        _ = v6
    } else {
        fmt.Println(v5, v2)
    }
    fmt.Println(v5)
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(fi7, v3, v2)
        for fi8 := 0; fi8 < 5; fi8++ {
            if true {
                switch (("" + "foo bar") + "hello") {
                case "hello":
                    v5 -= (-(v5))
                case "foo bar":
                    var v9 int8 = int8(v4)
                    _ = v9
                case "":
                    fmt.Println(v5, fi7)
                    v5 -= (v5 / ((v5) | 1))
                    _ = len(v1)
                }
                var v10 int16 = v5
                v5 = 0
                _ = v10
            } else {
                if ("a" != ("hello" + "hello")) {
                    v4 -= 18446744073709551615
                    var v11 uint8 = 23
                    _ = v11
                }
                if (((44 == 640) || v3) && (v2 != float32(-2.25))) {
                    fmt.Println(v5, fi8, v2)
                    v12 := GenS1{F0: v4, F1: 42, F2: (v4 % ((2) | 1))}
                    if (!v3) {
                        fmt.Println(v5)
                        if v3 {
                            switch fi7 {
                            case -537:
                                v13 := GenS1{F0: (v4 - (v4 % (((v4 & (^(v4)))) | 1))), F1: 127, F2: (v4 * (v4 | v4))}
                                _ = len(v1)
                                _ = v13
                            case -40:
                                delete(v1, fi7)
                            default:
                                v14 := GenS1{F0: v4, F1: 25, F2: 42}
                                _ = v14
                            }
                            if v3 {
                                var v15 uint8 = 1
                                if false {
                                    if (42 <= uint64(v4)) {
                                        if ((!(float32((3.14159 - 1e10)) >= v2)) || (2 >= 42)) {
                                            v16 := []int{(^(fi7)), (fi7 & fi8)}
                                            _ = v16
                                        } else {
                                            switch (fi8 | (fi8 | (fi7 - fi8))) {
                                            case 1:
                                                var v17 int = -1
                                                _ = v17
                                            default:
                                                v18 := []int{}
                                                _ = v18
                                            }
                                            if v3 {
                                                _ = len(v1)
                                                var v19 uint16 = 605
                                                _ = v19
                                            }
                                        }
                                        switch (fi8 / ((32767) | 1)) {
                                        case -793:
                                            v3 = v3
                                        case 0:
                                            fmt.Println(v5, v4)
                                        case 4294967295:
                                            var v20 uint64 = 183
                                            _ = v20
                                        }
                                        if v3 {
                                            var v21 int64 = int64(v4)
                                            if false {
                                                v22 := GenS1{F0: v4, F1: v15, F2: (^(v4))}
                                                v3 = v3
                                                var v23 uint64 = 4294967295
                                                _ = v22
                                                _ = v23
                                            }
                                            _ = v21
                                        } else {
                                            var v24 uint8 = v15
                                            v24 = 1
                                            _ = v24
                                        }
                                    }
                                    fmt.Println(v15)
                                    fmt.Println(v2, v5, v5)
                                } else {
                                    v25 := GenS2{F0: (^(v15)), F1: 156, F2: ""}
                                    _ = v25
                                }
                                switch ((("hello" + ("\n\t" + "")) + "") + "") {
                                case "hello":
                                    v4 += v4
                                case "a":
                                    var v26 bool = ((-2.25 + 0.0) <= (float64(0.1) + (1.5 * (1e10 * 0.0))))
                                    _ = v26
                                case "foo bar":
                                    switch "a" {
                                    case "\n\t":
                                        if v3 {
                                            v27 := GenS1{F0: v4, F1: 127, F2: v4}
                                            _ = v27
                                        }
                                    case "hello":
                                        if (((3.14159 * 1e10) - (((float64((float32(0.0) * v2)) - 3.14159) + 0.1) - 1e10)) < -2.25) {
                                            v4 = (v4 | v4)
                                        }
                                        switch "\n\t" {
                                        case "a":
                                            var v28 float32 = (v2 + (v2 + (-2.25 + (v2 / float32(1.5)))))
                                            _ = v28
                                        }
                                    case "":
                                        v15 -= (^(v15))
                                    default:
                                        v29 := GenS3{F0: 59, F1: (^(v15)), F2: -676}
                                        _ = v29
                                    }
                                }
                                _ = v15
                            } else {
                                var v30 string = "\n\t"
                                _ = v30
                            }
                        } else {
                            v31 := []int{}
                            v32 := GenS1{F0: (v4 << ((0) & 63)), F1: 0, F2: 666}
                            _ = v31
                            _ = v32
                        }
                    } else {
                        v5 = 0
                        var v33 int = fi7
                        _ = v33
                    }
                    _ = v12
                } else {
                    switch 1 {
                    case 9223372036854775807:
                        if (v3 && true) {
                            v34 := v1[fi7]
                            _ = v34
                        }
                        if (v3 && (!v3)) {
                            var v35 string = "a"
                            _ = v35
                        } else {
                            if false {
                                v4 += 2
                                v3 = (3.14159 < 1e10)
                                v36 := GenS2{F0: 200, F1: 276, F2: "\n\t"}
                                _ = v36
                            }
                            var v37 uint8 = 60
                            delete(v1, (fi7 | 38))
                            _ = v37
                        }
                        v4 -= v4
                    case 32768:
                        if false {
                            v2 += v2
                            v2 += (v2 * 3.14159)
                        } else {
                            v1[fi8] = "a"
                            fmt.Println(v4)
                        }
                    default:
                        v3 = (int8(v4) != 1)
                        var v38 uint64 = 42
                        var v39 int8 = 1
                        _ = v38
                        _ = v39
                    }
                    switch fi8 {
                    case 32768:
                        v40 := []int{(fi7 << ((42) & 63)), (fi8 ^ (fi8 >> (((v4 & (v4 >> ((v4) & 63)))) & 63)))}
                        _ = v40
                    default:
                        v41 := []int{(fi8 ^ fi7), (fi8 >> (((v4 + (v4 * (v4 << ((v4) & 63))))) & 63)), (fi8 + fi7), 42}
                        _ = v41
                    }
                    v4 -= 42
                }
            }
            var v42 uint64 = 538
            var v43 uint32 = 2147483647
            v5 += (v5 ^ (v5 / (((v5 | -32768)) | 1)))
            _ = fi8
            _ = v42
            _ = v43
        }
        var v44 uint16 = 32767
        _ = fi7
        _ = v44
    }
    for fi45 := 0; fi45 < 10; fi45++ {
        var v46 int32 = 836
        switch fi45 {
        case 500:
            for fi47 := 0; fi47 < 5; fi47++ {
                v48 := GenS4{F0: 1}
                var v49 bool = true
                v50 := GenS5{F0: 161, F1: 42}
                var v51 int16 = v5
                _ = fi47
                _ = v48
                _ = v49
                _ = v50
                _ = v51
            }
        case 32768:
            var v52 uint32 = 2147483647
            _ = v52
        default:
            fmt.Println(v46, v3, v3)
        }
        fmt.Println(v4)
        _ = fi45
        _ = v46
    }
    for fi53 := 0; fi53 < 9; fi53++ {
        v54 := GenS1{F0: v4, F1: 2, F2: v4}
        _ = fi53
        _ = v54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
