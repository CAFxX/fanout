package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int }
func genFn0(p0 string, p1 uint16) uint {
    var v6 int16 = 500
    v7 := []int{}
    fmt.Println(v6, p1, p0)
    if ((true && (2 != 0)) || (-2147483648 != 42)) {
        switch -28 {
        case 1:
            var v8 string = "a"
            _ = v8
        case 4294967295:
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := GenS1{F0: false, F1: 4294967295}
                p0 = p0
                v6 = 293
                _ = fi9
                _ = v10
            }
        case 703:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        default:
            p0 = (p0 + "")
        }
        for fi11 := 0; fi11 < 19; fi11++ {
            switch p0 {
            case "":
                v6 += 1
            case "foo bar":
                for fi12 := 0; fi12 < 20; fi12++ {
                    v13 := GenS2{F0: int(v6)}
                    fmt.Println(v6)
                    _ = fi12
                    _ = v13
                }
            default:
                v14 := GenS2{F0: (fi11 | fi11)}
                _ = v14
            }
            var v15 int64 = 9223372036854775807
            _ = fi11
            _ = v15
        }
        switch 1 {
        case -149:
            p1 = 1
        case 475:
            p0 = p0
        case -736:
            v16 := GenS2{F0: int(v6)}
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 780
}
func genFn1(p0 int32, p1 uint) int64 {
    var v17 bool = (false || ((false || (-181 == -1)) && false))
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS1{F0: v17, F1: (^(p1))}
        switch (fi18 ^ int(p1)) {
        case 32768:
            v20 := []int{(fi18 * -1), (fi18 << ((27) & 63))}
            _ = v20
        case 1:
            v21 := []int{fi18, (fi18 ^ 0)}
            _ = v21
        case 32767:
            if v17 {
                switch (fi18 << (((p1 * genFn0("foo bar", 1))) & 63)) {
                case 9223372036854775807:
                    v22 := GenS2{F0: (fi18 - (fi18 << (((p1 % (((^(p1))) | 1))) & 63)))}
                    _ = v22
                }
                if ((v17 && v17) || (!(false && (614 <= 0)))) {
                    switch fi18 {
                    case 0:
                        for fi23 := 0; fi23 < 20; fi23++ {
                            var v24 int16 = 1
                            if true {
                                var v25 float32 = 0.1
                                fmt.Println(p0, fi23)
                                _ = v25
                            }
                            v26 := []int{fi23}
                            _ = fi23
                            _ = v24
                            _ = v26
                        }
                    }
                    fmt.Println(v17, v17, v17)
                    var v27 uint = genFn0("\n\t", 42)
                    _ = v27
                }
                p1 += 298
            } else {
                var v28 float32 = 3.14159
                _ = v28
            }
        default:
            v29 := GenS1{F0: false, F1: p1}
            _ = v29
        }
        v30 := []int{fi18}
        _ = fi18
        _ = v19
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v17
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = (!true)
    v3 := make(map[uint32]int)
    v4 := []string{"hello", ""}
    v5 := []string{"\n\t"}
    v31 := GenS2{F0: 32768}
    var v32 int16 = -32768
    var v33 int = 195
    var v34 int = (v33 % ((-1) | 1))
    fmt.Println(v34, v33, v34)
    switch (v34 / (((v33 & v34)) | 1)) {
    case 1:
        fmt.Println(v32, v33, v33)
    }
    fmt.Println(v32)
    switch (("a" + ("" + (("" + ("" + "a")) + "a"))) + "") {
    case "a":
        var v35 float64 = 1e10
        _ = v35
    case "\n\t":
        fmt.Println(v32, v32, v2)
    case "":
        if true {
            for fi36 := 0; fi36 < 12; fi36++ {
                var v37 uint16 = 1
                _ = fi36
                _ = v37
            }
        } else {
            v34 = int(v32)
        }
    default:
        var v38 int16 = 42
        _ = v38
    }
    switch "hello" {
    case "a":
        v39 := GenS1{F0: (!false), F1: 65535}
        _ = v39
    case "\n\t":
        if false {
            v34 -= (v34 + v34)
            fmt.Println(v2)
        } else {
            fmt.Println(v32)
        }
    case "":
        var v40 uint8 = 42
        _ = v40
    }
    if ((1 <= 255) || (2 < 42)) {
        var v41 uint64 = 18446744073709551615
        fmt.Println(v2, v32)
        _ = v41
    } else {
        v34 += (v34 * 1)
    }
    for fi42 := 0; fi42 < 12; fi42++ {
        var v43 float32 = (((1e10 + -2.25) / 1.5) * 3.14159)
        var v44 uint64 = 497
        _ = fi42
        _ = v43
        _ = v44
    }
    if (!(265 == 1)) {
        if (v32 <= 42) {
            v2 = (2 < 1)
        }
        v3[273] = (v33 % (((v33 + v34)) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
