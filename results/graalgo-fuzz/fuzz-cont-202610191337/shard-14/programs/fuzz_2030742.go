package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint32; F1 uint16; F2 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool }
func genFn0(p0 uint32) uint64 {
    switch "foo bar" {
    case "hello":
        if (255 >= 888) {
            if false {
                switch -9223372036854775808 {
                case 459:
                    for fi5 := 0; fi5 < 4; fi5++ {
                        if true {
                            for fi6 := 0; fi6 < 16; fi6++ {
                                if false {
                                    var v7 int16 = -32768
                                    v8 := GenS1{F0: 605}
                                    _ = v7
                                    _ = v8
                                } else {
                                    v9 := []int{}
                                    v10 := GenS1{F0: (p0 * 69)}
                                    p0 = 81
                                    _ = v10
                                    _ = v9
                                }
                                fmt.Println(p0, fi5, fi6)
                                v11 := []int{}
                                _ = fi6
                                _ = v11
                            }
                            p0 += p0
                        }
                        if (false && true) {
                            v12 := []int{fi5, int(p0), fi5}
                            _ = v12
                        }
                        p0 = (p0 >> ((2) & 63))
                        _ = fi5
                    }
                }
            }
        } else {
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := []int{}
                var v15 int8 = 75
                _ = fi13
                _ = v14
                _ = v15
            }
            v16 := []int{42, 128, 4294967295, 128}
            var v17 int32 = -2147483648
            _ = v16
            _ = v17
        }
    case "foo bar":
        var v18 int32 = -1
        _ = v18
    case "\n\t":
        switch 9223372036854775807 {
        case 1:
            if (!(!(4294967295 < 134))) {
                switch -955 {
                case 9223372036854775807:
                    v19 := GenS1{F0: 42}
                    _ = v19
                case 2147483647:
                    v20 := []int{42}
                    _ = v20
                case -9223372036854775808:
                    if true {
                        var v21 int32 = -1
                        v22 := []int{2147483647}
                        var v23 int32 = v21
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        v24 := []int{-2147483648}
                        fmt.Println(p0)
                        p0 -= p0
                        _ = v24
                    }
                }
                if true {
                    switch -9223372036854775808 {
                    case 9223372036854775807:
                        v25 := []int{}
                        _ = v25
                    case 127:
                        v26 := []int{32768}
                        _ = v26
                    case 405:
                        p0 -= (p0 ^ p0)
                    }
                    switch 4294967295 {
                    case -1:
                        p0 = 2147483648
                    case -68:
                        for fi27 := 0; fi27 < 1; fi27++ {
                            for fi28 := 0; fi28 < 17; fi28++ {
                                v29 := GenS2{F0: "\n\t"}
                                var v30 uint64 = 42
                                _ = fi28
                                _ = v29
                                _ = v30
                            }
                            if false {
                                v31 := GenS3{F0: 134, F1: 42}
                                _ = v31
                            } else {
                                v32 := []int{(fi27 | fi27), (fi27 - (fi27 >> ((65535) & 63))), (fi27 + (-(fi27))), (fi27 * (fi27 / (((fi27 << ((0) & 63))) | 1)))}
                                _ = v32
                            }
                            p0 -= (^(p0))
                            _ = fi27
                        }
                    case 323:
                        for fi33 := 0; fi33 < 4; fi33++ {
                            p0 -= (p0 + p0)
                            p0 = p0
                            _ = fi33
                        }
                    default:
                        fmt.Println(p0)
                    }
                }
            } else {
                v34 := GenS4{F0: (("a" + "") + ("hello" + "a")), F1: "a"}
                _ = v34
            }
        case -519:
            if (!(1 < 126)) {
                p0 = 65535
                v35 := []int{-139, -525, 32768, -255}
                v36 := GenS5{F0: p0, F1: 65535, F2: (true && (((true && (0 <= 42)) && true) || (42 < 2147483647)))}
                _ = v35
                _ = v36
            }
        case 663:
            v37 := GenS6{F0: 0}
            _ = v37
        }
    default:
        v38 := []int{-529, -599, 575}
        _ = v38
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        p0 = p0
        _ = fi39
    }
    var v40 int8 = 83
    var v41 float32 = 1.5
    for fi42 := 0; fi42 < 5; fi42++ {
        v43 := GenS4{F0: ("hello" + "foo bar"), F1: ""}
        _ = fi42
        _ = v43
    }
    _ = p0
    _ = v40
    _ = v41
    return 522
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ("a" + "foo bar")
    var v3 bool = ((false || true) && true)
    v4 := make(map[string]bool)
    fmt.Println(v3, v2)
    v3 = v3
    var v44 uint16 = 854
    v3 = (!v3)
    v45 := v4["foo bar"]
    var v46 float32 = 1.5
    var v47 int8 = -78
    v48 := GenS3{F0: 145, F1: 1}
    v49 := GenS7{F0: ((v47 - v47) <= v47)}
    fmt.Println(v47)
    v46 = v46
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
