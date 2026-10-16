package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 int8; F2 int8 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := []int{fi4, fi4}
        _ = fi4
        _ = v5
    }
    var v6 float64 = 3.14159
    v7 := GenS1{F0: ("foo bar" + ("\n\t" + "foo bar")), F1: 90, F2: -1}
    _ = v6
    _ = v7
    return 0
}
func genFn1() int32 {
    switch ("\n\t" + ("foo bar" + "")) {
    case "\n\t":
        for fi8 := 0; fi8 < 20; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS1{F0: ("" + "\n\t"), F1: 4294967295, F2: 1}
                _ = fi9
                _ = v10
            }
            if false {
                fmt.Println(fi8)
            } else {
                var v11 float32 = -2.25
                _ = v11
            }
            v12 := []int{(fi8 << ((65535) & 63)), 4294967295}
            _ = fi8
            _ = v12
        }
    case "foo bar":
        if (!true) {
            var v13 int = 917
            for fi14 := 0; fi14 < 1; fi14++ {
                if false {
                    v15 := []int{fi14, (fi14 << ((696) & 63)), 1}
                    var v16 uint16 = 2
                    _ = v15
                    _ = v16
                }
                v17 := []int{2147483647, v13, fi14}
                _ = fi14
                _ = v17
            }
            var v18 uint32 = 0
            _ = v13
            _ = v18
        } else {
            if true {
                for fi19 := 0; fi19 < 17; fi19++ {
                    var v20 string = ("hello" + "\n\t")
                    if false {
                        v21 := []int{}
                        for fi22 := 0; fi22 < 20; fi22++ {
                            v20 = v20
                            v20 = v20
                            v20 = "\n\t"
                            v21 = append(v21, (^(fi22)))
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            _ = len(v21)
                            var v23 string = v20
                            _ = fi22
                            _ = v23
                        }
                        var v24 bool = true
                        _ = v21
                        _ = v24
                    }
                    v25 := GenS2{F0: 42}
                    _ = fi19
                    _ = v20
                    _ = v25
                }
                var v26 int32 = -2147483648
                _ = v26
            }
            var v27 int16 = 381
            _ = v27
        }
    case "a":
        switch 1 {
        case 503:
            var v28 int64 = 1
            _ = v28
        }
    }
    var v29 uint8 = 1
    if false {
        var v30 int8 = 0
        v30 -= (v30 + (v30 - v30))
        _ = v30
    }
    _ = v29
    return -222
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v31 := GenS1{F0: "hello", F1: 50, F2: -1}
    v1 = v1
    var v32 uint8 = 1
    v33 := GenS2{F0: int64(v32)}
    var v34 uint = 42
    v32 += v32
    for fi35 := 0; fi35 < 13; fi35++ {
        v32 = v32
        for fi36 := 0; fi36 < 1; fi36++ {
            var v37 uint8 = v32
            _ = fi36
            _ = v37
        }
        if (1.5 >= ((-2.25 - 1e10) + (1.5 * 1e10))) {
            var v38 bool = (!(!true))
            v34 += (^(v34))
            _ = v38
        } else {
            for fi39 := 0; fi39 < 5; fi39++ {
                if (727 != -1) {
                    v40 := GenS1{F0: "foo bar", F1: (v1 ^ 69), F2: genFn1()}
                    v41 := GenS3{F0: (1 != 1)}
                    _ = v40
                    _ = v41
                } else {
                    var v42 uint64 = 394
                    v43 := []int{-2147483648, fi39, fi35, int(v1)}
                    if false {
                        v42 -= 163
                        switch ("\n\t" + ("a" + "\n\t")) {
                        case "\n\t":
                            v44 := v3[0]
                            _ = v44
                        case "":
                            v34 = (v34 & v34)
                        case "hello":
                            var v45 uint8 = v32
                            _ = v45
                        default:
                            switch fi35 {
                            case 837:
                                v43 = append(v43, (fi39 >> ((260) & 63)))
                                var v46 int32 = 1
                                _ = v46
                            case 61:
                                if (((0.0 * 0.1) * float32((0.1 * (-2.25 * float64(3.14159))))) >= (float32(1.5) - (3.14159 + 1.5))) {
                                    v2[128] = true
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v47 uint = 613
                                        _ = v47
                                    case "":
                                        v48 := GenS2{F0: -1}
                                        _ = v48
                                    case "hello":
                                        switch "" {
                                        case "foo bar":
                                            var v49 uint32 = 42
                                            _ = v49
                                        case "\n\t":
                                            _ = len(v3)
                                        }
                                    default:
                                        v32 += (v32 | v32)
                                    }
                                } else {
                                    fmt.Println(v42, fi35)
                                }
                                var v50 uint64 = v42
                                var v51 uint = v34
                                var v52 bool = ((v32 << ((v34) & 63)) <= (v32 / ((v32) | 1)))
                                if (-1 >= 42) {
                                    var v53 int8 = 100
                                    var v54 float64 = 1e10
                                    _ = v53
                                    _ = v54
                                } else {
                                    switch (fi39 & 1) {
                                    case 475:
                                        if ((float64(3.14159) / 1.0) < 0.1) {
                                            v55 := GenS4{F0: -2147483648}
                                            fmt.Println(v42, v32, v51)
                                            var v56 int = 2147483647
                                            _ = v55
                                            _ = v56
                                        }
                                    case -1:
                                        switch ("a" + ((("hello" + ("a" + ("foo bar" + "\n\t"))) + "\n\t") + "")) {
                                        case "":
                                            v34 -= v51
                                        case "\n\t":
                                            fmt.Println(v32)
                                        case "foo bar":
                                            if (255 < 42) {
                                                v1 -= (v50 - (v1 ^ v50))
                                                v50 = 18446744073709551615
                                            }
                                        default:
                                            fmt.Println(v42, fi39)
                                        }
                                    }
                                    v52 = ((v52 || (1 <= -1)) || (((false || ((!v52) && ((v51 / (((v51 * (v34 << (((v51 & (v51 % ((2) | 1)))) & 63)))) | 1)) == 485))) && (4294967295 > 218)) || (-1 < 42)))
                                }
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                        }
                    } else {
                        fmt.Println(fi39)
                        switch ((("\n\t" + "hello") + "hello") + "foo bar") {
                        case "\n\t":
                            fmt.Println(v34, fi35)
                        case "foo bar":
                            v42 += 709
                        }
                    }
                    _ = v42
                    _ = v43
                }
                var v57 uint8 = (v32 << (((v34 % (((v34 ^ v34)) | 1))) & 63))
                _ = fi39
                _ = v57
            }
            switch fi35 {
            case -243:
                v34 += (v34 - v34)
            case 128:
                v58 := []int{(fi35 | fi35)}
                _ = v58
            case 32768:
                var v59 int = fi35
                _ = v59
            default:
                v32 = v32
            }
        }
        _ = fi35
    }
    var v60 uint64 = 843
    v32 += v32
    v1 = v1
    v61 := GenS5{F0: uint16(v32), F1: int8(v1), F2: 100}
    v1 -= v1
    fmt.Println(v60, v1, v34)
    v62 := v3[-1]
    var v63 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v60:", v60)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
