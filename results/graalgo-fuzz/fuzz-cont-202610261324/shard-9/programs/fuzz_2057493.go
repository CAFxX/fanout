package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 int; F1 int; F2 int8 }
type GenS6 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint64; F1 int16 }
func genFn0(p0 string, p1 uint8, p2 uint) int16 {
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint16 = 65535
        if (!((2147483648 < 4294967295) && ((42 > 9223372036854775807) || false))) {
            v9 := GenS1{F0: 42, F1: 0}
            var v10 int8 = -1
            _ = v10
            _ = v9
        } else {
            v11 := GenS1{F0: -1, F1: -763}
            _ = v11
        }
        var v12 int64 = int64(p2)
        _ = fi7
        _ = v12
        _ = v8
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        switch (-(fi13)) {
        case 4294967295:
            v14 := GenS2{F0: int(p2), F1: 0}
            _ = v14
        case -9223372036854775808:
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 int16 = -891
                _ = fi15
                _ = v16
            }
        }
        _ = fi13
    }
    if (true && false) {
        var v17 uint16 = 1
        _ = v17
    } else {
        if true {
            v18 := GenS2{F0: -455, F1: (p2 - 255)}
            for fi19 := 0; fi19 < 12; fi19++ {
                for fi20 := 0; fi20 < 7; fi20++ {
                    v21 := []int{}
                    if (p1 < p1) {
                        var v22 uint = (p2 % (((p2 >> (((p2 * (p2 - 759))) & 63))) | 1))
                        switch -9223372036854775808 {
                        case 127:
                            var v23 int = 128
                            _ = v23
                        case 9223372036854775807:
                            v24 := GenS3{F0: p2, F1: 127}
                            _ = v24
                        }
                        p1 = (p1 + p1)
                        p2 = p2
                        switch fi19 {
                        case 9223372036854775807:
                            _ = len(v21)
                        default:
                            var v25 int16 = 578
                            _ = v25
                        }
                        var v26 int = (fi19 + fi19)
                        var v27 int64 = 0
                        _ = v22
                        _ = v26
                        _ = v27
                    } else {
                        p2 = (p2 >> ((42) & 63))
                        p2 += p2
                    }
                    _ = fi20
                    _ = v21
                }
                var v28 int32 = 42
                _ = fi19
                _ = v28
            }
            p2 += 4294967295
            _ = v18
        } else {
            v29 := []int{-924, int(p1), 0}
            var v30 float64 = 3.14159
            v30 -= 1e10
            _ = v29
            _ = v30
        }
    }
    v31 := GenS3{F0: p2, F1: 127}
    v32 := []int{0}
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    _ = v32
    return 32767
}

func main() {
    var v1 float32 = (3.14159 - float32(0.1))
    var v2 int = 32768
    v3 := []string{"hello", "a", ("a" + ("hello" + "foo bar")), "\n\t"}
    var v4 int8 = 41
    v5 := make(map[uint32]int)
    var v6 int64 = 9223372036854775807
    v2 = (^(v2))
    for fi33 := 0; fi33 < 15; fi33++ {
        if true {
            var v34 bool = (true || (false || (false && (128 == 25))))
            _ = v34
        }
        _ = fi33
    }
    switch "" {
    case "a":
        if false {
            switch "\n\t" {
            case "":
                v35 := GenS4{F0: 65535, F1: 255}
                _ = v35
            }
            v2 = int(v6)
            _ = len(v5)
        } else {
            for fi36 := 0; fi36 < 6; fi36++ {
                v5[16] = fi36
                _ = fi36
            }
            v37 := v5[42]
            if true {
                v6 += 1
                v38 := GenS5{F0: v37, F1: v2, F2: v4}
                v4 = v4
                _ = v38
            }
            _ = v37
        }
    case "":
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    case "foo bar":
        if (false || false) {
            v39 := GenS4{F0: 0, F1: 824}
            for fi40 := 0; fi40 < 14; fi40++ {
                var v41 uint8 = 63
                for fi42 := 0; fi42 < 2; fi42++ {
                    if (42 <= 42) {
                        var v43 bool = false
                        var v44 uint32 = 1
                        if (-465 != -398) {
                            var v45 int16 = 32767
                            _ = v45
                        } else {
                            v46 := GenS3{F0: 271, F1: v4}
                            v2 += 32768
                            v3 = append(v3, "a")
                            _ = v46
                        }
                        _ = v43
                        _ = v44
                    }
                    if false {
                        v47 := GenS4{F0: 2147483647, F1: 42}
                        v6 = v6
                        _ = v47
                    } else {
                        fmt.Println(v2, v41, fi42)
                        v41 = (v41 * v41)
                        v1 = 3.14159
                    }
                    _ = fi42
                }
                _ = fi40
                _ = v41
            }
            _ = v39
        } else {
            v3 = append(v3, ("" + ("a" + "")))
            var v48 int = 1
            _ = v48
        }
    }
    var v49 uint8 = 255
    var v50 int32 = 42
    fmt.Println(v2)
    var v51 uint = 2
    v52 := GenS4{F0: 2147483648, F1: 2}
    v53 := GenS6{F0: v51, F1: 0, F2: 774}
    v54 := GenS7{F0: 4294967295, F1: -1}
    v50 += v50
    if (!true) {
        for fi55 := 0; fi55 < 2; fi55++ {
            for fi56 := 0; fi56 < 2; fi56++ {
                var v57 bool = true
                v1 = float32(3.14159)
                if (-32768 >= 0) {
                    var v58 int64 = v6
                    var v59 uint64 = 1
                    v51 = v51
                    _ = v58
                    _ = v59
                }
                _ = fi56
                _ = v57
            }
            _ = fi55
        }
        v51 = v51
    } else {
        v4 = 7
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
