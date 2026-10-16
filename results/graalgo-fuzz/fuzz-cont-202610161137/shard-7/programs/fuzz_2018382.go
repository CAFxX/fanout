package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int32; F2 uint8 }
func genFn0() int64 {
    var v7 int = 0
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint16 = 0
        v10 := GenS1{F0: ("" + "a"), F1: (0.1 > ((1.5 + (3.14159 / 1.5)) / 3.14159))}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint16 = 1
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int16 = 0
        _ = fi12
        _ = v13
    }
    _ = v11
    _ = v7
    return 0
}

func main() {
    var v1 uint16 = 65535
    var v2 int16 = 1000
    var v3 string = "\n\t"
    var v4 int32 = 2147483647
    v5 := make(map[int]string)
    v6 := make(map[int]bool)
    v2 += 32767
    delete(v6, 127)
    var v14 int32 = -405
    var v15 uint = 65535
    var v16 float32 = (0.0 * 0.0)
    v17 := []int{int(v2), 129}
    fmt.Println(v15, v4, v4)
    v5[-20] = "hello"
    switch "hello" {
    case "foo bar":
        fmt.Println(v1, v16, v16)
    }
    var v18 float32 = v16
    fmt.Println(v2)
    v19 := GenS1{F0: "a", F1: (!(1 < v14))}
    if (!false) {
        var v20 uint8 = 42
        _ = v20
    }
    v15 = (v15 | 65535)
    if false {
        v21 := GenS1{F0: "\n\t", F1: true}
        _ = v21
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            v24 := v5[fi23]
            v3 = "a"
            _ = fi23
            _ = v24
        }
        v15 = uint(v1)
        switch "a" {
        case "hello":
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 int16 = (v2 + v2)
                var v27 int = fi25
                v28 := GenS1{F0: "hello", F1: (v14 != v14)}
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
        case "foo bar":
            var v29 uint = (v15 ^ v15)
            _ = v29
        case "":
            v30 := GenS2{F0: (v4 & v14)}
            _ = v30
        default:
            var v31 string = ""
            _ = v31
        }
        _ = fi22
    }
    if false {
        v16 = (v18 * 0.1)
        if false {
            var v32 uint32 = 0
            if true {
                v33 := GenS3{F0: v3}
                if (9223372036854775807 == 32767) {
                    var v34 float64 = ((0.0 - (float64(v18) * 3.14159)) - ((((((-2.25 * 1.5) / 0.1) + (1.5 + (float64(v16) / (((3.14159 + 1e10) / 0.1) + -2.25)))) / 0.1) * 0.1) * -2.25))
                    v18 += v16
                    v35 := GenS3{F0: (((v3 + ("a" + (v3 + (((v3 + (v3 + "\n\t")) + v3) + v3)))) + v3) + v3)}
                    _ = v34
                    _ = v35
                } else {
                    var v36 uint32 = (v32 + v32)
                    v18 -= 0.1
                    fmt.Println(v2, v1, v4)
                    _ = v36
                }
                _ = v33
            }
            v37 := GenS2{F0: (v4 / ((v4) | 1))}
            _ = v32
            _ = v37
        } else {
            for fi38 := 0; fi38 < 19; fi38++ {
                for fi39 := 0; fi39 < 20; fi39++ {
                    var v40 string = v3
                    v4 -= v4
                    fmt.Println(v15, v18, v3)
                    _ = fi39
                    _ = v40
                }
                if true {
                    if (-604 > fi38) {
                        v16 += (v16 / -2.25)
                    }
                    var v41 uint8 = 105
                    var v42 string = v3
                    _ = v41
                    _ = v42
                }
                _ = fi38
            }
            var v43 float64 = 0.0
            _ = v43
        }
        switch 0 {
        case 3:
            v1 += v1
        case 250:
            for fi44 := 0; fi44 < 3; fi44++ {
                if (!false) {
                    switch -495 {
                    case 936:
                        switch fi44 {
                        case -536:
                            v45 := GenS2{F0: (v4 - v4)}
                            _ = v45
                        case 32767:
                            fmt.Println(v1, v14)
                        case -843:
                            var v46 uint32 = 756
                            _ = v46
                        default:
                            v3 = v3
                        }
                    case 2147483647:
                        v2 = v2
                    case 792:
                        if false {
                            fmt.Println(fi44, v4, v1)
                        } else {
                            v47 := GenS3{F0: v3}
                            v4 = 0
                            _ = v47
                        }
                    }
                    var v48 float64 = (1.5 / 3.14159)
                    _ = v48
                } else {
                    switch 4294967295 {
                    case -560:
                        v15 = v15
                    case 42:
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                    default:
                        var v49 uint8 = 150
                        _ = v49
                    }
                    for fi50 := 0; fi50 < 10; fi50++ {
                        v51 := GenS4{F0: (fi44 % ((fi44) | 1)), F1: -2147483648, F2: uint8(v1)}
                        v14 -= 1
                        v4 = v14
                        _ = fi50
                        _ = v51
                    }
                }
                _ = fi44
            }
        case 319:
            for fi52 := 0; fi52 < 14; fi52++ {
                var v53 int64 = 315
                if true {
                    var v54 float32 = ((3.14159 / v18) + (float32(-2.25) * (((v18 + v18) / 1.5) * (v16 + v18))))
                    _ = v54
                } else {
                    if false {
                        if ((fi52 < (fi52 << ((uint(v1)) & 63))) || true) {
                            v55 := GenS1{F0: ("" + ((v3 + v3) + "")), F1: true}
                            _ = v55
                        }
                        if true {
                            if true {
                                v56 := GenS2{F0: (v14 + 42)}
                                _ = v56
                            }
                            v4 = v4
                        } else {
                            var v57 int64 = (-(v53))
                            if true {
                                v57 = v57
                            }
                            var v58 string = (v3 + "foo bar")
                            _ = v57
                            _ = v58
                        }
                    } else {
                        if (v2 == (-(v2))) {
                            fmt.Println(v15)
                            var v59 int8 = 100
                            _ = v59
                        }
                        delete(v6, (fi52 | fi52))
                        if true {
                            fmt.Println(v18, v53, v4)
                        }
                    }
                }
                _ = fi52
                _ = v53
            }
        }
    }
    switch ("\n\t" + v3) {
    case "":
        v16 += 0.0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
