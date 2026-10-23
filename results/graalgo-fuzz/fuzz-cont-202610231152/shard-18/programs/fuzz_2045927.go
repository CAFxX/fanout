package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int; F1 uint64; F2 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint, p1 string) int16 {
    for fi4 := 0; fi4 < 1; fi4++ {
        p0 = 268
        for fi5 := 0; fi5 < 13; fi5++ {
            v6 := GenS1{F0: p1, F1: true}
            var v7 uint = p0
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    p1 = "hello"
    p0 += 18446744073709551615
    _ = p0
    _ = p1
    return 452
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    v8 := []int{0, 42}
    if false {
        v9 := GenS1{F0: "hello", F1: true}
        _ = v9
    } else {
        v10 := GenS1{F0: (("\n\t" + ("a" + "a")) + "\n\t"), F1: true}
        var v11 uint64 = 87
        _ = v10
        _ = v11
    }
    var v12 int64 = 0
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 int16 = 42
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: "a", F1: false}
    for fi16 := 0; fi16 < 1; fi16++ {
        if ((0.0 - 0.1) < 0.1) {
            if false {
                if (false && true) {
                    if len(v8) > 0 {
                        v8[0] = fi16
                    }
                } else {
                    if true {
                        if (42 >= 1) {
                            if (2147483648 < 2) {
                                if false {
                                    v12 += 0
                                }
                                v17 := GenS2{F0: 4294967295, F1: -1, F2: 1}
                                var v18 uint = 65535
                                _ = v17
                                _ = v18
                            }
                            if (!((3.14159 * 0.0) >= (1.5 - ((0.0 * -2.25) + (0.1 - -2.25))))) {
                                _ = len(v8)
                                _ = len(v8)
                                v19 := GenS3{F0: 887, F1: 504}
                                _ = v19
                            }
                        }
                    }
                    if (566 == 6) {
                        _ = len(v1)
                        v20 := GenS4{F0: -1, F1: 313, F2: -744}
                        v21 := GenS5{F0: 42}
                        _ = v20
                        _ = v21
                    } else {
                        var v22 uint32 = 0
                        _ = v22
                    }
                    v12 -= (v12 << ((1) & 63))
                }
                if true {
                    switch (("a" + "") + "a") {
                    case "\n\t":
                        var v23 uint = 64
                        _ = v23
                    case "hello":
                        var v24 uint64 = 646
                        _ = v24
                    case "a":
                        v25 := GenS4{F0: fi16, F1: 1, F2: 1}
                        _ = v25
                    default:
                        var v26 uint8 = 255
                        _ = v26
                    }
                    v12 = 0
                    _ = len(v3)
                }
                for fi27 := 0; fi27 < 13; fi27++ {
                    v12 -= v12
                    if false {
                        v2[(fi27 << ((255) & 63))] = 0
                    } else {
                        var v28 uint64 = 42
                        _ = v28
                    }
                    v2[155] = 364
                    v29 := GenS4{F0: 128, F1: 18446744073709551615, F2: 42}
                    fmt.Println(fi16, fi16, fi16)
                    _ = fi27
                    _ = v29
                }
            } else {
                switch "\n\t" {
                case "foo bar":
                    for fi30 := 0; fi30 < 17; fi30++ {
                        var v31 uint8 = 1
                        v31 += (^(v31))
                        var v32 uint8 = v31
                        var v33 int8 = 74
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                case "\n\t":
                    fmt.Println(v12)
                default:
                    for fi34 := 0; fi34 < 19; fi34++ {
                        var v35 int = (-(fi16))
                        var v36 int = 1
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
                for fi37 := 0; fi37 < 5; fi37++ {
                    v38 := GenS6{F0: true}
                    var v39 float64 = float64((float32(0.0) - 1.5))
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                var v40 uint8 = 85
                _ = v40
            }
            switch (("foo bar" + "") + ("foo bar" + "foo bar")) {
            case "\n\t":
                var v41 bool = false
                _ = v41
            case "foo bar":
                v1[2147483648] = ((-(v12)) == v12)
            default:
                if false {
                    v42 := GenS1{F0: "foo bar", F1: (-128 > -1)}
                    _ = v42
                } else {
                    v12 = (v12 * 0)
                    switch 32767 {
                    case -2147483648:
                        var v43 int8 = 1
                        _ = v43
                    case -859:
                        v44 := v3[(fi16 >> ((403) & 63))]
                        _ = v44
                    case 1:
                        v45 := GenS7{F0: true, F1: 32767}
                        _ = v45
                    }
                }
            }
        } else {
            v12 = -269
            switch 4294967295 {
            case -1:
                v12 += (v12 * (v12 * (v12 ^ (v12 & -921))))
            case 32768:
                var v46 uint8 = 106
                _ = v46
            case 0:
                if ((fi16 + fi16) <= fi16) {
                    fmt.Println(fi16, fi16)
                    _ = len(v8)
                    var v47 float32 = 0.0
                    _ = v47
                }
            }
        }
        _ = fi16
    }
    v12 -= (v12 - (v12 & -235))
    v12 -= (v12 ^ (v12 ^ (v12 & v12)))
    v8 = append(v8, -2147483648)
    switch 2147483647 {
    case 2147483647:
        if false {
            switch "a" {
            case "a":
                v12 = -1
            }
            var v48 int8 = 76
            fmt.Println(v48)
            _ = v48
        } else {
            var v49 uint32 = 4294967295
            _ = v49
        }
    }
    switch "hello" {
    case "hello":
        for fi50 := 0; fi50 < 7; fi50++ {
            for fi51 := 0; fi51 < 12; fi51++ {
                var v52 uint16 = 0
                switch ("a" + "hello") {
                case "hello":
                    v12 = (v12 - (v12 & (^(v12))))
                }
                _ = fi51
                _ = v52
            }
            _ = fi50
        }
    default:
        var v53 uint32 = 2
        _ = v53
    }
    fmt.Println(v12)
    var v54 float32 = (1e10 - 1.5)
    v12 += 1
    v12 += v12
    fmt.Println(v12, v12)
    v2[-2147483648] = 65535
    v12 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v54
    _ = v8
}
