package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
func genFn0(p0 uint, p1 uint8, p2 bool) uint16 {
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 int = fi5
        v7 := GenS1{F0: p0}
        if true {
            for fi8 := 0; fi8 < 5; fi8++ {
                fmt.Println(p0)
                var v9 int = v6
                v10 := GenS2{F0: 0, F1: 1}
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v11 int = 674
    v12 := []int{0, (v11 % ((128) | 1)), v11}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 936
}
func genFn1() uint16 {
    switch -9223372036854775808 {
    case -2147483648:
        v13 := []int{}
        _ = v13
    case -9223372036854775808:
        var v14 int16 = 0
        _ = v14
    default:
        var v15 string = ""
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := []int{fi16, fi16, 1}
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        _ = fi16
        _ = v17
    }
    var v18 uint64 = 649
    if (false && false) {
        v18 = 0
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := GenS2{F0: (v18 >> ((255) & 63)), F1: 9223372036854775807}
            _ = fi19
            _ = v20
        }
        switch 32767 {
        case 42:
            switch "\n\t" {
            case "foo bar":
                var v21 bool = (0 == 127)
                _ = v21
            case "\n\t":
                fmt.Println(v18)
            case "a":
                v18 -= v18
            }
        case -263:
            v22 := GenS2{F0: v18, F1: 0}
            _ = v22
        default:
            if true {
                if ((846 > 32767) && true) {
                    v18 = (^(v18))
                    var v23 int64 = 0
                    v18 += v18
                    _ = v23
                }
                v24 := []int{}
                fmt.Println(v18)
                _ = v24
            } else {
                var v25 bool = (3.14159 < (1e10 - float32(3.14159)))
                v25 = (32767 < -1)
                var v26 string = ("\n\t" + (("" + ("" + "foo bar")) + (("foo bar" + "\n\t") + ("\n\t" + (("foo bar" + ("hello" + "\n\t")) + "hello")))))
                _ = v25
                _ = v26
            }
        }
    } else {
        v18 -= v18
        v27 := GenS1{F0: 255}
        for fi28 := 0; fi28 < 14; fi28++ {
            switch "hello" {
            case "\n\t":
                var v29 uint8 = 127
                _ = v29
            case "":
                for fi30 := 0; fi30 < 16; fi30++ {
                    if ((true && true) && (65535 == genFn0(18446744073709551615, 2, ((0 <= 2147483648) && (((127 >= (fi30 | (fi30 ^ (fi28 | fi28)))) || true) && ((true && false) || (0 >= 42))))))) {
                        fmt.Println(fi30, fi30, fi30)
                        if (127 != 42) {
                            v18 -= v18
                        } else {
                            v18 -= (v18 - 1)
                            var v31 uint64 = 4294967295
                            _ = v31
                        }
                        if false {
                            v18 = v18
                        }
                    }
                    v32 := []int{(-(fi28)), (^(fi30)), (fi28 >> ((4294967295) & 63))}
                    v33 := GenS3{F0: 0}
                    var v34 uint64 = (v18 | v18)
                    _ = fi30
                    _ = v32
                    _ = v33
                    _ = v34
                }
            case "foo bar":
                v35 := []int{(fi28 & fi28), (fi28 * fi28), fi28, (-(fi28))}
                _ = v35
            }
            fmt.Println(v18)
            _ = fi28
        }
        _ = v27
    }
    v18 = 18446744073709551615
    _ = v18
    return 907
}

func main() {
    var v1 int16 = 510
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    v4 := make(map[string]uint32)
    if false {
        v1 += 1
    } else {
        switch "foo bar" {
        case "":
            v1 += (v1 + -32768)
        case "hello":
            if ((42 > 4294967295) || (69 != 48)) {
                v1 = -1
                v36 := GenS3{F0: 0}
                _ = v36
            } else {
                v2["hello"] = 323
                var v37 uint32 = 600
                fmt.Println(v1)
                _ = v37
            }
        }
        var v38 int64 = 597
        fmt.Println(v1, v38)
        _ = v38
    }
    v39 := []int{-2147483648, 9223372036854775807, 9223372036854775807, -2147483648}
    var v40 int16 = v1
    v1 = v1
    if false {
        for fi41 := 0; fi41 < 15; fi41++ {
            v1 = v1
            v42 := GenS4{F0: 13, F1: 2, F2: true}
            delete(v3, (("hello" + "") + ("\n\t" + "foo bar")))
            _ = fi41
            _ = v42
        }
        v43 := GenS2{F0: 0, F1: 42}
        v40 -= (^(v40))
        _ = v43
    }
    if (!false) {
        var v44 int32 = 1
        v45 := v2["\n\t"]
        _ = v44
        _ = v45
    }
    v1 = (v40 % ((v40) | 1))
    v46 := GenS3{F0: -1}
    v2[(("foo bar" + "\n\t") + ("" + ""))] = 135
    var v47 int32 = 1
    if false {
        for fi48 := 0; fi48 < 13; fi48++ {
            switch fi48 {
            case 32768:
                fmt.Println(v47)
            case 32767:
                var v49 string = ("a" + "hello")
                _ = v49
            }
            _ = fi48
        }
    } else {
        switch 2147483647 {
        case 4294967295:
            v1 += (-(v40))
        case -1:
            switch 0 {
            case 2147483647:
                v50 := GenS4{F0: -763, F1: 4294967295, F2: (-2.25 != float64(3.14159))}
                _ = v50
            case 42:
                v47 = v47
            default:
                v1 -= 1
            }
        default:
            if false {
                var v51 uint64 = 66
                _ = v51
            } else {
                v40 = (^(v40))
            }
        }
    }
    var v52 int64 = 159
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v46
    _ = v47
    _ = v52
}
