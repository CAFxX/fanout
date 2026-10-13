package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint }
func genFn0(p0 int16, p1 uint64) int {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int16 = -1
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 uint64 = (^(p1))
            switch fi9 {
            case 954:
                if true {
                    var v11 int8 = 127
                    v12 := GenS1{F0: false, F1: 1}
                    _ = v11
                    _ = v12
                }
            case 9223372036854775807:
                v8 += v8
            case 32768:
                p0 += (v8 << ((65535) & 63))
            }
            switch "\n\t" {
            case "foo bar":
                v13 := []int{(^(fi9)), (fi9 * 9223372036854775807), (fi9 + fi7)}
                _ = v13
            }
            _ = fi9
            _ = v10
        }
        if false {
            switch fi7 {
            case -9223372036854775808:
                var v14 uint64 = p1
                _ = v14
            }
            v15 := GenS2{F0: 186}
            _ = v15
        } else {
            var v16 float64 = 3.14159
            fmt.Println(fi7, p0)
            v16 -= (v16 / v16)
            _ = v16
        }
        _ = fi7
        _ = v8
    }
    var v17 int32 = 42
    _ = p0
    _ = p1
    _ = v17
    return 127
}
func genFn1() uint8 {
    if (((("" + "a") + "") + ("a" + "\n\t")) < "foo bar") {
        if false {
            var v18 uint = 0
            if true {
                if false {
                    v18 += v18
                    fmt.Println(v18)
                } else {
                    var v19 int16 = 1000
                    if true {
                        v18 -= v18
                        var v20 float32 = 1.5
                        var v21 bool = (-2147483648 == 1)
                        _ = v20
                        _ = v21
                    }
                    _ = v19
                }
                v22 := GenS2{F0: 2147483647}
                v23 := GenS3{F0: 757}
                _ = v22
                _ = v23
            } else {
                v18 = (^(v18))
            }
            if (!false) {
                var v24 uint64 = uint64(v18)
                _ = v24
            } else {
                switch 32767 {
                case 918:
                    var v25 float32 = ((1e10 / 1.0) - 1e10)
                    _ = v25
                case 1:
                    v18 = 2
                default:
                    var v26 bool = true
                    _ = v26
                }
            }
            _ = v18
        } else {
            var v27 int64 = 1
            switch 32767 {
            case -1:
                var v28 uint32 = 2147483648
                _ = v28
            case 518:
                v27 -= v27
            default:
                v29 := GenS4{F0: 1}
                _ = v29
            }
            for fi30 := 0; fi30 < 8; fi30++ {
                v31 := []int{genFn0(-32768, 0), fi30}
                v32 := GenS1{F0: true, F1: -2147483648}
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = v27
        }
        if true {
            var v33 uint = 255
            _ = v33
        }
        var v34 int8 = -1
        _ = v34
    } else {
        var v35 uint8 = 237
        for fi36 := 0; fi36 < 6; fi36++ {
            v37 := GenS4{F0: 32767}
            v35 = (v35 - 2)
            for fi38 := 0; fi38 < 19; fi38++ {
                var v39 int32 = 2147483647
                if (1000 <= 1) {
                    var v40 int32 = (v39 | 211)
                    var v41 uint32 = 2
                    fmt.Println(fi36)
                    _ = v40
                    _ = v41
                }
                if true {
                    if false {
                        v42 := GenS4{F0: -920}
                        var v43 uint32 = 1
                        switch (fi36 % ((fi36) | 1)) {
                        case -2147483648:
                            v43 += v43
                        case 9223372036854775807:
                            var v44 uint64 = 1
                            _ = v44
                        case 1:
                            v45 := []int{-272}
                            _ = v45
                        default:
                            switch "foo bar" {
                            case "foo bar":
                                v39 = (-(v39))
                            case "hello":
                                fmt.Println(v39)
                            }
                        }
                        _ = v42
                        _ = v43
                    } else {
                        fmt.Println(fi38)
                        var v46 float32 = ((-2.25 / -2.25) / 1.0)
                        switch ("hello" + ("hello" + "")) {
                        case "hello":
                            v47 := []int{}
                            _ = v47
                        case "foo bar":
                            if false {
                                if false {
                                    var v48 uint32 = 645
                                    fmt.Println(v48, fi36, v39)
                                    _ = v48
                                }
                                v49 := GenS4{F0: -1}
                                v50 := []int{9223372036854775807}
                                _ = v49
                                _ = v50
                            } else {
                                v46 -= 1e10
                                v46 = v46
                            }
                        }
                        _ = v46
                    }
                    var v51 int16 = -1
                    _ = v51
                } else {
                    switch fi38 {
                    case -9223372036854775808:
                        var v52 int64 = -1
                        _ = v52
                    }
                    switch (fi38 - (fi38 >> ((65535) & 63))) {
                    case 4294967295:
                        v53 := []int{fi36, fi36, -1, genFn0(-32768, 0)}
                        if (!true) {
                            _ = len(v53)
                            var v54 uint32 = 65535
                            var v55 int32 = 0
                            _ = v54
                            _ = v55
                        }
                        _ = v53
                    default:
                        v56 := GenS5{F0: v35, F1: (v35 << ((42) & 63)), F2: 255}
                        _ = v56
                    }
                }
                v39 += -227
                switch (fi36 % (((fi36 % ((genFn0(1, 877)) | 1))) | 1)) {
                case 128:
                    fmt.Println(fi38, v35)
                case 0:
                    var v57 uint32 = 2
                    _ = v57
                case -2147483648:
                    switch ("a" + "") {
                    case "":
                        v39 -= v39
                    default:
                        var v58 float32 = 1.5
                        var v59 uint8 = 2
                        _ = v58
                        _ = v59
                    }
                    v60 := []int{(fi38 - fi36)}
                    v35 = 2
                    _ = v60
                }
                _ = fi38
                _ = v39
            }
            _ = fi36
            _ = v37
        }
        _ = v35
    }
    var v61 float32 = float32((float64(-2.25) / ((float64((0.0 + (-2.25 * float32((0.1 / 1.0))))) * -2.25) / (0.0 + ((((0.0 + (0.0 + 0.1)) + 1e10) - 3.14159) - -2.25)))))
    for fi62 := 0; fi62 < 5; fi62++ {
        if false {
            v63 := []int{fi62, 32767}
            for fi64 := 0; fi64 < 20; fi64++ {
                v61 += (v61 / -2.25)
                v61 -= v61
                v61 = ((((v61 - 3.14159) / 1.5) * v61) * v61)
                _ = fi64
            }
            _ = v63
        }
        _ = fi62
    }
    _ = v61
    return 81
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 42
    var v3 string = ""
    v4 := make(map[int]bool)
    v5 := []int{1, 922, 2147483647, -1}
    var v6 string = (v3 + "foo bar")
    _ = len(v5)
    var v65 float64 = 1e10
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v66 int = 2147483647
    v67 := GenS1{F0: true, F1: -2147483648}
    fmt.Println(v2, v3, v65)
    v2 -= -1
    var v68 int16 = 0
    var v69 string = "\n\t"
    for fi70 := 0; fi70 < 2; fi70++ {
        var v71 string = v3
        _ = fi70
        _ = v71
    }
    var v72 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v72
}
