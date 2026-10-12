package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int32, p1 int16, p2 bool) uint32 {
    var v5 uint64 = 704
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: fi6, F1: 1, F2: ("a" + "hello")}
        var v8 float64 = (0.0 / 1.0)
        fmt.Println(p0)
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p2, p0)
    switch 575 {
    case 9223372036854775807:
        var v9 uint16 = 2
        _ = v9
    case -417:
        if (0.1 != (0.0 - (0.0 * ((3.14159 + -2.25) + (float64(-2.25) - float64(1.5)))))) {
            v10 := GenS1{F0: 32768, F1: 1, F2: "foo bar"}
            v11 := GenS1{F0: 128, F1: 42, F2: "foo bar"}
            _ = v10
            _ = v11
        } else {
            if false {
                v5 = (v5 >> ((65535) & 63))
            } else {
                if (-128 != 1) {
                    v12 := GenS1{F0: 1, F1: 871, F2: "\n\t"}
                    v13 := []int{32767, -805, -1}
                    _ = v12
                    _ = v13
                }
                fmt.Println(v5, p1, p1)
                var v14 uint = 182
                _ = v14
            }
            if false {
                var v15 float32 = (0.1 / (float32(1.5) / 0.1))
                v16 := GenS2{F0: 9223372036854775807}
                fmt.Println(v15, v5, p2)
                _ = v15
                _ = v16
            }
            v17 := []int{}
            _ = v17
        }
    default:
        var v18 int16 = 0
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 int16 = p1
        p2 = true
        switch fi19 {
        case -9223372036854775808:
            var v21 uint16 = 42
            _ = v21
        case -686:
            fmt.Println(p2, p0, v5)
        case 32768:
            var v22 float64 = -2.25
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    var v1 string = "foo bar"
    var v2 string = "foo bar"
    v3 := []int{-802, 4294967295, -2147483648}
    var v4 uint8 = 255
    var v23 int = -2147483648
    fmt.Println(v23)
    fmt.Println(v4, v4)
    v24 := GenS2{F0: 1}
    var v25 uint8 = (v4 % ((v4) | 1))
    _ = len(v3)
    v26 := GenS2{F0: 9223372036854775807}
    v1 = (v1 + "\n\t")
    if false {
        var v27 int8 = 42
        _ = v27
    } else {
        switch ("foo bar" + v1) {
        case "hello":
            for fi28 := 0; fi28 < 20; fi28++ {
                for fi29 := 0; fi29 < 15; fi29++ {
                    var v30 uint32 = 400
                    _ = fi29
                    _ = v30
                }
                _ = fi28
            }
        case "\n\t":
            v4 = v25
        default:
            for fi31 := 0; fi31 < 19; fi31++ {
                if true {
                    if ((true || ((-2.25 <= (3.14159 + (-2.25 + 3.14159))) && true)) && true) {
                        for fi32 := 0; fi32 < 4; fi32++ {
                            switch fi31 {
                            case 128:
                                var v33 int32 = 863
                                _ = v33
                            case -219:
                                v2 = ("foo bar" + v2)
                            case -1:
                                var v34 uint8 = (v25 ^ v25)
                                _ = v34
                            }
                            v23 -= fi32
                            _ = fi32
                        }
                        v1 = ((("\n\t" + v2) + v1) + ("" + (v1 + v2)))
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        for fi35 := 0; fi35 < 2; fi35++ {
                            var v36 int8 = 100
                            v4 += (^(v25))
                            _ = fi35
                            _ = v36
                        }
                        v37 := GenS3{F0: 2, F1: 32767}
                        _ = v37
                    }
                    v38 := GenS1{F0: v23, F1: 42, F2: ("a" + (v1 + ((v2 + v2) + ("foo bar" + ((v2 + v2) + ("foo bar" + v1))))))}
                    v25 -= v25
                    _ = v38
                }
                v39 := GenS4{F0: 42}
                fmt.Println(v1, v23)
                _ = fi31
                _ = v39
            }
        }
        var v40 string = (v1 + "a")
        if (false && false) {
            if false {
                v41 := GenS3{F0: 106, F1: 32767}
                if (1 == 4294967295) {
                    var v42 uint16 = 1
                    v4 = (v4 * v4)
                    _ = v42
                }
                var v43 int8 = 100
                _ = v41
                _ = v43
            }
            v23 += (v23 * 117)
        }
        _ = v40
    }
    v23 -= 755
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = "foo bar"
    v4 += (v4 - (v25 * 0))
    if (-2.25 == -2.25) {
        _ = len(v3)
        switch "\n\t" {
        case "a":
            var v44 string = v2
            _ = v44
        case "foo bar":
            v3 = append(v3, v23)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
