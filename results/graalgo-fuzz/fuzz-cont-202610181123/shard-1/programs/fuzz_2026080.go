package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
func genFn0() uint32 {
    v4 := GenS1{F0: 127, F1: 2}
    v5 := GenS1{F0: -128, F1: 42}
    var v6 int = -771
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int32, p1 uint64, p2 int8) int64 {
    var v7 int16 = -1
    v8 := GenS1{F0: 1, F1: 127}
    if false {
        v9 := []int{}
        var v10 string = "hello"
        if (p1 == (p1 ^ 24)) {
            var v11 uint64 = 350
            if len(v9) > 0 {
                v9[0] = 0
            }
            v10 = v10
            _ = v11
        } else {
            var v12 uint = 510
            v12 += 1
            _ = v12
        }
        _ = v10
        _ = v9
    }
    p0 += 2147483647
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint64 = 950
    var v3 int32 = -1
    switch -866 {
    case 9223372036854775807:
        v13 := []int{587, 32768, -1}
        _ = v13
    case 1:
        var v14 uint = 466
        _ = v14
    case 993:
        var v15 int16 = -1
        _ = v15
    default:
        v2 = 42
    }
    if (!(("\n\t" >= ("a" + ("" + "a"))) || (3.14159 >= (3.14159 / 1.0)))) {
        var v16 float32 = (1.5 * 0.1)
        var v17 float32 = v16
        v18 := GenS2{F0: 195}
        _ = v16
        _ = v17
        _ = v18
    }
    if (("hello" + "foo bar") < "a") {
        v1 = 255
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := GenS2{F0: 2}
            switch fi19 {
            case 9223372036854775807:
                var v21 int64 = genFn1(v3, 193, -128)
                _ = v21
            }
            _ = fi19
            _ = v20
        }
    } else {
        var v22 uint8 = 62
        var v23 uint16 = 571
        var v24 uint64 = v2
        _ = v22
        _ = v23
        _ = v24
    }
    if ((1e10 == 1.5) && true) {
        var v25 bool = (127 <= 135)
        _ = v25
    } else {
        var v26 int16 = 1000
        _ = v26
    }
    var v27 float64 = 0.1
    v1 = uint(v2)
    v2 = v2
    v28 := []int{7}
    v29 := GenS2{F0: 55}
    var v30 string = ("a" + ("" + "a"))
    if (false || (!true)) {
        var v31 string = v30
        v27 += float64(1e10)
        for fi32 := 0; fi32 < 12; fi32++ {
            var v33 string = v31
            var v34 int32 = -1
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v31
    }
    if (v30 < v30) {
        if ((!(("" == ("" + v30)) && (true && (float32(v27) <= (0.0 - ((0.0 + 1.5) * float32(v27))))))) || false) {
            var v35 int16 = 939
            _ = v35
        } else {
            v27 += 1.5
            v36 := GenS3{F0: 432}
            _ = v36
        }
    } else {
        if false {
            if len(v28) > 0 {
                v28[0] = 32767
            }
        } else {
            v37 := GenS1{F0: 127, F1: 2}
            _ = v37
        }
        switch v30 {
        case "":
            for fi38 := 0; fi38 < 17; fi38++ {
                var v39 int8 = -1
                v28 = append(v28, fi38)
                var v40 uint16 = 255
                _ = fi38
                _ = v39
                _ = v40
            }
        case "a":
            fmt.Println(v30, v1, v3)
        default:
            if false {
                v41 := GenS1{F0: -1, F1: 255}
                var v42 int64 = genFn1((v3 ^ v3), (v2 >> (((v1 & (v1 << (((v1 % ((255) | 1))) & 63)))) & 63)), 42)
                for fi43 := 0; fi43 < 7; fi43++ {
                    v27 += float64(3.14159)
                    if false {
                        var v44 int8 = 120
                        for fi45 := 0; fi45 < 11; fi45++ {
                            var v46 float32 = 3.14159
                            _ = fi45
                            _ = v46
                        }
                        _ = v44
                    }
                    _ = fi43
                }
                _ = v41
                _ = v42
            }
        }
    }
    v27 = float64(-2.25)
    var v47 uint16 = 689
    var v48 uint16 = v47
    fmt.Println(v1)
    if (v1 == (v1 - 827)) {
        _ = len(v28)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v47
    _ = v48
}
