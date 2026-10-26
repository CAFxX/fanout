package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 uint64; F1 uint32; F2 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 uint; F1 string; F2 int }
func genFn0() uint {
    if false {
        v7 := []int{-126, 1}
        _ = v7
    }
    var v8 uint = 110
    if true {
        if true {
            var v9 int32 = 222
            var v10 int64 = -1
            fmt.Println(v8, v8)
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: 4294967295}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 12; fi12++ {
            v8 += v8
            v8 = 2
            v8 = (^(v8))
            _ = fi12
        }
        switch -48 {
        case -1:
            v8 += (v8 - (v8 | 18446744073709551615))
        case 42:
            v8 -= (v8 % ((v8) | 1))
        case -581:
            v13 := GenS2{F0: 1}
            _ = v13
        }
        var v14 uint = (v8 >> (((v8 & (v8 * (v8 | 18446744073709551615)))) & 63))
        _ = v14
    }
    _ = v8
    return v8
}
func genFn1() int16 {
    v15 := GenS3{F0: 42, F1: -957}
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := []int{}
        var v18 float32 = 0.1
        var v19 int = -788
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 uint32 = 4294967295
    fmt.Println(v20)
    _ = v15
    _ = v20
    return -579
}

func main() {
    var v1 uint16 = 617
    var v2 uint8 = 1
    var v3 uint16 = (v1 - (v1 | (v1 * v1)))
    var v4 int64 = 42
    v5 := make(map[string]bool)
    var v6 float64 = 3.14159
    for fi21 := 0; fi21 < 7; fi21++ {
        v1 = 0
        _ = fi21
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        for fi23 := 0; fi23 < 18; fi23++ {
            fmt.Println(v4, v3, fi23)
            switch fi23 {
            case -1:
                if false {
                    fmt.Println(v2, v2, v6)
                    v6 += (((float64(1.5) * v6) - v6) + 1.5)
                }
            }
            _ = fi23
        }
        var v24 int16 = 1000
        for fi25 := 0; fi25 < 6; fi25++ {
            if false {
                v26 := GenS1{F0: 18446744073709551615}
                var v27 float64 = (1e10 + 3.14159)
                _ = v26
                _ = v27
            }
            v3 += 1
            _ = fi25
        }
        _ = fi22
        _ = v24
    }
    if false {
        fmt.Println(v3)
        for fi28 := 0; fi28 < 7; fi28++ {
            v29 := GenS4{F0: 519, F1: "\n\t"}
            switch fi28 {
            case 4294967295:
                if true {
                    v30 := []int{fi28}
                    v31 := v5["\n\t"]
                    _ = v30
                    _ = v31
                } else {
                    var v32 int32 = 0
                    var v33 int32 = (^(v32))
                    _ = v32
                    _ = v33
                }
            case -9223372036854775808:
                switch "" {
                case "foo bar":
                    if true {
                        switch (fi28 >> ((746) & 63)) {
                        case 554:
                            var v34 float32 = ((3.14159 * 1e10) * 1.5)
                            _ = v34
                        case 32767:
                            v4 = (v4 - v4)
                        case -9223372036854775808:
                            var v35 float32 = (-2.25 / 1.0)
                            _ = v35
                        }
                        var v36 int16 = 1000
                        v37 := []int{127, fi28, fi28}
                        _ = v36
                        _ = v37
                    } else {
                        v38 := []int{fi28, (fi28 - fi28)}
                        var v39 int = 4294967295
                        _ = v38
                        _ = v39
                    }
                case "":
                    var v40 float64 = float64(3.14159)
                    _ = v40
                default:
                    if (1.5 > 0.1) {
                        v2 = 127
                        v6 += v6
                    } else {
                        v41 := GenS2{F0: 32767}
                        v5[("a" + "")] = (false && true)
                        _ = v41
                    }
                }
            case 32767:
                var v42 float32 = 0.1
                _ = v42
            }
            if (!(23 > 1)) {
                for fi43 := 0; fi43 < 5; fi43++ {
                    var v44 int = fi28
                    _ = fi43
                    _ = v44
                }
                var v45 uint = 449
                v46 := GenS5{F0: 2, F1: 718, F2: fi28}
                _ = v45
                _ = v46
            }
            _ = fi28
            _ = v29
        }
        v47 := GenS2{F0: genFn1()}
        _ = v47
    } else {
        for fi48 := 0; fi48 < 7; fi48++ {
            switch "\n\t" {
            case "foo bar":
                for fi49 := 0; fi49 < 7; fi49++ {
                    fmt.Println(v3, v2)
                    switch ("\n\t" + "a") {
                    case "a":
                        v50 := GenS3{F0: uint8(v3), F1: (fi49 + 127)}
                        _ = v50
                    case "":
                        if false {
                            switch fi48 {
                            case 32768:
                                v3 -= (v3 * 255)
                            case -945:
                                v4 = (v4 | -9223372036854775808)
                            case 1:
                                var v51 int16 = -1
                                _ = v51
                            }
                            v52 := GenS6{F0: 4294967295}
                            v53 := GenS5{F0: 692, F1: 42, F2: fi48}
                            if true {
                                if ("a" > "foo bar") {
                                    var v54 int16 = int16(v2)
                                    v4 -= v4
                                    _ = v54
                                } else {
                                    var v55 uint = 18446744073709551615
                                    if false {
                                        v2 = (^(v2))
                                        switch "hello" {
                                        case "":
                                            fmt.Println(v2, v4)
                                            v3 -= 248
                                            if true {
                                                switch "hello" {
                                                case "hello":
                                                    v6 += v6
                                                default:
                                                    v55 += (v55 / (((v55 << ((65535) & 63))) | 1))
                                                }
                                                var v56 int16 = 721
                                                var v57 float64 = (((v6 / v6) - float64(3.14159)) - 3.14159)
                                                var v58 uint32 = 774
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            } else {
                                                v1 -= v1
                                                var v59 uint16 = v1
                                                _ = v59
                                            }
                                        default:
                                            v60 := GenS1{F0: 2}
                                            _ = v60
                                        }
                                        fmt.Println(fi48)
                                    }
                                    _ = v55
                                }
                                v61 := GenS3{F0: v2, F1: fi49}
                                _ = len(v5)
                                _ = v61
                            } else {
                                var v62 float64 = -2.25
                                _ = v62
                            }
                            v63 := v5["hello"]
                            var v64 bool = false
                            _ = v52
                            _ = v53
                            _ = v63
                            _ = v64
                        }
                    default:
                        v65 := GenS7{F0: uint32(v1), F1: 1}
                        _ = v65
                    }
                    switch ("foo bar" + "\n\t") {
                    case "foo bar":
                        var v66 int64 = (v4 - v4)
                        _ = v66
                    }
                    _ = fi49
                }
            }
            fmt.Println(v6, v3)
            var v67 int64 = v4
            _ = fi48
            _ = v67
        }
    }
    var v68 int32 = 2147483647
    v3 = v1
    v69 := GenS8{F0: 4294967295, F1: 1}
    var v70 int8 = 104
    var v71 int64 = 1
    switch ("" + ("" + ("hello" + "foo bar"))) {
    case "\n\t":
        v72 := GenS6{F0: 241}
        _ = v72
    case "":
        fmt.Println(v4)
    case "a":
        var v73 int = -174
        _ = v73
    }
    v74 := GenS1{F0: 68}
    v75 := GenS9{F0: 1, F1: (("foo bar" + "hello") + ("" + "foo bar")), F2: 32768}
    var v76 float64 = v6
    var v77 uint32 = 42
    fmt.Println(v76)
    v78 := GenS9{F0: 1, F1: "a", F2: -2147483648}
    for fi79 := 0; fi79 < 2; fi79++ {
        v4 = (v71 ^ -371)
        v2 = (^(v2))
        switch "foo bar" {
        case "hello":
            v70 -= v70
        case "a":
            var v80 float32 = 3.14159
            _ = v80
        default:
            var v81 uint8 = (v2 >> ((0) & 63))
            _ = v81
        }
        _ = fi79
    }
    if (1000 < 106) {
        fmt.Println(v70, v6, v70)
    } else {
        var v82 int32 = 2147483647
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v70
    _ = v71
    _ = v74
    _ = v75
    _ = v76
    _ = v77
    _ = v78
}
