package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 uint; F2 uint }
type GenS4 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint64; F1 uint; F2 int16 }
func genFn0(p0 uint8, p1 int64) uint {
    if true {
        v7 := GenS1{F0: 1}
        v8 := []int{42, -9223372036854775808, 0, 2147483647}
        _ = v7
        _ = v8
    } else {
        if (!(!true)) {
            if ((1e10 > 0.1) || true) {
                if (127 <= 92) {
                    v9 := []int{130, 127, 32767, 9223372036854775807}
                    var v10 uint32 = 65535
                    _ = v10
                    _ = v9
                } else {
                    v11 := []int{-9223372036854775808, 9223372036854775807, -1, 32767}
                    switch "foo bar" {
                    case "hello":
                        _ = len(v11)
                    default:
                        for fi12 := 0; fi12 < 16; fi12++ {
                            p1 -= p1
                            _ = fi12
                        }
                    }
                    switch "foo bar" {
                    case "foo bar":
                        if true {
                            var v13 string = "hello"
                            if true {
                                p0 += p0
                                var v14 uint32 = 2147483648
                                _ = v14
                            }
                            _ = v13
                        } else {
                            for fi15 := 0; fi15 < 1; fi15++ {
                                v11 = append(v11, (fi15 / (((-(fi15))) | 1)))
                                v11 = append(v11, (fi15 | fi15))
                                _ = fi15
                            }
                            p0 = 0
                        }
                    case "a":
                        p0 -= p0
                    }
                    _ = v11
                }
                for fi16 := 0; fi16 < 13; fi16++ {
                    for fi17 := 0; fi17 < 16; fi17++ {
                        v18 := []int{fi17, 32768, 9223372036854775807, (fi16 - (fi16 + (fi17 * (fi16 / (((fi16 & fi17)) | 1)))))}
                        _ = fi17
                        _ = v18
                    }
                    v19 := []int{(fi16 % ((2147483647) | 1)), fi16, (fi16 & (fi16 << ((386) & 63)))}
                    _ = fi16
                    _ = v19
                }
                p0 += 152
            }
            p1 += p1
            v20 := []int{547, 401}
            _ = v20
        } else {
            var v21 string = ("\n\t" + ("hello" + ("foo bar" + "")))
            var v22 int32 = 2147483647
            var v23 uint32 = 2
            _ = v21
            _ = v22
            _ = v23
        }
    }
    if true {
        for fi24 := 0; fi24 < 13; fi24++ {
            for fi25 := 0; fi25 < 5; fi25++ {
                p1 = (p1 ^ (p1 * p1))
                _ = fi25
            }
            v26 := GenS1{F0: 1}
            _ = fi24
            _ = v26
        }
        var v27 uint8 = (p0 << ((2) & 63))
        v28 := GenS2{F0: 1}
        _ = v27
        _ = v28
    } else {
        switch 259 {
        case 0:
            v29 := []int{}
            _ = v29
        case -318:
            for fi30 := 0; fi30 < 3; fi30++ {
                var v31 int16 = 42
                _ = fi30
                _ = v31
            }
        case -9223372036854775808:
            v32 := []int{714, 32768, -9223372036854775808}
            _ = v32
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 float32 = 1.5
    var v2 uint8 = 127
    v3 := make(map[string]uint32)
    var v4 uint32 = 65535
    var v5 uint32 = v4
    v6 := make(map[string]int)
    if (!false) {
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 float32 = (v1 - v1)
            if true {
                switch fi33 {
                case 200:
                    v1 += v1
                case 127:
                    var v35 int = 4294967295
                    _ = v35
                default:
                    switch (fi33 << ((553) & 63)) {
                    case 0:
                        v4 += 4294967295
                    }
                }
                if ((v2 < v2) && (856 <= 1000)) {
                    v36 := GenS3{F0: "foo bar", F1: 253, F2: genFn0(127, 9223372036854775807)}
                    _ = v36
                } else {
                    v37 := []int{(fi33 >> ((2) & 63)), fi33, 469}
                    _ = v37
                }
            } else {
                v2 = (v2 | 0)
            }
            if (1 < 42) {
                var v38 int32 = 1
                var v39 int8 = 1
                v40 := GenS1{F0: v38}
                _ = v38
                _ = v39
                _ = v40
            } else {
                var v41 uint32 = 2
                _ = v41
            }
            _ = fi33
            _ = v34
        }
    }
    var v42 uint8 = uint8(v5)
    fmt.Println(v4, v42, v1)
    v3[("" + "")] = (v5 << ((1) & 63))
    var v43 string = ("hello" + ("\n\t" + ("foo bar" + "hello")))
    v2 -= 39
    if false {
        _ = len(v3)
    }
    v44 := GenS2{F0: 353}
    switch v43 {
    case "a":
        var v45 uint = 255
        _ = v45
    case "":
        v46 := []int{}
        _ = v46
    case "\n\t":
        v47 := []int{2147483647}
        _ = v47
    }
    var v48 uint8 = (v42 * (v2 >> ((469) & 63)))
    if false {
        v2 = v2
        if (!true) {
            delete(v3, v43)
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v4, v43, v5)
            case "a":
                var v49 bool = true
                _ = v49
            case "":
                var v50 uint = 58
                _ = v50
            }
        } else {
            v51 := GenS4{F0: 1, F1: (v5 & 4294967295), F2: 4294967295}
            _ = v51
        }
        for fi52 := 0; fi52 < 14; fi52++ {
            v4 -= (v4 / ((1) | 1))
            v4 = 202
            _ = fi52
        }
    }
    var v53 int = 1
    switch (v53 % (((v53 | v53)) | 1)) {
    case -9223372036854775808:
        if (true || ("\n\t" < (v43 + "a"))) {
            var v54 uint16 = 32767
            var v55 uint64 = 4294967295
            _ = v54
            _ = v55
        }
    case 128:
        if (100 >= 42) {
            var v56 int = (v53 ^ (v53 >> ((18446744073709551615) & 63)))
            var v57 uint16 = 255
            _ = v56
            _ = v57
        }
    default:
        for fi58 := 0; fi58 < 14; fi58++ {
            var v59 uint32 = 245
            _ = fi58
            _ = v59
        }
    }
    for fi60 := 0; fi60 < 10; fi60++ {
        switch 42 {
        case 127:
            v4 = v5
        case -806:
            if true {
                var v61 int64 = 569
                _ = v61
            } else {
                var v62 int8 = 127
                _ = v62
            }
        default:
            switch v53 {
            case 2147483647:
                var v63 int16 = 452
                _ = v63
            case 873:
                v64 := GenS2{F0: 42}
                _ = v64
            }
        }
        v3[(v43 + "a")] = v4
        for fi65 := 0; fi65 < 2; fi65++ {
            if ("hello" != v43) {
                var v66 uint8 = 255
                _ = v66
            } else {
                if false {
                    v67 := GenS2{F0: 646}
                    var v68 int16 = 1
                    v69 := GenS5{F0: 85, F1: 969, F2: v68}
                    v68 = v68
                    if ((982 < 221) && true) {
                        switch 32767 {
                        case 1:
                            var v70 uint16 = 247
                            var v71 uint32 = (v4 / ((v5) | 1))
                            var v72 int32 = -813
                            _ = v70
                            _ = v71
                            _ = v72
                        default:
                            var v73 uint16 = 42
                            _ = v73
                        }
                        if (!(!(!false))) {
                            v43 = "hello"
                        }
                    } else {
                        var v74 int64 = 42
                        var v75 uint64 = 2
                        fmt.Println(fi65)
                        _ = v74
                        _ = v75
                    }
                    _ = v67
                    _ = v68
                    _ = v69
                }
            }
            var v76 float64 = 0.1
            v3[("\n\t" + v43)] = 207
            _ = fi65
            _ = v76
        }
        _ = fi60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v48
    _ = v5
    _ = v53
    _ = v6
}
