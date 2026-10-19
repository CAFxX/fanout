package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int8; F2 int32 }
type GenS4 struct { F0 int16; F1 uint16; F2 bool }
func genFn0() uint8 {
    v4 := []int{}
    var v5 uint8 = 42
    _ = v4
    _ = v5
    return 161
}
func genFn1(p0 uint64, p1 uint, p2 int8) bool {
    v6 := GenS1{F0: 1}
    v7 := []int{0, -599, 32768, 32767}
    for fi8 := 0; fi8 < 16; fi8++ {
        if (!(((^(p2)) != p2) && false)) {
            var v9 uint = 255
            _ = v9
        } else {
            for fi10 := 0; fi10 < 15; fi10++ {
                p1 += (^(p1))
                _ = len(v7)
                _ = fi10
            }
            var v11 int = int(p1)
            v11 += fi8
            _ = v11
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        switch "a" {
        case "foo bar":
            var v12 int32 = -1
            _ = v12
        case "":
            var v13 string = "hello"
            _ = v13
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return true
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = ("a" + "hello")
    v3 := make(map[int]string)
    var v14 uint8 = 128
    v14 -= 0
    v15 := []int{42, 445}
    if false {
        v14 += (v14 & v14)
        v16 := GenS1{F0: 65535}
        _ = v16
    } else {
        fmt.Println(v14, v2)
    }
    v17 := GenS1{F0: uint32(v14)}
    v2 = (v2 + (v2 + v2))
    for fi18 := 0; fi18 < 15; fi18++ {
        switch v2 {
        case "":
            for fi19 := 0; fi19 < 11; fi19++ {
                if false {
                    switch ("a" + ("" + ("a" + ((v2 + v2) + ("hello" + v2))))) {
                    case "a":
                        switch (^(fi18)) {
                        case -1:
                            v20 := GenS1{F0: 4294967295}
                            _ = v20
                        case -210:
                            var v21 int64 = 9223372036854775807
                            var v22 int16 = -32768
                            var v23 int8 = 63
                            _ = v21
                            _ = v22
                            _ = v23
                        case 32767:
                            fmt.Println(v2, fi18)
                        }
                    case "\n\t":
                        if ((591 <= 256) || (true || false)) {
                            if true {
                                v2 = v2
                            }
                            var v24 int8 = 8
                            if false {
                                fmt.Println(v24, fi19, v2)
                            }
                            _ = v24
                        }
                    }
                    v25 := GenS2{F0: 1}
                    _ = v25
                }
                v2 = (v2 + v2)
                v26 := GenS2{F0: 42}
                v27 := v1[596]
                switch v27 {
                case -844:
                    v28 := GenS1{F0: 624}
                    _ = v28
                case 4294967295:
                    v29 := GenS1{F0: 329}
                    _ = v29
                case 128:
                    var v30 uint = 370
                    var v31 string = ((v2 + (v2 + "hello")) + (v2 + v2))
                    _ = len(v1)
                    _ = v30
                    _ = v31
                default:
                    if true {
                        v32 := v1[731]
                        v33 := GenS3{F0: 0, F1: 77, F2: 42}
                        var v34 uint32 = 977
                        _ = v32
                        _ = v33
                        _ = v34
                    } else {
                        v35 := GenS3{F0: 42, F1: -1, F2: -775}
                        _ = len(v1)
                        _ = v35
                    }
                }
                _ = fi19
                _ = v26
                _ = v27
            }
        case "a":
            for fi36 := 0; fi36 < 7; fi36++ {
                v37 := GenS1{F0: 42}
                v2 = ("hello" + "\n\t")
                v38 := GenS4{F0: 0, F1: 0, F2: true}
                _ = fi36
                _ = v37
                _ = v38
            }
        case "hello":
            fmt.Println(v2, v14)
        default:
            v2 = v2
        }
        v2 = ((("hello" + v2) + "\n\t") + v2)
        var v39 uint8 = genFn0()
        _ = fi18
        _ = v39
    }
    v14 -= (^(v14))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
}
