package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint; F2 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string; F1 int; F2 int32 }
type GenS7 struct { F0 uint; F1 uint }
func genFn0(p0 uint64, p1 uint16) int {
    var v6 float32 = 1.5
    var v7 int8 = 42
    if (("a" + ((("hello" + (("\n\t" + "hello") + "")) + "hello") + "foo bar")) > ("a" + "a")) {
        if (p0 < 42) {
            var v8 uint16 = p1
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := []int{(fi9 ^ (fi9 * -2147483648)), fi9, 127}
                v8 -= p1
                _ = fi9
                _ = v10
            }
            var v11 uint32 = 2147483647
            _ = v11
            _ = v8
        } else {
            v12 := GenS1{F0: 2147483647}
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := GenS1{F0: 174}
                fmt.Println(v6)
                v6 -= (0.1 + (v6 - (v6 * (1.5 + float32(-2.25)))))
                _ = fi13
                _ = v14
            }
            _ = v12
        }
    }
    var v15 uint64 = 2
    if (true || true) {
        if true {
            if (v7 < (-(v7))) {
                v16 := GenS1{F0: 893}
                for fi17 := 0; fi17 < 8; fi17++ {
                    switch (fi17 & fi17) {
                    case 595:
                        v6 = (((v6 + v6) * v6) * (v6 + 1.5))
                    }
                    _ = fi17
                }
                v15 += v15
                _ = v16
            } else {
                if (-1 > 1000) {
                    v18 := GenS2{F0: 1, F1: ("foo bar" + "a")}
                    v19 := GenS3{F0: true, F1: p1}
                    _ = v18
                    _ = v19
                } else {
                    v20 := []int{127, 9223372036854775807, 567}
                    v15 = (p0 - (v15 - 18446744073709551615))
                    _ = v20
                }
                for fi21 := 0; fi21 < 8; fi21++ {
                    if false {
                        v22 := GenS4{F0: 696, F1: uint(p1), F2: 230}
                        var v23 int = (^(fi21))
                        fmt.Println(p0, p0)
                        _ = v22
                        _ = v23
                    }
                    p1 = p1
                    _ = fi21
                }
            }
            if false {
                v6 = float32(-2.25)
                fmt.Println(p0, p0)
                v24 := []int{1}
                _ = v24
            }
        } else {
            for fi25 := 0; fi25 < 13; fi25++ {
                v26 := []int{fi25, (fi25 & fi25), fi25}
                _ = fi25
                _ = v26
            }
            if true {
                switch "\n\t" {
                case "a":
                    for fi27 := 0; fi27 < 14; fi27++ {
                        v7 -= 6
                        p1 += p1
                        _ = fi27
                    }
                case "\n\t":
                    v15 = (p0 >> ((18446744073709551615) & 63))
                default:
                    v7 = 0
                }
                fmt.Println(p0, p0, v15)
            } else {
                v7 = -119
                if false {
                    switch "hello" {
                    case "":
                        var v28 float32 = v6
                        _ = v28
                    case "\n\t":
                        for fi29 := 0; fi29 < 10; fi29++ {
                            var v30 float64 = 0.0
                            fmt.Println(fi29, v15, p0)
                            _ = fi29
                            _ = v30
                        }
                    case "foo bar":
                        var v31 uint8 = 220
                        _ = v31
                    default:
                        var v32 float32 = ((((1e10 + (v6 * v6)) / (3.14159 / v6)) + (float32((((1e10 + float64(float32(1.5))) + 1e10) + (3.14159 * (float64(((v6 * (v6 + 1.5)) - v6)) - -2.25)))) * -2.25)) / (v6 / ((v6 / float32(-2.25)) * v6)))
                        _ = v32
                    }
                }
                fmt.Println(v15)
            }
            v15 += p0
        }
        v33 := []int{32768}
        _ = v33
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int, p1 uint32, p2 string) uint64 {
    fmt.Println(p0, p2, p0)
    switch (p0 - p0) {
    case -9223372036854775808:
        fmt.Println(p0)
    default:
        p2 = (("hello" + ("\n\t" + p2)) + p2)
    }
    v34 := []int{}
    switch "foo bar" {
    case "a":
        for fi35 := 0; fi35 < 20; fi35++ {
            var v36 int = genFn0(2, 0)
            v36 = (p0 / (((p0 << ((65535) & 63))) | 1))
            p1 = p1
            _ = fi35
            _ = v36
        }
    default:
        _ = len(v34)
    }
    if true {
        var v37 uint32 = 1
        var v38 string = p2
        if (!(v37 >= v37)) {
            switch p0 {
            case -9223372036854775808:
                p0 -= genFn0(42, 65535)
            default:
                v37 = (v37 >> ((65535) & 63))
            }
            p2 = "a"
        } else {
            var v39 uint32 = p1
            _ = v39
        }
        _ = v37
        _ = v38
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    return 18446744073709551615
}

func main() {
    var v1 int16 = -574
    v2 := make(map[uint32]int)
    var v3 bool = (!(!(!true)))
    v4 := []int{9223372036854775807, 32767}
    v5 := []uint32{2147483647}
    _ = len(v5)
    for fi40 := 0; fi40 < 16; fi40++ {
        if len(v5) > 0 {
            v5[0] = 65535
        }
        v3 = v3
        if true {
            for fi41 := 0; fi41 < 1; fi41++ {
                var v42 uint = 626
                v4 = append(v4, int(v1))
                _ = fi41
                _ = v42
            }
            for fi43 := 0; fi43 < 14; fi43++ {
                v44 := GenS5{F0: 2147483647}
                switch (fi43 / ((fi43) | 1)) {
                case 0:
                    switch genFn0(794, 42) {
                    case -2147483648:
                        var v45 bool = (("hello" + "a") >= ((("foo bar" + ("a" + ("foo bar" + "\n\t"))) + ("hello" + ("a" + ("" + "a")))) + ((("foo bar" + "foo bar") + "foo bar") + (("" + (("foo bar" + (("" + "foo bar") + ("foo bar" + "\n\t"))) + ((("a" + ("hello" + "\n\t")) + "foo bar") + "hello"))) + "a"))))
                        _ = v45
                    case -9223372036854775808:
                        var v46 uint = 18446744073709551615
                        v47 := GenS6{F0: "hello", F1: (fi43 & (fi40 >> ((547) & 63))), F2: -358}
                        _ = v46
                        _ = v47
                    case 128:
                        switch ("foo bar" + "a") {
                        case "a":
                            v5 = append(v5, 421)
                        case "\n\t":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        if ((v3 || (-73 < 1)) || ((v3 && (1 != 4294967295)) || (v1 >= (v1 + v1)))) {
                            if (!v3) {
                                v4 = append(v4, fi43)
                                v3 = ("" != ((("foo bar" + ("" + "")) + ((("hello" + "") + ((("foo bar" + "a") + ("foo bar" + "a")) + "foo bar")) + ("hello" + ("a" + "\n\t")))) + ""))
                            }
                        } else {
                            var v48 int32 = 424
                            _ = v48
                        }
                    }
                case -9223372036854775808:
                    v49 := GenS5{F0: 783}
                    _ = v49
                case 42:
                    var v50 bool = v3
                    _ = v50
                }
                _ = fi43
                _ = v44
            }
            var v51 uint16 = 65535
            _ = v51
        } else {
            for fi52 := 0; fi52 < 5; fi52++ {
                fmt.Println(fi40, fi52)
                v3 = v3
                _ = fi52
            }
            v1 = v1
            v53 := GenS1{F0: 2}
            _ = v53
        }
        _ = fi40
    }
    v1 = (v1 / ((v1) | 1))
    v54 := GenS3{F0: v3, F1: 32767}
    fmt.Println(v3, v1)
    for fi55 := 0; fi55 < 2; fi55++ {
        var v56 uint32 = 2147483647
        var v57 float32 = 0.1
        _ = fi55
        _ = v56
        _ = v57
    }
    v58 := GenS3{F0: (!(-2.25 >= 1e10)), F1: 255}
    v59 := GenS7{F0: 1, F1: 0}
    var v60 float32 = (0.1 * (-2.25 / 1.5))
    v61 := GenS4{F0: 0, F1: 65535, F2: 2}
    var v62 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v60:", v60)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v58
    _ = v59
    _ = v60
    _ = v61
    _ = v62
}
