package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int64; F1 int64; F2 uint }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 string; F1 uint32 }

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    var v3 string = ""
    var v4 int16 = 0
    var v5 int = -2147483648
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            v3 = v3
            _ = fi7
        }
        var v8 uint16 = 1
        var v9 int = v5
        _ = fi6
        _ = v8
        _ = v9
    }
    if false {
        for fi10 := 0; fi10 < 11; fi10++ {
            v4 += 1
            if true {
                var v11 int32 = -1
                _ = v11
            } else {
                v12 := GenS1{F0: -2147483648}
                switch (v3 + (((v3 + ("hello" + (v3 + v3))) + (v3 + v3)) + "a")) {
                case "":
                    v13 := GenS1{F0: 2147483647}
                    _ = v13
                case "a":
                    switch (fi10 & (v5 | fi10)) {
                    case 4294967295:
                        if (true && (-2.25 == -2.25)) {
                            switch ((v3 + (("\n\t" + (v3 + v3)) + v3)) + ("foo bar" + "hello")) {
                            case "\n\t":
                                var v14 uint8 = 103
                                _ = v14
                            default:
                                if len(v1) > 0 {
                                    v1[0] = "hello"
                                }
                            }
                            var v15 int = (v5 * (v5 << ((845) & 63)))
                            _ = v15
                        } else {
                            var v16 int16 = (^(v4))
                            switch (-(fi10)) {
                            case 4294967295:
                                if ((v5 > (v5 >> ((444) & 63))) && ((0.0 / -2.25) < 3.14159)) {
                                    v17 := v2[("foo bar" + "a")]
                                    if v17 {
                                        var v18 uint8 = 2
                                        _ = v18
                                    } else {
                                        for fi19 := 0; fi19 < 17; fi19++ {
                                            if ((((("hello" + (v3 + v3)) + "") + v3) + (("\n\t" + "a") + ("foo bar" + "a"))) <= "") {
                                                v3 = "foo bar"
                                                v16 += v16
                                                v20 := GenS1{F0: -1}
                                                v21 := GenS1{F0: 2147483647}
                                                _ = v20
                                                _ = v21
                                            } else {
                                                delete(v2, "hello")
                                                if (!false) {
                                                    var v22 int8 = -109
                                                    _ = v22
                                                }
                                                v5 = (fi10 * (fi10 % ((42) | 1)))
                                            }
                                            if v17 {
                                                switch ("\n\t" + v3) {
                                                case "foo bar":
                                                    var v23 uint8 = 42
                                                    _ = v23
                                                case "a":
                                                    delete(v2, v3)
                                                case "":
                                                    if v17 {
                                                        v24 := GenS1{F0: -2147483648}
                                                        _ = v24
                                                    } else {
                                                        v4 -= 42
                                                        switch "hello" {
                                                        case "hello":
                                                            var v25 int8 = 127
                                                            _ = v25
                                                        case "a":
                                                            v2[v3] = (((v5 > 1) || false) && false)
                                                        }
                                                        var v26 uint = 255
                                                        _ = v26
                                                    }
                                                default:
                                                    v16 += -1
                                                }
                                            } else {
                                                var v27 uint64 = 1
                                                _ = v27
                                            }
                                            _ = fi19
                                        }
                                    }
                                    if (fi10 >= (fi10 ^ (v5 + fi10))) {
                                        for fi28 := 0; fi28 < 16; fi28++ {
                                            var v29 bool = (!(!true))
                                            v4 = (v4 * (v16 % ((v16) | 1)))
                                            _ = fi28
                                            _ = v29
                                        }
                                    } else {
                                        var v30 int16 = (-(v4))
                                        v31 := GenS2{F0: -128, F1: 138}
                                        _ = v30
                                        _ = v31
                                    }
                                    _ = v17
                                } else {
                                    for fi32 := 0; fi32 < 4; fi32++ {
                                        var v33 int = 999
                                        _ = fi32
                                        _ = v33
                                    }
                                    if false {
                                        var v34 bool = true
                                        if true {
                                            if true {
                                                var v35 uint8 = 225
                                                _ = v35
                                            }
                                        } else {
                                            v4 = (v4 + v4)
                                        }
                                        _ = v34
                                    }
                                }
                            }
                            _ = v16
                        }
                    case 32767:
                        v3 = (v3 + ("foo bar" + (v3 + v3)))
                    case 1:
                        v2["hello"] = (!false)
                    default:
                        switch v3 {
                        case "a":
                            for fi36 := 0; fi36 < 10; fi36++ {
                                var v37 bool = true
                                v38 := GenS3{F0: -48, F1: 9223372036854775807, F2: 18446744073709551615}
                                switch v3 {
                                case "hello":
                                    var v39 string = ((v3 + "foo bar") + v3)
                                    _ = v39
                                default:
                                    var v40 uint64 = 42
                                    _ = v40
                                }
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                        case "\n\t":
                            var v41 int = int(v4)
                            _ = v41
                        }
                    }
                default:
                    v5 = fi10
                }
                if (true && ((!false) && true)) {
                    v42 := GenS4{F0: (v4 * (v4 + (v4 << ((0) & 63)))), F1: v4}
                    v4 = v4
                    _ = v42
                }
                _ = v12
            }
            _ = fi10
        }
        for fi43 := 0; fi43 < 16; fi43++ {
            v5 = (fi43 ^ (-(v5)))
            delete(v2, ("" + v3))
            _ = fi43
        }
    }
    v44 := GenS4{F0: v4, F1: -1}
    v5 += v5
    if false {
        v4 += (v4 >> ((539) & 63))
    }
    v45 := GenS5{F0: ((v3 + v3) + "foo bar"), F1: 198}
    var v46 uint = 0
    var v47 int16 = 1000
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
