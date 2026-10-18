package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS3 struct { F0 int8; F1 uint64; F2 int64 }
func genFn0(p0 uint64, p1 uint16, p2 bool) int64 {
    for fi5 := 0; fi5 < 20; fi5++ {
        p0 -= (p0 ^ (p0 % ((p0) | 1)))
        _ = fi5
    }
    p2 = true
    var v6 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 178
}
func genFn1(p0 uint64) uint {
    if false {
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := []int{}
                var v10 int16 = 1
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    } else {
        if (388 > 1) {
            switch 128 {
            case 9223372036854775807:
                var v11 bool = true
                _ = v11
            default:
                for fi12 := 0; fi12 < 7; fi12++ {
                    if true {
                        v13 := GenS1{F0: 4294967295}
                        for fi14 := 0; fi14 < 18; fi14++ {
                            v15 := GenS1{F0: 65535}
                            var v16 uint = 42
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        switch "" {
                        case "a":
                            if false {
                                for fi17 := 0; fi17 < 19; fi17++ {
                                    p0 += p0
                                    _ = fi17
                                }
                                v18 := GenS1{F0: 4294967295}
                                var v19 uint64 = p0
                                _ = v18
                                _ = v19
                            }
                        case "":
                            v20 := GenS2{F0: 193, F1: (p0 ^ p0), F2: 628}
                            _ = v20
                        case "hello":
                            p0 += p0
                        default:
                            v21 := []int{(^(fi12)), (^(fi12)), fi12}
                            _ = v21
                        }
                        _ = v13
                    }
                    v22 := []int{}
                    var v23 int8 = -1
                    _ = fi12
                    _ = v22
                    _ = v23
                }
            }
        }
        if true {
            var v24 uint8 = 236
            for fi25 := 0; fi25 < 20; fi25++ {
                v26 := []int{(fi25 % ((42) | 1))}
                v27 := GenS1{F0: 4294967295}
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = v24
        } else {
            if true {
                var v28 int16 = -32768
                var v29 uint8 = 2
                _ = v28
                _ = v29
            } else {
                v30 := GenS1{F0: 65535}
                v31 := []int{851, 42}
                _ = v30
                _ = v31
            }
        }
    }
    var v32 uint32 = 0
    var v33 string = ("" + ("foo bar" + "\n\t"))
    _ = p0
    _ = v32
    _ = v33
    return 72
}

func main() {
    var v1 string = ("" + "a")
    var v2 uint32 = 762
    v3 := []uint32{(v2 + (v2 & 2)), 289, v2}
    var v4 int32 = 42
    v34 := GenS2{F0: 2, F1: 357, F2: 42}
    v1 = v1
    v2 += v2
    if (!(0 > 1)) {
        v4 += 0
        var v35 uint8 = 2
        v2 = v2
        _ = v35
    } else {
        if true {
            v4 = (v4 | v4)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        if ((false || true) || true) {
            var v36 string = v1
            switch 9223372036854775807 {
            case -2147483648:
                fmt.Println(v4, v36, v2)
            }
            v1 = (v1 + "hello")
            _ = v36
        } else {
            fmt.Println(v2)
        }
    }
    if false {
        v3 = append(v3, 525)
        v2 = v2
    } else {
        for fi37 := 0; fi37 < 17; fi37++ {
            _ = len(v3)
            _ = fi37
        }
    }
    var v38 int32 = v4
    for fi39 := 0; fi39 < 5; fi39++ {
        v40 := GenS1{F0: 18446744073709551615}
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        var v42 uint64 = 18446744073709551615
        v42 += (v42 << ((18446744073709551615) & 63))
        v1 = ""
        _ = fi41
        _ = v42
    }
    v38 = -562
    for fi43 := 0; fi43 < 18; fi43++ {
        switch v1 {
        case "":
            var v44 int = (fi43 / (((^(fi43))) | 1))
            _ = v44
        case "hello":
            v45 := GenS2{F0: 182, F1: 287, F2: -583}
            _ = v45
        default:
            v46 := GenS3{F0: -90, F1: 42, F2: 42}
            _ = v46
        }
        v4 += (-(v4))
        if (false && (532 < 4294967295)) {
            v4 += -501
            switch v1 {
            case "foo bar":
                for fi47 := 0; fi47 < 12; fi47++ {
                    v3 = append(v3, 937)
                    fmt.Println(fi43, fi43)
                    _ = fi47
                }
            }
        }
        _ = fi43
    }
    v2 -= (v2 + 1)
    v48 := GenS1{F0: 596}
    v1 = "foo bar"
    v38 += (v4 | (v38 * 42))
    v49 := GenS1{F0: 99}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v48
    _ = v49
}
