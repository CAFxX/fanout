package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int32 }
type GenS2 struct { F0 int16; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0() bool {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: -108, F1: 127, F2: 0}
        var v9 uint = 255
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int64 = 42
    v10 = 9223372036854775807
    v10 -= (v10 - 989)
    v11 := GenS1{F0: 127, F1: 42, F2: 0}
    _ = v10
    _ = v11
    return false
}
func genFn1(p0 int16, p1 string, p2 int64) bool {
    p1 = "hello"
    var v12 int32 = 42
    if false {
        if true {
            p2 = 42
        }
    } else {
        v13 := GenS1{F0: 34, F1: 127, F2: v12}
        v14 := []int{32767}
        switch 32767 {
        case 9223372036854775807:
            p2 += p2
        case 0:
            var v15 uint16 = 770
            _ = v15
        }
        _ = v13
        _ = v14
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        p1 = (("hello" + p1) + p1)
        switch (fi16 | -2147483648) {
        case 42:
            v17 := GenS1{F0: 100, F1: 42, F2: 42}
            _ = v17
        case -2147483648:
            if false {
                if ((p0 % (((p0 | (p0 - p0))) | 1)) < 42) {
                    fmt.Println(v12)
                    p2 = (p2 - -9223372036854775808)
                    var v18 uint64 = 1
                    _ = v18
                }
                v12 = 986
                for fi19 := 0; fi19 < 7; fi19++ {
                    var v20 int32 = 0
                    switch fi16 {
                    case 42:
                        v21 := []int{(fi16 / ((fi19) | 1)), fi16, fi16, fi16}
                        _ = v21
                    default:
                        var v22 bool = false
                        _ = v22
                    }
                    _ = fi19
                    _ = v20
                }
            } else {
                switch fi16 {
                case 4294967295:
                    p0 = p0
                case -72:
                    p2 += 415
                default:
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v24 := []int{400, (fi16 >> ((796) & 63)), -2147483648, (-(fi23))}
                        _ = fi23
                        _ = v24
                    }
                }
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 int64 = (p2 >> ((893) & 63))
                    var v27 float64 = 0.1
                    if (!true) {
                        v27 = 0.1
                        v28 := []int{}
                        var v29 int64 = -816
                        _ = v28
                        _ = v29
                    } else {
                        v30 := []int{fi16}
                        switch p1 {
                        case "\n\t":
                            var v31 uint = 235
                            _ = v31
                        case "hello":
                            switch (p1 + ((("foo bar" + (p1 + p1)) + p1) + p1)) {
                            case "\n\t":
                                if false {
                                    if (((p1 + ("\n\t" + "hello")) + "\n\t") < ("\n\t" + p1)) {
                                        v32 := GenS2{F0: p0, F1: 0, F2: 65535}
                                        v30 = append(v30, int(p0))
                                        v27 = (float64(1.5) - v27)
                                        _ = v32
                                    }
                                    switch 555 {
                                    case 9223372036854775807:
                                        if true {
                                            p2 = 252
                                            v33 := GenS1{F0: 127, F1: 20, F2: v12}
                                            _ = v33
                                        } else {
                                            if len(v30) >= 2 {
                                                v30 = v30[1:len(v30)]
                                            }
                                            var v34 int32 = v12
                                            _ = len(v30)
                                            _ = v34
                                        }
                                    case -1:
                                        var v35 string = ((p1 + "foo bar") + ("" + "hello"))
                                        _ = v35
                                    }
                                } else {
                                    p1 = ((("" + p1) + "hello") + p1)
                                    v30 = append(v30, int(p0))
                                    fmt.Println(v12)
                                }
                            }
                        default:
                            v27 = (v27 - (float64(3.14159) * float64(((1.5 / (0.1 - -2.25)) + (-2.25 * (((float32(v27) - 3.14159) + 3.14159) / (((((((0.0 * (0.0 / ((float32(v27) - 0.0) * 1e10))) * 0.1) * (3.14159 / (1.5 - ((-2.25 * float32(0.0)) - 1.5)))) + -2.25) * ((0.1 + (1.5 / 1.5)) / 1e10)) * 3.14159) / ((1.5 - 3.14159) / ((3.14159 + 0.1) * float32(-2.25))))))))))
                        }
                        _ = v30
                    }
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                var v36 int32 = v12
                _ = v36
            }
        case 776:
            fmt.Println(fi16, p1, p1)
        }
        switch fi16 {
        case 128:
            v37 := []int{(fi16 / ((fi16) | 1)), fi16}
            _ = v37
        case -919:
            var v38 uint8 = 128
            _ = v38
        }
        _ = fi16
    }
    v12 = v12
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return false
}

func main() {
    var v1 int = -944
    v2 := make(map[uint32]int)
    var v3 uint16 = 32767
    var v4 int64 = -630
    v5 := make(map[int]uint32)
    v6 := make(map[int]string)
    for fi39 := 0; fi39 < 17; fi39++ {
        v40 := GenS2{F0: -32768, F1: 42, F2: v3}
        _ = fi39
        _ = v40
    }
    var v41 uint8 = 128
    delete(v2, 4294967295)
    v42 := GenS3{F0: 0, F1: (v3 / ((255) | 1))}
    v43 := GenS3{F0: 123, F1: 84}
    v41 -= (v41 - v41)
    var v44 int8 = 42
    fmt.Println(v41)
    switch (v1 + 4294967295) {
    case -1:
        v45 := v5[v1]
        _ = v45
    case 42:
        v1 = (-(v1))
    case -847:
        var v46 float64 = (float64(0.0) + (0.0 - ((0.1 / float64(0.1)) / float64(3.14159))))
        _ = v46
    }
    var v47 float64 = -2.25
    if true {
        var v48 bool = true
        _ = v48
    } else {
        var v49 int16 = -1
        _ = v49
    }
    fmt.Println(v41, v3, v1)
    if (true || true) {
        if false {
            for fi50 := 0; fi50 < 19; fi50++ {
                if ((-2.25 / 1.5) >= 0.1) {
                    for fi51 := 0; fi51 < 5; fi51++ {
                        switch v1 {
                        case -127:
                            var v52 int32 = 668
                            _ = v52
                        default:
                            v3 = v3
                        }
                        if (true && true) {
                            if (419 == 2) {
                                var v53 bool = ((v4 | 9223372036854775807) == v4)
                                _ = v53
                            }
                            v47 = (v47 - v47)
                            v54 := []int{}
                            _ = v54
                        } else {
                            v55 := GenS1{F0: (v44 * (v44 << ((42) & 63))), F1: int8(v1), F2: 2147483647}
                            _ = v55
                        }
                        var v56 int8 = v44
                        _ = fi51
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 16; fi57++ {
                        var v58 bool = (725 == 0)
                        v59 := []int{fi57, (fi57 << ((905) & 63)), -9223372036854775808}
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                }
                var v60 int16 = 0
                _ = fi50
                _ = v60
            }
            var v61 uint16 = (v3 % (((v3 * (v3 + (v3 / (((v3 * v3)) | 1))))) | 1))
            for fi62 := 0; fi62 < 7; fi62++ {
                if true {
                    var v63 uint8 = v41
                    _ = v63
                }
                v64 := GenS3{F0: 587, F1: v3}
                v47 = 0.0
                _ = fi62
                _ = v64
            }
            _ = v61
        } else {
            if ("foo bar" > ("" + (("a" + ("\n\t" + ((("foo bar" + (("\n\t" + "foo bar") + ("foo bar" + "foo bar"))) + ("" + "hello")) + (("" + ("\n\t" + "")) + "")))) + ("\n\t" + "foo bar")))) {
                v4 = (v4 & (-(v4)))
                v65 := []int{}
                _ = v65
            } else {
                var v66 int64 = 1
                _ = v66
            }
        }
        var v67 float64 = 0.0
        v68 := GenS3{F0: 616, F1: v3}
        _ = v67
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v47
    _ = v5
    _ = v6
}
