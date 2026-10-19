package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int16; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int; F2 int16 }
type GenS5 struct { F0 int32; F1 uint16 }
func genFn0(p0 bool, p1 uint) uint16 {
    p1 += 4294967295
    v4 := GenS1{F0: 42, F1: -826}
    var v5 uint16 = 718
    v6 := GenS2{F0: 43, F1: 696}
    v7 := GenS3{F0: -597, F1: -1, F2: 0}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (^(v5))
}
func genFn1() int16 {
    v8 := GenS4{F0: 1, F1: 0, F2: 0}
    var v9 uint64 = 1
    v9 -= 1
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int8 = 82
    v2 := make(map[uint32]bool)
    var v3 int16 = 32767
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS5{F0: 25, F1: 42}
        for fi12 := 0; fi12 < 12; fi12++ {
            if (!true) {
                if (("" + "foo bar") <= "\n\t") {
                    v1 -= v1
                }
                var v13 int16 = (v3 >> ((42) & 63))
                _ = v13
            }
            v14 := []int{(fi12 | (^(fi12)))}
            _ = fi12
            _ = v14
        }
        var v15 float64 = 3.14159
        _ = fi10
        _ = v11
        _ = v15
    }
    v16 := v2[103]
    switch "foo bar" {
    case "foo bar":
        v17 := GenS2{F0: 1, F1: -901}
        _ = v17
    case "a":
        switch "" {
        case "foo bar":
            v2[42] = (!true)
        case "\n\t":
            v1 = 0
        case "":
            for fi18 := 0; fi18 < 19; fi18++ {
                switch fi18 {
                case -9223372036854775808:
                    v16 = v16
                case 42:
                    v1 -= v1
                case 4294967295:
                    v19 := GenS3{F0: (-(v3)), F1: v3, F2: 615}
                    _ = v19
                default:
                    var v20 uint8 = 2
                    _ = v20
                }
                _ = fi18
            }
        default:
            var v21 string = "\n\t"
            _ = v21
        }
    }
    if v16 {
        var v22 uint = 2
        if ((v1 << ((v22) & 63)) != v1) {
            for fi23 := 0; fi23 < 18; fi23++ {
                for fi24 := 0; fi24 < 4; fi24++ {
                    if false {
                        delete(v2, 0)
                    }
                    v22 = (^(v22))
                    _ = fi24
                }
                switch 962 {
                case -2147483648:
                    v2[1] = v16
                case -1:
                    var v25 int64 = -1
                    _ = v25
                default:
                    v3 = (v3 >> (((v22 >> ((v22) & 63))) & 63))
                }
                _ = fi23
            }
        }
        if v16 {
            for fi26 := 0; fi26 < 13; fi26++ {
                for fi27 := 0; fi27 < 1; fi27++ {
                    var v28 int16 = 1
                    v29 := GenS1{F0: 159, F1: -9223372036854775808}
                    switch (-(fi26)) {
                    case -9223372036854775808:
                        var v30 bool = true
                        _ = v30
                    case 4294967295:
                        var v31 uint32 = 65535
                        _ = v31
                    default:
                        v3 = genFn1()
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                if (215 == 255) {
                    var v32 int16 = -32768
                    for fi33 := 0; fi33 < 16; fi33++ {
                        fmt.Println(v3)
                        var v34 int = -2147483648
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                } else {
                    switch 2147483647 {
                    case 128:
                        var v35 int32 = -2147483648
                        _ = v35
                    case -52:
                        for fi36 := 0; fi36 < 3; fi36++ {
                            if (1 > 2) {
                                _ = len(v2)
                            } else {
                                v3 -= v3
                                v37 := GenS2{F0: 2, F1: -266}
                                v1 += 42
                                fmt.Println(fi26, fi26)
                                var v38 int16 = int16(v22)
                                _ = v37
                                _ = v38
                            }
                            v1 -= v1
                            _ = fi36
                        }
                    }
                    if (42 != 42) {
                        var v39 string = "hello"
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 15; fi40++ {
                        v41 := v2[42]
                        _ = fi40
                        _ = v41
                    }
                }
                _ = fi26
            }
        } else {
            var v42 bool = (genFn1() > genFn1())
            for fi43 := 0; fi43 < 15; fi43++ {
                v44 := GenS2{F0: 18446744073709551615, F1: int64(v1)}
                _ = fi43
                _ = v44
            }
            v22 = (v22 << (((v22 * v22)) & 63))
            _ = v42
        }
        _ = v22
    }
    v1 = v1
    var v45 int16 = 1
    v46 := []int{42, 128}
    fmt.Println(v16, v16)
    for fi47 := 0; fi47 < 5; fi47++ {
        if len(v46) > 0 {
            v46[0] = -2147483648
        }
        var v48 uint64 = 0
        _ = fi47
        _ = v48
    }
    v1 += v1
    v2[0] = (65535 >= 32767)
    var v49 uint32 = 889
    v50 := GenS3{F0: (v3 & 1), F1: (v45 % ((574) | 1)), F2: 42}
    for fi51 := 0; fi51 < 1; fi51++ {
        for fi52 := 0; fi52 < 17; fi52++ {
            var v53 int64 = int64(v45)
            switch (fi52 - 139) {
            case -9223372036854775808:
                v1 = (v1 & (v1 & (v1 / ((0) | 1))))
            case 4294967295:
                var v54 int64 = int64(v1)
                _ = v54
            case 32768:
                if (!((!((^(fi52)) != (fi52 | fi52))) || (!v16))) {
                    v3 -= (v3 % ((v45) | 1))
                    if false {
                        var v55 float32 = 0.0
                        _ = len(v2)
                        delete(v2, 2147483648)
                        v49 -= 71
                        v45 = (v45 & (v3 - (-(v45))))
                        v53 += v53
                        _ = v55
                    }
                }
            }
            if len(v46) > 0 {
                v46[0] = fi52
            }
            _ = fi52
            _ = v53
        }
        _ = fi51
    }
    delete(v2, v49)
    var v56 int = 492
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v49:", v49)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v49
    _ = v50
    _ = v56
}
