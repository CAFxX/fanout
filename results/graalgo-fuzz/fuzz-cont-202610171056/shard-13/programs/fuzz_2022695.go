package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int16; F2 string }
func genFn0() uint64 {
    var v5 string = "hello"
    v6 := GenS1{F0: 43}
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint64, p1 string) int8 {
    if false {
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                p1 = p1
                _ = fi8
            }
            var v9 bool = (!((p0 | genFn0()) == (p0 + genFn0())))
            var v10 float64 = (((((1.5 - 0.0) * 0.1) + float64(1.5)) + ((float64(3.14159) - float64(3.14159)) * (0.0 + (3.14159 / 3.14159)))) + ((0.0 - (0.1 * (float64((-2.25 * 0.1)) * -2.25))) * 1.5))
            _ = fi7
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: 42}
        v12 := []int{-2147483648, 0, 1}
        _ = v11
        _ = v12
    } else {
        v13 := []int{}
        var v14 float64 = 1.5
        _ = v13
        _ = v14
    }
    var v15 uint64 = (p0 >> ((65535) & 63))
    _ = p0
    _ = p1
    _ = v15
    return 127
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 128
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]bool)
    for fi16 := 0; fi16 < 4; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            v2 += 167
            _ = fi17
        }
        var v18 uint64 = genFn0()
        _ = fi16
        _ = v18
    }
    v2 += 1
    v2 = v2
    v2 += v2
    if (("a" + "foo bar") <= "foo bar") {
        switch 9223372036854775807 {
        case 4294967295:
            if (false && (!("foo bar" != "foo bar"))) {
                for fi19 := 0; fi19 < 4; fi19++ {
                    _ = len(v3)
                    v2 = 237
                    if false {
                        v20 := GenS1{F0: 0}
                        _ = v20
                    }
                    _ = fi19
                }
                v2 = v2
                v21 := GenS1{F0: 807}
                _ = v21
            } else {
                v22 := v3[4294967295]
                var v23 int = -194
                switch (v22 + "hello") {
                case "hello":
                    switch (v23 | (v23 - 32767)) {
                    case 32767:
                        if true {
                            v1[(v23 >> ((42) & 63))] = 76
                            switch 32767 {
                            case 127:
                                fmt.Println(v2, v22)
                            case 128:
                                switch -1 {
                                case 954:
                                    if (481 != 65535) {
                                        if true {
                                            var v24 int = v23
                                            _ = v24
                                        } else {
                                            v23 -= v23
                                            for fi25 := 0; fi25 < 1; fi25++ {
                                                v23 -= v23
                                                v26 := GenS2{F0: 216, F1: -2147483648}
                                                for fi27 := 0; fi27 < 1; fi27++ {
                                                    if true {
                                                        if (18446744073709551615 >= 448) {
                                                            if false {
                                                                v28 := GenS1{F0: 0}
                                                                switch v23 {
                                                                case 9223372036854775807:
                                                                    var v29 bool = ((fi27 / (((fi27 | -9223372036854775808)) | 1)) <= fi27)
                                                                    _ = v29
                                                                case 42:
                                                                    v22 = (("\n\t" + v22) + (((v22 + (("" + "hello") + (v22 + v22))) + v22) + "foo bar"))
                                                                default:
                                                                    fmt.Println(v22, v22, fi25)
                                                                }
                                                                v30 := GenS3{F0: -1, F1: -1}
                                                                v2 = 127
                                                                v22 = "a"
                                                                _ = v28
                                                                _ = v30
                                                            } else {
                                                                v2 -= 127
                                                                var v31 string = v22
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    if ((-(v23)) != 42) {
                                                                        var v32 int32 = 1
                                                                        v32 = (v32 | v32)
                                                                        _ = v32
                                                                    } else {
                                                                        v2 -= (^(v2))
                                                                        v33 := []int{619, fi27}
                                                                        v34 := v3[65535]
                                                                        _ = v33
                                                                        _ = v34
                                                                    }
                                                                case "hello":
                                                                    var v35 uint64 = 4294967295
                                                                    _ = v35
                                                                default:
                                                                    v2 -= v2
                                                                }
                                                                var v36 int = int(v2)
                                                                if false {
                                                                    fmt.Println(v36, v2)
                                                                }
                                                                _ = v31
                                                                _ = v36
                                                            }
                                                            v37 := GenS4{F0: 42, F1: uint64(v2)}
                                                            v4[0] = true
                                                            _ = v37
                                                        }
                                                    }
                                                    v38 := GenS4{F0: 4294967295, F1: genFn0()}
                                                    _ = len(v3)
                                                    _ = fi27
                                                    _ = v38
                                                }
                                                _ = fi25
                                                _ = v26
                                            }
                                            for fi39 := 0; fi39 < 9; fi39++ {
                                                v40 := []int{(fi39 & fi39), (v23 & (v23 << ((65535) & 63)))}
                                                if len(v40) > 0 {
                                                    v40[0] = 9223372036854775807
                                                }
                                                v41 := GenS5{F0: 0}
                                                _ = fi39
                                                _ = v40
                                                _ = v41
                                            }
                                        }
                                        v42 := []int{}
                                        _ = v42
                                    }
                                case 9223372036854775807:
                                    var v43 uint64 = 4294967295
                                    _ = v43
                                case 32768:
                                    var v44 bool = (310 <= 0)
                                    _ = v44
                                }
                            }
                        } else {
                            var v45 bool = false
                            v46 := GenS6{F0: 127, F1: 0, F2: v22}
                            _ = v45
                            _ = v46
                        }
                    }
                }
                _ = v22
                _ = v23
            }
        case -657:
            var v47 bool = false
            _ = v47
        case 68:
            v48 := GenS5{F0: 1}
            _ = v48
        }
        for fi49 := 0; fi49 < 11; fi49++ {
            v2 -= v2
            _ = fi49
        }
    }
    var v50 int8 = genFn1(genFn0(), "a")
    var v51 uint = 18446744073709551615
    if ((-173 <= 0) && (-212 >= 32767)) {
        v52 := []int{-2147483648, 4294967295}
        _ = v52
    }
    var v53 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v53
}
