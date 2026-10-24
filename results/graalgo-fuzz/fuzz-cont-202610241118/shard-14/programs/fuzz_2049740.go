package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint16 }
type GenS2 struct { F0 int16; F1 int32; F2 bool }
type GenS3 struct { F0 uint32; F1 string; F2 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 string; F1 uint16; F2 int16 }
type GenS7 struct { F0 int; F1 string }
type GenS8 struct { F0 string; F1 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (((true && (15 != 205)) || (true || (!(!true)))) || ((-1 < -128) || (65535 < 255)))
    v3 := make(map[int]string)
    v4 := GenS1{F0: -6, F1: -9223372036854775808, F2: 53}
    switch -9223372036854775808 {
    case 41:
        var v5 uint16 = 0
        _ = v5
    case -2147483648:
        v2 = true
    case 9223372036854775807:
        var v6 float32 = 0.0
        _ = v6
    default:
        var v7 int16 = 1000
        _ = v7
    }
    if (1e10 >= float32(0.0)) {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int32 = 630
            v10 := GenS1{F0: -1, F1: 0, F2: 65535}
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        var v11 bool = v2
        _ = v11
    }
    if (42 <= 720) {
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 int8 = -128
            v3[(fi12 << ((569) & 63))] = "foo bar"
            var v14 int = (fi12 | fi12)
            _ = fi12
            _ = v13
            _ = v14
        }
        if (2147483647 >= 249) {
            switch "\n\t" {
            case "":
                if false {
                    if (!(!(-2147483648 != 1))) {
                        v2 = v2
                    } else {
                        v2 = (!(v2 && (99 < 552)))
                        var v15 int16 = 27
                        if v2 {
                            for fi16 := 0; fi16 < 6; fi16++ {
                                v17 := []int{-619, int(v15), 128}
                                v2 = (4294967295 != 883)
                                _ = fi16
                                _ = v17
                            }
                            if v2 {
                                v1[411] = ("hello" + "hello")
                                if (42 == -128) {
                                    fmt.Println(v2, v15)
                                    v18 := GenS2{F0: (v15 << ((705) & 63)), F1: -1, F2: (!((((1.5 - (1e10 - 3.14159)) * 1.5) < 3.14159) && (0 >= 0)))}
                                    fmt.Println(v15)
                                    _ = v18
                                } else {
                                    v19 := []int{}
                                    var v20 float64 = (1e10 * -2.25)
                                    v20 = v20
                                    _ = v19
                                    _ = v20
                                }
                                if (127 > 0) {
                                    switch "" {
                                    case "hello":
                                        v2 = v2
                                    case "\n\t":
                                        v21 := GenS3{F0: 4294967295, F1: "a", F2: 18446744073709551615}
                                        _ = v21
                                    case "":
                                        var v22 uint16 = 32767
                                        _ = v22
                                    default:
                                        fmt.Println(v2, v2)
                                    }
                                }
                            } else {
                                if (2147483647 <= -2147483648) {
                                    switch -1 {
                                    case 32767:
                                        if v2 {
                                            for fi23 := 0; fi23 < 16; fi23++ {
                                                var v24 int = fi23
                                                v25 := GenS3{F0: 42, F1: ("" + (("foo bar" + "hello") + "")), F2: 65535}
                                                v15 = v15
                                                _ = fi23
                                                _ = v24
                                                _ = v25
                                            }
                                            v26 := []int{0, 127, 32767}
                                            fmt.Println(v2, v2)
                                            _ = v26
                                        } else {
                                            for fi27 := 0; fi27 < 3; fi27++ {
                                                v15 -= v15
                                                _ = fi27
                                            }
                                        }
                                    case 42:
                                        switch (("foo bar" + "a") + "") {
                                        case "\n\t":
                                            if ((255 > 42) && v2) {
                                                v28 := GenS4{F0: ""}
                                                v29 := []int{-57}
                                                _ = v28
                                                _ = v29
                                            } else {
                                                switch 32767 {
                                                case 32768:
                                                    v30 := GenS5{F0: -1, F1: -1}
                                                    _ = v30
                                                }
                                                v31 := []int{58, -554}
                                                _ = v31
                                            }
                                        case "foo bar":
                                            v2 = (!(65535 == 413))
                                        case "hello":
                                            fmt.Println(v15, v15)
                                        }
                                    case -85:
                                        if v2 {
                                            v15 += v15
                                        }
                                    default:
                                        var v32 bool = (!v2)
                                        _ = v32
                                    }
                                    var v33 uint64 = 4294967295
                                    v2 = ((true && (!false)) || v2)
                                    _ = v33
                                }
                                fmt.Println(v15)
                                var v34 int = 9223372036854775807
                                _ = v34
                            }
                        } else {
                            v35 := GenS1{F0: -1, F1: -1, F2: 42}
                            delete(v3, -1)
                            v2 = (v2 || (6 != -128))
                            _ = v35
                        }
                        _ = v15
                    }
                    for fi36 := 0; fi36 < 13; fi36++ {
                        v2 = (42 == -9223372036854775808)
                        v37 := GenS6{F0: "foo bar", F1: 42, F2: 32767}
                        for fi38 := 0; fi38 < 16; fi38++ {
                            if (4294967295 == 4294967295) {
                                v39 := []int{}
                                _ = v39
                            } else {
                                var v40 float32 = 0.1
                                v41 := GenS7{F0: fi36, F1: "\n\t"}
                                v42 := []int{}
                                if (v2 || false) {
                                    var v43 uint = 221
                                    fmt.Println(v2)
                                    v40 = 1e10
                                    _ = v43
                                } else {
                                    if v2 {
                                        v2 = v2
                                        v2 = (v2 && ((-494 >= 0) || (646 <= 1)))
                                    } else {
                                        var v44 int64 = 42
                                        var v45 string = "\n\t"
                                        v44 += v44
                                        _ = v44
                                        _ = v45
                                    }
                                }
                                v2 = v2
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                            fmt.Println(fi38)
                            _ = fi38
                        }
                        _ = fi36
                        _ = v37
                    }
                }
            default:
                var v46 float32 = 1.5
                _ = v46
            }
            var v47 uint64 = 709
            _ = len(v3)
            _ = v47
        }
    }
    v2 = (42 >= 0)
    if v2 {
        fmt.Println(v2)
        v2 = v2
        if false {
            switch 128 {
            case 0:
                v2 = (644 >= 2)
            case 32768:
                fmt.Println(v2)
            case 128:
                v2 = ((-1 != 0) || ((-2.25 / 0.1) == 0.1))
            default:
                v48 := []int{}
                _ = v48
            }
            var v49 float64 = 1.5
            _ = v49
        } else {
            v2 = (282 == 190)
            v2 = (!(3.14159 <= 1.5))
            switch "\n\t" {
            case "hello":
                delete(v1, 133)
            default:
                v2 = v2
            }
        }
    } else {
        for fi50 := 0; fi50 < 1; fi50++ {
            delete(v1, 2147483647)
            _ = fi50
        }
        v51 := GenS8{F0: "foo bar", F1: 42}
        for fi52 := 0; fi52 < 2; fi52++ {
            fmt.Println(fi52, fi52)
            v53 := []int{}
            _ = fi52
            _ = v53
        }
        _ = v51
    }
    var v54 float64 = 0.1
    switch "" {
    case "hello":
        v55 := []int{127, 1}
        _ = v55
    }
    fmt.Println(v2, v54)
    v2 = ("\n\t" <= "a")
    v56 := GenS1{F0: -84, F1: 814, F2: 46}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v56
}
