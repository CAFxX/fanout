package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint; F2 bool }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int16; F1 string; F2 uint16 }
type GenS6 struct { F0 int16; F1 int16; F2 uint32 }
func genFn0(p0 int64, p1 int, p2 int16) string {
    var v5 int32 = 42
    switch "\n\t" {
    case "foo bar":
        if false {
            if true {
                if true {
                    v5 = -974
                } else {
                    var v6 int8 = 1
                    if ((-(p2)) == p2) {
                        switch p1 {
                        case 304:
                            v7 := GenS1{F0: v6, F1: 402, F2: (p0 ^ (p0 >> ((65535) & 63)))}
                            _ = v7
                        }
                        if (142 > 0) {
                            v8 := GenS2{F0: (p2 << ((415) & 63))}
                            _ = v8
                        } else {
                            p1 = 1
                            v9 := []int{p1, p1, (^(p1)), (^(p1))}
                            if len(v9) > 0 {
                                v9[0] = -80
                            }
                            _ = v9
                        }
                    }
                    _ = v6
                }
                p1 += p1
            }
        } else {
            var v10 uint64 = 982
            for fi11 := 0; fi11 < 8; fi11++ {
                v5 -= -2147483648
                var v12 string = ("" + "a")
                v13 := GenS3{F0: -128, F1: 42, F2: (p0 == p0)}
                _ = fi11
                _ = v12
                _ = v13
            }
            p2 = (p2 / ((42) | 1))
            _ = v10
        }
    case "\n\t":
        p2 += (^(p2))
    case "hello":
        if false {
            fmt.Println(v5, v5, p2)
            if false {
                for fi14 := 0; fi14 < 15; fi14++ {
                    if ((-(p0)) == p0) {
                        var v15 uint = 307
                        _ = v15
                    } else {
                        v16 := []int{(fi14 ^ (fi14 / ((-9223372036854775808) | 1))), p1, fi14, 16}
                        if len(v16) > 0 {
                            v16[0] = p1
                        }
                        _ = v16
                    }
                    _ = fi14
                }
                p1 = (p1 + 42)
                var v17 uint16 = 780
                _ = v17
            } else {
                if (true && false) {
                    var v18 float64 = 0.1
                    _ = v18
                } else {
                    v19 := []int{32767}
                    p1 += 4294967295
                    v19 = append(v19, (p1 * p1))
                    _ = v19
                }
                var v20 int8 = -128
                v21 := GenS4{F0: 2147483648, F1: 0}
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 12; fi22++ {
                fmt.Println(v5)
                var v23 uint16 = 757
                _ = fi22
                _ = v23
            }
        } else {
            var v24 uint16 = 94
            var v25 int32 = v5
            if false {
                v26 := []int{(p1 | 671), (p1 << ((65535) & 63)), p1, (p1 ^ 32767)}
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                switch ("a" + "\n\t") {
                case "hello":
                    v5 = v5
                case "":
                    v25 = v5
                }
                _ = v26
            }
            _ = v24
            _ = v25
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (("foo bar" + "hello") + "")
}

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    v4 := make(map[string]bool)
    if (v1 && v1) {
        if v1 {
            v27 := []int{}
            var v28 float64 = 1e10
            _ = v27
            _ = v28
        } else {
            if (v1 && (-2.25 == 1e10)) {
                var v29 int64 = 158
                v1 = ((980 < -1) || v1)
                v30 := GenS1{F0: 1, F1: 722, F2: v29}
                _ = v29
                _ = v30
            }
        }
        v31 := GenS5{F0: -32768, F1: ((("a" + "hello") + "") + "hello"), F2: 255}
        _ = v31
    } else {
        var v32 uint8 = 1
        _ = v32
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        switch "foo bar" {
        case "foo bar":
            v1 = true
        case "\n\t":
            for fi34 := 0; fi34 < 10; fi34++ {
                v1 = v1
                var v35 uint8 = 1
                _ = fi34
                _ = v35
            }
        default:
            for fi36 := 0; fi36 < 12; fi36++ {
                if ((131 >= 0) && (42 < 1)) {
                    v37 := GenS2{F0: 75}
                    v1 = true
                    var v38 uint8 = 128
                    delete(v4, (("hello" + "a") + "foo bar"))
                    _ = v37
                    _ = v38
                }
                var v39 int = fi36
                switch (fi33 + (^(fi36))) {
                case 30:
                    fmt.Println(v39, fi33)
                case 2147483647:
                    fmt.Println(v1, fi33, v39)
                case 128:
                    var v40 uint64 = 0
                    _ = v40
                }
                if v1 {
                    v39 -= 4294967295
                    v1 = false
                }
                _ = fi36
                _ = v39
            }
        }
        for fi41 := 0; fi41 < 10; fi41++ {
            if v1 {
                var v42 uint64 = 2
                v1 = v1
                v43 := GenS6{F0: 1000, F1: 32767, F2: 2147483648}
                _ = v42
                _ = v43
            }
            v44 := GenS3{F0: 28, F1: 784, F2: v1}
            _ = fi41
            _ = v44
        }
        _ = fi33
    }
    for fi45 := 0; fi45 < 2; fi45++ {
        v1 = v1
        var v46 uint16 = 65535
        v47 := GenS4{F0: 2147483647, F1: 813}
        _ = fi45
        _ = v46
        _ = v47
    }
    if (-1 < -1) {
        if v1 {
            v48 := GenS1{F0: 127, F1: 1, F2: -9223372036854775808}
            _ = v48
        }
    }
    var v49 uint = 147
    for fi50 := 0; fi50 < 16; fi50++ {
        v1 = false
        v1 = (759 < 396)
        _ = fi50
    }
    v49 -= (v49 % (((^(v49))) | 1))
    var v51 int = 42
    if (3.14159 != 0.0) {
        if (-9223372036854775808 != 386) {
            for fi52 := 0; fi52 < 4; fi52++ {
                if true {
                    delete(v4, (("hello" + (((("foo bar" + "\n\t") + (("foo bar" + "a") + ("hello" + "\n\t"))) + (((("hello" + ("\n\t" + "hello")) + "") + (("a" + "hello") + "a")) + "foo bar")) + ("foo bar" + "hello"))) + ""))
                } else {
                    v1 = v1
                    for fi53 := 0; fi53 < 12; fi53++ {
                        fmt.Println(v51)
                        if v1 {
                            var v54 string = ((("foo bar" + ("hello" + "a")) + (("a" + "hello") + "a")) + "hello")
                            _ = v54
                        }
                        _ = fi53
                    }
                }
                v55 := []int{fi52}
                if v1 {
                    delete(v2, "")
                    v49 = 42
                    v56 := GenS3{F0: 127, F1: 2, F2: (2 > 0)}
                    _ = v56
                } else {
                    if (!(v51 > v51)) {
                        v1 = (-2147483648 >= -391)
                        var v57 string = "foo bar"
                        _ = v57
                    }
                    if (!(1 >= 1000)) {
                        var v58 bool = v1
                        _ = v58
                    }
                }
                _ = fi52
                _ = v55
            }
            fmt.Println(v51, v1, v49)
            for fi59 := 0; fi59 < 7; fi59++ {
                v49 -= v49
                _ = fi59
            }
        }
        var v60 uint32 = 1
        _ = v60
    } else {
        fmt.Println(v1)
    }
    v51 = v51
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v51
}
