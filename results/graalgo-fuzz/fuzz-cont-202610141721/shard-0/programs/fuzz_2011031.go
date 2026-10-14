package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int16; F1 uint8; F2 int16 }
type GenS4 struct { F0 int32; F1 uint16; F2 bool }

func main() {
    var v1 float64 = (0.0 + 1.5)
    var v2 int8 = 127
    var v3 uint16 = 2
    v4 := []uint32{53, 403, 2147483648}
    var v5 int64 = 171
    v5 = v5
    v6 := GenS1{F0: (v2 >> ((952) & 63))}
    _ = len(v4)
    var v7 float64 = 3.14159
    v1 += v7
    v3 -= (^(v3))
    switch int(v2) {
    case -90:
        v2 -= v2
    case 42:
        if false {
            if len(v4) > 0 {
                v4[0] = 2
            }
            v8 := GenS2{F0: "\n\t", F1: -77}
            v5 += (v5 & v5)
            _ = v8
        }
    default:
        v9 := GenS1{F0: 1}
        _ = v9
    }
    v5 = 1
    v10 := GenS1{F0: v2}
    if true {
        if (true && (false && true)) {
            if (65535 == 339) {
                for fi11 := 0; fi11 < 11; fi11++ {
                    for fi12 := 0; fi12 < 4; fi12++ {
                        v1 = 1.5
                        v5 -= (v5 << ((uint(v3)) & 63))
                        _ = fi12
                    }
                    _ = fi11
                }
                if true {
                    if (!true) {
                        for fi13 := 0; fi13 < 9; fi13++ {
                            var v14 int16 = -32768
                            v15 := GenS1{F0: -106}
                            v16 := GenS1{F0: (^(v2))}
                            _ = fi13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                        v3 = (v3 * (v3 % ((1) | 1)))
                    } else {
                        for fi17 := 0; fi17 < 7; fi17++ {
                            var v18 uint64 = 2
                            var v19 int32 = int32(fi17)
                            if true {
                                var v20 uint16 = v3
                                _ = v20
                            }
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        switch 808 {
                        case 32767:
                            v21 := GenS3{F0: -308, F1: 128, F2: -1}
                            _ = v21
                        case 127:
                            var v22 uint8 = 255
                            _ = v22
                        }
                    }
                    for fi23 := 0; fi23 < 10; fi23++ {
                        var v24 bool = (fi23 < fi23)
                        var v25 float64 = (v1 / 1e10)
                        for fi26 := 0; fi26 < 11; fi26++ {
                            v4 = append(v4, 4294967295)
                            _ = fi26
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
            } else {
                var v27 int32 = -2
                var v28 int = -1
                _ = v27
                _ = v28
            }
        }
    } else {
        for fi29 := 0; fi29 < 6; fi29++ {
            v5 -= (v5 | v5)
            _ = fi29
        }
        if (true || (127 < 2)) {
            if len(v4) > 0 {
                v4[0] = 1
            }
            for fi30 := 0; fi30 < 19; fi30++ {
                switch "foo bar" {
                case "hello":
                    switch -115 {
                    case 32768:
                        var v31 uint8 = 2
                        _ = v31
                    default:
                        fmt.Println(fi30)
                    }
                }
                v3 = (v3 * (^(v3)))
                var v32 uint64 = 0
                _ = fi30
                _ = v32
            }
        } else {
            var v33 uint64 = 4294967295
            for fi34 := 0; fi34 < 19; fi34++ {
                if len(v4) > 0 {
                    v4[0] = 0
                }
                if len(v4) > 0 {
                    v4[0] = 401
                }
                _ = fi34
            }
            _ = v33
        }
        if (42 == 4294967295) {
            v1 -= ((3.14159 * -2.25) - 3.14159)
            v2 = v2
        } else {
            var v35 int16 = 1
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        var v37 string = "hello"
        if (65535 == 1) {
            var v38 int = fi36
            _ = v38
        } else {
            switch (v37 + (v37 + v37)) {
            case "":
                fmt.Println(v2, fi36, fi36)
            default:
                v39 := GenS1{F0: -128}
                _ = v39
            }
            v3 -= v3
        }
        if ((v3 ^ v3) < (v3 / ((v3) | 1))) {
            for fi40 := 0; fi40 < 6; fi40++ {
                fmt.Println(v37, v5, v5)
                if true {
                    var v41 float32 = ((0.1 - 1e10) - -2.25)
                    v42 := GenS4{F0: 1, F1: 800, F2: true}
                    _ = v41
                    _ = v42
                }
                _ = len(v4)
                _ = fi40
            }
        }
        _ = fi36
        _ = v37
    }
    if false {
        if (-128 > 0) {
            var v43 uint = 478
            _ = v43
        }
        var v44 int64 = v5
        _ = v44
    } else {
        fmt.Println(v2, v1)
    }
    var v45 int8 = 77
    if (0 > 0) {
        for fi46 := 0; fi46 < 20; fi46++ {
            if false {
                v47 := GenS4{F0: 2147483647, F1: (v3 - v3), F2: false}
                for fi48 := 0; fi48 < 16; fi48++ {
                    fmt.Println(v2)
                    v4 = append(v4, 65535)
                    if false {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v49 uint64 = 1
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v49
                    }
                    _ = fi48
                }
                for fi50 := 0; fi50 < 5; fi50++ {
                    v5 = -200
                    _ = fi50
                }
                _ = v47
            } else {
                var v51 uint64 = 42
                _ = v51
            }
            fmt.Println(v1)
            v52 := GenS1{F0: (v2 << ((255) & 63))}
            _ = fi46
            _ = v52
        }
        var v53 int64 = (v5 << ((2) & 63))
        _ = v53
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
