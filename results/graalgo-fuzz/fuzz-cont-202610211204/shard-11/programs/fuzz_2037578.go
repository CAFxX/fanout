package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 int64; F1 int8; F2 uint16 }
type GenS7 struct { F0 int; F1 int }
func genFn0() int {
    var v7 int64 = 42
    switch 2147483647 {
    case 251:
        var v8 int16 = -649
        _ = v8
    case 2147483647:
        fmt.Println(v7)
    case 32767:
        var v9 uint16 = 2
        _ = v9
    default:
        v7 -= (v7 % ((237) | 1))
    }
    var v10 float32 = 1.5
    var v11 int32 = 42
    v12 := GenS1{F0: "", F1: 573, F2: -1}
    _ = v10
    _ = v11
    _ = v12
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint32 = 565
    var v2 bool = true
    v3 := []string{"", "\n\t", "hello"}
    var v4 int = -279
    var v5 int8 = 1
    v6 := make(map[uint32]string)
    switch (("\n\t" + ("foo bar" + "")) + ("a" + "")) {
    case "foo bar":
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := GenS1{F0: ("foo bar" + ("a" + "")), F1: 1, F2: 260}
            if true {
                var v15 int8 = v5
                if (true && v2) {
                    for fi16 := 0; fi16 < 11; fi16++ {
                        switch ("\n\t" + "a") {
                        case "hello":
                            fmt.Println(v15, v4)
                        }
                        var v17 uint16 = 178
                        switch ("" + "") {
                        case "foo bar":
                            switch "a" {
                            case "a":
                                v18 := GenS1{F0: "hello", F1: 32767, F2: 42}
                                _ = v18
                            }
                        case "a":
                            v19 := GenS2{F0: (v17 >> ((4294967295) & 63)), F1: (v1 | v1), F2: 0}
                            _ = v19
                        case "":
                            v15 -= -128
                        default:
                            fmt.Println(fi13)
                        }
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        v20 := GenS1{F0: ("\n\t" + "hello"), F1: 0, F2: -1}
                        var v21 int = 1
                        v22 := GenS1{F0: "\n\t", F1: 0, F2: -1}
                        fmt.Println(v17, v4)
                        _ = fi16
                        _ = v17
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    delete(v6, (^(v1)))
                } else {
                    v23 := GenS2{F0: 42, F1: 0, F2: 18446744073709551615}
                    _ = v23
                }
                _ = v15
            } else {
                v2 = (!(!(208 < 127)))
                v6[(v1 >> ((0) & 63))] = "a"
                fmt.Println(v2, v4, v1)
            }
            _ = fi13
            _ = v14
        }
    default:
        switch "" {
        case "a":
            var v24 int16 = 0
            _ = v24
        case "\n\t":
            v25 := GenS3{F0: 999, F1: 127, F2: v1}
            _ = v25
        default:
            v4 -= -99
        }
    }
    v26 := GenS3{F0: 2, F1: 2, F2: v1}
    if (int64(v1) == 832) {
        var v27 uint16 = uint16(v1)
        v5 = v5
        v2 = false
        _ = v27
    } else {
        switch "hello" {
        case "a":
            fmt.Println(v5, v5, v5)
        case "foo bar":
            v2 = ((3.14159 <= 3.14159) && (!(((0 <= 4294967295) || false) || v2)))
        default:
            if (!(!v2)) {
                for fi28 := 0; fi28 < 10; fi28++ {
                    v4 = (v4 - -1)
                    if (9223372036854775807 <= 720) {
                        var v29 int = (fi28 - 566)
                        _ = v29
                    } else {
                        var v30 int64 = -583
                        v31 := GenS4{F0: (fi28 % (((fi28 / ((fi28) | 1))) | 1))}
                        _ = v30
                        _ = v31
                    }
                    var v32 int16 = 1
                    _ = fi28
                    _ = v32
                }
                if (true || ((0 >= 212) && (v2 || false))) {
                    fmt.Println(v2, v2)
                    if (!((!false) || v2)) {
                        var v33 int8 = (^(v5))
                        switch (-(v4)) {
                        case 128:
                            var v34 bool = (42 > 1)
                            _ = v34
                        case 882:
                            v33 -= v5
                        case -953:
                            v1 += v1
                        }
                        _ = v33
                    }
                }
            }
        }
        var v35 string = ""
        var v36 int64 = 42
        _ = v35
        _ = v36
    }
    switch "\n\t" {
    case "hello":
        _ = len(v6)
    default:
        for fi37 := 0; fi37 < 1; fi37++ {
            v1 += 4294967295
            for fi38 := 0; fi38 < 6; fi38++ {
                if v2 {
                    v39 := GenS2{F0: 2, F1: (v1 >> ((255) & 63)), F2: 1}
                    var v40 string = "a"
                    _ = v39
                    _ = v40
                } else {
                    switch genFn0() {
                    case 42:
                        var v41 bool = (int8(v1) >= (v5 << ((255) & 63)))
                        _ = v41
                    case 1:
                        fmt.Println(fi38)
                        v42 := GenS5{F0: 3, F1: -188}
                        v4 = (fi37 - 819)
                        _ = v42
                    case -1:
                        v5 = v5
                    default:
                        switch 4294967295 {
                        case -1:
                            var v43 uint16 = 2
                            _ = v43
                        case -196:
                            v3 = append(v3, ("\n\t" + "foo bar"))
                        case -599:
                            v44 := GenS4{F0: v4}
                            _ = v44
                        default:
                            v4 = fi37
                        }
                    }
                    var v45 uint64 = 1
                    v4 += fi38
                    var v46 int64 = -292
                    v2 = false
                    v47 := GenS6{F0: 1, F1: -2, F2: 1}
                    _ = v45
                    _ = v46
                    _ = v47
                }
                _ = fi38
            }
            _ = fi37
        }
    }
    var v48 float64 = 0.0
    var v49 string = "a"
    v50 := GenS3{F0: 284, F1: 128, F2: 65535}
    if (!(2147483647 > v1)) {
        v51 := GenS6{F0: -470, F1: (^(v5)), F2: 1}
        v5 += v5
        _ = v51
    } else {
        _ = len(v3)
        for fi52 := 0; fi52 < 9; fi52++ {
            var v53 float64 = v48
            _ = fi52
            _ = v53
        }
        if v2 {
            var v54 uint = 255
            _ = v54
        } else {
            v4 -= -2147483648
            var v55 uint64 = 18446744073709551615
            _ = v55
        }
    }
    _ = len(v6)
    var v56 int16 = 0
    fmt.Println(v48)
    switch "\n\t" {
    case "hello":
        var v57 uint8 = 2
        _ = v57
    case "a":
        v58 := GenS7{F0: genFn0(), F1: (v4 >> ((42) & 63))}
        _ = v58
    default:
        _ = len(v3)
    }
    var v59 uint32 = v1
    v56 += v56
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v60 int64 = -9223372036854775808
    switch v49 {
    case "\n\t":
        var v61 uint = 387
        _ = v61
    case "":
        var v62 string = "a"
        _ = v62
    case "hello":
        for fi63 := 0; fi63 < 18; fi63++ {
            for fi64 := 0; fi64 < 15; fi64++ {
                switch "hello" {
                case "":
                    switch "a" {
                    case "hello":
                        v65 := GenS3{F0: 18446744073709551615, F1: 128, F2: v1}
                        _ = v65
                    }
                }
                fmt.Println(v4)
                _ = fi64
            }
            v66 := v6[v59]
            _ = fi63
            _ = v66
        }
    default:
        if len(v3) > 0 {
            v3[0] = v49
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v56
    _ = v59
    _ = v6
    _ = v60
}
