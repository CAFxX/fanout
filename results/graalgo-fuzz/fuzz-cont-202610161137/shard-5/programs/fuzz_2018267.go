package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS4 struct { F0 bool; F1 int; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool; F1 int8 }
func genFn0(p0 bool) bool {
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := []int{fi5, 2147483647, -241, -2147483648}
        for fi7 := 0; fi7 < 5; fi7++ {
            if len(v6) > 0 {
                v6[0] = (fi7 | -2147483648)
            }
            p0 = ((!((-(fi5)) < (^(fi5)))) || (1 == 578))
            var v8 string = "hello"
            switch "\n\t" {
            case "hello":
                var v9 uint8 = 128
                _ = v9
            case "foo bar":
                var v10 uint16 = 0
                _ = v10
            case "":
                var v11 int64 = 9223372036854775807
                _ = v11
            default:
                _ = len(v6)
            }
            v12 := GenS1{F0: 127}
            _ = fi7
            _ = v12
            _ = v8
        }
        var v13 uint16 = 666
        _ = fi5
        _ = v13
        _ = v6
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := []int{42}
        switch fi14 {
        case 285:
            p0 = p0
        case -684:
            p0 = (3.14159 >= 3.14159)
        case 1:
            for fi16 := 0; fi16 < 17; fi16++ {
                p0 = p0
                if p0 {
                    var v17 uint32 = 854
                    if true {
                        var v18 uint64 = 771
                        _ = v18
                    } else {
                        if (-1 <= 42) {
                            v19 := GenS1{F0: 127}
                            _ = v19
                        } else {
                            v20 := GenS2{F0: 2, F1: 580}
                            var v21 string = "foo bar"
                            _ = v20
                            _ = v21
                        }
                    }
                    if len(v15) > 0 {
                        v15[0] = (^(fi14))
                    }
                    _ = v17
                }
                switch "" {
                case "\n\t":
                    p0 = (p0 && p0)
                    var v22 int = (fi14 % (((fi14 << ((204) & 63))) | 1))
                    _ = v22
                case "":
                    var v23 float64 = -2.25
                    _ = v23
                }
                _ = fi16
            }
        default:
            fmt.Println(fi14)
        }
        _ = fi14
        _ = v15
    }
    v24 := []int{32767, 9223372036854775807}
    p0 = (4294967295 < 135)
    if p0 {
        fmt.Println(p0)
    } else {
        var v25 uint64 = 0
        _ = v25
    }
    _ = p0
    _ = v24
    return (127 > -11)
}
func genFn1(p0 string, p1 uint) bool {
    v26 := GenS2{F0: p1, F1: 631}
    var v27 int16 = 0
    v28 := GenS3{F0: 378, F1: 1, F2: 2}
    for fi29 := 0; fi29 < 5; fi29++ {
        if false {
            fmt.Println(v27)
            v27 += (-(v27))
            v30 := GenS4{F0: (-846 < 1), F1: -153, F2: 105}
            _ = v30
        } else {
            fmt.Println(p1)
        }
        _ = fi29
    }
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    _ = v28
    return (!false)
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    var v3 string = ("" + "")
    var v4 bool = false
    fmt.Println(v3)
    switch v3 {
    case "":
        if (587 != 1) {
            fmt.Println(v3)
            v4 = v4
        }
    case "hello":
        switch 2147483647 {
        case 9223372036854775807:
            switch 0 {
            case 539:
                switch v3 {
                case "foo bar":
                    v4 = v4
                case "hello":
                    v4 = (v4 || (!(true && (65535 == 32767))))
                default:
                    v31 := GenS1{F0: 0}
                    _ = v31
                }
            }
        case 32768:
            v32 := GenS4{F0: v4, F1: 32767, F2: 4294967295}
            _ = v32
        }
    case "\n\t":
        for fi33 := 0; fi33 < 4; fi33++ {
            v3 = (("foo bar" + "hello") + "hello")
            v4 = (512 != 65535)
            for fi34 := 0; fi34 < 8; fi34++ {
                fmt.Println(fi33, v3, fi34)
                _ = fi34
            }
            _ = fi33
        }
    }
    fmt.Println(v4, v4)
    v35 := v1[-408]
    switch v3 {
    case "":
        for fi36 := 0; fi36 < 1; fi36++ {
            if ((v35 / (((fi36 << ((0) & 63))) | 1)) != v35) {
                v4 = (1 > 42)
                var v37 uint64 = 1
                fmt.Println(v35)
                _ = v37
            }
            v4 = (1000 <= -1)
            for fi38 := 0; fi38 < 14; fi38++ {
                v3 = v3
                _ = fi38
            }
            _ = fi36
        }
    default:
        if (!true) {
            v3 = v3
            var v39 uint16 = 2
            for fi40 := 0; fi40 < 16; fi40++ {
                v35 += v35
                v39 -= (v39 + v39)
                _ = fi40
            }
            _ = v39
        }
    }
    v41 := GenS3{F0: 42, F1: 42, F2: 1}
    v3 = v3
    if false {
        if (v4 && (1 <= -320)) {
            switch ("hello" + v3) {
            case "":
                v42 := []int{}
                _ = v42
            }
            fmt.Println(v3, v35)
        } else {
            v43 := GenS5{F0: -9223372036854775808}
            if (1.5 > 0.1) {
                for fi44 := 0; fi44 < 1; fi44++ {
                    switch "\n\t" {
                    case "a":
                        for fi45 := 0; fi45 < 5; fi45++ {
                            v4 = ((690 != 9223372036854775807) && (!(!v4)))
                            v46 := []int{(fi45 & v35), (v35 - 32768)}
                            var v47 int32 = -864
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                    default:
                        for fi48 := 0; fi48 < 17; fi48++ {
                            v49 := GenS6{F0: 0}
                            _ = fi48
                            _ = v49
                        }
                    }
                    fmt.Println(v4)
                    _ = fi44
                }
                for fi50 := 0; fi50 < 16; fi50++ {
                    v51 := []int{(^(v35)), v35, 127}
                    _ = fi50
                    _ = v51
                }
                for fi52 := 0; fi52 < 11; fi52++ {
                    if false {
                        v53 := []int{(^(v35)), v35, v35, 32767}
                        v54 := GenS7{F0: (v4 || v4), F1: 42}
                        _ = v53
                        _ = v54
                    } else {
                        v4 = false
                    }
                    var v55 uint32 = 65535
                    switch (("a" + (((v3 + ((v3 + (((v3 + v3) + v3) + "a")) + v3)) + v3) + v3)) + v3) {
                    case "hello":
                        v56 := []int{fi52}
                        _ = v56
                    }
                    _ = fi52
                    _ = v55
                }
            } else {
                v57 := GenS4{F0: true, F1: v35, F2: 4294967295}
                v58 := GenS1{F0: 128}
                _ = v57
                _ = v58
            }
            _ = v43
        }
        switch ((v3 + (((v3 + v3) + (v3 + ((v3 + v3) + v3))) + "foo bar")) + v3) {
        case "hello":
            for fi59 := 0; fi59 < 11; fi59++ {
                var v60 int = (fi59 - fi59)
                _ = fi59
                _ = v60
            }
        default:
            fmt.Println(v4, v3)
        }
        v61 := GenS7{F0: (2 <= 255), F1: -128}
        _ = v61
    }
    for fi62 := 0; fi62 < 8; fi62++ {
        var v63 float32 = 3.14159
        if (v4 && true) {
            v35 = (-(v35))
        } else {
            if v4 {
                v3 = ""
                v2[v35] = (v35 << ((1) & 63))
                v4 = (-746 <= -1)
            }
            for fi64 := 0; fi64 < 5; fi64++ {
                fmt.Println(v3)
                _ = fi64
            }
        }
        _ = fi62
        _ = v63
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v41
}
