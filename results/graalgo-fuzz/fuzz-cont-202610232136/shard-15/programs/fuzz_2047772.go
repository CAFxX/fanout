package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint16; F1 int8; F2 string }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint64; F2 int64 }
type GenS7 struct { F0 int16; F1 int32 }
func genFn0(p0 int64, p1 string, p2 uint8) int {
    v7 := GenS1{F0: 541, F1: false}
    fmt.Println(p0, p2)
    p2 += 1
    p1 = ((p1 + "") + "")
    fmt.Println(p2, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (true || (369 >= 843))
    v3 := []int{246}
    v4 := make(map[string]bool)
    v5 := make(map[uint32]bool)
    v6 := []uint32{4294967295, 418, 909, 65535}
    if ((4294967295 == 2) && true) {
        var v8 int32 = -1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v9 := GenS2{F0: 994, F1: -128, F2: "a"}
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            switch ((("a" + ("" + "foo bar")) + "\n\t") + ("" + "\n\t")) {
            case "hello":
                switch fi10 {
                case 231:
                    if v2 {
                        var v11 int16 = 0
                        switch ("hello" + "") {
                        case "":
                            switch "a" {
                            case "hello":
                                v2 = v2
                            case "\n\t":
                                var v12 uint8 = 0
                                _ = v12
                            }
                        case "\n\t":
                            for fi13 := 0; fi13 < 9; fi13++ {
                                v11 -= 932
                                _ = fi13
                            }
                        }
                        _ = v11
                    }
                default:
                    var v14 int64 = 578
                    _ = v14
                }
            case "":
                v2 = false
            default:
                for fi15 := 0; fi15 < 5; fi15++ {
                    v2 = ((!(929 != -194)) || v2)
                    _ = fi15
                }
            }
            if ((((float64(1e10) + ((0.0 / 1.5) / -2.25)) + 1e10) != 0.1) || ((214 > 52) && false)) {
                if v2 {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        fmt.Println(fi10, v2)
                        v2 = (!v2)
                        _ = fi16
                    }
                    v2 = true
                    var v17 string = "\n\t"
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        _ = len(v4)
                        if v2 {
                            _ = len(v1)
                        } else {
                            switch "\n\t" {
                            case "":
                                var v19 float64 = -2.25
                                _ = v19
                            }
                            switch "a" {
                            case "hello":
                                switch 628 {
                                case 128:
                                    v2 = true
                                }
                            case "":
                                var v20 int = fi18
                                _ = v20
                            default:
                                fmt.Println(fi18, v2)
                            }
                            switch fi18 {
                            case 32768:
                                var v21 string = "\n\t"
                                _ = v21
                            case 0:
                                var v22 int8 = 42
                                _ = v22
                            case 75:
                                switch 0 {
                                case 387:
                                    switch "\n\t" {
                                    case "hello":
                                        if false {
                                            v23 := GenS1{F0: 1, F1: (!v2)}
                                            _ = v23
                                        } else {
                                            v2 = v2
                                        }
                                    case "":
                                        v24 := GenS1{F0: 65535, F1: (-32768 > 0)}
                                        _ = v24
                                    case "a":
                                        var v25 int64 = -1
                                        _ = v25
                                    }
                                default:
                                    var v26 int32 = -1000
                                    _ = v26
                                }
                            }
                        }
                        v27 := GenS3{F0: 2, F1: 65535}
                        v28 := GenS2{F0: 42, F1: -111, F2: "a"}
                        v2 = v2
                        var v29 int = fi10
                        _ = fi18
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    var v30 float64 = 3.14159
                    for fi31 := 0; fi31 < 10; fi31++ {
                        var v32 bool = v2
                        switch (((("foo bar" + "a") + "a") + "a") + "\n\t") {
                        case "hello":
                            v2 = false
                        default:
                            v30 = 0.1
                        }
                        switch ("\n\t" + ("hello" + ("\n\t" + "foo bar"))) {
                        case "hello":
                            switch "a" {
                            case "a":
                                _ = len(v6)
                            default:
                                v33 := v4[("\n\t" + ("a" + "hello"))]
                                _ = v33
                            }
                        case "foo bar":
                            fmt.Println(v30, fi10)
                        }
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                }
                var v34 int = fi10
                var v35 int = (v34 * genFn0(133, "a", 42))
                _ = v34
                _ = v35
            }
            _ = fi10
        }
        for fi36 := 0; fi36 < 12; fi36++ {
            for fi37 := 0; fi37 < 6; fi37++ {
                v38 := GenS4{F0: -2147483648}
                if true {
                    var v39 float32 = (1.5 / (float32(0.0) - (0.1 - float32((3.14159 * (-2.25 / 3.14159))))))
                    _ = v39
                }
                fmt.Println(fi37, fi37)
                v40 := GenS1{F0: 652, F1: (!(-2.25 != float64((0.0 * (0.1 / 1e10)))))}
                _ = fi37
                _ = v38
                _ = v40
            }
            for fi41 := 0; fi41 < 20; fi41++ {
                if (255 != 0) {
                    v42 := GenS5{F0: -316}
                    var v43 int = fi36
                    v44 := GenS4{F0: 0}
                    _ = v42
                    _ = v43
                    _ = v44
                } else {
                    v1["hello"] = 65535
                    fmt.Println(fi36)
                }
                if (!v2) {
                    if v2 {
                        v2 = v2
                        v45 := GenS6{F0: fi41, F1: 0, F2: 1}
                        var v46 uint8 = 124
                        _ = v45
                        _ = v46
                    }
                    if len(v3) > 0 {
                        v3[0] = (-(fi36))
                    }
                    if v2 {
                        v2 = v2
                        var v47 uint32 = 250
                        fmt.Println(v2, fi36)
                        _ = v47
                    }
                } else {
                    var v48 bool = ((("foo bar" + "\n\t") + "\n\t") < "")
                    _ = v48
                }
                v2 = v2
                _ = fi41
            }
            _ = fi36
        }
        var v49 float64 = float64(0.0)
        _ = v49
    }
    fmt.Println(v2)
    if v2 {
        var v50 int16 = 0
        if true {
            var v51 int = genFn0(-710, "hello", 42)
            switch (v51 ^ (v51 + (v51 - (v51 & v51)))) {
            case -656:
                v51 = v51
            case -1:
                _ = len(v6)
            case 9223372036854775807:
                switch (-(v51)) {
                case 0:
                    v2 = v2
                default:
                    fmt.Println(v51, v51, v50)
                }
            }
            if v2 {
                v51 = 1
                if len(v6) > 0 {
                    v6[0] = 42
                }
            }
            _ = v51
        }
        v2 = v2
        _ = v50
    } else {
        v52 := GenS1{F0: 255, F1: false}
        v2 = v2
        v1["a"] = 617
        _ = v52
    }
    v1["a"] = 0
    if v2 {
        v2 = (-204 <= 940)
        fmt.Println(v2)
    }
    _ = len(v4)
    for fi53 := 0; fi53 < 14; fi53++ {
        v2 = (459 == 578)
        _ = fi53
    }
    v2 = (!(!(1 == -1)))
    v54 := GenS7{F0: -32768, F1: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
