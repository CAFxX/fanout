package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 bool; F2 string }
type GenS3 struct { F0 string; F1 int; F2 int }
type GenS4 struct { F0 uint; F1 uint32; F2 int64 }
type GenS5 struct { F0 int; F1 int16; F2 string }
type GenS6 struct { F0 uint; F1 uint16; F2 bool }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 string; F1 uint; F2 uint16 }
type GenS9 struct { F0 int64; F1 int64 }
func genFn0(p0 uint32, p1 string) uint64 {
    v6 := []int{628, -573, -2147483648}
    fmt.Println(p0)
    v7 := GenS1{F0: 4294967295}
    switch 4294967295 {
    case -535:
        fmt.Println(p0)
    case 127:
        switch p1 {
        case "hello":
            var v8 uint16 = 65535
            _ = v8
        case "\n\t":
            p0 -= p0
        }
    case 2147483647:
        if (0 < 9223372036854775807) {
            v9 := GenS2{F0: -1, F1: (0 >= 747), F2: p1}
            var v10 uint = 634
            _ = len(v6)
            _ = v10
            _ = v9
        }
    default:
        var v11 uint16 = 282
        _ = v11
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 bool) uint64 {
    p0 = p0
    v12 := GenS3{F0: "\n\t", F1: 1, F2: 587}
    _ = p0
    _ = v12
    return 18446744073709551615
}

func main() {
    var v1 float64 = 3.14159
    var v2 int16 = 42
    v3 := []string{("foo bar" + "")}
    var v4 int8 = -1
    var v5 int8 = -128
    var v13 bool = false
    switch ((("\n\t" + ("foo bar" + "hello")) + "a") + "") {
    case "":
        switch ("hello" + "\n\t") {
        case "a":
            v4 = (^(v4))
        case "foo bar":
            var v14 int = -2147483648
            _ = v14
        default:
            fmt.Println(v4, v1, v1)
        }
    case "\n\t":
        fmt.Println(v4, v1)
    default:
        if v13 {
            var v15 uint16 = 720
            for fi16 := 0; fi16 < 1; fi16++ {
                switch "foo bar" {
                case "a":
                    v17 := GenS4{F0: 65535, F1: 690, F2: -9223372036854775808}
                    _ = v17
                case "foo bar":
                    for fi18 := 0; fi18 < 6; fi18++ {
                        var v19 bool = false
                        v20 := GenS5{F0: fi18, F1: (v2 / ((v2) | 1)), F2: "hello"}
                        v3 = append(v3, "hello")
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                case "":
                    if (!v13) {
                        if (1 == 93) {
                            switch (fi16 << ((42) & 63)) {
                            case 127:
                                _ = len(v3)
                            default:
                                for fi21 := 0; fi21 < 17; fi21++ {
                                    if (-1 == 1) {
                                        v15 = (^(v15))
                                    }
                                    v22 := GenS6{F0: 0, F1: (v15 << ((1) & 63)), F2: (!(v13 && ("a" > ("foo bar" + (("a" + ("\n\t" + ("" + "a"))) + "a")))))}
                                    var v23 uint8 = 128
                                    var v24 uint = 901
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                            }
                            if (42 == 626) {
                                switch (("\n\t" + "\n\t") + "\n\t") {
                                case "a":
                                    fmt.Println(v2, v15, fi16)
                                case "hello":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "foo bar":
                                    v25 := GenS7{F0: "\n\t", F1: (-2.25 != 0.1)}
                                    _ = v25
                                }
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                            } else {
                                if false {
                                    var v26 int8 = 100
                                    if (v13 || v13) {
                                        var v27 uint8 = 1
                                        var v28 uint8 = 0
                                        v29 := GenS8{F0: "\n\t", F1: 2, F2: v15}
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        for fi30 := 0; fi30 < 13; fi30++ {
                                            v31 := GenS8{F0: ("a" + ("a" + (("\n\t" + "hello") + ("foo bar" + "hello")))), F1: 0, F2: v15}
                                            _ = fi30
                                            _ = v31
                                        }
                                    default:
                                        fmt.Println(v4, v13, v15)
                                    }
                                    _ = v26
                                } else {
                                    v15 -= (v15 >> ((565) & 63))
                                }
                            }
                            var v32 string = ((("\n\t" + "") + "a") + "a")
                            _ = v32
                        }
                    } else {
                        var v33 string = ("a" + "hello")
                        fmt.Println(v1, v13, v33)
                        v15 -= (v15 << ((2) & 63))
                        _ = v33
                    }
                }
                _ = fi16
            }
            var v34 int64 = 42
            _ = v15
            _ = v34
        }
    }
    var v35 uint8 = 2
    for fi36 := 0; fi36 < 4; fi36++ {
        if (2147483647 != 1) {
            for fi37 := 0; fi37 < 19; fi37++ {
                v38 := GenS5{F0: fi36, F1: v2, F2: "\n\t"}
                switch -75 {
                case 0:
                    if false {
                        v39 := GenS5{F0: fi37, F1: 32767, F2: ("hello" + "a")}
                        fmt.Println(v13, v1)
                        fmt.Println(v1, v13, v1)
                        _ = v39
                    }
                }
                var v40 int8 = 42
                _ = fi37
                _ = v38
                _ = v40
            }
            var v41 int64 = 9223372036854775807
            _ = v41
        }
        if len(v3) > 0 {
            v3[0] = ("\n\t" + ("\n\t" + ("" + ("" + "\n\t"))))
        }
        _ = fi36
    }
    var v42 uint64 = 1
    v35 -= v35
    var v43 float64 = (0.1 / 0.1)
    var v44 int8 = 100
    for fi45 := 0; fi45 < 1; fi45++ {
        switch fi45 {
        case -9223372036854775808:
            v46 := GenS9{F0: -958, F1: int64(fi45)}
            _ = v46
        case 4294967295:
            for fi47 := 0; fi47 < 20; fi47++ {
                if v13 {
                    if (!(609 > 65535)) {
                        v48 := GenS9{F0: 9223372036854775807, F1: 1}
                        var v49 bool = (18446744073709551615 <= 42)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v48
                        _ = v49
                    }
                } else {
                    v35 = 127
                    v35 -= v35
                }
                var v50 int16 = (v2 | 656)
                var v51 string = "a"
                v3 = append(v3, v51)
                _ = fi47
                _ = v50
                _ = v51
            }
        case 32767:
            if len(v3) > 0 {
                v3[0] = ("a" + ("foo bar" + "a"))
            }
        }
        _ = fi45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
