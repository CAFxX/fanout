package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint8; F2 uint8 }
func genFn0() bool {
    var v4 int8 = 18
    v5 := []int{855, -146}
    v4 -= -128
    switch ((("hello" + ("\n\t" + "a")) + "foo bar") + (("a" + ("foo bar" + "a")) + ("a" + ""))) {
    case "hello":
        if (128 != 143) {
            v4 = v4
            fmt.Println(v4)
        } else {
            switch 4294967295 {
            case 0:
                _ = len(v5)
            }
        }
    case "a":
        v6 := GenS1{F0: v4, F1: 157, F2: 42}
        _ = v6
    case "":
        switch "\n\t" {
        case "foo bar":
            v4 = v4
        }
    }
    _ = len(v5)
    _ = v4
    _ = v5
    return false
}
func genFn1() uint16 {
    var v7 bool = (((-498 <= 42) || (false && (-1 > 114))) || false)
    var v8 uint16 = 1
    v8 = 939
    v7 = v7
    fmt.Println(v8, v7)
    _ = v7
    _ = v8
    return (v8 - (^(v8)))
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = (("hello" + "hello") + "hello")
    var v3 bool = false
    fmt.Println(v3)
    var v9 uint8 = 0
    for fi10 := 0; fi10 < 11; fi10++ {
        if false {
            for fi11 := 0; fi11 < 2; fi11++ {
                v9 += 0
                switch ((v2 + v2) + ("\n\t" + (((v2 + ("\n\t" + v2)) + (v2 + (v2 + "a"))) + "a"))) {
                case "foo bar":
                    var v12 uint32 = 560
                    _ = v12
                }
                _ = len(v1)
                var v13 int32 = 1
                _ = fi11
                _ = v13
            }
            v14 := []int{(fi10 | (^(fi10))), (fi10 >> ((65535) & 63)), (fi10 | -88)}
            if v3 {
                var v15 int16 = -1
                var v16 string = v2
                v17 := GenS1{F0: 94, F1: -825, F2: 42}
                _ = v15
                _ = v16
                _ = v17
            }
            _ = v14
        } else {
            var v18 float32 = 0.1
            v19 := []int{(fi10 + fi10), (fi10 + (fi10 + (fi10 / ((fi10) | 1))))}
            _ = v18
            _ = v19
        }
        _ = fi10
    }
    v20 := []int{32768}
    if v3 {
        var v21 uint16 = 189
        v22 := GenS1{F0: 1, F1: int(v21), F2: 42}
        fmt.Println(v3, v9, v2)
        _ = v21
        _ = v22
    } else {
        if len(v20) > 0 {
            v20[0] = -2147483648
        }
    }
    v23 := GenS2{F0: ((1 > 18446744073709551615) && (0.0 > (1.5 + (1e10 - float32(1.5)))))}
    switch -708 {
    case 1:
        v20 = append(v20, 4294967295)
    case -267:
        fmt.Println(v2, v2)
    case -85:
        var v24 uint64 = 0
        _ = v24
    }
    if (255 <= 42) {
        for fi25 := 0; fi25 < 4; fi25++ {
            if (genFn1() > 973) {
                fmt.Println(fi25)
            }
            _ = fi25
        }
        for fi26 := 0; fi26 < 9; fi26++ {
            if v3 {
                v27 := v1[fi26]
                for fi28 := 0; fi28 < 19; fi28++ {
                    switch v2 {
                    case "a":
                        _ = len(v20)
                    case "hello":
                        if len(v20) > 0 {
                            v20[0] = (fi28 ^ (fi26 - (fi28 / ((fi26) | 1))))
                        }
                    case "":
                        fmt.Println(fi28, v27, v3)
                    }
                    var v29 uint32 = 986
                    _ = fi28
                    _ = v29
                }
                fmt.Println(v27, v27, v27)
                _ = v27
            } else {
                var v30 int = (^(fi26))
                if v3 {
                    v3 = (int8(v9) >= 1)
                } else {
                    v9 += (v9 - (v9 * v9))
                    v1[fi26] = v3
                }
                fmt.Println(v9, fi26, v2)
                _ = v30
            }
            for fi31 := 0; fi31 < 9; fi31++ {
                v32 := GenS3{F0: v3}
                if (v3 && (!v3)) {
                    var v33 uint64 = 42
                    switch v2 {
                    case "hello":
                        if (!(v3 && true)) {
                            fmt.Println(v2, v9)
                            fmt.Println(v3, v9)
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                        }
                    case "foo bar":
                        v34 := GenS4{F0: -1, F1: v33, F2: 171}
                        _ = v34
                    case "":
                        v33 = (v33 % ((42) | 1))
                    }
                    _ = v33
                }
                v20 = append(v20, (fi31 | fi26))
                v35 := GenS4{F0: -9223372036854775808, F1: 924, F2: 263}
                _ = fi31
                _ = v32
                _ = v35
            }
            _ = fi26
        }
    } else {
        var v36 int16 = 0
        fmt.Println(v9)
        _ = v36
    }
    v37 := GenS5{F0: 4294967295, F1: (v9 >> ((65535) & 63))}
    var v38 int64 = 9223372036854775807
    v9 = (v9 / ((131) | 1))
    v38 = int64(v9)
    var v39 int16 = 1
    switch 1 {
    case 694:
        for fi40 := 0; fi40 < 13; fi40++ {
            v41 := v1[fi40]
            fmt.Println(v41, fi40)
            v38 = (v38 * v38)
            _ = fi40
            _ = v41
        }
    case 0:
        delete(v1, -1)
    case 1:
        v9 = 26
    default:
        switch -26 {
        case 968:
            fmt.Println(v38, v2)
        case -1:
            fmt.Println(v3, v2, v39)
        case 75:
            v42 := GenS6{F0: v3, F1: v9, F2: (v9 << ((42) & 63))}
            _ = v42
        default:
            for fi43 := 0; fi43 < 3; fi43++ {
                switch fi43 {
                case 41:
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                case 2147483647:
                    var v44 uint = 255
                    _ = v44
                default:
                    v20 = append(v20, -1)
                }
                var v45 int32 = 2147483647
                _ = fi43
                _ = v45
            }
        }
    }
    switch v2 {
    case "hello":
        v9 -= v9
    case "a":
        v46 := v1[4294967295]
        _ = v46
    case "\n\t":
        switch (("" + v2) + v2) {
        case "":
            var v47 int8 = 42
            _ = v47
        case "foo bar":
            for fi48 := 0; fi48 < 8; fi48++ {
                fmt.Println(fi48)
                var v49 int64 = (-(v38))
                _ = fi48
                _ = v49
            }
        case "hello":
            for fi50 := 0; fi50 < 5; fi50++ {
                for fi51 := 0; fi51 < 6; fi51++ {
                    v9 = (^(v9))
                    _ = fi51
                }
                v3 = false
                _ = fi50
            }
        default:
            v3 = true
        }
    }
    v9 = v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v9
}
