package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint16 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            if true {
                var v8 uint16 = 42
                var v9 bool = (true || (0 < 127))
                _ = v8
                _ = v9
            } else {
                var v10 int16 = 32767
                _ = v10
            }
            _ = fi7
        }
        _ = fi6
    }
    var v11 uint64 = 1
    if ((42 == -2147483648) || false) {
        v11 = v11
        if true {
            v12 := []int{-143, 32768}
            _ = v12
        } else {
            v13 := []int{-835, 42}
            if (true && (false && true)) {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            } else {
                switch -1 {
                case 759:
                    var v14 uint16 = 0
                    _ = v14
                }
            }
            _ = v13
        }
    } else {
        if (-58 >= -40) {
            fmt.Println(v11)
            for fi15 := 0; fi15 < 16; fi15++ {
                v11 -= v11
                if (((fi15 + 1) <= fi15) || false) {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        fmt.Println(v11, fi15)
                        v11 -= (v11 & (v11 % ((4294967295) | 1)))
                        _ = fi16
                    }
                    if (true && false) {
                        switch (fi15 / ((fi15) | 1)) {
                        case 32768:
                            for fi17 := 0; fi17 < 16; fi17++ {
                                v11 -= (v11 & 0)
                                v11 = v11
                                var v18 bool = false
                                _ = fi17
                                _ = v18
                            }
                        case -1:
                            v19 := GenS1{F0: 2}
                            _ = v19
                        case 0:
                            var v20 uint32 = 128
                            _ = v20
                        default:
                            switch "" {
                            case "foo bar":
                                switch -972 {
                                case 9223372036854775807:
                                    fmt.Println(fi15, v11)
                                case 0:
                                    fmt.Println(v11, fi15)
                                case 232:
                                    v21 := GenS2{F0: fi15}
                                    _ = v21
                                }
                            default:
                                var v22 int = (fi15 << ((4294967295) & 63))
                                _ = v22
                            }
                        }
                        var v23 uint32 = 1
                        var v24 uint16 = 65535
                        _ = v23
                        _ = v24
                    }
                } else {
                    var v25 uint = 4294967295
                    v26 := []int{656, (^(fi15))}
                    v11 = (v11 ^ v11)
                    _ = v25
                    _ = v26
                }
                _ = fi15
            }
            if true {
                if (!(1 == 127)) {
                    fmt.Println(v11)
                }
                for fi27 := 0; fi27 < 8; fi27++ {
                    v11 -= v11
                    var v28 uint32 = 2147483647
                    _ = fi27
                    _ = v28
                }
            }
        }
        v29 := GenS1{F0: 630}
        _ = v29
    }
    v30 := []int{178}
    switch "a" {
    case "\n\t":
        for fi31 := 0; fi31 < 4; fi31++ {
            if false {
                var v32 uint8 = 166
                var v33 uint64 = v11
                if len(v30) > 0 {
                    v30[0] = (fi31 << ((398) & 63))
                }
                _ = v32
                _ = v33
            } else {
                if (!true) {
                    v11 = (v11 << ((327) & 63))
                    v30 = append(v30, (-(fi31)))
                }
                v11 = v11
            }
            v34 := GenS1{F0: 32767}
            _ = fi31
            _ = v34
        }
    }
    _ = v11
    _ = v30
    return 2
}

func main() {
    var v1 int = -774
    var v2 float32 = -2.25
    v3 := []int{(v1 & 963), v1, v1}
    var v4 int64 = 0
    var v5 int64 = (v4 >> ((65535) & 63))
    v35 := GenS3{F0: (70 != 0), F1: 959}
    v3 = append(v3, (v1 >> ((65535) & 63)))
    v3 = append(v3, (^(v1)))
    v36 := GenS2{F0: (^(v1))}
    var v37 uint8 = 2
    v2 = 0.1
    if (!true) {
        if false {
            var v38 float64 = (float64((v2 + 1.5)) - 1.5)
            v39 := GenS3{F0: true, F1: 705}
            _ = v38
            _ = v39
        } else {
            v40 := GenS4{F0: v37}
            v4 += int64(v37)
            _ = v40
        }
    } else {
        for fi41 := 0; fi41 < 18; fi41++ {
            switch "foo bar" {
            case "a":
                v1 -= fi41
            default:
                v1 = v1
            }
            v2 = ((v2 / v2) + (0.1 + float32((1e10 * float64(float32(0.1))))))
            _ = fi41
        }
        v3 = append(v3, (-(v1)))
        if true {
            var v42 uint16 = 255
            v5 = 177
            v42 -= (v42 >> ((1) & 63))
            _ = v42
        }
    }
    v43 := GenS3{F0: true, F1: -32768}
    if (!(642 != 2147483647)) {
        if len(v3) > 0 {
            v3[0] = v1
        }
    } else {
        v5 -= v5
        v3 = append(v3, 1)
        switch "" {
        case "foo bar":
            v44 := GenS1{F0: 65535}
            _ = v44
        case "hello":
            if (1 == v4) {
                var v45 bool = (42 != 2147483647)
                var v46 string = "\n\t"
                for fi47 := 0; fi47 < 7; fi47++ {
                    v48 := GenS5{F0: 255, F1: 2, F2: 0}
                    var v49 uint32 = 65535
                    if len(v3) > 0 {
                        v3[0] = (v1 << ((1) & 63))
                    }
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                _ = v45
                _ = v46
            }
        default:
            v1 = v1
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v2, v37)
    v5 = 9223372036854775807
    var v50 int32 = 183
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v43
    _ = v5
    _ = v50
}
