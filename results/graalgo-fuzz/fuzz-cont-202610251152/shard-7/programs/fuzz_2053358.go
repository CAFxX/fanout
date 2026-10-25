package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int32; F2 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 uint8; F1 uint8 }
func genFn0(p0 bool) string {
    if false {
        if (65535 <= 32767) {
            v7 := []int{-1, -406}
            var v8 string = "hello"
            _ = v7
            _ = v8
        } else {
            if (p0 || ((0 < -1) && (!(!p0)))) {
                v9 := GenS1{F0: -128}
                p0 = ((-325 >= 598) && (true && ((!(!((!p0) || (-2.25 <= 1e10)))) && (!((!(1e10 == 0.0)) && p0)))))
                _ = v9
            }
            fmt.Println(p0)
            var v10 int = -1
            _ = v10
        }
        if true {
            for fi11 := 0; fi11 < 19; fi11++ {
                p0 = p0
                var v12 uint8 = 128
                _ = fi11
                _ = v12
            }
        } else {
            var v13 uint32 = 1
            var v14 int = 128
            v15 := []int{42, 127, v14, (v14 % ((32767) | 1))}
            _ = v13
            _ = v14
            _ = v15
        }
        v16 := GenS2{F0: 1}
        _ = v16
    } else {
        if (42 > 1) {
            if (3 == 114) {
                v17 := []int{-2147483648, 32768, 42, -2147483648}
                _ = v17
            } else {
                if (p0 && p0) {
                    v18 := []int{9223372036854775807}
                    p0 = (p0 && (((1.5 - (0.1 * 3.14159)) - 1e10) >= (1.5 + 3.14159)))
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        v20 := []int{(^(fi19)), (fi19 ^ (fi19 + fi19)), (fi19 >> ((1) & 63)), 291}
                        _ = fi19
                        _ = v20
                    }
                    p0 = (2 <= 1)
                    var v21 uint = 18446744073709551615
                    _ = v21
                }
            }
            p0 = (32767 > 102)
            if (p0 || (128 >= -725)) {
                var v22 uint16 = 42
                var v23 uint32 = 621
                if (100 != 127) {
                    v24 := GenS3{F0: -1, F1: -1, F2: -32768}
                    v22 = v22
                    var v25 float32 = (1e10 + float32(1e10))
                    _ = v24
                    _ = v25
                } else {
                    v26 := []int{941, int(v23)}
                    var v27 float32 = (1.5 + 0.0)
                    _ = v26
                    _ = v27
                }
                _ = v22
                _ = v23
            }
        }
        if (1 != 4294967295) {
            if p0 {
                switch ("\n\t" + ((("a" + "a") + "hello") + "a")) {
                case "foo bar":
                    var v28 float32 = (1.5 - -2.25)
                    _ = v28
                }
                var v29 uint64 = 2
                if false {
                    v30 := GenS4{F0: (v29 / ((v29) | 1))}
                    fmt.Println(p0, v29)
                    var v31 float64 = (float64(-2.25) + 0.0)
                    _ = v30
                    _ = v31
                }
                _ = v29
            }
            if false {
                for fi32 := 0; fi32 < 16; fi32++ {
                    switch (fi32 ^ (fi32 | 32767)) {
                    case 688:
                        for fi33 := 0; fi33 < 8; fi33++ {
                            if (p0 && p0) {
                                var v34 uint16 = 67
                                _ = v34
                            }
                            fmt.Println(fi33, p0)
                            _ = fi33
                        }
                    case 142:
                        if p0 {
                            v35 := GenS5{F0: 255, F1: 166}
                            for fi36 := 0; fi36 < 18; fi36++ {
                                var v37 uint8 = 128
                                var v38 string = "\n\t"
                                v39 := GenS2{F0: -2147483648}
                                if (42 == 1) {
                                    if p0 {
                                        var v40 int8 = 42
                                        v38 = v38
                                        v41 := []int{636, -9223372036854775808}
                                        _ = v40
                                        _ = v41
                                    } else {
                                        v38 = "\n\t"
                                    }
                                    v42 := []int{(fi36 ^ fi32), (-(fi32)), 407, (fi32 * 32767)}
                                    _ = v42
                                } else {
                                    v43 := GenS5{F0: 32767, F1: v37}
                                    _ = v43
                                }
                                _ = fi36
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                            p0 = (("\n\t" + "a") <= "\n\t")
                            _ = v35
                        }
                    case 770:
                        switch 42 {
                        case 127:
                            var v44 uint32 = 65535
                            _ = v44
                        }
                    default:
                        fmt.Println(fi32, fi32)
                    }
                    v45 := GenS6{F0: 42}
                    _ = fi32
                    _ = v45
                }
            }
            fmt.Println(p0)
        }
        var v46 int = 4294967295
        _ = v46
    }
    var v47 int = 32768
    _ = p0
    _ = v47
    return "hello"
}

func main() {
    var v1 string = ("\n\t" + "hello")
    var v2 int8 = 127
    var v3 uint = 2
    v4 := []uint32{2, 42}
    v5 := make(map[int]string)
    var v6 uint64 = 2
    v2 += (v2 * (^(v2)))
    var v48 bool = true
    _ = len(v5)
    if (0 != 0) {
        for fi49 := 0; fi49 < 18; fi49++ {
            fmt.Println(v48)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi49
        }
        v50 := GenS7{F0: v48, F1: v6}
        _ = v50
    }
    var v51 int32 = 785
    if v48 {
        var v52 uint32 = 35
        v53 := v5[123]
        v1 = v53
        _ = v52
        _ = v53
    } else {
        var v54 int16 = 32767
        for fi55 := 0; fi55 < 2; fi55++ {
            if v48 {
                var v56 int32 = v51
                _ = v56
            } else {
                var v57 uint64 = (v6 / (((v6 + (v6 >> (((v3 | v3)) & 63)))) | 1))
                v58 := GenS7{F0: v48, F1: 0}
                v57 -= (v57 << ((uint(v57)) & 63))
                _ = v57
                _ = v58
            }
            if len(v4) > 0 {
                v4[0] = 631
            }
            _ = fi55
        }
        fmt.Println(v48, v2, v51)
        _ = v54
    }
    switch 0 {
    case -279:
        var v59 int32 = (v51 + 850)
        _ = v59
    case 652:
        v6 = (v6 & v6)
    }
    var v60 bool = v48
    for fi61 := 0; fi61 < 19; fi61++ {
        v62 := GenS5{F0: uint16(v3), F1: 93}
        _ = fi61
        _ = v62
    }
    v2 = v2
    switch ((v1 + v1) + (v1 + v1)) {
    case "hello":
        if (!v60) {
            v63 := GenS8{F0: 212, F1: 1}
            _ = v63
        }
    }
    var v64 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v51
    _ = v6
    _ = v60
    _ = v64
}
