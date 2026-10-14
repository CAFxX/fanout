package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint64; F2 int }
type GenS3 struct { F0 int64; F1 int32; F2 bool }
type GenS4 struct { F0 int; F1 uint; F2 uint8 }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 uint = 656
    v2 := make(map[uint32]bool)
    var v3 uint16 = 32767
    var v4 uint8 = 128
    var v5 string = "foo bar"
    var v6 uint32 = 0
    v3 = (v3 + uint16(v1))
    v6 += 2147483648
    v7 := []int{1, 128, 9223372036854775807, 128}
    switch 414 {
    case -404:
        v5 = ("hello" + (((v5 + ((v5 + "\n\t") + ("hello" + (("hello" + ("a" + v5)) + "\n\t")))) + v5) + "hello"))
    case 127:
        if false {
            if true {
                if (v4 <= v4) {
                    var v8 int = 127
                    for fi9 := 0; fi9 < 10; fi9++ {
                        if true {
                            var v10 uint16 = (v3 + v3)
                            _ = v10
                        }
                        v11 := GenS1{F0: fi9}
                        _ = fi9
                        _ = v11
                    }
                    _ = v8
                } else {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v13 := GenS1{F0: int(v6)}
                        v14 := GenS1{F0: (fi12 | (fi12 - (fi12 / (((fi12 - -9223372036854775808)) | 1))))}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    if false {
                        var v15 uint16 = 0
                        for fi16 := 0; fi16 < 3; fi16++ {
                            if false {
                                v6 -= uint32(v15)
                            } else {
                                v17 := GenS1{F0: int(v6)}
                                v3 -= uint16(v4)
                                var v18 float64 = 3.14159
                                _ = v17
                                _ = v18
                            }
                            for fi19 := 0; fi19 < 1; fi19++ {
                                var v20 float64 = float64(3.14159)
                                v15 -= 255
                                v20 = (((v20 + v20) * 3.14159) - (v20 * 0.1))
                                _ = fi19
                                _ = v20
                            }
                            _ = fi16
                        }
                        for fi21 := 0; fi21 < 16; fi21++ {
                            v15 = (^(v15))
                            var v22 uint64 = 194
                            _ = fi21
                            _ = v22
                        }
                        _ = v15
                    } else {
                        if ((-1 != 727) && (2 <= v4)) {
                            v23 := v2[(v6 & v6)]
                            _ = v23
                        }
                    }
                    var v24 bool = (!true)
                    _ = v24
                }
            } else {
                if len(v7) > 0 {
                    v7[0] = -56
                }
            }
            var v25 string = "a"
            _ = v25
        } else {
            if false {
                v6 = 65535
                v4 = 178
                fmt.Println(v1)
            } else {
                if true {
                    if ((!((v4 ^ uint8(v1)) > v4)) || false) {
                        for fi26 := 0; fi26 < 2; fi26++ {
                            var v27 string = ((v5 + ("" + ("foo bar" + ((v5 + v5) + "")))) + v5)
                            _ = fi26
                            _ = v27
                        }
                    }
                    v28 := GenS2{F0: 381, F1: 18, F2: -9223372036854775808}
                    _ = v28
                } else {
                    var v29 uint8 = v4
                    _ = v29
                }
                _ = len(v2)
            }
            switch int(v3) {
            case 9223372036854775807:
                var v30 uint32 = (v6 + v6)
                _ = v30
            case 2147483647:
                var v31 int8 = 1
                _ = v31
            }
            v32 := GenS1{F0: 32767}
            _ = v32
        }
    case 2147483647:
        for fi33 := 0; fi33 < 14; fi33++ {
            _ = len(v2)
            var v34 uint8 = 10
            if true {
                var v35 uint = (v1 % ((v1) | 1))
                for fi36 := 0; fi36 < 6; fi36++ {
                    v37 := GenS3{F0: 0, F1: 2147483647, F2: (180 < v35)}
                    switch v5 {
                    case "hello":
                        var v38 int = (fi33 % ((fi33) | 1))
                        _ = v38
                    }
                    v39 := GenS1{F0: (fi36 - (-(fi33)))}
                    v40 := GenS2{F0: 1, F1: 247, F2: fi33}
                    v4 -= (v34 + v34)
                    _ = fi36
                    _ = v37
                    _ = v39
                    _ = v40
                }
                var v41 int8 = -119
                _ = v35
                _ = v41
            }
            _ = fi33
            _ = v34
        }
    }
    if (!(true && true)) {
        switch v5 {
        case "foo bar":
            for fi42 := 0; fi42 < 13; fi42++ {
                for fi43 := 0; fi43 < 2; fi43++ {
                    v44 := GenS4{F0: fi42, F1: 255, F2: (v4 % ((42) | 1))}
                    switch 9223372036854775807 {
                    case -2147483648:
                        if ((!false) || (-280 <= fi43)) {
                            fmt.Println(v3, fi43, v3)
                        }
                    }
                    _ = fi43
                    _ = v44
                }
                _ = fi42
            }
        default:
            v6 = 4294967295
        }
        switch 128 {
        case 32768:
            v4 = uint8(v6)
        case 763:
            v7 = append(v7, 127)
        default:
            if true {
                v45 := GenS3{F0: 0, F1: -755, F2: ((v4 & uint8(v3)) == v4)}
                fmt.Println(v5)
                v2[(v6 * v6)] = true
                _ = v45
            } else {
                for fi46 := 0; fi46 < 11; fi46++ {
                    for fi47 := 0; fi47 < 19; fi47++ {
                        if true {
                            delete(v2, v6)
                            v5 = v5
                        } else {
                            _ = len(v7)
                            var v48 uint8 = v4
                            _ = v48
                        }
                        _ = fi47
                    }
                    var v49 int8 = 0
                    _ = fi46
                    _ = v49
                }
                var v50 uint = (v1 ^ (v1 << (((v1 ^ v1)) & 63)))
                var v51 string = ("" + "foo bar")
                _ = v50
                _ = v51
            }
        }
    } else {
        if true {
            fmt.Println(v4, v1, v4)
        }
    }
    v1 = v1
    if len(v7) > 0 {
        v7[0] = -772
    }
    v6 = v6
    var v52 int64 = 572
    delete(v2, v6)
    v52 = (v52 * v52)
    v53 := GenS5{F0: (v6 | (v6 >> (((v1 ^ (v1 & (v1 ^ v1)))) & 63)))}
    fmt.Println(v4, v5, v1)
    v54 := GenS4{F0: 127, F1: 255, F2: v4}
    for fi55 := 0; fi55 < 19; fi55++ {
        switch v5 {
        case "a":
            var v56 uint = (v1 % ((v1) | 1))
            _ = v56
        case "\n\t":
            var v57 uint64 = 2
            _ = v57
        case "foo bar":
            v58 := GenS2{F0: int16(fi55), F1: 368, F2: -52}
            _ = v58
        default:
            v5 = ("" + v5)
        }
        _ = fi55
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v6
    _ = v7
}
