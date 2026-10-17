package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 string; F2 string }
func genFn0(p0 int8, p1 uint64, p2 uint) uint64 {
    v6 := GenS1{F0: 42, F1: 926}
    if (!false) {
        for fi7 := 0; fi7 < 14; fi7++ {
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := []int{}
                _ = len(v9)
                var v10 uint32 = 2
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    } else {
        var v11 int16 = 42
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p1 & (p1 << ((p2) & 63)))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 float64 = 1e10
    var v4 uint32 = 4294967295
    var v5 uint = 18446744073709551615
    v12 := GenS2{F0: 42}
    v5 = 870
    v13 := GenS3{F0: 32767}
    var v14 uint64 = 2
    var v15 int32 = -2147483648
    var v16 string = ("foo bar" + ("foo bar" + ((("a" + ("hello" + "")) + "foo bar") + ("\n\t" + "foo bar"))))
    fmt.Println(v15, v5, v15)
    switch (v16 + "foo bar") {
    case "\n\t":
        if false {
            for fi17 := 0; fi17 < 10; fi17++ {
                if (((-32768 > 732) || true) || true) {
                    v18 := []int{(fi17 * (fi17 + (fi17 + fi17))), -372, fi17}
                    if false {
                        v16 = ""
                        if (!(9223372036854775807 > 9223372036854775807)) {
                            _ = len(v18)
                        } else {
                            var v19 uint8 = 127
                            v20 := GenS3{F0: 255}
                            var v21 float32 = 3.14159
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        switch v16 {
                        case "a":
                            v18 = append(v18, fi17)
                        case "hello":
                            v22 := GenS2{F0: -110}
                            _ = v22
                        case "foo bar":
                            v23 := GenS2{F0: -724}
                            _ = v23
                        }
                    } else {
                        var v24 int16 = 1000
                        var v25 int64 = -1
                        _ = v24
                        _ = v25
                    }
                    _ = v18
                } else {
                    for fi26 := 0; fi26 < 9; fi26++ {
                        var v27 uint8 = 72
                        var v28 uint8 = (^(v27))
                        v14 = v14
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    var v29 uint16 = 1
                    v30 := GenS1{F0: 42, F1: -134}
                    _ = v29
                    _ = v30
                }
                fmt.Println(fi17, v16, v16)
                v5 = 42
                _ = fi17
            }
            switch (((v16 + v16) + "hello") + "a") {
            case "\n\t":
                fmt.Println(v15)
            default:
                v15 = (v15 * (v15 | (^(v15))))
            }
            if true {
                var v31 bool = (false || (0 > v15))
                for fi32 := 0; fi32 < 11; fi32++ {
                    if ((1000 <= 1) || true) {
                        for fi33 := 0; fi33 < 20; fi33++ {
                            v34 := []int{fi33}
                            v4 = (v4 ^ 65535)
                            _ = fi33
                            _ = v34
                        }
                    }
                    v35 := GenS1{F0: 0, F1: -9223372036854775808}
                    _ = fi32
                    _ = v35
                }
                _ = v31
            }
        } else {
            var v36 float64 = (((((-2.25 + v3) - 1.5) * (float64(0.0) - 1.5)) * -2.25) + v3)
            v14 = (v14 | genFn0(1, 42, 0))
            var v37 uint32 = v4
            _ = v36
            _ = v37
        }
    case "":
        v4 += (^(v4))
    default:
        v38 := GenS3{F0: 65535}
        _ = v38
    }
    for fi39 := 0; fi39 < 10; fi39++ {
        v5 -= (v5 ^ (v5 >> ((v5) & 63)))
        var v40 int8 = 80
        _ = fi39
        _ = v40
    }
    fmt.Println(v3, v16, v16)
    v4 -= v4
    for fi41 := 0; fi41 < 15; fi41++ {
        for fi42 := 0; fi42 < 6; fi42++ {
            fmt.Println(v14, v3)
            switch fi41 {
            case -2147483648:
                switch (fi42 << ((v5) & 63)) {
                case -670:
                    v43 := []int{}
                    _ = v43
                case 519:
                    v44 := GenS3{F0: 255}
                    _ = v44
                case 0:
                    v45 := []int{fi41}
                    _ = v45
                default:
                    v46 := GenS4{F0: (v14 - v14), F1: v16, F2: v16}
                    _ = v46
                }
            default:
                switch (-(fi42)) {
                case -9223372036854775808:
                    v5 -= (v5 & 42)
                }
            }
            _ = fi42
        }
        _ = fi41
    }
    v47 := []int{-1, 0, 32768}
    var v48 uint32 = 120
    if false {
        if false {
            v5 = (^(v5))
            v49 := GenS2{F0: 301}
            var v50 int64 = 1
            _ = v49
            _ = v50
        } else {
            var v51 uint64 = 905
            _ = v51
        }
        var v52 string = ("" + "\n\t")
        _ = v52
    }
    var v53 int16 = -1
    v54 := GenS2{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v53
    _ = v54
}
