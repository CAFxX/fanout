package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int }
func genFn0() uint8 {
    if true {
        var v4 uint32 = 0
        v5 := []int{2147483647}
        v6 := GenS1{F0: 9223372036854775807}
        _ = v4
        _ = v5
        _ = v6
    } else {
        if ((2 < 199) || (false && (1 < 4294967295))) {
            for fi7 := 0; fi7 < 15; fi7++ {
                v8 := []int{}
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 float64 = (-2.25 + (0.0 - 1.5))
                v11 := GenS1{F0: -495}
                _ = fi9
                _ = v10
                _ = v11
            }
            switch -1 {
            case -2147483648:
                var v12 float32 = 1e10
                _ = v12
            }
        }
    }
    var v13 uint64 = 195
    fmt.Println(v13)
    v13 += v13
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := []int{fi14, 2147483647}
        for fi16 := 0; fi16 < 7; fi16++ {
            v13 = 0
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    _ = v13
    return 65
}

func main() {
    v1 := []uint32{0, 0, 65535}
    v2 := []int{-1}
    v3 := make(map[string]bool)
    var v17 int8 = 1
    if true {
        if true {
            if false {
                switch 32768 {
                case -1:
                    switch 42 {
                    case 93:
                        for fi18 := 0; fi18 < 19; fi18++ {
                            switch "hello" {
                            case "":
                                v17 = (v17 % ((0) | 1))
                            case "a":
                                v17 += (v17 ^ v17)
                            case "hello":
                                for fi19 := 0; fi19 < 17; fi19++ {
                                    var v20 string = "\n\t"
                                    v21 := GenS1{F0: -980}
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                            }
                            fmt.Println(fi18, v17)
                            _ = fi18
                        }
                    case 32767:
                        v22 := GenS1{F0: 184}
                        _ = v22
                    }
                }
            } else {
                var v23 uint16 = 32767
                var v24 bool = false
                v1 = append(v1, 65535)
                _ = v23
                _ = v24
            }
            fmt.Println(v17)
            switch ("a" + "foo bar") {
            case "\n\t":
                for fi25 := 0; fi25 < 18; fi25++ {
                    switch "a" {
                    case "hello":
                        if false {
                            var v26 bool = false
                            _ = v26
                        } else {
                            var v27 uint64 = 383
                            if false {
                                var v28 uint16 = 42
                                _ = v28
                            }
                            var v29 int8 = v17
                            _ = v27
                            _ = v29
                        }
                    case "":
                        v17 = (v17 << ((322) & 63))
                    case "a":
                        switch "hello" {
                        case "\n\t":
                            var v30 int16 = -32768
                            _ = v30
                        }
                    }
                    v17 -= (v17 + v17)
                    _ = fi25
                }
            case "a":
                fmt.Println(v17)
            }
        }
        v17 = -1
        v17 = (v17 % ((v17) | 1))
    }
    switch -217 {
    case -9223372036854775808:
        v17 = 100
    default:
        if (genFn0() != 1) {
            _ = len(v1)
        } else {
            if false {
                v17 = (v17 + -128)
            } else {
                var v31 uint32 = 2147483648
                _ = v31
            }
            v32 := GenS1{F0: 42}
            var v33 uint32 = 432
            _ = v32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        fmt.Println(v17, fi34)
        switch fi34 {
        case -2147483648:
            if (4294967295 == 4294967295) {
                var v35 int64 = -915
                switch 1 {
                case 0:
                    fmt.Println(v35)
                }
                switch (^(fi34)) {
                case 2147483647:
                    fmt.Println(v35)
                }
                _ = v35
            }
        case -9223372036854775808:
            v36 := GenS2{F0: 505, F1: ("hello" + "")}
            _ = v36
        }
        _ = fi34
    }
    v17 -= (v17 * (^(v17)))
    for fi37 := 0; fi37 < 16; fi37++ {
        v17 -= 0
        for fi38 := 0; fi38 < 3; fi38++ {
            v17 = v17
            fmt.Println(fi37, fi38, v17)
            _ = fi38
        }
        v17 -= v17
        _ = fi37
    }
    fmt.Println(v17)
    switch "\n\t" {
    case "":
        var v39 float32 = 0.1
        _ = v39
    case "foo bar":
        if true {
            v40 := v3[("foo bar" + "\n\t")]
            v41 := GenS3{F0: -348}
            _ = v40
            _ = v41
        }
    case "hello":
        v42 := GenS3{F0: -9223372036854775808}
        _ = v42
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi43 := 0; fi43 < 19; fi43++ {
        v44 := GenS2{F0: 2, F1: ""}
        fmt.Println(fi43)
        var v45 uint32 = 0
        _ = fi43
        _ = v44
        _ = v45
    }
    var v46 bool = true
    v46 = false
    v46 = true
    v46 = ("" < ("" + "foo bar"))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v46
}
