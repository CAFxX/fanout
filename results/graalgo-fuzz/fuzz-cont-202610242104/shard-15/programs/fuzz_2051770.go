package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 bool; F2 int8 }

func main() {
    var v1 uint32 = 42
    v2 := make(map[uint32]int)
    var v3 string = (("" + "a") + "")
    fmt.Println(v1)
    v2[228] = 127
    v3 = "hello"
    var v4 string = ("hello" + (("" + "hello") + v3))
    var v5 bool = true
    v5 = v5
    switch 793 {
    case 32768:
        var v6 uint16 = 42
        _ = v6
    case 127:
        v7 := GenS1{F0: 334, F1: 1, F2: 781}
        _ = v7
    case -2147483648:
        var v8 uint = 0
        _ = v8
    }
    if false {
        v5 = (0 == 1)
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 int = fi9
            v5 = (v5 && ((v1 & (v1 >> ((3) & 63))) >= v1))
            if (255 > 2) {
                var v11 int16 = 1000
                delete(v2, 2)
                _ = v11
            }
            _ = fi9
            _ = v10
        }
    } else {
        v1 = 42
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 uint64 = 0
            fmt.Println(v5, v5, v5)
            var v14 uint8 = 0
            _ = fi12
            _ = v13
            _ = v14
        }
        fmt.Println(v4)
    }
    switch v4 {
    case "foo bar":
        fmt.Println(v1)
    case "\n\t":
        if true {
            v15 := v2[v1]
            v16 := GenS2{F0: 160}
            fmt.Println(v1, v15, v15)
            _ = v15
            _ = v16
        }
    default:
        var v17 int = 741
        _ = v17
    }
    switch 32768 {
    case -1:
        v18 := GenS3{F0: 42, F1: 149}
        _ = v18
    case 42:
        var v19 bool = v5
        _ = v19
    case 4294967295:
        var v20 int16 = 1000
        _ = v20
    }
    if ((1 == 1) && v5) {
        for fi21 := 0; fi21 < 20; fi21++ {
            for fi22 := 0; fi22 < 20; fi22++ {
                v23 := []int{(fi22 >> ((42) & 63)), fi21, (fi22 - fi21)}
                if (v5 && (v1 >= v1)) {
                    if (0 == 2147483647) {
                        if false {
                            fmt.Println(v5, v4)
                            v24 := v2[(^(v1))]
                            _ = v24
                        } else {
                            var v25 int = (^(fi22))
                            var v26 uint8 = 64
                            if true {
                                v26 = v26
                                if (!(255 != 1)) {
                                    var v27 int = (v25 & (fi21 * (^(fi21))))
                                    v1 -= (v1 ^ v1)
                                    v28 := GenS4{F0: 756}
                                    _ = v27
                                    _ = v28
                                }
                            } else {
                                fmt.Println(v1, v5, v26)
                                v2[v1] = v25
                                switch v4 {
                                case "foo bar":
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                }
                            }
                            _ = v25
                            _ = v26
                        }
                        var v29 int = fi21
                        var v30 uint32 = 65535
                        _ = v29
                        _ = v30
                    }
                    switch "a" {
                    case "a":
                        v4 = "\n\t"
                        v4 = ("" + "hello")
                    case "hello":
                        fmt.Println(v1, fi21)
                    case "foo bar":
                        _ = len(v23)
                    default:
                        var v31 int64 = 0
                        _ = v31
                    }
                }
                if true {
                    var v32 uint64 = 2
                    v33 := GenS4{F0: 182}
                    if len(v23) > 0 {
                        v23[0] = 1
                    }
                    _ = v32
                    _ = v33
                }
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    }
    if (1e10 <= (-2.25 - (1.5 - float64(float32(1.5))))) {
        v34 := GenS2{F0: 2}
        if (-16 >= 127) {
            var v35 uint = 155
            _ = v35
        }
        _ = v34
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        v1 = (v1 / ((v1) | 1))
        v37 := []int{(fi36 >> ((18446744073709551615) & 63)), fi36, (fi36 + (fi36 / (((fi36 | fi36)) | 1)))}
        _ = fi36
        _ = v37
    }
    switch -256 {
    case -2147483648:
        switch -357 {
        case 4294967295:
            v38 := GenS5{F0: 127, F1: v5, F2: 100}
            _ = v38
        case -9223372036854775808:
            var v39 int = 42
            _ = v39
        case 1:
            v2[v1] = 32768
        }
    case -71:
        for fi40 := 0; fi40 < 17; fi40++ {
            fmt.Println(v1)
            if false {
                switch "foo bar" {
                case "foo bar":
                    for fi41 := 0; fi41 < 13; fi41++ {
                        var v42 uint64 = 42
                        v4 = ""
                        delete(v2, v1)
                        _ = fi41
                        _ = v42
                    }
                case "\n\t":
                    for fi43 := 0; fi43 < 10; fi43++ {
                        v1 = v1
                        v4 = "hello"
                        v1 -= (^(v1))
                        _ = fi43
                    }
                }
                for fi44 := 0; fi44 < 11; fi44++ {
                    fmt.Println(v3)
                    if (v5 || true) {
                        switch "a" {
                        case "hello":
                            var v45 int8 = 100
                            _ = v45
                        default:
                            v3 = (v4 + ("a" + v3))
                        }
                        var v46 int8 = 1
                        var v47 int16 = 0
                        _ = v46
                        _ = v47
                    } else {
                        switch fi44 {
                        case -887:
                            v3 = v4
                        case 136:
                            fmt.Println(fi44, v3, v3)
                        case 32767:
                            v4 = "a"
                        default:
                            v1 += 1
                        }
                        var v48 uint = 522
                        _ = v48
                    }
                    if true {
                        var v49 int8 = 100
                        _ = v49
                    } else {
                        v50 := GenS1{F0: 558, F1: 0, F2: -478}
                        v51 := GenS5{F0: (fi44 - fi40), F1: v5, F2: 1}
                        _ = v50
                        _ = v51
                    }
                    _ = fi44
                }
            }
            for fi52 := 0; fi52 < 20; fi52++ {
                var v53 int16 = 42
                _ = fi52
                _ = v53
            }
            _ = fi40
        }
    default:
        v2[(v1 | 116)] = 127
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
