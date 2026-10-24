package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int; F1 uint16; F2 bool }
type GenS5 struct { F0 uint8; F1 string; F2 int }
type GenS6 struct { F0 uint32 }
func genFn0() uint64 {
    var v7 string = (("foo bar" + "") + "")
    v7 = "hello"
    v7 = (v7 + (v7 + "foo bar"))
    fmt.Println(v7)
    var v8 uint16 = 0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 42
    v3 := make(map[string]uint32)
    v4 := []int{}
    var v5 uint8 = 127
    var v6 uint8 = (v5 + v5)
    if (119 > 42) {
        v4 = append(v4, 128)
        var v9 int64 = (^(v1))
        switch "foo bar" {
        case "a":
            var v10 bool = false
            _ = v10
        case "":
            var v11 uint16 = 255
            _ = v11
        }
        _ = v9
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        if (!(!(25 <= -904))) {
            fmt.Println(v5, v6, v6)
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v6 = v5
            for fi13 := 0; fi13 < 10; fi13++ {
                _ = len(v3)
                delete(v3, ("\n\t" + ("foo bar" + "foo bar")))
                _ = len(v3)
                _ = fi13
            }
        }
        _ = fi12
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint16 = 77
        for fi16 := 0; fi16 < 20; fi16++ {
            fmt.Println(v1)
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    _ = len(v4)
    v17 := GenS1{F0: false}
    v1 += (v1 % ((1) | 1))
    v5 = v6
    if ((0 > 0) || true) {
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS2{F0: (genFn0() == genFn0()), F1: 1}
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 uint8 = v5
                v1 -= v1
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 int32 = -574
                _ = fi22
                _ = v23
            }
            _ = fi18
            _ = v19
        }
        v5 = v6
        v24 := GenS1{F0: (-1 < -128)}
        _ = v24
    } else {
        for fi25 := 0; fi25 < 10; fi25++ {
            v26 := v3["foo bar"]
            for fi27 := 0; fi27 < 18; fi27++ {
                v6 += (v5 & v6)
                v26 = v26
                var v28 int16 = (v2 << ((1) & 63))
                var v29 int16 = 201
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = fi25
            _ = v26
        }
        if (1e10 != (1e10 / -2.25)) {
            for fi30 := 0; fi30 < 5; fi30++ {
                v3[("a" + "a")] = 65535
                v31 := GenS1{F0: (471 >= 18446744073709551615)}
                var v32 bool = ((127 > 0) && ("\n\t" > (((("\n\t" + ("a" + "")) + "hello") + "\n\t") + ("a" + ("" + "a")))))
                _ = fi30
                _ = v31
                _ = v32
            }
        } else {
            var v33 uint = 228
            _ = v33
        }
    }
    fmt.Println(v6)
    v6 = 127
    v34 := GenS2{F0: false, F1: 2}
    var v35 int16 = (v2 >> ((641) & 63))
    switch ("a" + "") {
    case "a":
        v35 = v2
    case "foo bar":
        for fi36 := 0; fi36 < 5; fi36++ {
            if (true && (552 < 4294967295)) {
                switch fi36 {
                case -129:
                    var v37 uint32 = 1
                    _ = v37
                case -2147483648:
                    v2 = (^(v35))
                default:
                    for fi38 := 0; fi38 < 12; fi38++ {
                        switch ("" + "foo bar") {
                        case "hello":
                            if true {
                                v39 := GenS3{F0: (v1 % ((v1) | 1)), F1: fi38}
                                v35 = (v35 >> ((397) & 63))
                                _ = v39
                            } else {
                                if true {
                                    var v40 uint8 = v5
                                    var v41 uint16 = 0
                                    _ = v40
                                    _ = v41
                                } else {
                                    var v42 uint32 = 2
                                    v1 -= v1
                                    _ = v42
                                }
                                var v43 uint16 = 247
                                v44 := GenS3{F0: (v1 * v1), F1: (fi38 % ((-389) | 1))}
                                _ = v43
                                _ = v44
                            }
                        case "a":
                            switch (fi38 % (((^(fi38))) | 1)) {
                            case 4294967295:
                                v45 := GenS4{F0: fi38, F1: 42, F2: false}
                                _ = v45
                            default:
                                var v46 int64 = 42
                                _ = v46
                            }
                        default:
                            v4 = append(v4, fi38)
                        }
                        _ = fi38
                    }
                }
                if (!((v2 / (((^(v35))) | 1)) == (v2 + (v35 - (v35 % ((v2) | 1)))))) {
                    v2 += 0
                    v2 -= (v35 & 990)
                } else {
                    v2 -= 0
                }
                var v47 uint32 = 981
                _ = v47
            }
            v48 := GenS2{F0: ((v5 % (((v5 >> ((0) & 63))) | 1)) <= v5), F1: 2}
            if len(v4) > 0 {
                v4[0] = (-(fi36))
            }
            _ = fi36
            _ = v48
        }
    }
    v5 += (v6 >> ((2) & 63))
    var v49 float64 = 0.0
    switch -777 {
    case 42:
        if (!((1 != -1) || (2147483647 < 1))) {
            for fi50 := 0; fi50 < 4; fi50++ {
                var v51 float64 = 1e10
                fmt.Println(v2)
                if (true || (-1 <= v35)) {
                    switch "a" {
                    case "\n\t":
                        var v52 uint = 360
                        _ = v52
                    default:
                        var v53 uint = 65535
                        _ = v53
                    }
                    var v54 float32 = ((0.0 + (1e10 + 0.1)) * 1e10)
                    var v55 float64 = 1e10
                    _ = v54
                    _ = v55
                } else {
                    var v56 uint8 = v5
                    _ = v56
                }
                _ = fi50
                _ = v51
            }
            v57 := GenS1{F0: false}
            _ = v57
        } else {
            if false {
                var v58 uint64 = 192
                var v59 uint64 = v58
                v60 := GenS5{F0: v6, F1: "a", F2: -1}
                _ = v58
                _ = v59
                _ = v60
            }
        }
    default:
        var v61 float64 = 1e10
        _ = v61
    }
    var v62 int64 = 42
    v63 := GenS6{F0: 260}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v62
    _ = v63
}
