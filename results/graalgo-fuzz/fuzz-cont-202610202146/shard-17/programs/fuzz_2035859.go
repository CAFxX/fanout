package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 int; F2 uint64 }
func genFn0(p0 uint8, p1 bool) bool {
    p0 -= p0
    v6 := GenS1{F0: -1, F1: 318}
    if false {
        v7 := GenS2{F0: 821, F1: 65535}
        p1 = p1
        var v8 int8 = 42
        _ = v7
        _ = v8
    } else {
        v9 := GenS3{F0: 42, F1: false}
        p1 = (0 > 885)
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(fi10)
        var v11 string = ("a" + ("a" + ("foo bar" + "\n\t")))
        _ = fi10
        _ = v11
    }
    var v12 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    return true
}

func main() {
    v1 := []string{("foo bar" + ("\n\t" + (("a" + "hello") + "foo bar")))}
    v2 := make(map[string]int)
    var v3 int = 42
    v4 := make(map[uint32]string)
    v5 := []string{("a" + ""), "foo bar"}
    v13 := v2["hello"]
    v4[2] = ("hello" + "a")
    var v14 string = ""
    var v15 uint32 = 2147483647
    v13 -= 2147483647
    for fi16 := 0; fi16 < 14; fi16++ {
        fmt.Println(v14)
        v13 -= (v13 & (v13 >> ((4294967295) & 63)))
        _ = fi16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v3 = v3
        if (true || (!((1.5 / (1e10 / 1.5)) == -2.25))) {
            v18 := GenS3{F0: -1, F1: true}
            if (v3 != (fi17 >> ((uint(v15)) & 63))) {
                switch -619 {
                case -1:
                    v13 = (fi17 >> ((255) & 63))
                default:
                    if false {
                        for fi19 := 0; fi19 < 18; fi19++ {
                            var v20 string = v14
                            v21 := GenS2{F0: -2147483648, F1: uint16(v15)}
                            _ = len(v5)
                            switch (fi19 >> ((734) & 63)) {
                            case -646:
                                var v22 uint8 = 128
                                _ = v22
                            case 308:
                                fmt.Println(fi17)
                            }
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        v15 += (v15 >> ((2) & 63))
                    }
                }
            }
            _ = v18
        }
        _ = fi17
    }
    if false {
        var v23 uint32 = v15
        v24 := GenS2{F0: (v13 / ((v3) | 1)), F1: 171}
        _ = v23
        _ = v24
    } else {
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v25 uint16 = 65535
            _ = v25
        }
        var v26 string = (v14 + "\n\t")
        v3 = (v13 % (((v3 % ((127) | 1))) | 1))
        _ = v26
    }
    if true {
        fmt.Println(v14, v3, v13)
        for fi27 := 0; fi27 < 20; fi27++ {
            switch fi27 {
            case -9223372036854775808:
                fmt.Println(fi27, v13)
            }
            fmt.Println(v3)
            _ = fi27
        }
    } else {
        var v28 uint32 = v15
        v3 -= v13
        if false {
            v14 = v14
            if (!(!true)) {
                v29 := GenS1{F0: -2147483648, F1: 664}
                fmt.Println(v14, v28, v28)
                _ = v29
            } else {
                var v30 uint16 = 65535
                _ = v30
            }
        } else {
            v28 = 42
        }
        _ = v28
    }
    var v31 float64 = (0.1 * 1.5)
    for fi32 := 0; fi32 < 13; fi32++ {
        v14 = ("foo bar" + "a")
        var v33 string = v14
        for fi34 := 0; fi34 < 2; fi34++ {
            v13 += (v13 << ((42) & 63))
            _ = fi34
        }
        _ = fi32
        _ = v33
    }
    v35 := GenS4{F0: 1, F1: 42, F2: 337}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v36 int16 = 1000
    if (!true) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            var v37 bool = ((v15 == (v15 ^ (v15 >> ((1) & 63)))) && (0 < (^(v15))))
            _ = v37
        }
    } else {
        var v38 uint = 65535
        _ = v38
    }
    switch v13 {
    case 953:
        for fi39 := 0; fi39 < 6; fi39++ {
            v40 := GenS5{F0: 65535}
            _ = fi39
            _ = v40
        }
    case 4294967295:
        switch (((v14 + "foo bar") + "a") + ((("" + v14) + ((v14 + "a") + ("" + v14))) + "a")) {
        case "foo bar":
            if len(v5) > 0 {
                v5[0] = v14
            }
        case "":
            fmt.Println(v13, v13, v13)
        case "a":
            if false {
                v41 := GenS6{F0: 892, F1: 4294967295, F2: 4294967295}
                var v42 uint64 = 18446744073709551615
                if len(v5) > 0 {
                    v5[0] = (("a" + ((("" + v14) + v14) + "a")) + v14)
                }
                _ = v41
                _ = v42
            } else {
                switch 2147483647 {
                case 32767:
                    switch (v14 + v14) {
                    case "":
                        fmt.Println(v14)
                    }
                }
                v43 := GenS3{F0: -1, F1: true}
                _ = v43
            }
        }
    case 127:
        if false {
            fmt.Println(v15, v31)
            _ = len(v4)
        }
    default:
        v44 := GenS5{F0: 65535}
        _ = v44
    }
    fmt.Println(v13, v3, v36)
    switch ("" + "") {
    case "hello":
        var v45 uint = 235
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
