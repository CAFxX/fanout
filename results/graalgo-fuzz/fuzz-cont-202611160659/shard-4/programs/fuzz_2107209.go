package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint8; F1 uint; F2 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int; F1 uint; F2 uint }
type GenS8 struct { F0 uint16; F1 int }
type GenS9 struct { F0 string; F1 bool }
type GenS10 struct { F0 uint32; F1 int32; F2 uint16 }
func genFn0(p0 int, p1 bool) string {
    var v4 float32 = 0.0
    if (!((float64(0.1) != 1.5) && true)) {
        fmt.Println(v4, p1, v4)
        var v5 int = 128
        v6 := []int{p0, (-(v5)), v5}
        _ = v5
        _ = v6
    } else {
        var v7 uint16 = 0
        var v8 int32 = 193
        p0 = (p0 % (((p0 + 2147483647)) | 1))
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: "a"}
    var v10 uint8 = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v4
    _ = v9
    return ""
}
func genFn1(p0 int32, p1 bool, p2 int8) uint {
    v11 := GenS1{F0: (("" + "\n\t") + "")}
    if p1 {
        fmt.Println(p2)
    }
    var v12 int = 1
    p0 = p0
    var v13 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 int = -9223372036854775808
    var v3 int32 = 0
    v2 += v2
    v14 := GenS1{F0: "\n\t"}
    switch ("" + "hello") {
    case "foo bar":
        var v15 uint64 = 366
        _ = v15
    case "":
        v16 := []int{}
        _ = v16
    }
    if false {
        if true {
            var v17 uint32 = 42
            switch "foo bar" {
            case "a":
                v1[""] = v2
            case "foo bar":
                for fi18 := 0; fi18 < 7; fi18++ {
                    if (!false) {
                        if true {
                            var v19 int = fi18
                            var v20 uint64 = 416
                            _ = v19
                            _ = v20
                        }
                        switch "hello" {
                        case "hello":
                            for fi21 := 0; fi21 < 5; fi21++ {
                                if ((false && (1.5 <= (0.1 * float32((3.14159 / float64(0.1)))))) || true) {
                                    v22 := GenS1{F0: "foo bar"}
                                    v23 := []int{}
                                    var v24 int8 = 100
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                } else {
                                    v25 := GenS1{F0: "hello"}
                                    _ = v25
                                }
                                fmt.Println(v2)
                                v3 -= v3
                                _ = fi21
                            }
                        case "a":
                            switch "foo bar" {
                            case "a":
                                var v26 string = "a"
                                _ = v26
                            default:
                                var v27 bool = false
                                _ = v27
                            }
                        }
                        fmt.Println(fi18, v17, v2)
                    } else {
                        var v28 uint = 1
                        if (!(42 != 954)) {
                            if (!(!true)) {
                                v29 := GenS2{F0: (v17 % ((2147483647) | 1)), F1: 32767}
                                v2 = fi18
                                _ = v29
                            }
                        } else {
                            if false {
                                v30 := GenS3{F0: (fi18 | 0), F1: v17, F2: 0}
                                v31 := GenS4{F0: 128, F1: (v28 << ((v28) & 63)), F2: 32767}
                                v17 += (^(v17))
                                _ = v30
                                _ = v31
                            }
                            var v32 int32 = (v3 % (((v3 % (((v3 + 42)) | 1))) | 1))
                            var v33 uint32 = (^(v17))
                            _ = v32
                            _ = v33
                        }
                        _ = v28
                    }
                    fmt.Println(v17)
                    v3 = v3
                    _ = fi18
                }
            default:
                if false {
                    if ("" <= "") {
                        var v34 float32 = ((0.1 + 0.1) - (1e10 - 1.5))
                        if true {
                            if ((!("a" <= "hello")) || false) {
                                v17 = (v17 << ((4294967295) & 63))
                                fmt.Println(v2)
                                v35 := GenS5{F0: "a"}
                                _ = v35
                            }
                        }
                        v3 = -213
                        _ = v34
                    } else {
                        var v36 bool = true
                        if (1 != int64(v3)) {
                            var v37 bool = true
                            if v36 {
                                if (v37 || (((v3 > (v3 | (v3 & (v3 % (((v3 / ((v3) | 1))) | 1))))) && false) && ("hello" > (("" + (("\n\t" + "hello") + "a")) + "")))) {
                                    v36 = false
                                    v38 := []int{v2, v2, (v2 + (v2 & (v2 + (v2 + (v2 + (^(v2)))))))}
                                    v36 = v37
                                    _ = v38
                                }
                                fmt.Println(v17, v17)
                                v36 = (42 >= 734)
                            } else {
                                for fi39 := 0; fi39 < 14; fi39++ {
                                    for fi40 := 0; fi40 < 17; fi40++ {
                                        var v41 uint = uint(v17)
                                        _ = fi40
                                        _ = v41
                                    }
                                    for fi42 := 0; fi42 < 1; fi42++ {
                                        if (v37 && v36) {
                                            v43 := []int{fi42}
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v44 uint8 = 2
                                                switch int(v44) {
                                                case -38:
                                                    _ = len(v43)
                                                case 42:
                                                    if (65535 != 56) {
                                                        v37 = ((v2 >= 32767) && v37)
                                                    } else {
                                                        var v45 int8 = 20
                                                        _ = v45
                                                    }
                                                }
                                                v44 += v44
                                                _ = v44
                                            case "\n\t":
                                                switch (fi39 / ((v2) | 1)) {
                                                case 128:
                                                    v43 = append(v43, (v2 >> ((0) & 63)))
                                                case -2147483648:
                                                    v17 = v17
                                                case 42:
                                                    fmt.Println(v2)
                                                }
                                            case "hello":
                                                v3 = v3
                                            default:
                                                _ = len(v1)
                                            }
                                            _ = v43
                                        } else {
                                            var v46 int32 = (v3 << ((42) & 63))
                                            _ = v46
                                        }
                                        if (1.5 == 3.14159) {
                                            if (-1 == -72) {
                                                if true {
                                                    var v47 int32 = 1
                                                    v48 := GenS6{F0: 574}
                                                    _ = v47
                                                    _ = v48
                                                }
                                                v49 := GenS7{F0: 1, F1: 255, F2: 27}
                                                fmt.Println(v3, fi39)
                                                var v50 int64 = 1
                                                v51 := GenS8{F0: 0, F1: (-(fi42))}
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                            }
                                            v52 := []int{fi42, (v2 ^ v2)}
                                            _ = v52
                                        } else {
                                            v53 := GenS9{F0: "", F1: (!((("hello" + "") + "foo bar") > "foo bar"))}
                                            var v54 uint64 = 42
                                            _ = v53
                                            _ = v54
                                        }
                                        fmt.Println(v2, fi39)
                                        _ = fi42
                                    }
                                    _ = fi39
                                }
                                fmt.Println(v2, v36, v2)
                            }
                            if (false || (v17 >= (v17 % (((^(v17))) | 1)))) {
                                var v55 string = "a"
                                _ = v55
                            } else {
                                for fi56 := 0; fi56 < 3; fi56++ {
                                    delete(v1, "foo bar")
                                    if v37 {
                                        var v57 uint16 = 280
                                        switch (v2 + v2) {
                                        case 127:
                                            switch (v2 + 2147483647) {
                                            case 128:
                                                fmt.Println(v57, v57)
                                            case -1:
                                                if true {
                                                    if v36 {
                                                        v1["foo bar"] = (fi56 << ((18446744073709551615) & 63))
                                                        fmt.Println(fi56, v57, v37)
                                                    } else {
                                                        v2 = v2
                                                        if (94 >= 1) {
                                                            for fi58 := 0; fi58 < 19; fi58++ {
                                                                fmt.Println(v17)
                                                                v37 = v37
                                                                _ = fi58
                                                            }
                                                            fmt.Println(v3, v36)
                                                            var v59 uint32 = 2147483648
                                                            _ = v59
                                                        }
                                                        v60 := []int{}
                                                        _ = v60
                                                    }
                                                }
                                            default:
                                                v36 = (160 > 2)
                                            }
                                        case 9223372036854775807:
                                            fmt.Println(v17, fi56)
                                        }
                                        _ = v57
                                    }
                                    v61 := GenS3{F0: (fi56 & v2), F1: v17, F2: (-(v3))}
                                    _ = fi56
                                    _ = v61
                                }
                                var v62 int8 = -110
                                _ = v62
                            }
                            _ = v37
                        } else {
                            for fi63 := 0; fi63 < 7; fi63++ {
                                if (1 == 2) {
                                    v17 = (v17 ^ (v17 / (((v17 % ((0) | 1))) | 1)))
                                }
                                _ = fi63
                            }
                        }
                        if (v3 > v3) {
                            for fi64 := 0; fi64 < 16; fi64++ {
                                for fi65 := 0; fi65 < 8; fi65++ {
                                    v66 := []int{fi64}
                                    _ = fi65
                                    _ = v66
                                }
                                var v67 int16 = 1000
                                var v68 int16 = (v67 + 0)
                                _ = fi64
                                _ = v67
                                _ = v68
                            }
                            switch (v2 >> ((65535) & 63)) {
                            case 424:
                                v69 := v1["\n\t"]
                                _ = v69
                            case -1:
                                switch (v2 % ((v2) | 1)) {
                                case 4294967295:
                                    if (225 <= -189) {
                                        v17 -= 65535
                                        var v70 uint = 2
                                        _ = v70
                                    } else {
                                        v3 = (v3 * v3)
                                        v36 = v36
                                    }
                                default:
                                    if v36 {
                                        fmt.Println(v17, v36, v3)
                                        v17 += v17
                                    } else {
                                        fmt.Println(v3, v17)
                                        var v71 int16 = 32767
                                        v72 := []int{int(v17), (v2 ^ v2)}
                                        _ = v71
                                        _ = v72
                                    }
                                }
                            case 2147483647:
                                var v73 int = (^(v2))
                                _ = v73
                            default:
                                var v74 float32 = -2.25
                                _ = v74
                            }
                            if false {
                                var v75 uint16 = 75
                                _ = v75
                            } else {
                                fmt.Println(v2)
                                fmt.Println(v3, v36, v17)
                                var v76 uint8 = 42
                                _ = v76
                            }
                        }
                        _ = v36
                    }
                    var v77 int16 = 32767
                    if (true || (true && true)) {
                        for fi78 := 0; fi78 < 12; fi78++ {
                            var v79 string = "foo bar"
                            _ = fi78
                            _ = v79
                        }
                        var v80 float64 = 0.1
                        _ = v80
                    } else {
                        switch v2 {
                        case 9223372036854775807:
                            v81 := GenS2{F0: 842, F1: 65535}
                            _ = v81
                        case 2147483647:
                            switch v2 {
                            case 128:
                                v82 := []int{(v2 & (^(v2))), v2, v2, 0}
                                _ = v82
                            case 770:
                                v1["a"] = (v2 * -2147483648)
                            default:
                                var v83 string = ((("\n\t" + "\n\t") + "\n\t") + "foo bar")
                                _ = v83
                            }
                        }
                        for fi84 := 0; fi84 < 15; fi84++ {
                            var v85 uint = 255
                            var v86 uint32 = v17
                            _ = fi84
                            _ = v85
                            _ = v86
                        }
                        v17 += v17
                    }
                    _ = v77
                }
            }
            v17 += (v17 - (v17 - (v17 ^ 320)))
            _ = v17
        }
        if (0 == 0) {
            switch ((("" + ((("a" + ("" + "hello")) + "a") + "a")) + "hello") + "a") {
            case "":
                fmt.Println(v2, v3)
            case "foo bar":
                fmt.Println(v3, v3)
            }
            var v87 string = (("a" + ("" + ("hello" + (("" + (("foo bar" + "foo bar") + ("foo bar" + ("\n\t" + "a")))) + "a")))) + "\n\t")
            v88 := GenS2{F0: 678, F1: 255}
            _ = v87
            _ = v88
        } else {
            fmt.Println(v2)
            var v89 float64 = (1.5 + 0.1)
            v90 := []int{}
            _ = v89
            _ = v90
        }
        fmt.Println(v2, v2)
    }
    v3 = v3
    var v91 uint = genFn1((v3 * 0), true, 127)
    v92 := []int{(v2 * (v2 ^ (v2 - 9223372036854775807)))}
    v93 := GenS10{F0: 347, F1: (v3 - 29), F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v91:", v91)
    fmt.Println("slicelen v92:", len(v92), v92)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v91
    _ = v92
    _ = v93
}
