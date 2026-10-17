package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int32; F2 int8 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint; F1 int16; F2 uint }
type GenS6 struct { F0 uint32; F1 int8; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (!true)
    var v3 uint64 = 18446744073709551615
    var v4 uint8 = 2
    var v5 bool = ("hello" < (("foo bar" + ("a" + "\n\t")) + ("foo bar" + (("\n\t" + "hello") + (("hello" + "\n\t") + "hello")))))
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: "\n\t"}
    fmt.Println(v4, v3, v2)
    fmt.Println(v5)
    v8 := GenS1{F0: "\n\t"}
    v9 := GenS2{F0: -281, F1: 0, F2: 42}
    if true {
        var v10 int = -277
        for fi11 := 0; fi11 < 3; fi11++ {
            if false {
                v12 := []int{(v10 << ((42) & 63)), (fi11 | (^(v10))), 0}
                _ = v12
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                v14 := GenS1{F0: "\n\t"}
                v4 = 127
                _ = fi13
                _ = v14
            }
            _ = fi11
        }
        v15 := []int{v10, -2147483648, (v10 * (v10 ^ (v10 << ((1) & 63))))}
        _ = v10
        _ = v15
    } else {
        var v16 uint64 = v3
        _ = v16
    }
    v17 := v1[(("a" + "a") + ("a" + "hello"))]
    for fi18 := 0; fi18 < 17; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            if v5 {
                fmt.Println(fi18)
                var v20 int = -220
                _ = v20
            }
            switch 42 {
            case 127:
                v21 := GenS2{F0: 42, F1: -1, F2: -128}
                v22 := GenS2{F0: 689, F1: 833, F2: 0}
                _ = v21
                _ = v22
            default:
                switch (fi18 ^ fi18) {
                case -9223372036854775808:
                    v2 = v5
                    if true {
                        fmt.Println(fi18)
                        if (9223372036854775807 >= -793) {
                            var v23 int64 = -892
                            v4 = v4
                            _ = v23
                        }
                        v24 := GenS3{F0: fi19, F1: 462}
                        var v25 bool = (((float64(1e10) * 1.5) + float64(float32(3.14159))) < 1.5)
                        _ = v24
                        _ = v25
                    } else {
                        var v26 uint32 = 115
                        _ = v26
                    }
                default:
                    fmt.Println(v4)
                }
            }
            _ = fi19
        }
        v5 = v5
        if (v17 <= (^(v17))) {
            v27 := v1["a"]
            v2 = (723 < 1)
            _ = v27
        } else {
            switch 1 {
            case -152:
                v28 := GenS1{F0: ((("foo bar" + (("" + "") + "hello")) + "hello") + "")}
                _ = v28
            case 748:
                var v29 int = fi18
                _ = v29
            case 0:
                fmt.Println(v4, v17, v5)
            default:
                fmt.Println(fi18)
            }
            var v30 bool = (((float32(1.5) + (0.0 - 0.0)) < -2.25) || (100 != 0))
            _ = v30
        }
        _ = fi18
    }
    fmt.Println(v4, v5)
    switch -8 {
    case 128:
        if v2 {
            switch "foo bar" {
            case "hello":
                var v31 uint16 = 255
                _ = v31
            case "":
                var v32 int64 = -657
                _ = v32
            case "a":
                v5 = true
            }
            fmt.Println(v17)
        }
    case -960:
        fmt.Println(v5, v17, v5)
    case 79:
        if false {
            v5 = (!(v5 || false))
            for fi33 := 0; fi33 < 11; fi33++ {
                for fi34 := 0; fi34 < 4; fi34++ {
                    fmt.Println(v3, fi34, v17)
                    if v5 {
                        v3 = v3
                        var v35 uint = uint(v17)
                        delete(v6, ((("foo bar" + "foo bar") + "hello") + "foo bar"))
                        _ = v35
                    } else {
                        v5 = (((0.1 / 1.0) <= float32(1e10)) || (-128 != 42))
                        var v36 int16 = -32768
                        _ = v36
                    }
                    v5 = (!v2)
                    _ = fi34
                }
                var v37 uint16 = 383
                for fi38 := 0; fi38 < 18; fi38++ {
                    var v39 int64 = -30
                    _ = fi38
                    _ = v39
                }
                _ = fi33
                _ = v37
            }
            v3 += (v3 / ((v3) | 1))
        } else {
            v17 -= v17
            v5 = false
            v40 := []int{42, -899, 2147483647}
            _ = v40
        }
    default:
        if v2 {
            v1["\n\t"] = v17
            v3 += (v3 | v3)
            if v5 {
                if v2 {
                    if ("foo bar" > ((("\n\t" + "foo bar") + ("hello" + ("" + ("" + ("hello" + "\n\t"))))) + "")) {
                        if v5 {
                            var v41 uint16 = 32767
                            if v2 {
                                v4 -= (v4 ^ (^(v4)))
                                v3 = 1
                            }
                            _ = v41
                        }
                        if (v5 && (-2.25 > 0.1)) {
                            for fi42 := 0; fi42 < 9; fi42++ {
                                switch 9223372036854775807 {
                                case 1:
                                    fmt.Println(v4, v17)
                                default:
                                    for fi43 := 0; fi43 < 5; fi43++ {
                                        var v44 bool = v5
                                        v45 := v1[""]
                                        _ = fi43
                                        _ = v44
                                        _ = v45
                                    }
                                }
                                _ = fi42
                            }
                        }
                        v46 := GenS2{F0: 0, F1: 749, F2: 19}
                        _ = v46
                    }
                    v47 := GenS3{F0: -755, F1: 571}
                    _ = v47
                }
                v17 += v17
                v48 := GenS1{F0: ("foo bar" + "\n\t")}
                _ = v48
            }
        }
    }
    for fi49 := 0; fi49 < 16; fi49++ {
        var v50 int64 = 0
        _ = fi49
        _ = v50
    }
    switch 42 {
    case -51:
        v51 := GenS4{F0: 83, F1: -2147483648}
        _ = v51
    case 4294967295:
        v5 = true
    case 127:
        v52 := GenS1{F0: "foo bar"}
        _ = v52
    default:
        v53 := GenS2{F0: -106, F1: 42, F2: int8(v4)}
        _ = v53
    }
    switch 607 {
    case 2147483647:
        v17 = 991
    case 42:
        v54 := GenS5{F0: 4294967295, F1: 1, F2: 0}
        _ = v54
    case 9223372036854775807:
        v55 := GenS2{F0: 42, F1: -943, F2: 127}
        _ = v55
    default:
        for fi56 := 0; fi56 < 2; fi56++ {
            v57 := GenS3{F0: fi56, F1: (fi56 % (((fi56 * (fi56 / (((fi56 << ((0) & 63))) | 1)))) | 1))}
            v5 = v2
            _ = fi56
            _ = v57
        }
    }
    var v58 int16 = 0
    v59 := GenS6{F0: 309, F1: 127, F2: 65535}
    if false {
        fmt.Println(v3, v2, v58)
        if (v5 || v2) {
            fmt.Println(v4, v5)
            fmt.Println(v2, v4)
        }
        v2 = v5
    } else {
        if (v4 == v4) {
            var v60 int16 = v58
            _ = v60
        }
    }
    v2 = v5
    var v61 int16 = v58
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v61
    _ = v7
    _ = v8
    _ = v9
}
