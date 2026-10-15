package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint32 }

func main() {
    v1 := []int{0, 32767, -958, 4294967295}
    v2 := []string{}
    var v3 float32 = -2.25
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 uint8 = 0
        _ = fi4
        _ = v5
    }
    if true {
        if (!false) {
            var v6 uint32 = 2147483647
            for fi7 := 0; fi7 < 10; fi7++ {
                v8 := GenS1{F0: 32767}
                if (true && true) {
                    var v9 uint16 = 1
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
            _ = v6
        }
        var v10 int8 = -49
        v3 = v3
        _ = v10
    } else {
        fmt.Println(v3)
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS1{F0: 1}
        _ = fi11
        _ = v12
    }
    if false {
        fmt.Println(v3)
        v13 := GenS2{F0: -1}
        for fi14 := 0; fi14 < 17; fi14++ {
            v3 += v3
            v2 = append(v2, ("foo bar" + "hello"))
            v3 = v3
            _ = fi14
        }
        _ = v13
    } else {
        v3 = (v3 / v3)
        var v15 uint = 18446744073709551615
        _ = v15
    }
    if (!(40 == 127)) {
        var v16 float64 = 0.1
        var v17 int64 = 0
        v17 += (v17 - (v17 ^ (v17 ^ v17)))
        _ = v16
        _ = v17
    } else {
        var v18 bool = (!(!true))
        var v19 float64 = -2.25
        v3 += v3
        _ = v18
        _ = v19
    }
    v2 = append(v2, "hello")
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        switch (fi20 ^ fi20) {
        case 494:
            var v21 uint16 = 255
            _ = v21
        case 128:
            if (!true) {
                _ = len(v1)
                v3 -= v3
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        default:
            v3 -= v3
        }
        _ = fi20
    }
    v22 := GenS3{F0: -920}
    if (0 >= 2147483647) {
        var v23 uint = 2
        if (false || false) {
            v24 := GenS3{F0: -32768}
            _ = v24
        }
        if false {
            if true {
                for fi25 := 0; fi25 < 18; fi25++ {
                    var v26 string = "\n\t"
                    v27 := GenS4{F0: (v23 % ((495) | 1)), F1: ("\n\t" + v26)}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                switch ("hello" + "a") {
                case "foo bar":
                    v3 = 1e10
                case "":
                    v28 := GenS5{F0: 1, F1: 127}
                    _ = v28
                case "a":
                    v29 := GenS5{F0: 42, F1: 82}
                    _ = v29
                }
                if true {
                    if false {
                        if false {
                            switch "a" {
                            case "hello":
                                for fi30 := 0; fi30 < 19; fi30++ {
                                    if false {
                                        fmt.Println(v3, fi30, fi30)
                                        fmt.Println(v3, v23, v23)
                                    }
                                    _ = fi30
                                }
                            }
                        } else {
                            if (2 <= 536) {
                                fmt.Println(v23, v23)
                            }
                            if false {
                                v3 -= (float32(-2.25) / v3)
                            } else {
                                var v31 uint32 = 2147483647
                                _ = v31
                            }
                        }
                    } else {
                        _ = len(v2)
                        v3 -= v3
                    }
                    if false {
                        v32 := GenS6{F0: 2147483647}
                        for fi33 := 0; fi33 < 7; fi33++ {
                            var v34 uint64 = 0
                            _ = len(v1)
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    }
                    for fi35 := 0; fi35 < 16; fi35++ {
                        for fi36 := 0; fi36 < 16; fi36++ {
                            v37 := GenS6{F0: -1}
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = ("" + ("a" + "a"))
                                }
                                var v38 bool = false
                                if len(v1) > 0 {
                                    v1[0] = 42
                                }
                                _ = v38
                            } else {
                                var v39 uint16 = 255
                                _ = len(v2)
                                _ = v39
                            }
                            v3 = 0.0
                            _ = fi36
                            _ = v37
                        }
                        v40 := GenS5{F0: 67, F1: 0}
                        if (!false) {
                            v41 := GenS7{F0: 0}
                            _ = v41
                        }
                        _ = fi35
                        _ = v40
                    }
                }
            } else {
                if (!(true && (("a" + "a") >= ("a" + "a")))) {
                    for fi42 := 0; fi42 < 13; fi42++ {
                        var v43 int16 = -1
                        if (false && (false || (true || false))) {
                            v44 := GenS4{F0: (v23 * (v23 << ((65535) & 63))), F1: ""}
                            _ = len(v1)
                            _ = v44
                        }
                        v43 += (v43 / ((-1) | 1))
                        _ = fi42
                        _ = v43
                    }
                    for fi45 := 0; fi45 < 17; fi45++ {
                        v46 := GenS7{F0: 439}
                        v47 := GenS8{F0: 1}
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 12; fi48++ {
                        if false {
                            var v49 uint32 = 65535
                            _ = v49
                        } else {
                            v50 := GenS9{F0: 2147483648}
                            v23 = v23
                            var v51 int32 = 1
                            _ = v50
                            _ = v51
                        }
                        _ = fi48
                    }
                } else {
                    fmt.Println(v23)
                    v23 = v23
                    var v52 uint = v23
                    _ = v52
                }
                v53 := GenS1{F0: 1000}
                _ = v53
            }
        }
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
}
