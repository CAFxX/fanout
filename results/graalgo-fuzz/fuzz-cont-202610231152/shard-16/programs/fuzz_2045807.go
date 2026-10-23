package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 string }
type GenS2 struct { F0 int32; F1 string; F2 int }
type GenS3 struct { F0 bool; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 string; F1 uint32; F2 uint32 }

func main() {
    v1 := []int{-9223372036854775808, -9223372036854775808, 290}
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    switch "" {
    case "a":
        var v4 int32 = -2147483648
        _ = v4
    case "\n\t":
        var v5 int = 42
        _ = v5
    }
    switch 212 {
    case 32768:
        if true {
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 uint = 255
                v8 := v2[4294967295]
                _ = fi6
                _ = v7
                _ = v8
            }
            if true {
                var v9 bool = (true && true)
                if (!(((float32(((1e10 * (0.1 + 1.5)) * 0.1)) + 0.1) - -2.25) == 0.0)) {
                    v9 = v9
                } else {
                    v10 := GenS1{F0: -1, F1: -290, F2: "a"}
                    if (0 != 1) {
                        switch (("foo bar" + ("\n\t" + ("" + ("foo bar" + "hello")))) + (("a" + "") + "hello")) {
                        case "":
                            v9 = (-1 != 0)
                        case "foo bar":
                            fmt.Println(v9)
                        case "\n\t":
                            for fi11 := 0; fi11 < 12; fi11++ {
                                if ("a" >= ("foo bar" + "foo bar")) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    for fi12 := 0; fi12 < 11; fi12++ {
                                        switch (("\n\t" + (("a" + ("foo bar" + ("\n\t" + ("" + "")))) + "\n\t")) + "\n\t") {
                                        case "":
                                            v13 := v2[42]
                                            _ = v13
                                        case "a":
                                            v1 = append(v1, fi11)
                                        default:
                                            v9 = (248 < 255)
                                        }
                                        v9 = v9
                                        v9 = (2 >= 77)
                                        _ = fi12
                                    }
                                }
                                var v14 uint = 244
                                _ = fi11
                                _ = v14
                            }
                        default:
                            var v15 uint16 = 80
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 9; fi16++ {
                            v17 := GenS1{F0: 543, F1: 42, F2: "hello"}
                            _ = fi16
                            _ = v17
                        }
                    }
                    var v18 string = ""
                    _ = v10
                    _ = v18
                }
                fmt.Println(v9)
                _ = v9
            }
            switch (("foo bar" + "foo bar") + ((("\n\t" + (("foo bar" + ((("hello" + ((("\n\t" + "a") + "hello") + "")) + "a") + "")) + "")) + (((("hello" + "") + "") + "") + "\n\t")) + "\n\t")) {
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "foo bar":
                v19 := GenS2{F0: 883, F1: ("hello" + "a"), F2: 127}
                _ = v19
            case "a":
                var v20 string = ("a" + "a")
                _ = v20
            default:
                var v21 uint8 = 0
                _ = v21
            }
        } else {
            for fi22 := 0; fi22 < 19; fi22++ {
                if false {
                    var v23 int32 = 0
                    v24 := GenS3{F0: false, F1: 2147483647, F2: (v23 + v23)}
                    _ = v23
                    _ = v24
                }
                for fi25 := 0; fi25 < 6; fi25++ {
                    var v26 uint16 = 154
                    var v27 uint32 = 2
                    var v28 bool = true
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                var v29 int8 = -1
                _ = fi22
                _ = v29
            }
        }
    }
    if (!false) {
        if true {
            var v30 string = "\n\t"
            _ = v30
        }
        var v31 uint16 = 394
        _ = v31
    }
    if true {
        var v32 uint32 = 42
        _ = v32
    } else {
        var v33 int16 = 32767
        v34 := GenS3{F0: false, F1: 272, F2: 2147483647}
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        var v36 uint16 = 255
        delete(v2, 65535)
        _ = fi35
        _ = v36
    }
    v37 := GenS1{F0: 0, F1: -909, F2: "\n\t"}
    var v38 uint32 = 1
    switch "" {
    case "a":
        fmt.Println(v38)
    case "foo bar":
        for fi39 := 0; fi39 < 18; fi39++ {
            var v40 uint64 = 0
            _ = fi39
            _ = v40
        }
    default:
        v38 = 0
    }
    for fi41 := 0; fi41 < 3; fi41++ {
        for fi42 := 0; fi42 < 13; fi42++ {
            if false {
                if false {
                    fmt.Println(v38, fi42)
                    var v43 int = 2147483647
                    _ = len(v3)
                    _ = v43
                } else {
                    v38 += v38
                    v38 += v38
                    v38 = (^(v38))
                }
                v44 := GenS4{F0: 207, F1: -781}
                switch (fi41 % ((-500) | 1)) {
                case -2147483648:
                    fmt.Println(v38, v38)
                case 32768:
                    delete(v3, "\n\t")
                case 751:
                    _ = len(v1)
                }
                _ = v44
            } else {
                v45 := GenS4{F0: 1, F1: -523}
                fmt.Println(fi42, v38, fi42)
                _ = v45
            }
            v38 -= 539
            _ = fi42
        }
        _ = fi41
    }
    fmt.Println(v38)
    var v46 int16 = 32767
    v47 := GenS5{F0: "", F1: 42, F2: (v38 * (v38 ^ 20))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v46
    _ = v47
}
