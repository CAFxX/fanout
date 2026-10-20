package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 int16; F2 uint16 }
func genFn0() int32 {
    switch ("a" + "a") {
    case "\n\t":
        switch 4294967295 {
        case 42:
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 int32 = 624
                v8 = v8
                _ = fi7
                _ = v8
            }
        }
    case "hello":
        v9 := GenS1{F0: 65535}
        _ = v9
    default:
        if false {
            var v10 int8 = -13
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        if ((3.14159 + 1.5) >= ((-2.25 / 1e10) - -2.25)) {
            fmt.Println(fi11)
        } else {
            v12 := GenS1{F0: 32767}
            var v13 uint32 = 218
            var v14 int = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := GenS2{F0: 717, F1: ""}
        v17 := GenS2{F0: 42, F1: "hello"}
        for fi18 := 0; fi18 < 9; fi18++ {
            if (4294967295 < 42) {
                var v19 float32 = 1e10
                v20 := GenS3{F0: 2147483648, F1: -1}
                fmt.Println(v19)
                _ = v19
                _ = v20
            }
            _ = fi18
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    var v21 bool = (true || true)
    var v22 uint16 = 65535
    _ = v21
    _ = v22
    return -1
}
func genFn1(p0 uint, p1 int64) bool {
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 int16 = 1000
        for fi25 := 0; fi25 < 6; fi25++ {
            switch "\n\t" {
            case "":
                p1 = p1
            case "a":
                if (108 <= 42) {
                    switch "a" {
                    case "":
                        var v26 uint8 = 255
                        _ = v26
                    case "hello":
                        var v27 int64 = 0
                        _ = v27
                    case "\n\t":
                        v28 := GenS3{F0: uint32(p0), F1: (p1 << ((p0) & 63))}
                        if (2147483648 != 1) {
                            switch (^(fi25)) {
                            case 127:
                                p0 = (p0 | p0)
                            }
                        }
                        _ = v28
                    }
                    var v29 uint8 = 1
                    v30 := GenS2{F0: 255, F1: "hello"}
                    switch (fi25 / (((fi25 | 9223372036854775807)) | 1)) {
                    case -2147483648:
                        v31 := GenS2{F0: 32767, F1: "foo bar"}
                        _ = v31
                    }
                    fmt.Println(v29, v24, v29)
                    _ = v29
                    _ = v30
                } else {
                    var v32 float64 = (3.14159 / 1.0)
                    _ = v32
                }
            }
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    v33 := []int{-377}
    var v34 int = 4294967295
    _ = p0
    _ = p1
    _ = v33
    _ = v34
    return false
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 80
    var v3 float64 = 0.0
    var v4 int64 = 9223372036854775807
    v5 := make(map[string]bool)
    var v6 string = "\n\t"
    switch 127 {
    case -1:
        if (!(!false)) {
            if false {
                if (!false) {
                    var v35 uint64 = 471
                    switch ("foo bar" + "\n\t") {
                    case "hello":
                        for fi36 := 0; fi36 < 13; fi36++ {
                            fmt.Println(v4)
                            switch (fi36 & (fi36 / ((fi36) | 1))) {
                            case -362:
                                switch 127 {
                                case -749:
                                    var v37 bool = true
                                    _ = v37
                                }
                            }
                            v35 = (v35 << ((0) & 63))
                            _ = fi36
                        }
                    case "\n\t":
                        fmt.Println(v35, v3, v2)
                    case "a":
                        for fi38 := 0; fi38 < 6; fi38++ {
                            var v39 uint8 = 193
                            v6 = (v6 + "foo bar")
                            _ = fi38
                            _ = v39
                        }
                    }
                    v2 += v2
                    _ = v35
                } else {
                    v40 := GenS3{F0: 499, F1: 308}
                    if (!true) {
                        var v41 uint64 = 4294967295
                        _ = v41
                    }
                    v3 = (float64(0.1) / 1e10)
                    _ = v40
                }
                v42 := GenS4{F0: int(v2)}
                _ = v42
            }
        } else {
            for fi43 := 0; fi43 < 18; fi43++ {
                for fi44 := 0; fi44 < 5; fi44++ {
                    var v45 int8 = -78
                    switch (fi44 ^ fi43) {
                    case 118:
                        if true {
                            var v46 uint16 = 255
                            _ = v46
                        }
                    case 127:
                        var v47 string = "hello"
                        _ = v47
                    default:
                        fmt.Println(v2, fi43)
                    }
                    if true {
                        v6 = (v6 + v6)
                    }
                    _ = fi44
                    _ = v45
                }
                for fi48 := 0; fi48 < 1; fi48++ {
                    v49 := GenS5{F0: 42}
                    _ = fi48
                    _ = v49
                }
                _ = fi43
            }
        }
    case 4294967295:
        fmt.Println(v4, v6)
    case 128:
        v50 := []int{}
        _ = v50
    }
    if ((545 < 4294967295) || true) {
        for fi51 := 0; fi51 < 8; fi51++ {
            v52 := GenS2{F0: 42, F1: "hello"}
            v53 := []int{fi51, int(v2), (fi51 + fi51), 32767}
            v5[""] = false
            _ = fi51
            _ = v52
            _ = v53
        }
        for fi54 := 0; fi54 < 5; fi54++ {
            v6 = "hello"
            if (42 < 0) {
                var v55 string = v6
                var v56 int8 = -1
                v3 = v3
                _ = v55
                _ = v56
            } else {
                fmt.Println(v6)
                var v57 uint16 = 65535
                _ = v57
            }
            _ = fi54
        }
        v3 += -2.25
    }
    if true {
        v58 := GenS4{F0: 32768}
        delete(v5, v6)
        v59 := GenS6{F0: 18446744073709551615, F1: -49, F2: v2}
        _ = v58
        _ = v59
    } else {
        v60 := GenS5{F0: 94}
        var v61 int8 = int8(v4)
        _ = v60
        _ = v61
    }
    v6 = v6
    v62 := GenS3{F0: 2147483648, F1: (v4 ^ (v4 - (v4 | -426)))}
    var v63 float32 = -2.25
    v4 -= v4
    if false {
        v64 := v1[733]
        _ = v64
    }
    var v65 int64 = v4
    v66 := GenS1{F0: v2}
    for fi67 := 0; fi67 < 19; fi67++ {
        v68 := GenS6{F0: 42, F1: -1, F2: 0}
        v69 := GenS6{F0: 42, F1: 892, F2: (v2 ^ 65535)}
        _ = fi67
        _ = v68
        _ = v69
    }
    var v70 uint = 663
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
    _ = v65
    _ = v66
    _ = v70
}
