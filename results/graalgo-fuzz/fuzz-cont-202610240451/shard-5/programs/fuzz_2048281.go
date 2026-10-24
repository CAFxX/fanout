package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 int; F2 uint8 }
type GenS7 struct { F0 int8; F1 int }
func genFn0(p0 uint64) bool {
    v5 := []int{865, -515, -9223372036854775808, 127}
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS1{F0: 940}
            _ = len(v5)
            var v10 int8 = 1
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    _ = p0
    _ = v5
    return (-2.25 > 1.5)
}
func genFn1() uint32 {
    var v11 string = (("hello" + "") + ("foo bar" + ""))
    v11 = v11
    var v12 int32 = 42
    v13 := []int{9223372036854775807, 2147483647}
    _ = v11
    _ = v12
    _ = v13
    return 216
}

func main() {
    v1 := []int{}
    v2 := []int{9223372036854775807}
    var v3 uint32 = 0
    var v4 uint32 = v3
    if (-2.25 <= 0.1) {
        var v14 int64 = 1
        if (6 != 100) {
            var v15 int = -2147483648
            fmt.Println(v14, v3, v4)
            _ = v15
        }
        v14 = v14
        _ = v14
    }
    switch 127 {
    case 9223372036854775807:
        if true {
            var v16 float32 = 1.5
            switch "\n\t" {
            case "\n\t":
                v17 := GenS2{F0: 243}
                _ = v17
            case "":
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 uint64 = 463
                    _ = fi18
                    _ = v19
                }
            }
            _ = v16
        } else {
            v2 = append(v2, 9223372036854775807)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case 127:
        v3 = v3
    case 997:
        if false {
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 uint32 = 478
                if ((826 < 1) || (!false)) {
                    var v22 float32 = (-2.25 * (0.1 - (1e10 * -2.25)))
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 uint32 = 42
                v4 = genFn1()
                _ = fi23
                _ = v24
            }
        } else {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        }
    default:
        if len(v2) > 0 {
            v2[0] = 32768
        }
    }
    if (!true) {
        for fi25 := 0; fi25 < 9; fi25++ {
            v26 := GenS3{F0: 901}
            switch ((("\n\t" + "a") + "foo bar") + (((("foo bar" + "hello") + "hello") + ("\n\t" + ("\n\t" + "hello"))) + ((("hello" + ("foo bar" + "a")) + ("\n\t" + "hello")) + "a"))) {
            case "\n\t":
                for fi27 := 0; fi27 < 6; fi27++ {
                    v3 = v3
                    if (9223372036854775807 >= -280) {
                        v3 -= (^(v3))
                        v4 = (v3 % (((^(v4))) | 1))
                        if ((1.5 - 1e10) >= (0.0 / (1e10 * (1.5 * -2.25)))) {
                            if false {
                                switch (fi27 << ((2) & 63)) {
                                case -539:
                                    var v28 int8 = 123
                                    _ = v28
                                case 32767:
                                    if (((467 != 0) || false) || true) {
                                        _ = len(v1)
                                        v4 += (v4 & (^(v4)))
                                        switch fi25 {
                                        case -2147483648:
                                            _ = len(v2)
                                        case 660:
                                            var v29 int32 = -2147483648
                                            _ = v29
                                        }
                                    } else {
                                        v4 = 65535
                                        v4 = (v4 >> ((18446744073709551615) & 63))
                                    }
                                }
                                switch "\n\t" {
                                case "":
                                    v3 += 2
                                case "foo bar":
                                    if false {
                                        v30 := GenS4{F0: 18446744073709551615, F1: false}
                                        _ = v30
                                    }
                                case "\n\t":
                                    var v31 uint32 = genFn1()
                                    _ = v31
                                }
                            }
                        } else {
                            switch ("\n\t" + "a") {
                            case "\n\t":
                                var v32 uint16 = 312
                                _ = v32
                            case "foo bar":
                                switch fi25 {
                                case 128:
                                    if (("a" > "a") || false) {
                                        var v33 uint32 = (^(v4))
                                        v34 := GenS5{F0: 1}
                                        _ = v33
                                        _ = v34
                                    } else {
                                        fmt.Println(fi25)
                                        v4 = genFn1()
                                        v35 := GenS6{F0: genFn1(), F1: (fi27 | fi25), F2: 206}
                                        var v36 int = fi25
                                        _ = v35
                                        _ = v36
                                    }
                                }
                            }
                            var v37 int = fi25
                            _ = len(v2)
                            switch ("foo bar" + "") {
                            case "":
                                switch v37 {
                                case 32768:
                                    v4 += (^(v3))
                                default:
                                    if (false && (9223372036854775807 == -1)) {
                                        v38 := GenS1{F0: v3}
                                        v3 = genFn1()
                                        _ = v38
                                    } else {
                                        if (65535 <= 421) {
                                            fmt.Println(fi25, v4, fi25)
                                            if ((false && (-186 != 0)) && true) {
                                                var v39 uint8 = 42
                                                switch "foo bar" {
                                                case "hello":
                                                    switch v37 {
                                                    case -162:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                case "":
                                                    var v40 float32 = (3.14159 * (3.14159 * 1e10))
                                                    fmt.Println(fi25, v3)
                                                    _ = v40
                                                }
                                                _ = v39
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = -1
                                                }
                                                var v41 int = (fi27 / ((fi25) | 1))
                                                _ = v41
                                            }
                                        }
                                        var v42 string = ("\n\t" + "foo bar")
                                        _ = v42
                                    }
                                }
                            case "hello":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "foo bar":
                                var v43 float32 = (((1e10 - 1e10) / float32(3.14159)) * 1.5)
                                _ = v43
                            }
                            switch "foo bar" {
                            case "a":
                                var v44 uint16 = 878
                                _ = v44
                            }
                            _ = v37
                        }
                    }
                    _ = fi27
                }
            }
            var v45 bool = (!false)
            _ = fi25
            _ = v26
            _ = v45
        }
        for fi46 := 0; fi46 < 13; fi46++ {
            if len(v1) > 0 {
                v1[0] = fi46
            }
            v47 := GenS7{F0: 42, F1: (fi46 >> ((282) & 63))}
            v4 += (v3 + 42)
            _ = fi46
            _ = v47
        }
        _ = len(v2)
    }
    v3 += (v4 | 1)
    var v48 int16 = -235
    v2 = append(v2, 32767)
    v48 -= -1
    v48 = (v48 >> ((42) & 63))
    var v49 uint16 = 0
    if (true && false) {
        var v50 string = ("a" + "\n\t")
        var v51 int64 = 0
        _ = v50
        _ = v51
    } else {
        v52 := GenS1{F0: v4}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
}
