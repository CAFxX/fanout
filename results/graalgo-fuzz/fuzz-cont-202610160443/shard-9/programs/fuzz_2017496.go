package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int8; F1 int64 }
func genFn0(p0 uint16) uint32 {
    var v4 int = -835
    v5 := GenS1{F0: -483, F1: 75, F2: 9223372036854775807}
    if false {
        switch "hello" {
        case "":
            var v6 int8 = 1
            _ = v6
        case "\n\t":
            switch (-(v4)) {
            case 917:
                if true {
                    for fi7 := 0; fi7 < 18; fi7++ {
                        v8 := GenS2{F0: ("a" + "a")}
                        _ = fi7
                        _ = v8
                    }
                    v4 += v4
                }
            case 9223372036854775807:
                var v9 uint8 = 127
                _ = v9
            case -552:
                v4 += 1
            default:
                if (1 >= 0) {
                    switch ("foo bar" + ("\n\t" + "a")) {
                    case "":
                        var v10 uint16 = 42
                        _ = v10
                    }
                }
            }
        case "foo bar":
            if false {
                var v11 uint8 = 6
                v12 := GenS2{F0: ("\n\t" + "a")}
                _ = v11
                _ = v12
            } else {
                switch v4 {
                case -2147483648:
                    for fi13 := 0; fi13 < 2; fi13++ {
                        v14 := GenS1{F0: 0, F1: 63, F2: 837}
                        v15 := GenS1{F0: 1000, F1: 956, F2: -784}
                        v16 := GenS1{F0: 0, F1: 1, F2: -9223372036854775808}
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                case -597:
                    var v17 uint8 = 1
                    _ = v17
                }
                v18 := GenS3{F0: -127, F1: true}
                _ = v18
            }
        default:
            for fi19 := 0; fi19 < 14; fi19++ {
                if (!(true || true)) {
                    p0 -= (p0 * p0)
                }
                if (!false) {
                    var v20 uint = 1
                    v21 := []int{(v4 * 528), fi19}
                    var v22 uint32 = 2
                    _ = v20
                    _ = v21
                    _ = v22
                }
                _ = fi19
            }
        }
    } else {
        v23 := []int{v4}
        _ = v23
    }
    _ = p0
    _ = v4
    _ = v5
    return 2147483647
}
func genFn1(p0 int, p1 uint, p2 int32) int16 {
    p0 = p0
    v24 := GenS2{F0: ("foo bar" + "a")}
    for fi25 := 0; fi25 < 13; fi25++ {
        if false {
            v26 := GenS2{F0: ""}
            v27 := []int{(fi25 / (((p0 + 9223372036854775807)) | 1)), (p0 | fi25)}
            _ = v26
            _ = v27
        } else {
            if (float32(0.0) > (0.1 / 3.14159)) {
                v28 := []int{(p0 % ((fi25) | 1)), p0, fi25}
                _ = v28
            }
            var v29 uint64 = 572
            _ = v29
        }
        _ = fi25
    }
    var v30 string = "a"
    if ((p1 | (p1 * p1)) > p1) {
        if (!false) {
            p2 = p2
            v31 := GenS1{F0: 1000, F1: -778, F2: -1}
            var v32 uint = (p1 << (((p1 ^ 4294967295)) & 63))
            _ = v31
            _ = v32
        }
        var v33 int16 = -32768
        _ = v33
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v30
    return 0
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 222
    if (42 != 291) {
        for fi34 := 0; fi34 < 6; fi34++ {
            v1 = v1
            _ = fi34
        }
        v3 = v3
    }
    v35 := GenS3{F0: -128, F1: (!((-32768 <= int16(v3)) || false))}
    if (1 == 1) {
        v36 := []int{int(v3), -834, -277, -169}
        _ = v36
    }
    for fi37 := 0; fi37 < 16; fi37++ {
        v2[2] = 718
        v38 := GenS1{F0: genFn1((fi37 << ((191) & 63)), 374, -77), F1: 1, F2: 9223372036854775807}
        fmt.Println(fi37)
        _ = fi37
        _ = v38
    }
    v3 -= (v3 + (^(v3)))
    v1 = v1
    for fi39 := 0; fi39 < 20; fi39++ {
        var v40 uint64 = 18446744073709551615
        v41 := GenS4{F0: 1, F1: 1}
        for fi42 := 0; fi42 < 16; fi42++ {
            v43 := GenS3{F0: 100, F1: false}
            var v44 float32 = (((1.5 - 1.5) / 1.5) * (1e10 + (float32((0.0 - 1.5)) / 3.14159)))
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi39
        _ = v40
        _ = v41
    }
    v3 = v3
    var v45 int = -754
    v3 = (^(v3))
    v45 -= (v45 / (((^(v45))) | 1))
    v2[42] = 590
    v3 -= v3
    v3 -= (v3 >> ((42) & 63))
    for fi46 := 0; fi46 < 11; fi46++ {
        for fi47 := 0; fi47 < 2; fi47++ {
            if true {
                v3 = (^(v3))
                var v48 uint32 = 65535
                var v49 int32 = -2147483648
                _ = v48
                _ = v49
            }
            _ = fi47
        }
        fmt.Println(v3, v1)
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v45
}
