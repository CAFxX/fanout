package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int; F1 uint16 }
func genFn0(p0 bool) int {
    if (1 >= 134) {
        p0 = (1000 >= 42)
    }
    for fi4 := 0; fi4 < 15; fi4++ {
        p0 = (p0 && (((false && (32767 > 0)) && true) && (748 == 0)))
        v5 := []int{fi4, (fi4 | fi4), fi4}
        p0 = (p0 || p0)
        _ = fi4
        _ = v5
    }
    if (32767 > 32767) {
        if p0 {
            p0 = (736 <= 42)
        } else {
            v6 := GenS1{F0: 32767, F1: 32767}
            v7 := GenS2{F0: (!true), F1: 2}
            _ = v6
            _ = v7
        }
        v8 := []int{}
        if len(v8) > 0 {
            v8[0] = -1
        }
        _ = v8
    } else {
        v9 := GenS2{F0: (!true), F1: 0}
        p0 = p0
        _ = v9
    }
    v10 := []int{2147483647, 42, 32767, -9223372036854775808}
    _ = p0
    _ = v10
    return -2147483648
}

func main() {
    v1 := []int{316, 2147483647, 9223372036854775807}
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS3{F0: (^(fi11)), F1: 0}
        if (!(2 != 2)) {
            for fi13 := 0; fi13 < 4; fi13++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(fi11)
                _ = fi13
            }
            var v14 int8 = 1
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    if (3.14159 == float32(1.5)) {
        var v15 bool = false
        v15 = false
        _ = v15
    }
    var v16 int16 = 0
    if (!(-18 >= 127)) {
        if (1 == -9223372036854775808) {
            if (((("\n\t" + "") + "a") <= "foo bar") || true) {
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 uint = 2
                    _ = fi17
                    _ = v18
                }
                switch ("a" + ("" + "")) {
                case "foo bar":
                    if true {
                        var v19 int64 = 42
                        _ = v19
                    }
                case "\n\t":
                    v3[("a" + "a")] = true
                default:
                    if false {
                        var v20 int8 = 121
                        switch ("\n\t" + "foo bar") {
                        case "hello":
                            _ = len(v1)
                        case "a":
                            if len(v1) > 0 {
                                v1[0] = -9223372036854775808
                            }
                        case "":
                            if (2 == 367) {
                                for fi21 := 0; fi21 < 1; fi21++ {
                                    var v22 uint = 820
                                    var v23 int32 = 2147483647
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                                for fi24 := 0; fi24 < 16; fi24++ {
                                    for fi25 := 0; fi25 < 13; fi25++ {
                                        v26 := v3[(((("foo bar" + ("\n\t" + "a")) + "foo bar") + ("a" + "hello")) + ("\n\t" + ("hello" + ("" + ("foo bar" + "hello")))))]
                                        v1 = append(v1, 32768)
                                        _ = fi25
                                        _ = v26
                                    }
                                    _ = fi24
                                }
                            }
                        default:
                            var v27 int16 = (v16 - (v16 << ((2) & 63)))
                            _ = v27
                        }
                        _ = v20
                    }
                }
                fmt.Println(v16)
            }
        }
        v28 := GenS3{F0: 112, F1: 255}
        _ = v28
    } else {
        if true {
            if (42 == -1) {
                v16 -= v16
                v29 := GenS3{F0: 42, F1: 65535}
                _ = v29
            }
            v16 += 0
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v16 += v16
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    _ = len(v1)
    delete(v3, (("\n\t" + "foo bar") + ""))
    var v30 int32 = 0
    v30 = v30
    if true {
        switch genFn0(false) {
        case 128:
            var v31 string = "foo bar"
            _ = v31
        }
        for fi32 := 0; fi32 < 20; fi32++ {
            if (!((int(v30) <= genFn0(false)) && (694 != 4294967295))) {
                var v33 int16 = -35
                _ = v33
            } else {
                for fi34 := 0; fi34 < 12; fi34++ {
                    var v35 uint64 = 2
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v36 uint16 = 1
                    switch int(v30) {
                    case 2147483647:
                        v1 = append(v1, int(v16))
                    case 128:
                        if false {
                            v37 := GenS3{F0: (fi32 * genFn0(false)), F1: (v36 + 42)}
                            var v38 uint64 = (v35 ^ 765)
                            v39 := GenS2{F0: true, F1: uint(v38)}
                            _ = v37
                            _ = v38
                            _ = v39
                        } else {
                            var v40 uint32 = 840
                            fmt.Println(v40, fi34, v16)
                            var v41 int32 = (^(v30))
                            _ = v40
                            _ = v41
                        }
                    case 32768:
                        delete(v3, "a")
                    }
                    v30 += -1
                    v1 = append(v1, -976)
                    var v42 float64 = (1.5 + 3.14159)
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v42
                }
                var v43 int64 = 42
                _ = v43
            }
            _ = fi32
        }
        var v44 int32 = 0
        _ = v44
    } else {
        if true {
            v30 = v30
            v45 := GenS1{F0: 915, F1: 42}
            v46 := GenS3{F0: -1, F1: 1}
            _ = v45
            _ = v46
        } else {
            for fi47 := 0; fi47 < 12; fi47++ {
                v16 -= v16
                for fi48 := 0; fi48 < 6; fi48++ {
                    v2["hello"] = (!((566 >= 2147483647) || true))
                    v30 = v30
                    v30 -= -2147483648
                    _ = fi48
                }
                _ = fi47
            }
            v49 := GenS3{F0: genFn0((-1 == 0)), F1: 65535}
            _ = v49
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
}
