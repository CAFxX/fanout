package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 string }
func genFn0() uint8 {
    v7 := []int{42, -1, 42, 128}
    v8 := GenS1{F0: "\n\t"}
    var v9 int8 = -1
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int = -9223372036854775808
    var v2 float64 = 0.0
    var v3 uint8 = 128
    v4 := make(map[uint32]uint32)
    var v5 int = v1
    var v6 int = v1
    var v10 bool = true
    var v11 uint64 = 42
    if (65535 < 263) {
        for fi12 := 0; fi12 < 8; fi12++ {
            switch "hello" {
            case "hello":
                var v13 uint64 = v11
                _ = v13
            default:
                fmt.Println(v3, v1)
            }
            _ = fi12
        }
        var v14 bool = (v3 <= 0)
        _ = v14
    } else {
        switch (v6 ^ v5) {
        case -692:
            var v15 bool = (!((v10 || (!false)) || v10))
            _ = v15
        case 0:
            var v16 uint8 = (^(v3))
            _ = v16
        case 9223372036854775807:
            fmt.Println(v11, v11, v2)
        default:
            v6 = v1
        }
    }
    v3 += 128
    v17 := GenS1{F0: (("hello" + ("foo bar" + "a")) + "hello")}
    var v18 uint8 = v3
    delete(v4, 1)
    if v10 {
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 float64 = v2
            if (v20 > v2) {
                v10 = false
                if (968 == 4294967295) {
                    v21 := []int{}
                    var v22 int64 = -693
                    v23 := GenS2{F0: 2147483647, F1: "\n\t"}
                    _ = v21
                    _ = v22
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 19; fi24++ {
                        var v25 uint = 4294967295
                        var v26 int64 = 9223372036854775807
                        if v10 {
                            v6 = v6
                            v3 = uint8(v11)
                        } else {
                            fmt.Println(fi19, v10)
                            var v27 string = "hello"
                            _ = v27
                        }
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    v3 -= v3
                }
            } else {
                if v10 {
                    v28 := []int{v1, v1}
                    switch (v6 << ((977) & 63)) {
                    case 1:
                        switch ("a" + (("" + ("foo bar" + "a")) + ("" + (("foo bar" + (("a" + ("\n\t" + ("hello" + ""))) + "foo bar")) + "foo bar")))) {
                        case "\n\t":
                            if v10 {
                                for fi29 := 0; fi29 < 1; fi29++ {
                                    var v30 float64 = ((v20 * 1e10) + ((v2 * 1e10) - float64(0.1)))
                                    var v31 bool = ((0.0 * 0.1) >= -2.25)
                                    v32 := GenS1{F0: "a"}
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                                v33 := GenS1{F0: "\n\t"}
                                for fi34 := 0; fi34 < 11; fi34++ {
                                    if v10 {
                                        var v35 uint = 0
                                        _ = len(v28)
                                        _ = v35
                                    }
                                    var v36 int64 = -1
                                    v36 -= (v36 - v36)
                                    _ = fi34
                                    _ = v36
                                }
                                _ = v33
                            } else {
                                if (v6 <= (^(v5))) {
                                    if (0 != -1) {
                                        v37 := GenS2{F0: 42, F1: "\n\t"}
                                        fmt.Println(fi19, v20, v2)
                                        _ = v37
                                    }
                                } else {
                                    var v38 uint32 = 2147483648
                                    _ = v38
                                }
                                var v39 float32 = 0.1
                                _ = v39
                            }
                        default:
                            var v40 uint16 = 37
                            _ = v40
                        }
                    case 9223372036854775807:
                        if len(v28) > 0 {
                            v28[0] = v1
                        }
                    }
                    _ = v28
                } else {
                    delete(v4, 0)
                }
            }
            for fi41 := 0; fi41 < 1; fi41++ {
                var v42 int16 = 42
                v11 = 18446744073709551615
                _ = fi41
                _ = v42
            }
            _ = fi19
            _ = v20
        }
    } else {
        var v43 int = v6
        _ = v43
    }
    if (4294967295 > v5) {
        var v44 int = int(v11)
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
