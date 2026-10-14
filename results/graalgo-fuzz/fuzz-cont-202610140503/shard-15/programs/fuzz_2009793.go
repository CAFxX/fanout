package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint64; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool }
func genFn0(p0 uint, p1 uint32, p2 int8) int8 {
    fmt.Println(p0, p1, p1)
    v6 := []int{32767, int(p0), 0, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -23
}

func main() {
    var v1 int16 = -637
    var v2 int32 = 151
    v3 := make(map[string]uint32)
    var v4 int8 = 100
    v5 := []string{"a", "a", "hello", "a"}
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: ""}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 uint32 = 223
    switch 2147483647 {
    case 2147483647:
        switch 9223372036854775807 {
        case 32767:
            if false {
                v11 := GenS2{F0: 383}
                if (("" >= "hello") || true) {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        _ = len(v5)
                        v13 := GenS3{F0: 42, F1: 728, F2: "a"}
                        _ = fi12
                        _ = v13
                    }
                    v4 = (^(v4))
                    var v14 float32 = 0.1
                    _ = v14
                } else {
                    v15 := GenS4{F0: 32767}
                    _ = v15
                }
                _ = v11
            } else {
                v4 = genFn0(43, v10, (v4 + v4))
                for fi16 := 0; fi16 < 14; fi16++ {
                    fmt.Println(v2, v1)
                    _ = fi16
                }
                if ((^(v10)) <= (v10 ^ v10)) {
                    if false {
                        v10 = (v10 / ((747) | 1))
                        v17 := GenS1{F0: "\n\t"}
                        var v18 int = 431
                        _ = v17
                        _ = v18
                    }
                    v1 -= (v1 * 0)
                } else {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        fmt.Println(fi19, fi19)
                        for fi20 := 0; fi20 < 15; fi20++ {
                            var v21 float64 = 1e10
                            if false {
                                var v22 float64 = v21
                                v4 = genFn0(1, v10, (v4 >> ((255) & 63)))
                                v1 += 32767
                                _ = v22
                            } else {
                                var v23 float32 = (((float32(v21) / -2.25) - float32(1e10)) / 1e10)
                                v23 -= v23
                                _ = v23
                            }
                            _ = fi20
                            _ = v21
                        }
                        if true {
                            v24 := GenS2{F0: v10}
                            var v25 float32 = (1.5 * float32(((-2.25 / ((0.1 * (float64(-2.25) + (3.14159 - 1.5))) + 3.14159)) * float64(-2.25))))
                            _ = v24
                            _ = v25
                        } else {
                            v26 := GenS5{F0: ("hello" + "\n\t"), F1: (v1 - 296)}
                            _ = v26
                        }
                        _ = fi19
                    }
                    v27 := v3[""]
                    var v28 float32 = ((3.14159 + 3.14159) * float32(1e10))
                    _ = v27
                    _ = v28
                }
            }
        }
    default:
        v29 := GenS3{F0: 0, F1: 2, F2: ""}
        _ = v29
    }
    var v30 uint8 = 1
    var v31 float32 = (1e10 - float32(1.5))
    var v32 uint8 = (^(v30))
    switch "a" {
    case "a":
        v33 := GenS6{F0: v4}
        _ = v33
    case "\n\t":
        v34 := GenS1{F0: "hello"}
        _ = v34
    }
    v35 := GenS6{F0: v4}
    fmt.Println(v1)
    v36 := GenS7{F0: 123}
    fmt.Println(v32)
    v5 = append(v5, "\n\t")
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch -843 {
    case -2147483648:
        switch ((("foo bar" + "a") + "") + "\n\t") {
        case "foo bar":
            v37 := GenS8{F0: false}
            _ = v37
        case "hello":
            var v38 bool = (!true)
            _ = v38
        case "a":
            v39 := GenS2{F0: v10}
            _ = v39
        default:
            fmt.Println(v10, v32)
        }
    case -567:
        for fi40 := 0; fi40 < 4; fi40++ {
            for fi41 := 0; fi41 < 1; fi41++ {
                var v42 float64 = float64(((v31 + float32(float64(0.1))) * v31))
                v43 := GenS3{F0: uint(v10), F1: 2, F2: "\n\t"}
                _ = fi41
                _ = v42
                _ = v43
            }
            _ = fi40
        }
    case 623:
        for fi44 := 0; fi44 < 8; fi44++ {
            v45 := GenS8{F0: true}
            var v46 uint8 = v32
            v47 := GenS7{F0: v46}
            _ = fi44
            _ = v45
            _ = v46
            _ = v47
        }
    default:
        var v48 uint32 = v10
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
