package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int; F1 uint32; F2 uint64 }
func genFn0(p0 bool, p1 uint16, p2 int16) uint64 {
    p1 -= 2
    v7 := []int{-812, 127, 39, 4294967295}
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: 406, F1: fi8}
        p1 += p1
        if (int64(p2) >= -9223372036854775808) {
            switch "foo bar" {
            case "":
                if (!true) {
                    for fi10 := 0; fi10 < 1; fi10++ {
                        v11 := GenS1{F0: 0, F1: (fi10 % (((fi8 % (((^(fi8))) | 1))) | 1))}
                        _ = fi10
                        _ = v11
                    }
                }
            case "foo bar":
                p2 += p2
            }
        }
        _ = fi8
        _ = v9
    }
    v12 := GenS2{F0: -1, F1: (p1 ^ (p1 >> ((0) & 63))), F2: 65535}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    return 4294967295
}
func genFn1(p0 int64, p1 uint, p2 int8) int32 {
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(p2)
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 int16 = -1
            v16 := []int{(fi13 ^ fi14), fi14, fi13}
            v17 := GenS3{F0: 660, F1: ("foo bar" + "foo bar")}
            fmt.Println(fi14, fi13, fi14)
            var v18 bool = true
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
            _ = v18
        }
        var v19 bool = false
        _ = fi13
        _ = v19
    }
    var v20 int64 = (p0 / (((p0 / (((p0 >> ((p1) & 63))) | 1))) | 1))
    var v21 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 0
}

func main() {
    var v1 int16 = 831
    var v2 float64 = -2.25
    var v3 uint64 = 0
    var v4 float32 = 1e10
    v5 := []int{42, 128}
    v6 := make(map[string]string)
    var v22 float32 = v4
    fmt.Println(v22)
    if (!((v2 * -2.25) != float64(v22))) {
        v23 := v6[""]
        v24 := GenS4{F0: true, F1: 1, F2: 773}
        if true {
            var v25 int = 32768
            _ = v25
        }
        _ = v23
        _ = v24
    }
    delete(v6, "")
    v3 = (^(v3))
    switch -546 {
    case 32767:
        for fi26 := 0; fi26 < 10; fi26++ {
            if false {
                if true {
                    v22 -= v22
                    for fi27 := 0; fi27 < 14; fi27++ {
                        switch fi26 {
                        case 0:
                            fmt.Println(fi26, v1, v22)
                        }
                        v6["foo bar"] = "hello"
                        var v28 int16 = (v1 & -32768)
                        _ = fi27
                        _ = v28
                    }
                } else {
                    var v29 int64 = -1
                    v22 -= v4
                    _ = v29
                }
                if false {
                    v30 := GenS5{F0: -893}
                    _ = v30
                }
            } else {
                if true {
                    if false {
                        if true {
                            for fi31 := 0; fi31 < 20; fi31++ {
                                v1 = (v1 | v1)
                                if (1.5 <= (v4 + (v4 / float32(v2)))) {
                                    var v32 int16 = (v1 & (^(v1)))
                                    switch "" {
                                    case "hello":
                                        switch (fi26 + (fi31 ^ (fi26 ^ fi26))) {
                                        case 2147483647:
                                            v3 = v3
                                        case -723:
                                            v33 := GenS6{F0: 4294967295, F1: 42, F2: 898}
                                            _ = v33
                                        case 99:
                                            v32 = 590
                                        }
                                    case "a":
                                        var v34 uint64 = v3
                                        _ = v34
                                    default:
                                        v3 -= 4294967295
                                    }
                                    _ = v32
                                }
                                v3 -= 695
                                var v35 float32 = (float32(float64(float32(float64(1e10)))) * v4)
                                _ = fi31
                                _ = v35
                            }
                        } else {
                            v36 := GenS7{F0: 2147483648}
                            fmt.Println(fi26, v1, fi26)
                            _ = v36
                        }
                        v6[(((("\n\t" + "\n\t") + ("" + "")) + "foo bar") + "a")] = ("hello" + "hello")
                    }
                    _ = len(v5)
                    if (2147483647 >= 1) {
                        var v37 int32 = int32(v3)
                        v1 = (v1 | v1)
                        _ = v37
                    }
                }
            }
            if true {
                fmt.Println(v3, v4, v3)
                fmt.Println(v3, fi26, v3)
            }
            _ = fi26
        }
    case -930:
        for fi38 := 0; fi38 < 6; fi38++ {
            if (false && (!((fi38 * (fi38 << ((284) & 63))) != (^(fi38))))) {
                for fi39 := 0; fi39 < 6; fi39++ {
                    fmt.Println(v4)
                    var v40 int = 0
                    _ = fi39
                    _ = v40
                }
                var v41 uint8 = 0
                _ = v41
            } else {
                for fi42 := 0; fi42 < 20; fi42++ {
                    v43 := GenS8{F0: ""}
                    _ = fi42
                    _ = v43
                }
            }
            switch ((("hello" + "hello") + "\n\t") + (("foo bar" + "a") + "foo bar")) {
            case "foo bar":
                v44 := GenS9{F0: ("\n\t" + "foo bar")}
                _ = v44
            case "\n\t":
                if len(v5) > 0 {
                    v5[0] = fi38
                }
            case "a":
                var v45 float64 = v2
                _ = v45
            }
            var v46 uint16 = 255
            _ = fi38
            _ = v46
        }
    case 127:
        var v47 int16 = v1
        _ = v47
    default:
        for fi48 := 0; fi48 < 3; fi48++ {
            fmt.Println(v22, v4)
            for fi49 := 0; fi49 < 20; fi49++ {
                v3 = (v3 + (v3 / (((v3 << ((65535) & 63))) | 1)))
                fmt.Println(v2, v2, v4)
                var v50 float32 = v4
                _ = fi49
                _ = v50
            }
            if ((-128 != int8(fi48)) && (376 >= 18446744073709551615)) {
                if ("hello" <= "hello") {
                    var v51 int8 = 0
                    _ = v51
                } else {
                    v22 = 0.0
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v52 := v6["a"]
                _ = v52
            }
            _ = fi48
        }
    }
    v22 += ((0.1 - v4) - v4)
    switch -260 {
    case 127:
        if (239 == 0) {
            if true {
                v4 += (float32((float64((float32(float64(v4)) + ((v4 / (v22 * v22)) + v22))) - (v2 + v2))) + 0.0)
                var v53 uint64 = (v3 << ((255) & 63))
                if (true && true) {
                    v54 := GenS3{F0: (v3 + (v3 << ((18446744073709551615) & 63))), F1: "a"}
                    var v55 uint32 = 0
                    v22 += v4
                    _ = v54
                    _ = v55
                }
                _ = v53
            } else {
                v5 = append(v5, 4294967295)
            }
            v3 += v3
        }
    case -2147483648:
        v2 += (v2 + v2)
    case 101:
        var v56 float32 = float32(v2)
        _ = v56
    }
    v57 := GenS10{F0: 0, F1: 804, F2: v3}
    v3 += genFn0(true, 949, v1)
    v58 := v6["foo bar"]
    if (425 == 255) {
        var v59 int16 = (v1 * v1)
        _ = v59
    }
    var v60 int = 858
    for fi61 := 0; fi61 < 17; fi61++ {
        var v62 float32 = -2.25
        v22 = (-2.25 * v22)
        if ((!true) || false) {
            v22 += 0.1
            if ((42 < 65535) && (!false)) {
                for fi63 := 0; fi63 < 8; fi63++ {
                    v64 := GenS5{F0: genFn1(42, 2, -29)}
                    var v65 uint = 0
                    _ = fi63
                    _ = v64
                    _ = v65
                }
                v60 = v60
            }
        } else {
            v60 += fi61
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = fi61
        _ = v62
    }
    switch v58 {
    case "":
        var v66 string = ((v58 + ((v58 + "hello") + "hello")) + v58)
        _ = v66
    case "foo bar":
        fmt.Println(v4, v60, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
    _ = v60
}
