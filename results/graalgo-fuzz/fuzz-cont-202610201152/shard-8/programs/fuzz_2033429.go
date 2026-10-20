package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint = 2
    v2 := []string{}
    v3 := []uint32{uint32(v1)}
    var v4 int = -153
    var v5 int = (v4 * (v4 % ((-368) | 1)))
    if (false && true) {
        if false {
            if true {
                var v6 bool = ("\n\t" == ("a" + ("foo bar" + "")))
                _ = v6
            } else {
                v7 := GenS1{F0: 65535}
                _ = v7
            }
            v4 = (v4 << (((v1 ^ v1)) & 63))
            if len(v3) > 0 {
                v3[0] = uint32(v1)
            }
        }
    }
    if (false || (661 == 42)) {
        v8 := GenS1{F0: 1}
        _ = v8
    } else {
        _ = len(v2)
    }
    v4 += v4
    if (!true) {
        if (true || (18446744073709551615 <= 42)) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v5 += v4
            _ = len(v3)
        }
    } else {
        var v9 float32 = 1.5
        if true {
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS1{F0: 375}
                var v12 int32 = 0
                switch fi10 {
                case 32768:
                    v13 := GenS1{F0: 1}
                    _ = v13
                default:
                    fmt.Println(v1)
                }
                _ = fi10
                _ = v11
                _ = v12
            }
            var v14 uint64 = 18446744073709551615
            _ = len(v3)
            _ = v14
        } else {
            v15 := GenS2{F0: -1, F1: 677}
            if ((42 < -9223372036854775808) && false) {
                var v16 int32 = -147
                _ = v16
            } else {
                for fi17 := 0; fi17 < 18; fi17++ {
                    if true {
                        var v18 bool = true
                        if len(v2) > 0 {
                            v2[0] = ((("foo bar" + "a") + "a") + "foo bar")
                        }
                        for fi19 := 0; fi19 < 20; fi19++ {
                            switch ("hello" + (((("" + "foo bar") + "foo bar") + "foo bar") + ("a" + "hello"))) {
                            case "a":
                                if v18 {
                                    v20 := GenS3{F0: 42, F1: 255}
                                    _ = v20
                                }
                            }
                            _ = fi19
                        }
                        _ = v18
                    }
                    if (-128 == -128) {
                        for fi21 := 0; fi21 < 20; fi21++ {
                            var v22 uint8 = 1
                            v22 += v22
                            switch ("a" + "") {
                            case "hello":
                                _ = len(v2)
                            case "":
                                var v23 bool = (!(741 <= 42))
                                _ = v23
                            }
                            _ = fi21
                            _ = v22
                        }
                    }
                    var v24 int8 = 100
                    _ = fi17
                    _ = v24
                }
                switch "hello" {
                case "hello":
                    if (true && false) {
                        var v25 float32 = 0.1
                        _ = v25
                    } else {
                        if true {
                            var v26 uint32 = 57
                            for fi27 := 0; fi27 < 1; fi27++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi27
                            }
                            if (true && false) {
                                for fi28 := 0; fi28 < 18; fi28++ {
                                    fmt.Println(v4, v5)
                                    _ = fi28
                                }
                                fmt.Println(v9)
                            }
                            _ = v26
                        } else {
                            v5 -= -2147483648
                            var v29 float32 = v9
                            _ = v29
                        }
                    }
                case "\n\t":
                    v5 += 2147483647
                }
                v2 = append(v2, "")
            }
            var v30 float64 = 0.1
            _ = v15
            _ = v30
        }
        _ = v9
    }
    fmt.Println(v1, v5)
    _ = len(v2)
    var v31 uint64 = 42
    switch "" {
    case "":
        var v32 int32 = -1
        _ = v32
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "hello":
        v1 += (v1 & (v1 + (v1 & v1)))
    }
    if true {
        v33 := GenS2{F0: (^(v4)), F1: 1}
        _ = v33
    } else {
        for fi34 := 0; fi34 < 17; fi34++ {
            for fi35 := 0; fi35 < 14; fi35++ {
                if len(v2) > 0 {
                    v2[0] = ("\n\t" + "foo bar")
                }
                var v36 string = "\n\t"
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
        fmt.Println(v31, v31)
    }
    v4 -= v4
    v37 := GenS3{F0: 613, F1: 1}
    var v38 string = ""
    var v39 int = v4
    v4 -= -989
    if false {
        v38 = (v38 + v38)
        fmt.Println(v38)
        switch v38 {
        case "a":
            switch (v38 + (v38 + "")) {
            case "foo bar":
                var v40 int16 = 1
                _ = v40
            case "\n\t":
                var v41 string = v38
                _ = v41
            case "":
                if (!true) {
                    v42 := GenS4{F0: 318}
                    var v43 int32 = -28
                    _ = v42
                    _ = v43
                }
            default:
                v5 -= -2147483648
            }
        default:
            var v44 int32 = 42
            _ = v44
        }
    } else {
        var v45 int8 = 65
        v5 = (^(v4))
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
