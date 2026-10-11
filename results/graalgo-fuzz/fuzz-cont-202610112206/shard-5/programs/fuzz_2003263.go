package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int16; F2 string }
type GenS5 struct { F0 int }
func genFn0(p0 int16) uint {
    v5 := GenS1{F0: 99}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 int64, p1 string) uint16 {
    v6 := []int{1}
    p1 = (((p1 + "") + ((p1 + (p1 + "foo bar")) + ("\n\t" + ("\n\t" + (("" + ((("a" + ("" + p1)) + p1) + "a")) + p1))))) + p1)
    p1 = ("hello" + p1)
    v7 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 571
}

func main() {
    var v1 int8 = 42
    var v2 uint16 = 0
    var v3 int16 = -320
    v4 := make(map[string]int)
    if (-380 < -294) {
        v3 = -415
        _ = len(v4)
        for fi8 := 0; fi8 < 17; fi8++ {
            if (!false) {
                v9 := []int{(fi8 >> ((2) & 63)), (fi8 ^ fi8), (-(fi8))}
                v1 += 42
                _ = v9
            }
            _ = fi8
        }
    } else {
        for fi10 := 0; fi10 < 5; fi10++ {
            v2 = v2
            _ = fi10
        }
    }
    var v11 int16 = (v3 * (v3 << ((65535) & 63)))
    if (!(true && false)) {
        for fi12 := 0; fi12 < 4; fi12++ {
            switch (fi12 & 32767) {
            case 42:
                switch "" {
                case "foo bar":
                    if (("\n\t" + "foo bar") >= ("" + ("hello" + "\n\t"))) {
                        for fi13 := 0; fi13 < 19; fi13++ {
                            v14 := []int{4294967295, (fi12 << ((568) & 63)), fi12}
                            var v15 int16 = -1
                            var v16 uint16 = (v2 << ((77) & 63))
                            _ = fi13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    }
                case "":
                    if true {
                        var v17 uint8 = 42
                        var v18 float32 = (0.0 - 0.1)
                        _ = v17
                        _ = v18
                    } else {
                        v19 := GenS2{F0: 202}
                        var v20 int64 = -9223372036854775808
                        _ = v19
                        _ = v20
                    }
                default:
                    v1 = -1
                }
            case 166:
                v21 := []int{}
                _ = v21
            case 1:
                if false {
                    v3 += v3
                }
            }
            switch ("foo bar" + "") {
            case "foo bar":
                v22 := []int{int(v3)}
                _ = v22
            }
            _ = fi12
        }
        var v23 int64 = -1
        _ = v23
    } else {
        if true {
            if (42 > 0) {
                var v24 int = 9223372036854775807
                _ = v24
            }
            v4[""] = -2147483648
        } else {
            v25 := GenS1{F0: 890}
            for fi26 := 0; fi26 < 13; fi26++ {
                switch 127 {
                case -9223372036854775808:
                    if (true || false) {
                        v4[("hello" + "foo bar")] = (fi26 ^ (fi26 % (((fi26 << ((42) & 63))) | 1)))
                    }
                default:
                    v27 := []int{}
                    _ = v27
                }
                _ = fi26
            }
            v1 = -10
            _ = v25
        }
    }
    var v28 bool = true
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        v1 = (^(v1))
    case "foo bar":
        fmt.Println(v3, v3, v11)
    }
    fmt.Println(v2, v11, v28)
    if false {
        v29 := []int{32768}
        switch "a" {
        case "a":
            switch "hello" {
            case "\n\t":
                v30 := GenS1{F0: 4294967295}
                _ = v30
            case "a":
                var v31 uint16 = v2
                _ = v31
            default:
                v32 := GenS3{F0: int8(v3)}
                _ = v32
            }
        case "foo bar":
            var v33 int64 = 295
            _ = v33
        default:
            if true {
                v34 := GenS4{F0: 42, F1: (v11 * (-(v3))), F2: ("" + "a")}
                if ((v28 && (!(false || false))) || v28) {
                    v35 := GenS5{F0: -9223372036854775808}
                    if (!true) {
                        if v28 {
                            switch ("hello" + "") {
                            case "\n\t":
                                v3 = (v11 - v3)
                            case "hello":
                                v2 = 2
                            case "":
                                fmt.Println(v2, v11)
                            default:
                                v36 := GenS4{F0: (v1 - v1), F1: v3, F2: ""}
                                _ = v36
                            }
                        }
                    }
                    v3 -= -64
                    _ = v35
                } else {
                    if (!false) {
                        switch ("foo bar" + "a") {
                        case "\n\t":
                            v37 := GenS2{F0: v2}
                            _ = v37
                        default:
                            delete(v4, "a")
                        }
                    } else {
                        fmt.Println(v2, v11)
                    }
                    var v38 uint8 = 97
                    var v39 uint32 = 546
                    _ = v38
                    _ = v39
                }
                switch -2147483648 {
                case 0:
                    var v40 uint = 255
                    _ = v40
                case -9223372036854775808:
                    v41 := v4["hello"]
                    _ = v41
                case 32768:
                    if (!v28) {
                        var v42 int64 = 5
                        for fi43 := 0; fi43 < 18; fi43++ {
                            for fi44 := 0; fi44 < 16; fi44++ {
                                v45 := v4[""]
                                _ = fi44
                                _ = v45
                            }
                            _ = fi43
                        }
                        _ = v42
                    } else {
                        v46 := v4[""]
                        _ = v46
                    }
                default:
                    v3 += (v11 / ((1) | 1))
                }
                _ = v34
            }
        }
        _ = v29
    }
    v4["a"] = -639
    v11 += v11
    v47 := GenS4{F0: (v1 % ((v1) | 1)), F1: (-(v11)), F2: "a"}
    v2 = v2
    if false {
        switch int(v1) {
        case 128:
            v1 = -128
        }
        v48 := []int{189}
        _ = v48
    }
    var v49 int16 = 1000
    v3 = v49
    var v50 int32 = 28
    fmt.Println(v50)
    switch 1 {
    case 4294967295:
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v47
    _ = v49
    _ = v50
}
