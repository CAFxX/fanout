package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint64) int32 {
    for fi4 := 0; fi4 < 18; fi4++ {
        if (42 >= 0) {
            p0 += 60
            var v5 int = fi4
            v6 := GenS1{F0: 255, F1: 1}
            _ = v5
            _ = v6
        }
        p0 += p0
        _ = fi4
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 += p0
        v8 := GenS1{F0: 255, F1: 219}
        _ = fi7
        _ = v8
    }
    switch 127 {
    case 1:
        if true {
            fmt.Println(p0)
            if (float32(3.14159) >= -2.25) {
                if (42 < -114) {
                    p0 = 0
                    if (255 < 42) {
                        v9 := GenS1{F0: 2, F1: 128}
                        if (!(1 <= 0)) {
                            v10 := GenS1{F0: 1, F1: 255}
                            _ = v10
                        } else {
                            var v11 uint32 = 1
                            var v12 bool = false
                            _ = v11
                            _ = v12
                        }
                        _ = v9
                    }
                }
                var v13 bool = (-128 < 103)
                _ = v13
            } else {
                var v14 int32 = 42
                v15 := []int{}
                _ = v14
                _ = v15
            }
        }
    case 9223372036854775807:
        var v16 string = "hello"
        _ = v16
    default:
        switch 127 {
        case 0:
            var v17 int16 = 1
            _ = v17
        case -1:
            var v18 uint64 = p0
            _ = v18
        case 965:
            v19 := []int{42, -971, 128, -1}
            _ = v19
        }
    }
    p0 += (p0 >> ((255) & 63))
    var v20 string = "foo bar"
    _ = p0
    _ = v20
    return -364
}
func genFn1(p0 uint) uint {
    p0 = (p0 ^ p0)
    var v21 float32 = 3.14159
    switch (("foo bar" + "hello") + ("hello" + ("hello" + "\n\t"))) {
    case "hello":
        for fi22 := 0; fi22 < 13; fi22++ {
            fmt.Println(v21, fi22)
            var v23 uint8 = 42
            v24 := []int{(fi22 % (((fi22 >> ((p0) & 63))) | 1)), fi22}
            _ = fi22
            _ = v23
            _ = v24
        }
    default:
        switch ("hello" + "foo bar") {
        case "":
            v25 := GenS1{F0: 255, F1: 128}
            _ = v25
        case "foo bar":
            v26 := GenS2{F0: (uint64(p0) >= 115)}
            _ = v26
        case "\n\t":
            v27 := []int{-2147483648, 509, -1}
            _ = v27
        default:
            var v28 string = ("a" + ((("" + ("" + "foo bar")) + ("hello" + "")) + ("\n\t" + "\n\t")))
            _ = v28
        }
    }
    switch ((("foo bar" + "hello") + "") + "hello") {
    case "\n\t":
        for fi29 := 0; fi29 < 3; fi29++ {
            v21 -= v21
            v30 := GenS1{F0: p0, F1: 42}
            _ = fi29
            _ = v30
        }
    case "":
        p0 -= 1
    }
    p0 = p0
    _ = p0
    _ = v21
    return 895
}

func main() {
    var v1 uint = 532
    var v2 float64 = 1.5
    v3 := make(map[string]bool)
    if (!((!(!true)) || true)) {
        fmt.Println(v2)
    } else {
        v2 += (1.5 * v2)
    }
    v1 += v1
    fmt.Println(v1, v1)
    var v31 int32 = -1
    _ = len(v3)
    v31 += 0
    v32 := GenS1{F0: (v1 << ((v1) & 63)), F1: 160}
    for fi33 := 0; fi33 < 18; fi33++ {
        fmt.Println(v31)
        for fi34 := 0; fi34 < 7; fi34++ {
            switch ("hello" + "a") {
            case "\n\t":
                v1 = 65535
            }
            _ = fi34
        }
        _ = fi33
    }
    if (false || (42 > 0)) {
        for fi35 := 0; fi35 < 12; fi35++ {
            switch (fi35 & fi35) {
            case 0:
                v31 -= (v31 << ((v1) & 63))
            case 1:
                v36 := GenS3{F0: -908, F1: (v1 * (v1 >> ((4294967295) & 63)))}
                _ = v36
            default:
                var v37 uint16 = 42
                _ = v37
            }
            v2 -= v2
            v38 := v3["\n\t"]
            _ = fi35
            _ = v38
        }
        var v39 uint = 65535
        _ = v39
    }
    v1 = 65535
    if (254 > 18446744073709551615) {
        for fi40 := 0; fi40 < 15; fi40++ {
            for fi41 := 0; fi41 < 19; fi41++ {
                v1 = genFn1((v1 << (((v1 << ((v1) & 63))) & 63)))
                fmt.Println(v1)
                switch fi41 {
                case 127:
                    v2 = 3.14159
                }
                switch "foo bar" {
                case "hello":
                    if false {
                        var v42 uint16 = 32767
                        switch "hello" {
                        case "a":
                            if false {
                                var v43 uint32 = 1
                                _ = v43
                            } else {
                                var v44 uint64 = 18446744073709551615
                                _ = v44
                            }
                        case "":
                            v3[(("a" + "hello") + "foo bar")] = false
                            v45 := GenS2{F0: (59 >= 42)}
                            v46 := GenS4{F0: v42}
                            _ = v45
                            _ = v46
                        }
                        var v47 int8 = 0
                        switch "\n\t" {
                        case "foo bar":
                            delete(v3, "foo bar")
                        default:
                            fmt.Println(v42, v1, v1)
                        }
                        _ = v42
                        _ = v47
                    } else {
                        if true {
                            if false {
                                var v48 int8 = 100
                                _ = v48
                            }
                            if true {
                                var v49 uint16 = 979
                                v31 -= 0
                                v1 = genFn1((v1 | (^(v1))))
                                _ = v49
                            }
                            v3["hello"] = false
                        }
                        var v50 float32 = -2.25
                        v2 -= (float64(-2.25) + 3.14159)
                        fmt.Println(v50, v50)
                        _ = v50
                    }
                case "\n\t":
                    v51 := []int{}
                    _ = v51
                default:
                    if false {
                        v31 = v31
                        v2 += v2
                        fmt.Println(v2, v31, v1)
                    }
                }
                _ = fi41
            }
            for fi52 := 0; fi52 < 11; fi52++ {
                if (-750 < 42) {
                    v53 := GenS4{F0: 1}
                    fmt.Println(v2, v1)
                    _ = v53
                } else {
                    v54 := GenS4{F0: 2}
                    fmt.Println(v2, v31)
                    _ = v54
                }
                var v55 string = ("a" + "a")
                v3[v55] = (510 == 655)
                _ = fi52
                _ = v55
            }
            _ = fi40
        }
    } else {
        var v56 int16 = -1
        v57 := GenS1{F0: 0, F1: 1}
        _ = v56
        _ = v57
    }
    v2 += 0.1
    var v58 uint = (v1 * (^(v1)))
    for fi59 := 0; fi59 < 6; fi59++ {
        if false {
            v60 := v3["hello"]
            _ = len(v3)
            _ = v60
        } else {
            var v61 uint16 = 888
            for fi62 := 0; fi62 < 3; fi62++ {
                var v63 float32 = 1.5
                _ = fi62
                _ = v63
            }
            var v64 uint64 = 2
            _ = v61
            _ = v64
        }
        v1 -= (v58 >> ((v58) & 63))
        if true {
            for fi65 := 0; fi65 < 8; fi65++ {
                v3["\n\t"] = false
                var v66 int16 = -1
                _ = fi65
                _ = v66
            }
            v31 += (^(v31))
            delete(v3, ("hello" + "foo bar"))
        } else {
            var v67 uint8 = 255
            var v68 uint64 = 358
            v31 = (v31 * (^(v31)))
            _ = v67
            _ = v68
        }
        _ = fi59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v58
}
