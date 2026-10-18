package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint8; F1 int32; F2 int16 }
type GenS5 struct { F0 uint16; F1 int32; F2 int64 }
func genFn0(p0 int, p1 int8, p2 int8) uint {
    var v7 int8 = (p1 % (((p2 % ((p1) | 1))) | 1))
    var v8 int = (p0 >> ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint8 = 85
    v2 := []int{128, 919}
    var v3 uint32 = 470
    var v4 float64 = 3.14159
    v5 := make(map[int]uint32)
    var v6 bool = (3.14159 == 0.1)
    switch ("hello" + "foo bar") {
    case "":
        var v9 int8 = 100
        _ = v9
    case "a":
        for fi10 := 0; fi10 < 7; fi10++ {
            if true {
                if false {
                    if (!(2147483647 < 0)) {
                        var v11 int8 = 1
                        var v12 float64 = v4
                        _ = v11
                        _ = v12
                    }
                    v13 := v5[2147483647]
                    _ = v13
                }
                v3 = (^(v3))
            } else {
                var v14 uint = 161
                _ = v14
            }
            fmt.Println(v3, v1)
            switch "\n\t" {
            case "foo bar":
                v1 = 127
            default:
                v15 := GenS1{F0: 42}
                _ = v15
            }
            _ = fi10
        }
    }
    var v16 int = -183
    v6 = (v6 && v6)
    switch v16 {
    case 9223372036854775807:
        fmt.Println(v4, v16, v6)
    case 1:
        for fi17 := 0; fi17 < 10; fi17++ {
            if (false && (!v6)) {
                if (0.0 <= 3.14159) {
                    v18 := GenS2{F0: 32767}
                    fmt.Println(v1, v16)
                    _ = v18
                } else {
                    if (42 >= 42) {
                        if ((4294967295 != 2) && (true && true)) {
                            var v19 uint64 = 2
                            if (1000 < 0) {
                                var v20 uint32 = (v3 + 170)
                                var v21 uint32 = v20
                                for fi22 := 0; fi22 < 4; fi22++ {
                                    switch ("hello" + ("foo bar" + "hello")) {
                                    case "":
                                        fmt.Println(v20)
                                    case "\n\t":
                                        v20 = (v21 + 2147483648)
                                    case "foo bar":
                                        var v23 uint16 = 385
                                        _ = v23
                                    default:
                                        var v24 float32 = 3.14159
                                        _ = v24
                                    }
                                    _ = fi22
                                }
                                _ = v20
                                _ = v21
                            } else {
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    var v26 uint16 = uint16(v1)
                                    if len(v2) > 0 {
                                        v2[0] = 42
                                    }
                                    _ = fi25
                                    _ = v26
                                }
                                for fi27 := 0; fi27 < 4; fi27++ {
                                    switch "a" {
                                    case "":
                                        v2 = append(v2, fi27)
                                    }
                                    _ = fi27
                                }
                            }
                            _ = v19
                        }
                        if v6 {
                            fmt.Println(v3, v3)
                            var v28 int32 = -555
                            if v6 {
                                switch v16 {
                                case 280:
                                    for fi29 := 0; fi29 < 5; fi29++ {
                                        var v30 float32 = (1.5 - 0.0)
                                        var v31 int16 = -205
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                    }
                                case -9223372036854775808:
                                    for fi32 := 0; fi32 < 5; fi32++ {
                                        var v33 string = "foo bar"
                                        v2 = append(v2, (fi32 - v16))
                                        v16 = int(v3)
                                        _ = fi32
                                        _ = v33
                                    }
                                case -2147483648:
                                    fmt.Println(v6, v1)
                                }
                                fmt.Println(v6)
                            }
                            _ = v28
                        }
                        for fi34 := 0; fi34 < 17; fi34++ {
                            var v35 uint64 = 774
                            fmt.Println(v1)
                            _ = fi34
                            _ = v35
                        }
                    } else {
                        fmt.Println(v6, fi17, fi17)
                        v36 := v5[v16]
                        _ = v36
                    }
                }
                v37 := v5[(fi17 % (((fi17 + v16)) | 1))]
                v37 = v3
                _ = v37
            } else {
                var v38 string = ""
                var v39 uint16 = 42
                _ = v38
                _ = v39
            }
            _ = fi17
        }
    case 4294967295:
        for fi40 := 0; fi40 < 2; fi40++ {
            fmt.Println(v16, v3)
            var v41 float64 = v4
            var v42 bool = (0 >= int8(v16))
            _ = fi40
            _ = v41
            _ = v42
        }
    default:
        switch ("a" + ("hello" + "foo bar")) {
        case "foo bar":
            var v43 uint16 = 42
            _ = v43
        case "a":
            v44 := GenS1{F0: 42}
            _ = v44
        case "hello":
            v6 = false
        default:
            v3 = 42
        }
    }
    var v45 int8 = 42
    if false {
        _ = len(v2)
    } else {
        if v6 {
            for fi46 := 0; fi46 < 13; fi46++ {
                var v47 int = fi46
                _ = fi46
                _ = v47
            }
            var v48 int = (v16 / ((0) | 1))
            _ = v48
        } else {
            switch "hello" {
            case "\n\t":
                v49 := GenS3{F0: 18446744073709551615, F1: (v45 + v45)}
                _ = v49
            case "foo bar":
                v16 = v16
            case "hello":
                switch (v16 + 9223372036854775807) {
                case 9223372036854775807:
                    for fi50 := 0; fi50 < 2; fi50++ {
                        v51 := GenS4{F0: v1, F1: 808, F2: 1000}
                        var v52 int8 = v45
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                case 32767:
                    var v53 float64 = (((0.1 / -2.25) * float64(3.14159)) / v4)
                    _ = v53
                case 868:
                    if (!v6) {
                        v54 := GenS5{F0: 65535, F1: 1000, F2: 9223372036854775807}
                        _ = v54
                    }
                default:
                    var v55 float32 = 1e10
                    _ = v55
                }
            default:
                v6 = true
            }
        }
        var v56 uint32 = v3
        fmt.Println(v1, v45, v16)
        _ = v56
    }
    v57 := GenS2{F0: -32768}
    delete(v5, (v16 / ((951) | 1)))
    fmt.Println(v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v57
    _ = v6
}
