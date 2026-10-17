package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS2 struct { F0 int8; F1 int64; F2 int32 }
type GenS3 struct { F0 int64; F1 uint }
func genFn0(p0 uint, p1 int16, p2 int16) int32 {
    v6 := GenS1{F0: 255, F1: 641, F2: 0}
    v7 := GenS2{F0: -85, F1: 42, F2: -2147483648}
    for fi8 := 0; fi8 < 20; fi8++ {
        p0 += (p0 % ((p0) | 1))
        var v9 uint = 65535
        var v10 int8 = 100
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int16 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint16, p1 int) uint8 {
    if true {
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 uint8 = 134
            _ = fi12
            _ = v13
        }
        p1 += (p1 + (p1 - (p1 % (((p1 | p1)) | 1))))
    }
    var v14 float32 = float32((0.1 - 1e10))
    var v15 uint8 = 128
    var v16 int16 = 0
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return (v15 << ((502) & 63))
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[string]int)
    var v3 int32 = 1
    v4 := []int{v1, (v1 / (((v1 & v1)) | 1)), 32768, -1}
    v5 := make(map[int]uint32)
    switch (v1 >> ((65) & 63)) {
    case 1:
        var v17 int16 = -526
        _ = v17
    default:
        v5[-836] = 1
    }
    switch (-(v1)) {
    case 517:
        if (!false) {
            v18 := GenS1{F0: genFn1(65535, (v1 % (((v1 | (v1 - v1))) | 1))), F1: 0, F2: 32767}
            _ = v18
        }
    case 2147483647:
        switch "foo bar" {
        case "\n\t":
            v1 = v1
        case "foo bar":
            var v19 uint8 = 255
            _ = v19
        default:
            fmt.Println(v3)
        }
    case 688:
        v20 := GenS3{F0: 42, F1: 42}
        _ = v20
    }
    v21 := v5[(v1 / ((v1) | 1))]
    v22 := GenS1{F0: 0, F1: 483, F2: 1}
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 int64 = 42
        for fi25 := 0; fi25 < 5; fi25++ {
            fmt.Println(fi23, v1)
            _ = fi25
        }
        for fi26 := 0; fi26 < 20; fi26++ {
            var v27 int = 4294967295
            _ = fi26
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    switch (v1 - (v1 % (((v1 >> ((838) & 63))) | 1))) {
    case -213:
        v3 = (v3 - (v3 >> ((65535) & 63)))
    default:
        for fi28 := 0; fi28 < 5; fi28++ {
            switch (fi28 + (^(v1))) {
            case 538:
                v29 := GenS1{F0: 2, F1: 2, F2: 473}
                _ = v29
            case 32767:
                var v30 uint64 = 2
                _ = v30
            case 1:
                v21 = v21
            }
            v31 := GenS1{F0: 0, F1: 0, F2: 683}
            v21 += (v21 >> ((777) & 63))
            _ = fi28
            _ = v31
        }
    }
    if false {
        v3 = -1
        if false {
            fmt.Println(v3, v1, v3)
            for fi32 := 0; fi32 < 15; fi32++ {
                if ((-843 != -147) || (268 == 960)) {
                    switch "foo bar" {
                    case "\n\t":
                        v33 := GenS2{F0: 127, F1: 42, F2: (v3 >> ((4294967295) & 63))}
                        _ = v33
                    default:
                        if false {
                            for fi34 := 0; fi34 < 19; fi34++ {
                                var v35 uint = 0
                                fmt.Println(v21)
                                fmt.Println(v1, v21, v21)
                                switch ("\n\t" + ("a" + "")) {
                                case "hello":
                                    fmt.Println(v21, v3, v3)
                                case "a":
                                    switch -540 {
                                    case 362:
                                        var v36 int = (v1 ^ (fi34 - (-(fi32))))
                                        _ = v36
                                    case 2147483647:
                                        var v37 uint64 = 4294967295
                                        _ = v37
                                    case -9223372036854775808:
                                        var v38 uint16 = 0
                                        _ = v38
                                    }
                                default:
                                    var v39 string = "\n\t"
                                    _ = v39
                                }
                                _ = fi34
                                _ = v35
                            }
                        } else {
                            if true {
                                fmt.Println(v1)
                                for fi40 := 0; fi40 < 12; fi40++ {
                                    var v41 uint64 = 18
                                    var v42 int16 = -593
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                }
                            }
                            for fi43 := 0; fi43 < 1; fi43++ {
                                delete(v5, (v1 + 0))
                                var v44 int32 = 42
                                v1 = 657
                                if len(v4) > 0 {
                                    v4[0] = (^(fi32))
                                }
                                if (-1 == 1) {
                                    v45 := GenS1{F0: 255, F1: 18446744073709551615, F2: 0}
                                    fmt.Println(v21, v21, v3)
                                    _ = v45
                                }
                                v44 -= (v3 ^ (v3 >> ((4294967295) & 63)))
                                var v46 int8 = -128
                                _ = fi43
                                _ = v44
                                _ = v46
                            }
                        }
                    }
                    var v47 uint8 = 42
                    _ = v47
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v48 uint32 = (v21 + v21)
                    var v49 int8 = 100
                    _ = v48
                    _ = v49
                }
                _ = fi32
            }
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                v50 := GenS1{F0: 137, F1: 629, F2: -1}
                v21 = 572
                v51 := GenS2{F0: -128, F1: -1, F2: v3}
                _ = v50
                _ = v51
            }
        }
        var v52 int8 = 100
        _ = v52
    } else {
        var v53 uint8 = 87
        v3 = v3
        v3 = 2147483647
        _ = v53
    }
    v4 = append(v4, (v1 ^ -157))
    if len(v4) > 0 {
        v4[0] = (v1 ^ v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
