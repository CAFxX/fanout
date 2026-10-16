package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 int64; F2 int16 }
type GenS6 struct { F0 int32; F1 int; F2 uint16 }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 uint8; F1 uint; F2 int8 }
func genFn0(p0 int32) uint16 {
    var v7 int8 = 100
    v8 := GenS1{F0: "hello"}
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint = 356
    var v2 int16 = 42
    var v3 int8 = 1
    v4 := []uint32{65535, 1}
    var v5 int8 = -1
    var v6 int16 = 1000
    switch "foo bar" {
    case "\n\t":
        v4 = append(v4, 65535)
    case "":
        v6 -= v6
    }
    if len(v4) > 0 {
        v4[0] = 87
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 float32 = 1.5
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int8 = -82
            if ((v6 % (((v2 & -32768)) | 1)) != v2) {
                if (42 >= 127) {
                    var v13 float32 = v10
                    _ = v13
                } else {
                    var v14 bool = (!((!true) && false))
                    _ = v14
                }
            }
            _ = fi11
            _ = v12
        }
        v4 = append(v4, 2147483647)
        _ = fi9
        _ = v10
    }
    var v15 uint = (v1 ^ (^(v1)))
    var v16 int = 32767
    var v17 bool = (v16 >= v16)
    v15 = v1
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            if (-9223372036854775808 == 1) {
                switch ("a" + "a") {
                case "\n\t":
                    v20 := GenS1{F0: "\n\t"}
                    _ = v20
                case "a":
                    switch (fi18 ^ (fi19 | (v16 & v16))) {
                    case -661:
                        switch "foo bar" {
                        case "hello":
                            var v21 float32 = 3.14159
                            _ = v21
                        default:
                            v1 = (v1 | (v1 >> ((765) & 63)))
                            if false {
                                var v22 int8 = v5
                                if (65535 != 42) {
                                    var v23 uint8 = 127
                                    v16 -= 514
                                    _ = v23
                                }
                                var v24 int64 = 0
                                _ = v22
                                _ = v24
                            } else {
                                fmt.Println(v3, v5)
                            }
                            _ = len(v4)
                        }
                    case 4294967295:
                        fmt.Println(fi18, v15)
                    default:
                        switch (("foo bar" + "a") + ("foo bar" + "\n\t")) {
                        case "hello":
                            var v25 float32 = 1e10
                            _ = v25
                        case "a":
                            var v26 int16 = (^(v2))
                            _ = v26
                        case "foo bar":
                            if v17 {
                                fmt.Println(v1, v3, v6)
                                v27 := GenS1{F0: ("a" + ("\n\t" + "foo bar"))}
                                var v28 uint64 = 677
                                _ = v27
                                _ = v28
                            } else {
                                v1 = v1
                            }
                        }
                    }
                    v29 := GenS1{F0: "a"}
                    _ = v29
                case "":
                    var v30 uint32 = 550
                    _ = v30
                }
                if v17 {
                    var v31 uint = v15
                    var v32 int16 = (v6 << (((v31 % (((v15 & (^(v1)))) | 1))) & 63))
                    var v33 uint64 = 1
                    _ = v31
                    _ = v32
                    _ = v33
                } else {
                    var v34 uint64 = 572
                    if (!((v15 - v15) != v15)) {
                        v17 = true
                    } else {
                        switch (fi19 >> (((v1 / ((v15) | 1))) & 63)) {
                        case 365:
                            v5 += (v3 << ((459) & 63))
                        default:
                            v15 -= (v15 << (((^(v1))) & 63))
                        }
                    }
                    _ = v34
                }
                fmt.Println(v17, v5, v16)
            } else {
                if (v17 || false) {
                    if len(v4) > 0 {
                        v4[0] = 418
                    }
                } else {
                    switch (^(fi18)) {
                    case 128:
                        v15 += v1
                    case 1:
                        v4 = append(v4, 65535)
                    default:
                        var v35 int64 = 358
                        _ = v35
                    }
                    if true {
                        switch (("hello" + ("foo bar" + ("foo bar" + ("hello" + "a")))) + "foo bar") {
                        case "hello":
                            var v36 bool = (!(954 == 868))
                            _ = v36
                        case "a":
                            v17 = (v3 >= 127)
                        default:
                            if (!(false && (-2147483648 <= -1))) {
                                if v17 {
                                    v6 = (v6 % ((v2) | 1))
                                    v5 -= v3
                                } else {
                                    if (42 <= -9223372036854775808) {
                                        if (148 != -9223372036854775808) {
                                            if v17 {
                                                if (v17 || (v3 <= v3)) {
                                                    var v37 uint16 = 211
                                                    _ = v37
                                                } else {
                                                    var v38 uint16 = 255
                                                    _ = v38
                                                }
                                                fmt.Println(v16)
                                                fmt.Println(fi19, v5)
                                                v39 := GenS2{F0: 32767, F1: 241}
                                                v40 := GenS3{F0: 2}
                                                v41 := GenS1{F0: "hello"}
                                                _ = v39
                                                _ = v40
                                                _ = v41
                                            } else {
                                                v42 := GenS4{F0: 0}
                                                if (191 != 1) {
                                                    v2 = -458
                                                    var v43 uint16 = 777
                                                    v15 = (^(v15))
                                                    v4 = append(v4, 4294967295)
                                                    _ = v43
                                                } else {
                                                    if v17 {
                                                        fmt.Println(v16, v15, v2)
                                                        v44 := GenS5{F0: (v6 % ((v2) | 1)), F1: 1, F2: v2}
                                                        _ = v44
                                                    }
                                                }
                                                var v45 int64 = 661
                                                _ = v42
                                                _ = v45
                                            }
                                            v17 = v17
                                            v46 := GenS6{F0: -381, F1: (v16 & 32767), F2: 83}
                                            _ = v46
                                        }
                                    } else {
                                        v5 = v3
                                    }
                                    v47 := GenS1{F0: "a"}
                                    _ = v47
                                }
                                _ = len(v4)
                            } else {
                                var v48 int8 = v3
                                v48 = v3
                                v15 += (v15 ^ (v15 ^ 65535))
                                fmt.Println(v17, v3, v48)
                                _ = v48
                            }
                        }
                    }
                    v49 := GenS7{F0: v17, F1: (fi18 & (fi19 + fi19))}
                    _ = v49
                }
                fmt.Println(v5, fi18, v1)
            }
            if len(v4) > 0 {
                v4[0] = 611
            }
            _ = fi19
        }
        var v50 uint8 = 42
        var v51 uint32 = 973
        _ = fi18
        _ = v50
        _ = v51
    }
    v5 = (v3 % ((v3) | 1))
    for fi52 := 0; fi52 < 5; fi52++ {
        switch fi52 {
        case 9223372036854775807:
            v17 = true
        case 128:
            v6 = (v2 - v6)
        case 4294967295:
            v2 -= (v2 * (v2 & (-(v2))))
        }
        switch fi52 {
        case 4294967295:
            switch v16 {
            case 1:
                fmt.Println(v16, v6, fi52)
            case 127:
                if v17 {
                    switch (v16 % (((fi52 << ((451) & 63))) | 1)) {
                    case 32768:
                        v6 -= (v6 - (-(v2)))
                    case 2147483647:
                        v53 := GenS6{F0: -2147483648, F1: v16, F2: 32767}
                        _ = v53
                    case -2147483648:
                        var v54 bool = (!true)
                        _ = v54
                    }
                    var v55 uint32 = 2
                    _ = v55
                } else {
                    if true {
                        _ = len(v4)
                    } else {
                        var v56 uint16 = 255
                        if (!v17) {
                            var v57 int8 = v5
                            if v17 {
                                for fi58 := 0; fi58 < 8; fi58++ {
                                    var v59 int32 = 934
                                    _ = fi58
                                    _ = v59
                                }
                                if len(v4) > 0 {
                                    v4[0] = 2147483647
                                }
                                for fi60 := 0; fi60 < 17; fi60++ {
                                    if true {
                                        v61 := GenS8{F0: uint8(v56), F1: v15, F2: 127}
                                        _ = v61
                                    } else {
                                        _ = len(v4)
                                        _ = len(v4)
                                        fmt.Println(fi60)
                                    }
                                    _ = fi60
                                }
                            }
                            for fi62 := 0; fi62 < 15; fi62++ {
                                var v63 int64 = 0
                                fmt.Println(v3, v56)
                                _ = fi62
                                _ = v63
                            }
                            _ = v57
                        }
                        _ = v56
                    }
                    v1 = v1
                    var v64 float32 = -2.25
                    _ = v64
                }
            default:
                v3 -= (v3 << ((349) & 63))
            }
        default:
            if len(v4) > 0 {
                v4[0] = 1
            }
        }
        var v65 uint32 = 0
        _ = fi52
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
