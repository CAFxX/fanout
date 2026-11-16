package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint8; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint64; F2 int8 }
func genFn0() int {
    v5 := []int{468, 32768, 32767}
    var v6 float32 = 1.5
    _ = v5
    _ = v6
    return -288
}
func genFn1(p0 uint16, p1 uint8) int32 {
    var v7 uint = 450
    v8 := GenS1{F0: 233}
    v9 := []int{}
    p1 = p1
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            v7 = (v7 << ((uint(p0)) & 63))
            _ = fi11
        }
        if true {
            switch 993 {
            case 42:
                v12 := GenS1{F0: 4294967295}
                _ = v12
            case 128:
                fmt.Println(p1, p0, p1)
            case -145:
                v13 := GenS2{F0: 42, F1: (p1 & p1), F2: 9223372036854775807}
                _ = v13
            default:
                v9 = append(v9, -290)
            }
            switch ("\n\t" + "foo bar") {
            case "hello":
                if len(v9) > 0 {
                    v9[0] = 1
                }
            }
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    var v3 uint = 105
    var v4 int = 1
    for fi14 := 0; fi14 < 15; fi14++ {
        if ((3.14159 / 1.0) == 0.1) {
            v15 := v2["\n\t"]
            var v16 float64 = 3.14159
            _ = v15
            _ = v16
        } else {
            var v17 string = "\n\t"
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS1{F0: 2147483647}
                var v20 int64 = 42
                v20 = v20
                var v21 uint8 = 127
                if true {
                    if true {
                        var v22 int16 = 1000
                        var v23 int8 = 0
                        _ = v22
                        _ = v23
                    } else {
                        var v24 int64 = v20
                        v25 := GenS1{F0: 4294967295}
                        var v26 float32 = 1e10
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                }
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            _ = v17
        }
        if true {
            fmt.Println(v4, fi14, v3)
            switch "a" {
            case "":
                for fi27 := 0; fi27 < 16; fi27++ {
                    v4 = (fi27 ^ (^(fi14)))
                    fmt.Println(fi27)
                    v4 -= 4294967295
                    fmt.Println(v3, fi27, fi27)
                    if true {
                        var v28 uint8 = 127
                        _ = v28
                    } else {
                        v3 = v3
                        if (1 == 2147483647) {
                            fmt.Println(fi14, fi14, v4)
                        } else {
                            var v29 float64 = (1e10 - 0.0)
                            if ("\n\t" < "\n\t") {
                                switch "a" {
                                case "":
                                    var v30 int64 = -9223372036854775808
                                    _ = v30
                                }
                                var v31 string = (("a" + ("" + ("hello" + "hello"))) + ("" + "a"))
                                _ = v31
                            } else {
                                if false {
                                    fmt.Println(fi27)
                                    var v32 uint16 = 0
                                    _ = v32
                                }
                                if false {
                                    v33 := GenS3{F0: 0}
                                    _ = v33
                                } else {
                                    switch ("\n\t" + (("" + "foo bar") + ("foo bar" + "hello"))) {
                                    case "":
                                        v34 := GenS3{F0: 1000}
                                        _ = v34
                                    case "foo bar":
                                        v35 := GenS2{F0: 0, F1: 42, F2: 0}
                                        _ = v35
                                    case "a":
                                        var v36 uint64 = 1
                                        _ = v36
                                    }
                                }
                            }
                            var v37 uint = (^(v3))
                            _ = v29
                            _ = v37
                        }
                    }
                    _ = fi27
                }
            case "\n\t":
                var v38 int32 = 42
                _ = v38
            }
            var v39 bool = false
            _ = v39
        } else {
            switch -342 {
            case -2147483648:
                switch (v4 << (((v3 | v3)) & 63)) {
                case 224:
                    if (!true) {
                        var v40 uint = (v3 + (v3 % (((v3 + v3)) | 1)))
                        v4 -= fi14
                        _ = v40
                    } else {
                        fmt.Println(v4, v4, v4)
                        v4 += -9223372036854775808
                        var v41 int16 = -1
                        _ = v41
                    }
                case 42:
                    switch (fi14 / (((v4 % ((genFn0()) | 1))) | 1)) {
                    case 0:
                        switch (v4 * (v4 % ((fi14) | 1))) {
                        case -245:
                            v4 -= (v4 % (((fi14 >> (((v3 + v3)) & 63))) | 1))
                        default:
                            switch "foo bar" {
                            case "foo bar":
                                switch genFn0() {
                                case -681:
                                    v42 := GenS2{F0: 0, F1: 128, F2: 42}
                                    _ = v42
                                case 32768:
                                    if false {
                                        v4 = (fi14 ^ 32767)
                                        v43 := []int{}
                                        _ = v43
                                    }
                                default:
                                    fmt.Println(v4, v3)
                                }
                            case "\n\t":
                                v44 := []int{(v4 << ((588) & 63)), -9223372036854775808, (v4 - (fi14 >> (((^(v3))) & 63))), genFn0()}
                                _ = v44
                            case "":
                                if false {
                                    switch (v4 / ((174) | 1)) {
                                    case -443:
                                        var v45 uint16 = 255
                                        _ = v45
                                    case 1:
                                        var v46 int32 = 2147483647
                                        _ = v46
                                    case 127:
                                        fmt.Println(v3, v4)
                                    default:
                                        for fi47 := 0; fi47 < 17; fi47++ {
                                            var v48 uint = v3
                                            var v49 int = (fi14 | (fi47 ^ (fi47 - genFn0())))
                                            _ = fi47
                                            _ = v48
                                            _ = v49
                                        }
                                    }
                                    v50 := []int{(v4 << ((545) & 63)), (fi14 + v4), (fi14 - v4), v4}
                                    for fi51 := 0; fi51 < 14; fi51++ {
                                        v3 -= (v3 >> (((v3 * 42)) & 63))
                                        _ = fi51
                                    }
                                    _ = v50
                                } else {
                                    v4 = genFn0()
                                    for fi52 := 0; fi52 < 14; fi52++ {
                                        var v53 uint8 = 255
                                        _ = fi52
                                        _ = v53
                                    }
                                    v4 += (fi14 ^ fi14)
                                }
                            default:
                                v54 := GenS3{F0: 42}
                                _ = v54
                            }
                        }
                    }
                }
            }
            var v55 uint8 = 0
            if (-2.25 != ((0.1 * 1e10) * ((((float32((3.14159 / 1.0)) * 0.1) / 0.1) / 1.5) / 0.1))) {
                v55 += (v55 | v55)
            } else {
                v56 := []int{(-(v4)), v4, (-(fi14)), (^(fi14))}
                fmt.Println(v55)
                _ = v56
            }
            _ = v55
        }
        _ = fi14
    }
    v57 := v1[(v4 - 127)]
    v58 := GenS2{F0: 751, F1: 255, F2: -9223372036854775808}
    fmt.Println(v4)
    v4 = (v4 >> ((v3) & 63))
    v59 := GenS4{F0: "\n\t"}
    v60 := GenS1{F0: 2}
    v57 = v57
    if (1 >= 1000) {
        v3 = v3
        if false {
            fmt.Println(v3, v4, v57)
            v61 := GenS5{F0: 42, F1: 18446744073709551615, F2: 100}
            for fi62 := 0; fi62 < 19; fi62++ {
                var v63 uint16 = 118
                v64 := GenS1{F0: 408}
                _ = fi62
                _ = v63
                _ = v64
            }
            _ = v61
        }
        fmt.Println(v3, v57, v3)
    } else {
        v1[v4] = genFn0()
        v4 += (v57 * 486)
    }
    var v65 int16 = 32767
    var v66 int = (v57 + v4)
    fmt.Println(v66, v3, v57)
    switch ((((((((("a" + "a") + "\n\t") + "foo bar") + "\n\t") + ("\n\t" + "\n\t")) + ((("foo bar" + ("foo bar" + "")) + (("hello" + "") + "")) + "\n\t")) + (("\n\t" + "") + "a")) + "foo bar") + "") {
    case "":
        for fi67 := 0; fi67 < 14; fi67++ {
            v66 += v57
            v4 -= genFn0()
            _ = fi67
        }
    default:
        if (v65 == (^(v65))) {
            switch (-(v4)) {
            case -260:
                var v68 bool = ((2147483648 == 2) && (0 != 4294967295))
                _ = v68
            case -9223372036854775808:
                fmt.Println(v57)
            case 32767:
                delete(v1, (-(v66)))
            }
        } else {
            if false {
                for fi69 := 0; fi69 < 2; fi69++ {
                    v65 -= (v65 / (((^(v65))) | 1))
                    v4 = v4
                    _ = fi69
                }
                var v70 int16 = (v65 % ((-1) | 1))
                _ = v70
            } else {
                var v71 uint32 = 454
                _ = v71
            }
            var v72 bool = false
            _ = v72
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v65
    _ = v66
}
