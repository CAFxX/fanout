package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int8; F2 uint64 }
func genFn0(p0 uint) uint64 {
    if true {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int = (fi7 % ((fi7) | 1))
            var v9 uint8 = 255
            if (1 == -84) {
                for fi10 := 0; fi10 < 20; fi10++ {
                    if true {
                        v11 := GenS1{F0: 18446744073709551615, F1: v8}
                        var v12 uint32 = 65535
                        fmt.Println(fi10, fi7, fi10)
                        if ("a" > "foo bar") {
                            v13 := []int{259}
                            var v14 int = (fi7 >> ((p0) & 63))
                            p0 = (p0 % ((p0) | 1))
                            _ = v13
                            _ = v14
                        }
                        _ = v11
                        _ = v12
                    } else {
                        fmt.Println(v8)
                        if false {
                            v15 := GenS1{F0: 921, F1: 2147483647}
                            _ = v15
                        }
                        v16 := []int{(fi10 << ((65535) & 63)), fi10}
                        _ = v16
                    }
                    _ = fi10
                }
                v8 = 128
            } else {
                v9 = 0
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v17 := GenS2{F0: 2147483648}
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 bool = (false && true)
        var v20 float32 = 1.5
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 float64 = 0.1
    switch ("" + ("hello" + (("foo bar" + ("a" + "a")) + "a"))) {
    case "\n\t":
        fmt.Println(v21, p0)
    case "foo bar":
        for fi22 := 0; fi22 < 7; fi22++ {
            switch (^(fi22)) {
            case 8:
                v23 := []int{fi22, (fi22 / ((fi22) | 1)), fi22}
                _ = v23
            case 32767:
                v24 := GenS2{F0: 893}
                _ = v24
            case 796:
                for fi25 := 0; fi25 < 4; fi25++ {
                    switch (fi22 & (fi25 & (fi25 | (-(fi25))))) {
                    case 1:
                        fmt.Println(p0, fi25)
                    case 9223372036854775807:
                        if (false || (!(-873 >= -700))) {
                            var v26 int8 = 100
                            _ = v26
                        } else {
                            fmt.Println(v21, fi25)
                            v21 = (float64(-2.25) - v21)
                            fmt.Println(v21, fi25)
                        }
                    default:
                        var v27 uint32 = 2147483648
                        _ = v27
                    }
                    var v28 uint64 = 4294967295
                    fmt.Println(fi22, p0, v28)
                    _ = fi25
                    _ = v28
                }
            default:
                v29 := GenS1{F0: 1, F1: -9223372036854775808}
                _ = v29
            }
            for fi30 := 0; fi30 < 10; fi30++ {
                v31 := GenS3{F0: p0, F1: 0, F2: 185}
                if true {
                    var v32 string = "\n\t"
                    switch v32 {
                    case "\n\t":
                        if false {
                            if false {
                                var v33 float64 = ((v21 + float64(1e10)) - (((v21 + v21) * v21) - float64(1.5)))
                                v33 += (float64(float32(v33)) / ((v33 * float64(0.1)) * 1.5))
                                var v34 uint16 = 65535
                                _ = v33
                                _ = v34
                            } else {
                                var v35 int8 = 8
                                var v36 uint8 = 42
                                _ = v35
                                _ = v36
                            }
                        } else {
                            if (740 > 42) {
                                p0 = 268
                                v21 += ((1e10 / (float64(float32(v21)) + v21)) + (1e10 - v21))
                                v32 = "hello"
                            } else {
                                v32 = "a"
                                switch ("\n\t" + "hello") {
                                case "":
                                    if ((fi30 | fi30) == fi22) {
                                        if false {
                                            p0 = (p0 ^ p0)
                                            var v37 string = v32
                                            _ = v37
                                        } else {
                                            v38 := []int{(fi30 ^ fi30), (^(fi22)), fi30}
                                            _ = v38
                                        }
                                        fmt.Println(v32, v32, p0)
                                    } else {
                                        var v39 int8 = 87
                                        _ = v39
                                    }
                                case "foo bar":
                                    v40 := GenS1{F0: 1, F1: (fi30 << ((p0) & 63))}
                                    _ = v40
                                case "\n\t":
                                    var v41 int8 = int8(fi30)
                                    _ = v41
                                }
                                var v42 uint16 = 42
                                _ = v42
                            }
                        }
                    case "":
                        var v43 uint64 = 2
                        _ = v43
                    case "foo bar":
                        v44 := GenS2{F0: 0}
                        _ = v44
                    }
                    _ = v32
                } else {
                    var v45 string = ("\n\t" + "")
                    v46 := []int{fi22, (fi30 / ((fi22) | 1)), fi30}
                    _ = v45
                    _ = v46
                }
                v47 := []int{(fi30 - (fi22 + (fi30 | (fi30 % ((fi30) | 1)))))}
                _ = fi30
                _ = v31
                _ = v47
            }
            _ = fi22
        }
    }
    _ = p0
    _ = v17
    _ = v21
    return 18446744073709551615
}
func genFn1(p0 int16) int8 {
    p0 = p0
    p0 = (^(p0))
    if false {
        fmt.Println(p0)
        p0 += p0
        var v48 bool = (-80 <= -1)
        _ = v48
    } else {
        p0 += 42
        for fi49 := 0; fi49 < 8; fi49++ {
            fmt.Println(p0, p0)
            _ = fi49
        }
    }
    var v50 uint64 = 2
    var v51 int64 = 42
    _ = p0
    _ = v50
    _ = v51
    return 42
}

func main() {
    var v1 string = ""
    var v2 int8 = 0
    v3 := []int{686, 42, 9223372036854775807, -2147483648}
    v4 := make(map[uint32]string)
    v5 := make(map[int]string)
    var v6 uint16 = 2
    fmt.Println(v1, v6)
    var v52 uint16 = v6
    for fi53 := 0; fi53 < 12; fi53++ {
        var v54 uint32 = 0
        _ = len(v4)
        _ = fi53
        _ = v54
    }
    for fi55 := 0; fi55 < 4; fi55++ {
        var v56 int64 = 39
        _ = len(v3)
        v2 -= v2
        _ = fi55
        _ = v56
    }
    fmt.Println(v52)
    var v57 uint16 = 32767
    fmt.Println(v6)
    v58 := v5[-1]
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    var v59 uint16 = (v52 * v57)
    var v60 int8 = v2
    if (true && false) {
        switch 448 {
        case 32767:
            switch 9223372036854775807 {
            case 42:
                if false {
                    var v61 uint64 = 18446744073709551615
                    var v62 uint64 = 4294967295
                    v1 = v1
                    _ = v61
                    _ = v62
                } else {
                    v63 := GenS1{F0: 631, F1: 548}
                    _ = v63
                }
            case 127:
                switch v1 {
                case "":
                    if (true && false) {
                        for fi64 := 0; fi64 < 5; fi64++ {
                            for fi65 := 0; fi65 < 18; fi65++ {
                                switch fi64 {
                                case 758:
                                    fmt.Println(fi64, fi65, v57)
                                case 193:
                                    v6 -= (v59 >> ((0) & 63))
                                case 2147483647:
                                    switch (-(fi64)) {
                                    case 128:
                                        v66 := v4[65535]
                                        _ = v66
                                    default:
                                        v60 = genFn1(-495)
                                    }
                                }
                                var v67 bool = true
                                _ = fi65
                                _ = v67
                            }
                            v60 = (v2 / (((v2 % (((v60 >> ((0) & 63))) | 1))) | 1))
                            _ = fi64
                        }
                        v68 := GenS2{F0: 1}
                        v1 = v58
                        _ = v68
                    }
                case "\n\t":
                    v52 += 32767
                case "a":
                    v4[4294967295] = "\n\t"
                }
            }
        case -1:
            v69 := v5[127]
            _ = v69
        default:
            switch v58 {
            case "\n\t":
                var v70 int16 = 1
                _ = v70
            default:
                if true {
                    v57 += v57
                }
            }
        }
        v71 := GenS1{F0: 2, F1: -9223372036854775808}
        if (!(1 < 0)) {
            var v72 bool = true
            var v73 uint16 = (v52 % (((v59 % (((v57 ^ v57)) | 1))) | 1))
            _ = v72
            _ = v73
        }
        _ = v71
    }
    var v74 float64 = 0.0
    v75 := GenS1{F0: 42, F1: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v74
    _ = v75
}
