package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint8 }
type GenS2 struct { F0 int; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int; F1 uint32; F2 int32 }
func genFn0(p0 int32, p1 int) uint32 {
    var v7 int = p1
    p1 += v7
    fmt.Println(v7)
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := []string{("foo bar" + "foo bar")}
    var v2 uint32 = 220
    var v3 int64 = 17
    v4 := make(map[string]int)
    v5 := []string{"hello", "\n\t"}
    v6 := make(map[string]string)
    var v8 uint8 = 1
    fmt.Println(v8, v8)
    v9 := GenS1{F0: 1000, F1: 1, F2: 255}
    if ((v8 - (^(v8))) == v8) {
        if len(v5) > 0 {
            v5[0] = ""
        }
        var v10 uint16 = 1
        if false {
            v4["foo bar"] = 152
            var v11 uint64 = 18446744073709551615
            v10 = v10
            _ = v11
        } else {
            var v12 uint32 = (v2 >> ((898) & 63))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v12
        }
        _ = v10
    }
    v13 := GenS2{F0: 32767, F1: 32767, F2: (v2 * 2147483648)}
    v2 = genFn0(602, -9223372036854775808)
    switch "a" {
    case "foo bar":
        if false {
            _ = len(v6)
        }
    case "a":
        v14 := GenS3{F0: 65535}
        _ = v14
    case "":
        if false {
            v15 := GenS4{F0: 0, F1: 65535}
            _ = v15
        } else {
            var v16 float64 = ((float64(1.5) * (0.0 * 0.1)) + 1.5)
            v17 := GenS5{F0: -2147483648}
            _ = v16
            _ = v17
        }
    default:
        v4[""] = int(v8)
    }
    fmt.Println(v8)
    v3 = 0
    var v18 int32 = -2147483648
    if true {
        v19 := GenS2{F0: 32767, F1: 32767, F2: v2}
        switch 4294967295 {
        case 948:
            v18 = (v18 ^ 581)
        case 264:
            v20 := GenS6{F0: 32767, F1: (false && (2 >= 4294967295))}
            _ = v20
        case -100:
            _ = len(v1)
        }
        _ = v19
    }
    switch ("foo bar" + (("\n\t" + "hello") + "a")) {
    case "":
        v18 += v18
    default:
        if (2 < 4294967295) {
            var v21 int = 9223372036854775807
            for fi22 := 0; fi22 < 5; fi22++ {
                for fi23 := 0; fi23 < 16; fi23++ {
                    var v24 int8 = 0
                    if true {
                        _ = len(v1)
                        fmt.Println(v18, v2)
                        fmt.Println(v8, v18)
                    }
                    v25 := GenS7{F0: v2, F1: 0, F2: 255}
                    var v26 int64 = (v3 >> ((497) & 63))
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                if (false && (!false)) {
                    var v27 uint16 = 259
                    v27 += 32767
                    var v28 float64 = (3.14159 / 1.0)
                    _ = v27
                    _ = v28
                } else {
                    switch "foo bar" {
                    case "hello":
                        if (!((float32(1.5) / (0.0 - (3.14159 - 1.5))) == float32(float64(3.14159)))) {
                            switch ((("a" + "foo bar") + ("a" + "")) + "a") {
                            case "a":
                                var v29 bool = false
                                _ = v29
                            case "":
                                var v30 uint = 881
                                _ = v30
                            }
                        } else {
                            v8 = (v8 >> ((576) & 63))
                            var v31 string = "\n\t"
                            _ = v31
                        }
                    default:
                        v32 := GenS3{F0: 42}
                        _ = v32
                    }
                    if (0 != 1) {
                        for fi33 := 0; fi33 < 14; fi33++ {
                            v3 = v3
                            switch (("foo bar" + ("\n\t" + "a")) + (("foo bar" + (("foo bar" + (("" + "a") + "")) + (("hello" + ("foo bar" + "foo bar")) + "foo bar"))) + ("a" + "foo bar"))) {
                            case "foo bar":
                                if false {
                                    if true {
                                        if (501 >= 497) {
                                            v2 = v2
                                            switch v21 {
                                            case 127:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            case 32768:
                                                v21 -= fi33
                                            }
                                            v8 = v8
                                        }
                                        if ((true || false) && true) {
                                            if false {
                                                v3 += -9223372036854775808
                                                v1 = append(v1, ("hello" + (("foo bar" + ("" + "foo bar")) + "hello")))
                                                var v34 bool = true
                                                _ = v34
                                            }
                                            switch fi33 {
                                            case 2147483647:
                                                if (0 >= 1) {
                                                    var v35 string = "hello"
                                                    if true {
                                                        if (!(!false)) {
                                                            v1 = append(v1, "foo bar")
                                                        } else {
                                                            var v36 uint16 = uint16(v2)
                                                            var v37 uint16 = (v36 * 32767)
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        var v38 string = ((((v35 + (v35 + v35)) + v35) + "hello") + ("foo bar" + ((v35 + ("" + "foo bar")) + "")))
                                                        v21 -= int(v2)
                                                        _ = v38
                                                    } else {
                                                        v21 = (fi33 % ((fi33) | 1))
                                                        v39 := GenS2{F0: v21, F1: 976, F2: v2}
                                                        var v40 float32 = (0.1 + float32(1.5))
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                    if (v3 < (v3 + v3)) {
                                                        var v41 uint8 = v8
                                                        if true {
                                                            var v42 float64 = ((1e10 - 3.14159) * ((1.5 * (1e10 / 1.0)) - 1e10))
                                                            v43 := GenS8{F0: (v21 | (fi33 >> ((4294967295) & 63)))}
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                        var v44 string = v35
                                                        v21 -= 42
                                                        v18 -= (v18 / (((v18 & 1)) | 1))
                                                        _ = v41
                                                        _ = v44
                                                    }
                                                    _ = v35
                                                }
                                            case 42:
                                                v8 = (v8 & 2)
                                            default:
                                                v18 -= 867
                                            }
                                        } else {
                                            var v45 int16 = 1000
                                            _ = v45
                                        }
                                        fmt.Println(v8)
                                    }
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = ("hello" + "hello")
                                    }
                                }
                            case "a":
                                v46 := GenS6{F0: 960, F1: true}
                                _ = v46
                            }
                            _ = fi33
                        }
                        fmt.Println(v8, v8)
                    }
                    switch ("hello" + ("\n\t" + ("" + ("a" + "foo bar")))) {
                    case "":
                        if (false || false) {
                            switch (("a" + "") + "a") {
                            case "\n\t":
                                for fi47 := 0; fi47 < 10; fi47++ {
                                    v18 = 71
                                    var v48 int16 = 456
                                    v49 := GenS7{F0: (v2 & (v2 << ((547) & 63))), F1: 4294967295, F2: 255}
                                    if false {
                                        if false {
                                            switch "\n\t" {
                                            case "":
                                                var v50 float64 = ((3.14159 + -2.25) - ((1e10 + (-2.25 / 3.14159)) + (0.0 + (3.14159 * (float64(3.14159) * (-2.25 * -2.25))))))
                                                _ = v50
                                            case "hello":
                                                _ = len(v5)
                                            case "foo bar":
                                                fmt.Println(fi47)
                                            }
                                            if true {
                                                v48 -= (v48 * (v48 + (v48 | v48)))
                                            } else {
                                                v4[(((("" + ((("foo bar" + ("" + "foo bar")) + ("hello" + "foo bar")) + ("a" + "hello"))) + ("a" + "hello")) + "hello") + "a")] = v21
                                                v21 = (fi22 << ((0) & 63))
                                            }
                                            v51 := GenS2{F0: (fi22 + fi22), F1: 255, F2: genFn0(int32(v8), int(v2))}
                                            var v52 int16 = v48
                                            _ = v51
                                            _ = v52
                                        } else {
                                            v53 := GenS9{F0: (v21 * fi22), F1: (^(v2)), F2: (v18 % (((v18 % ((2147483647) | 1))) | 1))}
                                            _ = v53
                                        }
                                    } else {
                                        v21 = 127
                                    }
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                            case "hello":
                                fmt.Println(v18, v2, v8)
                            case "foo bar":
                                var v54 float32 = 1e10
                                _ = v54
                            default:
                                v55 := v6["hello"]
                                _ = v55
                            }
                        } else {
                            var v56 string = ("\n\t" + "")
                            v2 = genFn0((v18 * -1), v21)
                            _ = v56
                        }
                    case "hello":
                        fmt.Println(v3, fi22)
                    }
                }
                _ = fi22
            }
            switch "foo bar" {
            case "\n\t":
                v18 = v18
            case "foo bar":
                for fi57 := 0; fi57 < 3; fi57++ {
                    var v58 uint = 114
                    v8 += 255
                    _ = fi57
                    _ = v58
                }
            }
            _ = v21
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
