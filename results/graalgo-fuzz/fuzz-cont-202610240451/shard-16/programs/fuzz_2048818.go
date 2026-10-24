package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int32; F1 int16; F2 int16 }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0(p0 int32) int64 {
    v7 := []int{int(p0), -9223372036854775808, 42, 128}
    var v8 int64 = 9223372036854775807
    _ = len(v7)
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 string, p1 uint32, p2 bool) uint64 {
    p0 = p0
    p2 = ((true && (32767 == 42)) && false)
    if (p2 && p2) {
        var v9 int16 = -413
        _ = v9
    } else {
        for fi10 := 0; fi10 < 19; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                p1 = p1
                p2 = p2
                if p2 {
                    var v12 uint = 255
                    _ = v12
                } else {
                    var v13 int = 38
                    v14 := GenS1{F0: p0, F1: (("" + p0) + ((p0 + "hello") + ((p0 + p0) + "")))}
                    _ = v13
                    _ = v14
                }
                _ = fi11
            }
            p0 = p0
            _ = fi10
        }
        var v15 uint = 255
        var v16 int32 = -979
        _ = v15
        _ = v16
    }
    switch p0 {
    case "foo bar":
        for fi17 := 0; fi17 < 18; fi17++ {
            p1 = 0
            p0 = p0
            p1 += (p1 / (((p1 | p1)) | 1))
            _ = fi17
        }
    case "":
        p2 = (!false)
    case "a":
        v18 := []int{}
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        if (p1 == p1) {
            fmt.Println(p2, fi19, fi19)
            var v20 int32 = 1
            _ = v20
        } else {
            v21 := GenS2{F0: 2147483647, F1: -32768, F2: 1000}
            fmt.Println(fi19)
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 int = fi19
                _ = fi22
                _ = v23
            }
            _ = v21
        }
        _ = fi19
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    var v1 int16 = 1
    var v2 float32 = 1.5
    var v3 int = 32767
    var v4 uint64 = 42
    v5 := []int{v3}
    v6 := make(map[string]string)
    var v24 bool = (!true)
    v24 = (!(v24 || v24))
    v1 = v1
    switch "\n\t" {
    case "\n\t":
        if false {
            v6[(("\n\t" + "") + ((("hello" + "") + ("\n\t" + ("a" + (("foo bar" + "\n\t") + ("" + "a"))))) + ""))] = ""
        }
    default:
        v2 += v2
    }
    var v25 int64 = 9223372036854775807
    for fi26 := 0; fi26 < 18; fi26++ {
        v4 = genFn1(("\n\t" + ""), 4294967295, (2 > 0))
        switch "\n\t" {
        case "foo bar":
            if v24 {
                v1 += (v1 + (v1 << ((218) & 63)))
                var v27 uint64 = (^(v4))
                for fi28 := 0; fi28 < 20; fi28++ {
                    v29 := GenS2{F0: 2147483647, F1: (v1 & (v1 & v1)), F2: (v1 ^ v1)}
                    switch v3 {
                    case 127:
                        v25 += genFn0(1)
                    }
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
        }
        var v30 int16 = v1
        _ = fi26
        _ = v30
    }
    var v31 int = 32767
    for fi32 := 0; fi32 < 3; fi32++ {
        v3 += (v3 >> ((65535) & 63))
        var v33 string = "\n\t"
        _ = fi32
        _ = v33
    }
    var v34 uint32 = 2147483647
    var v35 int16 = (v1 & (v1 & -331))
    switch "hello" {
    case "\n\t":
        if (0 > 2) {
            switch v31 {
            case -986:
                if ((("hello" + "foo bar") + "\n\t") > ("a" + "a")) {
                    if ((^(v3)) >= v3) {
                        var v36 int16 = v1
                        v37 := GenS3{F0: 255, F1: v4}
                        _ = v36
                        _ = v37
                    } else {
                        if v24 {
                            var v38 float32 = float32((0.1 * 3.14159))
                            var v39 int64 = -9223372036854775808
                            _ = v38
                            _ = v39
                        }
                        var v40 uint8 = 207
                        v24 = v24
                        _ = v40
                    }
                    fmt.Println(v34)
                }
            case 0:
                v31 = (v31 | (v31 * 2147483647))
            default:
                fmt.Println(v31, v4)
            }
        } else {
            if len(v5) > 0 {
                v5[0] = (v3 >> ((255) & 63))
            }
            fmt.Println(v3, v25)
        }
    case "":
        for fi41 := 0; fi41 < 19; fi41++ {
            var v42 int16 = (^(v1))
            if (v25 > 1) {
                v42 = -1
                v25 = v25
                if true {
                    if (-1 >= -19) {
                        for fi43 := 0; fi43 < 18; fi43++ {
                            _ = len(v5)
                            _ = fi43
                        }
                    }
                }
            }
            var v44 float32 = v2
            _ = fi41
            _ = v42
            _ = v44
        }
    case "hello":
        for fi45 := 0; fi45 < 15; fi45++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi45
        }
    }
    if true {
        var v46 float32 = v2
        v2 = (v46 / v46)
        _ = v46
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
