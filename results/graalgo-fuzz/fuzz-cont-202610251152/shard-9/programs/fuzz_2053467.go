package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int; F1 uint; F2 int16 }
type GenS3 struct { F0 int64; F1 bool; F2 int }
type GenS4 struct { F0 uint8; F1 string }
func genFn0(p0 uint32, p1 int32, p2 int32) uint32 {
    v5 := []int{841, 32767, 2147483647, 32767}
    var v6 int8 = 1
    _ = len(v5)
    switch 736 {
    case -9223372036854775808:
        if len(v5) > 0 {
            v5[0] = 635
        }
    }
    if false {
        var v7 int32 = p1
        v8 := GenS1{F0: 440, F1: -1}
        var v9 int16 = 68
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 62
}

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    v4 := []uint32{124, uint32(v1), 973, 1}
    v4 = append(v4, 0)
    if (!(false && (1.5 != 1e10))) {
        var v10 uint64 = 42
        v11 := GenS1{F0: 32767, F1: -891}
        fmt.Println(v1)
        _ = v10
        _ = v11
    }
    v1 -= v1
    if (127 <= 42) {
        if (!false) {
            v12 := GenS2{F0: 873, F1: 4294967295, F2: 918}
            _ = v12
        } else {
            var v13 float32 = 0.1
            _ = v13
        }
    }
    v1 = (v1 % (((v1 % (((^(v1))) | 1))) | 1))
    if (true || (!true)) {
        var v14 float32 = -2.25
        _ = v14
    }
    if true {
        v15 := GenS1{F0: 1, F1: -9223372036854775808}
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        _ = v15
    } else {
        _ = len(v2)
        for fi16 := 0; fi16 < 19; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                switch "a" {
                case "":
                    var v18 uint32 = 65535
                    _ = v18
                case "hello":
                    if (false && (true || false)) {
                        v1 = (^(v1))
                        fmt.Println(v1, fi17, v1)
                    }
                case "\n\t":
                    var v19 float32 = 3.14159
                    _ = v19
                }
                var v20 uint16 = 469
                switch (fi17 & -2147483648) {
                case 226:
                    var v21 uint64 = 0
                    _ = v21
                }
                _ = fi17
                _ = v20
            }
            switch (fi16 * fi16) {
            case 9223372036854775807:
                var v22 uint32 = 42
                _ = v22
            case -1:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 127:
                if ((-9223372036854775808 > (fi16 % (((fi16 ^ fi16)) | 1))) || ((1 != 0) && false)) {
                    if (("hello" > ("a" + "hello")) || (-73 >= 0)) {
                        var v23 int16 = 742
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 7; fi24++ {
                            if (-1 != 0) {
                                switch ("a" + (("foo bar" + ("\n\t" + "foo bar")) + ("\n\t" + "\n\t"))) {
                                case "\n\t":
                                    if true {
                                        v25 := GenS2{F0: 32768, F1: (^(v1)), F2: 1000}
                                        _ = len(v4)
                                        v26 := GenS2{F0: 4294967295, F1: (v1 - v1), F2: 32767}
                                        _ = v25
                                        _ = v26
                                    }
                                case "a":
                                    var v27 bool = true
                                    _ = v27
                                case "foo bar":
                                    if false {
                                        v28 := v3[65535]
                                        var v29 int8 = 103
                                        v30 := GenS3{F0: -761, F1: true, F2: (v28 & (fi24 / (((v28 | (fi24 / ((v28) | 1)))) | 1)))}
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    }
                                }
                                v31 := GenS1{F0: 65535, F1: 9223372036854775807}
                                _ = v31
                            } else {
                                v3[0] = fi16
                                v32 := v2[2147483647]
                                _ = v32
                            }
                            v1 = (^(v1))
                            v33 := GenS4{F0: 179, F1: "\n\t"}
                            _ = fi24
                            _ = v33
                        }
                        var v34 bool = (56 != 63)
                        _ = v34
                    }
                }
            }
            _ = fi16
        }
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
