package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int64; F1 int16; F2 string }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int; F1 int64 }

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]string)
    var v3 float32 = 0.0
    v4 := make(map[int]bool)
    if false {
        v3 -= 0.0
        if (1 != 19) {
            for fi5 := 0; fi5 < 1; fi5++ {
                var v6 int8 = 42
                fmt.Println(fi5, v3)
                _ = fi5
                _ = v6
            }
            switch 434 {
            case -331:
                if (!true) {
                    v1 = (v1 << ((18446744073709551615) & 63))
                    for fi7 := 0; fi7 < 7; fi7++ {
                        for fi8 := 0; fi8 < 2; fi8++ {
                            var v9 bool = true
                            var v10 uint32 = 42
                            switch "hello" {
                            case "a":
                                v11 := GenS1{F0: (v10 & 2147483648), F1: -9223372036854775808}
                                switch "a" {
                                case "a":
                                    switch ("\n\t" + "\n\t") {
                                    case "a":
                                        var v12 string = ((("a" + "\n\t") + "foo bar") + "a")
                                        _ = v12
                                    case "hello":
                                        fmt.Println(v10, fi8, v9)
                                        var v13 int8 = 127
                                        v14 := GenS1{F0: (v10 >> ((454) & 63)), F1: 42}
                                        _ = v13
                                        _ = v14
                                    default:
                                        var v15 uint32 = 65535
                                        _ = v15
                                    }
                                    v10 -= v10
                                    var v16 uint = 2
                                    v17 := GenS1{F0: (v10 | (^(v10))), F1: 1}
                                    _ = v16
                                    _ = v17
                                case "":
                                    var v18 uint64 = 2
                                    _ = v18
                                case "foo bar":
                                    fmt.Println(v9, v10)
                                }
                                _ = v11
                            case "hello":
                                v1 += 42
                            default:
                                var v19 uint8 = 129
                                _ = v19
                            }
                            _ = fi8
                            _ = v10
                            _ = v9
                        }
                        _ = fi7
                    }
                } else {
                    v20 := GenS1{F0: 1, F1: -62}
                    v1 = v1
                    if true {
                        fmt.Println(v1, v1)
                        if true {
                            for fi21 := 0; fi21 < 3; fi21++ {
                                v22 := GenS1{F0: 2147483647, F1: 521}
                                _ = fi21
                                _ = v22
                            }
                            v1 += v1
                            var v23 int8 = 36
                            _ = v23
                        }
                        v24 := GenS2{F0: 0, F1: 4294967295, F2: v1}
                        _ = v24
                    } else {
                        var v25 int = 42
                        _ = v25
                    }
                    _ = v20
                }
            }
            var v26 uint32 = 2147483648
            _ = v26
        }
    } else {
        var v27 uint32 = 2
        v28 := GenS3{F0: -9223372036854775808, F1: 763, F2: "foo bar"}
        _ = v27
        _ = v28
    }
    var v29 int16 = 1
    v1 = v1
    v30 := GenS4{F0: 42, F1: 970}
    if (!false) {
        var v31 int32 = 0
        switch -1 {
        case 127:
            var v32 uint = 626
            _ = v32
        case 0:
            switch "foo bar" {
            case "":
                v2[(((("" + "") + "hello") + "\n\t") + ("hello" + ""))] = "\n\t"
            case "foo bar":
                v29 = 32767
            default:
                v31 = v31
            }
        case 1:
            v33 := GenS4{F0: 948, F1: 4294967295}
            _ = v33
        }
        for fi34 := 0; fi34 < 12; fi34++ {
            switch (fi34 ^ (fi34 * fi34)) {
            case -1:
                var v35 float32 = ((v3 - v3) - (v3 + 0.1))
                _ = v35
            case 127:
                v36 := GenS1{F0: 1, F1: -835}
                _ = v36
            }
            _ = fi34
        }
        _ = v31
    } else {
        for fi37 := 0; fi37 < 4; fi37++ {
            v38 := []int{(fi37 ^ fi37), fi37}
            if true {
                delete(v2, "hello")
                for fi39 := 0; fi39 < 19; fi39++ {
                    v29 += (v29 * (v29 ^ (v29 & (v29 + (v29 % ((v29) | 1))))))
                    _ = fi39
                }
            }
            _ = fi37
            _ = v38
        }
    }
    v40 := GenS5{F0: -167, F1: 9223372036854775807}
    v1 += (v1 >> ((1) & 63))
    var v41 int8 = -1
    var v42 int16 = (v29 * v29)
    if (0 == 9223372036854775807) {
        var v43 int = 831
        var v44 uint16 = (v1 % (((v1 ^ (v1 % ((v1) | 1)))) | 1))
        _ = v43
        _ = v44
    } else {
        v42 = 42
    }
    var v45 int8 = (v41 ^ (v41 / ((v41) | 1)))
    var v46 int = 32767
    fmt.Println(v41)
    v29 -= (v42 ^ v29)
    var v47 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v45
    _ = v46
    _ = v47
}
