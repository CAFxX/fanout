package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int16 }
func genFn0() int {
    for fi6 := 0; fi6 < 19; fi6++ {
        if (!true) {
            v7 := []int{(fi6 + 1), (fi6 & (fi6 | (-(fi6)))), fi6}
            _ = v7
        } else {
            v8 := GenS1{F0: fi6, F1: -777, F2: 255}
            v9 := []int{fi6, (fi6 >> ((18446744073709551615) & 63))}
            _ = v8
            _ = v9
        }
        v10 := []int{fi6}
        _ = fi6
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        if false {
            v12 := GenS1{F0: fi11, F1: 448, F2: 65535}
            v13 := []int{}
            var v14 uint = 4294967295
            _ = v12
            _ = v13
            _ = v14
        } else {
            var v15 int32 = 66
            var v16 int = fi11
            var v17 int32 = v15
            _ = v15
            _ = v16
            _ = v17
        }
        v18 := []int{(fi11 * fi11), -2147483648}
        _ = fi11
        _ = v18
    }
    return 128
}
func genFn1(p0 uint, p1 int64) int64 {
    fmt.Println(p0, p0)
    fmt.Println(p1, p1)
    v19 := GenS1{F0: 765, F1: p1, F2: 65535}
    _ = p0
    _ = p1
    _ = v19
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 32767
    var v2 int64 = -9223372036854775808
    var v3 int32 = 2147483647
    var v4 int = 32767
    v5 := []uint32{503}
    var v20 int8 = -128
    _ = len(v5)
    v2 = genFn1(0, v2)
    v4 += v4
    switch "hello" {
    case "\n\t":
        if (true || (42 <= 155)) {
            v4 -= (^(v4))
        } else {
            switch (("a" + "a") + "a") {
            case "\n\t":
                v4 = 32767
            }
        }
    case "foo bar":
        switch "hello" {
        case "\n\t":
            v20 = -47
        case "foo bar":
            v5 = append(v5, 1)
        case "a":
            for fi21 := 0; fi21 < 8; fi21++ {
                var v22 string = ""
                var v23 float32 = (1.5 / 1.0)
                var v24 float64 = (0.0 - 3.14159)
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
        }
    case "a":
        for fi25 := 0; fi25 < 15; fi25++ {
            if (0.0 == -2.25) {
                v26 := GenS2{F0: (!(false && false))}
                for fi27 := 0; fi27 < 19; fi27++ {
                    if (!true) {
                        v28 := GenS2{F0: (v2 == v2)}
                        if false {
                            _ = len(v5)
                        }
                        if len(v5) > 0 {
                            v5[0] = 1
                        }
                        _ = v28
                    }
                    _ = fi27
                }
                _ = v26
            } else {
                var v29 float32 = (0.0 * 1e10)
                if ((v1 / ((v1) | 1)) <= v1) {
                    switch v4 {
                    case 2147483647:
                        if false {
                            v2 -= (v2 % (((v2 + v2)) | 1))
                            var v30 float32 = 3.14159
                            v2 = (v2 | (v2 + (v2 << ((2) & 63))))
                            _ = v30
                        } else {
                            v2 = (v2 >> ((494) & 63))
                        }
                    case -1:
                        var v31 int8 = (-(v20))
                        _ = v31
                    default:
                        var v32 bool = ((v1 ^ (v1 & (v1 % ((v1) | 1)))) != 232)
                        _ = v32
                    }
                    v2 -= 193
                }
                var v33 uint16 = (v1 & v1)
                _ = v29
                _ = v33
            }
            var v34 uint8 = 0
            _ = fi25
            _ = v34
        }
    }
    v3 += (v3 % ((v3) | 1))
    var v35 int16 = 42
    switch ("" + ((("\n\t" + "a") + "hello") + "\n\t")) {
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "\n\t":
        var v36 uint16 = (^(v1))
        _ = v36
    default:
        if true {
            v4 = (v4 % ((genFn0()) | 1))
            v4 = genFn0()
        } else {
            var v37 uint32 = 2147483648
            switch (v4 | v4) {
            case -455:
                if false {
                    v38 := GenS3{F0: -1, F1: 1}
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 7; fi39++ {
                        fmt.Println(v2, v35, fi39)
                        _ = fi39
                    }
                    v3 -= 0
                    if len(v5) > 0 {
                        v5[0] = (v37 % ((4294967295) | 1))
                    }
                }
            case 844:
                switch "\n\t" {
                case "\n\t":
                    var v40 int64 = v2
                    _ = v40
                case "hello":
                    var v41 uint16 = (^(v1))
                    _ = v41
                case "":
                    v42 := GenS3{F0: 32767, F1: 1}
                    _ = v42
                default:
                    switch (v4 - (^(v4))) {
                    case -481:
                        if false {
                            v43 := GenS1{F0: v4, F1: v2, F2: v1}
                            _ = v43
                        } else {
                            if false {
                                v44 := GenS4{F0: (v35 | (-(v35)))}
                                _ = v44
                            } else {
                                fmt.Println(v2)
                            }
                        }
                    case 128:
                        v3 = (v3 % ((v3) | 1))
                    default:
                        v3 += v3
                    }
                }
            case -9223372036854775808:
                v20 = (v20 * int8(v35))
            }
            v45 := GenS3{F0: 32767, F1: 42}
            _ = v37
            _ = v45
        }
    }
    fmt.Println(v4, v35, v3)
    if true {
        v2 -= (^(v2))
        v3 = (v3 & (v3 % ((v3) | 1)))
        fmt.Println(v1, v1, v3)
    } else {
        if false {
            var v46 float64 = 0.0
            _ = v46
        }
        for fi47 := 0; fi47 < 11; fi47++ {
            if true {
                if true {
                    var v48 uint32 = 0
                    _ = v48
                }
            } else {
                for fi49 := 0; fi49 < 11; fi49++ {
                    var v50 int = fi49
                    var v51 string = "a"
                    _ = fi49
                    _ = v50
                    _ = v51
                }
                v52 := GenS3{F0: (v35 / ((v35) | 1)), F1: 1}
                _ = v52
            }
            _ = fi47
        }
        v1 -= v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
