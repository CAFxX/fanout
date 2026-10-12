package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
func genFn0() uint8 {
    v7 := GenS1{F0: 1, F1: (((((0.0 * 3.14159) / 0.1) * float32(float64(3.14159))) >= -2.25) || (!("\n\t" == "a")))}
    var v8 int32 = -2147483648
    var v9 float64 = 1e10
    _ = v7
    _ = v8
    _ = v9
    return 127
}
func genFn1(p0 int8, p1 uint8, p2 uint32) int32 {
    if false {
        p2 += p2
        var v10 float64 = 0.0
        var v11 string = "\n\t"
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 3; fi12++ {
            if ((p0 % ((p0) | 1)) >= (-(p0))) {
                v13 := GenS2{F0: -148}
                _ = v13
            }
            v14 := []int{fi12, fi12, fi12, 476}
            p1 = (p1 & 174)
            _ = fi12
            _ = v14
        }
        p2 = p2
        var v15 uint8 = (p1 >> ((319) & 63))
        _ = v15
    }
    if (2147483647 > 0) {
        p1 -= p1
        var v16 float64 = float64((0.0 / (0.0 + (0.1 + (0.1 - 1e10)))))
        _ = v16
    } else {
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 float64 = 0.1
    v2 := []int{9223372036854775807}
    var v3 float64 = v1
    var v4 uint8 = 142
    v5 := []string{("" + ("" + "a"))}
    var v6 float64 = float64(-2.25)
    var v17 float64 = ((((0.1 - float64(3.14159)) * float64((0.1 / float32(v6)))) / 1.5) / (float64(3.14159) - (v1 + (float64(((3.14159 * -2.25) / (1.5 - 0.1))) / (float64(1e10) / v1)))))
    var v18 int16 = 438
    for fi19 := 0; fi19 < 12; fi19++ {
        if len(v5) > 0 {
            v5[0] = ("\n\t" + ("\n\t" + ("" + (("" + "foo bar") + "\n\t"))))
        }
        _ = fi19
    }
    if (12 >= -92) {
        fmt.Println(v17)
        v20 := GenS2{F0: 776}
        _ = v20
    } else {
        var v21 float64 = 1.5
        var v22 float32 = float32((v1 / v6))
        _ = v21
        _ = v22
    }
    fmt.Println(v18)
    if (((false || (((float64(0.0) - (v6 * 1.5)) * (v1 / v6)) > ((((v6 / v3) * 1e10) + v3) + float64(3.14159)))) && true) && (2147483648 > 873)) {
        var v23 int64 = -73
        switch 127 {
        case -42:
            if true {
                v24 := GenS1{F0: -1, F1: true}
                var v25 int = 9223372036854775807
                _ = v24
                _ = v25
            } else {
                v1 += v17
                v26 := GenS1{F0: -120, F1: true}
                _ = v26
            }
        case 721:
            v1 += v3
        case 42:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        switch "" {
        case "":
            v6 -= v1
        case "hello":
            fmt.Println(v6, v23, v23)
        }
        _ = v23
    } else {
        var v27 string = ""
        var v28 uint32 = 977
        for fi29 := 0; fi29 < 12; fi29++ {
            v4 = 2
            var v30 int16 = (v18 << ((790) & 63))
            switch ((v27 + v27) + v27) {
            case "":
                v17 = v3
            default:
                if false {
                    switch ("\n\t" + v27) {
                    case "hello":
                        v31 := GenS3{F0: -128}
                        _ = v31
                    case "a":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "foo bar":
                        for fi32 := 0; fi32 < 18; fi32++ {
                            v30 -= v18
                            _ = fi32
                        }
                    }
                } else {
                    var v33 int64 = 1
                    var v34 int64 = v33
                    _ = v33
                    _ = v34
                }
            }
            _ = fi29
            _ = v30
        }
        _ = v27
        _ = v28
    }
    v6 -= v17
    v18 = v18
    if ("foo bar" != "a") {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v35 := GenS1{F0: -21, F1: true}
        _ = v35
    }
    fmt.Println(v6, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
