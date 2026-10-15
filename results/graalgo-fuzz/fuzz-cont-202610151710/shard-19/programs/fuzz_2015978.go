package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int32; F1 uint; F2 string }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint; F1 uint8; F2 int32 }

func main() {
    var v1 int16 = 1000
    v2 := []uint32{130, 172}
    var v3 uint = 1
    var v4 bool = false
    var v5 uint = v3
    var v6 float32 = 1.5
    for fi7 := 0; fi7 < 15; fi7++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if false {
            _ = len(v2)
            v4 = (!(v4 && ((!(42 != 683)) && ((v4 || v4) && v4))))
            if (1 < -9223372036854775808) {
                v4 = (18446744073709551615 <= 769)
                var v8 int16 = v1
                _ = v8
            } else {
                var v9 string = ("a" + ((("hello" + (((("" + ("" + "hello")) + "\n\t") + ("\n\t" + (("\n\t" + "\n\t") + ""))) + "\n\t")) + (((("\n\t" + ("foo bar" + (((("foo bar" + "foo bar") + "a") + "a") + "foo bar"))) + (("" + (("hello" + "foo bar") + (("foo bar" + ((("a" + ("hello" + (("foo bar" + (("hello" + "") + (("a" + "hello") + ""))) + "a"))) + ("\n\t" + "foo bar")) + "\n\t")) + "a"))) + ("\n\t" + "a"))) + ("foo bar" + "foo bar")) + (("hello" + "\n\t") + "a"))) + "\n\t"))
                _ = v9
            }
        } else {
            var v10 float64 = 3.14159
            if v4 {
                var v11 uint16 = 1
                v5 -= v5
                _ = v11
            } else {
                v12 := GenS1{F0: v5}
                v2 = append(v2, 65535)
                _ = v12
            }
            _ = v10
        }
        _ = fi7
    }
    v4 = (42 > 128)
    for fi13 := 0; fi13 < 20; fi13++ {
        fmt.Println(v1)
        if (!v4) {
            v1 += v1
            switch -1 {
            case 42:
                var v14 int8 = 1
                _ = v14
            case 32768:
                v4 = false
            }
            var v15 int = (-(fi13))
            _ = v15
        }
        v16 := GenS2{F0: (18446744073709551615 >= v5), F1: v4}
        _ = fi13
        _ = v16
    }
    switch ("" + "") {
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        switch ("\n\t" + "") {
        case "":
            v17 := GenS2{F0: v4, F1: v4}
            _ = v17
        case "foo bar":
            for fi18 := 0; fi18 < 19; fi18++ {
                for fi19 := 0; fi19 < 4; fi19++ {
                    if ((729 != 656) && (v4 || v4)) {
                        v20 := GenS3{F0: -1, F1: v3, F2: ((("hello" + "foo bar") + ((("hello" + (("" + "a") + "hello")) + "foo bar") + "foo bar")) + "")}
                        v21 := GenS1{F0: 9}
                        _ = v20
                        _ = v21
                    }
                    if true {
                        if v4 {
                            if v4 {
                                if ("hello" != "foo bar") {
                                    if len(v2) > 0 {
                                        v2[0] = 2147483648
                                    }
                                    var v22 string = ("\n\t" + "a")
                                    v6 = v6
                                    v5 = v5
                                    _ = v22
                                }
                                var v23 bool = (!v4)
                                switch (fi18 * 4294967295) {
                                case 894:
                                    if (false && true) {
                                        v2 = append(v2, 65535)
                                    }
                                default:
                                    v3 -= 518
                                    v24 := GenS4{F0: 71, F1: (!(!(!(v4 || (2147483647 < 2147483647)))))}
                                    _ = v24
                                }
                                v25 := GenS5{F0: (fi18 & fi19), F1: -1}
                                _ = v23
                                _ = v25
                            } else {
                                if (127 >= 158) {
                                    switch fi18 {
                                    case 0:
                                        var v26 string = "foo bar"
                                        _ = v26
                                    }
                                    v5 -= v3
                                    v6 += v6
                                } else {
                                    var v27 uint16 = 188
                                    v28 := GenS3{F0: 42, F1: v5, F2: "a"}
                                    _ = v27
                                    _ = v28
                                }
                            }
                            v1 -= v1
                            var v29 uint16 = 2
                            _ = v29
                        } else {
                            var v30 int = fi18
                            switch (-(v30)) {
                            case 128:
                                v2 = append(v2, 2147483647)
                            }
                            _ = v30
                        }
                        v3 -= (v5 & 255)
                    }
                    var v31 float64 = 1e10
                    _ = fi19
                    _ = v31
                }
                var v32 uint8 = 1
                if false {
                    v32 -= (^(v32))
                    var v33 string = "\n\t"
                    _ = v33
                } else {
                    v32 += 0
                }
                _ = fi18
                _ = v32
            }
        default:
            for fi34 := 0; fi34 < 2; fi34++ {
                var v35 uint16 = 46
                _ = fi34
                _ = v35
            }
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 += (v1 / (((v1 - v1)) | 1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v36 := GenS2{F0: (100 == 100), F1: ((v1 / ((1) | 1)) < (v1 << ((42) & 63)))}
    for fi37 := 0; fi37 < 9; fi37++ {
        var v38 uint64 = 42
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 19; fi39++ {
        fmt.Println(v3, v4, v5)
        if len(v2) > 0 {
            v2[0] = 42
        }
        if len(v2) > 0 {
            v2[0] = 811
        }
        _ = fi39
    }
    for fi40 := 0; fi40 < 18; fi40++ {
        var v41 float64 = 0.1
        _ = fi40
        _ = v41
    }
    var v42 string = "foo bar"
    v43 := GenS6{F0: v3, F1: 0, F2: -1}
    if (9223372036854775807 < -952) {
        if v4 {
            for fi44 := 0; fi44 < 7; fi44++ {
                if (400 <= uint64(v3)) {
                    var v45 int8 = -42
                    _ = v45
                } else {
                    var v46 bool = v4
                    var v47 uint8 = 1
                    _ = v46
                    _ = v47
                }
                switch fi44 {
                case -526:
                    fmt.Println(v4, fi44)
                case 795:
                    v1 -= -748
                case 1:
                    if v4 {
                        switch fi44 {
                        case -880:
                            fmt.Println(v42, v4, fi44)
                        }
                        var v48 bool = v4
                        _ = v48
                    }
                default:
                    fmt.Println(v1)
                }
                fmt.Println(fi44)
                _ = fi44
            }
            v2 = append(v2, 65535)
            var v49 bool = (v4 || v4)
            _ = v49
        }
        v50 := GenS6{F0: v5, F1: 127, F2: -1}
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
