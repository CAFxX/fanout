package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 bool }
func genFn0(p0 string) int32 {
    fmt.Println(p0)
    if true {
        v5 := []int{}
        p0 = p0
        if len(v5) > 0 {
            v5[0] = 128
        }
        _ = v5
    } else {
        var v6 bool = (2 >= 2)
        p0 = ((p0 + "foo bar") + (("" + (p0 + p0)) + ""))
        _ = v6
    }
    switch 0 {
    case 32767:
        p0 = (p0 + "\n\t")
    case 0:
        fmt.Println(p0)
    }
    if (886 != -2147483648) {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 uint8 = 2
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := []int{(fi9 << ((0) & 63)), -9223372036854775808, 127}
            _ = fi9
            _ = v10
        }
        fmt.Println(p0)
    } else {
        v11 := []int{847, 257, 128}
        _ = len(v11)
        _ = v11
    }
    _ = p0
    return 0
}
func genFn1(p0 int64) uint {
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: 18446744073709551615, F1: (fi12 >> ((0) & 63))}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        if true {
            v15 := GenS1{F0: 2, F1: -1}
            if true {
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 uint32 = 927
                    var v18 int32 = 706
                    switch "foo bar" {
                    case "a":
                        switch 849 {
                        case -658:
                            fmt.Println(v17, v18)
                        case -2147483648:
                            var v19 uint16 = 952
                            _ = v19
                        }
                        var v20 uint16 = uint16(v17)
                        var v21 uint32 = v17
                        _ = v20
                        _ = v21
                    case "foo bar":
                        var v22 int16 = -449
                        _ = v22
                    case "":
                        if true {
                            v23 := GenS1{F0: 598, F1: (fi16 | fi14)}
                            switch (^(fi16)) {
                            case 0:
                                v17 += (v17 ^ v17)
                            case 32767:
                                v24 := GenS1{F0: 4294967295, F1: fi14}
                                p0 = (p0 / (((p0 & p0)) | 1))
                                _ = v24
                            case -333:
                                v18 -= v18
                            default:
                                v18 = v18
                            }
                            _ = v23
                        }
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                v25 := GenS1{F0: 0, F1: fi14}
                v26 := []int{(fi14 & fi14)}
                _ = v25
                _ = v26
            }
            if (true || false) {
                v27 := []int{fi14, fi14, (fi14 >> ((104) & 63)), fi14}
                p0 += p0
                p0 -= p0
                _ = v27
            } else {
                var v28 uint16 = 633
                _ = v28
            }
            _ = v15
        } else {
            if true {
                p0 += (^(p0))
                v29 := GenS2{F0: 355, F1: 42, F2: 32767}
                fmt.Println(p0)
                _ = v29
            } else {
                var v30 int8 = -128
                v30 = (v30 << ((2) & 63))
                if true {
                    p0 = p0
                } else {
                    p0 = (^(p0))
                    for fi31 := 0; fi31 < 1; fi31++ {
                        v30 = int8(p0)
                        v32 := GenS3{F0: 2, F1: -1}
                        p0 = (p0 - (^(p0)))
                        var v33 uint8 = 127
                        v33 = 45
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                }
                _ = v30
            }
            if (0 == 4294967295) {
                if ((float64(-2.25) * -2.25) >= float64(float32(0.1))) {
                    var v34 uint8 = 76
                    v35 := GenS1{F0: 2, F1: fi14}
                    _ = v34
                    _ = v35
                }
            }
            fmt.Println(p0, p0)
        }
        v36 := []int{(fi14 ^ (fi14 >> ((716) & 63)))}
        fmt.Println(p0, fi14)
        _ = fi14
        _ = v36
    }
    _ = p0
    return 255
}

func main() {
    var v1 string = "\n\t"
    var v2 bool = (false || (!(120 > 2)))
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 164
    v1 = "a"
    v1 = ""
    var v37 int32 = 2147483647
    v2 = ((!v2) && ((0 != -1) && (0 < 1000)))
    for fi38 := 0; fi38 < 8; fi38++ {
        var v39 uint16 = 32767
        _ = fi38
        _ = v39
    }
    delete(v3, (v4 | v4))
    v40 := GenS4{F0: ("hello" > v1)}
    var v41 float32 = 3.14159
    v1 = (v1 + ("" + "\n\t"))
    v2 = (2 <= 42)
    if (!(!false)) {
        v1 = (((("foo bar" + v1) + (v1 + v1)) + v1) + v1)
        if true {
            var v42 int8 = 127
            switch (v1 + (v1 + "foo bar")) {
            case "hello":
                var v43 int16 = -32768
                _ = v43
            }
            _ = v42
        } else {
            v1 = "foo bar"
            v1 = ("a" + ((v1 + v1) + v1))
        }
        v41 += (v41 + v41)
    } else {
        v2 = ((935 != 1) && (-9223372036854775808 == 9223372036854775807))
        if true {
            switch "\n\t" {
            case "a":
                var v44 int8 = -16
                _ = v44
            case "foo bar":
                v4 -= 447
            case "":
                v2 = true
            }
        } else {
            for fi45 := 0; fi45 < 5; fi45++ {
                delete(v3, 2147483648)
                var v46 int = (fi45 / (((fi45 | (fi45 - (fi45 % ((1) | 1))))) | 1))
                _ = fi45
                _ = v46
            }
            v1 = "a"
            if v2 {
                var v47 uint16 = 0
                v48 := GenS1{F0: 2, F1: -68}
                v41 += float32(0.0)
                _ = v47
                _ = v48
            } else {
                for fi49 := 0; fi49 < 8; fi49++ {
                    var v50 int = -1
                    switch (v50 / ((int(v4)) | 1)) {
                    case 32767:
                        v51 := []int{fi49, fi49, v50, v50}
                        _ = v51
                    }
                    _ = fi49
                    _ = v50
                }
                fmt.Println(v4, v37, v1)
                for fi52 := 0; fi52 < 7; fi52++ {
                    var v53 int8 = 100
                    _ = fi52
                    _ = v53
                }
            }
        }
        v1 = v1
    }
    switch 128 {
    case 128:
        for fi54 := 0; fi54 < 15; fi54++ {
            var v55 int = fi54
            for fi56 := 0; fi56 < 18; fi56++ {
                v55 = v55
                var v57 int32 = genFn0(v1)
                var v58 int8 = -128
                v59 := GenS4{F0: v2}
                var v60 bool = v2
                _ = fi56
                _ = v57
                _ = v58
                _ = v59
                _ = v60
            }
            var v61 string = v1
            _ = fi54
            _ = v55
            _ = v61
        }
    default:
        fmt.Println(v41)
    }
    for fi62 := 0; fi62 < 10; fi62++ {
        var v63 uint32 = v4
        _ = fi62
        _ = v63
    }
    var v64 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v64
}
