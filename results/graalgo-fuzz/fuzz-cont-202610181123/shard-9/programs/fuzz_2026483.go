package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint32; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string }
func genFn0(p0 int8, p1 uint8, p2 int8) int {
    if true {
        var v6 int64 = 1
        switch (("foo bar" + "hello") + "hello") {
        case "a":
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := []int{(fi7 - fi7)}
                _ = fi7
                _ = v8
            }
        case "":
            switch "a" {
            case "foo bar":
                v6 = v6
            }
        case "hello":
            for fi9 := 0; fi9 < 6; fi9++ {
                switch 1 {
                case -2147483648:
                    for fi10 := 0; fi10 < 14; fi10++ {
                        v6 -= (v6 << ((65535) & 63))
                        v6 += v6
                        p0 = 127
                        v11 := []int{int(v6), -923, fi9, (fi9 % (((fi10 + 128)) | 1))}
                        _ = fi10
                        _ = v11
                    }
                case 42:
                    v12 := []int{1, fi9}
                    _ = v12
                }
                if ((!true) && false) {
                    v13 := GenS1{F0: 42, F1: 42}
                    var v14 uint8 = p1
                    _ = v13
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 16; fi15++ {
                        var v16 int64 = (v6 * 42)
                        v17 := GenS2{F0: 4294967295, F1: 65535, F2: 2147483647}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    v18 := GenS1{F0: 57, F1: p2}
                    p1 = p1
                    _ = v18
                }
                _ = fi9
            }
        }
        _ = v6
    } else {
        if (255 <= 18446744073709551615) {
            for fi19 := 0; fi19 < 4; fi19++ {
                if false {
                    var v20 bool = false
                    var v21 int16 = -32768
                    _ = v20
                    _ = v21
                } else {
                    if (32767 > 32767) {
                        var v22 float32 = (float32((0.1 - (1e10 * 1.5))) - 0.0)
                        _ = v22
                    }
                    v23 := []int{fi19, fi19}
                    _ = v23
                }
                v24 := GenS3{F0: (-(p2)), F1: 2147483647, F2: "foo bar"}
                _ = fi19
                _ = v24
            }
        }
        switch "a" {
        case "hello":
            for fi25 := 0; fi25 < 10; fi25++ {
                if false {
                    v26 := GenS3{F0: p2, F1: 869, F2: (("a" + "\n\t") + "a")}
                    v27 := []int{(fi25 - (fi25 | fi25)), (fi25 ^ fi25), int(p1), fi25}
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS4{F0: 255}
                    _ = v28
                }
                v29 := GenS5{F0: p2, F1: 0}
                if (((("\n\t" + (("" + "hello") + "foo bar")) == "\n\t") && true) && ((195 > 32767) && (0.0 < float64(1.5)))) {
                    v30 := []int{(fi25 - 346), (fi25 - fi25), fi25, -2147483648}
                    if (3.14159 < 1.5) {
                        fmt.Println(p2, fi25, p1)
                        fmt.Println(p2, p2)
                    } else {
                        fmt.Println(fi25)
                        if (252 > p1) {
                            var v31 bool = (-2147483648 > 923)
                            p2 = (p2 >> ((65535) & 63))
                            for fi32 := 0; fi32 < 13; fi32++ {
                                if (!(0.1 > (float32(3.14159) - (3.14159 / 3.14159)))) {
                                    p1 += 255
                                }
                                v33 := GenS6{F0: -420}
                                _ = fi32
                                _ = v33
                            }
                            _ = v31
                        } else {
                            var v34 uint = 42
                            _ = v34
                        }
                        var v35 string = "foo bar"
                        _ = v35
                    }
                    _ = v30
                }
                _ = fi25
                _ = v29
            }
        case "\n\t":
            for fi36 := 0; fi36 < 15; fi36++ {
                var v37 int8 = (p2 % ((p0) | 1))
                if false {
                    p2 = v37
                } else {
                    p1 = 128
                    for fi38 := 0; fi38 < 8; fi38++ {
                        v39 := []int{128, (^(fi38))}
                        _ = fi38
                        _ = v39
                    }
                    var v40 int64 = 42
                    _ = v40
                }
                _ = fi36
                _ = v37
            }
        }
        var v41 string = ("\n\t" + "\n\t")
        _ = v41
    }
    var v42 uint8 = p1
    var v43 uint32 = 65535
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v42
    _ = v43
    return 42
}

func main() {
    var v1 bool = false
    var v2 float64 = (float64(1e10) * float64(-2.25))
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    v5 := []int{-2147483648}
    v2 = v2
    v44 := GenS7{F0: -4}
    var v45 int = 32767
    var v46 int8 = -7
    v46 += (v46 % ((v46) | 1))
    fmt.Println(v1)
    for fi47 := 0; fi47 < 5; fi47++ {
        v46 = -1
        for fi48 := 0; fi48 < 14; fi48++ {
            fmt.Println(fi48)
            if v1 {
                if ((1 < 1) || false) {
                    v4[907] = "hello"
                    var v49 uint8 = 42
                    switch v45 {
                    case 645:
                        var v50 uint = 255
                        _ = v50
                    case -760:
                        if (false && (826 < 42)) {
                            if (!false) {
                                v51 := GenS6{F0: 32767}
                                v52 := GenS5{F0: 100, F1: -1}
                                _ = v51
                                _ = v52
                            }
                        } else {
                            var v53 int16 = 42
                            var v54 int = (fi47 % (((fi48 ^ (-(fi47)))) | 1))
                            delete(v3, "")
                            _ = v53
                            _ = v54
                        }
                    }
                    if ((2147483647 <= 0) || false) {
                        switch "foo bar" {
                        case "":
                            switch fi48 {
                            case -679:
                                v46 += v46
                            }
                        }
                        var v55 uint32 = 474
                        switch (^(fi48)) {
                        case 236:
                            if len(v5) > 0 {
                                v5[0] = (fi48 * fi47)
                            }
                            v55 += v55
                        default:
                            v56 := GenS7{F0: 1}
                            _ = v56
                        }
                        _ = v55
                    } else {
                        v57 := GenS1{F0: 2, F1: 127}
                        _ = v57
                    }
                    fmt.Println(v2)
                    _ = v49
                }
                v45 += -9223372036854775808
                var v58 uint64 = 42
                _ = v58
            }
            if v1 {
                v59 := GenS7{F0: 665}
                _ = v59
            } else {
                var v60 float32 = (1.5 / 1e10)
                fmt.Println(fi48, fi47, v1)
                v61 := GenS8{F0: ("hello" + "hello")}
                _ = v60
                _ = v61
            }
            _ = fi48
        }
        _ = fi47
    }
    var v62 int = (v45 << ((255) & 63))
    fmt.Println(v62, v45, v45)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v62
}
