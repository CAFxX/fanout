package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 bool }
func genFn0(p0 int64, p1 uint32, p2 bool) int32 {
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int16 = -226
        var v7 int32 = -2147483648
        _ = fi5
        _ = v6
        _ = v7
    }
    p2 = p2
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 float32 = (((1.5 * -2.25) * (1e10 + 1.5)) * 0.0)
    var v2 int32 = 0
    var v3 uint8 = 0
    var v4 int8 = 119
    v8 := []int{}
    var v9 uint64 = 677
    var v10 string = ("\n\t" + "")
    if ((v3 + 42) == (v3 % (((v3 ^ (v3 - 218))) | 1))) {
        switch v10 {
        case "foo bar":
            var v11 uint = 18446744073709551615
            _ = v11
        case "a":
            if false {
                var v12 uint32 = 375
                _ = v12
            }
        default:
            v13 := GenS1{F0: 2147483647}
            _ = v13
        }
    }
    fmt.Println(v3)
    for fi14 := 0; fi14 < 15; fi14++ {
        if true {
            var v15 int32 = (v2 ^ 0)
            var v16 uint16 = 677
            if (v3 > (v3 - (v3 - v3))) {
                var v17 int8 = (v4 / ((-128) | 1))
                _ = v17
            }
            _ = v15
            _ = v16
        } else {
            if true {
                fmt.Println(v10, v1)
                for fi18 := 0; fi18 < 15; fi18++ {
                    if false {
                        v8 = append(v8, (fi18 ^ (fi18 / ((128) | 1))))
                        switch fi18 {
                        case 127:
                            fmt.Println(fi18, v2)
                        case 32768:
                            v19 := GenS2{F0: v2, F1: v9}
                            _ = v19
                        case 1:
                            v20 := GenS2{F0: 2147483647, F1: v9}
                            _ = v20
                        }
                        switch fi18 {
                        case 0:
                            var v21 float32 = float32(3.14159)
                            _ = v21
                        default:
                            v22 := GenS3{F0: 4294967295}
                            _ = v22
                        }
                    }
                    v4 = v4
                    _ = fi18
                }
                if len(v8) > 0 {
                    v8[0] = (-(fi14))
                }
            } else {
                if true {
                    if false {
                        switch ("hello" + v10) {
                        case "foo bar":
                            for fi23 := 0; fi23 < 5; fi23++ {
                                var v24 float64 = -2.25
                                _ = fi23
                                _ = v24
                            }
                        case "hello":
                            var v25 int64 = 988
                            _ = v25
                        case "a":
                            var v26 bool = false
                            _ = v26
                        }
                    }
                } else {
                    switch v10 {
                    case "a":
                        var v27 int32 = v2
                        _ = v27
                    default:
                        for fi28 := 0; fi28 < 18; fi28++ {
                            v29 := GenS3{F0: 65535}
                            v1 = (float32((3.14159 + 1.5)) - v1)
                            var v30 uint8 = v3
                            if true {
                                fmt.Println(v3, fi14, v3)
                                v31 := GenS4{F0: 495, F1: ((v10 + (((((v10 + v10) + v10) + "\n\t") + v10) + (v10 + ""))) + v10)}
                                v3 = (v30 / (((v30 & v3)) | 1))
                                _ = v31
                            } else {
                                if ((1.5 + 1.5) != 0.1) {
                                    var v32 float64 = 0.1
                                    var v33 uint16 = 42
                                    _ = len(v8)
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v34 uint = 18446744073709551615
                                    if (true || (!true)) {
                                        var v35 float64 = 3.14159
                                        fmt.Println(v9)
                                        if false {
                                            if (2 > 2) {
                                                var v36 uint64 = (v9 - (v9 << (((v34 * v34)) & 63)))
                                                var v37 float64 = (float64((v1 - ((-2.25 - float32((1.5 * float64(float32(3.14159))))) / -2.25))) * v35)
                                                switch v10 {
                                                case "foo bar":
                                                    var v38 int8 = v4
                                                    _ = v38
                                                case "hello":
                                                    if (v3 != (v30 | (v3 ^ v30))) {
                                                        v4 += (v4 - 100)
                                                        _ = len(v8)
                                                        switch (v10 + (((v10 + ("\n\t" + v10)) + v10) + "a")) {
                                                        case "foo bar":
                                                            _ = len(v8)
                                                        default:
                                                            v39 := GenS5{F0: 1}
                                                            _ = v39
                                                        }
                                                        v40 := GenS5{F0: 42}
                                                        v35 += (0.1 * 1.5)
                                                        _ = v40
                                                    }
                                                case "":
                                                    var v41 uint32 = 2147483647
                                                    fmt.Println(v36, v4)
                                                    _ = v41
                                                }
                                                _ = v36
                                                _ = v37
                                            } else {
                                                var v42 string = ("hello" + (v10 + v10))
                                                _ = v42
                                            }
                                            var v43 int8 = 127
                                            _ = v43
                                        } else {
                                            _ = len(v8)
                                        }
                                        _ = v35
                                    }
                                    if false {
                                        switch ("\n\t" + ("hello" + ((v10 + v10) + (v10 + v10)))) {
                                        case "foo bar":
                                            fmt.Println(v3, v30)
                                        default:
                                            v10 = v10
                                        }
                                        if ((!(!true)) && ((!(!false)) && (1 < v3))) {
                                            switch "a" {
                                            case "a":
                                                var v44 uint = 1
                                                _ = v44
                                            }
                                            v45 := GenS5{F0: 963}
                                            v8 = append(v8, fi14)
                                            _ = v45
                                        }
                                        fmt.Println(v1, fi14, v4)
                                    }
                                    _ = v34
                                }
                                switch v10 {
                                case "":
                                    if false {
                                        switch (fi14 >> ((4294967295) & 63)) {
                                        case 4294967295:
                                            var v46 float64 = -2.25
                                            _ = v46
                                        case 9223372036854775807:
                                            var v47 uint32 = 22
                                            _ = v47
                                        default:
                                            var v48 string = (("a" + (v10 + v10)) + (v10 + ""))
                                            _ = v48
                                        }
                                    }
                                case "a":
                                    v2 = (v2 * genFn0(-9223372036854775808, 2, (v9 > v9)))
                                default:
                                    if false {
                                        _ = len(v8)
                                    }
                                }
                                if len(v8) > 0 {
                                    v8[0] = (fi28 / ((fi14) | 1))
                                }
                            }
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    }
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
                fmt.Println(v9)
            }
            v49 := GenS4{F0: 1, F1: (v10 + "\n\t")}
            _ = v49
        }
        var v50 int32 = genFn0(42, 65535, false)
        v51 := GenS6{F0: (^(v3)), F1: (v3 >> ((uint(v3)) & 63))}
        _ = fi14
        _ = v50
        _ = v51
    }
    var v52 uint64 = v9
    var v53 int16 = -1
    switch 127 {
    case 4294967295:
        var v54 int64 = 438
        _ = v54
    case -250:
        v55 := GenS7{F0: false}
        _ = v55
    case 127:
        var v56 int64 = 1
        _ = v56
    }
    v53 = int16(v2)
    var v57 int32 = (v2 * (v2 % ((v2) | 1)))
    var v58 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v57
    _ = v58
    _ = v8
    _ = v9
}
