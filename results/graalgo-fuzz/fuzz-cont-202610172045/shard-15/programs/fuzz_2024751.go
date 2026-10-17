package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int8; F1 uint8; F2 string }
type GenS3 struct { F0 int8; F1 uint64; F2 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 260
    v3 := make(map[uint32]int)
    v2 = v2
    for fi4 := 0; fi4 < 12; fi4++ {
        switch (fi4 >> ((uint(v2)) & 63)) {
        case 42:
            v5 := GenS1{F0: 100, F1: "hello"}
            _ = v5
        case 0:
            switch ("hello" + "a") {
            case "foo bar":
                v6 := []int{(fi4 & 0)}
                _ = v6
            case "\n\t":
                fmt.Println(v2)
            case "hello":
                var v7 uint64 = (^(v2))
                _ = v7
            default:
                var v8 float64 = 3.14159
                _ = v8
            }
        case 127:
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 bool = false
                _ = fi9
                _ = v10
            }
        default:
            var v11 uint16 = 32767
            _ = v11
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            fmt.Println(fi4, v2)
            _ = fi12
        }
        _ = fi4
    }
    v13 := GenS1{F0: 127, F1: "hello"}
    fmt.Println(v2)
    fmt.Println(v2)
    v14 := GenS2{F0: 127, F1: 0, F2: ("" + (("\n\t" + "a") + "\n\t"))}
    if (2147483647 > 4294967295) {
        v2 -= 487
        v2 = v2
        switch 32768 {
        case 245:
            for fi15 := 0; fi15 < 4; fi15++ {
                if false {
                    switch -921 {
                    case -2147483648:
                        var v16 int = -1
                        _ = v16
                    case 0:
                        v2 -= v2
                    case 127:
                        if (-1 != -1) {
                            switch "\n\t" {
                            case "hello":
                                v2 += v2
                            case "a":
                                var v17 uint8 = 1
                                _ = v17
                            }
                        }
                    default:
                        v2 += 4294967295
                    }
                    for fi18 := 0; fi18 < 16; fi18++ {
                        fmt.Println(v2, fi18, fi18)
                        var v19 uint = 65535
                        _ = fi18
                        _ = v19
                    }
                    fmt.Println(v2, v2)
                }
                var v20 int = fi15
                v20 -= v20
                _ = fi15
                _ = v20
            }
        case 0:
            for fi21 := 0; fi21 < 13; fi21++ {
                v2 = v2
                switch "foo bar" {
                case "foo bar":
                    var v22 string = "hello"
                    _ = v22
                case "hello":
                    for fi23 := 0; fi23 < 17; fi23++ {
                        switch fi23 {
                        case -1:
                            var v24 float32 = 1.5
                            var v25 bool = true
                            var v26 int32 = 1
                            _ = v24
                            _ = v25
                            _ = v26
                        case -2147483648:
                            v2 = 18446744073709551615
                        case 32767:
                            if true {
                                delete(v1, 65535)
                                if true {
                                    var v27 int64 = -566
                                    _ = v27
                                } else {
                                    v28 := GenS2{F0: int8(fi23), F1: 2, F2: ""}
                                    var v29 uint8 = 127
                                    _ = v28
                                    _ = v29
                                }
                                v30 := []int{fi23, 0}
                                delete(v1, 502)
                                var v31 uint = 42
                                _ = v30
                                _ = v31
                            } else {
                                v32 := GenS3{F0: -28, F1: (v2 & (v2 / ((18446744073709551615) | 1))), F2: (!(!(!false)))}
                                var v33 uint = 509
                                _ = v32
                                _ = v33
                            }
                            v34 := GenS2{F0: -128, F1: 0, F2: ""}
                            var v35 uint32 = 42
                            _ = v34
                            _ = v35
                        default:
                            var v36 int32 = 400
                            _ = v36
                        }
                        fmt.Println(fi21)
                        _ = fi23
                    }
                case "a":
                    var v37 uint32 = 65535
                    _ = v37
                }
                for fi38 := 0; fi38 < 15; fi38++ {
                    var v39 string = ("hello" + "")
                    v2 -= (v2 >> ((18446744073709551615) & 63))
                    var v40 uint64 = v2
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                _ = fi21
            }
        }
    }
    if (!((v2 ^ (v2 | v2)) <= (^(v2)))) {
        if (0 != 42) {
            fmt.Println(v2)
            var v41 int32 = 1
            if (-375 >= 9223372036854775807) {
                fmt.Println(v2, v41)
                _ = len(v1)
            }
            _ = v41
        } else {
            v3[2] = 4294967295
            switch ("\n\t" + (("" + "a") + "")) {
            case "\n\t":
                v42 := []int{-85, 951, 0, -354}
                _ = v42
            case "a":
                v2 += v2
            }
            fmt.Println(v2)
        }
        v43 := []int{}
        if (false && true) {
            v44 := GenS3{F0: -128, F1: v2, F2: (((1e10 - -2.25) + 0.1) == (-2.25 + (0.0 - -2.25)))}
            if true {
                v45 := GenS3{F0: -1, F1: 1, F2: false}
                fmt.Println(v2)
                _ = v45
            } else {
                v46 := GenS1{F0: -1, F1: ""}
                v2 = 443
                v2 = v2
                _ = v46
            }
            _ = v44
        } else {
            var v47 string = "foo bar"
            for fi48 := 0; fi48 < 19; fi48++ {
                var v49 int64 = 175
                var v50 int16 = -1
                fmt.Println(v2)
                _ = fi48
                _ = v49
                _ = v50
            }
            if (65535 >= 160) {
                var v51 float64 = -2.25
                v47 = v47
                _ = v51
            }
            _ = v47
        }
        _ = v43
    }
    v2 = (v2 | 594)
    v2 += v2
    fmt.Println(v2)
    _ = len(v3)
    v2 = v2
    var v52 uint64 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v52
}
