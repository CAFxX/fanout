package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int64 }
type GenS2 struct { F0 int8; F1 uint; F2 uint32 }
func genFn0(p0 int, p1 int64) uint16 {
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := []int{}
        var v7 uint32 = 0
        _ = fi5
        _ = v6
        _ = v7
    }
    if (p0 < 42) {
        var v8 int = p0
        v9 := GenS1{F0: "", F1: 3, F2: 150}
        v10 := GenS1{F0: "foo bar", F1: 127, F2: p1}
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int16 = 42
        v13 := GenS1{F0: "hello", F1: 127, F2: -1}
        switch (p0 % ((int(p1)) | 1)) {
        case 1:
            var v14 int16 = (v12 * v12)
            _ = v14
        case 32767:
            var v15 int64 = (-(p1))
            _ = v15
        case 2147483647:
            p1 += (p1 ^ p1)
        default:
            var v16 bool = ((3.14159 < (1e10 / 1e10)) || ("foo bar" >= "a"))
            _ = v16
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    v1 := []int{128, 32767, 1}
    var v2 uint8 = 42
    v3 := []string{"\n\t", ("\n\t" + ((((("\n\t" + "\n\t") + ("\n\t" + "hello")) + "a") + "a") + ("" + "foo bar")))}
    v4 := []string{}
    if false {
        switch "a" {
        case "foo bar":
            v17 := GenS2{F0: 1, F1: 783, F2: 4294967295}
            _ = v17
        case "hello":
            fmt.Println(v2)
        case "":
            var v18 uint8 = v2
            _ = v18
        default:
            v1 = append(v1, 308)
        }
        var v19 int32 = -2147483648
        _ = v19
    } else {
        v2 = (v2 | (^(v2)))
        v20 := GenS1{F0: ("\n\t" + (((("a" + ("a" + "")) + ("" + "a")) + "foo bar") + ("foo bar" + "hello"))), F1: -128, F2: 9223372036854775807}
        switch 0 {
        case 128:
            v2 -= v2
        case 4294967295:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            if (!("\n\t" < "foo bar")) {
                var v21 uint16 = 701
                switch 1 {
                case 128:
                    var v22 uint8 = (v2 >> ((327) & 63))
                    _ = v22
                case -2147483648:
                    var v23 uint64 = 2
                    _ = v23
                case 127:
                    if false {
                        v24 := GenS1{F0: "hello", F1: -128, F2: -9223372036854775808}
                        _ = v24
                    } else {
                        v2 = (v2 * (^(v2)))
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if true {
                                if (2 >= 0) {
                                    for fi25 := 0; fi25 < 2; fi25++ {
                                        fmt.Println(fi25)
                                        _ = len(v1)
                                        fmt.Println(v21)
                                        _ = fi25
                                    }
                                    if (true && (false && (true || false))) {
                                        var v26 uint64 = 42
                                        var v27 int64 = 42
                                        var v28 bool = false
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                }
                                for fi29 := 0; fi29 < 6; fi29++ {
                                    switch (fi29 >> ((573) & 63)) {
                                    case 0:
                                        var v30 uint64 = 823
                                        _ = v30
                                    }
                                    var v31 int = fi29
                                    _ = fi29
                                    _ = v31
                                }
                                for fi32 := 0; fi32 < 5; fi32++ {
                                    var v33 uint64 = 509
                                    _ = fi32
                                    _ = v33
                                }
                            }
                        } else {
                            for fi34 := 0; fi34 < 1; fi34++ {
                                v21 = v21
                                _ = fi34
                            }
                            var v35 bool = false
                            for fi36 := 0; fi36 < 16; fi36++ {
                                switch 127 {
                                case 631:
                                    var v37 uint8 = 128
                                    _ = v37
                                default:
                                    v38 := GenS2{F0: 1, F1: 1, F2: 65535}
                                    _ = v38
                                }
                                _ = fi36
                            }
                            _ = v35
                        }
                    }
                default:
                    for fi39 := 0; fi39 < 13; fi39++ {
                        var v40 uint16 = 0
                        v2 += v2
                        var v41 uint32 = 506
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                }
                v21 -= 42
                _ = v21
            } else {
                var v42 int = 232
                var v43 uint64 = 18446744073709551615
                var v44 bool = (42 != 171)
                _ = v42
                _ = v43
                _ = v44
            }
        }
        _ = v20
    }
    if (0 <= -124) {
        fmt.Println(v2)
        switch -276 {
        case 32768:
            var v45 int = 229
            _ = v45
        }
    }
    for fi46 := 0; fi46 < 16; fi46++ {
        fmt.Println(fi46)
        _ = len(v1)
        _ = fi46
    }
    v2 = (v2 ^ v2)
    fmt.Println(v2)
    var v47 uint8 = v2
    var v48 string = "a"
    var v49 uint8 = v2
    if false {
        v49 = (v47 / ((42) | 1))
        for fi50 := 0; fi50 < 4; fi50++ {
            v2 += 252
            for fi51 := 0; fi51 < 7; fi51++ {
                var v52 float32 = 3.14159
                _ = fi51
                _ = v52
            }
            if (!false) {
                if (!(true || ((948 > 65535) || false))) {
                    v49 = (v2 | v49)
                }
                v49 += (v2 - (v49 ^ 1))
            }
            _ = fi50
        }
        fmt.Println(v49)
    }
    v53 := GenS1{F0: v48, F1: -1, F2: 1}
    if ((true && (!false)) || true) {
        if true {
            var v54 uint8 = (v2 + (^(v47)))
            _ = v54
        } else {
            _ = len(v3)
            _ = len(v4)
        }
    }
    var v55 uint16 = 411
    var v56 int32 = -482
    if true {
        v48 = ("\n\t" + "a")
        v47 -= (^(v49))
    } else {
        fmt.Println(v47, v49)
        v2 = (v47 | 1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v53
    _ = v55
    _ = v56
}
