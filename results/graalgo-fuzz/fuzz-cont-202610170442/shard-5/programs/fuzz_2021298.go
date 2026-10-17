package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 bool }
func genFn0(p0 uint16, p1 uint8, p2 bool) int32 {
    var v5 int = -189
    if (!(("a" + "hello") > "foo bar")) {
        var v6 uint64 = 4294967295
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint64 = 171
            var v9 uint64 = v6
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := []int{}
        _ = v10
        _ = v6
    } else {
        var v11 uint16 = p0
        p2 = p2
        fmt.Println(v5)
        _ = v11
    }
    v12 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v5
    return -2147483648
}
func genFn1(p0 string, p1 int16, p2 uint32) uint64 {
    p0 = "foo bar"
    v13 := GenS1{F0: 32768}
    v14 := GenS2{F0: uint16(p2)}
    p1 = (p1 / ((32767) | 1))
    var v15 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return 4294967295
}

func main() {
    var v1 uint32 = 0
    var v2 float32 = (0.1 / 0.1)
    v3 := []string{}
    v4 := make(map[string]string)
    var v16 uint8 = 128
    var v17 uint64 = 42
    fmt.Println(v17, v16, v2)
    fmt.Println(v2, v2)
    var v18 float32 = v2
    var v19 int16 = 0
    v20 := GenS2{F0: 255}
    if false {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        var v21 int32 = -2147483648
        for fi22 := 0; fi22 < 19; fi22++ {
            v16 = (v16 + 0)
            _ = fi22
        }
        _ = v21
    } else {
        v1 = (v1 | 4294967295)
        _ = len(v4)
    }
    switch ("foo bar" + (("\n\t" + "foo bar") + "\n\t")) {
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            var v23 int64 = -733
            _ = v23
        }
    default:
        v24 := GenS3{F0: 515}
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        _ = len(v3)
        if ("\n\t" > "hello") {
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 uint16 = 65535
                if (true && (1e10 < ((float64(v18) + (1.5 * -2.25)) / (1e10 - -2.25)))) {
                    var v28 bool = ((3.14159 + (1.5 - -2.25)) == (3.14159 - 0.0))
                    var v29 int = fi25
                    _ = v28
                    _ = v29
                } else {
                    fmt.Println(v2)
                    if (v16 == (^(v16))) {
                        var v30 bool = false
                        switch (fi26 & (fi26 & fi25)) {
                        case -879:
                            switch (fi25 >> ((714) & 63)) {
                            case 9223372036854775807:
                                var v31 int = fi26
                                _ = v31
                            case 2147483647:
                                if v30 {
                                    v16 += v16
                                    var v32 uint8 = v16
                                    v33 := GenS3{F0: -2147483648}
                                    _ = v32
                                    _ = v33
                                }
                            }
                        case 127:
                            switch "\n\t" {
                            case "":
                                if true {
                                    var v34 uint32 = (v1 >> ((255) & 63))
                                    var v35 int32 = 2147483647
                                    _ = v34
                                    _ = v35
                                }
                            default:
                                var v36 uint64 = uint64(v1)
                                _ = v36
                            }
                        case 4294967295:
                            fmt.Println(fi25, v27)
                        }
                        _ = v30
                    } else {
                        v17 = 744
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
                _ = fi26
                _ = v27
            }
            if (((float64(v18) + ((((0.1 * 0.0) * 1e10) + 1e10) - ((((-2.25 - 0.1) - 1.5) / 3.14159) / (1.5 / 1.0)))) / 3.14159) >= 0.1) {
                switch (fi25 & 32768) {
                case 0:
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                default:
                    v37 := GenS4{F0: 1, F1: 1, F2: (v16 >> ((65535) & 63))}
                    _ = v37
                }
            }
            v38 := GenS5{F0: 174, F1: 65535, F2: v17}
            _ = v38
        } else {
            if (2 > uint(v16)) {
                var v39 uint8 = v16
                var v40 uint8 = v39
                _ = v39
                _ = v40
            }
            var v41 uint64 = genFn1(("" + "hello"), (v19 | (v19 % ((v19) | 1))), v1)
            for fi42 := 0; fi42 < 6; fi42++ {
                var v43 uint32 = v1
                if false {
                    fmt.Println(fi25)
                }
                if ((v19 % (((v19 - v19)) | 1)) != (v19 | (v19 / ((-32768) | 1)))) {
                    v44 := v4[("foo bar" + "foo bar")]
                    switch 32767 {
                    case 289:
                        v45 := GenS3{F0: -2147483648}
                        _ = v45
                    default:
                        v46 := GenS4{F0: 9223372036854775807, F1: -128, F2: (v16 >> ((142) & 63))}
                        _ = v46
                    }
                    _ = v44
                } else {
                    switch (fi25 / ((fi25) | 1)) {
                    case 497:
                        if true {
                            fmt.Println(fi25)
                            fmt.Println(v1, v19, v43)
                            v47 := GenS6{F0: 0}
                            var v48 uint16 = 145
                            _ = v47
                            _ = v48
                        } else {
                            v17 += (v41 << ((uint(v41)) & 63))
                            var v49 bool = (false && (!true))
                            var v50 uint = 2
                            _ = v49
                            _ = v50
                        }
                    case 32768:
                        v51 := GenS7{F0: 0, F1: 377}
                        _ = v51
                    case 32767:
                        if (((v1 + (v43 << ((1) & 63))) == (v1 + 0)) || true) {
                            switch (fi42 - fi42) {
                            case -9223372036854775808:
                                delete(v4, ("" + ("hello" + ("hello" + "a"))))
                            }
                        } else {
                            v2 = v18
                            v41 = v17
                            if (("a" + "") != "\n\t") {
                                v1 += v43
                                fmt.Println(v2, v18, v2)
                            }
                        }
                    }
                }
                _ = fi42
                _ = v43
            }
            _ = v41
        }
        _ = fi25
    }
    for fi52 := 0; fi52 < 1; fi52++ {
        v53 := GenS5{F0: 2, F1: uint16(v17), F2: 42}
        _ = fi52
        _ = v53
    }
    v54 := GenS8{F0: (!(!false))}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v54
}
