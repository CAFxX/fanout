package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 bool; F1 int16; F2 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 uint32 }

func main() {
    v1 := []int{-401, 234, 4294967295, 1}
    var v2 uint16 = 0
    v3 := []string{"", ((("foo bar" + "foo bar") + ((("foo bar" + "\n\t") + "foo bar") + "foo bar")) + "hello"), "foo bar", "hello"}
    v2 = v2
    switch -2147483648 {
    case -843:
        switch "a" {
        case "":
            v4 := GenS1{F0: 159, F1: 105}
            _ = v4
        case "\n\t":
            switch ("\n\t" + ((("hello" + (("a" + (("hello" + "\n\t") + "foo bar")) + (("" + ("" + "hello")) + "\n\t"))) + "") + "")) {
            case "\n\t":
                switch "a" {
                case "a":
                    if (42 >= -9223372036854775808) {
                        switch ("a" + "a") {
                        case "":
                            fmt.Println(v2)
                        default:
                            if true {
                                var v5 uint = 842
                                _ = v5
                            } else {
                                v6 := GenS1{F0: 270, F1: -95}
                                _ = v6
                            }
                        }
                    }
                default:
                    for fi7 := 0; fi7 < 2; fi7++ {
                        for fi8 := 0; fi8 < 7; fi8++ {
                            v2 += 32767
                            _ = fi8
                        }
                        v2 += 65535
                        var v9 int32 = 0
                        _ = fi7
                        _ = v9
                    }
                }
            case "a":
                v10 := GenS1{F0: 593, F1: -1}
                _ = v10
            default:
                if ((-1 >= -622) && true) {
                    for fi11 := 0; fi11 < 13; fi11++ {
                        if len(v1) > 0 {
                            v1[0] = fi11
                        }
                        var v12 uint = 65535
                        switch (^(fi11)) {
                        case 67:
                            _ = len(v1)
                        case 127:
                            switch 37 {
                            case 1:
                                fmt.Println(v12)
                            }
                        }
                        _ = fi11
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 16; fi13++ {
                        for fi14 := 0; fi14 < 19; fi14++ {
                            var v15 bool = (1 != 2)
                            var v16 uint16 = v2
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        _ = fi13
                    }
                } else {
                    if (1 >= 0) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if false {
                            var v17 int = 42
                            if false {
                                switch -2147483648 {
                                case 1:
                                    var v18 string = "a"
                                    _ = v18
                                }
                                var v19 uint8 = 22
                                if (875 > 1) {
                                    var v20 float32 = 0.0
                                    _ = len(v3)
                                    switch v17 {
                                    case -9223372036854775808:
                                        var v21 int16 = -1
                                        _ = v21
                                    }
                                    _ = v20
                                }
                                _ = v19
                            } else {
                                switch 32768 {
                                case 2147483647:
                                    var v22 int16 = 0
                                    _ = v22
                                }
                                fmt.Println(v17)
                                for fi23 := 0; fi23 < 11; fi23++ {
                                    v17 += v17
                                    v17 += (fi23 * fi23)
                                    _ = fi23
                                }
                            }
                            _ = v17
                        }
                        v24 := GenS1{F0: 2, F1: 100}
                        _ = v24
                    } else {
                        switch 128 {
                        case 128:
                            if true {
                                v2 = 542
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        case 32767:
                            v25 := GenS1{F0: 0, F1: -128}
                            _ = v25
                        default:
                            for fi26 := 0; fi26 < 18; fi26++ {
                                v1 = append(v1, (fi26 - (fi26 << ((556) & 63))))
                                if (("foo bar" + ("hello" + "")) <= "a") {
                                    var v27 uint16 = v2
                                    if false {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v27 += (v2 >> ((702) & 63))
                                        if true {
                                            fmt.Println(v27)
                                            v2 = 886
                                            for fi28 := 0; fi28 < 5; fi28++ {
                                                v29 := GenS2{F0: 317, F1: v27}
                                                fmt.Println(v2, v2)
                                                switch "a" {
                                                case "":
                                                    var v30 int64 = 42
                                                    _ = v30
                                                default:
                                                    v27 += (v2 * 0)
                                                }
                                                _ = fi28
                                                _ = v29
                                            }
                                        }
                                    }
                                    v31 := GenS1{F0: 2, F1: 0}
                                    _ = v27
                                    _ = v31
                                }
                                _ = fi26
                            }
                        }
                    }
                    if (-835 >= 45) {
                        switch ("hello" + "hello") {
                        case "\n\t":
                            _ = len(v3)
                        case "":
                            v2 = 65535
                        default:
                            v1 = append(v1, -1)
                        }
                        var v32 uint16 = (v2 % ((v2) | 1))
                        if (!false) {
                            for fi33 := 0; fi33 < 10; fi33++ {
                                switch "a" {
                                case "\n\t":
                                    v3 = append(v3, "a")
                                }
                                if len(v1) > 0 {
                                    v1[0] = fi33
                                }
                                _ = fi33
                            }
                            v34 := GenS1{F0: 875, F1: 100}
                            _ = v34
                        } else {
                            v2 = (v2 << ((24) & 63))
                            v35 := GenS3{F0: (127 != -1), F1: -576, F2: 42}
                            v2 = (v2 ^ (v2 << ((18446744073709551615) & 63)))
                            _ = v35
                        }
                        _ = v32
                    } else {
                        if (39 >= 65535) {
                            var v36 int8 = 0
                            _ = v36
                        }
                        v37 := GenS1{F0: 42, F1: 93}
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                        _ = v37
                    }
                }
            }
        case "hello":
            _ = len(v3)
        }
    case -595:
        switch ("" + "hello") {
        case "foo bar":
            var v38 int8 = -128
            _ = v38
        default:
            var v39 int16 = -113
            _ = v39
        }
    case 42:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v40 := GenS1{F0: 214, F1: 9}
        _ = v40
    }
    v2 += 21
    var v41 int16 = 1000
    v42 := GenS4{F0: "\n\t"}
    if (127 < 64) {
        v3 = append(v3, "foo bar")
        switch "\n\t" {
        case "\n\t":
            var v43 int = -542
            _ = v43
        default:
            fmt.Println(v2)
        }
    }
    var v44 float32 = 0.0
    var v45 int64 = 42
    v46 := GenS5{F0: (v45 & int64(v2))}
    var v47 int32 = -205
    v48 := GenS6{F0: 255, F1: 65535}
    v2 = (^(v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
}
