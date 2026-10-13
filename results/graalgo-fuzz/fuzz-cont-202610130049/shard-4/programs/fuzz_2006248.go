package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{0, 42, 2, 111}
    var v3 uint8 = 38
    var v4 int64 = -173
    for fi5 := 0; fi5 < 20; fi5++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi5
    }
    if (!((871 > 0) && true)) {
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 float32 = 1.5
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 bool = (v3 >= 255)
                if v9 {
                    var v10 int8 = -128
                    var v11 string = "foo bar"
                    v12 := GenS1{F0: 4294967295, F1: fi8}
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    v4 = v4
                }
                _ = fi8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        v1[""] = 32767
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v13 := GenS2{F0: 214}
        _ = v13
    }
    v3 = (v3 ^ 1)
    var v14 int32 = -751
    switch "hello" {
    case "":
        var v15 int16 = -1
        _ = v15
    case "hello":
        if true {
            v3 = v3
        }
    case "a":
        if (!((v14 <= (-(v14))) || true)) {
            v16 := GenS1{F0: 4294967295, F1: 127}
            var v17 int64 = 0
            _ = v16
            _ = v17
        } else {
            var v18 bool = (127 == -128)
            fmt.Println(v18)
            _ = v18
        }
    }
    v4 += (v4 % (((v4 & (v4 ^ (v4 / (((v4 << ((18446744073709551615) & 63))) | 1))))) | 1))
    v4 -= (v4 >> ((743) & 63))
    for fi19 := 0; fi19 < 19; fi19++ {
        if len(v2) > 0 {
            v2[0] = 702
        }
        _ = len(v1)
        var v20 int8 = 1
        _ = fi19
        _ = v20
    }
    switch "" {
    case "foo bar":
        for fi21 := 0; fi21 < 5; fi21++ {
            v22 := GenS3{F0: 2147483647, F1: 42}
            fmt.Println(fi21, v4)
            _ = fi21
            _ = v22
        }
    case "hello":
        if ((v14 / (((-(v14))) | 1)) < v14) {
            v23 := GenS4{F0: 32768, F1: 106}
            _ = v23
        }
    default:
        if false {
            switch 32767 {
            case 9223372036854775807:
                if false {
                    v3 = v3
                }
            case -2147483648:
                if false {
                    for fi24 := 0; fi24 < 3; fi24++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi24
                    }
                    v1[("a" + "a")] = 2147483647
                }
            case 4294967295:
                for fi25 := 0; fi25 < 13; fi25++ {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        if true {
                            var v27 int32 = v14
                            if (!(false && false)) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v28 int64 = (v4 & 0)
                                if ("" >= "") {
                                    fmt.Println(fi26, v14, v14)
                                }
                                v28 = v4
                                var v29 uint16 = 0
                                var v30 float64 = -2.25
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            _ = v27
                        } else {
                            switch (fi26 ^ 32767) {
                            case 1:
                                if len(v2) > 0 {
                                    v2[0] = 2147483647
                                }
                            default:
                                _ = len(v2)
                            }
                        }
                        if (!(true || false)) {
                            if (true || (false && true)) {
                                v4 -= (v4 % (((-(v4))) | 1))
                                v31 := GenS3{F0: int(v14), F1: 128}
                                var v32 uint16 = 255
                                _ = v31
                                _ = v32
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = fi26
                    }
                    for fi33 := 0; fi33 < 2; fi33++ {
                        if ((!true) && false) {
                            v34 := GenS1{F0: 4294967295, F1: fi33}
                            fmt.Println(fi25, fi33, v14)
                            v3 = v3
                            _ = v34
                        }
                        v3 -= v3
                        switch "" {
                        case "hello":
                            if ((!(1 >= v3)) || (0.0 > 0.1)) {
                                v14 = -879
                            } else {
                                v35 := GenS5{F0: "hello"}
                                var v36 int64 = v4
                                var v37 int16 = int16(v14)
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                        case "":
                            v38 := GenS5{F0: (("a" + "") + ("a" + "\n\t"))}
                            _ = v38
                        case "a":
                            v39 := v1[""]
                            switch v39 {
                            case 9223372036854775807:
                                _ = len(v2)
                            case -668:
                                v3 = v3
                            case -704:
                                if (32767 < 189) {
                                    var v40 int8 = -128
                                    v41 := v1[""]
                                    v42 := GenS3{F0: (fi25 | (fi25 * fi25)), F1: v3}
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                }
                            }
                            _ = v39
                        }
                        fmt.Println(fi33, v4)
                        v43 := GenS6{F0: 1, F1: 65535}
                        _ = fi33
                        _ = v43
                    }
                    _ = fi25
                }
            }
            _ = len(v2)
        } else {
            v4 = v4
            v4 = v4
        }
    }
    switch "" {
    case "a":
        switch "foo bar" {
        case "":
            v44 := GenS2{F0: 42}
            _ = v44
        case "a":
            for fi45 := 0; fi45 < 11; fi45++ {
                var v46 int32 = (v14 & 1)
                v46 = v14
                _ = fi45
                _ = v46
            }
        case "\n\t":
            for fi47 := 0; fi47 < 14; fi47++ {
                fmt.Println(v14, v3)
                if (116 > 65535) {
                    v4 += (v4 << ((617) & 63))
                } else {
                    switch (("\n\t" + ("\n\t" + "a")) + "") {
                    case "a":
                        delete(v1, ("a" + ("" + ("a" + ("a" + "foo bar")))))
                    }
                }
                v14 = v14
                _ = fi47
            }
        default:
            v14 = (v14 | 42)
        }
    }
    fmt.Println(v4)
    v3 = 165
    switch ("foo bar" + ("hello" + "\n\t")) {
    case "a":
        v3 -= v3
    default:
        _ = len(v1)
    }
    fmt.Println(v4, v3, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
