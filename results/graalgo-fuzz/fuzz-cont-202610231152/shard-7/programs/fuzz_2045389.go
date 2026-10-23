package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 uint16; F2 int }
func genFn0(p0 uint8, p1 uint16) uint {
    if true {
        var v6 uint8 = p0
        _ = v6
    } else {
        for fi7 := 0; fi7 < 5; fi7++ {
            p0 -= 38
            v8 := GenS1{F0: "\n\t", F1: 127, F2: (p1 & (p1 ^ p1))}
            v9 := GenS2{F0: true, F1: -32768, F2: 4294967295}
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 int64 = 1
        _ = v10
    }
    var v11 int = 78
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            if false {
                v14 := []int{(v11 % (((fi13 | fi12)) | 1))}
                v11 -= v11
                _ = v14
            }
            p1 = 32767
            v15 := []int{v11, fi13, (fi12 >> ((0) & 63))}
            _ = fi13
            _ = v15
        }
        _ = fi12
    }
    fmt.Println(p0, p0, p1)
    if (!(208 > 1)) {
        var v16 string = ""
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v11
    return 1
}
func genFn1() uint16 {
    if true {
        var v17 uint64 = 624
        _ = v17
    } else {
        var v18 uint = 42
        v19 := GenS2{F0: false, F1: 32767, F2: 65535}
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 int32 = 958
        v22 := []int{-624}
        _ = fi20
        _ = v21
        _ = v22
    }
    return 2
}

func main() {
    v1 := []uint32{1, 629, 203}
    var v2 string = ""
    var v3 uint16 = 1
    v4 := make(map[uint32]bool)
    var v5 int64 = 299
    switch 127 {
    case 127:
        _ = len(v4)
    case -273:
        v23 := GenS2{F0: true, F1: -1, F2: 13}
        _ = v23
    default:
        v24 := GenS2{F0: true, F1: -605, F2: 2}
        _ = v24
    }
    if (!false) {
        v25 := GenS3{F0: 1, F1: (v3 + 65535)}
        fmt.Println(v5, v5, v2)
        for fi26 := 0; fi26 < 15; fi26++ {
            if true {
                v27 := GenS4{F0: 2}
                _ = len(v1)
                v28 := GenS1{F0: ("foo bar" + (v2 + v2)), F1: -125, F2: 973}
                _ = v27
                _ = v28
            } else {
                for fi29 := 0; fi29 < 15; fi29++ {
                    switch fi26 {
                    case -885:
                        if false {
                            delete(v4, 2)
                            v5 -= 42
                            if true {
                                var v30 uint16 = (^(v3))
                                switch "\n\t" {
                                case "a":
                                    _ = len(v4)
                                }
                                _ = v30
                            } else {
                                var v31 int8 = 100
                                var v32 uint = 1000
                                if ((!(false && ((fi26 * fi26) >= (fi26 >> ((v32) & 63))))) && (!false)) {
                                    switch (v2 + v2) {
                                    case "":
                                        if (-1 > (v5 | int64(v3))) {
                                            if (true && false) {
                                                v33 := GenS3{F0: (v32 * v32), F1: v3}
                                                if (!true) {
                                                    v31 += (v31 + -128)
                                                    if true {
                                                        fmt.Println(fi26)
                                                        if len(v1) > 0 {
                                                            v1[0] = 304
                                                        }
                                                    }
                                                    v3 = v3
                                                } else {
                                                    if len(v1) > 0 {
                                                        v1[0] = 65535
                                                    }
                                                    fmt.Println(fi26)
                                                }
                                                _ = v33
                                            }
                                        }
                                    case "foo bar":
                                        if false {
                                            if (true && false) {
                                                switch 32768 {
                                                case 32768:
                                                    _ = len(v1)
                                                case 4294967295:
                                                    v3 += 1
                                                case 0:
                                                    var v34 string = ("\n\t" + v2)
                                                    _ = v34
                                                }
                                                v5 = 1
                                            }
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v31 = v31
                                            _ = len(v1)
                                        }
                                    }
                                    var v35 int64 = -672
                                    _ = v35
                                } else {
                                    if false {
                                        var v36 uint16 = v3
                                        _ = v36
                                    }
                                }
                                _ = v31
                                _ = v32
                            }
                            fmt.Println(v3, v2)
                        }
                    case -2147483648:
                        var v37 int32 = 1
                        _ = v37
                    case 0:
                        v38 := GenS3{F0: 18446744073709551615, F1: v3}
                        _ = v38
                    }
                    _ = fi29
                }
                var v39 int32 = -313
                _ = v39
            }
            _ = fi26
        }
        _ = v25
    }
    v3 -= (v3 - v3)
    v40 := GenS5{F0: v5, F1: genFn1(), F2: 9223372036854775807}
    for fi41 := 0; fi41 < 10; fi41++ {
        switch fi41 {
        case -447:
            v3 = 2
        case -1:
            v2 = (v2 + v2)
        case 32768:
            for fi42 := 0; fi42 < 3; fi42++ {
                var v43 bool = (125 == -1)
                var v44 int8 = 42
                _ = fi42
                _ = v43
                _ = v44
            }
        }
        var v45 uint32 = 2147483647
        _ = fi41
        _ = v45
    }
    for fi46 := 0; fi46 < 19; fi46++ {
        for fi47 := 0; fi47 < 14; fi47++ {
            var v48 uint64 = 38
            fmt.Println(fi47, v2, fi46)
            var v49 uint = 65535
            _ = fi47
            _ = v48
            _ = v49
        }
        switch (v2 + "hello") {
        case "":
            v4[42] = false
        default:
            switch (fi46 ^ fi46) {
            case -754:
                for fi50 := 0; fi50 < 4; fi50++ {
                    v5 = (v5 - v5)
                    _ = fi50
                }
            case 4294967295:
                if true {
                    v51 := GenS1{F0: "a", F1: -1, F2: (v3 & (v3 << ((1) & 63)))}
                    _ = v51
                } else {
                    v5 = v5
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            default:
                v52 := GenS3{F0: 2, F1: genFn1()}
                _ = v52
            }
        }
        _ = fi46
    }
    v5 = (v5 >> ((0) & 63))
    fmt.Println(v2, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
