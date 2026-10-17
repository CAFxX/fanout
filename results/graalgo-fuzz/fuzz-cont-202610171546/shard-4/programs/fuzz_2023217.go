package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 int16 }
func genFn0(p0 bool, p1 int) int16 {
    v4 := GenS1{F0: 85, F1: (1 != -2147483648), F2: (p1 % (((p1 | p1)) | 1))}
    v5 := GenS1{F0: 980, F1: (false && p0), F2: 32768}
    switch "hello" {
    case "":
        switch "\n\t" {
        case "foo bar":
            fmt.Println(p1, p1)
        default:
            switch -1 {
            case 673:
                var v6 float64 = (0.0 / 1e10)
                _ = v6
            case 127:
                var v7 uint64 = 2
                _ = v7
            case 209:
                fmt.Println(p1, p0)
            }
        }
    case "a":
        v8 := GenS2{F0: 2, F1: 4294967295}
        _ = v8
    case "foo bar":
        if p0 {
            for fi9 := 0; fi9 < 20; fi9++ {
                for fi10 := 0; fi10 < 10; fi10++ {
                    if (141 <= 2) {
                        v11 := []int{(fi9 & -1), (fi10 - p1), (^(fi9)), 0}
                        p1 -= -9223372036854775808
                        _ = v11
                    }
                    p0 = (p0 && ("hello" != "\n\t"))
                    var v12 int = (p1 & fi9)
                    var v13 string = ("hello" + "foo bar")
                    v14 := GenS1{F0: 4294967295, F1: ((9223372036854775807 >= 0) && p0), F2: (^(v12))}
                    _ = fi10
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi9
            }
        } else {
            var v15 int32 = -322
            var v16 string = "\n\t"
            _ = v15
            _ = v16
        }
    default:
        if true {
            var v17 int8 = -59
            _ = v17
        } else {
            var v18 int = -90
            for fi19 := 0; fi19 < 12; fi19++ {
                v20 := []int{(fi19 / ((42) | 1))}
                _ = fi19
                _ = v20
            }
            for fi21 := 0; fi21 < 18; fi21++ {
                switch v18 {
                case 128:
                    p0 = ("hello" == ("foo bar" + "\n\t"))
                }
                if p0 {
                    fmt.Println(p1, p1, fi21)
                    v22 := GenS2{F0: 2, F1: 0}
                    switch (v18 << ((2) & 63)) {
                    case 9223372036854775807:
                        var v23 int8 = -128
                        _ = v23
                    case -980:
                        for fi24 := 0; fi24 < 13; fi24++ {
                            if (42 <= 40) {
                                var v25 int32 = 0
                                v26 := []int{v18, (fi24 << ((801) & 63))}
                                if ((!true) || false) {
                                    var v27 int8 = -123
                                    v18 = v18
                                    _ = len(v26)
                                    if p0 {
                                        switch 42 {
                                        case 9223372036854775807:
                                            switch (-(fi21)) {
                                            case 81:
                                                var v28 string = ""
                                                _ = v28
                                            case 32768:
                                                fmt.Println(fi24)
                                            case 127:
                                                v29 := GenS1{F0: 2, F1: false, F2: (^(p1))}
                                                _ = v29
                                            }
                                        case 128:
                                            if false {
                                                if len(v26) > 0 {
                                                    v26[0] = (^(fi21))
                                                }
                                                fmt.Println(v27)
                                                if true {
                                                    fmt.Println(p0, fi21, p1)
                                                    p0 = (28 != 255)
                                                    fmt.Println(v27, v18, v25)
                                                } else {
                                                    var v30 uint8 = 127
                                                    _ = v30
                                                }
                                            }
                                        case -9223372036854775808:
                                            var v31 int = (fi24 | (v18 * v18))
                                            _ = v31
                                        }
                                    }
                                    _ = v27
                                }
                                _ = v25
                                _ = v26
                            }
                            p0 = (!p0)
                            _ = fi24
                        }
                    }
                    _ = v22
                }
                _ = fi21
            }
            _ = v18
        }
    }
    if ((("foo bar" >= "hello") && p0) && (false || true)) {
        for fi32 := 0; fi32 < 13; fi32++ {
            p0 = (42 >= -209)
            _ = fi32
        }
        v33 := []int{(p1 / ((p1) | 1)), 128}
        if p0 {
            p1 = 128
            v34 := GenS3{F0: 42}
            var v35 uint32 = 65535
            _ = v34
            _ = v35
        } else {
            if (2 > 32767) {
                p0 = p0
                var v36 uint = 2
                _ = v36
            }
        }
        _ = v33
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 822
}

func main() {
    v1 := []int{42}
    var v2 int64 = 42
    var v3 uint8 = 84
    v2 += 9223372036854775807
    for fi37 := 0; fi37 < 5; fi37++ {
        fmt.Println(v3, v2)
        for fi38 := 0; fi38 < 11; fi38++ {
            fmt.Println(fi37, fi37, fi37)
            fmt.Println(v3, fi37, v2)
            _ = fi38
        }
        _ = fi37
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        v40 := GenS4{F0: 178}
        v3 = (v3 + (^(v3)))
        var v41 uint = 255
        _ = fi39
        _ = v40
        _ = v41
    }
    switch 1 {
    case -1:
        v42 := GenS3{F0: 1}
        _ = v42
    case 32767:
        for fi43 := 0; fi43 < 5; fi43++ {
            var v44 uint64 = 4294967295
            _ = fi43
            _ = v44
        }
    case 42:
        if len(v1) > 0 {
            v1[0] = 128
        }
    }
    v45 := GenS3{F0: 42}
    if false {
        v2 = (v2 >> ((354) & 63))
        if true {
            if (0.0 == (((0.1 - 3.14159) / -2.25) / 3.14159)) {
                var v46 uint = 1
                var v47 uint64 = 0
                _ = v46
                _ = v47
            }
            if ((^(v3)) > v3) {
                var v48 float64 = 1e10
                if false {
                    v1 = append(v1, -601)
                } else {
                    fmt.Println(v48, v48, v48)
                    v49 := GenS3{F0: 4294967295}
                    _ = v49
                }
                var v50 uint8 = v3
                _ = v48
                _ = v50
            } else {
                fmt.Println(v2, v2)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (65535 != 65535) {
                    v3 = 128
                    var v51 float32 = 3.14159
                    _ = v51
                }
            }
        } else {
            var v52 int16 = 1000
            _ = v52
        }
        if (127 == int8(v3)) {
            for fi53 := 0; fi53 < 6; fi53++ {
                fmt.Println(v2, fi53, v3)
                for fi54 := 0; fi54 < 7; fi54++ {
                    _ = len(v1)
                    var v55 int8 = -74
                    fmt.Println(fi54)
                    _ = fi54
                    _ = v55
                }
                _ = fi53
            }
            for fi56 := 0; fi56 < 10; fi56++ {
                var v57 bool = false
                for fi58 := 0; fi58 < 3; fi58++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi58
                }
                if true {
                    switch (fi56 << ((436) & 63)) {
                    case 9223372036854775807:
                        switch "hello" {
                        case "foo bar":
                            v3 = (^(v3))
                        case "hello":
                            v59 := GenS5{F0: (fi56 >> ((336) & 63)), F1: 32767}
                            _ = v59
                        }
                    case 42:
                        v60 := GenS5{F0: (fi56 - (fi56 | (-(fi56)))), F1: -32768}
                        _ = v60
                    }
                }
                _ = fi56
                _ = v57
            }
        } else {
            if len(v1) > 0 {
                v1[0] = 128
            }
            var v61 int32 = 2147483647
            _ = v61
        }
    } else {
        v3 = 0
        v62 := GenS3{F0: 102}
        var v63 int8 = 100
        _ = v62
        _ = v63
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi64 := 0; fi64 < 5; fi64++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi65 := 0; fi65 < 1; fi65++ {
            v2 -= -9223372036854775808
            v3 -= (v3 / (((v3 + 2)) | 1))
            _ = fi65
        }
        _ = fi64
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
}
