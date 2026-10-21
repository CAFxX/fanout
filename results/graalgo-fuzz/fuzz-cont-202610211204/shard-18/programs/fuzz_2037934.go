package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 bool; F2 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 uint16 }
func genFn0() string {
    var v7 uint8 = 128
    v7 += 127
    fmt.Println(v7)
    fmt.Println(v7)
    _ = v7
    return (("a" + "foo bar") + "foo bar")
}
func genFn1(p0 uint64, p1 uint) bool {
    var v8 float32 = ((1e10 + 0.0) + 3.14159)
    switch 1 {
    case 128:
        if true {
            var v9 int16 = 32767
            _ = v9
        }
    case 4294967295:
        v10 := GenS1{F0: 0}
        _ = v10
    default:
        v11 := []int{32767}
        _ = v11
    }
    if (-2147483648 < 969) {
        fmt.Println(p1, v8)
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint8 = 150
        v14 := []int{}
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v8
    return (true && true)
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 0
    v3 := make(map[uint32]int)
    var v4 int64 = 42
    var v5 uint16 = 868
    var v6 int16 = 35
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            v5 += v5
            _ = fi16
        }
        _ = fi15
    }
    var v17 uint32 = 2147483647
    fmt.Println(v4)
    if false {
        fmt.Println(v6, v17)
        v18 := GenS1{F0: (v4 ^ -1)}
        _ = v18
    } else {
        v19 := GenS2{F0: v17}
        v20 := []int{2147483647, -2147483648}
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 int16 = -1
        v23 := GenS3{F0: 707, F1: true}
        if (1 > 42) {
            for fi24 := 0; fi24 < 4; fi24++ {
                switch "" {
                case "hello":
                    if true {
                        v25 := GenS4{F0: "foo bar"}
                        var v26 int = (fi21 ^ 128)
                        v27 := []int{9223372036854775807, fi21}
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        v17 = uint32(v2)
                        v6 -= (-(v6))
                        v17 = (v17 - 2147483647)
                    }
                case "\n\t":
                    if (!true) {
                        switch (fi21 * fi24) {
                        case -327:
                            v22 = v6
                        case -2147483648:
                            v28 := GenS5{F0: -1, F1: (false || false), F2: 0}
                            _ = v28
                        }
                        switch "a" {
                        case "foo bar":
                            v5 = (v5 / ((v5) | 1))
                        case "\n\t":
                            fmt.Println(v17, v4)
                        case "":
                            switch -1 {
                            case 32768:
                                switch ((("\n\t" + (("\n\t" + "\n\t") + "a")) + ("hello" + "\n\t")) + ("hello" + "")) {
                                case "a":
                                    v4 += v4
                                case "foo bar":
                                    var v29 int8 = 0
                                    _ = v29
                                case "":
                                    var v30 uint = 255
                                    _ = v30
                                default:
                                    switch (^(fi21)) {
                                    case 0:
                                        v31 := GenS4{F0: "a"}
                                        _ = v31
                                    default:
                                        v32 := []int{-1, 128, 42, (fi21 >> ((18446744073709551615) & 63))}
                                        _ = v32
                                    }
                                }
                            }
                        default:
                            var v33 int16 = v6
                            _ = v33
                        }
                        if true {
                            fmt.Println(v17)
                            v22 -= (v6 * -129)
                            delete(v1, fi24)
                        } else {
                            var v34 int8 = -76
                            var v35 bool = ((fi24 + (fi24 - (fi24 - (fi24 & 4294967295)))) >= 448)
                            v1[(fi24 + fi21)] = fi21
                            _ = v34
                            _ = v35
                        }
                    }
                default:
                    fmt.Println(v6)
                }
                v36 := GenS6{F0: (fi24 ^ fi21)}
                if (true && false) {
                    _ = len(v1)
                    v6 += v6
                    var v37 uint8 = (^(v2))
                    _ = v37
                }
                v38 := GenS7{F0: 2}
                _ = fi24
                _ = v36
                _ = v38
            }
            fmt.Println(v22, v22)
            var v39 int32 = 2147483647
            _ = v39
        } else {
            if false {
                v40 := GenS7{F0: 65535}
                if ("\n\t" >= "hello") {
                    v6 = (-(v6))
                    if (!(2 == 42)) {
                        v22 += 42
                    }
                    switch int(v2) {
                    case 42:
                        delete(v1, (fi21 % ((302) | 1)))
                    case -9223372036854775808:
                        v17 -= v17
                    }
                }
                if true {
                    v5 -= (v5 % ((v5) | 1))
                }
                _ = v40
            }
            var v41 uint8 = v2
            v42 := GenS3{F0: v5, F1: true}
            _ = v41
            _ = v42
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi43 := 0; fi43 < 11; fi43++ {
        if true {
            fmt.Println(v6)
            var v44 bool = false
            v17 -= v17
            _ = v44
        } else {
            fmt.Println(v2)
        }
        v4 = (v4 >> ((454) & 63))
        _ = fi43
    }
    v45 := GenS8{F0: (v6 * (v6 & v6)), F1: (v5 >> ((409) & 63))}
    var v46 uint16 = (v5 ^ 0)
    v47 := GenS7{F0: 42}
    fmt.Println(v2, v2, v46)
    var v48 uint = 65535
    v5 = 2
    if false {
        var v49 int16 = 1000
        var v50 uint8 = 127
        var v51 uint = (v48 >> (((v48 | (v48 % (((v48 * (v48 + (v48 - (v48 + v48))))) | 1)))) & 63))
        _ = v49
        _ = v50
        _ = v51
    }
    if (v48 != v48) {
        v48 = (v48 % (((v48 << (((v48 ^ 18446744073709551615)) & 63))) | 1))
    } else {
        if false {
            switch "hello" {
            case "\n\t":
                v52 := GenS2{F0: (^(v17))}
                _ = v52
            }
            var v53 bool = ((v17 ^ v17) >= v17)
            _ = v53
        }
        if true {
            fmt.Println(v2, v46, v48)
        } else {
            v54 := []int{32767}
            _ = v54
        }
    }
    v48 = (v48 << (((v48 / ((v48) | 1))) & 63))
    fmt.Println(v2, v5, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
