package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint64; F2 int16 }
func genFn0(p0 string, p1 bool, p2 uint64) uint {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint16 = 945
        v9 := []int{(fi7 / (((fi7 & (-(fi7)))) | 1)), -1, (-(fi7))}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{9223372036854775807, -2147483648}
    var v11 int8 = -44
    var v12 int32 = 1
    fmt.Println(p0, p0, v12)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v4 bool = true
    var v5 int64 = 0
    v6 := make(map[uint32]bool)
    var v13 int8 = 100
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v13, v5, v13)
    }
    var v14 uint32 = 0
    switch ("hello" + "foo bar") {
    case "foo bar":
        v4 = v4
    case "":
        var v15 string = "\n\t"
        _ = v15
    default:
        v16 := GenS1{F0: (int16(v13) < -32768), F1: 32767, F2: 77}
        _ = v16
    }
    if v4 {
        if ("\n\t" > "hello") {
            switch (("" + "\n\t") + "") {
            case "\n\t":
                if v4 {
                    v5 -= v5
                    if (!v4) {
                        var v17 uint32 = v14
                        _ = v17
                    } else {
                        for fi18 := 0; fi18 < 8; fi18++ {
                            if v4 {
                                v19 := []int{(fi18 + fi18), fi18, 1, (fi18 ^ (fi18 | fi18))}
                                switch (-(fi18)) {
                                case 9223372036854775807:
                                    v13 = v13
                                case -39:
                                    for fi20 := 0; fi20 < 17; fi20++ {
                                        if true {
                                            switch (fi18 % (((fi20 - fi20)) | 1)) {
                                            case 82:
                                                v5 -= v5
                                            case 2147483647:
                                                _ = len(v19)
                                            case -1:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    switch fi18 {
                                                    case 4294967295:
                                                        fmt.Println(v13, fi20, v13)
                                                    default:
                                                        v21 := GenS1{F0: v4, F1: -501, F2: 128}
                                                        _ = v21
                                                    }
                                                case "":
                                                    v22 := GenS1{F0: ((228 < 32767) || (539 != 42)), F1: 32767, F2: 42}
                                                    _ = v22
                                                case "hello":
                                                    fmt.Println(v5, v13)
                                                }
                                            }
                                        } else {
                                            v4 = false
                                            var v23 uint64 = uint64(v14)
                                            var v24 int64 = (v5 & (v5 ^ (v5 >> ((0) & 63))))
                                            _ = v23
                                            _ = v24
                                        }
                                        _ = fi20
                                    }
                                }
                                for fi25 := 0; fi25 < 10; fi25++ {
                                    var v26 uint32 = v14
                                    if v4 {
                                        var v27 bool = true
                                        var v28 uint = 42
                                        v4 = (v28 < v28)
                                        _ = v27
                                        _ = v28
                                    } else {
                                        v4 = v4
                                        var v29 int16 = 0
                                        if (-2147483648 == fi18) {
                                            var v30 int32 = -1
                                            switch fi18 {
                                            case 32767:
                                                var v31 int64 = v5
                                                _ = v31
                                            case 1:
                                                v19 = append(v19, fi25)
                                            case -2147483648:
                                                var v32 uint8 = 1
                                                _ = v32
                                            default:
                                                v33 := GenS1{F0: true, F1: (v29 + (v29 * -781)), F2: 1}
                                                _ = v33
                                            }
                                            v34 := GenS1{F0: v4, F1: v29, F2: 127}
                                            _ = v30
                                            _ = v34
                                        }
                                        _ = v29
                                    }
                                    _ = fi25
                                    _ = v26
                                }
                                _ = v19
                            }
                            var v35 int = 4294967295
                            if true {
                                fmt.Println(v13, v4, fi18)
                                v36 := GenS2{F0: v4, F1: (^(v14))}
                                var v37 int32 = 1
                                _ = v36
                                _ = v37
                            } else {
                                v38 := GenS2{F0: v4, F1: 0}
                                _ = v38
                            }
                            _ = fi18
                            _ = v35
                        }
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v13)
                        case "a":
                            for fi39 := 0; fi39 < 18; fi39++ {
                                var v40 int16 = -993
                                _ = fi39
                                _ = v40
                            }
                        case "hello":
                            v14 = (^(v14))
                        }
                    }
                    v41 := GenS1{F0: (!((v13 % (((-(v13))) | 1)) >= 42)), F1: -919, F2: 127}
                    _ = v41
                } else {
                    v5 += v5
                }
            }
            if false {
                fmt.Println(v5)
                v14 = 42
            }
        }
    }
    var v42 string = "foo bar"
    v43 := GenS3{F0: v4, F1: 4294967295, F2: -1}
    if ((1.5 <= 1.5) || (1 < (v5 >> ((255) & 63)))) {
        v14 = 962
    } else {
        var v44 int16 = 0
        fmt.Println(v14, v13, v13)
        var v45 uint32 = v14
        _ = v44
        _ = v45
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v14, v14, v42)
    case "a":
        v46 := GenS1{F0: (false || v4), F1: 42, F2: 0}
        _ = v46
    }
    v14 += v14
    var v47 float32 = 1.5
    v13 = v13
    v4 = false
    if ((220 >= 4294967295) || v4) {
        var v48 uint16 = 42
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v47
    _ = v5
    _ = v6
}
