package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 string; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 bool }
func genFn0(p0 uint) string {
    for fi4 := 0; fi4 < 16; fi4++ {
        for fi5 := 0; fi5 < 11; fi5++ {
            switch (fi5 | fi4) {
            case 253:
                fmt.Println(fi5)
            case 0:
                v6 := []int{}
                _ = v6
            case -548:
                switch (fi5 & -207) {
                case 32767:
                    if false {
                        v7 := []int{(-(fi5))}
                        _ = v7
                    } else {
                        v8 := GenS1{F0: 2147483648, F1: (fi5 >> ((4294967295) & 63)), F2: ("a" + "\n\t")}
                        _ = v8
                    }
                case -183:
                    fmt.Println(fi5, fi4, fi4)
                default:
                    if (false && ("hello" == "a")) {
                        if true {
                            var v9 uint = (p0 - (^(p0)))
                            _ = v9
                        } else {
                            var v10 uint8 = 2
                            var v11 int = fi4
                            fmt.Println(fi4, v10, v11)
                            _ = v10
                            _ = v11
                        }
                    } else {
                        var v12 uint64 = 18446744073709551615
                        fmt.Println(v12, fi5, fi5)
                        if true {
                            if (!false) {
                                v13 := []int{}
                                v12 -= (v12 - (^(v12)))
                                _ = v13
                            } else {
                                v12 -= (v12 + (^(v12)))
                                var v14 uint64 = (v12 - 42)
                                _ = v14
                            }
                            v15 := GenS2{F0: -2147483648}
                            fmt.Println(v12, fi4, p0)
                            _ = v15
                        }
                        _ = v12
                    }
                    v16 := []int{(fi5 * (^(fi4)))}
                    var v17 uint16 = 0
                    _ = v16
                    _ = v17
                }
            default:
                if (!("\n\t" != "a")) {
                    var v18 bool = false
                    switch fi5 {
                    case 128:
                        var v19 float32 = (-2.25 + (((1.5 * 1e10) - -2.25) + -2.25))
                        _ = v19
                    case -2147483648:
                        v20 := GenS3{F0: 255, F1: "\n\t", F2: 128}
                        _ = v20
                    default:
                        if false {
                            v21 := GenS3{F0: 255, F1: "a", F2: 127}
                            var v22 int32 = 42
                            _ = v21
                            _ = v22
                        } else {
                            fmt.Println(p0)
                            v18 = (v18 || ((!(100 >= 1)) && true))
                        }
                    }
                    v23 := GenS4{F0: 128, F1: 1}
                    _ = v18
                    _ = v23
                }
            }
            switch ("\n\t" + "foo bar") {
            case "hello":
                if (!true) {
                    if true {
                        var v24 bool = (1 < 840)
                        switch "" {
                        case "hello":
                            v25 := []int{}
                            _ = v25
                        default:
                            var v26 uint32 = 0
                            _ = v26
                        }
                        v27 := GenS3{F0: uint8(p0), F1: ("hello" + "a"), F2: 42}
                        _ = v24
                        _ = v27
                    } else {
                        v28 := []int{(fi5 ^ 0), fi4, 32767}
                        v28 = append(v28, -2147483648)
                        var v29 int8 = 107
                        _ = v28
                        _ = v29
                    }
                    v30 := []int{0, (fi4 >> (((p0 + (p0 - p0))) & 63)), (fi4 ^ (fi4 - (-(fi4)))), (fi4 % ((928) | 1))}
                    fmt.Println(fi5, fi5, fi5)
                    p0 -= (p0 - p0)
                    switch (fi5 ^ (fi4 << ((p0) & 63))) {
                    case -2147483648:
                        p0 -= (p0 % (((p0 + p0)) | 1))
                    case 112:
                        p0 += (p0 % ((p0) | 1))
                    }
                    _ = v30
                }
            case "\n\t":
                var v31 float32 = 0.0
                _ = v31
            default:
                p0 = (p0 + p0)
            }
            p0 = (p0 / ((2) | 1))
            _ = fi5
        }
        var v32 int = (fi4 >> ((805) & 63))
        _ = fi4
        _ = v32
    }
    p0 -= p0
    var v33 uint8 = 2
    v34 := GenS5{F0: false}
    _ = p0
    _ = v33
    _ = v34
    return ""
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 886
    var v3 int64 = 0
    var v35 string = "hello"
    switch 42 {
    case -674:
        v35 = v35
    case -1:
        v36 := []int{9223372036854775807}
        _ = v36
    }
    var v37 bool = false
    if (2147483647 > -860) {
        switch v35 {
        case "a":
            var v38 uint32 = 42
            _ = v38
        default:
            v3 += v3
        }
        var v39 int64 = 9223372036854775807
        _ = v39
    }
    if false {
        v40 := []int{483, 128, 9223372036854775807, 2147483647}
        _ = v40
    } else {
        var v41 float64 = 0.1
        var v42 bool = (2147483648 > 65535)
        v37 = (65535 <= 65535)
        _ = v41
        _ = v42
    }
    v43 := GenS6{F0: (-1 == 231), F1: (-109 <= 127)}
    if v37 {
        switch ((v35 + v35) + "") {
        case "":
            fmt.Println(v37)
        case "hello":
            switch (("" + "") + ((v35 + v35) + (v35 + v35))) {
            case "foo bar":
                var v44 uint8 = 127
                _ = v44
            }
        default:
            v37 = v37
        }
        v45 := []int{240, 4294967295}
        _ = v45
    }
    var v46 uint8 = 203
    var v47 float64 = 1e10
    switch 1 {
    case 128:
        fmt.Println(v3, v47)
    case -759:
        var v48 uint32 = 0
        _ = v48
    case 4294967295:
        if (v37 || (!true)) {
            for fi49 := 0; fi49 < 8; fi49++ {
                v50 := []int{(fi49 << ((255) & 63)), 0}
                _ = fi49
                _ = v50
            }
            var v51 float32 = 1.5
            _ = v51
        } else {
            fmt.Println(v37, v47, v46)
            if false {
                v37 = ((1e10 + ((-2.25 + float32(float64(0.1))) / (1.5 - ((-2.25 + 0.1) * float32(v47))))) != 1.5)
                v46 = (v46 << ((uint(v2)) & 63))
                var v52 int = int(v3)
                _ = v52
            }
            v35 = (v35 + "hello")
        }
    }
    v53 := []int{2147483647, 127}
    fmt.Println(v35)
    v54 := GenS3{F0: (v46 | v46), F1: (v35 + (v35 + v35)), F2: (v46 >> ((18446744073709551615) & 63))}
    v2 = (^(v2))
    v47 += (v47 + v47)
    if (-1 > 127) {
        var v55 int32 = 1
        v46 = (v46 << ((395) & 63))
        _ = v55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v53:", len(v53), v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v43
    _ = v46
    _ = v47
    _ = v53
    _ = v54
}
