package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    var v3 uint32 = 2147483647
    var v4 bool = (25 < uint8(v3))
    v4 = v4
    switch 32767 {
    case 202:
        v5 := GenS1{F0: 22, F1: 0, F2: -625}
        _ = v5
    case -1:
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 float64 = 3.14159
            var v8 int64 = 42
            _ = fi6
            _ = v7
            _ = v8
        }
    case -9223372036854775808:
        if (72 > -1) {
            if v4 {
                v9 := v1[("hello" + ("foo bar" + "\n\t"))]
                var v10 int = -2147483648
                _ = v10
                _ = v9
            }
        } else {
            for fi11 := 0; fi11 < 15; fi11++ {
                v12 := GenS2{F0: (fi11 - fi11)}
                fmt.Println(fi11, fi11, v4)
                var v13 uint16 = 42
                _ = fi11
                _ = v12
                _ = v13
            }
            if (!((0 < 9223372036854775807) || true)) {
                v3 -= 112
                for fi14 := 0; fi14 < 8; fi14++ {
                    fmt.Println(v3, fi14)
                    _ = fi14
                }
            } else {
                if true {
                    switch ("foo bar" + "foo bar") {
                    case "":
                        var v15 float32 = -2.25
                        _ = v15
                    default:
                        if v4 {
                            v16 := GenS1{F0: -2147483648, F1: 964, F2: 1}
                            _ = v16
                        }
                    }
                    var v17 float32 = (0.0 / 1.0)
                    v3 = (v3 << ((255) & 63))
                    _ = v17
                }
            }
        }
    default:
        if (v4 || (4294967295 <= 894)) {
            if false {
                v18 := GenS3{F0: -1}
                if (false && v4) {
                    switch "\n\t" {
                    case "hello":
                        delete(v1, "hello")
                    }
                    if true {
                        fmt.Println(v3, v4)
                    } else {
                        v19 := []int{-74}
                        v4 = ((float64(3.14159) < 1e10) || ((!false) || v4))
                        v4 = false
                        _ = v19
                    }
                }
                v3 -= v3
                _ = v18
            } else {
                var v20 float64 = -2.25
                v21 := GenS4{F0: -1}
                _ = v20
                _ = v21
            }
            if (!v4) {
                for fi22 := 0; fi22 < 14; fi22++ {
                    switch ((((("a" + "hello") + "hello") + "foo bar") + (("foo bar" + "a") + "foo bar")) + ("a" + "hello")) {
                    case "hello":
                        v4 = ((((3.14159 + (0.0 / (1e10 * 3.14159))) + 1.5) * (1e10 / (3.14159 + 1e10))) < float64(-2.25))
                    default:
                        v4 = (!v4)
                    }
                    v3 = v3
                    _ = fi22
                }
                for fi23 := 0; fi23 < 2; fi23++ {
                    v3 = 319
                    if ("hello" >= ("hello" + (("a" + "\n\t") + "\n\t"))) {
                        var v24 float64 = 1e10
                        v4 = v4
                        _ = v24
                    } else {
                        if false {
                            var v25 uint = 4294967295
                            v26 := v2["a"]
                            v27 := []int{-2147483648, (v26 * 32768), fi23, -1}
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        v28 := GenS3{F0: -9223372036854775808}
                        _ = v28
                    }
                    _ = fi23
                }
            } else {
                v29 := []int{1, 2147483647}
                _ = v29
            }
        }
    }
    v4 = (-893 < 4294967295)
    v30 := []int{}
    v3 -= (v3 / (((^(v3))) | 1))
    v31 := GenS2{F0: 1}
    switch "" {
    case "":
        if (-1 >= 1000) {
            var v32 int8 = -124
            _ = v32
        }
    default:
        switch "" {
        case "foo bar":
            v4 = false
        default:
            if v4 {
                v33 := GenS4{F0: 0}
                switch "" {
                case "a":
                    var v34 int64 = 1
                    _ = v34
                case "":
                    v30 = append(v30, -1)
                default:
                    var v35 uint16 = 255
                    _ = v35
                }
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
                _ = v33
            }
        }
    }
    v30 = append(v30, 127)
    if ((!(!true)) || false) {
        for fi36 := 0; fi36 < 9; fi36++ {
            var v37 uint32 = v3
            _ = fi36
            _ = v37
        }
    }
    switch 9223372036854775807 {
    case 800:
        var v38 int32 = 42
        _ = v38
    case 39:
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
    case 32768:
        fmt.Println(v3, v3)
    default:
        v4 = v4
    }
    if (127 == 1) {
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
        for fi39 := 0; fi39 < 1; fi39++ {
            v40 := GenS3{F0: 42}
            _ = fi39
            _ = v40
        }
        for fi41 := 0; fi41 < 7; fi41++ {
            if v4 {
                v3 += 2147483648
                v3 = 0
                for fi42 := 0; fi42 < 7; fi42++ {
                    if v4 {
                        v4 = v4
                    }
                    var v43 bool = false
                    fmt.Println(v4, v43)
                    if len(v30) >= 2 {
                        v30 = v30[1:len(v30)]
                    }
                    switch fi41 {
                    case 32767:
                        fmt.Println(v43, v43, v3)
                    case 42:
                        v43 = ("\n\t" >= "")
                    }
                    _ = fi42
                    _ = v43
                }
            }
            for fi44 := 0; fi44 < 12; fi44++ {
                var v45 int8 = 101
                if ((((v3 <= v3) || (0 > 255)) && (!(!(false && (((((float64(((1.5 / 1.0) - (3.14159 / 1.5))) / 1.0) / 1.0) - 1e10) - (-2.25 - -2.25)) != float64(-2.25)))))) || false) {
                    var v46 int32 = 1
                    var v47 string = ("\n\t" + ("hello" + (("foo bar" + ("foo bar" + "\n\t")) + "\n\t")))
                    _ = v46
                    _ = v47
                }
                _ = fi44
                _ = v45
            }
            var v48 int32 = 30
            _ = fi41
            _ = v48
        }
    }
    v49 := GenS5{F0: (0 == 4294967295), F1: 32767}
    for fi50 := 0; fi50 < 6; fi50++ {
        v51 := GenS2{F0: fi50}
        var v52 int = fi50
        _ = fi50
        _ = v51
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v49
}
