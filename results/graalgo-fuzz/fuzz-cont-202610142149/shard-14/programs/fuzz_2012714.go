package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint }
type GenS3 struct { F0 string; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 int32; F1 int8; F2 uint16 }
type GenS7 struct { F0 int16; F1 int; F2 int16 }
type GenS8 struct { F0 uint8; F1 int16; F2 int }
type GenS9 struct { F0 uint8; F1 uint; F2 uint16 }
func genFn0() string {
    var v7 int8 = -128
    var v8 int64 = 0
    _ = v7
    _ = v8
    return ""
}
func genFn1() bool {
    var v9 string = ("foo bar" + "\n\t")
    if (true && (v9 != ((v9 + v9) + v9))) {
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 14; fi11++ {
                fmt.Println(fi11)
                v12 := GenS1{F0: 127}
                fmt.Println(v9, fi10, fi11)
                _ = fi11
                _ = v12
            }
            var v13 int16 = 616
            if (fi10 == (fi10 | fi10)) {
                v14 := GenS2{F0: (v13 | v13), F1: 121, F2: 65535}
                var v15 int8 = 11
                _ = v14
                _ = v15
            } else {
                v9 = (v9 + (("foo bar" + v9) + v9))
            }
            _ = fi10
            _ = v13
        }
        v9 = "\n\t"
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := GenS3{F0: (v9 + (v9 + ((v9 + ("\n\t" + (("foo bar" + v9) + "hello"))) + v9))), F1: 1, F2: 127}
        _ = fi16
        _ = v17
    }
    var v18 int32 = 898
    v9 = ("\n\t" + v9)
    _ = v18
    _ = v9
    return (974 >= 2)
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 255
    v3 := []string{"a", "\n\t"}
    v4 := make(map[uint32]uint32)
    v5 := []string{"hello", "a", "\n\t"}
    v6 := []int{128}
    v19 := GenS4{F0: (v2 % (((^(v2))) | 1)), F1: 897}
    var v20 bool = (!false)
    v21 := v4[206]
    fmt.Println(v2, v2, v2)
    v22 := GenS5{F0: 100, F1: (v2 >> ((286) & 63))}
    v20 = (!(3.14159 <= float32(0.1)))
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    var v23 int64 = 857
    var v24 int64 = v23
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 20; fi26++ {
            v2 += v2
            if false {
                var v27 float32 = (-2.25 / (-2.25 * (float32(0.1) - 3.14159)))
                v3 = append(v3, "")
                _ = v27
            } else {
                if (v2 != v2) {
                    delete(v1, (("hello" + "\n\t") + ""))
                    if (((1e10 == 3.14159) && v20) && v20) {
                        v28 := GenS6{F0: -111, F1: 97, F2: 1}
                        v29 := GenS5{F0: -1, F1: v2}
                        if (!(fi26 < (fi25 >> ((uint(v21)) & 63)))) {
                            v30 := GenS7{F0: 304, F1: fi25, F2: 881}
                            _ = v30
                        }
                        _ = v28
                        _ = v29
                    }
                } else {
                    var v31 float64 = (float64(((0.1 * 1e10) / 3.14159)) * (-2.25 * (-2.25 - float64(1e10))))
                    switch (-(fi26)) {
                    case 9223372036854775807:
                        switch "" {
                        case "":
                            switch (fi25 / ((fi26) | 1)) {
                            case 32768:
                                var v32 string = (("hello" + ((("hello" + "hello") + ("hello" + (("hello" + "") + "a"))) + ("foo bar" + ("a" + "a")))) + "a")
                                _ = v32
                            case 0:
                                v33 := GenS8{F0: (v2 >> ((0) & 63)), F1: 1, F2: fi25}
                                _ = v33
                            }
                        case "foo bar":
                            _ = len(v1)
                        default:
                            v31 -= ((3.14159 / v31) * (float64(3.14159) + (v31 * v31)))
                        }
                    case -9223372036854775808:
                        if v20 {
                            v21 -= 4294967295
                        }
                        fmt.Println(v23, fi26, v20)
                    }
                    fmt.Println(fi25, v24, v2)
                    _ = v31
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = fi26
        }
        _ = fi25
    }
    fmt.Println(v2, v21)
    v24 = v24
    v23 -= v23
    var v34 int16 = 1000
    if (!false) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi35 := 0; fi35 < 14; fi35++ {
            if (!true) {
                for fi36 := 0; fi36 < 8; fi36++ {
                    v37 := GenS2{F0: (v34 + v34), F1: (^(v2)), F2: 18446744073709551615}
                    _ = fi36
                    _ = v37
                }
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v24 = -9223372036854775808
            _ = fi35
        }
        fmt.Println(v34)
    }
    v38 := v1[""]
    if v20 {
        for fi39 := 0; fi39 < 20; fi39++ {
            v34 = (v34 * v34)
            v2 = v2
            _ = fi39
        }
        for fi40 := 0; fi40 < 20; fi40++ {
            v41 := v4[1]
            v23 = (v24 >> ((510) & 63))
            _ = fi40
            _ = v41
        }
        fmt.Println(v21)
    } else {
        if (-36 != 0) {
            v42 := GenS9{F0: v2, F1: 0, F2: 65535}
            switch -44 {
            case 141:
                if v20 {
                    v23 = (v24 + v23)
                    v23 += v24
                    var v43 int64 = (v24 >> ((48) & 63))
                    _ = v43
                }
            }
            _ = v42
        } else {
            v44 := v4[(^(v21))]
            var v45 int32 = 42
            v23 = 0
            _ = v44
            _ = v45
        }
        for fi46 := 0; fi46 < 14; fi46++ {
            v1[v38] = v38
            v38 = (("a" + (v38 + ("hello" + (("a" + v38) + v38)))) + (v38 + "a"))
            v47 := v1[(v38 + v38)]
            _ = fi46
            _ = v47
        }
        if len(v3) > 0 {
            v3[0] = ("hello" + "a")
        }
    }
    var v48 uint32 = (v21 >> ((255) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
