package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 string; F2 int64 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint32 }
func genFn0(p0 uint, p1 string, p2 int8) uint {
    var v4 int32 = 556
    v5 := []int{-9223372036854775808}
    var v6 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 uint16) uint {
    p0 += p0
    v7 := []int{32767, 4294967295}
    switch "foo bar" {
    case "a":
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 bool = true
            _ = fi8
            _ = v9
        }
    case "foo bar":
        var v10 int = 32768
        _ = v10
    case "hello":
        v11 := GenS1{F0: 286}
        _ = v11
    }
    var v12 int8 = -7
    var v13 uint16 = 1
    _ = p0
    _ = v12
    _ = v13
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = -2.25
    v3 := make(map[string]string)
    v1[0] = -861
    if (145 == 0) {
        v2 = v2
    } else {
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 float64 = 0.1
            v16 := GenS1{F0: 1}
            _ = fi14
            _ = v15
            _ = v16
        }
        switch -436 {
        case 127:
            if (-1 < -814) {
                switch ("" + ((("a" + ("a" + "")) + "foo bar") + "\n\t")) {
                case "foo bar":
                    for fi17 := 0; fi17 < 2; fi17++ {
                        v18 := GenS1{F0: 42}
                        _ = fi17
                        _ = v18
                    }
                case "\n\t":
                    v3["\n\t"] = "foo bar"
                case "hello":
                    var v19 uint64 = 1
                    _ = v19
                }
                fmt.Println(v2)
                v20 := GenS1{F0: 642}
                _ = v20
            }
        case 355:
            if (-670 < 0) {
                v2 = 0.1
                v2 += v2
                switch ("foo bar" + "hello") {
                case "":
                    delete(v3, "")
                case "\n\t":
                    v2 += (float64(float32(v2)) + v2)
                case "hello":
                    var v21 bool = true
                    _ = v21
                default:
                    if true {
                        if false {
                            for fi22 := 0; fi22 < 11; fi22++ {
                                v2 = (v2 - v2)
                                fmt.Println(fi22)
                                _ = fi22
                            }
                        }
                    } else {
                        v23 := GenS1{F0: -1}
                        v24 := GenS1{F0: -892}
                        for fi25 := 0; fi25 < 1; fi25++ {
                            v26 := GenS1{F0: -2147483648}
                            for fi27 := 0; fi27 < 7; fi27++ {
                                switch -202 {
                                case 2147483647:
                                    switch 9223372036854775807 {
                                    case 4294967295:
                                        var v28 int = fi25
                                        _ = v28
                                    case -9223372036854775808:
                                        v2 = float64(-2.25)
                                    case -1:
                                        if true {
                                            var v29 float32 = 3.14159
                                            v30 := []int{fi27, 9223372036854775807, (fi27 >> ((255) & 63))}
                                            _ = v29
                                            _ = v30
                                        }
                                    default:
                                        fmt.Println(fi27, fi25, fi25)
                                    }
                                    v2 += (float64(0.1) * 3.14159)
                                    v31 := GenS2{F0: "a", F1: "\n\t", F2: 845}
                                    _ = v31
                                case 298:
                                    if false {
                                        if false {
                                            v2 = float64((3.14159 + (1e10 + (3.14159 + 3.14159))))
                                            fmt.Println(fi27, fi25)
                                            fmt.Println(fi25, fi27, fi27)
                                        } else {
                                            fmt.Println(v2)
                                        }
                                        v2 += v2
                                        switch (fi25 / (((fi27 | (fi27 - fi25))) | 1)) {
                                        case 1:
                                            v2 = float64(-2.25)
                                        default:
                                            fmt.Println(fi25, v2, fi25)
                                        }
                                    }
                                case 4294967295:
                                    v2 += 1e10
                                }
                                _ = fi27
                            }
                            for fi32 := 0; fi32 < 5; fi32++ {
                                v33 := GenS2{F0: "\n\t", F1: ("" + ""), F2: 289}
                                v2 += v2
                                v34 := GenS3{F0: 2147483647, F1: 1}
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                            _ = fi25
                            _ = v26
                        }
                        _ = v23
                        _ = v24
                    }
                }
            } else {
                fmt.Println(v2)
            }
        case -831:
            fmt.Println(v2)
        }
    }
    var v35 uint8 = 101
    var v36 uint64 = 4294967295
    for fi37 := 0; fi37 < 1; fi37++ {
        var v38 uint64 = 18446744073709551615
        fmt.Println(v36)
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 6; fi39++ {
        for fi40 := 0; fi40 < 6; fi40++ {
            v35 += 255
            v41 := GenS1{F0: -218}
            if ((v36 < 596) && true) {
                fmt.Println(fi40, fi40, v35)
                fmt.Println(v2, v35)
                v35 += 132
                var v42 int32 = 42
                _ = v42
            }
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    v43 := GenS1{F0: -878}
    switch 2147483647 {
    case 4294967295:
        v36 += (v36 + (v36 ^ 922))
    case -1:
        for fi44 := 0; fi44 < 4; fi44++ {
            v45 := []int{fi44, (^(fi44)), (^(fi44))}
            v46 := GenS4{F0: (v36 - v36), F1: int8(v36), F2: 2}
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    _ = len(v1)
    var v47 uint16 = 1
    v48 := GenS1{F0: -1}
    for fi49 := 0; fi49 < 16; fi49++ {
        fmt.Println(v47, fi49, v2)
        for fi50 := 0; fi50 < 2; fi50++ {
            var v51 float32 = 1e10
            if true {
                if false {
                    v47 = (v47 + (v47 | (v47 << ((18446744073709551615) & 63))))
                }
                v52 := []int{127, (fi50 * (fi50 + fi50))}
                v36 += v36
                v2 = (3.14159 - (v2 / float64(float32(((float64(float32(v2)) + v2) * v2)))))
                _ = v52
            }
            v2 = v2
            _ = fi50
            _ = v51
        }
        _ = fi49
    }
    if (-1 != 42) {
        v36 = v36
        v36 = v36
    }
    var v53 bool = false
    if (!((-128 < 1) || (1 == -988))) {
        var v54 int16 = -385
        _ = v54
    } else {
        for fi55 := 0; fi55 < 19; fi55++ {
            v53 = (false || (!((!true) && ((uint32(v36) >= 443) || v53))))
            var v56 uint = genFn0(261, "hello", -1)
            _ = fi55
            _ = v56
        }
        v57 := GenS2{F0: "foo bar", F1: ("\n\t" + "\n\t"), F2: 42}
        v58 := []int{0}
        _ = v57
        _ = v58
    }
    v59 := v1[2147483647]
    if (584 > 65535) {
        v2 = 1e10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v47:", v47)
    fmt.Println("v53:", v53)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v43
    _ = v47
    _ = v48
    _ = v53
    _ = v59
}
