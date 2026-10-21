package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint32) uint32 {
    p0 -= (p0 - p0)
    if false {
        v5 := []int{0, -1}
        var v6 int = 32768
        _ = v5
        _ = v6
    }
    p0 = p0
    p0 = p0
    _ = p0
    return p0
}

func main() {
    var v1 int = -510
    var v2 uint8 = 127
    v3 := []string{(("hello" + ("foo bar" + "a")) + "\n\t"), "\n\t", "\n\t", (("a" + ("hello" + ((("" + ("" + "a")) + "") + ("\n\t" + "foo bar")))) + (("a" + (("hello" + ((("a" + "") + ("foo bar" + "\n\t")) + "hello")) + "foo bar")) + "foo bar"))}
    v4 := make(map[string]int)
    var v7 uint32 = 2
    for fi8 := 0; fi8 < 16; fi8++ {
        v7 += uint32(v2)
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: 42, F1: 755}
            v2 = v2
            if (true || (90 <= 42)) {
                v2 = (v2 ^ v2)
            }
            v4["foo bar"] = fi8
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := v4["hello"]
    v12 := GenS1{F0: 1, F1: 0}
    v13 := GenS1{F0: 68, F1: 0}
    var v14 int = v11
    fmt.Println(v14, v11)
    delete(v4, ((("foo bar" + "a") + "hello") + ("\n\t" + ("foo bar" + ""))))
    var v15 string = "a"
    for fi16 := 0; fi16 < 3; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            if false {
                v18 := GenS2{F0: 127, F1: v2}
                v4[v15] = (v14 ^ -226)
                v3 = append(v3, "a")
                v19 := GenS1{F0: -9223372036854775808, F1: 32767}
                _ = v18
                _ = v19
            }
            if (-909 < 32767) {
                if (false || false) {
                    var v20 float32 = 0.1
                    if ("foo bar" != v15) {
                        if true {
                            v2 = v2
                            var v21 uint16 = 2
                            _ = v21
                        }
                    } else {
                        fmt.Println(v15, v2)
                        var v22 uint = 507
                        delete(v4, v15)
                        _ = v22
                    }
                    v23 := GenS2{F0: 74, F1: 128}
                    v24 := GenS2{F0: 1, F1: 2}
                    _ = v20
                    _ = v23
                    _ = v24
                } else {
                    if false {
                        v25 := GenS1{F0: -377, F1: 42}
                        v7 = genFn0(genFn0((v7 + 2147483648)))
                        _ = v25
                    } else {
                        fmt.Println(v1, v14, fi16)
                    }
                }
                var v26 uint32 = v7
                _ = v26
            } else {
                fmt.Println(v15)
                var v27 int8 = 42
                _ = v27
            }
            fmt.Println(v2, v2)
            _ = fi17
        }
        v28 := GenS3{F0: 42}
        _ = fi16
        _ = v28
    }
    fmt.Println(v14, v7, v7)
    if (255 > 788) {
        for fi29 := 0; fi29 < 16; fi29++ {
            var v30 uint64 = 2
            var v31 uint32 = 453
            _ = fi29
            _ = v30
            _ = v31
        }
        delete(v4, "a")
        var v32 float32 = 1.5
        _ = v32
    } else {
        for fi33 := 0; fi33 < 9; fi33++ {
            switch -680 {
            case -150:
                switch v15 {
                case "foo bar":
                    var v34 int16 = 32767
                    _ = v34
                default:
                    var v35 uint = uint(v7)
                    _ = v35
                }
            }
            for fi36 := 0; fi36 < 8; fi36++ {
                var v37 uint32 = (v7 / (((v7 | (^(v7)))) | 1))
                v2 = (v2 << ((121) & 63))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi36
                _ = v37
            }
            if (!false) {
                for fi38 := 0; fi38 < 15; fi38++ {
                    var v39 uint = 0
                    var v40 bool = true
                    var v41 int64 = -1
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                }
                v42 := GenS3{F0: 18446744073709551615}
                v1 = 264
                _ = v42
            }
            _ = fi33
        }
        v7 += uint32(v2)
        v14 += v14
    }
    v43 := v4[v15]
    var v44 uint64 = uint64(v7)
    v11 += -9223372036854775808
    var v45 float32 = 1e10
    if (v44 >= v44) {
        v46 := GenS3{F0: 4294967295}
        fmt.Println(v44)
        _ = v46
    }
    v47 := v4[v15]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v47
    _ = v7
}
