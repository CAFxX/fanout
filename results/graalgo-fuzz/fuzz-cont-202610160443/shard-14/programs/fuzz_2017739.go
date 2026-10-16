package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int64; F1 int16; F2 string }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 string, p1 uint8, p2 int8) int32 {
    p2 = (p2 >> ((2) & 63))
    fmt.Println(p0, p1, p2)
    v7 := GenS1{F0: "hello", F1: p0}
    v8 := GenS1{F0: p0, F1: (("" + p0) + ("foo bar" + p0))}
    v9 := []int{-2147483648, 2147483647, -9223372036854775808, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint16, p1 uint, p2 int32) uint32 {
    p1 = p1
    if false {
        if false {
            if true {
                p2 = (p2 & int32(p1))
                var v10 uint = p1
                v11 := GenS1{F0: ("hello" + "\n\t"), F1: "hello"}
                _ = v10
                _ = v11
            } else {
                v12 := GenS2{F0: int64(p0), F1: 0, F2: "foo bar"}
                if ((p1 == (p1 / (((p1 ^ (p1 ^ 255))) | 1))) || (p1 <= uint(p0))) {
                    if false {
                        var v13 int16 = -533
                        var v14 int = 32767
                        _ = v13
                        _ = v14
                    } else {
                        var v15 bool = true
                        if (42 < -9223372036854775808) {
                            v16 := []int{128, 127, 316, 127}
                            _ = v16
                        }
                        _ = v15
                    }
                } else {
                    v17 := []int{127, 454, 127, -663}
                    _ = v17
                }
                _ = v12
            }
            v18 := []int{813, 128, 959}
            _ = v18
        }
        v19 := []int{321, 2147483647}
        _ = v19
    } else {
        v20 := []int{}
        v21 := GenS1{F0: "", F1: ""}
        _ = v20
        _ = v21
    }
    var v22 uint16 = (p0 | (p0 ^ p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 65535
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    v3 := []uint32{155}
    var v4 float32 = 0.0
    var v5 int32 = -1
    var v6 uint64 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v23 uint8 = 88
    if true {
        for fi24 := 0; fi24 < 3; fi24++ {
            v4 += (v4 + v4)
            _ = fi24
        }
        var v25 int = 4294967295
        _ = v25
    }
    v26 := GenS2{F0: -1, F1: 0, F2: ("" + ("foo bar" + "a"))}
    if len(v3) > 0 {
        v3[0] = 424
    }
    v23 -= 23
    _ = len(v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v27 uint = 255
    var v28 uint = v27
    v29 := v2[362]
    if false {
        v30 := GenS3{F0: v27, F1: -1}
        _ = v30
    }
    switch "\n\t" {
    case "foo bar":
        if (1 >= 1) {
            for fi31 := 0; fi31 < 19; fi31++ {
                fmt.Println(v5, fi31)
                _ = len(v3)
                v5 += 2147483647
                _ = fi31
            }
            v6 += 260
        } else {
            if true {
                if false {
                    v28 = uint(v6)
                    if (true || ((!(2 != 1)) && ((v23 - 96) != v23))) {
                        v23 -= v23
                        var v32 string = "a"
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 9; fi33++ {
                        var v34 int = 0
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi33
                        _ = v34
                    }
                } else {
                    v35 := v1[42]
                    for fi36 := 0; fi36 < 18; fi36++ {
                        v4 -= (1e10 + ((float32(0.1) * 0.1) - (float32(-2.25) * 3.14159)))
                        _ = fi36
                    }
                    v35 = v35
                    _ = v35
                }
                v37 := GenS2{F0: -631, F1: -600, F2: "hello"}
                _ = v37
            } else {
                v38 := v2[0]
                v39 := v2[258]
                _ = v38
                _ = v39
            }
            v5 = -1
            v40 := GenS4{F0: 2147483648}
            _ = v40
        }
    case "hello":
        v27 = 0
    case "\n\t":
        fmt.Println(v27, v27)
    }
    if (-258 > -1) {
        v28 -= 65535
        if true {
            v2[317] = v29
            switch int(v23) {
            case 856:
                var v41 uint16 = 255
                _ = v41
            default:
                v42 := GenS4{F0: 357}
                _ = v42
            }
            v43 := GenS1{F0: (("a" + "\n\t") + "foo bar"), F1: "a"}
            _ = v43
        }
    } else {
        v44 := GenS3{F0: 18446744073709551615, F1: 141}
        if false {
            var v45 int16 = -1
            var v46 int = (v29 % ((int(v5)) | 1))
            if true {
                var v47 int = 686
                _ = v47
            }
            _ = v45
            _ = v46
        } else {
            v48 := v1[genFn1(123, v27, (v5 ^ (v5 << ((v28) & 63))))]
            _ = v48
        }
        _ = v44
    }
    for fi49 := 0; fi49 < 12; fi49++ {
        for fi50 := 0; fi50 < 16; fi50++ {
            switch ("foo bar" + "a") {
            case "":
                var v51 uint8 = v23
                _ = v51
            case "hello":
                v52 := GenS1{F0: ("foo bar" + "\n\t"), F1: "a"}
                _ = v52
            }
            if len(v3) > 0 {
                v3[0] = 0
            }
            _ = fi50
        }
        if false {
            for fi53 := 0; fi53 < 4; fi53++ {
                var v54 int8 = -95
                var v55 float32 = v4
                fmt.Println(v27, v4)
                _ = fi53
                _ = v54
                _ = v55
            }
        } else {
            for fi56 := 0; fi56 < 4; fi56++ {
                var v57 uint64 = v6
                switch (-(fi56)) {
                case 4294967295:
                    fmt.Println(v4, v29, v57)
                }
                _ = fi56
                _ = v57
            }
            var v58 int64 = 42
            _ = v58
        }
        _ = fi49
    }
    for fi59 := 0; fi59 < 4; fi59++ {
        v3 = append(v3, 1)
        _ = fi59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
