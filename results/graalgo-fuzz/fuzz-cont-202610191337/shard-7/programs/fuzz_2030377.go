package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint = 8
    var v2 uint16 = 255
    v3 := make(map[string]int)
    v4 := make(map[int]uint32)
    switch "\n\t" {
    case "foo bar":
        if true {
            v5 := []int{}
            if true {
                for fi6 := 0; fi6 < 20; fi6++ {
                    fmt.Println(v1)
                    v5 = append(v5, fi6)
                    _ = fi6
                }
            }
            var v7 int64 = 1
            _ = v5
            _ = v7
        } else {
            v8 := []int{4294967295, -9223372036854775808}
            _ = v8
        }
    case "a":
        var v9 bool = false
        _ = v9
    default:
        var v10 uint = (^(v1))
        _ = v10
    }
    var v11 int16 = -972
    _ = len(v3)
    switch 667 {
    case -1:
        var v12 uint64 = 42
        _ = v12
    }
    var v13 int16 = 42
    v14 := GenS1{F0: "a", F1: 1}
    var v15 uint = uint(v2)
    for fi16 := 0; fi16 < 14; fi16++ {
        if (!(4294967295 < 2147483648)) {
            var v17 uint8 = 127
            switch "hello" {
            case "foo bar":
                v15 = (v1 << ((65535) & 63))
            default:
                v18 := v3["foo bar"]
                _ = v18
            }
            v19 := []int{fi16, 2147483647}
            _ = v17
            _ = v19
        }
        v20 := GenS2{F0: 400}
        switch "foo bar" {
        case "a":
            switch (fi16 | (fi16 + fi16)) {
            case 128:
                v21 := GenS3{F0: 255, F1: v2}
                _ = v21
            case 4294967295:
                v15 = (v15 & v1)
            case -40:
                switch fi16 {
                case 2147483647:
                    switch -664 {
                    case -2147483648:
                        var v22 uint16 = (v2 & v2)
                        _ = v22
                    }
                case 511:
                    var v23 int8 = 127
                    _ = v23
                default:
                    var v24 int32 = 2147483647
                    _ = v24
                }
            default:
                v25 := []int{(fi16 >> ((42) & 63)), (-(fi16)), 42}
                _ = v25
            }
        case "hello":
            v13 = (v13 * v13)
        case "\n\t":
            for fi26 := 0; fi26 < 2; fi26++ {
                if true {
                    fmt.Println(v1, v15, v15)
                    switch fi16 {
                    case 127:
                        var v27 uint16 = (v2 >> (((v1 << ((v15) & 63))) & 63))
                        _ = v27
                    case -2147483648:
                        v1 -= v1
                    case 42:
                        v15 += v1
                    default:
                        if ((v13 <= -104) || (318 != 0)) {
                            var v28 bool = ((false || (-960 == -1)) && true)
                            var v29 float32 = 0.0
                            v30 := []int{}
                            var v31 int = fi26
                            _ = v28
                            _ = v29
                            _ = v30
                            _ = v31
                        } else {
                            v32 := []int{fi26}
                            var v33 uint = v15
                            v15 = 0
                            _ = v32
                            _ = v33
                        }
                    }
                }
                v34 := GenS1{F0: "hello", F1: 0}
                v35 := GenS4{F0: 0}
                _ = fi26
                _ = v34
                _ = v35
            }
        }
        _ = fi16
        _ = v20
    }
    var v36 int64 = 379
    for fi37 := 0; fi37 < 20; fi37++ {
        switch ("a" + "") {
        case "\n\t":
            delete(v3, "hello")
        case "a":
            var v38 uint32 = 892
            _ = v38
        case "foo bar":
            if false {
                var v39 uint16 = 255
                for fi40 := 0; fi40 < 1; fi40++ {
                    var v41 string = ("hello" + ("foo bar" + "hello"))
                    var v42 float64 = 0.1
                    v43 := []int{(fi37 / (((-(fi40))) | 1)), fi37}
                    fmt.Println(v15)
                    v41 = "foo bar"
                    var v44 bool = true
                    if ((float64(3.14159) * (0.1 + (((1e10 * ((v42 * v42) - (v42 * 1e10))) - 0.1) - 1.5))) >= v42) {
                        v3[(v41 + v41)] = (fi37 + 32768)
                        var v45 uint16 = v2
                        var v46 uint = (v15 + 1)
                        _ = v45
                        _ = v46
                    }
                    if (v1 > 2) {
                        var v47 int32 = -314
                        v42 = (1e10 * (0.0 * (v42 / ((v42 / v42) - (((v42 + v42) * v42) / -2.25)))))
                        fmt.Println(v36, fi37)
                        _ = v47
                    } else {
                        var v48 uint32 = 2147483647
                        if v44 {
                            v15 = v1
                            if ((fi40 % ((fi40) | 1)) != (fi37 % ((fi37) | 1))) {
                                fmt.Println(fi37, v13)
                                var v49 int = fi37
                                _ = len(v43)
                                _ = v49
                            } else {
                                fmt.Println(v36, v13, v36)
                            }
                        }
                        var v50 uint32 = v48
                        _ = v48
                        _ = v50
                    }
                    _ = fi40
                    _ = v41
                    _ = v42
                    _ = v43
                    _ = v44
                }
                _ = v39
            } else {
                v51 := []int{(fi37 & fi37), fi37, -214}
                _ = v51
            }
        default:
            _ = len(v3)
        }
        v52 := []int{9223372036854775807}
        _ = fi37
        _ = v52
    }
    v53 := GenS2{F0: (v15 >> (((v15 << (((^(v15))) & 63))) & 63))}
    v15 = (v1 * uint(v2))
    if true {
        v54 := []int{}
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v53
}
