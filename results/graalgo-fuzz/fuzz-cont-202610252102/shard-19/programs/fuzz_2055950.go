package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 uint; F2 int16 }
type GenS7 struct { F0 uint; F1 uint32 }
type GenS8 struct { F0 string }
func genFn0(p0 int16, p1 uint8, p2 uint) int32 {
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: "foo bar"}
        if ((p0 & 1000) == 42) {
            switch (("foo bar" + "hello") + "hello") {
            case "a":
                fmt.Println(p2, p0, p2)
            }
            v8 := GenS1{F0: ""}
            _ = v8
        } else {
            var v9 uint8 = (^(p1))
            var v10 int32 = 42
            _ = v10
            _ = v9
        }
        if (933 < 428) {
            v11 := GenS1{F0: "\n\t"}
            _ = v11
        } else {
            var v12 bool = ((false && true) || true)
            p0 = p0
            var v13 uint16 = 2
            _ = v12
            _ = v13
        }
        _ = fi6
        _ = v7
    }
    var v14 int8 = 127
    v15 := []int{32768, 32767, 2147483647, 0}
    v14 += (v14 % (((v14 / ((v14) | 1))) | 1))
    v16 := GenS1{F0: "\n\t"}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    return 2147483647
}

func main() {
    v1 := []uint32{42, 393}
    v2 := []int{-2147483648, -9223372036854775808, 4294967295}
    v3 := []uint32{669}
    var v4 float64 = 3.14159
    v5 := make(map[string]uint32)
    v4 += v4
    v4 = 3.14159
    if true {
        v4 = float64(0.1)
        v17 := GenS2{F0: 255}
        _ = v17
    } else {
        var v18 int64 = -1
        fmt.Println(v18, v18)
        _ = v18
    }
    v4 += float64(float32(v4))
    if (v4 >= v4) {
        if false {
            v4 = v4
            fmt.Println(v4)
            v4 += 0.1
        }
        fmt.Println(v4)
    } else {
        fmt.Println(v4)
        if ((false && false) || (0 < 0)) {
            switch "a" {
            case "\n\t":
                switch "\n\t" {
                case "foo bar":
                    var v19 float64 = (((1.5 / v4) / float64(float32(1.5))) * float64((0.1 * (((1.5 - ((1e10 * 1.5) + 3.14159)) - 0.0) - 1e10))))
                    _ = v19
                case "hello":
                    v20 := GenS2{F0: 77}
                    _ = v20
                case "":
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                }
            case "":
                fmt.Println(v4)
            case "foo bar":
                for fi21 := 0; fi21 < 1; fi21++ {
                    v22 := GenS3{F0: ("\n\t" + "\n\t"), F1: fi21, F2: ("hello" + "")}
                    _ = fi21
                    _ = v22
                }
            default:
                var v23 int8 = 1
                _ = v23
            }
            v24 := GenS4{F0: 0}
            v4 += v4
            _ = v24
        } else {
            if (-805 <= 32767) {
                var v25 float64 = (v4 * (v4 * -2.25))
                _ = v25
            } else {
                var v26 uint = 972
                v4 = (0.1 / (-2.25 * 0.1))
                _ = v26
            }
        }
        var v27 int8 = 60
        _ = v27
    }
    var v28 uint64 = 4294967295
    v28 = v28
    fmt.Println(v4)
    if (2147483647 >= 32767) {
        var v29 float32 = 1.5
        var v30 int = 32767
        v30 = v30
        _ = v29
        _ = v30
    } else {
        fmt.Println(v4, v28)
        var v31 string = ("" + ((((("\n\t" + "a") + "") + "\n\t") + "a") + ""))
        _ = v31
    }
    switch 2147483647 {
    case 1:
        var v32 int = 0
        _ = v32
    default:
        if false {
            v33 := GenS5{F0: "a"}
            v34 := GenS3{F0: "foo bar", F1: 128, F2: ("" + "")}
            v28 = v28
            _ = v33
            _ = v34
        } else {
            v35 := GenS6{F0: 0, F1: 255, F2: -32768}
            v28 = (v28 ^ v28)
            _ = v35
        }
    }
    var v36 float64 = 1.5
    v36 = v36
    v36 = (3.14159 * v4)
    for fi37 := 0; fi37 < 15; fi37++ {
        switch ("foo bar" + "a") {
        case "hello":
            v36 += (-2.25 * 1.5)
        case "":
            v4 = v4
        case "foo bar":
            var v38 string = "\n\t"
            _ = v38
        default:
            v36 = ((((v36 - (-2.25 + v36)) - v36) * 3.14159) / (v4 * v36))
        }
        if (884 != 2) {
            fmt.Println(v36, v4, v28)
        } else {
            var v39 int16 = -1
            var v40 uint8 = 127
            _ = v39
            _ = v40
        }
        if (4294967295 == 0) {
            if true {
                var v41 float64 = v4
                _ = v41
            }
            v36 += v36
            for fi42 := 0; fi42 < 3; fi42++ {
                if false {
                    if (true || (127 >= 0)) {
                        var v43 uint32 = 4294967295
                        if true {
                            if false {
                                v44 := GenS7{F0: 65535, F1: (v43 % ((367) | 1))}
                                if true {
                                    if (!true) {
                                        var v45 int32 = 0
                                        v4 -= v36
                                        if false {
                                            switch fi42 {
                                            case 1:
                                                var v46 float32 = 0.1
                                                _ = v46
                                            case 32767:
                                                delete(v5, "\n\t")
                                            case 600:
                                                var v47 string = "\n\t"
                                                _ = v47
                                            default:
                                                v45 = (v45 & (^(v45)))
                                            }
                                        }
                                        var v48 uint16 = 413
                                        _ = v45
                                        _ = v48
                                    }
                                    var v49 int32 = 42
                                    _ = len(v5)
                                    var v50 int64 = -9223372036854775808
                                    _ = v49
                                    _ = v50
                                }
                                _ = v44
                            } else {
                                if true {
                                    var v51 uint = 42
                                    if ("" < ("hello" + ("hello" + "\n\t"))) {
                                        v5["foo bar"] = v43
                                        var v52 bool = (0 != 0)
                                        v43 = v43
                                        _ = v52
                                    }
                                    _ = v51
                                }
                            }
                            v53 := v5["a"]
                            _ = v53
                        }
                        _ = v43
                    } else {
                        if true {
                            var v54 int32 = 1
                            _ = v54
                        } else {
                            if true {
                                v55 := v5["a"]
                                if true {
                                    if (1000 < 42) {
                                        var v56 float32 = 0.0
                                        v57 := GenS8{F0: "a"}
                                        v1 = append(v1, (v55 - (v55 + (^(v55)))))
                                        v2 = append(v2, fi42)
                                        var v58 string = "\n\t"
                                        v59 := GenS1{F0: ("a" + "hello")}
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                        _ = v59
                                    } else {
                                        var v60 uint = 65535
                                        _ = v60
                                    }
                                    v61 := v5[""]
                                    v36 = (((float64(0.1) / v4) + 3.14159) + 1e10)
                                    var v62 float32 = 0.1
                                    var v63 uint8 = 42
                                    v4 += 0.0
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                }
                                var v64 float64 = v36
                                _ = len(v5)
                                _ = v55
                                _ = v64
                            }
                        }
                    }
                }
                if false {
                    fmt.Println(fi37, fi42, v36)
                } else {
                    var v65 uint16 = 0
                    v65 = (v65 % ((v65) | 1))
                    v4 = v36
                    _ = v65
                }
                v36 += (((v36 * float64(3.14159)) + v4) + (v4 + float64(1.5)))
                v28 = (v28 - v28)
                var v66 int32 = -764
                _ = fi42
                _ = v66
            }
        } else {
            v5["a"] = 2147483647
        }
        _ = fi37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
