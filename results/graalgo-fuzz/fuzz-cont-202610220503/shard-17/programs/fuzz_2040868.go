package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 string; F1 uint64; F2 int16 }
func genFn0(p0 uint8, p1 int8, p2 uint8) uint64 {
    v4 := GenS1{F0: -2147483648, F1: 13}
    for fi5 := 0; fi5 < 16; fi5++ {
        p0 -= p0
        var v6 int32 = 0
        fmt.Println(p2, p0)
        _ = fi5
        _ = v6
    }
    if (335 != 1) {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int8 = 127
            p0 = p2
            _ = fi7
            _ = v8
        }
        if ((4294967295 <= 4294967295) || (!(!true))) {
            v9 := GenS1{F0: 1, F1: 4294967295}
            _ = v9
        }
        var v10 int32 = 0
        _ = v10
    } else {
        p1 -= p1
        v11 := []int{2147483647}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}
func genFn1() int16 {
    var v12 int16 = 0
    v12 = -250
    for fi13 := 0; fi13 < 3; fi13++ {
        v12 += int16(fi13)
        v14 := []int{fi13, 128, fi13}
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = fi13
        _ = v14
    }
    _ = v12
    return (v12 * (v12 & (v12 % (((v12 >> ((4294967295) & 63))) | 1))))
}

func main() {
    var v1 uint8 = 2
    var v2 int8 = 52
    v3 := make(map[uint32]int)
    var v15 uint16 = 42
    v3[116] = 32767
    v15 += (v15 / (((v15 & v15)) | 1))
    var v16 int = 4294967295
    if false {
        v16 -= v16
    }
    switch ("a" + "a") {
    case "foo bar":
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 uint16 = (v15 ^ (^(v15)))
                _ = fi18
                _ = v19
            }
            switch "hello" {
            case "foo bar":
                v16 = v16
            case "hello":
                fmt.Println(v15, v1, v1)
            }
            if (!false) {
                var v20 bool = (((true || false) && ((!(v16 == (^(fi17)))) && (0 != 1))) || true)
                v21 := GenS2{F0: -32768, F1: (true || v20)}
                switch "\n\t" {
                case "":
                    for fi22 := 0; fi22 < 17; fi22++ {
                        v23 := []int{}
                        if v20 {
                            v1 = v1
                        } else {
                            var v24 int = fi22
                            fmt.Println(fi17, v16)
                            v20 = (!(v20 || false))
                            var v25 int8 = (v2 & 0)
                            var v26 int64 = 42
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS1{F0: 1, F1: 255}
                        _ = fi22
                        _ = v23
                        _ = v27
                    }
                case "a":
                    var v28 float32 = 0.0
                    _ = v28
                case "hello":
                    v20 = false
                default:
                    fmt.Println(v16)
                }
                _ = v20
                _ = v21
            } else {
                var v29 uint8 = v1
                v2 = (^(v2))
                _ = v29
            }
            _ = fi17
        }
    }
    v30 := v3[1]
    switch "" {
    case "\n\t":
        fmt.Println(v15)
    case "foo bar":
        if true {
            if false {
                v31 := v3[2147483648]
                if true {
                    var v32 int = -2147483648
                    if true {
                        var v33 int32 = 1
                        if (42 > 42) {
                            for fi34 := 0; fi34 < 14; fi34++ {
                                v16 = v16
                                switch (-(v30)) {
                                case 128:
                                    v35 := v3[2147483648]
                                    _ = v35
                                case -272:
                                    if true {
                                        if (65535 != 152) {
                                            var v36 uint = 910
                                            _ = v36
                                        } else {
                                            if (532 < 2) {
                                                fmt.Println(v33, v30, v30)
                                            }
                                            for fi37 := 0; fi37 < 13; fi37++ {
                                                v38 := GenS1{F0: (v33 & v33), F1: 42}
                                                _ = fi37
                                                _ = v38
                                            }
                                            v2 -= v2
                                        }
                                    } else {
                                        v39 := []int{(-(v16)), int(v15), v16}
                                        _ = v39
                                    }
                                case 42:
                                    var v40 uint16 = v15
                                    _ = v40
                                }
                                _ = fi34
                            }
                            v32 = v32
                            fmt.Println(v32, v2, v33)
                        }
                        v41 := v3[0]
                        _ = v33
                        _ = v41
                    } else {
                        v32 -= v31
                        if false {
                            v32 = 757
                            v31 = (v31 % ((4294967295) | 1))
                            v42 := v3[2147483648]
                            _ = v42
                        } else {
                            var v43 uint16 = v15
                            v43 -= (^(v15))
                            _ = v43
                        }
                    }
                    _ = v32
                }
                _ = v31
            }
            switch (("" + "") + "\n\t") {
            case "\n\t":
                v44 := GenS2{F0: 42, F1: true}
                _ = v44
            case "":
                v30 += v16
            case "a":
                fmt.Println(v15, v30, v16)
            default:
                var v45 bool = false
                _ = v45
            }
        }
    case "a":
        v46 := []int{v30, v16, (v30 & v16), (v16 >> ((65535) & 63))}
        _ = v46
    }
    if false {
        fmt.Println(v15)
        var v47 uint32 = 42
        _ = v47
    } else {
        var v48 uint32 = 42
        _ = v48
    }
    fmt.Println(v30)
    v49 := GenS3{F0: "", F1: genFn0(v1, 0, (v1 - v1)), F2: 1}
    for fi50 := 0; fi50 < 6; fi50++ {
        var v51 bool = false
        _ = fi50
        _ = v51
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v49
}
