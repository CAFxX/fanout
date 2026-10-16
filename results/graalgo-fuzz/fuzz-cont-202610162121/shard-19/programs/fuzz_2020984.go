package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS2 struct { F0 int8; F1 int32; F2 int16 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 uint32 }
func genFn0(p0 int8, p1 string) int8 {
    var v7 uint8 = 1
    p1 = p1
    v8 := []int{702, 9223372036854775807, 4294967295}
    fmt.Println(p0)
    var v9 int = 68
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (p0 * p0)
}
func genFn1(p0 int32, p1 int32, p2 int64) string {
    for fi10 := 0; fi10 < 3; fi10++ {
        if (("\n\t" + "a") <= "hello") {
            switch (fi10 * fi10) {
            case -9223372036854775808:
                var v11 uint = 24
                _ = v11
            case 32768:
                v12 := []int{}
                _ = v12
            case 409:
                v13 := []int{}
                _ = v13
            default:
                switch (fi10 / (((fi10 >> ((1) & 63))) | 1)) {
                case 154:
                    p0 = 1
                }
            }
        } else {
            var v14 uint32 = 166
            p0 -= p0
            _ = v14
        }
        var v15 uint16 = 255
        _ = fi10
        _ = v15
    }
    if false {
        v16 := []int{9223372036854775807}
        p1 += (p0 * (^(p1)))
        fmt.Println(p2, p2, p1)
        _ = v16
    }
    v17 := GenS1{F0: 128, F1: 18446744073709551615, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return "foo bar"
}

func main() {
    var v1 string = "hello"
    v2 := []uint32{2}
    v3 := make(map[uint32]int)
    v4 := []string{}
    v5 := []string{v1, "foo bar", v1}
    v6 := make(map[int]bool)
    var v18 bool = true
    v19 := GenS1{F0: 128, F1: 371, F2: 205}
    _ = len(v3)
    for fi20 := 0; fi20 < 19; fi20++ {
        if len(v4) > 0 {
            v4[0] = v1
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int32 = 42
            v23 := GenS2{F0: -128, F1: v22, F2: -32768}
            if (0.1 > -2.25) {
                var v24 int64 = 9223372036854775807
                v25 := GenS2{F0: 100, F1: 2147483647, F2: 32767}
                _ = v24
                _ = v25
            }
            v22 = (v22 - (v22 - 1))
            v26 := GenS3{F0: int32(fi20), F1: 2}
            var v27 uint16 = 661
            if (0 > 127) {
                if v18 {
                    v1 = ("\n\t" + v1)
                    v27 = (v27 >> ((1) & 63))
                    var v28 uint16 = (v27 - 281)
                    _ = v28
                } else {
                    switch (fi20 / (((fi20 >> ((0) & 63))) | 1)) {
                    case 9223372036854775807:
                        _ = len(v2)
                    case 32768:
                        var v29 uint64 = 596
                        _ = v29
                    case -1:
                        fmt.Println(fi21, v18)
                        v5 = append(v5, "\n\t")
                        if v18 {
                            v30 := GenS4{F0: 2, F1: 296, F2: 2}
                            _ = len(v5)
                            var v31 int8 = -80
                            _ = v30
                            _ = v31
                        } else {
                            if ((-2.25 >= (((1e10 + 1.5) / float32(1e10)) - 0.0)) || (-1 >= (v22 | (v22 << ((0) & 63))))) {
                                if true {
                                    fmt.Println(v27, v22)
                                    switch (-(fi21)) {
                                    case 2147483647:
                                        v27 = (v27 & v27)
                                    }
                                    fmt.Println(v22)
                                }
                                v1 = (("" + v1) + v1)
                                v32 := GenS5{F0: "a", F1: v1}
                                _ = v32
                            }
                            var v33 string = "a"
                            _ = v33
                        }
                    }
                    v1 = ""
                }
                switch v1 {
                case "a":
                    if (!(true && (1.5 > float64(float32(3.14159))))) {
                        if false {
                            if (genFn0(-1, ((((v1 + v1) + v1) + (v1 + "foo bar")) + v1)) == -67) {
                                var v34 uint = 142
                                v35 := GenS4{F0: 559, F1: 4294967295, F2: 255}
                                v4 = append(v4, "a")
                                _ = v34
                                _ = v35
                            } else {
                                v1 = "hello"
                            }
                            var v36 uint8 = 255
                            _ = v36
                        }
                        v2 = append(v2, 2147483648)
                        var v37 float32 = -2.25
                        _ = v37
                    }
                default:
                    var v38 int16 = 1000
                    _ = v38
                }
                v27 += 230
            } else {
                v1 = ((((v1 + "a") + "foo bar") + v1) + "hello")
                if false {
                    if v18 {
                        fmt.Println(v22, fi20, v27)
                    }
                    if false {
                        v39 := GenS6{F0: 255, F1: fi21}
                        v40 := GenS3{F0: 2147483647, F1: 2}
                        if ((v27 / (((^(v27))) | 1)) > v27) {
                            v41 := GenS2{F0: 0, F1: -580, F2: -364}
                            var v42 uint64 = 4294967295
                            _ = v41
                            _ = v42
                        }
                        _ = v39
                        _ = v40
                    }
                    v27 -= (v27 * (v27 ^ 42))
                    fmt.Println(v22)
                }
                v1 = v1
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            v18 = (513 > -1)
            switch (fi21 + fi20) {
            case 32768:
                v27 = v27
            case -1:
                if (!(104 == -128)) {
                    var v43 uint32 = 328
                    v44 := GenS2{F0: 1, F1: v22, F2: -1}
                    fmt.Println(v1, v22, v43)
                    _ = v43
                    _ = v44
                } else {
                    fmt.Println(v18)
                }
            }
            var v45 uint = 18446744073709551615
            _ = fi21
            _ = v22
            _ = v23
            _ = v26
            _ = v27
            _ = v45
        }
        for fi46 := 0; fi46 < 2; fi46++ {
            v47 := GenS7{F0: 0}
            v48 := GenS7{F0: 824}
            var v49 int16 = 0
            var v50 uint64 = 1
            _ = fi46
            _ = v47
            _ = v48
            _ = v49
            _ = v50
        }
        _ = fi20
    }
    var v51 string = v1
    for fi52 := 0; fi52 < 3; fi52++ {
        var v53 uint16 = 42
        v18 = v18
        _ = fi52
        _ = v53
    }
    var v54 int8 = -128
    v18 = (!v18)
    var v55 uint16 = 30
    if true {
        v51 = ("\n\t" + v51)
    }
    switch v51 {
    case "\n\t":
        var v56 int8 = (-(v54))
        _ = v56
    case "foo bar":
        fmt.Println(v18, v1)
    }
    var v57 uint = 18446744073709551615
    for fi58 := 0; fi58 < 11; fi58++ {
        if len(v5) > 0 {
            v5[0] = ((v51 + ((("foo bar" + v1) + "\n\t") + ((v1 + (v1 + v51)) + v1))) + v51)
        }
        switch (v1 + "a") {
        case "\n\t":
            var v59 string = v51
            _ = v59
        }
        _ = fi58
    }
    v54 += v54
    v55 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v54
    _ = v55
    _ = v57
    _ = v6
}
