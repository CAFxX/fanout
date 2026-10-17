package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int16, p1 uint8, p2 uint32) int8 {
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 uint = uint(p1)
        fmt.Println(fi4, fi4, v5)
        switch "a" {
        case "hello":
            switch (fi4 >> ((v5) & 63)) {
            case -560:
                var v6 int16 = p0
                _ = v6
            case 610:
                for fi7 := 0; fi7 < 18; fi7++ {
                    if ((9223372036854775807 >= 42) && (!(2 <= 4294967295))) {
                        var v8 uint = v5
                        _ = v8
                    }
                    p2 = (p2 >> ((2) & 63))
                    switch (("a" + ((("" + (("hello" + "\n\t") + "a")) + "\n\t") + ("\n\t" + ("" + ("hello" + ("a" + "\n\t")))))) + ("hello" + "foo bar")) {
                    case "foo bar":
                        if false {
                            var v9 uint8 = p1
                            var v10 uint = v5
                            p1 += 255
                            _ = v10
                            _ = v9
                        }
                    }
                    _ = fi7
                }
            case 9223372036854775807:
                fmt.Println(p2)
            default:
                v11 := GenS1{F0: -102}
                _ = v11
            }
        }
        _ = fi4
        _ = v5
    }
    v12 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 127
}

func main() {
    v1 := []int{32767, -9223372036854775808}
    v2 := []int{0, -746, -486, -1}
    v3 := make(map[string]uint32)
    var v13 uint = 0
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 int32 = 350
        _ = fi14
        _ = v15
    }
    v13 = (v13 ^ v13)
    v16 := GenS1{F0: 100}
    v13 -= 0
    for fi17 := 0; fi17 < 20; fi17++ {
        switch "hello" {
        case "foo bar":
            var v18 int8 = 127
            _ = v18
        case "hello":
            v13 += 563
        case "\n\t":
            var v19 uint = v13
            _ = v19
        }
        if true {
            if true {
                switch fi17 {
                case 127:
                    for fi20 := 0; fi20 < 5; fi20++ {
                        fmt.Println(fi20, fi17)
                        v13 += v13
                        _ = fi20
                    }
                case 2147483647:
                    for fi21 := 0; fi21 < 3; fi21++ {
                        v13 = 255
                        if (32767 <= 1000) {
                            v13 += (v13 ^ (v13 << (((v13 << ((v13) & 63))) & 63)))
                            delete(v3, "foo bar")
                            v13 = (^(v13))
                        } else {
                            var v22 int16 = 1
                            _ = v22
                        }
                        _ = fi21
                    }
                case 7:
                    if (("hello" + "\n\t") <= ((("foo bar" + "") + "\n\t") + "hello")) {
                        delete(v3, "")
                        var v23 int64 = 1
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 18; fi24++ {
                            if ((!false) || true) {
                                v25 := GenS1{F0: 10}
                                v13 += v13
                                _ = v25
                            }
                            v26 := GenS1{F0: -80}
                            v13 = (v13 >> (((v13 | (^(v13)))) & 63))
                            _ = fi24
                            _ = v26
                        }
                        var v27 string = (("hello" + ("foo bar" + "a")) + "\n\t")
                        _ = v27
                    }
                default:
                    v13 = 525
                }
            }
            var v28 string = ((("" + ("\n\t" + ((((("\n\t" + "") + "hello") + "foo bar") + (("foo bar" + "\n\t") + ("" + "hello"))) + ""))) + "foo bar") + ((("\n\t" + "hello") + ("foo bar" + "hello")) + "hello"))
            _ = v28
        }
        v29 := GenS1{F0: -1}
        _ = fi17
        _ = v29
    }
    var v30 float64 = 0.1
    switch "hello" {
    case "\n\t":
        fmt.Println(v13, v13)
    }
    _ = len(v3)
    if ((!false) && false) {
        fmt.Println(v13, v13)
    }
    v30 = float64(float32(v30))
    v30 = (v30 * v30)
    v13 += (v13 * v13)
    for fi31 := 0; fi31 < 8; fi31++ {
        v13 -= (v13 - (v13 & (^(v13))))
        v32 := GenS1{F0: 100}
        if (!(!true)) {
            var v33 uint32 = 2147483648
            _ = v33
        }
        _ = fi31
        _ = v32
    }
    var v34 int64 = 608
    v13 += 18446744073709551615
    for fi35 := 0; fi35 < 6; fi35++ {
        for fi36 := 0; fi36 < 14; fi36++ {
            switch ("foo bar" + ("hello" + "\n\t")) {
            case "a":
                v30 -= 3.14159
            case "":
                v13 += (v13 / ((2) | 1))
            default:
                fmt.Println(v13, fi35)
            }
            var v37 float64 = (v30 * v30)
            v3[("hello" + ("\n\t" + ("\n\t" + "foo bar")))] = 2
            _ = fi36
            _ = v37
        }
        for fi38 := 0; fi38 < 10; fi38++ {
            v30 = v30
            switch 128 {
            case -1:
                fmt.Println(fi35, v13)
            default:
                fmt.Println(fi35, fi38)
            }
            _ = fi38
        }
        v34 = (v34 | v34)
        _ = fi35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v34
}
