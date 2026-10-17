package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int }
func genFn0(p0 bool, p1 bool, p2 uint8) uint {
    p2 = (p2 << ((308) & 63))
    p2 = 157
    v7 := GenS1{F0: 27}
    p2 = (^(p2))
    switch "\n\t" {
    case "a":
        fmt.Println(p0, p2)
    case "":
        p0 = (18446744073709551615 > 0)
    case "\n\t":
        fmt.Println(p2, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1() uint16 {
    var v8 string = "\n\t"
    v8 = ""
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := []int{(-(fi9))}
        fmt.Println(v8)
        _ = fi9
        _ = v10
    }
    var v11 string = v8
    var v12 uint = 4294967295
    _ = v11
    _ = v12
    _ = v8
    return 32767
}

func main() {
    var v1 bool = (42 < 33)
    var v2 int64 = -521
    var v3 string = ""
    var v4 uint16 = 42
    var v5 uint32 = uint32(v4)
    v6 := []string{v3, v3, v3, v3}
    v13 := GenS1{F0: -128}
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 int16 = 0
        var v16 string = ("" + ("a" + ""))
        if len(v6) > 0 {
            v6[0] = (((v16 + "\n\t") + v3) + ((("foo bar" + (v16 + "a")) + v3) + ("hello" + (v16 + v16))))
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (true || v1) {
        v17 := GenS1{F0: -1}
        _ = v17
    } else {
        v18 := GenS1{F0: -1}
        _ = v18
    }
    var v19 string = v3
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 float32 = 3.14159
        var v22 uint16 = v4
        _ = fi20
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: 0}
    for fi24 := 0; fi24 < 19; fi24++ {
        switch (v19 + v19) {
        case "foo bar":
            fmt.Println(fi24)
        case "":
            var v25 string = (v19 + "hello")
            _ = v25
        default:
            if len(v6) > 0 {
                v6[0] = v19
            }
        }
        for fi26 := 0; fi26 < 19; fi26++ {
            if v1 {
                var v27 uint = 4294967295
                fmt.Println(v27, v27)
                _ = v27
            }
            var v28 int32 = -2147483648
            if ((!true) || (0 < (v4 & genFn1()))) {
                _ = len(v6)
                if len(v6) > 0 {
                    v6[0] = "\n\t"
                }
                switch fi24 {
                case -589:
                    var v29 int64 = (v2 % (((^(v2))) | 1))
                    _ = v29
                default:
                    switch 1 {
                    case 42:
                        v30 := GenS2{F0: v2, F1: ((v1 || ((v2 | v2) < v2)) && v1)}
                        _ = v30
                    case 4294967295:
                        v1 = true
                    }
                }
            } else {
                v31 := GenS2{F0: 0, F1: (42 >= 42)}
                v32 := GenS3{F0: ((255 == 866) || (genFn1() <= (v4 * v4)))}
                var v33 uint64 = uint64(v4)
                v28 = 0
                _ = v31
                _ = v32
                _ = v33
            }
            _ = fi26
            _ = v28
        }
        _ = fi24
    }
    switch ("\n\t" + v19) {
    case "":
        var v34 uint8 = 219
        _ = v34
    case "\n\t":
        if (304 < 4294967295) {
            fmt.Println(v3, v4, v3)
            switch v19 {
            case "a":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case "hello":
                v3 = v3
            case "\n\t":
                switch 936 {
                case -587:
                    if (!((1.5 == 1e10) && false)) {
                        var v35 int16 = 0
                        v6 = append(v6, "")
                        var v36 float32 = 1.5
                        _ = v35
                        _ = v36
                    } else {
                        if len(v6) > 0 {
                            v6[0] = v3
                        }
                    }
                case -601:
                    fmt.Println(v2, v3, v19)
                }
            }
            v1 = (!v1)
        } else {
            if (v1 && v1) {
                v37 := GenS1{F0: -13}
                _ = v37
            } else {
                var v38 uint32 = (^(v5))
                var v39 int8 = -128
                v40 := GenS4{F0: 127, F1: 42}
                _ = v38
                _ = v39
                _ = v40
            }
            for fi41 := 0; fi41 < 16; fi41++ {
                var v42 uint8 = uint8(v4)
                v43 := GenS1{F0: 100}
                _ = fi41
                _ = v42
                _ = v43
            }
        }
    case "a":
        v5 -= v5
    default:
        v19 = (v19 + v3)
    }
    fmt.Println(v19, v5, v4)
    switch 128 {
    case 32768:
        for fi44 := 0; fi44 < 3; fi44++ {
            var v45 uint32 = v5
            switch v3 {
            case "\n\t":
                if (!(1000 >= 1)) {
                    switch v3 {
                    case "hello":
                        fmt.Println(v5, v4)
                    case "foo bar":
                        for fi46 := 0; fi46 < 19; fi46++ {
                            v2 = v2
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if v1 {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                v6 = append(v6, (v3 + v3))
                            } else {
                                if (!v1) {
                                    var v47 uint32 = v45
                                    _ = v47
                                }
                                v5 = v45
                            }
                            _ = fi46
                        }
                    default:
                        v2 = int64(v4)
                    }
                    v3 = ((v19 + v3) + v19)
                }
            case "hello":
                if true {
                    var v48 int32 = 2147483647
                    _ = v48
                } else {
                    var v49 int8 = 0
                    if v1 {
                        v50 := GenS3{F0: ((fi44 == fi44) && (!((!v1) || true)))}
                        _ = v50
                    }
                    if (v2 >= (v2 >> ((2) & 63))) {
                        var v51 uint16 = v4
                        _ = v51
                    }
                    _ = v49
                }
            case "a":
                for fi52 := 0; fi52 < 9; fi52++ {
                    fmt.Println(v2)
                    var v53 int = (fi44 + int(v45))
                    switch (fi52 / (((v53 >> ((21) & 63))) | 1)) {
                    case -9223372036854775808:
                        if (77 != 0) {
                            v54 := GenS3{F0: (!v1)}
                            _ = v54
                        }
                    case 127:
                        var v55 float64 = (3.14159 * 1e10)
                        _ = v55
                    case 0:
                        _ = len(v6)
                    }
                    _ = fi52
                    _ = v53
                }
            default:
                if (v3 < v3) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v56 int64 = 0
                    _ = v56
                }
            }
            v4 -= (^(v4))
            _ = fi44
            _ = v45
        }
    case -9223372036854775808:
        switch 4294967295 {
        case 1:
            var v57 uint64 = 42
            _ = v57
        case -220:
            fmt.Println(v2, v3)
        case 42:
            v3 = v3
        }
    case 2147483647:
        var v58 uint = 2
        _ = v58
    default:
        var v59 uint64 = 358
        _ = v59
    }
    v1 = v1
    var v60 float32 = (-2.25 * 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
