package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int32; F1 int; F2 bool }
type GenS6 struct { F0 uint8; F1 string; F2 int32 }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 uint64; F1 uint64; F2 uint16 }
func genFn0() int32 {
    v6 := []int{32767}
    if true {
        v6 = append(v6, -1)
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 uint = 65535
            _ = fi7
            _ = v8
        }
        var v9 int = 763
        _ = v9
    }
    var v10 uint32 = 464
    var v11 int8 = -128
    v11 += v11
    _ = v10
    _ = v11
    _ = v6
    return 457
}
func genFn1(p0 int) string {
    v12 := GenS1{F0: 1, F1: 42}
    if true {
        if (32767 == 0) {
            fmt.Println(p0)
        } else {
            if false {
                var v13 uint64 = 42
                fmt.Println(p0)
                var v14 bool = (-2.25 > 3.14159)
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                v16 := GenS1{F0: -42, F1: 42}
                _ = fi15
                _ = v16
            }
            v17 := []int{p0, 128, p0, p0}
            _ = v17
        }
        switch "foo bar" {
        case "\n\t":
            var v18 uint64 = 904
            _ = v18
        case "":
            if (!(127 > 95)) {
                v19 := GenS2{F0: 42, F1: 2, F2: 4294967295}
                var v20 string = ("\n\t" + "")
                _ = v19
                _ = v20
            }
        default:
            v21 := GenS2{F0: 128, F1: 18446744073709551615, F2: 42}
            _ = v21
        }
    }
    v22 := []int{1, (p0 << ((391) & 63)), (^(p0))}
    if false {
        fmt.Println(p0)
    } else {
        if false {
            p0 = p0
        }
        p0 = p0
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        if false {
            var v24 string = ""
            var v25 int64 = -9223372036854775808
            v25 = (v25 ^ v25)
            _ = v24
            _ = v25
        } else {
            switch (fi23 + (-(p0))) {
            case -2147483648:
                v26 := GenS1{F0: 109, F1: -1}
                _ = v26
            case 42:
                var v27 float32 = float32(-2.25)
                _ = v27
            default:
                var v28 uint64 = 881
                _ = v28
            }
            p0 = 1
            for fi29 := 0; fi29 < 15; fi29++ {
                var v30 int64 = -9223372036854775808
                _ = fi29
                _ = v30
            }
        }
        p0 += p0
        if (946 < 705) {
            var v31 float32 = (3.14159 - 1.5)
            _ = v31
        } else {
            if false {
                var v32 float32 = 0.0
                v33 := GenS1{F0: 66, F1: -1}
                var v34 uint64 = 0
                _ = v32
                _ = v33
                _ = v34
            }
            switch ("\n\t" + ("" + ("" + "foo bar"))) {
            case "":
                p0 -= (p0 >> ((65535) & 63))
            case "foo bar":
                for fi35 := 0; fi35 < 3; fi35++ {
                    v36 := GenS3{F0: 1, F1: 2, F2: 18446744073709551615}
                    var v37 float32 = (0.0 - 0.0)
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            }
            p0 = -2147483648
        }
        _ = fi23
    }
    _ = p0
    _ = v12
    _ = v22
    return "a"
}

func main() {
    var v1 float32 = (-2.25 / (float32(0.1) * ((0.0 + 1e10) - (3.14159 + 1e10))))
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    v4 := make(map[string]string)
    var v5 uint = 65535
    switch "" {
    case "foo bar":
        v5 = 166
    case "":
        var v38 int = 32768
        _ = v38
    default:
        fmt.Println(v5, v5)
    }
    v39 := GenS1{F0: -5, F1: -9223372036854775808}
    var v40 bool = true
    switch 555 {
    case 42:
        var v41 float32 = v1
        _ = v41
    case 128:
        switch 2147483647 {
        case 9223372036854775807:
            v42 := GenS4{F0: 42, F1: ("hello" + "a")}
            _ = v42
        case -91:
            for fi43 := 0; fi43 < 17; fi43++ {
                v44 := GenS2{F0: 2, F1: 4294967295, F2: 179}
                v45 := GenS5{F0: -423, F1: 127, F2: (!false)}
                if false {
                    if false {
                        for fi46 := 0; fi46 < 3; fi46++ {
                            var v47 uint = v5
                            _ = fi46
                            _ = v47
                        }
                    }
                    var v48 uint32 = 65535
                    v1 += (3.14159 * (v1 * ((v1 - ((v1 - v1) + v1)) + float32((1e10 / 1e10)))))
                    _ = v48
                }
                _ = fi43
                _ = v44
                _ = v45
            }
        case -453:
            v40 = (!false)
        }
    case 110:
        var v49 uint32 = 2147483648
        _ = v49
    }
    fmt.Println(v1, v1, v5)
    v1 -= float32(-2.25)
    v50 := GenS3{F0: 4294967295, F1: 255, F2: 602}
    for fi51 := 0; fi51 < 8; fi51++ {
        if (!v40) {
            for fi52 := 0; fi52 < 13; fi52++ {
                v53 := GenS6{F0: 2, F1: "a", F2: -513}
                v4[(("\n\t" + "a") + (("foo bar" + "\n\t") + ("a" + "a")))] = "foo bar"
                v54 := v3[95]
                _ = fi52
                _ = v53
                _ = v54
            }
            for fi55 := 0; fi55 < 16; fi55++ {
                var v56 uint32 = 495
                v1 += v1
                _ = fi55
                _ = v56
            }
            fmt.Println(v40, v40, v40)
        }
        for fi57 := 0; fi57 < 3; fi57++ {
            if (1 >= 164) {
                fmt.Println(fi51, v1)
                switch 128 {
                case 9223372036854775807:
                    v58 := []int{(fi51 | (fi57 * 2147483647))}
                    _ = v58
                case 42:
                    v59 := GenS7{F0: ("a" + "hello"), F1: 949}
                    _ = v59
                case 32768:
                    v60 := GenS8{F0: 0, F1: 42, F2: 255}
                    _ = v60
                default:
                    var v61 uint = (v5 + 1)
                    _ = v61
                }
            }
            var v62 int = 2147483647
            _ = fi57
            _ = v62
        }
        _ = fi51
    }
    v63 := GenS4{F0: 42, F1: ""}
    var v64 uint64 = 1
    if ((v64 / (((v64 - v64)) | 1)) <= v64) {
        v65 := []int{}
        v1 -= (((v1 / float32((0.1 * (0.1 - (1.5 / ((((1.5 / (0.1 - 0.0)) + -2.25) + -2.25) - 1e10)))))) + v1) - v1)
        var v66 string = "\n\t"
        _ = v65
        _ = v66
    }
    fmt.Println(v64)
    v67 := GenS3{F0: (^(v64)), F1: 255, F2: v64}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v50
    _ = v63
    _ = v64
    _ = v67
}
