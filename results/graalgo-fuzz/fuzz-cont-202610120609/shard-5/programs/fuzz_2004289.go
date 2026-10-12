package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16 }
func genFn0(p0 int8, p1 bool, p2 uint64) int32 {
    var v7 int = -1
    fmt.Println(p2, p1)
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint64 = (^(p2))
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(p0, p2, p2)
        fmt.Println(fi10, p1)
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 int16 = 42
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v13 := GenS1{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v7
    return 2147483647
}
func genFn1(p0 int16) int8 {
    p0 = (p0 + (p0 % (((p0 + (p0 ^ (p0 >> ((950) & 63))))) | 1)))
    fmt.Println(p0)
    v14 := []int{9223372036854775807, 127}
    p0 -= 32767
    _ = p0
    _ = v14
    return 100
}

func main() {
    var v1 float64 = -2.25
    var v2 uint32 = 2147483647
    v3 := make(map[string]uint32)
    v4 := []string{}
    var v5 uint8 = 1
    var v6 uint32 = v2
    v1 = v1
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            v2 = v2
            switch "a" {
            case "a":
                switch (fi15 ^ (fi15 * fi15)) {
                case 42:
                    fmt.Println(v1)
                }
            case "hello":
                v5 += (v5 * v5)
            case "":
                v2 += (^(v2))
            }
            fmt.Println(fi15, v6, fi16)
            _ = fi16
        }
        var v17 uint16 = 65535
        v5 += (v5 - (v5 - 0))
        _ = fi15
        _ = v17
    }
    var v18 int8 = 100
    delete(v3, "foo bar")
    for fi19 := 0; fi19 < 20; fi19++ {
        for fi20 := 0; fi20 < 19; fi20++ {
            v1 += v1
            v18 -= (v18 | v18)
            var v21 int16 = -32768
            v5 += v5
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v3[("foo bar" + "\n\t")] = (v6 % ((v2) | 1))
        var v23 uint32 = (v2 / ((v2) | 1))
        _ = fi22
        _ = v23
    }
    delete(v3, "")
    if true {
        v24 := GenS1{F0: 41}
        v25 := GenS2{F0: 558}
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 9; fi26++ {
            for fi27 := 0; fi27 < 12; fi27++ {
                if ((v2 % ((v2) | 1)) != (v2 | v2)) {
                    v28 := GenS1{F0: (v5 + v5)}
                    _ = v28
                } else {
                    switch fi26 {
                    case 153:
                        v29 := GenS1{F0: (v5 % ((v5) | 1))}
                        _ = v29
                    case 0:
                        _ = len(v3)
                    case 964:
                        v30 := GenS3{F0: 2}
                        _ = v30
                    }
                    if false {
                        if false {
                            var v31 uint8 = (v5 + (v5 ^ 1))
                            v32 := v3["foo bar"]
                            _ = v31
                            _ = v32
                        } else {
                            var v33 int8 = -38
                            _ = v33
                        }
                    }
                    if ((875 < 255) && true) {
                        if (!false) {
                            v34 := GenS3{F0: 32767}
                            v5 = 33
                            if true {
                                v6 += 1
                            }
                            _ = v34
                        }
                        var v35 int32 = -1
                        _ = v35
                    }
                    switch fi27 {
                    case 2147483647:
                        _ = len(v4)
                    default:
                        delete(v3, ("hello" + "hello"))
                    }
                }
                _ = fi27
            }
            var v36 uint8 = 255
            _ = fi26
            _ = v36
        }
    }
    v6 = (v2 >> ((942) & 63))
    for fi37 := 0; fi37 < 1; fi37++ {
        var v38 int8 = (v18 << ((28) & 63))
        _ = fi37
        _ = v38
    }
    delete(v3, "a")
    v1 = float64(1.5)
    switch 2147483647 {
    case 2147483647:
        fmt.Println(v6)
    case 128:
        if ((65535 == 18446744073709551615) || false) {
            if true {
                v18 = 127
                switch -725 {
                case -2147483648:
                    fmt.Println(v2, v18)
                }
            } else {
                v3[((("foo bar" + "") + "\n\t") + "a")] = (v2 - (v2 * (^(v6))))
                if (!true) {
                    v39 := GenS2{F0: 42}
                    if (32767 >= 868) {
                        v40 := GenS4{F0: 42}
                        _ = v40
                    }
                    _ = v39
                } else {
                    fmt.Println(v5, v6, v18)
                    for fi41 := 0; fi41 < 1; fi41++ {
                        for fi42 := 0; fi42 < 13; fi42++ {
                            v1 = (float64(-2.25) - float64((1.5 - 1e10)))
                            _ = fi42
                        }
                        _ = len(v4)
                        var v43 uint64 = 18446744073709551615
                        _ = fi41
                        _ = v43
                    }
                }
            }
            _ = len(v4)
            for fi44 := 0; fi44 < 17; fi44++ {
                v3["\n\t"] = v2
                _ = fi44
            }
        }
    }
    if false {
        if ((!false) && ((v2 << ((563) & 63)) >= (v2 + (v2 | 2)))) {
            if false {
                fmt.Println(v18, v18)
                var v45 int = -9223372036854775808
                var v46 uint8 = 127
                _ = v45
                _ = v46
            } else {
                if false {
                    var v47 int8 = 100
                    for fi48 := 0; fi48 < 19; fi48++ {
                        var v49 float64 = 1e10
                        _ = fi48
                        _ = v49
                    }
                    for fi50 := 0; fi50 < 20; fi50++ {
                        fmt.Println(v18, v5)
                        var v51 uint64 = 4294967295
                        _ = fi50
                        _ = v51
                    }
                    _ = v47
                } else {
                    _ = len(v3)
                }
            }
            if (!false) {
                var v52 int = 2147483647
                v52 += v52
                v4 = append(v4, "")
                _ = v52
            } else {
                fmt.Println(v18)
                for fi53 := 0; fi53 < 9; fi53++ {
                    if (!((v5 + (v5 & (^(v5)))) <= (v5 | (v5 << ((4294967295) & 63))))) {
                        for fi54 := 0; fi54 < 5; fi54++ {
                            var v55 int32 = -622
                            _ = fi54
                            _ = v55
                        }
                        v2 = 380
                    } else {
                        for fi56 := 0; fi56 < 2; fi56++ {
                            v4 = append(v4, "\n\t")
                            var v57 int16 = 1000
                            _ = fi56
                            _ = v57
                        }
                        if (-951 == -695) {
                            v2 = (v2 >> ((42) & 63))
                        }
                    }
                    _ = fi53
                }
            }
        } else {
            if false {
                v58 := GenS5{F0: -32768}
                v59 := GenS2{F0: 1}
                v60 := GenS1{F0: v5}
                _ = v58
                _ = v59
                _ = v60
            }
            fmt.Println(v6)
        }
    } else {
        var v61 float64 = 3.14159
        _ = v61
    }
    fmt.Println(v2)
    fmt.Println(v2, v5)
    v18 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
