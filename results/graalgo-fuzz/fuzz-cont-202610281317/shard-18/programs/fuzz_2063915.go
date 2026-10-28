package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint8; F2 int }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint64, p1 string, p2 uint32) int64 {
    var v4 int16 = 32767
    v5 := []int{1, -1, 4294967295}
    var v6 uint16 = 42
    if (525 != (^(p2))) {
        var v7 int64 = int64(v4)
        var v8 uint16 = v6
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: -1, F1: fi9, F2: uint(p2)}
            _ = fi9
            _ = v10
        }
        if (!(!(209 == 9223372036854775807))) {
            fmt.Println(v6)
            fmt.Println(v6)
        }
        v6 += 65535
    }
    var v11 float64 = ((0.1 / -2.25) / 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v4
    _ = v5
    _ = v6
    return 341
}
func genFn1(p0 uint, p1 uint, p2 int) int32 {
    var v12 int = p2
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: 127, F1: (v12 << ((p0) & 63)), F2: p1}
        for fi15 := 0; fi15 < 16; fi15++ {
            if (1000 == 362) {
                if (true || false) {
                    v12 = (v12 % ((v12) | 1))
                    v12 = fi15
                    fmt.Println(fi15, v12, p2)
                }
                p2 += 32767
                var v16 uint64 = 2
                v17 := GenS1{F0: 100, F1: v12, F2: 1}
                _ = v16
                _ = v17
            }
            var v18 int = (p2 % ((32767) | 1))
            p2 -= v18
            _ = fi15
            _ = v18
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 0
}

func main() {
    var v1 float64 = (-2.25 * 0.0)
    var v2 bool = false
    var v3 string = ("a" + "a")
    v1 = (float64(0.0) / float64(float32(0.1)))
    if v2 {
        var v19 uint = 65535
        fmt.Println(v2, v19)
        v19 -= v19
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        switch fi20 {
        case 9223372036854775807:
            var v21 uint32 = 4294967295
            _ = v21
        default:
            v2 = v2
        }
        v2 = false
        for fi22 := 0; fi22 < 19; fi22++ {
            fmt.Println(v2, fi20)
            v1 -= (1e10 + v1)
            v23 := GenS1{F0: -1, F1: (fi20 / ((fi22) | 1)), F2: 515}
            _ = fi22
            _ = v23
        }
        _ = fi20
    }
    v2 = (42 < 2147483648)
    for fi24 := 0; fi24 < 2; fi24++ {
        for fi25 := 0; fi25 < 17; fi25++ {
            fmt.Println(fi25, fi24)
            _ = fi25
        }
        if v2 {
            var v26 uint = 2
            if (4294967295 == 185) {
                if false {
                    v27 := []int{(fi24 << ((v26) & 63)), 700, (fi24 + 127)}
                    if (0 < 42) {
                        switch v3 {
                        case "\n\t":
                            v28 := GenS2{F0: 89}
                            _ = v28
                        }
                        var v29 int64 = genFn0(183, v3, 42)
                        if v2 {
                            switch v3 {
                            case "":
                                if len(v27) > 0 {
                                    v27[0] = fi24
                                }
                            case "foo bar":
                                for fi30 := 0; fi30 < 19; fi30++ {
                                    v26 -= v26
                                    switch (v3 + ((v3 + v3) + ((v3 + "") + v3))) {
                                    case "":
                                        v29 -= (v29 - (v29 | v29))
                                    }
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                    v31 := GenS1{F0: 105, F1: (fi24 >> ((2) & 63)), F2: (v26 % ((v26) | 1))}
                                    fmt.Println(fi24, fi24)
                                    _ = fi30
                                    _ = v31
                                }
                            }
                            v32 := GenS3{F0: 0, F1: 0, F2: fi24}
                            _ = v32
                        } else {
                            v3 = v3
                            v29 = (v29 - (v29 ^ int64(fi24)))
                        }
                        _ = v29
                    } else {
                        var v33 string = (v3 + v3)
                        v3 = v3
                        _ = v33
                    }
                    if len(v27) > 0 {
                        v27[0] = (fi24 - 684)
                    }
                    _ = v27
                } else {
                    v1 = v1
                    for fi34 := 0; fi34 < 13; fi34++ {
                        v35 := GenS2{F0: 910}
                        _ = fi34
                        _ = v35
                    }
                }
                v3 = v3
            } else {
                v2 = (v2 || (32767 != 32767))
                if v2 {
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v2, v2)
                    case "\n\t":
                        for fi36 := 0; fi36 < 8; fi36++ {
                            fmt.Println(v1)
                            v1 = (float64(float32(v1)) + v1)
                            var v37 uint8 = 255
                            _ = fi36
                            _ = v37
                        }
                    case "foo bar":
                        if (((v3 + "") + (v3 + "hello")) > ((("" + ("hello" + v3)) + v3) + ("hello" + v3))) {
                            var v38 int32 = genFn1(980, v26, fi24)
                            v39 := []int{fi24, (fi24 | (fi24 ^ fi24)), fi24}
                            _ = v38
                            _ = v39
                        }
                    default:
                        var v40 uint8 = 255
                        _ = v40
                    }
                    v41 := GenS2{F0: 42}
                    _ = v41
                }
            }
            v42 := GenS4{F0: 1}
            _ = v26
            _ = v42
        } else {
            var v43 uint32 = 4294967295
            v44 := []int{}
            _ = v43
            _ = v44
        }
        _ = fi24
    }
    if (!v2) {
        switch 9223372036854775807 {
        case 0:
            v1 = ((1e10 + (((3.14159 + v1) / v1) - 3.14159)) * v1)
        }
        if v2 {
            v45 := GenS3{F0: -1, F1: 0, F2: 42}
            var v46 string = (("" + v3) + "hello")
            _ = v45
            _ = v46
        }
    }
    if (896 != 0) {
        var v47 float64 = v1
        var v48 uint = 255
        _ = v47
        _ = v48
    }
    var v49 string = (v3 + v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
}
