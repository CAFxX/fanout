package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint; F2 string }
type GenS3 struct { F0 bool; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS5 struct { F0 string; F1 int16; F2 uint8 }
type GenS6 struct { F0 int8; F1 uint16; F2 uint }
func genFn0(p0 uint8, p1 int16, p2 uint32) int8 {
    var v7 uint32 = p2
    v8 := []int{4294967295, -819}
    for fi9 := 0; fi9 < 5; fi9++ {
        switch fi9 {
        case 32767:
            fmt.Println(p2)
        default:
            var v10 float32 = 0.0
            _ = v10
        }
        _ = fi9
    }
    fmt.Println(p0, p2, v7)
    v11 := GenS1{F0: (v7 / ((uint32(p0)) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{-943, 9223372036854775807, -2147483648, 32767}
    var v3 int32 = -706
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    v6 := make(map[uint32]string)
    for fi12 := 0; fi12 < 19; fi12++ {
        v3 += v3
        for fi13 := 0; fi13 < 8; fi13++ {
            if true {
                v14 := GenS2{F0: ((("hello" + "\n\t") + "hello") + "a"), F1: 42, F2: "a"}
                v3 = v3
                _ = v14
            } else {
                if true {
                    fmt.Println(fi13)
                    if (4294967295 == 4294967295) {
                        if (127 > (fi13 * -9223372036854775808)) {
                            v3 -= v3
                            if len(v2) > 0 {
                                v2[0] = fi12
                            }
                        } else {
                            var v15 int64 = 566
                            _ = v15
                        }
                        switch "hello" {
                        case "hello":
                            if (!true) {
                                v2 = append(v2, fi13)
                                if (float64(((((3.14159 * 3.14159) - (0.0 - 0.1)) * 3.14159) * ((((0.0 + ((float32(float64(0.0)) - 3.14159) + 3.14159)) / 1e10) * 1.5) + -2.25))) <= 3.14159) {
                                    var v16 uint = 936
                                    v3 = v3
                                    _ = v16
                                }
                                fmt.Println(fi13, fi12)
                            }
                        default:
                            if (!(42 > 2147483648)) {
                                v17 := v5[(fi13 * fi12)]
                                var v18 uint8 = 0
                                _ = v17
                                _ = v18
                            } else {
                                v19 := GenS3{F0: (-903 >= 42), F1: 21, F2: 609}
                                var v20 float32 = float32((0.0 * ((1e10 / ((float64((((3.14159 - 0.0) + (0.1 + 0.1)) * 1e10)) - float64((0.0 * -2.25))) * (0.0 + float64(-2.25)))) + 0.1)))
                                _ = v19
                                _ = v20
                            }
                        }
                    } else {
                        fmt.Println(fi13, v3, fi13)
                    }
                    if (-9223372036854775808 == -307) {
                        fmt.Println(fi12)
                    } else {
                        switch (fi12 + -701) {
                        case -2147483648:
                            _ = len(v5)
                        }
                        delete(v1, (fi12 % ((fi13) | 1)))
                        fmt.Println(fi12, v3, fi13)
                        v21 := GenS4{F0: 589, F1: 65535, F2: 42}
                        if true {
                            delete(v1, (fi13 / ((fi12) | 1)))
                            _ = len(v2)
                        } else {
                            var v22 uint32 = 556
                            v3 -= (v3 << ((977) & 63))
                            _ = v22
                        }
                        v23 := GenS3{F0: ((false || ((!false) && false)) && false), F1: 73, F2: 0}
                        _ = v21
                        _ = v23
                    }
                } else {
                    v24 := GenS4{F0: 42, F1: 1, F2: 139}
                    var v25 uint = 824
                    var v26 int16 = 42
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
            if (-26 != 20) {
                var v27 int16 = 0
                v3 += (v3 + (^(v3)))
                var v28 int8 = 42
                _ = v27
                _ = v28
            }
            v29 := GenS3{F0: false, F1: 0, F2: 914}
            _ = fi13
            _ = v29
        }
        _ = fi12
    }
    fmt.Println(v3)
    var v30 uint32 = 229
    var v31 int = -2147483648
    _ = len(v1)
    if true {
        var v32 float32 = float32((-2.25 / float64(1e10)))
        if true {
            for fi33 := 0; fi33 < 17; fi33++ {
                var v34 int16 = 32767
                var v35 float64 = float64(1.5)
                for fi36 := 0; fi36 < 2; fi36++ {
                    v37 := GenS5{F0: "foo bar", F1: -148, F2: 2}
                    _ = fi36
                    _ = v37
                }
                _ = fi33
                _ = v34
                _ = v35
            }
            v31 = (v31 + 2147483647)
        } else {
            v38 := GenS2{F0: ("" + ("foo bar" + (("a" + "a") + "hello"))), F1: 65535, F2: "a"}
            v32 = 1e10
            _ = v38
        }
        var v39 uint32 = (v30 + (v30 ^ v30))
        _ = v32
        _ = v39
    } else {
        v31 -= v31
    }
    var v40 int32 = v3
    if true {
        for fi41 := 0; fi41 < 8; fi41++ {
            if true {
                var v42 int16 = 1000
                switch 2147483647 {
                case -514:
                    if true {
                        v40 -= v3
                    } else {
                        var v43 float32 = (((-2.25 - 3.14159) * 1.5) / ((3.14159 / (((1e10 * 1e10) + 0.1) + ((3.14159 * 3.14159) / ((1e10 * 1e10) * 1.5)))) * 0.1))
                        if (-2147483648 > -1) {
                            if ((v30 >> ((255) & 63)) > (v30 * (v30 | 2147483648))) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v30 = v30
                                var v44 int = (v31 * 32768)
                                _ = v44
                            } else {
                                var v45 int = (fi41 * (fi41 % ((int(v42)) | 1)))
                                _ = v45
                            }
                            fmt.Println(v3, fi41, v42)
                            for fi46 := 0; fi46 < 5; fi46++ {
                                var v47 bool = true
                                v48 := GenS4{F0: 255, F1: 244, F2: 87}
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                        }
                        _ = v43
                    }
                case 4294967295:
                    var v49 uint16 = 1
                    _ = v49
                }
                _ = v42
            }
            for fi50 := 0; fi50 < 17; fi50++ {
                _ = len(v2)
                if false {
                    v40 = (v40 ^ v40)
                    v30 = 676
                }
                v51 := GenS4{F0: 61, F1: 18446744073709551615, F2: 85}
                _ = fi50
                _ = v51
            }
            _ = fi41
        }
        fmt.Println(v40, v30, v3)
        fmt.Println(v40)
    } else {
        v30 += (v30 ^ 873)
        for fi52 := 0; fi52 < 12; fi52++ {
            var v53 int32 = 1
            for fi54 := 0; fi54 < 20; fi54++ {
                switch "a" {
                case "foo bar":
                    var v55 uint8 = 0
                    _ = v55
                case "a":
                    var v56 bool = true
                    _ = v56
                }
                v53 -= (v53 & 0)
                _ = fi54
            }
            for fi57 := 0; fi57 < 2; fi57++ {
                var v58 int = (fi57 / (((-(v31))) | 1))
                v59 := GenS6{F0: genFn0(2, 1, (v30 >> ((65535) & 63))), F1: 0, F2: 456}
                fmt.Println(fi52)
                v53 = v53
                _ = fi57
                _ = v58
                _ = v59
            }
            _ = fi52
            _ = v53
        }
    }
    var v60 int8 = 42
    var v61 uint = 411
    if true {
        v62 := GenS2{F0: "", F1: (v61 | v61), F2: ((("" + "") + "a") + ("hello" + ("hello" + "foo bar")))}
        _ = v62
    } else {
        var v63 float64 = -2.25
        _ = len(v5)
        _ = v63
    }
    for fi64 := 0; fi64 < 12; fi64++ {
        v60 -= 100
        v61 -= v61
        _ = fi64
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
