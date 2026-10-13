package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int64; F2 int }
type GenS3 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint32, p1 uint) int16 {
    fmt.Println(p0, p1)
    v6 := GenS1{F0: 18446744073709551615}
    var v7 int8 = 127
    var v8 int32 = 1
    switch -659 {
    case -9223372036854775808:
        if false {
            fmt.Println(v7, v8)
            var v9 string = "foo bar"
            _ = v9
        }
    case -501:
        v8 -= (-(v8))
    case 42:
        var v10 uint32 = p0
        _ = v10
    default:
        for fi11 := 0; fi11 < 20; fi11++ {
            switch (-(fi11)) {
            case 325:
                for fi12 := 0; fi12 < 9; fi12++ {
                    v13 := []int{(fi11 / ((fi12) | 1)), (^(fi11)), fi11, (fi11 / ((0) | 1))}
                    _ = fi12
                    _ = v13
                }
            case 241:
                v14 := []int{1, (fi11 / ((fi11) | 1)), (fi11 / ((2147483647) | 1))}
                _ = v14
            case -868:
                v15 := []int{fi11, fi11, (fi11 << (((p1 % (((^(p1))) | 1))) & 63)), (^(fi11))}
                _ = v15
            }
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 bool, p1 uint, p2 int) uint16 {
    var v16 int16 = 1
    var v17 float64 = 0.0
    var v18 uint16 = 42
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(p1, v18)
        _ = fi19
    }
    p0 = (6 > 58)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return (v18 * (v18 / ((v18) | 1)))
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[uint32]int)
    v3 := []uint32{4294967295, 4294967295, 2147483648, 838}
    var v4 int16 = 705
    var v5 uint32 = 0
    if (42 >= 4294967295) {
        v4 = v4
        var v20 float64 = 1e10
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v20
    } else {
        for fi21 := 0; fi21 < 16; fi21++ {
            if ((v5 | v5) <= (v5 - v5)) {
                v22 := GenS2{F0: -461, F1: -9223372036854775808, F2: (fi21 ^ (-(fi21)))}
                for fi23 := 0; fi23 < 20; fi23++ {
                    v2[42] = fi23
                    v24 := GenS1{F0: 0}
                    _ = fi23
                    _ = v24
                }
                switch (fi21 + (fi21 | (fi21 - -2147483648))) {
                case 4294967295:
                    v25 := GenS3{F0: false, F1: 1, F2: 2147483648}
                    _ = v25
                case 1:
                    for fi26 := 0; fi26 < 7; fi26++ {
                        if (1 != 32767) {
                            if true {
                                if false {
                                    v4 = (v4 << ((18446744073709551615) & 63))
                                    delete(v2, v5)
                                    _ = len(v2)
                                } else {
                                    fmt.Println(v1)
                                }
                                v2[(v5 | (v5 | v5))] = fi21
                            }
                            v5 -= (v5 | (v5 + (^(v5))))
                            v2[(v5 >> ((2) & 63))] = fi21
                        }
                        var v27 string = (("hello" + "\n\t") + ("foo bar" + "\n\t"))
                        var v28 bool = false
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                case 128:
                    v29 := GenS3{F0: (v4 >= (v4 & v4)), F1: 1, F2: (v5 & (v5 / ((42) | 1)))}
                    _ = v29
                }
                _ = v22
            } else {
                fmt.Println(fi21)
                v5 -= v5
            }
            v30 := GenS4{F0: 2147483647, F1: 42}
            _ = fi21
            _ = v30
        }
        v31 := GenS4{F0: v5, F1: 982}
        _ = v31
    }
    v4 = -32768
    if (false || ((!false) || (!(888 > 42)))) {
        v5 = (^(v5))
        v32 := v2[(v5 | 223)]
        var v33 uint64 = 795
        _ = v32
        _ = v33
    }
    var v34 uint = 65535
    switch ("a" + "foo bar") {
    case "":
        var v35 string = ("foo bar" + (("hello" + "") + "a"))
        _ = v35
    case "foo bar":
        var v36 float64 = -2.25
        _ = v36
    case "a":
        v37 := GenS3{F0: (v1 > (^(v1))), F1: 2, F2: v5}
        _ = v37
    }
    if (0 > 1) {
        switch 802 {
        case 860:
            for fi38 := 0; fi38 < 10; fi38++ {
                if true {
                    if true {
                        var v39 int8 = -1
                        fmt.Println(v4, v39, v5)
                        _ = v39
                    } else {
                        if true {
                            for fi40 := 0; fi40 < 19; fi40++ {
                                v1 += genFn1(false, v34, (fi38 / ((fi40) | 1)))
                                _ = fi40
                            }
                            v5 -= (v5 * (v5 ^ v5))
                        }
                        var v41 bool = (64 == 42)
                        _ = v41
                    }
                    v1 += (v1 | v1)
                    var v42 int = (^(fi38))
                    _ = v42
                } else {
                    v4 = v4
                }
                _ = fi38
            }
        case -2147483648:
            switch "\n\t" {
            case "hello":
                v43 := GenS3{F0: true, F1: uint8(v34), F2: (v5 % (((^(v5))) | 1))}
                _ = v43
            default:
                var v44 uint32 = 42
                _ = v44
            }
        case -668:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v1 = (^(v1))
        v45 := v2[(^(v5))]
        _ = v45
    } else {
        var v46 uint8 = 65
        v5 += uint32(v1)
        if false {
            if (true && true) {
                _ = len(v3)
            }
            for fi47 := 0; fi47 < 14; fi47++ {
                fmt.Println(v5, v46)
                _ = fi47
            }
        }
        _ = v46
    }
    _ = len(v2)
    fmt.Println(v34, v5)
    fmt.Println(v34)
    v5 += (v5 - 65535)
    var v48 bool = false
    v2[v5] = -9223372036854775808
    v5 = (v5 - (v5 << ((v34) & 63)))
    v49 := v2[2]
    var v50 uint32 = (v5 ^ 2147483648)
    v51 := v2[4294967295]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
