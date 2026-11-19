package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int8; F2 int16 }
type GenS4 struct { F0 int8; F1 int32; F2 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint; F1 uint8 }
type GenS8 struct { F0 string }
func genFn0() int8 {
    var v5 uint64 = 42
    fmt.Println(v5)
    _ = v5
    return 1
}
func genFn1(p0 uint8, p1 uint32) uint {
    var v6 uint = 2
    p0 += 1
    fmt.Println(p1)
    var v7 string = ((("" + ("" + "hello")) + "") + (((("\n\t" + "") + "hello") + "") + ("foo bar" + (("\n\t" + "hello") + "hello"))))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (v6 * 956)
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 0
    var v3 uint = 1
    v4 := []uint32{0, 126}
    if len(v4) > 0 {
        v4[0] = 42
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v3 = 2
        _ = fi8
    }
    var v9 uint16 = 32767
    switch "\n\t" {
    case "a":
        var v10 int64 = 1
        _ = v10
    default:
        fmt.Println(v9)
    }
    var v11 uint = genFn1(1, 1)
    var v12 float64 = 1.5
    switch "hello" {
    case "":
        v12 += v12
    case "hello":
        if true {
            if (false && true) {
                fmt.Println(v3, v11, v2)
            }
        } else {
            v13 := GenS1{F0: 42}
            switch 2147483647 {
            case -303:
                fmt.Println(v12, v11, v2)
            case 0:
                v3 -= (^(v3))
            }
            v14 := GenS2{F0: "\n\t"}
            _ = v13
            _ = v14
        }
    case "foo bar":
        if (int16(v3) < (v2 * -425)) {
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS3{F0: v3, F1: -72, F2: (v2 ^ v2)}
                _ = fi15
                _ = v16
            }
            switch 53 {
            case 127:
                switch 4294967295 {
                case -9223372036854775808:
                    for fi17 := 0; fi17 < 3; fi17++ {
                        for fi18 := 0; fi18 < 18; fi18++ {
                            var v19 uint16 = (v9 | (^(v9)))
                            if false {
                                v4 = append(v4, uint32(v19))
                            }
                            fmt.Println(v11, v11, fi18)
                            _ = fi18
                            _ = v19
                        }
                        var v20 uint16 = v9
                        for fi21 := 0; fi21 < 11; fi21++ {
                            switch -535 {
                            case -256:
                                v9 -= v20
                            }
                            _ = fi21
                        }
                        _ = fi17
                        _ = v20
                    }
                default:
                    v22 := GenS4{F0: -128, F1: 2147483647, F2: v11}
                    _ = v22
                }
            case -2147483648:
                if false {
                    for fi23 := 0; fi23 < 15; fi23++ {
                        v12 = 0.0
                        _ = fi23
                    }
                }
            default:
                v24 := GenS5{F0: v9}
                _ = v24
            }
        } else {
            switch "hello" {
            case "hello":
                for fi25 := 0; fi25 < 5; fi25++ {
                    for fi26 := 0; fi26 < 19; fi26++ {
                        _ = len(v1)
                        v3 = 147
                        var v27 string = ("a" + ("\n\t" + ("\n\t" + "\n\t")))
                        _ = fi26
                        _ = v27
                    }
                    var v28 int16 = (^(v2))
                    fmt.Println(v2, v11)
                    _ = fi25
                    _ = v28
                }
            case "":
                if len(v4) > 0 {
                    v4[0] = 632
                }
            case "a":
                for fi29 := 0; fi29 < 11; fi29++ {
                    _ = len(v4)
                    var v30 int64 = 42
                    _ = fi29
                    _ = v30
                }
            }
            if ((!(529 != 42)) && true) {
                v31 := GenS4{F0: -106, F1: 42, F2: (^(v11))}
                _ = v31
            } else {
                switch "\n\t" {
                case "a":
                    switch (((("foo bar" + "\n\t") + ("a" + "")) + "hello") + "") {
                    case "foo bar":
                        v32 := GenS6{F0: 42}
                        _ = v32
                    case "a":
                        fmt.Println(v12, v3, v2)
                    case "":
                        switch ((((("\n\t" + "\n\t") + "") + (("hello" + "hello") + "")) + ("" + ("a" + ((("\n\t" + ("hello" + "hello")) + ("\n\t" + ("a" + ""))) + "\n\t")))) + ("" + "a")) {
                        case "foo bar":
                            var v33 int8 = -128
                            _ = v33
                        case "hello":
                            var v34 bool = (846 < v11)
                            _ = v34
                        default:
                            fmt.Println(v9)
                        }
                    default:
                        for fi35 := 0; fi35 < 7; fi35++ {
                            if (false || true) {
                                v2 += v2
                            }
                            _ = fi35
                        }
                    }
                }
            }
        }
    default:
        var v36 uint16 = (v9 >> ((2) & 63))
        _ = v36
    }
    switch -111 {
    case 2147483647:
        v4 = append(v4, 155)
    case 4294967295:
        for fi37 := 0; fi37 < 12; fi37++ {
            v12 = v12
            fmt.Println(v3)
            v11 = (v3 & 4294967295)
            _ = fi37
        }
    default:
        var v38 string = (("foo bar" + ("hello" + "\n\t")) + "foo bar")
        _ = v38
    }
    if (false || true) {
        var v39 bool = false
        var v40 uint8 = 127
        _ = v39
        _ = v40
    }
    switch "hello" {
    case "hello":
        if false {
            var v41 string = "\n\t"
            switch "\n\t" {
            case "\n\t":
                v42 := GenS5{F0: v9}
                _ = v42
            case "hello":
                var v43 int8 = -1
                _ = v43
            default:
                if true {
                    v44 := GenS3{F0: (v11 << ((genFn1(128, 2147483648)) & 63)), F1: 0, F2: 7}
                    v2 = v2
                    v45 := GenS7{F0: 0, F1: 255}
                    _ = v44
                    _ = v45
                }
            }
            _ = v41
        }
    case "":
        if false {
            for fi46 := 0; fi46 < 1; fi46++ {
                var v47 bool = true
                _ = fi46
                _ = v47
            }
        }
    default:
        v48 := GenS8{F0: ""}
        _ = v48
    }
    switch ("foo bar" + ("a" + ("\n\t" + "foo bar"))) {
    case "":
        var v49 string = ("foo bar" + (("foo bar" + "") + "\n\t"))
        _ = v49
    default:
        if false {
            if (v2 > v2) {
                switch "hello" {
                case "\n\t":
                    if true {
                        v4 = append(v4, 1)
                        var v50 string = ""
                        _ = v50
                    }
                case "":
                    var v51 uint16 = uint16(v3)
                    _ = v51
                case "a":
                    for fi52 := 0; fi52 < 11; fi52++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi52
                    }
                }
            }
            v2 = v2
            v3 -= (^(v3))
        }
    }
    switch ("hello" + "") {
    case "foo bar":
        var v53 int64 = -877
        _ = v53
    case "a":
        delete(v1, -808)
    }
    v3 -= (v3 << ((0) & 63))
    switch -310 {
    case 42:
        var v54 uint = 2
        _ = v54
    }
    delete(v1, 0)
    v55 := GenS4{F0: 28, F1: 648, F2: (v11 << (((v11 * (v3 + uint(v9)))) & 63))}
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v9
}
