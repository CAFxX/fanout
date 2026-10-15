package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint16 }
func genFn0() int {
    v6 := GenS1{F0: "a"}
    if false {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 float32 = -2.25
            _ = fi7
            _ = v8
        }
        v9 := GenS2{F0: 1}
        v10 := []int{-2147483648, -1, 127, -720}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        if (42 == 0) {
            if true {
                switch "hello" {
                case "\n\t":
                    var v12 uint64 = 0
                    _ = v12
                case "hello":
                    v13 := []int{fi11, 9223372036854775807}
                    _ = v13
                case "a":
                    var v14 int16 = 42
                    _ = v14
                default:
                    for fi15 := 0; fi15 < 16; fi15++ {
                        v16 := []int{(fi15 - fi15), fi15}
                        _ = len(v16)
                        _ = fi15
                        _ = v16
                    }
                }
                var v17 bool = true
                if true {
                    v17 = v17
                    v18 := GenS2{F0: 600}
                    fmt.Println(fi11, fi11)
                    _ = v18
                } else {
                    v19 := []int{}
                    for fi20 := 0; fi20 < 14; fi20++ {
                        var v21 uint16 = 1
                        _ = fi20
                        _ = v21
                    }
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    _ = v19
                }
                _ = v17
            }
            v22 := []int{fi11, 638, 32768, fi11}
            v23 := GenS3{F0: 128}
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            fmt.Println(fi11, fi24)
            if (!true) {
                v25 := GenS4{F0: 66}
                var v26 uint16 = 757
                _ = v25
                _ = v26
            }
            switch "" {
            case "a":
                var v27 int16 = 1
                _ = v27
            case "foo bar":
                v28 := GenS3{F0: 255}
                _ = v28
            case "hello":
                fmt.Println(fi24)
            default:
                var v29 int8 = 65
                _ = v29
            }
            _ = fi24
        }
        _ = fi11
    }
    _ = v6
    return 633
}

func main() {
    var v1 int16 = 1
    v2 := []string{}
    var v3 uint32 = 952
    var v4 int32 = -873
    var v5 string = "foo bar"
    v2 = append(v2, (v5 + "foo bar"))
    var v30 uint64 = 1
    v5 = ("foo bar" + (v5 + v5))
    fmt.Println(v30, v30, v30)
    var v31 uint64 = 1
    v1 -= (v1 >> ((96) & 63))
    var v32 uint8 = 190
    v33 := GenS5{F0: (v3 * v3), F1: (v32 ^ v32), F2: (v1 & (v1 & (v1 | v1)))}
    var v34 int64 = 716
    v35 := GenS5{F0: (v3 << ((1) & 63)), F1: (v32 * v32), F2: v1}
    v4 -= v4
    fmt.Println(v3, v3)
    fmt.Println(v1, v4, v4)
    v1 += v1
    if ((0.0 < (1.5 * 0.0)) && (((!(!false)) && true) && false)) {
        var v36 int32 = (v4 / ((-362) | 1))
        if true {
            if true {
                v37 := GenS6{F0: 0}
                var v38 float32 = 0.1
                var v39 int16 = v1
                _ = v37
                _ = v38
                _ = v39
            }
            var v40 uint = 377
            fmt.Println(v32, v40)
            _ = v40
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v36
    } else {
        if ((false || (!(float32((1e10 * (1.5 - (0.0 + 1e10)))) != 0.1))) || (-2.25 <= 1e10)) {
            var v41 int16 = v1
            var v42 int8 = -1
            switch -9223372036854775808 {
            case -423:
                v41 += (^(v41))
            case 472:
                var v43 uint = 888
                _ = v43
            case 374:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = v41
            _ = v42
        } else {
            v32 -= (v32 * (^(v32)))
        }
        if ((!(-1 >= int64(v3))) || (!false)) {
            for fi44 := 0; fi44 < 8; fi44++ {
                var v45 float64 = float64(3.14159)
                v46 := GenS6{F0: 0}
                v1 = -518
                _ = fi44
                _ = v45
                _ = v46
            }
        } else {
            var v47 string = "foo bar"
            v48 := GenS6{F0: 42}
            _ = v47
            _ = v48
        }
        var v49 uint32 = 897
        _ = v49
    }
    for fi50 := 0; fi50 < 6; fi50++ {
        v5 = ((("" + v5) + ("foo bar" + ("a" + v5))) + (v5 + "hello"))
        v51 := GenS6{F0: 255}
        v31 += 2
        _ = fi50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
