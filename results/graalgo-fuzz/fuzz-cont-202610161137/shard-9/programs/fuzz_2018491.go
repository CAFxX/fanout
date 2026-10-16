package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint32 }
func genFn0(p0 string) int {
    var v5 int8 = 127
    v6 := []int{-476, 2147483647, 4294967295, 89}
    v7 := GenS1{F0: 255, F1: (((p0 + p0) >= p0) || (-2.25 >= 3.14159))}
    for fi8 := 0; fi8 < 1; fi8++ {
        p0 = (p0 + ((p0 + "foo bar") + ""))
        v6 = append(v6, -9223372036854775808)
        var v9 uint64 = 361
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v5 = (^(v5))
        _ = fi10
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 846
}
func genFn1(p0 bool, p1 uint8) string {
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int32 = -2147483648
        _ = fi11
        _ = v12
    }
    v13 := []int{2147483647}
    _ = p0
    _ = p1
    _ = v13
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    var v4 int32 = 42
    switch 42 {
    case -9223372036854775808:
        v4 = (v4 << ((0) & 63))
    case 4294967295:
        v4 = (-(v4))
    case 127:
        v14 := []int{}
        _ = v14
    }
    if false {
        if (0 != 42) {
            v4 = v4
            var v15 int16 = 42
            for fi16 := 0; fi16 < 7; fi16++ {
                v17 := []int{(-(fi16)), 2147483647, genFn0(("" + "a"))}
                v4 = (-(v4))
                if (!true) {
                    if ((true || (true || (203 != 1))) || true) {
                        v1[65535] = (("" + "") + ("a" + "foo bar"))
                    } else {
                        var v18 uint64 = 1
                        for fi19 := 0; fi19 < 2; fi19++ {
                            v20 := GenS1{F0: 1, F1: true}
                            var v21 string = (("" + "") + "hello")
                            switch v21 {
                            case "hello":
                                v22 := GenS2{F0: -1, F1: 42, F2: 706}
                                _ = v22
                            case "":
                                v15 = (v15 * (-(v15)))
                            case "foo bar":
                                var v23 uint16 = 32767
                                _ = v23
                            }
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        _ = v18
                    }
                    switch (fi16 + (fi16 - (fi16 ^ 42))) {
                    case -9223372036854775808:
                        v4 = -2147483648
                    case -2147483648:
                        for fi24 := 0; fi24 < 10; fi24++ {
                            v25 := GenS2{F0: 0, F1: 42, F2: 824}
                            _ = fi24
                            _ = v25
                        }
                    }
                    for fi26 := 0; fi26 < 19; fi26++ {
                        var v27 uint64 = 678
                        switch (fi16 & genFn0("foo bar")) {
                        case 1:
                            v28 := v2[genFn0("hello")]
                            _ = v28
                        case -1:
                            var v29 float64 = (float64((float32(3.14159) / 1e10)) * -2.25)
                            _ = v29
                        default:
                            var v30 int64 = 0
                            _ = v30
                        }
                        _ = len(v17)
                        var v31 int8 = 127
                        _ = fi26
                        _ = v27
                        _ = v31
                    }
                }
                _ = fi16
                _ = v17
            }
            _ = v15
        }
        v4 += 1
        var v32 int16 = -1
        _ = v32
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        var v34 int64 = -9223372036854775808
        _ = fi33
        _ = v34
    }
    switch -549 {
    case 42:
        if (2147483647 >= (v4 * v4)) {
            var v35 uint = 200
            v35 = (v35 | v35)
            _ = v35
        } else {
            for fi36 := 0; fi36 < 17; fi36++ {
                fmt.Println(fi36)
                var v37 float32 = 0.0
                v4 += v4
                _ = fi36
                _ = v37
            }
            v4 = -1
        }
    case 1:
        v38 := []int{int(v4), -1, genFn0(""), 32767}
        _ = v38
    default:
        v4 += 488
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        v40 := GenS3{F0: 20, F1: 65535}
        v4 = (v4 << ((852) & 63))
        v41 := GenS1{F0: 163, F1: (42 > -1)}
        _ = fi39
        _ = v40
        _ = v41
    }
    v42 := []int{899, -9223372036854775808, -9223372036854775808}
    var v43 uint64 = 18446744073709551615
    if false {
        for fi44 := 0; fi44 < 10; fi44++ {
            v43 += (v43 ^ (v43 % (((v43 & (v43 << ((481) & 63)))) | 1)))
            for fi45 := 0; fi45 < 16; fi45++ {
                var v46 int16 = 64
                _ = fi45
                _ = v46
            }
            _ = fi44
        }
    }
    if true {
        var v47 int32 = (v4 - (v4 ^ v4))
        for fi48 := 0; fi48 < 2; fi48++ {
            v47 = (v47 + v4)
            v4 = v4
            var v49 uint = 65535
            _ = fi48
            _ = v49
        }
        fmt.Println(v43, v4)
        _ = v47
    } else {
        for fi50 := 0; fi50 < 12; fi50++ {
            var v51 float64 = (-2.25 * -2.25)
            _ = fi50
            _ = v51
        }
        v52 := GenS4{F0: 4294967295, F1: 4294967295, F2: 688}
        var v53 uint16 = 32767
        _ = v52
        _ = v53
    }
    switch -2147483648 {
    case 134:
        v43 -= (^(v43))
    default:
        fmt.Println(v43, v43)
    }
    if (!true) {
        v54 := GenS3{F0: 1, F1: 18446744073709551615}
        v43 = (v43 ^ 4294967295)
        _ = v54
    } else {
        if false {
            var v55 int16 = 1
            var v56 uint16 = 32767
            _ = v55
            _ = v56
        } else {
            var v57 uint64 = (v43 % ((v43) | 1))
            _ = v57
        }
        var v58 int8 = 0
        v43 += v43
        _ = v58
    }
    v43 = 0
    if len(v42) >= 2 {
        v42 = v42[1:len(v42)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
}
