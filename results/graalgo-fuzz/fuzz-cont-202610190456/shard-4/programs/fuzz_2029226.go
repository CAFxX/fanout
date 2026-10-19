package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 uint32; F2 string }
func genFn0(p0 string, p1 uint8, p2 string) uint8 {
    if true {
        switch p0 {
        case "\n\t":
            v4 := GenS1{F0: 42, F1: 42, F2: 42}
            _ = v4
        case "":
            for fi5 := 0; fi5 < 20; fi5++ {
                for fi6 := 0; fi6 < 4; fi6++ {
                    v7 := []int{32768, (fi6 * (-(fi6)))}
                    v8 := GenS1{F0: p1, F1: 0, F2: 0}
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                var v9 int = (fi5 / ((4294967295) | 1))
                _ = fi5
                _ = v9
            }
        default:
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 bool = false
                _ = fi10
                _ = v11
            }
        }
    } else {
        for fi12 := 0; fi12 < 19; fi12++ {
            fmt.Println(p1)
            var v13 int64 = 1
            fmt.Println(v13, p2, p2)
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 bool = (false && (!(-128 == 119)))
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 20; fi16++ {
            p2 = p2
            switch (p2 + "a") {
            case "":
                var v17 bool = true
                _ = v17
            case "foo bar":
                fmt.Println(p0, p1)
            }
            _ = fi16
        }
    }
    p1 = 43
    p0 = "hello"
    v18 := GenS2{F0: 2, F1: -502}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 124
}
func genFn1(p0 uint64) uint16 {
    v19 := []int{}
    var v20 int = 9223372036854775807
    _ = p0
    _ = v19
    _ = v20
    return 391
}

func main() {
    var v1 uint64 = 4294967295
    var v2 string = ""
    var v3 uint32 = 305
    var v21 int8 = -1
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 uint8 = 69
        var v24 int32 = 42
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 string = ""
    v26 := []int{-483, 128, 127, 128}
    if false {
        _ = len(v26)
    } else {
        v27 := GenS3{F0: v25}
        if len(v26) > 0 {
            v26[0] = 42
        }
        if true {
            _ = len(v26)
            fmt.Println(v21, v3, v21)
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
        } else {
            v21 += v21
            var v28 uint16 = 32767
            _ = v28
        }
        _ = v27
    }
    fmt.Println(v3, v25, v25)
    v21 = 2
    var v29 uint32 = (v3 >> ((65535) & 63))
    if false {
        var v30 int = 310
        _ = len(v26)
        _ = v30
    }
    var v31 uint = 4294967295
    var v32 int32 = -1
    v25 = (v2 + "\n\t")
    fmt.Println(v29)
    var v33 int16 = -1
    for fi34 := 0; fi34 < 19; fi34++ {
        var v35 float32 = -2.25
        for fi36 := 0; fi36 < 3; fi36++ {
            switch v25 {
            case "":
                if len(v26) > 0 {
                    v26[0] = (fi36 & fi34)
                }
            case "foo bar":
                if true {
                    fmt.Println(v31)
                    v37 := GenS4{F0: "", F1: (v29 ^ v3), F2: v2}
                    _ = v37
                } else {
                    v25 = v2
                    var v38 float32 = 1e10
                    _ = v38
                }
            }
            _ = fi36
        }
        var v39 uint8 = genFn0("hello", 0, v25)
        _ = fi34
        _ = v35
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
