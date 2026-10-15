package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 int16 }
func genFn0(p0 int32, p1 uint64, p2 int64) uint16 {
    for fi4 := 0; fi4 < 7; fi4++ {
        fmt.Println(fi4, p1)
        _ = fi4
    }
    v5 := GenS1{F0: 649, F1: 954, F2: 973}
    var v6 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 65535
}

func main() {
    var v1 string = ("hello" + "a")
    var v2 int = 128
    var v3 int64 = 697
    var v7 uint32 = 802
    if false {
        v1 = (((("\n\t" + v1) + (("\n\t" + v1) + ("" + (v1 + v1)))) + v1) + "\n\t")
        var v8 int32 = -1
        var v9 uint16 = 0
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: (v7 % ((v7) | 1)), F1: 4294967295, F2: v7}
    var v11 int32 = -2147483648
    var v12 int16 = 565
    var v13 int = -615
    var v14 float32 = 3.14159
    v15 := []int{(v2 - 127), (-(v13)), v2, v2}
    v15 = append(v15, (v2 - (v2 << ((485) & 63))))
    var v16 int32 = v11
    switch "\n\t" {
    case "hello":
        if len(v15) > 0 {
            v15[0] = (v2 ^ 32768)
        }
    case "":
        var v17 float32 = 1.5
        _ = v17
    default:
        switch ("\n\t" + ("a" + v1)) {
        case "a":
            if (true && true) {
                v14 = (v14 / 3.14159)
                for fi18 := 0; fi18 < 14; fi18++ {
                    switch -1 {
                    case -97:
                        for fi19 := 0; fi19 < 9; fi19++ {
                            var v20 int16 = (v12 >> ((0) & 63))
                            var v21 bool = false
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    case -1:
                        if true {
                            if (false || (!(1.5 <= 3.14159))) {
                                v3 = (v3 & (v3 * (v3 ^ (v3 / (((v3 / ((v3) | 1))) | 1)))))
                                var v22 uint64 = 18446744073709551615
                                _ = v22
                            } else {
                                fmt.Println(v11)
                                if (((1.5 - (0.1 * ((1.5 - 1.5) - 1e10))) + 0.0) == 0.0) {
                                    if len(v15) > 0 {
                                        v15[0] = v2
                                    }
                                    var v23 bool = false
                                    var v24 int32 = (v11 % ((309) | 1))
                                    _ = v23
                                    _ = v24
                                } else {
                                    var v25 int32 = (-(v16))
                                    _ = v25
                                }
                            }
                        }
                    }
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    _ = fi18
                }
                for fi26 := 0; fi26 < 9; fi26++ {
                    fmt.Println(v2)
                    for fi27 := 0; fi27 < 20; fi27++ {
                        var v28 bool = true
                        v29 := GenS2{F0: 2, F1: 42, F2: 1}
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    _ = fi26
                }
            } else {
                v13 -= (v2 ^ 32768)
                var v30 uint64 = 915
                if (true || ((-(v12)) <= 42)) {
                    if (true && ((^(v30)) < 4294967295)) {
                        for fi31 := 0; fi31 < 19; fi31++ {
                            if (!(!(v12 != (v12 * (v12 | 0))))) {
                                if false {
                                    switch v1 {
                                    case "\n\t":
                                        for fi32 := 0; fi32 < 6; fi32++ {
                                            switch fi32 {
                                            case -546:
                                                var v33 int32 = (v11 + v11)
                                                _ = v33
                                            case 0:
                                                v34 := GenS2{F0: 32767, F1: v30, F2: 255}
                                                _ = v34
                                            case 553:
                                                if (((v30 | v30) != 18446744073709551615) || false) {
                                                    if len(v15) > 0 {
                                                        v15[0] = int(v3)
                                                    }
                                                    if len(v15) > 0 {
                                                        v15[0] = v2
                                                    }
                                                } else {
                                                    var v35 uint8 = 127
                                                    _ = v35
                                                }
                                            }
                                            _ = fi32
                                        }
                                    default:
                                        var v36 int16 = (^(v12))
                                        _ = v36
                                    }
                                    v2 = (v2 >> ((353) & 63))
                                }
                                v7 += (v7 - v7)
                                var v37 float64 = 0.1
                                _ = v37
                            }
                            if true {
                                v38 := GenS1{F0: 2147483648, F1: 42, F2: (v7 << ((42) & 63))}
                                v39 := GenS3{F0: 18446744073709551615, F1: int(v30), F2: v12}
                                _ = v38
                                _ = v39
                            }
                            for fi40 := 0; fi40 < 18; fi40++ {
                                var v41 int64 = int64(v12)
                                _ = fi40
                                _ = v41
                            }
                            _ = fi31
                        }
                        v42 := GenS3{F0: uint(v7), F1: (v13 & (v2 << ((65535) & 63))), F2: (v12 >> ((255) & 63))}
                        switch (v13 >> ((1) & 63)) {
                        case 1:
                            switch (-(v13)) {
                            case 9223372036854775807:
                                var v43 bool = true
                                _ = v43
                            case 378:
                                switch 42 {
                                case 2147483647:
                                    for fi44 := 0; fi44 < 5; fi44++ {
                                        var v45 uint64 = v30
                                        _ = fi44
                                        _ = v45
                                    }
                                default:
                                    var v46 int = (v13 / ((79) | 1))
                                    _ = v46
                                }
                            case 42:
                                fmt.Println(v13)
                            default:
                                v47 := GenS2{F0: 1, F1: v30, F2: uint8(v30)}
                                _ = v47
                            }
                        case 32767:
                            v16 -= (v16 >> ((669) & 63))
                        case 2147483647:
                            v15 = append(v15, (-(v13)))
                        }
                        _ = v42
                    }
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    v1 = ((v1 + v1) + (v1 + v1))
                }
                _ = v30
            }
        }
    }
    fmt.Println(v12, v13, v13)
    v1 = v1
    var v48 float64 = 1e10
    v15 = append(v15, -898)
    switch v2 {
    case 1:
        v3 -= 0
    }
    var v49 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v7
}
