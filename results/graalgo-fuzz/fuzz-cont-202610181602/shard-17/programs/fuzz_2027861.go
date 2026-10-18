package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 string; F2 uint }
type GenS6 struct { F0 int16; F1 string; F2 string }
func genFn0(p0 uint64, p1 uint8) int {
    fmt.Println(p1, p0)
    v6 := []int{14, 0}
    _ = p0
    _ = p1
    _ = v6
    return 127
}
func genFn1(p0 uint16, p1 bool) int16 {
    switch 123 {
    case 2147483647:
        p0 = (p0 & (p0 / ((p0) | 1)))
    case 127:
        p0 = p0
    case 42:
        if true {
            var v7 int8 = 100
            _ = v7
        }
    default:
        v8 := GenS1{F0: 4294967295}
        _ = v8
    }
    if false {
        fmt.Println(p0)
    }
    v9 := GenS2{F0: 100, F1: 0}
    _ = p0
    _ = p1
    _ = v9
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 string = ("foo bar" + "\n\t")
    v3 := make(map[int]int)
    var v4 int8 = 1
    v5 := []string{"", v2}
    for fi10 := 0; fi10 < 14; fi10++ {
        if (true || ((^(fi10)) > (^(fi10)))) {
            if ((432 > 372) && ((!false) && false)) {
                for fi11 := 0; fi11 < 3; fi11++ {
                    v4 = v4
                    _ = fi11
                }
            }
            if (!("" <= (("hello" + "foo bar") + v2))) {
                var v12 uint16 = 255
                if false {
                    fmt.Println(v4, v12)
                } else {
                    v13 := GenS3{F0: v12, F1: -1}
                    _ = v13
                }
                _ = v12
            }
            v14 := v1[(fi10 >> ((871) & 63))]
            _ = v14
        } else {
            for fi15 := 0; fi15 < 4; fi15++ {
                switch "a" {
                case "foo bar":
                    var v16 int = 32768
                    _ = v16
                case "a":
                    switch 9223372036854775807 {
                    case 4294967295:
                        fmt.Println(fi10, fi15, v2)
                    default:
                        if (!((v4 / ((v4) | 1)) > (v4 % ((v4) | 1)))) {
                            var v17 int16 = 0
                            _ = v17
                        }
                    }
                }
                delete(v1, 310)
                var v18 int8 = (v4 + 42)
                _ = fi15
                _ = v18
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                _ = len(v1)
                _ = fi19
            }
            v2 = v2
        }
        _ = fi10
    }
    var v20 float32 = -2.25
    v20 = (0.0 * float32((0.1 - ((float64(float32(float64(v20))) / 1e10) * ((1.5 + (1e10 - 1e10)) - (1.5 - (0.0 * 1.5)))))))
    for fi21 := 0; fi21 < 6; fi21++ {
        v4 -= (v4 % (((v4 | (v4 * 1))) | 1))
        _ = fi21
    }
    v4 += v4
    for fi22 := 0; fi22 < 10; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            v20 -= v20
            v1[128] = (v2 + v2)
            _ = fi23
        }
        _ = fi22
    }
    if false {
        if (!true) {
            var v24 uint64 = 2
            v20 += (1e10 + float32(-2.25))
            _ = v24
        }
        if (-244 > 9223372036854775807) {
            for fi25 := 0; fi25 < 8; fi25++ {
                if (255 <= 84) {
                    delete(v3, (-(fi25)))
                    switch v2 {
                    case "\n\t":
                        v26 := GenS1{F0: 4294967295}
                        _ = v26
                    case "":
                        v27 := GenS1{F0: 2}
                        _ = v27
                    case "a":
                        if (-1 <= 1000) {
                            var v28 uint64 = 18446744073709551615
                            for fi29 := 0; fi29 < 3; fi29++ {
                                v20 -= (((v20 * 0.1) * 3.14159) - v20)
                                var v30 int16 = -16
                                _ = fi29
                                _ = v30
                            }
                            _ = v28
                        } else {
                            for fi31 := 0; fi31 < 3; fi31++ {
                                v1[fi25] = "foo bar"
                                v20 += v20
                                var v32 bool = false
                                _ = fi31
                                _ = v32
                            }
                            var v33 string = "a"
                            _ = v33
                        }
                    }
                }
                fmt.Println(v2, v20, fi25)
                v34 := GenS3{F0: 1, F1: 42}
                _ = fi25
                _ = v34
            }
        } else {
            v20 -= v20
            fmt.Println(v2, v2, v2)
            var v35 uint64 = 2
            _ = v35
        }
    }
    switch 1 {
    case 127:
        if (!true) {
            v20 += v20
            v36 := GenS4{F0: 0}
            _ = v36
        }
    case -6:
        v37 := GenS5{F0: 1, F1: v2, F2: 282}
        _ = v37
    case 32767:
        fmt.Println(v20, v4)
    default:
        if true {
            var v38 uint64 = 42
            _ = v38
        } else {
            switch "a" {
            case "a":
                if (1 < 185) {
                    delete(v1, 618)
                    v39 := v3[-1]
                    _ = v39
                }
            case "\n\t":
                var v40 uint16 = 255
                _ = v40
            }
            v3[128] = -1
            for fi41 := 0; fi41 < 15; fi41++ {
                v4 = v4
                fmt.Println(v2, v4)
                for fi42 := 0; fi42 < 5; fi42++ {
                    if (v2 != ("hello" + "\n\t")) {
                        switch ((v2 + (v2 + "")) + v2) {
                        case "a":
                            var v43 int32 = -391
                            _ = v43
                        case "foo bar":
                            delete(v1, (fi41 / ((fi41) | 1)))
                        case "hello":
                            fmt.Println(v4, v4, v4)
                        }
                        v44 := GenS2{F0: int8(fi42), F1: 1000}
                        _ = v44
                    }
                    _ = fi42
                }
                _ = fi41
            }
        }
    }
    fmt.Println(v2, v4)
    switch -460 {
    case 890:
        if true {
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                var v45 uint = 817
                _ = v45
            case "":
                _ = len(v3)
            case "foo bar":
                v2 = (v2 + "a")
            default:
                v20 = 1e10
            }
            if false {
                v2 = (("" + v2) + ("\n\t" + (("\n\t" + (v2 + v2)) + v2)))
            }
            for fi46 := 0; fi46 < 1; fi46++ {
                if (false && (42 <= 536)) {
                    for fi47 := 0; fi47 < 19; fi47++ {
                        v48 := GenS5{F0: 1, F1: v2, F2: 784}
                        _ = fi47
                        _ = v48
                    }
                }
                var v49 string = v2
                var v50 int32 = 105
                _ = fi46
                _ = v49
                _ = v50
            }
        }
    case 627:
        if (true && false) {
            if ("foo bar" < v2) {
                _ = len(v1)
                var v51 uint16 = 129
                _ = v51
            }
        } else {
            v2 = (v2 + v2)
            fmt.Println(v2, v2)
            delete(v3, 2147483647)
        }
    case -1:
        for fi52 := 0; fi52 < 8; fi52++ {
            _ = len(v3)
            fmt.Println(fi52, v4)
            _ = fi52
        }
    }
    for fi53 := 0; fi53 < 1; fi53++ {
        switch "" {
        case "\n\t":
            var v54 int = (^(fi53))
            _ = v54
        case "foo bar":
            if (true && (2 != 204)) {
                if true {
                    var v55 string = (v2 + v2)
                    var v56 uint64 = 0
                    _ = v55
                    _ = v56
                }
            }
        case "hello":
            v57 := GenS4{F0: 0}
            _ = v57
        default:
            if (151 < 213) {
                v58 := v1[fi53]
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v58
            } else {
                var v59 uint32 = 1
                fmt.Println(v20, fi53, fi53)
                _ = v59
            }
        }
        if (-123 > 1) {
            if len(v5) > 0 {
                v5[0] = ((("hello" + (((v2 + v2) + "") + "")) + (("" + "") + "foo bar")) + v2)
            }
        } else {
            for fi60 := 0; fi60 < 20; fi60++ {
                fmt.Println(v20, fi53, fi53)
                fmt.Println(v20)
                fmt.Println(v4, fi60)
                v20 -= (v20 + -2.25)
                v61 := GenS1{F0: 18446744073709551615}
                _ = fi60
                _ = v61
            }
            switch fi53 {
            case -805:
                var v62 int64 = 42
                _ = v62
            }
            for fi63 := 0; fi63 < 9; fi63++ {
                switch 181 {
                case 0:
                    v64 := GenS4{F0: 1}
                    _ = v64
                }
                v4 -= 122
                var v65 string = v2
                var v66 int64 = -502
                var v67 string = (v65 + "foo bar")
                var v68 bool = (1 >= 504)
                _ = fi63
                _ = v65
                _ = v66
                _ = v67
                _ = v68
            }
        }
        if (2 != 558) {
            fmt.Println(v20, v4, fi53)
            var v69 int8 = v4
            _ = v69
        } else {
            v70 := GenS6{F0: 537, F1: "foo bar", F2: ""}
            _ = v70
        }
        _ = fi53
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
