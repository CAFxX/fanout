package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16; F1 int64 }
func genFn0(p0 uint8, p1 bool) bool {
    p0 = p0
    p1 = (65535 >= 718)
    var v4 uint = 4294967295
    p0 = (p0 >> (((v4 >> ((v4) & 63))) & 63))
    for fi5 := 0; fi5 < 14; fi5++ {
        if false {
            if (fi5 == -419) {
                v6 := []int{fi5}
                var v7 uint8 = p0
                for fi8 := 0; fi8 < 1; fi8++ {
                    var v9 uint8 = (v7 % (((p0 % ((v7) | 1))) | 1))
                    v10 := GenS1{F0: fi5, F1: v7, F2: 32767}
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = v6
                _ = v7
            }
            for fi11 := 0; fi11 < 10; fi11++ {
                v12 := []int{fi5}
                fmt.Println(v4)
                var v13 float64 = ((float64((1.5 + 0.1)) - 0.0) + 0.0)
                var v14 bool = ((!("\n\t" >= "\n\t")) && (!(!true)))
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        }
        if ("hello" < ("\n\t" + ((("" + "a") + "a") + ("hello" + "")))) {
            if (!p1) {
                var v15 int16 = -32768
                v15 = (v15 ^ 1)
                _ = v15
            }
            if ((fi5 | (fi5 << ((42) & 63))) == 32767) {
                v16 := []int{(fi5 * -2147483648), (fi5 + fi5), -2147483648}
                p1 = p1
                switch (("" + (("a" + ("\n\t" + "")) + "\n\t")) + ("a" + "")) {
                case "":
                    switch "" {
                    case "hello":
                        var v17 int16 = -705
                        _ = v17
                    }
                case "foo bar":
                    var v18 float32 = float32(-2.25)
                    _ = v18
                case "a":
                    var v19 int32 = 42
                    _ = v19
                default:
                    if (p1 && (!p1)) {
                        var v20 uint = 172
                        for fi21 := 0; fi21 < 7; fi21++ {
                            v22 := GenS1{F0: (^(fi21)), F1: 0, F2: 42}
                            _ = fi21
                            _ = v22
                        }
                        _ = v20
                    } else {
                        v23 := GenS1{F0: (fi5 << ((841) & 63)), F1: (p0 - 255), F2: 32767}
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        if (128 <= 1) {
                            if len(v16) > 0 {
                                v16[0] = (fi5 ^ (fi5 - (fi5 ^ (fi5 << ((v4) & 63)))))
                            }
                            if p1 {
                                v24 := GenS1{F0: -411, F1: 42, F2: 0}
                                p1 = false
                                _ = v24
                            } else {
                                v25 := GenS2{F0: (p0 * p0)}
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                                var v26 float64 = (3.14159 - (3.14159 * -2.25))
                                _ = v25
                                _ = v26
                            }
                        } else {
                            for fi27 := 0; fi27 < 20; fi27++ {
                                fmt.Println(p0, fi5, p1)
                                _ = fi27
                            }
                            switch -998 {
                            case -1:
                                var v28 string = ""
                                _ = v28
                            case 9223372036854775807:
                                p0 = p0
                            default:
                                var v29 int8 = -42
                                _ = v29
                            }
                        }
                        _ = v23
                    }
                }
                _ = v16
            } else {
                v30 := GenS1{F0: int(p0), F1: (p0 - uint8(v4)), F2: int16(p0)}
                fmt.Println(p1)
                _ = v30
            }
        } else {
            var v31 int32 = 1
            fmt.Println(p0, fi5, p0)
            p0 = 2
            _ = v31
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return p1
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"a", ("" + (("hello" + "a") + "foo bar"))}
    var v3 int16 = 848
    v32 := GenS1{F0: 4294967295, F1: 42, F2: (v3 + (v3 - 377))}
    if false {
        v3 -= (v3 / (((v3 - v3)) | 1))
        v3 -= v3
    }
    var v33 uint16 = 949
    fmt.Println(v3, v3)
    delete(v1, ("foo bar" + "\n\t"))
    for fi34 := 0; fi34 < 1; fi34++ {
        fmt.Println(v3, fi34)
        for fi35 := 0; fi35 < 9; fi35++ {
            v36 := GenS1{F0: -618, F1: 127, F2: v3}
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    var v37 bool = true
    switch -9223372036854775808 {
    case -2147483648:
        v37 = v37
    case -1:
        if (0 > 42) {
            var v38 int32 = -1
            if (!v37) {
                v33 = v33
            }
            delete(v1, "hello")
            _ = v38
        } else {
            var v39 uint32 = 2147483648
            _ = v39
        }
    default:
        if v37 {
            var v40 uint8 = 127
            for fi41 := 0; fi41 < 9; fi41++ {
                v37 = v37
                _ = fi41
            }
            _ = v40
        }
    }
    switch -9223372036854775808 {
    case -384:
        if (v37 && v37) {
            switch 2147483647 {
            case 32:
                for fi42 := 0; fi42 < 15; fi42++ {
                    var v43 uint = 4294967295
                    _ = fi42
                    _ = v43
                }
            case 9223372036854775807:
                fmt.Println(v33, v37)
            }
            v37 = (0 < 94)
        } else {
            v33 += 354
            switch "foo bar" {
            case "":
                v44 := GenS3{F0: 2}
                _ = v44
            case "hello":
                if (559 <= 219) {
                    for fi45 := 0; fi45 < 6; fi45++ {
                        if ((127 < 56) || (!v37)) {
                            switch (fi45 | fi45) {
                            case -800:
                                for fi46 := 0; fi46 < 19; fi46++ {
                                    v3 = 774
                                    switch -677 {
                                    case 32768:
                                        v47 := GenS3{F0: 0}
                                        _ = v47
                                    case -538:
                                        v1["foo bar"] = (v37 || true)
                                    case 862:
                                        var v48 int8 = 1
                                        _ = v48
                                    default:
                                        v37 = (0 >= 0)
                                    }
                                    _ = fi46
                                }
                            }
                        }
                        _ = fi45
                    }
                    var v49 float64 = (3.14159 * 3.14159)
                    for fi50 := 0; fi50 < 9; fi50++ {
                        for fi51 := 0; fi51 < 13; fi51++ {
                            var v52 uint32 = 0
                            v52 += v52
                            _ = fi51
                            _ = v52
                        }
                        v53 := GenS4{F0: 42, F1: -107}
                        _ = fi50
                        _ = v53
                    }
                    _ = v49
                } else {
                    fmt.Println(v3, v3)
                }
            case "foo bar":
                switch "a" {
                case "":
                    var v54 int64 = -1
                    _ = v54
                case "foo bar":
                    _ = len(v1)
                case "hello":
                    fmt.Println(v33)
                default:
                    var v55 float64 = 1e10
                    _ = v55
                }
            default:
                if (!(v33 >= v33)) {
                    if v37 {
                        var v56 int = -565
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 16; fi57++ {
                        for fi58 := 0; fi58 < 10; fi58++ {
                            if (-128 <= 127) {
                                v3 += v3
                                v3 += v3
                            } else {
                                switch ("\n\t" + "hello") {
                                case "\n\t":
                                    var v59 bool = v37
                                    _ = v59
                                case "foo bar":
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                case "hello":
                                    var v60 uint32 = 4294967295
                                    _ = v60
                                }
                                _ = len(v1)
                                var v61 uint8 = 127
                                _ = v61
                            }
                            v37 = v37
                            var v62 uint8 = 42
                            v63 := v1["hello"]
                            _ = fi58
                            _ = v62
                            _ = v63
                        }
                        _ = fi57
                    }
                    if true {
                        var v64 int16 = v3
                        v65 := GenS2{F0: 127}
                        _ = v64
                        _ = v65
                    }
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case 127:
        fmt.Println(v33, v3)
    default:
        for fi66 := 0; fi66 < 10; fi66++ {
            fmt.Println(fi66, v3, v37)
            var v67 int64 = 1
            var v68 string = "a"
            _ = fi66
            _ = v67
            _ = v68
        }
    }
    v33 += 1
    fmt.Println(v3, v37)
    if ((v37 || (42 < 0)) && false) {
        var v69 float64 = float64(float32(float64((3.14159 / 3.14159))))
        for fi70 := 0; fi70 < 17; fi70++ {
            v71 := GenS5{F0: 46}
            _ = len(v2)
            delete(v1, ("a" + "foo bar"))
            _ = fi70
            _ = v71
        }
        var v72 int16 = v3
        _ = v69
        _ = v72
    }
    v73 := GenS6{F0: (-(v3)), F1: 1}
    v37 = ("a" >= ("hello" + ("hello" + "foo bar")))
    v3 += (v3 << ((42) & 63))
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v73
}
