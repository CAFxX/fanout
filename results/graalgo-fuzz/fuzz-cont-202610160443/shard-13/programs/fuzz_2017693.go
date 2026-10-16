package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 string; F2 uint64 }
func genFn0(p0 int) int {
    var v4 uint32 = 2147483647
    if true {
        switch "foo bar" {
        case "hello":
            for fi5 := 0; fi5 < 6; fi5++ {
                v6 := []int{0}
                var v7 string = (("" + "foo bar") + (("hello" + "a") + ("foo bar" + "foo bar")))
                _ = fi5
                _ = v6
                _ = v7
            }
        default:
            v8 := GenS1{F0: -1}
            _ = v8
        }
        for fi9 := 0; fi9 < 5; fi9++ {
            p0 = p0
            v10 := []int{(^(p0))}
            var v11 int32 = -115
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    _ = p0
    _ = v4
    return (p0 ^ p0)
}
func genFn1(p0 int8, p1 uint64) uint16 {
    var v12 float64 = 1.5
    v13 := GenS2{F0: false}
    var v14 uint16 = 65535
    switch "hello" {
    case "foo bar":
        var v15 float32 = 1e10
        _ = v15
    case "":
        var v16 int16 = 32767
        _ = v16
    default:
        p1 = 0
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 uint8 = 255
        fmt.Println(p0, p1, v18)
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    return v14
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (float64(float32((((float64(0.0) * float64(float32(3.14159))) / (float64(1e10) * 3.14159)) - 1.5))) / 1.5)
    var v3 uint8 = 2
    fmt.Println(v2, v3)
    v19 := GenS2{F0: false}
    v3 = 215
    v2 = v2
    v20 := []int{}
    if len(v20) > 0 {
        v20[0] = 128
    }
    _ = len(v20)
    var v21 int16 = -614
    switch "hello" {
    case "":
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
    case "\n\t":
        var v22 string = "hello"
        _ = v22
    case "foo bar":
        if ((false || (100 <= 0)) || true) {
            var v23 uint16 = 334
            switch 1 {
            case 127:
                v24 := GenS1{F0: v21}
                _ = v24
            case 42:
                var v25 uint64 = 1
                _ = v25
            default:
                if true {
                    fmt.Println(v3, v21)
                }
            }
            _ = v23
        } else {
            if ((-(v21)) >= (v21 ^ (v21 ^ v21))) {
                var v26 uint16 = 65535
                var v27 int8 = 42
                v28 := GenS3{F0: 944, F1: -119}
                _ = v26
                _ = v27
                _ = v28
            } else {
                if false {
                    v29 := GenS4{F0: -780}
                    v30 := GenS4{F0: -9223372036854775808}
                    if (!true) {
                        v31 := GenS1{F0: (v21 + v21)}
                        _ = v31
                    } else {
                        v3 -= (v3 >> ((18446744073709551615) & 63))
                        v21 = -905
                    }
                    _ = v29
                    _ = v30
                }
                v2 -= v2
            }
            for fi32 := 0; fi32 < 16; fi32++ {
                v21 = -32768
                v21 = (v21 % ((v21) | 1))
                _ = fi32
            }
        }
    default:
        var v33 int32 = int32(v3)
        _ = v33
    }
    var v34 float32 = 0.1
    if (false || true) {
        v21 = v21
        var v35 uint8 = v3
        _ = v35
    } else {
        var v36 uint16 = 32767
        switch 32767 {
        case -9223372036854775808:
            for fi37 := 0; fi37 < 5; fi37++ {
                for fi38 := 0; fi38 < 10; fi38++ {
                    if (2147483648 < uint32(v36)) {
                        v36 = (^(v36))
                        v1[fi38] = "hello"
                        var v39 uint = 396
                        _ = v39
                    } else {
                        var v40 string = ""
                        switch (v40 + "hello") {
                        case "a":
                            v41 := GenS2{F0: ((v3 / ((v3) | 1)) < 105)}
                            _ = v41
                        default:
                            v42 := GenS5{F0: 1, F1: (v40 + v40), F2: 1}
                            _ = v42
                        }
                        _ = v40
                    }
                    switch "hello" {
                    case "foo bar":
                        v36 += 255
                    case "a":
                        var v43 float64 = v2
                        switch fi38 {
                        case 32767:
                            v44 := GenS1{F0: v21}
                            _ = v44
                        case 1:
                            var v45 float32 = (v34 - v34)
                            _ = v45
                        case 9223372036854775807:
                            var v46 bool = (!(!true))
                            _ = v46
                        }
                        _ = v43
                    default:
                        fmt.Println(fi38)
                    }
                    switch (fi37 >> ((65535) & 63)) {
                    case -669:
                        v21 = (^(v21))
                    case 42:
                        var v47 uint = 4294967295
                        _ = v47
                    default:
                        var v48 uint8 = (^(v3))
                        _ = v48
                    }
                    _ = fi38
                }
                var v49 int16 = (v21 ^ (v21 ^ (-(v21))))
                _ = fi37
                _ = v49
            }
        }
        _ = v36
    }
    for fi50 := 0; fi50 < 13; fi50++ {
        var v51 float32 = v34
        _ = fi50
        _ = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
}
