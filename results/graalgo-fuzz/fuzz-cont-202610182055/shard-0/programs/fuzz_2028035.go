package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int8; F1 int; F2 uint16 }
type GenS3 struct { F0 uint8; F1 string; F2 int16 }
type GenS4 struct { F0 int }
func genFn0(p0 int8) uint {
    var v6 int8 = (p0 & p0)
    v7 := GenS1{F0: 866, F1: (!false)}
    _ = p0
    _ = v6
    _ = v7
    return 99
}
func genFn1(p0 int64) string {
    for fi8 := 0; fi8 < 10; fi8++ {
        p0 = (p0 << ((0) & 63))
        _ = fi8
    }
    p0 += p0
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(fi9, fi9)
        v10 := []int{(^(fi9)), -9223372036854775808, fi9, fi9}
        _ = fi9
        _ = v10
    }
    _ = p0
    return "\n\t"
}

func main() {
    var v1 float32 = 1.5
    var v2 uint = 4294967295
    var v3 uint64 = 2
    var v4 int32 = 0
    var v5 bool = (((v2 | (v2 & 587)) > v2) && true)
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(fi11, v1, fi11)
        for fi12 := 0; fi12 < 6; fi12++ {
            switch fi12 {
            case 32768:
                var v13 int8 = -128
                _ = v13
            }
            _ = fi12
        }
        _ = fi11
    }
    v3 = v3
    for fi14 := 0; fi14 < 14; fi14++ {
        v5 = v5
        fmt.Println(v1, v5)
        fmt.Println(v4, v5, v4)
        _ = fi14
    }
    if ((v4 & v4) < (v4 >> (((v2 ^ (v2 | (v2 + v2)))) & 63))) {
        v1 = v1
    }
    if ((2 >= 1) && true) {
        v15 := GenS1{F0: (v2 + (v2 ^ (v2 - v2))), F1: ("hello" >= ("\n\t" + ("a" + "a")))}
        _ = v15
    } else {
        if v5 {
            v3 = (v3 ^ 120)
            var v16 uint64 = v3
            if v5 {
                var v17 int64 = 9223372036854775807
                v18 := GenS2{F0: 42, F1: int(v2), F2: 255}
                var v19 int64 = v17
                _ = v17
                _ = v18
                _ = v19
            } else {
                var v20 int = -554
                switch ("" + ("a" + ("" + "hello"))) {
                case "":
                    v1 += -2.25
                case "foo bar":
                    v21 := GenS1{F0: genFn0(42), F1: false}
                    _ = v21
                case "a":
                    var v22 uint = v2
                    _ = v22
                default:
                    v5 = false
                }
                _ = v20
            }
            _ = v16
        }
        var v23 bool = (!v5)
        var v24 uint = 352
        _ = v23
        _ = v24
    }
    var v25 int32 = (-(v4))
    var v26 int32 = v25
    v2 -= (^(v2))
    var v27 int32 = (v25 / ((-971) | 1))
    var v28 float32 = float32(1e10)
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 int8 = 8
        if true {
            for fi31 := 0; fi31 < 9; fi31++ {
                if v5 {
                    if (42 >= 65535) {
                        v32 := []int{fi29}
                        _ = v32
                    } else {
                        if (255 < 42) {
                            v28 = v1
                        }
                        v33 := []int{fi29, fi29}
                        _ = v33
                    }
                }
                var v34 uint64 = v3
                _ = fi31
                _ = v34
            }
        } else {
            v35 := GenS1{F0: (v2 | (v2 ^ (v2 & (^(v2))))), F1: (v5 && ((v5 || true) || (true && true)))}
            switch (fi29 - -930) {
            case 127:
                if v5 {
                    fmt.Println(v3, v30)
                }
            case 42:
                if v5 {
                    v36 := []int{(fi29 | (^(fi29))), (fi29 % (((fi29 + 291)) | 1)), fi29, (fi29 ^ (-(fi29)))}
                    _ = v36
                }
            case 0:
                if (!((!v5) || (1 > 446))) {
                    v1 = 0.0
                    fmt.Println(v3, v27)
                    v5 = (v28 > ((v28 * 3.14159) * (0.1 - -2.25)))
                }
            default:
                var v37 uint64 = (v3 * (v3 * 4294967295))
                _ = v37
            }
            v4 += (v25 << (((v2 + v2)) & 63))
            _ = v35
        }
        var v38 bool = ((4294967295 == (v2 + genFn0((-(v30))))) || v5)
        _ = fi29
        _ = v30
        _ = v38
    }
    if true {
        v39 := GenS1{F0: (v2 / (((v2 | v2)) | 1)), F1: (3.14159 > 1e10)}
        var v40 uint8 = 3
        _ = v39
        _ = v40
    } else {
        v41 := GenS3{F0: 169, F1: "hello", F2: -1}
        fmt.Println(v27, v1)
        _ = v41
    }
    v42 := []int{-722}
    var v43 int16 = -512
    if false {
        var v44 int64 = -1
        v4 = 1
        _ = v44
    } else {
        if (-2.25 < (0.0 * ((3.14159 - 0.1) * 1.5))) {
            v45 := GenS4{F0: 32767}
            v42 = append(v42, 1)
            _ = v45
        }
        fmt.Println(v27, v5)
        var v46 float64 = (-2.25 / ((3.14159 + (-2.25 - (0.0 + 0.1))) + 3.14159))
        _ = v46
    }
    if v5 {
        var v47 string = ("" + ("\n\t" + "\n\t"))
        v4 = (v27 | -608)
        var v48 int16 = v43
        _ = v47
        _ = v48
    }
    var v49 int8 = 100
    var v50 uint64 = (v3 | 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v49
    _ = v5
    _ = v50
}
