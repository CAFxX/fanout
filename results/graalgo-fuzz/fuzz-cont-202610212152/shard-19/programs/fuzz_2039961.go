package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int; F1 bool; F2 bool }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint64, p2 uint32) uint32 {
    var v7 int8 = 47
    var v8 bool = ((p2 != (^(p2))) || (("a" + "a") > "hello"))
    p0 = v8
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p2 - p2)
}
func genFn1(p0 int16, p1 int16, p2 uint64) int16 {
    v9 := []int{42, -308, 1, 127}
    var v10 int16 = (p1 << ((42) & 63))
    var v11 int8 = -1
    fmt.Println(p2, v11, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 32767
}

func main() {
    v1 := []int{32767}
    var v2 int64 = 699
    var v3 int8 = 1
    var v4 uint32 = 2147483647
    v5 := make(map[int]int)
    var v6 uint8 = 255
    var v12 int32 = -1
    v13 := GenS1{F0: (v6 | (v6 >> ((805) & 63))), F1: (v2 * v2), F2: 456}
    for fi14 := 0; fi14 < 12; fi14++ {
        switch fi14 {
        case 128:
            v5[4294967295] = (fi14 * (fi14 ^ (fi14 - int(v4))))
        case 32768:
            switch 9223372036854775807 {
            case 1:
                if true {
                    var v15 uint8 = v6
                    v1 = append(v1, (fi14 - (fi14 ^ (fi14 >> ((4294967295) & 63)))))
                    _ = v15
                }
            default:
                var v16 uint8 = v6
                _ = v16
            }
        }
        _ = fi14
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        fmt.Println(v3, v4, v2)
    default:
        if (("" + "") >= "\n\t") {
            switch "hello" {
            case "a":
                for fi17 := 0; fi17 < 6; fi17++ {
                    if (!(65535 != 1)) {
                        if true {
                            v4 = 972
                        }
                        if (!((false && true) && false)) {
                            v18 := GenS1{F0: v6, F1: 948, F2: 65535}
                            for fi19 := 0; fi19 < 19; fi19++ {
                                _ = len(v5)
                                switch fi17 {
                                case 764:
                                    v3 = 42
                                case -9223372036854775808:
                                    _ = len(v1)
                                default:
                                    var v20 int = 32768
                                    _ = v20
                                }
                                if false {
                                    if false {
                                        v21 := GenS2{F0: 0, F1: (fi17 / (((-(fi19))) | 1))}
                                        v3 += (-(v3))
                                        _ = v21
                                    } else {
                                        fmt.Println(v2, fi19)
                                    }
                                    var v22 uint16 = 32767
                                    fmt.Println(v4)
                                    _ = v22
                                }
                                _ = fi19
                            }
                            _ = v18
                        } else {
                            var v23 float64 = float64(-2.25)
                            var v24 string = "hello"
                            for fi25 := 0; fi25 < 17; fi25++ {
                                v24 = "hello"
                                _ = fi25
                            }
                            _ = v23
                            _ = v24
                        }
                        var v26 int64 = (v2 % (((v2 & v2)) | 1))
                        _ = v26
                    }
                    v3 -= v3
                    _ = fi17
                }
            case "\n\t":
                switch 510 {
                case 127:
                    var v27 uint32 = genFn0((!true), 975, v4)
                    _ = v27
                case 42:
                    for fi28 := 0; fi28 < 19; fi28++ {
                        if (2 != 316) {
                            v29 := GenS3{F0: fi28, F1: false, F2: (v12 == (^(v12)))}
                            if (!(v12 != 0)) {
                                if true {
                                    var v30 uint32 = v4
                                    _ = v30
                                } else {
                                    var v31 int32 = 987
                                    v12 += v31
                                    _ = v31
                                }
                                v32 := GenS4{F0: (v6 + (v6 * v6)), F1: v2}
                                _ = v32
                            } else {
                                _ = len(v1)
                            }
                            fmt.Println(v4, v2, v3)
                            _ = v29
                        }
                        fmt.Println(v2, v12)
                        _ = fi28
                    }
                case -1:
                    v2 = (v2 / (((v2 + 1)) | 1))
                default:
                    v33 := GenS5{F0: 65535}
                    _ = v33
                }
            case "":
                v34 := GenS6{F0: 4294967295}
                _ = v34
            default:
                v35 := GenS4{F0: 198, F1: (v2 / ((v2) | 1))}
                _ = v35
            }
            v1 = append(v1, -1)
        }
    }
    fmt.Println(v2, v12)
    _ = len(v5)
    _ = len(v5)
    fmt.Println(v3)
    for fi36 := 0; fi36 < 2; fi36++ {
        switch "foo bar" {
        case "hello":
            v5[fi36] = fi36
        }
        var v37 bool = false
        _ = fi36
        _ = v37
    }
    var v38 bool = true
    var v39 string = "a"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
