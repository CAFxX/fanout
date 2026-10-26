package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0(p0 int64, p1 int) int {
    p0 -= (p0 % ((p0) | 1))
    p1 = -827
    var v5 string = ("hello" + ("\n\t" + "hello"))
    for fi6 := 0; fi6 < 16; fi6++ {
        if true {
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := []int{(fi6 + fi7), (-(p1))}
                v5 = (v5 + ((v5 + "") + "\n\t"))
                _ = fi7
                _ = v8
            }
            var v9 uint16 = 667
            _ = v9
        } else {
            for fi10 := 0; fi10 < 10; fi10++ {
                v11 := GenS1{F0: 2, F1: 259}
                v12 := []int{1, (^(fi6)), fi10}
                v13 := GenS1{F0: 50, F1: 2}
                if (4294967295 == 42) {
                    switch v5 {
                    case "a":
                        if true {
                            v5 = "\n\t"
                            var v14 int8 = 0
                            _ = v14
                        } else {
                            var v15 uint8 = 0
                            var v16 bool = false
                            var v17 int16 = -1
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                    case "hello":
                        v12 = append(v12, p1)
                    }
                    if (76 != 42) {
                        v18 := GenS2{F0: 2147483647, F1: 255, F2: 9223372036854775807}
                        _ = v18
                    }
                    fmt.Println(p1, p1)
                    v19 := GenS2{F0: 0, F1: 0, F2: p0}
                    p0 = 0
                    _ = v19
                } else {
                    p1 = fi10
                    p1 = (p1 | (^(fi10)))
                    p1 = p1
                }
                switch -2147483648 {
                case -829:
                    var v20 uint32 = 2147483648
                    _ = v20
                case 354:
                    fmt.Println(p0)
                case 9223372036854775807:
                    if (!false) {
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        fmt.Println(fi6, p0)
                    }
                default:
                    v21 := GenS1{F0: 327, F1: 649}
                    _ = v21
                }
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v5
    return p1
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = -2.25
    var v3 string = "a"
    var v4 bool = (!true)
    for fi22 := 0; fi22 < 14; fi22++ {
        v4 = ((0 <= 2) || ((-2.25 + ((-2.25 * 1.5) + 3.14159)) > (((3.14159 / -2.25) + ((3.14159 / 3.14159) / float32((float64(1.5) / -2.25)))) - (0.0 + 0.0))))
        fmt.Println(v2, v3)
        if true {
            v23 := []int{genFn0(-1, (fi22 + (^(fi22)))), fi22, (-(fi22)), -327}
            for fi24 := 0; fi24 < 15; fi24++ {
                var v25 bool = (2147483648 > 657)
                if (1 < 65535) {
                    fmt.Println(v2, v25)
                    var v26 uint = 826
                    _ = v26
                }
                v2 -= float64(3.14159)
                _ = fi24
                _ = v25
            }
            _ = v23
        }
        _ = fi22
    }
    v2 += v2
    fmt.Println(v3, v2)
    for fi27 := 0; fi27 < 13; fi27++ {
        for fi28 := 0; fi28 < 4; fi28++ {
            fmt.Println(v4)
            _ = fi28
        }
        if (-32768 >= -101) {
            if (!((!false) || false)) {
                v29 := v1[(((v3 + "foo bar") + "\n\t") + "hello")]
                if v4 {
                    v29 = (v29 + "hello")
                    v29 = "a"
                    v4 = v4
                } else {
                    switch v3 {
                    case "foo bar":
                        var v30 uint = 42
                        _ = v30
                    }
                }
                _ = v29
            }
        }
        v31 := GenS2{F0: 2147483647, F1: 614, F2: 0}
        _ = fi27
        _ = v31
    }
    fmt.Println(v3, v2, v2)
    v32 := []int{genFn0(776, -9223372036854775808), -2147483648, 32768}
    for fi33 := 0; fi33 < 9; fi33++ {
        if (42 != 458) {
            v3 = ((v3 + v3) + (((("" + v3) + (v3 + ("\n\t" + ""))) + "foo bar") + v3))
            var v34 uint8 = 2
            _ = v34
        } else {
            for fi35 := 0; fi35 < 8; fi35++ {
                v2 = 0.0
                fmt.Println(fi35, v3, fi33)
                _ = fi35
            }
            for fi36 := 0; fi36 < 16; fi36++ {
                _ = len(v32)
                _ = fi36
            }
        }
        v37 := GenS3{F0: 899, F1: (0 < -2147483648)}
        _ = fi33
        _ = v37
    }
    switch 4294967295 {
    case 2147483647:
        fmt.Println(v3, v3, v3)
    case 0:
        var v38 bool = true
        _ = v38
    case 32767:
        switch 32767 {
        case -1:
            var v39 string = "hello"
            _ = v39
        case 76:
            var v40 int = 128
            _ = v40
        case 42:
            if v4 {
                switch 0 {
                case 2147483647:
                    var v41 uint64 = 855
                    _ = v41
                }
            }
        default:
            var v42 uint16 = 706
            _ = v42
        }
    }
    v43 := GenS2{F0: 1, F1: 255, F2: 869}
    v44 := GenS3{F0: 4294967295, F1: v4}
    for fi45 := 0; fi45 < 10; fi45++ {
        v2 -= (0.0 - ((0.0 / (v2 / float64(float32(v2)))) - (v2 + 0.0)))
        fmt.Println(v4, fi45)
        fmt.Println(v2, v2)
        _ = fi45
    }
    for fi46 := 0; fi46 < 18; fi46++ {
        for fi47 := 0; fi47 < 3; fi47++ {
            v4 = (0 > -128)
            v3 = v3
            _ = fi47
        }
        var v48 float64 = 1e10
        _ = fi46
        _ = v48
    }
    var v49 uint32 = 2
    v49 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v43
    _ = v44
    _ = v49
}
