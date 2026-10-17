package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 int; F2 uint }
type GenS5 struct { F0 int64 }
func genFn0() uint64 {
    if (220 <= -561) {
        var v7 int32 = 1
        if false {
            fmt.Println(v7)
            var v8 uint64 = 4294967295
            v8 = (v8 << ((42) & 63))
            _ = v8
        }
        _ = v7
    }
    var v9 uint16 = 32767
    var v10 uint16 = v9
    var v11 uint8 = 255
    _ = v10
    _ = v11
    _ = v9
    return 961
}

func main() {
    var v1 int32 = 42
    v2 := make(map[string]string)
    var v3 int64 = 9223372036854775807
    var v4 int32 = v1
    var v5 float64 = (float64(0.1) / (1e10 * float64(-2.25)))
    var v6 int8 = 42
    v1 = v4
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 uint = 18446744073709551615
            _ = len(v2)
            var v15 int8 = (-(v6))
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    switch -2147483648 {
    case 4294967295:
        for fi16 := 0; fi16 < 11; fi16++ {
            v17 := []int{}
            _ = fi16
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 19; fi18++ {
            v4 = 2147483647
            v19 := GenS1{F0: v3, F1: 32767, F2: (v6 / ((v6) | 1))}
            var v20 uint = 160
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v21 := GenS2{F0: 1000}
    if (2 == 2) {
        var v22 string = "\n\t"
        fmt.Println(v22)
        v4 -= v1
        _ = v22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        v1 -= (v4 * v4)
        if true {
            delete(v2, "a")
            fmt.Println(v5)
        } else {
            var v24 uint = 312
            if true {
                var v25 uint64 = genFn0()
                v26 := GenS1{F0: v3, F1: -1, F2: v6}
                _ = v25
                _ = v26
            }
            _ = v24
        }
        _ = fi23
    }
    switch ("\n\t" + "a") {
    case "foo bar":
        fmt.Println(v5)
    case "a":
        switch "hello" {
        case "\n\t":
            v5 -= v5
        case "":
            if (1 == 515) {
                v1 += (^(v1))
            }
        case "foo bar":
            for fi27 := 0; fi27 < 10; fi27++ {
                var v28 uint16 = 517
                if ("a" == (("" + "foo bar") + ("hello" + "a"))) {
                    var v29 bool = true
                    v6 = v6
                    v3 += (-(v3))
                    _ = v29
                } else {
                    v30 := GenS1{F0: 9223372036854775807, F1: 1000, F2: v6}
                    for fi31 := 0; fi31 < 6; fi31++ {
                        fmt.Println(fi31, v5)
                        v32 := GenS3{F0: (v6 / ((v6) | 1))}
                        var v33 int16 = 1000
                        v34 := []int{-9223372036854775808, (fi31 | (fi27 * fi31)), 1, (fi31 + (fi27 << ((2) & 63)))}
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
        default:
            if (false || (0 <= 596)) {
                v35 := GenS1{F0: (^(v3)), F1: 1, F2: v6}
                for fi36 := 0; fi36 < 20; fi36++ {
                    v3 -= 929
                    _ = fi36
                }
                v37 := GenS3{F0: (^(v6))}
                _ = v35
                _ = v37
            } else {
                var v38 string = ((("" + "a") + "hello") + ("foo bar" + "hello"))
                _ = v38
            }
        }
    }
    if (!true) {
        var v39 float32 = (1.5 * 1.5)
        switch "" {
        case "\n\t":
            v40 := GenS1{F0: 0, F1: 0, F2: v6}
            _ = v40
        default:
            if (4294967295 >= 42) {
                v1 -= 42
                if true {
                    switch ("foo bar" + "") {
                    case "\n\t":
                        v2["hello"] = "a"
                    case "foo bar":
                        delete(v2, ("a" + ("a" + "a")))
                    }
                    v41 := GenS4{F0: 210, F1: 2147483647, F2: 1}
                    _ = v41
                }
            } else {
                v42 := []int{42, 2147483647, -9223372036854775808}
                if (((v5 + (v5 + 0.0)) == 1.5) && false) {
                    delete(v2, "hello")
                    v43 := GenS5{F0: v3}
                    v42 = append(v42, 1)
                    _ = v43
                } else {
                    var v44 string = (("foo bar" + ("a" + "\n\t")) + (("a" + "hello") + "hello"))
                    _ = v44
                }
                switch -1 {
                case 651:
                    var v45 uint16 = 819
                    _ = v45
                case 32767:
                    var v46 int64 = (v3 ^ (v3 / ((v3) | 1)))
                    _ = v46
                case 1:
                    fmt.Println(v4)
                }
                _ = v42
            }
        }
        _ = v39
    }
    fmt.Println(v3, v3, v6)
    for fi47 := 0; fi47 < 6; fi47++ {
        v48 := []int{-1, (-(fi47)), (fi47 / ((fi47) | 1)), (fi47 << ((256) & 63))}
        v49 := GenS1{F0: (-(v3)), F1: 0, F2: v6}
        if true {
            switch 128 {
            case 42:
                for fi50 := 0; fi50 < 9; fi50++ {
                    var v51 float64 = 0.0
                    _ = fi50
                    _ = v51
                }
            case 2147483647:
                if false {
                    if len(v48) > 0 {
                        v48[0] = (fi47 >> ((18446744073709551615) & 63))
                    }
                    var v52 float64 = v5
                    switch (^(fi47)) {
                    case 4294967295:
                        for fi53 := 0; fi53 < 5; fi53++ {
                            var v54 float32 = (0.1 + float32(v5))
                            v1 = v1
                            var v55 int16 = 1
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                    case 1:
                        var v56 int8 = 44
                        _ = v56
                    case -9223372036854775808:
                        for fi57 := 0; fi57 < 8; fi57++ {
                            var v58 uint32 = 123
                            var v59 float64 = float64(float32(v52))
                            _ = fi57
                            _ = v58
                            _ = v59
                        }
                    default:
                        for fi60 := 0; fi60 < 12; fi60++ {
                            fmt.Println(v3)
                            _ = fi60
                        }
                    }
                    _ = v52
                }
            case 456:
                var v61 uint32 = 421
                _ = v61
            }
        } else {
            var v62 uint8 = 234
            v62 = v62
            var v63 int8 = (v6 & 127)
            _ = v62
            _ = v63
        }
        _ = fi47
        _ = v48
        _ = v49
    }
    v64 := GenS4{F0: 188, F1: 32767, F2: 20}
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
