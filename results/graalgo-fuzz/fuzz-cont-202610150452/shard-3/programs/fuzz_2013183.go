package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 int32; F1 int8; F2 bool }
func genFn0(p0 uint8, p1 string, p2 uint32) uint64 {
    switch -800 {
    case 32768:
        fmt.Println(p0, p1, p2)
    case 360:
        switch "\n\t" {
        case "\n\t":
            var v7 uint32 = 4294967295
            _ = v7
        case "foo bar":
            v8 := []int{-2147483648, 0}
            _ = v8
        }
    }
    v9 := GenS1{F0: (^(p0)), F1: 42}
    var v10 int16 = -32768
    for fi11 := 0; fi11 < 15; fi11++ {
        switch 1 {
        case 42:
            fmt.Println(v10)
        case -1:
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 int32 = 2147483647
                var v14 uint64 = 1
                switch 127 {
                case -9223372036854775808:
                    fmt.Println(fi11, fi11, fi11)
                case 4294967295:
                    p1 = p1
                case 42:
                    v15 := []int{fi12}
                    _ = v15
                }
                _ = fi12
                _ = v13
                _ = v14
            }
        case 0:
            v16 := []int{}
            _ = v16
        }
        var v17 int32 = 42
        if (!true) {
            switch p1 {
            case "\n\t":
                if true {
                    v18 := GenS2{F0: 13, F1: (^(p2))}
                    p0 += (^(p0))
                    switch fi11 {
                    case 128:
                        v19 := GenS2{F0: 0, F1: (p2 & (p2 | (^(p2))))}
                        _ = v19
                    default:
                        switch p1 {
                        case "":
                            var v20 int8 = 0
                            _ = v20
                        default:
                            v21 := GenS3{F0: p1, F1: 42}
                            _ = v21
                        }
                    }
                    _ = v18
                }
            case "foo bar":
                var v22 int8 = -77
                _ = v22
            default:
                v23 := GenS3{F0: "foo bar", F1: (p0 * p0)}
                _ = v23
            }
            for fi24 := 0; fi24 < 17; fi24++ {
                var v25 string = "\n\t"
                v26 := GenS2{F0: 9223372036854775807, F1: (p2 | (p2 & (^(p2))))}
                v27 := []int{(fi24 ^ -105)}
                if false {
                    if true {
                        switch 42 {
                        case 4294967295:
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                        }
                    } else {
                        var v28 string = ((p1 + (p1 + (p1 + ((v25 + v25) + v25)))) + v25)
                        _ = v28
                    }
                    v27 = append(v27, fi24)
                }
                v10 = v10
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            v10 = 278
        } else {
            v29 := GenS1{F0: p0, F1: p2}
            if true {
                var v30 int = int(v10)
                var v31 string = ""
                _ = v30
                _ = v31
            }
            _ = v29
        }
        _ = fi11
        _ = v17
    }
    p2 = 550
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 468
}
func genFn1() uint32 {
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 int32 = 1
        _ = fi32
        _ = v33
    }
    v34 := []int{987, 32767, 127}
    switch "" {
    case "":
        if false {
            var v35 uint16 = 32767
            for fi36 := 0; fi36 < 18; fi36++ {
                var v37 bool = (!(!true))
                var v38 uint = 18446744073709551615
                v37 = true
                _ = fi36
                _ = v37
                _ = v38
            }
            var v39 int8 = -1
            _ = v35
            _ = v39
        }
    case "hello":
        if true {
            if len(v34) > 0 {
                v34[0] = -510
            }
            var v40 int32 = 2147483647
            var v41 uint32 = 2147483648
            _ = v40
            _ = v41
        }
    case "\n\t":
        if true {
            for fi42 := 0; fi42 < 12; fi42++ {
                v34 = append(v34, (-(fi42)))
                var v43 int32 = -525
                for fi44 := 0; fi44 < 14; fi44++ {
                    fmt.Println(v43, v43, v43)
                    v43 -= v43
                    _ = fi44
                }
                _ = fi42
                _ = v43
            }
            switch 9223372036854775807 {
            case -9223372036854775808:
                if len(v34) > 0 {
                    v34[0] = 32767
                }
            }
            var v45 int = -837
            _ = v45
        }
    }
    _ = v34
    return 1
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]int)
    var v3 float64 = 1.5
    v4 := make(map[string]string)
    var v5 int8 = 0
    v6 := make(map[string]bool)
    v46 := []int{}
    for fi47 := 0; fi47 < 18; fi47++ {
        if true {
            switch ("\n\t" + ("" + ("foo bar" + "a"))) {
            case "hello":
                for fi48 := 0; fi48 < 8; fi48++ {
                    v3 -= v3
                    _ = fi48
                }
            case "":
                var v49 uint32 = (v1 * 49)
                _ = v49
            case "\n\t":
                v3 = v3
            }
            for fi50 := 0; fi50 < 17; fi50++ {
                v1 += (v1 ^ v1)
                _ = fi50
            }
            v1 = genFn1()
        } else {
            var v51 uint16 = 32767
            _ = v51
        }
        _ = fi47
    }
    v52 := GenS4{F0: 657, F1: -1}
    if (-179 >= -32768) {
        if false {
            if false {
                var v53 float32 = float32(((v3 * v3) + v3))
                _ = v53
            } else {
                v54 := GenS2{F0: 1, F1: (v1 * genFn1())}
                var v55 string = ""
                _ = v54
                _ = v55
            }
            v3 -= 0.0
            v56 := GenS2{F0: 9223372036854775807, F1: (^(v1))}
            _ = v56
        } else {
            v46 = append(v46, -441)
            v5 += -128
            var v57 uint64 = 0
            _ = v57
        }
        for fi58 := 0; fi58 < 7; fi58++ {
            v59 := GenS5{F0: -278, F1: v5, F2: false}
            _ = fi58
            _ = v59
        }
    }
    fmt.Println(v3)
    for fi60 := 0; fi60 < 19; fi60++ {
        delete(v6, ("foo bar" + "hello"))
        delete(v6, "a")
        _ = fi60
    }
    for fi61 := 0; fi61 < 18; fi61++ {
        var v62 float32 = -2.25
        _ = fi61
        _ = v62
    }
    for fi63 := 0; fi63 < 19; fi63++ {
        if len(v46) >= 2 {
            v46 = v46[1:len(v46)]
        }
        _ = fi63
    }
    for fi64 := 0; fi64 < 8; fi64++ {
        v65 := GenS5{F0: -964, F1: -1, F2: true}
        v1 = 2147483647
        for fi66 := 0; fi66 < 3; fi66++ {
            v1 = (v1 >> ((1) & 63))
            if len(v46) >= 2 {
                v46 = v46[1:len(v46)]
            }
            v1 = genFn1()
            _ = fi66
        }
        _ = fi64
        _ = v65
    }
    v3 -= (float64(1e10) - -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v52
    _ = v6
}
