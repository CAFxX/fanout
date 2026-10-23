package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 string }
type GenS2 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint16; F1 string; F2 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    var v5 uint64 = 220
    v5 = v5
    fmt.Println(v5)
    v6 := GenS1{F0: 42, F1: 2, F2: (((("" + ("foo bar" + "foo bar")) + "\n\t") + "hello") + "\n\t")}
    if (3.14159 <= float64((3.14159 * -2.25))) {
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: -1, F1: 42, F2: "a"}
            if true {
                var v9 int8 = 55
                _ = v9
            } else {
                fmt.Println(v5, fi7)
            }
            _ = fi7
            _ = v8
        }
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v5 -= v5
        v11 := GenS1{F0: 127, F1: 330, F2: ""}
        _ = fi10
        _ = v11
    }
    v5 = v5
    v5 -= 1
    if (((("\n\t" + ("\n\t" + "")) + ("" + "")) + ("hello" + "")) != ("a" + ("hello" + ""))) {
        for fi12 := 0; fi12 < 10; fi12++ {
            v5 = (^(v5))
            _ = fi12
        }
        switch "" {
        case "":
            var v13 uint64 = v5
            _ = v13
        default:
            var v14 uint8 = 2
            _ = v14
        }
        if (("a" < ("hello" + "a")) && true) {
            if true {
                for fi15 := 0; fi15 < 11; fi15++ {
                    v5 = (v5 * v5)
                    _ = fi15
                }
                v5 = v5
                v16 := GenS1{F0: 116, F1: 2, F2: ("foo bar" + ((("" + ("foo bar" + ("hello" + "foo bar"))) + "a") + ("\n\t" + "")))}
                _ = v16
            } else {
                if (288 != 377) {
                    v5 -= (v5 << ((0) & 63))
                    v5 = (^(v5))
                } else {
                    var v17 uint = 255
                    for fi18 := 0; fi18 < 19; fi18++ {
                        v19 := GenS1{F0: -68, F1: uint32(v5), F2: "\n\t"}
                        for fi20 := 0; fi20 < 19; fi20++ {
                            if false {
                                var v21 uint = v17
                                _ = v21
                            } else {
                                switch (fi20 + (^(fi18))) {
                                case -2147483648:
                                    fmt.Println(v5, fi18)
                                case 32768:
                                    v17 -= v17
                                case 4294967295:
                                    v22 := GenS2{F0: (150 > 1), F1: uint16(v17), F2: 60}
                                    _ = v22
                                default:
                                    var v23 string = "a"
                                    _ = v23
                                }
                                if ((v17 & (v17 % ((241) | 1))) == (v17 ^ (v17 >> (((v17 & v17)) & 63)))) {
                                    if (!(!(!(false && (false && ((v17 <= (^(v17))) || false)))))) {
                                        var v24 float64 = float64(0.0)
                                        var v25 int = fi18
                                        v26 := []int{int(v17), 127, 4294967295}
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    }
                                    var v27 int = (fi18 ^ (fi18 % (((fi18 ^ (fi18 + fi18))) | 1)))
                                    var v28 float64 = ((3.14159 * float64(0.1)) - 1.5)
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v29 string = "\n\t"
                                    _ = v29
                                }
                                v17 += v17
                            }
                            var v30 int64 = -1
                            v31 := GenS1{F0: 42, F1: 0, F2: "\n\t"}
                            fmt.Println(v5, v17)
                            switch 128 {
                            case 9223372036854775807:
                                v5 -= 42
                            case 4294967295:
                                v30 -= v30
                                fmt.Println(v5)
                                if false {
                                    var v32 float64 = ((-2.25 - -2.25) / 1e10)
                                    v33 := GenS1{F0: -1, F1: 4294967295, F2: "hello"}
                                    v34 := v1[(("hello" + "\n\t") + "foo bar")]
                                    v5 += v5
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                                v5 = v5
                            }
                            _ = fi20
                            _ = v30
                            _ = v31
                        }
                        _ = fi18
                        _ = v19
                    }
                    _ = v17
                }
                v35 := GenS3{F0: 32767, F1: ("\n\t" + ""), F2: -1}
                for fi36 := 0; fi36 < 15; fi36++ {
                    v37 := []int{(^(fi36))}
                    _ = len(v3)
                    var v38 uint8 = 127
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                _ = v35
            }
            v5 = (v5 / ((42) | 1))
        } else {
            for fi39 := 0; fi39 < 7; fi39++ {
                v40 := []int{(^(fi39)), fi39, (fi39 | 4294967295), (fi39 ^ -846)}
                _ = fi39
                _ = v40
            }
        }
    } else {
        v5 += (v5 >> ((65535) & 63))
        switch "" {
        case "\n\t":
            v5 = v5
        default:
            if (1 == 127) {
                fmt.Println(v5)
            }
        }
        var v41 uint32 = 613
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
