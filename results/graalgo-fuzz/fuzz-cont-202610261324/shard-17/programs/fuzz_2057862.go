package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 int32 }
func genFn0(p0 bool, p1 uint8) int32 {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: 2, F1: 18446744073709551615}
        v9 := GenS1{F0: 2, F1: 0}
        v10 := GenS1{F0: 2147483647, F1: 46}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    switch "a" {
    case "":
        var v11 bool = false
        _ = v11
    default:
        v12 := GenS1{F0: 65535, F1: 1}
        _ = v12
    }
    if p0 {
        if true {
            v13 := []int{42, 2147483647, 32768, -400}
            _ = v13
        }
        v14 := GenS2{F0: 9223372036854775807, F1: 42}
        _ = v14
    } else {
        if (3.14159 > (1.5 + ((1.5 + float64(0.0)) - ((((1.5 / ((-2.25 + 0.1) + 1.5)) - 1e10) * ((((3.14159 / 1.5) + (1.5 / -2.25)) + 1.5) / (float64(3.14159) / 0.1))) + 1e10)))) {
            v15 := GenS1{F0: 1, F1: 827}
            switch "hello" {
            case "foo bar":
                fmt.Println(p1)
            case "a":
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 uint8 = (p1 % (((p1 << ((526) & 63))) | 1))
                    _ = fi16
                    _ = v17
                }
            default:
                p1 = p1
            }
            _ = v15
        }
        v18 := GenS3{F0: 391}
        var v19 uint32 = uint32(p1)
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    return -9
}
func genFn1(p0 uint16) uint8 {
    switch 895 {
    case 128:
        p0 = (p0 >> ((1) & 63))
    case -1000:
        for fi20 := 0; fi20 < 12; fi20++ {
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 bool = ((701 <= 32767) || true)
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
    case 0:
        var v23 uint = 4294967295
        _ = v23
    }
    p0 -= 491
    switch 2147483647 {
    case 128:
        for fi24 := 0; fi24 < 11; fi24++ {
            v25 := GenS3{F0: 31}
            for fi26 := 0; fi26 < 10; fi26++ {
                v27 := []int{(fi26 << ((658) & 63))}
                switch 42 {
                case 0:
                    if false {
                        v28 := GenS3{F0: 0}
                        v29 := GenS4{F0: 127, F1: 2}
                        v30 := GenS5{F0: 0}
                        _ = v28
                        _ = v29
                        _ = v30
                    } else {
                        var v31 string = "foo bar"
                        if false {
                            var v32 string = v31
                            v33 := GenS2{F0: 42, F1: 2}
                            var v34 uint64 = 613
                            var v35 string = v32
                            v34 += v34
                            _ = v32
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        _ = v31
                    }
                    if true {
                        v36 := GenS2{F0: 9223372036854775807, F1: 602}
                        p0 = 65535
                        p0 = 255
                        _ = v36
                    }
                    fmt.Println(fi24, fi24)
                    if len(v27) > 0 {
                        v27[0] = (^(fi24))
                    }
                    if true {
                        p0 = p0
                    } else {
                        _ = len(v27)
                        v37 := GenS3{F0: 4294967295}
                        p0 += 32767
                        v38 := GenS2{F0: 659, F1: 334}
                        _ = v37
                        _ = v38
                    }
                case -922:
                    v39 := GenS2{F0: 1, F1: 42}
                    _ = v39
                }
                switch "a" {
                case "\n\t":
                    p0 = p0
                case "foo bar":
                    switch (fi24 % ((fi24) | 1)) {
                    case -2147483648:
                        p0 = p0
                    default:
                        var v40 float32 = float32(float64(3.14159))
                        _ = v40
                    }
                default:
                    if len(v27) >= 2 {
                        v27 = v27[1:len(v27)]
                    }
                }
                _ = fi26
                _ = v27
            }
            var v41 int = -525
            _ = fi24
            _ = v25
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        var v43 int16 = 42
        var v44 int64 = 1
        _ = fi42
        _ = v43
        _ = v44
    }
    var v45 uint = 42
    _ = p0
    _ = v45
    return 128
}

func main() {
    var v1 uint16 = 650
    v2 := make(map[int]string)
    var v3 float32 = -2.25
    var v4 uint8 = 128
    var v5 uint32 = 42
    var v6 uint16 = v1
    var v46 uint8 = v4
    for fi47 := 0; fi47 < 11; fi47++ {
        fmt.Println(v4, v5)
        _ = fi47
    }
    if ((!false) && true) {
        v46 += (v46 | genFn1((v6 | (v6 >> ((787) & 63)))))
        for fi48 := 0; fi48 < 17; fi48++ {
            _ = len(v2)
            _ = fi48
        }
    }
    var v49 string = (("\n\t" + ("hello" + "foo bar")) + "\n\t")
    if (v49 >= "") {
        switch 42 {
        case -846:
            var v50 float32 = v3
            _ = v50
        }
    }
    if false {
        switch 850 {
        case 291:
            v6 = v1
        default:
            v51 := []int{32768}
            _ = v51
        }
        for fi52 := 0; fi52 < 1; fi52++ {
            var v53 bool = (!false)
            _ = fi52
            _ = v53
        }
    } else {
        for fi54 := 0; fi54 < 15; fi54++ {
            var v55 uint64 = 957
            for fi56 := 0; fi56 < 15; fi56++ {
                v57 := []int{fi54, fi56}
                v5 -= (v5 / ((2147483647) | 1))
                _ = fi56
                _ = v57
            }
            _ = fi54
            _ = v55
        }
        v6 = 42
    }
    v58 := GenS2{F0: -1, F1: uint(v5)}
    v3 -= (float32(1.5) / (((v3 - ((((v3 * ((v3 - (float32(3.14159) + (((v3 - (float32(1.5) + v3)) - ((0.0 - v3) / 0.1)) * (v3 - v3)))) - 0.0)) / v3) + 1.5) + v3)) + float32(((float64(-2.25) / -2.25) + (1e10 + -2.25)))) + 3.14159))
    fmt.Println(v4)
    var v59 uint8 = genFn1(uint16(v5))
    var v60 int64 = 134
    fmt.Println(v5, v3, v1)
    var v61 string = v49
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v49
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
}
