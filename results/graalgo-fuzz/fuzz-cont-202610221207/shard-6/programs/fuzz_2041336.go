package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8 }
func genFn0(p0 int64) uint32 {
    if false {
        var v4 uint = 0
        _ = v4
    } else {
        fmt.Println(p0)
        if (!(2147483647 == 42)) {
            v5 := GenS1{F0: 1, F1: 637, F2: -9223372036854775808}
            for fi6 := 0; fi6 < 5; fi6++ {
                p0 += (p0 - 42)
                var v7 float64 = 0.1
                if true {
                    v8 := []int{(^(fi6)), fi6}
                    for fi9 := 0; fi9 < 18; fi9++ {
                        v7 += 0.1
                        _ = fi9
                    }
                    var v10 int64 = p0
                    _ = v10
                    _ = v8
                } else {
                    p0 -= p0
                    var v11 bool = (1 >= 299)
                    _ = v11
                }
                _ = fi6
                _ = v7
            }
            _ = v5
        }
    }
    v12 := GenS1{F0: 0, F1: 145, F2: -9223372036854775808}
    fmt.Println(p0)
    var v13 uint = 65535
    _ = p0
    _ = v12
    _ = v13
    return 5
}

func main() {
    v1 := []uint32{65535, 2147483648, 2147483647}
    v2 := []string{("\n\t" + "hello"), ("foo bar" + ("\n\t" + (("a" + "\n\t") + (("a" + ("" + "\n\t")) + "\n\t")))), "hello", (("a" + (("foo bar" + "") + "")) + ("\n\t" + "foo bar"))}
    var v3 bool = false
    switch 42 {
    case -9223372036854775808:
        v14 := GenS2{F0: -566}
        _ = v14
    }
    var v15 uint64 = 42
    if len(v2) > 0 {
        v2[0] = "a"
    }
    switch -9223372036854775808 {
    case 128:
        fmt.Println(v3)
    default:
        v3 = v3
    }
    var v16 uint16 = 0
    var v17 int16 = -285
    if true {
        if (140 > 1) {
            var v18 int16 = (v17 / (((v17 - v17)) | 1))
            for fi19 := 0; fi19 < 20; fi19++ {
                fmt.Println(v16, v17)
                if (!true) {
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                    for fi20 := 0; fi20 < 20; fi20++ {
                        _ = len(v1)
                        _ = fi20
                    }
                }
                _ = fi19
            }
            v21 := GenS3{F0: 288}
            _ = v18
            _ = v21
        } else {
            v22 := GenS2{F0: 0}
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := GenS4{F0: true}
                v25 := GenS5{F0: 118}
                _ = fi23
                _ = v24
                _ = v25
            }
            var v26 uint16 = (v16 / ((v16) | 1))
            _ = v22
            _ = v26
        }
        v15 += 2
        _ = len(v2)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!(v3 && (!((^(v16)) > 1)))) {
        var v27 uint64 = v15
        switch ("" + "foo bar") {
        case "":
            var v28 uint32 = 613
            _ = v28
        }
        fmt.Println(v15)
        _ = v27
    } else {
        switch "" {
        case "foo bar":
            v1 = append(v1, 2147483647)
        case "a":
            var v29 uint16 = (v16 % (((v16 % (((v16 % ((255) | 1))) | 1))) | 1))
            _ = v29
        }
    }
    var v30 int8 = -58
    v17 -= v17
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 string = ("hello" + "\n\t")
        v30 += (^(v30))
        _ = fi31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
}
