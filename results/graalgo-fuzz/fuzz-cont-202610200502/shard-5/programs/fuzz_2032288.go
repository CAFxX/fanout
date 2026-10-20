package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint32 }
func genFn0(p0 int, p1 uint16) int16 {
    p1 = p1
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int = 9223372036854775807
        var v6 uint16 = p1
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint16 = p1
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    var v1 int32 = 2147483647
    var v2 bool = false
    var v3 int32 = v1
    v8 := GenS1{F0: (-(v1))}
    for fi9 := 0; fi9 < 1; fi9++ {
        if (2 < 657) {
            var v10 string = ("foo bar" + ((("hello" + "foo bar") + "") + "a"))
            _ = v10
        }
        _ = fi9
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        v1 = v3
        _ = fi11
    }
    var v12 uint16 = 42
    v3 += 2147483647
    var v13 float32 = 1.5
    v14 := GenS1{F0: v3}
    v15 := GenS1{F0: v3}
    switch 42 {
    case 4294967295:
        if v2 {
            fmt.Println(v1, v1)
            v16 := []int{}
            v3 = (v3 - v3)
            _ = v16
        }
    default:
        var v17 uint8 = 1
        _ = v17
    }
    if true {
        switch "foo bar" {
        case "foo bar":
            v18 := []int{128, 42}
            _ = v18
        default:
            v19 := GenS1{F0: 42}
            _ = v19
        }
        for fi20 := 0; fi20 < 18; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                v3 = (v1 & (v3 << ((607) & 63)))
                _ = fi21
            }
            _ = fi20
        }
    } else {
        var v22 int = 32768
        switch v22 {
        case 128:
            v3 = (v3 * v1)
        case 4294967295:
            fmt.Println(v12, v12, v1)
        }
        if ((^(v12)) <= (v12 - v12)) {
            var v23 int8 = -30
            v1 = -2147483648
            var v24 float32 = ((v13 - (v13 - (((((float32(-2.25) * ((v13 / ((-2.25 / (v13 / float32(0.1))) - ((1.5 + (float32(1.5) * (v13 * (((((float32(float64(1.5)) - v13) + (v13 - 1.5)) / v13) + v13) * v13)))) - (v13 - 1.5)))) + (0.1 + (float32(-2.25) / ((3.14159 / (float32(0.0) + v13)) * v13))))) - v13) * float32(0.1)) + v13) - float32(1.5)))) / v13)
            _ = v23
            _ = v24
        }
        _ = v22
    }
    if v2 {
        var v25 float64 = 3.14159
        if v2 {
            v26 := []int{}
            _ = v26
        }
        _ = v25
    } else {
        var v27 int8 = 1
        var v28 float64 = (((float64(float32(float64((3.14159 + 1e10)))) + 1.5) * ((((1.5 + (1.5 + 0.0)) + (3.14159 - ((((0.0 + float64((1.5 + v13))) * 1.5) + 1e10) - float64((v13 / -2.25))))) / 3.14159) * 3.14159)) * (float64(v13) + 0.1))
        _ = v27
        _ = v28
    }
    v3 = (v1 | v3)
    if v2 {
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 uint = 255
            var v31 uint8 = 1
            _ = fi29
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 7; fi32++ {
            v33 := GenS2{F0: (v1 - (v3 % ((v1) | 1)))}
            var v34 int32 = v1
            switch "foo bar" {
            case "hello":
                v35 := GenS3{F0: -1}
                _ = v35
            case "\n\t":
                if v2 {
                    var v36 string = "foo bar"
                    _ = v36
                } else {
                    switch fi32 {
                    case 1:
                        for fi37 := 0; fi37 < 3; fi37++ {
                            v38 := []int{fi32}
                            fmt.Println(fi37, v3, v1)
                            var v39 uint8 = 204
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    }
                    if (255 > 65535) {
                        for fi40 := 0; fi40 < 18; fi40++ {
                            var v41 uint64 = 4294967295
                            _ = fi40
                            _ = v41
                        }
                    } else {
                        switch (-(fi32)) {
                        case 42:
                            var v42 int16 = -1
                            _ = v42
                        case 170:
                            switch ("foo bar" + "") {
                            case "":
                                var v43 int32 = 2147483647
                                _ = v43
                            case "a":
                                var v44 float64 = 0.1
                                _ = v44
                            }
                        }
                        v34 = int32(fi32)
                        fmt.Println(v3, v1)
                    }
                    v45 := GenS3{F0: 127}
                    _ = v45
                }
            }
            _ = fi32
            _ = v33
            _ = v34
        }
        var v46 string = "foo bar"
        _ = v46
    } else {
        switch "\n\t" {
        case "a":
            v1 += (v3 * (v1 * v3))
        case "foo bar":
            v47 := []int{127, 128, 4294967295}
            _ = v47
        }
    }
    var v48 int = 4294967295
    switch 32768 {
    case 0:
        fmt.Println(v13)
    }
    v49 := GenS4{F0: -1, F1: 446}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v8
}
