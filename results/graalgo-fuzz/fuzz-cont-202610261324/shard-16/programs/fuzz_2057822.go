package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint }
func genFn0() int16 {
    var v5 int16 = -468
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: 42, F1: fi6}
        var v8 uint32 = 793
        if true {
            if (false && ((fi6 << ((42) & 63)) > fi6)) {
                var v9 int16 = (v5 / ((v5) | 1))
                switch fi6 {
                case -1:
                    v10 := GenS2{F0: (898 == 32767)}
                    _ = v10
                case -833:
                    var v11 int16 = (v9 % (((v9 / (((v9 * (v9 << ((670) & 63)))) | 1))) | 1))
                    _ = v11
                default:
                    for fi12 := 0; fi12 < 15; fi12++ {
                        var v13 uint8 = 255
                        v9 += v9
                        if false {
                            var v14 int16 = 42
                            var v15 int64 = 1
                            var v16 uint8 = v13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                        v17 := GenS1{F0: 0, F1: (fi6 >> ((685) & 63))}
                        _ = fi12
                        _ = v13
                        _ = v17
                    }
                }
                v18 := []int{-2147483648}
                _ = v18
                _ = v9
            }
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v5)
    _ = v5
    return v5
}
func genFn1(p0 int8, p1 int16, p2 int64) uint16 {
    var v19 string = ("hello" + (("a" + ("foo bar" + "foo bar")) + ""))
    v20 := GenS2{F0: false}
    switch (v19 + ((v19 + "hello") + v19)) {
    case "a":
        p0 = (p0 >> ((2) & 63))
    }
    var v21 bool = true
    switch 127 {
    case -9223372036854775808:
        v22 := []int{9223372036854775807}
        _ = v22
    case 32768:
        switch 128 {
        case 127:
            for fi23 := 0; fi23 < 17; fi23++ {
                var v24 int16 = 1000
                var v25 float32 = 1e10
                _ = fi23
                _ = v24
                _ = v25
            }
        }
    case -544:
        fmt.Println(p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    _ = v21
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    var v3 uint64 = 0
    v4 := []uint32{2, 2147483648, 633, 2}
    if (-128 < 42) {
        var v26 int8 = 100
        delete(v2, 65535)
        _ = v26
    }
    fmt.Println(v3)
    v27 := GenS3{F0: 311, F1: 4294967295}
    v28 := GenS2{F0: (true && false)}
    for fi29 := 0; fi29 < 10; fi29++ {
        v30 := v2[0]
        v31 := GenS4{F0: ""}
        var v32 int = v30
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    var v33 int16 = 42
    if (false && true) {
        delete(v1, "a")
        v33 -= v33
    } else {
        v34 := GenS5{F0: 878}
        for fi35 := 0; fi35 < 10; fi35++ {
            switch fi35 {
            case 2147483647:
                switch 128 {
                case 32767:
                    for fi36 := 0; fi36 < 4; fi36++ {
                        _ = len(v1)
                        _ = fi36
                    }
                default:
                    if false {
                        if false {
                            for fi37 := 0; fi37 < 1; fi37++ {
                                switch ("a" + "\n\t") {
                                case "\n\t":
                                    switch -597 {
                                    case 9223372036854775807:
                                        var v38 int = fi37
                                        _ = v38
                                    case 521:
                                        fmt.Println(fi37, fi35)
                                    case -2147483648:
                                        fmt.Println(fi37)
                                    default:
                                        fmt.Println(fi37, fi37)
                                    }
                                default:
                                    var v39 int16 = v33
                                    _ = v39
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = fi37
                            }
                        }
                        var v40 uint = 42
                        _ = v40
                    } else {
                        delete(v1, "a")
                    }
                }
            }
            if (!(0 <= 127)) {
                v3 += v3
                for fi41 := 0; fi41 < 17; fi41++ {
                    v42 := v2[181]
                    _ = fi41
                    _ = v42
                }
            } else {
                v3 = (v3 & 2)
                v33 -= -1
            }
            _ = fi35
        }
        fmt.Println(v3, v33)
        _ = v34
    }
    for fi43 := 0; fi43 < 8; fi43++ {
        fmt.Println(v33, v33, v3)
        v33 += v33
        _ = len(v2)
        _ = fi43
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        var v45 int64 = 9223372036854775807
        v46 := v1[""]
        if true {
            var v47 bool = (!true)
            for fi48 := 0; fi48 < 11; fi48++ {
                var v49 float32 = (-2.25 * (((float32(0.1) * 1e10) - 3.14159) * 1e10))
                fmt.Println(fi48, v47, v33)
                _ = fi48
                _ = v49
            }
            _ = v47
        }
        _ = fi44
        _ = v45
        _ = v46
    }
    var v50 uint8 = 127
    fmt.Println(v33, v3, v3)
    v51 := GenS6{F0: 0}
    v52 := v1[""]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v50
    _ = v51
    _ = v52
}
