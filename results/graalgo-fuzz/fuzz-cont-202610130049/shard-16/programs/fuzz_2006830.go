package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int; F2 string }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 int32; F2 uint8 }
func genFn0(p0 bool, p1 int32) int16 {
    switch "a" {
    case "a":
        for fi5 := 0; fi5 < 2; fi5++ {
            p0 = (9223372036854775807 < 9223372036854775807)
            if (!false) {
                if (!(!(!p0))) {
                    var v6 int = 32768
                    _ = v6
                } else {
                    v7 := []int{}
                    switch 0 {
                    case -103:
                        if ((^(p1)) == (p1 - (p1 >> ((255) & 63)))) {
                            switch (fi5 % ((4294967295) | 1)) {
                            case 9223372036854775807:
                                fmt.Println(p0, fi5, p1)
                            }
                            fmt.Println(p0, fi5)
                        }
                    default:
                        p1 -= p1
                    }
                    _ = v7
                }
                v8 := []int{fi5, fi5, fi5}
                _ = v8
            }
            var v9 int32 = (p1 >> ((1) & 63))
            _ = fi5
            _ = v9
        }
    }
    fmt.Println(p1, p0)
    if (!(807 != 401)) {
        switch 927 {
        case -208:
            fmt.Println(p1, p0)
        case 2147483647:
            v10 := GenS1{F0: 2147483647}
            _ = v10
        }
        p0 = true
    }
    var v11 bool = (!p0)
    var v12 int16 = 1
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 191
}
func genFn1(p0 uint16, p1 bool) int8 {
    p0 = 65535
    var v13 int32 = 1
    _ = p0
    _ = p1
    _ = v13
    return 127
}

func main() {
    v1 := []uint32{0}
    v2 := []uint32{793, 2147483648}
    var v3 int64 = -296
    v4 := make(map[int]string)
    switch -1 {
    case 2147483647:
        for fi14 := 0; fi14 < 20; fi14++ {
            switch "a" {
            case "foo bar":
                var v15 uint8 = 128
                _ = v15
            case "":
                v16 := GenS2{F0: 416, F1: fi14, F2: "a"}
                _ = v16
            default:
                for fi17 := 0; fi17 < 7; fi17++ {
                    switch "\n\t" {
                    case "hello":
                        v3 = v3
                    default:
                        var v18 int8 = genFn1(255, true)
                        _ = v18
                    }
                    v19 := GenS2{F0: 65535, F1: (fi14 + fi17), F2: "\n\t"}
                    fmt.Println(fi14, fi17, fi14)
                    _ = fi17
                    _ = v19
                }
            }
            _ = fi14
        }
    default:
        var v20 string = "hello"
        _ = v20
    }
    fmt.Println(v3)
    v3 += (^(v3))
    v21 := GenS3{F0: -17, F1: 100}
    switch "" {
    case "":
        for fi22 := 0; fi22 < 19; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                if true {
                    v3 += 9223372036854775807
                }
                _ = fi23
            }
            for fi24 := 0; fi24 < 13; fi24++ {
                v3 += (^(v3))
                _ = fi24
            }
            _ = fi22
        }
    default:
        v4[32767] = ("hello" + "a")
    }
    fmt.Println(v3)
    v3 = (v3 * (v3 << ((991) & 63)))
    if ("" < ("\n\t" + "foo bar")) {
        v3 += (v3 % (((v3 & v3)) | 1))
        var v25 float32 = 3.14159
        var v26 uint = 0
        _ = v25
        _ = v26
    }
    v27 := GenS4{F0: -128}
    var v28 float64 = 0.0
    v4[9223372036854775807] = ("\n\t" + ("\n\t" + ("hello" + "\n\t")))
    if (genFn0((42 == (^(v3))), -2147483648) <= 42) {
        switch ((("hello" + ("\n\t" + "foo bar")) + "hello") + "a") {
        case "foo bar":
            var v29 int32 = -364
            _ = v29
        case "a":
            v3 = v3
        default:
            var v30 uint8 = 127
            _ = v30
        }
    }
    v1 = append(v1, 2147483648)
    v4[-2147483648] = ("hello" + ("a" + "hello"))
    if ("\n\t" == "") {
        var v31 int16 = -722
        var v32 uint16 = 32767
        if (v28 == (((0.0 * (0.1 + v28)) + (0.0 + 0.0)) / 1.0)) {
            for fi33 := 0; fi33 < 16; fi33++ {
                v34 := v4[fi33]
                _ = fi33
                _ = v34
            }
            v3 += v3
        }
        _ = v31
        _ = v32
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        for fi36 := 0; fi36 < 10; fi36++ {
            fmt.Println(v28, v28, fi36)
            var v37 uint = 1
            if (v37 <= v37) {
                v38 := GenS5{F0: 182, F1: -2147483648, F2: 0}
                v39 := v4[2147483647]
                v40 := v4[329]
                _ = v38
                _ = v39
                _ = v40
            }
            _ = fi36
            _ = v37
        }
        var v41 uint64 = 4294967295
        var v42 uint8 = 2
        _ = fi35
        _ = v41
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
