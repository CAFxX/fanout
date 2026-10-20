package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int32; F2 int16 }
type GenS5 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint16) uint16 {
    v5 := GenS1{F0: 32767}
    fmt.Println(p0)
    v6 := GenS2{F0: 127, F1: "\n\t"}
    p0 = p0
    _ = p0
    _ = v5
    _ = v6
    return p0
}
func genFn1(p0 uint8) int32 {
    var v7 int8 = 0
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int8 = 31
        var v10 int16 = -177
        _ = fi8
        _ = v10
        _ = v9
    }
    v7 += (v7 << ((18446744073709551615) & 63))
    v7 = 127
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 string = "\n\t"
        switch ((v12 + "a") + "hello") {
        case "foo bar":
            v7 += v7
        case "hello":
            v13 := GenS1{F0: -750}
            _ = v13
        default:
            v7 = -49
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v7
    return 0
}

func main() {
    v1 := []uint32{2147483647, 966, 2}
    var v2 int64 = -1
    var v3 int32 = -2147483648
    var v4 int64 = v2
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 uint32 = 238
            v17 := GenS1{F0: -32768}
            _ = fi15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            v4 = (v4 & v2)
            _ = fi18
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi14
    }
    var v19 uint16 = 500
    v4 -= (v4 % ((v2) | 1))
    if false {
        v20 := GenS2{F0: 71, F1: (("hello" + "\n\t") + ((("" + "") + "") + "hello"))}
        v1 = append(v1, 851)
        _ = v20
    } else {
        var v21 string = ""
        v4 -= (v2 * v4)
        var v22 int16 = 1
        _ = v21
        _ = v22
    }
    switch "" {
    case "hello":
        var v23 uint = 255
        _ = v23
    case "a":
        v24 := GenS3{F0: ""}
        _ = v24
    case "":
        v2 = (v4 >> ((65535) & 63))
    default:
        var v25 bool = (29 > 32768)
        _ = v25
    }
    v3 = v3
    if false {
        var v26 int = 9223372036854775807
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v27 := GenS3{F0: ""}
        _ = v26
        _ = v27
    }
    fmt.Println(v2)
    var v28 int64 = -1
    v29 := GenS4{F0: 481, F1: -1, F2: 0}
    v30 := GenS5{F0: 1, F1: 38}
    switch ("foo bar" + "foo bar") {
    case "a":
        if false {
            switch 128 {
            case -1:
                if (false || (!false)) {
                    var v31 string = ("\n\t" + ((((((("\n\t" + "a") + "\n\t") + "a") + "a") + "hello") + "hello") + (("a" + "hello") + (("" + (("a" + "a") + "foo bar")) + ("foo bar" + "")))))
                    _ = v31
                } else {
                    if true {
                        if (!(100 == 42)) {
                            if ("a" <= (("a" + "foo bar") + ("" + (("\n\t" + "hello") + "\n\t")))) {
                                v32 := GenS2{F0: 0, F1: "hello"}
                                var v33 uint8 = 127
                                _ = v32
                                _ = v33
                            }
                            var v34 uint32 = 42
                            for fi35 := 0; fi35 < 2; fi35++ {
                                v4 -= -351
                                fmt.Println(v19, fi35, v3)
                                _ = fi35
                            }
                            _ = v34
                        } else {
                            for fi36 := 0; fi36 < 18; fi36++ {
                                switch "\n\t" {
                                case "hello":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case "\n\t":
                                    var v37 int32 = (v3 << ((841) & 63))
                                    _ = v37
                                }
                                _ = fi36
                            }
                            v19 = (v19 | v19)
                        }
                        if true {
                            v3 -= (^(v3))
                        }
                        v3 += (v3 * v3)
                    }
                    var v38 string = ("foo bar" + "hello")
                    if (246 != 2) {
                        if (128 >= 0) {
                            v19 = (^(v19))
                            var v39 uint16 = 852
                            _ = v39
                        }
                    }
                    _ = v38
                }
            default:
                switch 9223372036854775807 {
                case 128:
                    var v40 int32 = (v3 - (v3 % ((v3) | 1)))
                    _ = v40
                }
            }
        } else {
            if (!(false && false)) {
                var v41 int = 42
                if true {
                    fmt.Println(v41, v3)
                    v1 = append(v1, 356)
                }
                v42 := GenS1{F0: 32767}
                _ = v41
                _ = v42
            }
        }
    default:
        if true {
            var v43 uint = 18446744073709551615
            for fi44 := 0; fi44 < 20; fi44++ {
                switch (fi44 % ((fi44) | 1)) {
                case -994:
                    _ = len(v1)
                }
                fmt.Println(v3, fi44)
                _ = fi44
            }
            _ = len(v1)
            _ = v43
        }
    }
    v28 = -9223372036854775808
    v3 = (v3 + (v3 >> ((65535) & 63)))
    if true {
        switch 643 {
        case -9:
            if false {
                fmt.Println(v19)
            }
        default:
            v45 := GenS4{F0: 979, F1: v3, F2: -926}
            _ = v45
        }
        v19 = 916
        if ((!true) && (42 != 255)) {
            v2 = (v4 % (((v28 & v2)) | 1))
            var v46 uint = 18446744073709551615
            for fi47 := 0; fi47 < 14; fi47++ {
                var v48 uint16 = (v19 - genFn0(v19))
                if len(v1) > 0 {
                    v1[0] = 19
                }
                var v49 int = fi47
                _ = fi47
                _ = v48
                _ = v49
            }
            _ = v46
        } else {
            if false {
                for fi50 := 0; fi50 < 6; fi50++ {
                    for fi51 := 0; fi51 < 9; fi51++ {
                        var v52 bool = false
                        var v53 float32 = 1e10
                        var v54 int8 = -19
                        _ = fi51
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                    v28 += (v2 - 0)
                    var v55 uint64 = 2
                    _ = fi50
                    _ = v55
                }
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                if ("hello" == (("hello" + "a") + "\n\t")) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            var v56 float32 = float32(float64((1.5 + (0.0 / 1e10))))
            var v57 int8 = 100
            _ = v56
            _ = v57
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
