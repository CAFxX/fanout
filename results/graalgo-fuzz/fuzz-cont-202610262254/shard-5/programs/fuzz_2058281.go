package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 uint64; F1 int }
func genFn0(p0 uint32) uint64 {
    if (-48 != -65) {
        v5 := GenS1{F0: 137, F1: 42}
        var v6 int64 = 9223372036854775807
        _ = v5
        _ = v6
    }
    if false {
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := []int{}
            switch "\n\t" {
            case "":
                var v9 uint32 = 1
                _ = v9
            case "foo bar":
                for fi10 := 0; fi10 < 20; fi10++ {
                    v11 := GenS1{F0: 1, F1: 127}
                    _ = fi10
                    _ = v11
                }
            case "\n\t":
                fmt.Println(p0)
            }
            if len(v8) > 0 {
                v8[0] = 4294967295
            }
            _ = fi7
            _ = v8
        }
    }
    if ((-897 >= -9223372036854775808) || false) {
        v12 := GenS2{F0: 867, F1: 1}
        _ = v12
    } else {
        var v13 int = 1
        var v14 int8 = 100
        switch (v13 ^ v13) {
        case 923:
            var v15 uint64 = 4294967295
            _ = v15
        case 685:
            v13 = 9223372036854775807
        case 468:
            if false {
                switch "\n\t" {
                case "":
                    for fi16 := 0; fi16 < 19; fi16++ {
                        if true {
                            var v17 int = (fi16 * fi16)
                            if (!(-232 < -1)) {
                                if ((3.14159 > (float32(0.1) / 1.5)) || ((((("" + ("hello" + ("hello" + ("foo bar" + ("" + ("hello" + ("a" + "hello"))))))) + "hello") >= "\n\t") && true) || true)) {
                                    if (!(true && (0 <= (^(p0))))) {
                                        fmt.Println(p0, v14)
                                        var v18 string = (("hello" + "a") + "foo bar")
                                        v13 -= fi16
                                        _ = v18
                                    }
                                    v13 = v13
                                } else {
                                    var v19 string = ""
                                    switch (("foo bar" + (v19 + "foo bar")) + (((((v19 + "\n\t") + ((v19 + v19) + ("foo bar" + v19))) + v19) + (v19 + "a")) + v19)) {
                                    case "hello":
                                        fmt.Println(fi16, fi16, v17)
                                    case "":
                                        switch v17 {
                                        case 808:
                                            v20 := GenS3{F0: 0, F1: 0}
                                            _ = v20
                                        case 1:
                                            var v21 uint32 = (p0 - 1)
                                            _ = v21
                                        default:
                                            if (uint(p0) <= 65535) {
                                                v22 := []int{fi16, (v17 * fi16), (v13 & v17)}
                                                for fi23 := 0; fi23 < 3; fi23++ {
                                                    p0 += (p0 | 194)
                                                    var v24 uint = 0
                                                    if (!false) {
                                                        var v25 uint16 = 255
                                                        var v26 int64 = 1
                                                        _ = v25
                                                        _ = v26
                                                    }
                                                    _ = fi23
                                                    _ = v24
                                                }
                                                _ = v22
                                            }
                                        }
                                    }
                                    if (-31 >= 1000) {
                                        var v27 int = 2147483647
                                        fmt.Println(v19)
                                        var v28 float64 = 1.5
                                        _ = v27
                                        _ = v28
                                    }
                                    _ = v19
                                }
                                if true {
                                    fmt.Println(v17, p0, fi16)
                                }
                            } else {
                                var v29 int = v13
                                v14 -= (v14 << ((844) & 63))
                                var v30 uint64 = 42
                                _ = v29
                                _ = v30
                            }
                            _ = v17
                        }
                        _ = fi16
                    }
                }
                var v31 float32 = 0.1
                v32 := []int{v13, (v13 & v13), (v13 + v13)}
                _ = v31
                _ = v32
            }
        }
        _ = v13
        _ = v14
    }
    v33 := GenS2{F0: -32768, F1: 437}
    v34 := GenS3{F0: 18446744073709551615, F1: 929}
    _ = p0
    _ = v33
    _ = v34
    return 0
}
func genFn1(p0 string, p1 uint) int64 {
    var v35 int = 668
    for fi36 := 0; fi36 < 17; fi36++ {
        for fi37 := 0; fi37 < 3; fi37++ {
            var v38 int8 = -32
            p1 = (^(p1))
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    var v39 string = ("\n\t" + ((("a" + p0) + p0) + p0))
    v40 := GenS4{F0: 32768}
    for fi41 := 0; fi41 < 7; fi41++ {
        v42 := GenS4{F0: (fi41 ^ -440)}
        _ = fi41
        _ = v42
    }
    _ = p0
    _ = p1
    _ = v35
    _ = v39
    _ = v40
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    v3 := []string{"\n\t", "", ("hello" + "\n\t"), ""}
    v4 := make(map[string]string)
    var v43 int64 = 214
    var v44 uint8 = 2
    fmt.Println(v43)
    var v45 uint64 = 0
    fmt.Println(v43, v45)
    fmt.Println(v43)
    v46 := GenS5{F0: -83, F1: 32768}
    v44 -= v44
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v44 += (v44 >> ((1) & 63))
    v43 = (v43 | (v43 | (v43 - 42)))
    _ = len(v3)
    fmt.Println(v44)
    v47 := GenS6{F0: (v45 + (v45 | (v45 - v45))), F1: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
