package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 int8 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 int; F2 int64 }

func main() {
    v1 := []int{9223372036854775807}
    v2 := []uint32{207, 42, 63}
    var v3 int16 = 42
    for fi4 := 0; fi4 < 15; fi4++ {
        var v5 uint64 = 1
        _ = fi4
        _ = v5
    }
    v2 = append(v2, 65535)
    if (!false) {
        fmt.Println(v3)
        var v6 uint16 = 2
        for fi7 := 0; fi7 < 20; fi7++ {
            if false {
                v8 := GenS1{F0: 127, F1: -759, F2: -982}
                fmt.Println(v3)
                _ = v8
            } else {
                var v9 bool = ((0 > 9223372036854775807) && ((!false) && (!true)))
                if ((810 != 0) && (86 <= uint(v6))) {
                    v9 = (("a" + "\n\t") >= (("a" + (("hello" + (((("foo bar" + ("foo bar" + "foo bar")) + "foo bar") + "foo bar") + "hello")) + "foo bar")) + "foo bar"))
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                }
                _ = v9
            }
            _ = fi7
        }
        _ = v6
    } else {
        if true {
            v3 -= v3
            if true {
                for fi10 := 0; fi10 < 11; fi10++ {
                    switch fi10 {
                    case 42:
                        var v11 int64 = 1
                        _ = v11
                    case 9223372036854775807:
                        if (65535 <= 752) {
                            if true {
                                switch (fi10 >> ((1) & 63)) {
                                case 32767:
                                    v3 = -1
                                case -121:
                                    if (true && true) {
                                        var v12 bool = (1e10 <= (1.5 + 3.14159))
                                        _ = v12
                                    }
                                case 353:
                                    var v13 int8 = 127
                                    _ = v13
                                }
                            } else {
                                for fi14 := 0; fi14 < 13; fi14++ {
                                    v15 := GenS1{F0: 0, F1: -1, F2: -187}
                                    _ = fi14
                                    _ = v15
                                }
                                var v16 bool = (-87 <= 0)
                                _ = v16
                            }
                        } else {
                            for fi17 := 0; fi17 < 12; fi17++ {
                                var v18 string = ("\n\t" + "hello")
                                _ = fi17
                                _ = v18
                            }
                        }
                    case 127:
                        for fi19 := 0; fi19 < 20; fi19++ {
                            if false {
                                var v20 int8 = 0
                                var v21 int16 = (-(v3))
                                _ = v20
                                _ = v21
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v22 uint32 = 739
                                if (!(true || true)) {
                                    var v23 int64 = 1
                                    _ = v23
                                } else {
                                    var v24 float32 = 0.0
                                    switch "hello" {
                                    case "foo bar":
                                        switch ("" + "") {
                                        case "a":
                                            v22 = (v22 * (v22 ^ (v22 << ((65535) & 63))))
                                            var v25 int32 = -1
                                            _ = v25
                                        case "":
                                            v1 = append(v1, fi19)
                                            var v26 uint = 190
                                            v27 := GenS2{F0: 847, F1: 4294967295, F2: -128}
                                            switch "foo bar" {
                                            case "":
                                                switch "foo bar" {
                                                case "":
                                                    var v28 uint = v26
                                                    _ = v28
                                                case "hello":
                                                    fmt.Println(v22, v3)
                                                }
                                            case "\n\t":
                                                v29 := GenS3{F0: false, F1: fi10}
                                                _ = v29
                                            }
                                            _ = v26
                                            _ = v27
                                        case "\n\t":
                                            var v30 int64 = -846
                                            _ = v30
                                        }
                                    default:
                                        v22 -= 855
                                    }
                                    fmt.Println(fi19)
                                    _ = v24
                                }
                                v31 := GenS2{F0: 18446744073709551615, F1: 804, F2: 0}
                                v3 = (v3 >> ((706) & 63))
                                _ = v22
                                _ = v31
                            }
                            _ = fi19
                        }
                    default:
                        var v32 int8 = 1
                        _ = v32
                    }
                    v3 += 32767
                    for fi33 := 0; fi33 < 19; fi33++ {
                        v3 += (v3 % ((v3) | 1))
                        fmt.Println(v3)
                        _ = len(v2)
                        var v34 uint = 18446744073709551615
                        _ = fi33
                        _ = v34
                    }
                    _ = fi10
                }
                switch 870 {
                case 128:
                    var v35 int64 = 0
                    _ = v35
                }
            } else {
                v36 := GenS4{F0: 193}
                fmt.Println(v3)
                var v37 int64 = -273
                _ = v36
                _ = v37
            }
            if true {
                var v38 float64 = 1.5
                fmt.Println(v3, v3)
                _ = v38
            }
        } else {
            if true {
                switch 621 {
                case 1:
                    var v39 uint = 0
                    _ = v39
                case 9223372036854775807:
                    v3 += (v3 << ((751) & 63))
                case 32768:
                    v3 -= -879
                }
            } else {
                fmt.Println(v3)
                v40 := GenS3{F0: (false && (!true)), F1: -1}
                _ = v40
            }
            fmt.Println(v3)
            v3 = (-(v3))
        }
    }
    if len(v2) > 0 {
        v2[0] = 2
    }
    var v41 int = 32768
    if true {
        var v42 uint32 = 1
        _ = v42
    } else {
        v41 -= (v41 << ((65535) & 63))
        var v43 int16 = v3
        _ = v43
    }
    _ = len(v2)
    var v44 uint8 = 42
    v45 := GenS5{F0: int64(v41)}
    v46 := GenS6{F0: -1, F1: 128, F2: -1}
    if false {
        v47 := GenS5{F0: 0}
        for fi48 := 0; fi48 < 11; fi48++ {
            var v49 int64 = 9223372036854775807
            _ = fi48
            _ = v49
        }
        _ = v47
    }
    v44 = v44
    if true {
        v3 -= -1
    } else {
        v44 -= (v44 << ((97) & 63))
    }
    var v50 float64 = 1.5
    for fi51 := 0; fi51 < 3; fi51++ {
        for fi52 := 0; fi52 < 17; fi52++ {
            v3 -= -1
            if len(v1) > 0 {
                v1[0] = fi51
            }
            _ = fi52
        }
        for fi53 := 0; fi53 < 6; fi53++ {
            fmt.Println(v50, v41, v41)
            var v54 uint32 = 4294967295
            _ = fi53
            _ = v54
        }
        v55 := GenS2{F0: 42, F1: 42, F2: 42}
        _ = fi51
        _ = v55
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v44
    _ = v45
    _ = v46
    _ = v50
}
