package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64 }
func genFn0(p0 int16, p1 int16, p2 uint16) int {
    switch 516 {
    case 0:
        switch 4294967295 {
        case 128:
            v6 := GenS1{F0: 65535}
            _ = v6
        case -2147483648:
            if false {
                var v7 int64 = -1
                if ((-2.25 + 3.14159) >= ((0.0 / -2.25) - -2.25)) {
                    v8 := []int{}
                    var v9 string = "hello"
                    v10 := GenS1{F0: 0}
                    _ = v10
                    _ = v8
                    _ = v9
                }
                v7 += v7
                _ = v7
            }
        }
    case -175:
        var v11 uint16 = p2
        _ = v11
    }
    fmt.Println(p0)
    for fi12 := 0; fi12 < 14; fi12++ {
        p2 = p2
        var v13 int16 = p0
        var v14 bool = (false || (true || (-1 <= 1)))
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        p0 += 32767
        v16 := GenS1{F0: 264}
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := []string{"", "a", ((("a" + "\n\t") + ("a" + "foo bar")) + "a"), "foo bar"}
    v2 := make(map[uint32]string)
    var v3 bool = (!false)
    var v4 uint = 255
    var v5 uint16 = 0
    v17 := GenS2{F0: 0, F1: 2147483648}
    v1 = append(v1, "a")
    _ = len(v2)
    switch genFn0(1000, 876, v5) {
    case 641:
        var v18 int = 2147483647
        _ = v18
    case 4294967295:
        v3 = (255 != 42)
    case -667:
        if v3 {
            v3 = (v3 || (-32768 > 1000))
            if (!false) {
                for fi19 := 0; fi19 < 12; fi19++ {
                    if v3 {
                        var v20 bool = (0.1 <= 3.14159)
                        v4 += 319
                        _ = v20
                    }
                    v21 := v2[913]
                    var v22 string = v21
                    _ = fi19
                    _ = v21
                    _ = v22
                }
            }
        } else {
            var v23 string = ("hello" + "a")
            v3 = v3
            var v24 uint = (v4 & 898)
            _ = v23
            _ = v24
        }
    default:
        v25 := GenS3{F0: (true && v3)}
        _ = v25
    }
    v2[316] = "hello"
    v5 += 409
    v4 = v4
    switch "hello" {
    case "":
        v3 = v3
    default:
        v4 = v4
    }
    v26 := GenS1{F0: (v4 >> (((v4 / (((^(v4))) | 1))) & 63))}
    v3 = (("" + "a") > "")
    v3 = (!v3)
    var v27 uint64 = 591
    var v28 int = 32767
    for fi29 := 0; fi29 < 13; fi29++ {
        delete(v2, 0)
        if true {
            if (v3 || v3) {
                var v30 string = ("foo bar" + "a")
                _ = v30
            }
        }
        v1 = append(v1, "hello")
        _ = fi29
    }
    if false {
        switch ("hello" + "") {
        case "\n\t":
            for fi31 := 0; fi31 < 3; fi31++ {
                var v32 bool = v3
                var v33 int16 = 32767
                switch "" {
                case "foo bar":
                    for fi34 := 0; fi34 < 18; fi34++ {
                        fmt.Println(v32, v3, fi31)
                        _ = fi34
                    }
                default:
                    var v35 int64 = -9223372036854775808
                    _ = v35
                }
                _ = fi31
                _ = v32
                _ = v33
            }
        case "":
            v36 := GenS3{F0: v3}
            _ = v36
        default:
            var v37 uint16 = (v5 >> (((v4 + v4)) & 63))
            _ = v37
        }
        var v38 uint32 = 2
        v27 = (v27 * v27)
        _ = v38
    } else {
        switch genFn0(-814, int16(v28), v5) {
        case 139:
            for fi39 := 0; fi39 < 6; fi39++ {
                if false {
                    var v40 int16 = 340
                    v28 = (fi39 ^ v28)
                    fmt.Println(v5, v3)
                    _ = v40
                } else {
                    var v41 int8 = -1
                    _ = v41
                }
                if false {
                    v42 := GenS4{F0: 18446744073709551615, F1: uint8(v27), F2: 178}
                    _ = len(v1)
                    for fi43 := 0; fi43 < 20; fi43++ {
                        v44 := GenS5{F0: v27}
                        var v45 int8 = 42
                        _ = len(v2)
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    _ = v42
                }
                v46 := GenS6{F0: -796}
                _ = fi39
                _ = v46
            }
        default:
            switch 0 {
            case -9223372036854775808:
                var v47 int16 = int16(v4)
                _ = v47
            default:
                fmt.Println(v3)
            }
        }
        var v48 uint32 = 337
        _ = len(v2)
        _ = v48
    }
    v49 := GenS2{F0: -9223372036854775808, F1: 486}
    var v50 uint32 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
}
