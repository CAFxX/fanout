package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint32 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int16 = -1
    v3 := []int{36, 143, -2147483648, 4294967295}
    v2 -= v2
    v4 := GenS1{F0: (^(v1)), F1: -1}
    if (!(!(100 != 42))) {
        var v5 int64 = 9223372036854775807
        _ = v5
    }
    v2 = v2
    var v6 string = (("hello" + ("\n\t" + "hello")) + "foo bar")
    switch 1 {
    case -671:
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 float32 = 1e10
            _ = fi7
            _ = v8
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 int8 = -66
        fmt.Println(v1, v6, v1)
        switch (fi9 << ((660) & 63)) {
        case 9223372036854775807:
            if true {
                if len(v3) > 0 {
                    v3[0] = (-(fi9))
                }
                v3 = append(v3, (-(fi9)))
            }
        case -9223372036854775808:
            v3 = append(v3, fi9)
        case 2147483647:
            if true {
                if false {
                    fmt.Println(v1, v1)
                    for fi11 := 0; fi11 < 9; fi11++ {
                        var v12 string = v6
                        _ = fi11
                        _ = v12
                    }
                    v13 := GenS2{F0: 0}
                    _ = v13
                } else {
                    switch v6 {
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = -9223372036854775808
                        }
                    }
                    var v14 uint32 = 136
                    v15 := GenS3{F0: -128}
                    _ = v14
                    _ = v15
                }
                fmt.Println(v10)
                _ = len(v3)
            }
        default:
            v2 = (-(v2))
        }
        _ = fi9
        _ = v10
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        v6 = (("foo bar" + (v6 + (v6 + v6))) + v6)
        _ = fi16
    }
    var v17 uint8 = 255
    if true {
        v1 += v1
        v2 = (v2 + (v2 + (-(v2))))
    } else {
        v6 = "hello"
        fmt.Println(v1, v1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 17; fi19++ {
            switch fi18 {
            case 1:
                if (!((4294967295 <= v1) && false)) {
                    var v20 uint64 = (v1 + v1)
                    if false {
                        if (!(65535 > 2147483648)) {
                            var v21 float32 = 1.5
                            var v22 int8 = 1
                            var v23 uint64 = 4294967295
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        switch 4294967295 {
                        case 42:
                            if len(v3) > 0 {
                                v3[0] = 151
                            }
                        case 32768:
                            switch v6 {
                            case "":
                                v24 := GenS4{F0: 32767, F1: 32767, F2: 807}
                                _ = v24
                            case "a":
                                if (0 < 65535) {
                                    v17 = uint8(v1)
                                    if len(v3) > 0 {
                                        v3[0] = fi19
                                    }
                                }
                            }
                        case 2147483647:
                            if false {
                                if false {
                                    v2 += (v2 - v2)
                                    v3 = append(v3, (fi18 % ((fi18) | 1)))
                                } else {
                                    v3 = append(v3, fi18)
                                }
                            }
                            _ = len(v3)
                        }
                        v25 := GenS1{F0: (v20 ^ (v1 >> ((12) & 63))), F1: 40}
                        _ = v25
                    }
                    if true {
                        v26 := GenS4{F0: 65535, F1: 1, F2: 42}
                        var v27 int = fi19
                        _ = v26
                        _ = v27
                    } else {
                        v6 = v6
                        if (true || false) {
                            fmt.Println(v1, fi18)
                        }
                    }
                    _ = v20
                } else {
                    v28 := GenS1{F0: (^(v1)), F1: 42}
                    _ = v28
                }
            default:
                v17 = v17
            }
            _ = fi19
        }
        if true {
            var v29 string = ""
            var v30 int32 = 994
            if ((v2 << ((0) & 63)) <= (-(v2))) {
                _ = len(v3)
                v6 = v29
                v30 += v30
            } else {
                var v31 int16 = -796
                var v32 uint16 = uint16(v1)
                var v33 int32 = v30
                _ = v31
                _ = v32
                _ = v33
            }
            _ = v29
            _ = v30
        } else {
            v17 -= (^(v17))
            for fi34 := 0; fi34 < 13; fi34++ {
                if true {
                    v35 := GenS5{F0: -128, F1: 2}
                    _ = v35
                }
                _ = fi34
            }
        }
        v36 := GenS1{F0: v1, F1: 42}
        _ = fi18
        _ = v36
    }
    if ("hello" != (v6 + ((("foo bar" + (("foo bar" + ((v6 + (v6 + v6)) + ((v6 + "a") + v6))) + (v6 + (v6 + ("foo bar" + v6))))) + "hello") + ("" + "foo bar")))) {
        v37 := GenS3{F0: 26}
        v2 = -1
        if (!((535 <= 42) && (1 > 1))) {
            v38 := GenS1{F0: 0, F1: 0}
            _ = v38
        }
        _ = v37
    } else {
        for fi39 := 0; fi39 < 10; fi39++ {
            if false {
                fmt.Println(v1, v1)
                fmt.Println(v17, v1)
                v2 += 550
            } else {
                _ = len(v3)
            }
            _ = fi39
        }
        var v40 uint16 = 0
        var v41 uint8 = (v17 * 2)
        _ = v40
        _ = v41
    }
    v42 := GenS6{F0: (^(v17)), F1: v1, F2: uint32(v1)}
    var v43 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v43:", v43)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v6
}
