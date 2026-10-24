package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int8; F2 int }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int32; F1 int8; F2 uint8 }
type GenS8 struct { F0 int8; F1 uint16 }
type GenS9 struct { F0 uint16; F1 uint32 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint32; F1 int16; F2 int64 }
func genFn0(p0 int, p1 string) int8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 string = ((p1 + "") + "")
        switch p1 {
        case "":
            var v8 bool = true
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    v9 := GenS1{F0: 4294967295}
    switch (p0 / ((p0) | 1)) {
    case -344:
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint32 = 2147483648
            if (18446744073709551615 >= 718) {
                for fi12 := 0; fi12 < 1; fi12++ {
                    v13 := GenS2{F0: 156, F1: -1, F2: 32767}
                    switch "foo bar" {
                    case "\n\t":
                        if true {
                            var v14 uint32 = v11
                            p0 -= (p0 - (fi10 + fi12))
                            if ((v11 & 65535) <= (v14 >> ((52) & 63))) {
                                v15 := GenS1{F0: 42}
                                _ = v15
                            } else {
                                switch fi10 {
                                case 1:
                                    v16 := []int{}
                                    _ = v16
                                default:
                                    var v17 int = 0
                                    _ = v17
                                }
                            }
                            _ = v14
                        }
                    case "hello":
                        v18 := GenS1{F0: 553}
                        _ = v18
                    case "foo bar":
                        v19 := GenS3{F0: 1, F1: 18446744073709551615}
                        var v20 uint32 = 0
                        var v21 string = p1
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    v22 := GenS1{F0: 42}
                    _ = fi12
                    _ = v13
                    _ = v22
                }
                v23 := []int{32768, (fi10 + int(v11)), 32768}
                if (42 != -388) {
                    v24 := GenS4{F0: 912}
                    fmt.Println(p1, fi10)
                    for fi25 := 0; fi25 < 15; fi25++ {
                        var v26 bool = (128 == 0)
                        if v26 {
                            if (3.14159 != ((float64(1e10) + 1.5) + 0.0)) {
                                var v27 float32 = (3.14159 - -2.25)
                                _ = v27
                            } else {
                                var v28 float32 = 0.1
                                p0 = (p0 ^ (-(fi25)))
                                var v29 uint64 = 177
                                _ = v28
                                _ = v29
                            }
                            fmt.Println(p1, p0)
                            fmt.Println(fi10)
                        } else {
                            if (1 <= 42) {
                                if (-56 > -97) {
                                    v23 = append(v23, (p0 | (fi25 + -1)))
                                    var v30 float64 = (0.1 / (-2.25 / float64(float32(1.5))))
                                    var v31 uint32 = 2
                                    _ = v30
                                    _ = v31
                                } else {
                                    var v32 int16 = -32768
                                    p0 = 42
                                    _ = v32
                                }
                                p0 += 0
                                fmt.Println(v11, v26, p1)
                            }
                            v11 += (v11 << ((2) & 63))
                        }
                        if v26 {
                            var v33 int = fi10
                            v34 := GenS5{F0: 72, F1: 42, F2: -128}
                            v35 := GenS1{F0: 0}
                            _ = v33
                            _ = v34
                            _ = v35
                        } else {
                            p0 += -19
                            if v26 {
                                v26 = (1 != -1)
                                if (!(18446744073709551615 != 255)) {
                                    _ = len(v23)
                                    fmt.Println(fi10, fi10, p1)
                                    v26 = false
                                }
                            } else {
                                switch ((p1 + "a") + p1) {
                                case "a":
                                    var v36 uint8 = 94
                                    _ = v36
                                case "foo bar":
                                    if (p1 >= ("" + "foo bar")) {
                                        v26 = true
                                        v37 := GenS6{F0: 128, F1: 1}
                                        var v38 uint32 = 597
                                        _ = v37
                                        _ = v38
                                    } else {
                                        var v39 uint32 = (v11 + 458)
                                        _ = v39
                                    }
                                }
                                v11 -= v11
                                var v40 int = -2147483648
                                _ = v40
                            }
                            switch (p1 + p1) {
                            case "a":
                                var v41 uint = 65535
                                var v42 float32 = 1e10
                                _ = v41
                                _ = v42
                            default:
                                switch p1 {
                                case "hello":
                                    var v43 bool = (100 > 42)
                                    _ = v43
                                case "a":
                                    switch (("hello" + (("a" + "") + ("a" + "\n\t"))) + "\n\t") {
                                    case "":
                                        switch (p1 + p1) {
                                        case "":
                                            fmt.Println(v11)
                                        default:
                                            switch -9223372036854775808 {
                                            case 32767:
                                                v23 = append(v23, 42)
                                            }
                                        }
                                    }
                                }
                            }
                        }
                        _ = fi25
                        _ = v26
                    }
                    _ = v24
                } else {
                    v44 := GenS7{F0: 42, F1: 53, F2: 1}
                    v45 := GenS6{F0: 115, F1: 0}
                    p0 -= (p0 % ((p0) | 1))
                    _ = v44
                    _ = v45
                }
                _ = v23
            } else {
                v11 = (v11 - v11)
                v46 := GenS2{F0: 0, F1: -1, F2: (fi10 >> ((255) & 63))}
                var v47 string = p1
                _ = v46
                _ = v47
            }
            for fi48 := 0; fi48 < 13; fi48++ {
                v11 -= 887
                fmt.Println(fi10, fi48, p1)
                p0 = 9223372036854775807
                _ = fi48
            }
            _ = fi10
            _ = v11
        }
    case 42:
        for fi49 := 0; fi49 < 15; fi49++ {
            v50 := []int{fi49, 42}
            v51 := GenS8{F0: -128, F1: 2}
            _ = fi49
            _ = v50
            _ = v51
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    var v1 bool = true
    v2 := make(map[string]int)
    var v3 float64 = (1e10 * float64(0.1))
    v4 := make(map[string]int)
    var v5 uint = 4294967295
    fmt.Println(v3, v5, v1)
    var v52 uint = (v5 << ((v5) & 63))
    var v53 uint = 18446744073709551615
    v4[("a" + (("hello" + ("a" + "")) + "\n\t"))] = 9223372036854775807
    if true {
        var v54 uint = (^(v52))
        var v55 string = ("" + (("foo bar" + "hello") + "\n\t"))
        if (float32(v3) <= 1.5) {
            var v56 float64 = -2.25
            fmt.Println(v56, v54)
            v52 += v54
            _ = v56
        }
        _ = v54
        _ = v55
    }
    var v57 uint64 = 18446744073709551615
    for fi58 := 0; fi58 < 16; fi58++ {
        v1 = (("foo bar" + ((("foo bar" + "") + "foo bar") + "foo bar")) != "\n\t")
        if v1 {
            var v59 int32 = -2147483648
            v60 := GenS9{F0: 32767, F1: 4294967295}
            _ = v59
            _ = v60
        }
        if false {
            if (-9223372036854775808 >= fi58) {
                var v61 uint32 = 4294967295
                for fi62 := 0; fi62 < 10; fi62++ {
                    v1 = (true || true)
                    _ = fi62
                }
                _ = v61
            }
            v3 += v3
        } else {
            v5 = (v53 * v52)
            if (!((v1 || (v1 && v1)) || (("\n\t" + "a") <= ("" + "hello")))) {
                v5 = (v5 % (((v53 - (^(v53)))) | 1))
            } else {
                v53 = 65535
                switch ("\n\t" + "a") {
                case "":
                    var v63 int8 = 54
                    _ = v63
                case "foo bar":
                    if ((fi58 - fi58) <= 752) {
                        v57 -= v57
                        switch (fi58 + (fi58 + (fi58 ^ fi58))) {
                        case 2147483647:
                            if (127 == 1) {
                                if (!v1) {
                                    switch "hello" {
                                    case "\n\t":
                                        if true {
                                            for fi64 := 0; fi64 < 14; fi64++ {
                                                v65 := v2["hello"]
                                                v66 := []int{480, 4294967295}
                                                var v67 int16 = -622
                                                _ = fi64
                                                _ = v65
                                                _ = v66
                                                _ = v67
                                            }
                                            var v68 int16 = -488
                                            var v69 uint64 = v57
                                            _ = v68
                                            _ = v69
                                        }
                                    case "":
                                        v5 = v5
                                    case "a":
                                        v70 := GenS1{F0: 4294967295}
                                        _ = v70
                                    default:
                                        fmt.Println(v53)
                                    }
                                }
                                if v1 {
                                    var v71 uint32 = 1
                                    _ = v71
                                } else {
                                    if false {
                                        var v72 uint16 = 534
                                        _ = v72
                                    }
                                }
                                fmt.Println(v53)
                            } else {
                                v1 = (!v1)
                                switch ("" + ("hello" + "foo bar")) {
                                case "a":
                                    v4[(("\n\t" + "") + "hello")] = -36
                                case "hello":
                                    v73 := []int{(fi58 | 9223372036854775807), 32768, (fi58 * (fi58 ^ fi58))}
                                    _ = v73
                                }
                                v5 -= v52
                            }
                        case -423:
                            var v74 float64 = float64((float32(-2.25) / 1.0))
                            _ = v74
                        case 0:
                            v57 = (v57 | (v57 & (v57 & v57)))
                        default:
                            var v75 uint32 = 71
                            _ = v75
                        }
                    }
                case "a":
                    var v76 uint64 = v57
                    _ = v76
                default:
                    for fi77 := 0; fi77 < 11; fi77++ {
                        v5 = (v5 + (v5 | (v5 & v5)))
                        fmt.Println(fi77, v53)
                        _ = fi77
                    }
                }
                fmt.Println(v1, fi58, v1)
            }
            v4[("" + (("hello" + "foo bar") + "foo bar"))] = (fi58 >> (((^(v52))) & 63))
        }
        _ = fi58
    }
    if v1 {
        var v78 int32 = 0
        _ = v78
    }
    v79 := GenS10{F0: 0}
    v53 = v52
    var v80 bool = v1
    var v81 uint16 = 458
    v82 := GenS11{F0: 2, F1: 1000, F2: -9223372036854775808}
    for fi83 := 0; fi83 < 8; fi83++ {
        v3 -= 3.14159
        v80 = v1
        fmt.Println(v53, v1)
        _ = fi83
    }
    for fi84 := 0; fi84 < 2; fi84++ {
        if true {
            for fi85 := 0; fi85 < 15; fi85++ {
                var v86 bool = (0 != 127)
                _ = fi85
                _ = v86
            }
        } else {
            for fi87 := 0; fi87 < 14; fi87++ {
                v4["foo bar"] = fi84
                v1 = v80
                var v88 uint = (v53 << ((v52) & 63))
                _ = fi87
                _ = v88
            }
            v89 := []int{(fi84 >> (((^(v5))) & 63)), (fi84 * fi84), fi84, 9223372036854775807}
            _ = v89
        }
        for fi90 := 0; fi90 < 2; fi90++ {
            v80 = (!(42 < 1))
            if false {
                var v91 int = 822
                switch v91 {
                case 195:
                    v92 := []int{(fi84 / (((v91 % ((128) | 1))) | 1)), (v91 * (-(v91))), v91, fi84}
                    _ = v92
                case 4294967295:
                    fmt.Println(v81, v5)
                case 42:
                    switch ("hello" + "hello") {
                    case "":
                        fmt.Println(v53, v5)
                    case "a":
                        var v93 string = ""
                        _ = v93
                    }
                    fmt.Println(v57, v1)
                }
                v5 -= (^(v53))
                _ = v91
            } else {
                v2["\n\t"] = (fi90 / ((32768) | 1))
                v94 := GenS2{F0: 100, F1: -76, F2: fi90}
                _ = v94
            }
            var v95 uint32 = 2
            _ = fi90
            _ = v95
        }
        for fi96 := 0; fi96 < 18; fi96++ {
            var v97 int16 = 32767
            _ = fi96
            _ = v97
        }
        _ = fi84
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v57:", v57)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v57
    _ = v79
    _ = v80
    _ = v81
    _ = v82
}
