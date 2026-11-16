package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int64 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint16 }
func genFn0(p0 bool, p1 int64) string {
    v6 := GenS1{F0: (0 != 2), F1: 0, F2: (p1 * (p1 | p1))}
    if ((0 != 42) || p0) {
        var v7 bool = p0
        p1 = (p1 * p1)
        if p0 {
            fmt.Println(p1, p1, p1)
        } else {
            switch 32767 {
            case -363:
                p1 -= p1
            case 42:
                for fi8 := 0; fi8 < 7; fi8++ {
                    for fi9 := 0; fi9 < 7; fi9++ {
                        p1 += 9223372036854775807
                        var v10 int = fi9
                        _ = fi9
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 9; fi11++ {
                        var v12 int = 2147483647
                        _ = fi11
                        _ = v12
                    }
                    _ = fi8
                }
            case 128:
                v13 := GenS2{F0: (42 >= 65535), F1: (p0 || p0)}
                _ = v13
            default:
                v14 := []int{}
                _ = v14
            }
            for fi15 := 0; fi15 < 3; fi15++ {
                p0 = true
                fmt.Println(fi15, p1)
                _ = fi15
            }
            p1 += 9223372036854775807
        }
        _ = v7
    } else {
        p1 = p1
        if (128 < 0) {
            var v16 bool = (4294967295 < 0)
            switch "" {
            case "foo bar":
                if false {
                    var v17 uint32 = 4294967295
                    _ = v17
                }
            case "hello":
                v18 := []int{4294967295, -55, 323}
                _ = v18
            case "a":
                v19 := []int{1, -957, -226, -982}
                _ = v19
            default:
                var v20 float32 = 3.14159
                _ = v20
            }
            _ = v16
        } else {
            var v21 int32 = -721
            _ = v21
        }
    }
    v22 := GenS1{F0: (768 == 1), F1: 127, F2: p1}
    v23 := []int{4294967295}
    v24 := GenS3{F0: 925}
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    _ = v24
    _ = v6
    return "a"
}

func main() {
    v1 := []string{("\n\t" + "a")}
    v2 := make(map[uint32]int)
    var v3 float64 = 3.14159
    var v4 uint8 = 0
    v5 := make(map[string]string)
    v3 = 1e10
    fmt.Println(v4, v4)
    switch -711 {
    case 1:
        v1 = append(v1, ("" + "hello"))
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        v26 := GenS3{F0: 736}
        for fi27 := 0; fi27 < 16; fi27++ {
            v3 = v3
            v3 = float64(((0.0 - 3.14159) * (1.5 * 3.14159)))
            _ = fi27
        }
        var v28 bool = (false || true)
        _ = fi25
        _ = v26
        _ = v28
    }
    fmt.Println(v3, v4)
    _ = len(v5)
    if false {
        v29 := v2[42]
        delete(v5, "hello")
        fmt.Println(v3, v3, v4)
        _ = v29
    } else {
        var v30 int32 = -1
        var v31 float64 = v3
        v1 = append(v1, ("foo bar" + ("foo bar" + "hello")))
        _ = v30
        _ = v31
    }
    fmt.Println(v3)
    var v32 int32 = -568
    v33 := GenS4{F0: 42, F1: (!false)}
    var v34 int8 = -128
    switch "\n\t" {
    case "hello":
        var v35 float64 = (v3 * 1.5)
        _ = v35
    case "\n\t":
        v36 := GenS4{F0: 32767, F1: (580 < 787)}
        _ = v36
    case "a":
        if (float64(1.5) >= ((float64(1.5) * v3) * -2.25)) {
            switch int(v4) {
            case 32767:
                var v37 int32 = v32
                _ = v37
            case 0:
                switch ("\n\t" + ("hello" + ("foo bar" + ("hello" + "foo bar")))) {
                case "a":
                    var v38 string = ("hello" + "hello")
                    _ = v38
                case "foo bar":
                    v39 := GenS1{F0: (0 < 65535), F1: (v4 ^ 0), F2: -1}
                    _ = v39
                }
            case -738:
                v34 += (v34 << ((255) & 63))
            default:
                var v40 uint64 = 839
                _ = v40
            }
        }
    }
    for fi41 := 0; fi41 < 16; fi41++ {
        var v42 uint64 = 4294967295
        fmt.Println(v42, v42)
        var v43 int8 = (v34 >> ((111) & 63))
        _ = fi41
        _ = v42
        _ = v43
    }
    v34 = v34
    if (1 >= 42) {
        var v44 string = ""
        _ = len(v2)
        _ = v44
    } else {
        v34 = v34
    }
    v45 := GenS5{F0: 523}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v45
    _ = v5
}
