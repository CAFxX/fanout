package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int32 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint32; F1 int }
func genFn0(p0 int32) int8 {
    var v5 string = "a"
    switch (v5 + v5) {
    case "hello":
        v6 := GenS1{F0: 127, F1: v5, F2: 42}
        _ = v6
    case "":
        fmt.Println(p0, v5)
    case "a":
        if true {
            var v7 uint16 = 255
            _ = v7
        } else {
            fmt.Println(v5)
        }
    }
    switch v5 {
    case "a":
        v8 := []int{127}
        _ = v8
    }
    v5 = ("a" + "a")
    _ = p0
    _ = v5
    return -127
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"\n\t", (("foo bar" + "foo bar") + (((("" + "hello") + "") + ("hello" + "a")) + ("a" + "foo bar"))), "foo bar", (("a" + "hello") + "a")}
    v3 := []uint32{483}
    var v4 int = -9223372036854775808
    v4 -= (v4 * 9)
    if (2 > 706) {
        switch "a" {
        case "foo bar":
            v9 := v1[("a" + ("" + (("foo bar" + "hello") + "hello")))]
            _ = v9
        case "hello":
            v4 = (-(v4))
        case "\n\t":
            v4 = v4
        default:
            var v10 uint16 = 2
            _ = v10
        }
    } else {
        if (!(-632 >= 1)) {
            delete(v1, "hello")
            var v11 uint8 = 255
            v12 := GenS2{F0: 302, F1: 2147483647, F2: 128}
            _ = v11
            _ = v12
        }
        if false {
            var v13 bool = (2 == 964)
            var v14 int = v4
            _ = len(v2)
            _ = v13
            _ = v14
        }
    }
    v4 += v4
    var v15 string = (("hello" + "a") + ("" + ("foo bar" + ("foo bar" + (("a" + "hello") + "hello")))))
    _ = len(v2)
    var v16 uint32 = 2147483647
    v16 -= v16
    if len(v2) > 0 {
        v2[0] = ""
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        _ = len(v1)
        switch v15 {
        case "foo bar":
            if false {
                v18 := v1["foo bar"]
                if (238 < 218) {
                    for fi19 := 0; fi19 < 20; fi19++ {
                        var v20 float64 = 1e10
                        _ = fi19
                        _ = v20
                    }
                    var v21 uint8 = 2
                    _ = v21
                }
                _ = v18
            }
        default:
            for fi22 := 0; fi22 < 6; fi22++ {
                if ((0 >= 1) || (0.1 < 3.14159)) {
                    fmt.Println(fi22, fi17)
                    delete(v1, "")
                    if len(v2) > 0 {
                        v2[0] = ("hello" + (v15 + v15))
                    }
                    v23 := GenS3{F0: 0}
                    if (-951 < 128) {
                        var v24 uint16 = 0
                        fmt.Println(fi22, v16, v4)
                        _ = v24
                    }
                    if false {
                        switch v15 {
                        case "hello":
                            if (fi17 == fi22) {
                                var v25 float32 = (3.14159 * 1e10)
                                fmt.Println(v25, v25)
                                switch v4 {
                                case -2147483648:
                                    var v26 string = "foo bar"
                                    _ = v26
                                case 2147483647:
                                    if false {
                                        var v27 uint64 = 4294967295
                                        v15 = ("foo bar" + ("hello" + v15))
                                        v28 := GenS4{F0: 255, F1: 1}
                                        _ = v27
                                        _ = v28
                                    }
                                }
                                var v29 string = v15
                                _ = v25
                                _ = v29
                            }
                        case "a":
                            if (!true) {
                                v4 -= (fi22 - (fi22 & (v4 & 0)))
                                var v30 float64 = (float64((3.14159 + ((1.5 + 3.14159) - 3.14159))) * ((((0.1 - 0.0) * 1e10) - (0.0 / 1.5)) * float64(float32(1e10))))
                                switch (v15 + v15) {
                                case "hello":
                                    switch "a" {
                                    case "":
                                        if (0 != 42) {
                                            v31 := GenS1{F0: 0, F1: "hello", F2: -1}
                                            _ = v31
                                        } else {
                                            if ((false && true) || false) {
                                                var v32 int8 = -1
                                                fmt.Println(fi17, v4, v32)
                                                _ = v32
                                            } else {
                                                if false {
                                                    _ = len(v1)
                                                    var v33 int16 = 166
                                                    v1["\n\t"] = 2
                                                    _ = v33
                                                }
                                                if (!false) {
                                                    v15 = v15
                                                }
                                                var v34 int64 = 0
                                                _ = v34
                                            }
                                            switch int(v16) {
                                            case -9223372036854775808:
                                                var v35 uint = 65535
                                                _ = v35
                                            case -788:
                                                v15 = "\n\t"
                                                v36 := GenS5{F0: (v16 * 246), F1: (v4 << ((77) & 63))}
                                                v1["a"] = 969
                                                if (false && false) {
                                                    switch v15 {
                                                    case "foo bar":
                                                        if (!(true && (637 < 42))) {
                                                            var v37 float64 = float64(0.0)
                                                            v15 = (((v15 + v15) + v15) + ("hello" + ""))
                                                            fmt.Println(fi17, fi17)
                                                            _ = v37
                                                        }
                                                    case "hello":
                                                        v4 -= 9223372036854775807
                                                    case "a":
                                                        v38 := GenS1{F0: 116, F1: v15, F2: 0}
                                                        _ = v38
                                                    default:
                                                        v39 := GenS3{F0: 2}
                                                        _ = v39
                                                    }
                                                    var v40 bool = false
                                                    _ = v40
                                                }
                                                if false {
                                                    v16 = 2147483647
                                                } else {
                                                    if true {
                                                        var v41 int32 = 864
                                                        _ = len(v2)
                                                        if true {
                                                            _ = len(v1)
                                                        }
                                                        _ = v41
                                                    } else {
                                                        if ((1e10 + 0.1) < -2.25) {
                                                            v42 := GenS2{F0: (v16 - (v16 << ((18446744073709551615) & 63))), F1: 129, F2: 1}
                                                            switch int(v16) {
                                                            case 42:
                                                                v43 := GenS2{F0: (v16 - 4294967295), F1: -1, F2: 128}
                                                                _ = v43
                                                            case 727:
                                                                if len(v2) > 0 {
                                                                    v2[0] = v15
                                                                }
                                                            case 2147483647:
                                                                v44 := GenS5{F0: v16, F1: (^(fi17))}
                                                                _ = v44
                                                            }
                                                            _ = v42
                                                        }
                                                        v16 = (^(v16))
                                                    }
                                                    v45 := GenS2{F0: (v16 | 42), F1: 594, F2: 200}
                                                    _ = v45
                                                }
                                                _ = v36
                                            case 0:
                                                if (-420 == 2147483647) {
                                                    v16 = v16
                                                }
                                            default:
                                                if true {
                                                    if (((true && ((128 >= 3) || (127 > 228))) && (-1 <= 42)) || (32767 > 374)) {
                                                        switch (((v15 + v15) + v15) + (v15 + v15)) {
                                                        case "hello":
                                                            v16 = (^(v16))
                                                        case "foo bar":
                                                            var v46 int64 = -1
                                                            _ = v46
                                                        case "":
                                                            v15 = ("\n\t" + v15)
                                                        }
                                                        v2 = append(v2, "a")
                                                    }
                                                    v15 = ((v15 + "foo bar") + v15)
                                                }
                                            }
                                            _ = len(v3)
                                            if (true && ((v4 ^ fi22) != 0)) {
                                                var v47 bool = (255 >= 255)
                                                _ = v47
                                            } else {
                                                fmt.Println(fi22)
                                                fmt.Println(v4, v16)
                                                v4 = (^(fi22))
                                                v48 := GenS5{F0: v16, F1: 9223372036854775807}
                                                _ = v48
                                            }
                                        }
                                    }
                                case "foo bar":
                                    var v49 int = 4294967295
                                    _ = v49
                                case "a":
                                    if false {
                                        switch v15 {
                                        case "":
                                            switch v4 {
                                            case 32767:
                                                if ((-1 == -32768) || true) {
                                                    var v50 float64 = 0.1
                                                    _ = v50
                                                } else {
                                                    switch v15 {
                                                    case "a":
                                                        _ = len(v1)
                                                        if false {
                                                            v16 += v16
                                                            v51 := GenS3{F0: 255}
                                                            v16 = 65535
                                                            _ = v51
                                                        } else {
                                                            var v52 uint8 = 255
                                                            _ = v52
                                                        }
                                                        _ = len(v1)
                                                        v15 = (v15 + v15)
                                                        var v53 uint32 = (v16 << ((1) & 63))
                                                        _ = v53
                                                    }
                                                    if false {
                                                        v54 := GenS1{F0: 127, F1: "hello", F2: 2147483647}
                                                        fmt.Println(v16)
                                                        var v55 int16 = 1000
                                                        _ = v54
                                                        _ = v55
                                                    } else {
                                                        var v56 int64 = -972
                                                        switch (-(fi22)) {
                                                        case -9223372036854775808:
                                                            v1[("a" + "\n\t")] = v16
                                                            var v57 string = v15
                                                            switch -547 {
                                                            case 32767:
                                                                v30 = float64(3.14159)
                                                            }
                                                            _ = len(v1)
                                                            var v58 bool = false
                                                            v16 = (v16 - v16)
                                                            var v59 int16 = int16(fi17)
                                                            _ = v57
                                                            _ = v58
                                                            _ = v59
                                                        case -470:
                                                            v30 = (float64((3.14159 + (1.5 * 1.5))) * -2.25)
                                                        default:
                                                            var v60 uint64 = 644
                                                            _ = v60
                                                        }
                                                        var v61 float64 = ((v30 - v30) * (v30 - (0.0 * (v30 + v30))))
                                                        _ = v56
                                                        _ = v61
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            v1[v15] = v16
                                        }
                                        var v62 uint16 = 332
                                        _ = v62
                                    }
                                }
                                _ = v30
                            }
                        }
                    }
                    fmt.Println(fi22, fi22)
                    _ = v23
                }
                var v63 bool = true
                var v64 bool = v63
                _ = fi22
                _ = v63
                _ = v64
            }
        }
        for fi65 := 0; fi65 < 5; fi65++ {
            fmt.Println(v15)
            v15 = (v15 + (v15 + "\n\t"))
            _ = fi65
        }
        _ = fi17
    }
    v66 := GenS2{F0: 1, F1: 1, F2: 127}
    var v67 int = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v67
}
