package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 bool }
func genFn0(p0 int32) int {
    var v5 uint32 = 1
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: "", F1: (p0 / (((^(p0))) | 1)), F2: (-2.25 <= float64(-2.25))}
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: "\n\t", F1: p0, F2: false}
            if false {
                fmt.Println(fi8, p0)
            }
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    switch -286 {
    case 4294967295:
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 float64 = 0.0
            if (!(!true)) {
                if (((0 >= 2) && ((1.5 != -2.25) && ((p0 & p0) < -231))) && false) {
                    v12 := []int{}
                    v12 = append(v12, (^(fi10)))
                    _ = v12
                }
            }
            p0 = p0
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 float32 = (3.14159 - 3.14159)
    var v2 int8 = 100
    var v3 int32 = 42
    var v4 bool = (0.1 >= 3.14159)
    if (100 == v2) {
        v13 := []int{32767, -1, -741}
        _ = v13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 float64 = 1e10
            _ = fi15
            _ = v16
        }
        switch ("" + "") {
        case "foo bar":
            var v17 float64 = (3.14159 - 0.0)
            _ = v17
        }
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 int64 = 0
            var v20 float64 = 0.0
            var v21 uint32 = 65535
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        _ = fi14
    }
    var v22 uint32 = 1
    v23 := GenS1{F0: ("a" + ((("\n\t" + ("" + ("hello" + "hello"))) + (((("" + ("hello" + "a")) + "") + "") + ("\n\t" + "\n\t"))) + "foo bar")), F1: v3, F2: v4}
    v2 = v2
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 uint16 = 353
        var v26 int64 = 1
        var v27 uint8 = 42
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    switch -2147483648 {
    case -9223372036854775808:
        if true {
            v22 = (^(v22))
            var v28 uint = 4294967295
            if v4 {
                v29 := GenS1{F0: "", F1: (v3 ^ (v3 - v3)), F2: ((!((!(v1 <= v1)) || ("foo bar" == ("foo bar" + "foo bar")))) && false)}
                _ = v29
            } else {
                v30 := []int{}
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
                _ = v30
            }
            _ = v28
        }
    case 9223372036854775807:
        if ((18446744073709551615 > 4294967295) || ((944 >= 808) || ((v22 ^ (v22 - 2147483648)) != (v22 - v22)))) {
            for fi31 := 0; fi31 < 18; fi31++ {
                switch (fi31 % ((fi31) | 1)) {
                case -477:
                    var v32 int16 = -653
                    _ = v32
                case -616:
                    v2 = v2
                case 801:
                    if (1 != 694) {
                        v33 := []int{fi31, genFn0(42), fi31}
                        if true {
                            var v34 float64 = 0.0
                            _ = len(v33)
                            _ = v34
                        }
                        var v35 int16 = 32767
                        _ = v33
                        _ = v35
                    } else {
                        v1 = v1
                    }
                default:
                    v36 := GenS1{F0: ("foo bar" + ((("a" + "a") + (("foo bar" + "\n\t") + "foo bar")) + (("" + "a") + "hello"))), F1: v3, F2: (fi31 > (fi31 | (fi31 / ((genFn0(v3)) | 1))))}
                    _ = v36
                }
                _ = fi31
            }
        }
    case -108:
        fmt.Println(v2, v2)
    default:
        v37 := GenS1{F0: "", F1: 42, F2: (v4 || false)}
        _ = v37
    }
    var v38 int = -53
    v39 := []int{(-(v38)), genFn0(v3)}
    v2 = (v2 | (v2 / ((v2) | 1)))
    if (!v4) {
        var v40 int = genFn0(1)
        _ = v40
    } else {
        switch (v38 + (v38 ^ 128)) {
        case 450:
            _ = len(v39)
        case 2147483647:
            v3 += -2147483648
        default:
            var v41 int64 = -1
            _ = v41
        }
    }
    fmt.Println(v38, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
