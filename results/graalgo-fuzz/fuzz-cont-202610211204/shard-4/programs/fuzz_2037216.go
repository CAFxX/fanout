package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint16 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 int8; F2 int16 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint8) uint8 {
    fmt.Println(p0)
    if (-1 == -749) {
        v6 := GenS1{F0: 35, F1: (2 > 920), F2: 713}
        v7 := GenS1{F0: 69, F1: (!((p0 != (p0 % (((^(p0))) | 1))) && true)), F2: 295}
        _ = v6
        _ = v7
    } else {
        var v8 string = "a"
        if (((!((-2.25 == 0.1) && true)) && true) && true) {
            var v9 uint8 = p0
            _ = v9
        }
        _ = v8
    }
    v10 := GenS2{F0: false, F1: "\n\t"}
    _ = p0
    _ = v10
    return p0
}
func genFn1() uint16 {
    switch 2147483647 {
    case 1:
        var v11 bool = (42 == 1)
        _ = v11
    case 127:
        var v12 int16 = 0
        _ = v12
    case -553:
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS2{F0: (-2.25 > ((((-2.25 / -2.25) / (0.1 - ((float64(-2.25) * (0.0 + 0.0)) + ((1e10 - 3.14159) + 1e10)))) / 1.0) * -2.25)), F1: ""}
            _ = fi13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        if (false || (100 < 127)) {
            switch "foo bar" {
            case "a":
                v16 := GenS2{F0: false, F1: ""}
                _ = v16
            default:
                for fi17 := 0; fi17 < 1; fi17++ {
                    if true {
                        fmt.Println(fi17, fi17)
                        v18 := GenS3{F0: 0}
                        v19 := []int{fi17, fi17, (fi15 >> ((150) & 63))}
                        _ = v18
                        _ = v19
                    } else {
                        var v20 uint32 = 2
                        _ = v20
                    }
                    _ = fi17
                }
            }
            v21 := []int{32768, 2147483647, (fi15 - fi15), (fi15 - 146)}
            switch "\n\t" {
            case "\n\t":
                if (0 == 42) {
                    v22 := GenS4{F0: 1, F1: 100, F2: 1000}
                    _ = v22
                } else {
                    var v23 string = (("a" + (("a" + "hello") + ((("foo bar" + ("a" + (("" + "") + (((((((("hello" + ("hello" + (("foo bar" + "a") + "\n\t"))) + (("\n\t" + ("\n\t" + "a")) + "")) + "\n\t") + ("hello" + "a")) + "") + "hello") + ("" + "a")) + "\n\t")))) + ("hello" + "a")) + "a"))) + ("a" + ((("foo bar" + "") + "\n\t") + "a")))
                    if (v23 == v23) {
                        if ((fi15 | (fi15 + fi15)) <= fi15) {
                            v24 := GenS3{F0: -1}
                            var v25 int8 = 127
                            _ = v24
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 6; fi26++ {
                                var v27 float32 = (1.5 + (1e10 - -2.25))
                                var v28 float64 = (0.1 / 1.5)
                                v29 := GenS4{F0: 18446744073709551615, F1: 127, F2: 42}
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                _ = fi26
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            var v30 int8 = 127
                            _ = v30
                        }
                    }
                    v31 := GenS1{F0: 0, F1: (((168 <= 670) && false) && (genFn0(genFn0(156)) > 178)), F2: 32767}
                    _ = v23
                    _ = v31
                }
            case "hello":
                var v32 uint16 = 2
                _ = v32
            }
            _ = v21
        }
        if true {
            if ((-279 > 1) && (0 >= 1)) {
                var v33 float64 = (float64(1e10) - 1.5)
                for fi34 := 0; fi34 < 9; fi34++ {
                    if true {
                        v33 -= 1.5
                        v35 := GenS4{F0: 42, F1: 1, F2: 1000}
                        var v36 string = ("a" + "\n\t")
                        _ = v35
                        _ = v36
                    } else {
                        fmt.Println(fi15, v33, fi34)
                        var v37 float32 = 1.5
                        if true {
                            v33 += (v33 / 1e10)
                        } else {
                            v38 := GenS5{F0: 0, F1: 4294967295}
                            _ = v38
                        }
                        _ = v37
                    }
                    if (2 == 4294967295) {
                        v39 := []int{-848, fi15, (fi15 ^ fi15)}
                        _ = v39
                    }
                    _ = fi34
                }
                _ = v33
            } else {
                var v40 uint = 665
                var v41 int32 = 0
                _ = v40
                _ = v41
            }
            for fi42 := 0; fi42 < 16; fi42++ {
                var v43 int32 = 42
                _ = fi42
                _ = v43
            }
            var v44 uint = 965
            _ = v44
        }
        _ = fi15
    }
    var v45 int8 = -28
    for fi46 := 0; fi46 < 10; fi46++ {
        var v47 uint16 = 0
        _ = fi46
        _ = v47
    }
    _ = v45
    return 32767
}

func main() {
    var v1 int = -2147483648
    var v2 uint64 = 645
    var v3 uint64 = v2
    v4 := make(map[uint32]uint32)
    v5 := []uint32{1, 2147483648}
    delete(v4, 0)
    var v48 int32 = int32(v1)
    var v49 uint16 = 133
    delete(v4, 1)
    fmt.Println(v2, v1, v48)
    var v50 float32 = -2.25
    v49 -= genFn1()
    v50 -= (float32(3.14159) - 1.5)
    v3 = v3
    var v51 int16 = -856
    _ = len(v5)
    if (uint64(v49) <= v3) {
        for fi52 := 0; fi52 < 3; fi52++ {
            var v53 string = "\n\t"
            for fi54 := 0; fi54 < 8; fi54++ {
                var v55 uint16 = (v49 << ((856) & 63))
                v3 = v3
                if len(v5) > 0 {
                    v5[0] = 4294967295
                }
                _ = fi54
                _ = v55
            }
            v48 = v48
            _ = fi52
            _ = v53
        }
        v48 += (v48 ^ (v48 << ((65535) & 63)))
    } else {
        var v56 float32 = 0.0
        _ = v56
    }
    v50 += 1e10
    var v57 int16 = (v51 % ((v51) | 1))
    for fi58 := 0; fi58 < 1; fi58++ {
        delete(v4, 0)
        _ = fi58
    }
    v59 := GenS6{F0: 41}
    v2 = (v3 + (^(v3)))
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v57
    _ = v59
}
