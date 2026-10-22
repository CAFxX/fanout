package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int16; F2 uint }
func genFn0() uint32 {
    var v6 float64 = 3.14159
    v7 := GenS1{F0: 243, F1: ("foo bar" + (((("a" + ("" + "hello")) + "foo bar") + "") + "foo bar"))}
    var v8 uint8 = 170
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int16, p1 uint64) uint16 {
    switch 0 {
    case 1:
        p1 = p1
    case 844:
        switch "hello" {
        case "foo bar":
            p1 = (p1 & (p1 >> ((972) & 63)))
        default:
            var v9 int64 = 0
            _ = v9
        }
    }
    var v10 bool = (309 != 531)
    var v11 int32 = 1
    fmt.Println(p1, v10)
    var v12 bool = ((2 == 905) || v10)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 string = ""
    v2 := []uint32{0, 860}
    v3 := make(map[uint32]bool)
    var v4 uint16 = 0
    var v5 uint64 = 4294967295
    var v13 uint16 = v4
    v1 = ("" + v1)
    v13 += v13
    if true {
        fmt.Println(v4, v1)
        v2 = append(v2, 2147483647)
    } else {
        var v14 uint16 = v4
        var v15 uint32 = 2
        var v16 int32 = -1
        _ = v14
        _ = v15
        _ = v16
    }
    v4 = (v4 / (((v4 >> ((65535) & 63))) | 1))
    if (false && false) {
        v13 += (^(v13))
    } else {
        var v17 uint8 = 2
        _ = v17
    }
    if (true || (true || (false || (1 < v5)))) {
        if true {
            switch ("foo bar" + "a") {
            case "a":
                if false {
                    for fi18 := 0; fi18 < 8; fi18++ {
                        var v19 bool = (((!(!true)) && true) || true)
                        _ = fi18
                        _ = v19
                    }
                    var v20 bool = (!true)
                    _ = v20
                }
            case "foo bar":
                v21 := GenS2{F0: (!false)}
                _ = v21
            }
            if (!(v5 == v5)) {
                fmt.Println(v1)
                for fi22 := 0; fi22 < 18; fi22++ {
                    if true {
                        v13 = 64
                        if (!false) {
                            v2 = append(v2, 715)
                            for fi23 := 0; fi23 < 13; fi23++ {
                                fmt.Println(v13, v5)
                                if ((^(v5)) == v5) {
                                    v4 = genFn1(874, v5)
                                    delete(v3, genFn0())
                                }
                                v5 -= 18446744073709551615
                                _ = fi23
                            }
                        }
                        var v24 int32 = -2147483648
                        _ = v24
                    }
                    fmt.Println(fi22)
                    switch (("a" + (v1 + "")) + ((v1 + "foo bar") + "a")) {
                    case "hello":
                        switch v1 {
                        case "":
                            if (-971 <= -1000) {
                                v1 = v1
                                for fi25 := 0; fi25 < 6; fi25++ {
                                    fmt.Println(v5, v4, v5)
                                    var v26 int32 = -654
                                    _ = fi25
                                    _ = v26
                                }
                                switch (v1 + v1) {
                                case "a":
                                    v27 := v3[65535]
                                    _ = v27
                                }
                            } else {
                                switch (v1 + "foo bar") {
                                case "foo bar":
                                    var v28 uint64 = 947
                                    _ = v28
                                case "":
                                    if true {
                                        var v29 int32 = -1
                                        v30 := GenS3{F0: (v1 + v1)}
                                        _ = v29
                                        _ = v30
                                    }
                                case "hello":
                                    var v31 uint = 18446744073709551615
                                    _ = v31
                                }
                                v4 -= v4
                            }
                        case "foo bar":
                            var v32 int = (fi22 & (fi22 + (fi22 | fi22)))
                            _ = v32
                        case "a":
                            var v33 int16 = -701
                            _ = v33
                        default:
                            v34 := GenS4{F0: v5, F1: 32767, F2: 2}
                            _ = v34
                        }
                    }
                    _ = fi22
                }
                v5 = 93
            }
        }
    }
    if (true || (!((("foo bar" + v1) >= "") || (1 >= 255)))) {
        var v35 string = (v1 + "\n\t")
        for fi36 := 0; fi36 < 7; fi36++ {
            for fi37 := 0; fi37 < 3; fi37++ {
                var v38 uint8 = 189
                _ = fi37
                _ = v38
            }
            _ = fi36
        }
        v4 += (v4 ^ v4)
        _ = v35
    } else {
        for fi39 := 0; fi39 < 13; fi39++ {
            var v40 int8 = 42
            var v41 uint64 = (v5 / ((v5) | 1))
            delete(v3, 74)
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        fmt.Println(v13, v5)
        v13 = genFn1(306, (v5 ^ v5))
        _ = fi42
    }
    v43 := GenS3{F0: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
