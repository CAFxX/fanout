package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 int64; F1 int }
func genFn0(p0 string, p1 uint64) int {
    var v6 float64 = float64(3.14159)
    switch (p0 + "hello") {
    case "hello":
        fmt.Println(p0, p1)
    case "foo bar":
        v7 := GenS1{F0: 100}
        _ = v7
    default:
        var v8 bool = (255 >= 981)
        _ = v8
    }
    v9 := []int{331, -2147483648, 279, 9223372036854775807}
    _ = len(v9)
    v6 += v6
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return 32768
}
func genFn1(p0 uint32) int64 {
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int = (fi10 - fi10)
        fmt.Println(fi10, p0, p0)
        p0 = p0
        _ = fi10
        _ = v11
    }
    p0 = 42
    _ = p0
    return -1
}

func main() {
    var v1 string = ("\n\t" + "\n\t")
    v2 := []string{(v1 + ("a" + v1))}
    var v3 bool = (2 > 42)
    v4 := make(map[uint32]string)
    var v5 int32 = 2147483647
    fmt.Println(v1, v1)
    if v3 {
        for fi12 := 0; fi12 < 1; fi12++ {
            if v3 {
                v13 := GenS1{F0: 42}
                var v14 float32 = -2.25
                _ = v13
                _ = v14
            } else {
                _ = len(v2)
            }
            if ((v5 ^ v5) >= 2147483647) {
                _ = len(v4)
            } else {
                if v3 {
                    if (-913 <= -781) {
                        v15 := GenS1{F0: 100}
                        v3 = v3
                        v5 = v5
                        _ = v15
                    } else {
                        var v16 int8 = 42
                        if (((float32(0.0) * 0.0) * 0.1) != ((3.14159 + 1.5) + -2.25)) {
                            if len(v2) > 0 {
                                v2[0] = v1
                            }
                            if v3 {
                                var v17 int16 = 42
                                v18 := GenS2{F0: 65535, F1: true}
                                _ = v17
                                _ = v18
                            } else {
                                var v19 uint64 = 2
                                var v20 float32 = -2.25
                                _ = v19
                                _ = v20
                            }
                        }
                        _ = v16
                    }
                    for fi21 := 0; fi21 < 15; fi21++ {
                        v4[2147483647] = "\n\t"
                        v22 := GenS3{F0: 18446744073709551615, F1: (v3 || (true && (42 >= 65535))), F2: 18446744073709551615}
                        if (!false) {
                            v23 := GenS4{F0: 272}
                            var v24 float32 = 3.14159
                            if v3 {
                                v25 := GenS5{F0: -622}
                                _ = v25
                            }
                            _ = v23
                            _ = v24
                        }
                        var v26 int = (fi12 + (fi12 - fi21))
                        _ = fi21
                        _ = v22
                        _ = v26
                    }
                }
                v1 = v1
            }
            fmt.Println(v5, v5, v1)
            _ = fi12
        }
        for fi27 := 0; fi27 < 19; fi27++ {
            switch (-(fi27)) {
            case 42:
                var v28 float32 = -2.25
                _ = v28
            case 32767:
                v5 -= v5
            }
            _ = fi27
        }
        fmt.Println(v1, v1)
    } else {
        var v29 float32 = 1.5
        _ = v29
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        _ = len(v4)
        if v3 {
            fmt.Println(v5, v1)
            switch v1 {
            case "foo bar":
                if (65535 != 0) {
                    var v31 int = fi30
                    _ = v31
                }
            case "hello":
                var v32 uint64 = 4294967295
                _ = v32
            }
        } else {
            var v33 int32 = v5
            for fi34 := 0; fi34 < 10; fi34++ {
                if (972 > 2) {
                    fmt.Println(v5)
                    if (0 < 734) {
                        var v35 int = genFn0("hello", 42)
                        if (1.5 >= 1.5) {
                            v36 := GenS5{F0: -282}
                            if v3 {
                                v3 = (1 == 0)
                                var v37 int64 = 9223372036854775807
                                _ = len(v4)
                                _ = v37
                            } else {
                                v4[42] = v1
                                v38 := v4[0]
                                v2 = append(v2, "a")
                                _ = v38
                            }
                            _ = v36
                        }
                        v3 = true
                        _ = v35
                    }
                    var v39 int = 2147483647
                    v40 := GenS6{F0: (v1 + "a")}
                    _ = v39
                    _ = v40
                }
                _ = fi34
            }
            fmt.Println(v5, v5, v33)
            _ = v33
        }
        _ = fi30
    }
    v4[2] = ((v1 + (v1 + "foo bar")) + v1)
    v1 = ("a" + ("a" + (v1 + "foo bar")))
    fmt.Println(v1, v3)
    delete(v4, 0)
    for fi41 := 0; fi41 < 17; fi41++ {
        for fi42 := 0; fi42 < 19; fi42++ {
            switch v1 {
            case "":
                if true {
                    fmt.Println(v1, v5)
                    v5 += -1
                    fmt.Println(v3, v3)
                    switch genFn0("", 18446744073709551615) {
                    case 171:
                        switch v1 {
                        case "a":
                            switch (v1 + ((("foo bar" + "hello") + "hello") + "foo bar")) {
                            case "hello":
                                if true {
                                    var v43 uint16 = 255
                                    if v3 {
                                        delete(v4, 2)
                                        var v44 int8 = 14
                                        if v3 {
                                            delete(v4, 2)
                                            var v45 int64 = -9223372036854775808
                                            _ = len(v4)
                                            _ = v45
                                        }
                                        _ = v44
                                    } else {
                                        v3 = v3
                                    }
                                    var v46 int32 = v5
                                    _ = v43
                                    _ = v46
                                } else {
                                    var v47 int8 = 64
                                    _ = v47
                                }
                            case "foo bar":
                                var v48 bool = v3
                                _ = v48
                            default:
                                fmt.Println(fi42, v3, fi42)
                            }
                        case "hello":
                            v1 = v1
                            v49 := GenS7{F0: -931}
                            if v3 {
                                v3 = (!((-8 < -1) && v3))
                                fmt.Println(v5, v1, v3)
                                var v50 int32 = -1
                                _ = v50
                            } else {
                                fmt.Println(fi42, v1)
                                var v51 int = 32767
                                if v3 {
                                    if v3 {
                                        var v52 uint32 = 720
                                        _ = v52
                                    } else {
                                        var v53 int64 = 42
                                        v53 -= v53
                                        _ = v53
                                    }
                                    v1 = v1
                                    _ = len(v2)
                                    v1 = "hello"
                                }
                                v54 := GenS8{F0: 0, F1: 18446744073709551615}
                                _ = v51
                                _ = v54
                            }
                            v5 -= (v5 / (((v5 ^ (-(v5)))) | 1))
                            _ = v49
                        case "foo bar":
                            if ((982 >= 65535) && ((v5 << ((255) & 63)) < (^(v5)))) {
                                switch fi41 {
                                case 2147483647:
                                    var v55 int8 = 86
                                    _ = v55
                                case -747:
                                    v56 := GenS7{F0: -32768}
                                    _ = v56
                                }
                                var v57 int8 = 100
                                _ = v57
                            }
                        }
                    }
                } else {
                    v1 = (v1 + (v1 + v1))
                    v1 = v1
                }
                v58 := GenS2{F0: 65535, F1: v3}
                _ = v58
            case "a":
                fmt.Println(v1, fi42, v3)
            case "hello":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                switch "\n\t" {
                case "":
                    v1 = "hello"
                case "a":
                    v59 := GenS5{F0: v5}
                    _ = v59
                default:
                    var v60 uint8 = 2
                    _ = v60
                }
            }
            switch (v1 + "\n\t") {
            case "a":
                var v61 uint16 = 128
                _ = v61
            case "":
                var v62 uint32 = 637
                _ = v62
            default:
                var v63 uint64 = 683
                _ = v63
            }
            v64 := GenS9{F0: 0, F1: -181}
            fmt.Println(v1, v5)
            _ = fi42
            _ = v64
        }
        _ = fi41
    }
    if v3 {
        var v65 int8 = -77
        fmt.Println(v1, v1)
        _ = v65
    } else {
        v3 = v3
    }
    v2 = append(v2, "")
    var v66 int32 = -913
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
}
