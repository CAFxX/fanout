package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 int16; F2 int8 }
func genFn0() uint32 {
    var v7 int8 = 127
    var v8 int = -2147483648
    _ = v7
    _ = v8
    return 852
}

func main() {
    v1 := []uint32{}
    v2 := []int{42, -9223372036854775808}
    v3 := []int{286}
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]bool)
    var v6 uint16 = 0
    var v9 int = 4294967295
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v2) > 0 {
        v2[0] = v9
    }
    v9 -= (-(v9))
    var v10 bool = (!(!true))
    for fi11 := 0; fi11 < 9; fi11++ {
        v6 -= (^(v6))
        v6 -= 32767
        var v12 float32 = 3.14159
        _ = fi11
        _ = v12
    }
    switch ("" + ("a" + "hello")) {
    case "\n\t":
        fmt.Println(v9)
    case "hello":
        v9 += (v9 - (v9 + v9))
    case "":
        var v13 uint64 = 1
        _ = v13
    }
    if ((!(!v10)) && false) {
        if v10 {
            var v14 int = v9
            v15 := GenS1{F0: 0, F1: 545, F2: (^(v6))}
            v16 := GenS2{F0: 65535}
            _ = v14
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 14; fi17++ {
                fmt.Println(v9, fi17, v9)
                _ = fi17
            }
            if (!v10) {
                v6 -= (v6 | v6)
            }
            _ = len(v5)
        }
        fmt.Println(v9, v10)
    } else {
        v9 += v9
    }
    v6 += 32767
    for fi18 := 0; fi18 < 17; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            v9 -= v9
            var v20 int8 = -128
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    if v10 {
        v9 = 657
        var v21 int32 = 42
        _ = v21
    } else {
        v5[4294967295] = (!((1.5 + float32(3.14159)) != 1.5))
        fmt.Println(v6, v9, v6)
        for fi22 := 0; fi22 < 11; fi22++ {
            if ((^(v6)) == v6) {
                v10 = false
                v23 := GenS3{F0: -972, F1: 65535}
                v6 = v6
                _ = v23
            } else {
                for fi24 := 0; fi24 < 6; fi24++ {
                    var v25 string = (((("" + ("a" + "hello")) + ("a" + ("\n\t" + (("a" + ("" + "a")) + ("\n\t" + ("" + ((("\n\t" + ("a" + "\n\t")) + "hello") + "hello"))))))) + "foo bar") + ("foo bar" + ("a" + "a")))
                    var v26 int32 = 0
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v6 -= 2
                v9 = (v9 ^ v9)
            }
            var v27 int = fi22
            for fi28 := 0; fi28 < 19; fi28++ {
                _ = len(v3)
                v9 = v27
                _ = fi28
            }
            _ = fi22
            _ = v27
        }
    }
    switch (v9 % ((v9) | 1)) {
    case -229:
        fmt.Println(v10, v10)
    case 42:
        v6 = v6
    case 32767:
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 string = "\n\t"
            for fi31 := 0; fi31 < 12; fi31++ {
                if v10 {
                    var v32 uint = 255
                    v33 := v5[2147483648]
                    _ = v32
                    _ = v33
                } else {
                    var v34 uint32 = 2147483648
                    switch v9 {
                    case 32767:
                        v6 -= (v6 | 1)
                    }
                    var v35 string = (v30 + (v30 + ("hello" + v30)))
                    _ = v34
                    _ = v35
                }
                _ = fi31
            }
            switch -799 {
            case -1:
                v6 -= v6
            case 9223372036854775807:
                if v10 {
                    v36 := GenS1{F0: -862, F1: -1, F2: v6}
                    _ = v36
                } else {
                    var v37 int16 = 736
                    fmt.Println(v37)
                    v38 := GenS4{F0: (^(v6))}
                    _ = v37
                    _ = v38
                }
            case 768:
                v6 = v6
            default:
                v9 = (v9 + 42)
            }
            _ = fi29
            _ = v30
        }
    default:
        fmt.Println(v6, v10, v6)
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        v40 := GenS2{F0: genFn0()}
        v41 := GenS5{F0: 136, F1: 21, F2: 0}
        v42 := GenS3{F0: 0, F1: 432}
        _ = fi39
        _ = v40
        _ = v41
        _ = v42
    }
    delete(v5, 969)
    var v43 int = v9
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v9
}
