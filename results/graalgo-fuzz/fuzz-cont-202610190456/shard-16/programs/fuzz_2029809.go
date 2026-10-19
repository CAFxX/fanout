package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint64 }
func genFn0() uint32 {
    var v4 bool = false
    if (true && v4) {
        var v5 float64 = 1.5
        var v6 float64 = v5
        _ = v5
        _ = v6
    } else {
        var v7 uint8 = 128
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        if v4 {
            fmt.Println(v4, fi8)
            var v9 int32 = -698
            _ = v9
        } else {
            v10 := []int{4294967295, fi8, fi8}
            for fi11 := 0; fi11 < 2; fi11++ {
                if ((2 > 65535) || v4) {
                    v4 = v4
                } else {
                    if len(v10) > 0 {
                        v10[0] = fi11
                    }
                }
                var v12 string = ""
                var v13 int = (fi11 ^ 0)
                _ = fi11
                _ = v12
                _ = v13
            }
            switch (fi8 & fi8) {
            case 0:
                var v14 int8 = 70
                _ = v14
            case 581:
                var v15 int64 = 0
                _ = v15
            case -544:
                switch "foo bar" {
                case "foo bar":
                    var v16 uint = 255
                    _ = v16
                }
            default:
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 int16 = -763
                    fmt.Println(v18)
                    fmt.Println(v18, v4)
                    _ = fi17
                    _ = v18
                }
            }
            _ = v10
        }
        v19 := []int{fi8, (fi8 | (-(fi8)))}
        fmt.Println(v4, v4)
        _ = fi8
        _ = v19
    }
    _ = v4
    return 2147483648
}
func genFn1(p0 string, p1 int32) uint64 {
    p1 = (-(p1))
    v20 := GenS1{F0: 42}
    p1 += 2147483647
    v21 := []int{559, 32767}
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 65535
    var v2 float32 = 0.1
    var v3 uint32 = v1
    v3 += genFn0()
    var v22 bool = ((255 != 18446744073709551615) && true)
    v3 += v3
    var v23 uint32 = (v3 % ((563) | 1))
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := GenS1{F0: 127}
        switch fi24 {
        case 9223372036854775807:
            var v26 int64 = -446
            _ = v26
        case -9223372036854775808:
            var v27 float64 = 0.1
            _ = v27
        default:
            var v28 uint8 = 2
            _ = v28
        }
        var v29 int = (fi24 % (((fi24 & fi24)) | 1))
        _ = fi24
        _ = v25
        _ = v29
    }
    switch 1 {
    case -770:
        v30 := GenS1{F0: 127}
        _ = v30
    default:
        v31 := []int{2147483647, 9223372036854775807}
        _ = v31
    }
    switch "\n\t" {
    case "a":
        var v32 int8 = 100
        _ = v32
    default:
        for fi33 := 0; fi33 < 17; fi33++ {
            if (!v22) {
                for fi34 := 0; fi34 < 18; fi34++ {
                    var v35 uint16 = 810
                    v36 := []int{(fi34 / ((fi34) | 1)), (fi33 | (fi33 >> ((168) & 63)))}
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            } else {
                for fi37 := 0; fi37 < 16; fi37++ {
                    v38 := []int{fi37, fi37, 723, -808}
                    _ = fi37
                    _ = v38
                }
                for fi39 := 0; fi39 < 8; fi39++ {
                    v22 = v22
                    v40 := GenS2{F0: 227, F1: 159}
                    if (("" < ("foo bar" + "hello")) && (-86 > 42)) {
                        v3 = genFn0()
                        v41 := GenS3{F0: ""}
                        fmt.Println(v2, v1)
                        _ = v41
                    }
                    _ = fi39
                    _ = v40
                }
                v42 := GenS1{F0: (fi33 % ((-1) | 1))}
                _ = v42
            }
            _ = fi33
        }
    }
    fmt.Println(v2, v1)
    if (114 <= 65535) {
        if v22 {
            fmt.Println(v1, v3, v1)
            for fi43 := 0; fi43 < 5; fi43++ {
                fmt.Println(v23, v22)
                _ = fi43
            }
            v44 := []int{42, 32768, -9223372036854775808}
            _ = v44
        }
        switch 9223372036854775807 {
        case 32768:
            v3 -= (^(v3))
        case 4294967295:
            var v45 uint32 = v1
            _ = v45
        case 128:
            if v22 {
                v46 := GenS3{F0: "\n\t"}
                v22 = v22
                _ = v46
            } else {
                if false {
                    if false {
                        for fi47 := 0; fi47 < 17; fi47++ {
                            var v48 int64 = 1
                            _ = fi47
                            _ = v48
                        }
                        fmt.Println(v23, v3, v3)
                    } else {
                        if (32767 > 2) {
                            v49 := []int{-2147483648, 274}
                            var v50 int8 = int8(v23)
                            v1 -= 4294967295
                            _ = v49
                            _ = v50
                        }
                    }
                    var v51 float32 = v2
                    v52 := GenS2{F0: 43, F1: 0}
                    _ = v51
                    _ = v52
                }
                var v53 uint32 = (v3 / ((v1) | 1))
                _ = v53
            }
        }
    } else {
        var v54 uint16 = 255
        var v55 float64 = 3.14159
        if v22 {
            v56 := GenS4{F0: 2, F1: 2}
            v57 := []int{}
            if v22 {
                var v58 uint32 = 88
                var v59 int8 = 0
                var v60 int32 = 1
                _ = v58
                _ = v59
                _ = v60
            }
            _ = v56
            _ = v57
        }
        _ = v54
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
