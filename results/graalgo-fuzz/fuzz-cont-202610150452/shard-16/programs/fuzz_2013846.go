package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 int32 }
func genFn0(p0 int8) uint16 {
    if (0 >= 2) {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int32 = 0
            _ = fi7
            _ = v8
        }
        fmt.Println(p0)
        v9 := GenS1{F0: (("foo bar" + (("\n\t" + "foo bar") + "a")) < "a"), F1: 0}
        _ = v9
    }
    if (-502 != 42) {
        var v10 uint8 = 2
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint8 = (v10 | v10)
            _ = fi11
            _ = v12
        }
        fmt.Println(p0, p0)
        _ = v10
    } else {
        v13 := []int{}
        var v14 int16 = 1000
        if true {
            if true {
                if len(v13) > 0 {
                    v13[0] = 4294967295
                }
            } else {
                for fi15 := 0; fi15 < 2; fi15++ {
                    if (true || true) {
                        v16 := GenS2{F0: 345, F1: 255}
                        var v17 float64 = 0.1
                        var v18 int = 0
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    _ = fi15
                }
                var v19 int32 = 1
                _ = v19
            }
            p0 = (-(p0))
        } else {
            var v20 string = "foo bar"
            _ = v20
        }
        _ = v13
        _ = v14
    }
    _ = p0
    return 42
}
func genFn1() int32 {
    v21 := []int{}
    v22 := GenS2{F0: 490, F1: 1}
    _ = v21
    _ = v22
    return -1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    var v3 uint = 42
    var v4 int64 = 1
    var v5 uint16 = 472
    var v6 int16 = 32767
    v23 := GenS1{F0: (1 < 42), F1: 2}
    var v24 int32 = -1
    v25 := GenS1{F0: false, F1: 770}
    _ = len(v1)
    v26 := v1[0]
    v27 := GenS2{F0: 1, F1: 127}
    switch -513 {
    case -1:
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 int16 = (v6 + v6)
            var v30 int8 = -115
            switch (("a" + "a") + "hello") {
            case "a":
                v31 := v1[(^(fi28))]
                _ = v31
            default:
                v2["a"] = ""
            }
            _ = fi28
            _ = v29
            _ = v30
        }
    default:
        for fi32 := 0; fi32 < 18; fi32++ {
            v33 := v2[(("hello" + "hello") + "a")]
            _ = fi32
            _ = v33
        }
    }
    if (42 <= 0) {
        var v34 int = -2147483648
        var v35 float64 = 1e10
        _ = v34
        _ = v35
    } else {
        for fi36 := 0; fi36 < 19; fi36++ {
            v37 := GenS3{F0: 0, F1: -16}
            v3 = v3
            for fi38 := 0; fi38 < 15; fi38++ {
                var v39 string = "a"
                if (true || ((-2.25 * (3.14159 - 3.14159)) < 1e10)) {
                    var v40 float32 = float32(3.14159)
                    v24 = v24
                    v5 -= (v5 - 192)
                    _ = v40
                } else {
                    var v41 int32 = (v24 % (((v24 / ((v24) | 1))) | 1))
                    switch v39 {
                    case "hello":
                        v5 += (v5 * (v5 % (((^(v5))) | 1)))
                    }
                    switch 1 {
                    case 127:
                        v42 := []int{fi38, (fi38 / (((fi36 >> ((0) & 63))) | 1)), (fi38 << (((v3 ^ (v3 / (((v3 * v3)) | 1)))) & 63)), fi38}
                        _ = v42
                    }
                    v41 = genFn1()
                    _ = v41
                }
                _ = fi38
                _ = v39
            }
            _ = fi36
            _ = v37
        }
        var v43 uint16 = genFn0(127)
        _ = v43
    }
    for fi44 := 0; fi44 < 7; fi44++ {
        if (!((^(v5)) >= v5)) {
            for fi45 := 0; fi45 < 9; fi45++ {
                var v46 uint8 = 194
                fmt.Println(v46)
                var v47 float32 = -2.25
                _ = fi45
                _ = v46
                _ = v47
            }
            v24 -= (-(v24))
            fmt.Println(v26, v5)
        } else {
            v6 -= (v6 ^ v6)
        }
        v4 = v4
        _ = fi44
    }
    v3 = (v3 * (v3 & v3))
    if (!(v26 && v26)) {
        var v48 float64 = ((-2.25 - (((0.0 / 1e10) / 1.5) + (float64(0.0) * 0.1))) * 0.1)
        _ = v48
    } else {
        if (!(127 < 127)) {
            switch "hello" {
            case "\n\t":
                v49 := GenS4{F0: genFn1()}
                _ = v49
            default:
                v26 = v26
            }
        } else {
            switch -171 {
            case 9223372036854775807:
                var v50 bool = v26
                _ = v50
            case 736:
                var v51 float32 = 1.5
                _ = v51
            case -9223372036854775808:
                v52 := GenS5{F0: 2, F1: (v24 | v24)}
                _ = v52
            }
            var v53 string = "hello"
            v54 := v1[32767]
            _ = v53
            _ = v54
        }
        var v55 float32 = -2.25
        var v56 float64 = -2.25
        _ = v55
        _ = v56
    }
    var v57 float32 = (float32((((float64(3.14159) / 1.0) * float64((-2.25 / float32((3.14159 * float64(0.1)))))) / 0.1)) / 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
}
