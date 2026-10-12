package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int32; F1 int64; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 42
    v3 := []uint32{42, 104, 2147483648, 2147483647}
    v4 := GenS1{F0: "\n\t", F1: (v2 << ((4294967295) & 63))}
    switch -525 {
    case 4294967295:
        v2 += v2
    case 32767:
        v2 += 0
    case -152:
        var v5 float32 = 0.1
        _ = v5
    default:
        var v6 uint = 42
        _ = v6
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        delete(v1, ((("" + ("foo bar" + "foo bar")) + ("foo bar" + "foo bar")) + "foo bar"))
        if (!true) {
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 int8 = 62
                switch fi7 {
                case 4294967295:
                    switch ((("\n\t" + "hello") + "") + ("a" + "hello")) {
                    case "hello":
                        v8 += (v8 | (^(v8)))
                    case "\n\t":
                        if false {
                            var v9 bool = (false && true)
                            for fi10 := 0; fi10 < 20; fi10++ {
                                switch -367 {
                                case 4294967295:
                                    v11 := GenS1{F0: "a", F1: 0}
                                    _ = v11
                                case -126:
                                    v2 -= -1
                                }
                                var v12 bool = ("a" == "a")
                                _ = fi10
                                _ = v12
                            }
                            v8 = -73
                            _ = v9
                        }
                    default:
                        v13 := GenS2{F0: 157, F1: false, F2: 42}
                        _ = v13
                    }
                }
                for fi14 := 0; fi14 < 2; fi14++ {
                    var v15 uint64 = 4294967295
                    switch fi7 {
                    case 32767:
                        var v16 uint32 = 1
                        _ = v16
                    case 0:
                        switch "foo bar" {
                        case "\n\t":
                            v17 := GenS3{F0: (713 <= (^(v2)))}
                            _ = v17
                        default:
                            var v18 int16 = 1
                            v8 -= (v8 + -97)
                            _ = v18
                        }
                    case -41:
                        v2 = (v2 ^ -1)
                    default:
                        var v19 uint64 = v15
                        _ = v19
                    }
                    _ = fi14
                    _ = v15
                }
                _ = fi7
                _ = v8
            }
        } else {
            var v20 bool = true
            v21 := GenS4{F0: 188, F1: 0}
            _ = v20
            _ = v21
        }
        if (2 >= 0) {
            v2 += (v2 & v2)
            switch "a" {
            case "a":
                v2 -= v2
            }
            var v22 int64 = v2
            _ = v22
        }
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v23 uint8 = 127
    if (false || true) {
        if (!(("foo bar" + "hello") < (("foo bar" + ("foo bar" + ("" + "\n\t"))) + ""))) {
            v23 += v23
            for fi24 := 0; fi24 < 14; fi24++ {
                if true {
                    var v25 uint16 = 1
                    for fi26 := 0; fi26 < 14; fi26++ {
                        var v27 float64 = 3.14159
                        _ = fi26
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 15; fi28++ {
                        var v29 float64 = 0.0
                        _ = fi28
                        _ = v29
                    }
                    _ = v25
                }
                _ = fi24
            }
        } else {
            v30 := GenS3{F0: (128 >= (^(v23)))}
            var v31 int64 = 123
            _ = v30
            _ = v31
        }
        var v32 bool = (255 != 440)
        v2 = v2
        _ = v32
    } else {
        for fi33 := 0; fi33 < 10; fi33++ {
            v2 = v2
            fmt.Println(v2, v23, v2)
            _ = fi33
        }
        var v34 int64 = (v2 | (v2 ^ v2))
        _ = v34
    }
    if true {
        switch ("a" + "a") {
        case "hello":
            v23 -= (v23 >> ((255) & 63))
        case "foo bar":
            if (!false) {
                var v35 uint8 = v23
                _ = v35
            } else {
                v36 := GenS4{F0: v23, F1: 4294967295}
                _ = v36
            }
        case "\n\t":
            var v37 uint64 = 18446744073709551615
            _ = v37
        }
        if (!((v23 & v23) == 255)) {
            v38 := GenS5{F0: -69}
            v39 := GenS4{F0: (v23 | v23), F1: 0}
            switch ("" + "") {
            case "":
                if (!false) {
                    switch "hello" {
                    case "":
                        v1["a"] = 42
                    default:
                        if true {
                            for fi40 := 0; fi40 < 14; fi40++ {
                                for fi41 := 0; fi41 < 7; fi41++ {
                                    v42 := GenS6{F0: 255}
                                    _ = fi41
                                    _ = v42
                                }
                                v43 := GenS7{F0: 32767, F1: ("hello" + ("a" + "\n\t"))}
                                _ = fi40
                                _ = v43
                            }
                        } else {
                            if (true || true) {
                                v2 = v2
                                switch "hello" {
                                case "hello":
                                    v23 = (v23 % (((v23 + v23)) | 1))
                                case "":
                                    _ = len(v1)
                                }
                            }
                            if false {
                                for fi44 := 0; fi44 < 8; fi44++ {
                                    var v45 int16 = 1000
                                    v2 = (v2 ^ v2)
                                    _ = fi44
                                    _ = v45
                                }
                                v46 := GenS8{F0: 607}
                                _ = v46
                            } else {
                                v23 -= v23
                                var v47 uint64 = 5
                                v48 := GenS1{F0: ("a" + "foo bar"), F1: -1}
                                _ = v47
                                _ = v48
                            }
                            v23 -= (^(v23))
                        }
                    }
                    var v49 uint16 = 1
                    _ = v49
                }
            case "foo bar":
                v50 := GenS9{F0: -2147483648, F1: (-(v2)), F2: 280}
                _ = v50
            }
            _ = v38
            _ = v39
        }
        switch (("foo bar" + "hello") + (("" + ("hello" + "a")) + (("" + "hello") + "\n\t"))) {
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = 42
            }
        default:
            var v51 uint8 = 110
            _ = v51
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
