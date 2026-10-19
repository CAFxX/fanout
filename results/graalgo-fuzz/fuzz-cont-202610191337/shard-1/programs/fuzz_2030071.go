package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 string }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int8; F1 int64 }
func genFn0() int32 {
    var v4 string = ""
    v5 := []int{-318, 522}
    _ = v4
    _ = v5
    return 2147483647
}
func genFn1(p0 bool, p1 int32) bool {
    p0 = true
    fmt.Println(p0)
    p0 = p0
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 uint64 = 4294967295
    v2 := []string{("hello" + "a"), ("hello" + (("hello" + "\n\t") + ("a" + "a"))), "foo bar"}
    var v3 uint = 65535
    v1 += v1
    var v6 uint64 = (v1 / ((uint64(v3)) | 1))
    if true {
        switch (("foo bar" + "") + "\n\t") {
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "foo bar":
            if false {
                fmt.Println(v1, v3, v1)
                var v7 float64 = 0.1
                var v8 uint32 = 400
                _ = v7
                _ = v8
            } else {
                switch 32767 {
                case -1:
                    var v9 int32 = genFn0()
                    _ = v9
                case 1:
                    v3 += 42
                case 42:
                    if len(v2) > 0 {
                        v2[0] = ("" + "a")
                    }
                }
            }
        default:
            for fi10 := 0; fi10 < 6; fi10++ {
                if (234 < 0) {
                    _ = len(v2)
                    for fi11 := 0; fi11 < 14; fi11++ {
                        fmt.Println(v6, v6, v6)
                        _ = fi11
                    }
                } else {
                    v3 += (v3 + v3)
                }
                _ = fi10
            }
        }
    } else {
        switch -17 {
        case 303:
            var v12 uint8 = 255
            _ = v12
        case 0:
            var v13 uint16 = uint16(v3)
            _ = v13
        case -384:
            var v14 uint8 = 2
            _ = v14
        default:
            _ = len(v2)
        }
        v15 := GenS1{F0: 128, F1: 42, F2: (("a" + "hello") + ("foo bar" + (("foo bar" + ("\n\t" + ("foo bar" + "foo bar"))) + "")))}
        _ = v15
    }
    v3 += (v3 << (((v3 + (^(v3)))) & 63))
    fmt.Println(v1, v6, v3)
    _ = len(v2)
    switch 0 {
    case 9223372036854775807:
        if true {
            if false {
                if ((true && false) || false) {
                    fmt.Println(v3, v1, v1)
                    fmt.Println(v3, v1)
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v3 = v3
                        var v17 int16 = 32767
                        fmt.Println(fi16)
                        _ = fi16
                        _ = v17
                    }
                } else {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        fmt.Println(fi18, v1)
                        fmt.Println(v3, fi18)
                        _ = fi18
                    }
                    var v19 int32 = -1
                    _ = v19
                }
                for fi20 := 0; fi20 < 2; fi20++ {
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    for fi21 := 0; fi21 < 6; fi21++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v6 = (v1 + v1)
                        _ = fi21
                    }
                    var v22 int64 = 1
                    _ = fi20
                    _ = v22
                }
            }
            if false {
                switch -9223372036854775808 {
                case 4294967295:
                    v2 = append(v2, "\n\t")
                case 32767:
                    var v23 string = ""
                    _ = v23
                case 9223372036854775807:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
        }
    case 2147483647:
        fmt.Println(v1, v6)
    case 128:
        v24 := GenS2{F0: (^(v1)), F1: true}
        _ = v24
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        v1 += v1
        v26 := GenS3{F0: 1, F1: 42}
        var v27 uint8 = 2
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
}
