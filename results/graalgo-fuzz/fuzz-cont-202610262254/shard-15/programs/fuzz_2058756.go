package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint32; F2 bool }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 uint64; F1 int32 }
func genFn0(p0 int64) uint {
    for fi7 := 0; fi7 < 12; fi7++ {
        switch fi7 {
        case 141:
            v8 := GenS1{F0: 1, F1: 42}
            _ = v8
        case 75:
            p0 -= (p0 | 9223372036854775807)
        }
        _ = fi7
    }
    if false {
        switch 42 {
        case -485:
            v9 := GenS1{F0: 248, F1: 255}
            _ = v9
        }
    }
    _ = p0
    return 255
}
func genFn1(p0 int64) bool {
    var v10 string = "hello"
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS1{F0: 65535, F1: 334}
        var v13 int16 = 1000
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 19, F1: 65535}
    _ = p0
    _ = v10
    _ = v14
    return false
}

func main() {
    var v1 float32 = 3.14159
    var v2 float32 = (1e10 + (v1 - ((0.0 + float32(-2.25)) + v1)))
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    var v5 uint16 = 0
    v6 := []int{890, 42, -9223372036854775808, 0}
    var v15 float64 = (-2.25 * (0.0 + 1.5))
    switch "foo bar" {
    case "":
        for fi16 := 0; fi16 < 3; fi16++ {
            _ = len(v3)
            v17 := GenS1{F0: v5, F1: v5}
            var v18 uint32 = 65535
            _ = fi16
            _ = v17
            _ = v18
        }
    case "\n\t":
        var v19 int = 32767
        _ = v19
    default:
        var v20 bool = (-506 > -699)
        _ = v20
    }
    var v21 float64 = 0.1
    switch 32767 {
    case 32767:
        switch "" {
        case "":
            var v22 float64 = 1.5
            _ = v22
        }
    case 980:
        v3["foo bar"] = 0
    case 9223372036854775807:
        var v23 int8 = 127
        _ = v23
    }
    switch "" {
    case "a":
        delete(v3, "foo bar")
    case "":
        v24 := GenS1{F0: 842, F1: (v5 / ((v5) | 1))}
        _ = v24
    case "hello":
        switch ("foo bar" + "hello") {
        case "":
            _ = len(v3)
        case "hello":
            v25 := GenS2{F0: 2147483647, F1: 255, F2: 100}
            _ = v25
        }
    default:
        v15 = v21
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 float32 = 0.1
        _ = fi26
        _ = v27
    }
    fmt.Println(v15, v15, v2)
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 bool = (1.5 >= v15)
        var v30 int16 = 32767
        v3[""] = (fi28 + fi28)
        _ = fi28
        _ = v29
        _ = v30
    }
    switch "\n\t" {
    case "foo bar":
        if true {
            switch "a" {
            case "foo bar":
                var v31 int16 = 42
                _ = v31
            case "":
                v2 += float32(1e10)
            case "\n\t":
                switch "" {
                case "foo bar":
                    var v32 float32 = -2.25
                    _ = v32
                default:
                    var v33 uint32 = 2147483648
                    _ = v33
                }
            }
        }
    case "a":
        for fi34 := 0; fi34 < 15; fi34++ {
            v35 := GenS3{F0: 44, F1: 0, F2: true}
            _ = fi34
            _ = v35
        }
    default:
        if len(v6) > 0 {
            v6[0] = 128
        }
    }
    var v36 int32 = -1
    v21 = ((float64(1.5) / v15) + (v21 + float64((v2 / (v1 * 1.5)))))
    for fi37 := 0; fi37 < 20; fi37++ {
        v15 += 0.1
        _ = fi37
    }
    if false {
        if (18446744073709551615 != 18446744073709551615) {
            switch 32768 {
            case 32768:
                var v38 float32 = v1
                _ = v38
            case 0:
                _ = len(v6)
            }
            var v39 float64 = -2.25
            switch -1 {
            case 32768:
                for fi40 := 0; fi40 < 11; fi40++ {
                    fmt.Println(v15, v15)
                    _ = fi40
                }
            case 42:
                v41 := GenS3{F0: 255, F1: 65535, F2: false}
                _ = v41
            }
            _ = v39
        } else {
            if (42 < -1) {
                if ((v36 % ((152) | 1)) > -263) {
                    for fi42 := 0; fi42 < 13; fi42++ {
                        v43 := v4[fi42]
                        if v43 {
                            v1 += v1
                            v43 = ("a" != "")
                            v6 = append(v6, (fi42 << ((4294967295) & 63)))
                        }
                        var v44 uint = genFn0(-1)
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                } else {
                    var v45 uint = 845
                    _ = v45
                }
                var v46 float64 = (v21 - v15)
                _ = v46
            }
            v47 := GenS1{F0: (v5 + 0), F1: (v5 & (v5 | (v5 << ((42) & 63))))}
            _ = v47
        }
        var v48 int = 387
        _ = v48
    } else {
        if (true && true) {
            if len(v6) > 0 {
                v6[0] = 0
            }
            for fi49 := 0; fi49 < 7; fi49++ {
                v21 = float64((float32(v21) - 3.14159))
                fmt.Println(v36, v2, v5)
                switch (fi49 / (((fi49 + -124)) | 1)) {
                case 32768:
                    v50 := GenS4{F0: 73, F1: "hello"}
                    _ = v50
                case 887:
                    var v51 uint16 = (v5 + (v5 % ((v5) | 1)))
                    _ = v51
                case 731:
                    _ = len(v3)
                }
                _ = fi49
            }
        }
        var v52 bool = (!false)
        v53 := GenS5{F0: 0, F1: v36}
        _ = v52
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
