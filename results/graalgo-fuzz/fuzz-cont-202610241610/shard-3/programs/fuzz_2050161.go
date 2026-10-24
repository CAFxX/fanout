package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
func genFn0() int32 {
    if ((100 == -66) || true) {
        var v7 int64 = -9223372036854775808
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := []int{(fi8 / ((4294967295) | 1)), (fi8 % ((fi8) | 1)), fi8, 4294967295}
            _ = len(v9)
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    if false {
        v10 := []int{1, 1}
        _ = v10
    }
    return -2147483648
}
func genFn1() int {
    var v11 uint = 2
    v12 := []int{32767, 62}
    if true {
        for fi13 := 0; fi13 < 3; fi13++ {
            fmt.Println(fi13, fi13)
            var v14 uint32 = 0
            _ = fi13
            _ = v14
        }
        var v15 uint = (v11 * v11)
        _ = v15
    } else {
        v11 = 1
    }
    _ = v11
    _ = v12
    return -9223372036854775808
}

func main() {
    v1 := []string{(("" + ("" + "a")) + ("a" + "a")), "a", ("" + (((("a" + "") + "a") + "a") + "\n\t")), "foo bar"}
    var v2 int16 = -1
    var v3 int16 = v2
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    var v6 float32 = (0.1 * -2.25)
    for fi16 := 0; fi16 < 16; fi16++ {
        v6 -= -2.25
        v3 = v2
        _ = fi16
    }
    if true {
        v2 -= v2
        for fi17 := 0; fi17 < 3; fi17++ {
            v2 += 42
            var v18 int = fi17
            _ = fi17
            _ = v18
        }
    }
    v6 = ((float32(3.14159) / v6) / ((1e10 / (v6 - (3.14159 / (v6 - ((0.1 * (((3.14159 + 0.0) - float32(1e10)) * v6)) / 0.1))))) / v6))
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 bool = false
        if v20 {
            for fi21 := 0; fi21 < 6; fi21++ {
                if true {
                    if (!(930 < 65535)) {
                        switch fi19 {
                        case 32767:
                            v20 = v20
                        case 0:
                            if ("\n\t" > "a") {
                                var v22 float64 = 1e10
                                if (!(float64(-2.25) == (v22 * (v22 * ((float64(v6) - v22) + (-2.25 - (((1e10 / v22) - 1e10) / 3.14159))))))) {
                                    _ = len(v1)
                                } else {
                                    var v23 int64 = -1
                                    v24 := GenS1{F0: 104, F1: -2147483648}
                                    _ = v23
                                    _ = v24
                                }
                                v22 = v22
                                _ = v22
                            }
                        }
                        if (("foo bar" + ("foo bar" + "foo bar")) == "a") {
                            v3 = (v2 / ((v3) | 1))
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v6)
                            }
                            v5["hello"] = false
                        }
                        v25 := GenS1{F0: 1, F1: 42}
                        _ = v25
                    } else {
                        v4["\n\t"] = (fi19 >> ((753) & 63))
                        var v26 uint8 = 255
                        var v27 bool = v20
                        _ = v26
                        _ = v27
                    }
                }
                _ = fi21
            }
            var v28 uint64 = 4294967295
            var v29 int16 = (v2 | (v3 - 297))
            _ = v28
            _ = v29
        } else {
            switch -708 {
            case 417:
                switch (fi19 * genFn1()) {
                case 630:
                    for fi30 := 0; fi30 < 1; fi30++ {
                        v31 := GenS1{F0: 1, F1: 42}
                        var v32 int32 = 1
                        var v33 uint64 = 1
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                case -1:
                    fmt.Println(v20)
                }
            case 9223372036854775807:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case -459:
                v6 += 1e10
            default:
                var v34 int64 = 42
                _ = v34
            }
            v3 = v3
            var v35 uint32 = 778
            _ = v35
        }
        _ = fi19
        _ = v20
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        _ = len(v1)
        _ = fi36
    }
    v6 = (v6 - v6)
    var v37 uint = 65535
    var v38 int = 32768
    v39 := GenS1{F0: 255, F1: 0}
    for fi40 := 0; fi40 < 7; fi40++ {
        v37 = 4294967295
        _ = fi40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
