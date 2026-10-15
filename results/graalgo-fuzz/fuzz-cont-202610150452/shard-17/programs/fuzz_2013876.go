package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int64; F1 int8; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 uint64; F2 int }
type GenS5 struct { F0 int16 }

func main() {
    var v1 int8 = 127
    var v2 uint32 = 0
    var v3 int64 = 0
    v4 := []int{}
    var v5 float64 = 1.5
    if true {
        fmt.Println(v3)
        v6 := GenS1{F0: ("foo bar" + "hello"), F1: 32768}
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: "a", F1: fi7}
            var v9 uint = 0
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    }
    v3 = (v3 >> ((4294967295) & 63))
    switch -528 {
    case 12:
        v10 := GenS2{F0: 794, F1: 1, F2: v1}
        _ = v10
    }
    v11 := GenS2{F0: v3, F1: v1, F2: (v1 / ((100) | 1))}
    if false {
        switch "foo bar" {
        case "":
            if len(v4) > 0 {
                v4[0] = 0
            }
        case "a":
            if false {
                var v12 uint = 255
                v4 = append(v4, -704)
                v1 = (v1 + (v1 / ((int8(v2)) | 1)))
                _ = v12
            } else {
                _ = len(v4)
                fmt.Println(v5, v5, v5)
                switch -518 {
                case 127:
                    v13 := GenS1{F0: "", F1: 32767}
                    _ = v13
                }
            }
        default:
            for fi14 := 0; fi14 < 11; fi14++ {
                if true {
                    for fi15 := 0; fi15 < 9; fi15++ {
                        if (!(true && true)) {
                            switch fi15 {
                            case 9223372036854775807:
                                v4 = append(v4, fi15)
                            case -396:
                                var v16 uint = 65535
                                _ = v16
                            default:
                                fmt.Println(fi14)
                            }
                            var v17 int64 = 500
                            fmt.Println(v2)
                            v2 = (v2 * 2147483648)
                            v18 := GenS3{F0: (v1 * v1)}
                            _ = v17
                            _ = v18
                        } else {
                            fmt.Println(v5, v2, v5)
                            fmt.Println(fi14, v2, v3)
                        }
                        _ = fi15
                    }
                    _ = len(v4)
                    v19 := GenS3{F0: (v1 + 4)}
                    _ = v19
                } else {
                    v2 = (v2 & v2)
                    var v20 uint32 = (v2 - (v2 / ((v2) | 1)))
                    if false {
                        if true {
                            switch "foo bar" {
                            case "\n\t":
                                var v21 int64 = v3
                                _ = v21
                            case "hello":
                                v3 += (v3 << ((18446744073709551615) & 63))
                            default:
                                var v22 uint64 = 18446744073709551615
                                _ = v22
                            }
                        } else {
                            v23 := GenS4{F0: 18446744073709551615, F1: 1, F2: int(v1)}
                            v2 += v20
                            _ = v23
                        }
                        switch (fi14 | (fi14 ^ fi14)) {
                        case -1:
                            if len(v4) > 0 {
                                v4[0] = fi14
                            }
                        case 817:
                            v2 = 0
                        case -215:
                            v4 = append(v4, (fi14 | fi14))
                        }
                    } else {
                        if len(v4) > 0 {
                            v4[0] = (fi14 | 1)
                        }
                        switch "hello" {
                        case "\n\t":
                            _ = len(v4)
                        default:
                            var v24 uint = 2
                            _ = v24
                        }
                        switch (-(fi14)) {
                        case -411:
                            var v25 uint64 = 42
                            _ = v25
                        case 486:
                            v26 := GenS3{F0: 1}
                            _ = v26
                        }
                    }
                    _ = v20
                }
                _ = fi14
            }
        }
        var v27 float64 = float64(1.5)
        _ = v27
    } else {
        if ((v3 / ((v3) | 1)) > (-(v3))) {
            var v28 bool = false
            var v29 float64 = (float64(1e10) + v5)
            v29 -= v5
            _ = v28
            _ = v29
        }
    }
    var v30 int = -691
    v4 = append(v4, 32767)
    fmt.Println(v30)
    if false {
        if len(v4) > 0 {
            v4[0] = v30
        }
        var v31 uint32 = v2
        _ = v31
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v32 int8 = v1
    v2 = v2
    if (v2 >= (v2 % ((v2) | 1))) {
        var v33 float32 = 0.0
        _ = v33
    } else {
        if false {
            var v34 uint8 = 0
            var v35 float32 = 3.14159
            _ = v34
            _ = v35
        } else {
            v32 = (v32 ^ 96)
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi36 := 0; fi36 < 6; fi36++ {
            var v37 int64 = (v3 & -1)
            v38 := GenS5{F0: -32768}
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    var v39 int8 = v32
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 -= (v1 ^ (v32 - (v1 % ((int8(v3)) | 1))))
    for fi40 := 0; fi40 < 3; fi40++ {
        if true {
            if ((v2 <= (v2 * (v2 - v2))) && (true || (("" + "") > ("hello" + "")))) {
                v1 -= int8(fi40)
                if (255 > 300) {
                    var v41 int32 = 0
                    var v42 uint32 = 976
                    _ = v41
                    _ = v42
                } else {
                    v43 := GenS3{F0: (v1 % ((1) | 1))}
                    fmt.Println(v32)
                    v2 += v2
                    _ = v43
                }
            }
            v3 += -9223372036854775808
            if len(v4) > 0 {
                v4[0] = 9223372036854775807
            }
        } else {
            if (2147483648 >= 0) {
                var v44 int16 = 32767
                _ = len(v4)
                var v45 uint = 1
                _ = v44
                _ = v45
            } else {
                for fi46 := 0; fi46 < 3; fi46++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v47 int16 = -32768
                    v32 -= v39
                    _ = fi46
                    _ = v47
                }
            }
        }
        _ = fi40
    }
    var v48 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v39
    _ = v4
    _ = v48
    _ = v5
}
