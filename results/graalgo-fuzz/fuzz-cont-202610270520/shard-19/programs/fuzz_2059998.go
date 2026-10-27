package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint; F1 int; F2 uint32 }

func main() {
    v1 := []string{}
    v2 := make(map[int]string)
    v3 := []string{("a" + "foo bar"), "a"}
    v4 := make(map[uint32]bool)
    var v5 uint64 = 697
    v5 += v5
    delete(v2, 2147483647)
    var v6 int64 = 9223372036854775807
    if ("a" > ("a" + "hello")) {
        switch "hello" {
        case "foo bar":
            fmt.Println(v5)
        case "":
            v3 = append(v3, ("a" + (("hello" + "") + "foo bar")))
        }
        if len(v3) > 0 {
            v3[0] = "a"
        }
    }
    var v7 int64 = (^(v6))
    if true {
        var v8 int8 = 42
        v9 := v2[2147483647]
        _ = v8
        _ = v9
    }
    if (!(925 < 146)) {
        if (true || true) {
            v10 := v4[750]
            fmt.Println(v10, v6)
            _ = v10
        } else {
            switch (("hello" + (("hello" + (("foo bar" + "foo bar") + "hello")) + "a")) + ("\n\t" + ("foo bar" + ("" + "a")))) {
            case "foo bar":
                v7 += v6
            case "\n\t":
                for fi11 := 0; fi11 < 14; fi11++ {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        var v13 bool = false
                        _ = fi12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 8; fi14++ {
                        v6 = 903
                        v7 = v6
                        var v15 uint8 = 128
                        _ = fi14
                        _ = v15
                    }
                    _ = fi11
                }
            }
        }
        var v16 uint = 2
        _ = v16
    } else {
        v7 = (v7 * v7)
        var v17 float32 = 0.0
        _ = v17
    }
    if ("hello" > "foo bar") {
        if true {
            delete(v2, 0)
        } else {
            for fi18 := 0; fi18 < 11; fi18++ {
                v7 = 42
                var v19 uint64 = (v5 ^ v5)
                if (129 <= 255) {
                    var v20 int16 = 933
                    fmt.Println(v5, fi18, v19)
                    if (225 != 127) {
                        var v21 uint = 737
                        var v22 uint32 = 299
                        _ = v21
                        _ = v22
                    } else {
                        var v23 uint8 = 1
                        fmt.Println(v20)
                        _ = v23
                    }
                    _ = v20
                }
                _ = fi18
                _ = v19
            }
            for fi24 := 0; fi24 < 10; fi24++ {
                v3 = append(v3, "")
                var v25 uint64 = v5
                var v26 int = fi24
                _ = fi24
                _ = v25
                _ = v26
            }
            switch "\n\t" {
            case "":
                for fi27 := 0; fi27 < 1; fi27++ {
                    if (!(2 > 255)) {
                        if false {
                            switch "" {
                            case "\n\t":
                                v5 = 0
                            case "":
                                if true {
                                    var v28 float32 = 3.14159
                                    _ = v28
                                }
                            default:
                                var v29 int16 = 42
                                _ = v29
                            }
                        } else {
                            v5 = 4294967295
                        }
                        v30 := GenS1{F0: fi27, F1: 42, F2: 1}
                        _ = v30
                    } else {
                        if false {
                            _ = len(v1)
                            v5 = (v5 + (v5 << ((65535) & 63)))
                        }
                    }
                    for fi31 := 0; fi31 < 10; fi31++ {
                        var v32 int8 = 29
                        _ = fi31
                        _ = v32
                    }
                    _ = fi27
                }
            case "a":
                var v33 int64 = v6
                _ = v33
            case "hello":
                v1 = append(v1, "")
            default:
                v34 := GenS1{F0: 0, F1: -1, F2: 42}
                _ = v34
            }
        }
        var v35 float64 = ((0.1 - 1e10) / 3.14159)
        _ = v35
    }
    if (false && true) {
        v2[-2147483648] = ""
    }
    for fi36 := 0; fi36 < 16; fi36++ {
        var v37 float32 = 1e10
        switch "foo bar" {
        case "hello":
            switch "\n\t" {
            case "":
                for fi38 := 0; fi38 < 20; fi38++ {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    if (42 > 1) {
                        v39 := GenS1{F0: fi38, F1: 42, F2: 127}
                        v37 = float32(1.5)
                        var v40 bool = (-128 > 42)
                        _ = v39
                        _ = v40
                    } else {
                        fmt.Println(fi38, fi38)
                        var v41 string = ("a" + "\n\t")
                        delete(v2, fi38)
                        _ = v41
                    }
                    var v42 int64 = (v6 ^ v7)
                    _ = fi38
                    _ = v42
                }
            }
        case "\n\t":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        switch (-(fi36)) {
        case 2147483647:
            var v43 uint = 18446744073709551615
            _ = v43
        case 128:
            if true {
                var v44 uint = 261
                switch -1 {
                case 32767:
                    var v45 int16 = -545
                    _ = v45
                default:
                    var v46 bool = (((-1 != 32767) || false) && (42 > 741))
                    _ = v46
                }
                _ = v44
            } else {
                var v47 uint32 = 4294967295
                v37 -= v37
                _ = v47
            }
        case 440:
            var v48 float64 = -2.25
            _ = v48
        default:
            for fi49 := 0; fi49 < 4; fi49++ {
                v50 := GenS2{F0: 1, F1: fi49, F2: 2147483648}
                v1 = append(v1, "foo bar")
                if true {
                    _ = len(v2)
                    v5 = v5
                    fmt.Println(v5, fi49)
                }
                v51 := v4[2147483648]
                _ = fi49
                _ = v50
                _ = v51
            }
        }
        _ = fi36
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
