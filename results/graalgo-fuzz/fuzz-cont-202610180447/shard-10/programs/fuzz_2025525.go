package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int32 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []uint32{2147483648, 1, 259}
    var v2 int32 = -292
    v3 := make(map[string]uint32)
    delete(v3, "a")
    _ = len(v3)
    fmt.Println(v2)
    var v4 uint = 2
    if true {
        _ = len(v1)
        var v5 int32 = v2
        if false {
            v6 := GenS1{F0: true, F1: 2, F2: 42}
            _ = v6
        } else {
            for fi7 := 0; fi7 < 1; fi7++ {
                var v8 int16 = 1
                _ = fi7
                _ = v8
            }
        }
        _ = v5
    }
    switch 2147483647 {
    case 32767:
        switch ("foo bar" + ("\n\t" + "foo bar")) {
        case "":
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 int8 = -37
                for fi11 := 0; fi11 < 18; fi11++ {
                    var v12 bool = false
                    var v13 bool = true
                    v12 = (v13 && (42 < 0))
                    var v14 uint16 = 1
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                if false {
                    v3["hello"] = 2147483648
                    v10 -= v10
                    v15 := GenS1{F0: false, F1: 128, F2: (v2 << ((2) & 63))}
                    _ = v15
                } else {
                    if false {
                        var v16 float64 = (3.14159 - 0.1)
                        switch (fi9 | 906) {
                        case 875:
                            if false {
                                switch int(v10) {
                                case 456:
                                    fmt.Println(fi9, v16)
                                default:
                                    v10 = 1
                                }
                                var v17 uint16 = 42
                                _ = v17
                            }
                        }
                        v16 = -2.25
                        _ = v16
                    }
                    v18 := GenS1{F0: true, F1: 128, F2: -763}
                    _ = v18
                }
                _ = fi9
                _ = v10
            }
        case "foo bar":
            v19 := GenS1{F0: (!(!(!(-9223372036854775808 != 1)))), F1: 42, F2: (v2 % ((-459) | 1))}
            _ = v19
        case "a":
            if false {
                v2 = v2
                for fi20 := 0; fi20 < 1; fi20++ {
                    for fi21 := 0; fi21 < 5; fi21++ {
                        delete(v3, ((("hello" + ("hello" + ("" + "a"))) + ("" + "\n\t")) + "\n\t"))
                        _ = len(v3)
                        _ = fi21
                    }
                    var v22 bool = true
                    _ = fi20
                    _ = v22
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (-128 == 127) {
                    for fi23 := 0; fi23 < 20; fi23++ {
                        var v24 float64 = 0.0
                        v4 = 301
                        for fi25 := 0; fi25 < 20; fi25++ {
                            v26 := GenS1{F0: (1000 > 929), F1: 42, F2: v2}
                            v24 += v24
                            _ = fi25
                            _ = v26
                        }
                        _ = fi23
                        _ = v24
                    }
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v28 := GenS1{F0: true, F1: 208, F2: -2147483648}
                        switch fi27 {
                        case -1:
                            v4 -= v4
                        }
                        fmt.Println(fi27, v2, v4)
                        _ = fi27
                        _ = v28
                    }
                }
            }
        }
    }
    if (-2.25 >= -2.25) {
        var v29 bool = ((v4 | (v4 >> (((v4 % ((v4) | 1))) & 63))) > (v4 | 4294967295))
        if false {
            for fi30 := 0; fi30 < 16; fi30++ {
                var v31 string = "hello"
                switch ((("hello" + v31) + "foo bar") + v31) {
                case "":
                    v32 := GenS1{F0: v29, F1: 255, F2: 467}
                    _ = v32
                case "foo bar":
                    switch fi30 {
                    case 128:
                        v33 := GenS1{F0: (false || v29), F1: 127, F2: 297}
                        _ = v33
                    case 989:
                        v34 := GenS2{F0: -9223372036854775808}
                        _ = v34
                    case 7:
                        v29 = (1 >= uint8(v4))
                    default:
                        var v35 uint32 = 65535
                        _ = v35
                    }
                }
                for fi36 := 0; fi36 < 16; fi36++ {
                    var v37 string = ""
                    _ = fi36
                    _ = v37
                }
                _ = fi30
                _ = v31
            }
        }
        _ = v29
    }
    if (true || false) {
        for fi38 := 0; fi38 < 8; fi38++ {
            var v39 int64 = 0
            for fi40 := 0; fi40 < 9; fi40++ {
                v39 -= v39
                var v41 float64 = ((float64(0.1) * 0.1) * 1e10)
                if (((v2 != (v2 << ((42) & 63))) || (18446744073709551615 > 0)) && true) {
                    if false {
                        var v42 float32 = (-2.25 * 3.14159)
                        var v43 int8 = -1
                        _ = v42
                        _ = v43
                    } else {
                        var v44 int32 = (v2 >> ((4294967295) & 63))
                        fmt.Println(v39)
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                        _ = v44
                    }
                }
                _ = fi40
                _ = v41
            }
            switch ("hello" + ("a" + "foo bar")) {
            case "a":
                v2 = v2
            default:
                v45 := GenS1{F0: (3.14159 <= 3.14159), F1: 1, F2: v2}
                _ = v45
            }
            _ = fi38
            _ = v39
        }
    } else {
        v2 = -1
    }
    v2 -= (v2 + (v2 << ((v4) & 63)))
    var v46 uint64 = 18446744073709551615
    v4 -= v4
    if (!(true && (32767 == 42))) {
        var v47 int64 = 42
        _ = v47
    }
    var v48 float64 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v48
}
