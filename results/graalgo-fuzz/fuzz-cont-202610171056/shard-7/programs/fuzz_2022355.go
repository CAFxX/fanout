package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 string; F2 int64 }
type GenS3 struct { F0 int32; F1 int; F2 uint }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint }
type GenS5 struct { F0 int32; F1 int8; F2 uint32 }
func genFn0(p0 int16, p1 int) bool {
    v6 := GenS1{F0: -1}
    p1 -= -2147483648
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: -1}
        p0 += int16(p1)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return true
}

func main() {
    var v1 uint32 = 600
    var v2 uint64 = 1
    var v3 int = int(v1)
    v4 := []int{v3}
    var v5 uint = 280
    for fi9 := 0; fi9 < 5; fi9++ {
        v1 += (v1 + (^(v1)))
        _ = fi9
    }
    var v10 int8 = 0
    if ((!false) && (false || (!true))) {
        fmt.Println(v1, v10, v1)
    }
    v4 = append(v4, v3)
    var v11 bool = ((("a" + ("\n\t" + ("a" + "hello"))) + ("hello" + ("hello" + ""))) >= ("hello" + "foo bar"))
    if (v11 || (1 >= v3)) {
        fmt.Println(v11, v1)
        _ = len(v4)
        var v12 int = v3
        _ = v12
    }
    if (2 == (v5 ^ (^(v5)))) {
        fmt.Println(v2, v11, v5)
        var v13 string = "\n\t"
        if true {
            if v11 {
                switch "foo bar" {
                case "":
                    v11 = (!(!(!(((^(v5)) >= (v5 >> (((v5 | 4294967295)) & 63))) || ((v5 / (((v5 * 673)) | 1)) != (v5 << ((v5) & 63)))))))
                default:
                    switch 269 {
                    case 9223372036854775807:
                        for fi14 := 0; fi14 < 15; fi14++ {
                            for fi15 := 0; fi15 < 6; fi15++ {
                                switch (v3 ^ (fi15 & 517)) {
                                case 226:
                                    var v16 string = "a"
                                    _ = v16
                                case -184:
                                    v13 = (v13 + v13)
                                    v10 += (v10 + v10)
                                    switch v13 {
                                    case "foo bar":
                                        var v17 bool = v11
                                        _ = v17
                                    }
                                    _ = len(v4)
                                case 128:
                                    var v18 int = 9223372036854775807
                                    _ = v18
                                default:
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(v2, v11, v11)
                                    case "\n\t":
                                        var v19 int32 = -481
                                        _ = v19
                                    case "":
                                        v13 = v13
                                    }
                                }
                                v20 := GenS1{F0: -2147483648}
                                v1 = v1
                                _ = fi15
                                _ = v20
                            }
                            switch -1 {
                            case 4294967295:
                                v13 = v13
                            case -1:
                                if (!((false && (true || true)) || (!((v13 + v13) != v13)))) {
                                    v4 = append(v4, v3)
                                }
                            default:
                                for fi21 := 0; fi21 < 14; fi21++ {
                                    if (!(false || ((v11 && (!(v11 && false))) || v11))) {
                                        var v22 uint = (v5 ^ 2)
                                        v23 := GenS1{F0: 1}
                                        _ = v22
                                        _ = v23
                                    } else {
                                        _ = len(v4)
                                        v24 := GenS1{F0: -22}
                                        v4 = append(v4, 0)
                                        _ = v24
                                    }
                                    var v25 uint32 = 4294967295
                                    v26 := GenS1{F0: -638}
                                    _ = fi21
                                    _ = v25
                                    _ = v26
                                }
                            }
                            fmt.Println(v13)
                            _ = fi14
                        }
                    case -1:
                        for fi27 := 0; fi27 < 14; fi27++ {
                            var v28 int64 = -1
                            _ = fi27
                            _ = v28
                        }
                    default:
                        v29 := GenS1{F0: 205}
                        _ = v29
                    }
                }
            } else {
                v30 := GenS2{F0: (v2 - (v2 & v2)), F1: "foo bar", F2: -1}
                _ = len(v4)
                _ = v30
            }
        }
        _ = v13
    } else {
        if len(v4) > 0 {
            v4[0] = 42
        }
    }
    fmt.Println(v3, v10, v11)
    var v31 string = ("\n\t" + ((("a" + "hello") + ("\n\t" + "")) + "hello"))
    v32 := GenS2{F0: (v2 >> (((v5 * (v5 + v5))) & 63)), F1: v31, F2: 810}
    v11 = (!(!(!(v11 || (!v11)))))
    if v11 {
        v33 := GenS1{F0: 1}
        _ = v33
    }
    var v34 int8 = v10
    var v35 int = (^(v3))
    var v36 uint64 = v2
    v37 := GenS3{F0: 1, F1: 772, F2: v5}
    for fi38 := 0; fi38 < 20; fi38++ {
        if (2 == 32767) {
            var v39 int32 = -873
            if (v11 || true) {
                switch "\n\t" {
                case "\n\t":
                    v31 = (v31 + v31)
                case "":
                    v40 := GenS4{F0: (v2 | (^(v36))), F1: v1, F2: 18446744073709551615}
                    _ = v40
                case "a":
                    var v41 int = v3
                    _ = v41
                default:
                    fmt.Println(v5, v31)
                }
            } else {
                var v42 int = (v3 | v35)
                _ = v42
            }
            v43 := GenS4{F0: (^(v36)), F1: v1, F2: (v5 | (v5 + v5))}
            _ = v39
            _ = v43
        } else {
            for fi44 := 0; fi44 < 15; fi44++ {
                fmt.Println(v5, fi44)
                if ((-274 < 42) && (float64(1e10) > 1.5)) {
                    v36 = v2
                }
                var v45 float64 = 0.1
                _ = fi44
                _ = v45
            }
            for fi46 := 0; fi46 < 18; fi46++ {
                var v47 string = (v31 + "hello")
                _ = fi46
                _ = v47
            }
        }
        for fi48 := 0; fi48 < 19; fi48++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi48
        }
        _ = fi38
    }
    v49 := GenS5{F0: -1, F1: v10, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v49
    _ = v5
}
