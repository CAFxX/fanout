package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint }
type GenS3 struct { F0 string; F1 int; F2 bool }
type GenS4 struct { F0 int16; F1 int8; F2 uint8 }
func genFn0() string {
    if false {
        var v6 float32 = -2.25
        var v7 int16 = 23
        _ = v6
        _ = v7
    }
    var v8 string = ""
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := []int{(-(fi9))}
        _ = fi9
        _ = v10
    }
    v8 = ("hello" + (v8 + "hello"))
    if true {
        for fi11 := 0; fi11 < 19; fi11++ {
            for fi12 := 0; fi12 < 3; fi12++ {
                if false {
                    fmt.Println(v8, fi12, fi12)
                    var v13 uint8 = 186
                    var v14 int32 = 0
                    _ = v13
                    _ = v14
                }
                v8 = (v8 + "\n\t")
                _ = fi12
            }
            for fi15 := 0; fi15 < 19; fi15++ {
                fmt.Println(fi11, fi15, v8)
                v16 := GenS1{F0: 1, F1: 1}
                v17 := GenS2{F0: 1, F1: 0, F2: 2}
                _ = fi15
                _ = v16
                _ = v17
            }
            for fi18 := 0; fi18 < 5; fi18++ {
                v19 := []int{fi18, (fi11 << ((2) & 63))}
                fmt.Println(fi11, v8, fi11)
                var v20 uint64 = 4294967295
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi11
        }
        switch ((v8 + v8) + (((v8 + v8) + v8) + ("a" + v8))) {
        case "\n\t":
            v8 = (v8 + "")
        default:
            if (v8 == v8) {
                if ((0 <= 9223372036854775807) && false) {
                    v8 = "foo bar"
                    if (false || (32767 >= 0)) {
                        v21 := GenS2{F0: 1, F1: 0, F2: 244}
                        if true {
                            v8 = ("\n\t" + v8)
                        }
                        var v22 uint16 = 882
                        _ = v21
                        _ = v22
                    }
                }
                for fi23 := 0; fi23 < 4; fi23++ {
                    v8 = ("hello" + (v8 + "\n\t"))
                    if false {
                        var v24 int64 = 639
                        _ = v24
                    }
                    _ = fi23
                }
            }
        }
    }
    _ = v8
    return v8
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 608
    var v3 uint32 = 2147483647
    v4 := []uint32{(v3 & (v3 - v3)), (v3 >> ((18446744073709551615) & 63))}
    var v5 int16 = -443
    var v25 bool = true
    if (((42 > 0) || v25) || v25) {
        var v26 uint32 = v3
        switch 9223372036854775807 {
        case 767:
            v2 = v2
        case 42:
            v5 -= v5
        }
        _ = v26
    }
    var v27 bool = (!v25)
    switch "a" {
    case "\n\t":
        var v28 bool = true
        _ = v28
    case "a":
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 int16 = v5
            v30 = 0
            if v25 {
                if (-77 <= 0) {
                    var v31 float64 = ((float64(-2.25) * 3.14159) + -2.25)
                    v32 := GenS1{F0: 255, F1: 100}
                    v5 += (v30 ^ -1)
                    _ = v31
                    _ = v32
                } else {
                    v4 = append(v4, 4294967295)
                    var v33 int8 = 100
                    _ = v33
                }
                v4 = append(v4, 4294967295)
                fmt.Println(v27, v3)
            } else {
                v34 := GenS2{F0: uint16(v3), F1: 42, F2: 42}
                fmt.Println(v30, v3)
                v35 := GenS3{F0: "foo bar", F1: (-(fi29)), F2: v27}
                _ = v34
                _ = v35
            }
            _ = fi29
            _ = v30
        }
    }
    v5 -= v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 -= v3
    if (127 < 42) {
        _ = len(v1)
        if (v3 == (v3 * (v3 & (v3 - v3)))) {
            var v36 int32 = (v2 + -2147483648)
            _ = v36
        }
    } else {
        var v37 int64 = 0
        _ = v37
    }
    switch ("a" + "") {
    case "":
        v38 := v1["a"]
        _ = v38
    case "a":
        v39 := GenS2{F0: 2, F1: 9, F2: 1}
        _ = v39
    default:
        if ((!(0 <= 100)) || false) {
            switch "" {
            case "a":
                if v27 {
                    var v40 float64 = (((0.0 - 1.5) + (((((0.0 - -2.25) + 3.14159) * 0.1) / 1e10) * 0.1)) * (-2.25 - (0.0 - 0.0)))
                    v41 := GenS1{F0: 42, F1: -1}
                    _ = v40
                    _ = v41
                }
            case "foo bar":
                if (-1 >= (v2 >> ((uint(v3)) & 63))) {
                    var v42 float64 = 0.1
                    _ = v42
                } else {
                    var v43 float32 = 1.5
                    _ = v43
                }
            default:
                v44 := GenS1{F0: 42, F1: 127}
                _ = v44
            }
        }
    }
    var v45 uint64 = 2
    for fi46 := 0; fi46 < 11; fi46++ {
        v27 = (v5 != v5)
        var v47 uint16 = 65535
        _ = fi46
        _ = v47
    }
    v27 = true
    var v48 int = -666
    delete(v1, "\n\t")
    fmt.Println(v45)
    for fi49 := 0; fi49 < 10; fi49++ {
        var v50 int32 = -2147483648
        var v51 uint32 = (v3 | 65535)
        v5 = (v5 % ((1) | 1))
        _ = fi49
        _ = v50
        _ = v51
    }
    v52 := GenS4{F0: (-(v5)), F1: 0, F2: 1}
    var v53 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v5
    _ = v52
    _ = v53
}
