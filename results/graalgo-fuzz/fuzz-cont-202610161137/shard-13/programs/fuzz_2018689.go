package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS3 struct { F0 string; F1 int64; F2 int }
type GenS4 struct { F0 uint64; F1 string; F2 int }
func genFn0(p0 int) int16 {
    v7 := GenS1{F0: 4294967295, F1: 0}
    var v8 uint = 918
    _ = p0
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 42
    var v4 int32 = 1
    var v5 string = ("foo bar" + "hello")
    v6 := make(map[string]bool)
    var v9 int = 2147483647
    v10 := GenS1{F0: 87, F1: 42}
    if true {
        var v11 int8 = 127
        switch "a" {
        case "":
            fmt.Println(v11, v4)
        }
        var v12 float64 = 0.0
        _ = v11
        _ = v12
    }
    if true {
        switch "\n\t" {
        case "hello":
            v5 = v5
        default:
            fmt.Println(v1, v1, v9)
        }
        v6[v5] = true
    }
    switch v9 {
    case -1:
        v13 := GenS2{F0: 42, F1: 40, F2: 42}
        _ = v13
    }
    if false {
        v14 := GenS2{F0: 2, F1: 42, F2: -124}
        if false {
            var v15 int16 = 1000
            if (2 < 42) {
                fmt.Println(v4, v3)
                if true {
                    if false {
                        v9 += 9223372036854775807
                        switch "" {
                        case "hello":
                            v5 = "hello"
                        }
                    } else {
                        v16 := []int{v9, int(v15), v9}
                        var v17 uint64 = (v3 | 380)
                        _ = v16
                        _ = v17
                    }
                    v18 := GenS1{F0: 107, F1: -324}
                    _ = v18
                }
            } else {
                if (2147483648 <= 65535) {
                    for fi19 := 0; fi19 < 18; fi19++ {
                        v5 = ""
                        _ = fi19
                    }
                    v20 := []int{}
                    v21 := GenS1{F0: 740, F1: int16(v4)}
                    _ = v20
                    _ = v21
                }
            }
            if false {
                switch 9223372036854775807 {
                case 499:
                    var v22 int8 = 42
                    _ = v22
                case 4294967295:
                    switch v5 {
                    case "":
                        var v23 float32 = 0.0
                        _ = v23
                    case "hello":
                        for fi24 := 0; fi24 < 4; fi24++ {
                            switch (fi24 << ((4294967295) & 63)) {
                            case 13:
                                for fi25 := 0; fi25 < 2; fi25++ {
                                    if (225 == 42) {
                                        v5 = ""
                                        if (!(((-2.25 - 0.0) + 1.5) > 1.5)) {
                                            var v26 int16 = 1000
                                            var v27 int = fi24
                                            v28 := GenS2{F0: 65535, F1: 255, F2: 5}
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        }
                                        _ = len(v2)
                                    } else {
                                        var v29 uint64 = (v3 & v3)
                                        v30 := GenS2{F0: 32767, F1: 255, F2: 42}
                                        _ = v29
                                        _ = v30
                                    }
                                    var v31 int8 = -1
                                    var v32 int16 = (v15 | v15)
                                    v33 := GenS2{F0: 65535, F1: 128, F2: v31}
                                    _ = fi25
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                            }
                            v34 := []int{(fi24 / (((v9 | -9223372036854775808)) | 1)), int(v4), (fi24 >> ((785) & 63)), (fi24 * v9)}
                            _ = fi24
                            _ = v34
                        }
                    default:
                        for fi35 := 0; fi35 < 10; fi35++ {
                            switch ((v5 + "\n\t") + (v5 + (((v5 + ("\n\t" + v5)) + (v5 + ("\n\t" + v5))) + v5))) {
                            case "":
                                v36 := GenS1{F0: 0, F1: v15}
                                _ = v36
                            }
                            switch (fi35 / ((v9) | 1)) {
                            case 996:
                                v5 = v5
                            case -539:
                                switch v9 {
                                case 270:
                                    var v37 float64 = (3.14159 - 3.14159)
                                    _ = v37
                                default:
                                    if (!false) {
                                        if true {
                                            var v38 string = "hello"
                                            _ = v38
                                        }
                                        v3 -= v3
                                        v1 -= v1
                                    }
                                }
                            case -9223372036854775808:
                                v39 := GenS2{F0: 0, F1: 178, F2: 100}
                                _ = v39
                            default:
                                v40 := GenS3{F0: "foo bar", F1: 853, F2: v9}
                                _ = v40
                            }
                            _ = fi35
                        }
                    }
                case 127:
                    var v41 uint64 = v3
                    _ = v41
                default:
                    for fi42 := 0; fi42 < 4; fi42++ {
                        fmt.Println(v4, v9)
                        _ = fi42
                    }
                }
                var v43 int32 = v4
                _ = v43
            } else {
                v4 = v4
                v44 := []int{418, (v9 % (((v9 - 0)) | 1)), 128}
                v4 = (^(v4))
                _ = v44
            }
            _ = v15
        } else {
            fmt.Println(v9, v1, v1)
            fmt.Println(v9, v9, v4)
            v5 = v5
        }
        var v45 uint8 = 127
        _ = v14
        _ = v45
    } else {
        var v46 uint64 = v3
        v1 = v1
        _ = v46
    }
    v47 := GenS4{F0: v3, F1: v5, F2: v9}
    v6[""] = true
    var v48 int16 = genFn0(v9)
    var v49 uint = 4294967295
    switch v5 {
    case "foo bar":
        if true {
            var v50 uint8 = 0
            v51 := []int{}
            var v52 int8 = 104
            _ = v50
            _ = v51
            _ = v52
        } else {
            var v53 uint16 = 2
            _ = v53
        }
    case "":
        fmt.Println(v49, v3, v48)
    }
    v54 := []int{-819, v9, v9}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v54
    _ = v6
    _ = v9
}
