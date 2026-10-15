package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 string; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint8, p1 int8, p2 uint8) bool {
    v6 := []int{}
    var v7 uint32 = uint32(p0)
    var v8 uint = 65535
    var v9 uint16 = 255
    v9 += uint16(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (0.1 < -2.25)
}
func genFn1(p0 bool) int8 {
    p0 = true
    if true {
        var v10 uint16 = 0
        v10 = 596
        _ = v10
    }
    p0 = true
    if (!(968 == 42)) {
        var v11 uint16 = 2
        for fi12 := 0; fi12 < 12; fi12++ {
            if ((28 != -1) && true) {
                v13 := []int{fi12, 35, fi12}
                switch "foo bar" {
                case "":
                    v11 += 32767
                case "a":
                    v14 := GenS1{F0: -128, F1: 42}
                    _ = v14
                case "foo bar":
                    v15 := GenS1{F0: -93, F1: (fi12 * int(v11))}
                    _ = v15
                default:
                    var v16 float32 = 1e10
                    _ = v16
                }
                _ = v13
            } else {
                v17 := []int{fi12, (fi12 % (((fi12 % ((-423) | 1))) | 1)), fi12}
                _ = v17
            }
            switch ("foo bar" + "foo bar") {
            case "a":
                v11 += 42
            case "":
                v18 := []int{9223372036854775807}
                _ = v18
            default:
                p0 = (2147483647 >= 65535)
            }
            v19 := []int{-1, fi12}
            _ = fi12
            _ = v19
        }
        if true {
            v20 := []int{472, 32767, 1, 942}
            for fi21 := 0; fi21 < 3; fi21++ {
                var v22 int64 = 1
                _ = fi21
                _ = v22
            }
            _ = v20
        }
        _ = v11
    } else {
        p0 = (!true)
        var v23 uint8 = 219
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 int64 = -9223372036854775808
            v26 := GenS1{F0: 27, F1: fi24}
            for fi27 := 0; fi27 < 10; fi27++ {
                v25 = 1
                _ = fi27
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = v23
    }
    var v28 uint = 2
    _ = p0
    _ = v28
    return 0
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]string)
    v3 := []string{(("\n\t" + ("foo bar" + "\n\t")) + "a"), ("foo bar" + "foo bar"), "hello"}
    v4 := make(map[int]int)
    var v5 int32 = 0
    var v29 int8 = 93
    if false {
        var v30 int64 = 42
        _ = v30
    } else {
        v31 := v2[""]
        if (v29 > genFn1((929 < 0))) {
            v31 = (v31 + ("hello" + ""))
        }
        v29 -= int8(v1)
        _ = v31
    }
    var v32 uint = 859
    var v33 bool = ((42 == -950) && (false && (0.0 == ((float64((1e10 - -2.25)) / 1.0) * -2.25))))
    v34 := GenS2{F0: 2, F1: (v5 - (v5 ^ (v5 ^ v5)))}
    for fi35 := 0; fi35 < 15; fi35++ {
        if v33 {
            var v36 uint64 = v1
            _ = v36
        } else {
            v37 := GenS3{F0: -32768}
            if v33 {
                v38 := GenS2{F0: 328, F1: v5}
                switch fi35 {
                case -9223372036854775808:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 4294967295:
                    if v33 {
                        var v39 float64 = 1e10
                        var v40 int16 = 42
                        _ = v39
                        _ = v40
                    } else {
                        v29 += 42
                        v41 := GenS3{F0: 483}
                        if false {
                            var v42 string = "foo bar"
                            v1 -= 4294967295
                            switch v42 {
                            case "hello":
                                if v33 {
                                    if (!(-154 < 1)) {
                                        v43 := GenS3{F0: -32768}
                                        v44 := GenS4{F0: 42, F1: "", F2: 232}
                                        _ = v43
                                        _ = v44
                                    }
                                }
                            case "foo bar":
                                v32 += (v32 >> (((v32 * v32)) & 63))
                            default:
                                var v45 int32 = v5
                                _ = v45
                            }
                            _ = v42
                        }
                        _ = v41
                    }
                case 32767:
                    v33 = true
                default:
                    v1 = (v1 / (((v1 * v1)) | 1))
                }
                v33 = (v33 && v33)
                _ = v38
            }
            _ = v37
        }
        v1 = (v1 | (v1 << (((^(v32))) & 63)))
        _ = fi35
    }
    for fi46 := 0; fi46 < 8; fi46++ {
        v47 := GenS4{F0: 238, F1: ("a" + (("\n\t" + "a") + (("\n\t" + "foo bar") + "\n\t"))), F2: 255}
        v48 := GenS2{F0: 2147483648, F1: (v5 + (v5 | (^(v5))))}
        _ = fi46
        _ = v47
        _ = v48
    }
    var v49 bool = false
    fmt.Println(v32, v49)
    _ = len(v3)
    for fi50 := 0; fi50 < 2; fi50++ {
        var v51 int64 = int64(v5)
        if v33 {
            for fi52 := 0; fi52 < 14; fi52++ {
                v53 := GenS1{F0: v29, F1: (fi50 << ((4294967295) & 63))}
                if v33 {
                    v5 -= v5
                    var v54 int8 = (v29 % ((v29) | 1))
                    v5 = v5
                    _ = v54
                } else {
                    delete(v2, "foo bar")
                }
                _ = len(v3)
                fmt.Println(v1, v29)
                _ = fi52
                _ = v53
            }
            v33 = (v1 >= v1)
        }
        _ = fi50
        _ = v51
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (true && false) {
        fmt.Println(v49)
        if ((66 == 128) || (1000 != -32768)) {
            v55 := GenS5{F0: 183, F1: 181}
            v1 -= (^(v1))
            fmt.Println(v32)
            _ = v55
        }
        switch (("hello" + ("foo bar" + ((("\n\t" + "foo bar") + "hello") + ""))) + ("a" + "hello")) {
        case "a":
            v5 = (v5 + 892)
        case "hello":
            for fi56 := 0; fi56 < 11; fi56++ {
                v33 = false
                _ = fi56
            }
        case "":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    } else {
        for fi57 := 0; fi57 < 16; fi57++ {
            fmt.Println(v33, v49, v49)
            switch 42 {
            case 2147483647:
                var v58 uint64 = v1
                _ = v58
            default:
                var v59 uint = 2
                _ = v59
            }
            var v60 uint8 = 127
            _ = fi57
            _ = v60
        }
    }
    v61 := GenS2{F0: uint32(v32), F1: int32(v29)}
    if v33 {
        fmt.Println(v49)
    } else {
        v33 = (128 == 1)
        if false {
            for fi62 := 0; fi62 < 7; fi62++ {
                switch "a" {
                case "hello":
                    for fi63 := 0; fi63 < 20; fi63++ {
                        v1 = v1
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                        _ = fi63
                    }
                case "\n\t":
                    var v64 bool = (4294967295 != 65535)
                    _ = v64
                default:
                    fmt.Println(v29, v5, v49)
                }
                v65 := GenS6{F0: ("hello" + "a"), F1: 1}
                _ = fi62
                _ = v65
            }
            var v66 uint32 = 4294967295
            _ = v66
        } else {
            v67 := GenS7{F0: 128, F1: 1}
            _ = v67
        }
        var v68 bool = (!(!false))
        _ = v68
    }
    v69 := GenS5{F0: 255, F1: 48}
    fmt.Println(v32, v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v49
    _ = v5
    _ = v61
    _ = v69
}
